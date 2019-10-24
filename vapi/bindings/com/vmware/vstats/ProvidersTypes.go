/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Data type definitions file for service: Providers.
 * Includes binding types of a structures and enumerations defined in the service.
 * Shared by client-side stubs and server-side skeletons to ensure type
 * compatibility.
 */
package vstats

import (
	"reflect"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/bindings"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/data"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/protocol"
)

// Resource type for data providers.
const Providers_RESOURCE_TYPE = "com.vmware.vstats.model.Provider"


// ``Summary`` class describes a statistical data provider.
type ProvidersSummary struct {
    // Provider identifier.
	Id string
    // The ID given to the provider by its respective inventory.
	IdValue string
    // Provider type.
	Type_ string
    // An optional designation of the scheme.
	Scheme *string
    // Schema-less metadata with extra information for the provider.
	Metadata *string
    // Timestamp which is obtained when querying counters from a provider and is used as since parameter when new counter listing is needed.
	TrackingSn *int64
}



func providersListInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func providersListOutputType() bindings.BindingType {
	return bindings.NewListType(bindings.NewReferenceType(ProvidersSummaryBindingType), reflect.TypeOf([]ProvidersSummary{}))
}

func providersListRestMetadata() protocol.OperationRestMetadata {
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
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
		"/stats/providers",
		resultHeaders,
		200,
		errorHeaders,
		map[string]int{"Error": 500,"Unauthenticated": 401,"Unauthorized": 403})
}


func ProvidersSummaryBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["id"] = bindings.NewIdType([]string{"com.vmware.vstats.model.Provider"}, "")
	fieldNameMap["id"] = "Id"
	fields["id_value"] = bindings.NewStringType()
	fieldNameMap["id_value"] = "IdValue"
	fields["type"] = bindings.NewStringType()
	fieldNameMap["type"] = "Type_"
	fields["scheme"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["scheme"] = "Scheme"
	fields["metadata"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["metadata"] = "Metadata"
	fields["tracking_sn"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fieldNameMap["tracking_sn"] = "TrackingSn"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vstats.providers.summary", fields, reflect.TypeOf(ProvidersSummary{}), fieldNameMap, validators)
}

