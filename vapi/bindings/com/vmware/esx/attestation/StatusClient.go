/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Interface file for service: Status
 * Used by client-side stubs.
 */
package attestation


// The ``Status`` interface provides methods to get the attestation service health status.
type StatusClient interface {

    // Return the servie status. 
    //
    //  Returns the attestation service health status.
    // @return The service status information.
    // @throws Error if there is a generic error.
    // @throws Unauthenticated if the caller is not authenticated.
    // @throws Unauthorized if the caller is not authorized.
	Get() (StatusInfo, error)
}
