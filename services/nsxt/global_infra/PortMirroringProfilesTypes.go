/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Code generated. DO NOT EDIT.

/*
 * Data type definitions file for service: PortMirroringProfiles.
 * Includes binding types of a structures and enumerations defined in the service.
 * Shared by client-side stubs and server-side skeletons to ensure type
 * compatibility.
 */

package global_infra

import (
	"reflect"
	"github.com/vmware/vsphere-automation-sdk-go/services/nsxt/model"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/bindings"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/data"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/protocol"
)





func portMirroringProfilesDeleteInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["port_mirroring_profile_id"] = bindings.NewStringType()
	fields["override"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fieldNameMap["port_mirroring_profile_id"] = "PortMirroringProfileId"
	fieldNameMap["override"] = "Override"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func portMirroringProfilesDeleteOutputType() bindings.BindingType {
	return bindings.NewVoidType()
}

func portMirroringProfilesDeleteRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	fields["port_mirroring_profile_id"] = bindings.NewStringType()
	fields["override"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fieldNameMap["port_mirroring_profile_id"] = "PortMirroringProfileId"
	fieldNameMap["override"] = "Override"
	paramsTypeMap["port_mirroring_profile_id"] = bindings.NewStringType()
	paramsTypeMap["override"] = bindings.NewOptionalType(bindings.NewBooleanType())
	paramsTypeMap["portMirroringProfileId"] = bindings.NewStringType()
	pathParams["port_mirroring_profile_id"] = "portMirroringProfileId"
	queryParams["override"] = "override"
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
		"/policy/api/v1/global-infra/port-mirroring-profiles/{portMirroringProfileId}",
		resultHeaders,
		204,
		errorHeaders,
		map[string]int{"InvalidRequest": 400,"Unauthorized": 403,"ServiceUnavailable": 503,"InternalServerError": 500,"NotFound": 404})
}

func portMirroringProfilesGetInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["port_mirroring_profile_id"] = bindings.NewStringType()
	fieldNameMap["port_mirroring_profile_id"] = "PortMirroringProfileId"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func portMirroringProfilesGetOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.PortMirroringProfileBindingType)
}

func portMirroringProfilesGetRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	fields["port_mirroring_profile_id"] = bindings.NewStringType()
	fieldNameMap["port_mirroring_profile_id"] = "PortMirroringProfileId"
	paramsTypeMap["port_mirroring_profile_id"] = bindings.NewStringType()
	paramsTypeMap["portMirroringProfileId"] = bindings.NewStringType()
	pathParams["port_mirroring_profile_id"] = "portMirroringProfileId"
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
		"/policy/api/v1/global-infra/port-mirroring-profiles/{portMirroringProfileId}",
		resultHeaders,
		200,
		errorHeaders,
		map[string]int{"InvalidRequest": 400,"Unauthorized": 403,"ServiceUnavailable": 503,"InternalServerError": 500,"NotFound": 404})
}

func portMirroringProfilesListInputType() bindings.StructType {
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

func portMirroringProfilesListOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.PortMirroringProfileListResultBindingType)
}

func portMirroringProfilesListRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
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
		"/policy/api/v1/global-infra/port-mirroring-profiles",
		resultHeaders,
		200,
		errorHeaders,
		map[string]int{"InvalidRequest": 400,"Unauthorized": 403,"ServiceUnavailable": 503,"InternalServerError": 500,"NotFound": 404})
}

func portMirroringProfilesPatchInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["port_mirroring_profile_id"] = bindings.NewStringType()
	fields["port_mirroring_profile"] = bindings.NewReferenceType(model.PortMirroringProfileBindingType)
	fields["override"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fieldNameMap["port_mirroring_profile_id"] = "PortMirroringProfileId"
	fieldNameMap["port_mirroring_profile"] = "PortMirroringProfile"
	fieldNameMap["override"] = "Override"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func portMirroringProfilesPatchOutputType() bindings.BindingType {
	return bindings.NewVoidType()
}

func portMirroringProfilesPatchRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	fields["port_mirroring_profile_id"] = bindings.NewStringType()
	fields["port_mirroring_profile"] = bindings.NewReferenceType(model.PortMirroringProfileBindingType)
	fields["override"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fieldNameMap["port_mirroring_profile_id"] = "PortMirroringProfileId"
	fieldNameMap["port_mirroring_profile"] = "PortMirroringProfile"
	fieldNameMap["override"] = "Override"
	paramsTypeMap["port_mirroring_profile_id"] = bindings.NewStringType()
	paramsTypeMap["override"] = bindings.NewOptionalType(bindings.NewBooleanType())
	paramsTypeMap["port_mirroring_profile"] = bindings.NewReferenceType(model.PortMirroringProfileBindingType)
	paramsTypeMap["portMirroringProfileId"] = bindings.NewStringType()
	pathParams["port_mirroring_profile_id"] = "portMirroringProfileId"
	queryParams["override"] = "override"
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
		"port_mirroring_profile",
		"PATCH",
		"/policy/api/v1/global-infra/port-mirroring-profiles/{portMirroringProfileId}",
		resultHeaders,
		204,
		errorHeaders,
		map[string]int{"InvalidRequest": 400,"Unauthorized": 403,"ServiceUnavailable": 503,"InternalServerError": 500,"NotFound": 404})
}

func portMirroringProfilesUpdateInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["port_mirroring_profile_id"] = bindings.NewStringType()
	fields["port_mirroring_profile"] = bindings.NewReferenceType(model.PortMirroringProfileBindingType)
	fields["override"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fieldNameMap["port_mirroring_profile_id"] = "PortMirroringProfileId"
	fieldNameMap["port_mirroring_profile"] = "PortMirroringProfile"
	fieldNameMap["override"] = "Override"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func portMirroringProfilesUpdateOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.PortMirroringProfileBindingType)
}

func portMirroringProfilesUpdateRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	fields["port_mirroring_profile_id"] = bindings.NewStringType()
	fields["port_mirroring_profile"] = bindings.NewReferenceType(model.PortMirroringProfileBindingType)
	fields["override"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fieldNameMap["port_mirroring_profile_id"] = "PortMirroringProfileId"
	fieldNameMap["port_mirroring_profile"] = "PortMirroringProfile"
	fieldNameMap["override"] = "Override"
	paramsTypeMap["port_mirroring_profile_id"] = bindings.NewStringType()
	paramsTypeMap["override"] = bindings.NewOptionalType(bindings.NewBooleanType())
	paramsTypeMap["port_mirroring_profile"] = bindings.NewReferenceType(model.PortMirroringProfileBindingType)
	paramsTypeMap["portMirroringProfileId"] = bindings.NewStringType()
	pathParams["port_mirroring_profile_id"] = "portMirroringProfileId"
	queryParams["override"] = "override"
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
		"port_mirroring_profile",
		"PUT",
		"/policy/api/v1/global-infra/port-mirroring-profiles/{portMirroringProfileId}",
		resultHeaders,
		200,
		errorHeaders,
		map[string]int{"InvalidRequest": 400,"Unauthorized": 403,"ServiceUnavailable": 503,"InternalServerError": 500,"NotFound": 404})
}


