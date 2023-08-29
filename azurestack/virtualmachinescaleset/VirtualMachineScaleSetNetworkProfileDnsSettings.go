// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package virtualmachinescaleset


type VirtualMachineScaleSetNetworkProfileDnsSettings struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurestack/1.0.0/docs/resources/virtual_machine_scale_set#dns_servers VirtualMachineScaleSet#dns_servers}.
	DnsServers *[]*string `field:"required" json:"dnsServers" yaml:"dnsServers"`
}

