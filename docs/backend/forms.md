# FEATURE DYNAMIC FORMS - DF

## Überblick

Das System für dynamische Formulare ermöglicht es, flexible Formulare zu erstellen, die verschiedene Fragetypen unterstützen und deren Struktur zur Laufzeit definiert werden kann. Die Implementierung basiert auf JSON-Strukturen für maximale Flexibilität.

## Architektur

### Datenmodelle

#### Form Model (`models/form.go`)

```go
type Form struct {
    ID                uuid.UUID       `json:"id" gorm:"primaryKey"`
    Title             string          `json:"title"`
    Description       string          `json:"description"`
    MaxSubmitsPerUser int             `json:"max_submits_per_user"`
    Body              json.RawMessage `json:"body"`
    CreatedAt         time.Time       `json:"created_at" gorm:"autoCreateTime"`
    UpdatedAt         time.Time       `json:"updated_at" gorm:"autoUpdateTime"`
}
```

**Felder Erklärung:**
- `ID`: Eindeutige UUID für das Formular
- `Title`: Anzeigename des Formulars
- `Description`: Beschreibung des Formulars
- `MaxSubmitsPerUser`: Maximale Anzahl von Einreichungen pro Benutzer
- `Body`: JSON-Struktur mit den Formularfragen (siehe JSON Schema unten)
- `CreatedAt/UpdatedAt`: Zeitstempel für Erstellung und letzte Änderung

#### Submission Model (`models/submission.go`)

```go
type Submission struct {
    ID        uuid.UUID       `json:"id" gorm:"primaryKey"`
    FormID    uuid.UUID       `json:"form_id"`
    Answer    json.RawMessage `json:"answer"`
    CreatedAt time.Time       `json:"created_at"`
}
```

**Felder Erklärung:**
- `ID`: Eindeutige UUID für die Einreichung
- `FormID`: Referenz auf das zugehörige Formular
- `Answer`: JSON-Struktur mit den Antworten (Key-Value Paare)
- `CreatedAt`: Zeitstempel der Einreichung

## JSON Schema für Formulare

### Form Body Struktur

```json
{
  "questions": [
    {
      "id": "q1",
      "type": "text",
      "label": "Was ist dein Name?",
      "required": true,
      "placeholder": "Gib deinen Namen ein"
    },
    {
      "id": "q2",
      "type": "email",
      "label": "Was ist deine E-Mail-Adresse?",
      "required": true,
      "validation": {
        "pattern": "^[\\w-\\.]+@([\\w-]+\\.)+[\\w-]{2,4}$"
      }
    },
    {
      "id": "q3",
      "type": "select",
      "label": "Wähle eine Option",
      "required": true,
      "options": [
        {"value": "option1", "label": "Option 1"},
        {"value": "option2", "label": "Option 2"}
      ]
    },
    {
      "id": "q4",
      "type": "checkbox",
      "label": "Mehrfachauswahl",
      "options": [
        {"value": "choice1", "label": "Auswahl 1"},
        {"value": "choice2", "label": "Auswahl 2"}
      ]
    },
    {
      "id": "q5",
      "type": "textarea",
      "label": "Kommentar",
      "required": false,
      "rows": 4
    }
  ]
}
```

### Unterstützte Fragetypen

| Typ | Beschreibung | Zusätzliche Eigenschaften |
|-----|--------------|---------------------------|
| `text` | Einzeiliges Textfeld | `placeholder`, `maxLength`, `minLength` |
| `email` | E-Mail Eingabefeld | `validation.pattern` |
| `number` | Zahlen Eingabefeld | `min`, `max`, `step` |
| `textarea` | Mehrzeiliges Textfeld | `rows`, `cols`, `maxLength` |
| `select` | Dropdown-Auswahl | `options[]` |
| `radio` | Radio-Button Gruppe | `options[]` |
| `checkbox` | Checkbox-Gruppe | `options[]` |
| `date` | Datums-Auswahl | `min`, `max` |
| `file` | Datei-Upload | `accept`, `maxSize` |

