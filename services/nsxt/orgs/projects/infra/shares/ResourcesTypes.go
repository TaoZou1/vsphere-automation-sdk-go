// Copyright © 2019-2021 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Data type definitions file for service: Resources.
// Includes binding types of a structures and enumerations defined in the service.
// Shared by client-side stubs and server-side skeletons to ensure type
// compatibility.

package shares

import (
	"github.com/TaoZou1/vsphere-automation-sdk-go/runtime/bindings"
	"github.com/TaoZou1/vsphere-automation-sdk-go/runtime/data"
	"github.com/TaoZou1/vsphere-automation-sdk-go/runtime/protocol"
	"github.com/TaoZou1/vsphere-automation-sdk-go/services/nsxt/model"
	"reflect"
)

func resourcesDeleteInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["org_id"] = bindings.NewStringType()
	fields["project_id"] = bindings.NewStringType()
	fields["share_id"] = bindings.NewStringType()
	fields["shared_resource_id"] = bindings.NewStringType()
	fieldNameMap["org_id"] = "OrgId"
	fieldNameMap["project_id"] = "ProjectId"
	fieldNameMap["share_id"] = "ShareId"
	fieldNameMap["shared_resource_id"] = "SharedResourceId"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func resourcesDeleteOutputType() bindings.BindingType {
	return bindings.NewVoidType()
}

func resourcesDeleteRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["org_id"] = bindings.NewStringType()
	fields["project_id"] = bindings.NewStringType()
	fields["share_id"] = bindings.NewStringType()
	fields["shared_resource_id"] = bindings.NewStringType()
	fieldNameMap["org_id"] = "OrgId"
	fieldNameMap["project_id"] = "ProjectId"
	fieldNameMap["share_id"] = "ShareId"
	fieldNameMap["shared_resource_id"] = "SharedResourceId"
	paramsTypeMap["share_id"] = bindings.NewStringType()
	paramsTypeMap["shared_resource_id"] = bindings.NewStringType()
	paramsTypeMap["project_id"] = bindings.NewStringType()
	paramsTypeMap["org_id"] = bindings.NewStringType()
	paramsTypeMap["orgId"] = bindings.NewStringType()
	paramsTypeMap["projectId"] = bindings.NewStringType()
	paramsTypeMap["shareId"] = bindings.NewStringType()
	paramsTypeMap["sharedResourceId"] = bindings.NewStringType()
	pathParams["share_id"] = "shareId"
	pathParams["shared_resource_id"] = "sharedResourceId"
	pathParams["project_id"] = "projectId"
	pathParams["org_id"] = "orgId"
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
		"/policy/api/v1/orgs/{orgId}/projects/{projectId}/infra/shares/{shareId}/resources/{sharedResourceId}",
		"",
		resultHeaders,
		204,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func resourcesGetInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["share_id"] = bindings.NewStringType()
	fields["shared_resource_id"] = bindings.NewStringType()
	fields["org_id"] = bindings.NewStringType()
	fields["project_id"] = bindings.NewStringType()
	fieldNameMap["share_id"] = "ShareId"
	fieldNameMap["shared_resource_id"] = "SharedResourceId"
	fieldNameMap["org_id"] = "OrgId"
	fieldNameMap["project_id"] = "ProjectId"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func resourcesGetOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.SharedResourceBindingType)
}

func resourcesGetRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["share_id"] = bindings.NewStringType()
	fields["shared_resource_id"] = bindings.NewStringType()
	fields["org_id"] = bindings.NewStringType()
	fields["project_id"] = bindings.NewStringType()
	fieldNameMap["share_id"] = "ShareId"
	fieldNameMap["shared_resource_id"] = "SharedResourceId"
	fieldNameMap["org_id"] = "OrgId"
	fieldNameMap["project_id"] = "ProjectId"
	paramsTypeMap["share_id"] = bindings.NewStringType()
	paramsTypeMap["shared_resource_id"] = bindings.NewStringType()
	paramsTypeMap["project_id"] = bindings.NewStringType()
	paramsTypeMap["org_id"] = bindings.NewStringType()
	paramsTypeMap["orgId"] = bindings.NewStringType()
	paramsTypeMap["projectId"] = bindings.NewStringType()
	paramsTypeMap["shareId"] = bindings.NewStringType()
	paramsTypeMap["sharedResourceId"] = bindings.NewStringType()
	pathParams["share_id"] = "shareId"
	pathParams["shared_resource_id"] = "sharedResourceId"
	pathParams["project_id"] = "projectId"
	pathParams["org_id"] = "orgId"
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
		"/policy/api/v1/orgs/{orgId}/projects/{projectId}/infra/shares/{shareId}/resources/{sharedResourceId}",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func resourcesListInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["share_id"] = bindings.NewStringType()
	fields["org_id"] = bindings.NewStringType()
	fields["project_id"] = bindings.NewStringType()
	fields["resource_type"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["share_id"] = "ShareId"
	fieldNameMap["org_id"] = "OrgId"
	fieldNameMap["project_id"] = "ProjectId"
	fieldNameMap["resource_type"] = "ResourceType"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func resourcesListOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.SharedResourceListResultBindingType)
}

func resourcesListRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["share_id"] = bindings.NewStringType()
	fields["org_id"] = bindings.NewStringType()
	fields["project_id"] = bindings.NewStringType()
	fields["resource_type"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["share_id"] = "ShareId"
	fieldNameMap["org_id"] = "OrgId"
	fieldNameMap["project_id"] = "ProjectId"
	fieldNameMap["resource_type"] = "ResourceType"
	paramsTypeMap["share_id"] = bindings.NewStringType()
	paramsTypeMap["resource_type"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["project_id"] = bindings.NewStringType()
	paramsTypeMap["org_id"] = bindings.NewStringType()
	paramsTypeMap["orgId"] = bindings.NewStringType()
	paramsTypeMap["projectId"] = bindings.NewStringType()
	paramsTypeMap["shareId"] = bindings.NewStringType()
	pathParams["share_id"] = "shareId"
	pathParams["project_id"] = "projectId"
	pathParams["org_id"] = "orgId"
	queryParams["resource_type"] = "resource_type"
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
		"/policy/api/v1/orgs/{orgId}/projects/{projectId}/infra/shares/{shareId}/resources",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func resourcesPatchInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["share_id"] = bindings.NewStringType()
	fields["shared_resource_id"] = bindings.NewStringType()
	fields["org_id"] = bindings.NewStringType()
	fields["project_id"] = bindings.NewStringType()
	fields["shared_resource"] = bindings.NewReferenceType(model.SharedResourceBindingType)
	fieldNameMap["share_id"] = "ShareId"
	fieldNameMap["shared_resource_id"] = "SharedResourceId"
	fieldNameMap["org_id"] = "OrgId"
	fieldNameMap["project_id"] = "ProjectId"
	fieldNameMap["shared_resource"] = "SharedResource"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func resourcesPatchOutputType() bindings.BindingType {
	return bindings.NewVoidType()
}

func resourcesPatchRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["share_id"] = bindings.NewStringType()
	fields["shared_resource_id"] = bindings.NewStringType()
	fields["org_id"] = bindings.NewStringType()
	fields["project_id"] = bindings.NewStringType()
	fields["shared_resource"] = bindings.NewReferenceType(model.SharedResourceBindingType)
	fieldNameMap["share_id"] = "ShareId"
	fieldNameMap["shared_resource_id"] = "SharedResourceId"
	fieldNameMap["org_id"] = "OrgId"
	fieldNameMap["project_id"] = "ProjectId"
	fieldNameMap["shared_resource"] = "SharedResource"
	paramsTypeMap["share_id"] = bindings.NewStringType()
	paramsTypeMap["shared_resource_id"] = bindings.NewStringType()
	paramsTypeMap["project_id"] = bindings.NewStringType()
	paramsTypeMap["org_id"] = bindings.NewStringType()
	paramsTypeMap["shared_resource"] = bindings.NewReferenceType(model.SharedResourceBindingType)
	paramsTypeMap["orgId"] = bindings.NewStringType()
	paramsTypeMap["projectId"] = bindings.NewStringType()
	paramsTypeMap["shareId"] = bindings.NewStringType()
	paramsTypeMap["sharedResourceId"] = bindings.NewStringType()
	pathParams["share_id"] = "shareId"
	pathParams["shared_resource_id"] = "sharedResourceId"
	pathParams["project_id"] = "projectId"
	pathParams["org_id"] = "orgId"
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
		"shared_resource",
		"PATCH",
		"/policy/api/v1/orgs/{orgId}/projects/{projectId}/infra/shares/{shareId}/resources/{sharedResourceId}",
		"",
		resultHeaders,
		204,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func resourcesUpdateInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["share_id"] = bindings.NewStringType()
	fields["shared_resource_id"] = bindings.NewStringType()
	fields["org_id"] = bindings.NewStringType()
	fields["project_id"] = bindings.NewStringType()
	fields["shared_resource"] = bindings.NewReferenceType(model.SharedResourceBindingType)
	fieldNameMap["share_id"] = "ShareId"
	fieldNameMap["shared_resource_id"] = "SharedResourceId"
	fieldNameMap["org_id"] = "OrgId"
	fieldNameMap["project_id"] = "ProjectId"
	fieldNameMap["shared_resource"] = "SharedResource"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func resourcesUpdateOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.SharedResourceBindingType)
}

func resourcesUpdateRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["share_id"] = bindings.NewStringType()
	fields["shared_resource_id"] = bindings.NewStringType()
	fields["org_id"] = bindings.NewStringType()
	fields["project_id"] = bindings.NewStringType()
	fields["shared_resource"] = bindings.NewReferenceType(model.SharedResourceBindingType)
	fieldNameMap["share_id"] = "ShareId"
	fieldNameMap["shared_resource_id"] = "SharedResourceId"
	fieldNameMap["org_id"] = "OrgId"
	fieldNameMap["project_id"] = "ProjectId"
	fieldNameMap["shared_resource"] = "SharedResource"
	paramsTypeMap["share_id"] = bindings.NewStringType()
	paramsTypeMap["shared_resource_id"] = bindings.NewStringType()
	paramsTypeMap["project_id"] = bindings.NewStringType()
	paramsTypeMap["org_id"] = bindings.NewStringType()
	paramsTypeMap["shared_resource"] = bindings.NewReferenceType(model.SharedResourceBindingType)
	paramsTypeMap["orgId"] = bindings.NewStringType()
	paramsTypeMap["projectId"] = bindings.NewStringType()
	paramsTypeMap["shareId"] = bindings.NewStringType()
	paramsTypeMap["sharedResourceId"] = bindings.NewStringType()
	pathParams["share_id"] = "shareId"
	pathParams["shared_resource_id"] = "sharedResourceId"
	pathParams["project_id"] = "projectId"
	pathParams["org_id"] = "orgId"
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
		"shared_resource",
		"PUT",
		"/policy/api/v1/orgs/{orgId}/projects/{projectId}/infra/shares/{shareId}/resources/{sharedResourceId}",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}
