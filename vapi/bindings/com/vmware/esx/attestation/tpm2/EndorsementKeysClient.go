/* Copyright © 2019 VMware, Inc. All Rights Reserved.
     SPDX-License-Identifier: BSD-2-Clause */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Interface file for service: EndorsementKeys
 * Used by client-side stubs.
 */

package tpm2

import (
)

// The ``EndorsementKeys`` interface provides methods to manage Trusted Platform Module (TPM) Endorsement Keys (EK).
type EndorsementKeysClient interface {


    // Return a list of configured TPM endorsement keys.
    // @return A list of configured endorsement keys.
    // @throws Error if there is a generic error.
    // @throws Unauthenticated if the caller is not authenticated.
    // @throws Unauthorized if the caller is not authorized.
    List() ([]EndorsementKeysSummary, error) 


    // Add a new TPM endorsement key.
    //
    // @param specParam The configuration.
    // @throws AlreadyExists if the endorsement key name exists.
    // @throws Error if there is a generic error.
    // @throws InvalidArgument if the configuration is invalid.
    // @throws Unauthenticated if the caller is not authenticated.
    // @throws Unauthorized if the caller is not authorized.
    Create(specParam EndorsementKeysCreateSpec) error 


    // Remove a TPM endorsement key.
    //
    // @param nameParam The endorsement key name.
    // The parameter must be an identifier for the resource type: ``com.vmware.esx.attestation.tpm2.endorsement_keys``.
    // @throws Error if there is a generic error.
    // @throws InvalidArgument if the name is invalid.
    // @throws NotFound if the name is not found.
    // @throws Unauthenticated if the caller is not authenticated.
    // @throws Unauthorized if the caller is not authorized.
    Delete(nameParam string) error 


    // Get the TPM endorsement key details.
    //
    // @param nameParam The endorsement key name.
    // The parameter must be an identifier for the resource type: ``com.vmware.esx.attestation.tpm2.endorsement_keys``.
    // @return The endorsement key info.
    // @throws Error if there is a generic error.
    // @throws InvalidArgument if the name is invalid.
    // @throws NotFound if the endorsement key is not found.
    // @throws Unauthenticated if the caller is not authenticated.
    // @throws Unauthorized if the caller is not authorized.
    Get(nameParam string) (EndorsementKeysInfo, error) 

}
