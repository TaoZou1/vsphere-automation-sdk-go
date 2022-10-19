// Copyright © 2019-2021 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Data type definitions file for service: EventLogServers.
// Includes binding types of a structures and enumerations defined in the service.
// Shared by client-side stubs and server-side skeletons to ensure type
// compatibility.

package identity_firewall_stores

import (
	"github.com/TaoZou1/vsphere-automation-sdk-go/runtime/bindings"
	"github.com/TaoZou1/vsphere-automation-sdk-go/runtime/data"
	"github.com/TaoZou1/vsphere-automation-sdk-go/runtime/protocol"
	"github.com/TaoZou1/vsphere-automation-sdk-go/services/nsxt/model"
	"reflect"
)

func eventLogServersDeleteInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["identity_firewall_store_id"] = bindings.NewStringType()
	fields["event_log_server_id"] = bindings.NewStringType()
	fields["enforcement_point_path"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["identity_firewall_store_id"] = "IdentityFirewallStoreId"
	fieldNameMap["event_log_server_id"] = "EventLogServerId"
	fieldNameMap["enforcement_point_path"] = "EnforcementPointPath"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func eventLogServersDeleteOutputType() bindings.BindingType {
	return bindings.NewVoidType()
}

func eventLogServersDeleteRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["identity_firewall_store_id"] = bindings.NewStringType()
	fields["event_log_server_id"] = bindings.NewStringType()
	fields["enforcement_point_path"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["identity_firewall_store_id"] = "IdentityFirewallStoreId"
	fieldNameMap["event_log_server_id"] = "EventLogServerId"
	fieldNameMap["enforcement_point_path"] = "EnforcementPointPath"
	paramsTypeMap["enforcement_point_path"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["event_log_server_id"] = bindings.NewStringType()
	paramsTypeMap["identity_firewall_store_id"] = bindings.NewStringType()
	paramsTypeMap["identityFirewallStoreId"] = bindings.NewStringType()
	paramsTypeMap["eventLogServerId"] = bindings.NewStringType()
	pathParams["identity_firewall_store_id"] = "identityFirewallStoreId"
	pathParams["event_log_server_id"] = "eventLogServerId"
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
		"/policy/api/v1/infra/identity-firewall-stores/{identityFirewallStoreId}/event-log-servers/{eventLogServerId}",
		"",
		resultHeaders,
		204,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func eventLogServersGetInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["identity_firewall_store_id"] = bindings.NewStringType()
	fields["event_log_server_id"] = bindings.NewStringType()
	fields["enforcement_point_path"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["identity_firewall_store_id"] = "IdentityFirewallStoreId"
	fieldNameMap["event_log_server_id"] = "EventLogServerId"
	fieldNameMap["enforcement_point_path"] = "EnforcementPointPath"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func eventLogServersGetOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.IdentityFirewallStoreEventLogServerBindingType)
}

func eventLogServersGetRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["identity_firewall_store_id"] = bindings.NewStringType()
	fields["event_log_server_id"] = bindings.NewStringType()
	fields["enforcement_point_path"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["identity_firewall_store_id"] = "IdentityFirewallStoreId"
	fieldNameMap["event_log_server_id"] = "EventLogServerId"
	fieldNameMap["enforcement_point_path"] = "EnforcementPointPath"
	paramsTypeMap["enforcement_point_path"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["event_log_server_id"] = bindings.NewStringType()
	paramsTypeMap["identity_firewall_store_id"] = bindings.NewStringType()
	paramsTypeMap["identityFirewallStoreId"] = bindings.NewStringType()
	paramsTypeMap["eventLogServerId"] = bindings.NewStringType()
	pathParams["identity_firewall_store_id"] = "identityFirewallStoreId"
	pathParams["event_log_server_id"] = "eventLogServerId"
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
		"/policy/api/v1/infra/identity-firewall-stores/{identityFirewallStoreId}/event-log-servers/{eventLogServerId}",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func eventLogServersPatchInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["identity_firewall_store_id"] = bindings.NewStringType()
	fields["event_log_server_id"] = bindings.NewStringType()
	fields["identity_firewall_store_event_log_server"] = bindings.NewReferenceType(model.IdentityFirewallStoreEventLogServerBindingType)
	fields["enforcement_point_path"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["identity_firewall_store_id"] = "IdentityFirewallStoreId"
	fieldNameMap["event_log_server_id"] = "EventLogServerId"
	fieldNameMap["identity_firewall_store_event_log_server"] = "IdentityFirewallStoreEventLogServer"
	fieldNameMap["enforcement_point_path"] = "EnforcementPointPath"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func eventLogServersPatchOutputType() bindings.BindingType {
	return bindings.NewVoidType()
}

func eventLogServersPatchRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["identity_firewall_store_id"] = bindings.NewStringType()
	fields["event_log_server_id"] = bindings.NewStringType()
	fields["identity_firewall_store_event_log_server"] = bindings.NewReferenceType(model.IdentityFirewallStoreEventLogServerBindingType)
	fields["enforcement_point_path"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["identity_firewall_store_id"] = "IdentityFirewallStoreId"
	fieldNameMap["event_log_server_id"] = "EventLogServerId"
	fieldNameMap["identity_firewall_store_event_log_server"] = "IdentityFirewallStoreEventLogServer"
	fieldNameMap["enforcement_point_path"] = "EnforcementPointPath"
	paramsTypeMap["identity_firewall_store_event_log_server"] = bindings.NewReferenceType(model.IdentityFirewallStoreEventLogServerBindingType)
	paramsTypeMap["enforcement_point_path"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["event_log_server_id"] = bindings.NewStringType()
	paramsTypeMap["identity_firewall_store_id"] = bindings.NewStringType()
	paramsTypeMap["identityFirewallStoreId"] = bindings.NewStringType()
	paramsTypeMap["eventLogServerId"] = bindings.NewStringType()
	pathParams["identity_firewall_store_id"] = "identityFirewallStoreId"
	pathParams["event_log_server_id"] = "eventLogServerId"
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
		"identity_firewall_store_event_log_server",
		"PATCH",
		"/policy/api/v1/infra/identity-firewall-stores/{identityFirewallStoreId}/event-log-servers/{eventLogServerId}",
		"",
		resultHeaders,
		204,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func eventLogServersUpdateInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["identity_firewall_store_id"] = bindings.NewStringType()
	fields["event_log_server_id"] = bindings.NewStringType()
	fields["identity_firewall_store_event_log_server"] = bindings.NewReferenceType(model.IdentityFirewallStoreEventLogServerBindingType)
	fields["enforcement_point_path"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["identity_firewall_store_id"] = "IdentityFirewallStoreId"
	fieldNameMap["event_log_server_id"] = "EventLogServerId"
	fieldNameMap["identity_firewall_store_event_log_server"] = "IdentityFirewallStoreEventLogServer"
	fieldNameMap["enforcement_point_path"] = "EnforcementPointPath"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func eventLogServersUpdateOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.IdentityFirewallStoreEventLogServerBindingType)
}

func eventLogServersUpdateRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["identity_firewall_store_id"] = bindings.NewStringType()
	fields["event_log_server_id"] = bindings.NewStringType()
	fields["identity_firewall_store_event_log_server"] = bindings.NewReferenceType(model.IdentityFirewallStoreEventLogServerBindingType)
	fields["enforcement_point_path"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["identity_firewall_store_id"] = "IdentityFirewallStoreId"
	fieldNameMap["event_log_server_id"] = "EventLogServerId"
	fieldNameMap["identity_firewall_store_event_log_server"] = "IdentityFirewallStoreEventLogServer"
	fieldNameMap["enforcement_point_path"] = "EnforcementPointPath"
	paramsTypeMap["identity_firewall_store_event_log_server"] = bindings.NewReferenceType(model.IdentityFirewallStoreEventLogServerBindingType)
	paramsTypeMap["enforcement_point_path"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["event_log_server_id"] = bindings.NewStringType()
	paramsTypeMap["identity_firewall_store_id"] = bindings.NewStringType()
	paramsTypeMap["identityFirewallStoreId"] = bindings.NewStringType()
	paramsTypeMap["eventLogServerId"] = bindings.NewStringType()
	pathParams["identity_firewall_store_id"] = "identityFirewallStoreId"
	pathParams["event_log_server_id"] = "eventLogServerId"
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
		"identity_firewall_store_event_log_server",
		"PUT",
		"/policy/api/v1/infra/identity-firewall-stores/{identityFirewallStoreId}/event-log-servers/{eventLogServerId}",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}
