package windowsvirtualmachinescaleset


type WindowsVirtualMachineScaleSetTerminateNotification struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurestack/1.0.0/docs/resources/windows_virtual_machine_scale_set#enabled WindowsVirtualMachineScaleSet#enabled}.
	Enabled interface{} `field:"required" json:"enabled" yaml:"enabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurestack/1.0.0/docs/resources/windows_virtual_machine_scale_set#timeout WindowsVirtualMachineScaleSet#timeout}.
	Timeout *string `field:"optional" json:"timeout" yaml:"timeout"`
}

