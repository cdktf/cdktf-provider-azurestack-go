// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dnstxtrecord


type DnsTxtRecordRecord struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurestack/1.0.0/docs/resources/dns_txt_record#value DnsTxtRecord#value}.
	Value *string `field:"required" json:"value" yaml:"value"`
}

