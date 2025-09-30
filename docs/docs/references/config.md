---
categories:
- reference
description: Reference documentation for YAML8n's configuration
title: Config
---

{{% snippet config_format YAML8n yaml8n %}}

## Configuration Values

{{% snippet "config_cli" %}}

{{% snippet config_key "checkCode" %}}

String, check/validate a specific [language code]({{< ref "/docs/references/translations#iso639codes" >}}).

**Default:** `""`

{{% snippet config_key "defaultCode" %}}

String, the default [language code]({{< ref "/docs/references/translations#iso639codes" >}}) to use if none are set in the translation file.  Used to globally set the `defaultCode` value across multiple translation files or within a monorepo.

**Default:** `""`

{{% snippet config_key "failWarn" %}}

Boolean, if warnings should cause failures.

**Default:** `false`

{{% snippet "config_httpClient" YAML8n %}}

{{% snippet config_key "iso639Codes" %}}

Map of ISO 639 codes to their pretty name.  YAML8n will warn on translations that are missing these codes.  Will be used if none are set in the translation file.  Used to globally set the `iso639Codes` value across multiple translation files or within a monorepo.

**Default:** `{}`

{{% snippet config_jsonnet true %}}

{{% snippet config_licenseKey YAML8n %}}

{{% snippet config_key "translations" %}}

String, the path to a translations YAML file.

**Default:** `""`
