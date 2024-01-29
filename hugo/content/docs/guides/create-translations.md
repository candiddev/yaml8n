---
categories:
- guide
description: How to create translations.
title: Create Translations
weight: 20
---

## Setup Languages

Setup your Translations by creating a YAML file containing the default translation, the list of languages you want to translate, and the output formats for the Translations.

See [Translations]({{< ref "/docs/references/translations" >}}) for details on the format.

Here is an example script:

{{< highlight bash >}}
cat > translations.yaml <<EOF
defaultCode: en
iso639Codes:
  de: Deutsch
  en: English
outputs:
- format: typescript
  package: index
  path: ../web/src/lib/yaml8n
translations:
EOF
{{< /highlight >}}

## Add Translations

You can add translations to this file by creating a new key under `translations:`:

{{< highlight yaml >}}
translations:
  HelloWorld:
    context: Greeting
    en: Hello World
{{< /highlight >}}

{{% alert title="Note" %}}
Start by copying strings from your existing codebase into this file when converting an existing codebase to use YAML8n.  Give the translations meaningful, specific names like WebModalPayment.  If you start sharing translations, add a prefix like Global.
{{% /alert %}}

You don't have to add every language code (in the example above, `de` is missing).  YAML8n will warn on missing translations so you can [add them later]({{< ref "/docs/guides/add-missing-translations" >}}).
