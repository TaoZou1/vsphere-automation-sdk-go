/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Code generated. DO NOT EDIT.

/*
 * Data type definitions file for service: PortQosProfileBindingMaps.
 * Includes binding types of a structures and enumerations defined in the service.
 * Shared by client-side stubs and server-side skeletons to ensure type
 * compatibility.
 */

package ports

import (
	"reflect"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/services/nsxt/model"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/bindings"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/data"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/protocol"
)





func portQosProfileBindingMapsDeleteInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["segment_id"] = bindings.NewStringType()
	fields["port_id"] = bindings.NewStringType()
	fields["port_qos_profile_binding_map_id"] = bindings.NewStringType()
	fieldNameMap["segment_id"] = "SegmentId"
	fieldNameMap["port_id"] = "PortId"
	fieldNameMap["port_qos_profile_binding_map_id"] = "PortQosProfileBindingMapId"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func portQosProfileBindingMapsDeleteOutputType() bindings.BindingType {
	return bindings.NewVoidType()
}

func portQosProfileBindingMapsDeleteRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["segment_id"] = bindings.NewStringType()
	fields["port_id"] = bindings.NewStringType()
	fields["port_qos_profile_binding_map_id"] = bindings.NewStringType()
	fieldNameMap["segment_id"] = "SegmentId"
	fieldNameMap["port_id"] = "PortId"
	fieldNameMap["port_qos_profile_binding_map_id"] = "PortQosProfileBindingMapId"
	paramsTypeMap["port_qos_profile_binding_map_id"] = bindings.NewStringType()
	paramsTypeMap["port_id"] = bindings.NewStringType()
	paramsTypeMap["segment_id"] = bindings.NewStringType()
	paramsTypeMap["segmentId"] = bindings.NewStringType()
	paramsTypeMap["portId"] = bindings.NewStringType()
	paramsTypeMap["portQosProfileBindingMapId"] = bindings.NewStringType()
	pathParams["segment_id"] = "segmentId"
	pathParams["port_qos_profile_binding_map_id"] = "portQosProfileBindingMapId"
	pathParams["port_id"] = "portId"
	resultHeaders := map[string]string{}
	errorHeaders := map[string]string{}
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
		"/policy/api/v1/infra/segments/{segmentId}/ports/{portId}/port-qos-profile-binding-maps/{portQosProfileBindingMapId}",
		"",
		resultHeaders,
		204,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400,"com.vmware.vapi.std.errors.unauthorized": 403,"com.vmware.vapi.std.errors.service_unavailable": 503,"com.vmware.vapi.std.errors.internal_server_error": 500,"com.vmware.vapi.std.errors.not_found": 404})
}

func portQosProfileBindingMapsGetInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["segment_id"] = bindings.NewStringType()
	fields["port_id"] = bindings.NewStringType()
	fields["port_qos_profile_binding_map_id"] = bindings.NewStringType()
	fieldNameMap["segment_id"] = "SegmentId"
	fieldNameMap["port_id"] = "PortId"
	fieldNameMap["port_qos_profile_binding_map_id"] = "PortQosProfileBindingMapId"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func portQosProfileBindingMapsGetOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.PortQosProfileBindingMapBindingType)
}

func portQosProfileBindingMapsGetRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["segment_id"] = bindings.NewStringType()
	fields["port_id"] = bindings.NewStringType()
	fields["port_qos_profile_binding_map_id"] = bindings.NewStringType()
	fieldNameMap["segment_id"] = "SegmentId"
	fieldNameMap["port_id"] = "PortId"
	fieldNameMap["port_qos_profile_binding_map_id"] = "PortQosProfileBindingMapId"
	paramsTypeMap["port_qos_profile_binding_map_id"] = bindings.NewStringType()
	paramsTypeMap["port_id"] = bindings.NewStringType()
	paramsTypeMap["segment_id"] = bindings.NewStringType()
	paramsTypeMap["segmentId"] = bindings.NewStringType()
	paramsTypeMap["portId"] = bindings.NewStringType()
	paramsTypeMap["portQosProfileBindingMapId"] = bindings.NewStringType()
	pathParams["segment_id"] = "segmentId"
	pathParams["port_qos_profile_binding_map_id"] = "portQosProfileBindingMapId"
	pathParams["port_id"] = "portId"
	resultHeaders := map[string]string{}
	errorHeaders := map[string]string{}
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
		"/policy/api/v1/infra/segments/{segmentId}/ports/{portId}/port-qos-profile-binding-maps/{portQosProfileBindingMapId}",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400,"com.vmware.vapi.std.errors.unauthorized": 403,"com.vmware.vapi.std.errors.service_unavailable": 503,"com.vmware.vapi.std.errors.internal_server_error": 500,"com.vmware.vapi.std.errors.not_found": 404})
}

func portQosProfileBindingMapsListInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["segment_id"] = bindings.NewStringType()
	fields["port_id"] = bindings.NewStringType()
	fields["cursor"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["included_fields"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["page_size"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fields["sort_ascending"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fields["sort_by"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["segment_id"] = "SegmentId"
	fieldNameMap["port_id"] = "PortId"
	fieldNameMap["cursor"] = "Cursor"
	fieldNameMap["included_fields"] = "IncludedFields"
	fieldNameMap["page_size"] = "PageSize"
	fieldNameMap["sort_ascending"] = "SortAscending"
	fieldNameMap["sort_by"] = "SortBy"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func portQosProfileBindingMapsListOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.PortQosProfileBindingMapListResultBindingType)
}

func portQosProfileBindingMapsListRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["segment_id"] = bindings.NewStringType()
	fields["port_id"] = bindings.NewStringType()
	fields["cursor"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["included_fields"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["page_size"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fields["sort_ascending"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fields["sort_by"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["segment_id"] = "SegmentId"
	fieldNameMap["port_id"] = "PortId"
	fieldNameMap["cursor"] = "Cursor"
	fieldNameMap["included_fields"] = "IncludedFields"
	fieldNameMap["page_size"] = "PageSize"
	fieldNameMap["sort_ascending"] = "SortAscending"
	fieldNameMap["sort_by"] = "SortBy"
	paramsTypeMap["included_fields"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["page_size"] = bindings.NewOptionalType(bindings.NewIntegerType())
	paramsTypeMap["port_id"] = bindings.NewStringType()
	paramsTypeMap["segment_id"] = bindings.NewStringType()
	paramsTypeMap["cursor"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["sort_by"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["sort_ascending"] = bindings.NewOptionalType(bindings.NewBooleanType())
	paramsTypeMap["segmentId"] = bindings.NewStringType()
	paramsTypeMap["portId"] = bindings.NewStringType()
	pathParams["segment_id"] = "segmentId"
	pathParams["port_id"] = "portId"
	queryParams["cursor"] = "cursor"
	queryParams["sort_ascending"] = "sort_ascending"
	queryParams["included_fields"] = "included_fields"
	queryParams["sort_by"] = "sort_by"
	queryParams["page_size"] = "page_size"
	resultHeaders := map[string]string{}
	errorHeaders := map[string]string{}
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
		"/policy/api/v1/infra/segments/{segmentId}/ports/{portId}/port-qos-profile-binding-maps",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400,"com.vmware.vapi.std.errors.unauthorized": 403,"com.vmware.vapi.std.errors.service_unavailable": 503,"com.vmware.vapi.std.errors.internal_server_error": 500,"com.vmware.vapi.std.errors.not_found": 404})
}

func portQosProfileBindingMapsPatchInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["segment_id"] = bindings.NewStringType()
	fields["port_id"] = bindings.NewStringType()
	fields["port_qos_profile_binding_map_id"] = bindings.NewStringType()
	fields["port_qos_profile_binding_map"] = bindings.NewReferenceType(model.PortQosProfileBindingMapBindingType)
	fieldNameMap["segment_id"] = "SegmentId"
	fieldNameMap["port_id"] = "PortId"
	fieldNameMap["port_qos_profile_binding_map_id"] = "PortQosProfileBindingMapId"
	fieldNameMap["port_qos_profile_binding_map"] = "PortQosProfileBindingMap"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func portQosProfileBindingMapsPatchOutputType() bindings.BindingType {
	return bindings.NewVoidType()
}

func portQosProfileBindingMapsPatchRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["segment_id"] = bindings.NewStringType()
	fields["port_id"] = bindings.NewStringType()
	fields["port_qos_profile_binding_map_id"] = bindings.NewStringType()
	fields["port_qos_profile_binding_map"] = bindings.NewReferenceType(model.PortQosProfileBindingMapBindingType)
	fieldNameMap["segment_id"] = "SegmentId"
	fieldNameMap["port_id"] = "PortId"
	fieldNameMap["port_qos_profile_binding_map_id"] = "PortQosProfileBindingMapId"
	fieldNameMap["port_qos_profile_binding_map"] = "PortQosProfileBindingMap"
	paramsTypeMap["port_qos_profile_binding_map_id"] = bindings.NewStringType()
	paramsTypeMap["port_id"] = bindings.NewStringType()
	paramsTypeMap["segment_id"] = bindings.NewStringType()
	paramsTypeMap["port_qos_profile_binding_map"] = bindings.NewReferenceType(model.PortQosProfileBindingMapBindingType)
	paramsTypeMap["segmentId"] = bindings.NewStringType()
	paramsTypeMap["portId"] = bindings.NewStringType()
	paramsTypeMap["portQosProfileBindingMapId"] = bindings.NewStringType()
	pathParams["segment_id"] = "segmentId"
	pathParams["port_qos_profile_binding_map_id"] = "portQosProfileBindingMapId"
	pathParams["port_id"] = "portId"
	resultHeaders := map[string]string{}
	errorHeaders := map[string]string{}
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
		"port_qos_profile_binding_map",
		"PATCH",
		"/policy/api/v1/infra/segments/{segmentId}/ports/{portId}/port-qos-profile-binding-maps/{portQosProfileBindingMapId}",
		"",
		resultHeaders,
		204,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400,"com.vmware.vapi.std.errors.unauthorized": 403,"com.vmware.vapi.std.errors.service_unavailable": 503,"com.vmware.vapi.std.errors.internal_server_error": 500,"com.vmware.vapi.std.errors.not_found": 404})
}

func portQosProfileBindingMapsUpdateInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["segment_id"] = bindings.NewStringType()
	fields["port_id"] = bindings.NewStringType()
	fields["port_qos_profile_binding_map_id"] = bindings.NewStringType()
	fields["port_qos_profile_binding_map"] = bindings.NewReferenceType(model.PortQosProfileBindingMapBindingType)
	fieldNameMap["segment_id"] = "SegmentId"
	fieldNameMap["port_id"] = "PortId"
	fieldNameMap["port_qos_profile_binding_map_id"] = "PortQosProfileBindingMapId"
	fieldNameMap["port_qos_profile_binding_map"] = "PortQosProfileBindingMap"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func portQosProfileBindingMapsUpdateOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.PortQosProfileBindingMapBindingType)
}

func portQosProfileBindingMapsUpdateRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["segment_id"] = bindings.NewStringType()
	fields["port_id"] = bindings.NewStringType()
	fields["port_qos_profile_binding_map_id"] = bindings.NewStringType()
	fields["port_qos_profile_binding_map"] = bindings.NewReferenceType(model.PortQosProfileBindingMapBindingType)
	fieldNameMap["segment_id"] = "SegmentId"
	fieldNameMap["port_id"] = "PortId"
	fieldNameMap["port_qos_profile_binding_map_id"] = "PortQosProfileBindingMapId"
	fieldNameMap["port_qos_profile_binding_map"] = "PortQosProfileBindingMap"
	paramsTypeMap["port_qos_profile_binding_map_id"] = bindings.NewStringType()
	paramsTypeMap["port_id"] = bindings.NewStringType()
	paramsTypeMap["segment_id"] = bindings.NewStringType()
	paramsTypeMap["port_qos_profile_binding_map"] = bindings.NewReferenceType(model.PortQosProfileBindingMapBindingType)
	paramsTypeMap["segmentId"] = bindings.NewStringType()
	paramsTypeMap["portId"] = bindings.NewStringType()
	paramsTypeMap["portQosProfileBindingMapId"] = bindings.NewStringType()
	pathParams["segment_id"] = "segmentId"
	pathParams["port_qos_profile_binding_map_id"] = "portQosProfileBindingMapId"
	pathParams["port_id"] = "portId"
	resultHeaders := map[string]string{}
	errorHeaders := map[string]string{}
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
		"port_qos_profile_binding_map",
		"PUT",
		"/policy/api/v1/infra/segments/{segmentId}/ports/{portId}/port-qos-profile-binding-maps/{portQosProfileBindingMapId}",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400,"com.vmware.vapi.std.errors.unauthorized": 403,"com.vmware.vapi.std.errors.service_unavailable": 503,"com.vmware.vapi.std.errors.internal_server_error": 500,"com.vmware.vapi.std.errors.not_found": 404})
}


