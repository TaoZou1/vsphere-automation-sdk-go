/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Interface file for service: Convergence
 * Used by client-side stubs.
 */
package hlm


// The ``Convergence`` check whether the community is converged. **Warning:** This interface is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
type ConvergenceClient interface {

    // Performs convergence checks for whole community in all serveices. This can be called after link finished to check whether the system is available. Or anytime want to know whether the system is converged after data been modified. **Warning:** This method is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    // @return ConvergenceConvergenceInfo objects. The convergence status per service per replica for all members of this community
    // @throws Unauthenticated If unable to authenticate to target node.
    // @throws Unauthorized If the caller is not authorized.
    // @throws Error If the system reports an error while responding to the request.
	Converge() (ConvergenceConvergenceInfo, error)
}
