package tests

import (
	"os"
	"testing"

	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/bindings/cis/tagging/category"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/bindings/vcenter/cluster"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/protocol/client"

	"fmt"

	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/utils"
)

var server = "https://sc2-10-184-170-114.eng.vmware.com"
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
	clusterClient := cluster.NewClusterClientImpl(connector)
	clusters, err := clusterClient.List(nil)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Number of clusters in vCenter: %b", len(clusters))
}

func TestTagsCategories(t *testing.T) {
	catClient := category.NewCategoryClientImpl(connector)
	tags, err := catClient.List()
	if err != nil {
		panic(err)
	}
	fmt.Printf("There are %b tags in vCenter", len(tags))
}
