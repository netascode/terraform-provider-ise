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

//template:begin testAccDataSource
func TestAccDataSourceIseEndpoint(t *testing.T) {
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttr("data.ise_endpoint.test", "name", "00:11:22:33:44:55"))
	checks = append(checks, resource.TestCheckResourceAttr("data.ise_endpoint.test", "description", "My endpoint"))
	checks = append(checks, resource.TestCheckResourceAttr("data.ise_endpoint.test", "mac", "00:11:22:33:44:55"))
	checks = append(checks, resource.TestCheckResourceAttr("data.ise_endpoint.test", "group_id", "3a88eec0-8c00-11e6-996c-525400b48521"))
	checks = append(checks, resource.TestCheckResourceAttr("data.ise_endpoint.test", "profile_id", "3a91a150-8c00-11e6-996c-525400b48521"))
	checks = append(checks, resource.TestCheckResourceAttr("data.ise_endpoint.test", "static_profile_assignment", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("data.ise_endpoint.test", "static_profile_assignment_defined", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("data.ise_endpoint.test", "static_group_assignment", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("data.ise_endpoint.test", "static_group_assignment_defined", "true"))
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceIseEndpointConfig(),
				Check:  resource.ComposeTestCheckFunc(checks...),
			},
		},
	})
}

//template:end testAccDataSource

//template:begin testPrerequisites
//template:end testPrerequisites

//template:begin testAccDataSourceConfig
func testAccDataSourceIseEndpointConfig() string {
	config := `resource "ise_endpoint" "test" {` + "\n"
	config += `	name = "00:11:22:33:44:55"` + "\n"
	config += `	description = "My endpoint"` + "\n"
	config += `	mac = "00:11:22:33:44:55"` + "\n"
	config += `	group_id = "3a88eec0-8c00-11e6-996c-525400b48521"` + "\n"
	config += `	profile_id = "3a91a150-8c00-11e6-996c-525400b48521"` + "\n"
	config += `	static_profile_assignment = true` + "\n"
	config += `	static_profile_assignment_defined = true` + "\n"
	config += `	static_group_assignment = true` + "\n"
	config += `	static_group_assignment_defined = true` + "\n"
	config += `}` + "\n"

	config += `
		data "ise_endpoint" "test" {
			id = ise_endpoint.test.id
		}
	`
	return config
}

//template:end testAccDataSourceConfig
