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

	"github.com/CiscoDevNet/terraform-provider-ise/internal/provider/helpers"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

//template:end imports

//template:begin types
type Endpoint struct {
	Id                             types.String `tfsdk:"id"`
	Name                           types.String `tfsdk:"name"`
	Description                    types.String `tfsdk:"description"`
	Mac                            types.String `tfsdk:"mac"`
	GroupId                        types.String `tfsdk:"group_id"`
	ProfileId                      types.String `tfsdk:"profile_id"`
	StaticProfileAssignment        types.Bool   `tfsdk:"static_profile_assignment"`
	StaticProfileAssignmentDefined types.Bool   `tfsdk:"static_profile_assignment_defined"`
	StaticGroupAssignment          types.Bool   `tfsdk:"static_group_assignment"`
	StaticGroupAssignmentDefined   types.Bool   `tfsdk:"static_group_assignment_defined"`
	CustomAttributes               types.Map    `tfsdk:"custom_attributes"`
	IdentityStore                  types.String `tfsdk:"identity_store"`
	IdentityStoreId                types.String `tfsdk:"identity_store_id"`
	PortalUser                     types.String `tfsdk:"portal_user"`
	MdmServerName                  types.String `tfsdk:"mdm_server_name"`
	MdmReachable                   types.Bool   `tfsdk:"mdm_reachable"`
	MdmEnrolled                    types.Bool   `tfsdk:"mdm_enrolled"`
	MdmComplianceStatus            types.Bool   `tfsdk:"mdm_compliance_status"`
	MdmOs                          types.String `tfsdk:"mdm_os"`
	MdmManufacturer                types.String `tfsdk:"mdm_manufacturer"`
	MdmModel                       types.String `tfsdk:"mdm_model"`
	MdmSerial                      types.String `tfsdk:"mdm_serial"`
	MdmEncrypted                   types.Bool   `tfsdk:"mdm_encrypted"`
	MdmPinlock                     types.Bool   `tfsdk:"mdm_pinlock"`
	MdmJailBroken                  types.Bool   `tfsdk:"mdm_jail_broken"`
	MdmImei                        types.String `tfsdk:"mdm_imei"`
	MdmPhoneNumber                 types.String `tfsdk:"mdm_phone_number"`
}

//template:end types

//template:begin getPath
func (data Endpoint) getPath() string {
	return "/ers/config/endpoint"
}

//template:end getPath

//template:begin getPathDelete

//template:end getPathDelete

