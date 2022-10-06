package networkinterfacebackendaddresspoolassociation


type NetworkInterfaceBackendAddressPoolAssociationTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurestack/r/network_interface_backend_address_pool_association#create NetworkInterfaceBackendAddressPoolAssociation#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurestack/r/network_interface_backend_address_pool_association#delete NetworkInterfaceBackendAddressPoolAssociation#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurestack/r/network_interface_backend_address_pool_association#read NetworkInterfaceBackendAddressPoolAssociation#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurestack/r/network_interface_backend_address_pool_association#update NetworkInterfaceBackendAddressPoolAssociation#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

