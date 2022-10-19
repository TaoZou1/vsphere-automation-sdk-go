// Copyright © 2019-2021 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Data type definitions file for service: SegmentMonitoringProfileBindingMaps.
// Includes binding types of a structures and enumerations defined in the service.
// Shared by client-side stubs and server-side skeletons to ensure type
// compatibility.

package segments

import (
	"github.com/TaoZou1/vsphere-automation-sdk-go/runtime/bindings"
	"github.com/TaoZou1/vsphere-automation-sdk-go/runtime/data"
	"github.com/TaoZou1/vsphere-automation-sdk-go/runtime/protocol"
	"github.com/TaoZou1/vsphere-automation-sdk-go/services/nsxt/model"
	"reflect"
)

func segmentMonitoringProfileBindingMapsDeleteInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["infra_segment_id"] = bindings.NewStringType()
	fields["segment_monitoring_profile_binding_map_id"] = bindings.NewStringType()
	fields["org_id"] = bindings.NewStringType()
	fields["project_id"] = bindings.NewStringType()
	fieldNameMap["infra_segment_id"] = "InfraSegmentId"
	fieldNameMap["segment_monitoring_profile_binding_map_id"] = "SegmentMonitoringProfileBindingMapId"
	fieldNameMap["org_id"] = "OrgId"
	fieldNameMap["project_id"] = "ProjectId"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func segmentMonitoringProfileBindingMapsDeleteOutputType() bindings.BindingType {
	return bindings.NewVoidType()
}

func segmentMonitoringProfileBindingMapsDeleteRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["infra_segment_id"] = bindings.NewStringType()
	fields["segment_monitoring_profile_binding_map_id"] = bindings.NewStringType()
	fields["org_id"] = bindings.NewStringType()
	fields["project_id"] = bindings.NewStringType()
	fieldNameMap["infra_segment_id"] = "InfraSegmentId"
	fieldNameMap["segment_monitoring_profile_binding_map_id"] = "SegmentMonitoringProfileBindingMapId"
	fieldNameMap["org_id"] = "OrgId"
	fieldNameMap["project_id"] = "ProjectId"
	paramsTypeMap["project_id"] = bindings.NewStringType()
	paramsTypeMap["org_id"] = bindings.NewStringType()
	paramsTypeMap["segment_monitoring_profile_binding_map_id"] = bindings.NewStringType()
	paramsTypeMap["infra_segment_id"] = bindings.NewStringType()
	paramsTypeMap["orgId"] = bindings.NewStringType()
	paramsTypeMap["projectId"] = bindings.NewStringType()
	paramsTypeMap["infraSegmentId"] = bindings.NewStringType()
	paramsTypeMap["segmentMonitoringProfileBindingMapId"] = bindings.NewStringType()
	pathParams["segment_monitoring_profile_binding_map_id"] = "segmentMonitoringProfileBindingMapId"
	pathParams["infra_segment_id"] = "infraSegmentId"
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
		"/policy/api/v1/orgs/{orgId}/projects/{projectId}/infra/segments/{infraSegmentId}/segment-monitoring-profile-binding-maps/{segmentMonitoringProfileBindingMapId}",
		"",
		resultHeaders,
		204,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func segmentMonitoringProfileBindingMapsGetInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["infra_segment_id"] = bindings.NewStringType()
	fields["segment_monitoring_profile_binding_map_id"] = bindings.NewStringType()
	fields["org_id"] = bindings.NewStringType()
	fields["project_id"] = bindings.NewStringType()
	fieldNameMap["infra_segment_id"] = "InfraSegmentId"
	fieldNameMap["segment_monitoring_profile_binding_map_id"] = "SegmentMonitoringProfileBindingMapId"
	fieldNameMap["org_id"] = "OrgId"
	fieldNameMap["project_id"] = "ProjectId"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func segmentMonitoringProfileBindingMapsGetOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.SegmentMonitoringProfileBindingMapBindingType)
}

func segmentMonitoringProfileBindingMapsGetRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["infra_segment_id"] = bindings.NewStringType()
	fields["segment_monitoring_profile_binding_map_id"] = bindings.NewStringType()
	fields["org_id"] = bindings.NewStringType()
	fields["project_id"] = bindings.NewStringType()
	fieldNameMap["infra_segment_id"] = "InfraSegmentId"
	fieldNameMap["segment_monitoring_profile_binding_map_id"] = "SegmentMonitoringProfileBindingMapId"
	fieldNameMap["org_id"] = "OrgId"
	fieldNameMap["project_id"] = "ProjectId"
	paramsTypeMap["project_id"] = bindings.NewStringType()
	paramsTypeMap["org_id"] = bindings.NewStringType()
	paramsTypeMap["segment_monitoring_profile_binding_map_id"] = bindings.NewStringType()
	paramsTypeMap["infra_segment_id"] = bindings.NewStringType()
	paramsTypeMap["orgId"] = bindings.NewStringType()
	paramsTypeMap["projectId"] = bindings.NewStringType()
	paramsTypeMap["infraSegmentId"] = bindings.NewStringType()
	paramsTypeMap["segmentMonitoringProfileBindingMapId"] = bindings.NewStringType()
	pathParams["segment_monitoring_profile_binding_map_id"] = "segmentMonitoringProfileBindingMapId"
	pathParams["infra_segment_id"] = "infraSegmentId"
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
		"/policy/api/v1/orgs/{orgId}/projects/{projectId}/infra/segments/{infraSegmentId}/segment-monitoring-profile-binding-maps/{segmentMonitoringProfileBindingMapId}",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func segmentMonitoringProfileBindingMapsListInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["org_id"] = bindings.NewStringType()
	fields["project_id"] = bindings.NewStringType()
	fields["infra_segment_id"] = bindings.NewStringType()
	fields["cursor"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["include_mark_for_delete_objects"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fields["included_fields"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["page_size"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fields["sort_ascending"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fields["sort_by"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["org_id"] = "OrgId"
	fieldNameMap["project_id"] = "ProjectId"
	fieldNameMap["infra_segment_id"] = "InfraSegmentId"
	fieldNameMap["cursor"] = "Cursor"
	fieldNameMap["include_mark_for_delete_objects"] = "IncludeMarkForDeleteObjects"
	fieldNameMap["included_fields"] = "IncludedFields"
	fieldNameMap["page_size"] = "PageSize"
	fieldNameMap["sort_ascending"] = "SortAscending"
	fieldNameMap["sort_by"] = "SortBy"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func segmentMonitoringProfileBindingMapsListOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.SegmentMonitoringProfileBindingMapListResultBindingType)
}

func segmentMonitoringProfileBindingMapsListRestMetadata() protocol.OperationRestMetadata {
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
	fields["infra_segment_id"] = bindings.NewStringType()
	fields["cursor"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["include_mark_for_delete_objects"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fields["included_fields"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["page_size"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fields["sort_ascending"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fields["sort_by"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["org_id"] = "OrgId"
	fieldNameMap["project_id"] = "ProjectId"
	fieldNameMap["infra_segment_id"] = "InfraSegmentId"
	fieldNameMap["cursor"] = "Cursor"
	fieldNameMap["include_mark_for_delete_objects"] = "IncludeMarkForDeleteObjects"
	fieldNameMap["included_fields"] = "IncludedFields"
	fieldNameMap["page_size"] = "PageSize"
	fieldNameMap["sort_ascending"] = "SortAscending"
	fieldNameMap["sort_by"] = "SortBy"
	paramsTypeMap["project_id"] = bindings.NewStringType()
	paramsTypeMap["org_id"] = bindings.NewStringType()
	paramsTypeMap["included_fields"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["page_size"] = bindings.NewOptionalType(bindings.NewIntegerType())
	paramsTypeMap["infra_segment_id"] = bindings.NewStringType()
	paramsTypeMap["include_mark_for_delete_objects"] = bindings.NewOptionalType(bindings.NewBooleanType())
	paramsTypeMap["cursor"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["sort_by"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["sort_ascending"] = bindings.NewOptionalType(bindings.NewBooleanType())
	paramsTypeMap["orgId"] = bindings.NewStringType()
	paramsTypeMap["projectId"] = bindings.NewStringType()
	paramsTypeMap["infraSegmentId"] = bindings.NewStringType()
	pathParams["project_id"] = "projectId"
	pathParams["infra_segment_id"] = "infraSegmentId"
	pathParams["org_id"] = "orgId"
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
		"/policy/api/v1/orgs/{orgId}/projects/{projectId}/infra/segments/{infraSegmentId}/segment-monitoring-profile-binding-maps",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func segmentMonitoringProfileBindingMapsPatchInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["infra_segment_id"] = bindings.NewStringType()
	fields["segment_monitoring_profile_binding_map_id"] = bindings.NewStringType()
	fields["org_id"] = bindings.NewStringType()
	fields["project_id"] = bindings.NewStringType()
	fields["segment_monitoring_profile_binding_map"] = bindings.NewReferenceType(model.SegmentMonitoringProfileBindingMapBindingType)
	fieldNameMap["infra_segment_id"] = "InfraSegmentId"
	fieldNameMap["segment_monitoring_profile_binding_map_id"] = "SegmentMonitoringProfileBindingMapId"
	fieldNameMap["org_id"] = "OrgId"
	fieldNameMap["project_id"] = "ProjectId"
	fieldNameMap["segment_monitoring_profile_binding_map"] = "SegmentMonitoringProfileBindingMap"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func segmentMonitoringProfileBindingMapsPatchOutputType() bindings.BindingType {
	return bindings.NewVoidType()
}

func segmentMonitoringProfileBindingMapsPatchRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["infra_segment_id"] = bindings.NewStringType()
	fields["segment_monitoring_profile_binding_map_id"] = bindings.NewStringType()
	fields["org_id"] = bindings.NewStringType()
	fields["project_id"] = bindings.NewStringType()
	fields["segment_monitoring_profile_binding_map"] = bindings.NewReferenceType(model.SegmentMonitoringProfileBindingMapBindingType)
	fieldNameMap["infra_segment_id"] = "InfraSegmentId"
	fieldNameMap["segment_monitoring_profile_binding_map_id"] = "SegmentMonitoringProfileBindingMapId"
	fieldNameMap["org_id"] = "OrgId"
	fieldNameMap["project_id"] = "ProjectId"
	fieldNameMap["segment_monitoring_profile_binding_map"] = "SegmentMonitoringProfileBindingMap"
	paramsTypeMap["project_id"] = bindings.NewStringType()
	paramsTypeMap["org_id"] = bindings.NewStringType()
	paramsTypeMap["segment_monitoring_profile_binding_map_id"] = bindings.NewStringType()
	paramsTypeMap["segment_monitoring_profile_binding_map"] = bindings.NewReferenceType(model.SegmentMonitoringProfileBindingMapBindingType)
	paramsTypeMap["infra_segment_id"] = bindings.NewStringType()
	paramsTypeMap["orgId"] = bindings.NewStringType()
	paramsTypeMap["projectId"] = bindings.NewStringType()
	paramsTypeMap["infraSegmentId"] = bindings.NewStringType()
	paramsTypeMap["segmentMonitoringProfileBindingMapId"] = bindings.NewStringType()
	pathParams["segment_monitoring_profile_binding_map_id"] = "segmentMonitoringProfileBindingMapId"
	pathParams["infra_segment_id"] = "infraSegmentId"
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
		"segment_monitoring_profile_binding_map",
		"PATCH",
		"/policy/api/v1/orgs/{orgId}/projects/{projectId}/infra/segments/{infraSegmentId}/segment-monitoring-profile-binding-maps/{segmentMonitoringProfileBindingMapId}",
		"",
		resultHeaders,
		204,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func segmentMonitoringProfileBindingMapsUpdateInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["infra_segment_id"] = bindings.NewStringType()
	fields["segment_monitoring_profile_binding_map_id"] = bindings.NewStringType()
	fields["org_id"] = bindings.NewStringType()
	fields["project_id"] = bindings.NewStringType()
	fields["segment_monitoring_profile_binding_map"] = bindings.NewReferenceType(model.SegmentMonitoringProfileBindingMapBindingType)
	fieldNameMap["infra_segment_id"] = "InfraSegmentId"
	fieldNameMap["segment_monitoring_profile_binding_map_id"] = "SegmentMonitoringProfileBindingMapId"
	fieldNameMap["org_id"] = "OrgId"
	fieldNameMap["project_id"] = "ProjectId"
	fieldNameMap["segment_monitoring_profile_binding_map"] = "SegmentMonitoringProfileBindingMap"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func segmentMonitoringProfileBindingMapsUpdateOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.SegmentMonitoringProfileBindingMapBindingType)
}

func segmentMonitoringProfileBindingMapsUpdateRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["infra_segment_id"] = bindings.NewStringType()
	fields["segment_monitoring_profile_binding_map_id"] = bindings.NewStringType()
	fields["org_id"] = bindings.NewStringType()
	fields["project_id"] = bindings.NewStringType()
	fields["segment_monitoring_profile_binding_map"] = bindings.NewReferenceType(model.SegmentMonitoringProfileBindingMapBindingType)
	fieldNameMap["infra_segment_id"] = "InfraSegmentId"
	fieldNameMap["segment_monitoring_profile_binding_map_id"] = "SegmentMonitoringProfileBindingMapId"
	fieldNameMap["org_id"] = "OrgId"
	fieldNameMap["project_id"] = "ProjectId"
	fieldNameMap["segment_monitoring_profile_binding_map"] = "SegmentMonitoringProfileBindingMap"
	paramsTypeMap["project_id"] = bindings.NewStringType()
	paramsTypeMap["org_id"] = bindings.NewStringType()
	paramsTypeMap["segment_monitoring_profile_binding_map_id"] = bindings.NewStringType()
	paramsTypeMap["segment_monitoring_profile_binding_map"] = bindings.NewReferenceType(model.SegmentMonitoringProfileBindingMapBindingType)
	paramsTypeMap["infra_segment_id"] = bindings.NewStringType()
	paramsTypeMap["orgId"] = bindings.NewStringType()
	paramsTypeMap["projectId"] = bindings.NewStringType()
	paramsTypeMap["infraSegmentId"] = bindings.NewStringType()
	paramsTypeMap["segmentMonitoringProfileBindingMapId"] = bindings.NewStringType()
	pathParams["segment_monitoring_profile_binding_map_id"] = "segmentMonitoringProfileBindingMapId"
	pathParams["infra_segment_id"] = "infraSegmentId"
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
		"segment_monitoring_profile_binding_map",
		"PUT",
		"/policy/api/v1/orgs/{orgId}/projects/{projectId}/infra/segments/{infraSegmentId}/segment-monitoring-profile-binding-maps/{segmentMonitoringProfileBindingMapId}",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}