new-project:
	./scripts/create-project.sh

select-and-build:
	./scripts/select-and-build.sh

build-project:
	@if [ -z "$(SLUG)" ]; then \
		echo "Error: SLUG is required. Usage: make build SLUG=your-project-slug"; \
		exit 1; \
	fi
	@if [ ! -d "projects/$(SLUG)" ]; then \
		echo "Error: Project '$(SLUG)' not found in projects directory"; \
		exit 1; \
	fi
	@echo "Building project: $(SLUG)"
	@cd projects/$(SLUG) && go build -o bin/$(SLUG) cmd/main.go
	@echo "✅ Project $(SLUG) built successfully in projects/$(SLUG)/bin/$(SLUG)"

