/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Interface file for service: Version
 * Used by client-side stubs.
 */
package nsx


// The ``Version`` interface provides information about supported NSX version for vCenter server. **Warning:** This interface is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
type VersionClient interface {

    // Get information about supported NSX version for vCenter Server. **Warning:** This method is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    // @return VersionInfo
    // @throws Unauthenticated if the user can not be authenticated.
    // @throws Unauthorized if the user doesn't have the required privileges.
    // @throws Error if there is some unknown internal error. The accompanying error message will give more details about the error.
	GetVersion() (VersionVersionInfo, error)
}
