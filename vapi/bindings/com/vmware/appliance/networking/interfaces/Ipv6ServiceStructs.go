/* Copyright © 2019 VMware, Inc. All Rights Reserved.
     SPDX-License-Identifier: BSD-2-Clause */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Data type definitions file for service: Ipv6.
 * Includes binding types of a structures and enumerations defined in the service.
 * Shared by client-side stubs and server-side skeletons to ensure type
 * compatibility.
 */

package interfaces

import (
    "reflect"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/bindings"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/data"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/protocol"
)



// The ``Origin`` enumeration class defines IPv6 address origin values.
//
// <p> See {@link com.vmware.vapi.bindings.ApiEnumeration enumerated types description}.
 
type Ipv6_Origin string

const (
    // The IPv6 address is assigned by a DHCP server. See RFC 4293.
     Ipv6_Origin_DHCP Ipv6_Origin = "DHCP"
    // The IPv6 address is assigned randomly by the system. See RFC 4293.
     Ipv6_Origin_RANDOM Ipv6_Origin = "RANDOM"
    // The IPv6 address was manually configured to a specified address, for example, by user configuration. See RFC 4293.
     Ipv6_Origin_MANUAL Ipv6_Origin = "MANUAL"
    // The IPv6 address is assigned by IPv6 Stateless Address Auto-configuration (SLAAC). See RFC 4293.
     Ipv6_Origin_LINKLAYER Ipv6_Origin = "LINKLAYER"
    // The IPv6 address is assigned by a mechanism other than manual, DHCP, SLAAC, or random. See RFC 4293.
     Ipv6_Origin_OTHER Ipv6_Origin = "OTHER"
)

func (o Ipv6_Origin) Ipv6_Origin() bool {
    switch o {
        case Ipv6_Origin_DHCP:
            return true
        case Ipv6_Origin_RANDOM:
            return true
        case Ipv6_Origin_MANUAL:
            return true
        case Ipv6_Origin_LINKLAYER:
            return true
        case Ipv6_Origin_OTHER:
            return true
        default:
            return false
    }
}




// The ``Status`` enumeration class defines IPv6 address status values. See RFC 4293.
//
// <p> See {@link com.vmware.vapi.bindings.ApiEnumeration enumerated types description}.
 
type Ipv6_Status string

const (
    // The IPv6 address is in the process of being verified as unique. An address in this state cannot be used for general communication. It can be used to determine the uniqueness of the address.
     Ipv6_Status_TENTATIVE Ipv6_Status = "TENTATIVE"
    // The status of this address cannot be determined.
     Ipv6_Status_UNKNOWN Ipv6_Status = "UNKNOWN"
    // The IPv6 address is inaccessible because the interface to which this address is assigned is not operational.
     Ipv6_Status_INACCESSIBLE Ipv6_Status = "INACCESSIBLE"
    // The IPv6 address is not a valid address. It should not appear as the destination or source address of a packet.
     Ipv6_Status_INVALID Ipv6_Status = "INVALID"
    // The IPv6 address is not unique on the link and cannot be used.
     Ipv6_Status_DUPLICATE Ipv6_Status = "DUPLICATE"
    // This is a valid IPv6 address that can appear as the destination or source address of a packet.
     Ipv6_Status_PREFERRED Ipv6_Status = "PREFERRED"
    // The is a valid but deprecated IPv6 address. This address cannot be used as a source address in new communications, although packets addressed to such an address are processed as expected.
     Ipv6_Status_DEPRECATED Ipv6_Status = "DEPRECATED"
    // The IPv6 address is available for use, subject to restrictions, while its uniqueness on a link is being verified.
     Ipv6_Status_OPTIMISTIC Ipv6_Status = "OPTIMISTIC"
)

func (s Ipv6_Status) Ipv6_Status() bool {
    switch s {
        case Ipv6_Status_TENTATIVE:
            return true
        case Ipv6_Status_UNKNOWN:
            return true
        case Ipv6_Status_INACCESSIBLE:
            return true
        case Ipv6_Status_INVALID:
            return true
        case Ipv6_Status_DUPLICATE:
            return true
        case Ipv6_Status_PREFERRED:
            return true
        case Ipv6_Status_DEPRECATED:
            return true
        case Ipv6_Status_OPTIMISTIC:
            return true
        default:
            return false
    }
}





// The ``Address`` class provides the structure used to name an IPv6 address.
 type Ipv6Address struct {
    // The IPv6 address, for example, fc00:10:20:83:20c:29ff:fe94:bb5a.
    Address string
    // The IPv6 CIDR prefix, for example, 64.
    Prefix int64
}



func (Ipv6Address Ipv6Address) Error() string {
    return "com.vmware.appliance.networking.interfaces.address"
}



// The ``AddressInfo`` class provides the structure that you can use to get information about an IPv6 address along with its origin and status.
 type Ipv6AddressInfo struct {
    // The Origin of the IPv6 address. For more information, see RFC 4293.
    Origin Ipv6_Origin
    // The Status of the IPv6 address. For more information, see RFC 4293.
    Status Ipv6_Status
    // The IPv6 address, for example, fc00:10:20:83:20c:29ff:fe94:bb5a.
    Address string
    // The IPv6 CIDR prefix, for example, 64.
    Prefix int64
}



func (Ipv6AddressInfo Ipv6AddressInfo) Error() string {
    return "com.vmware.appliance.networking.interfaces.address_info"
}



