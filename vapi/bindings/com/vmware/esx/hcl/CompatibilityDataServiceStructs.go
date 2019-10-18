/* Copyright © 2019 VMware, Inc. All Rights Reserved.
     SPDX-License-Identifier: BSD-2-Clause */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Data type definitions file for service: CompatibilityData.
 * Includes binding types of a structures and enumerations defined in the service.
 * Shared by client-side stubs and server-side skeletons to ensure type
 * compatibility.
 */

package hcl

import (
    "reflect"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/bindings"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/data"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/protocol"
    "time"
)





// The ``Status`` class contains properties to describe the information available for the compatibility data. **Warning:** This class is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
 type CompatibilityDataStatus struct {
    // Indicates when the data was last updated. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    UpdatedAt time.Time
    // Notifications returned by the operation. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    Notifications Notifications
}









func compatibilityDataGetInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func compatibilityDataGetOutputType() bindings.BindingType {
    return bindings.NewReferenceType(CompatibilityDataStatusBindingType)
}

func compatibilityDataGetRestMetadata() protocol.OperationRestMetadata {
    paramsTypeMap := map[string]bindings.BindingType{}
    pathParams := map[string]string{}
    queryParams := map[string]string{}
    headerParams := map[string]string{}
    resultHeaders := map[string]string{}
    errorHeaders := map[string]string{}
    errorHeaders["Unauthenticated.challenge"] = "WWW-Authenticate"
    return protocol.NewOperationRestMetadata(
      paramsTypeMap,
      pathParams,
      queryParams,
      headerParams,
      "",
      "GET",
      "/esx/hcl/compatibility-data/status",
       resultHeaders,
       200,
       errorHeaders,
       map[string]int{"NotAllowedInCurrentState": 400,"Unauthenticated": 401,"ResourceInaccessible": 500,"Error": 500})
}


func compatibilityDataUpdateInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func compatibilityDataUpdateOutputType() bindings.BindingType {
    return bindings.NewVoidType()
}

func compatibilityDataUpdateRestMetadata() protocol.OperationRestMetadata {
    paramsTypeMap := map[string]bindings.BindingType{}
    pathParams := map[string]string{}
    queryParams := map[string]string{}
    headerParams := map[string]string{}
    resultHeaders := map[string]string{}
    errorHeaders := map[string]string{}
    errorHeaders["Unauthenticated.challenge"] = "WWW-Authenticate"
    return protocol.NewOperationRestMetadata(
      paramsTypeMap,
      pathParams,
      queryParams,
      headerParams,
      "",
      "POST",
      "/esx/hcl/compatibility-data",
       resultHeaders,
       202,
       errorHeaders,
       map[string]int{"NotAllowedInCurrentState": 400,"Unauthenticated": 401,"ResourceInaccessible": 500,"Error": 500})
}



func CompatibilityDataStatusBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["updated_at"] = bindings.NewDateTimeType()
    fieldNameMap["updated_at"] = "UpdatedAt"
    fields["notifications"] = bindings.NewReferenceType(NotificationsBindingType)
    fieldNameMap["notifications"] = "Notifications"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.esx.hcl.compatibility_data.status",fields, reflect.TypeOf(CompatibilityDataStatus{}), fieldNameMap, validators)
}


