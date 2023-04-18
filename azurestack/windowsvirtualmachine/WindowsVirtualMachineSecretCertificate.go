package windowsvirtualmachine


type WindowsVirtualMachineSecretCertificate struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurestack/1.0.0/docs/resources/windows_virtual_machine#store WindowsVirtualMachine#store}.
	Store *string `field:"required" json:"store" yaml:"store"`
}

