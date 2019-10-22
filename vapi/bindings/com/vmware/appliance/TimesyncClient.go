/* Copyright © 2019 VMware, Inc. All Rights Reserved.
     SPDX-License-Identifier: BSD-2-Clause */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Interface file for service: Timesync
 * Used by client-side stubs.
 */

package appliance

import (
)

// ``Timesync`` interface provides methods Performs time synchronization configuration. This interface was added in vSphere API 6.7.
type TimesyncClient interface {


    // Set time synchronization mode. This method was added in vSphere API 6.7.
    //
    // @param modeParam Time synchronization mode.
    // @throws Error Generic error
    Set(modeParam Timesync_TimeSyncMode) error 


    // Get time synchronization mode. This method was added in vSphere API 6.7.
    // @return Time synchronization mode.
    // @throws Error Generic error
    Get() (Timesync_TimeSyncMode, error) 

}
