---
categories:
- feature
description: YAML8n uses AI to create translations automatically.
title: AI-Powered Translations
type: docs
---

YAML8n uses [Google Cloud Translation API](https://cloud.google.com/translate) to generate missing translations automatically:

{{< highlight console >}}
$ ./yaml8n translate translations.yaml
[NOTICE]
7 Warnings Found:
- EmailDailyAgendaHeader is missing a translation: hi
- EmailDailyAgendaHeader is missing a translation: nl
- EmailDailyAgendaHeader is missing a translation: zh
- EmailDailyAgendaHeader is missing a translation: ar
- EmailDailyAgendaHeader is missing a translation: de
- EmailDailyAgendaHeader is missing a translation: es
- EmailDailyAgendaHeader is missing a translation: fr

Need to translate 259 characters (type OK to continue): OK
Translating 48/86 - es...
Translating 48/86 - fr...
Translating 48/86 - hi...
Translating 48/86 - nl...
Translating 48/86 - zh...
Translating 48/86 - ar...
Translating 48/86 - de...
{{< /highlight >}}
