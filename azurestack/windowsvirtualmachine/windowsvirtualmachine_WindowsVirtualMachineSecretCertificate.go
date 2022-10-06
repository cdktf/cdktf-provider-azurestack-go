package windowsvirtualmachine


type WindowsVirtualMachineSecretCertificate struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurestack/r/windows_virtual_machine#store WindowsVirtualMachine#store}.
	Store *string `field:"required" json:"store" yaml:"store"`
}

