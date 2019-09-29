/*
 * Copyright 2019 VMware, Inc.  All rights reserved.
 */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Interface file for service: Replicated
 * Used by client-side stubs.
 */

package replicated

import (
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/bindings/vcenter/deployment"
)

// The ``Replicated`` interface provides methods to check if the configuring vCenter Server can be replicated to the remote PSC. This interface was added in vSphere API 6.7.
type ReplicatedClient interface {


    // Checks whether the provided remote PSC is reachable and can be replicated. This method was added in vSphere API 6.7.
    //
    // @param specParam Information to configure a replicated PSC.
    // @return Information about the success or failure of the checks that were performed.
    // @throws Unauthenticated if the caller is not authenticated.
    // @throws Unauthenticated if external PSC credentials are not valid when configuring PSC to replicate with an external existing PSC.
    // @throws InvalidArgument if passed arguments are invalid.
    // @throws NotAllowedInCurrentState if the appliance is not in INITIALIZED state.
    Check(specParam deployment.ReplicatedPscSpec) (deployment.CheckInfo, error) 

}
