/*
 * Copyright 2019 VMware, Inc.  All rights reserved.
 */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Interface file for service: Cluster
 * Used by client-side stubs.
 */

package cluster

import (
)

// The ``Cluster`` interface provides methods to manage clusters in the vCenter Server.
type ClusterClient interface {


    // Returns information about at most 1000 visible (subject to permission checks) clusters in vCenter matching the FilterSpec.
    //
    // @param filterParam Specification of matching clusters for which information should be returned.
    // If null, the behavior is equivalent to a FilterSpec with all properties null which means all clusters match the filter.
    // @return Commonly used information about the clusters matching the FilterSpec.
    // @throws UnableToAllocateResource if more than 1000 clusters match the FilterSpec.
    // @throws ServiceUnavailable if the system is unable to communicate with a service to complete the request.
    // @throws Unauthenticated if the user can not be authenticated.
    // @throws Unauthorized if the user doesn't have the required privileges.
    List(filterParam *FilterSpec) ([]Summary, error) 


    // Retrieves information about the cluster corresponding to ``cluster``.
    //
    // @param clusterParam Identifier of the cluster.
    // The parameter must be an identifier for the resource type: ``ClusterComputeResource``.
    // @return The Info instances that corresponds to the ``cluster``.
    // @throws NotFound if there is no cluster associated with ``cluster`` in the system.
    // @throws ServiceUnavailable if the system is unable to communicate with a service to complete the request.
    // @throws Unauthenticated if the session id is missing from the request or the corresponding session object cannot be found.
    // @throws Unauthorized if the user doesn't not have the required privileges.
    Get(clusterParam string) (Info, error) 

}
