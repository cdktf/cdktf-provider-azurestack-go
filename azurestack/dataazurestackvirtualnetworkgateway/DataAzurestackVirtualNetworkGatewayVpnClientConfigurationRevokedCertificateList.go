// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dataazurestackvirtualnetworkgateway

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-azurestack-go/azurestack/v6/jsii"

	"github.com/cdktf/cdktf-provider-azurestack-go/azurestack/v6/dataazurestackvirtualnetworkgateway/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataAzurestackVirtualNetworkGatewayVpnClientConfigurationRevokedCertificateList interface {
	cdktf.ComplexList
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	// The attribute on the parent resource this class is referencing.
	TerraformAttribute() *string
	SetTerraformAttribute(val *string)
	// The parent resource.
	TerraformResource() cdktf.IInterpolatingParent
	SetTerraformResource(val cdktf.IInterpolatingParent)
	// whether the list is wrapping a set (will add tolist() to be able to access an item via an index).
	WrapsSet() *bool
	SetWrapsSet(val *bool)
	// Experimental.
	ComputeFqn() *string
	Get(index *float64) DataAzurestackVirtualNetworkGatewayVpnClientConfigurationRevokedCertificateOutputReference
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DataAzurestackVirtualNetworkGatewayVpnClientConfigurationRevokedCertificateList
type jsiiProxy_DataAzurestackVirtualNetworkGatewayVpnClientConfigurationRevokedCertificateList struct {
	internal.Type__cdktfComplexList
}

func (j *jsiiProxy_DataAzurestackVirtualNetworkGatewayVpnClientConfigurationRevokedCertificateList) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurestackVirtualNetworkGatewayVpnClientConfigurationRevokedCertificateList) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurestackVirtualNetworkGatewayVpnClientConfigurationRevokedCertificateList) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurestackVirtualNetworkGatewayVpnClientConfigurationRevokedCertificateList) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurestackVirtualNetworkGatewayVpnClientConfigurationRevokedCertificateList) WrapsSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"wrapsSet",
		&returns,
	)
	return returns
}


func NewDataAzurestackVirtualNetworkGatewayVpnClientConfigurationRevokedCertificateList(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) DataAzurestackVirtualNetworkGatewayVpnClientConfigurationRevokedCertificateList {
	_init_.Initialize()

	if err := validateNewDataAzurestackVirtualNetworkGatewayVpnClientConfigurationRevokedCertificateListParameters(terraformResource, terraformAttribute, wrapsSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataAzurestackVirtualNetworkGatewayVpnClientConfigurationRevokedCertificateList{}

	_jsii_.Create(
		"@cdktf/provider-azurestack.dataAzurestackVirtualNetworkGateway.DataAzurestackVirtualNetworkGatewayVpnClientConfigurationRevokedCertificateList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		&j,
	)

	return &j
}

func NewDataAzurestackVirtualNetworkGatewayVpnClientConfigurationRevokedCertificateList_Override(d DataAzurestackVirtualNetworkGatewayVpnClientConfigurationRevokedCertificateList, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurestack.dataAzurestackVirtualNetworkGateway.DataAzurestackVirtualNetworkGatewayVpnClientConfigurationRevokedCertificateList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		d,
	)
}

func (j *jsiiProxy_DataAzurestackVirtualNetworkGatewayVpnClientConfigurationRevokedCertificateList)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataAzurestackVirtualNetworkGatewayVpnClientConfigurationRevokedCertificateList)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_DataAzurestackVirtualNetworkGatewayVpnClientConfigurationRevokedCertificateList)SetWrapsSet(val *bool) {
	if err := j.validateSetWrapsSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"wrapsSet",
		val,
	)
}

func (d *jsiiProxy_DataAzurestackVirtualNetworkGatewayVpnClientConfigurationRevokedCertificateList) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAzurestackVirtualNetworkGatewayVpnClientConfigurationRevokedCertificateList) Get(index *float64) DataAzurestackVirtualNetworkGatewayVpnClientConfigurationRevokedCertificateOutputReference {
	if err := d.validateGetParameters(index); err != nil {
		panic(err)
	}
	var returns DataAzurestackVirtualNetworkGatewayVpnClientConfigurationRevokedCertificateOutputReference

	_jsii_.Invoke(
		d,
		"get",
		[]interface{}{index},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAzurestackVirtualNetworkGatewayVpnClientConfigurationRevokedCertificateList) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := d.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		d,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAzurestackVirtualNetworkGatewayVpnClientConfigurationRevokedCertificateList) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

