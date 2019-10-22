/* Copyright © 2019 VMware, Inc. All Rights Reserved.
     SPDX-License-Identifier: BSD-2-Clause */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Interface file for service: ReplicationStatus
 * Used by client-side stubs.
 */

package topology

import (
)

// The ``ReplicationStatus`` interface provides methods to retrieve replication status information of vCenter and Platform Services Controller nodes of type VCSA_EMBEDDED/PSC_EXTERNAL (see NodesInfo#type).
type ReplicationStatusClient interface {


    // Returns the replication information of vCenter and Platform Services Controller nodes of type VCSA_EMBEDDED/PSC_EXTERNAL (see NodesInfo#type) matching the ReplicationStatusFilterSpec.
    //
    // @param filterParam  Specification of matching vCenter and Platform Services Controller nodes for which information should be returned.
    // If null, the behavior is equivalent to a ReplicationStatusFilterSpec with all properties null which means all vCenter and Platform Services Controller nodes of type VCSA_EMBEDDED/PSC_EXTERNAL match the filter.
    // @return Commonly used replication information about vCenter and Platform Services Controller nodes matching the ReplicationStatusFilterSpec.
    // @throws Unauthenticated  if the user can not be authenticated.
    // @throws Unauthorized  if the user doesn't have the required privileges.
    // @throws InvalidArgument  if the ReplicationStatusFilterSpec#nodes property contains a invalid value.
    List(filterParam *ReplicationStatusFilterSpec) ([]ReplicationStatusSummary, error) 

}
