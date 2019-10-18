/* Copyright © 2019 VMware, Inc. All Rights Reserved.
     SPDX-License-Identifier: BSD-2-Clause */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Interface file for service: Sync
 * Used by client-side stubs.
 */

package links

import (
)

// The ``Sync`` interface provides methods to create a sync session, get information on Sync. Usage beyond VMware Cloud on AWS is not supported. **Warning:** This interface is available as Technology Preview. These are early access APIs provided to test, automate and provide feedback on the feature. Since this can change based on feedback, VMware does not guarantee backwards compatibility and recommends against using them in production environments. Some Technology Preview APIs might only be applicable to specific environments.
type SyncClient interface {


    // Resets the sync state between the linked domains by initiating a fresh sync for all providers. If an existing sync is in progress this cancels the sync. Usage beyond VMware Cloud on AWS is not supported. **Warning:** This method is available as Technology Preview. These are early access APIs provided to test, automate and provide feedback on the feature. Since this can change based on feedback, VMware does not guarantee backwards compatibility and recommends against using them in production environments. Some Technology Preview APIs might only be applicable to specific environments.
    //
    // @param linkParam Unique identifier of the link.
    // The parameter must be an identifier for the resource type: ``com.vmware.vcenter.hvc.Links``.
    // @throws Error if the system reports an error while responding to the request.
    // @throws NotFound if the link Identifier associated with ``link`` does not exist.
    // @throws Unauthorized if the user is not authorized to perform this operation.
    Reset(linkParam string) error 


    // Remediates the sync state between the linked domains by initiating a fresh sync for certificate providers. If an existing sync is in progress this cancels the sync. Usage beyond VMware Cloud on AWS is not supported. **Warning:** This method is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    //
    // @param linkParam Unique identifier of the link.
    // The parameter must be an identifier for the resource type: ``com.vmware.vcenter.hvc.Links``.
    // @param credentialsParam Credentials to use for authentication to the remote system.
    // This parameter is currently required. In the future, it may become optional.
    // @throws Error if the system reports an error while responding to the request.
    // @throws NotFound if the link Identifier associated with ``link`` does not exist.
    // @throws Unauthorized if the user is not authorized to perform this operation.
    // @throws ResourceBusy if a sync is already running.
    // @throws TimedOut if the remediate action times out.
    Remediate(linkParam string, credentialsParam *SyncCredentials) (SyncRemediationStatus, error) 

}
