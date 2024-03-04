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
	_ datasource.DataSource              = &DeviceAdminTimeAndDateConditionDataSource{}
	_ datasource.DataSourceWithConfigure = &DeviceAdminTimeAndDateConditionDataSource{}
)

func NewDeviceAdminTimeAndDateConditionDataSource() datasource.DataSource {
	return &DeviceAdminTimeAndDateConditionDataSource{}
}

type DeviceAdminTimeAndDateConditionDataSource struct {
	client *ise.Client
}

func (d *DeviceAdminTimeAndDateConditionDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_device_admin_time_and_date_condition"
}

func (d *DeviceAdminTimeAndDateConditionDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: "This data source can read the Device Admin Time And Date Condition.",

		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				MarkdownDescription: "The id of the object",
				Optional:            true,
				Computed:            true,
			},
			"name": schema.StringAttribute{
				MarkdownDescription: "Condition name",
				Optional:            true,
				Computed:            true,
			},
			"description": schema.StringAttribute{
				MarkdownDescription: "Condition description",
				Computed:            true,
			},
			"is_negate": schema.BoolAttribute{
				MarkdownDescription: "Indicates whereas this condition is in negate mode",
				Computed:            true,
			},
			"week_days": schema.SetAttribute{
				MarkdownDescription: "Defines for which days this condition will be matched. List of weekdays - `Monday`, `Tuesday`, `Wednesday`, `Thursday`, `Friday`, `Saturday`, `Sunday`. Default - List of all week days.",
				ElementType:         types.StringType,
				Computed:            true,
			},
			"week_days_exception": schema.SetAttribute{
				MarkdownDescription: "Defines for which days this condition will NOT be matched. List of weekdays - `Monday`, `Tuesday`, `Wednesday`, `Thursday`, `Friday`, `Saturday`, `Sunday`.",
				ElementType:         types.StringType,
				Computed:            true,
			},
			"start_date": schema.StringAttribute{
				MarkdownDescription: "Start date",
				Computed:            true,
			},
			"end_date": schema.StringAttribute{
				MarkdownDescription: "End date",
				Computed:            true,
			},
			"exception_start_date": schema.StringAttribute{
				MarkdownDescription: "Exception start date",
				Computed:            true,
			},
			"exception_end_date": schema.StringAttribute{
				MarkdownDescription: "Exception end date",
				Computed:            true,
			},
			"start_time": schema.StringAttribute{
				MarkdownDescription: "Start time",
				Computed:            true,
			},
			"end_time": schema.StringAttribute{
				MarkdownDescription: "End time",
				Computed:            true,
			},
			"exception_start_time": schema.StringAttribute{
				MarkdownDescription: "Exception start time",
				Computed:            true,
			},
			"exception_end_time": schema.StringAttribute{
				MarkdownDescription: "Exception end time",
				Computed:            true,
			},
		},
	}
}
func (d *DeviceAdminTimeAndDateConditionDataSource) ConfigValidators(ctx context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{
		datasourcevalidator.ExactlyOneOf(
			path.MatchRoot("id"),
			path.MatchRoot("name"),
		),
	}
}

func (d *DeviceAdminTimeAndDateConditionDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.client = req.ProviderData.(*IseProviderData).Client
}

//template:end model

//template:begin read
func (d *DeviceAdminTimeAndDateConditionDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config DeviceAdminTimeAndDateCondition

	// Read config
	diags := req.Config.Get(ctx, &config)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Read", config.Id.ValueString()))
	if config.Id.IsNull() && !config.Name.IsNull() {
		for page := 1; ; page++ {
			res, err := d.client.Get(config.getPath())
			if err != nil {
				resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve objects, got error: %s", err))
				return
			}
			if value := res.Get("response"); len(value.Array()) > 0 {
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
