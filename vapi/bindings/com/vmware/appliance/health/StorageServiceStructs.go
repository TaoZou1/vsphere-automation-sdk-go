/* Copyright © 2019 VMware, Inc. All Rights Reserved.
     SPDX-License-Identifier: BSD-2-Clause */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Data type definitions file for service: Storage.
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
 
type Storage_HealthLevel string

const (
    // The service health is degraded. The service might have serious problems.
     Storage_HealthLevel_orange Storage_HealthLevel = "orange"
    // No health data is available for this service.
     Storage_HealthLevel_gray Storage_HealthLevel = "gray"
    // Service is healthy.
     Storage_HealthLevel_green Storage_HealthLevel = "green"
    // The service is unavaiable, not functioning properly, or will stop functioning soon.
     Storage_HealthLevel_red Storage_HealthLevel = "red"
    // The service is healthy state, but experiencing some levels of problems.
     Storage_HealthLevel_yellow Storage_HealthLevel = "yellow"
)

func (h Storage_HealthLevel) Storage_HealthLevel() bool {
    switch h {
        case Storage_HealthLevel_orange:
            return true
        case Storage_HealthLevel_gray:
            return true
        case Storage_HealthLevel_green:
            return true
        case Storage_HealthLevel_red:
            return true
        case Storage_HealthLevel_yellow:
            return true
        default:
            return false
    }
}








func storageGetInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func storageGetOutputType() bindings.BindingType {
    return bindings.NewEnumType("com.vmware.appliance.health.storage.health_level", reflect.TypeOf(Storage_HealthLevel(Storage_HealthLevel_orange)))
}

func storageGetRestMetadata() protocol.OperationRestMetadata {
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




