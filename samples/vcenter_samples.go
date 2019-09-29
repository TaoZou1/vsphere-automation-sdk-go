package main

import (
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/bindings/cis/tagging/category"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/bindings/vcenter/cluster"

	"fmt"

	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/utils"
)

var server = "http://sc2-rdops-vm08-dhcp-8-186.eng.vmware.com"
var username = "administrator@vsphere.local"
var password = "Admin!23"

func main() {

	connector, err := utils.NewVsphereConnector(server, username, password)
	if err != nil {
		panic(err)
	}

	catClient := category.NewCategoryClientImpl(connector)
	tags, err := catClient.List()
	if err != nil {
		panic(err)
	}
	fmt.Printf("There are %b tags in vCenter", len(tags))

	clusterClient := cluster.NewClusterClientImpl(connector)
	clusters, err := clusterClient.List(nil)
	fmt.Printf("Number of clusters in vCenter: %b", len(clusters))
}
