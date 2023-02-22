// Copyright © 2019-2021 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Data type definitions file for service: Tunnels.
// Includes binding types of a structures and enumerations defined in the service.
// Shared by client-side stubs and server-side skeletons to ensure type
// compatibility.

package locale_services

import (
	"github.com/vmware/vsphere-automation-sdk-go/runtime/bindings"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/data"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/protocol"
	"github.com/vmware/vsphere-automation-sdk-go/services/nsxt/model"
	"reflect"
)

func tunnelsDeleteInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["tier0_id"] = bindings.NewStringType()
	fields["locale_svc_id"] = bindings.NewStringType()
	fields["tunnel_id"] = bindings.NewStringType()
	fieldNameMap["tier0_id"] = "Tier0Id"
	fieldNameMap["locale_svc_id"] = "LocaleSvcId"
	fieldNameMap["tunnel_id"] = "TunnelId"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func tunnelsDeleteOutputType() bindings.BindingType {
	return bindings.NewVoidType()
}

func tunnelsDeleteRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["tier0_id"] = bindings.NewStringType()
	fields["locale_svc_id"] = bindings.NewStringType()
	fields["tunnel_id"] = bindings.NewStringType()
	fieldNameMap["tier0_id"] = "Tier0Id"
	fieldNameMap["locale_svc_id"] = "LocaleSvcId"
	fieldNameMap["tunnel_id"] = "TunnelId"
	paramsTypeMap["tier0_id"] = bindings.NewStringType()
	paramsTypeMap["tunnel_id"] = bindings.NewStringType()
	paramsTypeMap["locale_svc_id"] = bindings.NewStringType()
	paramsTypeMap["tier0Id"] = bindings.NewStringType()
	paramsTypeMap["localeSvcId"] = bindings.NewStringType()
	paramsTypeMap["tunnelId"] = bindings.NewStringType()
	pathParams["tunnel_id"] = "tunnelId"
	pathParams["tier0_id"] = "tier0Id"
	pathParams["locale_svc_id"] = "localeSvcId"
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
		"DELETE",
		"/policy/api/v1/infra/tier-0s/{tier0Id}/locale-services/{localeSvcId}/tunnels/{tunnelId}",
		"",
		resultHeaders,
		204,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func tunnelsGetInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["tier0_id"] = bindings.NewStringType()
	fields["locale_svc_id"] = bindings.NewStringType()
	fields["tunnel_id"] = bindings.NewStringType()
	fieldNameMap["tier0_id"] = "Tier0Id"
	fieldNameMap["locale_svc_id"] = "LocaleSvcId"
	fieldNameMap["tunnel_id"] = "TunnelId"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func tunnelsGetOutputType() bindings.BindingType {
	return bindings.NewDynamicStructType([]bindings.ReferenceType{bindings.NewReferenceType(model.TunnelBindingType)}, bindings.REST)
}

func tunnelsGetRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["tier0_id"] = bindings.NewStringType()
	fields["locale_svc_id"] = bindings.NewStringType()
	fields["tunnel_id"] = bindings.NewStringType()
	fieldNameMap["tier0_id"] = "Tier0Id"
	fieldNameMap["locale_svc_id"] = "LocaleSvcId"
	fieldNameMap["tunnel_id"] = "TunnelId"
	paramsTypeMap["tier0_id"] = bindings.NewStringType()
	paramsTypeMap["tunnel_id"] = bindings.NewStringType()
	paramsTypeMap["locale_svc_id"] = bindings.NewStringType()
	paramsTypeMap["tier0Id"] = bindings.NewStringType()
	paramsTypeMap["localeSvcId"] = bindings.NewStringType()
	paramsTypeMap["tunnelId"] = bindings.NewStringType()
	pathParams["tunnel_id"] = "tunnelId"
	pathParams["tier0_id"] = "tier0Id"
	pathParams["locale_svc_id"] = "localeSvcId"
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
		"/policy/api/v1/infra/tier-0s/{tier0Id}/locale-services/{localeSvcId}/tunnels/{tunnelId}",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func tunnelsListInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["tier0_id"] = bindings.NewStringType()
	fields["locale_svc_id"] = bindings.NewStringType()
	fields["cursor"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["included_fields"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["page_size"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fields["sort_ascending"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fields["sort_by"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["tier0_id"] = "Tier0Id"
	fieldNameMap["locale_svc_id"] = "LocaleSvcId"
	fieldNameMap["cursor"] = "Cursor"
	fieldNameMap["included_fields"] = "IncludedFields"
	fieldNameMap["page_size"] = "PageSize"
	fieldNameMap["sort_ascending"] = "SortAscending"
	fieldNameMap["sort_by"] = "SortBy"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func tunnelsListOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.PolicyTunnelListResultBindingType)
}

func tunnelsListRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["tier0_id"] = bindings.NewStringType()
	fields["locale_svc_id"] = bindings.NewStringType()
	fields["cursor"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["included_fields"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["page_size"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fields["sort_ascending"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fields["sort_by"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["tier0_id"] = "Tier0Id"
	fieldNameMap["locale_svc_id"] = "LocaleSvcId"
	fieldNameMap["cursor"] = "Cursor"
	fieldNameMap["included_fields"] = "IncludedFields"
	fieldNameMap["page_size"] = "PageSize"
	fieldNameMap["sort_ascending"] = "SortAscending"
	fieldNameMap["sort_by"] = "SortBy"
	paramsTypeMap["tier0_id"] = bindings.NewStringType()
	paramsTypeMap["included_fields"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["page_size"] = bindings.NewOptionalType(bindings.NewIntegerType())
	paramsTypeMap["locale_svc_id"] = bindings.NewStringType()
	paramsTypeMap["cursor"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["sort_by"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["sort_ascending"] = bindings.NewOptionalType(bindings.NewBooleanType())
	paramsTypeMap["tier0Id"] = bindings.NewStringType()
	paramsTypeMap["localeSvcId"] = bindings.NewStringType()
	pathParams["tier0_id"] = "tier0Id"
	pathParams["locale_svc_id"] = "localeSvcId"
	queryParams["cursor"] = "cursor"
	queryParams["sort_ascending"] = "sort_ascending"
	queryParams["included_fields"] = "included_fields"
	queryParams["sort_by"] = "sort_by"
	queryParams["page_size"] = "page_size"
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
		"/policy/api/v1/infra/tier-0s/{tier0Id}/locale-services/{localeSvcId}/tunnels",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func tunnelsPatchInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["tier0_id"] = bindings.NewStringType()
	fields["locale_svc_id"] = bindings.NewStringType()
	fields["tunnel_id"] = bindings.NewStringType()
	fields["tunnel"] = bindings.NewDynamicStructType([]bindings.ReferenceType{bindings.NewReferenceType(model.TunnelBindingType)}, bindings.REST)
	fieldNameMap["tier0_id"] = "Tier0Id"
	fieldNameMap["locale_svc_id"] = "LocaleSvcId"
	fieldNameMap["tunnel_id"] = "TunnelId"
	fieldNameMap["tunnel"] = "Tunnel"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func tunnelsPatchOutputType() bindings.BindingType {
	return bindings.NewVoidType()
}

func tunnelsPatchRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["tier0_id"] = bindings.NewStringType()
	fields["locale_svc_id"] = bindings.NewStringType()
	fields["tunnel_id"] = bindings.NewStringType()
	fields["tunnel"] = bindings.NewDynamicStructType([]bindings.ReferenceType{bindings.NewReferenceType(model.TunnelBindingType)}, bindings.REST)
	fieldNameMap["tier0_id"] = "Tier0Id"
	fieldNameMap["locale_svc_id"] = "LocaleSvcId"
	fieldNameMap["tunnel_id"] = "TunnelId"
	fieldNameMap["tunnel"] = "Tunnel"
	paramsTypeMap["tier0_id"] = bindings.NewStringType()
	paramsTypeMap["tunnel_id"] = bindings.NewStringType()
	paramsTypeMap["tunnel"] = bindings.NewDynamicStructType([]bindings.ReferenceType{bindings.NewReferenceType(model.TunnelBindingType)}, bindings.REST)
	paramsTypeMap["locale_svc_id"] = bindings.NewStringType()
	paramsTypeMap["tier0Id"] = bindings.NewStringType()
	paramsTypeMap["localeSvcId"] = bindings.NewStringType()
	paramsTypeMap["tunnelId"] = bindings.NewStringType()
	pathParams["tunnel_id"] = "tunnelId"
	pathParams["tier0_id"] = "tier0Id"
	pathParams["locale_svc_id"] = "localeSvcId"
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
		"tunnel",
		"PATCH",
		"/policy/api/v1/infra/tier-0s/{tier0Id}/locale-services/{localeSvcId}/tunnels/{tunnelId}",
		"",
		resultHeaders,
		204,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func tunnelsUpdateInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["tier0_id"] = bindings.NewStringType()
	fields["locale_svc_id"] = bindings.NewStringType()
	fields["tunnel_id"] = bindings.NewStringType()
	fields["tunnel"] = bindings.NewDynamicStructType([]bindings.ReferenceType{bindings.NewReferenceType(model.TunnelBindingType)}, bindings.REST)
	fieldNameMap["tier0_id"] = "Tier0Id"
	fieldNameMap["locale_svc_id"] = "LocaleSvcId"
	fieldNameMap["tunnel_id"] = "TunnelId"
	fieldNameMap["tunnel"] = "Tunnel"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func tunnelsUpdateOutputType() bindings.BindingType {
	return bindings.NewDynamicStructType([]bindings.ReferenceType{bindings.NewReferenceType(model.TunnelBindingType)}, bindings.REST)
}

func tunnelsUpdateRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["tier0_id"] = bindings.NewStringType()
	fields["locale_svc_id"] = bindings.NewStringType()
	fields["tunnel_id"] = bindings.NewStringType()
	fields["tunnel"] = bindings.NewDynamicStructType([]bindings.ReferenceType{bindings.NewReferenceType(model.TunnelBindingType)}, bindings.REST)
	fieldNameMap["tier0_id"] = "Tier0Id"
	fieldNameMap["locale_svc_id"] = "LocaleSvcId"
	fieldNameMap["tunnel_id"] = "TunnelId"
	fieldNameMap["tunnel"] = "Tunnel"
	paramsTypeMap["tier0_id"] = bindings.NewStringType()
	paramsTypeMap["tunnel_id"] = bindings.NewStringType()
	paramsTypeMap["tunnel"] = bindings.NewDynamicStructType([]bindings.ReferenceType{bindings.NewReferenceType(model.TunnelBindingType)}, bindings.REST)
	paramsTypeMap["locale_svc_id"] = bindings.NewStringType()
	paramsTypeMap["tier0Id"] = bindings.NewStringType()
	paramsTypeMap["localeSvcId"] = bindings.NewStringType()
	paramsTypeMap["tunnelId"] = bindings.NewStringType()
	pathParams["tunnel_id"] = "tunnelId"
	pathParams["tier0_id"] = "tier0Id"
	pathParams["locale_svc_id"] = "localeSvcId"
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
		"tunnel",
		"PUT",
		"/policy/api/v1/infra/tier-0s/{tier0Id}/locale-services/{localeSvcId}/tunnels/{tunnelId}",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}
