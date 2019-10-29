/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Code generated. DO NOT EDIT.

/*
 * Data type definitions file for service: Credential.
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





func credentialSetInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["provider"] = bindings.NewIdType([]string{"com.vmware.esx.kms.providers"}, "")
	fields["credential"] = bindings.NewSecretType()
	fieldNameMap["provider"] = "Provider"
	fieldNameMap["credential"] = "Credential"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func credentialSetOutputType() bindings.BindingType {
	return bindings.NewVoidType()
}

func credentialSetRestMetadata() protocol.OperationRestMetadata {
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	paramsTypeMap["provider"] = bindings.NewIdType([]string{"com.vmware.esx.kms.providers"}, "")
	paramsTypeMap["credential"] = bindings.NewSecretType()
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
		"credential",
		"PUT",
		"/esx/kms/providers/{provider}/credential",
		resultHeaders,
		204,
		errorHeaders,
		map[string]int{"InvalidArgument": 400,"NotFound": 404,"Unauthenticated": 401,"Unauthorized": 403,"Error": 500})
}


