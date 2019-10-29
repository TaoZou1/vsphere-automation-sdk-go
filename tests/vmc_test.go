package tests

import (
	"github.com/stretchr/testify/assert"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/utils"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/bindings/vmc"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/bindings/vmc/orgs"
	"testing"
)

var refreshToken = "DqPVNL902N3nAnuE4afk6Em5Fk36HGSLqwPPnKxPW1NYzeK00LPO48kKfZ3c5Jbl"
var orgID = "058f47c4-92aa-417f-8747-87f3ed61cb45"
var sddcID = "933d8724-836b-430a-82c0-2a347e0fb488"

func TestOrgName(t *testing.T) {

	connector, err := utils.NewVmcConnector(refreshToken, "", "")
	if err != nil {
		panic(err)
	}

	orgClient := vmc.NewDefaultOrgsClient(connector)
	org, err := orgClient.Get(orgID)

	if err != nil {
		panic(err)
	}
	assert.Equal(t, "CPBU-API/CLI/SDK", *org.DisplayName)
}

func TestListOrgs(t *testing.T) {

	connector, err := utils.NewVmcConnector(refreshToken, "", "")
	if err != nil {
		panic(err)
	}

	orgClient := vmc.NewDefaultOrgsClient(connector)
	orgs, err := orgClient.List()

	if err != nil {
		panic(err)
	}
	assert.Equal(t, 2, len(orgs))
}

// func TestTasks(t *testing.T) {
// 	connector, _ := utils.NewVmcConnector(refreshToken, "", "")
// 	taskClient := tasks.NewTasksClientImpl(connector)
// 	// filterPara := "(status eq 'STARTED')"
// 	tasks, err := taskClient.List(orgID, nil)
// 	if err != nil {
// 		panic(err)
// 	}
// 	assert.Equal(t, 10, len(tasks))
// }

func TestGetSDDC(t *testing.T) {

	connector, err := utils.NewVmcConnector(refreshToken, "", "")
	if err != nil {
		panic(err)
	}

	sddcClient := orgs.NewDefaultSddcsClient(connector)

	sddc, err := sddcClient.Get(orgID, sddcID)
	if err != nil {
		panic(err)
	}
	assert.Equal(t, "PowerCLI-Test-SDDC", *sddc.Name)
}

func TestAccountLinking(t *testing.T) {
	connector, err := utils.NewVmcConnector(refreshToken, "", "")
	if err != nil {
		panic(err)
	}
	accountLinkClient := orgs.NewDefaultAccountLinkClient(connector)
	err = accountLinkClient.Get(orgID)
	assert.Nil(t, err)
}
