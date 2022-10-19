// Copyright © 2019-2021 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Data type definitions file for service: Segments.
// Includes binding types of a structures and enumerations defined in the service.
// Shared by client-side stubs and server-side skeletons to ensure type
// compatibility.

package infra

import (
	"github.com/TaoZou1/vsphere-automation-sdk-go/runtime/bindings"
	"github.com/TaoZou1/vsphere-automation-sdk-go/runtime/data"
	"github.com/TaoZou1/vsphere-automation-sdk-go/runtime/protocol"
	"github.com/TaoZou1/vsphere-automation-sdk-go/services/nsxt/model"
	"reflect"
)

// Possible value for ``segmentType`` of method Segments#list.
const Segments_LIST_SEGMENT_TYPE_DVPORTGROUP = "DVPortgroup"

// Possible value for ``segmentType`` of method Segments#list.
const Segments_LIST_SEGMENT_TYPE_ALL = "ALL"

func segmentsDeleteInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["segment_id"] = bindings.NewStringType()
	fields["org_id"] = bindings.NewStringType()
	fields["project_id"] = bindings.NewStringType()
	fieldNameMap["segment_id"] = "SegmentId"
	fieldNameMap["org_id"] = "OrgId"
	fieldNameMap["project_id"] = "ProjectId"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func segmentsDeleteOutputType() bindings.BindingType {
	return bindings.NewVoidType()
}

func segmentsDeleteRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["segment_id"] = bindings.NewStringType()
	fields["org_id"] = bindings.NewStringType()
	fields["project_id"] = bindings.NewStringType()
	fieldNameMap["segment_id"] = "SegmentId"
	fieldNameMap["org_id"] = "OrgId"
	fieldNameMap["project_id"] = "ProjectId"
	paramsTypeMap["project_id"] = bindings.NewStringType()
	paramsTypeMap["org_id"] = bindings.NewStringType()
	paramsTypeMap["segment_id"] = bindings.NewStringType()
	paramsTypeMap["orgId"] = bindings.NewStringType()
	paramsTypeMap["projectId"] = bindings.NewStringType()
	paramsTypeMap["segmentId"] = bindings.NewStringType()
	pathParams["segment_id"] = "segmentId"
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
		"/policy/api/v1/orgs/{orgId}/projects/{projectId}/infra/segments/{segmentId}",
		"",
		resultHeaders,
		204,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func segmentsDelete0InputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["org_id"] = bindings.NewStringType()
	fields["project_id"] = bindings.NewStringType()
	fields["segment_id"] = bindings.NewStringType()
	fields["cascade"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fieldNameMap["org_id"] = "OrgId"
	fieldNameMap["project_id"] = "ProjectId"
	fieldNameMap["segment_id"] = "SegmentId"
	fieldNameMap["cascade"] = "Cascade"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func segmentsDelete0OutputType() bindings.BindingType {
	return bindings.NewVoidType()
}

func segmentsDelete0RestMetadata() protocol.OperationRestMetadata {
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
	fields["segment_id"] = bindings.NewStringType()
	fields["cascade"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fieldNameMap["org_id"] = "OrgId"
	fieldNameMap["project_id"] = "ProjectId"
	fieldNameMap["segment_id"] = "SegmentId"
	fieldNameMap["cascade"] = "Cascade"
	paramsTypeMap["project_id"] = bindings.NewStringType()
	paramsTypeMap["org_id"] = bindings.NewStringType()
	paramsTypeMap["cascade"] = bindings.NewOptionalType(bindings.NewBooleanType())
	paramsTypeMap["segment_id"] = bindings.NewStringType()
	paramsTypeMap["orgId"] = bindings.NewStringType()
	paramsTypeMap["projectId"] = bindings.NewStringType()
	paramsTypeMap["segmentId"] = bindings.NewStringType()
	pathParams["segment_id"] = "segmentId"
	pathParams["project_id"] = "projectId"
	pathParams["org_id"] = "orgId"
	queryParams["cascade"] = "cascade"
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
		"force=true",
		"",
		"DELETE",
		"/policy/api/v1/orgs/{orgId}/projects/{projectId}/infra/segments/{segmentId}",
		"",
		resultHeaders,
		204,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func segmentsDeletedhcpleasesInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["org_id"] = bindings.NewStringType()
	fields["project_id"] = bindings.NewStringType()
	fields["segment_id"] = bindings.NewStringType()
	fields["dhcp_delete_leases"] = bindings.NewReferenceType(model.DhcpDeleteLeasesBindingType)
	fields["enforcement_point_path"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["org_id"] = "OrgId"
	fieldNameMap["project_id"] = "ProjectId"
	fieldNameMap["segment_id"] = "SegmentId"
	fieldNameMap["dhcp_delete_leases"] = "DhcpDeleteLeases"
	fieldNameMap["enforcement_point_path"] = "EnforcementPointPath"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func segmentsDeletedhcpleasesOutputType() bindings.BindingType {
	return bindings.NewVoidType()
}

func segmentsDeletedhcpleasesRestMetadata() protocol.OperationRestMetadata {
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
	fields["segment_id"] = bindings.NewStringType()
	fields["dhcp_delete_leases"] = bindings.NewReferenceType(model.DhcpDeleteLeasesBindingType)
	fields["enforcement_point_path"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["org_id"] = "OrgId"
	fieldNameMap["project_id"] = "ProjectId"
	fieldNameMap["segment_id"] = "SegmentId"
	fieldNameMap["dhcp_delete_leases"] = "DhcpDeleteLeases"
	fieldNameMap["enforcement_point_path"] = "EnforcementPointPath"
	paramsTypeMap["enforcement_point_path"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["project_id"] = bindings.NewStringType()
	paramsTypeMap["org_id"] = bindings.NewStringType()
	paramsTypeMap["dhcp_delete_leases"] = bindings.NewReferenceType(model.DhcpDeleteLeasesBindingType)
	paramsTypeMap["segment_id"] = bindings.NewStringType()
	paramsTypeMap["orgId"] = bindings.NewStringType()
	paramsTypeMap["projectId"] = bindings.NewStringType()
	paramsTypeMap["segmentId"] = bindings.NewStringType()
	pathParams["segment_id"] = "segmentId"
	pathParams["project_id"] = "projectId"
	pathParams["org_id"] = "orgId"
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
		"action=delete_dhcp_leases",
		"dhcp_delete_leases",
		"POST",
		"/policy/api/v1/orgs/{orgId}/projects/{projectId}/infra/segments/{segmentId}",
		"",
		resultHeaders,
		204,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func segmentsGetInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["segment_id"] = bindings.NewStringType()
	fields["org_id"] = bindings.NewStringType()
	fields["project_id"] = bindings.NewStringType()
	fieldNameMap["segment_id"] = "SegmentId"
	fieldNameMap["org_id"] = "OrgId"
	fieldNameMap["project_id"] = "ProjectId"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func segmentsGetOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.SegmentBindingType)
}

func segmentsGetRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["segment_id"] = bindings.NewStringType()
	fields["org_id"] = bindings.NewStringType()
	fields["project_id"] = bindings.NewStringType()
	fieldNameMap["segment_id"] = "SegmentId"
	fieldNameMap["org_id"] = "OrgId"
	fieldNameMap["project_id"] = "ProjectId"
	paramsTypeMap["project_id"] = bindings.NewStringType()
	paramsTypeMap["org_id"] = bindings.NewStringType()
	paramsTypeMap["segment_id"] = bindings.NewStringType()
	paramsTypeMap["orgId"] = bindings.NewStringType()
	paramsTypeMap["projectId"] = bindings.NewStringType()
	paramsTypeMap["segmentId"] = bindings.NewStringType()
	pathParams["segment_id"] = "segmentId"
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
		"/policy/api/v1/orgs/{orgId}/projects/{projectId}/infra/segments/{segmentId}",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func segmentsListInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["org_id"] = bindings.NewStringType()
	fields["project_id"] = bindings.NewStringType()
	fields["cursor"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["include_mark_for_delete_objects"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fields["included_fields"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["page_size"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fields["segment_type"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["sort_ascending"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fields["sort_by"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["org_id"] = "OrgId"
	fieldNameMap["project_id"] = "ProjectId"
	fieldNameMap["cursor"] = "Cursor"
	fieldNameMap["include_mark_for_delete_objects"] = "IncludeMarkForDeleteObjects"
	fieldNameMap["included_fields"] = "IncludedFields"
	fieldNameMap["page_size"] = "PageSize"
	fieldNameMap["segment_type"] = "SegmentType"
	fieldNameMap["sort_ascending"] = "SortAscending"
	fieldNameMap["sort_by"] = "SortBy"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func segmentsListOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.SegmentListResultBindingType)
}

func segmentsListRestMetadata() protocol.OperationRestMetadata {
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
	fields["cursor"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["include_mark_for_delete_objects"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fields["included_fields"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["page_size"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fields["segment_type"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["sort_ascending"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fields["sort_by"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["org_id"] = "OrgId"
	fieldNameMap["project_id"] = "ProjectId"
	fieldNameMap["cursor"] = "Cursor"
	fieldNameMap["include_mark_for_delete_objects"] = "IncludeMarkForDeleteObjects"
	fieldNameMap["included_fields"] = "IncludedFields"
	fieldNameMap["page_size"] = "PageSize"
	fieldNameMap["segment_type"] = "SegmentType"
	fieldNameMap["sort_ascending"] = "SortAscending"
	fieldNameMap["sort_by"] = "SortBy"
	paramsTypeMap["segment_type"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["project_id"] = bindings.NewStringType()
	paramsTypeMap["org_id"] = bindings.NewStringType()
	paramsTypeMap["included_fields"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["page_size"] = bindings.NewOptionalType(bindings.NewIntegerType())
	paramsTypeMap["include_mark_for_delete_objects"] = bindings.NewOptionalType(bindings.NewBooleanType())
	paramsTypeMap["cursor"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["sort_by"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["sort_ascending"] = bindings.NewOptionalType(bindings.NewBooleanType())
	paramsTypeMap["orgId"] = bindings.NewStringType()
	paramsTypeMap["projectId"] = bindings.NewStringType()
	pathParams["project_id"] = "projectId"
	pathParams["org_id"] = "orgId"
	queryParams["cursor"] = "cursor"
	queryParams["sort_ascending"] = "sort_ascending"
	queryParams["included_fields"] = "included_fields"
	queryParams["segment_type"] = "segment_type"
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
		"/policy/api/v1/orgs/{orgId}/projects/{projectId}/infra/segments",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func segmentsPatchInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["segment_id"] = bindings.NewStringType()
	fields["org_id"] = bindings.NewStringType()
	fields["project_id"] = bindings.NewStringType()
	fields["segment"] = bindings.NewReferenceType(model.SegmentBindingType)
	fieldNameMap["segment_id"] = "SegmentId"
	fieldNameMap["org_id"] = "OrgId"
	fieldNameMap["project_id"] = "ProjectId"
	fieldNameMap["segment"] = "Segment"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func segmentsPatchOutputType() bindings.BindingType {
	return bindings.NewVoidType()
}

func segmentsPatchRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["segment_id"] = bindings.NewStringType()
	fields["org_id"] = bindings.NewStringType()
	fields["project_id"] = bindings.NewStringType()
	fields["segment"] = bindings.NewReferenceType(model.SegmentBindingType)
	fieldNameMap["segment_id"] = "SegmentId"
	fieldNameMap["org_id"] = "OrgId"
	fieldNameMap["project_id"] = "ProjectId"
	fieldNameMap["segment"] = "Segment"
	paramsTypeMap["project_id"] = bindings.NewStringType()
	paramsTypeMap["org_id"] = bindings.NewStringType()
	paramsTypeMap["segment_id"] = bindings.NewStringType()
	paramsTypeMap["segment"] = bindings.NewReferenceType(model.SegmentBindingType)
	paramsTypeMap["orgId"] = bindings.NewStringType()
	paramsTypeMap["projectId"] = bindings.NewStringType()
	paramsTypeMap["segmentId"] = bindings.NewStringType()
	pathParams["segment_id"] = "segmentId"
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
		"segment",
		"PATCH",
		"/policy/api/v1/orgs/{orgId}/projects/{projectId}/infra/segments/{segmentId}",
		"",
		resultHeaders,
		204,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func segmentsPatch0InputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["segment_id"] = bindings.NewStringType()
	fields["org_id"] = bindings.NewStringType()
	fields["project_id"] = bindings.NewStringType()
	fields["segment"] = bindings.NewReferenceType(model.SegmentBindingType)
	fieldNameMap["segment_id"] = "SegmentId"
	fieldNameMap["org_id"] = "OrgId"
	fieldNameMap["project_id"] = "ProjectId"
	fieldNameMap["segment"] = "Segment"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func segmentsPatch0OutputType() bindings.BindingType {
	return bindings.NewVoidType()
}

func segmentsPatch0RestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["segment_id"] = bindings.NewStringType()
	fields["org_id"] = bindings.NewStringType()
	fields["project_id"] = bindings.NewStringType()
	fields["segment"] = bindings.NewReferenceType(model.SegmentBindingType)
	fieldNameMap["segment_id"] = "SegmentId"
	fieldNameMap["org_id"] = "OrgId"
	fieldNameMap["project_id"] = "ProjectId"
	fieldNameMap["segment"] = "Segment"
	paramsTypeMap["project_id"] = bindings.NewStringType()
	paramsTypeMap["org_id"] = bindings.NewStringType()
	paramsTypeMap["segment_id"] = bindings.NewStringType()
	paramsTypeMap["segment"] = bindings.NewReferenceType(model.SegmentBindingType)
	paramsTypeMap["orgId"] = bindings.NewStringType()
	paramsTypeMap["projectId"] = bindings.NewStringType()
	paramsTypeMap["segmentId"] = bindings.NewStringType()
	pathParams["segment_id"] = "segmentId"
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
		"force=true",
		"segment",
		"PATCH",
		"/policy/api/v1/orgs/{orgId}/projects/{projectId}/infra/segments/{segmentId}",
		"",
		resultHeaders,
		204,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func segmentsUpdateInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["segment_id"] = bindings.NewStringType()
	fields["org_id"] = bindings.NewStringType()
	fields["project_id"] = bindings.NewStringType()
	fields["segment"] = bindings.NewReferenceType(model.SegmentBindingType)
	fieldNameMap["segment_id"] = "SegmentId"
	fieldNameMap["org_id"] = "OrgId"
	fieldNameMap["project_id"] = "ProjectId"
	fieldNameMap["segment"] = "Segment"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func segmentsUpdateOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.SegmentBindingType)
}

func segmentsUpdateRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["segment_id"] = bindings.NewStringType()
	fields["org_id"] = bindings.NewStringType()
	fields["project_id"] = bindings.NewStringType()
	fields["segment"] = bindings.NewReferenceType(model.SegmentBindingType)
	fieldNameMap["segment_id"] = "SegmentId"
	fieldNameMap["org_id"] = "OrgId"
	fieldNameMap["project_id"] = "ProjectId"
	fieldNameMap["segment"] = "Segment"
	paramsTypeMap["project_id"] = bindings.NewStringType()
	paramsTypeMap["org_id"] = bindings.NewStringType()
	paramsTypeMap["segment_id"] = bindings.NewStringType()
	paramsTypeMap["segment"] = bindings.NewReferenceType(model.SegmentBindingType)
	paramsTypeMap["orgId"] = bindings.NewStringType()
	paramsTypeMap["projectId"] = bindings.NewStringType()
	paramsTypeMap["segmentId"] = bindings.NewStringType()
	pathParams["segment_id"] = "segmentId"
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
		"segment",
		"PUT",
		"/policy/api/v1/orgs/{orgId}/projects/{projectId}/infra/segments/{segmentId}",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func segmentsUpdate0InputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["segment_id"] = bindings.NewStringType()
	fields["org_id"] = bindings.NewStringType()
	fields["project_id"] = bindings.NewStringType()
	fields["segment"] = bindings.NewReferenceType(model.SegmentBindingType)
	fieldNameMap["segment_id"] = "SegmentId"
	fieldNameMap["org_id"] = "OrgId"
	fieldNameMap["project_id"] = "ProjectId"
	fieldNameMap["segment"] = "Segment"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func segmentsUpdate0OutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.SegmentBindingType)
}

func segmentsUpdate0RestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["segment_id"] = bindings.NewStringType()
	fields["org_id"] = bindings.NewStringType()
	fields["project_id"] = bindings.NewStringType()
	fields["segment"] = bindings.NewReferenceType(model.SegmentBindingType)
	fieldNameMap["segment_id"] = "SegmentId"
	fieldNameMap["org_id"] = "OrgId"
	fieldNameMap["project_id"] = "ProjectId"
	fieldNameMap["segment"] = "Segment"
	paramsTypeMap["project_id"] = bindings.NewStringType()
	paramsTypeMap["org_id"] = bindings.NewStringType()
	paramsTypeMap["segment_id"] = bindings.NewStringType()
	paramsTypeMap["segment"] = bindings.NewReferenceType(model.SegmentBindingType)
	paramsTypeMap["orgId"] = bindings.NewStringType()
	paramsTypeMap["projectId"] = bindings.NewStringType()
	paramsTypeMap["segmentId"] = bindings.NewStringType()
	pathParams["segment_id"] = "segmentId"
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
		"force=true",
		"segment",
		"PUT",
		"/policy/api/v1/orgs/{orgId}/projects/{projectId}/infra/segments/{segmentId}",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}
