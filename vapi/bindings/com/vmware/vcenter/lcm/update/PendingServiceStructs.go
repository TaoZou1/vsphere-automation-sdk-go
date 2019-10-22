/* Copyright © 2019 VMware, Inc. All Rights Reserved.
     SPDX-License-Identifier: BSD-2-Clause */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Data type definitions file for service: Pending.
 * Includes binding types of a structures and enumerations defined in the service.
 * Shared by client-side stubs and server-side skeletons to ensure type
 * compatibility.
 */

package update

import (
    "reflect"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/bindings"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/data"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/protocol"
    "net/url"
    "time"
)

// Resource type for pending update
const Pending_RESOURCE_TYPE = "com.vmware.vcenter.lcm.update.pending"


// Level of severity for applying a given patch or update.
//
// <p> See {@link com.vmware.vapi.bindings.ApiEnumeration enumerated types description}.
 
type Pending_SeverityType string

const (
    // Vulnerabilities that can be exploited by an unauthenticated attacker from the Internet or those that break the guest/host Operating System isolation.
     Pending_SeverityType_CRITICAL Pending_SeverityType = "CRITICAL"
    // Vulnerabilities that are not rated critical but whose exploitation results in the complete compromise of confidentiality and/or integrity of user data and/or processing resources through user assistance or by authenticated attackers.
     Pending_SeverityType_IMPORTANT Pending_SeverityType = "IMPORTANT"
    // Vulnerabilities where the ability to exploit is mitigated to a significant degree by configuration or difficulty of exploitation, but in certain deployment scenarios could still lead to the compromise of confidentiality, integrity, or availability of user data and/or processing resources.
     Pending_SeverityType_MODERATE Pending_SeverityType = "MODERATE"
    // All other issues that may or maynot have a security impact. Vulnerabilities where exploitation is believed to be extremely difficult, or where successful exploitation would have minimal impact.
     Pending_SeverityType_LOW Pending_SeverityType = "LOW"
)

func (s Pending_SeverityType) Pending_SeverityType() bool {
    switch s {
        case Pending_SeverityType_CRITICAL:
            return true
        case Pending_SeverityType_IMPORTANT:
            return true
        case Pending_SeverityType_MODERATE:
            return true
        case Pending_SeverityType_LOW:
            return true
        default:
            return false
    }
}




// The ``Category`` enumeration class defines the type of payload this release has on top of previous release
//
// <p> See {@link com.vmware.vapi.bindings.ApiEnumeration enumerated types description}.
 
type Pending_Category string

const (
    // Fixes vulnerabilities, doesn't change functionality
     Pending_Category_SECURITY Pending_Category = "SECURITY"
    // Fixes bugs/vulnerabilities, doesn't change functionality
     Pending_Category_FIX Pending_Category = "FIX"
    // Changes product functionality
     Pending_Category_UPDATE Pending_Category = "UPDATE"
    // Introduces new features, significantly changes product functionality
     Pending_Category_UPGRADE Pending_Category = "UPGRADE"
)

func (c Pending_Category) Pending_Category() bool {
    switch c {
        case Pending_Category_SECURITY:
            return true
        case Pending_Category_FIX:
            return true
        case Pending_Category_UPDATE:
            return true
        case Pending_Category_UPGRADE:
            return true
        default:
            return false
    }
}




// The ``UpdateType`` enumeration class defines update type
//
// <p> See {@link com.vmware.vapi.bindings.ApiEnumeration enumerated types description}.
 
type Pending_UpdateType string

const (
    // Fixes bugs/vulnerabilities, doesn't change functionality
     Pending_UpdateType_PATCH Pending_UpdateType = "PATCH"
    // Changes product functionality
     Pending_UpdateType_UPDATE Pending_UpdateType = "UPDATE"
    // Introduces new features, significantly changes product functionality
     Pending_UpdateType_UPGRADE Pending_UpdateType = "UPGRADE"
)

func (u Pending_UpdateType) Pending_UpdateType() bool {
    switch u {
        case Pending_UpdateType_PATCH:
            return true
        case Pending_UpdateType_UPDATE:
            return true
        case Pending_UpdateType_UPGRADE:
            return true
        default:
            return false
    }
}





