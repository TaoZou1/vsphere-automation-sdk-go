/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Code generated. DO NOT EDIT.

/*
 * Interface file for service: GetRoot
 * Used by client-side stubs.
 */

package certificate_authority


// The ``GetRoot`` interface provides methods to get VMCA's trusted root cert. **Warning:** This interface is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
type GetRootClient interface {

    // Returns the rhttpproxy TLS certificate. **Warning:** This method is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    // @return Root certificate.
    // @throws Error if failed to fetch root certificate
	GetRoot() (string, error)
}
