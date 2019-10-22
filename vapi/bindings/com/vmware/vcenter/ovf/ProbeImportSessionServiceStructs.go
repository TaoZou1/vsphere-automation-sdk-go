/* Copyright © 2019 VMware, Inc. All Rights Reserved.
     SPDX-License-Identifier: BSD-2-Clause */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Data type definitions file for service: ProbeImportSession.
 * Includes binding types of a structures and enumerations defined in the service.
 * Shared by client-side stubs and server-side skeletons to ensure type
 * compatibility.
 */

package ovf

import (
    "reflect"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/bindings"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/data"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/protocol"
)



// The ``State`` enumeration class defines the states of ProbeImportSession.
//
// <p> See {@link com.vmware.vapi.bindings.ApiEnumeration enumerated types description}.
 
type ProbeImportSession_State string

const (
    // State of an import transfer that does not have any files available. The transfer needs the OVF descriptor to continue. If this is a PUSH transfer, the client must upload the OVF descriptor, and the transfer file list has one file info entry with a URL to which the client must upload the OVF descriptor using HTTP PUT. For pull transfers (including content library), the server is in the process of retrieving the OVF descriptor. 
    //
    //  Transition to the next state is done when the server has retrieved the complete OVF content and parsed it.
     ProbeImportSession_State_PROBE_IMPORT_OVF_TRANSFER ProbeImportSession_State = "PROBE_IMPORT_OVF_TRANSFER"
    // The file list contains a number of message bundles that need to be transferred to the server. If this is a PUSH transfer, the client must PUT the requested files to the server. 
    //
    //  In case the OVF descriptor does not specify any bundles this state is skipped. 
    //
    //  Transition to next state is done when the complete content of all message bundles has been retrieved by the server.
     ProbeImportSession_State_PROBE_IMPORT_MSG_BUNDLES_TRANSFER ProbeImportSession_State = "PROBE_IMPORT_MSG_BUNDLES_TRANSFER"
    // The server can be queried for OVF parameters, and the client can specify instantiation parameters. 
    //
    //  Specifying an OVF instantiation parameter might affect other OVF instantiation parameters.
     ProbeImportSession_State_PROBE_IMPORT_SELECTING_OVF_PARAMS ProbeImportSession_State = "PROBE_IMPORT_SELECTING_OVF_PARAMS"
    // The transfer failed.
     ProbeImportSession_State_PROBE_IMPORT_ERROR ProbeImportSession_State = "PROBE_IMPORT_ERROR"
)

func (s ProbeImportSession_State) ProbeImportSession_State() bool {
    switch s {
        case ProbeImportSession_State_PROBE_IMPORT_OVF_TRANSFER:
            return true
        case ProbeImportSession_State_PROBE_IMPORT_MSG_BUNDLES_TRANSFER:
            return true
        case ProbeImportSession_State_PROBE_IMPORT_SELECTING_OVF_PARAMS:
            return true
        case ProbeImportSession_State_PROBE_IMPORT_ERROR:
            return true
        default:
            return false
    }
}





// The ``Info`` class represents an import session.
 type ProbeImportSessionInfo struct {
    // The state for the current transfer.
    State ProbeImportSession_State
    // List of files that are part of this transfer. 
//
//  There is an entry for the OVF descriptor and optional manifest and certificate.
    Files []OvfFileInfo
    // List of errors.
    Errors []OvfError
    // List of warnings.
    Warnings []OvfWarning
    // List of information.
    Information []OvfInfo
}


// TODO some error doesn't have ErrorType. Resolve this later VAPI-3009
//





func probeImportSessionCreateProbeImportSessionInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["client_token"] = bindings.NewOptionalType(bindings.NewStringType())
    fields["create_spec"] = bindings.NewDynamicStructType([]bindings.ReferenceType {bindings.NewReferenceType(ImportSessionCreateSpecBindingType),}, bindings.JSONRPC)
    fieldNameMap["client_token"] = "ClientToken"
    fieldNameMap["create_spec"] = "CreateSpec"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func probeImportSessionCreateProbeImportSessionOutputType() bindings.BindingType {
    return bindings.NewIdType([]string {"com.vmware.vcenter.OvfProbeImportSession"}, "")
}

func probeImportSessionCreateProbeImportSessionRestMetadata() protocol.OperationRestMetadata {
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
       map[string]int{"InvalidArgument": 400,"NotFound": 404,"ResourceInaccessible": 500})
}


func probeImportSessionGetInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["id"] = bindings.NewIdType([]string {"com.vmware.vcenter.OvfProbeImportSession"}, "")
    fieldNameMap["id"] = "Id"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func probeImportSessionGetOutputType() bindings.BindingType {
    return bindings.NewReferenceType(ProbeImportSessionInfoBindingType)
}

func probeImportSessionGetRestMetadata() protocol.OperationRestMetadata {
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
       map[string]int{"NotFound": 404})
}


func probeImportSessionTryInstantiateInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["id"] = bindings.NewIdType([]string {"com.vmware.vcenter.OvfProbeImportSession"}, "")
    fields["instantiation_parameters"] = bindings.NewListType(bindings.NewDynamicStructType([]bindings.ReferenceType {bindings.NewReferenceType(OvfParamsBindingType),}, bindings.JSONRPC), reflect.TypeOf([]*data.StructValue{}))
    fieldNameMap["id"] = "Id"
    fieldNameMap["instantiation_parameters"] = "InstantiationParameters"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func probeImportSessionTryInstantiateOutputType() bindings.BindingType {
    return bindings.NewReferenceType(ImportSessionOvfValidationResultBindingType)
}

func probeImportSessionTryInstantiateRestMetadata() protocol.OperationRestMetadata {
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
       map[string]int{"NotFound": 404,"NotAllowedInCurrentState": 400,"Unauthorized": 403})
}


func probeImportSessionDeleteInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["id"] = bindings.NewIdType([]string {"com.vmware.vcenter.OvfProbeImportSession"}, "")
    fieldNameMap["id"] = "Id"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func probeImportSessionDeleteOutputType() bindings.BindingType {
    return bindings.NewVoidType()
}

func probeImportSessionDeleteRestMetadata() protocol.OperationRestMetadata {
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
       map[string]int{})
}



func ProbeImportSessionInfoBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["state"] = bindings.NewEnumType("com.vmware.vcenter.ovf.probe_import_session.state", reflect.TypeOf(ProbeImportSession_State(ProbeImportSession_State_PROBE_IMPORT_OVF_TRANSFER)))
    fieldNameMap["state"] = "State"
    fields["files"] = bindings.NewListType(bindings.NewReferenceType(OvfFileInfoBindingType), reflect.TypeOf([]OvfFileInfo{}))
    fieldNameMap["files"] = "Files"
    fields["errors"] = bindings.NewListType(bindings.NewReferenceType(OvfErrorBindingType), reflect.TypeOf([]OvfError{}))
    fieldNameMap["errors"] = "Errors"
    fields["warnings"] = bindings.NewListType(bindings.NewReferenceType(OvfWarningBindingType), reflect.TypeOf([]OvfWarning{}))
    fieldNameMap["warnings"] = "Warnings"
    fields["information"] = bindings.NewListType(bindings.NewReferenceType(OvfInfoBindingType), reflect.TypeOf([]OvfInfo{}))
    fieldNameMap["information"] = "Information"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.ovf.probe_import_session.info",fields, reflect.TypeOf(ProbeImportSessionInfo{}), fieldNameMap, validators)
}


