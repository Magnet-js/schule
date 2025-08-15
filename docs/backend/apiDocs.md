# 📝 Example API Documentation: ToDo Backend

This is a sample documentation for a simple REST API backend that manages ToDo items.

---

## 🌐 Base URL

example.com/api/v1

---

## 📚 Endpoints

### ✅ GET /health

**Description:** Returns a list of all ToDos.

**Response:**

- **Status Code:** `200 OK`
- **Content-Type:** `application/json`

**Example Response:**

```json
{
    "status": "ok"
}
```

### PUT /forms

**Description:** Create a new form based on the body

**Body:**
```json
{
    "title": "Schüler Stammdaten",
    "description": "Erfassung der grundlegenden Schülerdaten",
    "multi_viewable": true,
    "approve_needed": false,
    "body": {
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
                    {
                        "value": "10a",
                        "label": "10a"
                    },
                    {
                        "value": "10b",
                        "label": "10b"
                    },
                    {
                        "value": "11a",
                        "label": "11a"
                    }
                ]
            }
        ]
    },
    "form_editors": {
        "users": [
            "user2"
        ]
    },
    "form_viewer": {
        "groups": [
            "everyone"
        ]
    },
    "form_approvers": {
        "groups": [
            "teacher"
        ],
        "users": [
            "user1",
            "user2"
        ]
    }
}
```

**Response:**

- **Status Code:** `201 Created`
