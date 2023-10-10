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
	"strings"

	"github.com/CiscoDevNet/terraform-provider-ise/internal/provider/helpers"
	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/netascode/go-ise"
)

//template:end imports

//template:begin model

// Ensure provider defined types fully satisfy framework interfaces
var _ resource.Resource = &AuthorizationProfileResource{}
var _ resource.ResourceWithImportState = &AuthorizationProfileResource{}

func NewAuthorizationProfileResource() resource.Resource {
	return &AuthorizationProfileResource{}
}

type AuthorizationProfileResource struct {
	client *ise.Client
}

func (r *AuthorizationProfileResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_authorization_profile"
}

func (r *AuthorizationProfileResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: helpers.NewAttributeDescription("This resource can manage an authorization profiles policy element.").String,

		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				MarkdownDescription: "The id of the object",
				Computed:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"name": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("The name of the authorization profile").String,
				Required:            true,
			},
			"description": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Description").String,
				Optional:            true,
			},
			"vlan_name_id": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Vlan name or ID").String,
				Optional:            true,
			},
			"vlan_tag_id": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("Vlan tag ID").AddIntegerRangeDescription(0, 31).String,
				Optional:            true,
				Validators: []validator.Int64{
					int64validator.Between(0, 31),
				},
			},
			"web_redirection_type": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("This type must fit the `web_redirection_portal_name`").AddStringEnumDescription("CentralizedWebAuth", "HotSpot", "NativeSupplicanProvisioning", "ClientProvisioning").String,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("CentralizedWebAuth", "HotSpot", "NativeSupplicanProvisioning", "ClientProvisioning"),
				},
			},
			"web_redirection_acl": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Web redirection ACL").String,
				Optional:            true,
			},
			"web_redirection_portal_name": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("A portal that exist in the DB and fits the `web_redirection_type`").String,
				Optional:            true,
			},
			"web_redirection_static_ip_host_name_fqdn": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("IP, hostname or FQDN").String,
				Optional:            true,
			},
			"web_redirection_display_certificates_renewal_messages": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("This attribute is mandatory when `web_redirection_type` value is `CentralizedWebAuth`. For all other `web_redirection_type` values the field must be ignored.").String,
				Optional:            true,
			},
			"agentless_posture": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Agentless Posture.").String,
				Optional:            true,
			},
			"access_type": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Access type").AddStringEnumDescription("ACCESS_ACCEPT", "ACCESS_REJECT").AddDefaultValueDescription("ACCESS_ACCEPT").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("ACCESS_ACCEPT", "ACCESS_REJECT"),
				},
				Default: stringdefault.StaticString("ACCESS_ACCEPT"),
			},
			"profile_name": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Value needs to be an existing Network Device Profile").AddDefaultValueDescription("Cisco").String,
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString("Cisco"),
			},
			"airespace_acl": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Airespace ACL").String,
				Optional:            true,
			},
			"acl": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("ACL").String,
				Optional:            true,
			},
			"dacl_name": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("DACL name").String,
				Optional:            true,
			},
			"auto_smart_port": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Auto smart port").String,
				Optional:            true,
			},
			"interface_template": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Interface template").String,
				Optional:            true,
			},
			"ipv6_acl_filter": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("IPv6 ACL").String,
				Optional:            true,
			},
			"avc_profile": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("AVC profile").String,
				Optional:            true,
			},
			"asa_vpn": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("ASA VPN").String,
				Optional:            true,
			},
			"unique_identifier": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Unique identifier").String,
				Optional:            true,
			},
			"track_movement": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Track movement").AddDefaultValueDescription("false").String,
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"service_template": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Service template").AddDefaultValueDescription("false").String,
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"easywired_session_candidate": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Easy wired session candidate").AddDefaultValueDescription("false").String,
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"voice_domain_permission": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Voice domain permission").AddDefaultValueDescription("false").String,
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"neat": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("NEAT").AddDefaultValueDescription("false").String,
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"web_auth": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Web authentication (local)").AddDefaultValueDescription("false").String,
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"mac_sec_policy": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("MacSec policy").AddStringEnumDescription("MUST_SECURE", "MUST_NOT_SECURE", "SHOULD_SECURE").String,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("MUST_SECURE", "MUST_NOT_SECURE", "SHOULD_SECURE"),
				},
			},
			"reauthentication_connectivity": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Maintain Connectivity During Reauthentication").AddStringEnumDescription("DEFAULT", "RADIUS_REQUEST").String,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("DEFAULT", "RADIUS_REQUEST"),
				},
			},
			"reauthentication_timer": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("Reauthentication timer").AddIntegerRangeDescription(1, 65535).String,
				Optional:            true,
				Validators: []validator.Int64{
					int64validator.Between(1, 65535),
				},
			},
			"advanced_attributes": schema.ListNestedAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("List of advanced attributes").String,
				Optional:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"attribute_1_value_type": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Advanced attribute value type").AddStringEnumDescription("AdvancedDictionaryAttribute", "AttributeValue").String,
							Required:            true,
							Validators: []validator.String{
								stringvalidator.OneOf("AdvancedDictionaryAttribute", "AttributeValue"),
							},
						},
						"attribute_1_dictionary_name": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Dictionary name").String,
							Required:            true,
						},
						"attribute_1_name": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Attribute name").String,
							Required:            true,
						},
						"attribute_2_value_type": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Advanced attribute value type").AddStringEnumDescription("AdvancedDictionaryAttribute", "AttributeValue").String,
							Required:            true,
							Validators: []validator.String{
								stringvalidator.OneOf("AdvancedDictionaryAttribute", "AttributeValue"),
							},
						},
						"attribute_2_value": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Attribute value").String,
							Required:            true,
						},
					},
				},
			},
			"ipv6_dacl_name": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("IPv6 DACL name").String,
				Optional:            true,
			},
			"airespace_ipv6_acl": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Airespace IPv6 ACL").String,
				Optional:            true,
			},
		},
	}
}

func (r *AuthorizationProfileResource) Configure(_ context.Context, req resource.ConfigureRequest, _ *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	r.client = req.ProviderData.(*IseProviderData).Client
}

//template:end model

//template:begin create
func (r *AuthorizationProfileResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan AuthorizationProfile

	// Read plan
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Create", plan.Id.ValueString()))

	// Create object
	body := plan.toBody(ctx, AuthorizationProfile{})
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
func (r *AuthorizationProfileResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state AuthorizationProfile

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

//template:end read

//template:begin update
func (r *AuthorizationProfileResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan, state AuthorizationProfile

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

//template:end update

//template:begin delete
func (r *AuthorizationProfileResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state AuthorizationProfile

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

//template:end delete

//template:begin import
func (r *AuthorizationProfileResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

//template:end import
