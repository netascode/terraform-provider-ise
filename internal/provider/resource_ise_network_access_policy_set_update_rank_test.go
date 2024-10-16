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
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

//template:end imports

//template:begin testAcc
func TestAccIseNetworkAccessPolicySetUpdateRank(t *testing.T) {
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttr("ise_network_access_policy_set_update_rank.test", "rank", "0"))

	var steps []resource.TestStep
	steps = append(steps, resource.TestStep{
		Config: testAccIseNetworkAccessPolicySetUpdateRankPrerequisitesConfig + testAccIseNetworkAccessPolicySetUpdateRankConfig_all(),
		Check:  resource.ComposeTestCheckFunc(checks...),
	})
	steps = append(steps, resource.TestStep{
		ResourceName: "ise_network_access_policy_set_update_rank.test",
		ImportState:  true,
	})

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps:                    steps,
	})
}

//template:end testAcc

//template:begin testPrerequisites
const testAccIseNetworkAccessPolicySetUpdateRankPrerequisitesConfig = `
resource "ise_network_access_policy_set" "test" {
  name                      = "PolicySet1"
  service_name              = "Default Network Access"
  condition_type            = "ConditionAttributes"
  condition_is_negate       = false
  condition_attribute_name  = "Location"
  condition_attribute_value = "All Locations"
  condition_dictionary_name = "DEVICE"
  condition_operator        = "equals"
}

`

//template:end testPrerequisites

//template:begin testAccConfigMinimal
func testAccIseNetworkAccessPolicySetUpdateRankConfig_minimum() string {
	config := `resource "ise_network_access_policy_set_update_rank" "test" {` + "\n"
	config += `	policy_set_id = ise_network_access_policy_set.test.id` + "\n"
	config += `	rank = 0` + "\n"
	config += `}` + "\n"
	return config
}

//template:end testAccConfigMinimal

//template:begin testAccConfigAll
func testAccIseNetworkAccessPolicySetUpdateRankConfig_all() string {
	config := `resource "ise_network_access_policy_set_update_rank" "test" {` + "\n"
	config += `	policy_set_id = ise_network_access_policy_set.test.id` + "\n"
	config += `	rank = 0` + "\n"
	config += `}` + "\n"
	return config
}

//template:end testAccConfigAll
