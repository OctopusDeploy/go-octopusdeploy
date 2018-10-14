# Validate Octopus Objects

Date: 2018-10-14

## Status

:white_check_mark: Accepted

## Context

When constructing Octopus Deploy objects, such a project, some settings are mandatory, and others accept some specific properties, for example:

* A Projects **name** is mandatory `string` property.
* A Projects **Default failure mode** supports a string of values `EnvironmentDefault`, `Off` or `On`.

The API will reject Adds or Updates with incorrect or missing properties.

## Decision

* As this client will be used mainly for a Terraform provider, doing validation first using `terraform plan` would catch these errors before a `terraform apply` if they are validated via the client.

* There will be a single go file that contains all of the valid values for the different objects, which can then be used in the Terraform provider for validation. This saves storing the valid properties in both the provider and the client.

## Consequences

* **Harder:** Updates to the Octopus Deploy API will need updating in this client.

* **Easier:** There is a single place to manage these settings, making it easy to use them from the Terraform provide and keep things updated.
