/* Copyright © 2019 VMware, Inc. All Rights Reserved.
     SPDX-License-Identifier: BSD-2-Clause */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Interface file for service: Accounts
 * Used by client-side stubs.
 */

package vmwportal

import (
)

// The ``Accounts`` interface represents all the operations details of user accounts on my.vmware.com portal. **Warning:** This interface is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
type AccountsClient interface {


    // Gets the list of Entitlement Accounts (EA) attached to the user account on the my.vmware.com portal. **Warning:** This method is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    // @return List of accounts with account name and account ID.
    // @throws ResourceInaccessible if user's my.vmware.com session has expired.
    // @throws Unauthenticated if the user can not be authenticated.
    // @throws Error if there is some unknown internal error. The accompanying error message will give more details about the error.
    List() ([]AccountsSummary, error) 


    // Checks if given user account is entitled to download product binaries from the my.vmware.com portal. Call Accounts::list() to get all the entitlement accounts linked to given my.vmware.com user. **Warning:** This method is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    //
    // @param accountParam ID for which entitlement has to be checked.
    // The parameter must be an identifier for the resource type: ``com.vmware.vcenter.vmwportal.Account``.
    // @param productParam name for which entitlement has to be checked.
    // @throws NotFound if the entitlements are not found.
    // @throws Unauthenticated if the user can not be authenticated.
    // @throws ResourceInaccessible if user's my.vmware.com session has expired.
    // @throws Error if there is some unknown internal error. The accompanying error message will give more details about the error.
    Check(accountParam string, productParam AccountsCheckSpec) error 

}
