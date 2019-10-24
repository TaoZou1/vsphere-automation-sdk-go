/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Interface file for service: DeploymentType
 * Used by client-side stubs.
 */
package cluster


// The DeploymentType interface provides methods to get the deployment type of a vCenter High Availability Cluster (VCHA Cluster).
type DeploymentTypeClient interface {

    // Retrieves the deployment type of a VCHA cluster.
    // @return Info structure containing the deployment type information of the the VCHA cluster.
    // @throws Unauthorized If the user has insufficient privilege to perform the operation. Operation execution requires the System.Read privilege.
    // @throws Error If any other error occurs.
	Get() (DeploymentTypeInfo, error)
}
