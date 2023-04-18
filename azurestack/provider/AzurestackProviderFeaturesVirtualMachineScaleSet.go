package provider


type AzurestackProviderFeaturesVirtualMachineScaleSet struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurestack/1.0.0/docs#roll_instances_when_required AzurestackProvider#roll_instances_when_required}.
	RollInstancesWhenRequired interface{} `field:"required" json:"rollInstancesWhenRequired" yaml:"rollInstancesWhenRequired"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurestack/1.0.0/docs#force_delete AzurestackProvider#force_delete}.
	ForceDelete interface{} `field:"optional" json:"forceDelete" yaml:"forceDelete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurestack/1.0.0/docs#scale_to_zero_before_deletion AzurestackProvider#scale_to_zero_before_deletion}.
	ScaleToZeroBeforeDeletion interface{} `field:"optional" json:"scaleToZeroBeforeDeletion" yaml:"scaleToZeroBeforeDeletion"`
}

