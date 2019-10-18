/* Copyright © 2019 VMware, Inc. All Rights Reserved.
     SPDX-License-Identifier: BSD-2-Clause */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Data type definitions file for service: History.
 * Includes binding types of a structures and enumerations defined in the service.
 * Shared by client-side stubs and server-side skeletons to ensure type
 * compatibility.
 */

package update

import (
    "reflect"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/bindings/com/vmware/vapi/std"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/bindings"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/data"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/protocol"
    "net/url"
    "time"
)



// The ``Severity`` enumeration class defines the severity of the issues fixed in the update. **Warning:** This enumeration is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
//
// <p> See {@link com.vmware.vapi.bindings.ApiEnumeration enumerated types description}.
 
type HistorySeverity string

const (
    // Vulnerabilities that can be exploited by an unauthenticated attacker from the Internet or those that break the guest/host Operating System isolation. The exploitation results in the complete compromise of confidentiality, integrity, and availability of user data and/or processing resources without user interaction. Exploitation could be leveraged to propagate an Internet worm or execute arbitrary code between Virtual Machines and/or the Host Operating System. **Warning:** This constant field is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
     HistorySeverity_CRITICAL HistorySeverity = "CRITICAL"
    // Vulnerabilities that are not rated critical but whose exploitation results in the complete compromise of confidentiality and/or integrity of user data and/or processing resources through user assistance or by authenticated attackers. This rating also applies to those vulnerabilities which could lead to the complete compromise of availability when exploitation is by a remote unauthenticated attacker from the Internet or through a breach of virtual machine isolation. **Warning:** This constant field is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
     HistorySeverity_IMPORTANT HistorySeverity = "IMPORTANT"
    // Vulnerabilities where the ability to exploit is mitigated to a significant degree by configuration or difficulty of exploitation, but in certain deployment scenarios could still lead to the compromise of confidentiality, integrity, or availability of user data and/or processing resources. **Warning:** This constant field is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
     HistorySeverity_MODERATE HistorySeverity = "MODERATE"
)

func (s HistorySeverity) HistorySeverity() bool {
    switch s {
        case HistorySeverity_CRITICAL:
            return true
        case HistorySeverity_IMPORTANT:
            return true
        case HistorySeverity_MODERATE:
            return true
        default:
            return false
    }
}




// The ``Category`` enumeration class defines update type. **Warning:** This enumeration is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
//
// <p> See {@link com.vmware.vapi.bindings.ApiEnumeration enumerated types description}.
 
type HistoryCategory string

const (
    // Fixes vulnerabilities, doesn't change functionality. **Warning:** This constant field is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
     HistoryCategory_SECURITY HistoryCategory = "SECURITY"
    // Fixes bugs/vulnerabilities, doesn't change functionality. **Warning:** This constant field is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
     HistoryCategory_FIX HistoryCategory = "FIX"
    // Changes product functionality. **Warning:** This constant field is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
     HistoryCategory_UPDATE HistoryCategory = "UPDATE"
    // Introduces new features, significantly changes product functionality. **Warning:** This constant field is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
     HistoryCategory_UPGRADE HistoryCategory = "UPGRADE"
)

func (c HistoryCategory) HistoryCategory() bool {
    switch c {
        case HistoryCategory_SECURITY:
            return true
        case HistoryCategory_FIX:
            return true
        case HistoryCategory_UPDATE:
            return true
        case HistoryCategory_UPGRADE:
            return true
        default:
            return false
    }
}






// The ``Summary`` class contains information about the installed updates. **Warning:** This class is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
 type HistorySummary struct {
    // Name of the update. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    Name string
    // Installation date and time. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    InstallDate time.Time
    // Version of the update. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    Version string
    // The severity of the patch. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    Severity HistorySeverity
    // Update release date. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    ReleaseDate time.Time
    // Update category. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    UpdateType HistoryCategory
}







