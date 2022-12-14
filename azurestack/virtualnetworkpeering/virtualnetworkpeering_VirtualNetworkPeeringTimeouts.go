package virtualnetworkpeering


type VirtualNetworkPeeringTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurestack/r/virtual_network_peering#create VirtualNetworkPeering#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurestack/r/virtual_network_peering#delete VirtualNetworkPeering#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurestack/r/virtual_network_peering#read VirtualNetworkPeering#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurestack/r/virtual_network_peering#update VirtualNetworkPeering#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

