package main

import (
	"encoding/json"

	"github.com/go-resty/resty/v2"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceProjectCreate(d *schema.ResourceData, m interface{}) error {
	name := d.Get("name").(string)

	client := resty.New()
	rawResp, err := client.R().SetAuthToken(m.(*Config).Token).SetBody(map[string]interface{}{"name": name}).Post(m.(*Config).ApiOrigin + "/v1/projects/ensure-project")

	if err != nil {
		return err
	}

	var response map[string]interface{}
	json.Unmarshal([]byte(rawResp.Body()), &response)

	d.SetId(response["id"].(string))
	return nil
}

func resourceProjectRead(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceProjectUpdate(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceProjectDelete(d *schema.ResourceData, m interface{}) error {
	name := d.Get("name").(string)

	client := resty.New()
	_, err := client.R().SetAuthToken(m.(*Config).Token).Delete(m.(*Config).ApiOrigin + "/v1/projects/" + name)

	if err != nil {
		return err
	}

	return nil
}

func resourceProject() *schema.Resource {
	return &schema.Resource{
		Create: resourceProjectCreate,
		Read:   resourceProjectRead,
		Update: resourceProjectUpdate,
		Delete: resourceProjectDelete,

		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}
