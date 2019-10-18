/* Copyright © 2019 VMware, Inc. All Rights Reserved.
     SPDX-License-Identifier: BSD-2-Clause */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Data type definitions file for package: com.vmware.vcenter.namespaces.
 * Includes binding types of a top level structures and enumerations.
 * Shared by client-side stubs and server-side skeletons to ensure type
 * compatibility.
 */

package namespaces

import (
    "reflect"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/bindings"
)




// A ``Container`` holds the data that describes a container within a pod. **Warning:** This class is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
type Container struct {
    // The name of the container. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    Name string
    // The image the container is running. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    Image string
    // The container's current condition. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    Status string
    // Time at which the container was last (re-)started. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    StartedAt string
    // Time at which the container last terminated. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    FinishedAt string
    // Details about the container's current condition, if any. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    Reason string
}







// The ``ResourceQuotaOptionsV1`` class represents the resource quota limits which can be applied on the namespace. Refer to ` <https://kubernetes.io/docs/concepts/policy/resource-quotas>`_ for information related to the properties of this object and what they map to. **Warning:** This class is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
type ResourceQuotaOptionsV1 struct {
    // This is equivalent to 'limits.memory' option which is the maximum memory limit (in mebibytes) across all pods which exist in a non-terminal state in the namespace. This value translates to the memory limit on the ResourcePool in vCenter Server created for the namespace. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    MemoryLimit *int64
    // This represents the default memory limit (in mebibytes) for containers in the pod. This translates to default memory limit in a LimitRange object. Refer ` for information about LimitRange. <https://kubernetes.io/docs/tasks/administer-cluster/manage-resources/memory-default-namespace/>`_. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    MemoryLimitDefault *int64
    // This represents the default memory request (in mebibytes) for containers in the pod. This translates to default memory request in a LimitRange object. Refer ` for information about LimitRange. <https://kubernetes.io/docs/tasks/administer-cluster/manage-resources/memory-default-namespace/>`_. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    MemoryRequestDefault *int64
    // This is equivalent to 'limits.cpu' option which is the maximum CPU limit (in MHz) across all pods which exist in a non-terminal state in the namespace. If specified, this limit should be at least 10 MHz. This value translates to the CPU limit on the ResourcePool in vCenter Server created for the namespace. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    CpuLimit *int64
    // This represents the default cpu limit (in MHz) for containers in the pod. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    CpuLimitDefault *int64
    // This represents the default CPU request (in MHz) for containers in the pod. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    CpuRequestDefault *int64
    // This represents 'requests.storage' option which is the maximum storage request (in mebibytes) across all persistent volume claims from pods which exist in a non-terminal state in the namespace. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    StorageRequestLimit *int64
    // This represents 'pods' option which is the maximum number of pods which exist in a non-terminal state in the namespace. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    PodCount *int64
    // This represents 'count/services' option which is the maximum number of services in the namespace. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    ServiceCount *int64
    // This represents 'count/deployments.apps' option which is the maximum number of deployments in the namespace. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    DeploymentCount *int64
    // This represents 'count/daemonsets.apps' option which is the maximum number of DaemonSets in the namespace. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    DaemonSetCount *int64
    // This represents 'count/replicasets.apps' option which is the maximum number of ReplicaSets in the namespace. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    ReplicaSetCount *int64
    // This represents 'count/replicationcontrollers' option which is the maximum number of ReplicationControllers in the namespace. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    ReplicationControllerCount *int64
    // This represents 'count/statefulsets.apps' option which is the maximum number of StatefulSets in the namespace. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    StatefulSetCount *int64
    // This represents 'count/configmaps' option which is the maximum number of ConfigMaps in the namespace. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    ConfigMapCount *int64
    // This represents 'count/secrets' option which is the maximum number of secrets in the namespace. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    SecretCount *int64
    // This represents 'count/persistentvolumeclaims' option which is the maximum number of PersistentVolumeClaims in the namespace. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    PersistentVolumeClaimCount *int64
    // This represents 'count/jobs.batch' option which is the maximum number jobs in the namespace. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    JobCount *int64
}







