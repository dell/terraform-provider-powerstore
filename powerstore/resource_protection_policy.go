package powerstore

import (
	"context"
	"errors"
	"fmt"
	"log"

	"strings"
	client "terraform-provider-powerstore/client"
	"terraform-provider-powerstore/models"

	"github.com/dell/gopowerstore"
	"github.com/hashicorp/terraform-plugin-framework-validators/setvalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// newProtectionPolicyResource returns protection policy new resource instance
func newProtectionPolicyResource() resource.Resource {
	return &resourceProtectionPolicy{}
}

type resourceProtectionPolicy struct {
	client *client.Client
}

// Metadata defines resource interface Metadata method
func (r *resourceProtectionPolicy) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_protectionpolicy"
}

// Schema defines resource interface Schema method
func (r *resourceProtectionPolicy) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{

		MarkdownDescription: "ProtectionPolicy resource",

		Attributes: map[string]schema.Attribute{

			"id": schema.StringAttribute{
				Computed:            true,
				Description:         "Unique identifier of the policy.",
				MarkdownDescription: "Unique identifier of the policy.",
			},

			"name": schema.StringAttribute{
				Required:            true,
				Description:         "The name of the protection policy.",
				MarkdownDescription: "The name of the protection policy.",
				Validators: []validator.String{
					stringvalidator.LengthAtLeast(1),
				},
			},

			"description": schema.StringAttribute{
				Optional:            true,
				Computed:            true,
				Description:         "The description of the protection policy.",
				MarkdownDescription: "The description of the protection policy.",
				Validators: []validator.String{
					stringvalidator.LengthAtLeast(1),
				},
			},

			"snapshot_rule_ids": schema.SetAttribute{
				ElementType:         types.StringType,
				Optional:            true,
				Computed:            true,
				Description:         "List of the snapshot rule IDs that are associated with this policy.",
				MarkdownDescription: "List of the snapshot rule IDs that are associated with this policy.",
				Validators: []validator.Set{

					setvalidator.SizeAtLeast(1),

					setvalidator.ValueStringsAre(
						stringvalidator.LengthAtLeast(1),
					),
				},
			},

			"snapshot_rule_names": schema.SetAttribute{
				ElementType:         types.StringType,
				Optional:            true,
				Computed:            true,
				Description:         "List of the snapshot rule names that are associated with this policy.",
				MarkdownDescription: "List of the snapshot rule names that are associated with this policy.",
				Validators: []validator.Set{

					setvalidator.SizeAtLeast(1),

					setvalidator.ValueStringsAre(
						stringvalidator.LengthAtLeast(1),
					),
				},
			},

			"replication_rule_ids": schema.SetAttribute{
				ElementType:         types.StringType,
				Optional:            true,
				Computed:            true,
				Description:         "List of the replication rule IDs that are associated with this policy.",
				MarkdownDescription: "List of the replication rule IDs that are associated with this policy.",
				Validators: []validator.Set{

					setvalidator.SizeAtLeast(1),

					setvalidator.ValueStringsAre(
						stringvalidator.LengthAtLeast(1),
					),
				},
			},

			"replication_rule_names": schema.SetAttribute{
				ElementType:         types.StringType,
				Optional:            true,
				Computed:            true,
				Description:         "List of the replication rule names that are associated with this policy.",
				MarkdownDescription: "List of the replication rule names that are associated with this policy.",
				Validators: []validator.Set{

					setvalidator.SizeAtLeast(1),

					setvalidator.ValueStringsAre(
						stringvalidator.LengthAtLeast(1),
					),
				},
			},

			// todo, once thse fields are added in gopowerstore
			// then these will be picked up

			// "type": schema.StringAttribute{
			// 	Computed:            true,
			// 	Description:         "The type of the protection policy.",
			// 	MarkdownDescription: "The type of the protection policy.",
			// },
		},
	}
}

// Configure defines resource interface Configure method
func (r *resourceProtectionPolicy) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}

	client, ok := req.ProviderData.(*client.Client)

	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected Resource Configure Type",
			fmt.Sprintf("Expected *http.Client, got: %T. Please report this issue to the provider developers.", req.ProviderData),
		)

		return
	}

	r.client = client
}

