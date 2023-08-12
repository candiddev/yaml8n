# YAML8n

> Easy Localizations, Great Developer Experience

[![Integration](https://github.com/candiddev/yaml8n/actions/workflows/workflow.yaml/badge.svg?branch=main)](https://github.com/candiddev/yaml8n/actions/workflows/workflow.yaml)

YAML8n is an open source command line (CLI) tool for converting translations into type safe code.

YAML8n makes translating applications easy:

- Define your translations using YAML
- Use Google Cloud Translate to fill in missing languages
- Generate type safe code for your favorite programming languages

Visit https://yaml8n.dev for more information.

## License

The code in this repository is licensed under the [GNU AGPL](https://www.gnu.org/licenses/agpl-3.0.en.html).  Visit https://yaml8n.dev/pricing/ to purchase a license exemptions.

## Development

Our development process is mostly trunk-based with a `main` branch that folks can contribute to using pull requests.  We tag releases as necessary using CalVer.

### Repository Layout

- `./github:` Reusable GitHub Actions
- `./go:` YAML8n code
- `./hugo:` YAML8n website
- `./shell:` Development tooling
- `./shared:` Shared libraries from https://github.com/candiddev/shared

Make sure you initialize the shared submodule:

```bash
git submodule update --init
```

### CI/CD

We use GitHub Actions to lint, test, build, release, and deploy the code.  You can view the pipelines in the `.github/workflows` directory.  You should be able to run most workflows locally and validate your code before opening a pull request.

### Tooling

Visit [shared/README.md](shared/README.md) for more information.
