package provider


type AzurestackProviderFeaturesVirtualMachineScaleSet struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurestack#roll_instances_when_required AzurestackProvider#roll_instances_when_required}.
	RollInstancesWhenRequired interface{} `field:"required" json:"rollInstancesWhenRequired" yaml:"rollInstancesWhenRequired"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurestack#force_delete AzurestackProvider#force_delete}.
	ForceDelete interface{} `field:"optional" json:"forceDelete" yaml:"forceDelete"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurestack#scale_to_zero_before_deletion AzurestackProvider#scale_to_zero_before_deletion}.
	ScaleToZeroBeforeDeletion interface{} `field:"optional" json:"scaleToZeroBeforeDeletion" yaml:"scaleToZeroBeforeDeletion"`
}

