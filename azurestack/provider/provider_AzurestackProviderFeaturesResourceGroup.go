package provider


type AzurestackProviderFeaturesResourceGroup struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurestack#prevent_deletion_if_contains_resources AzurestackProvider#prevent_deletion_if_contains_resources}.
	PreventDeletionIfContainsResources interface{} `field:"optional" json:"preventDeletionIfContainsResources" yaml:"preventDeletionIfContainsResources"`
}

