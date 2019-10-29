package tests

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"

	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/bindings/cis/tagging"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/bindings/vcenter"
	// "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/bindings/vcenter/vm"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/protocol/client"

	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/utils"
)

var server = "https://vspheresdk-vc.eng.vmware.com"
var username = "administrator@vsphere.local"
var password = "Admin!23"

var connector client.Connector
var err error

func TestMain(m *testing.M) {
	// call flag.Parse() here if TestMain uses flags
	connector, err = utils.NewVsphereConnector(server, username, password)
	if err != nil {
		panic(err)
	}
	os.Exit(m.Run())
}

func TestClusters(t *testing.T) {
	clusterClient := vcenter.NewDefaultClusterClient(connector)
	clusters, err := clusterClient.List(nil)
	if err != nil {
		panic(err)
	}
	assert.Equal(t, "Cluster1", clusters[0].Name)
}

func TestTagsCategories(t *testing.T) {
	catClient := tagging.NewDefaultCategoryClient(connector)
	cats, err := catClient.List()
	if err != nil {
		panic(err)
	}
	assert.Equal(t, 1, len(cats))
}

func TestListVMs(t *testing.T) {
	vmClient := vcenter.NewDefaultVMClient(connector)
	vms, err := vmClient.List(nil)
	if err != nil {
		panic(err)
	}
	assert.True(t, len(vms) > 10)
}



// func Test(t *testing.T) {
// 	Client := (connector)
// 	, err := Client.List(nil)
// 	if err != nil {
// 		panic(err)
// 	}
// 	assert.True(t, len(vms) > 10)
// }

// func Test(t *testing.T) {
// 	Client := (connector)
// 	, err := Client.List(nil)
// 	if err != nil {
// 		panic(err)
// 	}
// 	assert.True(t, len(vms) > 10)
// }

// func Test(t *testing.T) {
// 	Client := (connector)
// 	, err := Client.List(nil)
// 	if err != nil {
// 		panic(err)
// 	}
// 	assert.True(t, len(vms) > 10)
// }

// func Test(t *testing.T) {
// 	Client := (connector)
// 	, err := Client.List(nil)
// 	if err != nil {
// 		panic(err)
// 	}
// 	assert.True(t, len(vms) > 10)
// }
