/* Copyright © 2019 VMware, Inc. All Rights Reserved.
     SPDX-License-Identifier: BSD-2-Clause */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Interface file for service: Timezone
 * Used by client-side stubs.
 */

package time

import (
)

// The ``Timezone`` interface provides methods to get and set the appliance timezone.
type TimezoneClient interface {


    // Set time zone.
    //
    // @param nameParam Time zone name.
    // @throws InvalidArgument if passed arguments are invalid.
    // @throws Error if any error occurs during the execution of the operation.
    Set(nameParam string) error 


    // Get time zone.
    // @return Time zone name.
    // @throws Error if timezone cannot be read.
    Get() (string, error) 

}
