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
			"vercel_dns":     resourceDns(),
			"vercel_domain":  resourceDoamin(),
			"vercel_project": resourceProject(),
		},
		Schema: map[string]*schema.Schema{
			"token": {
				Type:     schema.TypeString,
				Required: true,
			},
			"api_origin": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "https://api.vercel.com",
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
