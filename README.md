# Go Enterprise Starter (Git + CI/CD + Conv. Commits)

Base para projeto Go com Git “padrão empresa”: Conventional Commits, hooks versionados, lint/test automatizados e releases semânticas.

## Rodando local
```bash
make tidy fmt vet lint test
make hooks       # configura template e hooks
git config --get core.hooksPath  # deve ser scripts/git-hooks
```

## Estrutura
- `cmd/app/main.go`: entrypoint
- `internal/*`: pacotes internos
- `scripts/git-hooks/*`: hooks versionados
- `.golangci.yml`: configuração de lint
- `.pre-commit-config.yaml`: qualidade & segurança
- `.gitmessage`: template para commits
- `.releaserc`: release semântico (CI)

## CI
- Lint com `golangci-lint`
- Teste com `-race` e cobertura
- Release automático na `main` (semantic-release)
