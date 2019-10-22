/* Copyright © 2019 VMware, Inc. All Rights Reserved.
     SPDX-License-Identifier: BSD-2-Clause */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Data type definitions file for service: Validation.
 * Includes binding types of a structures and enumerations defined in the service.
 * Shared by client-side stubs and server-side skeletons to ensure type
 * compatibility.
 */

package lcm

import (
    "reflect"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/bindings"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/data"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/protocol"
)




 type ValidationApplianceNameRequest struct {
    // This subsection describes the ESX or VC on which to deploy the appliance. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    DestinationLocation ValidationApplianceNameDestinationLocation
    // Spec to describe the new appliance. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    DestinationAppliance ValidationApplianceNameDestinationAppliance
}


// TODO some error doesn't have ErrorType. Resolve this later VAPI-3009
//


// Data container for appliance name information used in validation of appliance name request. **Warning:** This class is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
 type ValidationApplianceNameDestinationAppliance struct {
    // The name of the appliance to deploy. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    ApplianceName string
}


// TODO some error doesn't have ErrorType. Resolve this later VAPI-3009
//


 type ValidationApplianceNameDestinationLocation struct {
    // This section describes the ESX host on which to deploy the appliance. Required if you are deploying the appliance directly on an ESX host. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    Esx *ValidationApplianceNameEsx
    // This subsection describes the vCenter on which to deploy the appliance. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    Vcenter *ValidationApplianceNameVc
}


// TODO some error doesn't have ErrorType. Resolve this later VAPI-3009
//


// This section describes the ESX host on which to deploy the appliance. Required if you are deploying the appliance directly on an ESX host. **Warning:** This class is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
 type ValidationApplianceNameEsx struct {
    // The configuration to connect to an ESX/VC. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    Connection Connection
}


// TODO some error doesn't have ErrorType. Resolve this later VAPI-3009
//


// The configuration of ESX inventory. **Warning:** This class is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
 type ValidationApplianceNameEsxInventory struct {
    // The path to the resource pool on the ESX host in which the appliance will be deployed. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    ResourcePoolPath *string
}


// TODO some error doesn't have ErrorType. Resolve this later VAPI-3009
//


// This subsection describes the vCenter on which to deploy the appliance. **Warning:** This class is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
 type ValidationApplianceNameVc struct {
    // The configuration to connect to an ESX/VC. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    Connection Connection
    // All names are case-sensitive. you can install the appliance to one of the following destinations: 1. A resource pool in a cluster, use 'cluster_path'. 2. A specific ESX host in a cluster, use 'host_path'. 3. A resource pool in a specific ESX host being managed by the current vCenter, use 'resource_pool_path'. You must always provide the 'network_name' key. To install a new appliance to a specific ESX host in a cluster, provide the 'host_path' key, and the 'datastore_name', e.g. 'host_path': '/MyDataCenter/host/MyCluster/10.20.30.40', 'datastore_name': 'Your Datastore'. To install a new appliance to a specific resource pool, provide the 'resource_pool_path', and the 'datastore_name', e.g. 'resource_pool_path': '/Your Datacenter Folder/Your Datacenter/host/Your Cluster/Resources/Your Resource Pool', 'datastore_name': 'Your Datastore'. To place a new appliance to a virtual machine Folder, provide the 'vm_folder_path', e.g. 'vm_folder_path': 'VM Folder 0/VM Folder1'. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    Inventory ValidationApplianceNameVcInventory
}


// TODO some error doesn't have ErrorType. Resolve this later VAPI-3009
//


// All names are case-sensitive. you can install the appliance to one of the following destinations: 1. A resource pool in a cluster, use 'cluster_path'. 2. A specific ESX host in a cluster, use 'host_path'. 3. A resource pool in a specific ESX host being managed by the current vCenter, use 'resource_pool_path'. You must always provide the 'network_name' key. To install a new appliance to a specific ESX host in a cluster, provide the 'host_path' key, and the 'datastore_name', e.g. 'host_path': '/MyDataCenter/host/MyCluster/10.20.30.40', 'datastore_name': 'Your Datastore'. To install a new appliance to a specific resource pool, provide the 'resource_pool_path', and the 'datastore_name', e.g. 'resource_pool_path': '/Your Datacenter Folder/Your Datacenter/host/Your Cluster/Resources/Your Resource Pool', 'datastore_name': 'Your Datastore'. To place a new appliance to a virtual machine Folder, provide the 'vm_folder_path', e.g. 'vm_folder_path': 'VM Folder 0/VM Folder1'. **Warning:** This class is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
 type ValidationApplianceNameVcInventory struct {
        VmFolderPath *string
    // Full path to resource pool. Format: /{datacenter folder}/{datacenter name}/host/{host name}/{cluster_name}/Resources/{resource pool}. e.g: /Your Datacenter Folder/Your Datacenter/host/Your Cluster/Resources/Your Resource Pool. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    ResourcePoolPath *string
    // Full path to the cluster. Format: /{datacenter folder}/{datacenter name}/host/{cluster_name}. e.g: /Your Datacenter Folder/Your Datacenter/host/Your Cluster. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    ClusterPath *string
    // Full path to an ESX host. Format: /{datacenter folder}/{datacenter name}/host/{host name}. e.g: /Your Datacenter Folder/Your Datacenter/host/Your Host. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    HostPath *string
}


