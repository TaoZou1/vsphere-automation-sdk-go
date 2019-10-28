/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Interface file for service: Environment
 * Used by client-side stubs.
 */
package appliance


// The ``Environment`` interface provides methods to get and set appliance environment.
type EnvironmentClient interface {

    // Sets the properties of the appliance environment.
    //
    // @param configParam Structure containing the values of the Environment.
    // @throws Error if any error occurs during the execution of the operation.
	Set(configParam EnvironmentConfig) error

    // Gets the properties of the appliance environment.
    // @return Structure containing the values of the Environment.
    // @throws Error if any error occurs during the execution of the operation.
	Get() (EnvironmentInfo, error)
}
