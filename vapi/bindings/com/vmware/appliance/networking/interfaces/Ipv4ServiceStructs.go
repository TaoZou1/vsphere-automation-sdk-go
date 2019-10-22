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

package interfaces

import (
    "reflect"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/bindings"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/data"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/protocol"
)



// The ``Mode`` enumeration class defines different IPv4 address assignment modes.
//
// <p> See {@link com.vmware.vapi.bindings.ApiEnumeration enumerated types description}.
 
type Ipv4_Mode string

const (
    // The IPv4 address is automatically assigned by a DHCP server.
     Ipv4_Mode_DHCP Ipv4_Mode = "DHCP"
    // The IPv4 address is static.
     Ipv4_Mode_STATIC Ipv4_Mode = "STATIC"
    // The IPv4 protocol is not configured.
     Ipv4_Mode_UNCONFIGURED Ipv4_Mode = "UNCONFIGURED"
)

func (m Ipv4_Mode) Ipv4_Mode() bool {
    switch m {
        case Ipv4_Mode_DHCP:
            return true
        case Ipv4_Mode_STATIC:
            return true
        case Ipv4_Mode_UNCONFIGURED:
            return true
        default:
            return false
    }
}





// The ``Config`` class provides defines the IPv4 configuration of a network interface.
 type Ipv4Config struct {
    // The Address assignment mode.
    Mode Ipv4_Mode
    // The IPv4 address, for example, "10.20.80.191".
    Address *string
    // The IPv4 CIDR prefix, for example, 24. See http://www.oav.net/mirrors/cidr.html for netmask-to-prefix conversion.
    Prefix *int64
    // The IPv4 address of the default gateway. This configures the global default gateway on the appliance with the specified gateway address and interface. This gateway replaces the existing default gateway configured on the appliance. However, if the gateway address is link-local, then it is added for that interface. This does not support configuration of multiple global default gateways through different interfaces.
    DefaultGateway *string
}


// TODO some error doesn't have ErrorType. Resolve this later VAPI-3009
//


// The ``Info`` class defines current IPv4 configuration state of a network interface.
 type Ipv4Info struct {
    // The specified network interface is configurable or not.
    Configurable bool
    // The Address assignment mode.
    Mode Ipv4_Mode
    // The IPv4 address, for example, "10.20.80.191".
    Address *string
    // The IPv4 CIDR prefix, for example, 24. See http://www.oav.net/mirrors/cidr.html for netmask-to-prefix conversion.
    Prefix *int64
    // The IPv4 address of the default gateway. This configures the global default gateway on the appliance with the specified gateway address and interface. This gateway replaces the existing default gateway configured on the appliance. However, if the gateway address is link-local, then it is added for that interface. This does not support configuration of multiple global default gateways through different interfaces.
    DefaultGateway *string
}


// TODO some error doesn't have ErrorType. Resolve this later VAPI-3009
//





func ipv4SetInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["interface_name"] = bindings.NewIdType([]string {"com.vmware.appliance.networking.interfaces"}, "")
    fields["config"] = bindings.NewReferenceType(Ipv4ConfigBindingType)
    fieldNameMap["interface_name"] = "InterfaceName"
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
       map[string]int{"NotFound": 404,"NotAllowedInCurrentState": 400,"ResourceBusy": 500,"Error": 500})
}


func ipv4GetInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["interface_name"] = bindings.NewIdType([]string {"com.vmware.appliance.networking.interfaces"}, "")
    fieldNameMap["interface_name"] = "InterfaceName"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func ipv4GetOutputType() bindings.BindingType {
    return bindings.NewReferenceType(Ipv4InfoBindingType)
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
       map[string]int{"NotFound": 404,"Error": 500})
}



func Ipv4ConfigBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["mode"] = bindings.NewEnumType("com.vmware.appliance.networking.interfaces.ipv4.mode", reflect.TypeOf(Ipv4_Mode(Ipv4_Mode_DHCP)))
    fieldNameMap["mode"] = "Mode"
    fields["address"] = bindings.NewOptionalType(bindings.NewStringType())
    fieldNameMap["address"] = "Address"
    fields["prefix"] = bindings.NewOptionalType(bindings.NewIntegerType())
    fieldNameMap["prefix"] = "Prefix"
    fields["default_gateway"] = bindings.NewOptionalType(bindings.NewStringType())
    fieldNameMap["default_gateway"] = "DefaultGateway"
    var validators = []bindings.Validator{}
    uv1 := bindings.NewUnionValidator("mode",
        map[string][]bindings.FieldData {
            "STATIC": []bindings.FieldData {
                 bindings.NewFieldData("address", true),
                 bindings.NewFieldData("prefix", true),
            },
            "DHCP": []bindings.FieldData {},
            "UNCONFIGURED": []bindings.FieldData {},
        },
    )
    validators = append(validators, uv1)
    return bindings.NewStructType("com.vmware.appliance.networking.interfaces.ipv4.config",fields, reflect.TypeOf(Ipv4Config{}), fieldNameMap, validators)
}

func Ipv4InfoBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["configurable"] = bindings.NewBooleanType()
    fieldNameMap["configurable"] = "Configurable"
    fields["mode"] = bindings.NewEnumType("com.vmware.appliance.networking.interfaces.ipv4.mode", reflect.TypeOf(Ipv4_Mode(Ipv4_Mode_DHCP)))
    fieldNameMap["mode"] = "Mode"
    fields["address"] = bindings.NewOptionalType(bindings.NewStringType())
    fieldNameMap["address"] = "Address"
    fields["prefix"] = bindings.NewOptionalType(bindings.NewIntegerType())
    fieldNameMap["prefix"] = "Prefix"
    fields["default_gateway"] = bindings.NewOptionalType(bindings.NewStringType())
    fieldNameMap["default_gateway"] = "DefaultGateway"
    var validators = []bindings.Validator{}
    uv1 := bindings.NewUnionValidator("mode",
        map[string][]bindings.FieldData {
            "STATIC": []bindings.FieldData {
                 bindings.NewFieldData("address", true),
                 bindings.NewFieldData("prefix", true),
                 bindings.NewFieldData("default_gateway", true),
            },
            "DHCP": []bindings.FieldData {
                 bindings.NewFieldData("address", true),
                 bindings.NewFieldData("prefix", true),
                 bindings.NewFieldData("default_gateway", true),
            },
            "UNCONFIGURED": []bindings.FieldData {},
        },
    )
    validators = append(validators, uv1)
    return bindings.NewStructType("com.vmware.appliance.networking.interfaces.ipv4.info",fields, reflect.TypeOf(Ipv4Info{}), fieldNameMap, validators)
}


