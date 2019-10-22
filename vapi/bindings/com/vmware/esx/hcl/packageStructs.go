/* Copyright © 2019 VMware, Inc. All Rights Reserved.
     SPDX-License-Identifier: BSD-2-Clause */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Data type definitions file for package: com.vmware.esx.hcl.
 * Includes binding types of a top level structures and enumerations.
 * Shared by client-side stubs and server-side skeletons to ensure type
 * compatibility.
 */

package hcl

import (
    "reflect"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/bindings/com/vmware/vapi/std"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/bindings"
    "time"
)


// The ``CompatibilityStatus`` enumeration class defines compatibility status of a given server or PCI device against a specific release of ESXi. **Warning:** This enumeration is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
//
// <p> See {@link com.vmware.vapi.bindings.ApiEnumeration enumerated types description}.

type CompatibilityStatus string

const (
    // When given hardware is certified for the specified ESXi release but no validation of the software of this hardware is performed. **Warning:** This constant field is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
     CompatibilityStatus_COMPATIBLE CompatibilityStatus = "COMPATIBLE"
    // When given hardware is not certified for the specified ESXi release. **Warning:** This constant field is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
     CompatibilityStatus_INCOMPATIBLE CompatibilityStatus = "INCOMPATIBLE"
    // When there is no information about specified hardware. **Warning:** This constant field is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
     CompatibilityStatus_UNAVAILABLE CompatibilityStatus = "UNAVAILABLE"
    // When given hardware is certified for the specified ESXi release. Its software is also validated and it is also certified. **Warning:** This constant field is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
     CompatibilityStatus_CERTIFIED CompatibilityStatus = "CERTIFIED"
    // When given hardware is certified for the specified ESXi release. Its software is also validated and it is not certified. **Warning:** This constant field is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
     CompatibilityStatus_NOT_CERTIFIED CompatibilityStatus = "NOT_CERTIFIED"
)

func (c CompatibilityStatus) CompatibilityStatus() bool {
    switch c {
        case CompatibilityStatus_COMPATIBLE:
            return true
        case CompatibilityStatus_INCOMPATIBLE:
            return true
        case CompatibilityStatus_UNAVAILABLE:
            return true
        case CompatibilityStatus_CERTIFIED:
            return true
        case CompatibilityStatus_NOT_CERTIFIED:
            return true
        default:
            return false
    }
}





// The ``Driver`` class contains properties describing information about a driver. **Warning:** This class is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
type Driver struct {
    // The name of the driver. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    Name string
    // The version of the driver. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    Version string
    // The vendor that produced the driver. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    Vendor *string
}


// TODO some error doesn't have ErrorType. Resolve this later VAPI-3009
//


// The ``Firmware`` class contains properties describing information about a firmware. **Warning:** This class is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
type Firmware struct {
    // The version of the firmware. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    Version string
    // The vendor that produced the firmware. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    Vendor *string
}


// TODO some error doesn't have ErrorType. Resolve this later VAPI-3009
//


// The ``PCIDevice`` class contains properties describing information about a single PCI device on a host. **Warning:** This class is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
type PCIDevice struct {
    // The device's address in a given ESXi host. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    PciAddress string
    // A unique number assigned to each computer hardware device that helps to identify the chipset manufacturer. For example, Dell, Broadcom, etc. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    Vid string
    // A unique number that identifies the specific device of the Vendor (VID). **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    Did string
    // A unique number that identifies the card manufacturer. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    Svid string
    // A unique number that identifies the specific device of Subsystem Vendor (SVID). **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    Ssid string
    // The name of the vendor. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    Vendor string
    // The name of the device model. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    ModelName string
    // Register that specifies the type of function the device performs. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    ClassCode string
    // Register that specifies the specific function the device performs. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    SubClassCode string
    // Currently installed driver used by the device. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    Driver Driver
    // Currently installed firmware used by the device. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    Firmware *Firmware
    // Shows whether the device is part of VSAN cluster or not. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    UsedByVsan bool
}


// TODO some error doesn't have ErrorType. Resolve this later VAPI-3009
//


// The ``Server`` class contains properties describing information about a server. **Warning:** This class is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
type Server struct {
    // The name of the server model. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    ModelName string
    // The name of the vendor. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    Vendor string
    // The CPU series name. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    CpuSeries string
    // The current CPU features. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    CpuFeatures string
    // Currently installed BIOS of the server. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    Bios Firmware
}


// TODO some error doesn't have ErrorType. Resolve this later VAPI-3009
//


// The ``Notification`` class contains properties to describe any info/warning/error messages that Tasks can raise. **Warning:** This class is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
type Notification struct {
    // The notification id. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    Id string
    // The time the notification was raised/found. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    Time time.Time
    // The notification message. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    Message std.LocalizableMessage
    // The resolution message, if any. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    Resolution *std.LocalizableMessage
}


