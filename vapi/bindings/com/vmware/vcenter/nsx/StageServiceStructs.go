/* Copyright © 2019 VMware, Inc. All Rights Reserved.
     SPDX-License-Identifier: BSD-2-Clause */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Data type definitions file for service: Stage.
 * Includes binding types of a structures and enumerations defined in the service.
 * Shared by client-side stubs and server-side skeletons to ensure type
 * compatibility.
 */

package nsx

import (
    "reflect"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/bindings/com/vmware/vapi/std"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/bindings"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/data"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/protocol"
    "time"
)



// The ``StageStatus`` enumeration class defines the overall staging status for a bundle. **Warning:** This enumeration is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
//
// <p> See {@link com.vmware.vapi.bindings.ApiEnumeration enumerated types description}.
 
type StageStageStatus string

const (
    // The staging is not yet started. **Warning:** This constant field is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
     StageStageStatus_NOT_STARTED StageStageStatus = "NOT_STARTED"
    // Staging is in progress. **Warning:** This constant field is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
     StageStageStatus_IN_PROGRESS StageStageStatus = "IN_PROGRESS"
    // Staging has been Cancelled by the user. **Warning:** This constant field is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
     StageStageStatus_CANCELLED StageStageStatus = "CANCELLED"
    // The Staging operation failed. **Warning:** This constant field is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
     StageStageStatus_FAILED StageStageStatus = "FAILED"
    // The Staging has been completed. **Warning:** This constant field is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
     StageStageStatus_COMPLETED StageStageStatus = "COMPLETED"
)

func (s StageStageStatus) StageStageStatus() bool {
    switch s {
        case StageStageStatus_NOT_STARTED:
            return true
        case StageStageStatus_IN_PROGRESS:
            return true
        case StageStageStatus_CANCELLED:
            return true
        case StageStageStatus_FAILED:
            return true
        case StageStageStatus_COMPLETED:
            return true
        default:
            return false
    }
}




// The ``StageType`` enumeration class defines the overall staging operation type. **Warning:** This enumeration is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
//
// <p> See {@link com.vmware.vapi.bindings.ApiEnumeration enumerated types description}.
 
type StageStageType string

const (
    // Download operation. **Warning:** This constant field is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
     StageStageType_DOWNLOAD StageStageType = "DOWNLOAD"
    // Upload operation. **Warning:** This constant field is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
     StageStageType_UPLOAD StageStageType = "UPLOAD"
)

func (s StageStageType) StageStageType() bool {
    switch s {
        case StageStageType_DOWNLOAD:
            return true
        case StageStageType_UPLOAD:
            return true
        default:
            return false
    }
}






// The ``Message`` class contains information about the errors and warnings if any. **Warning:** This class is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
 type StageMessage struct {
    // Unique identifier of the given operation. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    Id string
    // The time the notification was raised/found. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    Time time.Time
    // The resolution message, if any. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    Resolution *std.LocalizableMessage
    // The notification message. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    Message std.LocalizableMessage
}







// The ``ExecutionStatus`` class contains information about the overall staging operation status. **Warning:** This class is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
 type StageExecutionStatus struct {
    // The staging operation type. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    Type_ StageStageType
    // The overall staging operation status. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    Status StageStageStatus
    // The percent completion of the operation. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    PercentComplete float64
    // The list of warning messages. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    Warnings []StageMessage
    // The list of error messages. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    Errors []StageMessage
}









func stageDownloadInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func stageDownloadOutputType() bindings.BindingType {
    return bindings.NewVoidType()
}

func stageDownloadRestMetadata() protocol.OperationRestMetadata {
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
      "/vcenter/nsx/lcm/stage",
       resultHeaders,
       204,
       errorHeaders,
       map[string]int{"Unauthenticated": 401,"Unauthorized": 403,"NotAllowedInCurrentState": 400,"Error": 500})
}


func stageUploadInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func stageUploadOutputType() bindings.BindingType {
    return bindings.NewVoidType()
}

func stageUploadRestMetadata() protocol.OperationRestMetadata {
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
      "/vcenter/nsx/lcm/stage",
       resultHeaders,
       204,
       errorHeaders,
       map[string]int{"Unauthenticated": 401,"Unauthorized": 403,"NotAllowedInCurrentState": 400,"Error": 500})
}


func stageCancelInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func stageCancelOutputType() bindings.BindingType {
    return bindings.NewVoidType()
}

func stageCancelRestMetadata() protocol.OperationRestMetadata {
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
      "/vcenter/nsx/lcm/stage",
       resultHeaders,
       204,
       errorHeaders,
       map[string]int{"Unauthenticated": 401,"Unauthorized": 403,"NotAllowedInCurrentState": 400,"Error": 500})
}


func stageCheckStatusInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func stageCheckStatusOutputType() bindings.BindingType {
    return bindings.NewReferenceType(StageExecutionStatusBindingType)
}

func stageCheckStatusRestMetadata() protocol.OperationRestMetadata {
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
      "/vcenter/nsx/lcm/stage",
       resultHeaders,
       200,
       errorHeaders,
       map[string]int{"Unauthenticated": 401,"Unauthorized": 403,"Error": 500})
}



func StageMessageBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["id"] = bindings.NewIdType([]string {"com.vmware.vcenter.message"}, "")
    fieldNameMap["id"] = "Id"
    fields["time"] = bindings.NewDateTimeType()
    fieldNameMap["time"] = "Time"
    fields["resolution"] = bindings.NewOptionalType(bindings.NewReferenceType(std.LocalizableMessageBindingType))
    fieldNameMap["resolution"] = "Resolution"
    fields["message"] = bindings.NewReferenceType(std.LocalizableMessageBindingType)
    fieldNameMap["message"] = "Message"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.nsx.stage.message",fields, reflect.TypeOf(StageMessage{}), fieldNameMap, validators)
}

func StageExecutionStatusBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["type"] = bindings.NewEnumType("com.vmware.vcenter.nsx.stage.stage_type", reflect.TypeOf(StageStageType(StageStageType_DOWNLOAD)))
    fieldNameMap["type"] = "Type_"
    fields["status"] = bindings.NewEnumType("com.vmware.vcenter.nsx.stage.stage_status", reflect.TypeOf(StageStageStatus(StageStageStatus_NOT_STARTED)))
    fieldNameMap["status"] = "Status"
    fields["percent_complete"] = bindings.NewDoubleType()
    fieldNameMap["percent_complete"] = "PercentComplete"
    fields["warnings"] = bindings.NewListType(bindings.NewReferenceType(StageMessageBindingType), reflect.TypeOf([]StageMessage{}))
    fieldNameMap["warnings"] = "Warnings"
    fields["errors"] = bindings.NewListType(bindings.NewReferenceType(StageMessageBindingType), reflect.TypeOf([]StageMessage{}))
    fieldNameMap["errors"] = "Errors"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.nsx.stage.execution_status",fields, reflect.TypeOf(StageExecutionStatus{}), fieldNameMap, validators)
}


