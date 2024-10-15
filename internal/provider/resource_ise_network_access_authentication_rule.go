// Copyright © 2023 Cisco Systems, Inc. and its affiliates.
// All rights reserved.
//
// Licensed under the Mozilla Public License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://mozilla.org/MPL/2.0/
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// SPDX-License-Identifier: MPL-2.0

// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

//template:begin imports
import (
	"context"
	"fmt"
	"net/url"
	"strings"

	"github.com/CiscoDevNet/terraform-provider-ise/internal/provider/helpers"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/netascode/go-ise"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

//template:end imports

//template:begin header

// Ensure provider defined types fully satisfy framework interfaces
var _ resource.Resource = &NetworkAccessAuthenticationRuleResource{}
var _ resource.ResourceWithImportState = &NetworkAccessAuthenticationRuleResource{}

func NewNetworkAccessAuthenticationRuleResource() resource.Resource {
	return &NetworkAccessAuthenticationRuleResource{}
}

type NetworkAccessAuthenticationRuleResource struct {
	client *ise.Client
}

func (r *NetworkAccessAuthenticationRuleResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_network_access_authentication_rule"
}

//template:end header

//template:begin model
func (r *NetworkAccessAuthenticationRuleResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: helpers.NewAttributeDescription("This resource can manage a Network Access Authentication Rule.").String,

		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				MarkdownDescription: "The id of the object",
				Computed:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"policy_set_id": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Policy set ID").String,
				Required:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"name": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Rule name, [Valid characters are alphanumerics, underscore, hyphen, space, period, parentheses]").String,
				Required:            true,
			},
			"default": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Indicates if this rule is the default one").String,
				Optional:            true,
			},
			"rank": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("The rank (priority) in relation to other rules. Lower rank is higher priority.").String,
				Optional:            true,
			},
			"state": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("The state that the rule is in. A disabled rule cannot be matched.").AddStringEnumDescription("disabled", "enabled", "monitor").String,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("disabled", "enabled", "monitor"),
				},
			},
			"condition_type": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Indicates whether the record is the condition itself or a logical aggregation. Logical aggreation indicates that additional conditions are present under the children attribute.").AddStringEnumDescription("ConditionAndBlock", "ConditionAttributes", "ConditionOrBlock", "ConditionReference").String,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("ConditionAndBlock", "ConditionAttributes", "ConditionOrBlock", "ConditionReference"),
				},
			},
			"condition_id": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("UUID for condition").String,
				Optional:            true,
			},
			"condition_is_negate": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Indicates whereas this condition is in negate mode").String,
				Optional:            true,
			},
			"condition_attribute_name": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Dictionary attribute name").String,
				Optional:            true,
			},
			"condition_attribute_value": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Attribute value for condition. Value type is specified in dictionary object.").String,
				Optional:            true,
			},
			"condition_dictionary_name": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Dictionary name").String,
				Optional:            true,
			},
			"condition_dictionary_value": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Dictionary value").String,
				Optional:            true,
			},
			"condition_operator": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Equality operator").AddStringEnumDescription("contains", "endsWith", "equals", "greaterOrEquals", "greaterThan", "in", "ipEquals", "ipGreaterThan", "ipLessThan", "ipNotEquals", "lessOrEquals", "lessThan", "matches", "notContains", "notEndsWith", "notEquals", "notIn", "notStartsWith", "startsWith").String,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("contains", "endsWith", "equals", "greaterOrEquals", "greaterThan", "in", "ipEquals", "ipGreaterThan", "ipLessThan", "ipNotEquals", "lessOrEquals", "lessThan", "matches", "notContains", "notEndsWith", "notEquals", "notIn", "notStartsWith", "startsWith"),
				},
			},
			"children": schema.ListNestedAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("List of child conditions. `condition_type` must be one of `ConditionAndBlock` or `ConditionOrBlock`.").String,
				Optional:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"condition_type": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Indicates whether the record is the condition itself or a logical aggregation. Logical aggreation indicates that additional conditions are present under the children attribute.").AddStringEnumDescription("ConditionAndBlock", "ConditionAttributes", "ConditionOrBlock", "ConditionReference").String,
							Required:            true,
							Validators: []validator.String{
								stringvalidator.OneOf("ConditionAndBlock", "ConditionAttributes", "ConditionOrBlock", "ConditionReference"),
							},
						},
						"id": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("UUID for condition").String,
							Optional:            true,
						},
						"is_negate": schema.BoolAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Indicates whereas this condition is in negate mode").String,
							Optional:            true,
						},
						"attribute_name": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Dictionary attribute name").String,
							Optional:            true,
						},
						"attribute_value": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Attribute value for condition. Value type is specified in dictionary object.").String,
							Optional:            true,
						},
						"dictionary_name": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Dictionary name").String,
							Optional:            true,
						},
						"dictionary_value": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Dictionary value").String,
							Optional:            true,
						},
						"operator": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Equality operator").AddStringEnumDescription("contains", "endsWith", "equals", "greaterOrEquals", "greaterThan", "in", "ipEquals", "ipGreaterThan", "ipLessThan", "ipNotEquals", "lessOrEquals", "lessThan", "matches", "notContains", "notEndsWith", "notEquals", "notIn", "notStartsWith", "startsWith").String,
							Optional:            true,
							Validators: []validator.String{
								stringvalidator.OneOf("contains", "endsWith", "equals", "greaterOrEquals", "greaterThan", "in", "ipEquals", "ipGreaterThan", "ipLessThan", "ipNotEquals", "lessOrEquals", "lessThan", "matches", "notContains", "notEndsWith", "notEquals", "notIn", "notStartsWith", "startsWith"),
							},
						},
						"children": schema.ListNestedAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("List of child conditions. `condition_type` must be one of `ConditionAndBlock` or `ConditionOrBlock`.").String,
							Optional:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"condition_type": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Condition type.").AddStringEnumDescription("ConditionAttributes", "ConditionReference").String,
										Required:            true,
										Validators: []validator.String{
											stringvalidator.OneOf("ConditionAttributes", "ConditionReference"),
										},
									},
									"id": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("UUID for condition").String,
										Optional:            true,
									},
									"is_negate": schema.BoolAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Indicates whereas this condition is in negate mode").String,
										Optional:            true,
									},
									"attribute_name": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Dictionary attribute name").String,
										Optional:            true,
									},
									"attribute_value": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Attribute value for condition. Value type is specified in dictionary object.").String,
										Optional:            true,
									},
									"dictionary_name": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Dictionary name").String,
										Optional:            true,
									},
									"dictionary_value": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Dictionary value").String,
										Optional:            true,
									},
									"operator": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Equality operator").AddStringEnumDescription("contains", "endsWith", "equals", "greaterOrEquals", "greaterThan", "in", "ipEquals", "ipGreaterThan", "ipLessThan", "ipNotEquals", "lessOrEquals", "lessThan", "matches", "notContains", "notEndsWith", "notEquals", "notIn", "notStartsWith", "startsWith").String,
										Optional:            true,
										Validators: []validator.String{
											stringvalidator.OneOf("contains", "endsWith", "equals", "greaterOrEquals", "greaterThan", "in", "ipEquals", "ipGreaterThan", "ipLessThan", "ipNotEquals", "lessOrEquals", "lessThan", "matches", "notContains", "notEndsWith", "notEquals", "notIn", "notStartsWith", "startsWith"),
										},
									},
								},
							},
						},
					},
				},
			},
			"identity_source_name": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Identity source name from the identity stores").String,
				Optional:            true,
			},
			"if_auth_fail": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Action to perform when authentication fails such as Bad credentials, disabled user and so on").AddStringEnumDescription("REJECT", "DROP", "CONTINUE").String,
				Required:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("REJECT", "DROP", "CONTINUE"),
				},
			},
			"if_process_fail": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Action to perform when ISE is uanble to access the identity database").AddStringEnumDescription("REJECT", "DROP", "CONTINUE").String,
				Required:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("REJECT", "DROP", "CONTINUE"),
				},
			},
			"if_user_not_found": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Action to perform when user is not found in any of identity stores").AddStringEnumDescription("REJECT", "DROP", "CONTINUE").String,
				Required:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("REJECT", "DROP", "CONTINUE"),
				},
			},
		},
	}
}

