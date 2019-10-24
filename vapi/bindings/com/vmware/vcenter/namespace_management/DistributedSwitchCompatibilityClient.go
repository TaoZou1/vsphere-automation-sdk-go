/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Interface file for service: DistributedSwitchCompatibility
 * Used by client-side stubs.
 */
package namespace_management


// The ``DistributedSwitchCompatibility`` interface provides methods to get Namespaces compatibility information of Distributed Switches in this vCenter. **Warning:** This interface is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
type DistributedSwitchCompatibilityClient interface {

    // Returns Namespaces compatibility information of Distributed Switches in vCenter associated with the vCenter cluster, matching the DistributedSwitchCompatibilityFilterSpec. **Warning:** This method is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    //
    // @param clusterParam Identifier of a vCenter Cluster. Only Distributed Switches associated with the vCenter Cluster will be considered by the filter.
    // The parameter must be an identifier for the resource type: ``ClusterComputeResource``.
    // @param filterParam Specification of matching Distributed Switches for which information should be returned.
    // If null, the behavior is equivalent to a DistributedSwitchCompatibilityFilterSpec with all properties null which means all Distributed Switches match the filter.
    // @return Namespaces compatibility information for Distributed Switches matching the the DistributedSwitchCompatibilityFilterSpec.
	List(clusterParam string, filterParam *DistributedSwitchCompatibilityFilterSpec) ([]DistributedSwitchCompatibilitySummary, error)
}
