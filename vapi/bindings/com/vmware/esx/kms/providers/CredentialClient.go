/* Copyright © 2019 VMware, Inc. All Rights Reserved.
     SPDX-License-Identifier: BSD-2-Clause */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Interface file for service: Credential
 * Used by client-side stubs.
 */

package providers

import (
)

// The ``Credential`` interface provides methods to add a credential for key management server(s).
type CredentialClient interface {


    // Set key management server credential.
    //
    // @param providerParam Identifier of the provider.
    // The parameter must be an identifier for the resource type: ``com.vmware.esx.kms.providers``.
    // @param credentialParam KMIP KMS password or AWS access key.
    // @throws InvalidArgument if the provider Id is empty.
    // @throws NotFound if the provider is not found.
    // @throws Unauthenticated if the caller is not authenticated.
    // @throws Unauthorized if the caller is not authorized.
    // @throws Error if any other error occurs.
    Set(providerParam string, credentialParam string) error 

}
