/* Copyright © 2019 VMware, Inc. All Rights Reserved.
     SPDX-License-Identifier: BSD-2-Clause */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Data type definitions file for service: Routes.
 * Includes binding types of a structures and enumerations defined in the service.
 * Shared by client-side stubs and server-side skeletons to ensure type
 * compatibility.
 */

package networking

import (
    "reflect"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/bindings/com/vmware/vapi/std"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/bindings"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/data"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/protocol"
)



// ``Mode`` enumeration class Mode of the routes static or autoconfigured. **Warning:** This enumeration is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
//
// <p> See {@link com.vmware.vapi.bindings.ApiEnumeration enumerated types description}.
 
type Routes_Mode string

const (
    // Route installed by administrator. **Warning:** This constant field is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
     Routes_Mode_STATIC Routes_Mode = "STATIC"
    // Autoconfigured route. **Warning:** This constant field is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
     Routes_Mode_AUTO Routes_Mode = "AUTO"
)

func (m Routes_Mode) Routes_Mode() bool {
    switch m {
        case Routes_Mode_STATIC:
            return true
        case Routes_Mode_AUTO:
            return true
        default:
            return false
    }
}




// ``OverallStatus`` enumeration class Health indicator. **Warning:** This enumeration is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
//
// <p> See {@link com.vmware.vapi.bindings.ApiEnumeration enumerated types description}.
 
type Routes_OverallStatus string

const (
    // In case data has more than one test, this indicates not all tests were successful. **Warning:** This constant field is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
     Routes_OverallStatus_PARTIAL_FAILED Routes_OverallStatus = "PARTIAL_FAILED"
    // All tests were successful for given data. **Warning:** This constant field is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
     Routes_OverallStatus_SUCCESSFUL Routes_OverallStatus = "SUCCESSFUL"
    // All tests failed for given data. **Warning:** This constant field is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
     Routes_OverallStatus_FAILED Routes_OverallStatus = "FAILED"
)

func (o Routes_OverallStatus) Routes_OverallStatus() bool {
    switch o {
        case Routes_OverallStatus_PARTIAL_FAILED:
            return true
        case Routes_OverallStatus_SUCCESSFUL:
            return true
        case Routes_OverallStatus_FAILED:
            return true
        default:
            return false
    }
}




// ``Status`` enumeration class Individual test result. **Warning:** This enumeration is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
//
// <p> See {@link com.vmware.vapi.bindings.ApiEnumeration enumerated types description}.
 
type Routes_Status string

const (
    // message indicates the test failed. **Warning:** This constant field is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
     Routes_Status_FAILURE Routes_Status = "FAILURE"
    // message indicates that the test was successful. **Warning:** This constant field is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
     Routes_Status_SUCCESS Routes_Status = "SUCCESS"
)

func (s Routes_Status) Routes_Status() bool {
    switch s {
        case Routes_Status_FAILURE:
            return true
        case Routes_Status_SUCCESS:
            return true
        default:
            return false
    }
}





// ``Config`` class Structure that describes how routing is performed for a particular destination and prefix. A destination/prefix of 0.0.0.0/0 ( for IPv4) or ::/0 (for IPv6) refers to the default gateway. **Warning:** This class is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
 type RoutesConfig struct {
    // Destination address that defines this route. An empty string indicates 0.0.0.0. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    Destination *string
    // Destination CIDR prefix that defines this route. See http://www.oav.net/mirrors/cidr.html for netmask-to-prefix conversion. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    Prefix int64
    // Gateway address. An empty string indicates 0.0.0.0. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    Gateway *string
    // Output device interface, for example, "nic0". An empty string indicates that the rule applies to all interfaces(\*). **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    InterfaceName *string
}



//


// ``Info`` class Structure that describes how routing is performed for a particular destination and prefix. A destination/prefix of 0.0.0.0/0 ( for IPv4) or ::/0 (for IPv6) refers to the default gateway. **Warning:** This class is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
 type RoutesInfo struct {
    // Static provides information about installation of the route. True indicates the route was installed by the administrator. False indicates the route was autoconfigured. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    Mode Routes_Mode
    // Destination address that defines this route. An empty string indicates 0.0.0.0. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    Destination *string
    // Destination CIDR prefix that defines this route. See http://www.oav.net/mirrors/cidr.html for netmask-to-prefix conversion. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    Prefix int64
    // Gateway address. An empty string indicates 0.0.0.0. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    Gateway *string
    // Output device interface, for example, "nic0". An empty string indicates that the rule applies to all interfaces(\*). **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    InterfaceName *string
}



