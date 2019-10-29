/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Code generated. DO NOT EDIT.

/*
 * Data type definitions file for service: ServerCertificates.
 * Includes binding types of a structures and enumerations defined in the service.
 * Shared by client-side stubs and server-side skeletons to ensure type
 * compatibility.
 */

package providers

import (
	"reflect"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/bindings"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/data"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/protocol"
)



// The ``Info`` class contains information for trusted key management server certificates.
type ServerCertificatesInfo struct {
    // Certificates. 
    //
    //  PEM format.
	Certificates []string
}



func serverCertificatesSetInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["provider"] = bindings.NewIdType([]string{"com.vmware.esx.kms.providers"}, "")
	fields["certificates"] = bindings.NewListType(bindings.NewStringType(), reflect.TypeOf([]string{}))
	fieldNameMap["provider"] = "Provider"
	fieldNameMap["certificates"] = "Certificates"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func serverCertificatesSetOutputType() bindings.BindingType {
	return bindings.NewVoidType()
}

func serverCertificatesSetRestMetadata() protocol.OperationRestMetadata {
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	paramsTypeMap["provider"] = bindings.NewIdType([]string{"com.vmware.esx.kms.providers"}, "")
	paramsTypeMap["certificates"] = bindings.NewListType(bindings.NewStringType(), reflect.TypeOf([]string{}))
	paramsTypeMap["provider"] = bindings.NewIdType([]string{"com.vmware.esx.kms.providers"}, "")
	pathParams["provider"] = "provider"
	resultHeaders := map[string]string{}
	errorHeaders := map[string]string{}
	errorHeaders["Unauthenticated.challenge"] = "WWW-Authenticate"
	return protocol.NewOperationRestMetadata(
		paramsTypeMap,
		pathParams,
		queryParams,
		headerParams,
		"certificates",
		"PUT",
		"/esx/kms/providers/{provider}/server-certificates",
		resultHeaders,
		204,
		errorHeaders,
		map[string]int{"InvalidArgument": 400,"NotFound": 404,"Unauthenticated": 401,"Unauthorized": 403,"Unsupported": 400,"Error": 500})
}

func serverCertificatesGetInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["provider"] = bindings.NewIdType([]string{"com.vmware.esx.kms.providers"}, "")
	fieldNameMap["provider"] = "Provider"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func serverCertificatesGetOutputType() bindings.BindingType {
	return bindings.NewReferenceType(ServerCertificatesInfoBindingType)
}

func serverCertificatesGetRestMetadata() protocol.OperationRestMetadata {
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	paramsTypeMap["provider"] = bindings.NewIdType([]string{"com.vmware.esx.kms.providers"}, "")
	paramsTypeMap["provider"] = bindings.NewIdType([]string{"com.vmware.esx.kms.providers"}, "")
	pathParams["provider"] = "provider"
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
		"/esx/kms/providers/{provider}/server-certificates",
		resultHeaders,
		200,
		errorHeaders,
		map[string]int{"InvalidArgument": 400,"NotFound": 404,"Unauthenticated": 401,"Unauthorized": 403,"Error": 500})
}


func ServerCertificatesInfoBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["certificates"] = bindings.NewListType(bindings.NewStringType(), reflect.TypeOf([]string{}))
	fieldNameMap["certificates"] = "Certificates"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.esx.kms.providers.server_certificates.info", fields, reflect.TypeOf(ServerCertificatesInfo{}), fieldNameMap, validators)
}

