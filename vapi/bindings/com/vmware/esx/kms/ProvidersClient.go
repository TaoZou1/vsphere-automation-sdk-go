/* Copyright © 2019 VMware, Inc. All Rights Reserved.
     SPDX-License-Identifier: BSD-2-Clause */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Interface file for service: Providers
 * Used by client-side stubs.
 */

package kms

import (
)

// The ``Providers`` interface provides methods to create, update and delete Key Providers.
type ProvidersClient interface {


    // Return a list of Key Providers.
    // @return List of providers.
    // @throws Unauthenticated if the caller is not authenticated.
    // @throws Unauthorized if the caller is not authorized.
    // @throws Error if any other error occurs.
    List() ([]ProvidersSummary, error) 


    // Add a new Key Provider.
    //
    // @param specParam Provider information.
    // @throws AlreadyExists if the provider already exists.
    // @throws InvalidArgument if the spec is invalid.
    // @throws Unauthenticated if the caller is not authenticated.
    // @throws Unauthorized if the caller is not authorized.
    // @throws Error if any other error occurs.
    Create(specParam ProvidersCreateSpec) error 


    // Update an existing Key Provider.
    //
    // @param providerParam Identifier of the provider.
    // The parameter must be an identifier for the resource type: ``com.vmware.esx.kms.providers``.
    // @param specParam Provider information.
    // @throws InvalidArgument if the spec is invalid.
    // @throws NotFound if the provider is not found.
    // @throws Unauthenticated if the caller is not authenticated.
    // @throws Unauthorized if the caller is not authorized.
    // @throws Error if any other error occurs.
    Update(providerParam string, specParam ProvidersUpdateSpec) error 


    // Remove a Key Provider.
    //
    // @param providerParam Identifier of the provider.
    // The parameter must be an identifier for the resource type: ``com.vmware.esx.kms.providers``.
    // @throws InvalidArgument if the provider Id is empty.
    // @throws NotFound if the provider is not found.
    // @throws Unauthenticated if the caller is not authenticated.
    // @throws Unauthorized if the caller is not authorized.
    // @throws Error if any other error occurs.
    Delete(providerParam string) error 


    // Return information about a Key Provider.
    //
    // @param providerParam Identifier of the provider.
    // The parameter must be an identifier for the resource type: ``com.vmware.esx.kms.providers``.
    // @return Provider information.
    // @throws InvalidArgument if the provider Id is empty.
    // @throws NotFound if the provider is not found.
    // @throws Unauthenticated if the caller is not authenticated.
    // @throws Unauthorized if the caller is not authorized.
    // @throws Error if any other error occurs.
    Get(providerParam string) (ProvidersInfo, error) 

}
