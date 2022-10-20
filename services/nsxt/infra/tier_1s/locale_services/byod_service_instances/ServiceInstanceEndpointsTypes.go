// Copyright © 2019-2021 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Data type definitions file for service: ServiceInstanceEndpoints.
// Includes binding types of a structures and enumerations defined in the service.
// Shared by client-side stubs and server-side skeletons to ensure type
// compatibility.

package byod_service_instances

import (
	"github.com/vmware/vsphere-automation-sdk-go/runtime/bindings"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/data"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/protocol"
	"github.com/vmware/vsphere-automation-sdk-go/services/nsxt/model"
	"reflect"
)

func serviceInstanceEndpointsDeleteInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["tier1_id"] = bindings.NewStringType()
	fields["locale_service_id"] = bindings.NewStringType()
	fields["service_instance_id"] = bindings.NewStringType()
	fields["service_instance_endpoint_id"] = bindings.NewStringType()
	fieldNameMap["tier1_id"] = "Tier1Id"
	fieldNameMap["locale_service_id"] = "LocaleServiceId"
	fieldNameMap["service_instance_id"] = "ServiceInstanceId"
	fieldNameMap["service_instance_endpoint_id"] = "ServiceInstanceEndpointId"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func serviceInstanceEndpointsDeleteOutputType() bindings.BindingType {
	return bindings.NewVoidType()
}

func serviceInstanceEndpointsDeleteRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["tier1_id"] = bindings.NewStringType()
	fields["locale_service_id"] = bindings.NewStringType()
	fields["service_instance_id"] = bindings.NewStringType()
	fields["service_instance_endpoint_id"] = bindings.NewStringType()
	fieldNameMap["tier1_id"] = "Tier1Id"
	fieldNameMap["locale_service_id"] = "LocaleServiceId"
	fieldNameMap["service_instance_id"] = "ServiceInstanceId"
	fieldNameMap["service_instance_endpoint_id"] = "ServiceInstanceEndpointId"
	paramsTypeMap["locale_service_id"] = bindings.NewStringType()
	paramsTypeMap["tier1_id"] = bindings.NewStringType()
	paramsTypeMap["service_instance_id"] = bindings.NewStringType()
	paramsTypeMap["service_instance_endpoint_id"] = bindings.NewStringType()
	paramsTypeMap["tier1Id"] = bindings.NewStringType()
	paramsTypeMap["localeServiceId"] = bindings.NewStringType()
	paramsTypeMap["serviceInstanceId"] = bindings.NewStringType()
	paramsTypeMap["serviceInstanceEndpointId"] = bindings.NewStringType()
	pathParams["service_instance_endpoint_id"] = "serviceInstanceEndpointId"
	pathParams["tier1_id"] = "tier1Id"
	pathParams["service_instance_id"] = "serviceInstanceId"
	pathParams["locale_service_id"] = "localeServiceId"
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
		"/policy/api/v1/infra/tier-1s/{tier1Id}/locale-services/{localeServiceId}/byod-service-instances/{serviceInstanceId}/service-instance-endpoints/{serviceInstanceEndpointId}",
		"",
		resultHeaders,
		204,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func serviceInstanceEndpointsGetInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["tier1_id"] = bindings.NewStringType()
	fields["locale_service_id"] = bindings.NewStringType()
	fields["service_instance_id"] = bindings.NewStringType()
	fields["service_instance_endpoint_id"] = bindings.NewStringType()
	fieldNameMap["tier1_id"] = "Tier1Id"
	fieldNameMap["locale_service_id"] = "LocaleServiceId"
	fieldNameMap["service_instance_id"] = "ServiceInstanceId"
	fieldNameMap["service_instance_endpoint_id"] = "ServiceInstanceEndpointId"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func serviceInstanceEndpointsGetOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.ServiceInstanceEndpointBindingType)
}

func serviceInstanceEndpointsGetRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["tier1_id"] = bindings.NewStringType()
	fields["locale_service_id"] = bindings.NewStringType()
	fields["service_instance_id"] = bindings.NewStringType()
	fields["service_instance_endpoint_id"] = bindings.NewStringType()
	fieldNameMap["tier1_id"] = "Tier1Id"
	fieldNameMap["locale_service_id"] = "LocaleServiceId"
	fieldNameMap["service_instance_id"] = "ServiceInstanceId"
	fieldNameMap["service_instance_endpoint_id"] = "ServiceInstanceEndpointId"
	paramsTypeMap["locale_service_id"] = bindings.NewStringType()
	paramsTypeMap["tier1_id"] = bindings.NewStringType()
	paramsTypeMap["service_instance_id"] = bindings.NewStringType()
	paramsTypeMap["service_instance_endpoint_id"] = bindings.NewStringType()
	paramsTypeMap["tier1Id"] = bindings.NewStringType()
	paramsTypeMap["localeServiceId"] = bindings.NewStringType()
	paramsTypeMap["serviceInstanceId"] = bindings.NewStringType()
	paramsTypeMap["serviceInstanceEndpointId"] = bindings.NewStringType()
	pathParams["service_instance_endpoint_id"] = "serviceInstanceEndpointId"
	pathParams["tier1_id"] = "tier1Id"
	pathParams["service_instance_id"] = "serviceInstanceId"
	pathParams["locale_service_id"] = "localeServiceId"
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
		"/policy/api/v1/infra/tier-1s/{tier1Id}/locale-services/{localeServiceId}/byod-service-instances/{serviceInstanceId}/service-instance-endpoints/{serviceInstanceEndpointId}",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func serviceInstanceEndpointsListInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["tier1_id"] = bindings.NewStringType()
	fields["locale_service_id"] = bindings.NewStringType()
	fields["service_instance_id"] = bindings.NewStringType()
	fields["cursor"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["include_mark_for_delete_objects"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fields["included_fields"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["page_size"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fields["sort_ascending"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fields["sort_by"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["tier1_id"] = "Tier1Id"
	fieldNameMap["locale_service_id"] = "LocaleServiceId"
	fieldNameMap["service_instance_id"] = "ServiceInstanceId"
	fieldNameMap["cursor"] = "Cursor"
	fieldNameMap["include_mark_for_delete_objects"] = "IncludeMarkForDeleteObjects"
	fieldNameMap["included_fields"] = "IncludedFields"
	fieldNameMap["page_size"] = "PageSize"
	fieldNameMap["sort_ascending"] = "SortAscending"
	fieldNameMap["sort_by"] = "SortBy"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func serviceInstanceEndpointsListOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.ServiceInstanceEndpointListResultBindingType)
}

func serviceInstanceEndpointsListRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["tier1_id"] = bindings.NewStringType()
	fields["locale_service_id"] = bindings.NewStringType()
	fields["service_instance_id"] = bindings.NewStringType()
	fields["cursor"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["include_mark_for_delete_objects"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fields["included_fields"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["page_size"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fields["sort_ascending"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fields["sort_by"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["tier1_id"] = "Tier1Id"
	fieldNameMap["locale_service_id"] = "LocaleServiceId"
	fieldNameMap["service_instance_id"] = "ServiceInstanceId"
	fieldNameMap["cursor"] = "Cursor"
	fieldNameMap["include_mark_for_delete_objects"] = "IncludeMarkForDeleteObjects"
	fieldNameMap["included_fields"] = "IncludedFields"
	fieldNameMap["page_size"] = "PageSize"
	fieldNameMap["sort_ascending"] = "SortAscending"
	fieldNameMap["sort_by"] = "SortBy"
	paramsTypeMap["locale_service_id"] = bindings.NewStringType()
	paramsTypeMap["tier1_id"] = bindings.NewStringType()
	paramsTypeMap["included_fields"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["page_size"] = bindings.NewOptionalType(bindings.NewIntegerType())
	paramsTypeMap["service_instance_id"] = bindings.NewStringType()
	paramsTypeMap["include_mark_for_delete_objects"] = bindings.NewOptionalType(bindings.NewBooleanType())
	paramsTypeMap["cursor"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["sort_by"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["sort_ascending"] = bindings.NewOptionalType(bindings.NewBooleanType())
	paramsTypeMap["tier1Id"] = bindings.NewStringType()
	paramsTypeMap["localeServiceId"] = bindings.NewStringType()
	paramsTypeMap["serviceInstanceId"] = bindings.NewStringType()
	pathParams["tier1_id"] = "tier1Id"
	pathParams["service_instance_id"] = "serviceInstanceId"
	pathParams["locale_service_id"] = "localeServiceId"
	queryParams["cursor"] = "cursor"
	queryParams["sort_ascending"] = "sort_ascending"
	queryParams["included_fields"] = "included_fields"
	queryParams["sort_by"] = "sort_by"
	queryParams["include_mark_for_delete_objects"] = "include_mark_for_delete_objects"
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
		"/policy/api/v1/infra/tier-1s/{tier1Id}/locale-services/{localeServiceId}/byod-service-instances/{serviceInstanceId}/service-instance-endpoints",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func serviceInstanceEndpointsPatchInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["tier1_id"] = bindings.NewStringType()
	fields["locale_service_id"] = bindings.NewStringType()
	fields["service_instance_id"] = bindings.NewStringType()
	fields["service_instance_endpoint_id"] = bindings.NewStringType()
	fields["service_instance_endpoint"] = bindings.NewReferenceType(model.ServiceInstanceEndpointBindingType)
	fieldNameMap["tier1_id"] = "Tier1Id"
	fieldNameMap["locale_service_id"] = "LocaleServiceId"
	fieldNameMap["service_instance_id"] = "ServiceInstanceId"
	fieldNameMap["service_instance_endpoint_id"] = "ServiceInstanceEndpointId"
	fieldNameMap["service_instance_endpoint"] = "ServiceInstanceEndpoint"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func serviceInstanceEndpointsPatchOutputType() bindings.BindingType {
	return bindings.NewVoidType()
}

func serviceInstanceEndpointsPatchRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["tier1_id"] = bindings.NewStringType()
	fields["locale_service_id"] = bindings.NewStringType()
	fields["service_instance_id"] = bindings.NewStringType()
	fields["service_instance_endpoint_id"] = bindings.NewStringType()
	fields["service_instance_endpoint"] = bindings.NewReferenceType(model.ServiceInstanceEndpointBindingType)
	fieldNameMap["tier1_id"] = "Tier1Id"
	fieldNameMap["locale_service_id"] = "LocaleServiceId"
	fieldNameMap["service_instance_id"] = "ServiceInstanceId"
	fieldNameMap["service_instance_endpoint_id"] = "ServiceInstanceEndpointId"
	fieldNameMap["service_instance_endpoint"] = "ServiceInstanceEndpoint"
	paramsTypeMap["locale_service_id"] = bindings.NewStringType()
	paramsTypeMap["tier1_id"] = bindings.NewStringType()
	paramsTypeMap["service_instance_endpoint"] = bindings.NewReferenceType(model.ServiceInstanceEndpointBindingType)
	paramsTypeMap["service_instance_id"] = bindings.NewStringType()
	paramsTypeMap["service_instance_endpoint_id"] = bindings.NewStringType()
	paramsTypeMap["tier1Id"] = bindings.NewStringType()
	paramsTypeMap["localeServiceId"] = bindings.NewStringType()
	paramsTypeMap["serviceInstanceId"] = bindings.NewStringType()
	paramsTypeMap["serviceInstanceEndpointId"] = bindings.NewStringType()
	pathParams["service_instance_endpoint_id"] = "serviceInstanceEndpointId"
	pathParams["tier1_id"] = "tier1Id"
	pathParams["service_instance_id"] = "serviceInstanceId"
	pathParams["locale_service_id"] = "localeServiceId"
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
		"service_instance_endpoint",
		"PATCH",
		"/policy/api/v1/infra/tier-1s/{tier1Id}/locale-services/{localeServiceId}/byod-service-instances/{serviceInstanceId}/service-instance-endpoints/{serviceInstanceEndpointId}",
		"",
		resultHeaders,
		204,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func serviceInstanceEndpointsUpdateInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["tier1_id"] = bindings.NewStringType()
	fields["locale_service_id"] = bindings.NewStringType()
	fields["service_instance_id"] = bindings.NewStringType()
	fields["service_instance_endpoint_id"] = bindings.NewStringType()
	fields["service_instance_endpoint"] = bindings.NewReferenceType(model.ServiceInstanceEndpointBindingType)
	fieldNameMap["tier1_id"] = "Tier1Id"
	fieldNameMap["locale_service_id"] = "LocaleServiceId"
	fieldNameMap["service_instance_id"] = "ServiceInstanceId"
	fieldNameMap["service_instance_endpoint_id"] = "ServiceInstanceEndpointId"
	fieldNameMap["service_instance_endpoint"] = "ServiceInstanceEndpoint"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func serviceInstanceEndpointsUpdateOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.ServiceInstanceEndpointBindingType)
}

func serviceInstanceEndpointsUpdateRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["tier1_id"] = bindings.NewStringType()
	fields["locale_service_id"] = bindings.NewStringType()
	fields["service_instance_id"] = bindings.NewStringType()
	fields["service_instance_endpoint_id"] = bindings.NewStringType()
	fields["service_instance_endpoint"] = bindings.NewReferenceType(model.ServiceInstanceEndpointBindingType)
	fieldNameMap["tier1_id"] = "Tier1Id"
	fieldNameMap["locale_service_id"] = "LocaleServiceId"
	fieldNameMap["service_instance_id"] = "ServiceInstanceId"
	fieldNameMap["service_instance_endpoint_id"] = "ServiceInstanceEndpointId"
	fieldNameMap["service_instance_endpoint"] = "ServiceInstanceEndpoint"
	paramsTypeMap["locale_service_id"] = bindings.NewStringType()
	paramsTypeMap["tier1_id"] = bindings.NewStringType()
	paramsTypeMap["service_instance_endpoint"] = bindings.NewReferenceType(model.ServiceInstanceEndpointBindingType)
	paramsTypeMap["service_instance_id"] = bindings.NewStringType()
	paramsTypeMap["service_instance_endpoint_id"] = bindings.NewStringType()
	paramsTypeMap["tier1Id"] = bindings.NewStringType()
	paramsTypeMap["localeServiceId"] = bindings.NewStringType()
	paramsTypeMap["serviceInstanceId"] = bindings.NewStringType()
	paramsTypeMap["serviceInstanceEndpointId"] = bindings.NewStringType()
	pathParams["service_instance_endpoint_id"] = "serviceInstanceEndpointId"
	pathParams["tier1_id"] = "tier1Id"
	pathParams["service_instance_id"] = "serviceInstanceId"
	pathParams["locale_service_id"] = "localeServiceId"
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
		"service_instance_endpoint",
		"PUT",
		"/policy/api/v1/infra/tier-1s/{tier1Id}/locale-services/{localeServiceId}/byod-service-instances/{serviceInstanceId}/service-instance-endpoints/{serviceInstanceEndpointId}",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}
