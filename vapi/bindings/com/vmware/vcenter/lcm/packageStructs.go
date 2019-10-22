/* Copyright © 2019 VMware, Inc. All Rights Reserved.
     SPDX-License-Identifier: BSD-2-Clause */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Data type definitions file for package: com.vmware.vcenter.lcm.
 * Includes binding types of a top level structures and enumerations.
 * Shared by client-side stubs and server-side skeletons to ensure type
 * compatibility.
 */

package lcm

import (
    "reflect"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/bindings/com/vmware/cis/task"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/bindings/com/vmware/vapi/std"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/bindings"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/data"
    "time"
)


// The size of appliance to be deployed. **Warning:** This enumeration is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
//
// <p> See {@link com.vmware.vapi.bindings.ApiEnumeration enumerated types description}.

type ApplianceSize string

const (
    // Appliance size of 'tiny', Default vCPUs: 2, Memory: 8GB, VM: 100, Hosts: 10. **Warning:** This constant field is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
     ApplianceSize_TINY ApplianceSize = "TINY"
    // Appliance size of 'small', Default vCPUs: 4, Memory: 16GB, VM: 1000, Hosts: 100. **Warning:** This constant field is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
     ApplianceSize_SMALL ApplianceSize = "SMALL"
    // Appliance size of 'medium', Default vCPUs: 8, Memory: 24GB, VM: 4000, Hosts: 400. **Warning:** This constant field is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
     ApplianceSize_MEDIUM ApplianceSize = "MEDIUM"
    // Appliance size of 'large', Default vCPUs: 16, Memory: 32GB, VM: 10000, Hosts: 1000. **Warning:** This constant field is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
     ApplianceSize_LARGE ApplianceSize = "LARGE"
    // Appliance size of 'extra large', Default vCPUs: 24, Memory: 48GB, VM: 35000, Hosts: 2000. **Warning:** This constant field is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
     ApplianceSize_XLARGE ApplianceSize = "XLARGE"
)

func (a ApplianceSize) ApplianceSize() bool {
    switch a {
        case ApplianceSize_TINY:
            return true
        case ApplianceSize_SMALL:
            return true
        case ApplianceSize_MEDIUM:
            return true
        case ApplianceSize_LARGE:
            return true
        case ApplianceSize_XLARGE:
            return true
        default:
            return false
    }
}




// The type of appliance to be deployed. **Warning:** This enumeration is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
//
// <p> See {@link com.vmware.vapi.bindings.ApiEnumeration enumerated types description}.

type ApplianceType string

const (
    // Management node. **Warning:** This constant field is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
     ApplianceType_VCSA_EXTERNAL ApplianceType = "VCSA_EXTERNAL"
    // Embedded node. **Warning:** This constant field is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
     ApplianceType_VCSA_EMBEDDED ApplianceType = "VCSA_EMBEDDED"
    // Infrastructure node. **Warning:** This constant field is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
     ApplianceType_PSC ApplianceType = "PSC"
    // VMC node. **Warning:** This constant field is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
     ApplianceType_VMC ApplianceType = "VMC"
)

func (a ApplianceType) ApplianceType() bool {
    switch a {
        case ApplianceType_VCSA_EXTERNAL:
            return true
        case ApplianceType_VCSA_EMBEDDED:
            return true
        case ApplianceType_PSC:
            return true
        case ApplianceType_VMC:
            return true
        default:
            return false
    }
}




// The storage size of the appliance to be deployed. **Warning:** This enumeration is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
//
// <p> See {@link com.vmware.vapi.bindings.ApiEnumeration enumerated types description}.

type StorageSize string

const (
    // Large storage. **Warning:** This constant field is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
     StorageSize_LARGE StorageSize = "LARGE"
    // Extra large storage. **Warning:** This constant field is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
     StorageSize_XLARGE StorageSize = "XLARGE"
    // Regular storage. **Warning:** This constant field is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
     StorageSize_REGULAR StorageSize = "REGULAR"
)

func (s StorageSize) StorageSize() bool {
    switch s {
        case StorageSize_LARGE:
            return true
        case StorageSize_XLARGE:
            return true
        case StorageSize_REGULAR:
            return true
        default:
            return false
    }
}





// Connection information for source/destination location. **Warning:** This class is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
type Connection struct {
    // The IP address or DNS resolvable name of the ESX/VC host. If a DNS resolvable name is provided, it must be resolvable from the machine that is running the installer. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    Hostname string
    // A username with administrative privileges on the ESX/VC host. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    Username string
    // The password of the 'username' on the ESX/VC host. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    Password string
    // The port number for the ESX/VC. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    HttpsPort *int64
    // A flag to indicate whether the ssl verification is required. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    SslVerify *bool
    // Thumbprint for SSL verification. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    SslThumbprint *string
}



//


// Information about the appliance deployed. **Warning:** This class is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
type DeploymentInfo struct {
    // The name of the appliance. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    ApplianceName string
    // The FQDN of the appliance. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    ApplianceFqdn *string
    // The ip addresses of the appliance. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    ApplianceIps []string
}



//


// Container to control deployment. **Warning:** This class is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
type DeploymentOption struct {
    // The options control if a task should be skipped. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    SkipOptions map[DeploymentOption_SkipOptions]bool
}



//
    
    // Skippable tasks. **Warning:** This enumeration is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    //
    // <p> See {@link com.vmware.vapi.bindings.ApiEnumeration enumerated types description}.
     
    type DeploymentOption_SkipOptions string

    const (
        // Skips the sso check. This should only be used when performing precheck for install/upgrade of management node before infrastructure node is deployed. **Warning:** This constant field is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
         DeploymentOption_SkipOptions_SKIP_SSO_CHECK DeploymentOption_SkipOptions = "SKIP_SSO_CHECK"
    )

    func (s DeploymentOption_SkipOptions) DeploymentOption_SkipOptions() bool {
        switch s {
            case DeploymentOption_SkipOptions_SKIP_SSO_CHECK:
                return true
            default:
                return false
        }
    }



