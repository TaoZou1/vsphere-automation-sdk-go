/* Copyright © 2019 VMware, Inc. All Rights Reserved.
     SPDX-License-Identifier: BSD-2-Clause */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Data type definitions file for service: Networking.
 * Includes binding types of a structures and enumerations defined in the service.
 * Shared by client-side stubs and server-side skeletons to ensure type
 * compatibility.
 */

package appliance

import (
    "reflect"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/bindings/com/vmware/appliance/networking"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/bindings/com/vmware/appliance/networking/dns"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/bindings/com/vmware/appliance/networking/interfaces"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/bindings"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/data"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/protocol"
)




// The ``DNSInfo`` class contains information about the DNS configuration of a virtual appliance.
 type NetworkingDNSInfo struct {
    // DNS mode.
    Mode NetworkingDNSInfo_DNSMode
    // Hostname.
    Hostname string
    // Servers.
    Servers []string
}


// TODO some error doesn't have ErrorType. Resolve this later VAPI-3009
//
    
    // The ``DNSMode`` enumeration class describes the source of DNS servers.
    //
    // <p> See {@link com.vmware.vapi.bindings.ApiEnumeration enumerated types description}.
     
    type NetworkingDNSInfo_DNSMode string

    const (
        // The DNS servers addresses are obtained from a DHCP server.
         NetworkingDNSInfo_DNSMode_DHCP NetworkingDNSInfo_DNSMode = "DHCP"
        // The DNS servers addresses are specified explicitly.
         NetworkingDNSInfo_DNSMode_STATIC NetworkingDNSInfo_DNSMode = "STATIC"
    )

    func (d NetworkingDNSInfo_DNSMode) NetworkingDNSInfo_DNSMode() bool {
        switch d {
            case NetworkingDNSInfo_DNSMode_DHCP:
                return true
            case NetworkingDNSInfo_DNSMode_STATIC:
                return true
            default:
                return false
        }
    }



// The ``Info`` class contains information about the network configuration of a virtual appliance.
 type NetworkingInfo struct {
    // DNS configuration.
    Dns NetworkingDNSInfo
    // Interface configuration as a key-value map where key is a network interface name, for example, "nic0".
    Interfaces map[string]networking.InterfacesInterfaceInfo
}


// TODO some error doesn't have ErrorType. Resolve this later VAPI-3009
//


// The ``UpdateSpec`` class describes whether to enable or disable ipv6 on interfaces
 type NetworkingUpdateSpec struct {
    // IPv6 Enabled or not
    Ipv6Enabled *bool
}


// TODO some error doesn't have ErrorType. Resolve this later VAPI-3009
//


 type NetworkingChangeSpec struct {
    // New hostname to assign to the management network of vCenter appliance
    Hostname string
    // vCenter Server SSO administrator username
    SSOUser string
    // vCenter Server SSO administrator Password
    SSOPassword string
    // DNS Configuration to set for the machine
    Dns *dns.ServersDNSServerConfig
    // IPv4 Configuration to set for the machine
    Ipv4 *interfaces.Ipv4Config
    // IPv6 Configuration to set for the machine
    Ipv6 *interfaces.Ipv6Config
}


// TODO some error doesn't have ErrorType. Resolve this later VAPI-3009
//





func networkingGetInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func networkingGetOutputType() bindings.BindingType {
    return bindings.NewReferenceType(NetworkingInfoBindingType)
}

func networkingGetRestMetadata() protocol.OperationRestMetadata {
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


func networkingUpdateInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["spec"] = bindings.NewReferenceType(NetworkingUpdateSpecBindingType)
    fieldNameMap["spec"] = "Spec"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func networkingUpdateOutputType() bindings.BindingType {
    return bindings.NewVoidType()
}

func networkingUpdateRestMetadata() protocol.OperationRestMetadata {
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


func networkingResetInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func networkingResetOutputType() bindings.BindingType {
    return bindings.NewVoidType()
}

func networkingResetRestMetadata() protocol.OperationRestMetadata {
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


func networkingChangeInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["spec"] = bindings.NewReferenceType(NetworkingChangeSpecBindingType)
    fieldNameMap["spec"] = "Spec"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func networkingChangeOutputType() bindings.BindingType {
    return bindings.NewVoidType()
}

func networkingChangeRestMetadata() protocol.OperationRestMetadata {
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
       map[string]int{"Unsupported": 400,"InvalidArgument": 400,"Unauthenticated": 401,"NotAllowedInCurrentState": 400})
}



func NetworkingDNSInfoBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["mode"] = bindings.NewEnumType("com.vmware.appliance.networking.DNS_info.DNS_mode", reflect.TypeOf(NetworkingDNSInfo_DNSMode(NetworkingDNSInfo_DNSMode_DHCP)))
    fieldNameMap["mode"] = "Mode"
    fields["hostname"] = bindings.NewStringType()
    fieldNameMap["hostname"] = "Hostname"
    fields["servers"] = bindings.NewListType(bindings.NewStringType(), reflect.TypeOf([]string{}))
    fieldNameMap["servers"] = "Servers"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.appliance.networking.DNS_info",fields, reflect.TypeOf(NetworkingDNSInfo{}), fieldNameMap, validators)
}

func NetworkingInfoBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["dns"] = bindings.NewReferenceType(NetworkingDNSInfoBindingType)
    fieldNameMap["dns"] = "Dns"
    fields["interfaces"] = bindings.NewMapType(bindings.NewIdType([]string {"com.vmware.appliance.networking.interfaces"}, ""), bindings.NewReferenceType(networking.InterfacesInterfaceInfoBindingType),reflect.TypeOf(map[string]networking.InterfacesInterfaceInfo{}))
    fieldNameMap["interfaces"] = "Interfaces"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.appliance.networking.info",fields, reflect.TypeOf(NetworkingInfo{}), fieldNameMap, validators)
}

func NetworkingUpdateSpecBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["ipv6_enabled"] = bindings.NewOptionalType(bindings.NewBooleanType())
    fieldNameMap["ipv6_enabled"] = "Ipv6Enabled"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.appliance.networking.update_spec",fields, reflect.TypeOf(NetworkingUpdateSpec{}), fieldNameMap, validators)
}

func NetworkingChangeSpecBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["hostname"] = bindings.NewStringType()
    fieldNameMap["hostname"] = "Hostname"
    fields["SSO_user"] = bindings.NewStringType()
    fieldNameMap["SSO_user"] = "SSOUser"
    fields["SSO_password"] = bindings.NewSecretType()
    fieldNameMap["SSO_password"] = "SSOPassword"
    fields["dns"] = bindings.NewOptionalType(bindings.NewReferenceType(dns.ServersDNSServerConfigBindingType))
    fieldNameMap["dns"] = "Dns"
    fields["ipv4"] = bindings.NewOptionalType(bindings.NewReferenceType(interfaces.Ipv4ConfigBindingType))
    fieldNameMap["ipv4"] = "Ipv4"
    fields["ipv6"] = bindings.NewOptionalType(bindings.NewReferenceType(interfaces.Ipv6ConfigBindingType))
    fieldNameMap["ipv6"] = "Ipv6"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.appliance.networking.change_spec",fields, reflect.TypeOf(NetworkingChangeSpec{}), fieldNameMap, validators)
}


