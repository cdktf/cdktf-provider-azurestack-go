package windowsvirtualmachine


type WindowsVirtualMachineWinrmListener struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurestack/r/windows_virtual_machine#protocol WindowsVirtualMachine#protocol}.
	Protocol *string `field:"required" json:"protocol" yaml:"protocol"`
}

