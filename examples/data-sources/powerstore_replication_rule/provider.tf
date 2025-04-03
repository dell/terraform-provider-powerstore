/*
Copyright (c) 2025 Dell Inc., or its subsidiaries. All Rights Reserved.

Licensed under the Mozilla Public License Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://mozilla.org/MPL/2.0/


Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

terraform {
  required_providers {
    powerstore = {
      version = "1.2.0"
      source  = "registry.terraform.io/dell/powerstore"
    }
  }
}

provider "powerstore" {
  username = var.username
  password = var.password
  endpoint = var.endpoint
  insecure = true
  timeout  = var.timeout

  ## Provider can also be set using environment variables
  ## If environment variables are set it will override this configuration
  ## Example environment variables
  # POWERSTORE_USERNAME="username"
  # POWERSTORE_PASSWORD="password"
  # POWERSTORE_ENDPOINT="https://yourhost.host.com/api/rest"
  # POWERSTORE_INSECURE="false"
  # POWERSTORE_TIMEOUT="120"
}