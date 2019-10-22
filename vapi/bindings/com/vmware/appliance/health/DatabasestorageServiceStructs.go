/* Copyright © 2019 VMware, Inc. All Rights Reserved.
     SPDX-License-Identifier: BSD-2-Clause */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Data type definitions file for service: Databasestorage.
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



// ``HealthLevel`` enumeration class Defines service health levels.
//
// <p> See {@link com.vmware.vapi.bindings.ApiEnumeration enumerated types description}.
 
type Databasestorage_HealthLevel string

const (
    // The service health is degraded. The service might have serious problems.
     Databasestorage_HealthLevel_orange Databasestorage_HealthLevel = "orange"
    // No health data is available for this service.
     Databasestorage_HealthLevel_gray Databasestorage_HealthLevel = "gray"
    // The service is healthy.
     Databasestorage_HealthLevel_green Databasestorage_HealthLevel = "green"
    // The service is unavaiable, not functioning properly, or will stop functioning soon.
     Databasestorage_HealthLevel_red Databasestorage_HealthLevel = "red"
    // The service is healthy but experiencing some problems.
     Databasestorage_HealthLevel_yellow Databasestorage_HealthLevel = "yellow"
)

func (h Databasestorage_HealthLevel) Databasestorage_HealthLevel() bool {
    switch h {
        case Databasestorage_HealthLevel_orange:
            return true
        case Databasestorage_HealthLevel_gray:
            return true
        case Databasestorage_HealthLevel_green:
            return true
        case Databasestorage_HealthLevel_red:
            return true
        case Databasestorage_HealthLevel_yellow:
            return true
        default:
            return false
    }
}








func databasestorageGetInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func databasestorageGetOutputType() bindings.BindingType {
    return bindings.NewEnumType("com.vmware.appliance.health.databasestorage.health_level", reflect.TypeOf(Databasestorage_HealthLevel(Databasestorage_HealthLevel_orange)))
}

func databasestorageGetRestMetadata() protocol.OperationRestMetadata {
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




