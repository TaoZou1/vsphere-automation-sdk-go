/* Copyright © 2019 VMware, Inc. All Rights Reserved.
     SPDX-License-Identifier: BSD-2-Clause */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Data type definitions file for service: ExportSession.
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



// The ``State`` enumeration class defines the different states for an export session.
//
// <p> See {@link com.vmware.vapi.bindings.ApiEnumeration enumerated types description}.
 
type ExportSession_State string

const (
    // Files are being made available for download. During this state entries will be added to the file list as the files become ready.
     ExportSession_State_EXPORT_PREPARING ExportSession_State = "EXPORT_PREPARING"
    // All files are available for download, and the file list is complete. This state is used when the files are being downloaded by the client using HTTP GETs.
     ExportSession_State_EXPORT_READY ExportSession_State = "EXPORT_READY"
    // Files are in the process of being transferred. This state is used when the target is a content library.
     ExportSession_State_EXPORT_IN_PROGRESS ExportSession_State = "EXPORT_IN_PROGRESS"
    // Files have been transferred. This state is used when the target is a content library.
     ExportSession_State_EXPORT_COMPLETED ExportSession_State = "EXPORT_COMPLETED"
    // The transfer failed.
     ExportSession_State_EXPORT_ERROR ExportSession_State = "EXPORT_ERROR"
)

func (s ExportSession_State) ExportSession_State() bool {
    switch s {
        case ExportSession_State_EXPORT_PREPARING:
            return true
        case ExportSession_State_EXPORT_READY:
            return true
        case ExportSession_State_EXPORT_IN_PROGRESS:
            return true
        case ExportSession_State_EXPORT_COMPLETED:
            return true
        case ExportSession_State_EXPORT_ERROR:
            return true
        default:
            return false
    }
}




// The ``TargetType`` enumeration class defines the target types of export session. If ExportSessionCreateSpec#contentLibrary is provided, the OVF is exported to a local content library.
//
// <p> See {@link com.vmware.vapi.bindings.ApiEnumeration enumerated types description}.
 
type ExportSession_TargetType string

const (
    // A set of URLs are published where the files can be downloaded from.
     ExportSession_TargetType_DOWNLOAD_TARGET ExportSession_TargetType = "DOWNLOAD_TARGET"
    // The OVF package is directly transferred to a content library.
     ExportSession_TargetType_CONTENT_LIBRARY_TARGET ExportSession_TargetType = "CONTENT_LIBRARY_TARGET"
)

func (t ExportSession_TargetType) ExportSession_TargetType() bool {
    switch t {
        case ExportSession_TargetType_DOWNLOAD_TARGET:
            return true
        case ExportSession_TargetType_CONTENT_LIBRARY_TARGET:
            return true
        default:
            return false
    }
}




// The ``TargetContentType`` enumeration class defines the target content types of export session. This only applies to ExportSession_TargetType#ExportSessionTargetType_DOWNLOAD_TARGET.
//
// <p> See {@link com.vmware.vapi.bindings.ApiEnumeration enumerated types description}.
 
type ExportSession_TargetContentType string

const (
    // Export the OVF package as a set of files.
     ExportSession_TargetContentType_OVF_TARGET ExportSession_TargetContentType = "OVF_TARGET"
    // Export the OVF package as a single OVA file.
     ExportSession_TargetContentType_OVA_TARGET ExportSession_TargetContentType = "OVA_TARGET"
)

func (t ExportSession_TargetContentType) ExportSession_TargetContentType() bool {
    switch t {
        case ExportSession_TargetContentType_OVF_TARGET:
            return true
        case ExportSession_TargetContentType_OVA_TARGET:
            return true
        default:
            return false
    }
}





// The ``SourceInfo`` class specifies either a virtual machine or a virtual appliance.
 type ExportSessionSourceInfo struct {
    // The identifier of the virtual appliance to export.
    VappId *string
    // The identifier of the virtual machine to export.
    VmId *string
}


