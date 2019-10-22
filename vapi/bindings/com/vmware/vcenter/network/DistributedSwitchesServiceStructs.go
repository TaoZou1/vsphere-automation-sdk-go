/* Copyright © 2019 VMware, Inc. All Rights Reserved.
     SPDX-License-Identifier: BSD-2-Clause */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Data type definitions file for service: DistributedSwitches.
 * Includes binding types of a structures and enumerations defined in the service.
 * Shared by client-side stubs and server-side skeletons to ensure type
 * compatibility.
 */

package network

import (
    "reflect"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/bindings"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/data"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/protocol"
)

// The resource type for the vCenter Distributed Switch. **Warning:** This constant field is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
const DistributedSwitches_RESOURCE_TYPE = "DistributedVirtualSwitch"



// The ``FilterSpec`` class contains properties used to filter the results when listing distributed switches (see DistributedSwitches#list). A flag is also provided to specify that only distributed switches connected to ALL specified hosts in a cluster should satisfy the filter. If multiple properties are specified, only distributed switches matching all the properties match the filter. **Warning:** This class is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
 type DistributedSwitchesFilterSpec struct {
    // Minimum VDS version that distributed switches must have to match the filter. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    MinVersion *string
    // Minimum size of the Maximum Transmission Unit (MTU) that this VDS must permit for its network raffic, in order to match the filter. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    MinMtu *int64
    // Clusters that must contain hosts that connect to the VDS for the VDS to match the filter. When specifying clusters this way, then for each cluster provided, the value of DistributedSwitchesFilterSpec#connectedToAllHosts is used to control whether ANY hosts or ALL hosts in it must connect to the VDS in order for the VDS to match this filter. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    Clusters map[string]bool
    // Flag to alter the filtering behavior among clustered hosts. If false, a VDS will match a cluster if ANY hosts in the cluster are connected to it. If true, a VDS will only match a cluster if ALL hosts in the cluster are connected to it (see DistributedSwitchesFilterSpec#clusters). **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    ConnectedToAllHosts *bool
}


// TODO some error doesn't have ErrorType. Resolve this later VAPI-3009
//


// The ``Summary`` class contains commonly used information about a VDS. **Warning:** This class is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
 type DistributedSwitchesSummary struct {
    // Identifier of the vSphere Distributed Switch (VDS). **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    DistributedSwitch string
    // Name of the VDS. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    Name string
    // UUID of the VDS. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    Uuid string
    // The VDS implementation version of the VDS. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    Version string
}


// TODO some error doesn't have ErrorType. Resolve this later VAPI-3009
//





func distributedSwitchesListInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["filter"] = bindings.NewOptionalType(bindings.NewReferenceType(DistributedSwitchesFilterSpecBindingType))
    fieldNameMap["filter"] = "Filter"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func distributedSwitchesListOutputType() bindings.BindingType {
    return bindings.NewListType(bindings.NewReferenceType(DistributedSwitchesSummaryBindingType), reflect.TypeOf([]DistributedSwitchesSummary{}))
}

func distributedSwitchesListRestMetadata() protocol.OperationRestMetadata {
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
       map[string]int{"InvalidArgument": 400,"UnableToAllocateResource": 500,"Unauthenticated": 401,"Unauthorized": 403})
}



func DistributedSwitchesFilterSpecBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["min_version"] = bindings.NewOptionalType(bindings.NewStringType())
    fieldNameMap["min_version"] = "MinVersion"
    fields["min_mtu"] = bindings.NewOptionalType(bindings.NewIntegerType())
    fieldNameMap["min_mtu"] = "MinMtu"
    fields["clusters"] = bindings.NewOptionalType(bindings.NewSetType(bindings.NewIdType([]string {"ClusterComputeResource"}, ""), reflect.TypeOf(map[string]bool{})))
    fieldNameMap["clusters"] = "Clusters"
    fields["connected_to_all_hosts"] = bindings.NewOptionalType(bindings.NewBooleanType())
    fieldNameMap["connected_to_all_hosts"] = "ConnectedToAllHosts"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.network.distributed_switches.filter_spec",fields, reflect.TypeOf(DistributedSwitchesFilterSpec{}), fieldNameMap, validators)
}

func DistributedSwitchesSummaryBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["distributed_switch"] = bindings.NewIdType([]string {"DistributedVirtualSwitch"}, "")
    fieldNameMap["distributed_switch"] = "DistributedSwitch"
    fields["name"] = bindings.NewStringType()
    fieldNameMap["name"] = "Name"
    fields["uuid"] = bindings.NewStringType()
    fieldNameMap["uuid"] = "Uuid"
    fields["version"] = bindings.NewStringType()
    fieldNameMap["version"] = "Version"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.network.distributed_switches.summary",fields, reflect.TypeOf(DistributedSwitchesSummary{}), fieldNameMap, validators)
}


