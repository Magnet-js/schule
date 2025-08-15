.PHONY: db backend frontend stop clean

db:
	@echo "🚀 Starte Postgres..."
	@docker-compose up -d postgres
	@echo "✅ Postgres läuft auf localhost:5432"

backend: db
	@echo "🌐 Starte Backend..."
	@if [ -f .env ]; then \
		echo "📄 .env Datei gefunden, lade Umgebungsvariablen..."; \
		export $$(grep -v '^#' .env | xargs) && cd backend && go run main.go; \
	fi

frontend: db
	@echo "🌐 Starte Frontend..."
	@cd frontend && pnpm install && pnpm dev

stop:
	@echo "⏹️  Stoppe alles..."
	@docker-compose stop
	@pkill -f "go run main.go" 2>/dev/null || true
	@pkill -f "vite" 2>/dev/null || true
	@echo "✅ Alles gestoppt"

clean:
	@echo "🧹 Setze Datenbank zurück..."
	@docker-compose down -v
	@echo "✅ Datenbank zurückgesetzt"