//template:end model

//template:begin configure
func (r *NetworkAccessAuthenticationRuleResource) Configure(_ context.Context, req resource.ConfigureRequest, _ *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	r.client = req.ProviderData.(*IseProviderData).Client
}

//template:end configure

func (r *NetworkAccessAuthenticationRuleResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan NetworkAccessAuthenticationRule

	// Read plan
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Create", plan.Id.ValueString()))

	// Create object
	body := plan.toBody(ctx, NetworkAccessAuthenticationRule{})

	if plan.Name.ValueString() != "Default" {
		res, _, err := r.client.Post(plan.getPath(), body)
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to configure object (POST), got error: %s, %s", err, res.String()))
			return
		}
		plan.Id = types.StringValue(res.Get("response.rule.id").String())
	} else {
		res, err := r.client.Get(plan.getPath())
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve objects, got error: %s", err))
			return
		}
		if value := res.Get("response"); len(value.Array()) > 0 {
			value.ForEach(func(k, v gjson.Result) bool {
				if v.Get("rule.name").String() == plan.Name.ValueString() {
					plan.Id = types.StringValue(v.Get("rule.id").String())
					return false
				}
				return true
			})
		}
		res, err = r.client.Put(plan.getPath()+"/"+plan.Id.ValueString(), body)
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to configure object (PUT), got error: %s, %s", err, res.String()))
			return
		}
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Create finished successfully", plan.Id.ValueString()))

	diags = resp.State.Set(ctx, &plan)
	resp.Diagnostics.Append(diags...)
}

