/* Copyright © 2019 VMware, Inc. All Rights Reserved.
     SPDX-License-Identifier: BSD-2-Clause */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Data type definitions file for service: Associations.
 * Includes binding types of a structures and enumerations defined in the service.
 * Shared by client-side stubs and server-side skeletons to ensure type
 * compatibility.
 */

package vm

import (
    "reflect"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/bindings/com/vmware/vapi/std"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/bindings"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/data"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/protocol"
)




// The ``Status`` class contains properties that are used to describe the outcome of the add/remove operations. **Warning:** This class is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
 type AssociationsStatus struct {
    // This is true if the complete method succeeded without any errors. Otherwise it is false and all or some methods have failed. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    Success bool
    // The array of error messages. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    ErrorMessages []std.LocalizableMessage
}


// TODO some error doesn't have ErrorType. Resolve this later VAPI-3009
//





func associationsAddInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["tag"] = bindings.NewIdType([]string {"com.vmware.cis.tagging.Tag:VirtualMachine"}, "")
    fields["vms"] = bindings.NewSetType(bindings.NewIdType([]string {"VirtualMachine"}, ""), reflect.TypeOf(map[string]bool{}))
    fieldNameMap["tag"] = "Tag"
    fieldNameMap["vms"] = "Vms"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func associationsAddOutputType() bindings.BindingType {
    return bindings.NewReferenceType(AssociationsStatusBindingType)
}

func associationsAddRestMetadata() protocol.OperationRestMetadata {
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
       map[string]int{"NotFound": 404,"Unauthorized": 403})
}


func associationsRemoveInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["tag"] = bindings.NewIdType([]string {"com.vmware.cis.tagging.Tag:VirtualMachine"}, "")
    fields["vms"] = bindings.NewSetType(bindings.NewIdType([]string {"VirtualMachine"}, ""), reflect.TypeOf(map[string]bool{}))
    fieldNameMap["tag"] = "Tag"
    fieldNameMap["vms"] = "Vms"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func associationsRemoveOutputType() bindings.BindingType {
    return bindings.NewReferenceType(AssociationsStatusBindingType)
}

func associationsRemoveRestMetadata() protocol.OperationRestMetadata {
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
       map[string]int{"NotFound": 404,"Unauthorized": 403})
}


func associationsGetInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["tag"] = bindings.NewIdType([]string {"com.vmware.cis.tagging.Tag:VirtualMachine"}, "")
    fieldNameMap["tag"] = "Tag"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func associationsGetOutputType() bindings.BindingType {
    return bindings.NewSetType(bindings.NewIdType([]string {"VirtualMachine"}, ""), reflect.TypeOf(map[string]bool{}))
}

func associationsGetRestMetadata() protocol.OperationRestMetadata {
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
       map[string]int{"NotFound": 404,"Unauthorized": 403})
}



func AssociationsStatusBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["success"] = bindings.NewBooleanType()
    fieldNameMap["success"] = "Success"
    fields["error_messages"] = bindings.NewListType(bindings.NewReferenceType(std.LocalizableMessageBindingType), reflect.TypeOf([]std.LocalizableMessage{}))
    fieldNameMap["error_messages"] = "ErrorMessages"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.tags.vm.associations.status",fields, reflect.TypeOf(AssociationsStatus{}), fieldNameMap, validators)
}


