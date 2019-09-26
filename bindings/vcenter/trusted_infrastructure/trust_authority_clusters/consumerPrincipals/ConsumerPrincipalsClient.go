/*
 * Copyright 2019 VMware, Inc.  All rights reserved.
 */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Interface file for service: ConsumerPrincipals
 * Used by client-side stubs.
 */

package consumerPrincipals

import (
)

// The ``ConsumerPrincipals`` interface configures the token policies and STS trust necessary for the workload vCenter to query the trusted services for their status.
type ConsumerPrincipalsClient interface {


    // Creates a profile with the specified connection information on all hosts from a Trust Authority Cluster.
    //
    // @param clusterParam The ID of the Trust Authority Cluster to configure.
    // The parameter must be an identifier for the resource type: ``ClusterComputeResource``.
    // @param specParam The CreateSpec specifying the connection information.
    // @return a unique identifier of the profile
    // The return value will be an identifier for the resource type: ``com.vmware.esx.authentication.clientprofile``.
    // @throws AlreadyExists if a profile for the issuer already exists.
    // @throws Error if there is a generic error.
    // @throws NotFound if there is no such cluster.
    // @throws Unauthenticated if the user can not be authenticated.
    Create(clusterParam string, specParam CreateSpec) (string, error) 


    // Removes the read-only policy configured on ESX for a specific principal.
    //
    // @param clusterParam The ID of the Trust Authority Cluster to configure.
    // The parameter must be an identifier for the resource type: ``ClusterComputeResource``.
    // @param profileParam The ID of the connection profile to modify.
    // The parameter must be an identifier for the resource type: ``com.vmware.esx.authentication.clientprofile``.
    // @throws Error if there is a generic error.
    // @throws NotFound if there is no profile configured with that ID.
    // @throws Unauthenticated if the user can not be authenticated.
    Delete(clusterParam string, profileParam string) error 


    // Retrieve information for a specific profile.
    //
    // @param clusterParam The ID of the Trust Authority Cluster on which the profile is configured.
    // The parameter must be an identifier for the resource type: ``ClusterComputeResource``.
    // @param profileParam The ID of the profile.
    // The parameter must be an identifier for the resource type: ``com.vmware.esx.authentication.clientprofile``.
    // @return information for a specific profile.
    // @throws Error if there is a generic error.
    // @throws NotFound if there is no profile configured with that ID.
    // @throws Unauthenticated if the user can not be authenticated.
    Get(clusterParam string, profileParam string) (Info, error) 


    // Lists all policies configured on a specific cluster.
    //
    // @param clusterParam The ID of the Trust Authority Cluster on which the profile is configured.
    // The parameter must be an identifier for the resource type: ``ClusterComputeResource``.
    // @param specParam A FilterSpec specifying the profiles to be listed.
    // If {\\\\@term.unset} return all policies.
    // @return the list of profiles matching the filter for that cluster.
    // @throws Error if there is a generic error.
    // @throws NotFound if there is no profile configured with that ID.
    // @throws Unauthenticated if the user can not be authenticated.
    List(clusterParam string, specParam *FilterSpec) ([]Summary, error) 

}
