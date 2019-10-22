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

package appliance

import (
    "reflect"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/bindings"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/data"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/protocol"
)



// The ``State`` enumeration class defines valid Run State for services. This enumeration was added in vSphere API 6.7.
//
// <p> See {@link com.vmware.vapi.bindings.ApiEnumeration enumerated types description}.
 
type Services_State string

const (
    // Service Run State is Starting, it is still not functional. This constant field was added in vSphere API 6.7.
     Services_State_STARTING Services_State = "STARTING"
    // Service Run State is Stopping, it is not functional. This constant field was added in vSphere API 6.7.
     Services_State_STOPPING Services_State = "STOPPING"
    // Service Run State is Started, it is fully functional. This constant field was added in vSphere API 6.7.
     Services_State_STARTED Services_State = "STARTED"
    // Service Run State is Stopped. This constant field was added in vSphere API 6.7.
     Services_State_STOPPED Services_State = "STOPPED"
)

func (s Services_State) Services_State() bool {
    switch s {
        case Services_State_STARTING:
            return true
        case Services_State_STOPPING:
            return true
        case Services_State_STARTED:
            return true
        case Services_State_STOPPED:
            return true
        default:
            return false
    }
}





// The ``Info`` class contains information about a service. This class was added in vSphere API 6.7.
 type ServicesInfo struct {
    // Service description. This property was added in vSphere API 6.7.
    Description string
    // Running State. This property was added in vSphere API 6.7.
    State Services_State
}



//





func servicesStartInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["service"] = bindings.NewIdType([]string {"com.vmware.appliance.services"}, "")
    fieldNameMap["service"] = "Service"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func servicesStartOutputType() bindings.BindingType {
    return bindings.NewVoidType()
}

func servicesStartRestMetadata() protocol.OperationRestMetadata {
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
       map[string]int{"NotFound": 404,"NotAllowedInCurrentState": 400,"TimedOut": 500,"Error": 500})
}


func servicesStopInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["service"] = bindings.NewIdType([]string {"com.vmware.appliance.services"}, "")
    fieldNameMap["service"] = "Service"
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
       map[string]int{"NotFound": 404,"Error": 500,"NotAllowedInCurrentState": 400})
}


func servicesRestartInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["service"] = bindings.NewIdType([]string {"com.vmware.appliance.services"}, "")
    fieldNameMap["service"] = "Service"
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
       map[string]int{"NotFound": 404,"TimedOut": 500,"NotAllowedInCurrentState": 400,"Error": 500})
}


func servicesGetInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["service"] = bindings.NewIdType([]string {"com.vmware.appliance.services"}, "")
    fieldNameMap["service"] = "Service"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func servicesGetOutputType() bindings.BindingType {
    return bindings.NewReferenceType(ServicesInfoBindingType)
}

func servicesGetRestMetadata() protocol.OperationRestMetadata {
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
       map[string]int{"NotFound": 404,"Error": 500})
}


func servicesListInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func servicesListOutputType() bindings.BindingType {
    return bindings.NewMapType(bindings.NewIdType([]string {"com.vmware.appliance.services"}, ""), bindings.NewReferenceType(ServicesInfoBindingType),reflect.TypeOf(map[string]ServicesInfo{}))
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



func ServicesInfoBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["description"] = bindings.NewStringType()
    fieldNameMap["description"] = "Description"
    fields["state"] = bindings.NewEnumType("com.vmware.appliance.services.state", reflect.TypeOf(Services_State(Services_State_STARTING)))
    fieldNameMap["state"] = "State"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.appliance.services.info",fields, reflect.TypeOf(ServicesInfo{}), fieldNameMap, validators)
}


