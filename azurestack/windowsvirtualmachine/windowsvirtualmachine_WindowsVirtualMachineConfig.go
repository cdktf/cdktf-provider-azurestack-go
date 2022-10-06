package windowsvirtualmachine

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type WindowsVirtualMachineConfig struct {
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
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurestack/r/windows_virtual_machine#admin_password WindowsVirtualMachine#admin_password}.
	AdminPassword *string `field:"required" json:"adminPassword" yaml:"adminPassword"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurestack/r/windows_virtual_machine#admin_username WindowsVirtualMachine#admin_username}.
	AdminUsername *string `field:"required" json:"adminUsername" yaml:"adminUsername"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurestack/r/windows_virtual_machine#location WindowsVirtualMachine#location}.
	Location *string `field:"required" json:"location" yaml:"location"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurestack/r/windows_virtual_machine#name WindowsVirtualMachine#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurestack/r/windows_virtual_machine#network_interface_ids WindowsVirtualMachine#network_interface_ids}.
	NetworkInterfaceIds *[]*string `field:"required" json:"networkInterfaceIds" yaml:"networkInterfaceIds"`
	// os_disk block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurestack/r/windows_virtual_machine#os_disk WindowsVirtualMachine#os_disk}
	OsDisk *WindowsVirtualMachineOsDisk `field:"required" json:"osDisk" yaml:"osDisk"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurestack/r/windows_virtual_machine#resource_group_name WindowsVirtualMachine#resource_group_name}.
	ResourceGroupName *string `field:"required" json:"resourceGroupName" yaml:"resourceGroupName"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurestack/r/windows_virtual_machine#size WindowsVirtualMachine#size}.
	Size *string `field:"required" json:"size" yaml:"size"`
	// additional_capabilities block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurestack/r/windows_virtual_machine#additional_capabilities WindowsVirtualMachine#additional_capabilities}
	AdditionalCapabilities *WindowsVirtualMachineAdditionalCapabilities `field:"optional" json:"additionalCapabilities" yaml:"additionalCapabilities"`
	// additional_unattend_content block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurestack/r/windows_virtual_machine#additional_unattend_content WindowsVirtualMachine#additional_unattend_content}
	AdditionalUnattendContent interface{} `field:"optional" json:"additionalUnattendContent" yaml:"additionalUnattendContent"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurestack/r/windows_virtual_machine#allow_extension_operations WindowsVirtualMachine#allow_extension_operations}.
	AllowExtensionOperations interface{} `field:"optional" json:"allowExtensionOperations" yaml:"allowExtensionOperations"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurestack/r/windows_virtual_machine#availability_set_id WindowsVirtualMachine#availability_set_id}.
	AvailabilitySetId *string `field:"optional" json:"availabilitySetId" yaml:"availabilitySetId"`
	// boot_diagnostics block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurestack/r/windows_virtual_machine#boot_diagnostics WindowsVirtualMachine#boot_diagnostics}
	BootDiagnostics *WindowsVirtualMachineBootDiagnostics `field:"optional" json:"bootDiagnostics" yaml:"bootDiagnostics"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurestack/r/windows_virtual_machine#computer_name WindowsVirtualMachine#computer_name}.
	ComputerName *string `field:"optional" json:"computerName" yaml:"computerName"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurestack/r/windows_virtual_machine#custom_data WindowsVirtualMachine#custom_data}.
	CustomData *string `field:"optional" json:"customData" yaml:"customData"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurestack/r/windows_virtual_machine#enable_automatic_updates WindowsVirtualMachine#enable_automatic_updates}.
	EnableAutomaticUpdates interface{} `field:"optional" json:"enableAutomaticUpdates" yaml:"enableAutomaticUpdates"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurestack/r/windows_virtual_machine#encryption_at_host_enabled WindowsVirtualMachine#encryption_at_host_enabled}.
	EncryptionAtHostEnabled interface{} `field:"optional" json:"encryptionAtHostEnabled" yaml:"encryptionAtHostEnabled"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurestack/r/windows_virtual_machine#eviction_policy WindowsVirtualMachine#eviction_policy}.
	EvictionPolicy *string `field:"optional" json:"evictionPolicy" yaml:"evictionPolicy"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurestack/r/windows_virtual_machine#extensions_time_budget WindowsVirtualMachine#extensions_time_budget}.
	ExtensionsTimeBudget *string `field:"optional" json:"extensionsTimeBudget" yaml:"extensionsTimeBudget"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurestack/r/windows_virtual_machine#id WindowsVirtualMachine#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurestack/r/windows_virtual_machine#license_type WindowsVirtualMachine#license_type}.
	LicenseType *string `field:"optional" json:"licenseType" yaml:"licenseType"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurestack/r/windows_virtual_machine#max_bid_price WindowsVirtualMachine#max_bid_price}.
	MaxBidPrice *float64 `field:"optional" json:"maxBidPrice" yaml:"maxBidPrice"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurestack/r/windows_virtual_machine#patch_mode WindowsVirtualMachine#patch_mode}.
	PatchMode *string `field:"optional" json:"patchMode" yaml:"patchMode"`
	// plan block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurestack/r/windows_virtual_machine#plan WindowsVirtualMachine#plan}
	Plan *WindowsVirtualMachinePlan `field:"optional" json:"plan" yaml:"plan"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurestack/r/windows_virtual_machine#priority WindowsVirtualMachine#priority}.
	Priority *string `field:"optional" json:"priority" yaml:"priority"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurestack/r/windows_virtual_machine#provision_vm_agent WindowsVirtualMachine#provision_vm_agent}.
	ProvisionVmAgent interface{} `field:"optional" json:"provisionVmAgent" yaml:"provisionVmAgent"`
	// secret block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurestack/r/windows_virtual_machine#secret WindowsVirtualMachine#secret}
	Secret interface{} `field:"optional" json:"secret" yaml:"secret"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurestack/r/windows_virtual_machine#source_image_id WindowsVirtualMachine#source_image_id}.
	SourceImageId *string `field:"optional" json:"sourceImageId" yaml:"sourceImageId"`
	// source_image_reference block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurestack/r/windows_virtual_machine#source_image_reference WindowsVirtualMachine#source_image_reference}
	SourceImageReference *WindowsVirtualMachineSourceImageReference `field:"optional" json:"sourceImageReference" yaml:"sourceImageReference"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurestack/r/windows_virtual_machine#tags WindowsVirtualMachine#tags}.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurestack/r/windows_virtual_machine#timeouts WindowsVirtualMachine#timeouts}
	Timeouts *WindowsVirtualMachineTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurestack/r/windows_virtual_machine#timezone WindowsVirtualMachine#timezone}.
	Timezone *string `field:"optional" json:"timezone" yaml:"timezone"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurestack/r/windows_virtual_machine#virtual_machine_scale_set_id WindowsVirtualMachine#virtual_machine_scale_set_id}.
	VirtualMachineScaleSetId *string `field:"optional" json:"virtualMachineScaleSetId" yaml:"virtualMachineScaleSetId"`
	// winrm_listener block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurestack/r/windows_virtual_machine#winrm_listener WindowsVirtualMachine#winrm_listener}
	WinrmListener interface{} `field:"optional" json:"winrmListener" yaml:"winrmListener"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurestack/r/windows_virtual_machine#zone WindowsVirtualMachine#zone}.
	Zone *string `field:"optional" json:"zone" yaml:"zone"`
}

