---
categories:
- reference
description: Reference documentation for YAML8n's CLI
title: CLI
---

## Arguments

Arguments must be entered before commands.

### `-c [code]`

Check/validate a specific [language code](../references/translations#iso639codes).

### `-d`

Enable debug logging.

### `-j`

Output JSON instead of YAML.

### `-n`

Disable colored log output.

### `-w`

Fail on translation warnings, like missing translations.

## Commands

### `generate [path]`

Generate code for the [outputs](../translations/#outputs) specified within the [Translations](../translations/) located at `path`.

### `translate [path]`

Add missing translations for the [Translations](../translations/) located at `path` using [Google Cloud Translation API](https://cloud.google.com/translate).  See [Guides/Add Missing Translations](../../guides/add-missing-translations) for more information.

### `validate [path]`

Lint and validate the [Translations](../translations/) located at `path`.

### `version`

Print the current version of YAML8n.

### `watch [path]`

Watch the [Translations](../translations/) located at `path` for changes and generate code for the [outputs](../translations/#outputs) specified within on change.
