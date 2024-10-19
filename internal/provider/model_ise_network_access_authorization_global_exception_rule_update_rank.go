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
type NetworkAccessAuthorizationGlobalExceptionRuleUpdateRank struct {
	Id     types.String `tfsdk:"id"`
	RuleId types.String `tfsdk:"rule_id"`
	Rank   types.Int64  `tfsdk:"rank"`
}

//template:end types

//template:begin getPath
func (data NetworkAccessAuthorizationGlobalExceptionRuleUpdateRank) getPath() string {
	return "/api/v1/policy/network-access/policy-set/global-exception"
}

//template:end getPath

//template:begin getPathDelete

//template:end getPathDelete

//template:begin toBody
func (data NetworkAccessAuthorizationGlobalExceptionRuleUpdateRank) toBody(ctx context.Context, state NetworkAccessAuthorizationGlobalExceptionRuleUpdateRank) string {
	body := ""
	if !data.RuleId.IsNull() {
		body, _ = sjson.Set(body, "", data.RuleId.ValueString())
	}
	if !data.Rank.IsNull() {
		body, _ = sjson.Set(body, "rule.rank", data.Rank.ValueInt64())
	}
	return body
}

//template:end toBody

//template:begin fromBody
func (data *NetworkAccessAuthorizationGlobalExceptionRuleUpdateRank) fromBody(ctx context.Context, res gjson.Result) {
	if value := res.Get("response.rule.rank"); value.Exists() && value.Type != gjson.Null {
		data.Rank = types.Int64Value(value.Int())
	} else {
		data.Rank = types.Int64Null()
	}
}

//template:end fromBody

//template:begin updateFromBody
func (data *NetworkAccessAuthorizationGlobalExceptionRuleUpdateRank) updateFromBody(ctx context.Context, res gjson.Result) {
	if value := res.Get("response.rule.rank"); value.Exists() && !data.Rank.IsNull() {
		data.Rank = types.Int64Value(value.Int())
	} else {
		data.Rank = types.Int64Null()
	}
}

//template:end updateFromBody

//template:begin isNull
func (data *NetworkAccessAuthorizationGlobalExceptionRuleUpdateRank) isNull(ctx context.Context, res gjson.Result) bool {
	if !data.RuleId.IsNull() {
		return false
	}
	if !data.Rank.IsNull() {
		return false
	}
	return true
}

//template:end isNull