// Spec to describe the new appliance. **Warning:** This class is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
type DestinationAppliance struct {
    // The name of the appliance to deploy. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    ApplianceName string
    // The type of appliance to deploy. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    ApplianceType ApplianceType
    // A size descriptor based on the number of virtual machines which will be managed by the new vCenter appliance. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    ApplianceSize *ApplianceSize
    // The disk size of the new vCenter appliance. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    ApplianceDiskSize *StorageSize
    // Password must conform to the following requirements: 1. At least 8 characters. 2. No more than 20 characters. 3. At least 1 uppercase character. 4. At least 1 lowercase character. 5. At least 1 number. 6. At least 1 special character (e.g., '!', '(', '\\\\@', etc.). 7. Only visible A-Z, a-z, 0-9 and punctuation (spaces are not allowed). **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    RootPassword string
    // Whether to deploy the appliance with thin mode virtual disks. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    ThinDiskMode bool
    // The location of the OVA file. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    OvaLocation string
    // A flag to indicate whether the ssl verification is required. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    OvaLocationSslVerify *bool
    // SSL thumbprint of ssl verification. If provided, ssl_verify can be omitted or set to true. If omitted, ssl_verify must be false. If omitted and ssl_verify is true, an error will occur. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    OvaLocationSslThumbprint *string
    // The location of the OVF Tool. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    OvftoolLocation string
    // Flag to indicate whether or not to verify the SSL thumbprint of OVF Tool location. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    OvftoolLocationSslVerify *bool
    // SSL thumbprint of OVF Tool location to be verified. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    OvftoolLocationSslThumbprint *string
    // The configuration of vCenter services. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    Services DestinationApplianceService
    // The network settings of the appliance to be deployed. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    Network Network
    // Configuration of the vCSA time synchronization. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    Time Time
    // The OVF Tool arguments to be included. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    OvftoolArguments map[string]string
    // Spec used to configure an embedded vCenter Server. This field describes how the embedded vCenter Server appliance should be configured. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    VcsaEmbedded *VcsaEmbedded
    // Spec used to configure a vCenter Server registered with an external PSC. If unset, either vcsa_embedded or psc must be provided. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    Vmc *ExternalVcsa
}



//


// The configuration of vCenter services. **Warning:** This class is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
type DestinationApplianceService struct {
    // Port numbers on which the vCenter Server Appliance communicates with the other vSphere components. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    Rhttpproxy *ReverseProxy
    // Whether to enable SSH on the vCenter Appliance. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    Ssh Ssh
}



//


// Configuration of destination location. **Warning:** This class is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
type DestinationLocation struct {
    // This section describes the ESX host on which to deploy the appliance. Required if you are deploying the appliance directly on an ESX host. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    Esx *Esx
    // This subsection describes the vCenter on which to deploy the appliance. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    Vcenter *Vc
}



//


// Configuration of the replicated Single Sign-On for Embedded type deployment. **Warning:** This class is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
type EmbeddedReplicatedVcsa struct {
    // Administrator password of the existing Single Sign-On to be replicated. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    SsoAdminPassword string
    // Domain name for the remote appliance which is being replicated. For example, 'vsphere.local'. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    SsoDomainName string
    // The IP address or DNS resolvable name for the remote appliance. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    PartnerHostname string
    // A flag to indicate whether the ssl verification is required. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    SslVerify *bool
    // SHA1 thumbprint of the server SSL certificate which will be used for verification. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    SslThumbprint *string
    // The HTTPS port of the external PSC appliance. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    HttpsPort *int64
}



//


// Configuration of the standalone Single Sign-On for Embedded type deployment. **Warning:** This class is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
type EmbeddedStandaloneVcsa struct {
    // Password must conform to the following requirements: 1. At least 8 characters. 2. No more than 20 characters. 3. At least 1 uppercase character. 4. At least 1 lowercase character. 5. At least 1 number. 6. At least 1 special character (e.g., '!', '(', '\\\\@', etc.). 7. Only visible A-Z, a-z, 0-9 and punctuation (spaces are not allowed). **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    SsoAdminPassword string
    // The Single Sign-On domain name to be used to configure this vCenter Server Appliance. For example, 'vsphere.local'. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    SsoDomainName string
}



//


// Configuration of ESX. **Warning:** This class is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
type Esx struct {
    // The configuration to connect to an ESX/VC. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    Connection Connection
    // The configuration of ESX inventory. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    Inventory EsxInventory
}



//


// Configuration of ESX's inventory. **Warning:** This class is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
type EsxInventory struct {
    // The datastore on which to store the files of the appliance. This value has to be either a specific datastore name, or a specific datastore in a datastore cluster. The datastore must be accessible from the ESX host and must have at least 25 GB of free space. Otherwise, the new appliance might not power on. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    DatastoreName string
    // The network of the ESX host to which the new appliance should connect. Omit this parameter if the ESX host has one network. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    NetworkName *string
    // The path to the resource pool on the ESX host in which the appliance will be deployed. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    ResourcePoolPath *string
}



//


// Configuration of the external tools used. **Warning:** This class is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
type ExternalTool struct {
    // The name of the external tool. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    Name string
    // The host name of the external tool. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    Hostname *string
    // The location of the external tool. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    Location string
}



//


// Configuration of the Single Sign-On for Management type deployment. **Warning:** This class is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
type ExternalVcsa struct {
    // Administrator password of the external PSC to register with. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    SsoAdminPassword string
    // Domain name of the external PSC. For example, 'vsphere.local'. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    SsoDomainName string
    // The IP address or DNS resolvable name of the remote PSC to which this configuring vCenter Server will be registered. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    PscHostname string
    // A flag to indicate whether the SSL verification is required. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    SslVerify *bool
    // SHA1 thumbprint of the server SSL certificate which will be used for verification. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    SslThumbprint *string
    // The HTTPS port of the external PSC appliance. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    HttpsPort *int64
}



//


// The specification that represents a install operation. **Warning:** This class is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
type InstallSpec struct {
    // This subsection describes the ESX or VC on which to deploy the appliance. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    DestinationLocation DestinationLocation
    // Spec to describe the new appliance. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    DestinationAppliance DestinationAppliance
    // The time to start the execution of deploying an appliance. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    ScheduledStart *time.Time
    // The amount of minutes by which the execution is allowed to be delayed. This field will be ignored if scheduled start is not provided. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    DelayTolerance *int64
}



//


// Network configuration of the appliance to be deployed. **Warning:** This class is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
type Network struct {
    // Primary network identity. Can be either an IP address or a fully qualified domain name(FQDN). **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    Hostname *string
    // Network IP address family. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    IpFamily *TemporaryNetwork_IpType
    // Network mode. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    Mode TemporaryNetwork_NetworkMode
    // Network IP address. Required for static mode only. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    Ip *string
    // A comma-separated list of IP addresses of DNS servers. A JSON array such as ["1.2.3.4", "127.0.0.1"]. Required for static mode only. DNS servers must be reachable from the machine that runs CLI installer. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    DnsServers []string
    // Network prefix length. Required for static mode only. Remove if the mode is "dhcp". This is the number of bits set in the subnet mask; for instance, if the subnet mask is 255.255.255.0, there are 24 bits in the binary version of the subnet mask, so the prefix length is 24. If used, the values must be in the inclusive range of 0 to 32 for IPv4 and 0 to 128 for IPv6. Required for static mode only. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    Prefix *int64
    // Gateway of the network. Required for static mode only. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    Gateway *string
}



//


// The ``Notification`` class describes a notification that can be reported by the appliance task, which can be of type info, warning or errors.
type Notification struct {
    // The notification id.
    Id string
    // The time the notification was raised/found.
    Time *time.Time
    // The notification message.
    Message std.LocalizableMessage
    // The resolution message, if any.
    Resolution *std.LocalizableMessage
}



//


// The ``Notifications`` class contains info/warning/error messages that can be reported be the appliance task.
type Notifications struct {
    // Info notification messages reported.
    Info []Notification
    // Warning notification messages reported.
    Warnings []Notification
    // Error notification messages reported.
    Errors []Notification
}



