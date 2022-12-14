package virtualmachinescaleset


type VirtualMachineScaleSetOsProfileSecrets struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurestack/r/virtual_machine_scale_set#source_vault_id VirtualMachineScaleSet#source_vault_id}.
	SourceVaultId *string `field:"required" json:"sourceVaultId" yaml:"sourceVaultId"`
	// vault_certificates block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurestack/r/virtual_machine_scale_set#vault_certificates VirtualMachineScaleSet#vault_certificates}
	VaultCertificates interface{} `field:"optional" json:"vaultCertificates" yaml:"vaultCertificates"`
}

