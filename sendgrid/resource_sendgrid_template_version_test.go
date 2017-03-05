package sendgrid

import (
	"fmt"
	"io/ioutil"
	"strings"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
	"github.com/mitchellh/go-homedir"
	"github.com/opteemister/terraform-client-sendgrid"
)

func TestAccSendgridTemplateVersion_Basic(t *testing.T) {
	htmlContent, err := loadFileContent("./resources/test_template.html")
	if err != nil {
		t.Error("Can't read template file")
	}

	plainContent, err := loadFileContent("./resources/test_template_plain.html")
	if err != nil {
		t.Error("Can't read template file")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckSendgridTemplateVersionDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckSendgridTemplateVersionConfig,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSendgridTemplateVersionExists("sendgrid_template_version.foo"),
					resource.TestCheckResourceAttr(
						"sendgrid_template_version.foo", "name", "name for template version foo"),
					resource.TestCheckResourceAttr(
						"sendgrid_template_version.foo", "subject", "foo subject"),
					resource.TestCheckResourceAttr(
						"sendgrid_template_version.foo", "html_content_hash", getHash(string(htmlContent))),
					resource.TestCheckResourceAttr(
						"sendgrid_template_version.foo", "plain_content_hash", getHash(string(plainContent))),
				),
			},
		},
	})
}

func TestAccSendgridTemplateVersionNotActive(t *testing.T) {
	htmlContent, err := loadFileContent("./resources/test_template.html")
	if err != nil {
		t.Error("Can't read template file")
	}

	plainContent, err := loadFileContent("./resources/test_template_plain.html")
	if err != nil {
		t.Error("Can't read template file")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckSendgridTemplateVersionDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckSendgridTemplateVersionConfigNotActive,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSendgridTemplateVersionExists("sendgrid_template_version.foo"),
					resource.TestCheckResourceAttr(
						"sendgrid_template_version.foo", "name", "name for template version foo"),
					resource.TestCheckResourceAttr(
						"sendgrid_template_version.foo", "subject", "foo subject"),
					resource.TestCheckResourceAttr(
						"sendgrid_template_version.foo", "html_content_hash", getHash(string(htmlContent))),
					resource.TestCheckResourceAttr(
						"sendgrid_template_version.foo", "plain_content_hash", getHash(string(plainContent))),
					resource.TestCheckResourceAttr(
						"sendgrid_template_version.foo", "active", "false"),
				),
			},
		},
	})
}

func TestAccSendgridTemplateVersion_Updated(t *testing.T) {
	htmlContent, err := loadFileContent("./resources/test_template.html")
	if err != nil {
		t.Error("Can't read template file")
	}

	plainContent, err := loadFileContent("./resources/test_template_plain.html")
	if err != nil {
		t.Error("Can't read template file")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckSendgridTemplateVersionDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckSendgridTemplateVersionConfig,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSendgridTemplateVersionExists("sendgrid_template_version.foo"),
					resource.TestCheckResourceAttr(
						"sendgrid_template_version.foo", "name", "name for template version foo"),
					resource.TestCheckResourceAttr(
						"sendgrid_template_version.foo", "subject", "foo subject"),
					resource.TestCheckResourceAttr(
						"sendgrid_template_version.foo", "html_content_hash", getHash(string(htmlContent))),
					resource.TestCheckResourceAttr(
						"sendgrid_template_version.foo", "plain_content_hash", getHash(string(plainContent))),
					resource.TestCheckResourceAttr(
						"sendgrid_template_version.foo", "active", "true"),
				),
			},
			resource.TestStep{
				Config: testAccCheckSendgridTemplateVersionConfigUpdated,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSendgridTemplateVersionExists("sendgrid_template_version.foo"),
					resource.TestCheckResourceAttr(
						"sendgrid_template_version.foo", "name", "name for template version bar"),
					resource.TestCheckResourceAttr(
						"sendgrid_template_version.foo", "subject", "bar subject"),
					resource.TestCheckResourceAttr(
						"sendgrid_template_version.foo", "html_content_hash", getHash(string(htmlContent))),
					resource.TestCheckResourceAttr(
						"sendgrid_template_version.foo", "plain_content_hash", getHash(string(plainContent))),
					resource.TestCheckResourceAttr(
						"sendgrid_template_version.foo", "active", "false"),
				),
			},
		},
	})
}

