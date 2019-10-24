/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Data type definitions file for service: Clusters.
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



// The ``NetworkSpec`` contains information related to network configuration for the Tunnel Endpoints (TEPs) on the hosts in the VC cluster. **Warning:** This class is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
type ClustersNetworkSpec struct {
    // The VLAN identifier for this network. Possible values: 
    //
    // * A value of 0 specifies that the network is not associated with a VLAN.
    // * A value from 1 to 4094 specifies a VLAN ID for the network.
    //
    //  Any other value is invalid. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
	Vlan int64
    // The IP address allocation mode. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
	Mode IpAllocationMode
    // Specifies if a new NSX IP pool must be used for IP address allocation for this network interface. This field is required if ClustersNetworkSpec#mode is set to IP_POOL. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
	CreateIpPool *bool
    // New NSX IP pool information. If the ClustersNetworkSpec#mode is set to IP_POOL and ClustersNetworkSpec#createIpPool is set to true, then this field contains the specification for the new NSX IP pool to be created and used for the network interface. The IP address from this NSX IP pool is assigned to a new vmknic on each host in the vSphere cluster. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
	IpPoolSpec *IpPoolCreateSpec
    // ID(Name) of the existing NSX IP pool. If the ClustersNetworkSpec#mode is set to IP_POOL and ClustersNetworkSpec#createIpPool is set to false, then this field contains the ID of an existing NSX IP pool to be used for the network interface. The IP address from this NSX IP pool is assigned to a new vmknic on each host in the vSphere cluster. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
	IpPool *string
}

// The ``EnableSpec`` class contains the specification to configure NSX networking on a vSphere cluster. **Warning:** This class is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
type ClustersEnableSpec struct {
    // The vSphere Distributed Switch (VDS) to be used for hosts' transport node configuration. All the hosts in the vSphere cluster must be connected to this VDS. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
	DvSwitch string
    // The network specification for the hosts' transport node configuration. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
	OverlayNetwork ClustersNetworkSpec
}



func clustersEnableInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["cluster"] = bindings.NewIdType([]string{"ClusterComputeResource"}, "")
	fields["spec"] = bindings.NewReferenceType(ClustersEnableSpecBindingType)
	fieldNameMap["cluster"] = "Cluster"
	fieldNameMap["spec"] = "Spec"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func clustersEnableOutputType() bindings.BindingType {
	return bindings.NewVoidType()
}

func clustersEnableRestMetadata() protocol.OperationRestMetadata {
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	paramsTypeMap["cluster"] = bindings.NewIdType([]string{"ClusterComputeResource"}, "")
	paramsTypeMap["spec"] = bindings.NewReferenceType(ClustersEnableSpecBindingType)
	paramsTypeMap["cluster"] = bindings.NewIdType([]string{"ClusterComputeResource"}, "")
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
		"/vcenter/nsx/clusters/{cluster}",
		resultHeaders,
		204,
		errorHeaders,
		map[string]int{"NotFound": 404,"AlreadyExists": 400,"InvalidArgument": 400,"Error": 500})
}


func ClustersNetworkSpecBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["vlan"] = bindings.NewIntegerType()
	fieldNameMap["vlan"] = "Vlan"
	fields["mode"] = bindings.NewEnumType("com.vmware.vcenter.nsx.ip_allocation_mode", reflect.TypeOf(IpAllocationMode(IpAllocationMode_DHCP)))
	fieldNameMap["mode"] = "Mode"
	fields["create_ip_pool"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fieldNameMap["create_ip_pool"] = "CreateIpPool"
	fields["ip_pool_spec"] = bindings.NewOptionalType(bindings.NewReferenceType(IpPoolCreateSpecBindingType))
	fieldNameMap["ip_pool_spec"] = "IpPoolSpec"
	fields["ip_pool"] = bindings.NewOptionalType(bindings.NewIdType([]string{"com.vmware.nsx.pools.ip_pool"}, ""))
	fieldNameMap["ip_pool"] = "IpPool"
	var validators = []bindings.Validator{}
	uv1 := bindings.NewUnionValidator("mode",
		map[string][]bindings.FieldData{
			"IP_POOL": []bindings.FieldData{
				bindings.NewFieldData("create_ip_pool", true),
				bindings.NewFieldData("ip_pool_spec", false),
				bindings.NewFieldData("ip_pool", false),
			},
			"DHCP": []bindings.FieldData{},
		},
	)
	validators = append(validators, uv1)
	return bindings.NewStructType("com.vmware.vcenter.nsx.clusters.network_spec", fields, reflect.TypeOf(ClustersNetworkSpec{}), fieldNameMap, validators)
}

func ClustersEnableSpecBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["dv_switch"] = bindings.NewIdType([]string{"DistributedVirtualSwitch"}, "")
	fieldNameMap["dv_switch"] = "DvSwitch"
	fields["overlay_network"] = bindings.NewReferenceType(ClustersNetworkSpecBindingType)
	fieldNameMap["overlay_network"] = "OverlayNetwork"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vcenter.nsx.clusters.enable_spec", fields, reflect.TypeOf(ClustersEnableSpec{}), fieldNameMap, validators)
}

