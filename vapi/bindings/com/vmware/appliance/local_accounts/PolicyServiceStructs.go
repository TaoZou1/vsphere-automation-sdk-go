/* Copyright © 2019 VMware, Inc. All Rights Reserved.
     SPDX-License-Identifier: BSD-2-Clause */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Data type definitions file for service: Policy.
 * Includes binding types of a structures and enumerations defined in the service.
 * Shared by client-side stubs and server-side skeletons to ensure type
 * compatibility.
 */

package local_accounts

import (
    "reflect"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/bindings"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/data"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/protocol"
)




// The ``Info`` class defines the global password policy.
 type PolicyInfo struct {
    // Maximum number of days a password may be used. If the password is older than this, a password change will be forced.
    MaxDays *int64
    // Minimum number of days allowed between password changes. Any password changes attempted sooner than this will be rejected.
    MinDays *int64
    // Number of days warning given before a password expires. A zero means warning is given only upon the day of expiration.
    WarnDays *int64
}



func (PolicyInfo PolicyInfo) Error() string {
    return "com.vmware.appliance.local_accounts.info"
}






func policyGetInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func policyGetOutputType() bindings.BindingType {
    return bindings.NewReferenceType(PolicyInfoBindingType)
}

func policyGetRestMetadata() protocol.OperationRestMetadata {
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
       map[string]int{"Error": 500})
}


func policySetInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["policy"] = bindings.NewReferenceType(PolicyInfoBindingType)
    fieldNameMap["policy"] = "Policy"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func policySetOutputType() bindings.BindingType {
    return bindings.NewVoidType()
}

func policySetRestMetadata() protocol.OperationRestMetadata {
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
       map[string]int{"InvalidArgument": 400,"Error": 500})
}



func PolicyInfoBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["max_days"] = bindings.NewOptionalType(bindings.NewIntegerType())
    fieldNameMap["max_days"] = "MaxDays"
    fields["min_days"] = bindings.NewOptionalType(bindings.NewIntegerType())
    fieldNameMap["min_days"] = "MinDays"
    fields["warn_days"] = bindings.NewOptionalType(bindings.NewIntegerType())
    fieldNameMap["warn_days"] = "WarnDays"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.appliance.local_accounts.policy.info",fields, reflect.TypeOf(PolicyInfo{}), fieldNameMap, validators)
}


