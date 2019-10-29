/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Code generated. DO NOT EDIT.

/*
 * Interface file for service: Associations
 * Used by client-side stubs.
 */

package tagging


// The ``Associations`` interface provides methods to list tag associations.
type AssociationsClient interface {

    // Returns tag associations that match the specified iteration spec.
    //
    // @param iterateParam The specification of a page to be retrieved.
    // If null, the first page will be retrieved.
    // @return A page of the tag associations matching the iteration spec.
    // @throws InvalidArgument if AssociationsIterationSpec#marker is not a marker returned from an earlier invocation of this {\\\\@term operation).
    // @throws Unauthorized if the user doesn't have the required privileges.
	List(iterateParam *AssociationsIterationSpec) (AssociationsListResult, error)
}
