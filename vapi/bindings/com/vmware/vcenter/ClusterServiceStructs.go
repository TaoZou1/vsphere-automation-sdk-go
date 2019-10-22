/* Copyright © 2019 VMware, Inc. All Rights Reserved.
     SPDX-License-Identifier: BSD-2-Clause */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Data type definitions file for service: Cluster.
 * Includes binding types of a structures and enumerations defined in the service.
 * Shared by client-side stubs and server-side skeletons to ensure type
 * compatibility.
 */

package vcenter

import (
    "reflect"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/bindings"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/data"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/protocol"
)

// The resource type for the vCenter Cluster
const Cluster_RESOURCE_TYPE = "ClusterComputeResource"



// The ``FilterSpec`` class contains properties used to filter the results when listing clusters (see Cluster#list). If multiple properties are specified, only clusters matching all of the properties match the filter.
 type ClusterFilterSpec struct {
    // Identifiers of clusters that can match the filter.
    Clusters map[string]bool
    // Names that clusters must have to match the filter (see ClusterInfo#name).
    Names map[string]bool
    // Folders that must contain the cluster for the cluster to match the filter.
    Folders map[string]bool
    // Datacenters that must contain the cluster for the cluster to match the filter.
    Datacenters map[string]bool
}


// TODO some error doesn't have ErrorType. Resolve this later VAPI-3009
//


// The ``Summary`` class contains commonly used information about a cluster in vCenter Server.
 type ClusterSummary struct {
    // Identifier of the cluster.
    Cluster string
    // Name of the cluster.
    Name string
    // Flag indicating whether the vSphere HA feature is enabled for the cluster.
    HaEnabled bool
    // Flag indicating whether the vSphere DRS service is enabled for the cluster.
    DrsEnabled bool
}


// TODO some error doesn't have ErrorType. Resolve this later VAPI-3009
//


// The ``Info`` class contains information about a cluster in vCenter Server.
 type ClusterInfo struct {
    // The name of the cluster
    Name string
    // Identifier of the root resource pool of the cluster
    ResourcePool string
}


// TODO some error doesn't have ErrorType. Resolve this later VAPI-3009
//


// The ``CreateSpec`` class defines the information used to create a cluster. **Warning:** This class is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
 type ClusterCreateSpec struct {
    // The name of the cluster to be created. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    Name string
    // Identifier of the folder in which the new cluster is to be created. The folder must support vSphere compute resource as child entity. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    Parent string
}


// TODO some error doesn't have ErrorType. Resolve this later VAPI-3009
//





func clusterListInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["filter"] = bindings.NewOptionalType(bindings.NewReferenceType(ClusterFilterSpecBindingType))
    fieldNameMap["filter"] = "Filter"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func clusterListOutputType() bindings.BindingType {
    return bindings.NewListType(bindings.NewReferenceType(ClusterSummaryBindingType), reflect.TypeOf([]ClusterSummary{}))
}

func clusterListRestMetadata() protocol.OperationRestMetadata {
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
       map[string]int{"UnableToAllocateResource": 500,"ServiceUnavailable": 503,"Unauthenticated": 401,"Unauthorized": 403})
}


func clusterGetInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["cluster"] = bindings.NewIdType([]string {"ClusterComputeResource"}, "")
    fieldNameMap["cluster"] = "Cluster"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func clusterGetOutputType() bindings.BindingType {
    return bindings.NewReferenceType(ClusterInfoBindingType)
}

func clusterGetRestMetadata() protocol.OperationRestMetadata {
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
       map[string]int{"NotFound": 404,"ServiceUnavailable": 503,"Unauthenticated": 401,"Unauthorized": 403})
}


func clusterCreateInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["spec"] = bindings.NewReferenceType(ClusterCreateSpecBindingType)
    fieldNameMap["spec"] = "Spec"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func clusterCreateOutputType() bindings.BindingType {
    return bindings.NewIdType([]string {"ClusterComputeResource"}, "")
}

func clusterCreateRestMetadata() protocol.OperationRestMetadata {
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
       map[string]int{"AlreadyExists": 400,"InvalidArgument": 400,"InvalidElementType": 400,"Unauthenticated": 401,"Unauthorized": 403})
}


func clusterDeleteInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["cluster"] = bindings.NewIdType([]string {"ClusterComputeResource"}, "")
    fieldNameMap["cluster"] = "Cluster"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func clusterDeleteOutputType() bindings.BindingType {
    return bindings.NewVoidType()
}

func clusterDeleteRestMetadata() protocol.OperationRestMetadata {
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
       map[string]int{"Error": 500,"NotFound": 404,"ResourceInUse": 400,"Unauthenticated": 401,"Unauthorized": 403})
}



func ClusterFilterSpecBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["clusters"] = bindings.NewOptionalType(bindings.NewSetType(bindings.NewIdType([]string {"ClusterComputeResource"}, ""), reflect.TypeOf(map[string]bool{})))
    fieldNameMap["clusters"] = "Clusters"
    fields["names"] = bindings.NewOptionalType(bindings.NewSetType(bindings.NewStringType(), reflect.TypeOf(map[string]bool{})))
    fieldNameMap["names"] = "Names"
    fields["folders"] = bindings.NewOptionalType(bindings.NewSetType(bindings.NewIdType([]string {"Folder"}, ""), reflect.TypeOf(map[string]bool{})))
    fieldNameMap["folders"] = "Folders"
    fields["datacenters"] = bindings.NewOptionalType(bindings.NewSetType(bindings.NewIdType([]string {"Datacenter"}, ""), reflect.TypeOf(map[string]bool{})))
    fieldNameMap["datacenters"] = "Datacenters"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.cluster.filter_spec",fields, reflect.TypeOf(ClusterFilterSpec{}), fieldNameMap, validators)
}

func ClusterSummaryBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["cluster"] = bindings.NewIdType([]string {"ClusterComputeResource"}, "")
    fieldNameMap["cluster"] = "Cluster"
    fields["name"] = bindings.NewStringType()
    fieldNameMap["name"] = "Name"
    fields["ha_enabled"] = bindings.NewBooleanType()
    fieldNameMap["ha_enabled"] = "HaEnabled"
    fields["drs_enabled"] = bindings.NewBooleanType()
    fieldNameMap["drs_enabled"] = "DrsEnabled"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.cluster.summary",fields, reflect.TypeOf(ClusterSummary{}), fieldNameMap, validators)
}

func ClusterInfoBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["name"] = bindings.NewStringType()
    fieldNameMap["name"] = "Name"
    fields["resource_pool"] = bindings.NewIdType([]string {"ResourcePool"}, "")
    fieldNameMap["resource_pool"] = "ResourcePool"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.cluster.info",fields, reflect.TypeOf(ClusterInfo{}), fieldNameMap, validators)
}

func ClusterCreateSpecBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["name"] = bindings.NewStringType()
    fieldNameMap["name"] = "Name"
    fields["parent"] = bindings.NewIdType([]string {"Folder"}, "")
    fieldNameMap["parent"] = "Parent"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.cluster.create_spec",fields, reflect.TypeOf(ClusterCreateSpec{}), fieldNameMap, validators)
}


