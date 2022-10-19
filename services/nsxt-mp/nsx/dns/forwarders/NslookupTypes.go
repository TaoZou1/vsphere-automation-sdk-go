// Copyright © 2019-2021 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Data type definitions file for service: Nslookup.
// Includes binding types of a structures and enumerations defined in the service.
// Shared by client-side stubs and server-side skeletons to ensure type
// compatibility.

package forwarders

import (
	"github.com/TaoZou1/vsphere-automation-sdk-go/runtime/bindings"
	"github.com/TaoZou1/vsphere-automation-sdk-go/runtime/data"
	"github.com/TaoZou1/vsphere-automation-sdk-go/runtime/protocol"
	"github.com/TaoZou1/vsphere-automation-sdk-go/services/nsxt-mp/nsx/model"
	"reflect"
)

func nslookupGetInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["forwarder_id"] = bindings.NewStringType()
	fields["address"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["server_ip"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["source_ip"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["forwarder_id"] = "ForwarderId"
	fieldNameMap["address"] = "Address"
	fieldNameMap["server_ip"] = "ServerIp"
	fieldNameMap["source_ip"] = "SourceIp"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func nslookupGetOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.DnsAnswerBindingType)
}

func nslookupGetRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["forwarder_id"] = bindings.NewStringType()
	fields["address"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["server_ip"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["source_ip"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["forwarder_id"] = "ForwarderId"
	fieldNameMap["address"] = "Address"
	fieldNameMap["server_ip"] = "ServerIp"
	fieldNameMap["source_ip"] = "SourceIp"
	paramsTypeMap["server_ip"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["source_ip"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["forwarder_id"] = bindings.NewStringType()
	paramsTypeMap["address"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["forwarderId"] = bindings.NewStringType()
	pathParams["forwarder_id"] = "forwarderId"
	queryParams["address"] = "address"
	queryParams["server_ip"] = "server_ip"
	queryParams["source_ip"] = "source_ip"
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
		"/api/v1/dns/forwarders/{forwarderId}/nslookup",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}