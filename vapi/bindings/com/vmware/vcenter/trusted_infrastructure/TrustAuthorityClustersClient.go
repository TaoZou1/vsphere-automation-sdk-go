/* Copyright © 2019 VMware, Inc. All Rights Reserved.
     SPDX-License-Identifier: BSD-2-Clause */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Interface file for service: TrustAuthorityClusters
 * Used by client-side stubs.
 */

package trusted_infrastructure

import (
)

// The ``TrustAuthorityClusters`` interface manages all the Trust Authority Components on each Trust Authority Host in the cluster. The ``TrustAuthorityClusters`` interface transforms a ClusterComputeResource into Trust Authority Cluster and vice versa. This interface was added in vSphere API 7.0.
type TrustAuthorityClustersClient interface {


    // Updates the state of a cluster. This method was added in vSphere API 7.0.
    //
    // @param clusterParam Cluster id.
    // The parameter must be an identifier for the resource type: ``ClusterComputeResource``.
    // @param specParam The specification for update of a cluster.
    // @throws Error if there is a generic error.
    // @throws NotFound if ``spec`` doesn't match to any cluster compute resource.
    // @throws Unauthenticated if the user can not be authenticated.
    Update(clusterParam string, specParam TrustAuthorityClustersUpdateSpec) error 


    // Get the result of the last Update operation which matches the cluster id. This method was added in vSphere API 7.0.
    //
    // @param clusterParam Cluster id.
    // The parameter must be an identifier for the resource type: ``ClusterComputeResource``.
    // @return The TrustAuthorityClustersInfo instance which contains information about the state of the cluster.
    // @throws Error if there is a generic error.
    // @throws NotFound if ``cluster`` doesn't match to any ClusterComputeResource.
    // @throws Unauthenticated if the user can not be authenticated.
    Get(clusterParam string) (TrustAuthorityClustersInfo, error) 


    // Returns a list of clusters for this vCenter instance which matches the TrustAuthorityClustersFilterSpec. This method was added in vSphere API 7.0.
    //
    // @param specParam Return only clusters matching the specified filters.
    // If {\\\\@term.unset} return all clusters.
    // @return List of TrustAuthorityClustersSummary for a TrustAuthorityClusters.
    // @throws Error if there is a generic error.
    // @throws Unauthenticated if the user can not be authenticated.
    List(specParam *TrustAuthorityClustersFilterSpec) ([]TrustAuthorityClustersSummary, error) 

}
