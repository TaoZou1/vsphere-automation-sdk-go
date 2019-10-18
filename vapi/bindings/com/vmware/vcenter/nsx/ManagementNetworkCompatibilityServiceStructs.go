/* Copyright © 2019 VMware, Inc. All Rights Reserved.
     SPDX-License-Identifier: BSD-2-Clause */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Data type definitions file for service: ManagementNetworkCompatibility.
 * Includes binding types of a structures and enumerations defined in the service.
 * Shared by client-side stubs and server-side skeletons to ensure type
 * compatibility.
 */

package nsx

import (
    "reflect"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/bindings/com/vmware/vapi/std"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/bindings"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/data"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/protocol"
)





// The ``HostVmknicInfo`` class contains information about the network configuration of a VMKernel adapter of a host. **Warning:** This class is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
 type ManagementNetworkCompatibilityHostVmknicInfo struct {
    // Identifier of the Host. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    Host string
    // Name of the VMKernel NIC. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    Vmknic string
    // Subnet Mask for the VMKernel NIC on the network. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    SubnetMask string
    // Gateway IP address for the VMKernel NIC on the network. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    Gateway string
    // List of reasons for incompatibility of the ManagementNetworkCompatibilityHostVmknicInfo#vmknic for management network configuration. If array of reasons is not empty, then this ManagementNetworkCompatibilityHostVmknicInfo#vmknic cannot be used for the management network configuration. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    IncompatibleReasons []std.LocalizableMessage
    // List of violations of VMware Best Practice Guidelines for management network configuration. If array of violations is not empty, then this ManagementNetworkCompatibilityHostVmknicInfo#vmknic may be used for the management network configuration, but it might not work as expected. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    BestPracticeViolations []std.LocalizableMessage
}







// The ``Summary`` class contains information about the compatibility of a Distributed Virtual Portgroup (DVPG) and the network configuration of associated vmknics for setting up management networks. **Warning:** This class is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
 type ManagementNetworkCompatibilitySummary struct {
    // Identifier of the Distributed Portgroup (DVPG). **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    DistributedPortgroup string
    // Name of the Distributed Portgroup (DVPG). **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    Name string
    // The array of configurations of vmknics attached to the DVPG. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    HostVmknics []ManagementNetworkCompatibilityHostVmknicInfo
}









func managementNetworkCompatibilityListInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["cluster"] = bindings.NewIdType([]string {"ClusterComputeResource"}, "")
    fields["distributed_switch"] = bindings.NewIdType([]string {"DistributedVirtualSwitch"}, "")
    fieldNameMap["cluster"] = "Cluster"
    fieldNameMap["distributed_switch"] = "DistributedSwitch"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func managementNetworkCompatibilityListOutputType() bindings.BindingType {
    return bindings.NewListType(bindings.NewReferenceType(ManagementNetworkCompatibilitySummaryBindingType), reflect.TypeOf([]ManagementNetworkCompatibilitySummary{}))
}

func managementNetworkCompatibilityListRestMetadata() protocol.OperationRestMetadata {
    paramsTypeMap := map[string]bindings.BindingType{}
    pathParams := map[string]string{}
    queryParams := map[string]string{}
    headerParams := map[string]string{}
    paramsTypeMap["cluster"] = bindings.NewIdType([]string {"ClusterComputeResource"}, "")
    paramsTypeMap["distributed_switch"] = bindings.NewIdType([]string {"DistributedVirtualSwitch"}, "")
    queryParams["cluster"] = "cluster"
    queryParams["distributed_switch"] = "distributed_switch"
    resultHeaders := map[string]string{}
    errorHeaders := map[string]string{}
    return protocol.NewOperationRestMetadata(
      paramsTypeMap,
      pathParams,
      queryParams,
      headerParams,
      "",
      "GET",
      "/vcenter/nsx/management-network-compatibility",
       resultHeaders,
       200,
       errorHeaders,
       map[string]int{"NotFound": 404,"Error": 500})
}



func ManagementNetworkCompatibilityHostVmknicInfoBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["host"] = bindings.NewIdType([]string {"HostSystem"}, "")
    fieldNameMap["host"] = "Host"
    fields["vmknic"] = bindings.NewStringType()
    fieldNameMap["vmknic"] = "Vmknic"
    fields["subnet_mask"] = bindings.NewStringType()
    fieldNameMap["subnet_mask"] = "SubnetMask"
    fields["gateway"] = bindings.NewStringType()
    fieldNameMap["gateway"] = "Gateway"
    fields["incompatible_reasons"] = bindings.NewListType(bindings.NewReferenceType(std.LocalizableMessageBindingType), reflect.TypeOf([]std.LocalizableMessage{}))
    fieldNameMap["incompatible_reasons"] = "IncompatibleReasons"
    fields["best_practice_violations"] = bindings.NewListType(bindings.NewReferenceType(std.LocalizableMessageBindingType), reflect.TypeOf([]std.LocalizableMessage{}))
    fieldNameMap["best_practice_violations"] = "BestPracticeViolations"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.nsx.management_network_compatibility.host_vmknic_info",fields, reflect.TypeOf(ManagementNetworkCompatibilityHostVmknicInfo{}), fieldNameMap, validators)
}

func ManagementNetworkCompatibilitySummaryBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["distributed_portgroup"] = bindings.NewIdType([]string {"DistributedVirtualPortgroup"}, "")
    fieldNameMap["distributed_portgroup"] = "DistributedPortgroup"
    fields["name"] = bindings.NewStringType()
    fieldNameMap["name"] = "Name"
    fields["host_vmknics"] = bindings.NewListType(bindings.NewReferenceType(ManagementNetworkCompatibilityHostVmknicInfoBindingType), reflect.TypeOf([]ManagementNetworkCompatibilityHostVmknicInfo{}))
    fieldNameMap["host_vmknics"] = "HostVmknics"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.nsx.management_network_compatibility.summary",fields, reflect.TypeOf(ManagementNetworkCompatibilitySummary{}), fieldNameMap, validators)
}


