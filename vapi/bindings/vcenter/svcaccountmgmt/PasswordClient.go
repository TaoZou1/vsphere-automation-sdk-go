/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Interface file for service: Password
 * Used by client-side stubs.
 */
package svcaccountmgmt


// The ``Password`` interface provides methods to manage service account passwords. **Warning:** This interface is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
type PasswordClient interface {

    // Change the service account password. **Warning:** This method is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    //
    // @param changeSpecParam 
    // @return OutputSpec The newly generated service account password.
    // @throws NotFound if the service account bearing the input name does not exists.
    // @throws Error if failed due to generic exception.
	Change(changeSpecParam PasswordChangeSpec) (PasswordOutputSpec, error)

    // Reset the service account password. **Warning:** This method is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    //
    // @param resetSpecParam 
    // @return OutputSpec The newly generated service account password.
    // @throws NotFound if the service account bearing the input name does not exist.
    // @throws Error if failed due to generic exception.
	Reset(resetSpecParam PasswordResetSpec) (PasswordOutputSpec, error)
}
