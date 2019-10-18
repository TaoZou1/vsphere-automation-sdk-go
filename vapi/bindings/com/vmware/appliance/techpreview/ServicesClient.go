/* Copyright © 2019 VMware, Inc. All Rights Reserved.
     SPDX-License-Identifier: BSD-2-Clause */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Interface file for service: Services
 * Used by client-side stubs.
 */

package techpreview

import (
)

// ``Services`` interface provides methods Manages services.
type ServicesClient interface {


    // Get a list of all known services.
    // @return List of services.
    // @throws Error Generic error
    List() ([]ServicesServiceInfo, error) 


    // Stop a service
    //
    // @param nameParam Name of service.
    // @param timeoutParam Timeout in seconds. Zero (0) means no timeout.
    // @throws Error Generic error
    Stop(nameParam string, timeoutParam int64) error 


    // start or restart a service
    //
    // @param nameParam Name of the service to start or restart.
    // @param timeoutParam Timeout in seconds. Zero (0) means no timeout.
    // @throws Error Generic error
    Restart(nameParam string, timeoutParam int64) error 

}
