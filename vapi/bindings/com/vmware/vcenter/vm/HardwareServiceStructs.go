/* Copyright © 2019 VMware, Inc. All Rights Reserved.
     SPDX-License-Identifier: BSD-2-Clause */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Data type definitions file for service: Hardware.
 * Includes binding types of a structures and enumerations defined in the service.
 * Shared by client-side stubs and server-side skeletons to ensure type
 * compatibility.
 */

package vm

import (
    "reflect"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/bindings"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/data"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/protocol"
)



// The ``Version`` enumeration class defines the valid virtual hardware versions for a virtual machine. See https://kb.vmware.com/s/article/1003746 (Virtual machine hardware versions (1003746)).
//
// <p> See {@link com.vmware.vapi.bindings.ApiEnumeration enumerated types description}.
 
type Hardware_Version string

const (
    // Hardware version 3, first supported in ESXi 2.5.
     Hardware_Version_VMX_03 Hardware_Version = "VMX_03"
    // Hardware version 4, first supported in ESXi 3.0.
     Hardware_Version_VMX_04 Hardware_Version = "VMX_04"
    // Hardware version 6, first supported in WS 6.0.
     Hardware_Version_VMX_06 Hardware_Version = "VMX_06"
    // Hardware version 7, first supported in ESXi 4.0.
     Hardware_Version_VMX_07 Hardware_Version = "VMX_07"
    // Hardware version 8, first supported in ESXi 5.0.
     Hardware_Version_VMX_08 Hardware_Version = "VMX_08"
    // Hardware version 9, first supported in ESXi 5.1.
     Hardware_Version_VMX_09 Hardware_Version = "VMX_09"
    // Hardware version 10, first supported in ESXi 5.5.
     Hardware_Version_VMX_10 Hardware_Version = "VMX_10"
    // Hardware version 11, first supported in ESXi 6.0.
     Hardware_Version_VMX_11 Hardware_Version = "VMX_11"
    // Hardware version 12, first supported in Workstation 12.0.
     Hardware_Version_VMX_12 Hardware_Version = "VMX_12"
    // Hardware version 13, first supported in ESXi 6.5.
     Hardware_Version_VMX_13 Hardware_Version = "VMX_13"
    // Hardware version 14, first supported in ESXi 6.7. This constant field was added in vSphere API 6.7.
     Hardware_Version_VMX_14 Hardware_Version = "VMX_14"
    // Hardware version 15, first supported in ESXi 6.7.0 Update 2. This constant field was added in vSphere API 6.7.2.
     Hardware_Version_VMX_15 Hardware_Version = "VMX_15"
    // Hardware version 16, first supported in Workstation 15.0. This constant field was added in vSphere API 7.0.
     Hardware_Version_VMX_16 Hardware_Version = "VMX_16"
    // Hardware version 17, first supported in ESX 7.0. This constant field was added in vSphere API 7.0.
     Hardware_Version_VMX_17 Hardware_Version = "VMX_17"
    // Future hardware version, not supported yet. **Warning:** This constant field is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
     Hardware_Version_VMX_FUTURE Hardware_Version = "VMX_FUTURE"
)

func (v Hardware_Version) Hardware_Version() bool {
    switch v {
        case Hardware_Version_VMX_03:
            return true
        case Hardware_Version_VMX_04:
            return true
        case Hardware_Version_VMX_06:
            return true
        case Hardware_Version_VMX_07:
            return true
        case Hardware_Version_VMX_08:
            return true
        case Hardware_Version_VMX_09:
            return true
        case Hardware_Version_VMX_10:
            return true
        case Hardware_Version_VMX_11:
            return true
        case Hardware_Version_VMX_12:
            return true
        case Hardware_Version_VMX_13:
            return true
        case Hardware_Version_VMX_14:
            return true
        case Hardware_Version_VMX_15:
            return true
        case Hardware_Version_VMX_16:
            return true
        case Hardware_Version_VMX_17:
            return true
        case Hardware_Version_VMX_FUTURE:
            return true
        default:
            return false
    }
}




// The ``UpgradePolicy`` enumeration class defines the valid virtual hardware upgrade policies for a virtual machine.
//
// <p> See {@link com.vmware.vapi.bindings.ApiEnumeration enumerated types description}.
 
type Hardware_UpgradePolicy string

const (
    // Do not upgrade the virtual machine when it is powered on.
     Hardware_UpgradePolicy_NEVER Hardware_UpgradePolicy = "NEVER"
    // Run scheduled upgrade when the virtual machine is powered on after a clean shutdown of the guest operating system.
     Hardware_UpgradePolicy_AFTER_CLEAN_SHUTDOWN Hardware_UpgradePolicy = "AFTER_CLEAN_SHUTDOWN"
    // Run scheduled upgrade when the virtual machine is powered on.
     Hardware_UpgradePolicy_ALWAYS Hardware_UpgradePolicy = "ALWAYS"
)

