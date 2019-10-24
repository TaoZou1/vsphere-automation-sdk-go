/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Interface file for service: Capabilities
 * Used by client-side stubs.
 */
package vcha


// The ``Capabilities`` interface provides methods to get the capabilities of a vCenter High Availability (VCHA) node.
type CapabilitiesClient interface {

    // Gets the capabilities of the active node of a VCHA cluster.
    // @return Info structure containing the VCHA capabilities.
    // @throws Unauthorized If the user has insufficient privilege to perform the operation. Operation execution requires the System.Read privilege.
    // @throws Error If any other error occurs.
	Get() (CapabilitiesInfo, error)
}
