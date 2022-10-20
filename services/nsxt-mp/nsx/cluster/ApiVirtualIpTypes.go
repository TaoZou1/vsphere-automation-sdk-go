// Copyright © 2019-2021 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Data type definitions file for service: ApiVirtualIp.
// Includes binding types of a structures and enumerations defined in the service.
// Shared by client-side stubs and server-side skeletons to ensure type
// compatibility.

package cluster

import (
	"github.com/TaoZou1/vsphere-automation-sdk-go/runtime/bindings"
	"github.com/TaoZou1/vsphere-automation-sdk-go/runtime/data"
	"github.com/TaoZou1/vsphere-automation-sdk-go/runtime/protocol"
	"github.com/TaoZou1/vsphere-automation-sdk-go/services/nsxt-mp/nsx/model"
	"reflect"
)

// Possible value for ``force`` of method ApiVirtualIp#setvirtualip.
const ApiVirtualIp_SETVIRTUALIP_FORCE_TRUE = "true"

// Possible value for ``force`` of method ApiVirtualIp#setvirtualip.
const ApiVirtualIp_SETVIRTUALIP_FORCE_FALSE = "false"

func apiVirtualIpClearvirtualipInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func apiVirtualIpClearvirtualipOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.ClusterVirtualIpPropertiesBindingType)
}

func apiVirtualIpClearvirtualipRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	resultHeaders := map[string]string{}
	errorHeaders := map[string]map[string]string{}
	return protocol.NewOperationRestMetadata(
		fields,
		fieldNameMap,
		paramsTypeMap,
		pathParams,
		queryParams,
		headerParams,
		dispatchHeaderParams,
		bodyFieldsMap,
		"action=clear_virtual_ip",
		"",
		"POST",
		"/api/v1/cluster/api-virtual-ip",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func apiVirtualIpClearvirtualip6InputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func apiVirtualIpClearvirtualip6OutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.ClusterVirtualIpPropertiesBindingType)
}

func apiVirtualIpClearvirtualip6RestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	resultHeaders := map[string]string{}
	errorHeaders := map[string]map[string]string{}
	return protocol.NewOperationRestMetadata(
		fields,
		fieldNameMap,
		paramsTypeMap,
		pathParams,
		queryParams,
		headerParams,
		dispatchHeaderParams,
		bodyFieldsMap,
		"action=clear_virtual_ip6",
		"",
		"POST",
		"/api/v1/cluster/api-virtual-ip",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func apiVirtualIpGetInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func apiVirtualIpGetOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.ClusterVirtualIpPropertiesBindingType)
}

func apiVirtualIpGetRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	resultHeaders := map[string]string{}
	errorHeaders := map[string]map[string]string{}
	return protocol.NewOperationRestMetadata(
		fields,
		fieldNameMap,
		paramsTypeMap,
		pathParams,
		queryParams,
		headerParams,
		dispatchHeaderParams,
		bodyFieldsMap,
		"",
		"",
		"GET",
		"/api/v1/cluster/api-virtual-ip",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func apiVirtualIpSetvirtualipInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["force"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["ip6_address"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["ip_address"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["force"] = "Force"
	fieldNameMap["ip6_address"] = "Ip6Address"
	fieldNameMap["ip_address"] = "IpAddress"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func apiVirtualIpSetvirtualipOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.ClusterVirtualIpPropertiesBindingType)
}

func apiVirtualIpSetvirtualipRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["force"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["ip6_address"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["ip_address"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["force"] = "Force"
	fieldNameMap["ip6_address"] = "Ip6Address"
	fieldNameMap["ip_address"] = "IpAddress"
	paramsTypeMap["force"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["ip_address"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["ip6_address"] = bindings.NewOptionalType(bindings.NewStringType())
	queryParams["ip6_address"] = "ip6_address"
	queryParams["force"] = "force"
	queryParams["ip_address"] = "ip_address"
	resultHeaders := map[string]string{}
	errorHeaders := map[string]map[string]string{}
	return protocol.NewOperationRestMetadata(
		fields,
		fieldNameMap,
		paramsTypeMap,
		pathParams,
		queryParams,
		headerParams,
		dispatchHeaderParams,
		bodyFieldsMap,
		"action=set_virtual_ip",
		"",
		"POST",
		"/api/v1/cluster/api-virtual-ip",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.concurrent_change": 409, "com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}
