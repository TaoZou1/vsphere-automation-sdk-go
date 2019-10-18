/* Copyright © 2019 VMware, Inc. All Rights Reserved.
     SPDX-License-Identifier: BSD-2-Clause */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Data type definitions file for service: Password.
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





// The ``ChangeSpec`` class contains information that is required to change a service account password. **Warning:** This class is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
 type PasswordChangeSpec struct {
    // service account name. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    AccountName string
    // current password. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    OldPassword string
}







// The ``ResetSpec`` class contains information that is required to reset a service account password. **Warning:** This class is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
 type PasswordResetSpec struct {
    // service account name. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    AccountName string
}







// The ``OutputSpec`` class contains information that is returned on successful reset of the service account password. **Warning:** This class is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
 type PasswordOutputSpec struct {
    // service account password. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    NewPassword string
}









func passwordChangeInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["change_spec"] = bindings.NewReferenceType(PasswordChangeSpecBindingType)
    fieldNameMap["change_spec"] = "ChangeSpec"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func passwordChangeOutputType() bindings.BindingType {
    return bindings.NewReferenceType(PasswordOutputSpecBindingType)
}

func passwordChangeRestMetadata() protocol.OperationRestMetadata {
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
       map[string]int{"NotFound": 404,"InvalidArgument": 400,"Unauthorized": 403,"Error": 500})
}


func passwordResetInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["reset_spec"] = bindings.NewReferenceType(PasswordResetSpecBindingType)
    fieldNameMap["reset_spec"] = "ResetSpec"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func passwordResetOutputType() bindings.BindingType {
    return bindings.NewReferenceType(PasswordOutputSpecBindingType)
}

func passwordResetRestMetadata() protocol.OperationRestMetadata {
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



func PasswordChangeSpecBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["account_name"] = bindings.NewStringType()
    fieldNameMap["account_name"] = "AccountName"
    fields["old_password"] = bindings.NewSecretType()
    fieldNameMap["old_password"] = "OldPassword"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.svcaccountmgmt.password.change_spec",fields, reflect.TypeOf(PasswordChangeSpec{}), fieldNameMap, validators)
}

func PasswordResetSpecBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["account_name"] = bindings.NewStringType()
    fieldNameMap["account_name"] = "AccountName"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.svcaccountmgmt.password.reset_spec",fields, reflect.TypeOf(PasswordResetSpec{}), fieldNameMap, validators)
}

func PasswordOutputSpecBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["new_password"] = bindings.NewSecretType()
    fieldNameMap["new_password"] = "NewPassword"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.svcaccountmgmt.password.output_spec",fields, reflect.TypeOf(PasswordOutputSpec{}), fieldNameMap, validators)
}


