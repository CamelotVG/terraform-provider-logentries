package main

import (
	"github.com/hashicorp/terraform-plugin-sdk/plugin"
	"github.com/camelotvg/terraform-provider-logentries/logentries"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: logentries.Provider})
}
