package manageddisk


type ManagedDiskEncryption struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurestack/r/managed_disk#enabled ManagedDisk#enabled}.
	Enabled interface{} `field:"required" json:"enabled" yaml:"enabled"`
	// disk_encryption_key block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurestack/r/managed_disk#disk_encryption_key ManagedDisk#disk_encryption_key}
	DiskEncryptionKey *ManagedDiskEncryptionDiskEncryptionKey `field:"optional" json:"diskEncryptionKey" yaml:"diskEncryptionKey"`
	// key_encryption_key block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurestack/r/managed_disk#key_encryption_key ManagedDisk#key_encryption_key}
	KeyEncryptionKey *ManagedDiskEncryptionKeyEncryptionKey `field:"optional" json:"keyEncryptionKey" yaml:"keyEncryptionKey"`
}

