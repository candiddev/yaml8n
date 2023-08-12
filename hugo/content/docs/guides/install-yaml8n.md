---
categories:
- guide
title: Install YAML8n
weight: 10
---

Installing YAML8n depends on how you want to run it.  YAML8n is available as a [binary](#binary) or a [container](#container).

## Binary

YAML8n binaries are available on [GitHub](https://github.com/candiddev/yaml8n/releases).

{{< tabpane text=true >}}
{{< tab header="Linux amd64" >}}
{{< highlight bash >}}
curl -L https://github.com/candiddev/yaml8n/releases/latest/download/yaml8n_linux_amd64.tar.gz -O
curl -L https://github.com/candiddev/yaml8n/releases/latest/download/yaml8n_linux_amd64.tar.gz.sha256 -O
sha256sum -c yaml8n_linux_amd64.tar.gz.sha256
tar -xzf yaml8n_linux_amd64.tar.gz
{{< /highlight >}}
{{< /tab >}}

{{< tab header="Linux arm" >}}
{{< highlight bash >}}
curl -L https://github.com/candiddev/yaml8n/releases/latest/download/yaml8n_linux_arm.tar.gz -O
curl -L https://github.com/candiddev/yaml8n/releases/latest/download/yaml8n_linux_arm.tar.gz.sha256 -O
sha256sum -c yaml8n_linux_arm.tar.gz.sha256
tar -xzf yaml8n_linux_arm.tar.gz
{{< /highlight >}}
{{< /tab >}}

{{< tab header="Linux arm64" >}}
{{< highlight bash >}}
curl -L https://github.com/candiddev/yaml8n/releases/latest/download/yaml8n_linux_arm64.tar.gz -O
curl -L https://github.com/candiddev/yaml8n/releases/latest/download/yaml8n_linux_arm64.tar.gz.sha256 -O
sha256sum -c yaml8n_linux_arm64.tar.gz.sha256
tar -xzf yaml8n_linux_arm64.tar.gz
{{< /tab >}}
{{< /highlight >}}

{{< tab header="OSX amd64" >}}
{{< highlight bash >}}
curl -L https://github.com/candiddev/yaml8n/releases/latest/download/yaml8n_darwin_amd64.tar.gz -O
curl -L https://github.com/candiddev/yaml8n/releases/latest/download/yaml8n_darwin_amd64.tar.gz.sha256 -O
sha256sum -c yaml8n_darwin_amd64.tar.gz.sha256
tar -xzf yaml8n_darwin_amd64.tar.gz
{{< /highlight >}}
{{< /tab >}}

{{< tab header="OSX arm64" >}}
{{< highlight bash >}}
curl -L https://github.com/candiddev/yaml8n/releases/latest/download/yaml8n_darwin_arm64.tar.gz -O
curl -L https://github.com/candiddev/yaml8n/releases/latest/download/yaml8n_darwin_arm64.tar.gz.sha256 -O
sha256sum -c yaml8n_darwin_arm64.tar.gz.sha256
tar -xzf yaml8n_darwin_arm64.tar.gz
{{< /highlight >}}
{{< /tab >}}

{{< /tabpane >}}

## Container

YAML8n containers are available on [GitHub](https://github.com/candiddev/yaml8n/pkgs/container/yaml8n).

You can create an alias to run YAML8n as a container:

{{< highlight bash >}}
alias yaml8n='docker run -u $(id -u):$(id -g) -it --rm -v $(pwd):/work -w /work ghcr.io/candiddev/yaml8n:latest'
{{< /highlight >}}
