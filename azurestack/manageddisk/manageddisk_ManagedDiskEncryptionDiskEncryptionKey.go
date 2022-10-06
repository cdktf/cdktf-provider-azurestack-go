package manageddisk


type ManagedDiskEncryptionDiskEncryptionKey struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurestack/r/managed_disk#secret_url ManagedDisk#secret_url}.
	SecretUrl *string `field:"required" json:"secretUrl" yaml:"secretUrl"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurestack/r/managed_disk#source_vault_id ManagedDisk#source_vault_id}.
	SourceVaultId *string `field:"required" json:"sourceVaultId" yaml:"sourceVaultId"`
}

