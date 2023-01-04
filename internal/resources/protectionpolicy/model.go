package protectionpolicy

import (
	"context"
	"fmt"
	"log"
	"terraform-provider-powerstore/internal/powerstore"
	"terraform-provider-powerstore/internal/utils"

	"github.com/dell/gopowerstore"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
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

// todo
// client is need in this method
// because gopowerstore implementation does not include names in response
// all thie spaghettic code down below will be gone, once implementation includes names in response
func (m *model) saveSeverResponse(response gopowerstore.ProtectionPolicy, client *powerstore.Client) (diag1 diag.Diagnostics) {

	log.Printf("mayank response %+v", response)

	m.ID = types.StringValue(response.ID)
	m.Name = types.StringValue(response.Name)
	m.Description = types.StringValue(response.Description)

	// todo
	// kina duplicate code for replicatin and snapshot rules
	// duplicacy can be rmeoved by introducing interfaces or reflect or generics as both are differnt object 

	// if names are provided, then only names will be populated, else id
	planNamesMap := map[string]struct{}{}

	for _, name := range utils.SetValuesToStringList(m.SnapshotRuleNames, "") {
		planNamesMap[name] = struct{}{}
	}

	ids := []string{}
	names := []string{}
	for _, rule := range response.SnapshotRules {

		// todo should be present in response
		res, err := client.PStoreClient.GetSnapshotRule(context.Background(), rule.ID)
		if err != nil {
			diag1.AddError(
				"unable to fetch name for snapshotrule ID",
				fmt.Sprintf("unable to fetch name for snapshotrule ID: %s, error: %s", rule.ID, err.Error()),
			)
			return
		}

		if _, ok := planNamesMap[res.Name]; ok {
			names = append(names, res.Name)
		} else {
			ids = append(ids, rule.ID)
		}
	}

	var diag2 diag.Diagnostics

	m.SnapshotRuleIDs, diag2 = utils.StringListToSetValues(ids, "")
	diag1.Append(diag2...)
	m.SnapshotRuleNames, diag2 = utils.StringListToSetValues(names, "")
	diag1.Append(diag2...)

	planNamesMap = map[string]struct{}{}

	for _, name := range utils.SetValuesToStringList(m.ReplicationRuleNames, "") {
		planNamesMap[name] = struct{}{}
	}

	ids = []string{}
	names = []string{}
	for _, rule := range response.ReplicationRules {

		res, err := client.PStoreClient.GetReplicationRule(context.Background(), rule.ID)
		if err != nil {
			diag1.AddError(
				"unable to fetch name for replication ID",
				fmt.Sprintf("unable to fetch name for replication ID: %s, error: %s", rule.ID, err.Error()),
			)
			return
		}

		if _, ok := planNamesMap[res.Name]; ok {
			names = append(names, res.Name)
		} else {
			ids = append(ids, rule.ID)
		}
	}

	m.ReplicationRuleIDs, diag2 = utils.StringListToSetValues(ids, "")
	diag1.Append(diag2...)
	m.ReplicationRuleNames, diag2 = utils.StringListToSetValues(names, "")
	diag1.Append(diag2...)

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
