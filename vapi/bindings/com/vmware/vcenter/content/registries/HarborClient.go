/* Copyright © 2019 VMware, Inc. All Rights Reserved.
     SPDX-License-Identifier: BSD-2-Clause */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Interface file for service: Harbor
 * Used by client-side stubs.
 */

package registries

import (
)

// The ``Harbor`` interface provides methods to manage the lifecycle of an integrated Harbor container registry in vCenter. **Warning:** This interface is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
type HarborClient interface {


    // Creates a Harbor registry in the cluster. **Warning:** This method is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    //
    // @param clientTokenParam A unique token generated on the client for each creation request. The token should be a universally unique identifier (UUID), for example: ``b8a2a2e3-2314-43cd-a871-6ede0f429751``. This token can be used to guarantee idempotent creation.
    // If not specified, creation is not idempotent.
    // @param specParam Specification for creating the Harbor registry.
    // @return Identifier of the deployed registry.
    // The return value will be an identifier for the resource type: ``com.vmware.vcenter.content.Registry``.
    // @throws AlreadyExists if a Harbor already exists in the associated cluster.
    // @throws NotFound if resources/objects could not be located.
    // @throws Unsupported if HarborCreateSpec#cluster does not have vSphere namespace enabled.
    // @throws InvalidArgument if ``spec`` contains any errors.
    // @throws Error if the system reports an error while responding to the request.
    // @throws Unauthenticated if the user can not be authenticated.
    // @throws Unauthorized if the user that requested the method is not authorized to perform the method.
    Create(clientTokenParam *string, specParam HarborCreateSpec) (string, error) 


    // Delete the Harbor registry in the cluster. **Warning:** This method is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    //
    // @param registryParam Identifier of the registry.
    // The parameter must be an identifier for the resource type: ``com.vmware.vcenter.content.Registry``.
    // @throws NotFound if a registry specified by ``registry`` could not be found.
    // @throws Error if the system reports an error while responding to the request.
    // @throws Unauthenticated if the user can not be authenticated.
    // @throws Unauthorized if the user that requested the method is not authorized to perform the method.
    Delete(registryParam string) error 


    // Get detailed information of the Harbor registry. **Warning:** This method is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    //
    // @param registryParam Identifier of the registry.
    // The parameter must be an identifier for the resource type: ``com.vmware.vcenter.content.Registry``.
    // @return Information about the registry.
    // @throws NotFound if a Harbor registry specified by ``registry`` could not be found.
    // @throws Error if the system reports an error while responding to the request.
    // @throws Unauthenticated if the user can not be authenticated.
    // @throws Unauthorized if the user that requested the method is not authorized to perform the method.
    Get(registryParam string) (HarborInfo, error) 


    // Returns basic information of all Harbor registries. **Warning:** This method is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    // @return The list of basic information of all Harbor registries.
    // @throws Error if the system reports an error while responding to the request.
    // @throws Unauthenticated if the user can not be authenticated.
    // @throws Unauthorized if the user is not a member of the Administrators
    List() ([]HarborSummary, error) 

}