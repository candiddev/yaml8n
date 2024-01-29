---
categories:
- reference
description: Reference documentation for YAML8n's CLI
title: CLI
---

## Arguments

Arguments must be entered before commands.

### `-c [code]`

Check/validate a specific [language code]({{< ref "/docs/references/translations#iso639codes" >}}).

### `-d`

Enable debug logging.

### `-n`

Disable colored log output.

### `-w`

Fail on translation warnings, like missing translations.

## Commands

### `generate [path]`

Generate code for the [outputs]({{< ref "/docs/references/translations#outputs" >}}) specified within the [Translations]({{< ref "/docs/references/translations" >}}) located at `path`.

### `translate [path]`

Add missing translations for the [Translations]({{< ref "/docs/references/translations" >}}) located at `path` using [Google Cloud Translation API](https://cloud.google.com/translate).  See [Guides/Add Missing Translations]({{< ref "/docs/guides/add-missing-translations" >}}) for more information.

### `validate [path]`

Lint and validate the [Translations]({{< ref "/docs/references/translations" >}}) located at `path`.

### `version`

Print the current version of YAML8n.

### `watch [path]`

Watch the [Translations]({{< ref "/docs/references/translations" >}}) located at `path` for changes and generate code for the [outputs]({{< ref "/docs/references/translations#outputs" >}}) specified within on change.
