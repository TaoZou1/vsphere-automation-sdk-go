/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Interface file for service: Configuration
 * Used by client-side stubs.
 */
package vstats


// The ``Configuration`` interface provides methods to manipulate vStats service configuration.
type ConfigurationClient interface {

    // Update vStats service settings.
    //
    // @param updateSpecParam vStats service settings to update.
    // @throws Error if the system reports an error while responding to the request.
    // @throws InvalidArgument if any of the specified parameters are invalid.
    // @throws Unauthenticated if the user can not be authenticated.
    // @throws Unauthorized if the user does not have sufficient privileges.
	Update(updateSpecParam ConfigurationUpdateSpec) error

    // Returns log level information.
    // @return Log level information.
    // @throws Error if the system reports an error while responding to the request.
    // @throws Unauthenticated if the user can not be authenticated.
    // @throws Unauthorized if the user does not have sufficient privileges.
	Get() (ConfigurationInfo, error)
}
