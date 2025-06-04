// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package provider

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-azurestack-go/azurestack/v9/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-azurestack-go/azurestack/v9/provider/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/azurestack/1.0.0/docs azurestack}.
type AzurestackProvider interface {
	cdktf.TerraformProvider
	Alias() *string
	SetAlias(val *string)
	AliasInput() *string
	ArmEndpoint() *string
	SetArmEndpoint(val *string)
	ArmEndpointInput() *string
	AuxiliaryTenantIds() *[]*string
	SetAuxiliaryTenantIds(val *[]*string)
	AuxiliaryTenantIdsInput() *[]*string
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	ClientCertificatePassword() *string
	SetClientCertificatePassword(val *string)
	ClientCertificatePasswordInput() *string
	ClientCertificatePath() *string
	SetClientCertificatePath(val *string)
	ClientCertificatePathInput() *string
	ClientId() *string
	SetClientId(val *string)
	ClientIdInput() *string
	ClientSecret() *string
	SetClientSecret(val *string)
	ClientSecretInput() *string
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	DisableCorrelationRequestId() interface{}
	SetDisableCorrelationRequestId(val interface{})
	DisableCorrelationRequestIdInput() interface{}
	Environment() *string
	SetEnvironment(val *string)
	EnvironmentInput() *string
	Features() *AzurestackProviderFeatures
	SetFeatures(val *AzurestackProviderFeatures)
	FeaturesInput() *AzurestackProviderFeatures
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	// Experimental.
	MetaAttributes() *map[string]interface{}
	MetadataHost() *string
	SetMetadataHost(val *string)
	MetadataHostInput() *string
	MsiEndpoint() *string
	SetMsiEndpoint(val *string)
	MsiEndpointInput() *string
	// The tree node.
	Node() constructs.Node
	// Experimental.
	RawOverrides() interface{}
	SkipProviderRegistration() interface{}
	SetSkipProviderRegistration(val interface{})
	SkipProviderRegistrationInput() interface{}
	SubscriptionId() *string
	SetSubscriptionId(val *string)
	SubscriptionIdInput() *string
	TenantId() *string
	SetTenantId(val *string)
	TenantIdInput() *string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformProviderSource() *string
	// Experimental.
	TerraformResourceType() *string
	UseMsi() interface{}
	SetUseMsi(val interface{})
	UseMsiInput() interface{}
	// Experimental.
	AddOverride(path *string, value interface{})
	// Overrides the auto-generated logical ID with a specific ID.
	// Experimental.
	OverrideLogicalId(newLogicalId *string)
	ResetAlias()
	ResetArmEndpoint()
	ResetAuxiliaryTenantIds()
	ResetClientCertificatePassword()
	ResetClientCertificatePath()
	ResetClientId()
	ResetClientSecret()
	ResetDisableCorrelationRequestId()
	ResetEnvironment()
	ResetMetadataHost()
	ResetMsiEndpoint()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetSkipProviderRegistration()
	ResetSubscriptionId()
	ResetTenantId()
	ResetUseMsi()
	SynthesizeAttributes() *map[string]interface{}
	SynthesizeHclAttributes() *map[string]interface{}
	// Experimental.
	ToHclTerraform() interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for AzurestackProvider
type jsiiProxy_AzurestackProvider struct {
	internal.Type__cdktfTerraformProvider
}

func (j *jsiiProxy_AzurestackProvider) Alias() *string {
	var returns *string
	_jsii_.Get(
		j,
		"alias",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AzurestackProvider) AliasInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"aliasInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AzurestackProvider) ArmEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"armEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AzurestackProvider) ArmEndpointInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"armEndpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AzurestackProvider) AuxiliaryTenantIds() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"auxiliaryTenantIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AzurestackProvider) AuxiliaryTenantIdsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"auxiliaryTenantIdsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AzurestackProvider) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AzurestackProvider) ClientCertificatePassword() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientCertificatePassword",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AzurestackProvider) ClientCertificatePasswordInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientCertificatePasswordInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AzurestackProvider) ClientCertificatePath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientCertificatePath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AzurestackProvider) ClientCertificatePathInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientCertificatePathInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AzurestackProvider) ClientId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AzurestackProvider) ClientIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AzurestackProvider) ClientSecret() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientSecret",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AzurestackProvider) ClientSecretInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientSecretInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AzurestackProvider) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AzurestackProvider) DisableCorrelationRequestId() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"disableCorrelationRequestId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AzurestackProvider) DisableCorrelationRequestIdInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"disableCorrelationRequestIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AzurestackProvider) Environment() *string {
	var returns *string
	_jsii_.Get(
		j,
		"environment",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AzurestackProvider) EnvironmentInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"environmentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AzurestackProvider) Features() *AzurestackProviderFeatures {
	var returns *AzurestackProviderFeatures
	_jsii_.Get(
		j,
		"features",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AzurestackProvider) FeaturesInput() *AzurestackProviderFeatures {
	var returns *AzurestackProviderFeatures
	_jsii_.Get(
		j,
		"featuresInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AzurestackProvider) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AzurestackProvider) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AzurestackProvider) MetaAttributes() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"metaAttributes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AzurestackProvider) MetadataHost() *string {
	var returns *string
	_jsii_.Get(
		j,
		"metadataHost",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AzurestackProvider) MetadataHostInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"metadataHostInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AzurestackProvider) MsiEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"msiEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AzurestackProvider) MsiEndpointInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"msiEndpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AzurestackProvider) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AzurestackProvider) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AzurestackProvider) SkipProviderRegistration() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"skipProviderRegistration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AzurestackProvider) SkipProviderRegistrationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"skipProviderRegistrationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AzurestackProvider) SubscriptionId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"subscriptionId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AzurestackProvider) SubscriptionIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"subscriptionIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AzurestackProvider) TenantId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tenantId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AzurestackProvider) TenantIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tenantIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AzurestackProvider) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AzurestackProvider) TerraformProviderSource() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformProviderSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AzurestackProvider) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AzurestackProvider) UseMsi() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"useMsi",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AzurestackProvider) UseMsiInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"useMsiInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/azurestack/1.0.0/docs azurestack} Resource.
