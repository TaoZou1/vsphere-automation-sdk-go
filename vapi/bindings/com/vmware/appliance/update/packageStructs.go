/* Copyright © 2019 VMware, Inc. All Rights Reserved.
     SPDX-License-Identifier: BSD-2-Clause */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Data type definitions file for package: com.vmware.appliance.update.
 * Includes binding types of a top level structures and enumerations.
 * Shared by client-side stubs and server-side skeletons to ensure type
 * compatibility.
 */

package update

import (
    "reflect"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/bindings/com/vmware/vapi/std"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/bindings"
    "time"
)



// The ``ServiceInfo`` class describes a service to be stopped and started during the update installation.
type ServiceInfo struct {
    // Service ID
    Service string
    // Service description
    Description std.LocalizableMessage
}



func (ServiceInfo ServiceInfo) Error() string {
    return "com.vmware.appliance.update.service_info"
}



// The ``CommonInfo`` class defines common update information
type CommonInfo struct {
    // Description of the update. The short information what this update is. E.g. "Update2 for vCenter Server Appliance 6.5"
    Description std.LocalizableMessage
    // Update priority
    Priority CommonInfo_Priority
    // Update severity
    Severity CommonInfo_Severity
    // Update category
    UpdateType CommonInfo_Category
    // Update release date.
    ReleaseDate time.Time
    // Flag indicating whether reboot is required after update.
    RebootRequired bool
    // Download Size of update in Megabytes.
    Size int64
}



func (CommonInfo CommonInfo) Error() string {
    return "com.vmware.appliance.update.common_info"
}

    
    // The ``Priority`` enumeration class defines the update installation priority recommendations.
    //
    // <p> See {@link com.vmware.vapi.bindings.ApiEnumeration enumerated types description}.
     
    type CommonInfo_Priority string

    const (
        // Install ASAP
         CommonInfo_Priority_HIGH CommonInfo_Priority = "HIGH"
        // Install at the earliest convenience
         CommonInfo_Priority_MEDIUM CommonInfo_Priority = "MEDIUM"
        // Install at your discretion
         CommonInfo_Priority_LOW CommonInfo_Priority = "LOW"
    )

    func (p CommonInfo_Priority) CommonInfo_Priority() bool {
        switch p {
            case CommonInfo_Priority_HIGH:
                return true
            case CommonInfo_Priority_MEDIUM:
                return true
            case CommonInfo_Priority_LOW:
                return true
            default:
                return false
        }
    }

    
    // The ``Severity`` enumeration class defines the severity of the issues fixed in the update.
    //
    // <p> See {@link com.vmware.vapi.bindings.ApiEnumeration enumerated types description}.
     
    type CommonInfo_Severity string

    const (
        // Vulnerabilities that can be exploited by an unauthenticated attacker from the Internet or those that break the guest/host Operating System isolation. The exploitation results in the complete compromise of confidentiality, integrity, and availability of user data and/or processing resources without user interaction. Exploitation could be leveraged to propagate an Internet worm or execute arbitrary code between Virtual Machines and/or the Host Operating System.
         CommonInfo_Severity_CRITICAL CommonInfo_Severity = "CRITICAL"
        // Vulnerabilities that are not rated critical but whose exploitation results in the complete compromise of confidentiality and/or integrity of user data and/or processing resources through user assistance or by authenticated attackers. This rating also applies to those vulnerabilities which could lead to the complete compromise of availability when exploitation is by a remote unauthenticated attacker from the Internet or through a breach of virtual machine isolation.
         CommonInfo_Severity_IMPORTANT CommonInfo_Severity = "IMPORTANT"
        // Vulnerabilities where the ability to exploit is mitigated to a significant degree by configuration or difficulty of exploitation, but in certain deployment scenarios could still lead to the compromise of confidentiality, integrity, or availability of user data and/or processing resources.
         CommonInfo_Severity_MODERATE CommonInfo_Severity = "MODERATE"
        // All other issues that have a security impact. Vulnerabilities where exploitation is believed to be extremely difficult, or where successful exploitation would have minimal impact
         CommonInfo_Severity_LOW CommonInfo_Severity = "LOW"
    )

    func (s CommonInfo_Severity) CommonInfo_Severity() bool {
        switch s {
            case CommonInfo_Severity_CRITICAL:
                return true
            case CommonInfo_Severity_IMPORTANT:
                return true
            case CommonInfo_Severity_MODERATE:
                return true
            case CommonInfo_Severity_LOW:
                return true
            default:
                return false
        }
    }

    
    // The ``Category`` enumeration class defines update type
    //
    // <p> See {@link com.vmware.vapi.bindings.ApiEnumeration enumerated types description}.
     
    type CommonInfo_Category string

    const (
        // Fixes vulnerabilities, doesn't change functionality
         CommonInfo_Category_SECURITY CommonInfo_Category = "SECURITY"
        // Fixes bugs/vulnerabilities, doesn't change functionality
         CommonInfo_Category_FIX CommonInfo_Category = "FIX"
        // Changes product functionality
         CommonInfo_Category_UPDATE CommonInfo_Category = "UPDATE"
        // Introduces new features, significantly changes product functionality
         CommonInfo_Category_UPGRADE CommonInfo_Category = "UPGRADE"
    )

    func (c CommonInfo_Category) CommonInfo_Category() bool {
        switch c {
            case CommonInfo_Category_SECURITY:
                return true
            case CommonInfo_Category_FIX:
                return true
            case CommonInfo_Category_UPDATE:
                return true
            case CommonInfo_Category_UPGRADE:
                return true
            default:
                return false
        }
    }



