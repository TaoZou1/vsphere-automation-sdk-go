/* Copyright © 2019 VMware, Inc. All Rights Reserved.
     SPDX-License-Identifier: BSD-2-Clause */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Data type definitions file for service: HealthCheckSettings.
 * Includes binding types of a structures and enumerations defined in the service.
 * Shared by client-side stubs and server-side skeletons to ensure type
 * compatibility.
 */

package appliance

import (
    "reflect"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/bindings"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/data"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/protocol"
)




// ``SettingSpec`` class contains specification of vCenter Server health and its corresponding state. **Warning:** This class is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
 type HealthCheckSettingsSettingSpec struct {
    // Health check state of database during manual backup with this option set, the health status of the particular health check will be reported. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    DbHealthCheckStateManualBackup bool
    // Health check state of database during schedule backup with this option set, the health status of the particular health check will be reported. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    DbHealthCheckStateScheduledBackup bool
}



func (HealthCheckSettingsSettingSpec HealthCheckSettingsSettingSpec) Error() string {
    return "com.vmware.appliance.setting_spec"
}



 type HealthCheckSettingsUpdateSpec struct {
    // Health check state of database during manual backup with this option set, the health status of the particular health check will be enabled. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    DbHealthCheckStateManualBackup *bool
    // Health check state of database during schedule backup with this option set, the health status of the particular health check will be enabled. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    DbHealthCheckStateScheduledBackup *bool
}



func (HealthCheckSettingsUpdateSpec HealthCheckSettingsUpdateSpec) Error() string {
    return "com.vmware.appliance.update_spec"
}






func healthCheckSettingsGetInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func healthCheckSettingsGetOutputType() bindings.BindingType {
    return bindings.NewReferenceType(HealthCheckSettingsSettingSpecBindingType)
}

func healthCheckSettingsGetRestMetadata() protocol.OperationRestMetadata {
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


func healthCheckSettingsUpdateInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["health_settings"] = bindings.NewReferenceType(HealthCheckSettingsUpdateSpecBindingType)
    fieldNameMap["health_settings"] = "HealthSettings"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func healthCheckSettingsUpdateOutputType() bindings.BindingType {
    return bindings.NewVoidType()
}

func healthCheckSettingsUpdateRestMetadata() protocol.OperationRestMetadata {
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
       map[string]int{"Error": 500,"InvalidArgument": 400})
}



func HealthCheckSettingsSettingSpecBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["db_health_check_state_manual_backup"] = bindings.NewBooleanType()
    fieldNameMap["db_health_check_state_manual_backup"] = "DbHealthCheckStateManualBackup"
    fields["db_health_check_state_scheduled_backup"] = bindings.NewBooleanType()
    fieldNameMap["db_health_check_state_scheduled_backup"] = "DbHealthCheckStateScheduledBackup"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.appliance.health_check_settings.setting_spec",fields, reflect.TypeOf(HealthCheckSettingsSettingSpec{}), fieldNameMap, validators)
}

func HealthCheckSettingsUpdateSpecBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["db_health_check_state_manual_backup"] = bindings.NewOptionalType(bindings.NewBooleanType())
    fieldNameMap["db_health_check_state_manual_backup"] = "DbHealthCheckStateManualBackup"
    fields["db_health_check_state_scheduled_backup"] = bindings.NewOptionalType(bindings.NewBooleanType())
    fieldNameMap["db_health_check_state_scheduled_backup"] = "DbHealthCheckStateScheduledBackup"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.appliance.health_check_settings.update_spec",fields, reflect.TypeOf(HealthCheckSettingsUpdateSpec{}), fieldNameMap, validators)
}


