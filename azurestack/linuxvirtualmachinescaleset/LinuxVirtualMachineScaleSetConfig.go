package linuxvirtualmachinescaleset

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type LinuxVirtualMachineScaleSetConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurestack/1.0.0/docs/resources/linux_virtual_machine_scale_set#admin_username LinuxVirtualMachineScaleSet#admin_username}.
	AdminUsername *string `field:"required" json:"adminUsername" yaml:"adminUsername"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurestack/1.0.0/docs/resources/linux_virtual_machine_scale_set#instances LinuxVirtualMachineScaleSet#instances}.
	Instances *float64 `field:"required" json:"instances" yaml:"instances"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurestack/1.0.0/docs/resources/linux_virtual_machine_scale_set#location LinuxVirtualMachineScaleSet#location}.
	Location *string `field:"required" json:"location" yaml:"location"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurestack/1.0.0/docs/resources/linux_virtual_machine_scale_set#name LinuxVirtualMachineScaleSet#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// network_interface block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurestack/1.0.0/docs/resources/linux_virtual_machine_scale_set#network_interface LinuxVirtualMachineScaleSet#network_interface}
	NetworkInterface interface{} `field:"required" json:"networkInterface" yaml:"networkInterface"`
	// os_disk block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurestack/1.0.0/docs/resources/linux_virtual_machine_scale_set#os_disk LinuxVirtualMachineScaleSet#os_disk}
	OsDisk *LinuxVirtualMachineScaleSetOsDisk `field:"required" json:"osDisk" yaml:"osDisk"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurestack/1.0.0/docs/resources/linux_virtual_machine_scale_set#resource_group_name LinuxVirtualMachineScaleSet#resource_group_name}.
	ResourceGroupName *string `field:"required" json:"resourceGroupName" yaml:"resourceGroupName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurestack/1.0.0/docs/resources/linux_virtual_machine_scale_set#sku LinuxVirtualMachineScaleSet#sku}.
	Sku *string `field:"required" json:"sku" yaml:"sku"`
	// additional_capabilities block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurestack/1.0.0/docs/resources/linux_virtual_machine_scale_set#additional_capabilities LinuxVirtualMachineScaleSet#additional_capabilities}
	AdditionalCapabilities *LinuxVirtualMachineScaleSetAdditionalCapabilities `field:"optional" json:"additionalCapabilities" yaml:"additionalCapabilities"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurestack/1.0.0/docs/resources/linux_virtual_machine_scale_set#admin_password LinuxVirtualMachineScaleSet#admin_password}.
	AdminPassword *string `field:"optional" json:"adminPassword" yaml:"adminPassword"`
	// admin_ssh_key block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurestack/1.0.0/docs/resources/linux_virtual_machine_scale_set#admin_ssh_key LinuxVirtualMachineScaleSet#admin_ssh_key}
	AdminSshKey interface{} `field:"optional" json:"adminSshKey" yaml:"adminSshKey"`
	// automatic_instance_repair block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurestack/1.0.0/docs/resources/linux_virtual_machine_scale_set#automatic_instance_repair LinuxVirtualMachineScaleSet#automatic_instance_repair}
	AutomaticInstanceRepair *LinuxVirtualMachineScaleSetAutomaticInstanceRepair `field:"optional" json:"automaticInstanceRepair" yaml:"automaticInstanceRepair"`
	// automatic_os_upgrade_policy block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurestack/1.0.0/docs/resources/linux_virtual_machine_scale_set#automatic_os_upgrade_policy LinuxVirtualMachineScaleSet#automatic_os_upgrade_policy}
	AutomaticOsUpgradePolicy *LinuxVirtualMachineScaleSetAutomaticOsUpgradePolicy `field:"optional" json:"automaticOsUpgradePolicy" yaml:"automaticOsUpgradePolicy"`
	// boot_diagnostics block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurestack/1.0.0/docs/resources/linux_virtual_machine_scale_set#boot_diagnostics LinuxVirtualMachineScaleSet#boot_diagnostics}
	BootDiagnostics *LinuxVirtualMachineScaleSetBootDiagnostics `field:"optional" json:"bootDiagnostics" yaml:"bootDiagnostics"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurestack/1.0.0/docs/resources/linux_virtual_machine_scale_set#computer_name_prefix LinuxVirtualMachineScaleSet#computer_name_prefix}.
	ComputerNamePrefix *string `field:"optional" json:"computerNamePrefix" yaml:"computerNamePrefix"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurestack/1.0.0/docs/resources/linux_virtual_machine_scale_set#custom_data LinuxVirtualMachineScaleSet#custom_data}.
	CustomData *string `field:"optional" json:"customData" yaml:"customData"`
	// data_disk block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurestack/1.0.0/docs/resources/linux_virtual_machine_scale_set#data_disk LinuxVirtualMachineScaleSet#data_disk}
	DataDisk interface{} `field:"optional" json:"dataDisk" yaml:"dataDisk"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurestack/1.0.0/docs/resources/linux_virtual_machine_scale_set#disable_password_authentication LinuxVirtualMachineScaleSet#disable_password_authentication}.
	DisablePasswordAuthentication interface{} `field:"optional" json:"disablePasswordAuthentication" yaml:"disablePasswordAuthentication"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurestack/1.0.0/docs/resources/linux_virtual_machine_scale_set#do_not_run_extensions_on_overprovisioned_machines LinuxVirtualMachineScaleSet#do_not_run_extensions_on_overprovisioned_machines}.
	DoNotRunExtensionsOnOverprovisionedMachines interface{} `field:"optional" json:"doNotRunExtensionsOnOverprovisionedMachines" yaml:"doNotRunExtensionsOnOverprovisionedMachines"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurestack/1.0.0/docs/resources/linux_virtual_machine_scale_set#encryption_at_host_enabled LinuxVirtualMachineScaleSet#encryption_at_host_enabled}.
	EncryptionAtHostEnabled interface{} `field:"optional" json:"encryptionAtHostEnabled" yaml:"encryptionAtHostEnabled"`
	// extension block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurestack/1.0.0/docs/resources/linux_virtual_machine_scale_set#extension LinuxVirtualMachineScaleSet#extension}
	Extension interface{} `field:"optional" json:"extension" yaml:"extension"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurestack/1.0.0/docs/resources/linux_virtual_machine_scale_set#health_probe_id LinuxVirtualMachineScaleSet#health_probe_id}.
	HealthProbeId *string `field:"optional" json:"healthProbeId" yaml:"healthProbeId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurestack/1.0.0/docs/resources/linux_virtual_machine_scale_set#id LinuxVirtualMachineScaleSet#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurestack/1.0.0/docs/resources/linux_virtual_machine_scale_set#overprovision LinuxVirtualMachineScaleSet#overprovision}.
	Overprovision interface{} `field:"optional" json:"overprovision" yaml:"overprovision"`
	// plan block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurestack/1.0.0/docs/resources/linux_virtual_machine_scale_set#plan LinuxVirtualMachineScaleSet#plan}
	Plan *LinuxVirtualMachineScaleSetPlan `field:"optional" json:"plan" yaml:"plan"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurestack/1.0.0/docs/resources/linux_virtual_machine_scale_set#platform_fault_domain_count LinuxVirtualMachineScaleSet#platform_fault_domain_count}.
	PlatformFaultDomainCount *float64 `field:"optional" json:"platformFaultDomainCount" yaml:"platformFaultDomainCount"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurestack/1.0.0/docs/resources/linux_virtual_machine_scale_set#provision_vm_agent LinuxVirtualMachineScaleSet#provision_vm_agent}.
	ProvisionVmAgent interface{} `field:"optional" json:"provisionVmAgent" yaml:"provisionVmAgent"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurestack/1.0.0/docs/resources/linux_virtual_machine_scale_set#scale_in_policy LinuxVirtualMachineScaleSet#scale_in_policy}.
	ScaleInPolicy *string `field:"optional" json:"scaleInPolicy" yaml:"scaleInPolicy"`
	// secret block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurestack/1.0.0/docs/resources/linux_virtual_machine_scale_set#secret LinuxVirtualMachineScaleSet#secret}
	Secret interface{} `field:"optional" json:"secret" yaml:"secret"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurestack/1.0.0/docs/resources/linux_virtual_machine_scale_set#single_placement_group LinuxVirtualMachineScaleSet#single_placement_group}.
	SinglePlacementGroup interface{} `field:"optional" json:"singlePlacementGroup" yaml:"singlePlacementGroup"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurestack/1.0.0/docs/resources/linux_virtual_machine_scale_set#source_image_id LinuxVirtualMachineScaleSet#source_image_id}.
	SourceImageId *string `field:"optional" json:"sourceImageId" yaml:"sourceImageId"`
	// source_image_reference block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurestack/1.0.0/docs/resources/linux_virtual_machine_scale_set#source_image_reference LinuxVirtualMachineScaleSet#source_image_reference}
	SourceImageReference *LinuxVirtualMachineScaleSetSourceImageReference `field:"optional" json:"sourceImageReference" yaml:"sourceImageReference"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurestack/1.0.0/docs/resources/linux_virtual_machine_scale_set#tags LinuxVirtualMachineScaleSet#tags}.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// terminate_notification block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurestack/1.0.0/docs/resources/linux_virtual_machine_scale_set#terminate_notification LinuxVirtualMachineScaleSet#terminate_notification}
	TerminateNotification *LinuxVirtualMachineScaleSetTerminateNotification `field:"optional" json:"terminateNotification" yaml:"terminateNotification"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurestack/1.0.0/docs/resources/linux_virtual_machine_scale_set#timeouts LinuxVirtualMachineScaleSet#timeouts}
	Timeouts *LinuxVirtualMachineScaleSetTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurestack/1.0.0/docs/resources/linux_virtual_machine_scale_set#upgrade_mode LinuxVirtualMachineScaleSet#upgrade_mode}.
	UpgradeMode *string `field:"optional" json:"upgradeMode" yaml:"upgradeMode"`
}

