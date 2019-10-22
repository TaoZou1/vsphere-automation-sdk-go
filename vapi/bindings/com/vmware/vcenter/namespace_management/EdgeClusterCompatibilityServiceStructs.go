/* Copyright © 2019 VMware, Inc. All Rights Reserved.
     SPDX-License-Identifier: BSD-2-Clause */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Data type definitions file for service: EdgeClusterCompatibility.
 * Includes binding types of a structures and enumerations defined in the service.
 * Shared by client-side stubs and server-side skeletons to ensure type
 * compatibility.
 */

package namespace_management

import (
    "reflect"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/bindings/com/vmware/vapi/std"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/bindings"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/data"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/protocol"
)




// The ``Summary`` class contains information about an NSX-T Edge Cluster, including its compatibility with the vCenter Namespaces feature. **Warning:** This class is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
 type EdgeClusterCompatibilitySummary struct {
    // Identifier of the Edge Cluster. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    EdgeCluster string
    // Display name of the Edge Cluster. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    DisplayName string
    // Compatibility of this Edge Cluster with Namespaces feature. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    Compatible bool
    // List of reasons for incompatibility. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    IncompatibilityReasons []std.LocalizableMessage
}


// TODO some error doesn't have ErrorType. Resolve this later VAPI-3009
//


// The ``FilterSpec`` class contains properties used to filter the results when listing Edge Clusters (see EdgeClusterCompatibility#list) and their compatibility information. **Warning:** This class is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
 type EdgeClusterCompatibilityFilterSpec struct {
    // Compatibility criteria for matching the filter. If true, only Edge Clusters which are compatible with vSphere Namespaces match the filter. If false, only Edge Clusters which are incompatible with vSphere Namespaces match the filter. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    Compatible *bool
}


// TODO some error doesn't have ErrorType. Resolve this later VAPI-3009
//





func edgeClusterCompatibilityListInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["cluster"] = bindings.NewIdType([]string {"ClusterComputeResource"}, "")
    fields["distributed_switch"] = bindings.NewIdType([]string {"vSphereDistributedSwitch"}, "")
    fields["filter"] = bindings.NewOptionalType(bindings.NewReferenceType(EdgeClusterCompatibilityFilterSpecBindingType))
    fieldNameMap["cluster"] = "Cluster"
    fieldNameMap["distributed_switch"] = "DistributedSwitch"
    fieldNameMap["filter"] = "Filter"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func edgeClusterCompatibilityListOutputType() bindings.BindingType {
    return bindings.NewListType(bindings.NewReferenceType(EdgeClusterCompatibilitySummaryBindingType), reflect.TypeOf([]EdgeClusterCompatibilitySummary{}))
}

func edgeClusterCompatibilityListRestMetadata() protocol.OperationRestMetadata {
    paramsTypeMap := map[string]bindings.BindingType{}
    pathParams := map[string]string{}
    queryParams := map[string]string{}
    headerParams := map[string]string{}
    paramsTypeMap["filter.compatible"] = bindings.NewOptionalType(bindings.NewBooleanType())
    paramsTypeMap["cluster"] = bindings.NewIdType([]string {"ClusterComputeResource"}, "")
    paramsTypeMap["distributed_switch"] = bindings.NewIdType([]string {"vSphereDistributedSwitch"}, "")
    queryParams["cluster"] = "cluster"
    queryParams["filter.compatible"] = "compatible"
    queryParams["distributed_switch"] = "distributed_switch"
    resultHeaders := map[string]string{}
    errorHeaders := map[string]string{}
    errorHeaders["Unauthenticated.challenge"] = "WWW-Authenticate"
    return protocol.NewOperationRestMetadata(
      paramsTypeMap,
      pathParams,
      queryParams,
      headerParams,
      "",
      "GET",
      "/vcenter/namespace-management/edge-cluster-compatibility",
       resultHeaders,
       200,
       errorHeaders,
       map[string]int{"Error": 500,"InternalServerError": 500,"NotFound": 404,"Unauthenticated": 401,"Unauthorized": 403})
}



func EdgeClusterCompatibilitySummaryBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["edge_cluster"] = bindings.NewIdType([]string {"NSXEdgeCluster"}, "")
    fieldNameMap["edge_cluster"] = "EdgeCluster"
    fields["display_name"] = bindings.NewStringType()
    fieldNameMap["display_name"] = "DisplayName"
    fields["compatible"] = bindings.NewBooleanType()
    fieldNameMap["compatible"] = "Compatible"
    fields["incompatibility_reasons"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewReferenceType(std.LocalizableMessageBindingType), reflect.TypeOf([]std.LocalizableMessage{})))
    fieldNameMap["incompatibility_reasons"] = "IncompatibilityReasons"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.namespace_management.edge_cluster_compatibility.summary",fields, reflect.TypeOf(EdgeClusterCompatibilitySummary{}), fieldNameMap, validators)
}

func EdgeClusterCompatibilityFilterSpecBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["compatible"] = bindings.NewOptionalType(bindings.NewBooleanType())
    fieldNameMap["compatible"] = "Compatible"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.namespace_management.edge_cluster_compatibility.filter_spec",fields, reflect.TypeOf(EdgeClusterCompatibilityFilterSpec{}), fieldNameMap, validators)
}


