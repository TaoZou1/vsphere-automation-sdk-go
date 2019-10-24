/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Interface file for service: Members
 * Used by client-side stubs.
 */
package project


// The ``Members`` interface provides methods to manage Harbor project members. **Warning:** This interface is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
type MembersClient interface {

    // Create a project member for specified Harbor project. **Warning:** This method is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    //
    // @param registryParam Identifier of the Registry.
    // The parameter must be an identifier for the resource type: ``com.vmware.vcenter.content.Registry``.
    // @param projectParam Identifier for the project.
    // The parameter must be an identifier for the resource type: ``com.vmware.vcenter.content.Registry.Harbor.Project``.
    // @param specParam Information about the member to be created.
    // @return Identifier of the newly created Harbor project member.
    // The return value will be an identifier for the resource type: ``com.vmware.vcenter.content.Registry.Harbor.Project.Member``.
    // @throws InvalidArgument if ``spec`` contains any error.
    // @throws NotAllowedInCurrentState if Harbor is being deleted.
    // @throws NotFound if ``registry`` or ``project`` cannot be found.
    // @throws AlreadyExists if the specified user or group is already a member in the Harbor project.
    // @throws Error if the system reports an error while responding to the request.
    // @throws Unauthenticated if the user can not be authenticated.
    // @throws Unauthorized if the user that requested the method is not authorized to perform the method.
	Create(registryParam string, projectParam string, specParam MembersCreateSpec) (string, error)

    // Remove a project member from specified Harbor project. **Warning:** This method is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    //
    // @param registryParam Identifier of the Registry.
    // The parameter must be an identifier for the resource type: ``com.vmware.vcenter.content.Registry``.
    // @param projectParam Identifier for the Harbor project.
    // The parameter must be an identifier for the resource type: ``com.vmware.vcenter.content.Registry.Harbor.Project``.
    // @param memberParam The member for this operation.
    // The parameter must be an identifier for the resource type: ``com.vmware.vcenter.content.Registry.Harbor.Project.Member``.
    // @throws NotAllowedInCurrentState if Harbor registry is being deleted.
    // @throws NotFound if ``registry`` or ``project`` or ``domain`` or ``subject`` cannot be found.
    // @throws Error if the system reports an error while responding to the request.
    // @throws Unauthenticated if the user can not be authenticated.
    // @throws Unauthorized if the user that requested the method is not authorized to perform the method.
	Delete(registryParam string, projectParam string, memberParam string) error

    // Update access control for a member in a registry project. **Warning:** This method is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    //
    // @param registryParam Identifier of the Registry.
    // The parameter must be an identifier for the resource type: ``com.vmware.vcenter.content.Registry``.
    // @param projectParam Identifier for the project.
    // The parameter must be an identifier for the resource type: ``com.vmware.vcenter.content.Registry.Harbor.Project``.
    // @param memberParam The member for this operation.
    // The parameter must be an identifier for the resource type: ``com.vmware.vcenter.content.Registry.Harbor.Project.Member``.
    // @param specParam Information about the new member role to be assigned.
    // @throws InvalidArgument if ``spec`` contains any error.
    // @throws NotAllowedInCurrentState if Harbor registry is being deleted.
    // @throws NotFound if ``registry`` or ``project`` or ``domain`` or ``subject`` cannot be found.
    // @throws Error if the system reports an error while responding to the request.
    // @throws Unauthenticated if the user can not be authenticated.
    // @throws Unauthorized if the user that requested the method is not authorized to perform the method.
	Update(registryParam string, projectParam string, memberParam string, specParam MembersUpdateSpec) error

    // Get detailed Harbor project member information for specified user or group. **Warning:** This method is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    //
    // @param registryParam Identifier of the Registry.
    // The parameter must be an identifier for the resource type: ``com.vmware.vcenter.content.Registry``.
    // @param projectParam Identifier for the Harbor project.
    // The parameter must be an identifier for the resource type: ``com.vmware.vcenter.content.Registry.Harbor.Project``.
    // @param memberParam The member for this operation.
    // The parameter must be an identifier for the resource type: ``com.vmware.vcenter.content.Registry.Harbor.Project.Member``.
    // @return Detailed information about the member of the Harbor project.
    // @throws NotFound if ``registry`` or ``project`` or ``domain`` or ``subject`` cannot be found.
    // @throws Error if the system reports an error while responding to the request.
    // @throws Unauthenticated if the user can not be authenticated.
    // @throws Unauthorized if the user that requested the method is not authorized to perform the method.
	Get(registryParam string, projectParam string, memberParam string) (MembersInfo, error)

    // Returns all members basic information in the specified Harbor project. **Warning:** This method is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    //
    // @param registryParam Identifier of the Registry.
    // The parameter must be an identifier for the resource type: ``com.vmware.vcenter.content.Registry``.
    // @param projectParam Identifier for the Harbor project.
    // The parameter must be an identifier for the resource type: ``com.vmware.vcenter.content.Registry.Harbor.Project``.
    // @return The list of information of all project members.
    // @throws Error if the system reports an error while responding to the request.
    // @throws NotFound if ``registry`` or ``project`` cannot be found.
    // @throws Unauthenticated if the user can not be authenticated.
    // @throws Unauthorized if the user is not a member of the Administrators
	List(registryParam string, projectParam string) ([]MembersSummary, error)
}
