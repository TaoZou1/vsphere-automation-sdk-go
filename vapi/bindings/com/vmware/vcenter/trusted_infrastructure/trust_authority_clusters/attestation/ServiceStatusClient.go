/* Copyright © 2019 VMware, Inc. All Rights Reserved.
     SPDX-License-Identifier: BSD-2-Clause */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Interface file for service: ServiceStatus
 * Used by client-side stubs.
 */

package attestation

import (
)

// The ``ServiceStatus`` interface provides methods to get the Attestation Service health status.
type ServiceStatusClient interface {


    // Return the Attestation service health in the given cluster.
    //
    // @param clusterParam Identifier of the cluster.
    // The parameter must be an identifier for the resource type: ``ClusterComputeResource``.
    // @return The Attestation service health status in the entire cluster.
    // @throws Error For any other error.
    // @throws InvalidArgument If the cluster id is empty.
    // @throws NotFound If the cluster is not found.
    // @throws Unauthenticated If the caller is not authenticated.
    Get(clusterParam string) (ServiceStatusInfo, error) 

}
