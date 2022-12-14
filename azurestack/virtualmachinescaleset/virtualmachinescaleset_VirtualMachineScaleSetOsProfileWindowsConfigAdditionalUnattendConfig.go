package virtualmachinescaleset


type VirtualMachineScaleSetOsProfileWindowsConfigAdditionalUnattendConfig struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurestack/r/virtual_machine_scale_set#component VirtualMachineScaleSet#component}.
	Component *string `field:"required" json:"component" yaml:"component"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurestack/r/virtual_machine_scale_set#content VirtualMachineScaleSet#content}.
	Content *string `field:"required" json:"content" yaml:"content"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurestack/r/virtual_machine_scale_set#pass VirtualMachineScaleSet#pass}.
	Pass *string `field:"required" json:"pass" yaml:"pass"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurestack/r/virtual_machine_scale_set#setting_name VirtualMachineScaleSet#setting_name}.
	SettingName *string `field:"required" json:"settingName" yaml:"settingName"`
}

