/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Code generated. DO NOT EDIT.

/*
 * Interface file for service: EdgeClusters
 * Used by client-side stubs.
 */

package nsx


// The {\\\\@name Edge-clusters} interface provides methods to set up and tear down NSX Edge capabilities on a vSphere cluster. **Warning:** This interface is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
type EdgeClustersClient interface {

    // Enable NSX Edge capabilities in a vSphere cluster. **Warning:** This method is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    //
    // @param clusterParam Identifier for the compute resource associated with the NSX Edge node(s).
    // The parameter must be an identifier for the resource type: ``ClusterComputeResource``.
    // @param specParam Specification for setting up the NSX Edge node(s) on the vSphere cluster.
    // @throws AlreadyExists if the cluster already has NSX Edge capabilities enabled.
    // @throws InvalidArgument if ``spec`` contain any errors.
    // @throws NotAllowedInCurrentState if all the hosts in the cluster are not transport nodes.
    // @throws NotFound if EdgeClustersManagementNetworkSpec#network, EdgeClustersOverlayNetworkSpec#ipPool, EdgeClustersPlacementSpec#resourcePool, EdgeClustersNodeSpec#datastore or ``cluster`` could not be located.
    // @throws Error if the system reports an error while responding to the request.
	Enable(clusterParam string, specParam EdgeClustersEnableSpec) error
}
