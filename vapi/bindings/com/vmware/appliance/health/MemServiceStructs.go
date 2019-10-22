/* Copyright © 2019 VMware, Inc. All Rights Reserved.
     SPDX-License-Identifier: BSD-2-Clause */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Data type definitions file for service: Mem.
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
 
type Mem_HealthLevel string

const (
    // The service health is degraded. The service might have serious problems
     Mem_HealthLevel_orange Mem_HealthLevel = "orange"
    // No health data is available for this service.
     Mem_HealthLevel_gray Mem_HealthLevel = "gray"
    // Service is healthy.
     Mem_HealthLevel_green Mem_HealthLevel = "green"
    // The service is unavaiable, not functioning properly, or will stop functioning soon.
     Mem_HealthLevel_red Mem_HealthLevel = "red"
    // The service is healthy state, but experiencing some levels of problems.
     Mem_HealthLevel_yellow Mem_HealthLevel = "yellow"
)

func (h Mem_HealthLevel) Mem_HealthLevel() bool {
    switch h {
        case Mem_HealthLevel_orange:
            return true
        case Mem_HealthLevel_gray:
            return true
        case Mem_HealthLevel_green:
            return true
        case Mem_HealthLevel_red:
            return true
        case Mem_HealthLevel_yellow:
            return true
        default:
            return false
    }
}








func memGetInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func memGetOutputType() bindings.BindingType {
    return bindings.NewEnumType("com.vmware.appliance.health.mem.health_level", reflect.TypeOf(Mem_HealthLevel(Mem_HealthLevel_orange)))
}

func memGetRestMetadata() protocol.OperationRestMetadata {
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




