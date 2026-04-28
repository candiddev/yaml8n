---
categories:
- reference
description: Reference documentation for YAML8n's configuration
title: Config
---

{{% snippet config_format YAML8n yaml8n %}}

## Configuration Values

{{% snippet "config_cli" green %}}

### `checkCode` {#checkcode}

String, check/validate a specific [language code]({{< ref "/docs/references/translations#iso639codes" >}}).

**Default:** `""`

### `defaultCode` {#defaultcode}

String, the default [language code]({{< ref "/docs/references/translations#iso639codes" >}}) to use if none are set in the translation file.  Used to globally set the `defaultCode` value across multiple translation files or within a monorepo.

**Default:** `""`

### `failWarn` {#failwarn}

Boolean, if warnings should cause failures.

**Default:** `false`

{{% snippet "config_httpClient" YAML8n %}}

### `iso639Codes` {#iso639codes}

Map of ISO 639 codes to their pretty name.  YAML8n will warn on translations that are missing these codes.  Will be used if none are set in the translation file.  Used to globally set the `iso639Codes` value across multiple translation files or within a monorepo.

**Default:** `{}`

{{% snippet config_jsonnet true %}}

{{% snippet config_licenseKey YAML8n %}}

### `translations` {#translations}

String, the path to a translations YAML file.

**Default:** `""`
