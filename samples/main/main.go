// Copyright © 2023 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

package main

import (
	"fmt"
	"os"

	"github.com/vmware/vsphere-automation-sdk-go/samples/vmc/inventory"
	"github.com/vmware/vsphere-automation-sdk-go/samples/vmc/org"
	"github.com/vmware/vsphere-automation-sdk-go/samples/vmc/utils"
)

func main() {
	args := os.Args

	argLength := len(args[1:])
	fmt.Println(argLength)

	// Refresh token -> api token from csp account.
	var sampleToRun = args[1]
	var refreshToken = args[2]
	var orgId = args[3]

	// Create connector instance.
	connector, err := utils.NewVmcConnector(refreshToken, utils.ServerUrl, utils.TokenUrl)
	if err != nil {
		panic(err)
	}

	switch sampleToRun {
	case utils.GET_DEPLOYMENTS:
		inventory.GetDeployments(orgId, connector)
	case utils.GET_CORE_CLUSTERS:
		var deploymentId = ""
		if argLength == 3 {
			var deploymentResponse = inventory.GetDeployments(orgId, connector)
			deploymentId = *deploymentResponse.Content[0].Id
		} else {
			deploymentId = args[4]
		}
		inventory.GetCoreClusters(orgId, deploymentId, connector)
	case utils.GET_ORG_BY_ID:
		org.GetOrganizationById(orgId, connector)
	case utils.GET_DEPLOYMENT_SUMMARY:
		var deploymentId = ""
		if argLength == 3 {
			var deploymentResponse = inventory.GetDeployments(orgId, connector)
			deploymentId = *deploymentResponse.Content[0].Id
		} else {
			deploymentId = args[4]
		}
		inventory.GetDeploymentOverviewSummary(orgId, deploymentId, connector)
	default:
		fmt.Println("Invalid Sample!!!")
	}
}
