/* Copyright © 2019 VMware, Inc. All Rights Reserved.
     SPDX-License-Identifier: BSD-2-Clause */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Data type definitions file for service: Tools.
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



// Current run state of VMware Tools in the guest operating system. This enumeration was added in vSphere API 7.0.
//
// <p> See {@link com.vmware.vapi.bindings.ApiEnumeration enumerated types description}.
 
type Tools_RunState string

const (
    // VMware Tools is not running. This constant field was added in vSphere API 7.0.
     Tools_RunState_NOT_RUNNING Tools_RunState = "NOT_RUNNING"
    // VMware Tools is running. This constant field was added in vSphere API 7.0.
     Tools_RunState_RUNNING Tools_RunState = "RUNNING"
    // VMware Tools is running scripts as part of a state transition. This constant field was added in vSphere API 7.0.
     Tools_RunState_EXECUTING_SCRIPTS Tools_RunState = "EXECUTING_SCRIPTS"
)

func (r Tools_RunState) Tools_RunState() bool {
    switch r {
        case Tools_RunState_NOT_RUNNING:
            return true
        case Tools_RunState_RUNNING:
            return true
        case Tools_RunState_EXECUTING_SCRIPTS:
            return true
        default:
            return false
    }
}




// The ``UpgradePolicy`` enumeration class defines when Tools are auto-upgraded for a virtual machine. This enumeration was added in vSphere API 7.0.
//
// <p> See {@link com.vmware.vapi.bindings.ApiEnumeration enumerated types description}.
 
type Tools_UpgradePolicy string

const (
    // No auto-upgrades for Tools will be performed for this virtual machine. Users must manually invoke the Tools#upgrade method to update Tools. This constant field was added in vSphere API 7.0.
     Tools_UpgradePolicy_MANUAL Tools_UpgradePolicy = "MANUAL"
    // When the virtual machine is power-cycled, the system checks for a newer version of Tools when the virtual machine is powered on. If it is available, a Tools upgrade is automatically performed on the virtual machine and it is rebooted if necessary. This constant field was added in vSphere API 7.0.
     Tools_UpgradePolicy_UPGRADE_AT_POWER_CYCLE Tools_UpgradePolicy = "UPGRADE_AT_POWER_CYCLE"
)

func (u Tools_UpgradePolicy) Tools_UpgradePolicy() bool {
    switch u {
        case Tools_UpgradePolicy_MANUAL:
            return true
        case Tools_UpgradePolicy_UPGRADE_AT_POWER_CYCLE:
            return true
        default:
            return false
    }
}




// The ``VersionStatus`` enumeration class defines the version status types of VMware Tools installed in the guest operating system. This enumeration was added in vSphere API 7.0.
//
// <p> See {@link com.vmware.vapi.bindings.ApiEnumeration enumerated types description}.
 
type Tools_VersionStatus string

const (
    // VMware Tools has never been installed. This constant field was added in vSphere API 7.0.
     Tools_VersionStatus_NOT_INSTALLED Tools_VersionStatus = "NOT_INSTALLED"
    // VMware Tools is installed, and the version is current. This constant field was added in vSphere API 7.0.
     Tools_VersionStatus_CURRENT Tools_VersionStatus = "CURRENT"
    // VMware Tools is installed, but it is not managed by VMware. This includes open-vm-tools or OSPs which should be managed by the guest operating system. This constant field was added in vSphere API 7.0.
     Tools_VersionStatus_UNMANAGED Tools_VersionStatus = "UNMANAGED"
    // VMware Tools is installed, but the version is too old. This constant field was added in vSphere API 7.0.
     Tools_VersionStatus_TOO_OLD_UNSUPPORTED Tools_VersionStatus = "TOO_OLD_UNSUPPORTED"
    // VMware Tools is installed, supported, but a newer version is available. This constant field was added in vSphere API 7.0.
     Tools_VersionStatus_SUPPORTED_OLD Tools_VersionStatus = "SUPPORTED_OLD"
    // VMware Tools is installed, supported, and newer than the version available on the host. This constant field was added in vSphere API 7.0.
     Tools_VersionStatus_SUPPORTED_NEW Tools_VersionStatus = "SUPPORTED_NEW"
    // VMware Tools is installed, and the version is known to be too new to work correctly with this virtual machine. This constant field was added in vSphere API 7.0.
     Tools_VersionStatus_TOO_NEW Tools_VersionStatus = "TOO_NEW"
    // VMware Tools is installed, but the installed version is known to have a grave bug and should be immediately upgraded. This constant field was added in vSphere API 7.0.
     Tools_VersionStatus_BLACKLISTED Tools_VersionStatus = "BLACKLISTED"
)

