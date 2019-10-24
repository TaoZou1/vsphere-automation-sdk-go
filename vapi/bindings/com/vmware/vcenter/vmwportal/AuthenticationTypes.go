/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Data type definitions file for service: Authentication.
 * Includes binding types of a structures and enumerations defined in the service.
 * Shared by client-side stubs and server-side skeletons to ensure type
 * compatibility.
 */
package vmwportal

import (
	"reflect"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/bindings"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/data"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/protocol"
)

// The ``CreateSpec`` class contains the basic authentication credentials for my.vmware.com portal. **Warning:** This class is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
type AuthenticationCreateSpec struct {
    // my.vmware.com username. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
	Username string
    // my.vmware.com password. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
	Password string
}


func authenticationCreateInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["spec"] = bindings.NewReferenceType(AuthenticationCreateSpecBindingType)
	fieldNameMap["spec"] = "Spec"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func authenticationCreateOutputType() bindings.BindingType {
	return bindings.NewVoidType()
}

func authenticationCreateRestMetadata() protocol.OperationRestMetadata {
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	paramsTypeMap["spec"] = bindings.NewReferenceType(AuthenticationCreateSpecBindingType)
	resultHeaders := map[string]string{}
	errorHeaders := map[string]string{}
	errorHeaders["Unauthenticated.challenge"] = "WWW-Authenticate"
	return protocol.NewOperationRestMetadata(
		paramsTypeMap,
		pathParams,
		queryParams,
		headerParams,
		"spec",
		"POST",
		"/vcenter/vmwportal/authentication",
		resultHeaders,
		204,
		errorHeaders,
		map[string]int{"Unauthenticated": 401,"Unauthorized": 403,"Error": 500})
}


func AuthenticationCreateSpecBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["username"] = bindings.NewStringType()
	fieldNameMap["username"] = "Username"
	fields["password"] = bindings.NewSecretType()
	fieldNameMap["password"] = "Password"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vcenter.vmwportal.authentication.create_spec", fields, reflect.TypeOf(AuthenticationCreateSpec{}), fieldNameMap, validators)
}