// The ``Summary`` class contains basic information about the vCenter patch/update/upgrade
 type PendingSummary struct {
    // Identifier of the given vSphere update
    PendingUpdate string
    // Version of the vSphere update or patch
    Version string
    // Release date of the vSphere update or patch
    ReleaseDate time.Time
    // Severity of the issues fixed in the vSphere update or patch
    Severity Pending_SeverityType
    // Build number of the vCenter Release
    Build string
    // Type of the Release based on the current vCenter version
    UpdateType Pending_UpdateType
    // Category of the release based on features bundled on top of previous release
    Category Pending_Category
    // Flag to suggest a reboot after the release is applied
    RebootRequired bool
    // VAMI or ISO URL for update or upgrade execute phase redirection
    ExecuteURL url.URL
    // List of URI pointing to patch or update release notes
    ReleaseNotes []url.URL
}


// TODO some error doesn't have ErrorType. Resolve this later VAPI-3009
//


// The ``ListResult`` class contains information about the pending patch/updates for the given vCenter server.
 type PendingListResult struct {
    // Time when the software depo was last checked.
    LastCheckTime time.Time
    // Number of pending updates
    UpdateCount *int64
    // Number of pending upgrades
    UpgradeCount *int64
    // List of pending update details
    Updates []PendingSummary
}


// TODO some error doesn't have ErrorType. Resolve this later VAPI-3009
//


// The ``Info`` class contains detailed information about the vCenter patch/update.
 type PendingInfo struct {
    // Description of the vSphere update
    Description string
    // Identifier of the given vSphere update
    PendingUpdate string
    // Version of the vSphere update or patch
    Version string
    // Release date of the vSphere update or patch
    ReleaseDate time.Time
    // Severity of the issues fixed in the vSphere update or patch
    Severity Pending_SeverityType
    // Build number of the vCenter Release
    Build string
    // Type of the Release based on the current vCenter version
    UpdateType Pending_UpdateType
    // Category of the release based on features bundled on top of previous release
    Category Pending_Category
    // Flag to suggest a reboot after the release is applied
    RebootRequired bool
    // VAMI or ISO URL for update or upgrade execute phase redirection
    ExecuteURL url.URL
    // List of URI pointing to patch or update release notes
    ReleaseNotes []url.URL
}


// TODO some error doesn't have ErrorType. Resolve this later VAPI-3009
//





func pendingListInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func pendingListOutputType() bindings.BindingType {
    return bindings.NewReferenceType(PendingListResultBindingType)
}

func pendingListRestMetadata() protocol.OperationRestMetadata {
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
      "GET",
      "/vcenter/lcm/update/pending",
       resultHeaders,
       200,
       errorHeaders,
       map[string]int{"Unauthenticated": 401,"Error": 500})
}


func pendingGetInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["version"] = bindings.NewIdType([]string {"com.vmware.vcenter.lcm.update.pending"}, "")
    fieldNameMap["version"] = "Version"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func pendingGetOutputType() bindings.BindingType {
    return bindings.NewReferenceType(PendingInfoBindingType)
}

func pendingGetRestMetadata() protocol.OperationRestMetadata {
    paramsTypeMap := map[string]bindings.BindingType{}
    pathParams := map[string]string{}
    queryParams := map[string]string{}
    headerParams := map[string]string{}
    paramsTypeMap["version"] = bindings.NewIdType([]string {"com.vmware.vcenter.lcm.update.pending"}, "")
    paramsTypeMap["version"] = bindings.NewIdType([]string {"com.vmware.vcenter.lcm.update.pending"}, "")
    pathParams["version"] = "version"
    resultHeaders := map[string]string{}
    errorHeaders := map[string]string{}
    errorHeaders["Unauthenticated.challenge"] = "WWW-Authenticate"
    return protocol.NewOperationRestMetadata(
      paramsTypeMap,
      pathParams,
      queryParams,
      headerParams,
      "",
      "GET",
      "/vcenter/lcm/update/pending/{version}",
       resultHeaders,
       200,
       errorHeaders,
       map[string]int{"Unauthenticated": 401,"NotFound": 404,"Error": 500})
}



func PendingSummaryBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["pending_update"] = bindings.NewIdType([]string {"com.vmware.vcenter.lcm.update.pending"}, "")
    fieldNameMap["pending_update"] = "PendingUpdate"
    fields["version"] = bindings.NewStringType()
    fieldNameMap["version"] = "Version"
    fields["release_date"] = bindings.NewDateTimeType()
    fieldNameMap["release_date"] = "ReleaseDate"
    fields["severity"] = bindings.NewEnumType("com.vmware.vcenter.lcm.update.pending.severity_type", reflect.TypeOf(Pending_SeverityType(Pending_SeverityType_CRITICAL)))
    fieldNameMap["severity"] = "Severity"
    fields["build"] = bindings.NewStringType()
    fieldNameMap["build"] = "Build"
    fields["update_type"] = bindings.NewEnumType("com.vmware.vcenter.lcm.update.pending.update_type", reflect.TypeOf(Pending_UpdateType(Pending_UpdateType_PATCH)))
    fieldNameMap["update_type"] = "UpdateType"
    fields["category"] = bindings.NewEnumType("com.vmware.vcenter.lcm.update.pending.category", reflect.TypeOf(Pending_Category(Pending_Category_SECURITY)))
    fieldNameMap["category"] = "Category"
    fields["reboot_required"] = bindings.NewBooleanType()
    fieldNameMap["reboot_required"] = "RebootRequired"
    fields["execute_URL"] = bindings.NewUriType()
    fieldNameMap["execute_URL"] = "ExecuteURL"
    fields["release_notes"] = bindings.NewListType(bindings.NewUriType(), reflect.TypeOf([]url.URL{}))
    fieldNameMap["release_notes"] = "ReleaseNotes"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.lcm.update.pending.summary",fields, reflect.TypeOf(PendingSummary{}), fieldNameMap, validators)
}

func PendingListResultBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["last_check_time"] = bindings.NewDateTimeType()
    fieldNameMap["last_check_time"] = "LastCheckTime"
    fields["update_count"] = bindings.NewOptionalType(bindings.NewIntegerType())
    fieldNameMap["update_count"] = "UpdateCount"
    fields["upgrade_count"] = bindings.NewOptionalType(bindings.NewIntegerType())
    fieldNameMap["upgrade_count"] = "UpgradeCount"
    fields["updates"] = bindings.NewListType(bindings.NewReferenceType(PendingSummaryBindingType), reflect.TypeOf([]PendingSummary{}))
    fieldNameMap["updates"] = "Updates"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.lcm.update.pending.list_result",fields, reflect.TypeOf(PendingListResult{}), fieldNameMap, validators)
}

func PendingInfoBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["description"] = bindings.NewStringType()
    fieldNameMap["description"] = "Description"
    fields["pending_update"] = bindings.NewIdType([]string {"com.vmware.vcenter.lcm.update.pending"}, "")
    fieldNameMap["pending_update"] = "PendingUpdate"
    fields["version"] = bindings.NewStringType()
    fieldNameMap["version"] = "Version"
    fields["release_date"] = bindings.NewDateTimeType()
    fieldNameMap["release_date"] = "ReleaseDate"
    fields["severity"] = bindings.NewEnumType("com.vmware.vcenter.lcm.update.pending.severity_type", reflect.TypeOf(Pending_SeverityType(Pending_SeverityType_CRITICAL)))
    fieldNameMap["severity"] = "Severity"
    fields["build"] = bindings.NewStringType()
    fieldNameMap["build"] = "Build"
    fields["update_type"] = bindings.NewEnumType("com.vmware.vcenter.lcm.update.pending.update_type", reflect.TypeOf(Pending_UpdateType(Pending_UpdateType_PATCH)))
    fieldNameMap["update_type"] = "UpdateType"
    fields["category"] = bindings.NewEnumType("com.vmware.vcenter.lcm.update.pending.category", reflect.TypeOf(Pending_Category(Pending_Category_SECURITY)))
    fieldNameMap["category"] = "Category"
    fields["reboot_required"] = bindings.NewBooleanType()
    fieldNameMap["reboot_required"] = "RebootRequired"
    fields["execute_URL"] = bindings.NewUriType()
    fieldNameMap["execute_URL"] = "ExecuteURL"
    fields["release_notes"] = bindings.NewListType(bindings.NewUriType(), reflect.TypeOf([]url.URL{}))
    fieldNameMap["release_notes"] = "ReleaseNotes"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.lcm.update.pending.info",fields, reflect.TypeOf(PendingInfo{}), fieldNameMap, validators)
}


