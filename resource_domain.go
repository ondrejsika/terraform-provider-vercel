package main

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/ondrejsika/zeit-go"
)

func resourceDomainCreate(d *schema.ResourceData, m interface{}) error {
	domain := d.Get("domain").(string)
	expectedPrice := d.Get("expected_price").(int)
	removeDomainOnDestroy := d.Get("remove_domain_on_destroy").(bool)

	d.SetId(domain)
	d.Set("remove_domain_on_destroy", removeDomainOnDestroy)

	rawResp, err := zeit.NewOrigin(m.(*Config).Token, m.(*Config).ApiOrigin).BuyDomain(domain, expectedPrice)

	if err != nil {
		return fmt.Errorf("%s", err)
	}

	var response map[string]map[string]interface{}
	json.Unmarshal([]byte(rawResp.Body()), &response)

	if response["error"]["code"] == "not_available" {
		return fmt.Errorf("Domain %s is not available", domain)
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
	domain := d.Get("domain").(string)
	removeDomainOnDestroy := d.Get("remove_domain_on_destroy").(bool)

	if removeDomainOnDestroy {
		_, err := zeit.NewOrigin(m.(*Config).Token, m.(*Config).ApiOrigin).RemoveDomain(domain)

		if err != nil {
			return fmt.Errorf("%s", err)
		}
	}

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
			"remove_domain_on_destroy": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
		},
	}
}
