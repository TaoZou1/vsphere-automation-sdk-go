/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Data type definitions file for service: ProductCatalog.
 * Includes binding types of a structures and enumerations defined in the service.
 * Shared by client-side stubs and server-side skeletons to ensure type
 * compatibility.
 */
package discovery

import (
	"reflect"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/bindings"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/data"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/protocol"
)

// The ``Summary`` class contains information about each VMware product and its corresponding versions that can be associated with vCenter Server.
type ProductCatalogSummary struct {
    // A product name that the customer is aware of.
	Name string
    // List of versions the customer can select from.
	Versions []string
}


func productCatalogListInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func productCatalogListOutputType() bindings.BindingType {
	return bindings.NewListType(bindings.NewReferenceType(ProductCatalogSummaryBindingType), reflect.TypeOf([]ProductCatalogSummary{}))
}

func productCatalogListRestMetadata() protocol.OperationRestMetadata {
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
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
		"/vcenter/lcm/discovery/product-catalog",
		resultHeaders,
		200,
		errorHeaders,
		map[string]int{"Unauthenticated": 401,"Error": 500})
}


func ProductCatalogSummaryBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["name"] = bindings.NewStringType()
	fieldNameMap["name"] = "Name"
	fields["versions"] = bindings.NewListType(bindings.NewStringType(), reflect.TypeOf([]string{}))
	fieldNameMap["versions"] = "Versions"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vcenter.lcm.discovery.product_catalog.summary", fields, reflect.TypeOf(ProductCatalogSummary{}), fieldNameMap, validators)
}