//template:begin read
func (r *NetworkAccessAuthenticationRuleResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state NetworkAccessAuthenticationRule

	// Read state
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Read", state.Id.String()))
	res, err := r.client.Get(state.getPath() + "/" + url.QueryEscape(state.Id.ValueString()))
	if err != nil && strings.Contains(err.Error(), "StatusCode 404") {
		resp.State.RemoveResource(ctx)
		return
	} else if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve object (GET), got error: %s, %s", err, res.String()))
		return
	}

	// If every attribute is set to null we are dealing with an import operation and therefore reading all attributes
	if state.isNull(ctx, res) {
		state.fromBody(ctx, res)
	} else {
		state.updateFromBody(ctx, res)
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Read finished successfully", state.Id.ValueString()))

	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
}

//template:end read

func (r *NetworkAccessAuthenticationRuleResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan, state, existingData NetworkAccessAuthenticationRule

	// Read plan
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	// Read state
	diags = req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Update", plan.Id.ValueString()))
	body := plan.toBody(ctx, state)

	// Check if plan.Rank is null (i.e., not provided) and set rank to body from existingData to not reorder rule during update
	if plan.Rank.IsNull() {
		// Read existing attributes from the API
		res, err := r.client.Get(plan.getPath() + "/" + url.QueryEscape(plan.Id.ValueString()))
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve object (GET), got error: %s", err))
			return
		}
		existingData.fromBody(ctx, res)

		// If plan.Rank is not provided, set it from existingData.Rank
		body, _ = sjson.Set(body, "rule.rank", existingData.Rank.ValueInt64())
	}

	res, err := r.client.Put(plan.getPath()+"/"+url.QueryEscape(plan.Id.ValueString()), body)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to configure object (PUT), got error: %s, %s", err, res.String()))
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Update finished successfully", plan.Id.ValueString()))

	diags = resp.State.Set(ctx, &plan)
	resp.Diagnostics.Append(diags...)
}

//template:begin delete
func (r *NetworkAccessAuthenticationRuleResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state NetworkAccessAuthenticationRule

	// Read state
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Delete", state.Id.ValueString()))
	res, err := r.client.Delete(state.getPath() + "/" + url.QueryEscape(state.Id.ValueString()))
	if err != nil && !strings.Contains(res.String(), "Attempted to delete default") {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to delete object (DELETE), got error: %s, %s", err, res.String()))
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Delete finished successfully", state.Id.ValueString()))

	resp.State.RemoveResource(ctx)
}

//template:end delete

//template:begin import
func (r *NetworkAccessAuthenticationRuleResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	idParts := strings.Split(req.ID, ",")

	if len(idParts) != 2 || idParts[0] == "" || idParts[1] == "" {
		resp.Diagnostics.AddError(
			"Unexpected Import Identifier",
			fmt.Sprintf("Expected import identifier with format: <policy_set_id>,<id>. Got: %q", req.ID),
		)
		return
	}
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("policy_set_id"), idParts[0])...)
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("id"), idParts[1])...)
}

//template:end import
