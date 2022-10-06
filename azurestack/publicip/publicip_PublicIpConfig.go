package publicip

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type PublicIpConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count *float64 `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktf.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurestack/r/public_ip#location PublicIp#location}.
	Location *string `field:"required" json:"location" yaml:"location"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurestack/r/public_ip#name PublicIp#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurestack/r/public_ip#resource_group_name PublicIp#resource_group_name}.
	ResourceGroupName *string `field:"required" json:"resourceGroupName" yaml:"resourceGroupName"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurestack/r/public_ip#allocation_method PublicIp#allocation_method}.
	AllocationMethod *string `field:"optional" json:"allocationMethod" yaml:"allocationMethod"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurestack/r/public_ip#domain_name_label PublicIp#domain_name_label}.
	DomainNameLabel *string `field:"optional" json:"domainNameLabel" yaml:"domainNameLabel"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurestack/r/public_ip#id PublicIp#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurestack/r/public_ip#idle_timeout_in_minutes PublicIp#idle_timeout_in_minutes}.
	IdleTimeoutInMinutes *float64 `field:"optional" json:"idleTimeoutInMinutes" yaml:"idleTimeoutInMinutes"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurestack/r/public_ip#ip_version PublicIp#ip_version}.
	IpVersion *string `field:"optional" json:"ipVersion" yaml:"ipVersion"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurestack/r/public_ip#public_ip_address_allocation PublicIp#public_ip_address_allocation}.
	PublicIpAddressAllocation *string `field:"optional" json:"publicIpAddressAllocation" yaml:"publicIpAddressAllocation"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurestack/r/public_ip#reverse_fqdn PublicIp#reverse_fqdn}.
	ReverseFqdn *string `field:"optional" json:"reverseFqdn" yaml:"reverseFqdn"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurestack/r/public_ip#sku PublicIp#sku}.
	Sku *string `field:"optional" json:"sku" yaml:"sku"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurestack/r/public_ip#tags PublicIp#tags}.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurestack/r/public_ip#timeouts PublicIp#timeouts}
	Timeouts *PublicIpTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

