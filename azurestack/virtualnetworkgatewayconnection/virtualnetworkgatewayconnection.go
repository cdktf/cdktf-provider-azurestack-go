package virtualnetworkgatewayconnection

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"@cdktf/provider-azurestack.virtualNetworkGatewayConnection.VirtualNetworkGatewayConnection",
		reflect.TypeOf((*VirtualNetworkGatewayConnection)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberProperty{JsiiProperty: "authorizationKey", GoGetter: "AuthorizationKey"},
			_jsii_.MemberProperty{JsiiProperty: "authorizationKeyInput", GoGetter: "AuthorizationKeyInput"},
			_jsii_.MemberProperty{JsiiProperty: "cdktfStack", GoGetter: "CdktfStack"},
			_jsii_.MemberProperty{JsiiProperty: "connection", GoGetter: "Connection"},
			_jsii_.MemberProperty{JsiiProperty: "constructNodeMetadata", GoGetter: "ConstructNodeMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "count", GoGetter: "Count"},
			_jsii_.MemberProperty{JsiiProperty: "dependsOn", GoGetter: "DependsOn"},
			_jsii_.MemberProperty{JsiiProperty: "enableBgp", GoGetter: "EnableBgp"},
			_jsii_.MemberProperty{JsiiProperty: "enableBgpInput", GoGetter: "EnableBgpInput"},
			_jsii_.MemberProperty{JsiiProperty: "expressRouteCircuitId", GoGetter: "ExpressRouteCircuitId"},
			_jsii_.MemberProperty{JsiiProperty: "expressRouteCircuitIdInput", GoGetter: "ExpressRouteCircuitIdInput"},
			_jsii_.MemberProperty{JsiiProperty: "forEach", GoGetter: "ForEach"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberProperty{JsiiProperty: "friendlyUniqueId", GoGetter: "FriendlyUniqueId"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "id", GoGetter: "Id"},
			_jsii_.MemberProperty{JsiiProperty: "idInput", GoGetter: "IdInput"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "ipsecPolicy", GoGetter: "IpsecPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "ipsecPolicyInput", GoGetter: "IpsecPolicyInput"},
			_jsii_.MemberProperty{JsiiProperty: "lifecycle", GoGetter: "Lifecycle"},
			_jsii_.MemberProperty{JsiiProperty: "localNetworkGatewayId", GoGetter: "LocalNetworkGatewayId"},
			_jsii_.MemberProperty{JsiiProperty: "localNetworkGatewayIdInput", GoGetter: "LocalNetworkGatewayIdInput"},
			_jsii_.MemberProperty{JsiiProperty: "location", GoGetter: "Location"},
			_jsii_.MemberProperty{JsiiProperty: "locationInput", GoGetter: "LocationInput"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "nameInput", GoGetter: "NameInput"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "peerVirtualNetworkGatewayId", GoGetter: "PeerVirtualNetworkGatewayId"},
			_jsii_.MemberProperty{JsiiProperty: "peerVirtualNetworkGatewayIdInput", GoGetter: "PeerVirtualNetworkGatewayIdInput"},
			_jsii_.MemberProperty{JsiiProperty: "provider", GoGetter: "Provider"},
			_jsii_.MemberProperty{JsiiProperty: "provisioners", GoGetter: "Provisioners"},
			_jsii_.MemberMethod{JsiiMethod: "putIpsecPolicy", GoMethod: "PutIpsecPolicy"},
			_jsii_.MemberMethod{JsiiMethod: "putTimeouts", GoMethod: "PutTimeouts"},
			_jsii_.MemberProperty{JsiiProperty: "rawOverrides", GoGetter: "RawOverrides"},
			_jsii_.MemberMethod{JsiiMethod: "resetAuthorizationKey", GoMethod: "ResetAuthorizationKey"},
			_jsii_.MemberMethod{JsiiMethod: "resetEnableBgp", GoMethod: "ResetEnableBgp"},
			_jsii_.MemberMethod{JsiiMethod: "resetExpressRouteCircuitId", GoMethod: "ResetExpressRouteCircuitId"},
			_jsii_.MemberMethod{JsiiMethod: "resetId", GoMethod: "ResetId"},
			_jsii_.MemberMethod{JsiiMethod: "resetIpsecPolicy", GoMethod: "ResetIpsecPolicy"},
			_jsii_.MemberMethod{JsiiMethod: "resetLocalNetworkGatewayId", GoMethod: "ResetLocalNetworkGatewayId"},
			_jsii_.MemberMethod{JsiiMethod: "resetOverrideLogicalId", GoMethod: "ResetOverrideLogicalId"},
			_jsii_.MemberMethod{JsiiMethod: "resetPeerVirtualNetworkGatewayId", GoMethod: "ResetPeerVirtualNetworkGatewayId"},
			_jsii_.MemberMethod{JsiiMethod: "resetRoutingWeight", GoMethod: "ResetRoutingWeight"},
			_jsii_.MemberMethod{JsiiMethod: "resetSharedKey", GoMethod: "ResetSharedKey"},
			_jsii_.MemberMethod{JsiiMethod: "resetTags", GoMethod: "ResetTags"},
			_jsii_.MemberMethod{JsiiMethod: "resetTimeouts", GoMethod: "ResetTimeouts"},
			_jsii_.MemberMethod{JsiiMethod: "resetUsePolicyBasedTrafficSelectors", GoMethod: "ResetUsePolicyBasedTrafficSelectors"},
			_jsii_.MemberProperty{JsiiProperty: "resourceGroupName", GoGetter: "ResourceGroupName"},
			_jsii_.MemberProperty{JsiiProperty: "resourceGroupNameInput", GoGetter: "ResourceGroupNameInput"},
			_jsii_.MemberProperty{JsiiProperty: "routingWeight", GoGetter: "RoutingWeight"},
			_jsii_.MemberProperty{JsiiProperty: "routingWeightInput", GoGetter: "RoutingWeightInput"},
			_jsii_.MemberProperty{JsiiProperty: "sharedKey", GoGetter: "SharedKey"},
			_jsii_.MemberProperty{JsiiProperty: "sharedKeyInput", GoGetter: "SharedKeyInput"},
			_jsii_.MemberMethod{JsiiMethod: "synthesizeAttributes", GoMethod: "SynthesizeAttributes"},
			_jsii_.MemberProperty{JsiiProperty: "tags", GoGetter: "Tags"},
			_jsii_.MemberProperty{JsiiProperty: "tagsInput", GoGetter: "TagsInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformGeneratorMetadata", GoGetter: "TerraformGeneratorMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "terraformMetaArguments", GoGetter: "TerraformMetaArguments"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResourceType", GoGetter: "TerraformResourceType"},
			_jsii_.MemberProperty{JsiiProperty: "timeouts", GoGetter: "Timeouts"},
			_jsii_.MemberProperty{JsiiProperty: "timeoutsInput", GoGetter: "TimeoutsInput"},
			_jsii_.MemberMethod{JsiiMethod: "toMetadata", GoMethod: "ToMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "toTerraform", GoMethod: "ToTerraform"},
			_jsii_.MemberProperty{JsiiProperty: "type", GoGetter: "Type"},
			_jsii_.MemberProperty{JsiiProperty: "typeInput", GoGetter: "TypeInput"},
			_jsii_.MemberProperty{JsiiProperty: "usePolicyBasedTrafficSelectors", GoGetter: "UsePolicyBasedTrafficSelectors"},
			_jsii_.MemberProperty{JsiiProperty: "usePolicyBasedTrafficSelectorsInput", GoGetter: "UsePolicyBasedTrafficSelectorsInput"},
			_jsii_.MemberProperty{JsiiProperty: "virtualNetworkGatewayId", GoGetter: "VirtualNetworkGatewayId"},
			_jsii_.MemberProperty{JsiiProperty: "virtualNetworkGatewayIdInput", GoGetter: "VirtualNetworkGatewayIdInput"},
		},
		func() interface{} {
			j := jsiiProxy_VirtualNetworkGatewayConnection{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfTerraformResource)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-azurestack.virtualNetworkGatewayConnection.VirtualNetworkGatewayConnectionConfig",
		reflect.TypeOf((*VirtualNetworkGatewayConnectionConfig)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-azurestack.virtualNetworkGatewayConnection.VirtualNetworkGatewayConnectionIpsecPolicy",
		reflect.TypeOf((*VirtualNetworkGatewayConnectionIpsecPolicy)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-azurestack.virtualNetworkGatewayConnection.VirtualNetworkGatewayConnectionIpsecPolicyOutputReference",
		reflect.TypeOf((*VirtualNetworkGatewayConnectionIpsecPolicyOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "dhGroup", GoGetter: "DhGroup"},
			_jsii_.MemberProperty{JsiiProperty: "dhGroupInput", GoGetter: "DhGroupInput"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "ikeEncryption", GoGetter: "IkeEncryption"},
			_jsii_.MemberProperty{JsiiProperty: "ikeEncryptionInput", GoGetter: "IkeEncryptionInput"},
			_jsii_.MemberProperty{JsiiProperty: "ikeIntegrity", GoGetter: "IkeIntegrity"},
			_jsii_.MemberProperty{JsiiProperty: "ikeIntegrityInput", GoGetter: "IkeIntegrityInput"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "ipsecEncryption", GoGetter: "IpsecEncryption"},
			_jsii_.MemberProperty{JsiiProperty: "ipsecEncryptionInput", GoGetter: "IpsecEncryptionInput"},
			_jsii_.MemberProperty{JsiiProperty: "ipsecIntegrity", GoGetter: "IpsecIntegrity"},
			_jsii_.MemberProperty{JsiiProperty: "ipsecIntegrityInput", GoGetter: "IpsecIntegrityInput"},
			_jsii_.MemberProperty{JsiiProperty: "pfsGroup", GoGetter: "PfsGroup"},
			_jsii_.MemberProperty{JsiiProperty: "pfsGroupInput", GoGetter: "PfsGroupInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetSaDatasize", GoMethod: "ResetSaDatasize"},
			_jsii_.MemberMethod{JsiiMethod: "resetSaLifetime", GoMethod: "ResetSaLifetime"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "saDatasize", GoGetter: "SaDatasize"},
			_jsii_.MemberProperty{JsiiProperty: "saDatasizeInput", GoGetter: "SaDatasizeInput"},
			_jsii_.MemberProperty{JsiiProperty: "saLifetime", GoGetter: "SaLifetime"},
			_jsii_.MemberProperty{JsiiProperty: "saLifetimeInput", GoGetter: "SaLifetimeInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_VirtualNetworkGatewayConnectionIpsecPolicyOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-azurestack.virtualNetworkGatewayConnection.VirtualNetworkGatewayConnectionTimeouts",
		reflect.TypeOf((*VirtualNetworkGatewayConnectionTimeouts)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-azurestack.virtualNetworkGatewayConnection.VirtualNetworkGatewayConnectionTimeoutsOutputReference",
		reflect.TypeOf((*VirtualNetworkGatewayConnectionTimeoutsOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "create", GoGetter: "Create"},
			_jsii_.MemberProperty{JsiiProperty: "createInput", GoGetter: "CreateInput"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "delete", GoGetter: "Delete"},
			_jsii_.MemberProperty{JsiiProperty: "deleteInput", GoGetter: "DeleteInput"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "read", GoGetter: "Read"},
			_jsii_.MemberProperty{JsiiProperty: "readInput", GoGetter: "ReadInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetCreate", GoMethod: "ResetCreate"},
			_jsii_.MemberMethod{JsiiMethod: "resetDelete", GoMethod: "ResetDelete"},
			_jsii_.MemberMethod{JsiiMethod: "resetRead", GoMethod: "ResetRead"},
			_jsii_.MemberMethod{JsiiMethod: "resetUpdate", GoMethod: "ResetUpdate"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "update", GoGetter: "Update"},
			_jsii_.MemberProperty{JsiiProperty: "updateInput", GoGetter: "UpdateInput"},
		},
		func() interface{} {
			j := jsiiProxy_VirtualNetworkGatewayConnectionTimeoutsOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
}
