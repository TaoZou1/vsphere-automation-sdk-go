/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Code generated. DO NOT EDIT.

/*
 * Data type definitions file for service: Reservation.
 * Includes binding types of a structures and enumerations defined in the service.
 * Shared by client-side stubs and server-side skeletons to ensure type
 * compatibility.
 */

package tbrs

import (
	"reflect"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/services/vmc/model"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/bindings"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/data"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/protocol"
)





func reservationPostInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["org"] = bindings.NewStringType()
	fields["sddc_state"] = bindings.NewOptionalType(bindings.NewReferenceType(model.SddcStateRequestBindingType))
	fieldNameMap["org"] = "Org"
	fieldNameMap["sddc_state"] = "SddcState"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func reservationPostOutputType() bindings.BindingType {
	return bindings.NewListType(bindings.NewReferenceType(model.ReservationWindowBindingType), reflect.TypeOf([]model.ReservationWindow{}))
}

func reservationPostRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	fields["org"] = bindings.NewStringType()
	fields["sddc_state"] = bindings.NewOptionalType(bindings.NewReferenceType(model.SddcStateRequestBindingType))
	fieldNameMap["org"] = "Org"
	fieldNameMap["sddc_state"] = "SddcState"
	paramsTypeMap["org"] = bindings.NewStringType()
	paramsTypeMap["sddc_state"] = bindings.NewOptionalType(bindings.NewReferenceType(model.SddcStateRequestBindingType))
	paramsTypeMap["org"] = bindings.NewStringType()
	pathParams["org"] = "org"
	resultHeaders := map[string]string{}
	errorHeaders := map[string]string{}
	errorHeaders["Unauthenticated.challenge"] = "WWW-Authenticate"
	return protocol.NewOperationRestMetadata(
		fields,
		fieldNameMap,
		paramsTypeMap,
		pathParams,
		queryParams,
		headerParams,
		"",
		"sddc_state",
		"POST",
		"/vmc/api/orgs/{org}/tbrs/reservation",
		resultHeaders,
		200,
		errorHeaders,
		map[string]int{"Unauthenticated": 401,"InvalidRequest": 400,"Unauthorized": 403})
}

