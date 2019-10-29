/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Code generated. DO NOT EDIT.

/*
 * Interface file for service: VcTrusts
 * Used by client-side stubs.
 */

package trust


// The ``VcTrusts`` interface provides methods to create, modify, delete and read vCenter trusts. **Warning:** This interface is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
type VcTrustsClient interface {

    // Returns summary information for each trust of the local node. **Warning:** This method is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    // @return List of trust summaries.
    // @throws Unauthorized if authorization is not given to caller.
	List() ([]VcTrustsSummary, error)

    // Creates a new trust for domain in given spec. **Warning:** This method is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    //
    // @param specParam The information needed to create a trust.
    // @throws Unauthorized if authorization is not given to caller.
    // @throws AlreadyExists if trust exists for domain in given spec.
    // @throws InvalidArgument if groupMap key is not in UPN format, if group value is not found or if at least one upn suffix is not given.
	Create(specParam VcTrustsCreateSpec) error

    // Retrieve trust for given domain. **Warning:** This method is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    //
    // @param domainParam Unique identifier for a domain.
    // @return Trust information.
    // @throws Unauthorized if authorization is not given to caller.
    // @throws NotFound if a trust does not exist for given domain.
	Get(domainParam string) (VcTrustsInfo, error)

    // Deletes a trust for given domain. **Warning:** This method is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    //
    // @param domainParam Unique identifier for a domain.
    // @throws Unauthorized if authorization is not given to caller.
    // @throws NotFound if a trust does not exist for given domain.
    // @throws InvalidArgument if the local domain is specified.
	Delete(domainParam string) error

    // Update a trust for given domain with data in given UpdateSpec. **Warning:** This method is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    //
    // @param domainParam Unique identifier for a domain.
    // @param specParam Trust update spec which contains the data to be updated.
    // @throws Unauthorized if authorization is not given to caller.
    // @throws NotFound if a VcTrusts does not exist for given domain.
    // @throws InvalidArgument if no arguments are provided in update spec.
	Update(domainParam string, specParam VcTrustsUpdateSpec) error
}
