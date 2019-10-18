/* Copyright © 2019 VMware, Inc. All Rights Reserved.
     SPDX-License-Identifier: BSD-2-Clause */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Data type definitions file for service: IpPools.
 * Includes binding types of a structures and enumerations defined in the service.
 * Shared by client-side stubs and server-side skeletons to ensure type
 * compatibility.
 */

package nsx

import (
    "reflect"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/bindings"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/data"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/protocol"
)





// The ``Summary`` class contains the basic information about the NSX IP pool. **Warning:** This class is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
 type IpPoolsSummary struct {
    // The ID(Name) of the NSX IP pool. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    IpPool string
    // The VLAN ID associated with the NSX IP pool. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    Vlan int64
}







// The ``Info`` class contains the specification for NSX IP pool. **Warning:** This class is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
 type IpPoolsInfo struct {
    // The vlan ID that this NSX IP Pools is associated with. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    Vlan int64
    // The format for IP range specification. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    Notation Info_Ipv4RangeNotation
    // The CIDR string denoting the range of IP addresses and the prefix. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    Cidr *Ipv4Cidr
    // The IPv4 address range. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    IpRange *Ipv4Range
    // The IPv4 address of the gateway associated with the NSX IP pool. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    Gateway string
}




    
    // The ``Ipv4RangeNotation`` enumeration class defines different notations used to represent range of IP addresses. **Warning:** This enumeration is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    //
    // <p> See {@link com.vmware.vapi.bindings.ApiEnumeration enumerated types description}.
     
    type Info_Ipv4RangeNotation string

    const (
        // The IPv4 addresses are represented as a CIDR. **Warning:** This constant field is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
         Info_Ipv4RangeNotation_CIDR Info_Ipv4RangeNotation = "CIDR"
        // The IPv4 addresses are represented as a range of IP addresses. **Warning:** This constant field is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
         Info_Ipv4RangeNotation_IP_RANGE Info_Ipv4RangeNotation = "IP_RANGE"
    )

    func (i Info_Ipv4RangeNotation) Info_Ipv4RangeNotation() bool {
        switch i {
            case Info_Ipv4RangeNotation_CIDR:
                return true
            case Info_Ipv4RangeNotation_IP_RANGE:
                return true
            default:
                return false
        }
    }






func ipPoolsListInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func ipPoolsListOutputType() bindings.BindingType {
    return bindings.NewListType(bindings.NewReferenceType(IpPoolsSummaryBindingType), reflect.TypeOf([]IpPoolsSummary{}))
}

func ipPoolsListRestMetadata() protocol.OperationRestMetadata {
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
      "GET",
      "/vcenter/nsx/ip-pools",
       resultHeaders,
       200,
       errorHeaders,
       map[string]int{"Error": 500})
}


func ipPoolsGetInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["ip_pool"] = bindings.NewIdType([]string {"com.vmware.nsx.pools.ip_pool"}, "")
    fieldNameMap["ip_pool"] = "IpPool"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func ipPoolsGetOutputType() bindings.BindingType {
    return bindings.NewReferenceType(IpPoolsInfoBindingType)
}

func ipPoolsGetRestMetadata() protocol.OperationRestMetadata {
    paramsTypeMap := map[string]bindings.BindingType{}
    pathParams := map[string]string{}
    queryParams := map[string]string{}
    headerParams := map[string]string{}
    paramsTypeMap["ip_pool"] = bindings.NewIdType([]string {"com.vmware.nsx.pools.ip_pool"}, "")
    paramsTypeMap["ipPool"] = bindings.NewIdType([]string {"com.vmware.nsx.pools.ip_pool"}, "")
    pathParams["ip_pool"] = "ipPool"
    resultHeaders := map[string]string{}
    errorHeaders := map[string]string{}
    return protocol.NewOperationRestMetadata(
      paramsTypeMap,
      pathParams,
      queryParams,
      headerParams,
      "",
      "GET",
      "/vcenter/nsx/ip-pools/{ipPool}",
       resultHeaders,
       200,
       errorHeaders,
       map[string]int{"NotFound": 404,"Error": 500})
}



func IpPoolsSummaryBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["ip_pool"] = bindings.NewIdType([]string {"com.vmware.nsx.pools.ip_pool"}, "")
    fieldNameMap["ip_pool"] = "IpPool"
    fields["vlan"] = bindings.NewIntegerType()
    fieldNameMap["vlan"] = "Vlan"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.nsx.ip_pools.summary",fields, reflect.TypeOf(IpPoolsSummary{}), fieldNameMap, validators)
}

func IpPoolsInfoBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["vlan"] = bindings.NewIntegerType()
    fieldNameMap["vlan"] = "Vlan"
    fields["notation"] = bindings.NewEnumType("com.vmware.vcenter.nsx.ip_pools.info.ipv4_range_notation", reflect.TypeOf(Info_Ipv4RangeNotation(Info_Ipv4RangeNotation_CIDR)))
    fieldNameMap["notation"] = "Notation"
    fields["cidr"] = bindings.NewOptionalType(bindings.NewReferenceType(Ipv4CidrBindingType))
    fieldNameMap["cidr"] = "Cidr"
    fields["ip_range"] = bindings.NewOptionalType(bindings.NewReferenceType(Ipv4RangeBindingType))
    fieldNameMap["ip_range"] = "IpRange"
    fields["gateway"] = bindings.NewStringType()
    fieldNameMap["gateway"] = "Gateway"
    var validators = []bindings.Validator{}
    uv1 := bindings.NewUnionValidator("notation",
        map[string][]bindings.FieldData {
            "CIDR": []bindings.FieldData {
                 bindings.NewFieldData("cidr", true),
            },
            "IP_RANGE": []bindings.FieldData {
                 bindings.NewFieldData("ip_range", true),
            },
        },
    )
    validators = append(validators, uv1)
    return bindings.NewStructType("com.vmware.vcenter.nsx.ip_pools.info",fields, reflect.TypeOf(IpPoolsInfo{}), fieldNameMap, validators)
}


