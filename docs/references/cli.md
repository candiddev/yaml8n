---
categories:
- reference
description: Reference documentation for YAML8n's CLI
title: CLI
---

{{% snippet cli_arguments %}}

{{% snippet cli_commands YAML8n %}}

{{% cli_autocomplete %}}

{{% snippet cli_config %}}

{{% snippet cli_docs %}}

{{% snippet cli_eula YAML8n %}}

### `generate [path]`

Generate code for the [outputs]({{< ref "/docs/references/translations#outputs" >}}) specified within the [Translations]({{< ref "/docs/references/translations" >}}) located at `path`.

{{% snippet cli_jq %}}

### `translate [path]`

Add missing translations for the [Translations]({{< ref "/docs/references/translations" >}}) located at `path` using [Google Cloud Translation API](https://cloud.google.com/translate).  See [Guides/Add Missing Translations]({{< ref "/docs/guides/add-missing-translations" >}}) for more information.

### `validate [path]`

Lint and validate the [Translations]({{< ref "/docs/references/translations" >}}) located at `path`.

{{% cli_version %}}

### `watch [path]`

Watch the [Translations]({{< ref "/docs/references/translations" >}}) located at `path` for changes and generate code for the [outputs]({{< ref "/docs/references/translations#outputs" >}}) specified within on change.
