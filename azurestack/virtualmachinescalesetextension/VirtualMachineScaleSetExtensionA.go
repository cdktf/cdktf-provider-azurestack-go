// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package virtualmachinescalesetextension

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-azurestack-go/azurestack/v6/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-azurestack-go/azurestack/v6/virtualmachinescalesetextension/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/azurestack/1.0.0/docs/resources/virtual_machine_scale_set_extension azurestack_virtual_machine_scale_set_extension}.
type VirtualMachineScaleSetExtensionA interface {
	cdktf.TerraformResource
	AutoUpgradeMinorVersion() interface{}
	SetAutoUpgradeMinorVersion(val interface{})
	AutoUpgradeMinorVersionInput() interface{}
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	// Experimental.
	Connection() interface{}
	// Experimental.
	SetConnection(val interface{})
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	// Experimental.
	Count() interface{}
	// Experimental.
	SetCount(val interface{})
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	ForceUpdateTag() *string
	SetForceUpdateTag(val *string)
	ForceUpdateTagInput() *string
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	Id() *string
	SetId(val *string)
	IdInput() *string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	Name() *string
	SetName(val *string)
	NameInput() *string
	// The tree node.
	Node() constructs.Node
	ProtectedSettings() *string
	SetProtectedSettings(val *string)
	ProtectedSettingsInput() *string
	// Experimental.
	Provider() cdktf.TerraformProvider
	// Experimental.
	SetProvider(val cdktf.TerraformProvider)
	// Experimental.
	Provisioners() *[]interface{}
	// Experimental.
	SetProvisioners(val *[]interface{})
	Publisher() *string
	SetPublisher(val *string)
	PublisherInput() *string
	// Experimental.
	RawOverrides() interface{}
	Settings() *string
	SetSettings(val *string)
	SettingsInput() *string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timeouts() VirtualMachineScaleSetExtensionTimeoutsOutputReference
	TimeoutsInput() interface{}
	Type() *string
	SetType(val *string)
	TypeHandlerVersion() *string
	SetTypeHandlerVersion(val *string)
	TypeHandlerVersionInput() *string
	TypeInput() *string
	VirtualMachineScaleSetId() *string
	SetVirtualMachineScaleSetId(val *string)
	VirtualMachineScaleSetIdInput() *string
	// Experimental.
	AddOverride(path *string, value interface{})
	// Experimental.
	GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{}
	// Experimental.
	GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable
	// Experimental.
	GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool
	// Experimental.
	GetListAttribute(terraformAttribute *string) *[]*string
	// Experimental.
	GetNumberAttribute(terraformAttribute *string) *float64
	// Experimental.
	GetNumberListAttribute(terraformAttribute *string) *[]*float64
	// Experimental.
	GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64
	// Experimental.
	GetStringAttribute(terraformAttribute *string) *string
	// Experimental.
	GetStringMapAttribute(terraformAttribute *string) *map[string]*string
	// Experimental.
	InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable
	// Overrides the auto-generated logical ID with a specific ID.
	// Experimental.
	OverrideLogicalId(newLogicalId *string)
	PutTimeouts(value *VirtualMachineScaleSetExtensionTimeouts)
	ResetAutoUpgradeMinorVersion()
	ResetForceUpdateTag()
	ResetId()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetProtectedSettings()
	ResetSettings()
	ResetTimeouts()
	SynthesizeAttributes() *map[string]interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for VirtualMachineScaleSetExtensionA
type jsiiProxy_VirtualMachineScaleSetExtensionA struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_VirtualMachineScaleSetExtensionA) AutoUpgradeMinorVersion() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"autoUpgradeMinorVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachineScaleSetExtensionA) AutoUpgradeMinorVersionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"autoUpgradeMinorVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachineScaleSetExtensionA) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachineScaleSetExtensionA) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachineScaleSetExtensionA) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachineScaleSetExtensionA) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachineScaleSetExtensionA) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachineScaleSetExtensionA) ForceUpdateTag() *string {
	var returns *string
	_jsii_.Get(
		j,
		"forceUpdateTag",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachineScaleSetExtensionA) ForceUpdateTagInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"forceUpdateTagInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachineScaleSetExtensionA) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachineScaleSetExtensionA) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachineScaleSetExtensionA) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachineScaleSetExtensionA) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachineScaleSetExtensionA) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachineScaleSetExtensionA) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachineScaleSetExtensionA) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachineScaleSetExtensionA) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachineScaleSetExtensionA) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachineScaleSetExtensionA) ProtectedSettings() *string {
	var returns *string
	_jsii_.Get(
		j,
		"protectedSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachineScaleSetExtensionA) ProtectedSettingsInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"protectedSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachineScaleSetExtensionA) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachineScaleSetExtensionA) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachineScaleSetExtensionA) Publisher() *string {
	var returns *string
	_jsii_.Get(
		j,
		"publisher",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachineScaleSetExtensionA) PublisherInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"publisherInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachineScaleSetExtensionA) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachineScaleSetExtensionA) Settings() *string {
	var returns *string
	_jsii_.Get(
		j,
		"settings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachineScaleSetExtensionA) SettingsInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"settingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachineScaleSetExtensionA) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachineScaleSetExtensionA) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachineScaleSetExtensionA) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachineScaleSetExtensionA) Timeouts() VirtualMachineScaleSetExtensionTimeoutsOutputReference {
	var returns VirtualMachineScaleSetExtensionTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachineScaleSetExtensionA) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachineScaleSetExtensionA) Type() *string {
	var returns *string
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachineScaleSetExtensionA) TypeHandlerVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"typeHandlerVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachineScaleSetExtensionA) TypeHandlerVersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"typeHandlerVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachineScaleSetExtensionA) TypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"typeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachineScaleSetExtensionA) VirtualMachineScaleSetId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"virtualMachineScaleSetId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachineScaleSetExtensionA) VirtualMachineScaleSetIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"virtualMachineScaleSetIdInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/azurestack/1.0.0/docs/resources/virtual_machine_scale_set_extension azurestack_virtual_machine_scale_set_extension} Resource.
func NewVirtualMachineScaleSetExtensionA(scope constructs.Construct, id *string, config *VirtualMachineScaleSetExtensionAConfig) VirtualMachineScaleSetExtensionA {
	_init_.Initialize()

	if err := validateNewVirtualMachineScaleSetExtensionAParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_VirtualMachineScaleSetExtensionA{}

	_jsii_.Create(
		"@cdktf/provider-azurestack.virtualMachineScaleSetExtension.VirtualMachineScaleSetExtensionA",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/azurestack/1.0.0/docs/resources/virtual_machine_scale_set_extension azurestack_virtual_machine_scale_set_extension} Resource.
func NewVirtualMachineScaleSetExtensionA_Override(v VirtualMachineScaleSetExtensionA, scope constructs.Construct, id *string, config *VirtualMachineScaleSetExtensionAConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurestack.virtualMachineScaleSetExtension.VirtualMachineScaleSetExtensionA",
		[]interface{}{scope, id, config},
		v,
	)
}

