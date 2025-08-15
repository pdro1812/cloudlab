# Estratégia de Branching (Trunk-Based)

- `main`: sempre estável; cada merge dispara pipeline completo e (opcionalmente) release.
- Branches curtas: `feat/*`, `fix/*`, `chore/*`, `hotfix/*`.
- `hotfix/*` pode ir direto para `main` com priorização de revisão.
- **Proteções**: 1–2 revisores obrigatórios, status checks verdes, commits assinados (opcional).

## Mensagens de commit (Conventional Commits)
Ex.: `feat(auth): adiciona login com OAuth2`

## Versionamento
- SemVer: `MAJOR.MINOR.PATCH`
- `feat` -> `minor`, `fix` -> `patch`, `BREAKING CHANGE` -> `major`.
