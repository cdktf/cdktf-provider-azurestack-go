// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package windowsvirtualmachine


type WindowsVirtualMachineBootDiagnostics struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurestack/1.0.0/docs/resources/windows_virtual_machine#storage_account_uri WindowsVirtualMachine#storage_account_uri}.
	StorageAccountUri *string `field:"required" json:"storageAccountUri" yaml:"storageAccountUri"`
}

