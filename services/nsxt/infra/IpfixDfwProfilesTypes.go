/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Code generated. DO NOT EDIT.

/*
 * Data type definitions file for service: IpfixDfwProfiles.
 * Includes binding types of a structures and enumerations defined in the service.
 * Shared by client-side stubs and server-side skeletons to ensure type
 * compatibility.
 */

package infra

import (
	"reflect"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/services/nsxt/model"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/bindings"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/data"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/protocol"
)





func ipfixDfwProfilesDeleteInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["ipfix_dfw_profile_id"] = bindings.NewStringType()
	fieldNameMap["ipfix_dfw_profile_id"] = "IpfixDfwProfileId"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func ipfixDfwProfilesDeleteOutputType() bindings.BindingType {
	return bindings.NewVoidType()
}

func ipfixDfwProfilesDeleteRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	fields["ipfix_dfw_profile_id"] = bindings.NewStringType()
	fieldNameMap["ipfix_dfw_profile_id"] = "IpfixDfwProfileId"
	paramsTypeMap["ipfix_dfw_profile_id"] = bindings.NewStringType()
	paramsTypeMap["ipfixDfwProfileId"] = bindings.NewStringType()
	pathParams["ipfix_dfw_profile_id"] = "ipfixDfwProfileId"
	resultHeaders := map[string]string{}
	errorHeaders := map[string]string{}
	return protocol.NewOperationRestMetadata(
		fields,
		fieldNameMap,
		paramsTypeMap,
		pathParams,
		queryParams,
		headerParams,
		"",
		"",
		"DELETE",
		"/policy/api/v1/infra/ipfix-dfw-profiles/{ipfixDfwProfileId}",
		resultHeaders,
		204,
		errorHeaders,
		map[string]int{"InvalidRequest": 400,"Unauthorized": 403,"ServiceUnavailable": 503,"InternalServerError": 500,"NotFound": 404})
}

func ipfixDfwProfilesGetInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["ipfix_dfw_profile_id"] = bindings.NewStringType()
	fieldNameMap["ipfix_dfw_profile_id"] = "IpfixDfwProfileId"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func ipfixDfwProfilesGetOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.IPFIXDFWProfileBindingType)
}

func ipfixDfwProfilesGetRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	fields["ipfix_dfw_profile_id"] = bindings.NewStringType()
	fieldNameMap["ipfix_dfw_profile_id"] = "IpfixDfwProfileId"
	paramsTypeMap["ipfix_dfw_profile_id"] = bindings.NewStringType()
	paramsTypeMap["ipfixDfwProfileId"] = bindings.NewStringType()
	pathParams["ipfix_dfw_profile_id"] = "ipfixDfwProfileId"
	resultHeaders := map[string]string{}
	errorHeaders := map[string]string{}
	return protocol.NewOperationRestMetadata(
		fields,
		fieldNameMap,
		paramsTypeMap,
		pathParams,
		queryParams,
		headerParams,
		"",
		"",
		"GET",
		"/policy/api/v1/infra/ipfix-dfw-profiles/{ipfixDfwProfileId}",
		resultHeaders,
		200,
		errorHeaders,
		map[string]int{"InvalidRequest": 400,"Unauthorized": 403,"ServiceUnavailable": 503,"InternalServerError": 500,"NotFound": 404})
}

func ipfixDfwProfilesListInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["cursor"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["include_mark_for_delete_objects"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fields["included_fields"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["page_size"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fields["sort_ascending"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fields["sort_by"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["cursor"] = "Cursor"
	fieldNameMap["include_mark_for_delete_objects"] = "IncludeMarkForDeleteObjects"
	fieldNameMap["included_fields"] = "IncludedFields"
	fieldNameMap["page_size"] = "PageSize"
	fieldNameMap["sort_ascending"] = "SortAscending"
	fieldNameMap["sort_by"] = "SortBy"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func ipfixDfwProfilesListOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.IPFIXDFWProfileListResultBindingType)
}

func ipfixDfwProfilesListRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	fields["cursor"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["include_mark_for_delete_objects"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fields["included_fields"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["page_size"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fields["sort_ascending"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fields["sort_by"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["cursor"] = "Cursor"
	fieldNameMap["include_mark_for_delete_objects"] = "IncludeMarkForDeleteObjects"
	fieldNameMap["included_fields"] = "IncludedFields"
	fieldNameMap["page_size"] = "PageSize"
	fieldNameMap["sort_ascending"] = "SortAscending"
	fieldNameMap["sort_by"] = "SortBy"
	paramsTypeMap["included_fields"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["page_size"] = bindings.NewOptionalType(bindings.NewIntegerType())
	paramsTypeMap["include_mark_for_delete_objects"] = bindings.NewOptionalType(bindings.NewBooleanType())
	paramsTypeMap["cursor"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["sort_by"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["sort_ascending"] = bindings.NewOptionalType(bindings.NewBooleanType())
	queryParams["cursor"] = "cursor"
	queryParams["sort_ascending"] = "sort_ascending"
	queryParams["included_fields"] = "included_fields"
	queryParams["sort_by"] = "sort_by"
	queryParams["include_mark_for_delete_objects"] = "include_mark_for_delete_objects"
	queryParams["page_size"] = "page_size"
	resultHeaders := map[string]string{}
	errorHeaders := map[string]string{}
	return protocol.NewOperationRestMetadata(
		fields,
		fieldNameMap,
		paramsTypeMap,
		pathParams,
		queryParams,
		headerParams,
		"",
		"",
		"GET",
		"/policy/api/v1/infra/ipfix-dfw-profiles",
		resultHeaders,
		200,
		errorHeaders,
		map[string]int{"InvalidRequest": 400,"Unauthorized": 403,"ServiceUnavailable": 503,"InternalServerError": 500,"NotFound": 404})
}

func ipfixDfwProfilesPatchInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["ipfix_dfw_profile_id"] = bindings.NewStringType()
	fields["i_PFIXDFW_profile"] = bindings.NewReferenceType(model.IPFIXDFWProfileBindingType)
	fieldNameMap["ipfix_dfw_profile_id"] = "IpfixDfwProfileId"
	fieldNameMap["i_PFIXDFW_profile"] = "IPFIXDFWProfile"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func ipfixDfwProfilesPatchOutputType() bindings.BindingType {
	return bindings.NewVoidType()
}

func ipfixDfwProfilesPatchRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	fields["ipfix_dfw_profile_id"] = bindings.NewStringType()
	fields["i_PFIXDFW_profile"] = bindings.NewReferenceType(model.IPFIXDFWProfileBindingType)
	fieldNameMap["ipfix_dfw_profile_id"] = "IpfixDfwProfileId"
	fieldNameMap["i_PFIXDFW_profile"] = "IPFIXDFWProfile"
	paramsTypeMap["ipfix_dfw_profile_id"] = bindings.NewStringType()
	paramsTypeMap["i_PFIXDFW_profile"] = bindings.NewReferenceType(model.IPFIXDFWProfileBindingType)
	paramsTypeMap["ipfixDfwProfileId"] = bindings.NewStringType()
	pathParams["ipfix_dfw_profile_id"] = "ipfixDfwProfileId"
	resultHeaders := map[string]string{}
	errorHeaders := map[string]string{}
	return protocol.NewOperationRestMetadata(
		fields,
		fieldNameMap,
		paramsTypeMap,
		pathParams,
		queryParams,
		headerParams,
		"",
		"i_PFIXDFW_profile",
		"PATCH",
		"/policy/api/v1/infra/ipfix-dfw-profiles/{ipfixDfwProfileId}",
		resultHeaders,
		204,
		errorHeaders,
		map[string]int{"InvalidRequest": 400,"Unauthorized": 403,"ServiceUnavailable": 503,"InternalServerError": 500,"NotFound": 404})
}

func ipfixDfwProfilesUpdateInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["ipfix_dfw_profile_id"] = bindings.NewStringType()
	fields["i_PFIXDFW_profile"] = bindings.NewReferenceType(model.IPFIXDFWProfileBindingType)
	fieldNameMap["ipfix_dfw_profile_id"] = "IpfixDfwProfileId"
	fieldNameMap["i_PFIXDFW_profile"] = "IPFIXDFWProfile"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func ipfixDfwProfilesUpdateOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.IPFIXDFWProfileBindingType)
}

func ipfixDfwProfilesUpdateRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	fields["ipfix_dfw_profile_id"] = bindings.NewStringType()
	fields["i_PFIXDFW_profile"] = bindings.NewReferenceType(model.IPFIXDFWProfileBindingType)
	fieldNameMap["ipfix_dfw_profile_id"] = "IpfixDfwProfileId"
	fieldNameMap["i_PFIXDFW_profile"] = "IPFIXDFWProfile"
	paramsTypeMap["ipfix_dfw_profile_id"] = bindings.NewStringType()
	paramsTypeMap["i_PFIXDFW_profile"] = bindings.NewReferenceType(model.IPFIXDFWProfileBindingType)
	paramsTypeMap["ipfixDfwProfileId"] = bindings.NewStringType()
	pathParams["ipfix_dfw_profile_id"] = "ipfixDfwProfileId"
	resultHeaders := map[string]string{}
	errorHeaders := map[string]string{}
	return protocol.NewOperationRestMetadata(
		fields,
		fieldNameMap,
		paramsTypeMap,
		pathParams,
		queryParams,
		headerParams,
		"",
		"i_PFIXDFW_profile",
		"PUT",
		"/policy/api/v1/infra/ipfix-dfw-profiles/{ipfixDfwProfileId}",
		resultHeaders,
		200,
		errorHeaders,
		map[string]int{"InvalidRequest": 400,"Unauthorized": 403,"ServiceUnavailable": 503,"InternalServerError": 500,"NotFound": 404})
}

