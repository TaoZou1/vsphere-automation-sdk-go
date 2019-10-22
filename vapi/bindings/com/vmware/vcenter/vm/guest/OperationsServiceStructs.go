/* Copyright © 2019 VMware, Inc. All Rights Reserved.
     SPDX-License-Identifier: BSD-2-Clause */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Data type definitions file for service: Operations.
 * Includes binding types of a structures and enumerations defined in the service.
 * Shared by client-side stubs and server-side skeletons to ensure type
 * compatibility.
 */

package guest

import (
    "reflect"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/bindings"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/data"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/protocol"
)




// Guest operating system operation status information. **Warning:** This class is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
 type OperationsInfo struct {
    // Guest operations availability. If map with bool value, the virtual machine is ready to process guest operations. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    GuestOperationsReady bool
    // Interactive guest operations availability. If map with bool value, the virtual machine is ready to process interactive guest operations. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    InteractiveGuestOperationsReady bool
}



//





func operationsGetInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["vm"] = bindings.NewIdType([]string {"VirtualMachine"}, "")
    fieldNameMap["vm"] = "Vm"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func operationsGetOutputType() bindings.BindingType {
    return bindings.NewReferenceType(OperationsInfoBindingType)
}

func operationsGetRestMetadata() protocol.OperationRestMetadata {
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
       map[string]int{"NotFound": 404})
}



func OperationsInfoBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["guest_operations_ready"] = bindings.NewBooleanType()
    fieldNameMap["guest_operations_ready"] = "GuestOperationsReady"
    fields["interactive_guest_operations_ready"] = bindings.NewBooleanType()
    fieldNameMap["interactive_guest_operations_ready"] = "InteractiveGuestOperationsReady"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.vm.guest.operations.info",fields, reflect.TypeOf(OperationsInfo{}), fieldNameMap, validators)
}


