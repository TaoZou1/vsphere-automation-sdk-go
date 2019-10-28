/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Interface file for service: VM
 * Used by client-side stubs.
 */
package tags


// The ``VM`` interface provides methods to manage which tags can be associated with virtual machines. **Warning:** This interface is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
type VMClient interface {

    // Returns information about the tags that are associated with virtual machines on this vCenter Server, where the tags need to match VMFilterSpec. **Warning:** This method is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    //
    // @param filterParam The specification of matching tag associations.
    // If null, the behavior is equivalent to a VMFilterSpec with all properties null, which means all tag associations match the filter.
    // @return ListResult with list of tag summaries, see VM.ListResult.
    // @throws Unauthorized if the user doesn't have the required privileges.
	List(filterParam *VMFilterSpec) (VMListResult, error)

    // Deletes any existing associations of this tag with virtual machines. **Warning:** This method is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    //
    // @param tagParam Identifier of the tag of which the associations with virtual machines will be deleted.
    // The parameter must be an identifier for the resource type: ``com.vmware.cis.tagging.Tag:VirtualMachine``.
    // @throws NotFound if a tag with this identifier does not exist.
    // @throws Unauthorized if the user doesn't have the required privileges.
	Delete(tagParam string) error
}
