/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Data type definitions file for service: OfferInstances.
 * Includes binding types of a structures and enumerations defined in the service.
 * Shared by client-side stubs and server-side skeletons to ensure type
 * compatibility.
 */
package subscriptions

import (
	"reflect"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/bindings/com/vmware/vmc/model"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/bindings"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/data"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/protocol"
)


func offerInstancesListInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["org"] = bindings.NewStringType()
	fields["region"] = bindings.NewStringType()
	fields["product_type"] = bindings.NewStringType()
	fields["product"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["type"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["org"] = "Org"
	fieldNameMap["region"] = "Region"
	fieldNameMap["product_type"] = "ProductType"
	fieldNameMap["product"] = "Product"
	fieldNameMap["type"] = "Type_"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func offerInstancesListOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.OfferInstancesHolderBindingType)
}

func offerInstancesListRestMetadata() protocol.OperationRestMetadata {
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	paramsTypeMap["region"] = bindings.NewStringType()
	paramsTypeMap["org"] = bindings.NewStringType()
	paramsTypeMap["product_type"] = bindings.NewStringType()
	paramsTypeMap["product"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["type"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["org"] = bindings.NewStringType()
	pathParams["org"] = "org"
	queryParams["product"] = "product"
	queryParams["product_type"] = "product_type"
	queryParams["region"] = "region"
	queryParams["type"] = "type"
	resultHeaders := map[string]string{}
	errorHeaders := map[string]string{}
	errorHeaders["Unauthenticated.challenge"] = "WWW-Authenticate"
	return protocol.NewOperationRestMetadata(
		paramsTypeMap,
		pathParams,
		queryParams,
		headerParams,
		"",
		"GET",
		"/vmc/api/orgs/{org}/subscriptions/offer-instances",
		resultHeaders,
		200,
		errorHeaders,
		map[string]int{"Unauthenticated": 401,"InvalidRequest": 400,"Unauthorized": 403})
}


