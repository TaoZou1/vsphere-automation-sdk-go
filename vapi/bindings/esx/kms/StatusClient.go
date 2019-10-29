/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Code generated. DO NOT EDIT.

/*
 * Interface file for service: Status
 * Used by client-side stubs.
 */

package kms


// The ``Status`` interface provides methods to get the key management service health status.
type StatusClient interface {

    // Returns the key management service health status.
    // @return The service status information.
    // @throws Unauthenticated if the caller is not authenticated.
    // @throws Unauthorized if the caller is not authorized.
	Get() (StatusInfo, error)
}
