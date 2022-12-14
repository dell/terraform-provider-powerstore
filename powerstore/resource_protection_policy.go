package powerstore

import (
	"context"
	"log"

	"strings"
	"terraform-provider-powerstore/models"

	"github.com/dell/gopowerstore"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type resourceProtectionPolicyType struct{}

// GetSchema returns the schema for this resource.
func (r resourceProtectionPolicyType) GetSchema(_ context.Context) (tfsdk.Schema, diag.Diagnostics) {
	return tfsdk.Schema{
		Attributes: map[string]tfsdk.Attribute{
			"id": {
				Type:                types.StringType,
				Computed:            true,
				Description:         "The ID of the protection policy,",
				MarkdownDescription: "The ID of the protection policy",
			},
			"name": {
				Type:                types.StringType,
				Required:            true,
				Description:         "The name of the protectio policy.",
				MarkdownDescription: "The name of the protection policy.",
			},
			"description": {
				Type:                types.StringType,
				Computed:            true,
				Optional:            true,
				Description:         "The description of the protection policy.",
				MarkdownDescription: "The description of the protection policy.",
			},
			"type": {
				Type:                types.StringType,
				Computed:            true,
				Optional:            true,
				Description:         "The type of the protection policy.",
				MarkdownDescription: "The type of the protection policy.",
			},
			"managed_by": {
				Type:                types.StringType,
				Computed:            true,
				Optional:            true,
				Description:         "Entity that owns and manages this instance.",
				MarkdownDescription: "Entity that owns and manages this instance.",
			},
			"managed_by_id": {
				Type:                types.StringType,
				Computed:            true,
				Optional:            true,
				Description:         "ID of the managing entity.",
				MarkdownDescription: "ID of the managing entity.",
			},
			"is_read_only": {
				Type:                types.BoolType,
				Computed:            true,
				Optional:            true,
				Description:         "The is_read_only of the protection policy.",
				MarkdownDescription: "The is_read_only of the protection policy.",
			},
			"is_replica": {
				Type:                types.BoolType,
				Computed:            true,
				Optional:            true,
				Description:         "The is_replica of the protection policy.",
				MarkdownDescription: "The is_replica of the protection policy.",
			},
			"snapshot_rule_ids": {
				Type:                types.ListType{ElemType: types.StringType},
				Computed:            true,
				Optional:            true,
				Description:         "The snapshot_rule_ids of the protection policy.",
				MarkdownDescription: "The snapshot_rule_ids of the protection policy.",
			},
			"replication_rule_ids": {
				Type:                types.ListType{ElemType: types.StringType},
				Computed:            true,
				Optional:            true,
				Description:         "The replication_rule_ids of the protection policy.",
				MarkdownDescription: "The replication_rule_ids of the protection policy.",
			},
		},
	}, nil
}

// NewResource is a wrapper around provider
func (r resourceProtectionPolicyType) NewResource(_ context.Context, p tfsdk.Provider) (tfsdk.Resource, diag.Diagnostics) {
	return resourceProtectionPolicy{
		p: *(p.(*Pstoreprovider)),
	}, nil
}

type resourceProtectionPolicy struct {
	p Pstoreprovider
}

func (r resourceProtectionPolicy) Create(ctx context.Context, req tfsdk.CreateResourceRequest, resp *tfsdk.CreateResourceResponse) {
	if !r.p.configured {
		resp.Diagnostics.AddError(
			"Provider not configured",
			"The provider hasn't been configured before apply, likely because it depends on an unknown value from another resource. This leads to weird stuff happening, so we'd prefer if you didn't do that. Thanks!",
		)
		return
	}
	var plan models.ProtectionPolicy

	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	protectionPolicyCreate := r.planToProtectionPolicyParam(plan)

	//Create New ProtectionPolicy
	polCreateResponse, err := r.p.client.PStoreClient.CreateProtectionPolicy(context.Background(), protectionPolicyCreate)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error creating protection policy",
			"Could not create protection policy, unexpected error: "+err.Error(),
		)
		return
	}

	//Get Protection Policy details using ID retrived above
	polResponse, err := r.p.client.PStoreClient.GetProtectionPolicy(context.Background(), polCreateResponse.ID)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error getting protection policy after creation",
			"Could not get protection policy, unexpected error: "+err.Error(),
		)
		return
	}

	result := models.ProtectionPolicy{}
	r.updatePolicyState(&result, polResponse, &plan)

	diags = resp.State.Set(ctx, result)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	log.Printf("Done with Create")
}

