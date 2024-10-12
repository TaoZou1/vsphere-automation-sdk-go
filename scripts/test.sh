#! /bin/bash
set -e

vmcStageRefreshToken=$1
vmcOrgId=$2

pushd samples

go run main/main.go GetOrganizationById  $vmcStageRefreshToken $vmcOrgId
go run main/main.go GetDeploymentOverviewSummary $vmcStageRefreshToken $vmcOrgId
go run main/main.go GetCoreClusters $vmcStageRefreshToken $vmcOrgId
