# Bitrise Step: Fetch Infisical Secrets

A Bitrise step that allows you to fetch secrets from Infisical and make them available in your Bitrise workflows.

## Description

This step enables secure access to your Infisical secrets during Bitrise CI/CD pipelines. It fetches the secrets from your Infisical instance and makes them available as environment variables for subsequent steps in your workflow.

## Setup

1. Add this step to your `bitrise.yml` workflow
2. Configure the required inputs in your project or workflow settings
3. The secrets will be automatically fetched and made available to your build process

## Inputs

*Configure these input values in your workflow:*

| Input | Description                                            | Required | Default                   |
|-------|--------------------------------------------------------|----------|---------------------------|
| infisical_client | Infisical Universal Auth Client ID                     | yes      | -                         |
| infisical_client_secret | Infisical Universal Auth Secret                        | yes      | -                         |
| infisical_project_id | Infisical project identifier                           | yes      | -                         |
| infisical_env | Infisical environment slug                             | yes      | -                         | 
| infisical_url | URL for your infisical application (should end in /api) | no       | https://app.infisical.com |
| infisical_path | The path to your secrets                               | no       | /                         |

## Usage

Add this step to your `bitrise.yml`: