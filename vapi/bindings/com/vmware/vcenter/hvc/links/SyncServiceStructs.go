/* Copyright © 2019 VMware, Inc. All Rights Reserved.
     SPDX-License-Identifier: BSD-2-Clause */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Data type definitions file for service: Sync.
 * Includes binding types of a structures and enumerations defined in the service.
 * Shared by client-side stubs and server-side skeletons to ensure type
 * compatibility.
 */

package links

import (
    "reflect"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/bindings/com/vmware/vapi/std"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/bindings"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/data"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/protocol"
)




// The ``Info`` class contains information about the result of remediate. **Warning:** This class is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
 type SyncRemediationStatus struct {
        StatusMessage *std.LocalizableMessage
}


// TODO some error doesn't have ErrorType. Resolve this later VAPI-3009
//


// The ``Credentials`` interface specifies user credentials to make a successful connection to remote endpoint. **Warning:** This class is available as Technology Preview. These are early access APIs provided to test, automate and provide feedback on the feature. Since this can change based on feedback, VMware does not guarantee backwards compatibility and recommends against using them in production environments. Some Technology Preview APIs might only be applicable to specific environments.
 type SyncCredentials struct {
    // Name of the user to authenticate. **Warning:** This property is available as Technology Preview. These are early access APIs provided to test, automate and provide feedback on the feature. Since this can change based on feedback, VMware does not guarantee backwards compatibility and recommends against using them in production environments. Some Technology Preview APIs might only be applicable to specific environments.
    UserName string
    // Password for the user. **Warning:** This property is available as Technology Preview. These are early access APIs provided to test, automate and provide feedback on the feature. Since this can change based on feedback, VMware does not guarantee backwards compatibility and recommends against using them in production environments. Some Technology Preview APIs might only be applicable to specific environments.
    Password string
}


// TODO some error doesn't have ErrorType. Resolve this later VAPI-3009
//





func syncResetInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["link"] = bindings.NewIdType([]string {"com.vmware.vcenter.hvc.Links"}, "")
    fieldNameMap["link"] = "Link"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func syncResetOutputType() bindings.BindingType {
    return bindings.NewVoidType()
}

func syncResetRestMetadata() protocol.OperationRestMetadata {
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
       map[string]int{"Error": 500,"NotFound": 404,"Unauthorized": 403})
}


func syncRemediateInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["link"] = bindings.NewIdType([]string {"com.vmware.vcenter.hvc.Links"}, "")
    fields["credentials"] = bindings.NewOptionalType(bindings.NewReferenceType(SyncCredentialsBindingType))
    fieldNameMap["link"] = "Link"
    fieldNameMap["credentials"] = "Credentials"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func syncRemediateOutputType() bindings.BindingType {
    return bindings.NewReferenceType(SyncRemediationStatusBindingType)
}

func syncRemediateRestMetadata() protocol.OperationRestMetadata {
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
       map[string]int{"Error": 500,"NotFound": 404,"Unauthorized": 403,"ResourceBusy": 500,"TimedOut": 500})
}



func SyncRemediationStatusBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["status_message"] = bindings.NewOptionalType(bindings.NewReferenceType(std.LocalizableMessageBindingType))
    fieldNameMap["status_message"] = "StatusMessage"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.hvc.links.sync.remediation_status",fields, reflect.TypeOf(SyncRemediationStatus{}), fieldNameMap, validators)
}

func SyncCredentialsBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["user_name"] = bindings.NewStringType()
    fieldNameMap["user_name"] = "UserName"
    fields["password"] = bindings.NewSecretType()
    fieldNameMap["password"] = "Password"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.hvc.links.sync.credentials",fields, reflect.TypeOf(SyncCredentials{}), fieldNameMap, validators)
}


