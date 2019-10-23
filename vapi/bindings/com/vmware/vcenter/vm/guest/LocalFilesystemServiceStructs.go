/* Copyright © 2019 VMware, Inc. All Rights Reserved.
     SPDX-License-Identifier: BSD-2-Clause */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Data type definitions file for service: LocalFilesystem.
 * Includes binding types of a structures and enumerations defined in the service.
 * Shared by client-side stubs and server-side skeletons to ensure type
 * compatibility.
 */

package guest

import (
    "reflect"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/bindings"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/data"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/protocol"
)




// Describes the virtual disk backing a local guest disk.
 type LocalFilesystemVirtualDiskMapping struct {
    // The virtual disk.
    Disk string
}



func (LocalFilesystemVirtualDiskMapping LocalFilesystemVirtualDiskMapping) Error() string {
    return "com.vmware.vcenter.vm.guest.virtual_disk_mapping"
}



// The ``Info`` class contains information about a local file system configured in the guest operating system.
 type LocalFilesystemInfo struct {
    // Total capacity of the file system, in bytes.
    Capacity int64
    // Free space on the file system, in bytes.
    FreeSpace int64
    // Filesystem type, if known. For example, ext3 or NTFS.
    Filesystem *string
    // VirtualDisks backing the guest partition, if known.
    Mappings []LocalFilesystemVirtualDiskMapping
}



func (LocalFilesystemInfo LocalFilesystemInfo) Error() string {
    return "com.vmware.vcenter.vm.guest.info"
}






func localFilesystemGetInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["vm"] = bindings.NewIdType([]string {"VirtualMachine"}, "")
    fieldNameMap["vm"] = "Vm"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func localFilesystemGetOutputType() bindings.BindingType {
    return bindings.NewMapType(bindings.NewStringType(), bindings.NewReferenceType(LocalFilesystemInfoBindingType),reflect.TypeOf(map[string]LocalFilesystemInfo{}))
}

func localFilesystemGetRestMetadata() protocol.OperationRestMetadata {
    paramsTypeMap := map[string]bindings.BindingType{}
    pathParams := map[string]string{}
    queryParams := map[string]string{}
    headerParams := map[string]string{}
    paramsTypeMap["vm"] = bindings.NewIdType([]string {"VirtualMachine"}, "")
    paramsTypeMap["vm"] = bindings.NewIdType([]string {"VirtualMachine"}, "")
    pathParams["vm"] = "vm"
    resultHeaders := map[string]string{}
    errorHeaders := map[string]string{}
    return protocol.NewOperationRestMetadata(
      paramsTypeMap,
      pathParams,
      queryParams,
      headerParams,
      "",
      "GET",
      "/vcenter/vm/{vm}/guest/local-filesystem",
       resultHeaders,
       200,
       errorHeaders,
       map[string]int{"Error": 500,"NotFound": 404,"ServiceUnavailable": 503})
}



func LocalFilesystemVirtualDiskMappingBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["disk"] = bindings.NewIdType([]string {"com.vmware.vcenter.vm.hardware.Disk"}, "")
    fieldNameMap["disk"] = "Disk"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.vm.guest.local_filesystem.virtual_disk_mapping",fields, reflect.TypeOf(LocalFilesystemVirtualDiskMapping{}), fieldNameMap, validators)
}

func LocalFilesystemInfoBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["capacity"] = bindings.NewIntegerType()
    fieldNameMap["capacity"] = "Capacity"
    fields["free_space"] = bindings.NewIntegerType()
    fieldNameMap["free_space"] = "FreeSpace"
    fields["filesystem"] = bindings.NewOptionalType(bindings.NewStringType())
    fieldNameMap["filesystem"] = "Filesystem"
    fields["mappings"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewReferenceType(LocalFilesystemVirtualDiskMappingBindingType), reflect.TypeOf([]LocalFilesystemVirtualDiskMapping{})))
    fieldNameMap["mappings"] = "Mappings"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.vm.guest.local_filesystem.info",fields, reflect.TypeOf(LocalFilesystemInfo{}), fieldNameMap, validators)
}


