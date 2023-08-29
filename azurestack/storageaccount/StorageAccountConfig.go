// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package storageaccount

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type StorageAccountConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktf.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurestack/1.0.0/docs/resources/storage_account#account_replication_type StorageAccount#account_replication_type}.
	AccountReplicationType *string `field:"required" json:"accountReplicationType" yaml:"accountReplicationType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurestack/1.0.0/docs/resources/storage_account#account_tier StorageAccount#account_tier}.
	AccountTier *string `field:"required" json:"accountTier" yaml:"accountTier"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurestack/1.0.0/docs/resources/storage_account#location StorageAccount#location}.
	Location *string `field:"required" json:"location" yaml:"location"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurestack/1.0.0/docs/resources/storage_account#name StorageAccount#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurestack/1.0.0/docs/resources/storage_account#resource_group_name StorageAccount#resource_group_name}.
	ResourceGroupName *string `field:"required" json:"resourceGroupName" yaml:"resourceGroupName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurestack/1.0.0/docs/resources/storage_account#account_encryption_source StorageAccount#account_encryption_source}.
	AccountEncryptionSource *string `field:"optional" json:"accountEncryptionSource" yaml:"accountEncryptionSource"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurestack/1.0.0/docs/resources/storage_account#account_kind StorageAccount#account_kind}.
	AccountKind *string `field:"optional" json:"accountKind" yaml:"accountKind"`
	// custom_domain block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurestack/1.0.0/docs/resources/storage_account#custom_domain StorageAccount#custom_domain}
	CustomDomain *StorageAccountCustomDomain `field:"optional" json:"customDomain" yaml:"customDomain"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurestack/1.0.0/docs/resources/storage_account#enable_blob_encryption StorageAccount#enable_blob_encryption}.
	EnableBlobEncryption interface{} `field:"optional" json:"enableBlobEncryption" yaml:"enableBlobEncryption"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurestack/1.0.0/docs/resources/storage_account#enable_https_traffic_only StorageAccount#enable_https_traffic_only}.
	EnableHttpsTrafficOnly interface{} `field:"optional" json:"enableHttpsTrafficOnly" yaml:"enableHttpsTrafficOnly"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurestack/1.0.0/docs/resources/storage_account#id StorageAccount#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurestack/1.0.0/docs/resources/storage_account#tags StorageAccount#tags}.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurestack/1.0.0/docs/resources/storage_account#timeouts StorageAccount#timeouts}
	Timeouts *StorageAccountTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

