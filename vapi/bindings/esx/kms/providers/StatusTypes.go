/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Code generated. DO NOT EDIT.

/*
 * Data type definitions file for service: Status.
 * Includes binding types of a structures and enumerations defined in the service.
 * Shared by client-side stubs and server-side skeletons to ensure type
 * compatibility.
 */

package providers

import (
	"reflect"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/bindings/vapi/std"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/bindings"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/data"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/protocol"
)



// The ``Health`` enumeration class defines the possible health states.
//
// <p> See {@link com.vmware.vapi.bindings.ApiEnumeration enumerated types description}.
type StatusHealth string

const (
    // NONE. 
    //
    //  No status available.
	StatusHealth_NONE StatusHealth = "NONE"
    // OK. 
    //
    //  Health is normal.
	StatusHealth_OK StatusHealth = "OK"
    // WARNING 
    //
    //  Health is normal, however there is an issue that requires attention.
	StatusHealth_WARNING StatusHealth = "WARNING"
    // ERROR 
    //
    //  Not healthy.
	StatusHealth_ERROR StatusHealth = "ERROR"
)

func (h StatusHealth) StatusHealth() bool {
	switch h {
	case StatusHealth_NONE:
		return true
	case StatusHealth_OK:
		return true
	case StatusHealth_WARNING:
		return true
	case StatusHealth_ERROR:
		return true
	default:
		return false
	}
}


// The ``ServerInfo`` class contains properties that describe the status of a key server.
type StatusServerInfo struct {
    // The connection status health of the server.
	Health StatusHealth
    // Details regarding the health of the server connection. 
    //
    //  When the service ``Health`` is not StatusHealth#StatusHealth_OK, this property will provide an actionable description of the issue.
	Details []std.LocalizableMessage
    // Whether this client trusts the server.
	ClientTrustServer bool
    // Whether the server trusts this client.
	ServerTrustClient bool
    // Name of the server.
	Name string
}

// The ``Info`` class contains properties that describe the status of the Key Provider.
type StatusInfo struct {
    // The health of the provider.
	Health StatusHealth
    // Details regarding the health of the provider. 
    //
    //  When the service ``Health`` is not StatusHealth#StatusHealth_OK, this property will provide an actionable description of the issue.
	Details []std.LocalizableMessage
    // Health of key servers.
	Servers []StatusServerInfo
}



func statusGetInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["provider"] = bindings.NewIdType([]string{"com.vmware.esx.kms.providers"}, "")
	fieldNameMap["provider"] = "Provider"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func statusGetOutputType() bindings.BindingType {
	return bindings.NewReferenceType(StatusInfoBindingType)
}

func statusGetRestMetadata() protocol.OperationRestMetadata {
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	paramsTypeMap["provider"] = bindings.NewIdType([]string{"com.vmware.esx.kms.providers"}, "")
	paramsTypeMap["provider"] = bindings.NewIdType([]string{"com.vmware.esx.kms.providers"}, "")
	pathParams["provider"] = "provider"
	resultHeaders := map[string]string{}
	errorHeaders := map[string]string{}
	errorHeaders["Unauthenticated.challenge"] = "WWW-Authenticate"
	return protocol.NewOperationRestMetadata(
		paramsTypeMap,
		pathParams,
		queryParams,
		headerParams,
		"",
		"GET",
		"/esx/kms/providers/{provider}/status",
		resultHeaders,
		200,
		errorHeaders,
		map[string]int{"InvalidArgument": 400,"NotFound": 404,"Unauthenticated": 401,"Unauthorized": 403,"Error": 500})
}


func StatusServerInfoBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["health"] = bindings.NewEnumType("com.vmware.esx.kms.providers.status.health", reflect.TypeOf(StatusHealth(StatusHealth_NONE)))
	fieldNameMap["health"] = "Health"
	fields["details"] = bindings.NewListType(bindings.NewReferenceType(std.LocalizableMessageBindingType), reflect.TypeOf([]std.LocalizableMessage{}))
	fieldNameMap["details"] = "Details"
	fields["client_trust_server"] = bindings.NewBooleanType()
	fieldNameMap["client_trust_server"] = "ClientTrustServer"
	fields["server_trust_client"] = bindings.NewBooleanType()
	fieldNameMap["server_trust_client"] = "ServerTrustClient"
	fields["name"] = bindings.NewStringType()
	fieldNameMap["name"] = "Name"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.esx.kms.providers.status.server_info", fields, reflect.TypeOf(StatusServerInfo{}), fieldNameMap, validators)
}

func StatusInfoBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["health"] = bindings.NewEnumType("com.vmware.esx.kms.providers.status.health", reflect.TypeOf(StatusHealth(StatusHealth_NONE)))
	fieldNameMap["health"] = "Health"
	fields["details"] = bindings.NewListType(bindings.NewReferenceType(std.LocalizableMessageBindingType), reflect.TypeOf([]std.LocalizableMessage{}))
	fieldNameMap["details"] = "Details"
	fields["servers"] = bindings.NewListType(bindings.NewReferenceType(StatusServerInfoBindingType), reflect.TypeOf([]StatusServerInfo{}))
	fieldNameMap["servers"] = "Servers"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.esx.kms.providers.status.info", fields, reflect.TypeOf(StatusInfo{}), fieldNameMap, validators)
}

