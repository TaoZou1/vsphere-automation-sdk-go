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

package namespace_management

import (
    "reflect"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/bindings/com/vmware/vapi/std"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/bindings"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/data"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/protocol"
)



// The ``ConfigStatus`` enumeration class describes the status of reaching the desired state configuration for the cluster. **Warning:** This enumeration is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
//
// <p> See {@link com.vmware.vapi.bindings.ApiEnumeration enumerated types description}.
 
type Clusters_ConfigStatus string

const (
    // The Namespace configuration is being applied to the cluster. **Warning:** This constant field is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
     Clusters_ConfigStatus_CONFIGURING Clusters_ConfigStatus = "CONFIGURING"
    // The Namespace configuration is being removed from the cluster. **Warning:** This constant field is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
     Clusters_ConfigStatus_REMOVING Clusters_ConfigStatus = "REMOVING"
    // The cluster is configured correctly with the Namespace configuration. **Warning:** This constant field is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
     Clusters_ConfigStatus_RUNNING Clusters_ConfigStatus = "RUNNING"
    // Failed to apply the Namespace configuration to the cluster, user intervention needed. **Warning:** This constant field is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
     Clusters_ConfigStatus_ERROR Clusters_ConfigStatus = "ERROR"
)

func (c Clusters_ConfigStatus) Clusters_ConfigStatus() bool {
    switch c {
        case Clusters_ConfigStatus_CONFIGURING:
            return true
        case Clusters_ConfigStatus_REMOVING:
            return true
        case Clusters_ConfigStatus_RUNNING:
            return true
        case Clusters_ConfigStatus_ERROR:
            return true
        default:
            return false
    }
}




// The ``KubernetesStatus`` enumeration class describes the cluster's ability to deploy pods. **Warning:** This enumeration is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
//
// <p> See {@link com.vmware.vapi.bindings.ApiEnumeration enumerated types description}.
 
type Clusters_KubernetesStatus string

const (
    // The cluster is able to accept pods. **Warning:** This constant field is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
     Clusters_KubernetesStatus_READY Clusters_KubernetesStatus = "READY"
    // The cluster may be able to accept pods, but has warning messages. **Warning:** This constant field is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
     Clusters_KubernetesStatus_WARNING Clusters_KubernetesStatus = "WARNING"
    // The cluster may not be able to accept pods and has error messages. **Warning:** This constant field is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
     Clusters_KubernetesStatus_ERROR Clusters_KubernetesStatus = "ERROR"
)

func (k Clusters_KubernetesStatus) Clusters_KubernetesStatus() bool {
    switch k {
        case Clusters_KubernetesStatus_READY:
            return true
        case Clusters_KubernetesStatus_WARNING:
            return true
        case Clusters_KubernetesStatus_ERROR:
            return true
        default:
            return false
    }
}




// Identifies the network plugin that cluster networking functionalities for this vSphere Namespaces Cluster. **Warning:** This enumeration is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
//
// <p> See {@link com.vmware.vapi.bindings.ApiEnumeration enumerated types description}.
 
type Clusters_NetworkProvider string

const (
    // NSX-T Container Plugin. **Warning:** This constant field is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
     Clusters_NetworkProvider_NSXT_CONTAINER_PLUGIN Clusters_NetworkProvider = "NSXT_CONTAINER_PLUGIN"
)

func (n Clusters_NetworkProvider) Clusters_NetworkProvider() bool {
    switch n {
        case Clusters_NetworkProvider_NSXT_CONTAINER_PLUGIN:
            return true
        default:
            return false
    }
}





// The ``Message`` class contains the information about the object configuration. **Warning:** This class is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
 type ClustersMessage struct {
    // Type of the message. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    Severity ClustersMessage_Severity
    // Details about the message. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    Details *std.LocalizableMessage
}



func (ClustersMessage ClustersMessage) Error() string {
    return "com.vmware.vcenter.namespace_management.message"
}

    
    // The ``Severity`` enumeration class represents the severity of the message. **Warning:** This enumeration is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    //
    // <p> See {@link com.vmware.vapi.bindings.ApiEnumeration enumerated types description}.
     
    type ClustersMessage_Severity string

    const (
        // Informational message. This may be accompanied by vCenter event. **Warning:** This constant field is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
         ClustersMessage_Severity_INFO ClustersMessage_Severity = "INFO"
        // Warning message. This may be accompanied by vCenter event. **Warning:** This constant field is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
         ClustersMessage_Severity_WARNING ClustersMessage_Severity = "WARNING"
        // Error message. This is accompanied by vCenter event and/or alarm. **Warning:** This constant field is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
         ClustersMessage_Severity_ERROR ClustersMessage_Severity = "ERROR"
    )

    func (s ClustersMessage_Severity) ClustersMessage_Severity() bool {
        switch s {
            case ClustersMessage_Severity_INFO:
                return true
            case ClustersMessage_Severity_WARNING:
                return true
            case ClustersMessage_Severity_ERROR:
                return true
            default:
                return false
        }
    }



// The ``Stats`` class contains the basic runtime statistics about a vSphere Namespaces-enabled cluster. **Warning:** This class is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
 type ClustersStats struct {
    // Overall CPU usage of the cluster, in MHz. This is the sum of CPU usage across all worker nodes in the cluster. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    CpuUsed int64
    // Total CPU capacity in the cluster available for vSphere Namespaces, in MHz. This is the sum of CPU capacities from all worker nodes in the cluster. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    CpuCapacity int64
    // Overall memory usage of the cluster, in mebibytes. This is the sum of memory usage across all worker nodes in the cluster. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    MemoryUsed int64
    // Total memory capacity of the cluster available for vSphere Namespaces, in mebibytes. This is the sum of memory capacities from all worker nodesin the cluster. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    MemoryCapacity int64
    // Overall storage used by the cluster, in mebibytes. This is the sum of storage used across all worker nodes in the cluster. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    StorageUsed int64
    // Overall storage capacity of the cluster available for vSphere Namespaces, in mebibytes. This is the sum of total storage available from all worker nodes in the cluster. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    StorageCapacity int64
}



func (ClustersStats ClustersStats) Error() string {
    return "com.vmware.vcenter.namespace_management.stats"
}



// The ``Summary`` class contains the basic information about the cluster statistics and status related to vSphere Namespaces. **Warning:** This class is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
 type ClustersSummary struct {
    // Identifier for the cluster. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    Cluster string
    // Name of the cluster. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    ClusterName string
    // Basic runtime statistics for the cluster. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    Stats ClustersStats
    // Current setting for ``ConfigStatus``. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    ConfigStatus Clusters_ConfigStatus
    // Current setting for ``KubernetesStatus``. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    KubernetesStatus Clusters_KubernetesStatus
}



func (ClustersSummary ClustersSummary) Error() string {
    return "com.vmware.vcenter.namespace_management.summary"
}