//template:begin toBody
func (data Endpoint) toBody(ctx context.Context, state Endpoint) string {
	body := ""
	if !data.Name.IsNull() {
		body, _ = sjson.Set(body, "ERSEndPoint.name", data.Name.ValueString())
	}
	if !data.Description.IsNull() {
		body, _ = sjson.Set(body, "ERSEndPoint.description", data.Description.ValueString())
	}
	if !data.Mac.IsNull() {
		body, _ = sjson.Set(body, "ERSEndPoint.mac", data.Mac.ValueString())
	}
	if !data.GroupId.IsNull() {
		body, _ = sjson.Set(body, "ERSEndPoint.groupId", data.GroupId.ValueString())
	}
	if !data.ProfileId.IsNull() {
		body, _ = sjson.Set(body, "ERSEndPoint.profileId", data.ProfileId.ValueString())
	}
	if !data.StaticProfileAssignment.IsNull() {
		body, _ = sjson.Set(body, "ERSEndPoint.staticProfileAssignment", data.StaticProfileAssignment.ValueBool())
	}
	if !data.StaticProfileAssignmentDefined.IsNull() {
		body, _ = sjson.Set(body, "ERSEndPoint.staticProfileAssignmentDefined", data.StaticProfileAssignmentDefined.ValueBool())
	}
	if !data.StaticGroupAssignment.IsNull() {
		body, _ = sjson.Set(body, "ERSEndPoint.staticGroupAssignment", data.StaticGroupAssignment.ValueBool())
	}
	if !data.StaticGroupAssignmentDefined.IsNull() {
		body, _ = sjson.Set(body, "ERSEndPoint.staticGroupAssignmentDefined", data.StaticGroupAssignmentDefined.ValueBool())
	}
	if !data.CustomAttributes.IsNull() {
		var values map[string]string
		data.CustomAttributes.ElementsAs(ctx, &values, false)
		body, _ = sjson.Set(body, "[ERSEndPoint customAttributes].customAttributes", values)
	}
	if !data.IdentityStore.IsNull() {
		body, _ = sjson.Set(body, "ERSEndPoint.identityStore", data.IdentityStore.ValueString())
	}
	if !data.IdentityStoreId.IsNull() {
		body, _ = sjson.Set(body, "ERSEndPoint.identityStoreId", data.IdentityStoreId.ValueString())
	}
	if !data.PortalUser.IsNull() {
		body, _ = sjson.Set(body, "ERSEndPoint.portalUser", data.PortalUser.ValueString())
	}
	if !data.MdmServerName.IsNull() {
		body, _ = sjson.Set(body, "ERSEndPoint.mdmAttributes.mdmServerName", data.MdmServerName.ValueString())
	}
	if !data.MdmReachable.IsNull() {
		body, _ = sjson.Set(body, "ERSEndPoint.mdmAttributes.mdmReachable", data.MdmReachable.ValueBool())
	}
	if !data.MdmEnrolled.IsNull() {
		body, _ = sjson.Set(body, "ERSEndPoint.mdmAttributes.mdmEnrolled", data.MdmEnrolled.ValueBool())
	}
	if !data.MdmComplianceStatus.IsNull() {
		body, _ = sjson.Set(body, "ERSEndPoint.mdmAttributes.mdmComplianceStatus", data.MdmComplianceStatus.ValueBool())
	}
	if !data.MdmOs.IsNull() {
		body, _ = sjson.Set(body, "ERSEndPoint.mdmAttributes.mdmOS", data.MdmOs.ValueString())
	}
	if !data.MdmManufacturer.IsNull() {
		body, _ = sjson.Set(body, "ERSEndPoint.mdmAttributes.mdmManufacturer", data.MdmManufacturer.ValueString())
	}
	if !data.MdmModel.IsNull() {
		body, _ = sjson.Set(body, "ERSEndPoint.mdmAttributes.mdmModel", data.MdmModel.ValueString())
	}
	if !data.MdmSerial.IsNull() {
		body, _ = sjson.Set(body, "ERSEndPoint.mdmAttributes.mdmSerial", data.MdmSerial.ValueString())
	}
	if !data.MdmEncrypted.IsNull() {
		body, _ = sjson.Set(body, "ERSEndPoint.mdmAttributes.mdmEncrypted", data.MdmEncrypted.ValueBool())
	}
	if !data.MdmPinlock.IsNull() {
		body, _ = sjson.Set(body, "ERSEndPoint.mdmAttributes.mdmPinlock", data.MdmPinlock.ValueBool())
	}
	if !data.MdmJailBroken.IsNull() {
		body, _ = sjson.Set(body, "ERSEndPoint.mdmAttributes.mdmJailBroken", data.MdmJailBroken.ValueBool())
	}
	if !data.MdmImei.IsNull() {
		body, _ = sjson.Set(body, "ERSEndPoint.mdmAttributes.mdmIMEI", data.MdmImei.ValueString())
	}
	if !data.MdmPhoneNumber.IsNull() {
		body, _ = sjson.Set(body, "ERSEndPoint.mdmAttributes.mdmPhoneNumber", data.MdmPhoneNumber.ValueString())
	}
	return body
}

//template:end toBody

