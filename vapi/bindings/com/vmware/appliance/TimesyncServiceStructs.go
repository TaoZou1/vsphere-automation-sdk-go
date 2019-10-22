/* Copyright © 2019 VMware, Inc. All Rights Reserved.
     SPDX-License-Identifier: BSD-2-Clause */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Data type definitions file for service: Timesync.
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



// The ``TimeSyncMode`` enumeration class defines time synchronization modes. This enumeration was added in vSphere API 6.7.
//
// <p> See {@link com.vmware.vapi.bindings.ApiEnumeration enumerated types description}.
 
type Timesync_TimeSyncMode string

const (
    // Time synchronization is disabled. This constant field was added in vSphere API 6.7.
     Timesync_TimeSyncMode_DISABLED Timesync_TimeSyncMode = "DISABLED"
    // NTP-based time synchronization. This constant field was added in vSphere API 6.7.
     Timesync_TimeSyncMode_NTP Timesync_TimeSyncMode = "NTP"
    // VMware Tool-based time synchronization. This constant field was added in vSphere API 6.7.
     Timesync_TimeSyncMode_HOST Timesync_TimeSyncMode = "HOST"
)

func (t Timesync_TimeSyncMode) Timesync_TimeSyncMode() bool {
    switch t {
        case Timesync_TimeSyncMode_DISABLED:
            return true
        case Timesync_TimeSyncMode_NTP:
            return true
        case Timesync_TimeSyncMode_HOST:
            return true
        default:
            return false
    }
}








func timesyncSetInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["mode"] = bindings.NewEnumType("com.vmware.appliance.timesync.time_sync_mode", reflect.TypeOf(Timesync_TimeSyncMode(Timesync_TimeSyncMode_DISABLED)))
    fieldNameMap["mode"] = "Mode"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func timesyncSetOutputType() bindings.BindingType {
    return bindings.NewVoidType()
}

func timesyncSetRestMetadata() protocol.OperationRestMetadata {
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


func timesyncGetInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func timesyncGetOutputType() bindings.BindingType {
    return bindings.NewEnumType("com.vmware.appliance.timesync.time_sync_mode", reflect.TypeOf(Timesync_TimeSyncMode(Timesync_TimeSyncMode_DISABLED)))
}

func timesyncGetRestMetadata() protocol.OperationRestMetadata {
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




