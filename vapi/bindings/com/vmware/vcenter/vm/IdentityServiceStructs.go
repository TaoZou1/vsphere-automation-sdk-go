/* Copyright © 2019 VMware, Inc. All Rights Reserved.
     SPDX-License-Identifier: BSD-2-Clause */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Data type definitions file for service: Identity.
 * Includes binding types of a structures and enumerations defined in the service.
 * Shared by client-side stubs and server-side skeletons to ensure type
 * compatibility.
 */

package vm

import (
    "reflect"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/bindings"
)




// The ``Info`` class contains information about the identity of a virtual machine.
 type IdentityInfo struct {
    // Virtual machine name.
    Name string
    // 128-bit SMBIOS UUID of a virtual machine represented as a hexadecimal string in "12345678-abcd-1234-cdef-123456789abc" format.
    BiosUuid string
    // VirtualCenter-specific 128-bit UUID of a virtual machine, represented as a hexademical string. This identifier is used by VirtualCenter to uniquely identify all virtual machine instances, including those that may share the same SMBIOS UUID.
    InstanceUuid string
}



func (IdentityInfo IdentityInfo) Error() string {
    return "com.vmware.vcenter.vm.info"
}







func IdentityInfoBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["name"] = bindings.NewStringType()
    fieldNameMap["name"] = "Name"
    fields["bios_uuid"] = bindings.NewStringType()
    fieldNameMap["bios_uuid"] = "BiosUuid"
    fields["instance_uuid"] = bindings.NewStringType()
    fieldNameMap["instance_uuid"] = "InstanceUuid"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.vm.identity.info",fields, reflect.TypeOf(IdentityInfo{}), fieldNameMap, validators)
}


