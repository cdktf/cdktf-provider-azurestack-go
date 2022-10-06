package virtualnetworkgateway


type VirtualNetworkGatewayBgpSettings struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurestack/r/virtual_network_gateway#asn VirtualNetworkGateway#asn}.
	Asn *float64 `field:"optional" json:"asn" yaml:"asn"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurestack/r/virtual_network_gateway#peering_address VirtualNetworkGateway#peering_address}.
	PeeringAddress *string `field:"optional" json:"peeringAddress" yaml:"peeringAddress"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurestack/r/virtual_network_gateway#peer_weight VirtualNetworkGateway#peer_weight}.
	PeerWeight *float64 `field:"optional" json:"peerWeight" yaml:"peerWeight"`
}

