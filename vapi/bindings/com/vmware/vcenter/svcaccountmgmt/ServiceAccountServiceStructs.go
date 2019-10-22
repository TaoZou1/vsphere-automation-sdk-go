/* Copyright © 2019 VMware, Inc. All Rights Reserved.
     SPDX-License-Identifier: BSD-2-Clause */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Data type definitions file for service: ServiceAccount.
 * Includes binding types of a structures and enumerations defined in the service.
 * Shared by client-side stubs and server-side skeletons to ensure type
 * compatibility.
 */

package svcaccountmgmt

import (
    "reflect"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/bindings"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/data"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/protocol"
)




// The ``CreateSpec`` class contains information to create a new service account. **Warning:** This class is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
 type ServiceAccountCreateSpec struct {
    // service account name. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    AccountName string
}


// TODO some error doesn't have ErrorType. Resolve this later VAPI-3009
//


// The ``DeleteSpec`` class contains information that is required for the deletion of the service account. **Warning:** This class is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
 type ServiceAccountDeleteSpec struct {
    // service account name. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    AccountName string
}


// TODO some error doesn't have ErrorType. Resolve this later VAPI-3009
//


// The ``OutputSpec`` class contains information that is returned on successful creation of the service account. **Warning:** This class is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
 type ServiceAccountOutputSpec struct {
    // service account password. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    AccountPassword string
}


// TODO some error doesn't have ErrorType. Resolve this later VAPI-3009
//





func serviceAccountCreateInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["create_spec"] = bindings.NewReferenceType(ServiceAccountCreateSpecBindingType)
    fieldNameMap["create_spec"] = "CreateSpec"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func serviceAccountCreateOutputType() bindings.BindingType {
    return bindings.NewReferenceType(ServiceAccountOutputSpecBindingType)
}

func serviceAccountCreateRestMetadata() protocol.OperationRestMetadata {
    paramsTypeMap := map[string]bindings.BindingType{}
    pathParams := map[string]string{}
    queryParams := map[string]string{}
    headerParams := map[string]string{}
    resultHeaders := map[string]string{}
    errorHeaders := map[string]string{}
    return protocol.NewOperationRestMetadata(
      paramsTypeMap,
      pathParams,
      queryParams,
      headerParams,
      "",
      "null",
      "",
       resultHeaders,
       0,
       errorHeaders,
       map[string]int{"InvalidArgument": 400,"AlreadyExists": 400,"Error": 500})
}


func serviceAccountDeleteInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["delete_spec"] = bindings.NewReferenceType(ServiceAccountDeleteSpecBindingType)
    fieldNameMap["delete_spec"] = "DeleteSpec"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func serviceAccountDeleteOutputType() bindings.BindingType {
    return bindings.NewVoidType()
}

func serviceAccountDeleteRestMetadata() protocol.OperationRestMetadata {
    paramsTypeMap := map[string]bindings.BindingType{}
    pathParams := map[string]string{}
    queryParams := map[string]string{}
    headerParams := map[string]string{}
    resultHeaders := map[string]string{}
    errorHeaders := map[string]string{}
    return protocol.NewOperationRestMetadata(
      paramsTypeMap,
      pathParams,
      queryParams,
      headerParams,
      "",
      "null",
      "",
       resultHeaders,
       0,
       errorHeaders,
       map[string]int{"NotFound": 404,"InvalidArgument": 400,"Error": 500})
}



func ServiceAccountCreateSpecBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["account_name"] = bindings.NewStringType()
    fieldNameMap["account_name"] = "AccountName"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.svcaccountmgmt.service_account.create_spec",fields, reflect.TypeOf(ServiceAccountCreateSpec{}), fieldNameMap, validators)
}

func ServiceAccountDeleteSpecBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["account_name"] = bindings.NewStringType()
    fieldNameMap["account_name"] = "AccountName"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.svcaccountmgmt.service_account.delete_spec",fields, reflect.TypeOf(ServiceAccountDeleteSpec{}), fieldNameMap, validators)
}

func ServiceAccountOutputSpecBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["account_password"] = bindings.NewSecretType()
    fieldNameMap["account_password"] = "AccountPassword"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.svcaccountmgmt.service_account.output_spec",fields, reflect.TypeOf(ServiceAccountOutputSpec{}), fieldNameMap, validators)
}


