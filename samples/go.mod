module github.com/vmware/vsphere-automation-sdk-go/samples

go 1.17

require (
	github.com/vmware/vsphere-automation-sdk-go/runtime v0.7.0
	github.com/vmware/vsphere-automation-sdk-go/services/vmwarecloud/vmc_core_inventory v0.1.0
	github.com/vmware/vsphere-automation-sdk-go/services/vmwarecloud/vmc_core_org v0.1.0
)

require (
	github.com/beevik/etree v1.1.0 // indirect
	github.com/gibson042/canonicaljson-go v1.0.3 // indirect
	github.com/golang-jwt/jwt/v4 v4.3.0 // indirect
	github.com/google/uuid v1.2.0 // indirect
	github.com/vmware/vsphere-automation-sdk-go/lib v0.7.0 // indirect
	golang.org/x/text v0.3.8 // indirect
)

replace (
	github.com/vmware/vsphere-automation-sdk-go/lib => ../lib
	github.com/vmware/vsphere-automation-sdk-go/runtime => ../runtime
	github.com/vmware/vsphere-automation-sdk-go/services/vmwarecloud/vmc_core_inventory => ../services/vmwarecloud/vmc_core_inventory
	github.com/vmware/vsphere-automation-sdk-go/services/vmwarecloud/vmc_core_org => ../services/vmwarecloud/vmc_core_org
)
