package provider


type AzurestackProviderFeatures struct {
	// resource_group block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurestack#resource_group AzurestackProvider#resource_group}
	ResourceGroup *AzurestackProviderFeaturesResourceGroup `field:"optional" json:"resourceGroup" yaml:"resourceGroup"`
	// virtual_machine block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurestack#virtual_machine AzurestackProvider#virtual_machine}
	VirtualMachine *AzurestackProviderFeaturesVirtualMachine `field:"optional" json:"virtualMachine" yaml:"virtualMachine"`
	// virtual_machine_scale_set block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurestack#virtual_machine_scale_set AzurestackProvider#virtual_machine_scale_set}
	VirtualMachineScaleSet *AzurestackProviderFeaturesVirtualMachineScaleSet `field:"optional" json:"virtualMachineScaleSet" yaml:"virtualMachineScaleSet"`
}