// Delete deletes the protection policy
func (r resourceProtectionPolicy) Delete(ctx context.Context, req tfsdk.DeleteResourceRequest, resp *tfsdk.DeleteResourceResponse) {
	log.Printf("Started with the Delete")

	var state models.ProtectionPolicy
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	//Get Protection Policy ID from state
	protectionPolicyID := state.ID.Value

	//Delete Protection Policy by calling API
	_, err := r.p.client.PStoreClient.DeleteProtectionPolicy(context.Background(), protectionPolicyID)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error deleting protection policy",
			"Could not delete protectionPolicyID "+protectionPolicyID+": "+err.Error(),
		)
		return
	}

	log.Printf("Done with Delete")
}

// Read fetch info about asked protection policy
func (r resourceProtectionPolicy) Read(ctx context.Context, req tfsdk.ReadResourceRequest, resp *tfsdk.ReadResourceResponse) {
	log.Printf("Reading Protection Policy")
	var state models.ProtectionPolicy
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	//Get protection policy details from API and update what is in state from what the API returns
	id := state.ID.Value
	response, err := r.p.client.PStoreClient.GetProtectionPolicy(context.Background(), id)

	if err != nil {
		resp.Diagnostics.AddError(
			"Error reading protection policy",
			"Could not read protection policy with error "+id+": "+err.Error(),
		)
		return
	}

	r.updatePolicyState(&state, response, &state)

	//Set state
	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	log.Printf("Done with Read")
}

// Update updates protection policy
func (r resourceProtectionPolicy) Update(ctx context.Context, req tfsdk.UpdateResourceRequest, resp *tfsdk.UpdateResourceResponse) {
}

func (r resourceProtectionPolicy) planToProtectionPolicyParam(plan models.ProtectionPolicy) *gopowerstore.ProtectionPolicyCreate {

	var replicationRuleIds []string
	for _, replicationRule := range plan.ReplicationRuleIDs.Elems {
		replicationRuleIds = append(replicationRuleIds, strings.Trim(replicationRule.String(), "\""))
	}

	var snapshotRuleIds []string
	for _, snapshotRule := range plan.SnapshotRuleIDs.Elems {
		snapshotRuleIds = append(snapshotRuleIds, strings.Trim(snapshotRule.String(), "\""))
	}

	protectionPolicyCreate := &gopowerstore.ProtectionPolicyCreate{
		Name:               plan.Name.Value,
		Description:        plan.Description.Value,
		ReplicationRuleIds: replicationRuleIds,
		SnapshotRuleIds:    snapshotRuleIds,
	}
	return protectionPolicyCreate
}

func (r resourceProtectionPolicy) updatePolicyState(polState *models.ProtectionPolicy, polResponse gopowerstore.ProtectionPolicy, polPlan *models.ProtectionPolicy) {
	// Update value from Protection Policy Response to State
	polState.ID.Value = polResponse.ID
	polState.Name.Value = polResponse.Name

	var replicationRuleIds []string
	for _, replicationRule := range polResponse.ReplicationRules {
		replicationRuleIds = append(replicationRuleIds, replicationRule.ID)
	}
	replicationList := []attr.Value{}
	for i := 0; i < len(replicationRuleIds); i++ {
		replicationList = append(replicationList, types.String{Value: string(replicationRuleIds[i])})
	}
	polState.ReplicationRuleIDs = types.List{
		ElemType: types.StringType,
		Elems:    replicationList,
	}

	var snapshotRuleIds []string
	for _, snapshotRule := range polResponse.SnapshotRules {
		snapshotRuleIds = append(snapshotRuleIds, snapshotRule.ID)
	}
	snapshotList := []attr.Value{}
	for i := 0; i < len(snapshotRuleIds); i++ {
		snapshotList = append(snapshotList, types.String{Value: string(snapshotRuleIds[i])})
	}
	polState.SnapshotRuleIDs = types.List{
		ElemType: types.StringType,
		Elems:    snapshotList,
	}
}
