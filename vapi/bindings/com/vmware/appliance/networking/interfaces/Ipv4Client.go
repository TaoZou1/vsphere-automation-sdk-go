/* Copyright © 2019 VMware, Inc. All Rights Reserved.
     SPDX-License-Identifier: BSD-2-Clause */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Interface file for service: Ipv4
 * Used by client-side stubs.
 */

package interfaces

import (
)

// The ``Ipv4`` interface provides methods to perform IPv4 network configuration for interfaces.
type Ipv4Client interface {


    // Set IPv4 network configuration for specific network interface.
    //
    // @param interfaceNameParam Network interface to update, for example, "nic0".
    // The parameter must be an identifier for the resource type: ``com.vmware.appliance.networking.interfaces``.
    // @param configParam The IPv4 Network configuration to set.
    // @throws NotFound if the specified NIC is not found.
    // @throws NotAllowedInCurrentState if the IP is used as PNID
    // @throws ResourceBusy if the specified NIC is busy.
    // @throws Error Generic error.
    Set(interfaceNameParam string, configParam Ipv4Config) error 


    // Get IPv4 network configuration for specific NIC.
    //
    // @param interfaceNameParam The Network interface to query, for example, "nic0".
    // The parameter must be an identifier for the resource type: ``com.vmware.appliance.networking.interfaces``.
    // @return The IPv4 configuration for the queried NIC.
    // @throws NotFound if the specified NIC is not found.
    // @throws Error Generic error.
    Get(interfaceNameParam string) (Ipv4Info, error) 

}
