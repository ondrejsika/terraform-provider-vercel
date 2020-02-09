package main

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/ondrejsika/zeit-go"
)

func resourceDomainCreate(d *schema.ResourceData, m interface{}) error {
	domain := d.Get("domain").(string)
	expectedPrice := d.Get("expected_price").(int)

	d.SetId(domain)

	_, err := zeit.NewOrigin(m.(*Config).Token, m.(*Config).ApiOrigin).BuyDomain(domain, expectedPrice)

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
