/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Code generated. DO NOT EDIT.

/*
 * Interface file for service: VcIdentity
 * Used by client-side stubs.
 */

package identity


// The ``VcIdentity`` interface provides methods to read and modify local vCenter identity. **Warning:** This interface is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
type VcIdentityClient interface {

    // Retrieve VcIdentity for local domain. **Warning:** This method is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    // @return VcIdentity.
    // @throws Unauthorized if authorization is not given to caller.
	Get() (VcIdentityInfo, error)

    // Update the VcIdentity for local domain with data in given UpdateSpec. **Warning:** This method is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    //
    // @param specParam VcIdentity update spec which contains the data to be updated.
    // @throws Unauthorized if authorization is not given to caller.
    // @throws InvalidArgument if no arguments are provided in update spec.
	Update(specParam VcIdentityUpdateSpec) error
}
