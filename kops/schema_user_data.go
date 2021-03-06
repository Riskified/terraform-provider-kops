package kops

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func schemaUserData() *schema.Schema {
	return &schema.Schema{
		Type:     schema.TypeList,
		Optional: true,
		Computed: true,
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				"name":    schemaStringRequired(),
				"type":    schemaStringRequired(),
				"content": schemaStringRequired(),
			},
		},
	}
}
