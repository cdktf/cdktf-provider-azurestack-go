package windowsvirtualmachinescaleset


type WindowsVirtualMachineScaleSetAdditionalUnattendContent struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurestack/r/windows_virtual_machine_scale_set#content WindowsVirtualMachineScaleSet#content}.
	Content *string `field:"required" json:"content" yaml:"content"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurestack/r/windows_virtual_machine_scale_set#setting WindowsVirtualMachineScaleSet#setting}.
	Setting *string `field:"required" json:"setting" yaml:"setting"`
}

