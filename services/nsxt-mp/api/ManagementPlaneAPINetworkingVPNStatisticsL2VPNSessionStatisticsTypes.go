// Copyright © 2019-2021 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Data type definitions file for service: ManagementPlaneAPINetworkingVPNStatisticsL2VPNSessionStatistics.
// Includes binding types of a structures and enumerations defined in the service.
// Shared by client-side stubs and server-side skeletons to ensure type
// compatibility.

package api

import (
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/runtime/bindings"
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/runtime/data"
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/runtime/protocol"
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/services/nsxt-mp/model"
	"reflect"
)

// Possible value for ``source`` of method ManagementPlaneAPINetworkingVPNStatisticsL2VPNSessionStatistics#getL2VPNSessionStatistics.
const ManagementPlaneAPINetworkingVPNStatisticsL2VPNSessionStatistics_GET_L2VPNSESSION_STATISTICS_SOURCE_REALTIME = "realtime"

// Possible value for ``source`` of method ManagementPlaneAPINetworkingVPNStatisticsL2VPNSessionStatistics#getL2VPNSessionStatistics.
const ManagementPlaneAPINetworkingVPNStatisticsL2VPNSessionStatistics_GET_L2VPNSESSION_STATISTICS_SOURCE_CACHED = "cached"

func managementPlaneAPINetworkingVPNStatisticsL2VPNSessionStatisticsGetL2VPNSessionStatisticsInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["session_id"] = bindings.NewStringType()
	fields["source"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["session_id"] = "SessionId"
	fieldNameMap["source"] = "Source"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func managementPlaneAPINetworkingVPNStatisticsL2VPNSessionStatisticsGetL2VPNSessionStatisticsOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.L2VPNSessionStatisticsBindingType)
}

func managementPlaneAPINetworkingVPNStatisticsL2VPNSessionStatisticsGetL2VPNSessionStatisticsRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["session_id"] = bindings.NewStringType()
	fields["source"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["session_id"] = "SessionId"
	fieldNameMap["source"] = "Source"
	paramsTypeMap["source"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["session_id"] = bindings.NewStringType()
	paramsTypeMap["sessionId"] = bindings.NewStringType()
	pathParams["session_id"] = "sessionId"
	queryParams["source"] = "source"
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
		"/api/v1/vpn/l2vpn/sessions/{sessionId}/statistics",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}
