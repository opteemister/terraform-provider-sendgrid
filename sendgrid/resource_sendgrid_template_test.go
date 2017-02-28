package sendgrid

import (
	"fmt"
	"strings"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
	"github.com/opteemister/terraform-client-sendgrid"
)

func TestAccSendgridTemplate_Basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckSendgridTemplateDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckSendgridTemplateConfig,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSendgridTemplateExists("sendgrid_template.foo"),
					resource.TestCheckResourceAttr(
						"sendgrid_template.foo", "name", "name for template foo"),
				),
			},
		},
	})
}

func TestAccSendgridTemplate_Updated(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckSendgridTemplateDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckSendgridTemplateConfig,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSendgridTemplateExists("sendgrid_template.foo"),
					resource.TestCheckResourceAttr(
						"sendgrid_template.foo", "name", "name for template foo"),
				),
			},
			resource.TestStep{
				Config: testAccCheckSendgridTemplateConfigUpdated,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSendgridTemplateExists("sendgrid_template.foo"),
					resource.TestCheckResourceAttr(
						"sendgrid_template.foo", "name", "name for template bar"),
				),
			},
		},
	})
}

func testAccCheckSendgridTemplateDestroy(s *terraform.State) error {
	fmt.Println("testAccCheckSendgridTemplateDestroy")
	client := testAccProvider.Meta().(*sendgrid_client.Client)

	if err := destroyHelper(s, client); err != nil {
		fmt.Println("testAccCheckSendgridTemplateDestroy error: ", err)
		return err
	}
	return nil
}

func testAccCheckSendgridTemplateExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		client := testAccProvider.Meta().(*sendgrid_client.Client)
		if err := existsHelper(s, client); err != nil {
			return err
		}
		return nil
	}
}

const testAccCheckSendgridTemplateConfig = `
resource "sendgrid_template" "foo" {
  name = "name for template foo"
}
`

const testAccCheckSendgridTemplateConfigUpdated = `
resource "sendgrid_template" "foo" {
  name = "name for template bar"
}
`

func destroyHelper(s *terraform.State, client *sendgrid_client.Client) error {
	fmt.Println("destroyHelper")
	for _, r := range s.RootModule().Resources {
		fmt.Println(r.Type)
		if r.Type == "sendgrid_template" {
			fmt.Println("Delete")
			if _, err := client.GetTemplate(r.Primary.ID); err != nil {
				if strings.Contains(err.Error(), "404") {
					continue
				}
				return fmt.Errorf("Received an error retrieving template %s", err)
			}
			return fmt.Errorf("Template still exists")
		}
	}
	return nil
}

func existsHelper(s *terraform.State, client *sendgrid_client.Client) error {
	for _, r := range s.RootModule().Resources {
		if r.Type == "sendgrid_template" {
			if _, err := client.GetTemplate(r.Primary.ID); err != nil {
				return fmt.Errorf("Received an error retrieving template %s", err)
			}
		}
	}
	return nil
}