func TestAccSendgridTemplateVersion_UpdatedContent(t *testing.T) {
	content1 := "content1"
	content2 := "content2"
	plain_content1 := "plain_content1"
	plain_content2 := "plain_content2"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckSendgridTemplateVersionDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				PreConfig: func() {
					writeContent("./resources/temp_template.html", content1)
					writeContent("./resources/temp_template_plain.html", plain_content1)
				},
				Config: testAccCheckSendgridTemplateVersionConfigUpdatedContent,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSendgridTemplateVersionExists("sendgrid_template_version.foo"),
					resource.TestCheckResourceAttr(
						"sendgrid_template_version.foo", "name", "name for template version foo"),
					resource.TestCheckResourceAttr(
						"sendgrid_template_version.foo", "subject", "foo subject"),
					resource.TestCheckResourceAttr(
						"sendgrid_template_version.foo", "html_content_hash", getHash(content1)),
					resource.TestCheckResourceAttr(
						"sendgrid_template_version.foo", "plain_content_hash", getHash(plain_content1)),
					resource.TestCheckResourceAttr(
						"sendgrid_template_version.foo", "active", "true"),
				),
			},
			resource.TestStep{
				PreConfig: func() {
					writeContent("./resources/temp_template.html", content2)
					writeContent("./resources/temp_template_plain.html", plain_content2)
				},
				Config: testAccCheckSendgridTemplateVersionConfigUpdatedContent,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSendgridTemplateVersionExists("sendgrid_template_version.foo"),
					resource.TestCheckResourceAttr(
						"sendgrid_template_version.foo", "name", "name for template version foo"),
					resource.TestCheckResourceAttr(
						"sendgrid_template_version.foo", "subject", "foo subject"),
					resource.TestCheckResourceAttr(
						"sendgrid_template_version.foo", "html_content_hash", getHash(content2)),
					resource.TestCheckResourceAttr(
						"sendgrid_template_version.foo", "plain_content_hash", getHash(plain_content2)),
					resource.TestCheckResourceAttr(
						"sendgrid_template_version.foo", "active", "true"),
				),
			},
		},
	})
}

func writeContent(file string, content string) error {
	filename, err := homedir.Expand(file)
	if err != nil {
		fmt.Println("File %s can't be expand. %s", file, err)
		return err
	}
	err = ioutil.WriteFile(filename, []byte(content), 0644)
	if err != nil {
		fmt.Println("File %s can't be written. %s", filename, err)
		return err
	}
	return nil
}

func testAccCheckSendgridTemplateVersionDestroy(s *terraform.State) error {
	client := testAccProvider.Meta().(*sendgrid_client.Client)

	if err := destroyVersionHelper(s, client); err != nil {
		return err
	}
	return nil
}

func testAccCheckSendgridTemplateVersionExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		client := testAccProvider.Meta().(*sendgrid_client.Client)
		if err := existsVersionHelper(s, client); err != nil {
			return err
		}
		return nil
	}
}

const testAccCheckSendgridTemplateVersionConfig = `
resource "sendgrid_template" "foo" {
  name = "name for template foo"
}

resource "sendgrid_template_version" "foo" {
  name = "name for template version foo"
	template_id = "${sendgrid_template.foo.id}"
	subject = "foo subject"
	html_content_file = "./resources/test_template.html"
	plain_content_file = "./resources/test_template_plain.html"
}
`

const testAccCheckSendgridTemplateVersionConfigNotActive = `
resource "sendgrid_template" "foo" {
  name = "name for template foo"
}

resource "sendgrid_template_version" "foo" {
  name = "name for template version foo"
	template_id = "${sendgrid_template.foo.id}"
	subject = "foo subject"
	html_content_file = "./resources/test_template.html"
	plain_content_file = "./resources/test_template_plain.html"
  active = false
}
`

const testAccCheckSendgridTemplateVersionConfigUpdated = `
resource "sendgrid_template" "foo" {
  name = "name for template foo"
}

resource "sendgrid_template_version" "foo" {
  name = "name for template version bar"
	template_id = "${sendgrid_template.foo.id}"
	subject = "bar subject"
	html_content_file = "./resources/test_template.html"
	plain_content_file = "./resources/test_template_plain.html"
  active = false
}
`

const testAccCheckSendgridTemplateVersionConfigUpdatedContent = `
resource "sendgrid_template" "foo" {
  name = "name for template foo"
}

resource "sendgrid_template_version" "foo" {
  name = "name for template version foo"
	template_id = "${sendgrid_template.foo.id}"
	subject = "foo subject"
	html_content_file = "./resources/temp_template.html"
	plain_content_file = "./resources/temp_template_plain.html"
	html_content_hash = "${base64sha256(file("./resources/temp_template.html"))}"
	plain_content_hash = "${base64sha256(file("./resources/temp_template_plain.html"))}"
}
`

func destroyVersionHelper(s *terraform.State, client *sendgrid_client.Client) error {
	for _, r := range s.RootModule().Resources {
		fmt.Println(r.Type)
		if r.Type == "sendgrid_template_version" {
			fmt.Println("Delete")
			if _, err := client.GetTemplateVersion(r.Primary.Attributes["template_id"], r.Primary.ID); err != nil {
				if strings.Contains(err.Error(), "404") {
					continue
				}
				return fmt.Errorf("Received an error retrieving template version %s", err)
			}
			return fmt.Errorf("Template version still exists")
		}
	}
	return destroyHelper(s, client)
}

func existsVersionHelper(s *terraform.State, client *sendgrid_client.Client) error {
	for _, r := range s.RootModule().Resources {
		fmt.Println(r.Type)
		if r.Type == "sendgrid_template_version" {
			if _, err := client.GetTemplateVersion(r.Primary.Attributes["template_id"], r.Primary.ID); err != nil {
				return fmt.Errorf("Received an error retrieving template version %s", err)
			}
		}
	}
	return existsHelper(s, client)
}
