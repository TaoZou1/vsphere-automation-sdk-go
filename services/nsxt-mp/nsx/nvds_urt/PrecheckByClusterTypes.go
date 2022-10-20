// Copyright © 2019-2021 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Data type definitions file for service: PrecheckByCluster.
// Includes binding types of a structures and enumerations defined in the service.
// Shared by client-side stubs and server-side skeletons to ensure type
// compatibility.

package nvds_urt

import (
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/runtime/bindings"
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/runtime/data"
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/runtime/protocol"
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/services/nsxt-mp/nsx/model"
	"reflect"
)

func precheckByClusterCreateInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["cluster_id"] = bindings.NewStringType()
	fields["tolerate_different_configurations"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fieldNameMap["cluster_id"] = "ClusterId"
	fieldNameMap["tolerate_different_configurations"] = "TolerateDifferentConfigurations"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func precheckByClusterCreateOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.NvdsUpgradePrecheckIdBindingType)
}

func precheckByClusterCreateRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["cluster_id"] = bindings.NewStringType()
	fields["tolerate_different_configurations"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fieldNameMap["cluster_id"] = "ClusterId"
	fieldNameMap["tolerate_different_configurations"] = "TolerateDifferentConfigurations"
	paramsTypeMap["cluster_id"] = bindings.NewStringType()
	paramsTypeMap["tolerate_different_configurations"] = bindings.NewOptionalType(bindings.NewBooleanType())
	paramsTypeMap["clusterId"] = bindings.NewStringType()
	pathParams["cluster_id"] = "clusterId"
	queryParams["tolerate_different_configurations"] = "tolerate_different_configurations"
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
		"POST",
		"/api/v1/nvds-urt/precheck-by-cluster/{clusterId}",
		"",
		resultHeaders,
		201,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func precheckByClusterGetInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["cluster_id"] = bindings.NewStringType()
	fieldNameMap["cluster_id"] = "ClusterId"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func precheckByClusterGetOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.NvdsUpgradePrecheckIdBindingType)
}

func precheckByClusterGetRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["cluster_id"] = bindings.NewStringType()
	fieldNameMap["cluster_id"] = "ClusterId"
	paramsTypeMap["cluster_id"] = bindings.NewStringType()
	paramsTypeMap["clusterId"] = bindings.NewStringType()
	pathParams["cluster_id"] = "clusterId"
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
		"/api/v1/nvds-urt/precheck-by-cluster/{clusterId}",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}
