/* Copyright © 2019 VMware, Inc. All Rights Reserved.
     SPDX-License-Identifier: BSD-2-Clause */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Data type definitions file for service: Update.
 * Includes binding types of a structures and enumerations defined in the service.
 * Shared by client-side stubs and server-side skeletons to ensure type
 * compatibility.
 */

package system

import (
    "reflect"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/bindings"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/data"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/protocol"
)



// ``AutoUpdateNotification`` enumeration class Defines state for automatic update notification.
//
// <p> See {@link com.vmware.vapi.bindings.ApiEnumeration enumerated types description}.
 
type Update_AutoUpdateNotification string

const (
    // Automatic update notification is disabled. Disable periodically query the configured url for updates.
     Update_AutoUpdateNotification_disabled Update_AutoUpdateNotification = "disabled"
    // Automatic update notification is enabled. Enable periodically query the configured url for updates.
     Update_AutoUpdateNotification_enabled Update_AutoUpdateNotification = "enabled"
)

func (a Update_AutoUpdateNotification) Update_AutoUpdateNotification() bool {
    switch a {
        case Update_AutoUpdateNotification_disabled:
            return true
        case Update_AutoUpdateNotification_enabled:
            return true
        default:
            return false
    }
}




// ``UpdateDay`` enumeration class Defines days to query for updates.
//
// <p> See {@link com.vmware.vapi.bindings.ApiEnumeration enumerated types description}.
 
type Update_UpdateDay string

const (
    // query for updates on Monday
     Update_UpdateDay_Monday Update_UpdateDay = "Monday"
    // query for updates on Tuesday
     Update_UpdateDay_Tuesday Update_UpdateDay = "Tuesday"
    // query for updates on Friday
     Update_UpdateDay_Friday Update_UpdateDay = "Friday"
    // query for updates on Wednesday
     Update_UpdateDay_Wednesday Update_UpdateDay = "Wednesday"
    // query for updates on Thursday
     Update_UpdateDay_Thursday Update_UpdateDay = "Thursday"
    // query for updates on Saturday
     Update_UpdateDay_Saturday Update_UpdateDay = "Saturday"
    // query for updates on Sunday
     Update_UpdateDay_Sunday Update_UpdateDay = "Sunday"
    // query for updates everyday
     Update_UpdateDay_Everyday Update_UpdateDay = "Everyday"
)

func (u Update_UpdateDay) Update_UpdateDay() bool {
    switch u {
        case Update_UpdateDay_Monday:
            return true
        case Update_UpdateDay_Tuesday:
            return true
        case Update_UpdateDay_Friday:
            return true
        case Update_UpdateDay_Wednesday:
            return true
        case Update_UpdateDay_Thursday:
            return true
        case Update_UpdateDay_Saturday:
            return true
        case Update_UpdateDay_Sunday:
            return true
        case Update_UpdateDay_Everyday:
            return true
        default:
            return false
    }
}





// ``UpdateStructSet`` class Structure to set url update repository.
 type UpdateUpdateStructSet struct {
    // Current appliance update repository URL. Enter "default" to reset the url to the default url.
    CurrentURL string
    // Check for update at the pre-configured repository URL.
    CheckUpdates Update_AutoUpdateNotification
    // time to query for updates Format: HH:MM:SS Military (24 hour) Time Format
    Time string
    // day to query for updates
    Day Update_UpdateDay
    // username for the url update repository
    Username string
    // password for the url update repository
    Password string
}


// TODO some error doesn't have ErrorType. Resolve this later VAPI-3009
//


// ``UpdateStructGet`` class Structure to get url update repository.
 type UpdateUpdateStructGet struct {
    // Current appliance update repository URL.
    CurrentURL string
    // Default appliance update repository URL.
    DefaultURL string
    // Check for update at the pre-configured repository URL.
    CheckUpdates Update_AutoUpdateNotification
    // time to query for updates Format: HH:MM:SS Military (24 hour) Time Format
    Time string
    // day to query for updates
    Day Update_UpdateDay
    // timestamp of latest update installation
    LatestUpdateInstallTime string
    // timestamp of latest query to update repository
    LatestUpdateQueryTime string
    // username for the url update repository
    Username string
    // password for the url update repository
    Password string
}


// TODO some error doesn't have ErrorType. Resolve this later VAPI-3009
//





func updateSetInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["config"] = bindings.NewReferenceType(UpdateUpdateStructSetBindingType)
    fieldNameMap["config"] = "Config"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func updateSetOutputType() bindings.BindingType {
    return bindings.NewVoidType()
}

func updateSetRestMetadata() protocol.OperationRestMetadata {
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


func updateGetInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func updateGetOutputType() bindings.BindingType {
    return bindings.NewReferenceType(UpdateUpdateStructGetBindingType)
}

func updateGetRestMetadata() protocol.OperationRestMetadata {
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



func UpdateUpdateStructSetBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["current_URL"] = bindings.NewStringType()
    fieldNameMap["current_URL"] = "CurrentURL"
    fields["check_updates"] = bindings.NewEnumType("com.vmware.appliance.techpreview.system.update.auto_update_notification", reflect.TypeOf(Update_AutoUpdateNotification(Update_AutoUpdateNotification_disabled)))
    fieldNameMap["check_updates"] = "CheckUpdates"
    fields["time"] = bindings.NewStringType()
    fieldNameMap["time"] = "Time"
    fields["day"] = bindings.NewEnumType("com.vmware.appliance.techpreview.system.update.update_day", reflect.TypeOf(Update_UpdateDay(Update_UpdateDay_Monday)))
    fieldNameMap["day"] = "Day"
    fields["username"] = bindings.NewStringType()
    fieldNameMap["username"] = "Username"
    fields["password"] = bindings.NewSecretType()
    fieldNameMap["password"] = "Password"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.appliance.techpreview.system.update.update_struct_set",fields, reflect.TypeOf(UpdateUpdateStructSet{}), fieldNameMap, validators)
}

func UpdateUpdateStructGetBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["current_URL"] = bindings.NewStringType()
    fieldNameMap["current_URL"] = "CurrentURL"
    fields["default_URL"] = bindings.NewStringType()
    fieldNameMap["default_URL"] = "DefaultURL"
    fields["check_updates"] = bindings.NewEnumType("com.vmware.appliance.techpreview.system.update.auto_update_notification", reflect.TypeOf(Update_AutoUpdateNotification(Update_AutoUpdateNotification_disabled)))
    fieldNameMap["check_updates"] = "CheckUpdates"
    fields["time"] = bindings.NewStringType()
    fieldNameMap["time"] = "Time"
    fields["day"] = bindings.NewEnumType("com.vmware.appliance.techpreview.system.update.update_day", reflect.TypeOf(Update_UpdateDay(Update_UpdateDay_Monday)))
    fieldNameMap["day"] = "Day"
    fields["latest_update_install_time"] = bindings.NewStringType()
    fieldNameMap["latest_update_install_time"] = "LatestUpdateInstallTime"
    fields["latest_update_query_time"] = bindings.NewStringType()
    fieldNameMap["latest_update_query_time"] = "LatestUpdateQueryTime"
    fields["username"] = bindings.NewStringType()
    fieldNameMap["username"] = "Username"
    fields["password"] = bindings.NewStringType()
    fieldNameMap["password"] = "Password"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.appliance.techpreview.system.update.update_struct_get",fields, reflect.TypeOf(UpdateUpdateStructGet{}), fieldNameMap, validators)
}


