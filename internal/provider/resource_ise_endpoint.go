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
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/netascode/go-ise"
)

//template:end imports

//template:begin header

// Ensure provider defined types fully satisfy framework interfaces
var _ resource.Resource = &EndpointResource{}
var _ resource.ResourceWithImportState = &EndpointResource{}

func NewEndpointResource() resource.Resource {
	return &EndpointResource{}
}

type EndpointResource struct {
	client *ise.Client
}

func (r *EndpointResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_endpoint"
}

//template:end header

//template:begin model
func (r *EndpointResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: helpers.NewAttributeDescription("This resource can manage an Endpoint.").String,

		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				MarkdownDescription: "The id of the object",
				Computed:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"name": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("The name of the endpoint").String,
				Required:            true,
			},
			"description": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Description").String,
				Optional:            true,
			},
			"mac": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("MAC address of the endpoint").String,
				Required:            true,
			},
			"group_id": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Identity Group ID").String,
				Required:            true,
			},
			"profile_id": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Profile ID").String,
				Optional:            true,
			},
			"static_profile_assignment": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Static Profile Assignment").String,
				Required:            true,
			},
			"static_profile_assignment_defined": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Static Profile Assignment Defined").AddDefaultValueDescription("true").String,
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(true),
			},
			"static_group_assignment": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Static Group Assignment").String,
				Required:            true,
			},
			"static_group_assignment_defined": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("staticGroupAssignmentDefined").AddDefaultValueDescription("true").String,
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(true),
			},
			"custom_attributes": schema.MapAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Custom Attributes").String,
				ElementType:         types.StringType,
				Optional:            true,
			},
			"identity_store": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Identity Store").String,
				Optional:            true,
			},
			"identity_store_id": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Identity Store Id").String,
				Optional:            true,
			},
			"portal_user": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Portal User").String,
				Optional:            true,
			},
			"mdm_server_name": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Mdm Server Name").String,
				Optional:            true,
			},
			"mdm_reachable": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Mdm Reachable").String,
				Optional:            true,
			},
			"mdm_enrolled": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Mdm Enrolled").String,
				Optional:            true,
			},
			"mdm_compliance_status": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Mdm Compliance Status").String,
				Optional:            true,
			},
			"mdm_os": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Mdm OS").String,
				Optional:            true,
			},
			"mdm_manufacturer": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Mdm Manufacturer").String,
				Optional:            true,
			},
			"mdm_model": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Mdm Model").String,
				Optional:            true,
			},
			"mdm_serial": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Mdm Serial").String,
				Optional:            true,
			},
			"mdm_encrypted": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Mdm Encrypted").String,
				Optional:            true,
			},
			"mdm_pinlock": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Mdm Pinlock").String,
				Optional:            true,
			},
			"mdm_jail_broken": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Mdm JailBroken").String,
				Optional:            true,
			},
			"mdm_imei": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Mdm IMEI").String,
				Optional:            true,
			},
			"mdm_phone_number": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Mdm PhoneNumber").String,
				Optional:            true,
			},
		},
	}
}

//template:end model

//template:begin configure
func (r *EndpointResource) Configure(_ context.Context, req resource.ConfigureRequest, _ *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	r.client = req.ProviderData.(*IseProviderData).Client
}

//template:end configure

//template:begin create
func (r *EndpointResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan Endpoint

	// Read plan
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Create", plan.Id.ValueString()))

	// Create object
	body := plan.toBody(ctx, Endpoint{})
	res, location, err := r.client.Post(plan.getPath(), body)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to configure object (POST), got error: %s, %s", err, res.String()))
		return
	}
	locationElements := strings.Split(location, "/")
	plan.Id = types.StringValue(locationElements[len(locationElements)-1])

	tflog.Debug(ctx, fmt.Sprintf("%s: Create finished successfully", plan.Id.ValueString()))

	diags = resp.State.Set(ctx, &plan)
	resp.Diagnostics.Append(diags...)
}

//template:end create

//template:begin read
func (r *EndpointResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state Endpoint

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

//template:begin update
func (r *EndpointResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan, state Endpoint

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

	res, err := r.client.Put(plan.getPath()+"/"+url.QueryEscape(plan.Id.ValueString()), body)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to configure object (PUT), got error: %s, %s", err, res.String()))
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Update finished successfully", plan.Id.ValueString()))

	diags = resp.State.Set(ctx, &plan)
	resp.Diagnostics.Append(diags...)
}

//template:end update

//template:begin delete
func (r *EndpointResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state Endpoint

	// Read state
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Delete", state.Id.ValueString()))
	res, err := r.client.Delete(state.getPath() + "/" + url.QueryEscape(state.Id.ValueString()))
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to delete object (DELETE), got error: %s, %s", err, res.String()))
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Delete finished successfully", state.Id.ValueString()))

	resp.State.RemoveResource(ctx)
}

//template:end delete

//template:begin import
func (r *EndpointResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

//template:end import
