// Copyright © 2023 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

package inventory

import (
	"fmt"

	"github.com/vmware/vsphere-automation-sdk-go/runtime/protocol/client"
	"github.com/vmware/vsphere-automation-sdk-go/services/vmwarecloud/vmc_core_inventory/api/inventory/core"
	"github.com/vmware/vsphere-automation-sdk-go/services/vmwarecloud/vmc_core_inventory/model"
)

func GetDeployments(orgId string, connector client.Connector) model.DeploymentResponse {

	includeDeletedResourcesParam := false
	deploymentsClient := core.NewDeploymentsClient(connector)
	pageParam := int64(0)
	sizeParam := int64(100)
	summaryParam := make([]string, 2)
	summaryParam[0] = "DeploymentOverviewSummary"
	summaryParam[1] = "DeploymentVcenterSummary"
	deploymentResponse, err := deploymentsClient.GetCoreDeployments(orgId, summaryParam, &includeDeletedResourcesParam, &pageParam, &sizeParam, nil)
	if err != nil {
		fmt.Println("Error: ", err)
		panic(err)
	}

	fmt.Println("Response size: ", len(deploymentResponse.Content))
	for _, deployment := range deploymentResponse.Content {
		fmt.Println("Name: ", deployment.Name)
		if deployment.Summary.DeploymentOverviewSummary != nil {
			fmt.Println("HostCount Summary: ", *deployment.Summary.DeploymentOverviewSummary.HostCount)
		}
		if deployment.Summary.DeploymentVcenterSummary != nil {
			fmt.Println("VC URL: ", *deployment.Summary.DeploymentVcenterSummary.VcUrl)
		}
	}

	return deploymentResponse
}
