/* Copyright © 2019 VMware, Inc. All Rights Reserved.
     SPDX-License-Identifier: BSD-2-Clause */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Data type definitions file for service: System.
 * Includes binding types of a structures and enumerations defined in the service.
 * Shared by client-side stubs and server-side skeletons to ensure type
 * compatibility.
 */

package health

import (
    "reflect"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/bindings"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/data"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/protocol"
)



// ``HealthLevel`` enumeration class Defines health levels.
//
// <p> See {@link com.vmware.vapi.bindings.ApiEnumeration enumerated types description}.
 
type System_HealthLevel string

const (
    // The service health is degraded. The service might have serious problems.
     System_HealthLevel_orange System_HealthLevel = "orange"
    // No health data is available for this service.
     System_HealthLevel_gray System_HealthLevel = "gray"
    // Service is healthy.
     System_HealthLevel_green System_HealthLevel = "green"
    // The service is unavaiable, not functioning properly, or will stop functioning soon.
     System_HealthLevel_red System_HealthLevel = "red"
    // The service is healthy state, but experiencing some levels of problems.
     System_HealthLevel_yellow System_HealthLevel = "yellow"
)

func (h System_HealthLevel) System_HealthLevel() bool {
    switch h {
        case System_HealthLevel_orange:
            return true
        case System_HealthLevel_gray:
            return true
        case System_HealthLevel_green:
            return true
        case System_HealthLevel_red:
            return true
        case System_HealthLevel_yellow:
            return true
        default:
            return false
    }
}








func systemLastcheckInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func systemLastcheckOutputType() bindings.BindingType {
    return bindings.NewDateTimeType()
}

func systemLastcheckRestMetadata() protocol.OperationRestMetadata {
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


func systemGetInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func systemGetOutputType() bindings.BindingType {
    return bindings.NewEnumType("com.vmware.appliance.health.system.health_level", reflect.TypeOf(System_HealthLevel(System_HealthLevel_orange)))
}

func systemGetRestMetadata() protocol.OperationRestMetadata {
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