// Create defines resource interface Create method
func (r *resourceProtectionPolicy) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {

	var plan models.ProtectionPolicy

	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	protectionPolicyCreate, err := r.planToProtectionPolicyParam(plan)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error creating protection policy",
			err.Error(),
		)
		return
	}

	//Create New ProtectionPolicy
	polCreateResponse, err := r.client.PStoreClient.CreateProtectionPolicy(context.Background(), protectionPolicyCreate)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error creating protection policy",
			"Could not create protection policy, unexpected error: "+err.Error(),
		)
		return
	}

	//Get Protection Policy details using ID retrived above
	polResponse, err := r.client.PStoreClient.GetProtectionPolicy(context.Background(), polCreateResponse.ID)
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

// Delete defines resource interface Delete method
func (r *resourceProtectionPolicy) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	log.Printf("Started with the Delete")

	var state models.ProtectionPolicy
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	//Get Protection Policy ID from state
	protectionPolicyID := state.ID.ValueString()

	//Delete Protection Policy by calling API
	_, err := r.client.PStoreClient.DeleteProtectionPolicy(context.Background(), protectionPolicyID)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error deleting protection policy",
			"Could not delete protectionPolicyID "+protectionPolicyID+": "+err.Error(),
		)
		return
	}

	log.Printf("Done with Delete")
}

// Read defines resource interface Read method
func (r *resourceProtectionPolicy) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	log.Printf("Reading Protection Policy")
	var state models.ProtectionPolicy
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	//Get protection policy details from API and update what is in state from what the API returns
	id := state.ID.ValueString()
	response, err := r.client.PStoreClient.GetProtectionPolicy(context.Background(), id)

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

// Update defines resource interface Update method
func (r *resourceProtectionPolicy) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	log.Printf("Started Update")

	//Get plan values
	var plan models.ProtectionPolicy
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	//Get current state
	var state models.ProtectionPolicy
	diags = req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	protectionPolicyUpdate, err := r.planToProtectionPolicyParam(plan)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error updating protection policy",
			err.Error(),
		)
		return
	}

	//Get Protection Policy ID from state
	protectionPolicyID := state.ID.ValueString()

	//Update Protection Policy by calling API
	_, err = r.client.PStoreClient.ModifyProtectionPolicy(context.Background(), protectionPolicyUpdate, protectionPolicyID)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error updating protection policy",
			"Could not update protectionPolicyID "+protectionPolicyID+": "+err.Error(),
		)
		return
	}

	//Get Protection Policy details
	getRes, err := r.client.PStoreClient.GetProtectionPolicy(context.Background(), protectionPolicyID)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error getting protection policy after update",
			"Could not get protection policy, unexpected error: "+err.Error(),
		)
		return
	}

	r.updatePolicyState(&state, getRes, &plan)

	diags = resp.State.Set(ctx, state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	log.Printf("Successfully done with Update")
}

func (r resourceProtectionPolicy) planToProtectionPolicyParam(plan models.ProtectionPolicy) (*gopowerstore.ProtectionPolicyCreate, error) {
	valid, err := r.fetchByName(&plan)
	if !valid {
		return nil, err
	}

	var replicationRuleIds []string
	for _, replicationRule := range plan.ReplicationRuleIDs.Elements() {
		replicationRuleIds = append(replicationRuleIds, strings.Trim(replicationRule.String(), "\""))
	}

	var snapshotRuleIds []string
	for _, snapshotRule := range plan.SnapshotRuleIDs.Elements() {
		snapshotRuleIds = append(snapshotRuleIds, strings.Trim(snapshotRule.String(), "\""))
	}

	protectionPolicyCreate := &gopowerstore.ProtectionPolicyCreate{
		Name:               plan.Name.ValueString(),
		Description:        plan.Description.ValueString(),
		ReplicationRuleIds: replicationRuleIds,
		SnapshotRuleIds:    snapshotRuleIds,
	}
	return protectionPolicyCreate, nil
}

