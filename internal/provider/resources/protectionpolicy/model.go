package protectionpolicy

import (
	"context"
	"fmt"
	"log"
	"reflect"
	"terraform-provider-powerstore/internal/powerstore"
	"terraform-provider-powerstore/internal/utils"

	"github.com/dell/gopowerstore"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

// Model specifies protectionpolicy properties
type model struct {
	ID          types.String `tfsdk:"id"`
	Name        types.String `tfsdk:"name"`
	Description types.String `tfsdk:"description"`

	SnapshotRuleIDs    types.Set `tfsdk:"snapshot_rule_ids"`
	ReplicationRuleIDs types.Set `tfsdk:"replication_rule_ids"`

	SnapshotRuleNames    types.Set `tfsdk:"snapshot_rule_names"`
	ReplicationRuleNames types.Set `tfsdk:"replication_rule_names"`

	// only to get by server
	// todo
	// currently following are not implemented in gopowerstore
	// Type        types.String `tfsdk:"type"`
	// ManagedBy   types.String `tfsdk:"managed_by"`
	// ManagedByID types.String `tfsdk:"managed_by_id"`
	// IsReadOnly  types.Bool   `tfsdk:"is_read_only"`
	// IsReplica   types.Bool   `tfsdk:"is_replica"`
}

func (m *model) saveSeverResponse(response gopowerstore.ProtectionPolicy, client *powerstore.Client) (diag1 diag.Diagnostics) {

	log.Printf("mayank response %+v", response)

	m.ID = types.StringValue(response.ID)
	m.Name = types.StringValue(response.Name)
	m.Description = types.StringValue(response.Description)

	diag1.Append(m.populateIdsAndNames(true, response, client)...)
	diag1.Append(m.populateIdsAndNames(false, response, client)...)

	return
}

func (m *model) serverRequest() *gopowerstore.ProtectionPolicyCreate {

	return &gopowerstore.ProtectionPolicyCreate{
		Name:        m.Name.ValueString(),
		Description: m.Description.ValueString(),
		SnapshotRuleIds: append(
			utils.SetValuesToStringList(m.SnapshotRuleIDs, ""),
			utils.SetValuesToStringList(m.SnapshotRuleNames, "name:")...,
		),
		ReplicationRuleIds: append(
			utils.SetValuesToStringList(m.ReplicationRuleIDs, ""),
			utils.SetValuesToStringList(m.ReplicationRuleNames, "name:")...,
		),
	}
}

// todo
// client is need in this method
// because gopowerstore implementation does not include names in response
// all this code down below will be gone, once implementation includes names in response
func (m *model) populateIdsAndNames(snapShotRules bool, response gopowerstore.ProtectionPolicy, client *powerstore.Client) (diag1 diag.Diagnostics) {

	var userDefinedNames basetypes.SetValue
	var rulesValue reflect.Value

	if snapShotRules {
		userDefinedNames = m.SnapshotRuleNames
		rulesValue = reflect.ValueOf(response.SnapshotRules)
	} else {
		userDefinedNames = m.ReplicationRuleNames
		rulesValue = reflect.ValueOf(response.ReplicationRules)
	}

	// if names are provided, then only names will be populated, else id
	// cannot be done in reverse, like for ids
	// as response always have ids
	// so in case we don't have plan like in import or read, then we won't be able to populate names correctly
	planNamesMap := map[string]struct{}{}

	for _, name := range utils.SetValuesToStringList(userDefinedNames, "") {
		planNamesMap[name] = struct{}{}
	}

	ids := []string{}
	names := []string{}
	for i := 0; i < rulesValue.Len(); i++ {

		ruleID := rulesValue.Index(i).FieldByName("ID").Interface().(string)

		var name string

		if snapShotRules {
			resp, err := client.PStoreClient.GetSnapshotRule(context.Background(), ruleID)
			if err != nil {
				diag1.AddError(
					"unable to fetch name for snapshot rule via ID",
					fmt.Sprintf("unable to fetch name for snapshot rule ID: %s, error: %s", ruleID, err.Error()),
				)
				return
			}
			name = resp.Name
		} else {
			resp, err := client.PStoreClient.GetReplicationRule(context.Background(), ruleID)
			if err != nil {
				diag1.AddError(
					"unable to fetch name for replication rule via ID",
					fmt.Sprintf("unable to fetch name for replication rule ID: %s, error: %s", ruleID, err.Error()),
				)
				return
			}
			name = resp.Name
		}

		if _, ok := planNamesMap[name]; ok {
			names = append(names, name)
		} else {
			ids = append(ids, ruleID)
		}
	}

	var diag2 diag.Diagnostics

	if snapShotRules {
		m.SnapshotRuleIDs, diag2 = utils.StringListToSetValues(ids, "")
		diag1.Append(diag2...)
		m.SnapshotRuleNames, diag2 = utils.StringListToSetValues(names, "")
		diag1.Append(diag2...)
	} else {
		m.ReplicationRuleIDs, diag2 = utils.StringListToSetValues(ids, "")
		diag1.Append(diag2...)
		m.ReplicationRuleNames, diag2 = utils.StringListToSetValues(names, "")
		diag1.Append(diag2...)
	}

	return
}
