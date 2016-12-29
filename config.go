package sendgrid

import (
	"log"

	"github.com/romanlaguta/terraform-client-sendgrid"
)

// Config holds API key to authenticate to Sendgrid.
type Config struct {
	APIKey string
}

// Client returns a new Sendgrid client.
func (c *Config) Client() *sendgrid_client.Client {

	client := sendgrid_client.NewClient(c.APIKey)
	log.Printf("[INFO] Sendgrid Client configured ")

	return client
}
