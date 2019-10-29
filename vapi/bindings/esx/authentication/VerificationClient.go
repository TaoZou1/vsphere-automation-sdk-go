/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Code generated. DO NOT EDIT.

/*
 * Interface file for service: Verification
 * Used by client-side stubs.
 */

package authentication


// ESX Authentication interface 
//
//  This ``Verification`` interface is providing a verification tool for tokens issued by Esx Authentication service.
type VerificationClient interface {

    // Verifies the provided token
    // @return username from the token
	Verify() (string, error)
}
