/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Code generated. DO NOT EDIT.

/*
 * Interface file for service: Staged
 * Used by client-side stubs.
 */

package update


// The ``Staged`` interface provides methods to get the status of the staged update.
type StagedClient interface {

    // Gets the current status of the staged update
    // @return Info structure with information about staged update
    // @throws Error Generic error
    // @throws Unauthenticated session is not authenticated
    // @throws Unauthorized session is not authorized to perform this operation
    // @throws NotAllowedInCurrentState if nothing is staged
	Get() (StagedInfo, error)

    // Deletes the staged update
    // @throws Error Generic error
    // @throws Unauthenticated session is not authenticated
    // @throws Unauthorized session is not authorized to perform this operation
	Delete() error
}
