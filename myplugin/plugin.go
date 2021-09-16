package custom

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

type Config struct {
	ApiKey string
}

func resourcePluginSourceRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	d.SetId("some_id_for_set_values")

	// here we can do what we want and set for output
	d.Set("data_source_schema_variable", config.ApiKey)
	return nil
}

func dataSourcePlugin() *schema.Resource {
	return &schema.Resource{
		Read: resourcePluginSourceRead,
		Schema: map[string]*schema.Schema{
			"data_source_schema_variable": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func providerConfig(d *schema.ResourceData) (interface{}, error) {
	c := &Config{
		ApiKey: d.Get("api_key").(string),
	}
	return c, nil
}

func Provider() *schema.Provider {
	return &schema.Provider{
		DataSourcesMap: map[string]*schema.Resource{
			"plugin_data": dataSourcePlugin(),
		},
		Schema: map[string]*schema.Schema{
			"api_key": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"provider_schema_variable": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
		ConfigureFunc: providerConfig,
	}
}
