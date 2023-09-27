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

	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

type InternalUser struct {
	Id                   types.String `tfsdk:"id"`
	Name                 types.String `tfsdk:"name"`
	Password             types.String `tfsdk:"password"`
	ChangePassword       types.Bool   `tfsdk:"change_password"`
	Email                types.String `tfsdk:"email"`
	AccountNameAlias     types.String `tfsdk:"account_name_alias"`
	EnablePassword       types.String `tfsdk:"enable_password"`
	Enabled              types.Bool   `tfsdk:"enabled"`
	PasswordNeverExpires types.Bool   `tfsdk:"password_never_expires"`
	FirstName            types.String `tfsdk:"first_name"`
	LastName             types.String `tfsdk:"last_name"`
	PasswordIdStore      types.String `tfsdk:"password_id_store"`
	Description          types.String `tfsdk:"description"`
}

func (data InternalUser) toBody(ctx context.Context, state InternalUser) string {
	body := ""
	if !data.Name.IsNull() {
		body, _ = sjson.Set(body, "InternalUser.name", data.Name.ValueString())
	}
	if !data.Password.IsNull() && data.Password != state.Password {
		body, _ = sjson.Set(body, "InternalUser.password", data.Password.ValueString())
	}
	if !data.ChangePassword.IsNull() {
		body, _ = sjson.Set(body, "InternalUser.changePassword", data.ChangePassword.ValueBool())
	}
	if !data.Email.IsNull() {
		body, _ = sjson.Set(body, "InternalUser.email", data.Email.ValueString())
	}
	if !data.AccountNameAlias.IsNull() {
		body, _ = sjson.Set(body, "InternalUser.accountNameAlias", data.AccountNameAlias.ValueString())
	}
	if !data.EnablePassword.IsNull() {
		body, _ = sjson.Set(body, "InternalUser.enablePassword", data.EnablePassword.ValueString())
	}
	if !data.Enabled.IsNull() {
		body, _ = sjson.Set(body, "InternalUser.enabled", data.Enabled.ValueBool())
	}
	if !data.PasswordNeverExpires.IsNull() {
		body, _ = sjson.Set(body, "InternalUser.passwordNeverExpires", data.PasswordNeverExpires.ValueBool())
	}
	if !data.FirstName.IsNull() {
		body, _ = sjson.Set(body, "InternalUser.firstName", data.FirstName.ValueString())
	}
	if !data.LastName.IsNull() {
		body, _ = sjson.Set(body, "InternalUser.lastName", data.LastName.ValueString())
	}
	if !data.PasswordIdStore.IsNull() {
		body, _ = sjson.Set(body, "InternalUser.passwordIDStore", data.PasswordIdStore.ValueString())
	}
	if !data.Description.IsNull() {
		body, _ = sjson.Set(body, "InternalUser.description", data.Description.ValueString())
	}
	return body
}

func (data *InternalUser) fromBody(ctx context.Context, res gjson.Result) {
	if value := res.Get("InternalUser.name"); value.Exists() {
		data.Name = types.StringValue(value.String())
	} else {
		data.Name = types.StringNull()
	}
	if value := res.Get("InternalUser.changePassword"); value.Exists() {
		data.ChangePassword = types.BoolValue(value.Bool())
	} else {
		data.ChangePassword = types.BoolNull()
	}
	if value := res.Get("InternalUser.email"); value.Exists() {
		data.Email = types.StringValue(value.String())
	} else {
		data.Email = types.StringNull()
	}
	if value := res.Get("InternalUser.accountNameAlias"); value.Exists() {
		data.AccountNameAlias = types.StringValue(value.String())
	} else {
		data.AccountNameAlias = types.StringNull()
	}
	if value := res.Get("InternalUser.enabled"); value.Exists() {
		data.Enabled = types.BoolValue(value.Bool())
	} else {
		data.Enabled = types.BoolNull()
	}
	if value := res.Get("InternalUser.passwordNeverExpires"); value.Exists() {
		data.PasswordNeverExpires = types.BoolValue(value.Bool())
	} else {
		data.PasswordNeverExpires = types.BoolNull()
	}
	if value := res.Get("InternalUser.firstName"); value.Exists() {
		data.FirstName = types.StringValue(value.String())
	} else {
		data.FirstName = types.StringNull()
	}
	if value := res.Get("InternalUser.lastName"); value.Exists() {
		data.LastName = types.StringValue(value.String())
	} else {
		data.LastName = types.StringNull()
	}
	if value := res.Get("InternalUser.passwordIDStore"); value.Exists() {
		data.PasswordIdStore = types.StringValue(value.String())
	} else {
		data.PasswordIdStore = types.StringNull()
	}
	if value := res.Get("InternalUser.description"); value.Exists() {
		data.Description = types.StringValue(value.String())
	} else {
		data.Description = types.StringNull()
	}
}

func (data *InternalUser) updateFromBody(ctx context.Context, res gjson.Result) {
	if value := res.Get("InternalUser.name"); value.Exists() && !data.Name.IsNull() {
		data.Name = types.StringValue(value.String())
	} else {
		data.Name = types.StringNull()
	}
	if value := res.Get("InternalUser.changePassword"); value.Exists() && !data.ChangePassword.IsNull() {
		data.ChangePassword = types.BoolValue(value.Bool())
	} else {
		data.ChangePassword = types.BoolNull()
	}
	if value := res.Get("InternalUser.email"); value.Exists() && !data.Email.IsNull() {
		data.Email = types.StringValue(value.String())
	} else {
		data.Email = types.StringNull()
	}
	if value := res.Get("InternalUser.accountNameAlias"); value.Exists() && !data.AccountNameAlias.IsNull() {
		data.AccountNameAlias = types.StringValue(value.String())
	} else {
		data.AccountNameAlias = types.StringNull()
	}
	if value := res.Get("InternalUser.enabled"); value.Exists() && !data.Enabled.IsNull() {
		data.Enabled = types.BoolValue(value.Bool())
	} else {
		data.Enabled = types.BoolNull()
	}
	if value := res.Get("InternalUser.passwordNeverExpires"); value.Exists() && !data.PasswordNeverExpires.IsNull() {
		data.PasswordNeverExpires = types.BoolValue(value.Bool())
	} else {
		data.PasswordNeverExpires = types.BoolNull()
	}
	if value := res.Get("InternalUser.firstName"); value.Exists() && !data.FirstName.IsNull() {
		data.FirstName = types.StringValue(value.String())
	} else {
		data.FirstName = types.StringNull()
	}
	if value := res.Get("InternalUser.lastName"); value.Exists() && !data.LastName.IsNull() {
		data.LastName = types.StringValue(value.String())
	} else {
		data.LastName = types.StringNull()
	}
	if value := res.Get("InternalUser.passwordIDStore"); value.Exists() && !data.PasswordIdStore.IsNull() {
		data.PasswordIdStore = types.StringValue(value.String())
	} else {
		data.PasswordIdStore = types.StringNull()
	}
	if value := res.Get("InternalUser.description"); value.Exists() && !data.Description.IsNull() {
		data.Description = types.StringValue(value.String())
	} else {
		data.Description = types.StringNull()
	}
}
