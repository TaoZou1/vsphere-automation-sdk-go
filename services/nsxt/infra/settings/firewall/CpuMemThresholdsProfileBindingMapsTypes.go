// Copyright © 2019-2021 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Data type definitions file for service: CpuMemThresholdsProfileBindingMaps.
// Includes binding types of a structures and enumerations defined in the service.
// Shared by client-side stubs and server-side skeletons to ensure type
// compatibility.

package firewall

import (
	"github.com/vmware/vsphere-automation-sdk-go/runtime/bindings"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/data"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/protocol"
	"github.com/vmware/vsphere-automation-sdk-go/services/nsxt/model"
	"reflect"
)

func cpuMemThresholdsProfileBindingMapsDeleteInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["cpu_mem_thresholds_profile_binding_map_id"] = bindings.NewStringType()
	fieldNameMap["cpu_mem_thresholds_profile_binding_map_id"] = "CpuMemThresholdsProfileBindingMapId"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func cpuMemThresholdsProfileBindingMapsDeleteOutputType() bindings.BindingType {
	return bindings.NewVoidType()
}

func cpuMemThresholdsProfileBindingMapsDeleteRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["cpu_mem_thresholds_profile_binding_map_id"] = bindings.NewStringType()
	fieldNameMap["cpu_mem_thresholds_profile_binding_map_id"] = "CpuMemThresholdsProfileBindingMapId"
	paramsTypeMap["cpu_mem_thresholds_profile_binding_map_id"] = bindings.NewStringType()
	paramsTypeMap["cpuMemThresholdsProfileBindingMapId"] = bindings.NewStringType()
	pathParams["cpu_mem_thresholds_profile_binding_map_id"] = "cpuMemThresholdsProfileBindingMapId"
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
		"/policy/api/v1/infra/settings/firewall/cpu-mem-thresholds-profile-binding-maps/{cpuMemThresholdsProfileBindingMapId}",
		"",
		resultHeaders,
		204,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func cpuMemThresholdsProfileBindingMapsGetInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["cpu_mem_thresholds_profile_binding_map_id"] = bindings.NewStringType()
	fieldNameMap["cpu_mem_thresholds_profile_binding_map_id"] = "CpuMemThresholdsProfileBindingMapId"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func cpuMemThresholdsProfileBindingMapsGetOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.PolicyFirewallCPUMemThresholdsProfileBindingMapBindingType)
}

func cpuMemThresholdsProfileBindingMapsGetRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["cpu_mem_thresholds_profile_binding_map_id"] = bindings.NewStringType()
	fieldNameMap["cpu_mem_thresholds_profile_binding_map_id"] = "CpuMemThresholdsProfileBindingMapId"
	paramsTypeMap["cpu_mem_thresholds_profile_binding_map_id"] = bindings.NewStringType()
	paramsTypeMap["cpuMemThresholdsProfileBindingMapId"] = bindings.NewStringType()
	pathParams["cpu_mem_thresholds_profile_binding_map_id"] = "cpuMemThresholdsProfileBindingMapId"
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
		"/policy/api/v1/infra/settings/firewall/cpu-mem-thresholds-profile-binding-maps/{cpuMemThresholdsProfileBindingMapId}",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func cpuMemThresholdsProfileBindingMapsListInputType() bindings.StructType {
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

func cpuMemThresholdsProfileBindingMapsListOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.PolicyFirewallCPUMemThresholdsProfileBindingMapListResultBindingType)
}

func cpuMemThresholdsProfileBindingMapsListRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
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
		"/policy/api/v1/infra/settings/firewall/cpu-mem-thresholds-profile-binding-maps",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func cpuMemThresholdsProfileBindingMapsPatchInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["cpu_mem_thresholds_profile_binding_map_id"] = bindings.NewStringType()
	fields["policy_firewall_CPU_mem_thresholds_profile_binding_map"] = bindings.NewReferenceType(model.PolicyFirewallCPUMemThresholdsProfileBindingMapBindingType)
	fieldNameMap["cpu_mem_thresholds_profile_binding_map_id"] = "CpuMemThresholdsProfileBindingMapId"
	fieldNameMap["policy_firewall_CPU_mem_thresholds_profile_binding_map"] = "PolicyFirewallCPUMemThresholdsProfileBindingMap"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func cpuMemThresholdsProfileBindingMapsPatchOutputType() bindings.BindingType {
	return bindings.NewVoidType()
}

