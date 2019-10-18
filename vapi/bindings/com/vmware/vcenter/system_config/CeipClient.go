/* Copyright © 2019 VMware, Inc. All Rights Reserved.
     SPDX-License-Identifier: BSD-2-Clause */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Interface file for service: Ceip
 * Used by client-side stubs.
 */

package system_config

import (
)

// The ``Ceip`` interface provides methods to get and set the CEIP (Customer Experience Improvement Program) value. If CEIP is set to True, the customer agrees to participate in the CEIP program. Setting the CEIP value to True or False affects the whole MxN deployment. This interface was added in vSphere API 6.7.
type CeipClient interface {


    // Set CEIP (Customer Experience Improvement Program) value. This method was added in vSphere API 6.7.
    //
    // @param valueParam CEIP boolean value to be set.
    // @throws Error if any error occurs during the execution of the set operation.
    Set(valueParam bool) error 


    // Get CEIP (Customer Experience Improvement Program) value. This method was added in vSphere API 6.7.
    // @return CEIP bool value.
    // @throws Error if any error occurs during the execution of the get operation.
    Get() (bool, error) 

}
