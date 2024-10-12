// Copyright © 2019-2021 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Data type definitions file for service: Children.
// Includes binding types of a structures and enumerations defined in the service.
// Shared by client-side stubs and server-side skeletons to ensure type
// compatibility.

package groups

import (
	"github.com/vmware/vsphere-automation-sdk-go/runtime/bindings"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/data"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/protocol"
	"github.com/vmware/vsphere-automation-sdk-go/services/nsxt/model"
	"reflect"
)

func childrenListInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["identity_firewall_store_id"] = bindings.NewStringType()
	fields["group_id"] = bindings.NewStringType()
	fields["include_users"] = bindings.NewBooleanType()
	fieldNameMap["identity_firewall_store_id"] = "IdentityFirewallStoreId"
	fieldNameMap["group_id"] = "GroupId"
	fieldNameMap["include_users"] = "IncludeUsers"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func childrenListOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.IdentityFirewallStoreUserOrGrouptListResultsBindingType)
}

func childrenListRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["identity_firewall_store_id"] = bindings.NewStringType()
	fields["group_id"] = bindings.NewStringType()
	fields["include_users"] = bindings.NewBooleanType()
	fieldNameMap["identity_firewall_store_id"] = "IdentityFirewallStoreId"
	fieldNameMap["group_id"] = "GroupId"
	fieldNameMap["include_users"] = "IncludeUsers"
	paramsTypeMap["include_users"] = bindings.NewBooleanType()
	paramsTypeMap["group_id"] = bindings.NewStringType()
	paramsTypeMap["identity_firewall_store_id"] = bindings.NewStringType()
	paramsTypeMap["identityFirewallStoreId"] = bindings.NewStringType()
	paramsTypeMap["groupId"] = bindings.NewStringType()
	pathParams["identity_firewall_store_id"] = "identityFirewallStoreId"
	pathParams["group_id"] = "groupId"
	queryParams["include_users"] = "include_users"
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
		"/policy/api/v1/infra/identity-firewall-stores/{identityFirewallStoreId}/groups/{groupId}/children",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}
