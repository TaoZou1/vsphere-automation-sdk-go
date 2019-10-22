/* Copyright © 2019 VMware, Inc. All Rights Reserved.
     SPDX-License-Identifier: BSD-2-Clause */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Interface file for service: Services
 * Used by client-side stubs.
 */

package attestation

import (
)

// The ``Services`` interface manages endpoint configuration about the attestation service and its relation to the Key Management Server (KMS) services.
type ServicesClient interface {


    // Returns the list of all attestation service instances.
    //
    // @param filterParam Return only services matching the specified filters. If the filter doesn't match, return an empty list.
    // If {\\\\@term.unset} return all services.
    // @return List of all attestation service instances.
    // @throws Error if there is a problem accessing the stored data.
    // @throws Unauthenticated if the user can not be authenticated.
    List(filterParam *ServicesFilterSpec) ([]ServicesSummary, error) 


    // Returns the detailed information about an attestation service instance.
    //
    // @param serviceParam the attestation service instance unique identifier.
    // The parameter must be an identifier for the resource type: ``com.vmware.esx.trusted_infrastructure.attestation.service``.
    // @return Detailed information about the specified attestation service instance.
    // @throws NotFound if there is no attestation service instance with the specified identifier.
    // @throws Error if there is a problem accessing the stored data.
    // @throws Unauthenticated if the user can not be authenticated.
    Get(serviceParam string) (ServicesInfo, error) 


    // Adds a new attestation service instance.
    //
    // @param specParam The ServicesCreateSpec for the new service.
    // @return Identifier of the newly registered attestation service instance.
    // The return value will be an identifier for the resource type: ``com.vmware.esx.trusted_infrastructure.attestation.service``.
    // @throws AlreadyExists if there is already an attestation service instance with the same trusted_infrastructure.NetworkAddress.
    // @throws InvalidArgument if the ServicesCreateSpec contains invalid data.
    // @throws UnableToAllocateResource if the host does not have a required license.
    // @throws Error if there is a problem storing the data.
    // @throws Unauthenticated if the user can not be authenticated.
    Create(specParam ServicesCreateSpec) (string, error) 


    // Removes a currently configured attestation service instance.
    //
    // @param serviceParam the attestation service instance unique identifier.
    // The parameter must be an identifier for the resource type: ``com.vmware.esx.trusted_infrastructure.attestation.service``.
    // @throws NotFound if the service is not found.
    // @throws Error if there is a problem storing the data.
    // @throws Unauthenticated if the user can not be authenticated.
    Delete(serviceParam string) error 

}