//


// Spec used to configure a Platform Services Controller. This section describes how the Platform Services Controller appliance should be configured. If unset, either ``#vcsaEmbedded`` or ``#vcsaExternal`` must be provided. **Warning:** This class is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
type Psc struct {
    // Spec used to configure a standalone Platform Services Controller. This section describes how the standalone PSC should be configured. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    Standalone *PscStandalone
    // Spec used to configure a replicated Platform Services Controller. This section describes how the replicated PSC should be configured. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    Replicated *PscReplicated
    // This key describes the enabling option for the VMware's Customer Experience Improvement Program (CEIP). By default we have ``ceipEnabled``: true, which indicates that you are joining CEIP. If you prefer not to participate in the VMware's CEIP for this product, you must disable CEIP by setting ``ceipEnabled``: false. You may join or leave VMware's CEIP for this product at any time. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    CeipEnabled bool
}



//


// Configuration of the replicated Single Sign-On for PSC type deployment. **Warning:** This class is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
type PscReplicated struct {
    // Administrator password of the PSC to be replicated. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    SsoAdminPassword string
    // Domain name of the remote PSC. For example, 'vsphere.local'. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    SsoDomainName string
    // The IP address or DNS resolvable name of the remote PSC. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    PartnerHostname string
    // A flag to indicate whether the SSL verification is required. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    SslVerify *bool
    // SHA1 thumbprint of the server SSL certificate which will be used for verification. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    SslThumbprint *string
    // The HTTPS port of the remote PSC. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    HttpsPort *int64
    // Site name of the newly deployed PSC. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    SsoSiteName *string
}



//


// Configuration of the standalone Single Sign-On for Embedded type deployment. **Warning:** This class is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
type PscStandalone struct {
    // Password must conform to the following requirements: 1. At least 8 characters. 2. No more than 20 characters. 3. At least 1 uppercase character. 4. At least 1 lowercase character. 5. At least 1 number. 6. At least 1 special character (e.g., '!', '(', '\\\\@', etc.). 7. Only visible A-Z, a-z, 0-9 and punctuation (spaces are not allowed). **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    SsoAdminPassword string
    // Domain name of the newly deployed PSC. For example, 'vsphere.local'. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    SsoDomainName string
    // Site name of the PSC. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    SsoSiteName *string
}



//


// Container of info, warning and error messages associated with a single task. **Warning:** This class is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
type Result struct {
    // Info notification messages reported. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    Info []Notification
    // Warning notification messages reported. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    Warnings []Notification
    // Error notification messages reported. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    Errors []Notification
}



//


// Port numbers on which the vCenter Server Appliance communicates with the other vSphere components. **Warning:** This class is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
type ReverseProxy struct {
    // Reverse proxy http port. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    HttpPort *int64
    // Reverse proxy https port. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    HttpsPort *int64
}



//


// Setting to enable SSH on the deployed appliance. **Warning:** This class is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
type Ssh struct {
    // Whether to enable SSH. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    Enabled *bool
}



//


// Container that contains the status information about a single task. **Warning:** This class is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
type SubTaskInfo struct {
    // The progress info of this deployment task. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    Progress task.Progress
    // The time that the last update is registered. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    LastUpdatedTime time.Time
    // Result of the task. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    Result *Result
    // External tools used for the deployment. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    ExternalTools []ExternalTool
    // Description of the operation associated with the task.
    Description std.LocalizableMessage
    // Identifier of the service containing the operation.
    Service string
    // Identifier of the operation associated with the task.
    Operation string
    // Parent of the current task.
    Parent *string
    // Identifier of the target created by the operation or an existing one the operation performed on.
    Target *std.DynamicID
    // Status of the operation associated with the task.
    Status task.Status
    // Flag to indicate whether or not the operation can be cancelled. The value may change as the operation progresses.
    Cancelable bool
    // Description of the error if the operation status is "FAILED".
    Error *data.ErrorValue
    // Time when the operation is started.
    StartTime *time.Time
    // Time when the operation is completed.
    EndTime *time.Time
    // Name of the user who performed the operation.
    User *string
}



//


// The container that contains the status information of a deployment. **Warning:** This class is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
type TaskInfo struct {
    // The path of the metadata file. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    MetadataFile string
    // The state of appliance being deployed. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    State *string
    // The total progress of the deployment operation. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    Progress *task.Progress
    // The time that the last update is registered. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    LastUpdatedTime time.Time
    // The ordered list of subtasks for this deployment operation. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    SubtaskOrder [][]string
    // The map of the deployment subtasks and their status information. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    Subtasks map[string]SubTaskInfo
    // Information about the appliance deployed. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    ApplianceInfo *DeploymentInfo
    // The result of validation or recommendation requests. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    Result data.DataValue
    // Additional information that a response may contain. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    AdditionalInfo *string
    // Description of the operation associated with the task.
    Description std.LocalizableMessage
    // Identifier of the service containing the operation.
    Service string
    // Identifier of the operation associated with the task.
    Operation string
    // Parent of the current task.
    Parent *string
    // Identifier of the target created by the operation or an existing one the operation performed on.
    Target *std.DynamicID
    // Status of the operation associated with the task.
    Status task.Status
    // Flag to indicate whether or not the operation can be cancelled. The value may change as the operation progresses.
    Cancelable bool
    // Description of the error if the operation status is "FAILED".
    Error *data.ErrorValue
    // Time when the operation is started.
    StartTime *time.Time
    // Time when the operation is completed.
    EndTime *time.Time
    // Name of the user who performed the operation.
    User *string
}



//


// Configuration of the temporary network which is used during upgrade/migrate. **Warning:** This class is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
type TemporaryNetwork struct {
    // Network IP address family. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    IpFamily *TemporaryNetwork_IpType
    // Network mode. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    Mode TemporaryNetwork_NetworkMode
    // Network IP address. Required for static mode only. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    Ip *string
    // A comma-separated list of IP addresses of DNS servers. A JSON array such as ["1.2.3.4", "127.0.0.1"]. Required for static mode only. DNS servers must be reachable from the machine that runs CLI installer. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    DnsServers []string
    // Network prefix length. Required for static mode only. Remove if the mode is "dhcp". This is the number of bits set in the subnet mask; for instance, if the subnet mask is 255.255.255.0, there are 24 bits in the binary version of the subnet mask, so the prefix length is 24. If used, the values must be in the inclusive range of 0 to 32 for IPv4 and 0 to 128 for IPv6. Required for static mode only. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    Prefix *int64
    // Gateway of the network. Required for static mode only. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    Gateway *string
}