// The ``Summary`` class contains the essential information about the update
type Summary struct {
    // Version in form of X.Y.Z.P. e.g. 6.5.1.5400
    Version string
    // Name of the update. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    Name *string
    // Description of the update. The short information what this update is. E.g. "Update2 for vCenter Server Appliance 6.5"
    Description std.LocalizableMessage
    // Update priority
    Priority CommonInfo_Priority
    // Update severity
    Severity CommonInfo_Severity
    // Update category
    UpdateType CommonInfo_Category
    // Update release date.
    ReleaseDate time.Time
    // Flag indicating whether reboot is required after update.
    RebootRequired bool
    // Download Size of update in Megabytes.
    Size int64
}



func (Summary Summary) Error() string {
    return "com.vmware.appliance.update.summary"
}







func ServiceInfoBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["service"] = bindings.NewIdType([]string {"com.vmware.appliance.service"}, "")
    fieldNameMap["service"] = "Service"
    fields["description"] = bindings.NewReferenceType(std.LocalizableMessageBindingType)
    fieldNameMap["description"] = "Description"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.appliance.update.service_info",fields, reflect.TypeOf(ServiceInfo{}), fieldNameMap, validators)
}

func CommonInfoBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["description"] = bindings.NewReferenceType(std.LocalizableMessageBindingType)
    fieldNameMap["description"] = "Description"
    fields["priority"] = bindings.NewEnumType("com.vmware.appliance.update.common_info.priority", reflect.TypeOf(CommonInfo_Priority(CommonInfo_Priority_HIGH)))
    fieldNameMap["priority"] = "Priority"
    fields["severity"] = bindings.NewEnumType("com.vmware.appliance.update.common_info.severity", reflect.TypeOf(CommonInfo_Severity(CommonInfo_Severity_CRITICAL)))
    fieldNameMap["severity"] = "Severity"
    fields["update_type"] = bindings.NewEnumType("com.vmware.appliance.update.common_info.category", reflect.TypeOf(CommonInfo_Category(CommonInfo_Category_SECURITY)))
    fieldNameMap["update_type"] = "UpdateType"
    fields["release_date"] = bindings.NewDateTimeType()
    fieldNameMap["release_date"] = "ReleaseDate"
    fields["reboot_required"] = bindings.NewBooleanType()
    fieldNameMap["reboot_required"] = "RebootRequired"
    fields["size"] = bindings.NewIntegerType()
    fieldNameMap["size"] = "Size"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.appliance.update.common_info",fields, reflect.TypeOf(CommonInfo{}), fieldNameMap, validators)
}

func SummaryBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["version"] = bindings.NewIdType([]string {"com.vmware.appliance.update.pending"}, "")
    fieldNameMap["version"] = "Version"
    fields["name"] = bindings.NewOptionalType(bindings.NewStringType())
    fieldNameMap["name"] = "Name"
    fields["description"] = bindings.NewReferenceType(std.LocalizableMessageBindingType)
    fieldNameMap["description"] = "Description"
    fields["priority"] = bindings.NewEnumType("com.vmware.appliance.update.common_info.priority", reflect.TypeOf(CommonInfo_Priority(CommonInfo_Priority_HIGH)))
    fieldNameMap["priority"] = "Priority"
    fields["severity"] = bindings.NewEnumType("com.vmware.appliance.update.common_info.severity", reflect.TypeOf(CommonInfo_Severity(CommonInfo_Severity_CRITICAL)))
    fieldNameMap["severity"] = "Severity"
    fields["update_type"] = bindings.NewEnumType("com.vmware.appliance.update.common_info.category", reflect.TypeOf(CommonInfo_Category(CommonInfo_Category_SECURITY)))
    fieldNameMap["update_type"] = "UpdateType"
    fields["release_date"] = bindings.NewDateTimeType()
    fieldNameMap["release_date"] = "ReleaseDate"
    fields["reboot_required"] = bindings.NewBooleanType()
    fieldNameMap["reboot_required"] = "RebootRequired"
    fields["size"] = bindings.NewIntegerType()
    fieldNameMap["size"] = "Size"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.appliance.update.summary",fields, reflect.TypeOf(Summary{}), fieldNameMap, validators)
}


