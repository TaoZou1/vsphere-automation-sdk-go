/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Interface file for service: ClusterCompatibility
 * Used by client-side stubs.
 */
package namespace_management


// The ``ClusterCompatibility`` interface provides methods to get Namespace-related compatibility information for clusters in this vCenter. **Warning:** This interface is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
type ClusterCompatibilityClient interface {

    // Returns Namespaces compatibility information for all clusters in vCenter matching the ClusterCompatibilityFilterSpec. The result contains only visible (subject to permission checks) clusters. **Warning:** This method is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    //
    // @param filterParam Specification of matching clusters for which information should be returned.
    // If null, the behavior is equivalent to a ClusterCompatibilityFilterSpec with all properties null which means all clusters match the filter.
    // @return Namespaces compatibility information for the clusters matching the the ClusterCompatibilityFilterSpec.
	List(filterParam *ClusterCompatibilityFilterSpec) ([]ClusterCompatibilitySummary, error)
}
