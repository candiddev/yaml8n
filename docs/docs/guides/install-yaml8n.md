---
categories:
- guide
description: How to install YAML8n
title: Install YAML8n
weight: 10
---

Installing YAML8n depends on how you want to run it.  YAML8n is available as a [binary](#binary) or a [container](#container).

## Binary

YAML8n binaries are available for various architectures and operating systems:

{{% release %}}

{{% alert title="Updating YAML8n" color="primary" %}}
YAML8n can be updated by replacing the binary with the latest version.
{{% /alert %}}


## Container

YAML8n containers are available on [GitHub](https://github.com/candiddev/yaml8n/pkgs/container/yaml8n).

You can create an alias to run YAML8n as a container:

{{< highlight bash >}}
alias yaml8n='docker run -u $(id -u):$(id -g) -it --rm -v $(pwd):/work -w /work ghcr.io/candiddev/yaml8n:latest'
{{< /highlight >}}

## SBOM

YAML8n ships with a Software Bill of Materials (SBOM) manifest generated using [CycloneDX](https://cyclonedx.org/).  The `.bom.json` manifest is available with the other [Binary Assets](#binary).
