/*
 * Copyright 2019 VMware, Inc.  All rights reserved.
 */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Interface file for service: ResourcePool
 * Used by client-side stubs.
 */

package resourcePool

import (
)

// The ResourcePool interface provides methods for manipulating a vCenter Server resource pool. 
//
//  This interface does not include virtual appliances in the inventory of resource pools even though part of the behavior of a virtual appliance is to act like a resource pool.
type ResourcePoolClient interface {


    // Retrieves information about the resource pool indicated by ``resource_pool``.
    //
    // @param resourcePoolParam Identifier of the resource pool for which information should be retrieved.
    // The parameter must be an identifier for the resource type: ``ResourcePool``.
    // @return Information about the resource pool.
    // @throws NotFound If the resource pool indicated by ``resource_pool`` does not exist.
    // @throws ServiceUnavailable If the system is unable to communicate with a service to complete the request.
    // @throws Unauthenticated If the user can not be authenticated.
    // @throws Unauthorized If the user doesn't have the required privileges.
    Get(resourcePoolParam string) (Info, error) 


    // Returns information about at most 1000 visible (subject to permission checks) resource pools in vCenter matching the FilterSpec.
    //
    // @param filterParam Specification of matching resource pools for which information should be returned.
    // If null, the behavior is equivalent to a FilterSpec with all properties null which means all resource pools match the filter.
    // @return Commonly used information about the resource pools matching the FilterSpec.
    // @throws UnableToAllocateResource If more than 1000 resource pools match the FilterSpec.
    // @throws ServiceUnavailable If the system is unable to communicate with a service to complete the request.
    // @throws Unauthenticated If the user can not be authenticated.
    // @throws Unauthorized If the user doesn't have the required privileges.
    List(filterParam *FilterSpec) ([]Summary, error) 

}