// TODO some error doesn't have ErrorType. Resolve this later VAPI-3009
//





func validationCheckApplianceNameInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["spec"] = bindings.NewReferenceType(ValidationApplianceNameRequestBindingType)
    fieldNameMap["spec"] = "Spec"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func validationCheckApplianceNameOutputType() bindings.BindingType {
    return bindings.NewBooleanType()
}

func validationCheckApplianceNameRestMetadata() protocol.OperationRestMetadata {
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
       map[string]int{})
}



func ValidationApplianceNameRequestBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["destination_location"] = bindings.NewReferenceType(ValidationApplianceNameDestinationLocationBindingType)
    fieldNameMap["destination_location"] = "DestinationLocation"
    fields["destination_appliance"] = bindings.NewReferenceType(ValidationApplianceNameDestinationApplianceBindingType)
    fieldNameMap["destination_appliance"] = "DestinationAppliance"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.lcm.validation.appliance_name_request",fields, reflect.TypeOf(ValidationApplianceNameRequest{}), fieldNameMap, validators)
}

func ValidationApplianceNameDestinationApplianceBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["appliance_name"] = bindings.NewStringType()
    fieldNameMap["appliance_name"] = "ApplianceName"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.lcm.validation.appliance_name_destination_appliance",fields, reflect.TypeOf(ValidationApplianceNameDestinationAppliance{}), fieldNameMap, validators)
}

func ValidationApplianceNameDestinationLocationBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["esx"] = bindings.NewOptionalType(bindings.NewReferenceType(ValidationApplianceNameEsxBindingType))
    fieldNameMap["esx"] = "Esx"
    fields["vcenter"] = bindings.NewOptionalType(bindings.NewReferenceType(ValidationApplianceNameVcBindingType))
    fieldNameMap["vcenter"] = "Vcenter"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.lcm.validation.appliance_name_destination_location",fields, reflect.TypeOf(ValidationApplianceNameDestinationLocation{}), fieldNameMap, validators)
}

func ValidationApplianceNameEsxBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["connection"] = bindings.NewReferenceType(ConnectionBindingType)
    fieldNameMap["connection"] = "Connection"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.lcm.validation.appliance_name_esx",fields, reflect.TypeOf(ValidationApplianceNameEsx{}), fieldNameMap, validators)
}

func ValidationApplianceNameEsxInventoryBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["resource_pool_path"] = bindings.NewOptionalType(bindings.NewStringType())
    fieldNameMap["resource_pool_path"] = "ResourcePoolPath"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.lcm.validation.appliance_name_esx_inventory",fields, reflect.TypeOf(ValidationApplianceNameEsxInventory{}), fieldNameMap, validators)
}

func ValidationApplianceNameVcBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["connection"] = bindings.NewReferenceType(ConnectionBindingType)
    fieldNameMap["connection"] = "Connection"
    fields["inventory"] = bindings.NewReferenceType(ValidationApplianceNameVcInventoryBindingType)
    fieldNameMap["inventory"] = "Inventory"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.lcm.validation.appliance_name_vc",fields, reflect.TypeOf(ValidationApplianceNameVc{}), fieldNameMap, validators)
}

func ValidationApplianceNameVcInventoryBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["vm_folder_path"] = bindings.NewOptionalType(bindings.NewStringType())
    fieldNameMap["vm_folder_path"] = "VmFolderPath"
    fields["resource_pool_path"] = bindings.NewOptionalType(bindings.NewStringType())
    fieldNameMap["resource_pool_path"] = "ResourcePoolPath"
    fields["cluster_path"] = bindings.NewOptionalType(bindings.NewStringType())
    fieldNameMap["cluster_path"] = "ClusterPath"
    fields["host_path"] = bindings.NewOptionalType(bindings.NewStringType())
    fieldNameMap["host_path"] = "HostPath"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.lcm.validation.appliance_name_vc_inventory",fields, reflect.TypeOf(ValidationApplianceNameVcInventory{}), fieldNameMap, validators)
}


