/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Code generated. DO NOT EDIT.

/*
 * Data type definitions file for service: PeerConfig.
 * Includes binding types of a structures and enumerations defined in the service.
 * Shared by client-side stubs and server-side skeletons to ensure type
 * compatibility.
 */

package l2vpns

import (
	"reflect"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/services/nsxt/model"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/bindings"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/data"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/protocol"
)





func peerConfigGetInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["tier0_id"] = bindings.NewStringType()
	fields["locale_service_id"] = bindings.NewStringType()
	fields["l2vpn_id"] = bindings.NewStringType()
	fields["enforcement_point_path"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["tier0_id"] = "Tier0Id"
	fieldNameMap["locale_service_id"] = "LocaleServiceId"
	fieldNameMap["l2vpn_id"] = "L2vpnId"
	fieldNameMap["enforcement_point_path"] = "EnforcementPointPath"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func peerConfigGetOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.AggregateL2VpnPeerConfigBindingType)
}

func peerConfigGetRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	fields["tier0_id"] = bindings.NewStringType()
	fields["locale_service_id"] = bindings.NewStringType()
	fields["l2vpn_id"] = bindings.NewStringType()
	fields["enforcement_point_path"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["tier0_id"] = "Tier0Id"
	fieldNameMap["locale_service_id"] = "LocaleServiceId"
	fieldNameMap["l2vpn_id"] = "L2vpnId"
	fieldNameMap["enforcement_point_path"] = "EnforcementPointPath"
	paramsTypeMap["tier0_id"] = bindings.NewStringType()
	paramsTypeMap["enforcement_point_path"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["locale_service_id"] = bindings.NewStringType()
	paramsTypeMap["l2vpn_id"] = bindings.NewStringType()
	paramsTypeMap["tier0Id"] = bindings.NewStringType()
	paramsTypeMap["localeServiceId"] = bindings.NewStringType()
	paramsTypeMap["l2vpnId"] = bindings.NewStringType()
	pathParams["tier0_id"] = "tier0Id"
	pathParams["locale_service_id"] = "localeServiceId"
	pathParams["l2vpn_id"] = "l2vpnId"
	queryParams["enforcement_point_path"] = "enforcement_point_path"
	resultHeaders := map[string]string{}
	errorHeaders := map[string]string{}
	return protocol.NewOperationRestMetadata(
		fields,
		fieldNameMap,
		paramsTypeMap,
		pathParams,
		queryParams,
		headerParams,
		"",
		"",
		"GET",
		"/policy/api/v1/infra/tier-0s/{tier0Id}/locale-services/{localeServiceId}/l2vpn-context/l2vpns/{l2vpnId}/peer-config",
		resultHeaders,
		200,
		errorHeaders,
		map[string]int{"InvalidRequest": 400,"Unauthorized": 403,"ServiceUnavailable": 503,"InternalServerError": 500,"NotFound": 404})
}