func NewAzurestackProvider(scope constructs.Construct, id *string, config *AzurestackProviderConfig) AzurestackProvider {
	_init_.Initialize()

	if err := validateNewAzurestackProviderParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_AzurestackProvider{}

	_jsii_.Create(
		"@cdktf/provider-azurestack.provider.AzurestackProvider",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/azurestack/1.0.0/docs azurestack} Resource.
func NewAzurestackProvider_Override(a AzurestackProvider, scope constructs.Construct, id *string, config *AzurestackProviderConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurestack.provider.AzurestackProvider",
		[]interface{}{scope, id, config},
		a,
	)
}

func (j *jsiiProxy_AzurestackProvider)SetAlias(val *string) {
	_jsii_.Set(
		j,
		"alias",
		val,
	)
}

func (j *jsiiProxy_AzurestackProvider)SetArmEndpoint(val *string) {
	_jsii_.Set(
		j,
		"armEndpoint",
		val,
	)
}

func (j *jsiiProxy_AzurestackProvider)SetAuxiliaryTenantIds(val *[]*string) {
	_jsii_.Set(
		j,
		"auxiliaryTenantIds",
		val,
	)
}

func (j *jsiiProxy_AzurestackProvider)SetClientCertificatePassword(val *string) {
	_jsii_.Set(
		j,
		"clientCertificatePassword",
		val,
	)
}

func (j *jsiiProxy_AzurestackProvider)SetClientCertificatePath(val *string) {
	_jsii_.Set(
		j,
		"clientCertificatePath",
		val,
	)
}

func (j *jsiiProxy_AzurestackProvider)SetClientId(val *string) {
	_jsii_.Set(
		j,
		"clientId",
		val,
	)
}

func (j *jsiiProxy_AzurestackProvider)SetClientSecret(val *string) {
	_jsii_.Set(
		j,
		"clientSecret",
		val,
	)
}

func (j *jsiiProxy_AzurestackProvider)SetDisableCorrelationRequestId(val interface{}) {
	if err := j.validateSetDisableCorrelationRequestIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"disableCorrelationRequestId",
		val,
	)
}

func (j *jsiiProxy_AzurestackProvider)SetEnvironment(val *string) {
	_jsii_.Set(
		j,
		"environment",
		val,
	)
}

func (j *jsiiProxy_AzurestackProvider)SetFeatures(val *AzurestackProviderFeatures) {
	if err := j.validateSetFeaturesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"features",
		val,
	)
}

func (j *jsiiProxy_AzurestackProvider)SetMetadataHost(val *string) {
	_jsii_.Set(
		j,
		"metadataHost",
		val,
	)
}

func (j *jsiiProxy_AzurestackProvider)SetMsiEndpoint(val *string) {
	_jsii_.Set(
		j,
		"msiEndpoint",
		val,
	)
}