## Submission Antworten Format

Die Antworten werden als JSON-Objekt gespeichert, wobei die Frage-IDs als Keys verwendet werden:

```json
{
  "q1": "Max Mustermann",
  "q2": "max@example.com",
  "q3": "option1",
  "q4": ["choice1", "choice2"],
  "q5": "Das ist mein Kommentar"
}
```

## Best Practices

### 1. Eindeutige Frage-IDs
- Verwende beschreibende und eindeutige IDs für Fragen (z.B. `student_name`, `email_address`)
- Vermeide Sonderzeichen in IDs, nutze nur `a-z`, `0-9` und `_`

### 2. Validierung
- Definiere `required` Felder explizit
- Nutze `validation.pattern` für komplexe Validierungen
- Setze `maxLength` für Textfelder um Spam zu verhindern

### 3. Benutzerfreundlichkeit
- Verwende aussagekräftige `label` Texte
- Nutze `placeholder` für Eingabehilfen
- Gruppiere verwandte Fragen logisch

### 4. Performance
- Begrenze die Anzahl der Fragen pro Formular (< 50)
- Nutze `MaxSubmitsPerUser` um Spam zu verhindern
- Implementiere Indizierung auf `FormID` in der Submissions-Tabelle

### 5. Datenintegrität
- Verwende Foreign Key Constraints zwischen Form und Submission
- Implementiere Soft-Delete für Formulare mit existierenden Submissions
- Validiere JSON-Struktur beim Erstellen/Updaten von Formularen

## Beispiel Implementation (aus seed.go)

```go
// Beispiel eines gut strukturierten Formulars
&models.Form{
    ID: uuid.New(),
    Title: "Schüler Stammdaten",
    Description: "Erfassung der grundlegenden Schülerdaten",
    Body: json.RawMessage(`{
        "questions": [
            {
                "id": "student_name",
                "type": "text",
                "label": "Vollständiger Name",
                "required": true,
                "maxLength": 100
            },
            {
                "id": "email",
                "type": "email",
                "label": "E-Mail-Adresse",
                "required": true,
                "validation": {
                    "pattern": "^[\\w-\\.]+@([\\w-]+\\.)+[\\w-]{2,4}$"
                }
            },
            {
                "id": "class",
                "type": "select",
                "label": "Klasse",
                "required": true,
                "options": [
                    {"value": "10a", "label": "10a"},
                    {"value": "10b", "label": "10b"},
                    {"value": "11a", "label": "11a"}
                ]
            }
        ]
    }`),
    MaxSubmitsPerUser: 1,
    CreatedAt: time.Now(),
    UpdatedAt: time.Now()
}
```

## Zukünftige Erweiterungen

### Geplante Features
- [ ] UserID Integration in Submissions
- [ ] Berechtigungssystem für Formular-Zugriff
- [ ] Conditional Logic (bedingte Fragen)
- [ ] File Upload Support
- [ ] Form Templates
- [ ] Export/Import Funktionalität
- [ ] Analytics und Reporting

### Technische Schulden
- UserID Feld in Submission Model hinzufügen
- CreatorID Feld in Form Model implementieren
- Soft-Delete Implementierung
- JSON Schema Validierung auf Backend-Seite
- Audit Trail für Formular-Änderungen

## Migration Strategie

Bei Änderungen an der JSON-Struktur:
1. Backward Compatibility gewährleisten
2. Versionierung der Schema-Struktur
3. Migration Scripts für bestehende Daten
4. Graceful Handling von Legacy-Formularen

## Testing

### Unit Tests
- JSON Schema Validierung
- CRUD Operationen für Forms/Submissions
- Edge Cases für verschiedene Fragetypen

### Integration Tests
- End-to-End Formular-Workflow
- Performance Tests mit großen Formularen
- Concurrent Submissions Testing