// The ``ResourceQuotaOptionsV1Update`` class represents the changes to resource quota limits which are set on the namespace. Refer to ` <\a> Kubernetes Resource Quota <https://kubernetes.io/docs/concepts/policy/resource-quotas>`_ for information related to the properties of this object and what they map to. **Warning:** This class is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
type ResourceQuotaOptionsV1Update struct {
    // This represents the new value for 'limits.memory' option which is equivalent to the maximum memory limit (in mebibytes) across all pods in the namespace. This field is ignored if ResourceQuotaOptionsV1Update#memoryLimitUnset is set to ``true``. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    MemoryLimit *int64
    // This represents the intent of the change to ResourceQuotaOptionsV1Update#memoryLimit. If this field is set to ``true``, the existing memory limit on the ResourcePool is removed. If this field is set to ``false``, the existing memory limit will be changed to the value specified in ResourceQuotaOptionsV1Update#memoryLimit, if any. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    MemoryLimitUnset *bool
    // This represents the new value for the default memory limit (in mebibytes) for containers in the pod. This field is ignored if ResourceQuotaOptionsV1Update#memoryLimitDefaultUnset is set to ``true``. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    MemoryLimitDefault *int64
    // This represents the intent of the change to ResourceQuotaOptionsV1Update#memoryLimitDefault. If this field is set to ``true``, the existing default memory limit on containers in the pod is removed. If this field is set to ``false``, the existing default memory limit will be changed to the value specified in ResourceQuotaOptionsV1Update#memoryLimitDefault, if any. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    MemoryLimitDefaultUnset *bool
    // This represents the new value for the default memory request (in mebibytes) for containers in the pod. This field is ignored if ResourceQuotaOptionsV1Update#memoryRequestDefaultUnset is set to ``true``. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    MemoryRequestDefault *int64
    // This represents the intent of the change to ResourceQuotaOptionsV1Update#memoryRequestDefault. If this field is set to ``true``, the existing default memory request on containers in the pod will be removed. If this field is set to ``false``, the existing default memory request will be changed to the value specified in ResourceQuotaOptionsV1Update#memoryRequestDefault, if any. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    MemoryRequestDefaultUnset *bool
    // This represents the new value for 'limits.cpu' option which is equivalent to the maximum CPU limit (in MHz) across all pods in the namespace. This field is ignored if ResourceQuotaOptionsV1Update#cpuLimitUnset is set to ``true``. If specified, this limit should be at least 10 MHz. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    CpuLimit *int64
    // This represents the intent of the change to ResourceQuotaOptionsV1Update#cpuLimit. If this field is set to ``true``, the existing CPU limit on the ResourcePool is removed. If this field is set to ``false``, the existing CPU limit will be changed to the value specified in ResourceQuotaOptionsV1Update#cpuLimit, if any. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    CpuLimitUnset *bool
    // This represents the new value for the default CPU limit (in Mhz) for containers in the pod. This field is ignored if ResourceQuotaOptionsV1Update#cpuLimitDefaultUnset is set to ``true``. If specified, this limit should be at least 10 MHz. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    CpuLimitDefault *int64
    // This represents the intent of the change to ResourceQuotaOptionsV1Update#cpuLimitDefault. If this field is set to ``true``, the existing default CPU limit on containers in the pod is removed. If this field is set to ``false``, the existing default CPU limit will be changed to the value specified in ResourceQuotaOptionsV1Update#cpuLimitDefault, if any. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    CpuLimitDefaultUnset *bool
    // This represents the new value for the default CPU request (in Mhz) for containers in the pod. This field is ignored if ResourceQuotaOptionsV1Update#cpuRequestDefaultUnset is set to ``true``. If specified, this field should be at least 10 MHz. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    CpuRequestDefault *int64
    // This represents the intent of the change to ResourceQuotaOptionsV1Update#cpuRequestDefault. If this field is set to ``true``, the existing default CPU request on containers in the pod is removed. If this field is set to ``false``, the existing default CPU request will be changed to the value specified in ResourceQuotaOptionsV1Update#cpuRequestDefault, if any. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    CpuRequestDefaultUnset *bool
    // This represents the new value for 'requests.storage' which is the limit on storage requests (in mebibytes) across all persistent volume claims from pods in the namespace. This field is ignored if if ResourceQuotaOptionsV1Update#storageRequestLimitUnset is set to ``true``. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    StorageRequestLimit *int64
    // This represents the intent of the change to ResourceQuotaOptionsV1Update#storageRequestLimit. If this field is set to ``true``, the existing storage request limit will be reset. If this field is set to ``false``, the existing storage request limit will be changed to the value specified in ResourceQuotaOptionsV1Update#storageRequestLimit, if any. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    StorageRequestLimitUnset *bool
    // This represents the new value for 'podCount' option which is the maximum number of pods in the namespace. This field is ignored if ResourceQuotaOptionsV1Update#podCountUnset is set to ``true``. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    PodCount *int64
    // This represents the intent of the change to ResourceQuotaOptionsV1Update#podCount. If this field is set to ``true``, the existing 'podCount' limit on the namespace will be reset. If this field is set to ``false``, the existing CPU limit will be changed to the value specified in ResourceQuotaOptionsV1Update#podCount, if any. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    PodCountUnset *bool
    // This represents the new value for 'serviceCount' option which is the maximum number of services in the namespace. This field is ignored if ResourceQuotaOptionsV1Update#serviceCountUnset is set to ``true``. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    ServiceCount *int64
    // This represents the intent of the change to ResourceQuotaOptionsV1Update#serviceCount. If this field is set to ``true``, the existing 'serviceCount' limit on the namespace will be reset. If this field is set to ``false``, the existing service count limit will be changed to the value specified in ResourceQuotaOptionsV1Update#serviceCount, if any. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    ServiceCountUnset *bool
    // This represents the new value for 'deploymentCount' option which is the maximum number of deployments in the namespace. This field is ignored if ResourceQuotaOptionsV1Update#deploymentCountUnset is set to ``true``. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    DeploymentCount *int64
    // This represents the intent of the change to ResourceQuotaOptionsV1Update#deploymentCount. If this field is set to ``true``, the existing 'deploymentCount' limit on the namespace will be reset. If this field is set to ``false``, the existing deployment count limit will be changed to the value specified in ResourceQuotaOptionsV1Update#deploymentCount, if any. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    DeploymentCountUnset *bool
    // This represents the new value for 'daemonSetCount' option which is the maximum number of DaemonSets in the namespace. This field is ignored if ResourceQuotaOptionsV1Update#daemonSetCountUnset is set to ``true``. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    DaemonSetCount *int64
    // This represents the intent of the change to ResourceQuotaOptionsV1Update#daemonSetCount. If this field is set to ``true``, the existing 'daemonSetCount' limit on the namespace will be reset. If this field is set to ``false``, the existing daemonset count limit will be changed to the value specified in ResourceQuotaOptionsV1Update#daemonSetCount, if any. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    DaemonSetCountUnset *bool
    // This represents the new value for 'replicaSetCount' option which is the maximum number of ReplicaSets in the namespace. This field is ignored if ResourceQuotaOptionsV1Update#replicaSetCountUnset is set to ``true``. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    ReplicaSetCount *int64
    // This represents the intent of the change to ResourceQuotaOptionsV1Update#replicaSetCount. If this field is set to ``true``, the existing 'replicaSetCount' limit on the namespace will be reset. If this field is set to ``false``, the existing replicaset count limit will be changed to the value specified in ResourceQuotaOptionsV1Update#replicaSetCount, if any. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    ReplicaSetCountUnset *bool
    // This represents the new value for 'replicationControllerCount' option which is the maximum number of ReplicationControllers in the namespace. This field is ignored if ResourceQuotaOptionsV1Update#replicationControllerCountUnset is set to ``true``. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    ReplicationControllerCount *int64
    // This represents the intent of the change to ResourceQuotaOptionsV1Update#replicationControllerCount. If this field is set to ``true``, the existing 'replicationControllerCount' limit on the namespace will be reset. If this field is set to ``false``, the existing replicationcontroller count limit will be changed to the value specified in ResourceQuotaOptionsV1Update#replicationControllerCount, if any. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    ReplicationControllerCountUnset *bool
    // This represents the new value for 'statefulSetCount' option which is the maximum number of StatefulSets in the namespace. This field is ignored if ResourceQuotaOptionsV1Update#statefulSetCountUnset is set to ``true``. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    StatefulSetCount *int64
    // This represents the intent of the change to ResourceQuotaOptionsV1Update#statefulSetCount. If this field is set to ``true``, the existing 'statefulSetCount' limit on the namespace will be reset. If this field is set to ``false``, the existing statefulset count limit will be changed to the value specified in ResourceQuotaOptionsV1Update#statefulSetCount, if any. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    StatefulSetCountUnset *bool
    // This represents the new value for 'configMapCount' option which is the maximum number of ConfigMaps in the namespace. This field is ignored if ResourceQuotaOptionsV1Update#configMapCountUnset is set to ``true``. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    ConfigMapCount *int64
    // This represents the intent of the change to ResourceQuotaOptionsV1Update#configMapCount. If this field is set to ``true``, the existing 'configMapCount' limit on the namespace will be reset. If this field is set to ``false``, the existing configmap count limit will be changed to the value specified in ResourceQuotaOptionsV1Update#configMapCount, if any. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    ConfigMapCountUnset *bool
    // This represents the new value for 'secretCount' option which is the maximum number of secrets in the namespace. This field is ignored if ResourceQuotaOptionsV1Update#secretCountUnset is set to ``true``. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    SecretCount *int64
    // This represents the intent of the change to ResourceQuotaOptionsV1Update#secretCount. If this field is set to ``true``, the existing 'secretCount' limit on the namespace will be reset. If this field is set to ``false``, the existing secret count limit will be changed to the value specified in ResourceQuotaOptionsV1Update#secretCount, if any. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    SecretCountUnset *bool
    // This represents the new value for 'persistentVolumeClaimCount' option which is the maximum number of PersistentVolumeClaims in the namespace. This field is ignored if ResourceQuotaOptionsV1Update#persistentVolumeClaimCountUnset is set to ``true``. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    PersistentVolumeClaimCount *int64
    // This represents the intent of the change to ResourceQuotaOptionsV1Update#persistentVolumeClaimCount. If this field is set to ``true``, the existing 'persistentVolumeClaimCount' limit on the namespace will be reset. If this field is set to ``false``, the existing replicationcontroller count limit will be changed to the value specified in ResourceQuotaOptionsV1Update#persistentVolumeClaimCount, if any. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    PersistentVolumeClaimCountUnset *bool
    // This represents the new value for 'jobCount' option which is the maximum number of jobs in the namespace. This field is ignored if ResourceQuotaOptionsV1Update#jobCountUnset is set to ``true``. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    JobCount *int64
    // This represents the intent of the change to ResourceQuotaOptionsV1Update#jobCount. If this field is set to ``true``, the existing 'jobCount' limit on the namespace will be reset. If this field is set to ``false``, the existing secret count limit will be changed to the value specified in ResourceQuotaOptionsV1Update#jobCount, if any. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    JobCountUnset *bool
}










func ContainerBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["name"] = bindings.NewStringType()
    fieldNameMap["name"] = "Name"
    fields["image"] = bindings.NewStringType()
    fieldNameMap["image"] = "Image"
    fields["status"] = bindings.NewStringType()
    fieldNameMap["status"] = "Status"
    fields["started_at"] = bindings.NewStringType()
    fieldNameMap["started_at"] = "StartedAt"
    fields["finished_at"] = bindings.NewStringType()
    fieldNameMap["finished_at"] = "FinishedAt"
    fields["reason"] = bindings.NewStringType()
    fieldNameMap["reason"] = "Reason"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.namespaces.container",fields, reflect.TypeOf(Container{}), fieldNameMap, validators)
}

func ResourceQuotaOptionsV1BindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["memory_limit"] = bindings.NewOptionalType(bindings.NewIntegerType())
    fieldNameMap["memory_limit"] = "MemoryLimit"
    fields["memory_limit_default"] = bindings.NewOptionalType(bindings.NewIntegerType())
    fieldNameMap["memory_limit_default"] = "MemoryLimitDefault"
    fields["memory_request_default"] = bindings.NewOptionalType(bindings.NewIntegerType())
    fieldNameMap["memory_request_default"] = "MemoryRequestDefault"
    fields["cpu_limit"] = bindings.NewOptionalType(bindings.NewIntegerType())
    fieldNameMap["cpu_limit"] = "CpuLimit"
    fields["cpu_limit_default"] = bindings.NewOptionalType(bindings.NewIntegerType())
    fieldNameMap["cpu_limit_default"] = "CpuLimitDefault"
    fields["cpu_request_default"] = bindings.NewOptionalType(bindings.NewIntegerType())
    fieldNameMap["cpu_request_default"] = "CpuRequestDefault"
    fields["storage_request_limit"] = bindings.NewOptionalType(bindings.NewIntegerType())
    fieldNameMap["storage_request_limit"] = "StorageRequestLimit"
    fields["pod_count"] = bindings.NewOptionalType(bindings.NewIntegerType())
    fieldNameMap["pod_count"] = "PodCount"
    fields["service_count"] = bindings.NewOptionalType(bindings.NewIntegerType())
    fieldNameMap["service_count"] = "ServiceCount"
    fields["deployment_count"] = bindings.NewOptionalType(bindings.NewIntegerType())
    fieldNameMap["deployment_count"] = "DeploymentCount"
    fields["daemon_set_count"] = bindings.NewOptionalType(bindings.NewIntegerType())
    fieldNameMap["daemon_set_count"] = "DaemonSetCount"
    fields["replica_set_count"] = bindings.NewOptionalType(bindings.NewIntegerType())
    fieldNameMap["replica_set_count"] = "ReplicaSetCount"
    fields["replication_controller_count"] = bindings.NewOptionalType(bindings.NewIntegerType())
    fieldNameMap["replication_controller_count"] = "ReplicationControllerCount"
    fields["stateful_set_count"] = bindings.NewOptionalType(bindings.NewIntegerType())
    fieldNameMap["stateful_set_count"] = "StatefulSetCount"
    fields["config_map_count"] = bindings.NewOptionalType(bindings.NewIntegerType())
    fieldNameMap["config_map_count"] = "ConfigMapCount"
    fields["secret_count"] = bindings.NewOptionalType(bindings.NewIntegerType())
    fieldNameMap["secret_count"] = "SecretCount"
    fields["persistent_volume_claim_count"] = bindings.NewOptionalType(bindings.NewIntegerType())
    fieldNameMap["persistent_volume_claim_count"] = "PersistentVolumeClaimCount"
    fields["job_count"] = bindings.NewOptionalType(bindings.NewIntegerType())
    fieldNameMap["job_count"] = "JobCount"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.namespaces.resource_quota_options_v1",fields, reflect.TypeOf(ResourceQuotaOptionsV1{}), fieldNameMap, validators)
}

