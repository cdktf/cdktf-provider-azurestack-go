package lb


type LbFrontendIpConfiguration struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurestack/r/lb#name Lb#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurestack/r/lb#private_ip_address Lb#private_ip_address}.
	PrivateIpAddress *string `field:"optional" json:"privateIpAddress" yaml:"privateIpAddress"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurestack/r/lb#private_ip_address_allocation Lb#private_ip_address_allocation}.
	PrivateIpAddressAllocation *string `field:"optional" json:"privateIpAddressAllocation" yaml:"privateIpAddressAllocation"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurestack/r/lb#public_ip_address_id Lb#public_ip_address_id}.
	PublicIpAddressId *string `field:"optional" json:"publicIpAddressId" yaml:"publicIpAddressId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurestack/r/lb#subnet_id Lb#subnet_id}.
	SubnetId *string `field:"optional" json:"subnetId" yaml:"subnetId"`
}

