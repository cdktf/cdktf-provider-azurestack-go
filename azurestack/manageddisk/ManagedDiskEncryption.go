// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package manageddisk


type ManagedDiskEncryption struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurestack/1.0.0/docs/resources/managed_disk#enabled ManagedDisk#enabled}.
	Enabled interface{} `field:"required" json:"enabled" yaml:"enabled"`
	// disk_encryption_key block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurestack/1.0.0/docs/resources/managed_disk#disk_encryption_key ManagedDisk#disk_encryption_key}
	DiskEncryptionKey *ManagedDiskEncryptionDiskEncryptionKey `field:"optional" json:"diskEncryptionKey" yaml:"diskEncryptionKey"`
	// key_encryption_key block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurestack/1.0.0/docs/resources/managed_disk#key_encryption_key ManagedDisk#key_encryption_key}
	KeyEncryptionKey *ManagedDiskEncryptionKeyEncryptionKey `field:"optional" json:"keyEncryptionKey" yaml:"keyEncryptionKey"`
}

