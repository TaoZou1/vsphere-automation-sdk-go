/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Code generated. DO NOT EDIT.

/*
 * Data type definitions file for service: Accounts.
 * Includes binding types of a structures and enumerations defined in the service.
 * Shared by client-side stubs and server-side skeletons to ensure type
 * compatibility.
 */

package nsx

import (
	"reflect"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/bindings"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/data"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/protocol"
)



// The ``Spec`` class contains the basic authentication credentials for my.vmware.com portal. **Warning:** This class is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
type AccountsSpec struct {
    // my.vmware.com username. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
	Username string
    // my.vmware.com password. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
	Password string
}

// The ``AccountsInfo`` class contains the information about account. **Warning:** This class is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
type AccountsAccountsInfo struct {
    // Account ID. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
	AccountId string
    // Account Name. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
	AccountName string
}



func accountsAuthenticateInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["spec"] = bindings.NewReferenceType(AccountsSpecBindingType)
	fieldNameMap["spec"] = "Spec"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func accountsAuthenticateOutputType() bindings.BindingType {
	return bindings.NewVoidType()
}

func accountsAuthenticateRestMetadata() protocol.OperationRestMetadata {
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	paramsTypeMap["spec"] = bindings.NewReferenceType(AccountsSpecBindingType)
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
		"/vcenter/vmwportal",
		resultHeaders,
		204,
		errorHeaders,
		map[string]int{"Unauthenticated": 401,"Unauthorized": 403,"Error": 500})
}

func accountsGetAccountsInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func accountsGetAccountsOutputType() bindings.BindingType {
	return bindings.NewListType(bindings.NewReferenceType(AccountsAccountsInfoBindingType), reflect.TypeOf([]AccountsAccountsInfo{}))
}

func accountsGetAccountsRestMetadata() protocol.OperationRestMetadata {
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
		"/vcenter/vmwportal/accounts",
		resultHeaders,
		200,
		errorHeaders,
		map[string]int{"Unauthenticated": 401,"Unauthorized": 403,"Error": 500})
}

func accountsCheckInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["accounts_id"] = bindings.NewIdType([]string{"ACCOUNT"}, "")
	fieldNameMap["accounts_id"] = "AccountsId"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func accountsCheckOutputType() bindings.BindingType {
	return bindings.NewVoidType()
}

func accountsCheckRestMetadata() protocol.OperationRestMetadata {
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	paramsTypeMap["accounts_id"] = bindings.NewIdType([]string{"ACCOUNT"}, "")
	paramsTypeMap["accounts_id"] = bindings.NewIdType([]string{"ACCOUNT"}, "")
	pathParams["accounts_id"] = "accounts_id"
	resultHeaders := map[string]string{}
	errorHeaders := map[string]string{}
	errorHeaders["Unauthenticated.challenge"] = "WWW-Authenticate"
	return protocol.NewOperationRestMetadata(
		paramsTypeMap,
		pathParams,
		queryParams,
		headerParams,
		"",
		"POST",
		"/vcenter/vmwportal/accounts/{accounts_id}",
		resultHeaders,
		204,
		errorHeaders,
		map[string]int{"Unauthenticated": 401,"Unauthorized": 403,"Error": 500})
}


func AccountsSpecBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["username"] = bindings.NewStringType()
	fieldNameMap["username"] = "Username"
	fields["password"] = bindings.NewSecretType()
	fieldNameMap["password"] = "Password"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vcenter.nsx.accounts.spec", fields, reflect.TypeOf(AccountsSpec{}), fieldNameMap, validators)
}

func AccountsAccountsInfoBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["account_id"] = bindings.NewStringType()
	fieldNameMap["account_id"] = "AccountId"
	fields["account_name"] = bindings.NewStringType()
	fieldNameMap["account_name"] = "AccountName"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vcenter.nsx.accounts.accounts_info", fields, reflect.TypeOf(AccountsAccountsInfo{}), fieldNameMap, validators)
}