//template:begin fromBody
func (data *Endpoint) fromBody(ctx context.Context, res gjson.Result) {
	if value := res.Get("ERSEndPoint.name"); value.Exists() {
		data.Name = types.StringValue(value.String())
	} else {
		data.Name = types.StringNull()
	}
	if value := res.Get("ERSEndPoint.description"); value.Exists() {
		data.Description = types.StringValue(value.String())
	} else {
		data.Description = types.StringNull()
	}
	if value := res.Get("ERSEndPoint.mac"); value.Exists() {
		data.Mac = types.StringValue(value.String())
	} else {
		data.Mac = types.StringNull()
	}
	if value := res.Get("ERSEndPoint.groupId"); value.Exists() {
		data.GroupId = types.StringValue(value.String())
	} else {
		data.GroupId = types.StringNull()
	}
	if value := res.Get("ERSEndPoint.profileId"); value.Exists() {
		data.ProfileId = types.StringValue(value.String())
	} else {
		data.ProfileId = types.StringNull()
	}
	if value := res.Get("ERSEndPoint.staticProfileAssignment"); value.Exists() {
		data.StaticProfileAssignment = types.BoolValue(value.Bool())
	} else {
		data.StaticProfileAssignment = types.BoolNull()
	}
	if value := res.Get("ERSEndPoint.staticProfileAssignmentDefined"); value.Exists() {
		data.StaticProfileAssignmentDefined = types.BoolValue(value.Bool())
	} else {
		data.StaticProfileAssignmentDefined = types.BoolValue(true)
	}
	if value := res.Get("ERSEndPoint.staticGroupAssignment"); value.Exists() {
		data.StaticGroupAssignment = types.BoolValue(value.Bool())
	} else {
		data.StaticGroupAssignment = types.BoolNull()
	}
	if value := res.Get("ERSEndPoint.staticGroupAssignmentDefined"); value.Exists() {
		data.StaticGroupAssignmentDefined = types.BoolValue(value.Bool())
	} else {
		data.StaticGroupAssignmentDefined = types.BoolValue(true)
	}
	if value := res.Get("[ERSEndPoint customAttributes].customAttributes"); value.Exists() {
		data.CustomAttributes = helpers.GetStringMap(value.Map())
	} else {
		data.CustomAttributes = types.MapNull(types.StringType)
	}
	if value := res.Get("ERSEndPoint.identityStore"); value.Exists() {
		data.IdentityStore = types.StringValue(value.String())
	} else {
		data.IdentityStore = types.StringNull()
	}
	if value := res.Get("ERSEndPoint.identityStoreId"); value.Exists() {
		data.IdentityStoreId = types.StringValue(value.String())
	} else {
		data.IdentityStoreId = types.StringNull()
	}
	if value := res.Get("ERSEndPoint.portalUser"); value.Exists() {
		data.PortalUser = types.StringValue(value.String())
	} else {
		data.PortalUser = types.StringNull()
	}
	if value := res.Get("ERSEndPoint.mdmAttributes.mdmServerName"); value.Exists() {
		data.MdmServerName = types.StringValue(value.String())
	} else {
		data.MdmServerName = types.StringNull()
	}
	if value := res.Get("ERSEndPoint.mdmAttributes.mdmReachable"); value.Exists() {
		data.MdmReachable = types.BoolValue(value.Bool())
	} else {
		data.MdmReachable = types.BoolNull()
	}
	if value := res.Get("ERSEndPoint.mdmAttributes.mdmEnrolled"); value.Exists() {
		data.MdmEnrolled = types.BoolValue(value.Bool())
	} else {
		data.MdmEnrolled = types.BoolNull()
	}
	if value := res.Get("ERSEndPoint.mdmAttributes.mdmComplianceStatus"); value.Exists() {
		data.MdmComplianceStatus = types.BoolValue(value.Bool())
	} else {
		data.MdmComplianceStatus = types.BoolNull()
	}
	if value := res.Get("ERSEndPoint.mdmAttributes.mdmOS"); value.Exists() {
		data.MdmOs = types.StringValue(value.String())
	} else {
		data.MdmOs = types.StringNull()
	}
	if value := res.Get("ERSEndPoint.mdmAttributes.mdmManufacturer"); value.Exists() {
		data.MdmManufacturer = types.StringValue(value.String())
	} else {
		data.MdmManufacturer = types.StringNull()
	}
	if value := res.Get("ERSEndPoint.mdmAttributes.mdmModel"); value.Exists() {
		data.MdmModel = types.StringValue(value.String())
	} else {
		data.MdmModel = types.StringNull()
	}
	if value := res.Get("ERSEndPoint.mdmAttributes.mdmSerial"); value.Exists() {
		data.MdmSerial = types.StringValue(value.String())
	} else {
		data.MdmSerial = types.StringNull()
	}
	if value := res.Get("ERSEndPoint.mdmAttributes.mdmEncrypted"); value.Exists() {
		data.MdmEncrypted = types.BoolValue(value.Bool())
	} else {
		data.MdmEncrypted = types.BoolNull()
	}
	if value := res.Get("ERSEndPoint.mdmAttributes.mdmPinlock"); value.Exists() {
		data.MdmPinlock = types.BoolValue(value.Bool())
	} else {
		data.MdmPinlock = types.BoolNull()
	}
	if value := res.Get("ERSEndPoint.mdmAttributes.mdmJailBroken"); value.Exists() {
		data.MdmJailBroken = types.BoolValue(value.Bool())
	} else {
		data.MdmJailBroken = types.BoolNull()
	}
	if value := res.Get("ERSEndPoint.mdmAttributes.mdmIMEI"); value.Exists() {
		data.MdmImei = types.StringValue(value.String())
	} else {
		data.MdmImei = types.StringNull()
	}
	if value := res.Get("ERSEndPoint.mdmAttributes.mdmPhoneNumber"); value.Exists() {
		data.MdmPhoneNumber = types.StringValue(value.String())
	} else {
		data.MdmPhoneNumber = types.StringNull()
	}
}

//template:end fromBody