func (j *jsiiProxy_VirtualMachineScaleSetExtensionA)SetAutoUpgradeMinorVersion(val interface{}) {
	if err := j.validateSetAutoUpgradeMinorVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"autoUpgradeMinorVersion",
		val,
	)
}

func (j *jsiiProxy_VirtualMachineScaleSetExtensionA)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_VirtualMachineScaleSetExtensionA)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_VirtualMachineScaleSetExtensionA)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_VirtualMachineScaleSetExtensionA)SetForceUpdateTag(val *string) {
	if err := j.validateSetForceUpdateTagParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"forceUpdateTag",
		val,
	)
}

func (j *jsiiProxy_VirtualMachineScaleSetExtensionA)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_VirtualMachineScaleSetExtensionA)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_VirtualMachineScaleSetExtensionA)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_VirtualMachineScaleSetExtensionA)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_VirtualMachineScaleSetExtensionA)SetProtectedSettings(val *string) {
	if err := j.validateSetProtectedSettingsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"protectedSettings",
		val,
	)
}

func (j *jsiiProxy_VirtualMachineScaleSetExtensionA)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_VirtualMachineScaleSetExtensionA)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_VirtualMachineScaleSetExtensionA)SetPublisher(val *string) {
	if err := j.validateSetPublisherParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"publisher",
		val,
	)
}

