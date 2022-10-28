package main

import (
	"context"
	"flag"
	"github.com/hashicorp/terraform-plugin-framework/providerserver"
	"log"
	"terraform-provider-powerstore/powerstore"
)

var (
	version string = "dev"
)

func main() {

	var debug bool

	flag.BoolVar(&debug, "debug", false, "set to true to run the provider with support for debuggers like delve")
	flag.Parse()

	err := providerserver.Serve(context.Background(), powerstore.New(version), providerserver.ServeOpts{
		Address: "powerstore.com/powerstoreprovider/powerstore",
		Debug:   debug,
	})

	if err != nil {
		log.Fatal(err.Error())
	}
}
