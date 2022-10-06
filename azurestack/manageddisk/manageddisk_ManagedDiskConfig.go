package manageddisk

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ManagedDiskConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count *float64 `field:"optional" json:"count" yaml:"count"`
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
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurestack/r/managed_disk#create_option ManagedDisk#create_option}.
	CreateOption *string `field:"required" json:"createOption" yaml:"createOption"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurestack/r/managed_disk#location ManagedDisk#location}.
	Location *string `field:"required" json:"location" yaml:"location"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurestack/r/managed_disk#name ManagedDisk#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurestack/r/managed_disk#resource_group_name ManagedDisk#resource_group_name}.
	ResourceGroupName *string `field:"required" json:"resourceGroupName" yaml:"resourceGroupName"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurestack/r/managed_disk#storage_account_type ManagedDisk#storage_account_type}.
	StorageAccountType *string `field:"required" json:"storageAccountType" yaml:"storageAccountType"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurestack/r/managed_disk#disk_size_gb ManagedDisk#disk_size_gb}.
	DiskSizeGb *float64 `field:"optional" json:"diskSizeGb" yaml:"diskSizeGb"`
	// encryption block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurestack/r/managed_disk#encryption ManagedDisk#encryption}
	Encryption *ManagedDiskEncryption `field:"optional" json:"encryption" yaml:"encryption"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurestack/r/managed_disk#hyper_v_generation ManagedDisk#hyper_v_generation}.
	HyperVGeneration *string `field:"optional" json:"hyperVGeneration" yaml:"hyperVGeneration"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurestack/r/managed_disk#id ManagedDisk#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurestack/r/managed_disk#image_reference_id ManagedDisk#image_reference_id}.
	ImageReferenceId *string `field:"optional" json:"imageReferenceId" yaml:"imageReferenceId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurestack/r/managed_disk#os_type ManagedDisk#os_type}.
	OsType *string `field:"optional" json:"osType" yaml:"osType"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurestack/r/managed_disk#source_resource_id ManagedDisk#source_resource_id}.
	SourceResourceId *string `field:"optional" json:"sourceResourceId" yaml:"sourceResourceId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurestack/r/managed_disk#source_uri ManagedDisk#source_uri}.
	SourceUri *string `field:"optional" json:"sourceUri" yaml:"sourceUri"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurestack/r/managed_disk#storage_account_id ManagedDisk#storage_account_id}.
	StorageAccountId *string `field:"optional" json:"storageAccountId" yaml:"storageAccountId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurestack/r/managed_disk#tags ManagedDisk#tags}.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurestack/r/managed_disk#timeouts ManagedDisk#timeouts}
	Timeouts *ManagedDiskTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

