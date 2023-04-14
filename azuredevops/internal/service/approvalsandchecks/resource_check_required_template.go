package approvalsandchecks

import (
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
	"github.com/microsoft/azure-devops-go-api/azuredevops/pipelineschecks"
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
			Elem: getRequiredTemplateSchema(),
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
			Description:      "type of repository host",
		},
		"repository": {
			Type:        schema.TypeString,
			Description: "repository in project/repository-name",
		},
		"ref": {
			Type:        schema.TypeString,
			Description: "git ref, for example refs/tags/release/v1",
		},
		"path_to_required_yaml_template": {
			Type:        schema.TypeString,
			Description: "path to the yaml file",
		},
	}
}


// as said here https://stackoverflow.com/questions/61471634/add-remove-pipeline-checks-using-rest-api the doc do not exists and only sniffint http requests is available to discover body schema
func expandRequiredTemplateCheck(d *schema.ResourceData) (*pipelineschecks.CheckConfiguration, string, error) {
	// inputs := map[string]interface{}{
	// 	"allowedBranches":          d.Get("allowed_branches").(string),
	// 	"ensureProtectionOfBranch": strconv.FormatBool(d.Get("verify_branch_protection").(bool)),
	// 	"allowUnknownStatusBranch": strconv.FormatBool(d.Get("ignore_unknown_protection_status").(bool)),
	// }

  re
	inputs := map[string]interface{}{
		"required_yaml_templates": []map[string]interface{}{
map[string]interface{}{

      }
    },
	}

	return doBaseExpansion(d, inputs, evaluateBranchProtectionDef)
}
