/* Copyright © 2019 VMware, Inc. All Rights Reserved.
     SPDX-License-Identifier: BSD-2-Clause */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Data type definitions file for package: com.vmware.vcenter.nsx.
 * Includes binding types of a top level structures and enumerations.
 * Shared by client-side stubs and server-side skeletons to ensure type
 * compatibility.
 */

package nsx

import (
    "reflect"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/bindings"
)


// The ``ApplianceSize`` enumeration class defines the virtual appliance size according to the scale of inventory it can manage. **Warning:** This enumeration is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
//
// <p> See {@link com.vmware.vapi.bindings.ApiEnumeration enumerated types description}.

type ApplianceSize string

const (
    // Appliance size of 'tiny', Default vCPUs: 2, Memory: 8GB, Inventory sise to be managed is VMs: 100, Hosts: 10. **Warning:** This constant field is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
     ApplianceSize_TINY ApplianceSize = "TINY"
    // Appliance size of 'small', Default vCPUs: 4, Memory: 16GB, Inventory sise to be managed is VMs: 1000, Hosts: 100. **Warning:** This constant field is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
     ApplianceSize_SMALL ApplianceSize = "SMALL"
    // Appliance size of 'medium', Default vCPUs: 8, Memory: 24GB, Inventory sise to be managed is VMs: 4000 , Hosts: 400. **Warning:** This constant field is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
     ApplianceSize_MEDIUM ApplianceSize = "MEDIUM"
    // Appliance size of 'large', Default vCPUs: 16, Memory: 32GB, Inventory sise to be managed is VMs: 10000, Hosts: 1000. **Warning:** This constant field is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
     ApplianceSize_LARGE ApplianceSize = "LARGE"
    // Appliance size of 'extra large', Default vCPUs: 24, Memory: 48GB, Inventory sise to be managed is VMs: 35000, Hosts: 2000. **Warning:** This constant field is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
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




// The ``IpAllocationMode`` enumeration class defines different address allocation modes. **Warning:** This enumeration is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
//
// <p> See {@link com.vmware.vapi.bindings.ApiEnumeration enumerated types description}.

type IpAllocationMode string

const (
    // The address is automatically assigned by a DHCP server. **Warning:** This constant field is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
     IpAllocationMode_DHCP IpAllocationMode = "DHCP"
    // The address is assigned from an NSX IP pool. **Warning:** This constant field is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
     IpAllocationMode_IP_POOL IpAllocationMode = "IP_POOL"
)

func (i IpAllocationMode) IpAllocationMode() bool {
    switch i {
        case IpAllocationMode_DHCP:
            return true
        case IpAllocationMode_IP_POOL:
            return true
        default:
            return false
    }
}





// The ``ApplianceConfig`` class contains the configuration specifications of an existing or a new appliance. **Warning:** This class is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
type ApplianceConfig struct {
    // The name of the appliance to deploy. Should not collide with the name of already existing VMs on the same Folder. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    ApplianceName string
    // A size descriptor based on the number of virtual machines which will be managed by the vCenter appliance. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    ApplianceSize *ApplianceSize
    // This password is for appliance OS root user Password must conform to the following requirements: 1. At least 8 characters. 2. No more than 20 characters. 3. At least 1 uppercase character. 4. At least 1 lowercase character. 5. At least 1 number. 6. At least 1 special character (e.g., '!', '(', '\\\\@', etc.). 7. Only visible A-Z, a-z, 0-9 and punctuation (spaces are not allowed). **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    RootPassword string
    // A thin provisioned virtual disk has space allocated and zeroed on demand as the space is used. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    Thin bool
    // Hostname of the appliance. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    Hostname string
    // Network details of the appliance. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    Network Network
    // Determines if the appliance deployment will be enrolled in the customer experience improvement program (CEIP). Enrollment means that the appliance will upload usage statistics to VMware. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    Ceip *bool
}


// TODO some error doesn't have ErrorType. Resolve this later VAPI-3009
//


// The ``Connection`` class contains information required to connect to a vCenter server. The connection to the vCenter server always uses the HTTPS protocol. **Warning:** This class is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
type Connection struct {
    // IP Address or DNS of the vCenter. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    Hostname string
    // Port number. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    Port *int64
    // SHA1 hash of the server SSL certificate. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    SslThumbprint *string
    // Username to access the server. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    Username *string
    // Password for the specified user. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    Password *string
}


// TODO some error doesn't have ErrorType. Resolve this later VAPI-3009
//


// The ``IpPoolCreateSpec`` class contains the specification to create an NSX IP pool. **Warning:** This class is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
type IpPoolCreateSpec struct {
    // The ID(Name) of the NSX IP pool. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    Name string
    // The CIDR string denoting the range of IP addresses and the prefix length. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    Cidr *Ipv4Cidr
    // The IPv4 address range. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    IpRange *Ipv4Range
    // The IPv4 address of the gateway associated with the NSX IP pool. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    Gateway string
}


// TODO some error doesn't have ErrorType. Resolve this later VAPI-3009
//


// The ``Ipv4Cidr`` class contains the specification for representing CIDR notation of IP range. For example, this can be used to represent 256 IP addresses using 10.10.10.0/24. **Warning:** This class is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
type Ipv4Cidr struct {
    // The IPv4 address. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    Address string
    // The prefix length in the IPv4 CIDR notation. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    Prefix int64
}


// TODO some error doesn't have ErrorType. Resolve this later VAPI-3009
//


// The ``Ipv4Range`` contains the specification to configure multiple interfaces in IPV4. The range of IPv4 addresses is derived by incrementing the startingAddress to the specified addressCount. **Warning:** This class is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
type Ipv4Range struct {
    // The IPv4 address denoting the start of the range. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    StartingAddress string
    // A positive number denoting the number of IP addresses in the range. Addresses are derived by incrementing Ipv4Range#startingAddress. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    AddressCount int64
    // The subnet mask to be set. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    SubnetMask *string
}


// TODO some error doesn't have ErrorType. Resolve this later VAPI-3009
//


// The ``Network`` class contains IP information used to configure a network interface. **Warning:** This class is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
type Network struct {
    // Network IP address. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    Ip string
    // A comma-separated list of IP addresses of DNS servers. A JSON array such as ["1.2.3.4", "127.0.0.1"]. DNS servers must be reachable. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    DnsServers []string
    // The subnet mask of the network. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    Subnet string
    // Gateway of the network. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    Gateway string
}


// TODO some error doesn't have ErrorType. Resolve this later VAPI-3009
//


// The ``PlacementDetails`` class contains information to describe the inventory placement of an appliance. **Warning:** This class is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
type PlacementDetails struct {
    // The identifier of the cluster of which the appliance is a member. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    Cluster string
    // The name of the cluster of which the appliance is a member. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    ClusterName string
    // The identifier of the Network object used to attach the appliance. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    Network string
    // The name of the Network object used to attach the appliance. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    NetworkName string
    // The identifier of the datastore to put all the virtual disks on. The datastore must be accessible from the management vCenter. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    Datastore string
    // The name of the datastore. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    DatastoreName string
}


// TODO some error doesn't have ErrorType. Resolve this later VAPI-3009
//






func ApplianceConfigBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["appliance_name"] = bindings.NewStringType()
    fieldNameMap["appliance_name"] = "ApplianceName"
    fields["appliance_size"] = bindings.NewOptionalType(bindings.NewEnumType("com.vmware.vcenter.nsx.appliance_size", reflect.TypeOf(ApplianceSize(ApplianceSize_TINY))))
    fieldNameMap["appliance_size"] = "ApplianceSize"
    fields["root_password"] = bindings.NewSecretType()
    fieldNameMap["root_password"] = "RootPassword"
    fields["thin"] = bindings.NewBooleanType()
    fieldNameMap["thin"] = "Thin"
    fields["hostname"] = bindings.NewStringType()
    fieldNameMap["hostname"] = "Hostname"
    fields["network"] = bindings.NewReferenceType(NetworkBindingType)
    fieldNameMap["network"] = "Network"
    fields["ceip"] = bindings.NewOptionalType(bindings.NewBooleanType())
    fieldNameMap["ceip"] = "Ceip"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.nsx.appliance_config",fields, reflect.TypeOf(ApplianceConfig{}), fieldNameMap, validators)
}

func ConnectionBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["hostname"] = bindings.NewStringType()
    fieldNameMap["hostname"] = "Hostname"
    fields["port"] = bindings.NewOptionalType(bindings.NewIntegerType())
    fieldNameMap["port"] = "Port"
    fields["ssl_thumbprint"] = bindings.NewOptionalType(bindings.NewStringType())
    fieldNameMap["ssl_thumbprint"] = "SslThumbprint"
    fields["username"] = bindings.NewOptionalType(bindings.NewStringType())
    fieldNameMap["username"] = "Username"
    fields["password"] = bindings.NewOptionalType(bindings.NewSecretType())
    fieldNameMap["password"] = "Password"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.nsx.connection",fields, reflect.TypeOf(Connection{}), fieldNameMap, validators)
}

func IpPoolCreateSpecBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["name"] = bindings.NewIdType([]string {"com.vmware.nsx.pools.ip_pool"}, "")
    fieldNameMap["name"] = "Name"
    fields["cidr"] = bindings.NewOptionalType(bindings.NewReferenceType(Ipv4CidrBindingType))
    fieldNameMap["cidr"] = "Cidr"
    fields["ip_range"] = bindings.NewOptionalType(bindings.NewReferenceType(Ipv4RangeBindingType))
    fieldNameMap["ip_range"] = "IpRange"
    fields["gateway"] = bindings.NewStringType()
    fieldNameMap["gateway"] = "Gateway"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.nsx.ip_pool_create_spec",fields, reflect.TypeOf(IpPoolCreateSpec{}), fieldNameMap, validators)
}

