/*
Copyright (c) 2026 Dell Inc., or its subsidiaries. All Rights Reserved.

Licensed under the Mozilla Public License Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://mozilla.org/MPL/2.0/


Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package powerstore

import (
	"context"
	"fmt"
	"net/url"
	"time"

	"terraform-provider-powerstore/client"
	"terraform-provider-powerstore/clientgen"
	"terraform-provider-powerstore/models"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var (
	_ datasource.DataSource              = &recycleBinDataSource{}
	_ datasource.DataSourceWithConfigure = &recycleBinDataSource{}
)

// newRecycleBinDataSource returns the recycle bin data source object
func newRecycleBinDataSource() datasource.DataSource {
	return &recycleBinDataSource{}
}

type recycleBinDataSource struct {
	client *clientgen.APIClient
}

func (d *recycleBinDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_recycle_bin"
}

func (d *recycleBinDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description:         "This datasource is used to query the recycle bin contents from a PowerStore array. The information fetched from this datasource can be used for recovering or permanently deleting items.",
		MarkdownDescription: "This datasource is used to query the recycle bin contents from a PowerStore array. The information fetched from this datasource can be used for recovering or permanently deleting items.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description:         "Unique identifier of a specific recycle bin item. Conflicts with `resource_type` and `filter_expression`.",
				MarkdownDescription: "Unique identifier of a specific recycle bin item. Conflicts with `resource_type` and `filter_expression`.",
				Optional:            true,
				Computed:            true,
				Validators: []validator.String{
					stringvalidator.ConflictsWith(path.MatchRoot("resource_type")),
					stringvalidator.ConflictsWith(path.MatchRoot("filter_expression")),
					stringvalidator.LengthAtLeast(1),
				},
			},
			"resource_type": schema.StringAttribute{
				Description:         "Filter recycle bin items by resource type. Valid values are `volume` and `volume_group`. Conflicts with `id`.",
				MarkdownDescription: "Filter recycle bin items by resource type. Valid values are `volume` and `volume_group`. Conflicts with `id`.",
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("volume", "volume_group"),
					stringvalidator.ConflictsWith(path.MatchRoot("id")),
				},
			},
			"filter_expression": schema.StringAttribute{
				Description:         "PowerStore filter expression to filter recycle bin items. Conflicts with `id`.",
				MarkdownDescription: "PowerStore filter expression to filter recycle bin items. Conflicts with `id`.",
				Optional:            true,
				CustomType:          models.FilterExpressionType{},
			},
			"recycle_bin_items": schema.ListNestedAttribute{
				Description:         "List of recycle bin items.",
				MarkdownDescription: "List of recycle bin items.",
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.StringAttribute{
							Description:         "Unique identifier for the recycle bin item.",
							MarkdownDescription: "Unique identifier for the recycle bin item.",
							Computed:            true,
						},
						"name": schema.StringAttribute{
							Description:         "The name of the deleted object.",
							MarkdownDescription: "The name of the deleted object.",
							Computed:            true,
						},
						"resource_type": schema.StringAttribute{
							Description:         "Type of the storage object (volume or volume_group).",
							MarkdownDescription: "Type of the storage object (`volume` or `volume_group`).",
							Computed:            true,
						},
						"logical_provisioned": schema.Int64Attribute{
							Description:         "Provisioned size of the object in bytes.",
							MarkdownDescription: "Provisioned size of the object in bytes.",
							Computed:            true,
						},
						"logical_used": schema.Int64Attribute{
							Description:         "Logical space used by the object in bytes.",
							MarkdownDescription: "Logical space used by the object in bytes.",
							Computed:            true,
						},
						"appliance_id": schema.StringAttribute{
							Description:         "The appliance where this resource is located.",
							MarkdownDescription: "The appliance where this resource is located.",
							Computed:            true,
						},
						"deletion_timestamp": schema.StringAttribute{
							Description:         "Time when the object was moved to the recycle bin.",
							MarkdownDescription: "Time when the object was moved to the recycle bin.",
							Computed:            true,
						},
						"expiration_timestamp": schema.StringAttribute{
							Description:         "Time when the object will be auto-purged.",
							MarkdownDescription: "Time when the object will be auto-purged.",
							Computed:            true,
						},
						"resource_type_l10n": schema.StringAttribute{
							Description:         "Localized message string corresponding to resource_type.",
							MarkdownDescription: "Localized message string corresponding to resource_type.",
							Computed:            true,
						},
					},
				},
			},
		},
	}
}

func (d *recycleBinDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	c, ok := req.ProviderData.(*client.Client)
	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected Data Source Configure Type",
			fmt.Sprintf("Expected *client.Client, got: %T.", req.ProviderData),
		)
		return
	}
	d.client = c.GenClient
}

// getRecycleBinItems fetches all recycle bin items
func (d *recycleBinDataSource) getRecycleBinItems(ctx context.Context) ([]clientgen.RecycleBinInstance, error) {
	result, _, err := d.client.RecycleBinApi.GetAllRecycleBins(ctx).Execute()
	return result, err
}

// getRecycleBinItem fetches a specific recycle bin item by ID
func (d *recycleBinDataSource) getRecycleBinItem(ctx context.Context, id string) (*clientgen.RecycleBinInstance, error) {
	result, _, err := d.client.RecycleBinApi.GetRecycleBinById(ctx, id).Execute()
	return result, err
}

// getRecycleBinItemsByFilter fetches recycle bin items with query filters
func (d *recycleBinDataSource) getRecycleBinItemsByFilter(ctx context.Context, filters map[string]string) ([]clientgen.RecycleBinInstance, error) {
	queries := make(url.Values)
	for k, v := range filters {
		queries.Set(k, v)
	}
	result, _, err := d.client.RecycleBinApi.GetAllRecycleBins(ctx).Queries(queries).Execute()
	return result, err
}

func (d *recycleBinDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var state models.RecycleBinDataSourceModel
	var items []clientgen.RecycleBinInstance
	var err error

	diags := req.Config.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)

	if state.ID.ValueString() != "" {
		item, getErr := d.getRecycleBinItem(ctx, state.ID.ValueString())
		if getErr != nil {
			resp.Diagnostics.AddError(
				"Unable to Read PowerStore Recycle Bin Item",
				getErr.Error(),
			)
			return
		}
		if item != nil {
			items = append(items, *item)
		}
	} else if state.Filters.ValueString() != "" {
		filterMap := convertQueriesToMap(state.Filters.ValueQueries())
		items, err = d.getRecycleBinItemsByFilter(ctx, filterMap)
	} else if state.ResourceType.ValueString() != "" {
		filterMap := map[string]string{
			"resource_type": "eq." + state.ResourceType.ValueString(),
		}
		items, err = d.getRecycleBinItemsByFilter(ctx, filterMap)
	} else {
		items, err = d.getRecycleBinItems(ctx)
	}

	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to Read PowerStore Recycle Bin Items",
			err.Error(),
		)
		return
	}

	state.RecycleBinItems = mapRecycleBinItemsToState(items)
	state.ID = types.StringValue("placeholder")

	diags = resp.State.Set(ctx, state)
	resp.Diagnostics.Append(diags...)
}

// mapRecycleBinItemsToState converts clientgen RecycleBinInstance items to Terraform state models.
func mapRecycleBinItemsToState(items []clientgen.RecycleBinInstance) []models.RecycleBinItemModel {
	var result []models.RecycleBinItemModel
	for _, item := range items {
		model := models.RecycleBinItemModel{}
		if item.Id != nil {
			model.ID = types.StringValue(*item.Id)
		}
		if item.Name != nil {
			model.Name = types.StringValue(*item.Name)
		}
		if item.ResourceType != nil {
			model.ResourceType = types.StringValue(string(*item.ResourceType))
		}
		if item.LogicalProvisioned != nil {
			model.LogicalProvisioned = types.Int64Value(*item.LogicalProvisioned)
		}
		if item.LogicalUsed != nil {
			model.LogicalUsed = types.Int64Value(*item.LogicalUsed)
		}
		if item.ApplianceId != nil {
			model.ApplianceID = types.StringValue(*item.ApplianceId)
		}
		if item.DeletionTimestamp != nil {
			model.DeletionTimestamp = types.StringValue(item.DeletionTimestamp.Format(time.RFC3339))
		}
		if item.ExpirationTimestamp != nil {
			model.ExpirationTimestamp = types.StringValue(item.ExpirationTimestamp.Format(time.RFC3339))
		}
		if item.ResourceTypeL10n != nil {
			model.ResourceTypeL10N = types.StringValue(*item.ResourceTypeL10n)
		}
		result = append(result, model)
	}
	return result
}
