// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package dataazurestackresources

// Building without runtime type checking enabled, so all the below just return nil

func (d *jsiiProxy_DataAzurestackResourcesResourcesList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (d *jsiiProxy_DataAzurestackResourcesResourcesList) validateGetParameters(index *float64) error {
	return nil
}

func (d *jsiiProxy_DataAzurestackResourcesResourcesList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_DataAzurestackResourcesResourcesList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_DataAzurestackResourcesResourcesList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_DataAzurestackResourcesResourcesList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewDataAzurestackResourcesResourcesListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