// The ``Info`` class contains detailed information about the cluster statistics and status related to vSphere Namespaces. **Warning:** This class is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
 type ClustersInfo struct {
    // Basic runtime statistics for the cluster. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    StatSummary ClustersStats
    // Current setting for ``ConfigStatus``. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    ConfigStatus Clusters_ConfigStatus
    // Current set of messages associated with the object. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    Messages []ClustersMessage
    // Current setting for ``KubernetesStatus``. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    KubernetesStatus Clusters_KubernetesStatus
    // Current set of messages associated with the object. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    KubernetesStatusMessages []ClustersMessage
    // Kubernetes API Server IP address on the management network. This is a floating IP and assigned to one of the control plane VMs on the management network. This endpoint is used by vSphere components. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    ApiServerManagementEndpoint string
    // Kubernetes API Server IP address via cluster network. This is the IP address of the Kubernetes LoadBalancer type service fronting the apiservers. This endpoint is the one configured in kubeconfig after login, and used for most human and application interaction with Kubernetes. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    ApiServerClusterEndpoint string
    // Identifier of the Kubernetes API servers. These are the IP addresses of the VM instances for the Kubernetes control plane on the management network. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    ApiServers map[string]bool
    // PEM-encoded x509 certificate used by TLS endpoint on Kubernetes API servers when accessed from the management network, e.g. from ESX servers or VCSA. This certificate is only valid for use with the apiServerManagementEndpoint. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    TlsManagementEndpointCertificate *string
    // PEM-encoded x509 certificate used by TLS endpoint on Kubernetes API servers when accessed via the load balancer, e.g. devops user on corporate network. This certificate is only valid for use with the apiServerClusterEndpoint. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    TlsEndpointCertificate *string
    // The provider of cluster networking for this vSphere Namespaces cluster. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    NetworkProvider Clusters_NetworkProvider
    // Specification for the NSX Container Plugin cluster network. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    NcpClusterNetworkInfo *ClustersNCPClusterNetworkInfo
    // CIDR block from which Kubernetes allocates service cluster IP addresses. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    ServiceCidr Ipv4Cidr
    // List of DNS server IP addresses to use on Kubernetes API server, specified in order of preference. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    MasterDNS []string
    // List of DNS server IP addresses to use for pods that execute on the worker nodes (which are native pods on ESXi hosts in the vSphere Namespaces Supervisor). **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    WorkerDNS []string
    // List of domains (for example "vmware.com") to be searched when trying to lookup a host name on Kubernetes API server, specified in order of preference. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    MasterDNSSearchDomains []string
}



func (ClustersInfo ClustersInfo) Error() string {
    return "com.vmware.vcenter.namespace_management.info"
}



// The ``Ipv4Range`` contains specification to configure multiple interfaces in IPv4. The range of IPv4 addresses is derived by incrementing the startingAddress to the specified addressCount. To use the object for a single IPv4 address specification, set addressCount to 1. **Warning:** This class is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
 type ClustersIpv4Range struct {
    // The IPv4 address denoting the start of the range. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    StartingAddress string
    // The number of IP addresses in the range. Addresses are derived by incrementing ClustersIpv4Range#startingAddress. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    AddressCount int64
    // Subnet mask to be set. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    SubnetMask string
    // The IPv4 address of the gateway associated with the range indicated by ClustersIpv4Range#startingAddress and ClustersIpv4Range#addressCount. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    Gateway string
}



func (ClustersIpv4Range ClustersIpv4Range) Error() string {
    return "com.vmware.vcenter.namespace_management.ipv4_range"
}



// The ``NetworkSpec`` contains information related to network configuration for one or more interfaces. **Warning:** This class is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
 type ClustersNetworkSpec struct {
    // Optionally specify the Floating IP used by the HA master cluster in the DHCP case. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    FloatingIP *string
    // Identifier for the network. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    Network string
    // The address assignment mode. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    Mode ClustersNetworkSpec_Ipv4Mode
    // Settings for the interfaces on the network. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    AddressRange *ClustersIpv4Range
}



func (ClustersNetworkSpec ClustersNetworkSpec) Error() string {
    return "com.vmware.vcenter.namespace_management.network_spec"
}

    
    // The ``Ipv4Mode`` enumeration class defines various IPv4 address assignment modes. **Warning:** This enumeration is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    //
    // <p> See {@link com.vmware.vapi.bindings.ApiEnumeration enumerated types description}.
     
    type ClustersNetworkSpec_Ipv4Mode string

    const (
        // The address is automatically assigned by a DHCP server. **Warning:** This constant field is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
         ClustersNetworkSpec_Ipv4Mode_DHCP ClustersNetworkSpec_Ipv4Mode = "DHCP"
        // The address is static. **Warning:** This constant field is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
         ClustersNetworkSpec_Ipv4Mode_STATICRANGE ClustersNetworkSpec_Ipv4Mode = "STATICRANGE"
    )

    func (i ClustersNetworkSpec_Ipv4Mode) ClustersNetworkSpec_Ipv4Mode() bool {
        switch i {
            case ClustersNetworkSpec_Ipv4Mode_DHCP:
                return true
            case ClustersNetworkSpec_Ipv4Mode_STATICRANGE:
                return true
            default:
                return false
        }
    }



// The ``ImageRegistry`` class contains the specification required to configure container image registry endpoint. **Warning:** This class is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
 type ClustersImageRegistry struct {
    // IP address or the hostname of container image registry. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    Hostname string
    // Port number of the container image registry. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    Port *int64
}



func (ClustersImageRegistry ClustersImageRegistry) Error() string {
    return "com.vmware.vcenter.namespace_management.image_registry"
}



// The ``ImageStorageSpec`` class contains the specification required to configure storage used for container images. **Warning:** This class is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
 type ClustersImageStorageSpec struct {
    // Identifier of the storage policy. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    StoragePolicy string
}



func (ClustersImageStorageSpec ClustersImageStorageSpec) Error() string {
    return "com.vmware.vcenter.namespace_management.image_storage_spec"
}



// The ``NCPClusterNetworkInfo`` class contains the NSX Container Plugin-specific cluster networking configuration. **Warning:** This class is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
 type ClustersNCPClusterNetworkInfo struct {
    // CIDR blocks from which Kubernetes allocates pod IP addresses. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    PodCidrs []Ipv4Cidr
    // CIDR blocks from which NSX assigns IP addresses for Kubernetes Ingresses and Kubernetes Services of type LoadBalancer. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    IngressCidrs []Ipv4Cidr
    // CIDR blocks from which NSX assigns IP addresses used for performing SNAT from container IPs to external IPs. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    EgressCidrs []Ipv4Cidr
    // vSphere Distributed Switch used to connect this cluster. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    ClusterDistributedSwitch string
    // NSX Edge Cluster to be used for Kubernetes Services of type LoadBalancer, Kubernetes Ingresses, and NSX SNAT. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    NsxEdgeCluster string
    // PEM-encoded x509 certificate used by NSX as a default fallback certificate for Kubernetes Ingress services. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    DefaultIngressTlsCertificate string
}



func (ClustersNCPClusterNetworkInfo ClustersNCPClusterNetworkInfo) Error() string {
    return "com.vmware.vcenter.namespace_management.NCP_cluster_network_info"
}



// The ``NCPClusterNetworkEnableSpec`` class encapsulates the NSX Container Plugin-specific cluster networking configuration parameters for the vSphere Namespaces Cluster Enable operation. **Warning:** This class is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
 type ClustersNCPClusterNetworkEnableSpec struct {
    // CIDR blocks from which Kubernetes allocates pod IP addresses. This range should not overlap with those in null, ClustersNCPClusterNetworkEnableSpec#ingressCidrs, ClustersNCPClusterNetworkEnableSpec#egressCidrs, or other services running in the datacenter. All Pod CIDR blocks must be of at least subnet size /23. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    PodCidrs []Ipv4Cidr
    // CIDR blocks from which NSX assigns IP addresses for Kubernetes Ingresses and Kubernetes Services of type LoadBalancer. These ranges should not overlap with those in ClustersNCPClusterNetworkEnableSpec#podCidrs, null, ClustersNCPClusterNetworkEnableSpec#egressCidrs, or other services running in the datacenter. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    IngressCidrs []Ipv4Cidr
    // CIDR blocks from which NSX assigns IP addresses used for performing SNAT from container IPs to external IPs. These ranges should not overlap with those in ClustersNCPClusterNetworkEnableSpec#podCidrs, null, ClustersNCPClusterNetworkEnableSpec#ingressCidrs, or other services running in the datacenter. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    EgressCidrs []Ipv4Cidr
    // vSphere Distributed Switch used to connect this cluster. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    ClusterDistributedSwitch string
    // NSX Edge Cluster to be used for Kubernetes Services of type LoadBalancer, Kubernetes Ingresses, and NSX SNAT. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    NsxEdgeCluster string
}