func (j *jsiiProxy_AzurestackProvider)SetSkipProviderRegistration(val interface{}) {
	if err := j.validateSetSkipProviderRegistrationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"skipProviderRegistration",
		val,
	)
}

func (j *jsiiProxy_AzurestackProvider)SetSubscriptionId(val *string) {
	_jsii_.Set(
		j,
		"subscriptionId",
		val,
	)
}

func (j *jsiiProxy_AzurestackProvider)SetTenantId(val *string) {
	_jsii_.Set(
		j,
		"tenantId",
		val,
	)
}

func (j *jsiiProxy_AzurestackProvider)SetUseMsi(val interface{}) {
	if err := j.validateSetUseMsiParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"useMsi",
		val,
	)
}

// Generates CDKTF code for importing a AzurestackProvider resource upon running "cdktf plan <stack-name>".
func AzurestackProvider_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateAzurestackProvider_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurestack.provider.AzurestackProvider",
		"generateConfigForImport",
		[]interface{}{scope, importToId, importFromId, provider},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Use this method instead of `instanceof` to properly detect `Construct`
// instances, even when the construct library is symlinked.
//
// Explanation: in JavaScript, multiple copies of the `constructs` library on
// disk are seen as independent, completely different libraries. As a
// consequence, the class `Construct` in each copy of the `constructs` library
// is seen as a different class, and an instance of one class will not test as
// `instanceof` the other class. `npm install` will not create installations
// like this, but users may manually symlink construct libraries together or
// use a monorepo tool: in those cases, multiple copies of the `constructs`
// library can be accidentally installed, and `instanceof` will behave
// unpredictably. It is safest to avoid using `instanceof`, and using
// this type-testing method instead.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
func AzurestackProvider_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAzurestackProvider_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurestack.provider.AzurestackProvider",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func AzurestackProvider_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAzurestackProvider_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurestack.provider.AzurestackProvider",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func AzurestackProvider_IsTerraformProvider(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAzurestackProvider_IsTerraformProviderParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurestack.provider.AzurestackProvider",
		"isTerraformProvider",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func AzurestackProvider_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-azurestack.provider.AzurestackProvider",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (a *jsiiProxy_AzurestackProvider) AddOverride(path *string, value interface{}) {
	if err := a.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (a *jsiiProxy_AzurestackProvider) OverrideLogicalId(newLogicalId *string) {
	if err := a.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (a *jsiiProxy_AzurestackProvider) ResetAlias() {
	_jsii_.InvokeVoid(
		a,
		"resetAlias",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AzurestackProvider) ResetArmEndpoint() {
	_jsii_.InvokeVoid(
		a,
		"resetArmEndpoint",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AzurestackProvider) ResetAuxiliaryTenantIds() {
	_jsii_.InvokeVoid(
		a,
		"resetAuxiliaryTenantIds",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AzurestackProvider) ResetClientCertificatePassword() {
	_jsii_.InvokeVoid(
		a,
		"resetClientCertificatePassword",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AzurestackProvider) ResetClientCertificatePath() {
	_jsii_.InvokeVoid(
		a,
		"resetClientCertificatePath",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AzurestackProvider) ResetClientId() {
	_jsii_.InvokeVoid(
		a,
		"resetClientId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AzurestackProvider) ResetClientSecret() {
	_jsii_.InvokeVoid(
		a,
		"resetClientSecret",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AzurestackProvider) ResetDisableCorrelationRequestId() {
	_jsii_.InvokeVoid(
		a,
		"resetDisableCorrelationRequestId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AzurestackProvider) ResetEnvironment() {
	_jsii_.InvokeVoid(
		a,
		"resetEnvironment",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AzurestackProvider) ResetMetadataHost() {
	_jsii_.InvokeVoid(
		a,
		"resetMetadataHost",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AzurestackProvider) ResetMsiEndpoint() {
	_jsii_.InvokeVoid(
		a,
		"resetMsiEndpoint",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AzurestackProvider) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		a,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AzurestackProvider) ResetSkipProviderRegistration() {
	_jsii_.InvokeVoid(
		a,
		"resetSkipProviderRegistration",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AzurestackProvider) ResetSubscriptionId() {
	_jsii_.InvokeVoid(
		a,
		"resetSubscriptionId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AzurestackProvider) ResetTenantId() {
	_jsii_.InvokeVoid(
		a,
		"resetTenantId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AzurestackProvider) ResetUseMsi() {
	_jsii_.InvokeVoid(
		a,
		"resetUseMsi",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AzurestackProvider) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AzurestackProvider) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AzurestackProvider) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AzurestackProvider) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AzurestackProvider) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AzurestackProvider) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