func (j *jsiiProxy_VirtualMachineScaleSetExtensionA)SetSettings(val *string) {
	if err := j.validateSetSettingsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"settings",
		val,
	)
}

func (j *jsiiProxy_VirtualMachineScaleSetExtensionA)SetType(val *string) {
	if err := j.validateSetTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"type",
		val,
	)
}

func (j *jsiiProxy_VirtualMachineScaleSetExtensionA)SetTypeHandlerVersion(val *string) {
	if err := j.validateSetTypeHandlerVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"typeHandlerVersion",
		val,
	)
}

func (j *jsiiProxy_VirtualMachineScaleSetExtensionA)SetVirtualMachineScaleSetId(val *string) {
	if err := j.validateSetVirtualMachineScaleSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"virtualMachineScaleSetId",
		val,
	)
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
func VirtualMachineScaleSetExtensionA_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateVirtualMachineScaleSetExtensionA_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurestack.virtualMachineScaleSetExtension.VirtualMachineScaleSetExtensionA",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func VirtualMachineScaleSetExtensionA_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateVirtualMachineScaleSetExtensionA_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurestack.virtualMachineScaleSetExtension.VirtualMachineScaleSetExtensionA",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func VirtualMachineScaleSetExtensionA_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateVirtualMachineScaleSetExtensionA_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurestack.virtualMachineScaleSetExtension.VirtualMachineScaleSetExtensionA",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func VirtualMachineScaleSetExtensionA_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-azurestack.virtualMachineScaleSetExtension.VirtualMachineScaleSetExtensionA",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (v *jsiiProxy_VirtualMachineScaleSetExtensionA) AddOverride(path *string, value interface{}) {
	if err := v.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (v *jsiiProxy_VirtualMachineScaleSetExtensionA) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := v.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		v,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VirtualMachineScaleSetExtensionA) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := v.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		v,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VirtualMachineScaleSetExtensionA) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := v.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		v,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VirtualMachineScaleSetExtensionA) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := v.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		v,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VirtualMachineScaleSetExtensionA) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := v.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		v,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VirtualMachineScaleSetExtensionA) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := v.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		v,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VirtualMachineScaleSetExtensionA) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := v.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		v,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VirtualMachineScaleSetExtensionA) GetStringAttribute(terraformAttribute *string) *string {
	if err := v.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		v,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VirtualMachineScaleSetExtensionA) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := v.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		v,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VirtualMachineScaleSetExtensionA) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := v.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		v,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VirtualMachineScaleSetExtensionA) OverrideLogicalId(newLogicalId *string) {
	if err := v.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (v *jsiiProxy_VirtualMachineScaleSetExtensionA) PutTimeouts(value *VirtualMachineScaleSetExtensionTimeouts) {
	if err := v.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (v *jsiiProxy_VirtualMachineScaleSetExtensionA) ResetAutoUpgradeMinorVersion() {
	_jsii_.InvokeVoid(
		v,
		"resetAutoUpgradeMinorVersion",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VirtualMachineScaleSetExtensionA) ResetForceUpdateTag() {
	_jsii_.InvokeVoid(
		v,
		"resetForceUpdateTag",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VirtualMachineScaleSetExtensionA) ResetId() {
	_jsii_.InvokeVoid(
		v,
		"resetId",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VirtualMachineScaleSetExtensionA) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		v,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VirtualMachineScaleSetExtensionA) ResetProtectedSettings() {
	_jsii_.InvokeVoid(
		v,
		"resetProtectedSettings",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VirtualMachineScaleSetExtensionA) ResetSettings() {
	_jsii_.InvokeVoid(
		v,
		"resetSettings",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VirtualMachineScaleSetExtensionA) ResetTimeouts() {
	_jsii_.InvokeVoid(
		v,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VirtualMachineScaleSetExtensionA) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		v,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VirtualMachineScaleSetExtensionA) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		v,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VirtualMachineScaleSetExtensionA) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		v,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VirtualMachineScaleSetExtensionA) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		v,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

