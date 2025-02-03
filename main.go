/*
Copyright (c) 2024 Dell Inc., or its subsidiaries. All Rights Reserved.

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

package main

import (
	"context"
	"flag"
	"log"
	"terraform-provider-powerstore/powerstore"

	"github.com/hashicorp/terraform-plugin-framework/providerserver"
)

// Run the docs generation tool, check its repository for more information on how it works and how docs
// can be customized.
//go:generate go run github.com/hashicorp/terraform-plugin-docs/cmd/tfplugindocs@v0.15.0

// Run the copyright generation tool
//go:generate go run -tags tools ./tools

var (
	version string = "dev"
)

func main() {

	var debug bool

	flag.BoolVar(&debug, "debug", false, "set to true to run the provider with support for debuggers like delve")
	flag.Parse()

	err := providerserver.Serve(context.Background(), powerstore.New(version), providerserver.ServeOpts{
		Address: "registry.terraform.io/dell/powerstore",
		Debug:   debug,
	})

	if err != nil {
		log.Fatal(err.Error())
	}
}
