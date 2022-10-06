package virtualmachinescaleset


type VirtualMachineScaleSetNetworkProfileDnsSettings struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurestack/r/virtual_machine_scale_set#dns_servers VirtualMachineScaleSet#dns_servers}.
	DnsServers *[]*string `field:"required" json:"dnsServers" yaml:"dnsServers"`
}

