package sendgrid

import (
	"fmt"
	"log"

	"errors"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"
)

func Provider() terraform.ResourceProvider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"api_key": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("SENDGRID_API_KEY", nil),
			},
		},

		ResourcesMap: map[string]*schema.Resource{
			"sendgrid_template": resourceSendgridTemplate(),
		},

		ConfigureFunc: providerConfigure,
	}
}

func providerConfigure(d *schema.ResourceData) (interface{}, error) {

	config := Config{
		APIKey: d.Get("api_key").(string),
	}

	log.Println("[INFO] Initializing Sendgrid client")
	client := config.Client()
	fmt.Println("Validate template")
	ok, err := client.Validate()

	if err != nil {
		return client, err
	}

	if ok == false {
		return client, errors.New(`No valid credential sources found for Sendgrid Provider. Please see https://terraform.io/docs/providers/sendgrid/index.html for more information on providing credentials for the Sendgrid Provider`)
	}

	return client, nil
}
