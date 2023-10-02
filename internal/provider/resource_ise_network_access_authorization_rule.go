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

import (
	"context"
	"fmt"
	"strings"

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
	"github.com/netascode/terraform-provider-ise/internal/provider/helpers"
)

// Ensure provider defined types fully satisfy framework interfaces
var _ resource.Resource = &NetworkAccessAuthorizationRuleResource{}
var _ resource.ResourceWithImportState = &NetworkAccessAuthorizationRuleResource{}

func NewNetworkAccessAuthorizationRuleResource() resource.Resource {
	return &NetworkAccessAuthorizationRuleResource{}
}

type NetworkAccessAuthorizationRuleResource struct {
	client *ise.Client
}

func (r *NetworkAccessAuthorizationRuleResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_network_access_authorization_rule"
}

func (r *NetworkAccessAuthorizationRuleResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: helpers.NewAttributeDescription("This resource can manage a Network Access Authorization Rule.").String,

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
				MarkdownDescription: helpers.NewAttributeDescription("Indicates whether the record is the condition itself or a logical aggregation. Logical aggreation indicates that additional conditions are present under the children attribute.").AddStringEnumDescription("ConditionAndBlock", "ConditionAttributes", "ConditionOrBlock", "ConditionReference", "LibraryConditionAndBlock", "LibraryConditionAttributes", "LibraryConditionOrBlock", "TimeAndDateCondition").String,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("ConditionAndBlock", "ConditionAttributes", "ConditionOrBlock", "ConditionReference", "LibraryConditionAndBlock", "LibraryConditionAttributes", "LibraryConditionOrBlock", "TimeAndDateCondition"),
				},
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
			"profile": schema.ListAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("The authorization profile(s)").String,
				ElementType:         types.StringType,
				Optional:            true,
			},
			"security_group": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Security group used in authorization policies").String,
				Optional:            true,
			},
		},
	}
}

func (r *NetworkAccessAuthorizationRuleResource) Configure(_ context.Context, req resource.ConfigureRequest, _ *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	r.client = req.ProviderData.(*IseProviderData).Client
}

func (r *NetworkAccessAuthorizationRuleResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan NetworkAccessAuthorizationRule

	// Read plan
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Create", plan.Id.ValueString()))

	// Create object
	body := plan.toBody(ctx, NetworkAccessAuthorizationRule{})
	res, _, err := r.client.Post(plan.getPath(), body)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to configure object (POST), got error: %s, %s", err, res.String()))
		return
	}
	plan.Id = types.StringValue(res.Get("response.rule.id").String())

	tflog.Debug(ctx, fmt.Sprintf("%s: Create finished successfully", plan.Id.ValueString()))

	diags = resp.State.Set(ctx, &plan)
	resp.Diagnostics.Append(diags...)
}

func (r *NetworkAccessAuthorizationRuleResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state NetworkAccessAuthorizationRule

	// Read state
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Read", state.Id.String()))

	res, err := r.client.Get(state.getPath() + "/" + state.Id.ValueString())
	if err != nil && strings.Contains(err.Error(), "StatusCode 404") {
		resp.State.RemoveResource(ctx)
		return
	} else if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve object (GET), got error: %s, %s", err, res.String()))
		return
	}

	state.updateFromBody(ctx, res)

	tflog.Debug(ctx, fmt.Sprintf("%s: Read finished successfully", state.Id.ValueString()))

	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
}

func (r *NetworkAccessAuthorizationRuleResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan, state NetworkAccessAuthorizationRule

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

	res, err := r.client.Put(plan.getPath()+"/"+plan.Id.ValueString(), body)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to configure object (PUT), got error: %s, %s", err, res.String()))
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Update finished successfully", plan.Id.ValueString()))

	diags = resp.State.Set(ctx, &plan)
	resp.Diagnostics.Append(diags...)
}

func (r *NetworkAccessAuthorizationRuleResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state NetworkAccessAuthorizationRule

	// Read state
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Delete", state.Id.ValueString()))
	res, err := r.client.Delete(state.getPath() + "/" + state.Id.ValueString())
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to delete object (DELETE), got error: %s, %s", err, res.String()))
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Delete finished successfully", state.Id.ValueString()))

	resp.State.RemoveResource(ctx)
}

func (r *NetworkAccessAuthorizationRuleResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}
