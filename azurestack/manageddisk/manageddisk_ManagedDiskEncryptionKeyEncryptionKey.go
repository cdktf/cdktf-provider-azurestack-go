package manageddisk


type ManagedDiskEncryptionKeyEncryptionKey struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurestack/r/managed_disk#key_url ManagedDisk#key_url}.
	KeyUrl *string `field:"required" json:"keyUrl" yaml:"keyUrl"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurestack/r/managed_disk#source_vault_id ManagedDisk#source_vault_id}.
	SourceVaultId *string `field:"required" json:"sourceVaultId" yaml:"sourceVaultId"`
}

