/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Code generated. DO NOT EDIT.

/*
 * Interface file for service: SessionManager
 * Used by client-side stubs.
 */

package sessions


// Session management interface for the publish API
type SessionManagerClient interface {

    // Logs a user and returns a session id. Authentication information should be provided in the SecurityContext.
	Login() (string, error)

    // Invalidates a session. Session Id should be provided in the SecurityContext.
	Logout() error

    // Keeps a session id alive. Session Id should be provided in the SecurityContext.
	KeepAlive() error
}
