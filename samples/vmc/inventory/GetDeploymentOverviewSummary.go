// Copyright © 2023 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

package inventory

import (
	"fmt"

	"github.com/vmware/vsphere-automation-sdk-go/runtime/protocol/client"
	"github.com/vmware/vsphere-automation-sdk-go/services/vmwarecloud/vmc_core_inventory/api/inventory/core/deployments/summaries"
	"github.com/vmware/vsphere-automation-sdk-go/services/vmwarecloud/vmc_core_inventory/model"
)

func GetDeploymentOverviewSummary(orgId, deploumentId string, connector client.Connector) model.DeploymentOverviewSummary {

	deploymentOverviewSummaryClient := summaries.NewDeploymentOverviewSummaryClient(connector)
	deploymentOverviewSummaryResponse, err := deploymentOverviewSummaryClient.GetCoreDeploymentOverviewSummary(orgId, deploumentId)
	if err != nil {
		panic(err)
	}

	fmt.Println("Name: ", deploymentOverviewSummaryResponse.Name)
	fmt.Println("Resource Id: ", deploymentOverviewSummaryResponse.ResourceId)
	fmt.Println("Cluster Count: ", *deploymentOverviewSummaryResponse.ClusterCount)
	fmt.Println("Host Count: ", *deploymentOverviewSummaryResponse.HostCount)

	return deploymentOverviewSummaryResponse
}
