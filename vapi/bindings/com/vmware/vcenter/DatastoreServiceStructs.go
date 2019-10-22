/* Copyright © 2019 VMware, Inc. All Rights Reserved.
     SPDX-License-Identifier: BSD-2-Clause */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Data type definitions file for service: Datastore.
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

// The resource type for the vCenter datastore
const Datastore_RESOURCE_TYPE = "Datastore"


// The ``Type`` enumeration class defines the supported types of vCenter datastores.
//
// <p> See {@link com.vmware.vapi.bindings.ApiEnumeration enumerated types description}.
 
type Datastore_Type string

const (
    // VMware File System (ESX Server only).
     Datastore_Type_VMFS Datastore_Type = "VMFS"
    // Network file system v3 (linux & esx servers only).
     Datastore_Type_NFS Datastore_Type = "NFS"
    // Network file system v4.1 (linux & esx servers only).
     Datastore_Type_NFS41 Datastore_Type = "NFS41"
    // Common Internet File System.
     Datastore_Type_CIFS Datastore_Type = "CIFS"
    // Virtual SAN (ESX Server only).
     Datastore_Type_VSAN Datastore_Type = "VSAN"
    // Flash Read Cache (ESX Server only).
     Datastore_Type_VFFS Datastore_Type = "VFFS"
    // vSphere Virtual Volume (ESX Server only).
     Datastore_Type_VVOL Datastore_Type = "VVOL"
)

func (t Datastore_Type) Datastore_Type() bool {
    switch t {
        case Datastore_Type_VMFS:
            return true
        case Datastore_Type_NFS:
            return true
        case Datastore_Type_NFS41:
            return true
        case Datastore_Type_CIFS:
            return true
        case Datastore_Type_VSAN:
            return true
        case Datastore_Type_VFFS:
            return true
        case Datastore_Type_VVOL:
            return true
        default:
            return false
    }
}





// The ``Info`` class contains information about a datastore.
 type DatastoreInfo struct {
    // Name of the datastore.
    Name string
    // Type (VMFS, NFS, NFS41, CIFS, VSAN, VFFS, VVOL) of the datastore.
    Type_ Datastore_Type
    // Whether or not this datastore is accessible.
    Accessible bool
    // Available space of this datastore, in bytes. 
//
//  The server periodically updates this value.
    FreeSpace *int64
    // Whether or not ore than one host in the datacenter has been configured with access to the datastore.
    MultipleHostAccess bool
    // Whether or not the datastore supports thin provisioning on a per file basis. When thin provisioning is used, backing storage is lazily allocated.
    ThinProvisioningSupported bool
}



//


// The ``FilterSpec`` class contains properties used to filter the results when listing datastores (see Datastore#list). If multiple properties are specified, only datastores matching all of the properties match the filter.
 type DatastoreFilterSpec struct {
    // Identifiers of datastores that can match the filter.
    Datastores map[string]bool
    // Names that datastores must have to match the filter (see DatastoreInfo#name).
    Names map[string]bool
    // Types that datastores must have to match the filter (see DatastoreSummary#type).
    Types map[Datastore_Type]bool
    // Folders that must contain the datastore for the datastore to match the filter.
    Folders map[string]bool
    // Datacenters that must contain the datastore for the datastore to match the filter.
    Datacenters map[string]bool
}



//


// The ``Summary`` class contains commonly used information about a datastore.
 type DatastoreSummary struct {
    // Identifier of the datastore.
    Datastore string
    // Name of the datastore.
    Name string
    // Type (VMFS, NFS, NFS41, CIFS, VSAN, VFFS, VVOL) of the datatore.
    Type_ Datastore_Type
    // Available space of this datastore, in bytes. 
//
//  The server periodically updates this value.
    FreeSpace *int64
    // Capacity of this datastore, in bytes. 
//
//  The server periodically updates this value.
    Capacity *int64
}



//





func datastoreGetInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["datastore"] = bindings.NewIdType([]string {"Datastore"}, "")
    fieldNameMap["datastore"] = "Datastore"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func datastoreGetOutputType() bindings.BindingType {
    return bindings.NewReferenceType(DatastoreInfoBindingType)
}

