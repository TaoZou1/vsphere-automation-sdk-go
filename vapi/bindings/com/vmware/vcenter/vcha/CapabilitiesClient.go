/* Copyright © 2019 VMware, Inc. All Rights Reserved.
     SPDX-License-Identifier: BSD-2-Clause */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Interface file for service: Capabilities
 * Used by client-side stubs.
 */

package vcha

import (
)

// The ``Capabilities`` interface provides methods to get the capabilities of a vCenter High Availability (VCHA) node. This interface was added in vSphere API 6.7.1.
type CapabilitiesClient interface {


    // Gets the capabilities of the active node of a VCHA cluster. This method was added in vSphere API 6.7.1.
    // @return Info structure containing the VCHA capabilities.
    // @throws Unauthorized If the user has insufficient privilege to perform the operation. Operation execution requires the System.Read privilege.
    // @throws Error If any other error occurs.
    Get() (CapabilitiesInfo, error) 

}
