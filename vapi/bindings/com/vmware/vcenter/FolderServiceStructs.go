/* Copyright © 2019 VMware, Inc. All Rights Reserved.
     SPDX-License-Identifier: BSD-2-Clause */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Data type definitions file for service: Folder.
 * Includes binding types of a structures and enumerations defined in the service.
 * Shared by client-side stubs and server-side skeletons to ensure type
 * compatibility.
 */

package vcenter

import (
    "reflect"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/bindings"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/data"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/protocol"
)

// The resource type for the vCenter folder
const Folder_RESOURCE_TYPE = "Folder"


// The ``Type`` enumeration class defines the type of a vCenter Server folder. The type of a folder determines what what kinds of children can be contained in the folder.
//
// <p> See {@link com.vmware.vapi.bindings.ApiEnumeration enumerated types description}.
 
type Folder_Type string

const (
    // A folder that can contain datacenters.
     Folder_Type_DATACENTER Folder_Type = "DATACENTER"
    // A folder that can contain datastores.
     Folder_Type_DATASTORE Folder_Type = "DATASTORE"
    // A folder that can contain compute resources (hosts and clusters).
     Folder_Type_HOST Folder_Type = "HOST"
    // A folder that can contain networkds.
     Folder_Type_NETWORK Folder_Type = "NETWORK"
    // A folder that can contain virtual machines.
     Folder_Type_VIRTUAL_MACHINE Folder_Type = "VIRTUAL_MACHINE"
)

func (t Folder_Type) Folder_Type() bool {
    switch t {
        case Folder_Type_DATACENTER:
            return true
        case Folder_Type_DATASTORE:
            return true
        case Folder_Type_HOST:
            return true
        case Folder_Type_NETWORK:
            return true
        case Folder_Type_VIRTUAL_MACHINE:
            return true
        default:
            return false
    }
}





// The ``FilterSpec`` class contains properties used to filter the results when listing folders (see Folder#list). If multiple properties are specified, only folders matching all of the properties match the filter.
 type FolderFilterSpec struct {
    // Identifiers of folders that can match the filter.
    Folders map[string]bool
    // Names that folders must have to match the filter (see FolderSummary#name).
    Names map[string]bool
    // Type that folders must have to match the filter (see FolderSummary#type).
    Type_ *Folder_Type
    // Folders that must contain the folder for the folder to match the filter.
    ParentFolders map[string]bool
    // Datacenters that must contain the folder for the folder to match the filter.
    Datacenters map[string]bool
}



//


// The ``Summary`` class contains commonly used information about a folder.
 type FolderSummary struct {
    // Identifier of the folder.
    Folder string
    // Name of the vCenter Server folder.
    Name string
    // Type (DATACENTER, DATASTORE, HOST, NETWORK, VIRTUAL_MACHINE) of the vCenter Server folder.
    Type_ Folder_Type
}



//





func folderListInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["filter"] = bindings.NewOptionalType(bindings.NewReferenceType(FolderFilterSpecBindingType))
    fieldNameMap["filter"] = "Filter"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func folderListOutputType() bindings.BindingType {
    return bindings.NewListType(bindings.NewReferenceType(FolderSummaryBindingType), reflect.TypeOf([]FolderSummary{}))
}

func folderListRestMetadata() protocol.OperationRestMetadata {
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
       map[string]int{"InvalidArgument": 400,"UnableToAllocateResource": 500,"ServiceUnavailable": 503,"Unauthenticated": 401,"Unauthorized": 403})
}



func FolderFilterSpecBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["folders"] = bindings.NewOptionalType(bindings.NewSetType(bindings.NewIdType([]string {"Folder"}, ""), reflect.TypeOf(map[string]bool{})))
    fieldNameMap["folders"] = "Folders"
    fields["names"] = bindings.NewOptionalType(bindings.NewSetType(bindings.NewStringType(), reflect.TypeOf(map[string]bool{})))
    fieldNameMap["names"] = "Names"
    fields["type"] = bindings.NewOptionalType(bindings.NewEnumType("com.vmware.vcenter.folder.type", reflect.TypeOf(Folder_Type(Folder_Type_DATACENTER))))
    fieldNameMap["type"] = "Type_"
    fields["parent_folders"] = bindings.NewOptionalType(bindings.NewSetType(bindings.NewIdType([]string {"Folder"}, ""), reflect.TypeOf(map[string]bool{})))
    fieldNameMap["parent_folders"] = "ParentFolders"
    fields["datacenters"] = bindings.NewOptionalType(bindings.NewSetType(bindings.NewIdType([]string {"Datacenter"}, ""), reflect.TypeOf(map[string]bool{})))
    fieldNameMap["datacenters"] = "Datacenters"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.folder.filter_spec",fields, reflect.TypeOf(FolderFilterSpec{}), fieldNameMap, validators)
}

func FolderSummaryBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["folder"] = bindings.NewIdType([]string {"Folder"}, "")
    fieldNameMap["folder"] = "Folder"
    fields["name"] = bindings.NewStringType()
    fieldNameMap["name"] = "Name"
    fields["type"] = bindings.NewEnumType("com.vmware.vcenter.folder.type", reflect.TypeOf(Folder_Type(Folder_Type_DATACENTER)))
    fieldNameMap["type"] = "Type_"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.folder.summary",fields, reflect.TypeOf(FolderSummary{}), fieldNameMap, validators)
}


