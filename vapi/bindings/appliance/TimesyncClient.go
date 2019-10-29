/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Code generated. DO NOT EDIT.

/*
 * Interface file for service: Timesync
 * Used by client-side stubs.
 */

package appliance


// ``Timesync`` interface provides methods Performs time synchronization configuration.
type TimesyncClient interface {

    // Set time synchronization mode.
    //
    // @param modeParam Time synchronization mode.
    // @throws Error Generic error
	Set(modeParam TimesyncTimeSyncMode) error

    // Get time synchronization mode.
    // @return Time synchronization mode.
    // @throws Error Generic error
	Get() (TimesyncTimeSyncMode, error)
}