func (v Tools_VersionStatus) Tools_VersionStatus() bool {
    switch v {
        case Tools_VersionStatus_NOT_INSTALLED:
            return true
        case Tools_VersionStatus_CURRENT:
            return true
        case Tools_VersionStatus_UNMANAGED:
            return true
        case Tools_VersionStatus_TOO_OLD_UNSUPPORTED:
            return true
        case Tools_VersionStatus_SUPPORTED_OLD:
            return true
        case Tools_VersionStatus_SUPPORTED_NEW:
            return true
        case Tools_VersionStatus_TOO_NEW:
            return true
        case Tools_VersionStatus_BLACKLISTED:
            return true
        default:
            return false
    }
}




// The ``ToolsInstallType`` enumeration class defines the installation type of the Tools in the guest operating system. This enumeration was added in vSphere API 7.0.
//
// <p> See {@link com.vmware.vapi.bindings.ApiEnumeration enumerated types description}.
 
type Tools_ToolsInstallType string

const (
    // Installation type is not known. Most likely tools have been installed by OSPs or open-vm-tools, but a version that does not report its install type or an install type that we do not recognize. This constant field was added in vSphere API 7.0.
     Tools_ToolsInstallType_UNKNOWN Tools_ToolsInstallType = "UNKNOWN"
    // MSI is the installation type used for VMware Tools on Windows. This constant field was added in vSphere API 7.0.
     Tools_ToolsInstallType_MSI Tools_ToolsInstallType = "MSI"
    // Tools have been installed by the tar installer. This constant field was added in vSphere API 7.0.
     Tools_ToolsInstallType_TAR Tools_ToolsInstallType = "TAR"
    // OSPs are RPM or Debian packages tailored for the OS in the VM. See http://packages.vmware.com. This constant field was added in vSphere API 7.0.
     Tools_ToolsInstallType_OSP Tools_ToolsInstallType = "OSP"
    // open-vm-tools are the open-source version of VMware Tools, may have been packaged by the OS vendor. This constant field was added in vSphere API 7.0.
     Tools_ToolsInstallType_OPEN_VM_TOOLS Tools_ToolsInstallType = "OPEN_VM_TOOLS"
)

func (t Tools_ToolsInstallType) Tools_ToolsInstallType() bool {
    switch t {
        case Tools_ToolsInstallType_UNKNOWN:
            return true
        case Tools_ToolsInstallType_MSI:
            return true
        case Tools_ToolsInstallType_TAR:
            return true
        case Tools_ToolsInstallType_OSP:
            return true
        case Tools_ToolsInstallType_OPEN_VM_TOOLS:
            return true
        default:
            return false
    }
}





// The ``Info`` class describes the VMWare Tools properties of a virtual machine. This class was added in vSphere API 7.0.
 type ToolsInfo struct {
    // Set if the virtual machine supports auto-upgrading Tools via Tools_UpgradePolicy. This property was added in vSphere API 7.0.
    AutoUpdateSupported bool
    // Number of attempts that have been made to install or upgrade the version of Tools installed on this virtual machine. This property was added in vSphere API 7.0.
    InstallAttemptCount *int64
    // Error that happened, if any, during last attempt to upgrade or install Tools. This property was added in vSphere API 7.0.
    Error *data.ErrorValue
    // Version of VMware Tools installed on the guest operating system. This property was added in vSphere API 7.0.
    VersionNumber *int64
    // Version of VMware Tools installed on the guest operating system. This is a human-readable value that should not be parsed. This property was added in vSphere API 7.0.
    Version *string
    // Tools upgrade policy setting for the virtual machine. Tools_UpgradePolicy. This property was added in vSphere API 7.0.
    UpgradePolicy Tools_UpgradePolicy
    // Current version status of VMware Tools in the guest operating system, if known. This property was added in vSphere API 7.0.
    VersionStatus *Tools_VersionStatus
    // Current installation type of VMware Tools in the guest operating system. This property was added in vSphere API 7.0.
    InstallType *Tools_ToolsInstallType
    // Current run state of VMware Tools in the guest operating system. This property was added in vSphere API 7.0.
    RunState Tools_RunState
}



//


