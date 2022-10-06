package windowsvirtualmachine


type WindowsVirtualMachineBootDiagnostics struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurestack/r/windows_virtual_machine#storage_account_uri WindowsVirtualMachine#storage_account_uri}.
	StorageAccountUri *string `field:"required" json:"storageAccountUri" yaml:"storageAccountUri"`
}

