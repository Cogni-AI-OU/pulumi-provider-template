# .github Directory

Use this as the entry point for agent work, and follow linked catalogs when relevant.

## Directory-Specific Agent files

TBA

## Additional key files

- [.github/actionlint-matcher.json](actionlint-matcher.json): problem matchers used in workflows.
- [.github/pre-commit-matcher.json](pre-commit-matcher.json): problem matchers used in workflows.
- [.github/FIREWALL.md](FIREWALL.md): firewall configuration and recommended hosts for agents.

## Hardened NEVER List

- **NEVER create `.github/README.md`**: GitHub renders `.github/README.md` with the highest priority. Creating it will
  override the main `README.md` on the repository homepage and profile page.

## Troubleshooting

TBA

## Additional notes

- Keep this Agent file up-to-date.
