/* Copyright © 2019 VMware, Inc. All Rights Reserved.
     SPDX-License-Identifier: BSD-2-Clause */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Data type definitions file for service: Health.
 * Includes binding types of a structures and enumerations defined in the service.
 * Shared by client-side stubs and server-side skeletons to ensure type
 * compatibility.
 */

package registries

import (
    "reflect"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/bindings/com/vmware/vapi/std"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/bindings"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/data"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/protocol"
)



// The ``Status`` enumeration class describes the status of the container registry. **Warning:** This enumeration is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
//
// <p> See {@link com.vmware.vapi.bindings.ApiEnumeration enumerated types description}.
 
type Health_Status string

const (
    // Container registry is starting. **Warning:** This constant field is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
     Health_Status_STARTING Health_Status = "STARTING"
    // Container registry is running. **Warning:** This constant field is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
     Health_Status_RUNNING Health_Status = "RUNNING"
    // Container registry is running with some warning, for example, storage reaches the threshold configuration. **Warning:** This constant field is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
     Health_Status_WARNING Health_Status = "WARNING"
    // Container registry failed to start or stopped with fatal error. **Warning:** This constant field is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
     Health_Status_ERROR Health_Status = "ERROR"
    // Container registry is being deleted. **Warning:** This constant field is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
     Health_Status_DELETING Health_Status = "DELETING"
)

func (s Health_Status) Health_Status() bool {
    switch s {
        case Health_Status_STARTING:
            return true
        case Health_Status_RUNNING:
            return true
        case Health_Status_WARNING:
            return true
        case Health_Status_ERROR:
            return true
        case Health_Status_DELETING:
            return true
        default:
            return false
    }
}





// The ``Info`` class contains health information about a container registry. **Warning:** This class is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
 type HealthInfo struct {
    // Container registry status. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    Status Health_Status
    // Details about the status. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    Details *std.LocalizableMessage
}


// TODO some error doesn't have ErrorType. Resolve this later VAPI-3009
//





func healthGetInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["registry"] = bindings.NewIdType([]string {"com.vmware.vcenter.content.Registry"}, "")
    fieldNameMap["registry"] = "Registry"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func healthGetOutputType() bindings.BindingType {
    return bindings.NewReferenceType(HealthInfoBindingType)
}

func healthGetRestMetadata() protocol.OperationRestMetadata {
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
       map[string]int{"Error": 500,"NotFound": 404,"Unauthenticated": 401,"Unauthorized": 403})
}



func HealthInfoBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["status"] = bindings.NewEnumType("com.vmware.vcenter.content.registries.health.status", reflect.TypeOf(Health_Status(Health_Status_STARTING)))
    fieldNameMap["status"] = "Status"
    fields["details"] = bindings.NewOptionalType(bindings.NewReferenceType(std.LocalizableMessageBindingType))
    fieldNameMap["details"] = "Details"
    var validators = []bindings.Validator{}
    uv1 := bindings.NewUnionValidator("status",
        map[string][]bindings.FieldData {
            "WARNING": []bindings.FieldData {
                 bindings.NewFieldData("details", true),
            },
            "ERROR": []bindings.FieldData {
                 bindings.NewFieldData("details", true),
            },
            "STARTING": []bindings.FieldData {},
            "RUNNING": []bindings.FieldData {},
            "DELETING": []bindings.FieldData {},
        },
    )
    validators = append(validators, uv1)
    return bindings.NewStructType("com.vmware.vcenter.content.registries.health.info",fields, reflect.TypeOf(HealthInfo{}), fieldNameMap, validators)
}


