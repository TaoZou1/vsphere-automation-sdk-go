/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Interface file for service: Update
 * Used by client-side stubs.
 */
package system


// ``Update`` interface provides methods Performs update repository configuration.
type UpdateClient interface {

    // Set update repository configuration.
    //
    // @param configParam update related configuration
    // @throws Error Generic error
	Set(configParam UpdateUpdateStructSet) error

    // Get update repository configuration.
    // @return update related configuration
    // @throws Error Generic error
	Get() (UpdateUpdateStructGet, error)
}
