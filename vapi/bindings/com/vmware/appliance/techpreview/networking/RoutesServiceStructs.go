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
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/bindings"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/data"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/protocol"
)



// ``TestStatus`` enumeration class Health indicator
//
// <p> See {@link com.vmware.vapi.bindings.ApiEnumeration enumerated types description}.
 
type RoutesTestStatus string

const (
    // In case data has more than one test, this indicates not all tests were successful
     RoutesTestStatus_orange RoutesTestStatus = "orange"
    // All tests were successful for given data
     RoutesTestStatus_green RoutesTestStatus = "green"
    // All tests failed for given data
     RoutesTestStatus_red RoutesTestStatus = "red"
)

func (t RoutesTestStatus) RoutesTestStatus() bool {
    switch t {
        case RoutesTestStatus_orange:
            return true
        case RoutesTestStatus_green:
            return true
        case RoutesTestStatus_red:
            return true
        default:
            return false
    }
}




// ``MessageStatus`` enumeration class Individual test result
//
// <p> See {@link com.vmware.vapi.bindings.ApiEnumeration enumerated types description}.
 
type RoutesMessageStatus string

const (
    // message indicates the test failed.
     RoutesMessageStatus_failure RoutesMessageStatus = "failure"
    // message indicates that the test was successful.
     RoutesMessageStatus_success RoutesMessageStatus = "success"
)

func (m RoutesMessageStatus) RoutesMessageStatus() bool {
    switch m {
        case RoutesMessageStatus_failure:
            return true
        case RoutesMessageStatus_success:
            return true
        default:
            return false
    }
}






// ``RouteReadOnly`` class Structure that describes how routing is performed for a particular destination and prefix. A destination/prefix of 0.0.0.0/0 ( for IPv4) or ::/0 (for IPv6) refers to the default gateway.
 type RoutesRouteReadOnly struct {
    // Destination address that defines this route.
    Destination string
    // Destination CIDR prefix that defines this route. See http://www.oav.net/mirrors/cidr.html for netmask-to-prefix conversion.
    Prefix int64
    // Gateway address.
    Gateway string
    // Output device interface, for example, "nic0".
    InterfaceName string
    // Static provides information about installation of the route. True indicates the route was installed by the administrator. False indicates the route was autoconfigured
    IsStatic bool
}







// ``Route`` class Structure that describes how routing is performed for a particular destination and prefix. A destination/prefix of 0.0.0.0/0 ( for IPv4) or ::/0 (for IPv6) refers to the default gateway.
 type RoutesRoute struct {
    // Destination address that defines this route.
    Destination string
    // Destination CIDR prefix that defines this route. See http://www.oav.net/mirrors/cidr.html for netmask-to-prefix conversion.
    Prefix int64
    // Gateway address.
    Gateway string
    // Output device interface, for example, "nic0".
    InterfaceName string
}







// ``Message`` class Test result and message
 type RoutesMessage struct {
    // message
    Message string
    // result of the test
    Result RoutesMessageStatus
}







// ``TestStatusInfo`` class Overall test result
 type RoutesTestStatusInfo struct {
    // Overall status of tests run.
    Status RoutesTestStatus
    // messages
    Messages []RoutesMessage
}









func routesTestInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["gateways"] = bindings.NewListType(bindings.NewStringType(), reflect.TypeOf([]string{}))
    fieldNameMap["gateways"] = "Gateways"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func routesTestOutputType() bindings.BindingType {
    return bindings.NewReferenceType(RoutesTestStatusInfoBindingType)
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


func routesAddInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["route"] = bindings.NewReferenceType(RoutesRouteBindingType)
    fieldNameMap["route"] = "Route"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func routesAddOutputType() bindings.BindingType {
    return bindings.NewVoidType()
}

func routesAddRestMetadata() protocol.OperationRestMetadata {
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
    fields["routes"] = bindings.NewListType(bindings.NewReferenceType(RoutesRouteBindingType), reflect.TypeOf([]RoutesRoute{}))
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


func routesListInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func routesListOutputType() bindings.BindingType {
    return bindings.NewListType(bindings.NewReferenceType(RoutesRouteReadOnlyBindingType), reflect.TypeOf([]RoutesRouteReadOnly{}))
}

func routesListRestMetadata() protocol.OperationRestMetadata {
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


func routesDeleteInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["route"] = bindings.NewReferenceType(RoutesRouteBindingType)
    fieldNameMap["route"] = "Route"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func routesDeleteOutputType() bindings.BindingType {
    return bindings.NewVoidType()
}

func routesDeleteRestMetadata() protocol.OperationRestMetadata {
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



func RoutesRouteReadOnlyBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["destination"] = bindings.NewStringType()
    fieldNameMap["destination"] = "Destination"
    fields["prefix"] = bindings.NewIntegerType()
    fieldNameMap["prefix"] = "Prefix"
    fields["gateway"] = bindings.NewStringType()
    fieldNameMap["gateway"] = "Gateway"
    fields["interface_name"] = bindings.NewStringType()
    fieldNameMap["interface_name"] = "InterfaceName"
    fields["is_static"] = bindings.NewBooleanType()
    fieldNameMap["is_static"] = "IsStatic"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.appliance.techpreview.networking.routes.route_read_only",fields, reflect.TypeOf(RoutesRouteReadOnly{}), fieldNameMap, validators)
}

func RoutesRouteBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["destination"] = bindings.NewStringType()
    fieldNameMap["destination"] = "Destination"
    fields["prefix"] = bindings.NewIntegerType()
    fieldNameMap["prefix"] = "Prefix"
    fields["gateway"] = bindings.NewStringType()
    fieldNameMap["gateway"] = "Gateway"
    fields["interface_name"] = bindings.NewStringType()
    fieldNameMap["interface_name"] = "InterfaceName"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.appliance.techpreview.networking.routes.route",fields, reflect.TypeOf(RoutesRoute{}), fieldNameMap, validators)
}

func RoutesMessageBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["message"] = bindings.NewStringType()
    fieldNameMap["message"] = "Message"
    fields["result"] = bindings.NewEnumType("com.vmware.appliance.techpreview.networking.routes.message_status", reflect.TypeOf(RoutesMessageStatus(RoutesMessageStatus_failure)))
    fieldNameMap["result"] = "Result"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.appliance.techpreview.networking.routes.message",fields, reflect.TypeOf(RoutesMessage{}), fieldNameMap, validators)
}

func RoutesTestStatusInfoBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["status"] = bindings.NewEnumType("com.vmware.appliance.techpreview.networking.routes.test_status", reflect.TypeOf(RoutesTestStatus(RoutesTestStatus_orange)))
    fieldNameMap["status"] = "Status"
    fields["messages"] = bindings.NewListType(bindings.NewReferenceType(RoutesMessageBindingType), reflect.TypeOf([]RoutesMessage{}))
    fieldNameMap["messages"] = "Messages"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.appliance.techpreview.networking.routes.test_status_info",fields, reflect.TypeOf(RoutesTestStatusInfo{}), fieldNameMap, validators)
}

