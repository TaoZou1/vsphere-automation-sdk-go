/* Copyright © 2019 VMware, Inc. All Rights Reserved.
     SPDX-License-Identifier: BSD-2-Clause */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Data type definitions file for service: TechPreview.
 * Includes binding types of a structures and enumerations defined in the service.
 * Shared by client-side stubs and server-side skeletons to ensure type
 * compatibility.
 */

package features

import (
    "reflect"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/bindings"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/data"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/protocol"
)



// The ``Status`` enumeration class defines the possible states of a feature state switch. **Warning:** This enumeration is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
//
// <p> See {@link com.vmware.vapi.bindings.ApiEnumeration enumerated types description}.
 
type TechPreview_Status string

const (
    // Feature state switch is disabled. **Warning:** This constant field is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
     TechPreview_Status_DISABLED TechPreview_Status = "DISABLED"
    // Feature state switch is enabled. **Warning:** This constant field is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
     TechPreview_Status_ENABLED TechPreview_Status = "ENABLED"
)

func (s TechPreview_Status) TechPreview_Status() bool {
    switch s {
        case TechPreview_Status_DISABLED:
            return true
        case TechPreview_Status_ENABLED:
            return true
        default:
            return false
    }
}








func techPreviewGetInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["host"] = bindings.NewIdType([]string {"HostSystem"}, "")
    fields["features"] = bindings.NewOptionalType(bindings.NewSetType(bindings.NewStringType(), reflect.TypeOf(map[string]bool{})))
    fieldNameMap["host"] = "Host"
    fieldNameMap["features"] = "Features"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func techPreviewGetOutputType() bindings.BindingType {
    return bindings.NewMapType(bindings.NewStringType(), bindings.NewEnumType("com.vmware.vcenter.host.features.tech_preview.status", reflect.TypeOf(TechPreview_Status(TechPreview_Status_DISABLED))),reflect.TypeOf(map[string]TechPreview_Status{}))
}

func techPreviewGetRestMetadata() protocol.OperationRestMetadata {
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
      "null",
      "",
       resultHeaders,
       0,
       errorHeaders,
       map[string]int{"Error": 500,"NotFound": 404,"Unauthenticated": 401})
}


func techPreviewUpdateInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["host"] = bindings.NewIdType([]string {"HostSystem"}, "")
    fields["feature_status"] = bindings.NewMapType(bindings.NewStringType(), bindings.NewEnumType("com.vmware.vcenter.host.features.tech_preview.status", reflect.TypeOf(TechPreview_Status(TechPreview_Status_DISABLED))),reflect.TypeOf(map[string]TechPreview_Status{}))
    fieldNameMap["host"] = "Host"
    fieldNameMap["feature_status"] = "FeatureStatus"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func techPreviewUpdateOutputType() bindings.BindingType {
    return bindings.NewVoidType()
}

func techPreviewUpdateRestMetadata() protocol.OperationRestMetadata {
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
      "null",
      "",
       resultHeaders,
       0,
       errorHeaders,
       map[string]int{"Error": 500,"NotFound": 404,"Unauthenticated": 401})
}