//template:begin updateFromBody
func (data *Endpoint) updateFromBody(ctx context.Context, res gjson.Result) {
	if value := res.Get("ERSEndPoint.name"); value.Exists() && !data.Name.IsNull() {
		data.Name = types.StringValue(value.String())
	} else {
		data.Name = types.StringNull()
	}
	if value := res.Get("ERSEndPoint.description"); value.Exists() && !data.Description.IsNull() {
		data.Description = types.StringValue(value.String())
	} else {
		data.Description = types.StringNull()
	}
	if value := res.Get("ERSEndPoint.mac"); value.Exists() && !data.Mac.IsNull() {
		data.Mac = types.StringValue(value.String())
	} else {
		data.Mac = types.StringNull()
	}
	if value := res.Get("ERSEndPoint.groupId"); value.Exists() && !data.GroupId.IsNull() {
		data.GroupId = types.StringValue(value.String())
	} else {
		data.GroupId = types.StringNull()
	}
	if value := res.Get("ERSEndPoint.profileId"); value.Exists() && !data.ProfileId.IsNull() {
		data.ProfileId = types.StringValue(value.String())
	} else {
		data.ProfileId = types.StringNull()
	}
	if value := res.Get("ERSEndPoint.staticProfileAssignment"); value.Exists() && !data.StaticProfileAssignment.IsNull() {
		data.StaticProfileAssignment = types.BoolValue(value.Bool())
	} else {
		data.StaticProfileAssignment = types.BoolNull()
	}
	if value := res.Get("ERSEndPoint.staticProfileAssignmentDefined"); value.Exists() && !data.StaticProfileAssignmentDefined.IsNull() {
		data.StaticProfileAssignmentDefined = types.BoolValue(value.Bool())
	} else if data.StaticProfileAssignmentDefined.ValueBool() != true {
		data.StaticProfileAssignmentDefined = types.BoolNull()
	}
	if value := res.Get("ERSEndPoint.staticGroupAssignment"); value.Exists() && !data.StaticGroupAssignment.IsNull() {
		data.StaticGroupAssignment = types.BoolValue(value.Bool())
	} else {
		data.StaticGroupAssignment = types.BoolNull()
	}
	if value := res.Get("ERSEndPoint.staticGroupAssignmentDefined"); value.Exists() && !data.StaticGroupAssignmentDefined.IsNull() {
		data.StaticGroupAssignmentDefined = types.BoolValue(value.Bool())
	} else if data.StaticGroupAssignmentDefined.ValueBool() != true {
		data.StaticGroupAssignmentDefined = types.BoolNull()
	}
	if value := res.Get("[ERSEndPoint customAttributes].customAttributes"); value.Exists() && !data.CustomAttributes.IsNull() {
		data.CustomAttributes = helpers.GetStringMap(value.Map())
	} else {
		data.CustomAttributes = types.MapNull(types.StringType)
	}
	if value := res.Get("ERSEndPoint.identityStore"); value.Exists() && !data.IdentityStore.IsNull() {
		data.IdentityStore = types.StringValue(value.String())
	} else {
		data.IdentityStore = types.StringNull()
	}
	if value := res.Get("ERSEndPoint.identityStoreId"); value.Exists() && !data.IdentityStoreId.IsNull() {
		data.IdentityStoreId = types.StringValue(value.String())
	} else {
		data.IdentityStoreId = types.StringNull()
	}
	if value := res.Get("ERSEndPoint.portalUser"); value.Exists() && !data.PortalUser.IsNull() {
		data.PortalUser = types.StringValue(value.String())
	} else {
		data.PortalUser = types.StringNull()
	}
	if value := res.Get("ERSEndPoint.mdmAttributes.mdmServerName"); value.Exists() && !data.MdmServerName.IsNull() {
		data.MdmServerName = types.StringValue(value.String())
	} else {
		data.MdmServerName = types.StringNull()
	}
	if value := res.Get("ERSEndPoint.mdmAttributes.mdmReachable"); value.Exists() && !data.MdmReachable.IsNull() {
		data.MdmReachable = types.BoolValue(value.Bool())
	} else {
		data.MdmReachable = types.BoolNull()
	}
	if value := res.Get("ERSEndPoint.mdmAttributes.mdmEnrolled"); value.Exists() && !data.MdmEnrolled.IsNull() {
		data.MdmEnrolled = types.BoolValue(value.Bool())
	} else {
		data.MdmEnrolled = types.BoolNull()
	}
	if value := res.Get("ERSEndPoint.mdmAttributes.mdmComplianceStatus"); value.Exists() && !data.MdmComplianceStatus.IsNull() {
		data.MdmComplianceStatus = types.BoolValue(value.Bool())
	} else {
		data.MdmComplianceStatus = types.BoolNull()
	}
	if value := res.Get("ERSEndPoint.mdmAttributes.mdmOS"); value.Exists() && !data.MdmOs.IsNull() {
		data.MdmOs = types.StringValue(value.String())
	} else {
		data.MdmOs = types.StringNull()
	}
	if value := res.Get("ERSEndPoint.mdmAttributes.mdmManufacturer"); value.Exists() && !data.MdmManufacturer.IsNull() {
		data.MdmManufacturer = types.StringValue(value.String())
	} else {
		data.MdmManufacturer = types.StringNull()
	}
	if value := res.Get("ERSEndPoint.mdmAttributes.mdmModel"); value.Exists() && !data.MdmModel.IsNull() {
		data.MdmModel = types.StringValue(value.String())
	} else {
		data.MdmModel = types.StringNull()
	}
	if value := res.Get("ERSEndPoint.mdmAttributes.mdmSerial"); value.Exists() && !data.MdmSerial.IsNull() {
		data.MdmSerial = types.StringValue(value.String())
	} else {
		data.MdmSerial = types.StringNull()
	}
	if value := res.Get("ERSEndPoint.mdmAttributes.mdmEncrypted"); value.Exists() && !data.MdmEncrypted.IsNull() {
		data.MdmEncrypted = types.BoolValue(value.Bool())
	} else {
		data.MdmEncrypted = types.BoolNull()
	}
	if value := res.Get("ERSEndPoint.mdmAttributes.mdmPinlock"); value.Exists() && !data.MdmPinlock.IsNull() {
		data.MdmPinlock = types.BoolValue(value.Bool())
	} else {
		data.MdmPinlock = types.BoolNull()
	}
	if value := res.Get("ERSEndPoint.mdmAttributes.mdmJailBroken"); value.Exists() && !data.MdmJailBroken.IsNull() {
		data.MdmJailBroken = types.BoolValue(value.Bool())
	} else {
		data.MdmJailBroken = types.BoolNull()
	}
	if value := res.Get("ERSEndPoint.mdmAttributes.mdmIMEI"); value.Exists() && !data.MdmImei.IsNull() {
		data.MdmImei = types.StringValue(value.String())
	} else {
		data.MdmImei = types.StringNull()
	}
	if value := res.Get("ERSEndPoint.mdmAttributes.mdmPhoneNumber"); value.Exists() && !data.MdmPhoneNumber.IsNull() {
		data.MdmPhoneNumber = types.StringValue(value.String())
	} else {
		data.MdmPhoneNumber = types.StringNull()
	}
}

