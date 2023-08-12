---
categories:
- reference
title: Translations
---

YAML8n primarily reads in a JSON or YAML file containing translations.  It can be named anything, but in the documentation it is called `translations.yaml`.

The file has content similar to this:

{{< highlight yaml >}}
defaultCode: en
iso639Codes:
  de: Deutsch
  en: English
outputs:
- format: typescript
  package: index
  path: ../web/src/lib/yaml8n
translations:
  ActionAdd:
    context: Add something
    de: Hinzufügen
    en: Add
{{< /highlight >}}

## Format

### `defaultCode`

String, specifies the default ISO 639 code if a translation isn't found.  Must be specified.

### `iso639Codes`

Map of ISO 639 codes to their pretty name.  YAML8n will warn on translations that are missing these codes.  At least one must be specified.

### `outputs`

A list of output formats for generating translations.  The format for the output is specified below.  At least one must be specified.

#### `format`

String, the format of the generated code.  Supported values are `go`, `go-i18n`, and `typescript`.  Must be specified.

#### `package`

String, the name of the package for the generated code.  Means different things for each language--in Go, this is the package name, in TypeScript, this is the name of the file it creates.  If not specified, the default is `yaml8n`.

#### `path`

String, the directory to output the translations to.  Can be an absolute path or relative to the configuration.  Must be specified.  YAML8n will create the directory if it does not exist.

### `translations`

Map of names to translations.  The name of the translation can be anything, the generated code will use this name for variable names so it must be valid within your language.

The translations are a map of [iso639Codes](#iso639codes) or the word `context` to a string.  For iso639Codes, the string must be the corresponding translation for the language.  For `context`, the string should be contextual information on where the translation will appear to help with translating.  YAML8n will warn if missing `context`.
