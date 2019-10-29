/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Code generated. DO NOT EDIT.

/*
 * Interface file for service: Ipv6
 * Used by client-side stubs.
 */

package networking


// ``Ipv6`` interface provides methods Performs IPV4 network configuration for interfaces.
type Ipv6Client interface {

    // Set IPv6 network configuration.
    //
    // @param configParam IPv6 configuration.
    // @throws Error Generic error
	Set(configParam []Ipv6IPv6Config) error

    // Get IPv6 network configuration for all configured interfaces.
    // @return IPv6 configuration for each interface.
    // @throws Error Generic error
	List() ([]Ipv6IPv6ConfigReadOnly, error)

    // Get IPv6 network configuration for interfaces.
    //
    // @param interfacesParam Network interfaces to query, for example, "nic0".
    // @return IPv6 configuration.
    // @throws Error Generic error
	Get(interfacesParam []string) ([]Ipv6IPv6ConfigReadOnly, error)
}
