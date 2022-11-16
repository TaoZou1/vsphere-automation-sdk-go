// Copyright © 2019-2021 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Data type definitions file for service: RoleBindings.
// Includes binding types of a structures and enumerations defined in the service.
// Shared by client-side stubs and server-side skeletons to ensure type
// compatibility.

package aaa

import (
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/runtime/bindings"
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/runtime/data"
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/runtime/protocol"
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/services/nsxt/model"
	"reflect"
)

// Possible value for ``identitySourceType`` of method RoleBindings#delete.
const RoleBindings_DELETE_IDENTITY_SOURCE_TYPE_VIDM = "VIDM"

// Possible value for ``identitySourceType`` of method RoleBindings#delete.
const RoleBindings_DELETE_IDENTITY_SOURCE_TYPE_LDAP = "LDAP"

// Possible value for ``identitySourceType`` of method RoleBindings#delete.
const RoleBindings_DELETE_IDENTITY_SOURCE_TYPE_OIDC = "OIDC"

// Possible value for ``type`` of method RoleBindings#delete.
const RoleBindings_DELETE_TYPE_REMOTE_USER = "remote_user"

// Possible value for ``type`` of method RoleBindings#delete.
const RoleBindings_DELETE_TYPE_REMOTE_GROUP = "remote_group"

// Possible value for ``type`` of method RoleBindings#delete.
const RoleBindings_DELETE_TYPE_LOCAL_USER = "local_user"

// Possible value for ``type`` of method RoleBindings#delete.
const RoleBindings_DELETE_TYPE_PRINCIPAL_IDENTITY = "principal_identity"

// Possible value for ``identitySourceType`` of method RoleBindings#get.
const RoleBindings_GET_IDENTITY_SOURCE_TYPE_VIDM = "VIDM"

// Possible value for ``identitySourceType`` of method RoleBindings#get.
const RoleBindings_GET_IDENTITY_SOURCE_TYPE_LDAP = "LDAP"

// Possible value for ``identitySourceType`` of method RoleBindings#get.
const RoleBindings_GET_IDENTITY_SOURCE_TYPE_OIDC = "OIDC"

// Possible value for ``type`` of method RoleBindings#get.
const RoleBindings_GET_TYPE_REMOTE_USER = "remote_user"

// Possible value for ``type`` of method RoleBindings#get.
const RoleBindings_GET_TYPE_REMOTE_GROUP = "remote_group"

// Possible value for ``type`` of method RoleBindings#get.
const RoleBindings_GET_TYPE_LOCAL_USER = "local_user"

// Possible value for ``type`` of method RoleBindings#get.
const RoleBindings_GET_TYPE_PRINCIPAL_IDENTITY = "principal_identity"

// Possible value for ``identitySourceType`` of method RoleBindings#list.
const RoleBindings_LIST_IDENTITY_SOURCE_TYPE_VIDM = "VIDM"

// Possible value for ``identitySourceType`` of method RoleBindings#list.
const RoleBindings_LIST_IDENTITY_SOURCE_TYPE_LDAP = "LDAP"

// Possible value for ``identitySourceType`` of method RoleBindings#list.
const RoleBindings_LIST_IDENTITY_SOURCE_TYPE_OIDC = "OIDC"

// Possible value for ``type`` of method RoleBindings#list.
const RoleBindings_LIST_TYPE_REMOTE_USER = "remote_user"

// Possible value for ``type`` of method RoleBindings#list.
const RoleBindings_LIST_TYPE_REMOTE_GROUP = "remote_group"

// Possible value for ``type`` of method RoleBindings#list.
const RoleBindings_LIST_TYPE_LOCAL_USER = "local_user"

// Possible value for ``type`` of method RoleBindings#list.
const RoleBindings_LIST_TYPE_PRINCIPAL_IDENTITY = "principal_identity"

func roleBindingsDeleteInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["org_id"] = bindings.NewStringType()
	fields["project_id"] = bindings.NewStringType()
	fields["vpc_id"] = bindings.NewStringType()
	fields["binding_id"] = bindings.NewStringType()
	fields["cursor"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["identity_source_id"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["identity_source_type"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["included_fields"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["name"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["page_size"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fields["path"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["role"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["root_path"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["sort_ascending"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fields["sort_by"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["type"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["org_id"] = "OrgId"
	fieldNameMap["project_id"] = "ProjectId"
	fieldNameMap["vpc_id"] = "VpcId"
	fieldNameMap["binding_id"] = "BindingId"
	fieldNameMap["cursor"] = "Cursor"
	fieldNameMap["identity_source_id"] = "IdentitySourceId"
	fieldNameMap["identity_source_type"] = "IdentitySourceType"
	fieldNameMap["included_fields"] = "IncludedFields"
	fieldNameMap["name"] = "Name"
	fieldNameMap["page_size"] = "PageSize"
	fieldNameMap["path"] = "Path"
	fieldNameMap["role"] = "Role"
	fieldNameMap["root_path"] = "RootPath"
	fieldNameMap["sort_ascending"] = "SortAscending"
	fieldNameMap["sort_by"] = "SortBy"
	fieldNameMap["type"] = "Type_"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func roleBindingsDeleteOutputType() bindings.BindingType {
	return bindings.NewVoidType()
}

func roleBindingsDeleteRestMetadata() protocol.OperationRestMetadata {
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
	fields["vpc_id"] = bindings.NewStringType()
	fields["binding_id"] = bindings.NewStringType()
	fields["cursor"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["identity_source_id"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["identity_source_type"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["included_fields"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["name"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["page_size"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fields["path"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["role"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["root_path"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["sort_ascending"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fields["sort_by"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["type"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["org_id"] = "OrgId"
	fieldNameMap["project_id"] = "ProjectId"
	fieldNameMap["vpc_id"] = "VpcId"
	fieldNameMap["binding_id"] = "BindingId"
	fieldNameMap["cursor"] = "Cursor"
	fieldNameMap["identity_source_id"] = "IdentitySourceId"
	fieldNameMap["identity_source_type"] = "IdentitySourceType"
	fieldNameMap["included_fields"] = "IncludedFields"
	fieldNameMap["name"] = "Name"
	fieldNameMap["page_size"] = "PageSize"
	fieldNameMap["path"] = "Path"
	fieldNameMap["role"] = "Role"
	fieldNameMap["root_path"] = "RootPath"
	fieldNameMap["sort_ascending"] = "SortAscending"
	fieldNameMap["sort_by"] = "SortBy"
	fieldNameMap["type"] = "Type_"
	paramsTypeMap["identity_source_type"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["name"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["vpc_id"] = bindings.NewStringType()
	paramsTypeMap["included_fields"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["type"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["cursor"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["role"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["sort_by"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["project_id"] = bindings.NewStringType()
	paramsTypeMap["org_id"] = bindings.NewStringType()
	paramsTypeMap["root_path"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["page_size"] = bindings.NewOptionalType(bindings.NewIntegerType())
	paramsTypeMap["identity_source_id"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["path"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["binding_id"] = bindings.NewStringType()
	paramsTypeMap["sort_ascending"] = bindings.NewOptionalType(bindings.NewBooleanType())
	paramsTypeMap["orgId"] = bindings.NewStringType()
	paramsTypeMap["projectId"] = bindings.NewStringType()
	paramsTypeMap["vpcId"] = bindings.NewStringType()
	paramsTypeMap["bindingId"] = bindings.NewStringType()
	pathParams["vpc_id"] = "vpcId"
	pathParams["binding_id"] = "bindingId"
	pathParams["project_id"] = "projectId"
	pathParams["org_id"] = "orgId"
	queryParams["cursor"] = "cursor"
	queryParams["identity_source_id"] = "identity_source_id"
	queryParams["path"] = "path"
	queryParams["root_path"] = "root_path"
	queryParams["role"] = "role"
	queryParams["sort_ascending"] = "sort_ascending"
	queryParams["included_fields"] = "included_fields"
	queryParams["name"] = "name"
	queryParams["sort_by"] = "sort_by"
	queryParams["type"] = "type"
	queryParams["identity_source_type"] = "identity_source_type"
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
		"DELETE",
		"/policy/api/v1/orgs/{orgId}/projects/{projectId}/vpcs/{vpcId}/aaa/role-bindings/{bindingId}",
		"",
		resultHeaders,
		204,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func roleBindingsGetInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["org_id"] = bindings.NewStringType()
	fields["project_id"] = bindings.NewStringType()
	fields["vpc_id"] = bindings.NewStringType()
	fields["binding_id"] = bindings.NewStringType()
	fields["cursor"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["identity_source_id"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["identity_source_type"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["included_fields"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["name"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["page_size"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fields["path"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["role"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["root_path"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["sort_ascending"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fields["sort_by"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["type"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["org_id"] = "OrgId"
	fieldNameMap["project_id"] = "ProjectId"
	fieldNameMap["vpc_id"] = "VpcId"
	fieldNameMap["binding_id"] = "BindingId"
	fieldNameMap["cursor"] = "Cursor"
	fieldNameMap["identity_source_id"] = "IdentitySourceId"
	fieldNameMap["identity_source_type"] = "IdentitySourceType"
	fieldNameMap["included_fields"] = "IncludedFields"
	fieldNameMap["name"] = "Name"
	fieldNameMap["page_size"] = "PageSize"
	fieldNameMap["path"] = "Path"
	fieldNameMap["role"] = "Role"
	fieldNameMap["root_path"] = "RootPath"
	fieldNameMap["sort_ascending"] = "SortAscending"
	fieldNameMap["sort_by"] = "SortBy"
	fieldNameMap["type"] = "Type_"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func roleBindingsGetOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.RoleBindingBindingType)
}

func roleBindingsGetRestMetadata() protocol.OperationRestMetadata {
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
	fields["vpc_id"] = bindings.NewStringType()
	fields["binding_id"] = bindings.NewStringType()
	fields["cursor"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["identity_source_id"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["identity_source_type"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["included_fields"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["name"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["page_size"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fields["path"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["role"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["root_path"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["sort_ascending"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fields["sort_by"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["type"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["org_id"] = "OrgId"
	fieldNameMap["project_id"] = "ProjectId"
	fieldNameMap["vpc_id"] = "VpcId"
	fieldNameMap["binding_id"] = "BindingId"
	fieldNameMap["cursor"] = "Cursor"
	fieldNameMap["identity_source_id"] = "IdentitySourceId"
	fieldNameMap["identity_source_type"] = "IdentitySourceType"
	fieldNameMap["included_fields"] = "IncludedFields"
	fieldNameMap["name"] = "Name"
	fieldNameMap["page_size"] = "PageSize"
	fieldNameMap["path"] = "Path"
	fieldNameMap["role"] = "Role"
	fieldNameMap["root_path"] = "RootPath"
	fieldNameMap["sort_ascending"] = "SortAscending"
	fieldNameMap["sort_by"] = "SortBy"
	fieldNameMap["type"] = "Type_"
	paramsTypeMap["identity_source_type"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["name"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["vpc_id"] = bindings.NewStringType()
	paramsTypeMap["included_fields"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["type"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["cursor"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["role"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["sort_by"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["project_id"] = bindings.NewStringType()
	paramsTypeMap["org_id"] = bindings.NewStringType()
	paramsTypeMap["root_path"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["page_size"] = bindings.NewOptionalType(bindings.NewIntegerType())
	paramsTypeMap["identity_source_id"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["path"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["binding_id"] = bindings.NewStringType()
	paramsTypeMap["sort_ascending"] = bindings.NewOptionalType(bindings.NewBooleanType())
	paramsTypeMap["orgId"] = bindings.NewStringType()
	paramsTypeMap["projectId"] = bindings.NewStringType()
	paramsTypeMap["vpcId"] = bindings.NewStringType()
	paramsTypeMap["bindingId"] = bindings.NewStringType()
	pathParams["vpc_id"] = "vpcId"
	pathParams["binding_id"] = "bindingId"
	pathParams["project_id"] = "projectId"
	pathParams["org_id"] = "orgId"
	queryParams["cursor"] = "cursor"
	queryParams["identity_source_id"] = "identity_source_id"
	queryParams["path"] = "path"
	queryParams["root_path"] = "root_path"
	queryParams["role"] = "role"
	queryParams["sort_ascending"] = "sort_ascending"
	queryParams["included_fields"] = "included_fields"
	queryParams["name"] = "name"
	queryParams["sort_by"] = "sort_by"
	queryParams["type"] = "type"
	queryParams["identity_source_type"] = "identity_source_type"
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
		"/policy/api/v1/orgs/{orgId}/projects/{projectId}/vpcs/{vpcId}/aaa/role-bindings/{bindingId}",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func roleBindingsListInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["org_id"] = bindings.NewStringType()
	fields["project_id"] = bindings.NewStringType()
	fields["vpc_id"] = bindings.NewStringType()
	fields["cursor"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["identity_source_id"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["identity_source_type"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["included_fields"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["name"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["page_size"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fields["path"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["role"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["root_path"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["sort_ascending"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fields["sort_by"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["type"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["org_id"] = "OrgId"
	fieldNameMap["project_id"] = "ProjectId"
	fieldNameMap["vpc_id"] = "VpcId"
	fieldNameMap["cursor"] = "Cursor"
	fieldNameMap["identity_source_id"] = "IdentitySourceId"
	fieldNameMap["identity_source_type"] = "IdentitySourceType"
	fieldNameMap["included_fields"] = "IncludedFields"
	fieldNameMap["name"] = "Name"
	fieldNameMap["page_size"] = "PageSize"
	fieldNameMap["path"] = "Path"
	fieldNameMap["role"] = "Role"
	fieldNameMap["root_path"] = "RootPath"
	fieldNameMap["sort_ascending"] = "SortAscending"
	fieldNameMap["sort_by"] = "SortBy"
	fieldNameMap["type"] = "Type_"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func roleBindingsListOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.RoleBindingListResultBindingType)
}

func roleBindingsListRestMetadata() protocol.OperationRestMetadata {
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
	fields["vpc_id"] = bindings.NewStringType()
	fields["cursor"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["identity_source_id"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["identity_source_type"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["included_fields"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["name"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["page_size"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fields["path"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["role"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["root_path"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["sort_ascending"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fields["sort_by"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["type"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["org_id"] = "OrgId"
	fieldNameMap["project_id"] = "ProjectId"
	fieldNameMap["vpc_id"] = "VpcId"
	fieldNameMap["cursor"] = "Cursor"
	fieldNameMap["identity_source_id"] = "IdentitySourceId"
	fieldNameMap["identity_source_type"] = "IdentitySourceType"
	fieldNameMap["included_fields"] = "IncludedFields"
	fieldNameMap["name"] = "Name"
	fieldNameMap["page_size"] = "PageSize"
	fieldNameMap["path"] = "Path"
	fieldNameMap["role"] = "Role"
	fieldNameMap["root_path"] = "RootPath"
	fieldNameMap["sort_ascending"] = "SortAscending"
	fieldNameMap["sort_by"] = "SortBy"
	fieldNameMap["type"] = "Type_"
	paramsTypeMap["identity_source_type"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["name"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["vpc_id"] = bindings.NewStringType()
	paramsTypeMap["included_fields"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["type"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["cursor"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["role"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["sort_by"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["project_id"] = bindings.NewStringType()
	paramsTypeMap["org_id"] = bindings.NewStringType()
	paramsTypeMap["root_path"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["page_size"] = bindings.NewOptionalType(bindings.NewIntegerType())
	paramsTypeMap["identity_source_id"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["path"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["sort_ascending"] = bindings.NewOptionalType(bindings.NewBooleanType())
	paramsTypeMap["orgId"] = bindings.NewStringType()
	paramsTypeMap["projectId"] = bindings.NewStringType()
	paramsTypeMap["vpcId"] = bindings.NewStringType()
	pathParams["vpc_id"] = "vpcId"
	pathParams["project_id"] = "projectId"
	pathParams["org_id"] = "orgId"
	queryParams["cursor"] = "cursor"
	queryParams["identity_source_id"] = "identity_source_id"
	queryParams["path"] = "path"
	queryParams["root_path"] = "root_path"
	queryParams["role"] = "role"
	queryParams["sort_ascending"] = "sort_ascending"
	queryParams["included_fields"] = "included_fields"
	queryParams["name"] = "name"
	queryParams["sort_by"] = "sort_by"
	queryParams["type"] = "type"
	queryParams["identity_source_type"] = "identity_source_type"
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
		"/policy/api/v1/orgs/{orgId}/projects/{projectId}/vpcs/{vpcId}/aaa/role-bindings",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}
