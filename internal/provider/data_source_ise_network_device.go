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

	"github.com/hashicorp/terraform-plugin-framework-validators/datasourcevalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/netascode/go-ise"
	"github.com/tidwall/gjson"
)

//template:end imports

//template:begin model

// Ensure the implementation satisfies the expected interfaces.
var (
	_ datasource.DataSource              = &NetworkDeviceDataSource{}
	_ datasource.DataSourceWithConfigure = &NetworkDeviceDataSource{}
)

func NewNetworkDeviceDataSource() datasource.DataSource {
	return &NetworkDeviceDataSource{}
}

type NetworkDeviceDataSource struct {
	client *ise.Client
}

func (d *NetworkDeviceDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_network_device"
}

func (d *NetworkDeviceDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: "This data source can read the Network Device.",

		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				MarkdownDescription: "The id of the object",
				Optional:            true,
				Computed:            true,
			},
			"name": schema.StringAttribute{
				MarkdownDescription: "The name of the network device",
				Optional:            true,
				Computed:            true,
			},
			"description": schema.StringAttribute{
				MarkdownDescription: "Description",
				Computed:            true,
			},
			"authentication_enable_key_wrap": schema.BoolAttribute{
				MarkdownDescription: "Enable key wrap",
				Computed:            true,
			},
			"authentication_encryption_key": schema.StringAttribute{
				MarkdownDescription: "Encryption key",
				Computed:            true,
			},
			"authentication_encryption_key_format": schema.StringAttribute{
				MarkdownDescription: "Key input format",
				Computed:            true,
			},
			"authentication_message_authenticator_code_key": schema.StringAttribute{
				MarkdownDescription: "Message authenticator code key",
				Computed:            true,
			},
			"authentication_network_protocol": schema.StringAttribute{
				MarkdownDescription: "Network protocol",
				Computed:            true,
			},
			"authentication_radius_shared_secret": schema.StringAttribute{
				MarkdownDescription: "RADIUS shared secret",
				Computed:            true,
			},
			"authentication_enable_multi_secret": schema.BoolAttribute{
				MarkdownDescription: "Enable multiple RADIUS shared secrets",
				Computed:            true,
			},
			"authentication_second_radius_shared_secret": schema.StringAttribute{
				MarkdownDescription: "Second RADIUS shared secret",
				Computed:            true,
			},
			"authentication_dtls_required": schema.BoolAttribute{
				MarkdownDescription: "Enforce use of DTLS",
				Computed:            true,
			},
			"coa_port": schema.Int64Attribute{
				MarkdownDescription: "CoA port",
				Computed:            true,
			},
			"dtls_dns_name": schema.StringAttribute{
				MarkdownDescription: "This value is used to verify the client identity contained in the X.509 RADIUS/DTLS client certificate",
				Computed:            true,
			},
			"ips": schema.ListNestedAttribute{
				MarkdownDescription: "List of IP subnets",
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"ipaddress": schema.StringAttribute{
							MarkdownDescription: "It can be either single ip address or ip range address",
							Computed:            true,
						},
						"ipaddress_exclude": schema.StringAttribute{
							MarkdownDescription: "It can be either single ip address or ip range address",
							Computed:            true,
						},
						"mask": schema.StringAttribute{
							MarkdownDescription: "Subnet mask length",
							Computed:            true,
						},
					},
				},
			},
			"network_device_groups": schema.SetAttribute{
				MarkdownDescription: "List of network device groups, e.g. `Device Type#All Device Types#ACCESS`",
				ElementType:         types.StringType,
				Computed:            true,
			},
			"model_name": schema.StringAttribute{
				MarkdownDescription: "Model name",
				Computed:            true,
			},
			"software_version": schema.StringAttribute{
				MarkdownDescription: "Software version",
				Computed:            true,
			},
			"profile_name": schema.StringAttribute{
				MarkdownDescription: "Profile name",
				Computed:            true,
			},
			"snmp_link_trap_query": schema.BoolAttribute{
				MarkdownDescription: "SNMP link Trap Query",
				Computed:            true,
			},
			"snmp_mac_trap_query": schema.BoolAttribute{
				MarkdownDescription: "SNMP MAC Trap Query",
				Computed:            true,
			},
			"snmp_originating_policy_service_node": schema.StringAttribute{
				MarkdownDescription: "Originating Policy Services Node",
				Computed:            true,
			},
			"snmp_polling_interval": schema.Int64Attribute{
				MarkdownDescription: "SNMP Polling Interval in seconds",
				Computed:            true,
			},
			"snmp_ro_community": schema.StringAttribute{
				MarkdownDescription: "SNMP RO Community",
				Computed:            true,
			},
			"snmp_version": schema.StringAttribute{
				MarkdownDescription: "SNMP version",
				Computed:            true,
			},
			"tacacs_connect_mode_options": schema.StringAttribute{
				MarkdownDescription: "Connect mode options",
				Computed:            true,
			},
			"tacacs_shared_secret": schema.StringAttribute{
				MarkdownDescription: "Shared secret",
				Computed:            true,
			},
			"trustsec_device_id": schema.StringAttribute{
				MarkdownDescription: "TrustSec device ID",
				Computed:            true,
			},
			"trustsec_device_password": schema.StringAttribute{
				MarkdownDescription: "TrustSec device password",
				Computed:            true,
			},
			"trustsec_rest_api_username": schema.StringAttribute{
				MarkdownDescription: "REST API username",
				Computed:            true,
			},
			"trustsec_rest_api_password": schema.StringAttribute{
				MarkdownDescription: "REST API password",
				Computed:            true,
			},
			"trustsec_enable_mode_password": schema.StringAttribute{
				MarkdownDescription: "Enable mode password",
				Computed:            true,
			},
			"trustsec_exec_mode_password": schema.StringAttribute{
				MarkdownDescription: "EXEC mode password",
				Computed:            true,
			},
			"trustsec_exec_mode_username": schema.StringAttribute{
				MarkdownDescription: "EXEC mode username",
				Computed:            true,
			},
			"trustsec_include_when_deploying_sgt_updates": schema.BoolAttribute{
				MarkdownDescription: "Include this device when deploying Security Group Tag Mapping Updates",
				Computed:            true,
			},
			"trustsec_download_enviroment_data_every_x_seconds": schema.Int64Attribute{
				MarkdownDescription: "Download environment data every X seconds",
				Computed:            true,
			},
			"trustsec_download_peer_authorization_policy_every_x_seconds": schema.Int64Attribute{
				MarkdownDescription: "Download peer authorization policy every X seconds",
				Computed:            true,
			},
			"trustsec_download_sgacl_lists_every_x_seconds": schema.Int64Attribute{
				MarkdownDescription: "Download SGACL lists every X seconds",
				Computed:            true,
			},
			"trustsec_other_sga_devices_to_trust_this_device": schema.BoolAttribute{
				MarkdownDescription: "Other TrustSec devices to trust this device",
				Computed:            true,
			},
			"trustsec_re_authentication_every_x_seconds": schema.Int64Attribute{
				MarkdownDescription: "Re-authenticate every X seconds",
				Computed:            true,
			},
			"trustsec_send_configuration_to_device": schema.BoolAttribute{
				MarkdownDescription: "Send configuration to device",
				Computed:            true,
			},
			"trustsec_send_configuration_to_device_using": schema.StringAttribute{
				MarkdownDescription: "Send configuration to device using",
				Computed:            true,
			},
			"trustsec_coa_source_host": schema.StringAttribute{
				MarkdownDescription: "CoA source host",
				Computed:            true,
			},
		},
	}
}
func (d *NetworkDeviceDataSource) ConfigValidators(ctx context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{
		datasourcevalidator.ExactlyOneOf(
			path.MatchRoot("id"),
			path.MatchRoot("name"),
		),
	}
}

