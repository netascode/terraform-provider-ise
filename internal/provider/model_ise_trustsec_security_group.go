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

	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

//template:end imports

//template:begin types
type TrustSecSecurityGroup struct {
	Id              types.String `tfsdk:"id"`
	Name            types.String `tfsdk:"name"`
	Description     types.String `tfsdk:"description"`
	Value           types.Int64  `tfsdk:"value"`
	PropogateToApic types.Bool   `tfsdk:"propogate_to_apic"`
	IsReadOnly      types.Bool   `tfsdk:"is_read_only"`
}

//template:end types

//template:begin getPath
func (data TrustSecSecurityGroup) getPath() string {
	return "/ers/config/sgt"
}

//template:end getPath

//template:begin getPathDelete

//template:end getPathDelete

//template:begin toBody
func (data TrustSecSecurityGroup) toBody(ctx context.Context, state TrustSecSecurityGroup) string {
	body := ""
	if !data.Name.IsNull() {
		body, _ = sjson.Set(body, "Sgt.name", data.Name.ValueString())
	}
	if !data.Description.IsNull() {
		body, _ = sjson.Set(body, "Sgt.description", data.Description.ValueString())
	}
	if !data.Value.IsNull() {
		body, _ = sjson.Set(body, "Sgt.value", data.Value.ValueInt64())
	}
	if !data.PropogateToApic.IsNull() {
		body, _ = sjson.Set(body, "Sgt.propogateToApic", data.PropogateToApic.ValueBool())
	}
	if !data.IsReadOnly.IsNull() {
		body, _ = sjson.Set(body, "Sgt.isReadOnly", data.IsReadOnly.ValueBool())
	}
	return body
}

//template:end toBody

//template:begin fromBody
func (data *TrustSecSecurityGroup) fromBody(ctx context.Context, res gjson.Result) {
	if value := res.Get("Sgt.name"); value.Exists() && value.Type != gjson.Null {
		data.Name = types.StringValue(value.String())
	} else {
		data.Name = types.StringNull()
	}
	if value := res.Get("Sgt.description"); value.Exists() && value.Type != gjson.Null {
		data.Description = types.StringValue(value.String())
	} else {
		data.Description = types.StringNull()
	}
	if value := res.Get("Sgt.value"); value.Exists() && value.Type != gjson.Null {
		data.Value = types.Int64Value(value.Int())
	} else {
		data.Value = types.Int64Null()
	}
	if value := res.Get("Sgt.propogateToApic"); value.Exists() && value.Type != gjson.Null {
		data.PropogateToApic = types.BoolValue(value.Bool())
	} else {
		data.PropogateToApic = types.BoolNull()
	}
}

//template:end fromBody

//template:begin updateFromBody
func (data *TrustSecSecurityGroup) updateFromBody(ctx context.Context, res gjson.Result) {
	if value := res.Get("Sgt.name"); value.Exists() && !data.Name.IsNull() {
		data.Name = types.StringValue(value.String())
	} else {
		data.Name = types.StringNull()
	}
	if value := res.Get("Sgt.description"); value.Exists() && !data.Description.IsNull() {
		data.Description = types.StringValue(value.String())
	} else {
		data.Description = types.StringNull()
	}
	if value := res.Get("Sgt.value"); value.Exists() && !data.Value.IsNull() {
		data.Value = types.Int64Value(value.Int())
	} else {
		data.Value = types.Int64Null()
	}
	if value := res.Get("Sgt.propogateToApic"); value.Exists() && !data.PropogateToApic.IsNull() {
		data.PropogateToApic = types.BoolValue(value.Bool())
	} else {
		data.PropogateToApic = types.BoolNull()
	}
}

//template:end updateFromBody

//template:begin isNull
func (data *TrustSecSecurityGroup) isNull(ctx context.Context, res gjson.Result) bool {
	if !data.Name.IsNull() {
		return false
	}
	if !data.Description.IsNull() {
		return false
	}
	if !data.Value.IsNull() {
		return false
	}
	if !data.PropogateToApic.IsNull() {
		return false
	}
	if !data.IsReadOnly.IsNull() {
		return false
	}
	return true
}

//template:end isNull
