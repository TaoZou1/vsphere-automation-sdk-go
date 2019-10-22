/* Copyright © 2019 VMware, Inc. All Rights Reserved.
     SPDX-License-Identifier: BSD-2-Clause */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Data type definitions file for service: ClusterSizeInfo.
 * Includes binding types of a structures and enumerations defined in the service.
 * Shared by client-side stubs and server-side skeletons to ensure type
 * compatibility.
 */

package namespace_management

import (
    "reflect"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/bindings"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/data"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/protocol"
)




// The ``VmInfo`` class contains the information about the configuration of the virtual machines which would be created for Namespaces. **Warning:** This class is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
 type ClusterSizeInfoVmInfo struct {
    // Number of CPU cores. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    Count int64
    // Number of CPU cores per socket. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    CoresPerSocket int64
    // Memory size, in mebibytes. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    Memory int64
    // Overall capacity of the disks in the virtual machine, in mebibytes. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    Capacity int64
}


// TODO some error doesn't have ErrorType. Resolve this later VAPI-3009
//


// The ``Info`` class contains the information about limits associated with a ``SizingHint``. **Warning:** This class is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
 type ClusterSizeInfoInfo struct {
    // The maximum number of supported pods. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    NumSupportedPods int64
    // The maximum number of supported services. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    NumSupportedServices int64
    // Default CIDR range from which Kubernetes allocates service cluster IP addresses. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    DefaultServiceCidr Ipv4Cidr
    // Default CIDR range from which Kubernetes allocates pod IP addresses. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    DefaultPodCidr Ipv4Cidr
    // Information about Kubernetes API server virtual machine configuration. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    MasterVmInfo ClusterSizeInfoVmInfo
    // Information about worker virtual machine configuration. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    WorkerVmInfo *ClusterSizeInfoVmInfo
}


// TODO some error doesn't have ErrorType. Resolve this later VAPI-3009
//





func clusterSizeInfoGetInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func clusterSizeInfoGetOutputType() bindings.BindingType {
    return bindings.NewMapType(bindings.NewEnumType("com.vmware.vcenter.namespace_management.sizing_hint", reflect.TypeOf(SizingHint(SizingHint_TINY))), bindings.NewReferenceType(ClusterSizeInfoInfoBindingType),reflect.TypeOf(map[SizingHint]ClusterSizeInfoInfo{}))
}

func clusterSizeInfoGetRestMetadata() protocol.OperationRestMetadata {
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
      "GET",
      "/vcenter/namespace-management/cluster-size-info",
       resultHeaders,
       200,
       errorHeaders,
       map[string]int{"Error": 500,"Unauthenticated": 401})
}



func ClusterSizeInfoVmInfoBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["count"] = bindings.NewIntegerType()
    fieldNameMap["count"] = "Count"
    fields["cores_per_socket"] = bindings.NewIntegerType()
    fieldNameMap["cores_per_socket"] = "CoresPerSocket"
    fields["memory"] = bindings.NewIntegerType()
    fieldNameMap["memory"] = "Memory"
    fields["capacity"] = bindings.NewIntegerType()
    fieldNameMap["capacity"] = "Capacity"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.namespace_management.cluster_size_info.vm_info",fields, reflect.TypeOf(ClusterSizeInfoVmInfo{}), fieldNameMap, validators)
}

func ClusterSizeInfoInfoBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["num_supported_pods"] = bindings.NewIntegerType()
    fieldNameMap["num_supported_pods"] = "NumSupportedPods"
    fields["num_supported_services"] = bindings.NewIntegerType()
    fieldNameMap["num_supported_services"] = "NumSupportedServices"
    fields["default_service_cidr"] = bindings.NewReferenceType(Ipv4CidrBindingType)
    fieldNameMap["default_service_cidr"] = "DefaultServiceCidr"
    fields["default_pod_cidr"] = bindings.NewReferenceType(Ipv4CidrBindingType)
    fieldNameMap["default_pod_cidr"] = "DefaultPodCidr"
    fields["master_vm_info"] = bindings.NewReferenceType(ClusterSizeInfoVmInfoBindingType)
    fieldNameMap["master_vm_info"] = "MasterVmInfo"
    fields["worker_vm_info"] = bindings.NewOptionalType(bindings.NewReferenceType(ClusterSizeInfoVmInfoBindingType))
    fieldNameMap["worker_vm_info"] = "WorkerVmInfo"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.namespace_management.cluster_size_info.info",fields, reflect.TypeOf(ClusterSizeInfoInfo{}), fieldNameMap, validators)
}


