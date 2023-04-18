package provider


type AzurestackProviderFeaturesVirtualMachine struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurestack/1.0.0/docs#delete_os_disk_on_deletion AzurestackProvider#delete_os_disk_on_deletion}.
	DeleteOsDiskOnDeletion interface{} `field:"optional" json:"deleteOsDiskOnDeletion" yaml:"deleteOsDiskOnDeletion"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurestack/1.0.0/docs#graceful_shutdown AzurestackProvider#graceful_shutdown}.
	GracefulShutdown interface{} `field:"optional" json:"gracefulShutdown" yaml:"gracefulShutdown"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurestack/1.0.0/docs#skip_shutdown_and_force_delete AzurestackProvider#skip_shutdown_and_force_delete}.
	SkipShutdownAndForceDelete interface{} `field:"optional" json:"skipShutdownAndForceDelete" yaml:"skipShutdownAndForceDelete"`
}

