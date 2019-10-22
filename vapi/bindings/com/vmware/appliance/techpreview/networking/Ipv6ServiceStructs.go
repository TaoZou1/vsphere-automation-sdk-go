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

package networking

import (
    "reflect"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/bindings"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/data"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/protocol"
)



// ``IPv6AddressOrigin`` enumeration class Defines IPv6 address origin values.
//
// <p> See {@link com.vmware.vapi.bindings.ApiEnumeration enumerated types description}.
 
type Ipv6_IPv6AddressOrigin string

const (
    // The IPv6 address is assigned by a DHCP server. See RFC 4293.
     Ipv6_IPv6AddressOrigin_dhcp Ipv6_IPv6AddressOrigin = "dhcp"
    // The IPv6 address is assigned randomly by the system. See RFC 4293.
     Ipv6_IPv6AddressOrigin_random Ipv6_IPv6AddressOrigin = "random"
    // The IPv6 address was manually configured to a specified address, for, example, by user configuration. See RFC 4293.
     Ipv6_IPv6AddressOrigin_manual Ipv6_IPv6AddressOrigin = "manual"
    // The IPv6 address is assigned by a mechanism other than manual, DHCP, SLAAC, or random. See RFC 4293.
     Ipv6_IPv6AddressOrigin_other Ipv6_IPv6AddressOrigin = "other"
    // The IPv6 address is assigned by IPv6 Stateless Address Auto-configuration (SLAAC). See RFC 4293.
     Ipv6_IPv6AddressOrigin_linklayer Ipv6_IPv6AddressOrigin = "linklayer"
)

func (i Ipv6_IPv6AddressOrigin) Ipv6_IPv6AddressOrigin() bool {
    switch i {
        case Ipv6_IPv6AddressOrigin_dhcp:
            return true
        case Ipv6_IPv6AddressOrigin_random:
            return true
        case Ipv6_IPv6AddressOrigin_manual:
            return true
        case Ipv6_IPv6AddressOrigin_other:
            return true
        case Ipv6_IPv6AddressOrigin_linklayer:
            return true
        default:
            return false
    }
}




// ``IPv6AddressStatus`` enumeration class Defines IPv6 address status values.
//
// <p> See {@link com.vmware.vapi.bindings.ApiEnumeration enumerated types description}.
 
type Ipv6_IPv6AddressStatus string

const (
    // This IPv6 address is in the process of being verified as unique. Do not use addresses in this state for general communication. You can use them to determine the uniqueness of the address. See RFC 4293.
     Ipv6_IPv6AddressStatus_tentative Ipv6_IPv6AddressStatus = "tentative"
    // The status of this address cannot be determined. See RFC 4293.
     Ipv6_IPv6AddressStatus_unknown Ipv6_IPv6AddressStatus = "unknown"
    // This IPv6 address is inaccessible because the interface to which this address is assigned is not operational. See RFC 4293.
     Ipv6_IPv6AddressStatus_inaccessible Ipv6_IPv6AddressStatus = "inaccessible"
    // This IPv6 address is not a valid address. It should not appear as the destination or source address of a packet. See RFC 4293.
     Ipv6_IPv6AddressStatus_invalid Ipv6_IPv6AddressStatus = "invalid"
    // This IPv6 address is not unique on the link. Do use this IPv6 address. See RFC 4293.
     Ipv6_IPv6AddressStatus_duplicate Ipv6_IPv6AddressStatus = "duplicate"
    // This is a valid IPv6 address that can appear as the destination or source address of a packet. See RFC 4293.
     Ipv6_IPv6AddressStatus_preferred Ipv6_IPv6AddressStatus = "preferred"
    // This is a valid but deprecated IPv6 address. Do not use this IPv6 address as a source address in new communications, although packets addressed to such an address are processed as expected. See RFC 4293.
     Ipv6_IPv6AddressStatus_deprecated Ipv6_IPv6AddressStatus = "deprecated"
    // This IPv6 address is available for use, subject to restrictions, while its uniqueness on a link is being verified. See RFC 4293.
     Ipv6_IPv6AddressStatus_optimistic Ipv6_IPv6AddressStatus = "optimistic"
)

