/* Copyright © 2019 VMware, Inc. All Rights Reserved.
     SPDX-License-Identifier: BSD-2-Clause */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Interface file for service: Databasestorage
 * Used by client-side stubs.
 */

package health

import (
)

// ``Databasestorage`` interface provides methods Get database storage health.
type DatabasestorageClient interface {


    // Get database storage health.
    // @return Database storage health
    // @throws Error Generic error
    Get() (Databasestorage_HealthLevel, error) 

}
