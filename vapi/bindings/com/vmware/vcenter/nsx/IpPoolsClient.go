/* Copyright © 2019 VMware, Inc. All Rights Reserved.
     SPDX-License-Identifier: BSD-2-Clause */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Interface file for service: IpPools
 * Used by client-side stubs.
 */

package nsx

import (
)

// The ``IpPools`` interface provides methods to read NSX IP pools used on NSX enabled vSphere clusters. **Warning:** This interface is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
type IpPoolsClient interface {


    // Returns the list of the NSX IP pools. **Warning:** This method is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    // @return IpPoolsSummary Summary of NSX IP pools created
    // @throws Error if the system reports an error while responding to the request.
    List() ([]IpPoolsSummary, error) 


    // Read the NSX IP pool. **Warning:** This method is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    //
    // @param ipPoolParam The NSX IP pool ID(Name).
    // The parameter must be an identifier for the resource type: ``com.vmware.nsx.pools.ip_pool``.
    // @return IpPoolsInfo NSX IP pool information for the specified ipPool.
    // @throws NotFound if ``ip_pool`` not found.
    // @throws Error if the system reports an error while responding to the request.
    Get(ipPoolParam string) (IpPoolsInfo, error) 

}
