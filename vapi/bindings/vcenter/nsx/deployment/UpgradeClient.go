/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Interface file for service: Upgrade
 * Used by client-side stubs.
 */
package deployment


// The ``Upgrade`` interface represents all the operations of NSX Upgrade from vCenter server. **Warning:** This interface is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
type UpgradeClient interface {

    // Gets NSX deployment information, this includes the list of availble clusters their upgrade status, available releases and progress of any running operation. **Warning:** This method is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    // @return Info
    // @throws Unauthenticated if the user can not be authenticated.
    // @throws Unauthorized if the user doesn't have the required privileges.
    // @throws NotAllowedInCurrentState if the operation is not allowed in the current state.
    // @throws Error if there is some unknown internal error. The accompanying error message will give more details about the error.
	Get() (UpgradeInfo, error)

    // Runs checks for available NSX upgrade. This is an asynchronous call with fast return. The result of the operation can be queried by calling the API /vcenter/nsx/deployment/upgrade. **Warning:** This method is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    // @throws Unauthenticated if the user can not be authenticated.
    // @throws Unauthorized if the user doesn't have the required privileges.
    // @throws NotAllowedInCurrentState if the operation is not allowed in the current state.
    // @throws Error if there is some unknown internal error. The accompanying error message will give more details about the error.
	Check() error

    // Applies available NSX upgrade. This is an asynchronous call with fast return. The result of the operation can be queried by calling the API /vcenter/nsx/deployment/upgrade. **Warning:** This method is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    //
    // @param modeParam The mode of upgrade operation
    // @throws Unauthenticated if the user can not be authenticated.
    // @throws Unauthorized if the user doesn't have the required privileges.
    // @throws NotAllowedInCurrentState if the operation is not allowed in the current state.
    // @throws AlreadyInDesiredState if the operation is already in progress.
    // @throws Error if there is some unknown internal error. The accompanying error message will give more details about the error.
	Start(modeParam UpgradeExecutionMode) error

    // Pauses already running NSX upgrade. This is an asynchronous call with fast return. The result of the operation can be queried by calling the API /vcenter/nsx/deployment/upgrade. **Warning:** This method is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    // @throws Unauthenticated if the user can not be authenticated.
    // @throws Unauthorized if the user doesn't have the required privileges.
    // @throws NotAllowedInCurrentState if the operation is not allowed in the current state.
    // @throws AlreadyInDesiredState if the operation is already in progress.
    // @throws Error if there is some unknown internal error. The accompanying error message will give more details about the error.
	Pause() error

    // Resumes already paused NSX upgrade. This is an asynchronous call with fast return. The result of the operation can be queried by calling the API /vcenter/nsx/deployment/upgrade. **Warning:** This method is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    // @throws Unauthenticated if the user can not be authenticated.
    // @throws Unauthorized if the user doesn't have the required privileges.
    // @throws NotAllowedInCurrentState if the operation is not allowed in the current state.
    // @throws Error if there is some unknown internal error. The accompanying error message will give more details about the error.
	Resume() error
}
