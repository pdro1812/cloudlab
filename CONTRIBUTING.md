# Contribuindo

## Fluxo de branches (trunk-based)
- Crie branch a partir de `main`: `feat/nome-curto`, `fix/bug-123`, `chore/dep-update`.
- Commits no padrão **Conventional Commits** (ex.: `feat(api): adiciona endpoint /health`).
- Abra Pull Request para `main`. Requer revisão e checks verdes.
- Preferir **squash merge** para manter histórico limpo.

## Padrão de commit
- `feat`: nova funcionalidade
- `fix`: correção de bug
- `docs`, `chore`, `refactor`, `perf`, `test`, `build`, `ci`
- Escopo opcional: `feat(auth): ...`
- Mensagem no imperativo, breve e descritiva.

## Testes e qualidade
- Rode `make test` e `make lint` localmente antes do PR.
- Não commitar segredos (.env). Use variáveis de ambiente no provedor (Actions/Secrets).

## Releases
- `semantic-release` cria tag, versiona e gera CHANGELOG automaticamente ao fazer merge na `main`.
