package lb


type LbTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurestack/r/lb#create Lb#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurestack/r/lb#delete Lb#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurestack/r/lb#read Lb#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurestack/r/lb#update Lb#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

