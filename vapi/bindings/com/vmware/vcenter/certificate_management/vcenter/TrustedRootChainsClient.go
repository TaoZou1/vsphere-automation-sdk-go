/* Copyright © 2019 VMware, Inc. All Rights Reserved.
     SPDX-License-Identifier: BSD-2-Clause */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Interface file for service: TrustedRootChains
 * Used by client-side stubs.
 */

package vcenter

import (
)

// The ``TrustedRootChains`` interface provides methods to create, modify, delete and read trusted root certificate chains.
type TrustedRootChainsClient interface {


    // Returns summary information for each trusted root certificate chain.
    // @return List of trusted root certificate chains summaries.
    // @throws Unauthorized if authorization is not given to caller.
    List() ([]TrustedRootChainsSummary, error) 


    // Creates a new trusted root certificate chain from the CreateSpec.
    //
    // @param specParam The information needed to create a trusted root certificate chain.
    // @return The unique identifier for the new trusted root chain.
    // @throws Unauthorized if authorization is not given to caller.
    // @throws AlreadyExists if a trusted root certificate chain exists with id in given spec.
    Create(specParam TrustedRootChainsCreateSpec) (string, error) 


    // Retrieve a trusted root certificate chain for a given identifier.
    //
    // @param chainParam Unique identifier for a trusted root cert chain.
    // @return TrustedRootChain.
    // @throws Unauthorized if authorization is not given to caller.
    // @throws NotFound if a trusted root certificate chain does not exist for given id.
    Get(chainParam string) (TrustedRootChainsInfo, error) 


    // Deletes trusted root certificate chain for a given identifier.
    //
    // @param chainParam Unique identifier for a trusted root cert chain.
    // @throws Unauthorized if authorization is not given to caller.
    // @throws NotFound if a trusted root certificate chain does not exist for given id.
    Delete(chainParam string) error 

}
