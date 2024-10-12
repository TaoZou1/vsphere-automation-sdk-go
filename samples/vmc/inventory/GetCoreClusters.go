// Copyright © 2023 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

package inventory

import (
	"fmt"

	"github.com/vmware/vsphere-automation-sdk-go/runtime/protocol/client"
	"github.com/vmware/vsphere-automation-sdk-go/services/vmwarecloud/vmc_core_inventory/api/inventory/core"
	"github.com/vmware/vsphere-automation-sdk-go/services/vmwarecloud/vmc_core_inventory/model"
)

func GetCoreClusters(orgId, deploymentId string, connector client.Connector) model.ClusterResponse {

	includeDeletedResourcesParam := false
	deploymentsClient := core.NewClustersClient(connector)
	pageParam := int64(0)
	sizeParam := int64(100)
	clusterResponse, err := deploymentsClient.GetCoreClusters(orgId, &deploymentId, &includeDeletedResourcesParam, &pageParam, &sizeParam, nil)
	if err != nil {
		panic(err)
	}

	clusterId := clusterResponse.Content[0].Id
	clusterName := clusterResponse.Content[0].Name
	fmt.Println("Cluster Id: ", *clusterId)
	fmt.Println("Name: ", clusterName)

	if clusterId == nil {
		panic(err)
	}

	return clusterResponse
}
