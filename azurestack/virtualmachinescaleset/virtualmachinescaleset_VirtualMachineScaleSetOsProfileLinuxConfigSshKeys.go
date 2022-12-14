package virtualmachinescaleset


type VirtualMachineScaleSetOsProfileLinuxConfigSshKeys struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurestack/r/virtual_machine_scale_set#path VirtualMachineScaleSet#path}.
	Path *string `field:"required" json:"path" yaml:"path"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurestack/r/virtual_machine_scale_set#key_data VirtualMachineScaleSet#key_data}.
	KeyData *string `field:"optional" json:"keyData" yaml:"keyData"`
}

