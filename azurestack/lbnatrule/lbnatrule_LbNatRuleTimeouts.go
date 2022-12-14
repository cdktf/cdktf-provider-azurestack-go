package lbnatrule


type LbNatRuleTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurestack/r/lb_nat_rule#create LbNatRule#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurestack/r/lb_nat_rule#delete LbNatRule#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurestack/r/lb_nat_rule#read LbNatRule#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurestack/r/lb_nat_rule#update LbNatRule#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

