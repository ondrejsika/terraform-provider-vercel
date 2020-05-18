package main

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/ondrejsika/vercel-go"
)

func resourceDnsCreate(d *schema.ResourceData, m interface{}) error {
	domain := d.Get("domain").(string)
	name := d.Get("name").(string)
	value := d.Get("value").(string)
	type_ := d.Get("type").(string)

	response, err := vercel.NewOrigin(m.(*Config).Token, m.(*Config).ApiOrigin).CreateRecord(domain, type_, name, value)

	if err != nil {
		return err
	}

	d.SetId(response.UID)
	return nil
}

func resourceDnsRead(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceDnsDelete(d *schema.ResourceData, m interface{}) error {
	uid := d.Id()
	domain := d.Get("domain").(string)

	err := vercel.NewOrigin(m.(*Config).Token, m.(*Config).ApiOrigin).DeleteRecord(domain, uid)

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