func (r resourceProtectionPolicy) updatePolicyState(polState *models.ProtectionPolicy, polResponse gopowerstore.ProtectionPolicy, polPlan *models.ProtectionPolicy) {
	// Update value from Protection Policy Response to State
	polState.ID = types.StringValue(polResponse.ID)
	polState.Name = types.StringValue(polResponse.Name)
	polState.Description = types.StringValue(polResponse.Description)

	//Update ReplicationRuleIDs value from Response to State
	var replicationRuleIds []string
	for _, replicationRule := range polResponse.ReplicationRules {
		replicationRuleIds = append(replicationRuleIds, replicationRule.ID)
	}
	replicationIDList := []attr.Value{}
	for _, replicationRuleID := range replicationRuleIds {
		replicationIDList = append(replicationIDList, types.StringValue(string(replicationRuleID)))
	}
	polState.ReplicationRuleIDs, _ = types.SetValue(types.StringType, replicationIDList)

	//Update ReplicationRuleNames value from Plan to State
	var replicationRuleNames []string
	for _, replicationRuleName := range polPlan.ReplicationRuleNames.Elements() {
		replicationRuleNames = append(replicationRuleNames, strings.Trim(replicationRuleName.String(), "\""))
	}
	replicationNameList := []attr.Value{}
	for _, replicationRuleName := range replicationRuleNames {
		replicationNameList = append(replicationNameList, types.StringValue(string(replicationRuleName)))
	}

	polState.ReplicationRuleNames, _ = types.SetValue(types.StringType, replicationNameList)

	//Update SnapshotRuleIDs value from Response to State
	var snapshotRuleIds []string
	for _, snapshotRule := range polResponse.SnapshotRules {
		snapshotRuleIds = append(snapshotRuleIds, snapshotRule.ID)
	}
	snapshotIDList := []attr.Value{}
	for _, snapshotRuleID := range snapshotRuleIds {
		snapshotIDList = append(snapshotIDList, types.StringValue(string(snapshotRuleID)))
	}
	polState.SnapshotRuleIDs, _ = types.SetValue(types.StringType, snapshotIDList)

	//Update SnapshotRuleNames value from Plan to State
	var snapshotRuleNames []string
	for _, snapshotRuleName := range polPlan.SnapshotRuleNames.Elements() {
		snapshotRuleNames = append(snapshotRuleNames, strings.Trim(snapshotRuleName.String(), "\""))
	}
	snapshotNameList := []attr.Value{}
	for _, snapshotRuleName := range snapshotRuleNames {
		snapshotNameList = append(snapshotNameList, types.StringValue(string(snapshotRuleName)))
	}
	polState.SnapshotRuleNames, _ = types.SetValue(types.StringType, snapshotNameList)
}

func (r resourceProtectionPolicy) fetchByName(plan *models.ProtectionPolicy) (valid bool, err error) {
	var snapshotRuleIds []string
	if len(plan.SnapshotRuleIDs.Elements()) != 0 && len(plan.SnapshotRuleNames.Elements()) != 0 {
		return false, errors.New("either of snapshot rule id or snapshot rule name should be present")
	} else if len(plan.SnapshotRuleNames.Elements()) != 0 {
		for _, snapshotRuleName := range plan.SnapshotRuleNames.Elements() {
			snapshotRule, _ := r.client.PStoreClient.GetSnapshotRuleByName(context.Background(), strings.Trim(snapshotRuleName.String(), "\""))
			snapshotRuleIds = append(snapshotRuleIds, strings.Trim(snapshotRule.ID, "\""))

			snapshotList := []attr.Value{}
			for i := 0; i < len(snapshotRuleIds); i++ {
				snapshotList = append(snapshotList, types.StringValue(string(snapshotRuleIds[i])))
			}

			plan.SnapshotRuleIDs, _ = types.SetValue(types.StringType, snapshotList)
		}
	}

	var replicationRuleIds []string
	if len(plan.ReplicationRuleIDs.Elements()) != 0 && len(plan.ReplicationRuleNames.Elements()) != 0 {
		return false, errors.New("either of replication rule id or replication rule name should be present")
	} else if len(plan.ReplicationRuleNames.Elements()) != 0 {
		for _, replicationRuleName := range plan.ReplicationRuleNames.Elements() {
			replicationRule, _ := r.client.PStoreClient.GetReplicationRuleByName(context.Background(), strings.Trim(replicationRuleName.String(), "\""))
			replicationRuleIds = append(replicationRuleIds, strings.Trim(replicationRule.ID, "\""))
		}

		replicationList := []attr.Value{}
		for i := 0; i < len(replicationRuleIds); i++ {
			replicationList = append(replicationList, types.StringValue(string(replicationRuleIds[i])))
		}

		plan.ReplicationRuleIDs, _ = types.SetValue(types.StringType, replicationList)
	}

	return true, nil
}
