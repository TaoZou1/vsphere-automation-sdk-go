/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Code generated. DO NOT EDIT.

/*
 * Interface file for service: Certificate
 * Used by client-side stubs.
 */

package attestation


// The ``Certificate`` interface provides methods to get or update the attestation report signing certificate. 
//
//  The result of attestation is a signed report that verifies the attested host's authenticity. This interface allows the report signing certificate to be read by anyone who needs to trust the signature.
type CertificateClient interface {

    // Return the certificate. 
    //
    //  Returns the certificate that is used to sign attestation reports.
    // @return The certificate information.
    // @throws Error if there is a generic error.
    // @throws Unauthenticated if the caller is not authenticated.
    // @throws Unauthorized if the caller is not authorized.
	Get() (CertificateInfo, error)

    // Set the certificate. 
    //
    //  Sets the certificate that is used to sign attestation reports. The existing certificate will be replaced by the one specified. 
    //
    //  This is method is used after generating a CSR, and the specified certificate must match a pending CSR.
    //
    // @param specParam The new certificate.
    // @throws Error if there is a generic error.
    // @throws InvalidArgument if the certificate is invalid.
    // @throws Unauthenticated if the caller is not authenticated.
    // @throws Unauthorized if the caller is not authorized.
	Set(specParam CertificateSetSpec) error
}
