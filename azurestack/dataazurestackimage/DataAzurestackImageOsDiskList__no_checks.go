// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package dataazurestackimage

// Building without runtime type checking enabled, so all the below just return nil

func (d *jsiiProxy_DataAzurestackImageOsDiskList) validateGetParameters(index *float64) error {
	return nil
}

func (d *jsiiProxy_DataAzurestackImageOsDiskList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_DataAzurestackImageOsDiskList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_DataAzurestackImageOsDiskList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_DataAzurestackImageOsDiskList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewDataAzurestackImageOsDiskListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