func datastoreGetRestMetadata() protocol.OperationRestMetadata {
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
       map[string]int{"NotFound": 404,"ServiceUnavailable": 503,"Unauthenticated": 401,"Unauthorized": 403})
}


func datastoreListInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["filter"] = bindings.NewOptionalType(bindings.NewReferenceType(DatastoreFilterSpecBindingType))
    fieldNameMap["filter"] = "Filter"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func datastoreListOutputType() bindings.BindingType {
    return bindings.NewListType(bindings.NewReferenceType(DatastoreSummaryBindingType), reflect.TypeOf([]DatastoreSummary{}))
}

func datastoreListRestMetadata() protocol.OperationRestMetadata {
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



func DatastoreInfoBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["name"] = bindings.NewStringType()
    fieldNameMap["name"] = "Name"
    fields["type"] = bindings.NewEnumType("com.vmware.vcenter.datastore.type", reflect.TypeOf(Datastore_Type(Datastore_Type_VMFS)))
    fieldNameMap["type"] = "Type_"
    fields["accessible"] = bindings.NewBooleanType()
    fieldNameMap["accessible"] = "Accessible"
    fields["free_space"] = bindings.NewOptionalType(bindings.NewIntegerType())
    fieldNameMap["free_space"] = "FreeSpace"
    fields["multiple_host_access"] = bindings.NewBooleanType()
    fieldNameMap["multiple_host_access"] = "MultipleHostAccess"
    fields["thin_provisioning_supported"] = bindings.NewBooleanType()
    fieldNameMap["thin_provisioning_supported"] = "ThinProvisioningSupported"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.datastore.info",fields, reflect.TypeOf(DatastoreInfo{}), fieldNameMap, validators)
}

func DatastoreFilterSpecBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["datastores"] = bindings.NewOptionalType(bindings.NewSetType(bindings.NewIdType([]string {"Datastore"}, ""), reflect.TypeOf(map[string]bool{})))
    fieldNameMap["datastores"] = "Datastores"
    fields["names"] = bindings.NewOptionalType(bindings.NewSetType(bindings.NewStringType(), reflect.TypeOf(map[string]bool{})))
    fieldNameMap["names"] = "Names"
    fields["types"] = bindings.NewOptionalType(bindings.NewSetType(bindings.NewEnumType("com.vmware.vcenter.datastore.type", reflect.TypeOf(Datastore_Type(Datastore_Type_VMFS))), reflect.TypeOf(map[Datastore_Type]bool{})))
    fieldNameMap["types"] = "Types"
    fields["folders"] = bindings.NewOptionalType(bindings.NewSetType(bindings.NewIdType([]string {"Folder"}, ""), reflect.TypeOf(map[string]bool{})))
    fieldNameMap["folders"] = "Folders"
    fields["datacenters"] = bindings.NewOptionalType(bindings.NewSetType(bindings.NewIdType([]string {"Datacenter"}, ""), reflect.TypeOf(map[string]bool{})))
    fieldNameMap["datacenters"] = "Datacenters"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.datastore.filter_spec",fields, reflect.TypeOf(DatastoreFilterSpec{}), fieldNameMap, validators)
}

func DatastoreSummaryBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["datastore"] = bindings.NewIdType([]string {"Datastore"}, "")
    fieldNameMap["datastore"] = "Datastore"
    fields["name"] = bindings.NewStringType()
    fieldNameMap["name"] = "Name"
    fields["type"] = bindings.NewEnumType("com.vmware.vcenter.datastore.type", reflect.TypeOf(Datastore_Type(Datastore_Type_VMFS)))
    fieldNameMap["type"] = "Type_"
    fields["free_space"] = bindings.NewOptionalType(bindings.NewIntegerType())
    fieldNameMap["free_space"] = "FreeSpace"
    fields["capacity"] = bindings.NewOptionalType(bindings.NewIntegerType())
    fieldNameMap["capacity"] = "Capacity"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.datastore.summary",fields, reflect.TypeOf(DatastoreSummary{}), fieldNameMap, validators)
}


