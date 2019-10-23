/* Copyright © 2019 VMware, Inc. All Rights Reserved.
     SPDX-License-Identifier: BSD-2-Clause */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Data type definitions file for service: Policy.
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
)



// The ``DayOfWeek`` enumeration class defines the set of days
//
// <p> See {@link com.vmware.vapi.bindings.ApiEnumeration enumerated types description}.
 
type Policy_DayOfWeek string

const (
    // Monday
     Policy_DayOfWeek_MONDAY Policy_DayOfWeek = "MONDAY"
    // Tuesday
     Policy_DayOfWeek_TUESDAY Policy_DayOfWeek = "TUESDAY"
    // Wednesday
     Policy_DayOfWeek_WEDNESDAY Policy_DayOfWeek = "WEDNESDAY"
    // Thursday
     Policy_DayOfWeek_THURSDAY Policy_DayOfWeek = "THURSDAY"
    // Friday
     Policy_DayOfWeek_FRIDAY Policy_DayOfWeek = "FRIDAY"
    // Saturday
     Policy_DayOfWeek_SATURDAY Policy_DayOfWeek = "SATURDAY"
    // Sunday
     Policy_DayOfWeek_SUNDAY Policy_DayOfWeek = "SUNDAY"
)

func (d Policy_DayOfWeek) Policy_DayOfWeek() bool {
    switch d {
        case Policy_DayOfWeek_MONDAY:
            return true
        case Policy_DayOfWeek_TUESDAY:
            return true
        case Policy_DayOfWeek_WEDNESDAY:
            return true
        case Policy_DayOfWeek_THURSDAY:
            return true
        case Policy_DayOfWeek_FRIDAY:
            return true
        case Policy_DayOfWeek_SATURDAY:
            return true
        case Policy_DayOfWeek_SUNDAY:
            return true
        default:
            return false
    }
}





// The ``Time`` class defines weekday and time the automatic check for new updates will be run
 type PolicyTime struct {
    // weekday to check for updates.
    Day Policy_DayOfWeek
    // Hour: 0-24
    Hour int64
    // Minute: 0-59
    Minute int64
}



func (PolicyTime PolicyTime) Error() string {
    return "com.vmware.appliance.update.time"
}



// The ``Info`` class defines automatic update checking and staging policy.
 type PolicyInfo struct {
    // Current appliance update custom repository URL.
    CustomURL *string
    // Current appliance update default repository URL.
    DefaultURL string
    // Username for the update repository
    Username *string
    // Schedule when the automatic check will be run.
    CheckSchedule []PolicyTime
    // Automatically stage the latest update if available.
    AutoStage bool
    // Is the appliance updated automatically. If map with bool value the appliance may ignore the check schedule or auto-stage settings.
    AutoUpdate bool
    // Whether API client should allow the user to start update manually
    ManualControl bool
}



func (PolicyInfo PolicyInfo) Error() string {
    return "com.vmware.appliance.update.info"
}



// The ``Config`` class defines automatic update checking and staging policy.
 type PolicyConfig struct {
    // Current appliance update repository URL.
    CustomURL *string
    // Username for the update repository
    Username *string
    // Password for the update repository
    Password *string
    // Schedule when the automatic check will be run.
    CheckSchedule []PolicyTime
    // Automatically stage the latest update if available.
    AutoStage bool
}



func (PolicyConfig PolicyConfig) Error() string {
    return "com.vmware.appliance.update.config"
}






func policyGetInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func policyGetOutputType() bindings.BindingType {
    return bindings.NewReferenceType(PolicyInfoBindingType)
}

func policyGetRestMetadata() protocol.OperationRestMetadata {
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
       map[string]int{"Error": 500,"Unauthenticated": 401,"Unauthorized": 403})
}


func policySetInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["policy"] = bindings.NewReferenceType(PolicyConfigBindingType)
    fieldNameMap["policy"] = "Policy"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func policySetOutputType() bindings.BindingType {
    return bindings.NewVoidType()
}

func policySetRestMetadata() protocol.OperationRestMetadata {
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
       map[string]int{"Error": 500,"Unauthenticated": 401,"Unauthorized": 403})
}



func PolicyTimeBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["day"] = bindings.NewEnumType("com.vmware.appliance.update.policy.day_of_week", reflect.TypeOf(Policy_DayOfWeek(Policy_DayOfWeek_MONDAY)))
    fieldNameMap["day"] = "Day"
    fields["hour"] = bindings.NewIntegerType()
    fieldNameMap["hour"] = "Hour"
    fields["minute"] = bindings.NewIntegerType()
    fieldNameMap["minute"] = "Minute"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.appliance.update.policy.time",fields, reflect.TypeOf(PolicyTime{}), fieldNameMap, validators)
}

func PolicyInfoBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["custom_URL"] = bindings.NewOptionalType(bindings.NewStringType())
    fieldNameMap["custom_URL"] = "CustomURL"
    fields["default_URL"] = bindings.NewStringType()
    fieldNameMap["default_URL"] = "DefaultURL"
    fields["username"] = bindings.NewOptionalType(bindings.NewStringType())
    fieldNameMap["username"] = "Username"
    fields["check_schedule"] = bindings.NewListType(bindings.NewReferenceType(PolicyTimeBindingType), reflect.TypeOf([]PolicyTime{}))
    fieldNameMap["check_schedule"] = "CheckSchedule"
    fields["auto_stage"] = bindings.NewBooleanType()
    fieldNameMap["auto_stage"] = "AutoStage"
    fields["auto_update"] = bindings.NewBooleanType()
    fieldNameMap["auto_update"] = "AutoUpdate"
    fields["manual_control"] = bindings.NewBooleanType()
    fieldNameMap["manual_control"] = "ManualControl"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.appliance.update.policy.info",fields, reflect.TypeOf(PolicyInfo{}), fieldNameMap, validators)
}

func PolicyConfigBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["custom_URL"] = bindings.NewOptionalType(bindings.NewStringType())
    fieldNameMap["custom_URL"] = "CustomURL"
    fields["username"] = bindings.NewOptionalType(bindings.NewStringType())
    fieldNameMap["username"] = "Username"
    fields["password"] = bindings.NewOptionalType(bindings.NewSecretType())
    fieldNameMap["password"] = "Password"
    fields["check_schedule"] = bindings.NewListType(bindings.NewReferenceType(PolicyTimeBindingType), reflect.TypeOf([]PolicyTime{}))
    fieldNameMap["check_schedule"] = "CheckSchedule"
    fields["auto_stage"] = bindings.NewBooleanType()
    fieldNameMap["auto_stage"] = "AutoStage"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.appliance.update.policy.config",fields, reflect.TypeOf(PolicyConfig{}), fieldNameMap, validators)
}


