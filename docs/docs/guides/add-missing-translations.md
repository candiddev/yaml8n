---
categories:
- guide
description: How to add missing translations
title: Add Missing Translations
weight: 40
---

YAML8n can add missing translations using [Google Cloud Translation API](https://cloud.google.com/translate).  Follow these steps after creating your [Translations]({{< ref "/docs/references/translations" >}}) to use {{% cli translate %}}.

## Create a Google Cloud Project

Visit [console.cloud.google.com] and create a new Google Cloud project to use with Cloud Translate.  You'll have to add a credit card or some form of payment, but the monthly free tier allowance for Cloud Translate is really generous.  You may not have to pay anything to use this service.

## Enable the Cloud Translate API

Visit https://console.cloud.google.com/translation and select the project you created.  It should prompt you to enable the AutoML API, tap **Enable API**.

## Authenticate YAML8n to Google Cloud

YAML8n needs a way to authenticate to Google Cloud.  There are two ways to do this: via a Service Account or via gcloud application default credentials.

### Using a Service Account

Visit https://console.cloud.google.com/iam-admin/serviceaccounts, select the project you created, and tap **CREATE SERVICE ACCOUNT** at the top.

On the first screen, give the service account a name like `yaml8n-user`.  Tap **CREATE AND CONTINUE**.

On the second screen, add the role **Cloud Translation API User**.  Tap **DONE**.

Back at the service account list, tap the service account you just created.

On the **Service account details** screen, tap **KEYS**.

On the Keys screen, tap **ADD KEY**, tap **Create new key**, tap **CREATE** (keep the default JSON selection).  Save this key somewhere on your device.

You'll need to set an environment variable to use the key with YAML8n:

{{< highlight bash >}}
GOOGLE_APPLICATION_CREDENTIALS=path/to/key yaml8n translate translations.yaml
{{< /highlight >}}

### Using gcloud

Follow [these instructions](https://cloud.google.com/sdk/docs/install) to install the gcloud CLI.

Once installed, run `gcloud auth application-default login`.  You should be able to run `yaml8n translate translations.yaml` now.
