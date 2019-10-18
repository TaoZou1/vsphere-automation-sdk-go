/* Copyright © 2019 VMware, Inc. All Rights Reserved.
     SPDX-License-Identifier: BSD-2-Clause */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Interface file for service: Attestation
 * Used by client-side stubs.
 */

package trust_authority_hosts

import (
)

// The ``Attestation`` interface contains information necessary to connect to the hosts running Attestation Service. This interface was added in vSphere API 7.0.
type AttestationClient interface {


    // Returns the connection info about the Attestation Service running on the specified host. This method was added in vSphere API 7.0.
    //
    // @param hostParam \\\\@{link com.vmware.vcenter.Host} id.
    // The parameter must be an identifier for the resource type: ``HostSystem``.
    // @return The AttestationInfo instance which contains the information necessary to connect to the Attestation Service.
    // @throws Error if service's TLS certificate chain is not valid.
    // @throws NotFound if ``host`` doesn't match to any Host.
    // @throws ResourceInaccessible if connection to ``host`` failed.
    // @throws Unauthenticated if the user can not be authenticated.
    Get(hostParam string) (AttestationInfo, error) 


    // Returns a list of the hosts running a Attestation Service matching the specified AttestationFilterSpec. This method was added in vSphere API 7.0.
    //
    // @param specParam Return details about Attestation Services matching the filter.
    // If {\\\\@term.unset} return all registered Attestation Services.
    // @param projectionParam The type of the returned summary - brief, normal, or full.
    // If {\\\\@term.unset} a normal projection will be used.
    // @return List of AttestationSummary of Attestation Services.
    // @throws Error if there is a generic error.
    // @throws InvalidArgument if the response data will exceed the message limit.
    // @throws Unauthenticated if the user can not be authenticated.
    List(specParam *AttestationFilterSpec, projectionParam *AttestationSummaryType) ([]AttestationSummary, error) 

}
