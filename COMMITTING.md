# Guia prático: como commitar em um projeto Go (padrão empresa)

## TL;DR
1. `git switch -c feat/minha-feature`
2. `make tidy fmt vet lint test`
3. `git add -p` (adicione só o que faz sentido)
4. `git commit` (use o template e o padrão abaixo)
5. `git push -u origin feat/minha-feature` e abra PR

---

## Padrão de mensagem (Conventional Commits)
```
<type>(<scope>): <subject>

<body opcional>
BREAKING CHANGE: <detalhe>
Refs: #123
```
**type**: feat | fix | docs | style | refactor | perf | test | build | ci | chore | revert  
**scope**: módulo/pacote, ex.: `api`, `auth`, `db`, `health`  
**subject**: imperativo e curto (até ~72 chars)

### Exemplos
- `feat(api): adiciona endpoint /health`
- `fix(repo): corrige condição de corrida no worker`
- `refactor(auth): simplifica validação de token`
- `perf(cache): reduz alocações no hot path`
- `docs(readme): adiciona seção de setup`
- `revert: "feat(api): adiciona endpoint /health"`

## Tamanho e foco do commit
- Um commit por mudança coesa.
- Evite commits gigantescos.
- Inclua testes quando corrigir bug ou adicionar feature.

## Hooks locais (recomendado)
Instale os hooks versionados e template:
```
make hooks
pre-commit install --hook-type pre-commit --hook-type commit-msg
```
O hook `pre-commit` roda:
- `gofmt -s` (falha se houver format por fazer)
- `go vet`
- `golangci-lint` (se instalado)
- `go test -short`

O hook `commit-msg` valida o padrão Conventional Commits.
- Se você tiver Node, usa `commitlint`.
- Sem Node, aplica regex simples (fallback).

## Checklist antes do commit
- [ ] Código formatado (`make fmt`)
- [ ] Lint e vet passaram (`make lint vet`)
- [ ] Testes locais passaram (`make test`)
- [ ] Não inclui segredos nem .env
- [ ] Mensagem de commit no padrão

## Dicas para Go
- Rodar `go mod tidy` sempre que adicionar/atualizar deps.
- Use `-race` nos testes da CI para detectar data races.
- Prefira funções pequenas e pacotes coesos em `internal/`.

