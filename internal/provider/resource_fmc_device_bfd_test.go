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

package provider

// Section below is generated&owned by "gen/generator.go". //template:begin imports
import (
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin testAcc

func TestAccFmcDeviceBFD(t *testing.T) {
	if os.Getenv("TF_VAR_device_id") == "" {
		t.Skip("skipping test, set environment variable TF_VAR_device_id")
	}
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttrSet("fmc_device_bfd.test", "type"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_device_bfd.test", "hop_type", "SINGLE_HOP"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_device_bfd.test", "bfd_template_id", "76d24097-41c4-4558-a4d0-a8c07ac08470"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_device_bfd.test", "interface_logical_name", "outside"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_device_bfd.test", "destination_host_object_id", "76d24097-41c4-4558-a4d0-a8c07ac08470"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_device_bfd.test", "source_host_object_id", "76d24097-41c4-4558-a4d0-a8c07ac08470"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_device_bfd.test", "interface_id", "76d24097-41c4-4558-a4d0-a8c07ac08470"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_device_bfd.test", "slow_timer", ""))

	var steps []resource.TestStep
	if os.Getenv("SKIP_MINIMUM_TEST") == "" {
		steps = append(steps, resource.TestStep{
			Config: testAccFmcDeviceBFDPrerequisitesConfig + testAccFmcDeviceBFDConfig_minimum(),
		})
	}
	steps = append(steps, resource.TestStep{
		Config: testAccFmcDeviceBFDPrerequisitesConfig + testAccFmcDeviceBFDConfig_all(),
		Check:  resource.ComposeTestCheckFunc(checks...),
	})

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		ErrorCheck:               func(err error) error { return testAccErrorCheck(t, err) },
		Steps:                    steps,
	})
}

// End of section. //template:end testAcc

// Section below is generated&owned by "gen/generator.go". //template:begin testPrerequisites

const testAccFmcDeviceBFDPrerequisitesConfig = `
variable "device_id" { default = null } // tests will set $TF_VAR_device_id
`

// End of section. //template:end testPrerequisites

// Section below is generated&owned by "gen/generator.go". //template:begin testAccConfigMinimal

func testAccFmcDeviceBFDConfig_minimum() string {
	config := `resource "fmc_device_bfd" "test" {` + "\n"
	config += `	device_id = var.device_id` + "\n"
	config += `	hop_type = "SINGLE_HOP"` + "\n"
	config += `	bfd_template_id = "76d24097-41c4-4558-a4d0-a8c07ac08470"` + "\n"
	config += `}` + "\n"
	return config
}

// End of section. //template:end testAccConfigMinimal

// Section below is generated&owned by "gen/generator.go". //template:begin testAccConfigAll

func testAccFmcDeviceBFDConfig_all() string {
	config := `resource "fmc_device_bfd" "test" {` + "\n"
	config += `	device_id = var.device_id` + "\n"
	config += `	hop_type = "SINGLE_HOP"` + "\n"
	config += `	bfd_template_id = "76d24097-41c4-4558-a4d0-a8c07ac08470"` + "\n"
	config += `	interface_logical_name = "outside"` + "\n"
	config += `	destination_host_object_id = "76d24097-41c4-4558-a4d0-a8c07ac08470"` + "\n"
	config += `	source_host_object_id = "76d24097-41c4-4558-a4d0-a8c07ac08470"` + "\n"
	config += `	interface_id = "76d24097-41c4-4558-a4d0-a8c07ac08470"` + "\n"
	config += `	slow_timer = ` + "\n"
	config += `}` + "\n"
	return config
}

// End of section. //template:end testAccConfigAll