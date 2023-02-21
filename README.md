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
# Terraform Provider for PowerStore

[![Contributor Covenant](https://img.shields.io/badge/Contributor%20Covenant-v2.0%20adopted-ff69b4.svg)](about/CODE_OF_CONDUCT.md)
[![License](https://img.shields.io/badge/License-MPL_2.0-blue.svg)](LICENSE)


The Terraform Provider for PowerStore is a plugin for Terraform that allows the resource management for PowerStore appliance. This provider is built by Dell Technologies CTIO (Chief Technology & Innovation Office) team. For more details on PowerStore, please refer to PowerStore Official webpage [here][powerstore-website].

For general information about Terraform, visit the [official website][tf-website] and the [GitHub project page][tf-github].

[tf-website]: https://terraform.io
[tf-github]: https://github.com/hashicorp/terraform
[powerstore-website]: https://www.delltechnologies.com/en-in/storage/powerstore-storage-appliance.htm

## Requirements

- [Terraform](https://www.terraform.io/downloads.html) 0.13.x and later
- [Go](https://golang.org/doc/install) 1.13.x and later (to build the provider plugin) 

## Development

*Note*: This project uses [Go modules](https://blog.golang.org/using-go-modules) making it safe to work with it outside of your existing [GOPATH](http://golang.org/doc/code.html#GOPATH). We assuming the dir you choose is in variable $DIR

### steps

1. clone repo and build the provider binary
```bash
$ mkdir -p $DIR && cd $DIR
$ git clone https://github.com/dell/terraform-provider-powerstore.git
$ go mod download
$ go build
```

2. configure terraform to use this new build provider binary instead of fetching from [terraform registry](https://registry.terraform.io/)
    1. create a file `.terraformrc` in home directory and copy the following content
    2. replace DIR_absolute_path with the $DIR content i.e. the absolute path where repo is cloned

```bash
$ cat ~/.terraformrc
provider_installation {

  dev_overrides {
      "powerstore.com/powerstoreprovider/powerstore" = "DIR_absolute_path"
  }

  direct {}
}
```

3. run example
```bash
$ cd $DIR/examples
# create a .auto.tfvars file with variables and value
# this var file will be auto loaded
$ cat $DIR/examples/.auto.tfvars
username=""
password=""
endpoint=""
$ terraform apply
...
^C
```

4. for debugging
```bash
export TF_LOG_PATH=/tmp/logfile
export GOPOWERSTORE_DEBUG="TRUE"
export TF_LOG="debug"
```

## Documentation

The documentation for the provider resources can found [here](https://github.com/dell/terraform-provider-powerstore/tree/main/docs/resources)

## Roadmap

Our roadmap for Terraform provider for PowerStore resources can be found [here](https://github.com/dell/terraform-provider-powerstore/tree/main/docs/ROADMAP.md)

## Contributing

The Terraform PowerStore provider is open-source and community supported.

For issues, questions or feedback, join the [Dell EMC Automation community](https://www.dell.com/community/Automation/bd-p/Automation).


