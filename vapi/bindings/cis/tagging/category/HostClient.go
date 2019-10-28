/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Interface file for service: Host
 * Used by client-side stubs.
 */
package category


// The ``Host`` interface provides methods to list categories that have hosts as an attachable type. **Warning:** This interface is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
type HostClient interface {

    // Returns information about the categories that have hosts as an attachable type. **Warning:** This method is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    // @return The list of categories that have hosts as an attachable type.
	List() ([]HostSummary, error)
}
