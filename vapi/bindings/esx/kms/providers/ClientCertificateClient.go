/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Code generated. DO NOT EDIT.

/*
 * Interface file for service: ClientCertificate
 * Used by client-side stubs.
 */

package providers


// The ``ClientCertificate`` interface provides methods to add and retrieve client certificate.
type ClientCertificateClient interface {

    // Generate a new self signed client certificate. Existing client certificate is overwritten. The key management server will use this certificate to validate the client connection.
    //
    // @param providerParam Identifier of the provider.
    // The parameter must be an identifier for the resource type: ``com.vmware.esx.kms.providers``.
    // @throws InvalidArgument if the provider Id is empty.
    // @throws NotFound if the provider is not found.
    // @throws Unauthenticated if the caller is not authenticated.
    // @throws Unauthorized if the caller is not authorized.
    // @throws Error if any other error occurs.
	Create(providerParam string) error

    // Return the existing client certificate.
    //
    // @param providerParam Identifier of the provider.
    // The parameter must be an identifier for the resource type: ``com.vmware.esx.kms.providers``.
    // @param specParam Specification to include additional information.
    // If null, the behavior is equivalent to a ClientCertificateGetSpec with all properties null
    // @return A structure containing the certificate.
    // @throws InvalidArgument if the provider Id is empty.
    // @throws NotFound if the provider is not found.
    // @throws Unauthenticated if the caller is not authenticated.
    // @throws Unauthorized if the caller is not authorized.
    // @throws Error if any other error occurs.
	Get(providerParam string, specParam *ClientCertificateGetSpec) (ClientCertificateInfo, error)

    // Set the client certificate. 
    //
    //  The key management server will use this certificate to validate the client connection. If a client certificate already exists, it will be replaced. 
    //
    //  An optional private key can be specified if the certificate has already been provisioned.
    //
    // @param providerParam Identifier of the provider.
    // The parameter must be an identifier for the resource type: ``com.vmware.esx.kms.providers``.
    // @param specParam Specification describing new certifiate and private key.
    // @throws InvalidArgument if the certificate or private key is invalid or provider Id is empty.
    // @throws NotFound if the provider is not found.
    // @throws Unauthenticated if the caller is not authenticated.
    // @throws Unauthorized if the caller is not authorized.
    // @throws Error if any other error occurs.
	Set(providerParam string, specParam ClientCertificateSetSpec) error
}
