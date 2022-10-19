// Copyright © 2019-2021 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Data type definitions file for service: ClientSslProfiles.
// Includes binding types of a structures and enumerations defined in the service.
// Shared by client-side stubs and server-side skeletons to ensure type
// compatibility.

package loadbalancer

import (
	"github.com/TaoZou1/vsphere-automation-sdk-go/runtime/bindings"
	"github.com/TaoZou1/vsphere-automation-sdk-go/runtime/data"
	"github.com/TaoZou1/vsphere-automation-sdk-go/runtime/protocol"
	"github.com/TaoZou1/vsphere-automation-sdk-go/services/nsxt-mp/nsx/model"
	"reflect"
)

func clientSslProfilesCreateInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["lb_client_ssl_profile"] = bindings.NewReferenceType(model.LbClientSslProfileBindingType)
	fieldNameMap["lb_client_ssl_profile"] = "LbClientSslProfile"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func clientSslProfilesCreateOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.LbClientSslProfileBindingType)
}

func clientSslProfilesCreateRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["lb_client_ssl_profile"] = bindings.NewReferenceType(model.LbClientSslProfileBindingType)
	fieldNameMap["lb_client_ssl_profile"] = "LbClientSslProfile"
	paramsTypeMap["lb_client_ssl_profile"] = bindings.NewReferenceType(model.LbClientSslProfileBindingType)
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
		"lb_client_ssl_profile",
		"POST",
		"/api/v1/loadbalancer/client-ssl-profiles",
		"",
		resultHeaders,
		201,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func clientSslProfilesDeleteInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["client_ssl_profile_id"] = bindings.NewStringType()
	fieldNameMap["client_ssl_profile_id"] = "ClientSslProfileId"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func clientSslProfilesDeleteOutputType() bindings.BindingType {
	return bindings.NewVoidType()
}

func clientSslProfilesDeleteRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["client_ssl_profile_id"] = bindings.NewStringType()
	fieldNameMap["client_ssl_profile_id"] = "ClientSslProfileId"
	paramsTypeMap["client_ssl_profile_id"] = bindings.NewStringType()
	paramsTypeMap["clientSslProfileId"] = bindings.NewStringType()
	pathParams["client_ssl_profile_id"] = "clientSslProfileId"
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
		"/api/v1/loadbalancer/client-ssl-profiles/{clientSslProfileId}",
		"",
		resultHeaders,
		204,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func clientSslProfilesGetInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["client_ssl_profile_id"] = bindings.NewStringType()
	fieldNameMap["client_ssl_profile_id"] = "ClientSslProfileId"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func clientSslProfilesGetOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.LbClientSslProfileBindingType)
}

func clientSslProfilesGetRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["client_ssl_profile_id"] = bindings.NewStringType()
	fieldNameMap["client_ssl_profile_id"] = "ClientSslProfileId"
	paramsTypeMap["client_ssl_profile_id"] = bindings.NewStringType()
	paramsTypeMap["clientSslProfileId"] = bindings.NewStringType()
	pathParams["client_ssl_profile_id"] = "clientSslProfileId"
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
		"/api/v1/loadbalancer/client-ssl-profiles/{clientSslProfileId}",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func clientSslProfilesListInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["cursor"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["included_fields"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["page_size"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fields["sort_ascending"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fields["sort_by"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["cursor"] = "Cursor"
	fieldNameMap["included_fields"] = "IncludedFields"
	fieldNameMap["page_size"] = "PageSize"
	fieldNameMap["sort_ascending"] = "SortAscending"
	fieldNameMap["sort_by"] = "SortBy"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func clientSslProfilesListOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.LbClientSslProfileListResultBindingType)
}

func clientSslProfilesListRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["cursor"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["included_fields"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["page_size"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fields["sort_ascending"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fields["sort_by"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["cursor"] = "Cursor"
	fieldNameMap["included_fields"] = "IncludedFields"
	fieldNameMap["page_size"] = "PageSize"
	fieldNameMap["sort_ascending"] = "SortAscending"
	fieldNameMap["sort_by"] = "SortBy"
	paramsTypeMap["included_fields"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["page_size"] = bindings.NewOptionalType(bindings.NewIntegerType())
	paramsTypeMap["cursor"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["sort_by"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["sort_ascending"] = bindings.NewOptionalType(bindings.NewBooleanType())
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
		"/api/v1/loadbalancer/client-ssl-profiles",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func clientSslProfilesUpdateInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["client_ssl_profile_id"] = bindings.NewStringType()
	fields["lb_client_ssl_profile"] = bindings.NewReferenceType(model.LbClientSslProfileBindingType)
	fieldNameMap["client_ssl_profile_id"] = "ClientSslProfileId"
	fieldNameMap["lb_client_ssl_profile"] = "LbClientSslProfile"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func clientSslProfilesUpdateOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.LbClientSslProfileBindingType)
}

func clientSslProfilesUpdateRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["client_ssl_profile_id"] = bindings.NewStringType()
	fields["lb_client_ssl_profile"] = bindings.NewReferenceType(model.LbClientSslProfileBindingType)
	fieldNameMap["client_ssl_profile_id"] = "ClientSslProfileId"
	fieldNameMap["lb_client_ssl_profile"] = "LbClientSslProfile"
	paramsTypeMap["client_ssl_profile_id"] = bindings.NewStringType()
	paramsTypeMap["lb_client_ssl_profile"] = bindings.NewReferenceType(model.LbClientSslProfileBindingType)
	paramsTypeMap["clientSslProfileId"] = bindings.NewStringType()
	pathParams["client_ssl_profile_id"] = "clientSslProfileId"
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
		"lb_client_ssl_profile",
		"PUT",
		"/api/v1/loadbalancer/client-ssl-profiles/{clientSslProfileId}",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}
