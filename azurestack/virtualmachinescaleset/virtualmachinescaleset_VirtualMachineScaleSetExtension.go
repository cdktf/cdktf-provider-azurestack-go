package virtualmachinescaleset


type VirtualMachineScaleSetExtension struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurestack/r/virtual_machine_scale_set#name VirtualMachineScaleSet#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurestack/r/virtual_machine_scale_set#publisher VirtualMachineScaleSet#publisher}.
	Publisher *string `field:"required" json:"publisher" yaml:"publisher"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurestack/r/virtual_machine_scale_set#type VirtualMachineScaleSet#type}.
	Type *string `field:"required" json:"type" yaml:"type"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurestack/r/virtual_machine_scale_set#type_handler_version VirtualMachineScaleSet#type_handler_version}.
	TypeHandlerVersion *string `field:"required" json:"typeHandlerVersion" yaml:"typeHandlerVersion"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurestack/r/virtual_machine_scale_set#auto_upgrade_minor_version VirtualMachineScaleSet#auto_upgrade_minor_version}.
	AutoUpgradeMinorVersion interface{} `field:"optional" json:"autoUpgradeMinorVersion" yaml:"autoUpgradeMinorVersion"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurestack/r/virtual_machine_scale_set#protected_settings VirtualMachineScaleSet#protected_settings}.
	ProtectedSettings *string `field:"optional" json:"protectedSettings" yaml:"protectedSettings"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurestack/r/virtual_machine_scale_set#settings VirtualMachineScaleSet#settings}.
	Settings *string `field:"optional" json:"settings" yaml:"settings"`
}