func (i Ipv6_IPv6AddressStatus) Ipv6_IPv6AddressStatus() bool {
    switch i {
        case Ipv6_IPv6AddressStatus_tentative:
            return true
        case Ipv6_IPv6AddressStatus_unknown:
            return true
        case Ipv6_IPv6AddressStatus_inaccessible:
            return true
        case Ipv6_IPv6AddressStatus_invalid:
            return true
        case Ipv6_IPv6AddressStatus_duplicate:
            return true
        case Ipv6_IPv6AddressStatus_preferred:
            return true
        case Ipv6_IPv6AddressStatus_deprecated:
            return true
        case Ipv6_IPv6AddressStatus_optimistic:
            return true
        default:
            return false
    }
}





// ``IPv6AddressReadOnly`` class Structure that you can use to get information about an IPv6 address along with its origin and status.
 type Ipv6IPv6AddressReadOnly struct {
    // IPv6 address, for example, fc00:10:20:83:20c:29ff:fe94:bb5a.
    Address string
    // IPv6 CIDR prefix, for example, 64.
    Prefix int64
    // Origin of the IPv6 address. For more information, see RFC 4293.
    Origin Ipv6_IPv6AddressOrigin
    // Status of the IPv6 address. For more information, see RFC 4293.
    Status Ipv6_IPv6AddressStatus
}



//


// ``IPv6ConfigReadOnly`` class Structure that defines an existing IPv6 configuration on a particular interface. This structure is read only.
 type Ipv6IPv6ConfigReadOnly struct {
    // Network interface, for example, "nic0" queried.
    InterfaceName string
    // Address assigned by a DHCP server.
    Dhcp bool
    // Address is assigned by Stateless Address Autoconfiguration (SLAAC).
    Autoconf bool
    // A list of all addresses assigned to this interface. The origin field of each address determines how the address was assigned, for example, statically, by DHCP, SLAAC.
    Addresses []Ipv6IPv6AddressReadOnly
    // Default gateway. This configures the global IPv6 default gateway on the appliance with the specified gateway address and interface. This gateway replaces the existing default gateway configured on the appliance. However, if the gateway address is link-local, then it is added for that interface. This does not support configuration of multiple global default gateways through different interfaces.
    DefaultGateway string
    // This indicates if the network configuration can be updated for the interface.
    Updateable bool
}



//


// ``IPv6Address`` class Structure used to name an IPv6 address.
 type Ipv6IPv6Address struct {
    // IPv6 address, for example, fc00:10:20:83:20c:29ff:fe94:bb5a.
    Address string
    // IPv6 CIDR prefix, for example, 64.
    Prefix int64
}



//


// ``IPv6Config`` class Structure that you can use to configure IPv6 on a particular interface. Because IPv6 permits multiple addresses per interface, addresses can be assigned by DHCP, SLAAC, and can also be statically assigned.
 type Ipv6IPv6Config struct {
    // Network interface, for example, "nic0" to configure.
    InterfaceName string
    // Address assigned by a DHCP server.
    Dhcp bool
    // Address is assigned by Stateless Address Autoconfiguration (SLAAC).
    Autoconf bool
    // A list of addresses to be statically assigned.
    Addresses []Ipv6IPv6Address
    // Default gateway for static IP address assignment. This configures the global IPv6 default gateway on the appliance with the specified gateway address and interface. This gateway replaces the existing default gateway configured on the appliance. However, if the gateway address is link-local, then it is added for that interface. This does not support configuration of multiple global default gateways through different interfaces.
    DefaultGateway string
}



//





func ipv6SetInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["config"] = bindings.NewListType(bindings.NewReferenceType(Ipv6IPv6ConfigBindingType), reflect.TypeOf([]Ipv6IPv6Config{}))
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
       map[string]int{"Error": 500})
}


func ipv6ListInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func ipv6ListOutputType() bindings.BindingType {
    return bindings.NewListType(bindings.NewReferenceType(Ipv6IPv6ConfigReadOnlyBindingType), reflect.TypeOf([]Ipv6IPv6ConfigReadOnly{}))
}

