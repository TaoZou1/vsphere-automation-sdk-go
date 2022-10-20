// Copyright © 2019-2021 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Data type definitions file for service: EdgeClusters.
// Includes binding types of a structures and enumerations defined in the service.
// Shared by client-side stubs and server-side skeletons to ensure type
// compatibility.

package nsx

import (
	"github.com/TaoZou1/vsphere-automation-sdk-go/runtime/bindings"
	"github.com/TaoZou1/vsphere-automation-sdk-go/runtime/data"
	"github.com/TaoZou1/vsphere-automation-sdk-go/runtime/protocol"
	"github.com/TaoZou1/vsphere-automation-sdk-go/services/nsxt-mp/nsx/model"
	"reflect"
)

func edgeClustersCreateInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["edge_cluster"] = bindings.NewReferenceType(model.EdgeClusterBindingType)
	fieldNameMap["edge_cluster"] = "EdgeCluster"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func edgeClustersCreateOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.EdgeClusterBindingType)
}

func edgeClustersCreateRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["edge_cluster"] = bindings.NewReferenceType(model.EdgeClusterBindingType)
	fieldNameMap["edge_cluster"] = "EdgeCluster"
	paramsTypeMap["edge_cluster"] = bindings.NewReferenceType(model.EdgeClusterBindingType)
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
		"edge_cluster",
		"POST",
		"/api/v1/edge-clusters",
		"",
		resultHeaders,
		201,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func edgeClustersDeleteInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["edge_cluster_id"] = bindings.NewStringType()
	fieldNameMap["edge_cluster_id"] = "EdgeClusterId"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func edgeClustersDeleteOutputType() bindings.BindingType {
	return bindings.NewVoidType()
}

func edgeClustersDeleteRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["edge_cluster_id"] = bindings.NewStringType()
	fieldNameMap["edge_cluster_id"] = "EdgeClusterId"
	paramsTypeMap["edge_cluster_id"] = bindings.NewStringType()
	paramsTypeMap["edgeClusterId"] = bindings.NewStringType()
	pathParams["edge_cluster_id"] = "edgeClusterId"
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
		"/api/v1/edge-clusters/{edgeClusterId}",
		"",
		resultHeaders,
		204,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func edgeClustersGetInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["edge_cluster_id"] = bindings.NewStringType()
	fieldNameMap["edge_cluster_id"] = "EdgeClusterId"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func edgeClustersGetOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.EdgeClusterBindingType)
}

func edgeClustersGetRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["edge_cluster_id"] = bindings.NewStringType()
	fieldNameMap["edge_cluster_id"] = "EdgeClusterId"
	paramsTypeMap["edge_cluster_id"] = bindings.NewStringType()
	paramsTypeMap["edgeClusterId"] = bindings.NewStringType()
	pathParams["edge_cluster_id"] = "edgeClusterId"
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
		"/api/v1/edge-clusters/{edgeClusterId}",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func edgeClustersListInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["cursor"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["included_fields"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["page_size"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fields["sort_ascending"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fields["sort_by"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["cursor"] = "Cursor"
	fieldNameMap["included_fields"] = "IncludedFields"
	fieldNameMap["page_size"] = "PageSize"
	fieldNameMap["sort_ascending"] = "SortAscending"
	fieldNameMap["sort_by"] = "SortBy"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func edgeClustersListOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.EdgeClusterListResultBindingType)
}

func edgeClustersListRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["cursor"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["included_fields"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["page_size"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fields["sort_ascending"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fields["sort_by"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["cursor"] = "Cursor"
	fieldNameMap["included_fields"] = "IncludedFields"
	fieldNameMap["page_size"] = "PageSize"
	fieldNameMap["sort_ascending"] = "SortAscending"
	fieldNameMap["sort_by"] = "SortBy"
	paramsTypeMap["included_fields"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["page_size"] = bindings.NewOptionalType(bindings.NewIntegerType())
	paramsTypeMap["cursor"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["sort_by"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["sort_ascending"] = bindings.NewOptionalType(bindings.NewBooleanType())
	queryParams["cursor"] = "cursor"
	queryParams["sort_ascending"] = "sort_ascending"
	queryParams["included_fields"] = "included_fields"
	queryParams["sort_by"] = "sort_by"
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
		"/api/v1/edge-clusters",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func edgeClustersRelocateandremoveInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["edge_cluster_id"] = bindings.NewStringType()
	fields["edge_cluster_member_index"] = bindings.NewReferenceType(model.EdgeClusterMemberIndexBindingType)
	fieldNameMap["edge_cluster_id"] = "EdgeClusterId"
	fieldNameMap["edge_cluster_member_index"] = "EdgeClusterMemberIndex"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func edgeClustersRelocateandremoveOutputType() bindings.BindingType {
	return bindings.NewVoidType()
}

func edgeClustersRelocateandremoveRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["edge_cluster_id"] = bindings.NewStringType()
	fields["edge_cluster_member_index"] = bindings.NewReferenceType(model.EdgeClusterMemberIndexBindingType)
	fieldNameMap["edge_cluster_id"] = "EdgeClusterId"
	fieldNameMap["edge_cluster_member_index"] = "EdgeClusterMemberIndex"
	paramsTypeMap["edge_cluster_id"] = bindings.NewStringType()
	paramsTypeMap["edge_cluster_member_index"] = bindings.NewReferenceType(model.EdgeClusterMemberIndexBindingType)
	paramsTypeMap["edgeClusterId"] = bindings.NewStringType()
	pathParams["edge_cluster_id"] = "edgeClusterId"
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
		"action=relocate_and_remove",
		"edge_cluster_member_index",
		"POST",
		"/api/v1/edge-clusters/{edgeClusterId}",
		"",
		resultHeaders,
		204,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func edgeClustersReplacetransportnodeInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["edge_cluster_id"] = bindings.NewStringType()
	fields["edge_cluster_member_transport_node"] = bindings.NewReferenceType(model.EdgeClusterMemberTransportNodeBindingType)
	fieldNameMap["edge_cluster_id"] = "EdgeClusterId"
	fieldNameMap["edge_cluster_member_transport_node"] = "EdgeClusterMemberTransportNode"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func edgeClustersReplacetransportnodeOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.EdgeClusterBindingType)
}

func edgeClustersReplacetransportnodeRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["edge_cluster_id"] = bindings.NewStringType()
	fields["edge_cluster_member_transport_node"] = bindings.NewReferenceType(model.EdgeClusterMemberTransportNodeBindingType)
	fieldNameMap["edge_cluster_id"] = "EdgeClusterId"
	fieldNameMap["edge_cluster_member_transport_node"] = "EdgeClusterMemberTransportNode"
	paramsTypeMap["edge_cluster_id"] = bindings.NewStringType()
	paramsTypeMap["edge_cluster_member_transport_node"] = bindings.NewReferenceType(model.EdgeClusterMemberTransportNodeBindingType)
	paramsTypeMap["edgeClusterId"] = bindings.NewStringType()
	pathParams["edge_cluster_id"] = "edgeClusterId"
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
		"action=replace_transport_node",
		"edge_cluster_member_transport_node",
		"POST",
		"/api/v1/edge-clusters/{edgeClusterId}",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func edgeClustersUpdateInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["edge_cluster_id"] = bindings.NewStringType()
	fields["edge_cluster"] = bindings.NewReferenceType(model.EdgeClusterBindingType)
	fieldNameMap["edge_cluster_id"] = "EdgeClusterId"
	fieldNameMap["edge_cluster"] = "EdgeCluster"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func edgeClustersUpdateOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.EdgeClusterBindingType)
}

func edgeClustersUpdateRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["edge_cluster_id"] = bindings.NewStringType()
	fields["edge_cluster"] = bindings.NewReferenceType(model.EdgeClusterBindingType)
	fieldNameMap["edge_cluster_id"] = "EdgeClusterId"
	fieldNameMap["edge_cluster"] = "EdgeCluster"
	paramsTypeMap["edge_cluster_id"] = bindings.NewStringType()
	paramsTypeMap["edge_cluster"] = bindings.NewReferenceType(model.EdgeClusterBindingType)
	paramsTypeMap["edgeClusterId"] = bindings.NewStringType()
	pathParams["edge_cluster_id"] = "edgeClusterId"
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
		"edge_cluster",
		"PUT",
		"/api/v1/edge-clusters/{edgeClusterId}",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}
