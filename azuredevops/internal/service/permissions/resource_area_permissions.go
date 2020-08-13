package permissions

import (
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	"github.com/microsoft/azure-devops-go-api/azuredevops/workitemtracking"
	"github.com/terraform-providers/terraform-provider-azuredevops/azuredevops/internal/client"
	securityhelper "github.com/terraform-providers/terraform-provider-azuredevops/azuredevops/internal/service/permissions/utils"
)

func ResourceAreaPermissions() *schema.Resource {
	return &schema.Resource{
		Create: resourceAreaPermissionsCreateOrUpdate,
		Read:   resourceAreaPermissionsRead,
		Update: resourceAreaPermissionsCreateOrUpdate,
		Delete: resourceAreaPermissionsDelete,
		Schema: securityhelper.CreatePermissionResourceSchema(map[string]*schema.Schema{
			"project_id": {
				Type:         schema.TypeString,
				ValidateFunc: validation.IsUUID,
				Required:     true,
				ForceNew:     true,
			},
			"path": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringIsNotWhiteSpace,
				ForceNew:     true,
				Optional:     true,
			},
		}),
	}
}

func resourceAreaPermissionsCreateOrUpdate(d *schema.ResourceData, m interface{}) error {
	clients := m.(*client.AggregatedClient)

	sn, aclToken, err := securityhelper.InitializeSecurityNamespaceAndToken(d, clients, securityhelper.SecurityNamespaceIDValues.CSS, createAreaToken)
	if err != nil {
		return err
	}

	if err = securityhelper.SetPrincipalPermissions(d, sn, aclToken, nil, false); err != nil {
		return err
	}

	return resourceAreaPermissionsRead(d, m)
}

func resourceAreaPermissionsRead(d *schema.ResourceData, m interface{}) error {
	clients := m.(*client.AggregatedClient)

	sn, aclToken, err := securityhelper.InitializeSecurityNamespaceAndToken(d, clients, securityhelper.SecurityNamespaceIDValues.CSS, createAreaToken)
	if err != nil {
		return err
	}

	principalPermissions, err := securityhelper.GetPrincipalPermissions(d, sn, aclToken)
	if err != nil {
		return err
	}
	if principalPermissions == nil {
		d.SetId("")
		log.Printf("[INFO] Permissions for ACL token %q not found. Removing from state", *aclToken)
		return nil
	}

	d.Set("permissions", principalPermissions.Permissions)
	return nil
}

func resourceAreaPermissionsDelete(d *schema.ResourceData, m interface{}) error {
	clients := m.(*client.AggregatedClient)

	sn, aclToken, err := securityhelper.InitializeSecurityNamespaceAndToken(d, clients, securityhelper.SecurityNamespaceIDValues.CSS, createAreaToken)
	if err != nil {
		return err
	}

	if err := securityhelper.SetPrincipalPermissions(d, sn, aclToken, &securityhelper.PermissionTypeValues.NotSet, true); err != nil {
		return err
	}

	d.SetId("")
	return nil
}

func createAreaToken(d *schema.ResourceData, clients *client.AggregatedClient) (*string, error) {
	projectID := d.Get("project_id").(string)
	path := d.Get("path").(string)
	aclToken, err := securityhelper.CreateClassificationNodeSecurityToken(clients.Ctx, clients.WorkItemTrackingClient, workitemtracking.TreeStructureGroupValues.Areas, projectID, path)
	if err != nil {
		return nil, err
	}
	return &aclToken, nil
}
