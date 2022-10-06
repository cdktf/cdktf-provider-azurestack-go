package dnszone

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DnsZoneConfig struct {
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
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurestack/r/dns_zone#name DnsZone#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurestack/r/dns_zone#resource_group_name DnsZone#resource_group_name}.
	ResourceGroupName *string `field:"required" json:"resourceGroupName" yaml:"resourceGroupName"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurestack/r/dns_zone#id DnsZone#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// soa_record block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurestack/r/dns_zone#soa_record DnsZone#soa_record}
	SoaRecord *DnsZoneSoaRecord `field:"optional" json:"soaRecord" yaml:"soaRecord"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurestack/r/dns_zone#tags DnsZone#tags}.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurestack/r/dns_zone#timeouts DnsZone#timeouts}
	Timeouts *DnsZoneTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

