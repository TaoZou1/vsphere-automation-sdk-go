/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Code generated. DO NOT EDIT.

/*
 * Interface file for service: TlsCsr
 * Used by client-side stubs.
 */

package vcenter


// The ``TlsCsr`` interface provides methods to generate certificate signing request.
type TlsCsrClient interface {

    // Generates a CSR with the given Spec.
    //
    // @param specParam The information needed to create a CSR.
    // @return A Certificate Signing Request.
    // @throws Error If CSR could not be created for given spec for a generic error.
	Create(specParam TlsCsrSpec) (TlsCsrInfo, error)
}
