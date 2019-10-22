/* Copyright © 2019 VMware, Inc. All Rights Reserved.
     SPDX-License-Identifier: BSD-2-Clause */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Data type definitions file for service: Interfaces.
 * Includes binding types of a structures and enumerations defined in the service.
 * Shared by client-side stubs and server-side skeletons to ensure type
 * compatibility.
 */

package networking

import (
    "reflect"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/bindings/com/vmware/appliance/networking/interfaces"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/bindings"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/data"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/protocol"
)



// ``InterfaceStatus`` enumeration class Defines interface status
//
// <p> See {@link com.vmware.vapi.bindings.ApiEnumeration enumerated types description}.
 
type Interfaces_InterfaceStatus string

const (
    // The interface is down.
     Interfaces_InterfaceStatus_down Interfaces_InterfaceStatus = "down"
    // The interface is up.
     Interfaces_InterfaceStatus_up Interfaces_InterfaceStatus = "up"
)

func (i Interfaces_InterfaceStatus) Interfaces_InterfaceStatus() bool {
    switch i {
        case Interfaces_InterfaceStatus_down:
            return true
        case Interfaces_InterfaceStatus_up:
            return true
        default:
            return false
    }
}





// ``InterfaceInfo`` class Structure that defines properties and status of a network interface.
 type InterfacesInterfaceInfo struct {
    // Interface name, for example, "nic0", "nic1".
    Name string
    // Interface status.
    Status Interfaces_InterfaceStatus
    // MAC address. For example 00:0C:29:94:BB:5A.
    Mac string
    // IPv4 Address information.
    Ipv4 *interfaces.Ipv4Info
    // IPv6 Address information.
    Ipv6 *interfaces.Ipv6Info
}


// TODO some error doesn't have ErrorType. Resolve this later VAPI-3009
//





func interfacesListInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func interfacesListOutputType() bindings.BindingType {
    return bindings.NewListType(bindings.NewReferenceType(InterfacesInterfaceInfoBindingType), reflect.TypeOf([]InterfacesInterfaceInfo{}))
}

func interfacesListRestMetadata() protocol.OperationRestMetadata {
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


func interfacesGetInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["interface_name"] = bindings.NewIdType([]string {"com.vmware.appliance.networking.interfaces"}, "")
    fieldNameMap["interface_name"] = "InterfaceName"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func interfacesGetOutputType() bindings.BindingType {
    return bindings.NewReferenceType(InterfacesInterfaceInfoBindingType)
}

func interfacesGetRestMetadata() protocol.OperationRestMetadata {
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
       map[string]int{"NotFound": 404,"Error": 500})
}



func InterfacesInterfaceInfoBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["name"] = bindings.NewStringType()
    fieldNameMap["name"] = "Name"
    fields["status"] = bindings.NewEnumType("com.vmware.appliance.networking.interfaces.interface_status", reflect.TypeOf(Interfaces_InterfaceStatus(Interfaces_InterfaceStatus_down)))
    fieldNameMap["status"] = "Status"
    fields["mac"] = bindings.NewStringType()
    fieldNameMap["mac"] = "Mac"
    fields["ipv4"] = bindings.NewOptionalType(bindings.NewReferenceType(interfaces.Ipv4InfoBindingType))
    fieldNameMap["ipv4"] = "Ipv4"
    fields["ipv6"] = bindings.NewOptionalType(bindings.NewReferenceType(interfaces.Ipv6InfoBindingType))
    fieldNameMap["ipv6"] = "Ipv6"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.appliance.networking.interfaces.interface_info",fields, reflect.TypeOf(InterfacesInterfaceInfo{}), fieldNameMap, validators)
}


