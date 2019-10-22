/* Copyright © 2019 VMware, Inc. All Rights Reserved.
     SPDX-License-Identifier: BSD-2-Clause */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Data type definitions file for service: EdgeClusters.
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




// The ``ManagementNetworkSpec`` contains information related to network configuration for the management network interface on the NSX Edge node. **Warning:** This class is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
 type EdgeClustersManagementNetworkSpec struct {
    // Identifier for the network. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    Network string
    // Subnet mask for the interface on the network. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    SubnetMask string
    // Gateway IP address for the interface on the network. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    Gateway string
}


// TODO some error doesn't have ErrorType. Resolve this later VAPI-3009
//


// The ``OverlayNetworkSpec`` contains information related to network configuration for the overlay network interface on the NSX Edge node. **Warning:** This class is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
 type EdgeClustersOverlayNetworkSpec struct {
    // The VLAN identifier for the network. Possible values: 
//
// * A value from 1 to 4094 specifies a VLAN ID for the network.
//
//  Any other value (including 0) is invalid. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    Vlan int64
    // Specifies if a new NSX IP pool must be used for IP address allocation for this network interface. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    CreateIpPool bool
    // New NSX IP pool information. If EdgeClustersOverlayNetworkSpec#createIpPool is set to true, then this field should contain the specificaton for the new NSX IP pool to be created and used for the network interface. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    IpPoolSpec *IpPoolCreateSpec
    // ID(Name) of the existing NSX IP pool. If EdgeClustersOverlayNetworkSpec#createIpPool is set to false, then this field should contain the ID(Name) of an existing NSX IP pool to be used for the network interface. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    IpPool *string
}


// TODO some error doesn't have ErrorType. Resolve this later VAPI-3009
//


// The ``UplinkNetworkSpec`` contains information related to configuration of the uplink network interface on the NSX Edge node(s). **Warning:** This class is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
 type EdgeClustersUplinkNetworkSpec struct {
    // The VLAN identifier for the network. Possible values: 
//
// * A value from 1 to 4094 specifies a VLAN ID for the network.
//
//  Any other value (including 0) is invalid. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    Vlan int64
    // Subnet mask for the interface on the network. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    SubnetMask string
    // Gateway IP address for the interface on the network. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    Gateway string
}


// TODO some error doesn't have ErrorType. Resolve this later VAPI-3009
//


// The ``PlacementSpec`` contains the placement specification required to set up NSX Edge node(s). **Warning:** This class is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
 type EdgeClustersPlacementSpec struct {
    // Specifies if a new resource pool should be created. If {#member createResourcePool} is set to true, a new resource pool with the name EdgeClustersPlacementSpec#resourcePoolName will be created and used for edge node deployment. If {#member createResourcePool} is set to false and {#member resourcePool} is set, then the existing resource pool identified by EdgeClustersPlacementSpec#resourcePool will be used for edge node deployment. If {#member createResourcePool} is set to false and EdgeClustersPlacementSpec#resourcePool unset, then the root resource pool will be associated with the edge node(s). **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    CreateResourcePool bool
    // Identifier for the resource pool associated with the NSX Edge node(s). **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    ResourcePool *string
    // Name of the new resource pool to be created and associated with the NSX Edge node(s). **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    ResourcePoolName *string
}


// TODO some error doesn't have ErrorType. Resolve this later VAPI-3009
//


// The ``NodeSpec`` class contains the specification required to set up NSX Edge node. **Warning:** This class is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
 type EdgeClustersNodeSpec struct {
    // Name of the virtual machine for the NSX Edge node. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    Name string
    // Password of the virtual machine for the NSX Edge node. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    Password string
    // A fully qualified domain name (FQDN) for the NSX Edge node. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    Hostname string
    // IP address for the interface on the management network. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    ManagementIpAddress string
    // Identifier for the datastore associated with the NSX Edge node. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    Datastore string
}


// TODO some error doesn't have ErrorType. Resolve this later VAPI-3009
//


