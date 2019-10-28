/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Interface file for service: EnhancedSecurity
 * Used by client-side stubs.
 */
package security


// The ``EnhancedSecurity`` interface provides methods to enable/disable enhanced security in the appliance for specific environment (GovCloud).
type EnhancedSecurityClient interface {

    // Enable/Disable advanced security (IL4/IL5).
    //
    // @param enabledParam 
    // @throws Error Generic error
	Set(enabledParam bool) error
}
