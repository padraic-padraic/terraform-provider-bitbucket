package main

import (
	"flag"

	"github.com/hashicorp/terraform-plugin-sdk/v2/plugin"
	"github.com/padraic-padraic/terraform-provider-bitbucket/v2/bitbucket"
)

func main() {
	var debug bool

	flag.BoolVar(&debug, "debug", false, "set to true to run the provider with support for debuggers like delve")
	flag.Parse()

	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: bitbucket.Provider,
		ProviderAddr: "DrFaust92/bitbucket",
		Debug:        debug,
	})
}
