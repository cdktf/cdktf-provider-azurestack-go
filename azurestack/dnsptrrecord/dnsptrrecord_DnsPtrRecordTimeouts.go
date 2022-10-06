package dnsptrrecord


type DnsPtrRecordTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurestack/r/dns_ptr_record#create DnsPtrRecord#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurestack/r/dns_ptr_record#delete DnsPtrRecord#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurestack/r/dns_ptr_record#read DnsPtrRecord#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurestack/r/dns_ptr_record#update DnsPtrRecord#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

