/* Copyright © 2019 VMware, Inc. All Rights Reserved.
     SPDX-License-Identifier: BSD-2-Clause */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Data type definitions file for service: Ipv4.
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



// ``IPv4Mode`` enumeration class Defines different IPv4 modes.
//
// <p> See {@link com.vmware.vapi.bindings.ApiEnumeration enumerated types description}.
 
type Ipv4_IPv4Mode string

const (
    // IPv4 address is automatically assigned by a DHCP server.
     Ipv4_IPv4Mode_dhcp Ipv4_IPv4Mode = "dhcp"
    // IPv4 address is static.
     Ipv4_IPv4Mode_is_static Ipv4_IPv4Mode = "is_static"
    // The IPv4 protocol is not configured.
     Ipv4_IPv4Mode_unconfigured Ipv4_IPv4Mode = "unconfigured"
)

func (i Ipv4_IPv4Mode) Ipv4_IPv4Mode() bool {
    switch i {
        case Ipv4_IPv4Mode_dhcp:
            return true
        case Ipv4_IPv4Mode_is_static:
            return true
        case Ipv4_IPv4Mode_unconfigured:
            return true
        default:
            return false
    }
}





// ``IPv4Config`` class Structure that defines the IPv4 configuration state of a network interface.
 type Ipv4IPv4Config struct {
    // Interface name, for example, "nic0", "nic1".
    InterfaceName string
    // Address assignment mode.
    Mode Ipv4_IPv4Mode
    // IPv4 address, for example, "10.20.80.191". Set this argument to an empty string "", if the mode is "unconfigured" or "dhcp".
    Address string
    // IPv4 CIDR prefix, for example , 24. See http://www.oav.net/mirrors/cidr.html for netmask-to-prefix conversion. Set this argument to 0 if the mode is "unconfigured" or "dhcp".
    Prefix int64
    // IPv4 address of the default gateway. This default gateway value is used if the mode argument is set to "static" This configures the global default gateway on the appliance with the specified gateway address and interface. This gateway replaces the existing default gateway configured on the appliance. However, if the gateway address is link-local, then it is added for that interface. This does not support configuration of multiple global default gateways through different interfaces.
    DefaultGateway string
}


// TODO some error doesn't have ErrorType. Resolve this later VAPI-3009
//


// ``IPv4ConfigReadOnly`` class Structure that defines the IPv4 configuration state of a network interface.
 type Ipv4IPv4ConfigReadOnly struct {
    // Interface name, for example, "nic0", "nic1".
    InterfaceName string
    // Address assignment mode.
    Mode Ipv4_IPv4Mode
    // IPv4 address, for example, "10.20.80.191". Set this argument to an empty string "", if the mode is "unconfigured" or "dhcp".
    Address string
    // IPv4 CIDR prefix, for example , 24. See http://www.oav.net/mirrors/cidr.html for netmask-to-prefix conversion. Set this argument to 0 if the mode is "unconfigured" or "dhcp".
    Prefix int64
    // IPv4 address of the default gateway. This default gateway value is used if the mode argument is set to "static" This configures the global default gateway on the appliance with the specified gateway address and interface. This gateway replaces the existing default gateway configured on the appliance. However, if the gateway address is link-local, then it is added for that interface. This does not support configuration of multiple global default gateways through different interfaces.
    DefaultGateway string
    // This indicates if the network configuration can be updated for the interface.
    Updateable bool
}


// TODO some error doesn't have ErrorType. Resolve this later VAPI-3009
//





func ipv4RenewInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["interfaces"] = bindings.NewListType(bindings.NewStringType(), reflect.TypeOf([]string{}))
    fieldNameMap["interfaces"] = "Interfaces"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func ipv4RenewOutputType() bindings.BindingType {
    return bindings.NewVoidType()
}

func ipv4RenewRestMetadata() protocol.OperationRestMetadata {
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


func ipv4SetInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["config"] = bindings.NewListType(bindings.NewReferenceType(Ipv4IPv4ConfigBindingType), reflect.TypeOf([]Ipv4IPv4Config{}))
    fieldNameMap["config"] = "Config"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func ipv4SetOutputType() bindings.BindingType {
    return bindings.NewVoidType()
}

func ipv4SetRestMetadata() protocol.OperationRestMetadata {
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


func ipv4ListInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func ipv4ListOutputType() bindings.BindingType {
    return bindings.NewListType(bindings.NewReferenceType(Ipv4IPv4ConfigReadOnlyBindingType), reflect.TypeOf([]Ipv4IPv4ConfigReadOnly{}))
}

func ipv4ListRestMetadata() protocol.OperationRestMetadata {
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


func ipv4GetInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["interfaces"] = bindings.NewListType(bindings.NewStringType(), reflect.TypeOf([]string{}))
    fieldNameMap["interfaces"] = "Interfaces"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func ipv4GetOutputType() bindings.BindingType {
    return bindings.NewListType(bindings.NewReferenceType(Ipv4IPv4ConfigReadOnlyBindingType), reflect.TypeOf([]Ipv4IPv4ConfigReadOnly{}))
}

func ipv4GetRestMetadata() protocol.OperationRestMetadata {
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



func Ipv4IPv4ConfigBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["interface_name"] = bindings.NewStringType()
    fieldNameMap["interface_name"] = "InterfaceName"
    fields["mode"] = bindings.NewEnumType("com.vmware.appliance.techpreview.networking.ipv4.I_pv4_mode", reflect.TypeOf(Ipv4_IPv4Mode(Ipv4_IPv4Mode_dhcp)))
    fieldNameMap["mode"] = "Mode"
    fields["address"] = bindings.NewStringType()
    fieldNameMap["address"] = "Address"
    fields["prefix"] = bindings.NewIntegerType()
    fieldNameMap["prefix"] = "Prefix"
    fields["default_gateway"] = bindings.NewStringType()
    fieldNameMap["default_gateway"] = "DefaultGateway"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.appliance.techpreview.networking.ipv4.I_pv4_config",fields, reflect.TypeOf(Ipv4IPv4Config{}), fieldNameMap, validators)
}

func Ipv4IPv4ConfigReadOnlyBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["interface_name"] = bindings.NewStringType()
    fieldNameMap["interface_name"] = "InterfaceName"
    fields["mode"] = bindings.NewEnumType("com.vmware.appliance.techpreview.networking.ipv4.I_pv4_mode", reflect.TypeOf(Ipv4_IPv4Mode(Ipv4_IPv4Mode_dhcp)))
    fieldNameMap["mode"] = "Mode"
    fields["address"] = bindings.NewStringType()
    fieldNameMap["address"] = "Address"
    fields["prefix"] = bindings.NewIntegerType()
    fieldNameMap["prefix"] = "Prefix"
    fields["default_gateway"] = bindings.NewStringType()
    fieldNameMap["default_gateway"] = "DefaultGateway"
    fields["updateable"] = bindings.NewBooleanType()
    fieldNameMap["updateable"] = "Updateable"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.appliance.techpreview.networking.ipv4.I_pv4_config_read_only",fields, reflect.TypeOf(Ipv4IPv4ConfigReadOnly{}), fieldNameMap, validators)
}


