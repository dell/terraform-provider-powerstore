<!--
Copyright (c) 2022 Dell Inc., or its subsidiaries. All Rights Reserved.
Licensed under the Mozilla Public License Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at
    http://mozilla.org/MPL/2.0/
Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
-->
# Terraform Provider for Dell Technologies PowerStore

[![Contributor Covenant](https://img.shields.io/badge/Contributor%20Covenant-v2.0%20adopted-ff69b4.svg)](about/CODE_OF_CONDUCT.md)
[![License](https://img.shields.io/badge/License-MPL_2.0-blue.svg)](LICENSE)


The Terraform Provider for Dell Technologies (Dell) PowerStore allows Data Center and IT administrators to use Hashicorp Terraform to automate and orchestrate the provisioning and management of Dell PowerStore storage systems.

The Terraform Provider can be used to manage volumes, snapshot rules, protection policies and storage containers.

## Table of contents

* [Support](#support)
* [License](#license)
* [Prerequisites](#prerequisites)
* [List of Resources in Terraform Provider for Dell PowerStore](#list-of-resources-in-terraform-provider-for-dell-powerstore)
* [List of DataSources in Terraform Provider for Dell PowerStore](#list-of-datasources-in-terraform-provider-for-dell-powerstore)
* [Releasing, Maintenance and Deprecation](#releasing-maintenance-and-deprecation)

## Support
For any Terraform Provider for Dell PowerStore issues, questions or feedback, please follow our [support process](https://github.com/dell/dell-terraform-providers/blob/main/docs/SUPPORT.md)

## License
The Terraform Provider for PowerStore is released and licensed under the MPL-2.0 license. See [LICENSE](https://github.com/dell/terraform-provider-powerstore/blob/main/LICENSE) for the full terms.

## Prerequisites

| **Terraform Provider** | **PowerStore Version** | **OS** | **Terraform** | **Golang**
|---------------------|-----------------------|-------|--------------------|--------------------------|
| v1.0.0 | 3.0 | Ubuntu 22.04 <br> RHEL 8.x <br> RHEL 7.x | 1.3.2 <br> 1.2.9 <br> | 1.19.x

## List of Resources in Terraform Provider for Dell PowerStore
  * Volume
  * Snapshot Rule
  * Protection Policy
  * Storage Container

## List of DataSources in Terraform Provider for Dell PowerStore
  * Volume

## Installation of Terraform Provider for Dell PowerFStore

## Installation from Terraform Registry

The provider will be fetched from the Terraform registry and installed by Terraform automatically.
Create a file called `main.tf` in your workspace with the following contents

```terraform
terraform {
  required_providers {
    powerstore = {
      version = "1.0.0"
      source = "registry.terraform.io/dell/powerstore"
    }
  }
}
```
Then, in that workspace, run
```
terraform init
```

## Installation from source code

Dependencies: Go 1.19.x, make, Terraform 1.2.9/1.3.2
<br>
<br>
Run
```
git clone https://github.com/dell/terraform-provider-powerstore.git
cd terraform-provider-powerstore
make install
```
Then follow [installation from Terraform registry](#installation-from-terraform-registry)

## SSL Certificate Verification

For SSL verifcation on RHEL, these steps can be performed:
 * Copy the CA certificate to the `/etc/pki/ca-trust/source/anchors` path of the host by any external means.
 * Import the SSL certificate to host by running
```
update-ca-trust extract
```
For SSL verification on Ubuntu, these steps can be performed:
 * Copy the CA certificate to the `/etc/ssl/certs` path of the host by any external means.
 * Import the SSL certificate to host by running:
 ```
  update-ca-certificates
```

## Releasing, Maintenance and Deprecation

Terraform Provider for Dell Technnologies PowerStore follows [Semantic Versioning](https://semver.org/).

New version will be released regularly if significant changes(bug fixes or new features) are made in the provider.

Released code versions are located on tags with names of the form "vx.y.z" where x.y.z corresponds to the version number.