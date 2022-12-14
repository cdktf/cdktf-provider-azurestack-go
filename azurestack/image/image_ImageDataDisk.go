package image


type ImageDataDisk struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurestack/r/image#blob_uri Image#blob_uri}.
	BlobUri *string `field:"optional" json:"blobUri" yaml:"blobUri"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurestack/r/image#caching Image#caching}.
	Caching *string `field:"optional" json:"caching" yaml:"caching"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurestack/r/image#lun Image#lun}.
	Lun *float64 `field:"optional" json:"lun" yaml:"lun"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurestack/r/image#managed_disk_id Image#managed_disk_id}.
	ManagedDiskId *string `field:"optional" json:"managedDiskId" yaml:"managedDiskId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurestack/r/image#size_gb Image#size_gb}.
	SizeGb *float64 `field:"optional" json:"sizeGb" yaml:"sizeGb"`
}