func (u Hardware_UpgradePolicy) Hardware_UpgradePolicy() bool {
    switch u {
        case Hardware_UpgradePolicy_NEVER:
            return true
        case Hardware_UpgradePolicy_AFTER_CLEAN_SHUTDOWN:
            return true
        case Hardware_UpgradePolicy_ALWAYS:
            return true
        default:
            return false
    }
}




// The ``UpgradeStatus`` enumeration class defines the valid virtual hardware upgrade statuses for a virtual machine.
//
// <p> See {@link com.vmware.vapi.bindings.ApiEnumeration enumerated types description}.
 
type Hardware_UpgradeStatus string

const (
    // No scheduled upgrade has been attempted.
     Hardware_UpgradeStatus_NONE Hardware_UpgradeStatus = "NONE"
    // Upgrade is scheduled but has not yet been run.
     Hardware_UpgradeStatus_PENDING Hardware_UpgradeStatus = "PENDING"
    // The most recent scheduled upgrade was successful.
     Hardware_UpgradeStatus_SUCCESS Hardware_UpgradeStatus = "SUCCESS"
    // The most recent scheduled upgrade was not successful.
     Hardware_UpgradeStatus_FAILED Hardware_UpgradeStatus = "FAILED"
)

func (u Hardware_UpgradeStatus) Hardware_UpgradeStatus() bool {
    switch u {
        case Hardware_UpgradeStatus_NONE:
            return true
        case Hardware_UpgradeStatus_PENDING:
            return true
        case Hardware_UpgradeStatus_SUCCESS:
            return true
        case Hardware_UpgradeStatus_FAILED:
            return true
        default:
            return false
    }
}





// The ``Info`` class contains information related to the virtual hardware of a virtual machine.
 type HardwareInfo struct {
    // Virtual hardware version.
    Version Hardware_Version
    // Scheduled upgrade policy.
    UpgradePolicy Hardware_UpgradePolicy
    // Target hardware version to be used on the next scheduled virtual hardware upgrade.
    UpgradeVersion *Hardware_Version
    // Scheduled upgrade status.
    UpgradeStatus Hardware_UpgradeStatus
    // Reason for the scheduled upgrade failure.
    UpgradeError *data.ErrorValue
}



//


// The ``UpdateSpec`` class describes the updates to virtual hardware settings of a virtual machine.
 type HardwareUpdateSpec struct {
    // Scheduled upgrade policy. 
//
//  If set to Hardware_UpgradePolicy#HardwareUpgradePolicy_NEVER, the HardwareInfo#upgradeVersion property will be reset to null.
    UpgradePolicy *Hardware_UpgradePolicy
    // Target hardware version to be used on the next scheduled virtual hardware upgrade. 
//
//  If specified, this property must represent a newer virtual hardware version than the current virtual hardware version reported in HardwareInfo#version.
    UpgradeVersion *Hardware_Version
}



//





func hardwareGetInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["vm"] = bindings.NewIdType([]string {"VirtualMachine"}, "")
    fieldNameMap["vm"] = "Vm"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func hardwareGetOutputType() bindings.BindingType {
    return bindings.NewReferenceType(HardwareInfoBindingType)
}

func hardwareGetRestMetadata() protocol.OperationRestMetadata {
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
       map[string]int{"Error": 500,"NotFound": 404,"ResourceInaccessible": 500,"ServiceUnavailable": 503,"Unauthenticated": 401,"Unauthorized": 403})
}


func hardwareUpdateInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["vm"] = bindings.NewIdType([]string {"VirtualMachine"}, "")
    fields["spec"] = bindings.NewReferenceType(HardwareUpdateSpecBindingType)
    fieldNameMap["vm"] = "Vm"
    fieldNameMap["spec"] = "Spec"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func hardwareUpdateOutputType() bindings.BindingType {
    return bindings.NewVoidType()
}

func hardwareUpdateRestMetadata() protocol.OperationRestMetadata {
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
       map[string]int{"Error": 500,"NotFound": 404,"AlreadyInDesiredState": 400,"InvalidArgument": 400,"Unsupported": 400,"ResourceBusy": 500,"ResourceInaccessible": 500,"ServiceUnavailable": 503,"Unauthenticated": 401,"Unauthorized": 403})
}


func hardwareUpgradeInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["vm"] = bindings.NewIdType([]string {"VirtualMachine"}, "")
    fields["version"] = bindings.NewOptionalType(bindings.NewEnumType("com.vmware.vcenter.vm.hardware.version", reflect.TypeOf(Hardware_Version(Hardware_Version_VMX_03))))
    fieldNameMap["vm"] = "Vm"
    fieldNameMap["version"] = "Version"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func hardwareUpgradeOutputType() bindings.BindingType {
    return bindings.NewVoidType()
}

func hardwareUpgradeRestMetadata() protocol.OperationRestMetadata {
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
       map[string]int{"Error": 500,"NotFound": 404,"NotAllowedInCurrentState": 400,"AlreadyInDesiredState": 400,"InvalidArgument": 400,"Unsupported": 400,"ResourceBusy": 500,"ResourceInaccessible": 500,"ServiceUnavailable": 503,"Unauthenticated": 401,"Unauthorized": 403})
}



func HardwareInfoBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["version"] = bindings.NewEnumType("com.vmware.vcenter.vm.hardware.version", reflect.TypeOf(Hardware_Version(Hardware_Version_VMX_03)))
    fieldNameMap["version"] = "Version"
    fields["upgrade_policy"] = bindings.NewEnumType("com.vmware.vcenter.vm.hardware.upgrade_policy", reflect.TypeOf(Hardware_UpgradePolicy(Hardware_UpgradePolicy_NEVER)))
    fieldNameMap["upgrade_policy"] = "UpgradePolicy"
    fields["upgrade_version"] = bindings.NewOptionalType(bindings.NewEnumType("com.vmware.vcenter.vm.hardware.version", reflect.TypeOf(Hardware_Version(Hardware_Version_VMX_03))))
    fieldNameMap["upgrade_version"] = "UpgradeVersion"
    fields["upgrade_status"] = bindings.NewEnumType("com.vmware.vcenter.vm.hardware.upgrade_status", reflect.TypeOf(Hardware_UpgradeStatus(Hardware_UpgradeStatus_NONE)))
    fieldNameMap["upgrade_status"] = "UpgradeStatus"
    fields["upgrade_error"] = bindings.NewOptionalType(bindings.NewAnyErrorType())
    fieldNameMap["upgrade_error"] = "UpgradeError"
    var validators = []bindings.Validator{}
    uv1 := bindings.NewUnionValidator("upgrade_policy",
        map[string][]bindings.FieldData {
            "AFTER_CLEAN_SHUTDOWN": []bindings.FieldData {
                 bindings.NewFieldData("upgrade_version", true),
            },
            "ALWAYS": []bindings.FieldData {
                 bindings.NewFieldData("upgrade_version", true),
            },
            "NEVER": []bindings.FieldData {},
        },
    )
    validators = append(validators, uv1)
    uv2 := bindings.NewUnionValidator("upgrade_status",
        map[string][]bindings.FieldData {
            "FAILED": []bindings.FieldData {
                 bindings.NewFieldData("upgrade_error", true),
            },
            "NONE": []bindings.FieldData {},
            "PENDING": []bindings.FieldData {},
            "SUCCESS": []bindings.FieldData {},
        },
    )
    validators = append(validators, uv2)
    return bindings.NewStructType("com.vmware.vcenter.vm.hardware.info",fields, reflect.TypeOf(HardwareInfo{}), fieldNameMap, validators)
}

func HardwareUpdateSpecBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["upgrade_policy"] = bindings.NewOptionalType(bindings.NewEnumType("com.vmware.vcenter.vm.hardware.upgrade_policy", reflect.TypeOf(Hardware_UpgradePolicy(Hardware_UpgradePolicy_NEVER))))
    fieldNameMap["upgrade_policy"] = "UpgradePolicy"
    fields["upgrade_version"] = bindings.NewOptionalType(bindings.NewEnumType("com.vmware.vcenter.vm.hardware.version", reflect.TypeOf(Hardware_Version(Hardware_Version_VMX_03))))
    fieldNameMap["upgrade_version"] = "UpgradeVersion"
    var validators = []bindings.Validator{}
    uv1 := bindings.NewUnionValidator("upgrade_policy",
        map[string][]bindings.FieldData {
            "AFTER_CLEAN_SHUTDOWN": []bindings.FieldData {
                 bindings.NewFieldData("upgrade_version", false),
            },
            "ALWAYS": []bindings.FieldData {
                 bindings.NewFieldData("upgrade_version", false),
            },
            "NEVER": []bindings.FieldData {},
        },
    )
    validators = append(validators, uv1)
    return bindings.NewStructType("com.vmware.vcenter.vm.hardware.update_spec",fields, reflect.TypeOf(HardwareUpdateSpec{}), fieldNameMap, validators)
}


