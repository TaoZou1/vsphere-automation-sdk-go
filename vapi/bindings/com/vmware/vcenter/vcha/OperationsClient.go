/* Copyright © 2019 VMware, Inc. All Rights Reserved.
     SPDX-License-Identifier: BSD-2-Clause */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Interface file for service: Operations
 * Used by client-side stubs.
 */

package vcha

import (
)

// The ``Operations`` interface provides methods to figure out the currently active operations and currently disabled operations in a vCenter High Availability (VCHA) cluster.
type OperationsClient interface {


    // Retrieves the current active and disabled operations of a VCHA cluster.
    // @return Info structure containing the current running and disabled operations of a VCHA cluster.
    // @throws Unauthorized If the user has insufficient privilege to perform the operation. Operation execution requires the System.Read privilege.
    // @throws Error If any other error occurs.
    Get() (OperationsInfo, error) 

}
