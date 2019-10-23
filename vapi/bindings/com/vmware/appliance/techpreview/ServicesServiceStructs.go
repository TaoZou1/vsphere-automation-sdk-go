/* Copyright © 2019 VMware, Inc. All Rights Reserved.
     SPDX-License-Identifier: BSD-2-Clause */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Data type definitions file for service: Services.
 * Includes binding types of a structures and enumerations defined in the service.
 * Shared by client-side stubs and server-side skeletons to ensure type
 * compatibility.
 */

package techpreview

import (
    "reflect"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/bindings"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/data"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/protocol"
)



// ``ServiceOps`` enumeration class Defines service operations.
//
// <p> See {@link com.vmware.vapi.bindings.ApiEnumeration enumerated types description}.
 
type Services_ServiceOps string

const (
    // The service status can be generated.
     Services_ServiceOps_status Services_ServiceOps = "status"
    // The service can be stopped.
     Services_ServiceOps_stop Services_ServiceOps = "stop"
    // The service can be started or restarted.
     Services_ServiceOps_restart Services_ServiceOps = "restart"
)

func (s Services_ServiceOps) Services_ServiceOps() bool {
    switch s {
        case Services_ServiceOps_status:
            return true
        case Services_ServiceOps_stop:
            return true
        case Services_ServiceOps_restart:
            return true
        default:
            return false
    }
}





// ``ServiceInfo`` class Structure that describes a service and the operations you can perform on a service.
 type ServicesServiceInfo struct {
    // The name of the service, for example, "vmware-vpxd". This name is unique per machine.
    Name string
    // A description of the service. The description can be empty.
    Description string
    // The operations encoded in an array, supported by the service.
    Operations []Services_ServiceOps
}



func (ServicesServiceInfo ServicesServiceInfo) Error() string {
    return "com.vmware.appliance.techpreview.service_info"
}






func servicesListInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func servicesListOutputType() bindings.BindingType {
    return bindings.NewListType(bindings.NewReferenceType(ServicesServiceInfoBindingType), reflect.TypeOf([]ServicesServiceInfo{}))
}

func servicesListRestMetadata() protocol.OperationRestMetadata {
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


func servicesStopInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["name"] = bindings.NewStringType()
    fields["timeout"] = bindings.NewIntegerType()
    fieldNameMap["name"] = "Name"
    fieldNameMap["timeout"] = "Timeout"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func servicesStopOutputType() bindings.BindingType {
    return bindings.NewVoidType()
}

func servicesStopRestMetadata() protocol.OperationRestMetadata {
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


func servicesRestartInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["name"] = bindings.NewStringType()
    fields["timeout"] = bindings.NewIntegerType()
    fieldNameMap["name"] = "Name"
    fieldNameMap["timeout"] = "Timeout"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func servicesRestartOutputType() bindings.BindingType {
    return bindings.NewVoidType()
}

func servicesRestartRestMetadata() protocol.OperationRestMetadata {
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



func ServicesServiceInfoBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["name"] = bindings.NewStringType()
    fieldNameMap["name"] = "Name"
    fields["description"] = bindings.NewStringType()
    fieldNameMap["description"] = "Description"
    fields["operations"] = bindings.NewListType(bindings.NewEnumType("com.vmware.appliance.techpreview.services.service_ops", reflect.TypeOf(Services_ServiceOps(Services_ServiceOps_status))), reflect.TypeOf([]Services_ServiceOps{}))
    fieldNameMap["operations"] = "Operations"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.appliance.techpreview.services.service_info",fields, reflect.TypeOf(ServicesServiceInfo{}), fieldNameMap, validators)
}