func (d *NetworkDeviceDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.client = req.ProviderData.(*IseProviderData).Client
}

//template:end model

//template:begin read
func (d *NetworkDeviceDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config NetworkDevice

	// Read config
	diags := req.Config.Get(ctx, &config)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Read", config.Id.ValueString()))
	if config.Id.IsNull() && !config.Name.IsNull() {
		for page := 1; ; page++ {
			res, err := d.client.Get(fmt.Sprintf("%s?size=100&page=%v", config.getPath(), page))
			if err != nil {
				resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve objects, got error: %s", err))
				return
			}
			if value := res.Get("SearchResult.resources"); len(value.Array()) > 0 {
				value.ForEach(func(k, v gjson.Result) bool {
					if config.Name.ValueString() == v.Get("name").String() {
						config.Id = types.StringValue(v.Get("id").String())
						tflog.Debug(ctx, fmt.Sprintf("%s: Found object with name '%v', id: %v", config.Id.String(), config.Name.ValueString(), config.Id.String()))
						return false
					}
					return true
				})
			}
			if !config.Id.IsNull() || !res.Get("SearchResult.nextPage").Exists() {
				break
			}
		}

		if config.Id.IsNull() {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to find object with name: %s", config.Name.ValueString()))
			return
		}
	}

	res, err := d.client.Get(config.getPath() + "/" + url.QueryEscape(config.Id.ValueString()))
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve object, got error: %s", err))
		return
	}

	config.fromBody(ctx, res)

	tflog.Debug(ctx, fmt.Sprintf("%s: Read finished successfully", config.Id.ValueString()))

	diags = resp.State.Set(ctx, &config)
	resp.Diagnostics.Append(diags...)
}

//template:end read
