// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package windowsvirtualmachine


type WindowsVirtualMachineWinrmListener struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurestack/1.0.0/docs/resources/windows_virtual_machine#protocol WindowsVirtualMachine#protocol}.
	Protocol *string `field:"required" json:"protocol" yaml:"protocol"`
}

