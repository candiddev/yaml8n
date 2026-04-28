---
categories:
- guide
description: How to integrate YAML8n within your CI/CD pipelines.
title: CI/CD
weight: 50
---

YAML8n works really well in a Continuous Integration/Continuous Delivery (CI/CD) pipeline.

## Check Generated Code

It's recommended to run {{% cli generate %}} follow by a `git diff` to check if the output code changes.  This should fail builds, as they aren't using the latest translations.

{{< highlight bash >}}
yaml8n generate translations.yaml
git diff --exit-code outputs/yaml8n.ts
{{< /highlight >}}

## Validate Translations

You can validate the [Translations]({{< ref "/docs/references/translations" >}}) by running {{% cli validate %}}.  This command will walk the translations and ensure they have the correct syntax.

Additionally, you can pass the argument `-w` to fail on warnings like missing translations.

{{< highlight bash >}}
yaml8n validate translations.yaml
{{< /highlight >}}