func ResourceQuotaOptionsV1UpdateBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["memory_limit"] = bindings.NewOptionalType(bindings.NewIntegerType())
    fieldNameMap["memory_limit"] = "MemoryLimit"
    fields["memory_limit_unset"] = bindings.NewOptionalType(bindings.NewBooleanType())
    fieldNameMap["memory_limit_unset"] = "MemoryLimitUnset"
    fields["memory_limit_default"] = bindings.NewOptionalType(bindings.NewIntegerType())
    fieldNameMap["memory_limit_default"] = "MemoryLimitDefault"
    fields["memory_limit_default_unset"] = bindings.NewOptionalType(bindings.NewBooleanType())
    fieldNameMap["memory_limit_default_unset"] = "MemoryLimitDefaultUnset"
    fields["memory_request_default"] = bindings.NewOptionalType(bindings.NewIntegerType())
    fieldNameMap["memory_request_default"] = "MemoryRequestDefault"
    fields["memory_request_default_unset"] = bindings.NewOptionalType(bindings.NewBooleanType())
    fieldNameMap["memory_request_default_unset"] = "MemoryRequestDefaultUnset"
    fields["cpu_limit"] = bindings.NewOptionalType(bindings.NewIntegerType())
    fieldNameMap["cpu_limit"] = "CpuLimit"
    fields["cpu_limit_unset"] = bindings.NewOptionalType(bindings.NewBooleanType())
    fieldNameMap["cpu_limit_unset"] = "CpuLimitUnset"
    fields["cpu_limit_default"] = bindings.NewOptionalType(bindings.NewIntegerType())
    fieldNameMap["cpu_limit_default"] = "CpuLimitDefault"
    fields["cpu_limit_default_unset"] = bindings.NewOptionalType(bindings.NewBooleanType())
    fieldNameMap["cpu_limit_default_unset"] = "CpuLimitDefaultUnset"
    fields["cpu_request_default"] = bindings.NewOptionalType(bindings.NewIntegerType())
    fieldNameMap["cpu_request_default"] = "CpuRequestDefault"
    fields["cpu_request_default_unset"] = bindings.NewOptionalType(bindings.NewBooleanType())
    fieldNameMap["cpu_request_default_unset"] = "CpuRequestDefaultUnset"
    fields["storage_request_limit"] = bindings.NewOptionalType(bindings.NewIntegerType())
    fieldNameMap["storage_request_limit"] = "StorageRequestLimit"
    fields["storage_request_limit_unset"] = bindings.NewOptionalType(bindings.NewBooleanType())
    fieldNameMap["storage_request_limit_unset"] = "StorageRequestLimitUnset"
    fields["pod_count"] = bindings.NewOptionalType(bindings.NewIntegerType())
    fieldNameMap["pod_count"] = "PodCount"
    fields["pod_count_unset"] = bindings.NewOptionalType(bindings.NewBooleanType())
    fieldNameMap["pod_count_unset"] = "PodCountUnset"
    fields["service_count"] = bindings.NewOptionalType(bindings.NewIntegerType())
    fieldNameMap["service_count"] = "ServiceCount"
    fields["service_count_unset"] = bindings.NewOptionalType(bindings.NewBooleanType())
    fieldNameMap["service_count_unset"] = "ServiceCountUnset"
    fields["deployment_count"] = bindings.NewOptionalType(bindings.NewIntegerType())
    fieldNameMap["deployment_count"] = "DeploymentCount"
    fields["deployment_count_unset"] = bindings.NewOptionalType(bindings.NewBooleanType())
    fieldNameMap["deployment_count_unset"] = "DeploymentCountUnset"
    fields["daemon_set_count"] = bindings.NewOptionalType(bindings.NewIntegerType())
    fieldNameMap["daemon_set_count"] = "DaemonSetCount"
    fields["daemon_set_count_unset"] = bindings.NewOptionalType(bindings.NewBooleanType())
    fieldNameMap["daemon_set_count_unset"] = "DaemonSetCountUnset"
    fields["replica_set_count"] = bindings.NewOptionalType(bindings.NewIntegerType())
    fieldNameMap["replica_set_count"] = "ReplicaSetCount"
    fields["replica_set_count_unset"] = bindings.NewOptionalType(bindings.NewBooleanType())
    fieldNameMap["replica_set_count_unset"] = "ReplicaSetCountUnset"
    fields["replication_controller_count"] = bindings.NewOptionalType(bindings.NewIntegerType())
    fieldNameMap["replication_controller_count"] = "ReplicationControllerCount"
    fields["replication_controller_count_unset"] = bindings.NewOptionalType(bindings.NewBooleanType())
    fieldNameMap["replication_controller_count_unset"] = "ReplicationControllerCountUnset"
    fields["stateful_set_count"] = bindings.NewOptionalType(bindings.NewIntegerType())
    fieldNameMap["stateful_set_count"] = "StatefulSetCount"
    fields["stateful_set_count_unset"] = bindings.NewOptionalType(bindings.NewBooleanType())
    fieldNameMap["stateful_set_count_unset"] = "StatefulSetCountUnset"
    fields["config_map_count"] = bindings.NewOptionalType(bindings.NewIntegerType())
    fieldNameMap["config_map_count"] = "ConfigMapCount"
    fields["config_map_count_unset"] = bindings.NewOptionalType(bindings.NewBooleanType())
    fieldNameMap["config_map_count_unset"] = "ConfigMapCountUnset"
    fields["secret_count"] = bindings.NewOptionalType(bindings.NewIntegerType())
    fieldNameMap["secret_count"] = "SecretCount"
    fields["secret_count_unset"] = bindings.NewOptionalType(bindings.NewBooleanType())
    fieldNameMap["secret_count_unset"] = "SecretCountUnset"
    fields["persistent_volume_claim_count"] = bindings.NewOptionalType(bindings.NewIntegerType())
    fieldNameMap["persistent_volume_claim_count"] = "PersistentVolumeClaimCount"
    fields["persistent_volume_claim_count_unset"] = bindings.NewOptionalType(bindings.NewBooleanType())
    fieldNameMap["persistent_volume_claim_count_unset"] = "PersistentVolumeClaimCountUnset"
    fields["job_count"] = bindings.NewOptionalType(bindings.NewIntegerType())
    fieldNameMap["job_count"] = "JobCount"
    fields["job_count_unset"] = bindings.NewOptionalType(bindings.NewBooleanType())
    fieldNameMap["job_count_unset"] = "JobCountUnset"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.namespaces.resource_quota_options_v1_update",fields, reflect.TypeOf(ResourceQuotaOptionsV1Update{}), fieldNameMap, validators)
}