//
    
    // Network IP address family. **Warning:** This enumeration is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    //
    // <p> See {@link com.vmware.vapi.bindings.ApiEnumeration enumerated types description}.
     
    type TemporaryNetwork_IpType string

    const (
        // IPv4 Type of IP address. **Warning:** This constant field is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
         TemporaryNetwork_IpType_IPV4 TemporaryNetwork_IpType = "IPV4"
        // IPv6 Type of IP address. **Warning:** This constant field is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
         TemporaryNetwork_IpType_IPV6 TemporaryNetwork_IpType = "IPV6"
    )

    func (i TemporaryNetwork_IpType) TemporaryNetwork_IpType() bool {
        switch i {
            case TemporaryNetwork_IpType_IPV4:
                return true
            case TemporaryNetwork_IpType_IPV6:
                return true
            default:
                return false
        }
    }

    
    // Network mode. **Warning:** This enumeration is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    //
    // <p> See {@link com.vmware.vapi.bindings.ApiEnumeration enumerated types description}.
     
    type TemporaryNetwork_NetworkMode string

    const (
        // DHCP mode. **Warning:** This constant field is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
         TemporaryNetwork_NetworkMode_DHCP TemporaryNetwork_NetworkMode = "DHCP"
        // Static IP mode. **Warning:** This constant field is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
         TemporaryNetwork_NetworkMode_STATIC TemporaryNetwork_NetworkMode = "STATIC"
    )

    func (n TemporaryNetwork_NetworkMode) TemporaryNetwork_NetworkMode() bool {
        switch n {
            case TemporaryNetwork_NetworkMode_DHCP:
                return true
            case TemporaryNetwork_NetworkMode_STATIC:
                return true
            default:
                return false
        }
    }



// NTP setting of the appliance to be deployed. **Warning:** This class is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
type Time struct {
    // To configure NTP time synchronization for the appliance, set the value to a comma - separated list of host names or IP addresses of Network Time Protocol(NTP) servers. If "ntp_servers" is not provided, the appliance clock will be synced to the ESX. For example: ["time.nist.gov"]. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    NtpServers []string
}



//


// Configuration of the VC that hosts/will host an appliance. **Warning:** This class is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
type Vc struct {
    // The configuration to connect to an ESX/VC. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    Connection Connection
    // All names are case-sensitive. you can install the appliance to one of the following destinations: 1. A resource pool in a cluster, use 'cluster_path'. 2. A specific ESX host in a cluster, use 'host_path'. 3. A resource pool in a specific ESX host being managed by the current vCenter, use 'resource_pool_path'. You must always provide the 'network_name' key. To install a new appliance to a specific ESX host in a cluster, provide the 'host_path' key, and the 'datastore_name', e.g. 'host_path': '/MyDataCenter/host/MyCluster/10.20.30.40', 'datastore_name': 'Your Datastore'. To install a new appliance to a specific resource pool, provide the 'resource_pool_path', and the 'datastore_name', e.g. 'resource_pool_path': '/Your Datacenter Folder/Your Datacenter/host/Your Cluster/Resources/Your Resource Pool', 'datastore_name': 'Your Datastore'. To place a new appliance to a virtual machine Folder, provide the 'vm_folder_path', e.g. 'vm_folder_path': 'VM Folder 0/VM Folder1'. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    Inventory VcInventory
}



//


// Inventory information about a VCenter. **Warning:** This class is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
type VcInventory struct {
    // Path of the VM folder. VM folder must be visible by the Data Center of the compute resourceFormat:{vm_folder1}/{vm_folder2}e.g.:'VM Folder 0/VM Folder1'. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    VmFolderPath *string
    // Full path to resource pool. Format: /{datacenter folder}/{datacenter name}/host/{host name}/{cluster_name}/Resources/{resource pool}. e.g: Your Datacenter Folder/Your Datacenter/host/Your Cluster/Resources/Your Resource Pool. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    ResourcePoolPath *string
    // Full path to the cluster. Format: /{datacenter folder}/{datacenter name}/host/{cluster_name}. e.g: /Your Datacenter Folder/Your Datacenter/host/Your Cluster. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    ClusterPath *string
        HostPath *string
    // The datastore on which to store the files of the appliance. This value has to be either a specific datastore name, or a specific datastore in a datastore cluster. The datastore must be accessible from the ESX host and must have at least 25 GB of free space. Otherwise, the new appliance might not power on. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    DatastoreName *string
    // The datastore cluster on which to store the files of the appliance. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    DatastoreClusterName *string
    // Name of the network. e.g. VM Network. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    NetworkName string
}



//


// Spec used to configure an embedded vCenter Server. This field describes how the embedded vCenter Server appliance should be configured. **Warning:** This class is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
type VcsaEmbedded struct {
    // Spec used to configure a standalone embedded vCenter Server. This field describes how the standalone vCenter Server appliance should be configured. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    Standalone *EmbeddedStandaloneVcsa
    // Spec used to configure a replicated embedded vCenter Server. This field describes how the replicated vCenter Server appliance should be configured. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    Replicated *EmbeddedReplicatedVcsa
    // This key describes the enabling option for the VMware's Customer Experience Improvement Program (CEIP). By default we have ``ceipEnabled``: true, which indicates that you are joining CEIP. If you prefer not to participate in the VMware's CEIP for this product, you must disable CEIP by setting ``ceipEnabled``: false. You may join or leave VMware's CEIP for this product at any time. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    CeipEnabled bool
}



//






func ConnectionBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["hostname"] = bindings.NewStringType()
    fieldNameMap["hostname"] = "Hostname"
    fields["username"] = bindings.NewStringType()
    fieldNameMap["username"] = "Username"
    fields["password"] = bindings.NewSecretType()
    fieldNameMap["password"] = "Password"
    fields["https_port"] = bindings.NewOptionalType(bindings.NewIntegerType())
    fieldNameMap["https_port"] = "HttpsPort"
    fields["ssl_verify"] = bindings.NewOptionalType(bindings.NewBooleanType())
    fieldNameMap["ssl_verify"] = "SslVerify"
    fields["ssl_thumbprint"] = bindings.NewOptionalType(bindings.NewStringType())
    fieldNameMap["ssl_thumbprint"] = "SslThumbprint"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.lcm.connection",fields, reflect.TypeOf(Connection{}), fieldNameMap, validators)
}

func DeploymentInfoBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["appliance_name"] = bindings.NewStringType()
    fieldNameMap["appliance_name"] = "ApplianceName"
    fields["appliance_fqdn"] = bindings.NewOptionalType(bindings.NewStringType())
    fieldNameMap["appliance_fqdn"] = "ApplianceFqdn"
    fields["appliance_ips"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewStringType(), reflect.TypeOf([]string{})))
    fieldNameMap["appliance_ips"] = "ApplianceIps"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.lcm.deployment_info",fields, reflect.TypeOf(DeploymentInfo{}), fieldNameMap, validators)
}

func DeploymentOptionBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["skip_options"] = bindings.NewMapType(bindings.NewEnumType("com.vmware.vcenter.lcm.deployment_option.skip_options", reflect.TypeOf(DeploymentOption_SkipOptions(DeploymentOption_SkipOptions_SKIP_SSO_CHECK))), bindings.NewBooleanType(),reflect.TypeOf(map[DeploymentOption_SkipOptions]bool{}))
    fieldNameMap["skip_options"] = "SkipOptions"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.lcm.deployment_option",fields, reflect.TypeOf(DeploymentOption{}), fieldNameMap, validators)
}

func DestinationApplianceBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["appliance_name"] = bindings.NewStringType()
    fieldNameMap["appliance_name"] = "ApplianceName"
    fields["appliance_type"] = bindings.NewEnumType("com.vmware.vcenter.lcm.appliance_type", reflect.TypeOf(ApplianceType(ApplianceType_VCSA_EXTERNAL)))
    fieldNameMap["appliance_type"] = "ApplianceType"
    fields["appliance_size"] = bindings.NewOptionalType(bindings.NewEnumType("com.vmware.vcenter.lcm.appliance_size", reflect.TypeOf(ApplianceSize(ApplianceSize_TINY))))
    fieldNameMap["appliance_size"] = "ApplianceSize"
    fields["appliance_disk_size"] = bindings.NewOptionalType(bindings.NewEnumType("com.vmware.vcenter.lcm.storage_size", reflect.TypeOf(StorageSize(StorageSize_LARGE))))
    fieldNameMap["appliance_disk_size"] = "ApplianceDiskSize"
    fields["root_password"] = bindings.NewSecretType()
    fieldNameMap["root_password"] = "RootPassword"
    fields["thin_disk_mode"] = bindings.NewBooleanType()
    fieldNameMap["thin_disk_mode"] = "ThinDiskMode"
    fields["ova_location"] = bindings.NewStringType()
    fieldNameMap["ova_location"] = "OvaLocation"
    fields["ova_location_ssl_verify"] = bindings.NewOptionalType(bindings.NewBooleanType())
    fieldNameMap["ova_location_ssl_verify"] = "OvaLocationSslVerify"
    fields["ova_location_ssl_thumbprint"] = bindings.NewOptionalType(bindings.NewStringType())
    fieldNameMap["ova_location_ssl_thumbprint"] = "OvaLocationSslThumbprint"
    fields["ovftool_location"] = bindings.NewStringType()
    fieldNameMap["ovftool_location"] = "OvftoolLocation"
    fields["ovftool_location_ssl_verify"] = bindings.NewOptionalType(bindings.NewBooleanType())
    fieldNameMap["ovftool_location_ssl_verify"] = "OvftoolLocationSslVerify"
    fields["ovftool_location_ssl_thumbprint"] = bindings.NewOptionalType(bindings.NewStringType())
    fieldNameMap["ovftool_location_ssl_thumbprint"] = "OvftoolLocationSslThumbprint"
    fields["services"] = bindings.NewReferenceType(DestinationApplianceServiceBindingType)
    fieldNameMap["services"] = "Services"
    fields["network"] = bindings.NewReferenceType(NetworkBindingType)
    fieldNameMap["network"] = "Network"
    fields["time"] = bindings.NewReferenceType(TimeBindingType)
    fieldNameMap["time"] = "Time"
    fields["ovftool_arguments"] = bindings.NewOptionalType(bindings.NewMapType(bindings.NewStringType(), bindings.NewStringType(),reflect.TypeOf(map[string]string{})))
    fieldNameMap["ovftool_arguments"] = "OvftoolArguments"
    fields["vcsa_embedded"] = bindings.NewOptionalType(bindings.NewReferenceType(VcsaEmbeddedBindingType))
    fieldNameMap["vcsa_embedded"] = "VcsaEmbedded"
    fields["vmc"] = bindings.NewOptionalType(bindings.NewReferenceType(ExternalVcsaBindingType))
    fieldNameMap["vmc"] = "Vmc"
    var validators = []bindings.Validator{}
    uv1 := bindings.NewUnionValidator("appliance_type",
        map[string][]bindings.FieldData {
            "VCSA_EMBEDDED": []bindings.FieldData {
                 bindings.NewFieldData("appliance_size", false),
                 bindings.NewFieldData("appliance_disk_size", false),
                 bindings.NewFieldData("vcsa_embedded", true),
            },
            "VMC": []bindings.FieldData {
                 bindings.NewFieldData("vmc", true),
            },
            "VCSA_EXTERNAL": []bindings.FieldData {},
            "PSC": []bindings.FieldData {},
        },
    )
    validators = append(validators, uv1)
    return bindings.NewStructType("com.vmware.vcenter.lcm.destination_appliance",fields, reflect.TypeOf(DestinationAppliance{}), fieldNameMap, validators)
}

func DestinationApplianceServiceBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["rhttpproxy"] = bindings.NewOptionalType(bindings.NewReferenceType(ReverseProxyBindingType))
    fieldNameMap["rhttpproxy"] = "Rhttpproxy"
    fields["ssh"] = bindings.NewReferenceType(SshBindingType)
    fieldNameMap["ssh"] = "Ssh"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.lcm.destination_appliance_service",fields, reflect.TypeOf(DestinationApplianceService{}), fieldNameMap, validators)
}

func DestinationLocationBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["esx"] = bindings.NewOptionalType(bindings.NewReferenceType(EsxBindingType))
    fieldNameMap["esx"] = "Esx"
    fields["vcenter"] = bindings.NewOptionalType(bindings.NewReferenceType(VcBindingType))
    fieldNameMap["vcenter"] = "Vcenter"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.lcm.destination_location",fields, reflect.TypeOf(DestinationLocation{}), fieldNameMap, validators)
}

func EmbeddedReplicatedVcsaBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["sso_admin_password"] = bindings.NewSecretType()
    fieldNameMap["sso_admin_password"] = "SsoAdminPassword"
    fields["sso_domain_name"] = bindings.NewStringType()
    fieldNameMap["sso_domain_name"] = "SsoDomainName"
    fields["partner_hostname"] = bindings.NewStringType()
    fieldNameMap["partner_hostname"] = "PartnerHostname"
    fields["ssl_verify"] = bindings.NewOptionalType(bindings.NewBooleanType())
    fieldNameMap["ssl_verify"] = "SslVerify"
    fields["ssl_thumbprint"] = bindings.NewOptionalType(bindings.NewStringType())
    fieldNameMap["ssl_thumbprint"] = "SslThumbprint"
    fields["https_port"] = bindings.NewOptionalType(bindings.NewIntegerType())
    fieldNameMap["https_port"] = "HttpsPort"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.lcm.embedded_replicated_vcsa",fields, reflect.TypeOf(EmbeddedReplicatedVcsa{}), fieldNameMap, validators)
}

func EmbeddedStandaloneVcsaBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["sso_admin_password"] = bindings.NewSecretType()
    fieldNameMap["sso_admin_password"] = "SsoAdminPassword"
    fields["sso_domain_name"] = bindings.NewStringType()
    fieldNameMap["sso_domain_name"] = "SsoDomainName"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.lcm.embedded_standalone_vcsa",fields, reflect.TypeOf(EmbeddedStandaloneVcsa{}), fieldNameMap, validators)
}

func EsxBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["connection"] = bindings.NewReferenceType(ConnectionBindingType)
    fieldNameMap["connection"] = "Connection"
    fields["inventory"] = bindings.NewReferenceType(EsxInventoryBindingType)
    fieldNameMap["inventory"] = "Inventory"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.lcm.esx",fields, reflect.TypeOf(Esx{}), fieldNameMap, validators)
}

func EsxInventoryBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["datastore_name"] = bindings.NewStringType()
    fieldNameMap["datastore_name"] = "DatastoreName"
    fields["network_name"] = bindings.NewOptionalType(bindings.NewStringType())
    fieldNameMap["network_name"] = "NetworkName"
    fields["resource_pool_path"] = bindings.NewOptionalType(bindings.NewStringType())
    fieldNameMap["resource_pool_path"] = "ResourcePoolPath"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.lcm.esx_inventory",fields, reflect.TypeOf(EsxInventory{}), fieldNameMap, validators)
}

func ExternalToolBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["name"] = bindings.NewStringType()
    fieldNameMap["name"] = "Name"
    fields["hostname"] = bindings.NewOptionalType(bindings.NewStringType())
    fieldNameMap["hostname"] = "Hostname"
    fields["location"] = bindings.NewStringType()
    fieldNameMap["location"] = "Location"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.lcm.external_tool",fields, reflect.TypeOf(ExternalTool{}), fieldNameMap, validators)
}

func ExternalVcsaBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["sso_admin_password"] = bindings.NewSecretType()
    fieldNameMap["sso_admin_password"] = "SsoAdminPassword"
    fields["sso_domain_name"] = bindings.NewStringType()
    fieldNameMap["sso_domain_name"] = "SsoDomainName"
    fields["psc_hostname"] = bindings.NewStringType()
    fieldNameMap["psc_hostname"] = "PscHostname"
    fields["ssl_verify"] = bindings.NewOptionalType(bindings.NewBooleanType())
    fieldNameMap["ssl_verify"] = "SslVerify"
    fields["ssl_thumbprint"] = bindings.NewOptionalType(bindings.NewStringType())
    fieldNameMap["ssl_thumbprint"] = "SslThumbprint"
    fields["https_port"] = bindings.NewOptionalType(bindings.NewIntegerType())
    fieldNameMap["https_port"] = "HttpsPort"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.lcm.external_vcsa",fields, reflect.TypeOf(ExternalVcsa{}), fieldNameMap, validators)
}

func InstallSpecBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["destination_location"] = bindings.NewReferenceType(DestinationLocationBindingType)
    fieldNameMap["destination_location"] = "DestinationLocation"
    fields["destination_appliance"] = bindings.NewReferenceType(DestinationApplianceBindingType)
    fieldNameMap["destination_appliance"] = "DestinationAppliance"
    fields["scheduled_start"] = bindings.NewOptionalType(bindings.NewDateTimeType())
    fieldNameMap["scheduled_start"] = "ScheduledStart"
    fields["delay_tolerance"] = bindings.NewOptionalType(bindings.NewIntegerType())
    fieldNameMap["delay_tolerance"] = "DelayTolerance"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.lcm.install_spec",fields, reflect.TypeOf(InstallSpec{}), fieldNameMap, validators)
}

func NetworkBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["hostname"] = bindings.NewOptionalType(bindings.NewStringType())
    fieldNameMap["hostname"] = "Hostname"
    fields["ip_family"] = bindings.NewOptionalType(bindings.NewEnumType("com.vmware.vcenter.lcm.temporary_network.ip_type", reflect.TypeOf(TemporaryNetwork_IpType(TemporaryNetwork_IpType_IPV4))))
    fieldNameMap["ip_family"] = "IpFamily"
    fields["mode"] = bindings.NewEnumType("com.vmware.vcenter.lcm.temporary_network.network_mode", reflect.TypeOf(TemporaryNetwork_NetworkMode(TemporaryNetwork_NetworkMode_DHCP)))
    fieldNameMap["mode"] = "Mode"
    fields["ip"] = bindings.NewOptionalType(bindings.NewStringType())
    fieldNameMap["ip"] = "Ip"
    fields["dns_servers"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewStringType(), reflect.TypeOf([]string{})))
    fieldNameMap["dns_servers"] = "DnsServers"
    fields["prefix"] = bindings.NewOptionalType(bindings.NewIntegerType())
    fieldNameMap["prefix"] = "Prefix"
    fields["gateway"] = bindings.NewOptionalType(bindings.NewStringType())
    fieldNameMap["gateway"] = "Gateway"
    var validators = []bindings.Validator{}
    uv1 := bindings.NewUnionValidator("mode",
        map[string][]bindings.FieldData {
            "STATIC": []bindings.FieldData {
                 bindings.NewFieldData("hostname", false),
                 bindings.NewFieldData("ip", true),
                 bindings.NewFieldData("dns_servers", true),
                 bindings.NewFieldData("prefix", true),
                 bindings.NewFieldData("gateway", true),
            },
            "DHCP": []bindings.FieldData {},
        },
    )
    validators = append(validators, uv1)
    return bindings.NewStructType("com.vmware.vcenter.lcm.network",fields, reflect.TypeOf(Network{}), fieldNameMap, validators)
}

func NotificationBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["id"] = bindings.NewStringType()
    fieldNameMap["id"] = "Id"
    fields["time"] = bindings.NewOptionalType(bindings.NewDateTimeType())
    fieldNameMap["time"] = "Time"
    fields["message"] = bindings.NewReferenceType(std.LocalizableMessageBindingType)
    fieldNameMap["message"] = "Message"
    fields["resolution"] = bindings.NewOptionalType(bindings.NewReferenceType(std.LocalizableMessageBindingType))
    fieldNameMap["resolution"] = "Resolution"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.lcm.notification",fields, reflect.TypeOf(Notification{}), fieldNameMap, validators)
}

func NotificationsBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["info"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewReferenceType(NotificationBindingType), reflect.TypeOf([]Notification{})))
    fieldNameMap["info"] = "Info"
    fields["warnings"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewReferenceType(NotificationBindingType), reflect.TypeOf([]Notification{})))
    fieldNameMap["warnings"] = "Warnings"
    fields["errors"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewReferenceType(NotificationBindingType), reflect.TypeOf([]Notification{})))
    fieldNameMap["errors"] = "Errors"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.lcm.notifications",fields, reflect.TypeOf(Notifications{}), fieldNameMap, validators)
}

func PscBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["standalone"] = bindings.NewOptionalType(bindings.NewReferenceType(PscStandaloneBindingType))
    fieldNameMap["standalone"] = "Standalone"
    fields["replicated"] = bindings.NewOptionalType(bindings.NewReferenceType(PscReplicatedBindingType))
    fieldNameMap["replicated"] = "Replicated"
    fields["ceip_enabled"] = bindings.NewBooleanType()
    fieldNameMap["ceip_enabled"] = "CeipEnabled"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.lcm.psc",fields, reflect.TypeOf(Psc{}), fieldNameMap, validators)
}

func PscReplicatedBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["sso_admin_password"] = bindings.NewSecretType()
    fieldNameMap["sso_admin_password"] = "SsoAdminPassword"
    fields["sso_domain_name"] = bindings.NewStringType()
    fieldNameMap["sso_domain_name"] = "SsoDomainName"
    fields["partner_hostname"] = bindings.NewStringType()
    fieldNameMap["partner_hostname"] = "PartnerHostname"
    fields["ssl_verify"] = bindings.NewOptionalType(bindings.NewBooleanType())
    fieldNameMap["ssl_verify"] = "SslVerify"
    fields["ssl_thumbprint"] = bindings.NewOptionalType(bindings.NewStringType())
    fieldNameMap["ssl_thumbprint"] = "SslThumbprint"
    fields["https_port"] = bindings.NewOptionalType(bindings.NewIntegerType())
    fieldNameMap["https_port"] = "HttpsPort"
    fields["sso_site_name"] = bindings.NewOptionalType(bindings.NewStringType())
    fieldNameMap["sso_site_name"] = "SsoSiteName"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.lcm.psc_replicated",fields, reflect.TypeOf(PscReplicated{}), fieldNameMap, validators)
}

func PscStandaloneBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["sso_admin_password"] = bindings.NewSecretType()
    fieldNameMap["sso_admin_password"] = "SsoAdminPassword"
    fields["sso_domain_name"] = bindings.NewStringType()
    fieldNameMap["sso_domain_name"] = "SsoDomainName"
    fields["sso_site_name"] = bindings.NewOptionalType(bindings.NewStringType())
    fieldNameMap["sso_site_name"] = "SsoSiteName"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.lcm.psc_standalone",fields, reflect.TypeOf(PscStandalone{}), fieldNameMap, validators)
}

func ResultBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["info"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewReferenceType(NotificationBindingType), reflect.TypeOf([]Notification{})))
    fieldNameMap["info"] = "Info"
    fields["warnings"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewReferenceType(NotificationBindingType), reflect.TypeOf([]Notification{})))
    fieldNameMap["warnings"] = "Warnings"
    fields["errors"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewReferenceType(NotificationBindingType), reflect.TypeOf([]Notification{})))
    fieldNameMap["errors"] = "Errors"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.lcm.result",fields, reflect.TypeOf(Result{}), fieldNameMap, validators)
}

func ReverseProxyBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["http_port"] = bindings.NewOptionalType(bindings.NewIntegerType())
    fieldNameMap["http_port"] = "HttpPort"
    fields["https_port"] = bindings.NewOptionalType(bindings.NewIntegerType())
    fieldNameMap["https_port"] = "HttpsPort"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.lcm.reverse_proxy",fields, reflect.TypeOf(ReverseProxy{}), fieldNameMap, validators)
}

func SshBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["enabled"] = bindings.NewOptionalType(bindings.NewBooleanType())
    fieldNameMap["enabled"] = "Enabled"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.lcm.ssh",fields, reflect.TypeOf(Ssh{}), fieldNameMap, validators)
}

func SubTaskInfoBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["progress"] = bindings.NewReferenceType(task.ProgressBindingType)
    fieldNameMap["progress"] = "Progress"
    fields["last_updated_time"] = bindings.NewDateTimeType()
    fieldNameMap["last_updated_time"] = "LastUpdatedTime"
    fields["result"] = bindings.NewOptionalType(bindings.NewReferenceType(ResultBindingType))
    fieldNameMap["result"] = "Result"
    fields["external_tools"] = bindings.NewListType(bindings.NewReferenceType(ExternalToolBindingType), reflect.TypeOf([]ExternalTool{}))
    fieldNameMap["external_tools"] = "ExternalTools"
    fields["description"] = bindings.NewReferenceType(std.LocalizableMessageBindingType)
    fieldNameMap["description"] = "Description"
    fields["service"] = bindings.NewIdType([]string {"com.vmware.vapi.service"}, "")
    fieldNameMap["service"] = "Service"
    fields["operation"] = bindings.NewIdType([]string {"com.vmware.vapi.operation"}, "")
    fieldNameMap["operation"] = "Operation"
    fields["parent"] = bindings.NewOptionalType(bindings.NewIdType([]string {"com.vmware.cis.task"}, ""))
    fieldNameMap["parent"] = "Parent"
    fields["target"] = bindings.NewOptionalType(bindings.NewReferenceType(std.DynamicIDBindingType))
    fieldNameMap["target"] = "Target"
    fields["status"] = bindings.NewEnumType("com.vmware.cis.task.status", reflect.TypeOf(task.Status(task.Status_PENDING)))
    fieldNameMap["status"] = "Status"
    fields["cancelable"] = bindings.NewBooleanType()
    fieldNameMap["cancelable"] = "Cancelable"
    fields["error"] = bindings.NewOptionalType(bindings.NewAnyErrorType())
    fieldNameMap["error"] = "Error"
    fields["start_time"] = bindings.NewOptionalType(bindings.NewDateTimeType())
    fieldNameMap["start_time"] = "StartTime"
    fields["end_time"] = bindings.NewOptionalType(bindings.NewDateTimeType())
    fieldNameMap["end_time"] = "EndTime"
    fields["user"] = bindings.NewOptionalType(bindings.NewStringType())
    fieldNameMap["user"] = "User"
    var validators = []bindings.Validator{}
    uv1 := bindings.NewUnionValidator("status",
        map[string][]bindings.FieldData {
            "FAILED": []bindings.FieldData {
                 bindings.NewFieldData("error", false),
                 bindings.NewFieldData("start_time", true),
                 bindings.NewFieldData("end_time", true),
            },
            "RUNNING": []bindings.FieldData {
                 bindings.NewFieldData("start_time", true),
            },
            "BLOCKED": []bindings.FieldData {
                 bindings.NewFieldData("start_time", true),
            },
            "SUCCEEDED": []bindings.FieldData {
                 bindings.NewFieldData("start_time", true),
                 bindings.NewFieldData("end_time", true),
            },
            "PENDING": []bindings.FieldData {},
        },
    )
    validators = append(validators, uv1)
    return bindings.NewStructType("com.vmware.vcenter.lcm.sub_task_info",fields, reflect.TypeOf(SubTaskInfo{}), fieldNameMap, validators)
}

func TaskInfoBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["metadata_file"] = bindings.NewStringType()
    fieldNameMap["metadata_file"] = "MetadataFile"
    fields["state"] = bindings.NewOptionalType(bindings.NewStringType())
    fieldNameMap["state"] = "State"
    fields["progress"] = bindings.NewOptionalType(bindings.NewReferenceType(task.ProgressBindingType))
    fieldNameMap["progress"] = "Progress"
    fields["last_updated_time"] = bindings.NewDateTimeType()
    fieldNameMap["last_updated_time"] = "LastUpdatedTime"
    fields["subtask_order"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewListType(bindings.NewStringType(), reflect.TypeOf([]string{})), reflect.TypeOf([][]string{})))
    fieldNameMap["subtask_order"] = "SubtaskOrder"
    fields["subtasks"] = bindings.NewOptionalType(bindings.NewMapType(bindings.NewStringType(), bindings.NewReferenceType(SubTaskInfoBindingType),reflect.TypeOf(map[string]SubTaskInfo{})))
    fieldNameMap["subtasks"] = "Subtasks"
    fields["appliance_info"] = bindings.NewOptionalType(bindings.NewReferenceType(DeploymentInfoBindingType))
    fieldNameMap["appliance_info"] = "ApplianceInfo"
    fields["result"] = bindings.NewOptionalType(bindings.NewOpaqueType())
    fieldNameMap["result"] = "Result"
    fields["additional_info"] = bindings.NewOptionalType(bindings.NewStringType())
    fieldNameMap["additional_info"] = "AdditionalInfo"
    fields["description"] = bindings.NewReferenceType(std.LocalizableMessageBindingType)
    fieldNameMap["description"] = "Description"
    fields["service"] = bindings.NewIdType([]string {"com.vmware.vapi.service"}, "")
    fieldNameMap["service"] = "Service"
    fields["operation"] = bindings.NewIdType([]string {"com.vmware.vapi.operation"}, "")
    fieldNameMap["operation"] = "Operation"
    fields["parent"] = bindings.NewOptionalType(bindings.NewIdType([]string {"com.vmware.cis.task"}, ""))
    fieldNameMap["parent"] = "Parent"
    fields["target"] = bindings.NewOptionalType(bindings.NewReferenceType(std.DynamicIDBindingType))
    fieldNameMap["target"] = "Target"
    fields["status"] = bindings.NewEnumType("com.vmware.cis.task.status", reflect.TypeOf(task.Status(task.Status_PENDING)))
    fieldNameMap["status"] = "Status"
    fields["cancelable"] = bindings.NewBooleanType()
    fieldNameMap["cancelable"] = "Cancelable"
    fields["error"] = bindings.NewOptionalType(bindings.NewAnyErrorType())
    fieldNameMap["error"] = "Error"
    fields["start_time"] = bindings.NewOptionalType(bindings.NewDateTimeType())
    fieldNameMap["start_time"] = "StartTime"
    fields["end_time"] = bindings.NewOptionalType(bindings.NewDateTimeType())
    fieldNameMap["end_time"] = "EndTime"
    fields["user"] = bindings.NewOptionalType(bindings.NewStringType())
    fieldNameMap["user"] = "User"
    var validators = []bindings.Validator{}
    uv1 := bindings.NewUnionValidator("status",
        map[string][]bindings.FieldData {
            "RUNNING": []bindings.FieldData {
                 bindings.NewFieldData("progress", true),
                 bindings.NewFieldData("start_time", true),
            },
            "FAILED": []bindings.FieldData {
                 bindings.NewFieldData("progress", true),
                 bindings.NewFieldData("error", false),
                 bindings.NewFieldData("start_time", true),
                 bindings.NewFieldData("end_time", true),
            },
            "SUCCEEDED": []bindings.FieldData {
                 bindings.NewFieldData("progress", true),
                 bindings.NewFieldData("start_time", true),
                 bindings.NewFieldData("end_time", true),
            },
            "BLOCKED": []bindings.FieldData {
                 bindings.NewFieldData("progress", true),
                 bindings.NewFieldData("start_time", true),
            },
            "PENDING": []bindings.FieldData {},
        },
    )
    validators = append(validators, uv1)
    return bindings.NewStructType("com.vmware.vcenter.lcm.task_info",fields, reflect.TypeOf(TaskInfo{}), fieldNameMap, validators)
}

func TemporaryNetworkBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["ip_family"] = bindings.NewOptionalType(bindings.NewEnumType("com.vmware.vcenter.lcm.temporary_network.ip_type", reflect.TypeOf(TemporaryNetwork_IpType(TemporaryNetwork_IpType_IPV4))))
    fieldNameMap["ip_family"] = "IpFamily"
    fields["mode"] = bindings.NewEnumType("com.vmware.vcenter.lcm.temporary_network.network_mode", reflect.TypeOf(TemporaryNetwork_NetworkMode(TemporaryNetwork_NetworkMode_DHCP)))
    fieldNameMap["mode"] = "Mode"
    fields["ip"] = bindings.NewOptionalType(bindings.NewStringType())
    fieldNameMap["ip"] = "Ip"
    fields["dns_servers"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewStringType(), reflect.TypeOf([]string{})))
    fieldNameMap["dns_servers"] = "DnsServers"
    fields["prefix"] = bindings.NewOptionalType(bindings.NewIntegerType())
    fieldNameMap["prefix"] = "Prefix"
    fields["gateway"] = bindings.NewOptionalType(bindings.NewStringType())
    fieldNameMap["gateway"] = "Gateway"
    var validators = []bindings.Validator{}
    uv1 := bindings.NewUnionValidator("mode",
        map[string][]bindings.FieldData {
            "STATIC": []bindings.FieldData {
                 bindings.NewFieldData("ip", true),
                 bindings.NewFieldData("dns_servers", true),
                 bindings.NewFieldData("prefix", true),
                 bindings.NewFieldData("gateway", true),
            },
            "DHCP": []bindings.FieldData {},
        },
    )
    validators = append(validators, uv1)
    return bindings.NewStructType("com.vmware.vcenter.lcm.temporary_network",fields, reflect.TypeOf(TemporaryNetwork{}), fieldNameMap, validators)
}

func TimeBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["ntp_servers"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewStringType(), reflect.TypeOf([]string{})))
    fieldNameMap["ntp_servers"] = "NtpServers"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.lcm.time",fields, reflect.TypeOf(Time{}), fieldNameMap, validators)
}

func VcBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["connection"] = bindings.NewReferenceType(ConnectionBindingType)
    fieldNameMap["connection"] = "Connection"
    fields["inventory"] = bindings.NewReferenceType(VcInventoryBindingType)
    fieldNameMap["inventory"] = "Inventory"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.lcm.vc",fields, reflect.TypeOf(Vc{}), fieldNameMap, validators)
}

func VcInventoryBindingType() bindings.BindingType {
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
    fields["datastore_name"] = bindings.NewOptionalType(bindings.NewStringType())
    fieldNameMap["datastore_name"] = "DatastoreName"
    fields["datastore_cluster_name"] = bindings.NewOptionalType(bindings.NewStringType())
    fieldNameMap["datastore_cluster_name"] = "DatastoreClusterName"
    fields["network_name"] = bindings.NewStringType()
    fieldNameMap["network_name"] = "NetworkName"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.lcm.vc_inventory",fields, reflect.TypeOf(VcInventory{}), fieldNameMap, validators)
}

func VcsaEmbeddedBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["standalone"] = bindings.NewOptionalType(bindings.NewReferenceType(EmbeddedStandaloneVcsaBindingType))
    fieldNameMap["standalone"] = "Standalone"
    fields["replicated"] = bindings.NewOptionalType(bindings.NewReferenceType(EmbeddedReplicatedVcsaBindingType))
    fieldNameMap["replicated"] = "Replicated"
    fields["ceip_enabled"] = bindings.NewBooleanType()
    fieldNameMap["ceip_enabled"] = "CeipEnabled"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.lcm.vcsa_embedded",fields, reflect.TypeOf(VcsaEmbedded{}), fieldNameMap, validators)
}


