/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Interface file for service: Registries
 * Used by client-side stubs.
 */
package registries


// The ``Registries`` interface provides methods to list all container registries available in the vCenter. **Warning:** This interface is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
type RegistriesClient interface {

    // Returns basic information of all container registries in the vCenter. **Warning:** This method is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    // @return The list of basic information of all container registries.
    // @throws Error if the system reports an error while responding to the request.
    // @throws Unauthenticated if the user can not be authenticated.
    // @throws Unauthorized if the user is not a member of the Administrators
	List() ([]RegistriesSummary, error)
}
