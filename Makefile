.PHONY: lint test

lint:
	@echo "Rodando linters..."
	@if command -v npm >/dev/null && [ -f package.json ]; then npm run lint; else echo "npm lint skip"; fi
	@if [ -f .pre-commit-config.yaml ]; then pre-commit run --all-files; else echo "pre-commit skip"; fi

test:
	@echo "Rodando testes..."
	@if command -v npm >/dev/null && [ -f package.json ]; then npm test; else echo "npm test skip"; fi
	@if [ -f pytest.ini ] || [ -f pyproject.toml ]; then pytest -q; else echo "pytest skip"; fi
