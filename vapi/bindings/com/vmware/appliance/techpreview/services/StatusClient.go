/* Copyright © 2019 VMware, Inc. All Rights Reserved.
     SPDX-License-Identifier: BSD-2-Clause */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Interface file for service: Status
 * Used by client-side stubs.
 */

package services

import (
)

// ``Status`` interface provides methods Get status of a service.
type StatusClient interface {


    // Get status of a service.
    //
    // @param nameParam Name of a service.
    // @param timeoutParam Timeout in seconds. Zero (0) means no timeout.
    // @return Status of the service.
    // @throws Error Generic error
    Get(nameParam string, timeoutParam int64) (StatusServiceStatus, error) 

}
