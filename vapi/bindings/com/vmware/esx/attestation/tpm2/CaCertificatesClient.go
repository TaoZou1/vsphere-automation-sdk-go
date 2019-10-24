/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Interface file for service: CaCertificates
 * Used by client-side stubs.
 */
package tpm2


// The ``CaCertificates`` interface provides methods to manage Trusted Platform Module (TPM) CA certificates. 
//
//  Endorsement Keys are typically packaged in a certificate that is signed by a certificate authority (CA). This interface allows the CA certificate to be registered with the attestation service in order to validate TPM EK certificates when presented at attestation time.
type CaCertificatesClient interface {

    // Return a list of configured TPM CA certificates.
    // @return A list of configured TPM CA certificates.
    // @throws Error if there is a generic error.
    // @throws Unauthenticated if the caller is not authenticated.
    // @throws Unauthorized if the caller is not authorized.
	List() ([]CaCertificatesSummary, error)

    // Add a new TPM CA certificate.
    //
    // @param specParam The new CA certificate details.
    // @throws AlreadyExists if the certificate name exists.
    // @throws Error if there is a generic error.
    // @throws InvalidArgument if the configuration is invalid.
    // @throws Unauthenticated if the caller is not authenticated.
    // @throws Unauthorized if the caller is not authorized.
	Create(specParam CaCertificatesCreateSpec) error

    // Remove a TPM CA certificate.
    //
    // @param nameParam The CA certificate name.
    // The parameter must be an identifier for the resource type: ``com.vmware.esx.attestation.tpm2.ca_certificates``.
    // @throws Error if there is a generic error.
    // @throws InvalidArgument if the name is invalid.
    // @throws NotFound if the name is not found.
    // @throws Unauthenticated if the caller is not authenticated.
    // @throws Unauthorized if the caller is not authorized.
	Delete(nameParam string) error

    // Get the TPM CA certificate details.
    //
    // @param nameParam The CA certificate name.
    // The parameter must be an identifier for the resource type: ``com.vmware.esx.attestation.tpm2.ca_certificates``.
    // @return CA certificate info.
    // @throws Error if there is a generic error.
    // @throws InvalidArgument if the name is invalid.
    // @throws NotFound if the CA certificate is not found.
    // @throws Unauthenticated if the caller is not authenticated.
    // @throws Unauthorized if the caller is not authorized.
	Get(nameParam string) (CaCertificatesInfo, error)
}