// TODO some error doesn't have ErrorType. Resolve this later VAPI-3009
//


// The ``Notifications`` class contains info/warning/error messages that can be reported be the task. **Warning:** This class is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
type Notifications struct {
    // Info notification messages reported. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    Info []Notification
    // Warning notification messages reported. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    Warnings []Notification
    // Error notification messages reported. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    Errors []Notification
}


// TODO some error doesn't have ErrorType. Resolve this later VAPI-3009
//






func DriverBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["name"] = bindings.NewStringType()
    fieldNameMap["name"] = "Name"
    fields["version"] = bindings.NewStringType()
    fieldNameMap["version"] = "Version"
    fields["vendor"] = bindings.NewOptionalType(bindings.NewStringType())
    fieldNameMap["vendor"] = "Vendor"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.esx.hcl.driver",fields, reflect.TypeOf(Driver{}), fieldNameMap, validators)
}

func FirmwareBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["version"] = bindings.NewStringType()
    fieldNameMap["version"] = "Version"
    fields["vendor"] = bindings.NewOptionalType(bindings.NewStringType())
    fieldNameMap["vendor"] = "Vendor"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.esx.hcl.firmware",fields, reflect.TypeOf(Firmware{}), fieldNameMap, validators)
}

func PCIDeviceBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["pci_address"] = bindings.NewStringType()
    fieldNameMap["pci_address"] = "PciAddress"
    fields["vid"] = bindings.NewStringType()
    fieldNameMap["vid"] = "Vid"
    fields["did"] = bindings.NewStringType()
    fieldNameMap["did"] = "Did"
    fields["svid"] = bindings.NewStringType()
    fieldNameMap["svid"] = "Svid"
    fields["ssid"] = bindings.NewStringType()
    fieldNameMap["ssid"] = "Ssid"
    fields["vendor"] = bindings.NewStringType()
    fieldNameMap["vendor"] = "Vendor"
    fields["model_name"] = bindings.NewStringType()
    fieldNameMap["model_name"] = "ModelName"
    fields["class_code"] = bindings.NewStringType()
    fieldNameMap["class_code"] = "ClassCode"
    fields["sub_class_code"] = bindings.NewStringType()
    fieldNameMap["sub_class_code"] = "SubClassCode"
    fields["driver"] = bindings.NewReferenceType(DriverBindingType)
    fieldNameMap["driver"] = "Driver"
    fields["firmware"] = bindings.NewOptionalType(bindings.NewReferenceType(FirmwareBindingType))
    fieldNameMap["firmware"] = "Firmware"
    fields["used_by_vsan"] = bindings.NewBooleanType()
    fieldNameMap["used_by_vsan"] = "UsedByVsan"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.esx.hcl.PCI_device",fields, reflect.TypeOf(PCIDevice{}), fieldNameMap, validators)
}

func ServerBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["model_name"] = bindings.NewStringType()
    fieldNameMap["model_name"] = "ModelName"
    fields["vendor"] = bindings.NewStringType()
    fieldNameMap["vendor"] = "Vendor"
    fields["cpu_series"] = bindings.NewStringType()
    fieldNameMap["cpu_series"] = "CpuSeries"
    fields["cpu_features"] = bindings.NewStringType()
    fieldNameMap["cpu_features"] = "CpuFeatures"
    fields["bios"] = bindings.NewReferenceType(FirmwareBindingType)
    fieldNameMap["bios"] = "Bios"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.esx.hcl.server",fields, reflect.TypeOf(Server{}), fieldNameMap, validators)
}

func NotificationBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["id"] = bindings.NewStringType()
    fieldNameMap["id"] = "Id"
    fields["time"] = bindings.NewDateTimeType()
    fieldNameMap["time"] = "Time"
    fields["message"] = bindings.NewReferenceType(std.LocalizableMessageBindingType)
    fieldNameMap["message"] = "Message"
    fields["resolution"] = bindings.NewOptionalType(bindings.NewReferenceType(std.LocalizableMessageBindingType))
    fieldNameMap["resolution"] = "Resolution"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.esx.hcl.notification",fields, reflect.TypeOf(Notification{}), fieldNameMap, validators)
}

func NotificationsBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["info"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewReferenceType(NotificationBindingType), reflect.TypeOf([]Notification{})))
    fieldNameMap["info"] = "Info"
    fields["warnings"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewReferenceType(NotificationBindingType), reflect.TypeOf([]Notification{})))
    fieldNameMap["warnings"] = "Warnings"
    fields["errors"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewReferenceType(NotificationBindingType), reflect.TypeOf([]Notification{})))
    fieldNameMap["errors"] = "Errors"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.esx.hcl.notifications",fields, reflect.TypeOf(Notifications{}), fieldNameMap, validators)
}


