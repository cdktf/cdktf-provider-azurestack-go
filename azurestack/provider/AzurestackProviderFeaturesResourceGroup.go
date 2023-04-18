package provider


type AzurestackProviderFeaturesResourceGroup struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurestack/1.0.0/docs#prevent_deletion_if_contains_resources AzurestackProvider#prevent_deletion_if_contains_resources}.
	PreventDeletionIfContainsResources interface{} `field:"optional" json:"preventDeletionIfContainsResources" yaml:"preventDeletionIfContainsResources"`
}

