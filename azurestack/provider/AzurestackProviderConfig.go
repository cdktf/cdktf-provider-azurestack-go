package provider


type AzurestackProviderConfig struct {
	// features block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurestack/1.0.0/docs#features AzurestackProvider#features}
	Features *AzurestackProviderFeatures `field:"required" json:"features" yaml:"features"`
	// Alias name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurestack/1.0.0/docs#alias AzurestackProvider#alias}
	Alias *string `field:"optional" json:"alias" yaml:"alias"`
	// The Hostname which should be used for the Azure Metadata Service.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurestack/1.0.0/docs#arm_endpoint AzurestackProvider#arm_endpoint}
	ArmEndpoint *string `field:"optional" json:"armEndpoint" yaml:"armEndpoint"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurestack/1.0.0/docs#auxiliary_tenant_ids AzurestackProvider#auxiliary_tenant_ids}.
	AuxiliaryTenantIds *[]*string `field:"optional" json:"auxiliaryTenantIds" yaml:"auxiliaryTenantIds"`
	// The password associated with the Client Certificate. For use when authenticating as a Service Principal using a Client Certificate.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurestack/1.0.0/docs#client_certificate_password AzurestackProvider#client_certificate_password}
	ClientCertificatePassword *string `field:"optional" json:"clientCertificatePassword" yaml:"clientCertificatePassword"`
	// The path to the Client Certificate associated with the Service Principal for use when authenticating as a Service Principal using a Client Certificate.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurestack/1.0.0/docs#client_certificate_path AzurestackProvider#client_certificate_path}
	ClientCertificatePath *string `field:"optional" json:"clientCertificatePath" yaml:"clientCertificatePath"`
	// The Client ID which should be used.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurestack/1.0.0/docs#client_id AzurestackProvider#client_id}
	ClientId *string `field:"optional" json:"clientId" yaml:"clientId"`
	// The Client Secret which should be used. For use When authenticating as a Service Principal using a Client Secret.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurestack/1.0.0/docs#client_secret AzurestackProvider#client_secret}
	ClientSecret *string `field:"optional" json:"clientSecret" yaml:"clientSecret"`
	// This will disable the x-ms-correlation-request-id header.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurestack/1.0.0/docs#disable_correlation_request_id AzurestackProvider#disable_correlation_request_id}
	DisableCorrelationRequestId interface{} `field:"optional" json:"disableCorrelationRequestId" yaml:"disableCorrelationRequestId"`
	// The Cloud Environment which should be used.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurestack/1.0.0/docs#environment AzurestackProvider#environment}
	Environment *string `field:"optional" json:"environment" yaml:"environment"`
	// The Hostname which should be used for the Azure Metadata Service.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurestack/1.0.0/docs#metadata_host AzurestackProvider#metadata_host}
	MetadataHost *string `field:"optional" json:"metadataHost" yaml:"metadataHost"`
	// The path to a custom endpoint for Managed Service Identity - in most circumstances this should be detected automatically.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurestack/1.0.0/docs#msi_endpoint AzurestackProvider#msi_endpoint}
	MsiEndpoint *string `field:"optional" json:"msiEndpoint" yaml:"msiEndpoint"`
	// Should the AzureStack Provider skip registering all of the Resource Providers that it supports, if they're not already registered?
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurestack/1.0.0/docs#skip_provider_registration AzurestackProvider#skip_provider_registration}
	SkipProviderRegistration interface{} `field:"optional" json:"skipProviderRegistration" yaml:"skipProviderRegistration"`
	// The Subscription ID which should be used.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurestack/1.0.0/docs#subscription_id AzurestackProvider#subscription_id}
	SubscriptionId *string `field:"optional" json:"subscriptionId" yaml:"subscriptionId"`
	// The Tenant ID which should be used.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurestack/1.0.0/docs#tenant_id AzurestackProvider#tenant_id}
	TenantId *string `field:"optional" json:"tenantId" yaml:"tenantId"`
	// Allowed Managed Service Identity be used for Authentication.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurestack/1.0.0/docs#use_msi AzurestackProvider#use_msi}
	UseMsi interface{} `field:"optional" json:"useMsi" yaml:"useMsi"`
}

