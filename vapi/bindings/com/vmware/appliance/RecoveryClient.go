/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Interface file for service: Recovery
 * Used by client-side stubs.
 */
package appliance


// The ``Recovery`` interface provides methods to invoke an appliance recovery (backup and restore).
type RecoveryClient interface {

    // Gets the properties of the appliance Recovery subsystem.
    // @return Structure containing the properties of the Recovery subsystem.
    // @throws Error if any error occurs during the execution of the operation.
	Get() (RecoveryInfo, error)
}
