/* Copyright © 2019 VMware, Inc. All Rights Reserved.
     SPDX-License-Identifier: BSD-2-Clause */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Data type definitions file for service: Environment.
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



// The ``ManagedBy`` enumeration class defines the possible entities managing the appliance.
//
// <p> See {@link com.vmware.vapi.bindings.ApiEnumeration enumerated types description}.
 
type Environment_ManagedBy string

const (
    // Managed by VMware
     Environment_ManagedBy_VMWARE Environment_ManagedBy = "VMWARE"
    // Managed by the user
     Environment_ManagedBy_USER Environment_ManagedBy = "USER"
)

func (m Environment_ManagedBy) Environment_ManagedBy() bool {
    switch m {
        case Environment_ManagedBy_VMWARE:
            return true
        case Environment_ManagedBy_USER:
            return true
        default:
            return false
    }
}




// The ``DeployedBy`` enumeration class defines the possible entities deploying the appliance.
//
// <p> See {@link com.vmware.vapi.bindings.ApiEnumeration enumerated types description}.
 
type Environment_DeployedBy string

const (
    // Deployed by VMware
     Environment_DeployedBy_VMWARE Environment_DeployedBy = "VMWARE"
    // Deployed by the user
     Environment_DeployedBy_USER Environment_DeployedBy = "USER"
)

func (d Environment_DeployedBy) Environment_DeployedBy() bool {
    switch d {
        case Environment_DeployedBy_VMWARE:
            return true
        case Environment_DeployedBy_USER:
            return true
        default:
            return false
    }
}




// The ``Provider`` enumeration class defines the possible providers. Now it has only AWS, in the future we foresee soft-layer, google cloud, azure, etc.
//
// <p> See {@link com.vmware.vapi.bindings.ApiEnumeration enumerated types description}.
 
type Environment_Provider string

const (
    // The appliance is located in AWS-backed VMC
     Environment_Provider_AWS Environment_Provider = "AWS"
    // The appliance location is UNKNOWN
     Environment_Provider_UNKNOWN Environment_Provider = "UNKNOWN"
)

func (p Environment_Provider) Environment_Provider() bool {
    switch p {
        case Environment_Provider_AWS:
            return true
        case Environment_Provider_UNKNOWN:
            return true
        default:
            return false
    }
}





// The ``Display`` class contains information what needs to be displayed in the user interface.
 type EnvironmentDisplay struct {
    // Display database monitoring.
    DatabaseMonitoring bool
    // Display SSO status.
    SsoStatus bool
}



func (EnvironmentDisplay EnvironmentDisplay) Error() string {
    return "com.vmware.appliance.display"
}



// The ``Info`` class contains the information about the appliance environment.
 type EnvironmentInfo struct {
    // Who manages this appliance.
    ManagedBy Environment_ManagedBy
    // Who deployed this appliance.
    DeployedBy Environment_DeployedBy
    // The cloud provider where this appliance is running.
    Provider Environment_Provider
    // What UI should display.
    Display EnvironmentDisplay
}



func (EnvironmentInfo EnvironmentInfo) Error() string {
    return "com.vmware.appliance.info"
}



// The ``Config`` class describes the configurable settings for the appliance environment.
 type EnvironmentConfig struct {
    // Who manages this appliance.
    ManagedBy Environment_ManagedBy
    // Who have deployed this appliance.
    DeployedBy Environment_DeployedBy
    // The cloud provider where this appliance is running.
    Provider Environment_Provider
    // What UI should display.
    Display EnvironmentDisplay
}



func (EnvironmentConfig EnvironmentConfig) Error() string {
    return "com.vmware.appliance.config"
}






func environmentSetInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["config"] = bindings.NewReferenceType(EnvironmentConfigBindingType)
    fieldNameMap["config"] = "Config"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func environmentSetOutputType() bindings.BindingType {
    return bindings.NewVoidType()
}

func environmentSetRestMetadata() protocol.OperationRestMetadata {
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


func environmentGetInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func environmentGetOutputType() bindings.BindingType {
    return bindings.NewReferenceType(EnvironmentInfoBindingType)
}

func environmentGetRestMetadata() protocol.OperationRestMetadata {
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



func EnvironmentDisplayBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["database_monitoring"] = bindings.NewBooleanType()
    fieldNameMap["database_monitoring"] = "DatabaseMonitoring"
    fields["sso_status"] = bindings.NewBooleanType()
    fieldNameMap["sso_status"] = "SsoStatus"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.appliance.environment.display",fields, reflect.TypeOf(EnvironmentDisplay{}), fieldNameMap, validators)
}

func EnvironmentInfoBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["managed_by"] = bindings.NewEnumType("com.vmware.appliance.environment.managed_by", reflect.TypeOf(Environment_ManagedBy(Environment_ManagedBy_VMWARE)))
    fieldNameMap["managed_by"] = "ManagedBy"
    fields["deployed_by"] = bindings.NewEnumType("com.vmware.appliance.environment.deployed_by", reflect.TypeOf(Environment_DeployedBy(Environment_DeployedBy_VMWARE)))
    fieldNameMap["deployed_by"] = "DeployedBy"
    fields["provider"] = bindings.NewEnumType("com.vmware.appliance.environment.provider", reflect.TypeOf(Environment_Provider(Environment_Provider_AWS)))
    fieldNameMap["provider"] = "Provider"
    fields["display"] = bindings.NewReferenceType(EnvironmentDisplayBindingType)
    fieldNameMap["display"] = "Display"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.appliance.environment.info",fields, reflect.TypeOf(EnvironmentInfo{}), fieldNameMap, validators)
}

func EnvironmentConfigBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["managed_by"] = bindings.NewEnumType("com.vmware.appliance.environment.managed_by", reflect.TypeOf(Environment_ManagedBy(Environment_ManagedBy_VMWARE)))
    fieldNameMap["managed_by"] = "ManagedBy"
    fields["deployed_by"] = bindings.NewEnumType("com.vmware.appliance.environment.deployed_by", reflect.TypeOf(Environment_DeployedBy(Environment_DeployedBy_VMWARE)))
    fieldNameMap["deployed_by"] = "DeployedBy"
    fields["provider"] = bindings.NewEnumType("com.vmware.appliance.environment.provider", reflect.TypeOf(Environment_Provider(Environment_Provider_AWS)))
    fieldNameMap["provider"] = "Provider"
    fields["display"] = bindings.NewReferenceType(EnvironmentDisplayBindingType)
    fieldNameMap["display"] = "Display"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.appliance.environment.config",fields, reflect.TypeOf(EnvironmentConfig{}), fieldNameMap, validators)
}