// The (\\\\@name UpdateSpec} class describes the VMware Tools properties of a virtual machine that can be updated. This class was added in vSphere API 7.0.
 type ToolsUpdateSpec struct {
    // Tools upgrade policy setting for the virtual machine. Tools_UpgradePolicy. This property was added in vSphere API 7.0.
    UpgradePolicy *Tools_UpgradePolicy
}



//





func toolsGetInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["vm"] = bindings.NewIdType([]string {"VirtualMachine"}, "")
    fieldNameMap["vm"] = "Vm"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func toolsGetOutputType() bindings.BindingType {
    return bindings.NewReferenceType(ToolsInfoBindingType)
}

func toolsGetRestMetadata() protocol.OperationRestMetadata {
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
       map[string]int{"Error": 500,"NotFound": 404})
}


func toolsUpdateInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["vm"] = bindings.NewIdType([]string {"VirtualMachine"}, "")
    fields["spec"] = bindings.NewReferenceType(ToolsUpdateSpecBindingType)
    fieldNameMap["vm"] = "Vm"
    fieldNameMap["spec"] = "Spec"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func toolsUpdateOutputType() bindings.BindingType {
    return bindings.NewVoidType()
}

func toolsUpdateRestMetadata() protocol.OperationRestMetadata {
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
       map[string]int{"Error": 500,"InvalidArgument": 400,"NotFound": 404})
}


func toolsUpgradeInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["vm"] = bindings.NewIdType([]string {"VirtualMachine"}, "")
    fields["command_line_options"] = bindings.NewOptionalType(bindings.NewStringType())
    fieldNameMap["vm"] = "Vm"
    fieldNameMap["command_line_options"] = "CommandLineOptions"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func toolsUpgradeOutputType() bindings.BindingType {
    return bindings.NewVoidType()
}

func toolsUpgradeRestMetadata() protocol.OperationRestMetadata {
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
       map[string]int{"NotFound": 404,"ServiceUnavailable": 503,"NotAllowedInCurrentState": 400,"AlreadyInDesiredState": 400,"Error": 500})
}



func ToolsInfoBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["auto_update_supported"] = bindings.NewBooleanType()
    fieldNameMap["auto_update_supported"] = "AutoUpdateSupported"
    fields["install_attempt_count"] = bindings.NewOptionalType(bindings.NewIntegerType())
    fieldNameMap["install_attempt_count"] = "InstallAttemptCount"
    fields["error"] = bindings.NewOptionalType(bindings.NewAnyErrorType())
    fieldNameMap["error"] = "Error"
    fields["version_number"] = bindings.NewOptionalType(bindings.NewIntegerType())
    fieldNameMap["version_number"] = "VersionNumber"
    fields["version"] = bindings.NewOptionalType(bindings.NewStringType())
    fieldNameMap["version"] = "Version"
    fields["upgrade_policy"] = bindings.NewEnumType("com.vmware.vcenter.vm.tools.upgrade_policy", reflect.TypeOf(Tools_UpgradePolicy(Tools_UpgradePolicy_MANUAL)))
    fieldNameMap["upgrade_policy"] = "UpgradePolicy"
    fields["version_status"] = bindings.NewOptionalType(bindings.NewEnumType("com.vmware.vcenter.vm.tools.version_status", reflect.TypeOf(Tools_VersionStatus(Tools_VersionStatus_NOT_INSTALLED))))
    fieldNameMap["version_status"] = "VersionStatus"
    fields["install_type"] = bindings.NewOptionalType(bindings.NewEnumType("com.vmware.vcenter.vm.tools.tools_install_type", reflect.TypeOf(Tools_ToolsInstallType(Tools_ToolsInstallType_UNKNOWN))))
    fieldNameMap["install_type"] = "InstallType"
    fields["run_state"] = bindings.NewEnumType("com.vmware.vcenter.vm.tools.run_state", reflect.TypeOf(Tools_RunState(Tools_RunState_NOT_RUNNING)))
    fieldNameMap["run_state"] = "RunState"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.vm.tools.info",fields, reflect.TypeOf(ToolsInfo{}), fieldNameMap, validators)
}

func ToolsUpdateSpecBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["upgrade_policy"] = bindings.NewOptionalType(bindings.NewEnumType("com.vmware.vcenter.vm.tools.upgrade_policy", reflect.TypeOf(Tools_UpgradePolicy(Tools_UpgradePolicy_MANUAL))))
    fieldNameMap["upgrade_policy"] = "UpgradePolicy"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.vm.tools.update_spec",fields, reflect.TypeOf(ToolsUpdateSpec{}), fieldNameMap, validators)
}