func (ClustersNCPClusterNetworkEnableSpec ClustersNCPClusterNetworkEnableSpec) Error() string {
    return "com.vmware.vcenter.namespace_management.NCP_cluster_network_enable_spec"
}



// The ``NCPClusterNetworkUpdateSpec`` class encapsulates the NSX Container Plugin-specific cluster networking configuration parameters for the vSphere Namespaces Cluster Update operation. **Warning:** This class is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
 type ClustersNCPClusterNetworkUpdateSpec struct {
    // CIDR blocks from which Kubernetes allocates pod IP addresses. This range should not overlap with those in null, ClustersNCPClusterNetworkUpdateSpec#ingressCidrs, ClustersNCPClusterNetworkUpdateSpec#egressCidrs, or other services running in the datacenter. An update operation only allows for addition of new CIDR blocks to the existing list. All Pod CIDR blocks must be of at least subnet size /23. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    PodCidrs []Ipv4Cidr
    // CIDR blocks from which NSX assigns IP addresses for Kubernetes Ingresses and Kubernetes Services of type LoadBalancer. These ranges should not overlap with those in ClustersNCPClusterNetworkUpdateSpec#podCidrs, null, ClustersNCPClusterNetworkUpdateSpec#egressCidrs, or other services running in the datacenter. An update operation only allows for addition of new CIDR blocks to the existing list. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    IngressCidrs []Ipv4Cidr
    // CIDR blocks from which NSX assigns IP addresses used for performing SNAT from container IPs to external IPs. These ranges should not overlap with those in ClustersNCPClusterNetworkUpdateSpec#podCidrs, null, ClustersNCPClusterNetworkUpdateSpec#ingressCidrs, or other services running in the datacenter. An update operation only allows for addition of new CIDR blocks to the existing list. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    EgressCidrs []Ipv4Cidr
    // PEM-encoded x509 certificate used by NSX as a default fallback certificate for Kubernetes Ingress services. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    DefaultIngressTlsCertificate *string
}



func (ClustersNCPClusterNetworkUpdateSpec ClustersNCPClusterNetworkUpdateSpec) Error() string {
    return "com.vmware.vcenter.namespace_management.NCP_cluster_network_update_spec"
}



// The ``NCPClusterNetworkSetSpec`` class encapsulates the NSX Container Plugin-specific cluster networking configuration parameters for the vSphere Namespaces Cluster Set operation. **Warning:** This class is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
 type ClustersNCPClusterNetworkSetSpec struct {
    // CIDR blocks from which Kubernetes allocates pod IP addresses. This range should not overlap with those in null, ClustersNCPClusterNetworkSetSpec#ingressCidrs, ClustersNCPClusterNetworkSetSpec#egressCidrs, or other services running in the datacenter. A set operation only allows for addition of new CIDR blocks to the existing list. All Pod CIDR blocks must be of at least subnet size /23. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    PodCidrs []Ipv4Cidr
    // CIDR blocks from which NSX assigns IP addresses for Kubernetes Ingresses and Kubernetes Services of type LoadBalancer. These ranges should not overlap with those in ClustersNCPClusterNetworkSetSpec#podCidrs, null, ClustersNCPClusterNetworkSetSpec#egressCidrs, or other services running in the datacenter. A set operation only allows for addition of new CIDR blocks to the existing list. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    IngressCidrs []Ipv4Cidr
    // CIDR blocks from which NSX assigns IP addresses used for performing SNAT from container IPs to external IPs. These ranges should not overlap with those in ClustersNCPClusterNetworkSetSpec#podCidrs, null, ClustersNCPClusterNetworkSetSpec#ingressCidrs, or other services running in the datacenter. A set operation only allows for addition of new CIDR blocks to the existing list. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    EgressCidrs []Ipv4Cidr
    // PEM-encoded x509 certificate used by NSX as a default fallback certificate for Kubernetes Ingress services. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    DefaultIngressTlsCertificate string
}



func (ClustersNCPClusterNetworkSetSpec ClustersNCPClusterNetworkSetSpec) Error() string {
    return "com.vmware.vcenter.namespace_management.NCP_cluster_network_set_spec"
}



// The ``EnableSpec`` class contains the specification required to enable vSphere Namespaces on a cluster. **Warning:** This class is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
 type ClustersEnableSpec struct {
    // This affects the size and resources allocated to the Kubernetes API server and the worker nodes. It also affects the suggested default serviceCidr and podCidrs. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    SizeHint SizingHint
    // CIDR block from which Kubernetes allocates service cluster IP addresses. This range should not overlap with those in null, null, null, or other services running in the datacenter. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    ServiceCidr Ipv4Cidr
    // The provider of cluster networking for this vSphere Namespaces cluster. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    NetworkProvider Clusters_NetworkProvider
    // Specification for the NSX Container Plugin cluster network. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    NcpClusterNetworkSpec *ClustersNCPClusterNetworkEnableSpec
    // Specification for the management network on Kubernetes API server. ClustersNetworkSpec#mode must be STATICRANGE as we require Kubernetes API server to have a stable address. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    MasterManagementNetwork ClustersNetworkSpec
    // TODO (VKAL-2577): Remove. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    MasterClusterNetwork *ClustersNetworkSpec
    // TODO (VKAL-2577): Remove. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    WorkerClusterNetwork *ClustersNetworkSpec
    // TODO (VKAL-2577): Remove. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    MasterAdditionalNetworks []ClustersNetworkSpec
    // TODO (VKAL-2577): Remove. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    WorkerAdditionalNetworks []ClustersNetworkSpec
    // List of DNS server IP addresses to use on Kubernetes API server, specified in order of preference. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    MasterDNS []string
    // List of DNS server IP addresses to use on the worker nodes, specified in order of preference. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    WorkerDNS []string
    // List of domains (for example "vmware.com") to be searched when trying to lookup a host name on Kubernetes API server, specified in order of preference. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    MasterDNSSearchDomains []string
    // List of domains (for example "vmware.com") to be searched when trying to lookup a host name on worker nodes, specified in order of preference. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    WorkerDNSSearchDomains []string
    // List of NTP server DNS names or IP addresses to use on Kubernetes API server, specified in order of preference. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    MasterNTPServers []string
    // Identifier of storage policy associated with Kubernetes API server. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    MasterStoragePolicy string
    // Identifier of storage policy associated with ephemeral disks of all the Kubernetes Pods in the cluster. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    EphemeralStoragePolicy string
    // Disclaimer to be displayed prior to login via the Kubectl plugin. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    LoginBanner *string
    // List of additional DNS names to associate with the Kubernetes API server. These DNS names are embedded in the TLS certificate presented by the API server. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    MasterDNSNames []string
    // Specification for storage to be used for container images. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    ImageStorage ClustersImageStorageSpec
    // Default image registry to use when Kubernetes Pod container specification does not specify it as part of the container image name. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    DefaultImageRegistry *ClustersImageRegistry
    // Default image repository to use when Kubernetes Pod container specification does not specify it as part of the container image name. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    DefaultImageRepository *string
}



func (ClustersEnableSpec ClustersEnableSpec) Error() string {
    return "com.vmware.vcenter.namespace_management.enable_spec"
}



