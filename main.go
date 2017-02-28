package main

import (
	"github.com/hashicorp/terraform/tree/0-8-stable/plugin"
	"github.com/opteemister/terraform-provider-sendgrid/sendgrid"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: sendgrid.Provider,
	})
}
