package storageaccount


type StorageAccountTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurestack/r/storage_account#create StorageAccount#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurestack/r/storage_account#delete StorageAccount#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurestack/r/storage_account#read StorageAccount#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurestack/r/storage_account#update StorageAccount#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

