---
categories:
- guide
description: How to generate type safe code from your translations.
title: Generate Code
weight: 30
---

You can start generating code once you have [created your translations]({{< ref "/docs/guides/create-translations" >}}).

## On Demand

Running {{% cli generate %}} will generate code for all of the [`outputs`]({{< ref "/docs/references/translations#outputs" >}}) listed in your [Translations]({{< ref "/docs/references/translations" >}}).

## Watch

Instead of generating code on demand, you can have YAML8n watch your [Translations]({{< ref "/docs/references/translations" >}}) and generate code upon saving them using {{% cli watch %}}.

{{% alert title="Note" %}}
You should absolutely use this while developing!  It makes it easy to add translations and rapidly use new variables/functions.
{{% /alert %}}
