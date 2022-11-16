// Copyright © 2019-2021 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Data type definitions file for service: RoleAssignmentPermissionConfig.
// Includes binding types of a structures and enumerations defined in the service.
// Shared by client-side stubs and server-side skeletons to ensure type
// compatibility.

package roles

import (
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/runtime/bindings"
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/runtime/data"
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/runtime/protocol"
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/services/nsxt/model"
	"reflect"
)

func roleAssignmentPermissionConfigUpdateInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["role"] = bindings.NewStringType()
	fields["role_assignment_permission_config"] = bindings.NewReferenceType(model.RoleAssignmentPermissionConfigBindingType)
	fieldNameMap["role"] = "Role"
	fieldNameMap["role_assignment_permission_config"] = "RoleAssignmentPermissionConfig"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func roleAssignmentPermissionConfigUpdateOutputType() bindings.BindingType {
	return bindings.NewVoidType()
}

func roleAssignmentPermissionConfigUpdateRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["role"] = bindings.NewStringType()
	fields["role_assignment_permission_config"] = bindings.NewReferenceType(model.RoleAssignmentPermissionConfigBindingType)
	fieldNameMap["role"] = "Role"
	fieldNameMap["role_assignment_permission_config"] = "RoleAssignmentPermissionConfig"
	paramsTypeMap["role"] = bindings.NewStringType()
	paramsTypeMap["role_assignment_permission_config"] = bindings.NewReferenceType(model.RoleAssignmentPermissionConfigBindingType)
	paramsTypeMap["role"] = bindings.NewStringType()
	pathParams["role"] = "role"
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
		"role_assignment_permission_config",
		"PUT",
		"/policy/api/v1/aaa/roles/{role}/role-assignment-permission-config",
		"",
		resultHeaders,
		204,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}
