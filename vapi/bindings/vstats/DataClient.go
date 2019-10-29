/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Code generated. DO NOT EDIT.

/*
 * Interface file for service: Data
 * Used by client-side stubs.
 */

package vstats


// The ``Data`` interface provides methods to query measurement and statistic data.
type DataClient interface {

    // Returns DataDataPointsResult matching the filter parameters. 
    //
    //  ``/vstats/data/dp?types=VM&types=VCPU`` 
    //
    //  ``/vstats/data/dp?rsrcs=type.HOST=host-16&rsrcs=type.VM=vm-31``
    //
    // @param filterParam Specification to match DataPoints.
    // When map with bool value filtering will be applied to the result.
    // @return Data points matching the filter.
    // @throws Error if the system reports an error while responding to the request.
    // @throws InvalidArgument if any of the specified parameters are invalid.
    // @throws Unauthenticated if the user can not be authenticated.
    // @throws Unauthorized if the user does not have sufficient privileges.
	QueryDataPoints(filterParam *DataFilterSpec) (DataDataPointsResult, error)
}