//


// ``RouteStatus`` class Test result and message. **Warning:** This class is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
 type RoutesRouteStatus struct {
    // message. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    Message std.LocalizableMessage
    // result of the test. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    Status Routes_Status
}



//


// ``TestStatus`` class Overall test result. **Warning:** This class is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
 type RoutesTestStatus struct {
    // Overall status of tests run. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    OverallStatus Routes_OverallStatus
    // messages. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    RouteStatus []RoutesRouteStatus
}



//





func routesTestInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["gateways"] = bindings.NewListType(bindings.NewStringType(), reflect.TypeOf([]string{}))
    fieldNameMap["gateways"] = "Gateways"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func routesTestOutputType() bindings.BindingType {
    return bindings.NewReferenceType(RoutesTestStatusBindingType)
}

func routesTestRestMetadata() protocol.OperationRestMetadata {
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


func routesSetInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["routes"] = bindings.NewListType(bindings.NewReferenceType(RoutesConfigBindingType), reflect.TypeOf([]RoutesConfig{}))
    fieldNameMap["routes"] = "Routes"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func routesSetOutputType() bindings.BindingType {
    return bindings.NewVoidType()
}

func routesSetRestMetadata() protocol.OperationRestMetadata {
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


func routesGetInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func routesGetOutputType() bindings.BindingType {
    return bindings.NewListType(bindings.NewReferenceType(RoutesInfoBindingType), reflect.TypeOf([]RoutesInfo{}))
}

func routesGetRestMetadata() protocol.OperationRestMetadata {
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



func RoutesConfigBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["destination"] = bindings.NewOptionalType(bindings.NewStringType())
    fieldNameMap["destination"] = "Destination"
    fields["prefix"] = bindings.NewIntegerType()
    fieldNameMap["prefix"] = "Prefix"
    fields["gateway"] = bindings.NewOptionalType(bindings.NewStringType())
    fieldNameMap["gateway"] = "Gateway"
    fields["interface_name"] = bindings.NewOptionalType(bindings.NewStringType())
    fieldNameMap["interface_name"] = "InterfaceName"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.appliance.networking.routes.config",fields, reflect.TypeOf(RoutesConfig{}), fieldNameMap, validators)
}

func RoutesInfoBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["mode"] = bindings.NewEnumType("com.vmware.appliance.networking.routes.mode", reflect.TypeOf(Routes_Mode(Routes_Mode_STATIC)))
    fieldNameMap["mode"] = "Mode"
    fields["destination"] = bindings.NewOptionalType(bindings.NewStringType())
    fieldNameMap["destination"] = "Destination"
    fields["prefix"] = bindings.NewIntegerType()
    fieldNameMap["prefix"] = "Prefix"
    fields["gateway"] = bindings.NewOptionalType(bindings.NewStringType())
    fieldNameMap["gateway"] = "Gateway"
    fields["interface_name"] = bindings.NewOptionalType(bindings.NewStringType())
    fieldNameMap["interface_name"] = "InterfaceName"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.appliance.networking.routes.info",fields, reflect.TypeOf(RoutesInfo{}), fieldNameMap, validators)
}

func RoutesRouteStatusBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["message"] = bindings.NewReferenceType(std.LocalizableMessageBindingType)
    fieldNameMap["message"] = "Message"
    fields["status"] = bindings.NewEnumType("com.vmware.appliance.networking.routes.status", reflect.TypeOf(Routes_Status(Routes_Status_FAILURE)))
    fieldNameMap["status"] = "Status"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.appliance.networking.routes.route_status",fields, reflect.TypeOf(RoutesRouteStatus{}), fieldNameMap, validators)
}

func RoutesTestStatusBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["overall_status"] = bindings.NewEnumType("com.vmware.appliance.networking.routes.overall_status", reflect.TypeOf(Routes_OverallStatus(Routes_OverallStatus_PARTIAL_FAILED)))
    fieldNameMap["overall_status"] = "OverallStatus"
    fields["route_status"] = bindings.NewListType(bindings.NewReferenceType(RoutesRouteStatusBindingType), reflect.TypeOf([]RoutesRouteStatus{}))
    fieldNameMap["route_status"] = "RouteStatus"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.appliance.networking.routes.test_status",fields, reflect.TypeOf(RoutesTestStatus{}), fieldNameMap, validators)
}


