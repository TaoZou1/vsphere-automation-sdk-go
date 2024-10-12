// Copyright © 2019-2021 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Data type definitions file for service: IdentityFirewallStores.
// Includes binding types of a structures and enumerations defined in the service.
// Shared by client-side stubs and server-side skeletons to ensure type
// compatibility.

package infra

import (
	"github.com/vmware/vsphere-automation-sdk-go/runtime/bindings"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/data"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/protocol"
	"github.com/vmware/vsphere-automation-sdk-go/services/nsxt/model"
	"reflect"
)

func identityFirewallStoresDeleteInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["identity_firewall_store_id"] = bindings.NewStringType()
	fields["enforcement_point_path"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["identity_firewall_store_id"] = "IdentityFirewallStoreId"
	fieldNameMap["enforcement_point_path"] = "EnforcementPointPath"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func identityFirewallStoresDeleteOutputType() bindings.BindingType {
	return bindings.NewVoidType()
}

func identityFirewallStoresDeleteRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["identity_firewall_store_id"] = bindings.NewStringType()
	fields["enforcement_point_path"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["identity_firewall_store_id"] = "IdentityFirewallStoreId"
	fieldNameMap["enforcement_point_path"] = "EnforcementPointPath"
	paramsTypeMap["enforcement_point_path"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["identity_firewall_store_id"] = bindings.NewStringType()
	paramsTypeMap["identityFirewallStoreId"] = bindings.NewStringType()
	pathParams["identity_firewall_store_id"] = "identityFirewallStoreId"
	queryParams["enforcement_point_path"] = "enforcement_point_path"
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
		"/policy/api/v1/infra/identity-firewall-stores/{identityFirewallStoreId}",
		"",
		resultHeaders,
		204,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func identityFirewallStoresGetInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["identity_firewall_store_id"] = bindings.NewStringType()
	fields["enforcement_point_path"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["identity_firewall_store_id"] = "IdentityFirewallStoreId"
	fieldNameMap["enforcement_point_path"] = "EnforcementPointPath"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func identityFirewallStoresGetOutputType() bindings.BindingType {
	return bindings.NewDynamicStructType([]bindings.ReferenceType{bindings.NewReferenceType(model.IdentityFirewallStoreBindingType)}, bindings.REST)
}

func identityFirewallStoresGetRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["identity_firewall_store_id"] = bindings.NewStringType()
	fields["enforcement_point_path"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["identity_firewall_store_id"] = "IdentityFirewallStoreId"
	fieldNameMap["enforcement_point_path"] = "EnforcementPointPath"
	paramsTypeMap["enforcement_point_path"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["identity_firewall_store_id"] = bindings.NewStringType()
	paramsTypeMap["identityFirewallStoreId"] = bindings.NewStringType()
	pathParams["identity_firewall_store_id"] = "identityFirewallStoreId"
	queryParams["enforcement_point_path"] = "enforcement_point_path"
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
		"/policy/api/v1/infra/identity-firewall-stores/{identityFirewallStoreId}",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func identityFirewallStoresListInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["cursor"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["enforcement_point_path"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["included_fields"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["page_size"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fields["sort_ascending"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fields["sort_by"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["cursor"] = "Cursor"
	fieldNameMap["enforcement_point_path"] = "EnforcementPointPath"
	fieldNameMap["included_fields"] = "IncludedFields"
	fieldNameMap["page_size"] = "PageSize"
	fieldNameMap["sort_ascending"] = "SortAscending"
	fieldNameMap["sort_by"] = "SortBy"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func identityFirewallStoresListOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.IdentityFirewallStoreListResultsBindingType)
}

func identityFirewallStoresListRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["cursor"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["enforcement_point_path"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["included_fields"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["page_size"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fields["sort_ascending"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fields["sort_by"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["cursor"] = "Cursor"
	fieldNameMap["enforcement_point_path"] = "EnforcementPointPath"
	fieldNameMap["included_fields"] = "IncludedFields"
	fieldNameMap["page_size"] = "PageSize"
	fieldNameMap["sort_ascending"] = "SortAscending"
	fieldNameMap["sort_by"] = "SortBy"
	paramsTypeMap["enforcement_point_path"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["included_fields"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["page_size"] = bindings.NewOptionalType(bindings.NewIntegerType())
	paramsTypeMap["cursor"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["sort_by"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["sort_ascending"] = bindings.NewOptionalType(bindings.NewBooleanType())
	queryParams["cursor"] = "cursor"
	queryParams["sort_ascending"] = "sort_ascending"
	queryParams["included_fields"] = "included_fields"
	queryParams["enforcement_point_path"] = "enforcement_point_path"
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
		"/policy/api/v1/infra/identity-firewall-stores",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func identityFirewallStoresPatchInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["identity_firewall_store_id"] = bindings.NewStringType()
	fields["identity_firewall_store"] = bindings.NewDynamicStructType([]bindings.ReferenceType{bindings.NewReferenceType(model.IdentityFirewallStoreBindingType)}, bindings.REST)
	fields["enforcement_point_path"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["identity_firewall_store_id"] = "IdentityFirewallStoreId"
	fieldNameMap["identity_firewall_store"] = "IdentityFirewallStore"
	fieldNameMap["enforcement_point_path"] = "EnforcementPointPath"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func identityFirewallStoresPatchOutputType() bindings.BindingType {
	return bindings.NewVoidType()
}

func identityFirewallStoresPatchRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["identity_firewall_store_id"] = bindings.NewStringType()
	fields["identity_firewall_store"] = bindings.NewDynamicStructType([]bindings.ReferenceType{bindings.NewReferenceType(model.IdentityFirewallStoreBindingType)}, bindings.REST)
	fields["enforcement_point_path"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["identity_firewall_store_id"] = "IdentityFirewallStoreId"
	fieldNameMap["identity_firewall_store"] = "IdentityFirewallStore"
	fieldNameMap["enforcement_point_path"] = "EnforcementPointPath"
	paramsTypeMap["enforcement_point_path"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["identity_firewall_store"] = bindings.NewDynamicStructType([]bindings.ReferenceType{bindings.NewReferenceType(model.IdentityFirewallStoreBindingType)}, bindings.REST)
	paramsTypeMap["identity_firewall_store_id"] = bindings.NewStringType()
	paramsTypeMap["identityFirewallStoreId"] = bindings.NewStringType()
	pathParams["identity_firewall_store_id"] = "identityFirewallStoreId"
	queryParams["enforcement_point_path"] = "enforcement_point_path"
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
		"identity_firewall_store",
		"PATCH",
		"/policy/api/v1/infra/identity-firewall-stores/{identityFirewallStoreId}",
		"",
		resultHeaders,
		204,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func identityFirewallStoresUpdateInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["identity_firewall_store_id"] = bindings.NewStringType()
	fields["identity_firewall_store"] = bindings.NewDynamicStructType([]bindings.ReferenceType{bindings.NewReferenceType(model.IdentityFirewallStoreBindingType)}, bindings.REST)
	fields["enforcement_point_path"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["identity_firewall_store_id"] = "IdentityFirewallStoreId"
	fieldNameMap["identity_firewall_store"] = "IdentityFirewallStore"
	fieldNameMap["enforcement_point_path"] = "EnforcementPointPath"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func identityFirewallStoresUpdateOutputType() bindings.BindingType {
	return bindings.NewDynamicStructType([]bindings.ReferenceType{bindings.NewReferenceType(model.IdentityFirewallStoreBindingType)}, bindings.REST)
}

func identityFirewallStoresUpdateRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["identity_firewall_store_id"] = bindings.NewStringType()
	fields["identity_firewall_store"] = bindings.NewDynamicStructType([]bindings.ReferenceType{bindings.NewReferenceType(model.IdentityFirewallStoreBindingType)}, bindings.REST)
	fields["enforcement_point_path"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["identity_firewall_store_id"] = "IdentityFirewallStoreId"
	fieldNameMap["identity_firewall_store"] = "IdentityFirewallStore"
	fieldNameMap["enforcement_point_path"] = "EnforcementPointPath"
	paramsTypeMap["enforcement_point_path"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["identity_firewall_store"] = bindings.NewDynamicStructType([]bindings.ReferenceType{bindings.NewReferenceType(model.IdentityFirewallStoreBindingType)}, bindings.REST)
	paramsTypeMap["identity_firewall_store_id"] = bindings.NewStringType()
	paramsTypeMap["identityFirewallStoreId"] = bindings.NewStringType()
	pathParams["identity_firewall_store_id"] = "identityFirewallStoreId"
	queryParams["enforcement_point_path"] = "enforcement_point_path"
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
		"identity_firewall_store",
		"PUT",
		"/policy/api/v1/infra/identity-firewall-stores/{identityFirewallStoreId}",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}
