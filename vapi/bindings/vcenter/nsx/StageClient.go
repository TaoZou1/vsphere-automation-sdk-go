/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Code generated. DO NOT EDIT.

/*
 * Interface file for service: Stage
 * Used by client-side stubs.
 */

package nsx


// The ``Stage`` interface represents all the operations of to Stage NSX bundels on vCenter server. **Warning:** This interface is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
type StageClient interface {

    // Download NSX binaries from the my.vmware.com portal on vCenter Server. **Warning:** This method is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    // @throws Unauthenticated if the user can not be authenticated.
    // @throws Unauthorized if the user doesn't have the required privileges.
    // @throws NotAllowedInCurrentState if the operation is not allowed in the current state.
    // @throws Error if there is some unknown internal error. The accompanying error message will give more details about the error.
	Download() error

    // Upload NSX binaries to vCenter server. **Warning:** This method is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    // @throws Unauthenticated if the user can not be authenticated.
    // @throws Unauthorized if the user doesn't have the required privileges.
    // @throws NotAllowedInCurrentState if the operation is not allowed in the current state.
    // @throws Error if there is some unknown internal error. The accompanying error message will give more details about the error.
	Upload() error

    // Cancel running NSX download or upload process. **Warning:** This method is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    // @throws Unauthenticated if the user can not be authenticated.
    // @throws Unauthorized if the user doesn't have the required privileges.
    // @throws NotAllowedInCurrentState if the operation is not allowed in the current state.
    // @throws Error if there is some unknown internal error. The accompanying error message will give more details about the error.
	Cancel() error

    // Get operation status. **Warning:** This method is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    // @return ExecutionStatus
    // @throws Unauthenticated if the user can not be authenticated.
    // @throws Unauthorized if the user doesn't have the required privileges.
    // @throws Error if there is some unknown internal error. The accompanying error message will give more details about the error.
	CheckStatus() (StageExecutionStatus, error)
}
