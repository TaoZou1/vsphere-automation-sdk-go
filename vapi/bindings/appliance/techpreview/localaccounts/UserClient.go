/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Interface file for service: User
 * Used by client-side stubs.
 */
package localaccounts


// ``User`` interface provides methods Perform operations on local user account.
type UserClient interface {

    // Delete a local user account.
    //
    // @param usernameParam User login name.
    // @throws Error Generic error
	Delete(usernameParam string) error

    // Create a new local user account.
    //
    // @param configParam User configuration.
    // @throws Error Generic error
	Add(configParam UserNewUserConfig) error

    // Update local user account properties role, full name, enabled status and password
    //
    // @param configParam User configuration.
    // @throws Error Generic error
	Set(configParam UserUserConfig) error

    // List of local accounts
    // @return User configuration.
    // @throws Error Generic error
	List() ([]UserUserConfigGet, error)

    // Get the local user account information.
    //
    // @param usernameParam User login name.
    // @return local user account information
    // @throws Error Generic error
	Get(usernameParam string) (UserUserConfigGet, error)
}
