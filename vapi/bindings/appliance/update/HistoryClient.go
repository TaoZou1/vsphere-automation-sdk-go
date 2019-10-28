/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Interface file for service: History
 * Used by client-side stubs.
 */
package update


// The ``History`` interface provides methods to get update history. **Warning:** This interface is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
type HistoryClient interface {

    // Returns information about the list of updates installed. **Warning:** This method is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    // @return List of Summary structure with information about installed updates
    // @throws Error Generic error
    // @throws Unauthenticated session is not authenticated
    // @throws Unauthorized session is not authorized to perform this operation
    // @throws NotAllowedInCurrentState if nothing is installed
	List() ([]HistorySummary, error)

    // Returns detailed information about a particular installed update. **Warning:** This method is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    //
    // @param versionParam 
    // The parameter must be an identifier for the resource type: ``com.vmware.appliance.update.history``.
    // @throws Error Generic error
    // @throws Unauthenticated session is not authenticated
    // @throws Unauthorized session is not authorized to perform this operation
    // @throws NotAllowedInCurrentState if nothing is installed
	Get(versionParam string) (HistoryInfo, error)
}