// The ``UplinkNodeSpec`` class contains the specification required to set up an NSX Edge node configured with the uplink network. **Warning:** This class is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
 type EdgeClustersUplinkNodeSpec struct {
    // IP address for the interface on the uplink network. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    UplinkIpAddress string
    // Name of the virtual machine for the NSX Edge node. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    Name string
    // Password of the virtual machine for the NSX Edge node. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    Password string
    // A fully qualified domain name (FQDN) for the NSX Edge node. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    Hostname string
    // IP address for the interface on the management network. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    ManagementIpAddress string
    // Identifier for the datastore associated with the NSX Edge node. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    Datastore string
}


// TODO some error doesn't have ErrorType. Resolve this later VAPI-3009
//


// The ``EnableSpec`` class contains the specification required to set up NSX Edge node(s). **Warning:** This class is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
 type EdgeClustersEnableSpec struct {
    // List of specifications for the NSX Edge node(s). This list is only is allowed to have a maximum of two elements. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    EdgeNodes []EdgeClustersUplinkNodeSpec
    // Specification for the placement associated with the NSX Edge node(s). **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    Placement EdgeClustersPlacementSpec
    // Specification for the overlay network for the NSX Edge node(s). **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    OverlayNetwork EdgeClustersOverlayNetworkSpec
    // IP address of the NSX Edge node on the management network interface. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    ManagementNetwork EdgeClustersManagementNetworkSpec
    // IP address of the NSX Edge node on the uplink network interface. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    UplinkNetwork EdgeClustersUplinkNetworkSpec
}


// TODO some error doesn't have ErrorType. Resolve this later VAPI-3009
//





func edgeClustersEnableInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["cluster"] = bindings.NewIdType([]string {"ClusterComputeResource"}, "")
    fields["spec"] = bindings.NewReferenceType(EdgeClustersEnableSpecBindingType)
    fieldNameMap["cluster"] = "Cluster"
    fieldNameMap["spec"] = "Spec"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func edgeClustersEnableOutputType() bindings.BindingType {
    return bindings.NewVoidType()
}

func edgeClustersEnableRestMetadata() protocol.OperationRestMetadata {
    paramsTypeMap := map[string]bindings.BindingType{}
    pathParams := map[string]string{}
    queryParams := map[string]string{}
    headerParams := map[string]string{}
    paramsTypeMap["cluster"] = bindings.NewIdType([]string {"ClusterComputeResource"}, "")
    paramsTypeMap["spec"] = bindings.NewReferenceType(EdgeClustersEnableSpecBindingType)
    paramsTypeMap["cluster"] = bindings.NewIdType([]string {"ClusterComputeResource"}, "")
    pathParams["cluster"] = "cluster"
    resultHeaders := map[string]string{}
    errorHeaders := map[string]string{}
    return protocol.NewOperationRestMetadata(
      paramsTypeMap,
      pathParams,
      queryParams,
      headerParams,
      "spec",
      "POST",
      "/vcenter/nsx/edge-clusters/{cluster}",
       resultHeaders,
       204,
       errorHeaders,
       map[string]int{"AlreadyExists": 400,"InvalidArgument": 400,"NotAllowedInCurrentState": 400,"NotFound": 404,"Error": 500})
}



func EdgeClustersManagementNetworkSpecBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["network"] = bindings.NewIdType([]string {"Network"}, "")
    fieldNameMap["network"] = "Network"
    fields["subnet_mask"] = bindings.NewStringType()
    fieldNameMap["subnet_mask"] = "SubnetMask"
    fields["gateway"] = bindings.NewStringType()
    fieldNameMap["gateway"] = "Gateway"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.nsx.edge_clusters.management_network_spec",fields, reflect.TypeOf(EdgeClustersManagementNetworkSpec{}), fieldNameMap, validators)
}

func EdgeClustersOverlayNetworkSpecBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["vlan"] = bindings.NewIntegerType()
    fieldNameMap["vlan"] = "Vlan"
    fields["create_ip_pool"] = bindings.NewBooleanType()
    fieldNameMap["create_ip_pool"] = "CreateIpPool"
    fields["ip_pool_spec"] = bindings.NewOptionalType(bindings.NewReferenceType(IpPoolCreateSpecBindingType))
    fieldNameMap["ip_pool_spec"] = "IpPoolSpec"
    fields["ip_pool"] = bindings.NewOptionalType(bindings.NewIdType([]string {"com.vmware.nsx.pools.ip_pool"}, ""))
    fieldNameMap["ip_pool"] = "IpPool"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.nsx.edge_clusters.overlay_network_spec",fields, reflect.TypeOf(EdgeClustersOverlayNetworkSpec{}), fieldNameMap, validators)
}

func EdgeClustersUplinkNetworkSpecBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["vlan"] = bindings.NewIntegerType()
    fieldNameMap["vlan"] = "Vlan"
    fields["subnet_mask"] = bindings.NewStringType()
    fieldNameMap["subnet_mask"] = "SubnetMask"
    fields["gateway"] = bindings.NewStringType()
    fieldNameMap["gateway"] = "Gateway"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.nsx.edge_clusters.uplink_network_spec",fields, reflect.TypeOf(EdgeClustersUplinkNetworkSpec{}), fieldNameMap, validators)
}

func EdgeClustersPlacementSpecBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["create_resource_pool"] = bindings.NewBooleanType()
    fieldNameMap["create_resource_pool"] = "CreateResourcePool"
    fields["resource_pool"] = bindings.NewOptionalType(bindings.NewIdType([]string {"ResourcePool"}, ""))
    fieldNameMap["resource_pool"] = "ResourcePool"
    fields["resource_pool_name"] = bindings.NewOptionalType(bindings.NewStringType())
    fieldNameMap["resource_pool_name"] = "ResourcePoolName"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.nsx.edge_clusters.placement_spec",fields, reflect.TypeOf(EdgeClustersPlacementSpec{}), fieldNameMap, validators)
}

func EdgeClustersNodeSpecBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["name"] = bindings.NewStringType()
    fieldNameMap["name"] = "Name"
    fields["password"] = bindings.NewSecretType()
    fieldNameMap["password"] = "Password"
    fields["hostname"] = bindings.NewStringType()
    fieldNameMap["hostname"] = "Hostname"
    fields["management_ip_address"] = bindings.NewStringType()
    fieldNameMap["management_ip_address"] = "ManagementIpAddress"
    fields["datastore"] = bindings.NewIdType([]string {"Datastore"}, "")
    fieldNameMap["datastore"] = "Datastore"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.nsx.edge_clusters.node_spec",fields, reflect.TypeOf(EdgeClustersNodeSpec{}), fieldNameMap, validators)
}

func EdgeClustersUplinkNodeSpecBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["uplink_ip_address"] = bindings.NewStringType()
    fieldNameMap["uplink_ip_address"] = "UplinkIpAddress"
    fields["name"] = bindings.NewStringType()
    fieldNameMap["name"] = "Name"
    fields["password"] = bindings.NewSecretType()
    fieldNameMap["password"] = "Password"
    fields["hostname"] = bindings.NewStringType()
    fieldNameMap["hostname"] = "Hostname"
    fields["management_ip_address"] = bindings.NewStringType()
    fieldNameMap["management_ip_address"] = "ManagementIpAddress"
    fields["datastore"] = bindings.NewIdType([]string {"Datastore"}, "")
    fieldNameMap["datastore"] = "Datastore"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.nsx.edge_clusters.uplink_node_spec",fields, reflect.TypeOf(EdgeClustersUplinkNodeSpec{}), fieldNameMap, validators)
}

func EdgeClustersEnableSpecBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["edge_nodes"] = bindings.NewListType(bindings.NewReferenceType(EdgeClustersUplinkNodeSpecBindingType), reflect.TypeOf([]EdgeClustersUplinkNodeSpec{}))
    fieldNameMap["edge_nodes"] = "EdgeNodes"
    fields["placement"] = bindings.NewReferenceType(EdgeClustersPlacementSpecBindingType)
    fieldNameMap["placement"] = "Placement"
    fields["overlay_network"] = bindings.NewReferenceType(EdgeClustersOverlayNetworkSpecBindingType)
    fieldNameMap["overlay_network"] = "OverlayNetwork"
    fields["management_network"] = bindings.NewReferenceType(EdgeClustersManagementNetworkSpecBindingType)
    fieldNameMap["management_network"] = "ManagementNetwork"
    fields["uplink_network"] = bindings.NewReferenceType(EdgeClustersUplinkNetworkSpecBindingType)
    fieldNameMap["uplink_network"] = "UplinkNetwork"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.nsx.edge_clusters.enable_spec",fields, reflect.TypeOf(EdgeClustersEnableSpec{}), fieldNameMap, validators)
}


