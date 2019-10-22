/* Copyright © 2019 VMware, Inc. All Rights Reserved.
     SPDX-License-Identifier: BSD-2-Clause */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Interface file for service: Csr
 * Used by client-side stubs.
 */

package client_certificate

import (
)

// The ``Csr`` interface provides methods to create a certificate signing request(CSR).
type CsrClient interface {


    // Generate a certificate signing request (CSR) for the client certificate. This overwrites any existing CSR. 
    //
    //  The CSR will allow the certificate to be signed by a third party. Once the CSR is signed, certificate may be added as a client certificate. 
    //
    //  Calling the API repeatedly will result in a generating a new CSR each time.
    //
    // @param providerParam Identifier of the provider.
    // The parameter must be an identifier for the resource type: ``com.vmware.esx.kms.providers``.
    // @return A structure containing the CSR.
    // @throws InvalidArgument if the provider Id is empty.
    // @throws NotFound if the provider is not found.
    // @throws Unauthenticated if the caller is not authenticated.
    // @throws Unauthorized if the caller is not authorized.
    // @throws Error if any other error occurs.
    Create(providerParam string) (CsrInfo, error) 


    // Get existing certificate signing request (CSR) for the client certificate and optionally the private key. 
    //
    //  The CSR will allow the certificate to be signed by a third party. Once the CSR is signed, certificate may be added as a client certificate. 
    //
    // @param providerParam Identifier of the provider.
    // The parameter must be an identifier for the resource type: ``com.vmware.esx.kms.providers``.
    // @param specParam Specification to include additional information.
    // If null, the behavior is equivalent to a CsrGetSpec with all properties null
    // @return A structure containing the CSR.
    // @throws InvalidArgument if the provider Id is empty.
    // @throws NotFound if the provider not found.
    // @throws Unauthenticated if the caller is not authenticated.
    // @throws Unauthorized if the caller is not authorized.
    // @throws Error if any other error occurs.
    Get(providerParam string, specParam *CsrGetSpec) (CsrInfo, error) 


    // Set the certificate signing request (CSR) and private key. 
    //
    //  The CSR will allow the certificate to be signed by a third party. Once the CSR is signed, certificate may be added as a client certificate. 
    //
    // @param providerParam Identifier of the provider.
    // The parameter must be an identifier for the resource type: ``com.vmware.esx.kms.providers``.
    // @param specParam Specification describing new CSR and private key.
    // @throws InvalidArgument if the provider Id is empty.
    // @throws NotFound if the provider not found.
    // @throws Unauthenticated if the caller is not authenticated.
    // @throws Unauthorized if the caller is not authorized.
    // @throws Error if any other error occurs.
    Set(providerParam string, specParam CsrSetSpec) error 

}