// TODO some error doesn't have ErrorType. Resolve this later VAPI-3009
//


// The ``CreateSpec`` class contains export parameters.
 type ExportSessionCreateSpec struct {
    // The download target type.
    TargetType ExportSession_TargetType
    // The download target content type. Default is ExportSession_TargetContentType#ExportSessionTargetContentType_OVF_TARGET.
    TargetContentType *ExportSession_TargetContentType
    // List of selected export flags. The supported flags are: 
//
// * BIOS_UUID: Include BIOS UUIDs.
// * IP: Include IP settings.
// * CDROM: Include attached CDROM media.
// * Other supported flags can be obtained using ExportFlag#list.
    ExportFlags []string
    // If non-empty, this specifies the name in the OVF descriptor.
    Name *string
    // If non-empty, this specifies the description to put in the OVF descriptor.
    Description *string
    // The identifier of the content library.
    ContentLibrary *string
    // The identifier of the content library item to update. If omitted, a new item will be created.
    ContentLibraryItem *string
}


// TODO some error doesn't have ErrorType. Resolve this later VAPI-3009
//


// The ``Info`` class represents an export session.
 type ExportSessionInfo struct {
    // The state for the current export.
    State ExportSession_State
    // Progress of export. Only set if target is ExportSession_TargetType#ExportSessionTargetType_CONTENT_LIBRARY_TARGET.
    Progress *int64
    // List of files that are part of this export. 
//
//  There is an entry for the OVF descriptor, files defined in the file reference section of the OVF descriptor, and optional manifest and certificate.
    Files []OvfFileInfo
    // List of errors.
    Errors []OvfError
    // List of warnings.
    Warnings []OvfWarning
    // List of information.
    Information []OvfInfo
    // This is the identifier of the content library item that is created in case the target type is ExportSession_TargetType#ExportSessionTargetType_CONTENT_LIBRARY_TARGET.
    LibraryItemId *string
}


// TODO some error doesn't have ErrorType. Resolve this later VAPI-3009
//


// The ``PreviewSpec`` class contains information about preview parameters.
 type ExportSessionPreviewSpec struct {
    // List of selected export flags. These are the same as when creating an export session.
    ExportFlags []string
    // The expected target type, the default type is ExportSession_TargetType#ExportSessionTargetType_DOWNLOAD_TARGET.
    TargetType *ExportSession_TargetType
}


// TODO some error doesn't have ErrorType. Resolve this later VAPI-3009
//


// The ``PreviewFile`` class contains information about a file that will be exported according to a preview.
 type ExportSessionPreviewFile struct {
    // The name of the file.
    Name string
}


// TODO some error doesn't have ErrorType. Resolve this later VAPI-3009
//


// The ``Preview`` class contains information about the result of an export preview.
 type ExportSessionPreview struct {
    // The files that will need to be transferred during the export. This includes disks, CD images, and the descriptor itself.
    Files []ExportSessionPreviewFile
}


// TODO some error doesn't have ErrorType. Resolve this later VAPI-3009
//





func exportSessionCreateInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["client_token"] = bindings.NewOptionalType(bindings.NewStringType())
    fields["source"] = bindings.NewReferenceType(ExportSessionSourceInfoBindingType)
    fields["create_spec"] = bindings.NewDynamicStructType([]bindings.ReferenceType {bindings.NewReferenceType(ExportSessionCreateSpecBindingType),}, bindings.JSONRPC)
    fieldNameMap["client_token"] = "ClientToken"
    fieldNameMap["source"] = "Source"
    fieldNameMap["create_spec"] = "CreateSpec"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func exportSessionCreateOutputType() bindings.BindingType {
    return bindings.NewIdType([]string {"com.vmware.vcenter.OvfExportSession"}, "")
}

func exportSessionCreateRestMetadata() protocol.OperationRestMetadata {
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
       map[string]int{"NotFound": 404,"InvalidArgument": 400,"NotAllowedInCurrentState": 400,"ResourceInaccessible": 500,"ResourceBusy": 500})
}


func exportSessionGetInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["id"] = bindings.NewIdType([]string {"com.vmware.vcenter.OvfExportSession"}, "")
    fieldNameMap["id"] = "Id"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func exportSessionGetOutputType() bindings.BindingType {
    return bindings.NewReferenceType(ExportSessionInfoBindingType)
}

func exportSessionGetRestMetadata() protocol.OperationRestMetadata {
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


func exportSessionDeleteInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["id"] = bindings.NewIdType([]string {"com.vmware.vcenter.OvfExportSession"}, "")
    fieldNameMap["id"] = "Id"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func exportSessionDeleteOutputType() bindings.BindingType {
    return bindings.NewVoidType()
}

func exportSessionDeleteRestMetadata() protocol.OperationRestMetadata {
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


func exportSessionProgressInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["id"] = bindings.NewIdType([]string {"com.vmware.vcenter.OvfExportSession"}, "")
    fields["percent"] = bindings.NewIntegerType()
    fieldNameMap["id"] = "Id"
    fieldNameMap["percent"] = "Percent"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func exportSessionProgressOutputType() bindings.BindingType {
    return bindings.NewVoidType()
}

func exportSessionProgressRestMetadata() protocol.OperationRestMetadata {
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
       map[string]int{"NotFound": 404,"InvalidArgument": 400,"NotAllowedInCurrentState": 400})
}


func exportSessionPreviewInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["source"] = bindings.NewReferenceType(ExportSessionSourceInfoBindingType)
    fields["preview_spec"] = bindings.NewDynamicStructType([]bindings.ReferenceType {bindings.NewReferenceType(ExportSessionPreviewSpecBindingType),}, bindings.JSONRPC)
    fieldNameMap["source"] = "Source"
    fieldNameMap["preview_spec"] = "PreviewSpec"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func exportSessionPreviewOutputType() bindings.BindingType {
    return bindings.NewReferenceType(ExportSessionPreviewBindingType)
}

func exportSessionPreviewRestMetadata() protocol.OperationRestMetadata {
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
       map[string]int{"NotFound": 404,"NotAllowedInCurrentState": 400,"ResourceInaccessible": 500})
}



func ExportSessionSourceInfoBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["vapp_id"] = bindings.NewOptionalType(bindings.NewIdType([]string {"VirtualApp"}, ""))
    fieldNameMap["vapp_id"] = "VappId"
    fields["vm_id"] = bindings.NewOptionalType(bindings.NewIdType([]string {"VirtualMachine"}, ""))
    fieldNameMap["vm_id"] = "VmId"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.ovf.export_session.source_info",fields, reflect.TypeOf(ExportSessionSourceInfo{}), fieldNameMap, validators)
}

func ExportSessionCreateSpecBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["target_type"] = bindings.NewEnumType("com.vmware.vcenter.ovf.export_session.target_type", reflect.TypeOf(ExportSession_TargetType(ExportSession_TargetType_DOWNLOAD_TARGET)))
    fieldNameMap["target_type"] = "TargetType"
    fields["target_content_type"] = bindings.NewOptionalType(bindings.NewEnumType("com.vmware.vcenter.ovf.export_session.target_content_type", reflect.TypeOf(ExportSession_TargetContentType(ExportSession_TargetContentType_OVF_TARGET))))
    fieldNameMap["target_content_type"] = "TargetContentType"
    fields["export_flags"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewStringType(), reflect.TypeOf([]string{})))
    fieldNameMap["export_flags"] = "ExportFlags"
    fields["name"] = bindings.NewOptionalType(bindings.NewStringType())
    fieldNameMap["name"] = "Name"
    fields["description"] = bindings.NewOptionalType(bindings.NewStringType())
    fieldNameMap["description"] = "Description"
    fields["content_library"] = bindings.NewOptionalType(bindings.NewIdType([]string {"com.vmware.content.Library"}, ""))
    fieldNameMap["content_library"] = "ContentLibrary"
    fields["content_library_item"] = bindings.NewOptionalType(bindings.NewIdType([]string {"com.vmware.content.library.Item"}, ""))
    fieldNameMap["content_library_item"] = "ContentLibraryItem"
    var validators = []bindings.Validator{}
    uv1 := bindings.NewUnionValidator("target_type",
        map[string][]bindings.FieldData {
            "DOWNLOAD_TARGET": []bindings.FieldData {
                 bindings.NewFieldData("target_content_type", true),
            },
            "CONTENT_LIBRARY_TARGET": []bindings.FieldData {
                 bindings.NewFieldData("content_library", true),
            },
        },
    )
    validators = append(validators, uv1)
    uv2 := bindings.NewUnionValidator("target_content_type",
        map[string][]bindings.FieldData {
            "OVF_TARGET": []bindings.FieldData {},
            "OVA_TARGET": []bindings.FieldData {},
        },
    )
    validators = append(validators, uv2)
    return bindings.NewStructType("com.vmware.vcenter.ovf.export_session.create_spec",fields, reflect.TypeOf(ExportSessionCreateSpec{}), fieldNameMap, validators)
}

func ExportSessionInfoBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["state"] = bindings.NewEnumType("com.vmware.vcenter.ovf.export_session.state", reflect.TypeOf(ExportSession_State(ExportSession_State_EXPORT_PREPARING)))
    fieldNameMap["state"] = "State"
    fields["progress"] = bindings.NewOptionalType(bindings.NewIntegerType())
    fieldNameMap["progress"] = "Progress"
    fields["files"] = bindings.NewListType(bindings.NewReferenceType(OvfFileInfoBindingType), reflect.TypeOf([]OvfFileInfo{}))
    fieldNameMap["files"] = "Files"
    fields["errors"] = bindings.NewListType(bindings.NewReferenceType(OvfErrorBindingType), reflect.TypeOf([]OvfError{}))
    fieldNameMap["errors"] = "Errors"
    fields["warnings"] = bindings.NewListType(bindings.NewReferenceType(OvfWarningBindingType), reflect.TypeOf([]OvfWarning{}))
    fieldNameMap["warnings"] = "Warnings"
    fields["information"] = bindings.NewListType(bindings.NewReferenceType(OvfInfoBindingType), reflect.TypeOf([]OvfInfo{}))
    fieldNameMap["information"] = "Information"
    fields["library_item_id"] = bindings.NewOptionalType(bindings.NewIdType([]string {"com.vmware.content.library.Item"}, ""))
    fieldNameMap["library_item_id"] = "LibraryItemId"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.ovf.export_session.info",fields, reflect.TypeOf(ExportSessionInfo{}), fieldNameMap, validators)
}

func ExportSessionPreviewSpecBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["export_flags"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewStringType(), reflect.TypeOf([]string{})))
    fieldNameMap["export_flags"] = "ExportFlags"
    fields["target_type"] = bindings.NewOptionalType(bindings.NewEnumType("com.vmware.vcenter.ovf.export_session.target_type", reflect.TypeOf(ExportSession_TargetType(ExportSession_TargetType_DOWNLOAD_TARGET))))
    fieldNameMap["target_type"] = "TargetType"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.ovf.export_session.preview_spec",fields, reflect.TypeOf(ExportSessionPreviewSpec{}), fieldNameMap, validators)
}

func ExportSessionPreviewFileBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["name"] = bindings.NewStringType()
    fieldNameMap["name"] = "Name"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.ovf.export_session.preview_file",fields, reflect.TypeOf(ExportSessionPreviewFile{}), fieldNameMap, validators)
}

func ExportSessionPreviewBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["files"] = bindings.NewListType(bindings.NewReferenceType(ExportSessionPreviewFileBindingType), reflect.TypeOf([]ExportSessionPreviewFile{}))
    fieldNameMap["files"] = "Files"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.ovf.export_session.preview",fields, reflect.TypeOf(ExportSessionPreview{}), fieldNameMap, validators)
}