func ipv6ListRestMetadata() protocol.OperationRestMetadata {
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


func ipv6GetInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["interfaces"] = bindings.NewListType(bindings.NewStringType(), reflect.TypeOf([]string{}))
    fieldNameMap["interfaces"] = "Interfaces"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func ipv6GetOutputType() bindings.BindingType {
    return bindings.NewListType(bindings.NewReferenceType(Ipv6IPv6ConfigReadOnlyBindingType), reflect.TypeOf([]Ipv6IPv6ConfigReadOnly{}))
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
       map[string]int{"Error": 500})
}



func Ipv6IPv6AddressReadOnlyBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["address"] = bindings.NewStringType()
    fieldNameMap["address"] = "Address"
    fields["prefix"] = bindings.NewIntegerType()
    fieldNameMap["prefix"] = "Prefix"
    fields["origin"] = bindings.NewEnumType("com.vmware.appliance.techpreview.networking.ipv6.I_pv6_address_origin", reflect.TypeOf(Ipv6_IPv6AddressOrigin(Ipv6_IPv6AddressOrigin_dhcp)))
    fieldNameMap["origin"] = "Origin"
    fields["status"] = bindings.NewEnumType("com.vmware.appliance.techpreview.networking.ipv6.I_pv6_address_status", reflect.TypeOf(Ipv6_IPv6AddressStatus(Ipv6_IPv6AddressStatus_tentative)))
    fieldNameMap["status"] = "Status"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.appliance.techpreview.networking.ipv6.I_pv6_address_read_only",fields, reflect.TypeOf(Ipv6IPv6AddressReadOnly{}), fieldNameMap, validators)
}

func Ipv6IPv6ConfigReadOnlyBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["interface_name"] = bindings.NewStringType()
    fieldNameMap["interface_name"] = "InterfaceName"
    fields["dhcp"] = bindings.NewBooleanType()
    fieldNameMap["dhcp"] = "Dhcp"
    fields["autoconf"] = bindings.NewBooleanType()
    fieldNameMap["autoconf"] = "Autoconf"
    fields["addresses"] = bindings.NewListType(bindings.NewReferenceType(Ipv6IPv6AddressReadOnlyBindingType), reflect.TypeOf([]Ipv6IPv6AddressReadOnly{}))
    fieldNameMap["addresses"] = "Addresses"
    fields["default_gateway"] = bindings.NewStringType()
    fieldNameMap["default_gateway"] = "DefaultGateway"
    fields["updateable"] = bindings.NewBooleanType()
    fieldNameMap["updateable"] = "Updateable"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.appliance.techpreview.networking.ipv6.I_pv6_config_read_only",fields, reflect.TypeOf(Ipv6IPv6ConfigReadOnly{}), fieldNameMap, validators)
}

func Ipv6IPv6AddressBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["address"] = bindings.NewStringType()
    fieldNameMap["address"] = "Address"
    fields["prefix"] = bindings.NewIntegerType()
    fieldNameMap["prefix"] = "Prefix"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.appliance.techpreview.networking.ipv6.I_pv6_address",fields, reflect.TypeOf(Ipv6IPv6Address{}), fieldNameMap, validators)
}

func Ipv6IPv6ConfigBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["interface_name"] = bindings.NewStringType()
    fieldNameMap["interface_name"] = "InterfaceName"
    fields["dhcp"] = bindings.NewBooleanType()
    fieldNameMap["dhcp"] = "Dhcp"
    fields["autoconf"] = bindings.NewBooleanType()
    fieldNameMap["autoconf"] = "Autoconf"
    fields["addresses"] = bindings.NewListType(bindings.NewReferenceType(Ipv6IPv6AddressBindingType), reflect.TypeOf([]Ipv6IPv6Address{}))
    fieldNameMap["addresses"] = "Addresses"
    fields["default_gateway"] = bindings.NewStringType()
    fieldNameMap["default_gateway"] = "DefaultGateway"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.appliance.techpreview.networking.ipv6.I_pv6_config",fields, reflect.TypeOf(Ipv6IPv6Config{}), fieldNameMap, validators)
}