// The ``UpdateSpec`` class contains the specification required to update the configuration on the Cluster. This class is applied partially, and only the specified fields will replace or modify their existing counterparts. **Warning:** This class is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
 type ClustersUpdateSpec struct {
    // This affects the size and resources allocated to the Kubernetes API server. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    SizeHint *SizingHint
    // The provider of cluster networking for this vSphere Namespaces cluster. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    NetworkProvider *Clusters_NetworkProvider
    // Updated specification for the cluster network configuration. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    NcpClusterNetworkSpec *ClustersNCPClusterNetworkUpdateSpec
    // List of DNS server IP addresses to use on Kubernetes API server, specified in order of preference. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    MasterDNS []string
    // List of DNS server IP addresses to use on the worker nodes, specified in order of preference. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    WorkerDNS []string
    // List of domains (for example "vmware.com") to be searched when trying to lookup a host name on Kubernetes API server, specified in order of preference. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    MasterDNSSearchDomains []string
    // List of domains (for example "vmware.com") to be searched when trying to lookup a host name on worker nodes, specified in order of preference. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    WorkerDNSSearchDomains []string
    // List of NTP server DNS names or IP addresses to use on Kubernetes API server, specified in order of preference. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    MasterNTPServers []string
    // Identifier of storage policy associated with Kubernetes API server. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    MasterStoragePolicy *string
    // Identifier of storage policy associated with ephemeral disks of all the Kubernetes Pods in the cluster. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    EphemeralStoragePolicy *string
    // Disclaimer to be displayed prior to login via the Kubectl plugin. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    LoginBanner *string
    // Specification for storage to be used for container images. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    ImageStorage *ClustersImageStorageSpec
    // Default image registry to use when Kubernetes Pod container specification does not specify it as part of the container image name. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    DefaultImageRegistry *ClustersImageRegistry
    // Default image repository to use when Kubernetes Pod container specification does not specify it as part of the container image name. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    DefaultImageRepository *string
    // Certificate issued for Kubernetes API Server. Certificate used must be created by signing the Certificate Signing Request obtained from null Because a ``CertificateSigningRequest`` is created on an existing Namespaces-enabled ``Cluster``, you must use the ``UpdateSpec`` to specify this ``tlsEndpointCertificate`` on an existing ``Cluster`` rather than during initially enabling Namespaces on a ``Cluster``. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    TlsEndpointCertificate *string
}



func (ClustersUpdateSpec ClustersUpdateSpec) Error() string {
    return "com.vmware.vcenter.namespace_management.update_spec"
}



// The ``SetSpec`` class contains the specification required to set a new configuration on the Cluster. This class is applied in entirety, replacing the current specification fully. **Warning:** This class is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
 type ClustersSetSpec struct {
    // This affects the size and resources allocated to the Kubernetes API server. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    SizeHint SizingHint
    // The provider of cluster networking for this vSphere Namespaces cluster. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    NetworkProvider Clusters_NetworkProvider
    // Specification for the NSX Container Plugin cluster network. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    NcpClusterNetworkSpec *ClustersNCPClusterNetworkSetSpec
    // List of DNS server IP addresses to use on Kubernetes API server, specified in order of preference. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    MasterDNS []string
    // List of DNS server IP addresses to use on the worker nodes, specified in order of preference. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    WorkerDNS []string
    // List of domains (for example "vmware.com") to be searched when trying to lookup a host name on Kubernetes API server, specified in order of preference. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    MasterDNSSearchDomains []string
    // List of domains (for example "vmware.com") to be searched when trying to lookup a host name on worker nodes, specified in order of preference. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    WorkerDNSSearchDomains []string
    // List of NTP server DNS names or IP addresses to use on Kubernetes API server, specified in order of preference. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    MasterNTPServers []string
    // Identifier of storage policy associated with Kubernetes API server. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    MasterStoragePolicy string
    // Identifier of storage policy associated with ephemeral disks of all the Kubernetes Pods in the cluster. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    EphemeralStoragePolicy string
    // Disclaimer to be displayed prior to login via the Kubectl plugin. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    LoginBanner *string
    // Specification for storage to be used for container images. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    ImageStorage ClustersImageStorageSpec
    // Default image registry to use when Kubernetes Pod container specification does not specify it as part of the container image name. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    DefaultImageRegistry *ClustersImageRegistry
    // Default image repository to use when Kubernetes Pod container specification does not specify it as part of the container image name. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    DefaultImageRepository *string
}



func (ClustersSetSpec ClustersSetSpec) Error() string {
    return "com.vmware.vcenter.namespace_management.set_spec"
}






func clustersEnableInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["cluster"] = bindings.NewIdType([]string {"ClusterComputeResource"}, "")
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
    paramsTypeMap["cluster"] = bindings.NewIdType([]string {"ClusterComputeResource"}, "")
    paramsTypeMap["spec"] = bindings.NewReferenceType(ClustersEnableSpecBindingType)
    paramsTypeMap["cluster"] = bindings.NewIdType([]string {"ClusterComputeResource"}, "")
    pathParams["cluster"] = "cluster"
    resultHeaders := map[string]string{}
    errorHeaders := map[string]string{}
    errorHeaders["Unauthenticated.challenge"] = "WWW-Authenticate"
    return protocol.NewOperationRestMetadata(
      paramsTypeMap,
      pathParams,
      queryParams,
      headerParams,
      "spec",
      "POST",
      "/vcenter/namespace-management/clusters/{cluster}",
       resultHeaders,
       204,
       errorHeaders,
       map[string]int{"AlreadyExists": 400,"Error": 500,"NotFound": 404,"InvalidArgument": 400,"Unauthenticated": 401,"Unauthorized": 403})
}


func clustersDisableInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["cluster"] = bindings.NewIdType([]string {"ClusterComputeResource"}, "")
    fieldNameMap["cluster"] = "Cluster"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func clustersDisableOutputType() bindings.BindingType {
    return bindings.NewVoidType()
}

func clustersDisableRestMetadata() protocol.OperationRestMetadata {
    paramsTypeMap := map[string]bindings.BindingType{}
    pathParams := map[string]string{}
    queryParams := map[string]string{}
    headerParams := map[string]string{}
    paramsTypeMap["cluster"] = bindings.NewIdType([]string {"ClusterComputeResource"}, "")
    paramsTypeMap["cluster"] = bindings.NewIdType([]string {"ClusterComputeResource"}, "")
    pathParams["cluster"] = "cluster"
    resultHeaders := map[string]string{}
    errorHeaders := map[string]string{}
    errorHeaders["Unauthenticated.challenge"] = "WWW-Authenticate"
    return protocol.NewOperationRestMetadata(
      paramsTypeMap,
      pathParams,
      queryParams,
      headerParams,
      "",
      "POST",
      "/vcenter/namespace-management/clusters/{cluster}",
       resultHeaders,
       204,
       errorHeaders,
       map[string]int{"Error": 500,"NotFound": 404,"Unauthenticated": 401,"Unauthorized": 403})
}


func clustersGetInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["cluster"] = bindings.NewIdType([]string {"ClusterComputeResource"}, "")
    fieldNameMap["cluster"] = "Cluster"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func clustersGetOutputType() bindings.BindingType {
    return bindings.NewReferenceType(ClustersInfoBindingType)
}

func clustersGetRestMetadata() protocol.OperationRestMetadata {
    paramsTypeMap := map[string]bindings.BindingType{}
    pathParams := map[string]string{}
    queryParams := map[string]string{}
    headerParams := map[string]string{}
    paramsTypeMap["cluster"] = bindings.NewIdType([]string {"ClusterComputeResource"}, "")
    paramsTypeMap["cluster"] = bindings.NewIdType([]string {"ClusterComputeResource"}, "")
    pathParams["cluster"] = "cluster"
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
      "/vcenter/namespace-management/clusters/{cluster}",
       resultHeaders,
       200,
       errorHeaders,
       map[string]int{"NotFound": 404,"Unsupported": 400,"Error": 500,"Unauthenticated": 401,"Unauthorized": 403})
}


func clustersListInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func clustersListOutputType() bindings.BindingType {
    return bindings.NewListType(bindings.NewReferenceType(ClustersSummaryBindingType), reflect.TypeOf([]ClustersSummary{}))
}

func clustersListRestMetadata() protocol.OperationRestMetadata {
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
      "/vcenter/namespace-management/clusters",
       resultHeaders,
       200,
       errorHeaders,
       map[string]int{"Error": 500,"Unauthenticated": 401,"Unauthorized": 403})
}


func clustersSetInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["cluster"] = bindings.NewIdType([]string {"ClusterComputeResource"}, "")
    fields["spec"] = bindings.NewReferenceType(ClustersSetSpecBindingType)
    fieldNameMap["cluster"] = "Cluster"
    fieldNameMap["spec"] = "Spec"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func clustersSetOutputType() bindings.BindingType {
    return bindings.NewVoidType()
}

func clustersSetRestMetadata() protocol.OperationRestMetadata {
    paramsTypeMap := map[string]bindings.BindingType{}
    pathParams := map[string]string{}
    queryParams := map[string]string{}
    headerParams := map[string]string{}
    paramsTypeMap["cluster"] = bindings.NewIdType([]string {"ClusterComputeResource"}, "")
    paramsTypeMap["spec"] = bindings.NewReferenceType(ClustersSetSpecBindingType)
    paramsTypeMap["cluster"] = bindings.NewIdType([]string {"ClusterComputeResource"}, "")
    pathParams["cluster"] = "cluster"
    resultHeaders := map[string]string{}
    errorHeaders := map[string]string{}
    errorHeaders["Unauthenticated.challenge"] = "WWW-Authenticate"
    return protocol.NewOperationRestMetadata(
      paramsTypeMap,
      pathParams,
      queryParams,
      headerParams,
      "spec",
      "PUT",
      "/vcenter/namespace-management/clusters/{cluster}",
       resultHeaders,
       204,
       errorHeaders,
       map[string]int{"Error": 500,"InvalidArgument": 400,"NotAllowedInCurrentState": 400,"NotFound": 404,"Unauthenticated": 401,"Unauthorized": 403})
}


func clustersUpdateInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["cluster"] = bindings.NewIdType([]string {"ClusterComputeResource"}, "")
    fields["spec"] = bindings.NewReferenceType(ClustersUpdateSpecBindingType)
    fieldNameMap["cluster"] = "Cluster"
    fieldNameMap["spec"] = "Spec"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func clustersUpdateOutputType() bindings.BindingType {
    return bindings.NewVoidType()
}

func clustersUpdateRestMetadata() protocol.OperationRestMetadata {
    paramsTypeMap := map[string]bindings.BindingType{}
    pathParams := map[string]string{}
    queryParams := map[string]string{}
    headerParams := map[string]string{}
    paramsTypeMap["cluster"] = bindings.NewIdType([]string {"ClusterComputeResource"}, "")
    paramsTypeMap["spec"] = bindings.NewReferenceType(ClustersUpdateSpecBindingType)
    paramsTypeMap["cluster"] = bindings.NewIdType([]string {"ClusterComputeResource"}, "")
    pathParams["cluster"] = "cluster"
    resultHeaders := map[string]string{}
    errorHeaders := map[string]string{}
    errorHeaders["Unauthenticated.challenge"] = "WWW-Authenticate"
    return protocol.NewOperationRestMetadata(
      paramsTypeMap,
      pathParams,
      queryParams,
      headerParams,
      "spec",
      "PATCH",
      "/vcenter/namespace-management/clusters/{cluster}",
       resultHeaders,
       204,
       errorHeaders,
       map[string]int{"Error": 500,"InvalidArgument": 400,"NotAllowedInCurrentState": 400,"NotFound": 404,"Unauthenticated": 401,"Unauthorized": 403})
}



func ClustersMessageBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["severity"] = bindings.NewEnumType("com.vmware.vcenter.namespace_management.clusters.message.severity", reflect.TypeOf(ClustersMessage_Severity(ClustersMessage_Severity_INFO)))
    fieldNameMap["severity"] = "Severity"
    fields["details"] = bindings.NewOptionalType(bindings.NewReferenceType(std.LocalizableMessageBindingType))
    fieldNameMap["details"] = "Details"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.namespace_management.clusters.message",fields, reflect.TypeOf(ClustersMessage{}), fieldNameMap, validators)
}

func ClustersStatsBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["cpu_used"] = bindings.NewIntegerType()
    fieldNameMap["cpu_used"] = "CpuUsed"
    fields["cpu_capacity"] = bindings.NewIntegerType()
    fieldNameMap["cpu_capacity"] = "CpuCapacity"
    fields["memory_used"] = bindings.NewIntegerType()
    fieldNameMap["memory_used"] = "MemoryUsed"
    fields["memory_capacity"] = bindings.NewIntegerType()
    fieldNameMap["memory_capacity"] = "MemoryCapacity"
    fields["storage_used"] = bindings.NewIntegerType()
    fieldNameMap["storage_used"] = "StorageUsed"
    fields["storage_capacity"] = bindings.NewIntegerType()
    fieldNameMap["storage_capacity"] = "StorageCapacity"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.namespace_management.clusters.stats",fields, reflect.TypeOf(ClustersStats{}), fieldNameMap, validators)
}

func ClustersSummaryBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["cluster"] = bindings.NewIdType([]string {"ClusterComputeResource"}, "")
    fieldNameMap["cluster"] = "Cluster"
    fields["cluster_name"] = bindings.NewIdType([]string {"ClusterComputeResource.name"}, "")
    fieldNameMap["cluster_name"] = "ClusterName"
    fields["stats"] = bindings.NewReferenceType(ClustersStatsBindingType)
    fieldNameMap["stats"] = "Stats"
    fields["config_status"] = bindings.NewEnumType("com.vmware.vcenter.namespace_management.clusters.config_status", reflect.TypeOf(Clusters_ConfigStatus(Clusters_ConfigStatus_CONFIGURING)))
    fieldNameMap["config_status"] = "ConfigStatus"
    fields["kubernetes_status"] = bindings.NewEnumType("com.vmware.vcenter.namespace_management.clusters.kubernetes_status", reflect.TypeOf(Clusters_KubernetesStatus(Clusters_KubernetesStatus_READY)))
    fieldNameMap["kubernetes_status"] = "KubernetesStatus"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.namespace_management.clusters.summary",fields, reflect.TypeOf(ClustersSummary{}), fieldNameMap, validators)
}

func ClustersInfoBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["stat_summary"] = bindings.NewReferenceType(ClustersStatsBindingType)
    fieldNameMap["stat_summary"] = "StatSummary"
    fields["config_status"] = bindings.NewEnumType("com.vmware.vcenter.namespace_management.clusters.config_status", reflect.TypeOf(Clusters_ConfigStatus(Clusters_ConfigStatus_CONFIGURING)))
    fieldNameMap["config_status"] = "ConfigStatus"
    fields["messages"] = bindings.NewListType(bindings.NewReferenceType(ClustersMessageBindingType), reflect.TypeOf([]ClustersMessage{}))
    fieldNameMap["messages"] = "Messages"
    fields["kubernetes_status"] = bindings.NewEnumType("com.vmware.vcenter.namespace_management.clusters.kubernetes_status", reflect.TypeOf(Clusters_KubernetesStatus(Clusters_KubernetesStatus_READY)))
    fieldNameMap["kubernetes_status"] = "KubernetesStatus"
    fields["kubernetes_status_messages"] = bindings.NewListType(bindings.NewReferenceType(ClustersMessageBindingType), reflect.TypeOf([]ClustersMessage{}))
    fieldNameMap["kubernetes_status_messages"] = "KubernetesStatusMessages"
    fields["api_server_management_endpoint"] = bindings.NewStringType()
    fieldNameMap["api_server_management_endpoint"] = "ApiServerManagementEndpoint"
    fields["api_server_cluster_endpoint"] = bindings.NewStringType()
    fieldNameMap["api_server_cluster_endpoint"] = "ApiServerClusterEndpoint"
    fields["api_servers"] = bindings.NewSetType(bindings.NewStringType(), reflect.TypeOf(map[string]bool{}))
    fieldNameMap["api_servers"] = "ApiServers"
    fields["tls_management_endpoint_certificate"] = bindings.NewOptionalType(bindings.NewStringType())
    fieldNameMap["tls_management_endpoint_certificate"] = "TlsManagementEndpointCertificate"
    fields["tls_endpoint_certificate"] = bindings.NewOptionalType(bindings.NewStringType())
    fieldNameMap["tls_endpoint_certificate"] = "TlsEndpointCertificate"
    fields["network_provider"] = bindings.NewEnumType("com.vmware.vcenter.namespace_management.clusters.network_provider", reflect.TypeOf(Clusters_NetworkProvider(Clusters_NetworkProvider_NSXT_CONTAINER_PLUGIN)))
    fieldNameMap["network_provider"] = "NetworkProvider"
    fields["ncp_cluster_network_info"] = bindings.NewOptionalType(bindings.NewReferenceType(ClustersNCPClusterNetworkInfoBindingType))
    fieldNameMap["ncp_cluster_network_info"] = "NcpClusterNetworkInfo"
    fields["service_cidr"] = bindings.NewReferenceType(Ipv4CidrBindingType)
    fieldNameMap["service_cidr"] = "ServiceCidr"
    fields["master_DNS"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewStringType(), reflect.TypeOf([]string{})))
    fieldNameMap["master_DNS"] = "MasterDNS"
    fields["worker_DNS"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewStringType(), reflect.TypeOf([]string{})))
    fieldNameMap["worker_DNS"] = "WorkerDNS"
    fields["master_DNS_search_domains"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewStringType(), reflect.TypeOf([]string{})))
    fieldNameMap["master_DNS_search_domains"] = "MasterDNSSearchDomains"
    var validators = []bindings.Validator{}
    uv1 := bindings.NewUnionValidator("network_provider",
        map[string][]bindings.FieldData {
            "NSXT_CONTAINER_PLUGIN": []bindings.FieldData {
                 bindings.NewFieldData("ncp_cluster_network_info", true),
            },
        },
    )
    validators = append(validators, uv1)
    return bindings.NewStructType("com.vmware.vcenter.namespace_management.clusters.info",fields, reflect.TypeOf(ClustersInfo{}), fieldNameMap, validators)
}

func ClustersIpv4RangeBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["starting_address"] = bindings.NewStringType()
    fieldNameMap["starting_address"] = "StartingAddress"
    fields["address_count"] = bindings.NewIntegerType()
    fieldNameMap["address_count"] = "AddressCount"
    fields["subnet_mask"] = bindings.NewStringType()
    fieldNameMap["subnet_mask"] = "SubnetMask"
    fields["gateway"] = bindings.NewStringType()
    fieldNameMap["gateway"] = "Gateway"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.namespace_management.clusters.ipv4_range",fields, reflect.TypeOf(ClustersIpv4Range{}), fieldNameMap, validators)
}

func ClustersNetworkSpecBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["floating_IP"] = bindings.NewOptionalType(bindings.NewStringType())
    fieldNameMap["floating_IP"] = "FloatingIP"
    fields["network"] = bindings.NewIdType([]string {"Network"}, "")
    fieldNameMap["network"] = "Network"
    fields["mode"] = bindings.NewEnumType("com.vmware.vcenter.namespace_management.clusters.network_spec.ipv4_mode", reflect.TypeOf(ClustersNetworkSpec_Ipv4Mode(ClustersNetworkSpec_Ipv4Mode_DHCP)))
    fieldNameMap["mode"] = "Mode"
    fields["address_range"] = bindings.NewOptionalType(bindings.NewReferenceType(ClustersIpv4RangeBindingType))
    fieldNameMap["address_range"] = "AddressRange"
    var validators = []bindings.Validator{}
    uv1 := bindings.NewUnionValidator("mode",
        map[string][]bindings.FieldData {
            "DHCP": []bindings.FieldData {
                 bindings.NewFieldData("floating_IP", false),
            },
            "STATICRANGE": []bindings.FieldData {
                 bindings.NewFieldData("address_range", true),
            },
        },
    )
    validators = append(validators, uv1)
    return bindings.NewStructType("com.vmware.vcenter.namespace_management.clusters.network_spec",fields, reflect.TypeOf(ClustersNetworkSpec{}), fieldNameMap, validators)
}

func ClustersImageRegistryBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["hostname"] = bindings.NewStringType()
    fieldNameMap["hostname"] = "Hostname"
    fields["port"] = bindings.NewOptionalType(bindings.NewIntegerType())
    fieldNameMap["port"] = "Port"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.namespace_management.clusters.image_registry",fields, reflect.TypeOf(ClustersImageRegistry{}), fieldNameMap, validators)
}

func ClustersImageStorageSpecBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["storage_policy"] = bindings.NewIdType([]string {"SpsStorageProfile"}, "")
    fieldNameMap["storage_policy"] = "StoragePolicy"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.namespace_management.clusters.image_storage_spec",fields, reflect.TypeOf(ClustersImageStorageSpec{}), fieldNameMap, validators)
}

func ClustersNCPClusterNetworkInfoBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["pod_cidrs"] = bindings.NewListType(bindings.NewReferenceType(Ipv4CidrBindingType), reflect.TypeOf([]Ipv4Cidr{}))
    fieldNameMap["pod_cidrs"] = "PodCidrs"
    fields["ingress_cidrs"] = bindings.NewListType(bindings.NewReferenceType(Ipv4CidrBindingType), reflect.TypeOf([]Ipv4Cidr{}))
    fieldNameMap["ingress_cidrs"] = "IngressCidrs"
    fields["egress_cidrs"] = bindings.NewListType(bindings.NewReferenceType(Ipv4CidrBindingType), reflect.TypeOf([]Ipv4Cidr{}))
    fieldNameMap["egress_cidrs"] = "EgressCidrs"
    fields["cluster_distributed_switch"] = bindings.NewIdType([]string {"vSphereDistributedSwitch"}, "")
    fieldNameMap["cluster_distributed_switch"] = "ClusterDistributedSwitch"
    fields["nsx_edge_cluster"] = bindings.NewIdType([]string {"NSXEdgeCluster"}, "")
    fieldNameMap["nsx_edge_cluster"] = "NsxEdgeCluster"
    fields["default_ingress_tls_certificate"] = bindings.NewStringType()
    fieldNameMap["default_ingress_tls_certificate"] = "DefaultIngressTlsCertificate"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.namespace_management.clusters.NCP_cluster_network_info",fields, reflect.TypeOf(ClustersNCPClusterNetworkInfo{}), fieldNameMap, validators)
}

