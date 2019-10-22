/* Copyright © 2019 VMware, Inc. All Rights Reserved.
     SPDX-License-Identifier: BSD-2-Clause */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Interface file for service: ServerCertificates
 * Used by client-side stubs.
 */

package providers

import (
)

// The ``ServerCertificates`` interface provides methods to add and retrieve trusted server certificates. 
//
//  A provider must be configured with a trusted server certificate before performing any key operations.
type ServerCertificatesClient interface {


    // Add trusted server certificate(s). 
    //
    //  The client will use these certificates to validate the server connection. The existing list of trusted certificates will be overwritten. 
    //
    //  The client will not trust the server connection until a server certificate has been set.
    //
    // @param providerParam Identifier of the provider.
    // The parameter must be an identifier for the resource type: ``com.vmware.esx.kms.providers``.
    // @param certificatesParam The server certificates to trust, PEM.
    // @throws InvalidArgument if one or more certificates are invalid or the provider Id is empty.
    // @throws NotFound if the provider is not found.
    // @throws Unauthenticated if the caller is not authenticated.
    // @throws Unauthorized if the caller is not authorized.
    // @throws Unsupported if the certificate count reaches the max limit.
    // @throws Error if any other error occurs.
    Set(providerParam string, certificatesParam []string) error 


    // Return trusted server certificates.
    //
    // @param providerParam Identifier of the provider.
    // The parameter must be an identifier for the resource type: ``com.vmware.esx.kms.providers``.
    // @return A structure containing Server certificates.
    // @throws InvalidArgument if the provider Id is empty.
    // @throws NotFound if the provider is not found.
    // @throws Unauthenticated if the caller is not authenticated.
    // @throws Unauthorized if the caller is not authorized.
    // @throws Error if any other error occurs.
    Get(providerParam string) (ServerCertificatesInfo, error) 

}
