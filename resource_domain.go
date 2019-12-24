package main

import (
	"github.com/go-resty/resty/v2"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceDomainCreate(d *schema.ResourceData, m interface{}) error {
	domain := d.Get("domain").(string)
	expectedPrice := d.Get("expected_price").(int)

	d.SetId(domain)

	client := resty.New()
	_, err := client.R().SetAuthToken(m.(*Config).Token).SetBody(map[string]interface{}{"name": domain, "expectedPrice": expectedPrice}).Post(m.(*Config).ApiOrigin + "/v2/domains/buy")

	if err != nil {
		return err
	}

	return nil
}

func resourceDomainRead(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceDomainUpdate(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceDomainDelete(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceDoamin() *schema.Resource {
	return &schema.Resource{
		Create: resourceDomainCreate,
		Read:   resourceDomainRead,
		Update: resourceDomainUpdate,
		Delete: resourceDomainDelete,

		Schema: map[string]*schema.Schema{
			"domain": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"expected_price": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
		},
	}
}
