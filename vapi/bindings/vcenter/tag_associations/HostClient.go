/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Code generated. DO NOT EDIT.

/*
 * Interface file for service: Host
 * Used by client-side stubs.
 */

package tag_associations


// The ``Host`` interface provides methods to query tag associations that have a host identifier. **Warning:** This interface is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
type HostClient interface {

    // Returns tag associations that match the specified filter. **Warning:** This method is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    //
    // @param iterateParam The specification of a page to be retrieved.
    // If null, the first page will be retrieved.
    // @param filterParam The specification of matching tag associations.
    // If null, the behavior is equivalent to a HostFilterSpec with all properties null, which means all tag associations match the filter.
    // @return The tag associations matching the HostFilterSpec.
    // @throws InvalidArgument if HostIterationSpec#marker is not a marker returned from an earlier invocation of this {\\\\@term operation).
    // @throws Unauthorized if the user doesn't have the required privileges.
	List(iterateParam *HostIterationSpec, filterParam *HostFilterSpec) (HostListResult, error)
}
