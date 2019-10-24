/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Interface file for service: Modules
 * Used by client-side stubs.
 */
package cluster


// The ``Modules`` interface provides methods to manage cluster modules. A cluster module contains a subset of the objects of a single vCenter cluster. A cluster module forms a scope on which compute-policies that relate objects operate; e.g., instead of considering all virtual machines in the cluster, the policy only considers all virtual machines in a cluster module. **Warning:** This interface is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
type ModulesClient interface {

    // Creates a new module in a vCenter cluster. **Warning:** This method is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    //
    // @param specParam Specification for the new module to be created.
    // @return The identifier of the newly created module. Use this identifier to get or destroy the module.
    // The return value will be an identifier for the resource type: ``com.vmware.vcenter.cluster.modules``.
    // @throws InvalidArgument if a parameter passed in the spec is invalid.
	Create(specParam ModulesCreateSpec) (string, error)

    // Returns information about the modules available in this vCenter server. **Warning:** This method is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    // @return The list of modules available on this vCenter server.
	List() (ModulesListResult, error)

    // Deletes a specific module. **Warning:** This method is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    //
    // @param moduleParam Identifier of the module to be deleted.
    // The parameter must be an identifier for the resource type: ``com.vmware.vcenter.cluster.modules``.
    // @throws NotFound if a module with this identifier does not exist.
	Delete(moduleParam string) error
}
