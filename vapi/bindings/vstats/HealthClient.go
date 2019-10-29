/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Code generated. DO NOT EDIT.

/*
 * Interface file for service: Health
 * Used by client-side stubs.
 */

package vstats


// The ``Health`` interface provides access to vStats health.
type HealthClient interface {

    // Returns information about service health.
    // @return Service health information.
    // @throws Error if the system reports an error while responding to the request.
    // @throws Unauthenticated if the user can not be authenticated.
    // @throws Unauthorized if the user does not have sufficient privileges.
	Get() (HealthInfo, error)
}
