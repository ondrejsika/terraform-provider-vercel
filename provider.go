package main

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

type Config struct {
	Token     string
	ApiOrigin string
}

func Provider() *schema.Provider {
	p := &schema.Provider{
		ResourcesMap: map[string]*schema.Resource{
			"zeit_dns":     resourceDns(),
			"zeit_domain":  resourceDoamin(),
			"zeit_project": resourceProject(),
		},
		Schema: map[string]*schema.Schema{
			"token": {
				Type:     schema.TypeString,
				Required: true,
			},
			"api_origin": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "https://api.zeit.co",
			},
		},
	}
	p.ConfigureFunc = func(d *schema.ResourceData) (interface{}, error) {

		config := Config{
			Token:     d.Get("token").(string),
			ApiOrigin: d.Get("api_origin").(string),
		}
		return &config, nil
	}
	return p
}
