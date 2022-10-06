package linuxvirtualmachine


type LinuxVirtualMachineBootDiagnostics struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurestack/r/linux_virtual_machine#storage_account_uri LinuxVirtualMachine#storage_account_uri}.
	StorageAccountUri *string `field:"required" json:"storageAccountUri" yaml:"storageAccountUri"`
}

