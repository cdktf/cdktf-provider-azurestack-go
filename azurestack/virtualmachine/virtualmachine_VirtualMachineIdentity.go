package virtualmachine


type VirtualMachineIdentity struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurestack/r/virtual_machine#type VirtualMachine#type}.
	Type *string `field:"required" json:"type" yaml:"type"`
}