// The ``Info`` class contains detailed information about the installed updates. **Warning:** This class is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
 type HistoryInfo struct {
    // Name of the update. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    Name string
    // Installation date and time. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    InstallDate time.Time
    // Product name of the update. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    Product string
    // Name of vendor who provided the update. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    Vendor string
    // URL for the knowledge base article describing this update. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    KnowledgeBase url.URL
    // URL of the update repository from which this update was downloaded. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    UpdateRepo url.URL
    // The build number of the update. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    Build int64
    // Time when the EULA was accepted for the update. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    EulaAcceptTime time.Time
    // Tags associated with update (CVE Score). **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    Tags *string
    // Description of the update. The short information what this update is. E.g. "Update2 for vCenter Server Appliance 6.5". This property was added in vSphere API 6.7.
    Description std.LocalizableMessage
    // Update priority. This property was added in vSphere API 6.7.
    Priority CommonInfo_Priority
    // Update severity. This property was added in vSphere API 6.7.
    Severity CommonInfo_Severity
    // Update category. This property was added in vSphere API 6.7.
    UpdateType CommonInfo_Category
    // Update release date. This property was added in vSphere API 6.7.
    ReleaseDate time.Time
    // Flag indicating whether reboot is required after update. This property was added in vSphere API 6.7.
    RebootRequired bool
    // Download Size of update in Megabytes. This property was added in vSphere API 6.7.
    Size int64
}









func historyListInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func historyListOutputType() bindings.BindingType {
    return bindings.NewListType(bindings.NewReferenceType(HistorySummaryBindingType), reflect.TypeOf([]HistorySummary{}))
}

func historyListRestMetadata() protocol.OperationRestMetadata {
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
       map[string]int{"Error": 500,"Unauthenticated": 401,"Unauthorized": 403,"NotAllowedInCurrentState": 400})
}


func historyGetInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["version"] = bindings.NewIdType([]string {"com.vmware.appliance.update.history"}, "")
    fieldNameMap["version"] = "Version"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func historyGetOutputType() bindings.BindingType {
    return bindings.NewReferenceType(HistoryInfoBindingType)
}

func historyGetRestMetadata() protocol.OperationRestMetadata {
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
       map[string]int{"NotFound": 404,"Error": 500,"Unauthenticated": 401,"Unauthorized": 403,"NotAllowedInCurrentState": 400})
}



func HistorySummaryBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["name"] = bindings.NewStringType()
    fieldNameMap["name"] = "Name"
    fields["install_date"] = bindings.NewDateTimeType()
    fieldNameMap["install_date"] = "InstallDate"
    fields["version"] = bindings.NewIdType([]string {"com.vmware.appliance.update.history"}, "")
    fieldNameMap["version"] = "Version"
    fields["severity"] = bindings.NewEnumType("com.vmware.appliance.update.history.severity", reflect.TypeOf(HistorySeverity(HistorySeverity_CRITICAL)))
    fieldNameMap["severity"] = "Severity"
    fields["release_date"] = bindings.NewDateTimeType()
    fieldNameMap["release_date"] = "ReleaseDate"
    fields["update_type"] = bindings.NewEnumType("com.vmware.appliance.update.history.category", reflect.TypeOf(HistoryCategory(HistoryCategory_SECURITY)))
    fieldNameMap["update_type"] = "UpdateType"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.appliance.update.history.summary",fields, reflect.TypeOf(HistorySummary{}), fieldNameMap, validators)
}

func HistoryInfoBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["name"] = bindings.NewStringType()
    fieldNameMap["name"] = "Name"
    fields["install_date"] = bindings.NewDateTimeType()
    fieldNameMap["install_date"] = "InstallDate"
    fields["product"] = bindings.NewStringType()
    fieldNameMap["product"] = "Product"
    fields["vendor"] = bindings.NewStringType()
    fieldNameMap["vendor"] = "Vendor"
    fields["knowledge_base"] = bindings.NewUriType()
    fieldNameMap["knowledge_base"] = "KnowledgeBase"
    fields["update_repo"] = bindings.NewUriType()
    fieldNameMap["update_repo"] = "UpdateRepo"
    fields["build"] = bindings.NewIntegerType()
    fieldNameMap["build"] = "Build"
    fields["eula_accept_time"] = bindings.NewDateTimeType()
    fieldNameMap["eula_accept_time"] = "EulaAcceptTime"
    fields["tags"] = bindings.NewOptionalType(bindings.NewStringType())
    fieldNameMap["tags"] = "Tags"
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
    return bindings.NewStructType("com.vmware.appliance.update.history.info",fields, reflect.TypeOf(HistoryInfo{}), fieldNameMap, validators)
}