//template:end updateFromBody

//template:begin isNull
func (data *Endpoint) isNull(ctx context.Context, res gjson.Result) bool {
	if !data.Name.IsNull() {
		return false
	}
	if !data.Description.IsNull() {
		return false
	}
	if !data.Mac.IsNull() {
		return false
	}
	if !data.GroupId.IsNull() {
		return false
	}
	if !data.ProfileId.IsNull() {
		return false
	}
	if !data.StaticProfileAssignment.IsNull() {
		return false
	}
	if !data.StaticProfileAssignmentDefined.IsNull() {
		return false
	}
	if !data.StaticGroupAssignment.IsNull() {
		return false
	}
	if !data.StaticGroupAssignmentDefined.IsNull() {
		return false
	}
	if !data.CustomAttributes.IsNull() {
		return false
	}
	if !data.IdentityStore.IsNull() {
		return false
	}
	if !data.IdentityStoreId.IsNull() {
		return false
	}
	if !data.PortalUser.IsNull() {
		return false
	}
	if !data.MdmServerName.IsNull() {
		return false
	}
	if !data.MdmReachable.IsNull() {
		return false
	}
	if !data.MdmEnrolled.IsNull() {
		return false
	}
	if !data.MdmComplianceStatus.IsNull() {
		return false
	}
	if !data.MdmOs.IsNull() {
		return false
	}
	if !data.MdmManufacturer.IsNull() {
		return false
	}
	if !data.MdmModel.IsNull() {
		return false
	}
	if !data.MdmSerial.IsNull() {
		return false
	}
	if !data.MdmEncrypted.IsNull() {
		return false
	}
	if !data.MdmPinlock.IsNull() {
		return false
	}
	if !data.MdmJailBroken.IsNull() {
		return false
	}
	if !data.MdmImei.IsNull() {
		return false
	}
	if !data.MdmPhoneNumber.IsNull() {
		return false
	}
	return true
}

//template:end isNull
