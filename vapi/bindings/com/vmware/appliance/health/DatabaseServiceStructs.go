/* Copyright © 2019 VMware, Inc. All Rights Reserved.
     SPDX-License-Identifier: BSD-2-Clause */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Data type definitions file for service: Database.
 * Includes binding types of a structures and enumerations defined in the service.
 * Shared by client-side stubs and server-side skeletons to ensure type
 * compatibility.
 */

package health

import (
    "reflect"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/bindings/com/vmware/vapi/std"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/bindings"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/data"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/protocol"
)




// The ``Info`` class contains information about the health of the the database. **Warning:** This class is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
 type DatabaseInfo struct {
    // Database health status. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    Status DatabaseInfo_Status
    // Messages describing any issues with the database, along with their severity. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    Messages []DatabaseMessage
}



func (DatabaseInfo DatabaseInfo) Error() string {
    return "com.vmware.appliance.health.info"
}

    
    // The ``Status`` enumeration class describes the health of the database. **Warning:** This enumeration is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    //
    // <p> See {@link com.vmware.vapi.bindings.ApiEnumeration enumerated types description}.
     
    type DatabaseInfo_Status string

    const (
        // The database is corrupted and vCenter server functionality will be impacted. **Warning:** This constant field is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
         DatabaseInfo_Status_UNHEALTHY DatabaseInfo_Status = "UNHEALTHY"
        // The database has issues but the impact on vCenter Server is low. **Warning:** This constant field is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
         DatabaseInfo_Status_DEGRADED DatabaseInfo_Status = "DEGRADED"
        // The database is healthy. **Warning:** This constant field is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
         DatabaseInfo_Status_HEALTHY DatabaseInfo_Status = "HEALTHY"
    )

    func (s DatabaseInfo_Status) DatabaseInfo_Status() bool {
        switch s {
            case DatabaseInfo_Status_UNHEALTHY:
                return true
            case DatabaseInfo_Status_DEGRADED:
                return true
            case DatabaseInfo_Status_HEALTHY:
                return true
            default:
                return false
        }
    }



// The ``Message`` class contains a database health message along with its severity. **Warning:** This class is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
 type DatabaseMessage struct {
    // Severity of the message. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    Severity DatabaseMessage_Severity
    // Message describing the issue with the database. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    Message std.LocalizableMessage
}



func (DatabaseMessage DatabaseMessage) Error() string {
    return "com.vmware.appliance.health.message"
}

    
    // The ``MessageSeverity`` enumeration class defines the levels of severity for a message. **Warning:** This enumeration is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    //
    // <p> See {@link com.vmware.vapi.bindings.ApiEnumeration enumerated types description}.
     
    type DatabaseMessage_Severity string

    const (
        // Error message. **Warning:** This constant field is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
         DatabaseMessage_Severity_ERROR DatabaseMessage_Severity = "ERROR"
        // Warning message. **Warning:** This constant field is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
         DatabaseMessage_Severity_WARNING DatabaseMessage_Severity = "WARNING"
    )

    func (s DatabaseMessage_Severity) DatabaseMessage_Severity() bool {
        switch s {
            case DatabaseMessage_Severity_ERROR:
                return true
            case DatabaseMessage_Severity_WARNING:
                return true
            default:
                return false
        }
    }






func databaseGetInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func databaseGetOutputType() bindings.BindingType {
    return bindings.NewReferenceType(DatabaseInfoBindingType)
}

func databaseGetRestMetadata() protocol.OperationRestMetadata {
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
       map[string]int{"Error": 500})
}



func DatabaseInfoBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["status"] = bindings.NewEnumType("com.vmware.appliance.health.database.info.status", reflect.TypeOf(DatabaseInfo_Status(DatabaseInfo_Status_UNHEALTHY)))
    fieldNameMap["status"] = "Status"
    fields["messages"] = bindings.NewListType(bindings.NewReferenceType(DatabaseMessageBindingType), reflect.TypeOf([]DatabaseMessage{}))
    fieldNameMap["messages"] = "Messages"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.appliance.health.database.info",fields, reflect.TypeOf(DatabaseInfo{}), fieldNameMap, validators)
}

func DatabaseMessageBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["severity"] = bindings.NewEnumType("com.vmware.appliance.health.database.message.severity", reflect.TypeOf(DatabaseMessage_Severity(DatabaseMessage_Severity_ERROR)))
    fieldNameMap["severity"] = "Severity"
    fields["message"] = bindings.NewReferenceType(std.LocalizableMessageBindingType)
    fieldNameMap["message"] = "Message"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.appliance.health.database.message",fields, reflect.TypeOf(DatabaseMessage{}), fieldNameMap, validators)
}


