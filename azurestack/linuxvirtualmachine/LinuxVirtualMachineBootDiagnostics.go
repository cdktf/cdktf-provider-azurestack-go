package linuxvirtualmachine


type LinuxVirtualMachineBootDiagnostics struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurestack/1.0.0/docs/resources/linux_virtual_machine#storage_account_uri LinuxVirtualMachine#storage_account_uri}.
	StorageAccountUri *string `field:"required" json:"storageAccountUri" yaml:"storageAccountUri"`
}

