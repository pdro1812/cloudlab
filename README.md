# Git Enterprise Starter

Estrutura-base para fluxo Git de nível produção (trunk-based) com automatizações de CI/CD, versionamento semântico e governança.

## Principais componentes
- **Branching**: `main` (produção), `develop` (opcional), `feat/*`, `fix/*`, `chore/*`, `hotfix/*`.
- **Proteções**: PR obrigatório, 1–2 reviews, status checks verdes, squash merge, regras de nome de branch.
- **Commits**: Conventional Commits + `commitlint`.
- **Versionamento**: SemVer + `semantic-release` (tag, changelog, release notes).
- **Qualidade**: Linters + testes automatizados + análise estática.
- **Segurança**: CODEOWNERS, Dependabot (opcional), políticas de secrets.
- **Templates**: Issues, PR, contribuição, segurança.

Veja `docs/branching.md` e `CONTRIBUTING.md`.
