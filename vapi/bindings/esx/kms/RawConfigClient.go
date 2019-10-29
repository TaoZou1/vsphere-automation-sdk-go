/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Code generated. DO NOT EDIT.

/*
 * Interface file for service: RawConfig
 * Used by client-side stubs.
 */

package kms


// The ``RawConfig`` interface provides methods to get or put the key management service persistent user configuration. 
//
//  This interface is most useful when either replicating existing configuration settings to a new host, or when applying a known desired configuration across several hosts in a cluster.
type RawConfigClient interface {

    // Return the configuration information. 
    //
    //  Returns the key management service persistent user configuration.
    //
    // @param type_Param The type of RawConfigInfo that will be returned.
    // If {\\\\@term.unset} RawConfigInfoType#RawConfigInfoType_BRIEF will be used.
    // @return The configuration.
    // @throws Error if there is a generic error.
    // @throws InvalidArgument if the info type is invalid.
    // @throws Unauthenticated if the caller is not authenticated.
    // @throws Unauthorized if the caller is not authorized.
	Get(type_Param *RawConfigInfoType) (RawConfigInfo, error)

    // Set the configuration. 
    //
    //  Overwrites all existing persistent user configuration with the specified configuration.
    //
    // @param specParam The configuration.
    // @throws Error if there is a generic error.
    // @throws InvalidArgument if the configuration is invalid.
    // @throws Unauthenticated if the caller is not authenticated.
    // @throws Unauthorized if the caller is not authorized.
	Set(specParam RawConfigSetSpec) error
}
