/* Copyright © 2019 VMware, Inc. All Rights Reserved.
     SPDX-License-Identifier: BSD-2-Clause */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Interface file for service: Version
 * Used by client-side stubs.
 */

package system

import (
)

// ``Version`` interface provides methods Get the appliance version.
type VersionClient interface {


    // Get the version.
    // @return version information about the appliance
    // @throws Error Generic error
    Get() (VersionVersionStruct, error) 

}
