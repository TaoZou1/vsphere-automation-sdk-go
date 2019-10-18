/* Copyright © 2019 VMware, Inc. All Rights Reserved.
     SPDX-License-Identifier: BSD-2-Clause */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Interface file for service: TimeSeries
 * Used by client-side stubs.
 */

package stats

import (
)

// The ``TimeSeries`` interface provides methods to gather statistical values for clusters, namespaces and pods. **Warning:** This interface is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
type TimeSeriesClient interface {


    // Gather statistical values for a cluster, namespace, or pod. **Warning:** This method is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    //
    // @param specParam Specification of the statistical values that should be returned.
    // @return A list of TimeSeries values for each counter specified in the request.
    // @throws Error if the system reports an error while responding to the request.
    // @throws InvalidArgument if the start time in TimeSeriesSpec#start is invalid, or the end time in TimeSeriesSpec#end is invalid.
    // @throws NotFound if the specified cluster in TimeSeriesSpec#cluster or the namespace in TimeSeriesSpec#namespace or TimeSeriesSpec#pod does not exist.
    // @throws Unsupported if the specified cluster in TimeSeriesSpec#cluster is not enabled for Namespaces.
    // @throws Unauthenticated if the user can not be authenticated.
    // @throws Unauthorized if the user does not have System.Read privilege.
    Get(specParam TimeSeriesSpec) ([]TimeSeriesTimeSeries, error) 

}
