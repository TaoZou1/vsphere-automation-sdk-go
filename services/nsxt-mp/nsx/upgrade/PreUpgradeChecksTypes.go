// Copyright © 2019-2021 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Data type definitions file for service: PreUpgradeChecks.
// Includes binding types of a structures and enumerations defined in the service.
// Shared by client-side stubs and server-side skeletons to ensure type
// compatibility.

package upgrade

import (
	"github.com/TaoZou1/vsphere-automation-sdk-go/runtime/bindings"
	"github.com/TaoZou1/vsphere-automation-sdk-go/runtime/data"
	"github.com/TaoZou1/vsphere-automation-sdk-go/runtime/protocol"
	"reflect"
)

func preUpgradeChecksAcknowledgeInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["pre_upgrade_check_id"] = bindings.NewStringType()
	fieldNameMap["pre_upgrade_check_id"] = "PreUpgradeCheckId"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func preUpgradeChecksAcknowledgeOutputType() bindings.BindingType {
	return bindings.NewVoidType()
}

func preUpgradeChecksAcknowledgeRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["pre_upgrade_check_id"] = bindings.NewStringType()
	fieldNameMap["pre_upgrade_check_id"] = "PreUpgradeCheckId"
	paramsTypeMap["pre_upgrade_check_id"] = bindings.NewStringType()
	paramsTypeMap["preUpgradeCheckId"] = bindings.NewStringType()
	pathParams["pre_upgrade_check_id"] = "preUpgradeCheckId"
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
		"action=acknowledge",
		"",
		"POST",
		"/api/v1/upgrade/pre-upgrade-checks/{preUpgradeCheckId}",
		"",
		resultHeaders,
		204,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func preUpgradeChecksResolveInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["pre_upgrade_check_id"] = bindings.NewStringType()
	fieldNameMap["pre_upgrade_check_id"] = "PreUpgradeCheckId"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func preUpgradeChecksResolveOutputType() bindings.BindingType {
	return bindings.NewVoidType()
}

func preUpgradeChecksResolveRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["pre_upgrade_check_id"] = bindings.NewStringType()
	fieldNameMap["pre_upgrade_check_id"] = "PreUpgradeCheckId"
	paramsTypeMap["pre_upgrade_check_id"] = bindings.NewStringType()
	paramsTypeMap["preUpgradeCheckId"] = bindings.NewStringType()
	pathParams["pre_upgrade_check_id"] = "preUpgradeCheckId"
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
		"action=resolve",
		"",
		"POST",
		"/api/v1/upgrade/pre-upgrade-checks/{preUpgradeCheckId}",
		"",
		resultHeaders,
		204,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}
