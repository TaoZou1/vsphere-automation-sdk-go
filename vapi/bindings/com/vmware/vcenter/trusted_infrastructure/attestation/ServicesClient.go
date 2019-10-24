/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Interface file for service: Services
 * Used by client-side stubs.
 */
package attestation


// The ``Services`` interface contains information about the registered instances of the Attestation Service.
type ServicesClient interface {

    // Returns the list of all ``Services`` instances for this vCenter.
    //
    // @param specParam Return only services matching the specified filters.
    // If {\\\\@term.unset} return all services.
    // @return List of all ``Services`` instances for this vCenter.
    // @throws Error if an error occurred while getting the data.
    // @throws Unauthenticated if the user can not be authenticated.
	List(specParam *ServicesFilterSpec) ([]ServicesSummary, error)

    // Returns the detailed information about an ``Services`` instance for this vCenter.
    //
    // @param serviceParam the ``Services`` instance unique identifier.
    // The parameter must be an identifier for the resource type: ``com.vmware.vcenter.trusted_platform.attestation.Service``.
    // @return Detailed information about the specified ``Services`` instance.
    // @throws Error if an error occurred while getting the data.
    // @throws NotFound if there is no ``Services`` instance with the specified ID.
    // @throws Unauthenticated if the user can not be authenticated.
	Get(serviceParam string) (ServicesInfo, error)

    // Adds a new ``Services`` instance to this vCenter.
    //
    // @param specParam The CreateSpec for the new service.
    // @return ID of the newly registered Attestation Service instance.
    // The return value will be an identifier for the resource type: ``com.vmware.vcenter.trusted_platform.attestation.Service``.
    // @throws AlreadyExists if there is already a ``Services`` instance with the same Address.
    // @throws Error if there is a generic error.
    // @throws InvalidArgument if the CreateSpec contains invalid data.
    // @throws Unauthenticated if the user can not be authenticated.
	Create(specParam ServicesCreateSpec) (string, error)

    // Removes a currently configured ``Services`` instance from this vCenter.
    //
    // @param serviceParam the ``Services`` instance unique identifier.
    // The parameter must be an identifier for the resource type: ``com.vmware.vcenter.trusted_platform.attestation.Service``.
    // @throws Error if an error occurred while deleting the service.
    // @throws Error if there is a generic error.
    // @throws NotFound if the ``Services`` instance is not found.
    // @throws ResourceBusy if the ``Services`` instance is used by a configuration on a cluster level.
    // @throws Unauthenticated if the user can not be authenticated.
	Delete(serviceParam string) error
}
