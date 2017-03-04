package sendgrid

import (
	"fmt"
	"strings"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/opteemister/terraform-client-sendgrid"
)

func resourceSendgridTemplate() *schema.Resource {
	return &schema.Resource{
		Create: resourceSendgridTemplateCreate,
		Read:   resourceSendgridTemplateRead,
		Update: resourceSendgridTemplateUpdate,
		Delete: resourceSendgridTemplateDelete,
		Exists: resourceSendgridTemplateExists,
		Importer: &schema.ResourceImporter{
			State: resourceSendgridTemplateImport,
		},

		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func buildTemplateStruct(d *schema.ResourceData) *sendgrid_client.Template {
	m := sendgrid_client.Template{
		Name: d.Get("name").(string),
	}

	return &m
}

func resourceSendgridTemplateExists(d *schema.ResourceData, meta interface{}) (b bool, e error) {
	// Exists - This is called to verify a resource still exists. It is called prior to Read,
	// and lowers the burden of Read to be able to assume the resource exists.
	client := meta.(*sendgrid_client.Client)

	fmt.Println("Exist template")
	if _, err := client.GetTemplate(d.Id()); err != nil {
		if strings.Contains(err.Error(), "404 Not Found") {
			return false, nil
		}
		return false, fmt.Errorf("error check existance template: %s", err.Error())
	}

	return true, nil
}

func resourceSendgridTemplateCreate(d *schema.ResourceData, meta interface{}) error {

	client := meta.(*sendgrid_client.Client)

	m := buildTemplateStruct(d)
	fmt.Println("Create template1")
	m, err := client.CreateTemplate(m)
	if err != nil {
		return fmt.Errorf("error updating template: %s", err.Error())
	}
	fmt.Println("Create template2")
	d.SetId(m.Id)

	return nil
}

func resourceSendgridTemplateRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*sendgrid_client.Client)

	fmt.Println("Read template")
	m, err := client.GetTemplate(d.Id())
	if err != nil {
		return fmt.Errorf("error reading template: %s", err.Error())
	}
	fmt.Println("[DEBUG] Template: %v", m)
	d.Set("name", m.Name)

	return nil
}

func resourceSendgridTemplateUpdate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*sendgrid_client.Client)

	m := &sendgrid_client.Template{}

	if attr, ok := d.GetOk("name"); ok {
		m.Name = attr.(string)
	}
	fmt.Println("Update template")
	if err := client.UpdateTemplate(d.Id(), m); err != nil {
		return fmt.Errorf("error updating Template: %s", err.Error())
	}

	return resourceSendgridTemplateRead(d, meta)
}

func resourceSendgridTemplateDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*sendgrid_client.Client)

	fmt.Println("Delete template")
	if err := client.DeleteTemplate(d.Id()); err != nil {
		return fmt.Errorf("error deleting template: %s", err.Error())
	}

	return nil
}

func resourceSendgridTemplateImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	fmt.Println("Import template")
	if err := resourceSendgridTemplateRead(d, meta); err != nil {
		return nil, fmt.Errorf("error importing template: %s", err.Error())
	}
	return []*schema.ResourceData{d}, nil
}
