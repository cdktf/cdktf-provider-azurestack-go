//go:build no_runtime_type_checking
// +build no_runtime_type_checking

package provider

// Building without runtime type checking enabled, so all the below just return nil

func (a *jsiiProxy_AzurestackProvider) validateAddOverrideParameters(path *string, value interface{}) error {
	return nil
}

func (a *jsiiProxy_AzurestackProvider) validateOverrideLogicalIdParameters(newLogicalId *string) error {
	return nil
}

func validateAzurestackProvider_IsConstructParameters(x interface{}) error {
	return nil
}

func (j *jsiiProxy_AzurestackProvider) validateSetDisableCorrelationRequestIdParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_AzurestackProvider) validateSetFeaturesParameters(val *AzurestackProviderFeatures) error {
	return nil
}

func (j *jsiiProxy_AzurestackProvider) validateSetSkipProviderRegistrationParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_AzurestackProvider) validateSetUseMsiParameters(val interface{}) error {
	return nil
}

func validateNewAzurestackProviderParameters(scope constructs.Construct, id *string, config *AzurestackProviderConfig) error {
	return nil
}