func ClustersNCPClusterNetworkEnableSpecBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["pod_cidrs"] = bindings.NewListType(bindings.NewReferenceType(Ipv4CidrBindingType), reflect.TypeOf([]Ipv4Cidr{}))
    fieldNameMap["pod_cidrs"] = "PodCidrs"
    fields["ingress_cidrs"] = bindings.NewListType(bindings.NewReferenceType(Ipv4CidrBindingType), reflect.TypeOf([]Ipv4Cidr{}))
    fieldNameMap["ingress_cidrs"] = "IngressCidrs"
    fields["egress_cidrs"] = bindings.NewListType(bindings.NewReferenceType(Ipv4CidrBindingType), reflect.TypeOf([]Ipv4Cidr{}))
    fieldNameMap["egress_cidrs"] = "EgressCidrs"
    fields["cluster_distributed_switch"] = bindings.NewIdType([]string {"vSphereDistributedSwitch"}, "")
    fieldNameMap["cluster_distributed_switch"] = "ClusterDistributedSwitch"
    fields["nsx_edge_cluster"] = bindings.NewIdType([]string {"NSXEdgeCluster"}, "")
    fieldNameMap["nsx_edge_cluster"] = "NsxEdgeCluster"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.namespace_management.clusters.NCP_cluster_network_enable_spec",fields, reflect.TypeOf(ClustersNCPClusterNetworkEnableSpec{}), fieldNameMap, validators)
}

func ClustersNCPClusterNetworkUpdateSpecBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["pod_cidrs"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewReferenceType(Ipv4CidrBindingType), reflect.TypeOf([]Ipv4Cidr{})))
    fieldNameMap["pod_cidrs"] = "PodCidrs"
    fields["ingress_cidrs"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewReferenceType(Ipv4CidrBindingType), reflect.TypeOf([]Ipv4Cidr{})))
    fieldNameMap["ingress_cidrs"] = "IngressCidrs"
    fields["egress_cidrs"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewReferenceType(Ipv4CidrBindingType), reflect.TypeOf([]Ipv4Cidr{})))
    fieldNameMap["egress_cidrs"] = "EgressCidrs"
    fields["default_ingress_tls_certificate"] = bindings.NewOptionalType(bindings.NewStringType())
    fieldNameMap["default_ingress_tls_certificate"] = "DefaultIngressTlsCertificate"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.namespace_management.clusters.NCP_cluster_network_update_spec",fields, reflect.TypeOf(ClustersNCPClusterNetworkUpdateSpec{}), fieldNameMap, validators)
}

func ClustersNCPClusterNetworkSetSpecBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["pod_cidrs"] = bindings.NewListType(bindings.NewReferenceType(Ipv4CidrBindingType), reflect.TypeOf([]Ipv4Cidr{}))
    fieldNameMap["pod_cidrs"] = "PodCidrs"
    fields["ingress_cidrs"] = bindings.NewListType(bindings.NewReferenceType(Ipv4CidrBindingType), reflect.TypeOf([]Ipv4Cidr{}))
    fieldNameMap["ingress_cidrs"] = "IngressCidrs"
    fields["egress_cidrs"] = bindings.NewListType(bindings.NewReferenceType(Ipv4CidrBindingType), reflect.TypeOf([]Ipv4Cidr{}))
    fieldNameMap["egress_cidrs"] = "EgressCidrs"
    fields["default_ingress_tls_certificate"] = bindings.NewStringType()
    fieldNameMap["default_ingress_tls_certificate"] = "DefaultIngressTlsCertificate"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.namespace_management.clusters.NCP_cluster_network_set_spec",fields, reflect.TypeOf(ClustersNCPClusterNetworkSetSpec{}), fieldNameMap, validators)
}

func ClustersEnableSpecBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["size_hint"] = bindings.NewEnumType("com.vmware.vcenter.namespace_management.sizing_hint", reflect.TypeOf(SizingHint(SizingHint_TINY)))
    fieldNameMap["size_hint"] = "SizeHint"
    fields["service_cidr"] = bindings.NewReferenceType(Ipv4CidrBindingType)
    fieldNameMap["service_cidr"] = "ServiceCidr"
    fields["network_provider"] = bindings.NewEnumType("com.vmware.vcenter.namespace_management.clusters.network_provider", reflect.TypeOf(Clusters_NetworkProvider(Clusters_NetworkProvider_NSXT_CONTAINER_PLUGIN)))
    fieldNameMap["network_provider"] = "NetworkProvider"
    fields["ncp_cluster_network_spec"] = bindings.NewOptionalType(bindings.NewReferenceType(ClustersNCPClusterNetworkEnableSpecBindingType))
    fieldNameMap["ncp_cluster_network_spec"] = "NcpClusterNetworkSpec"
    fields["master_management_network"] = bindings.NewReferenceType(ClustersNetworkSpecBindingType)
    fieldNameMap["master_management_network"] = "MasterManagementNetwork"
    fields["master_cluster_network"] = bindings.NewOptionalType(bindings.NewReferenceType(ClustersNetworkSpecBindingType))
    fieldNameMap["master_cluster_network"] = "MasterClusterNetwork"
    fields["worker_cluster_network"] = bindings.NewOptionalType(bindings.NewReferenceType(ClustersNetworkSpecBindingType))
    fieldNameMap["worker_cluster_network"] = "WorkerClusterNetwork"
    fields["master_additional_networks"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewReferenceType(ClustersNetworkSpecBindingType), reflect.TypeOf([]ClustersNetworkSpec{})))
    fieldNameMap["master_additional_networks"] = "MasterAdditionalNetworks"
    fields["worker_additional_networks"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewReferenceType(ClustersNetworkSpecBindingType), reflect.TypeOf([]ClustersNetworkSpec{})))
    fieldNameMap["worker_additional_networks"] = "WorkerAdditionalNetworks"
    fields["master_DNS"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewStringType(), reflect.TypeOf([]string{})))
    fieldNameMap["master_DNS"] = "MasterDNS"
    fields["worker_DNS"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewStringType(), reflect.TypeOf([]string{})))
    fieldNameMap["worker_DNS"] = "WorkerDNS"
    fields["master_DNS_search_domains"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewStringType(), reflect.TypeOf([]string{})))
    fieldNameMap["master_DNS_search_domains"] = "MasterDNSSearchDomains"
    fields["worker_DNS_search_domains"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewStringType(), reflect.TypeOf([]string{})))
    fieldNameMap["worker_DNS_search_domains"] = "WorkerDNSSearchDomains"
    fields["master_NTP_servers"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewStringType(), reflect.TypeOf([]string{})))
    fieldNameMap["master_NTP_servers"] = "MasterNTPServers"
    fields["master_storage_policy"] = bindings.NewIdType([]string {"SpsStorageProfile"}, "")
    fieldNameMap["master_storage_policy"] = "MasterStoragePolicy"
    fields["ephemeral_storage_policy"] = bindings.NewIdType([]string {"SpsStorageProfile"}, "")
    fieldNameMap["ephemeral_storage_policy"] = "EphemeralStoragePolicy"
    fields["login_banner"] = bindings.NewOptionalType(bindings.NewStringType())
    fieldNameMap["login_banner"] = "LoginBanner"
    fields["Master_DNS_names"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewStringType(), reflect.TypeOf([]string{})))
    fieldNameMap["Master_DNS_names"] = "MasterDNSNames"
    fields["image_storage"] = bindings.NewReferenceType(ClustersImageStorageSpecBindingType)
    fieldNameMap["image_storage"] = "ImageStorage"
    fields["default_image_registry"] = bindings.NewOptionalType(bindings.NewReferenceType(ClustersImageRegistryBindingType))
    fieldNameMap["default_image_registry"] = "DefaultImageRegistry"
    fields["default_image_repository"] = bindings.NewOptionalType(bindings.NewStringType())
    fieldNameMap["default_image_repository"] = "DefaultImageRepository"
    var validators = []bindings.Validator{}
    uv1 := bindings.NewUnionValidator("network_provider",
        map[string][]bindings.FieldData {
            "NSXT_CONTAINER_PLUGIN": []bindings.FieldData {
                 bindings.NewFieldData("ncp_cluster_network_spec", true),
            },
        },
    )
    validators = append(validators, uv1)
    return bindings.NewStructType("com.vmware.vcenter.namespace_management.clusters.enable_spec",fields, reflect.TypeOf(ClustersEnableSpec{}), fieldNameMap, validators)
}

func ClustersUpdateSpecBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["size_hint"] = bindings.NewOptionalType(bindings.NewEnumType("com.vmware.vcenter.namespace_management.sizing_hint", reflect.TypeOf(SizingHint(SizingHint_TINY))))
    fieldNameMap["size_hint"] = "SizeHint"
    fields["network_provider"] = bindings.NewOptionalType(bindings.NewEnumType("com.vmware.vcenter.namespace_management.clusters.network_provider", reflect.TypeOf(Clusters_NetworkProvider(Clusters_NetworkProvider_NSXT_CONTAINER_PLUGIN))))
    fieldNameMap["network_provider"] = "NetworkProvider"
    fields["ncp_cluster_network_spec"] = bindings.NewOptionalType(bindings.NewReferenceType(ClustersNCPClusterNetworkUpdateSpecBindingType))
    fieldNameMap["ncp_cluster_network_spec"] = "NcpClusterNetworkSpec"
    fields["master_DNS"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewStringType(), reflect.TypeOf([]string{})))
    fieldNameMap["master_DNS"] = "MasterDNS"
    fields["worker_DNS"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewStringType(), reflect.TypeOf([]string{})))
    fieldNameMap["worker_DNS"] = "WorkerDNS"
    fields["master_DNS_search_domains"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewStringType(), reflect.TypeOf([]string{})))
    fieldNameMap["master_DNS_search_domains"] = "MasterDNSSearchDomains"
    fields["worker_DNS_search_domains"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewStringType(), reflect.TypeOf([]string{})))
    fieldNameMap["worker_DNS_search_domains"] = "WorkerDNSSearchDomains"
    fields["master_NTP_servers"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewStringType(), reflect.TypeOf([]string{})))
    fieldNameMap["master_NTP_servers"] = "MasterNTPServers"
    fields["master_storage_policy"] = bindings.NewOptionalType(bindings.NewIdType([]string {"SpsStorageProfile"}, ""))
    fieldNameMap["master_storage_policy"] = "MasterStoragePolicy"
    fields["ephemeral_storage_policy"] = bindings.NewOptionalType(bindings.NewIdType([]string {"SpsStorageProfile"}, ""))
    fieldNameMap["ephemeral_storage_policy"] = "EphemeralStoragePolicy"
    fields["login_banner"] = bindings.NewOptionalType(bindings.NewStringType())
    fieldNameMap["login_banner"] = "LoginBanner"
    fields["image_storage"] = bindings.NewOptionalType(bindings.NewReferenceType(ClustersImageStorageSpecBindingType))
    fieldNameMap["image_storage"] = "ImageStorage"
    fields["default_image_registry"] = bindings.NewOptionalType(bindings.NewReferenceType(ClustersImageRegistryBindingType))
    fieldNameMap["default_image_registry"] = "DefaultImageRegistry"
    fields["default_image_repository"] = bindings.NewOptionalType(bindings.NewStringType())
    fieldNameMap["default_image_repository"] = "DefaultImageRepository"
    fields["tls_endpoint_certificate"] = bindings.NewOptionalType(bindings.NewStringType())
    fieldNameMap["tls_endpoint_certificate"] = "TlsEndpointCertificate"
    var validators = []bindings.Validator{}
    uv1 := bindings.NewUnionValidator("network_provider",
        map[string][]bindings.FieldData {
            "NSXT_CONTAINER_PLUGIN": []bindings.FieldData {
                 bindings.NewFieldData("ncp_cluster_network_spec", false),
            },
        },
    )
    validators = append(validators, uv1)
    return bindings.NewStructType("com.vmware.vcenter.namespace_management.clusters.update_spec",fields, reflect.TypeOf(ClustersUpdateSpec{}), fieldNameMap, validators)
}

func ClustersSetSpecBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["size_hint"] = bindings.NewEnumType("com.vmware.vcenter.namespace_management.sizing_hint", reflect.TypeOf(SizingHint(SizingHint_TINY)))
    fieldNameMap["size_hint"] = "SizeHint"
    fields["network_provider"] = bindings.NewEnumType("com.vmware.vcenter.namespace_management.clusters.network_provider", reflect.TypeOf(Clusters_NetworkProvider(Clusters_NetworkProvider_NSXT_CONTAINER_PLUGIN)))
    fieldNameMap["network_provider"] = "NetworkProvider"
    fields["ncp_cluster_network_spec"] = bindings.NewOptionalType(bindings.NewReferenceType(ClustersNCPClusterNetworkSetSpecBindingType))
    fieldNameMap["ncp_cluster_network_spec"] = "NcpClusterNetworkSpec"
    fields["master_DNS"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewStringType(), reflect.TypeOf([]string{})))
    fieldNameMap["master_DNS"] = "MasterDNS"
    fields["worker_DNS"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewStringType(), reflect.TypeOf([]string{})))
    fieldNameMap["worker_DNS"] = "WorkerDNS"
    fields["master_DNS_search_domains"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewStringType(), reflect.TypeOf([]string{})))
    fieldNameMap["master_DNS_search_domains"] = "MasterDNSSearchDomains"
    fields["worker_DNS_search_domains"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewStringType(), reflect.TypeOf([]string{})))
    fieldNameMap["worker_DNS_search_domains"] = "WorkerDNSSearchDomains"
    fields["master_NTP_servers"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewStringType(), reflect.TypeOf([]string{})))
    fieldNameMap["master_NTP_servers"] = "MasterNTPServers"
    fields["master_storage_policy"] = bindings.NewIdType([]string {"SpsStorageProfile"}, "")
    fieldNameMap["master_storage_policy"] = "MasterStoragePolicy"
    fields["ephemeral_storage_policy"] = bindings.NewIdType([]string {"SpsStorageProfile"}, "")
    fieldNameMap["ephemeral_storage_policy"] = "EphemeralStoragePolicy"
    fields["login_banner"] = bindings.NewOptionalType(bindings.NewStringType())
    fieldNameMap["login_banner"] = "LoginBanner"
    fields["image_storage"] = bindings.NewReferenceType(ClustersImageStorageSpecBindingType)
    fieldNameMap["image_storage"] = "ImageStorage"
    fields["default_image_registry"] = bindings.NewOptionalType(bindings.NewReferenceType(ClustersImageRegistryBindingType))
    fieldNameMap["default_image_registry"] = "DefaultImageRegistry"
    fields["default_image_repository"] = bindings.NewOptionalType(bindings.NewStringType())
    fieldNameMap["default_image_repository"] = "DefaultImageRepository"
    var validators = []bindings.Validator{}
    uv1 := bindings.NewUnionValidator("network_provider",
        map[string][]bindings.FieldData {
            "NSXT_CONTAINER_PLUGIN": []bindings.FieldData {
                 bindings.NewFieldData("ncp_cluster_network_spec", true),
            },
        },
    )
    validators = append(validators, uv1)
    return bindings.NewStructType("com.vmware.vcenter.namespace_management.clusters.set_spec",fields, reflect.TypeOf(ClustersSetSpec{}), fieldNameMap, validators)
}


