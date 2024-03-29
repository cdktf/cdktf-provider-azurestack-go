// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package windowsvirtualmachine


type WindowsVirtualMachineAdditionalCapabilities struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurestack/1.0.0/docs/resources/windows_virtual_machine#ultra_ssd_enabled WindowsVirtualMachine#ultra_ssd_enabled}.
	UltraSsdEnabled interface{} `field:"optional" json:"ultraSsdEnabled" yaml:"ultraSsdEnabled"`
}

