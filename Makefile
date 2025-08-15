.PHONY: tidy fmt vet lint test cover precommit-install hooks

GO ?= go

tidy:
	@$(GO) mod tidy

fmt:
	@echo "gofmt -s -w ."
	@gofmt -s -w .
	@command -v goimports >/dev/null 2>&1 && goimports -w . || true

vet:
	@$(GO) vet ./...

lint:
	@command -v golangci-lint >/dev/null 2>&1 && golangci-lint run || (echo "golangci-lint não instalado, skip"; exit 0)

test:
	@$(GO) test ./... -race -count=1 -short -coverprofile=coverage.out

cover:
	@$(GO) tool cover -func=coverage.out | tail -n1

precommit-install:
	@pre-commit install --hook-type pre-commit --hook-type commit-msg || true

hooks:
	@git config commit.template .gitmessage
	@git config core.hooksPath scripts/git-hooks
	@echo "Hooks habilitados via core.hooksPath -> scripts/git-hooks"
