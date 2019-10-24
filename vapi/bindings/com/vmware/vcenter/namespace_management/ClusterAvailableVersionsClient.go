/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Interface file for service: ClusterAvailableVersions
 * Used by client-side stubs.
 */
package namespace_management


// The {\\\\@name cluster-available-versions} interface provides methods to retrieve available upgrade versions of WCP and detailed information about each upgrade. **Warning:** This interface is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
type ClusterAvailableVersionsClient interface {

    // Get information about each available upgrade. **Warning:** This method is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    // @return Information for each upgrade.
    // @throws Error if the system reports an error while responding to the request.
    // @throws Unauthenticated if the user can not be authenticated.
    // @throws Unauthorized if the user does not have System.Read privilege.
	List() ([]ClusterAvailableVersionsSummary, error)
}
