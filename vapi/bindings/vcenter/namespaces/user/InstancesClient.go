/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Code generated. DO NOT EDIT.

/*
 * Interface file for service: Instances
 * Used by client-side stubs.
 */

package user


// The ``Instances`` interface provides methods to access namespaces for non-administrative users. **Warning:** This interface is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
type InstancesClient interface {

    // Returns namespaces that user making the call is authorized to access. **Warning:** This method is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    // @return List of Namespace identifiers together with the API endpoint for each namespace.
    // @throws Error if the system reports an error while responding to the request.
    // @throws Unauthenticated if the user can not be authenticated.
	List() ([]InstancesSummary, error)
}
