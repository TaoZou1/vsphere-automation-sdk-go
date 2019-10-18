/* Copyright © 2019 VMware, Inc. All Rights Reserved.
     SPDX-License-Identifier: BSD-2-Clause */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Interface file for service: Applmgmt
 * Used by client-side stubs.
 */

package health

import (
)

// ``Applmgmt`` interface provides methods Get health status of applmgmt services.
type ApplmgmtClient interface {


    // Get health status of applmgmt services.
    // @return health status
    // @throws Error Generic error
    Get() (string, error) 

}