func cpuMemThresholdsProfileBindingMapsPatchRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["cpu_mem_thresholds_profile_binding_map_id"] = bindings.NewStringType()
	fields["policy_firewall_CPU_mem_thresholds_profile_binding_map"] = bindings.NewReferenceType(model.PolicyFirewallCPUMemThresholdsProfileBindingMapBindingType)
	fieldNameMap["cpu_mem_thresholds_profile_binding_map_id"] = "CpuMemThresholdsProfileBindingMapId"
	fieldNameMap["policy_firewall_CPU_mem_thresholds_profile_binding_map"] = "PolicyFirewallCPUMemThresholdsProfileBindingMap"
	paramsTypeMap["policy_firewall_CPU_mem_thresholds_profile_binding_map"] = bindings.NewReferenceType(model.PolicyFirewallCPUMemThresholdsProfileBindingMapBindingType)
	paramsTypeMap["cpu_mem_thresholds_profile_binding_map_id"] = bindings.NewStringType()
	paramsTypeMap["cpuMemThresholdsProfileBindingMapId"] = bindings.NewStringType()
	pathParams["cpu_mem_thresholds_profile_binding_map_id"] = "cpuMemThresholdsProfileBindingMapId"
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
		"policy_firewall_CPU_mem_thresholds_profile_binding_map",
		"PATCH",
		"/policy/api/v1/infra/settings/firewall/cpu-mem-thresholds-profile-binding-maps/{cpuMemThresholdsProfileBindingMapId}",
		"",
		resultHeaders,
		204,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func cpuMemThresholdsProfileBindingMapsUpdateInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["cpu_mem_thresholds_profile_binding_map_id"] = bindings.NewStringType()
	fields["policy_firewall_CPU_mem_thresholds_profile_binding_map"] = bindings.NewReferenceType(model.PolicyFirewallCPUMemThresholdsProfileBindingMapBindingType)
	fieldNameMap["cpu_mem_thresholds_profile_binding_map_id"] = "CpuMemThresholdsProfileBindingMapId"
	fieldNameMap["policy_firewall_CPU_mem_thresholds_profile_binding_map"] = "PolicyFirewallCPUMemThresholdsProfileBindingMap"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func cpuMemThresholdsProfileBindingMapsUpdateOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.PolicyFirewallCPUMemThresholdsProfileBindingMapBindingType)
}

func cpuMemThresholdsProfileBindingMapsUpdateRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["cpu_mem_thresholds_profile_binding_map_id"] = bindings.NewStringType()
	fields["policy_firewall_CPU_mem_thresholds_profile_binding_map"] = bindings.NewReferenceType(model.PolicyFirewallCPUMemThresholdsProfileBindingMapBindingType)
	fieldNameMap["cpu_mem_thresholds_profile_binding_map_id"] = "CpuMemThresholdsProfileBindingMapId"
	fieldNameMap["policy_firewall_CPU_mem_thresholds_profile_binding_map"] = "PolicyFirewallCPUMemThresholdsProfileBindingMap"
	paramsTypeMap["policy_firewall_CPU_mem_thresholds_profile_binding_map"] = bindings.NewReferenceType(model.PolicyFirewallCPUMemThresholdsProfileBindingMapBindingType)
	paramsTypeMap["cpu_mem_thresholds_profile_binding_map_id"] = bindings.NewStringType()
	paramsTypeMap["cpuMemThresholdsProfileBindingMapId"] = bindings.NewStringType()
	pathParams["cpu_mem_thresholds_profile_binding_map_id"] = "cpuMemThresholdsProfileBindingMapId"
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
		"policy_firewall_CPU_mem_thresholds_profile_binding_map",
		"PUT",
		"/policy/api/v1/infra/settings/firewall/cpu-mem-thresholds-profile-binding-maps/{cpuMemThresholdsProfileBindingMapId}",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}
