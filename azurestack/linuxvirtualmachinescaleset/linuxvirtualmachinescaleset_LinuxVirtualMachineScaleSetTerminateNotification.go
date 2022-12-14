package linuxvirtualmachinescaleset


type LinuxVirtualMachineScaleSetTerminateNotification struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurestack/r/linux_virtual_machine_scale_set#enabled LinuxVirtualMachineScaleSet#enabled}.
	Enabled interface{} `field:"required" json:"enabled" yaml:"enabled"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurestack/r/linux_virtual_machine_scale_set#timeout LinuxVirtualMachineScaleSet#timeout}.
	Timeout *string `field:"optional" json:"timeout" yaml:"timeout"`
}

