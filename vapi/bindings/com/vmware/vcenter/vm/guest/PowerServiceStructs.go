/* Copyright © 2019 VMware, Inc. All Rights Reserved.
     SPDX-License-Identifier: BSD-2-Clause */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Data type definitions file for service: Power.
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



// Possible guest power states.
//
// <p> See {@link com.vmware.vapi.bindings.ApiEnumeration enumerated types description}.
 
type Power_State string

const (
    // The guest OS is running.
     Power_State_RUNNING Power_State = "RUNNING"
    // The guest OS is shutting down
     Power_State_SHUTTING_DOWN Power_State = "SHUTTING_DOWN"
    // The guest OS is resetting
     Power_State_RESETTING Power_State = "RESETTING"
    // The guest OS is in standby.
     Power_State_STANDBY Power_State = "STANDBY"
    // The guest OS is not running.
     Power_State_NOT_RUNNING Power_State = "NOT_RUNNING"
    // The guest OS power state is unknown.
     Power_State_UNAVAILABLE Power_State = "UNAVAILABLE"
)

func (s Power_State) Power_State() bool {
    switch s {
        case Power_State_RUNNING:
            return true
        case Power_State_SHUTTING_DOWN:
            return true
        case Power_State_RESETTING:
            return true
        case Power_State_STANDBY:
            return true
        case Power_State_NOT_RUNNING:
            return true
        case Power_State_UNAVAILABLE:
            return true
        default:
            return false
    }
}





// Information about the guest operating system power state.
 type PowerInfo struct {
    // The power state of the guest operating system.
    State Power_State
    // Flag indicating if the virtual machine is ready to process soft power operations.
    OperationsReady bool
}



func (PowerInfo PowerInfo) Error() string {
    return "com.vmware.vcenter.vm.guest.info"
}






func powerGetInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["vm"] = bindings.NewIdType([]string {"VirtualMachine"}, "")
    fieldNameMap["vm"] = "Vm"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func powerGetOutputType() bindings.BindingType {
    return bindings.NewReferenceType(PowerInfoBindingType)
}

func powerGetRestMetadata() protocol.OperationRestMetadata {
    paramsTypeMap := map[string]bindings.BindingType{}
    pathParams := map[string]string{}
    queryParams := map[string]string{}
    headerParams := map[string]string{}
    paramsTypeMap["vm"] = bindings.NewIdType([]string {"VirtualMachine"}, "")
    paramsTypeMap["vm"] = bindings.NewIdType([]string {"VirtualMachine"}, "")
    pathParams["vm"] = "vm"
    resultHeaders := map[string]string{}
    errorHeaders := map[string]string{}
    return protocol.NewOperationRestMetadata(
      paramsTypeMap,
      pathParams,
      queryParams,
      headerParams,
      "",
      "GET",
      "/vcenter/vm/{vm}/guest/power",
       resultHeaders,
       200,
       errorHeaders,
       map[string]int{"Error": 500,"NotFound": 404})
}


func powerShutdownInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["vm"] = bindings.NewIdType([]string {"VirtualMachine"}, "")
    fieldNameMap["vm"] = "Vm"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func powerShutdownOutputType() bindings.BindingType {
    return bindings.NewVoidType()
}

func powerShutdownRestMetadata() protocol.OperationRestMetadata {
    paramsTypeMap := map[string]bindings.BindingType{}
    pathParams := map[string]string{}
    queryParams := map[string]string{}
    headerParams := map[string]string{}
    paramsTypeMap["vm"] = bindings.NewIdType([]string {"VirtualMachine"}, "")
    paramsTypeMap["vm"] = bindings.NewIdType([]string {"VirtualMachine"}, "")
    pathParams["vm"] = "vm"
    resultHeaders := map[string]string{}
    errorHeaders := map[string]string{}
    return protocol.NewOperationRestMetadata(
      paramsTypeMap,
      pathParams,
      queryParams,
      headerParams,
      "",
      "POST",
      "/vcenter/vm/{vm}/guest/power",
       resultHeaders,
       204,
       errorHeaders,
       map[string]int{"Error": 500,"NotFound": 404,"AlreadyInDesiredState": 400,"ServiceUnavailable": 503,"NotAllowedInCurrentState": 400,"ResourceBusy": 500,"Unsupported": 400})
}


func powerRebootInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["vm"] = bindings.NewIdType([]string {"VirtualMachine"}, "")
    fieldNameMap["vm"] = "Vm"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func powerRebootOutputType() bindings.BindingType {
    return bindings.NewVoidType()
}

func powerRebootRestMetadata() protocol.OperationRestMetadata {
    paramsTypeMap := map[string]bindings.BindingType{}
    pathParams := map[string]string{}
    queryParams := map[string]string{}
    headerParams := map[string]string{}
    paramsTypeMap["vm"] = bindings.NewIdType([]string {"VirtualMachine"}, "")
    paramsTypeMap["vm"] = bindings.NewIdType([]string {"VirtualMachine"}, "")
    pathParams["vm"] = "vm"
    resultHeaders := map[string]string{}
    errorHeaders := map[string]string{}
    return protocol.NewOperationRestMetadata(
      paramsTypeMap,
      pathParams,
      queryParams,
      headerParams,
      "",
      "POST",
      "/vcenter/vm/{vm}/guest/power",
       resultHeaders,
       204,
       errorHeaders,
       map[string]int{"Error": 500,"NotFound": 404,"NotAllowedInCurrentState": 400,"ServiceUnavailable": 503,"ResourceBusy": 500,"Unsupported": 400})
}


func powerStandbyInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["vm"] = bindings.NewIdType([]string {"VirtualMachine"}, "")
    fieldNameMap["vm"] = "Vm"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func powerStandbyOutputType() bindings.BindingType {
    return bindings.NewVoidType()
}

func powerStandbyRestMetadata() protocol.OperationRestMetadata {
    paramsTypeMap := map[string]bindings.BindingType{}
    pathParams := map[string]string{}
    queryParams := map[string]string{}
    headerParams := map[string]string{}
    paramsTypeMap["vm"] = bindings.NewIdType([]string {"VirtualMachine"}, "")
    paramsTypeMap["vm"] = bindings.NewIdType([]string {"VirtualMachine"}, "")
    pathParams["vm"] = "vm"
    resultHeaders := map[string]string{}
    errorHeaders := map[string]string{}
    return protocol.NewOperationRestMetadata(
      paramsTypeMap,
      pathParams,
      queryParams,
      headerParams,
      "",
      "POST",
      "/vcenter/vm/{vm}/guest/power",
       resultHeaders,
       204,
       errorHeaders,
       map[string]int{"Error": 500,"AlreadyInDesiredState": 400,"NotFound": 404,"ServiceUnavailable": 503,"NotAllowedInCurrentState": 400,"ResourceBusy": 500,"Unsupported": 400})
}



func PowerInfoBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["state"] = bindings.NewEnumType("com.vmware.vcenter.vm.guest.power.state", reflect.TypeOf(Power_State(Power_State_RUNNING)))
    fieldNameMap["state"] = "State"
    fields["operations_ready"] = bindings.NewBooleanType()
    fieldNameMap["operations_ready"] = "OperationsReady"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.vm.guest.power.info",fields, reflect.TypeOf(PowerInfo{}), fieldNameMap, validators)
}