// The ``Config`` class provides the structure that you can use to configure IPv6 on a particular interface. Addresses can be assigned by DHCP, SLAAC or STATIC, as IPv6 permits multiple addresses per interface.
 type Ipv6Config struct {
    // An address will be assigned by a DHCP server.
    Dhcp bool
    // An address will be assigned by Stateless Address Autoconfiguration (SLAAC).
    Autoconf bool
    // The list of addresses to be statically assigned.
    Addresses []Ipv6Address
    // The default gateway for static IP address assignment. This configures the global IPv6 default gateway on the appliance with the specified gateway address and interface. This gateway replaces the existing default gateway configured on the appliance. However, if the gateway address is link-local, then it is added for that interface. This does not support configuration of multiple global default gateways through different interfaces.
    DefaultGateway string
}



func (Ipv6Config Ipv6Config) Error() string {
    return "com.vmware.appliance.networking.interfaces.config"
}



// The ``Info`` class provides the structure that defines an existing IPv6 configuration on a particular interface. This structure is read only.
 type Ipv6Info struct {
    // DHCP is on.
    Dhcp bool
    // Stateless Address Autoconfiguration (SLAAC) is on.
    Autoconf bool
    // List of addresses with their origins and statuses.
    Addresses []Ipv6AddressInfo
    // The default gateway for static IP address assignment. This configures the global IPv6 default gateway on the appliance with the specified gateway address and interface. This gateway replaces the existing default gateway configured on the appliance. However, if the gateway address is link-local, then it is added for that interface. This does not support configuration of multiple global default gateways through different interfaces.
    DefaultGateway string
    // Is NIC configurable or not.
    Configurable bool
}



func (Ipv6Info Ipv6Info) Error() string {
    return "com.vmware.appliance.networking.interfaces.info"
}






func ipv6SetInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["interface_name"] = bindings.NewIdType([]string {"com.vmware.appliance.networking.interfaces"}, "")
    fields["config"] = bindings.NewReferenceType(Ipv6ConfigBindingType)
    fieldNameMap["interface_name"] = "InterfaceName"
    fieldNameMap["config"] = "Config"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func ipv6SetOutputType() bindings.BindingType {
    return bindings.NewVoidType()
}

func ipv6SetRestMetadata() protocol.OperationRestMetadata {
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
       map[string]int{"ResourceBusy": 500,"NotFound": 404,"Error": 500})
}


func ipv6GetInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["interface_name"] = bindings.NewIdType([]string {"com.vmware.appliance.networking.interfaces"}, "")
    fieldNameMap["interface_name"] = "InterfaceName"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func ipv6GetOutputType() bindings.BindingType {
    return bindings.NewReferenceType(Ipv6InfoBindingType)
}

func ipv6GetRestMetadata() protocol.OperationRestMetadata {
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



func Ipv6AddressBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["address"] = bindings.NewStringType()
    fieldNameMap["address"] = "Address"
    fields["prefix"] = bindings.NewIntegerType()
    fieldNameMap["prefix"] = "Prefix"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.appliance.networking.interfaces.ipv6.address",fields, reflect.TypeOf(Ipv6Address{}), fieldNameMap, validators)
}

func Ipv6AddressInfoBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["origin"] = bindings.NewEnumType("com.vmware.appliance.networking.interfaces.ipv6.origin", reflect.TypeOf(Ipv6_Origin(Ipv6_Origin_DHCP)))
    fieldNameMap["origin"] = "Origin"
    fields["status"] = bindings.NewEnumType("com.vmware.appliance.networking.interfaces.ipv6.status", reflect.TypeOf(Ipv6_Status(Ipv6_Status_TENTATIVE)))
    fieldNameMap["status"] = "Status"
    fields["address"] = bindings.NewStringType()
    fieldNameMap["address"] = "Address"
    fields["prefix"] = bindings.NewIntegerType()
    fieldNameMap["prefix"] = "Prefix"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.appliance.networking.interfaces.ipv6.address_info",fields, reflect.TypeOf(Ipv6AddressInfo{}), fieldNameMap, validators)
}

func Ipv6ConfigBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["dhcp"] = bindings.NewBooleanType()
    fieldNameMap["dhcp"] = "Dhcp"
    fields["autoconf"] = bindings.NewBooleanType()
    fieldNameMap["autoconf"] = "Autoconf"
    fields["addresses"] = bindings.NewListType(bindings.NewReferenceType(Ipv6AddressBindingType), reflect.TypeOf([]Ipv6Address{}))
    fieldNameMap["addresses"] = "Addresses"
    fields["default_gateway"] = bindings.NewStringType()
    fieldNameMap["default_gateway"] = "DefaultGateway"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.appliance.networking.interfaces.ipv6.config",fields, reflect.TypeOf(Ipv6Config{}), fieldNameMap, validators)
}

func Ipv6InfoBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["dhcp"] = bindings.NewBooleanType()
    fieldNameMap["dhcp"] = "Dhcp"
    fields["autoconf"] = bindings.NewBooleanType()
    fieldNameMap["autoconf"] = "Autoconf"
    fields["addresses"] = bindings.NewListType(bindings.NewReferenceType(Ipv6AddressInfoBindingType), reflect.TypeOf([]Ipv6AddressInfo{}))
    fieldNameMap["addresses"] = "Addresses"
    fields["default_gateway"] = bindings.NewStringType()
    fieldNameMap["default_gateway"] = "DefaultGateway"
    fields["configurable"] = bindings.NewBooleanType()
    fieldNameMap["configurable"] = "Configurable"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.appliance.networking.interfaces.ipv6.info",fields, reflect.TypeOf(Ipv6Info{}), fieldNameMap, validators)
}


