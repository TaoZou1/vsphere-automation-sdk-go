/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Interface file for service: Nodes
 * Used by client-side stubs.
 */
package topology


// The ``Nodes`` interface provides methods to retrieve information about all linked vCenter Server nodes.
type NodesClient interface {

    // This operation retrieves detailed information about all linked vCenter Server nodes.
    // @return List of vCenter node details with. See NodesVcNode.
    // @throws Unauthenticated  if the user can not be authenticated.
    // @throws Error  if there are any errors in retrieving the vCenter data.
	List() ([]NodesVcNode, error)
}
