/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Interface file for service: ServiceAccount
 * Used by client-side stubs.
 */
package svcaccountmgmt


// The ``ServiceAccount`` interface provides methods to manage service accounts. **Warning:** This interface is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
type ServiceAccountClient interface {

    // Create a service account with the input configuration parameters. **Warning:** This method is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    //
    // @param createSpecParam 
    // @return OutputSpec The newly created service account information.
    // @throws AlreadyExists if a service account with the input name already exists.
    // @throws Error if failed due to generic exception.
	Create(createSpecParam ServiceAccountCreateSpec) (ServiceAccountOutputSpec, error)

    // Delete the service account with the input name. **Warning:** This method is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    //
    // @param deleteSpecParam 
    // @throws NotFound if the service account bearing the input name does not exist.
    // @throws Error if failed due to generic exception.
	Delete(deleteSpecParam ServiceAccountDeleteSpec) error
}
