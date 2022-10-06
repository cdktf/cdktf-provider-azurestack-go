package subnet


type SubnetTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurestack/r/subnet#create Subnet#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurestack/r/subnet#delete Subnet#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurestack/r/subnet#read Subnet#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurestack/r/subnet#update Subnet#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

