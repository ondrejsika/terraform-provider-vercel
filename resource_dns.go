package main

import (
	"encoding/json"

	"github.com/go-resty/resty/v2"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceDnsCreate(d *schema.ResourceData, m interface{}) error {
	domain := d.Get("domain").(string)
	name := d.Get("name").(string)
	value := d.Get("value").(string)
	type_ := d.Get("type").(string)

	client := resty.New()
	rawResp, err := client.R().SetAuthToken(m.(*Config).Token).SetBody(map[string]interface{}{"name": name, "value": value, "type": type_}).Post(m.(*Config).ApiOrigin + "/v2/domains/" + domain + "/records")

	if err != nil {
		return err
	}

	var response map[string]interface{}
	json.Unmarshal([]byte(rawResp.Body()), &response)

	d.SetId(response["uid"].(string))
	return nil
}

func resourceDnsRead(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceDnsDelete(d *schema.ResourceData, m interface{}) error {
	uid := d.Id()
	domain := d.Get("domain").(string)

	client := resty.New()
	_, err := client.R().SetAuthToken(m.(*Config).Token).Delete(m.(*Config).ApiOrigin + "/v2/domains/" + domain + "/records/" + uid)

	if err != nil {
		return err
	}
	return nil
}

func resourceDns() *schema.Resource {
	return &schema.Resource{
		Create: resourceDnsCreate,
		Read:   resourceDnsRead,
		Delete: resourceDnsDelete,

		Schema: map[string]*schema.Schema{
			"domain": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"type": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"value": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
		},
	}
}
