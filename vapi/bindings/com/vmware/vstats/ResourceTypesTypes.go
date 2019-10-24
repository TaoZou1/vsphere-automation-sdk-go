/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Data type definitions file for service: ResourceTypes.
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

// Resource type for resource types.
const ResourceTypes_RESOURCE_TYPE = "com.vmware.vstats.model.RsrcType"
// The ``Summary`` class contains information of addressable resource.
type ResourceTypesSummary struct {
    // Resource type.
	Type_ string
}


func resourceTypesListInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func resourceTypesListOutputType() bindings.BindingType {
	return bindings.NewListType(bindings.NewReferenceType(ResourceTypesSummaryBindingType), reflect.TypeOf([]ResourceTypesSummary{}))
}

func resourceTypesListRestMetadata() protocol.OperationRestMetadata {
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
		"/stats/rsrc-types",
		resultHeaders,
		200,
		errorHeaders,
		map[string]int{"Error": 500,"Unauthenticated": 401,"Unauthorized": 403})
}


func ResourceTypesSummaryBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["type"] = bindings.NewIdType([]string{"com.vmware.vstats.model.RsrcType"}, "")
	fieldNameMap["type"] = "Type_"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vstats.resource_types.summary", fields, reflect.TypeOf(ResourceTypesSummary{}), fieldNameMap, validators)
}

