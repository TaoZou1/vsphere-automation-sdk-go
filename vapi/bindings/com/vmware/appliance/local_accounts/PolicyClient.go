/* Copyright © 2019 VMware, Inc. All Rights Reserved.
     SPDX-License-Identifier: BSD-2-Clause */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Interface file for service: Policy
 * Used by client-side stubs.
 */

package local_accounts

import (
)

// The ``Policy`` interface provides methods to manage local user accounts
type PolicyClient interface {


    // Get the global password policy.
    // @return Global password policy
    // @throws Error Generic error
    Get() (PolicyInfo, error) 


    // Set the global password policy.
    //
    // @param policyParam Global password policy
    // @throws InvalidArgument if passed policy values are < -1 or > 99999
    // @throws Error Generic error
    Set(policyParam PolicyInfo) error 

}
