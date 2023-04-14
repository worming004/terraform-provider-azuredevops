package approvalsandchecks

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
	"github.com/microsoft/terraform-provider-azuredevops/azuredevops/internal/utils/suppress"
)

func ResourceCheckRequiredTemplate() *schema.Resource {
	r := genBaseCheckResource(nil, nil)

	r.Schema["required_yaml_templates"] = &schema.Schema{
		Type:     schema.TypeList,
		Optional: false,
		MinItems: 1,
		Elem: &schema.Schema{
			Type: schema.TypeSet,
			Elem: map[string]*schema.Schema{},
		},
	}

	return r
}

func getRequiredTemplateSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"repository_host": {
			Type:             schema.TypeString,
			DiffSuppressFunc: suppress.CaseDifference,
			ValidateFunc:     validation.StringInSlice([]string{"Azure Repos", "GitHub", "Bitbucket"}, false),
			Default:          "Azure Repos",
			Description:      "repository host",
		},
	}
}