func Ipv4CidrBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["address"] = bindings.NewStringType()
    fieldNameMap["address"] = "Address"
    fields["prefix"] = bindings.NewIntegerType()
    fieldNameMap["prefix"] = "Prefix"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.nsx.ipv4_cidr",fields, reflect.TypeOf(Ipv4Cidr{}), fieldNameMap, validators)
}

func Ipv4RangeBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["starting_address"] = bindings.NewStringType()
    fieldNameMap["starting_address"] = "StartingAddress"
    fields["address_count"] = bindings.NewIntegerType()
    fieldNameMap["address_count"] = "AddressCount"
    fields["subnet_mask"] = bindings.NewOptionalType(bindings.NewStringType())
    fieldNameMap["subnet_mask"] = "SubnetMask"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.nsx.ipv4_range",fields, reflect.TypeOf(Ipv4Range{}), fieldNameMap, validators)
}

func NetworkBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["ip"] = bindings.NewStringType()
    fieldNameMap["ip"] = "Ip"
    fields["dns_servers"] = bindings.NewListType(bindings.NewStringType(), reflect.TypeOf([]string{}))
    fieldNameMap["dns_servers"] = "DnsServers"
    fields["subnet"] = bindings.NewStringType()
    fieldNameMap["subnet"] = "Subnet"
    fields["gateway"] = bindings.NewStringType()
    fieldNameMap["gateway"] = "Gateway"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.nsx.network",fields, reflect.TypeOf(Network{}), fieldNameMap, validators)
}

func PlacementDetailsBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["cluster"] = bindings.NewIdType([]string {"ClusterComputeResource:VCenter"}, "")
    fieldNameMap["cluster"] = "Cluster"
    fields["cluster_name"] = bindings.NewStringType()
    fieldNameMap["cluster_name"] = "ClusterName"
    fields["network"] = bindings.NewIdType([]string {"Network:VCenter"}, "")
    fieldNameMap["network"] = "Network"
    fields["network_name"] = bindings.NewStringType()
    fieldNameMap["network_name"] = "NetworkName"
    fields["datastore"] = bindings.NewIdType([]string {"Datastore:VCenter"}, "")
    fieldNameMap["datastore"] = "Datastore"
    fields["datastore_name"] = bindings.NewStringType()
    fieldNameMap["datastore_name"] = "DatastoreName"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.nsx.placement_details",fields, reflect.TypeOf(PlacementDetails{}), fieldNameMap, validators)
}


