package virtualmachine


type VirtualMachineOsProfileLinuxConfig struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurestack/r/virtual_machine#disable_password_authentication VirtualMachine#disable_password_authentication}.
	DisablePasswordAuthentication interface{} `field:"required" json:"disablePasswordAuthentication" yaml:"disablePasswordAuthentication"`
	// ssh_keys block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurestack/r/virtual_machine#ssh_keys VirtualMachine#ssh_keys}
	SshKeys interface{} `field:"optional" json:"sshKeys" yaml:"sshKeys"`
}

