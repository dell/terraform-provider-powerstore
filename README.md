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
* [Documentation](#documentation)
* [New to Terraform?](#new-to-terraform)

## Support
For any Terraform Provider for Dell PowerStore issues, questions or feedback, please follow our [support process](https://github.com/dell/dell-terraform-providers/blob/main/docs/SUPPORT.md)

## License
The Terraform Provider for PowerStore is released and licensed under the MPL-2.0 license. See [LICENSE](https://github.com/dell/terraform-provider-powerstore/blob/main/LICENSE) for the full terms.

## Prerequisites

| **Terraform Provider** | **PowerStore Version** | **OS** | **Terraform** | **Golang**
|---------------------|-----------------------|-------|--------------------|--------------------------|
| v1.1.3 | 3.2/3.5/3.6/4.0 | Ubuntu 22.04 <br> RHEL 9.x | 1.5.x <br> 1.6.x <br> | 1.22.x

## List of Resources in Terraform Provider for Dell PowerStore
  * [Volume](docs/resources/volume.md)
  * [Volume Group](docs/resources/volumegroup.md)
  * [Volume Snapshot](docs/resources/volume_snapshot.md)
  * [Volume Group Snapshot](docs/resources/volumegroup_snapshot.md)
  * [Snapshot Rule](docs/resources/snapshotrule.md)
  * [Protection Policy](docs/resources/protectionpolicy.md)
  * [Storage Container](docs/resources/storagecontainer.md)
  * [Host](docs/resources/host.md)
  * [Host Group](docs/resources/hostgroup.md)

## List of DataSources in Terraform Provider for Dell PowerStore
  * [Volume](docs/data-sources/volume.md)
  * [Volume Group](docs/data-sources/volumegroup.md)
  * [Volume Snapshot](docs/data-sources/volume_snapshot.md)
  * [Volume Group Snapshot](docs/data-sources/volumegroup_snapshot.md)
  * [Host](docs/data-sources/host.md)
  * [Host Group](docs/data-sources/hostgroup.md)
  * [Snapshot Rule](docs/data-sources/snapshotrule.md)
  * [Protection Policy](docs/data-sources/protectionpolicy.md)


## Installation of Terraform Provider for Dell PowerStore

## Installation from Terraform Registry

The provider will be fetched from the Terraform registry and installed by Terraform automatically.
Create a file called `main.tf` in your workspace with the following contents

```terraform
terraform {
  required_providers {
    powerstore = {
      version = "1.1.3"
      source = "registry.terraform.io/dell/powerstore"
    }
  }
}
```
Then, in that workspace, run
```
terraform init
```

If you are upgrading from a previous version, set the version of powerstore in the required providers block to "1.1.0" as shown above.
Then, in your workspace, run
```
terraform init -upgrade
```
For more details on how to upgrade provider versions, please check out https://developer.hashicorp.com/terraform/tutorials/configuration-language/provider-versioning

## Installation from source code

Dependencies: Go 1.22.x, make, Terraform 1.5.x/1.6.x
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

## Documentation
For more detailed information, please refer to [Dell Terraform Providers Documentation](https://dell.github.io/terraform-docs/).

## New to Terraform?
**Here are some helpful links to get you started if you are new to terraform before using our provider:**

- Intro to Terraform: https://developer.hashicorp.com/terraform/intro 
- Providers: https://developer.hashicorp.com/terraform/language/providers 
- Resources: https://developer.hashicorp.com/terraform/language/resources
- Datasources: https://developer.hashicorp.com/terraform/language/data-sources
- Import: https://developer.hashicorp.com/terraform/language/import
- Variables: https://developer.hashicorp.com/terraform/language/values/variables
- Modules: https://developer.hashicorp.com/terraform/language/modules
- State: https://developer.hashicorp.com/terraform/language/state