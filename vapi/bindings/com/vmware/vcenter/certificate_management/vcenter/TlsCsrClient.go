/* Copyright © 2019 VMware, Inc. All Rights Reserved.
     SPDX-License-Identifier: BSD-2-Clause */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Interface file for service: TlsCsr
 * Used by client-side stubs.
 */

package vcenter

import (
)

// The ``TlsCsr`` interface provides methods to generate certificate signing request. This interface was added in vSphere API 6.7.2.
type TlsCsrClient interface {


    // Generates a CSR with the given Spec. This method was added in vSphere API 6.7.2.
    //
    // @param specParam The information needed to create a CSR.
    // @return A Certificate Signing Request.
    // @throws Error If CSR could not be created for given spec for a generic error.
    Create(specParam TlsCsrSpec) (TlsCsrInfo, error) 

}
