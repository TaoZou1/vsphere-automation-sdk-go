package client

import (
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/bindings"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/core"
)

type Connector interface {
	GetApiProvider() core.APIProvider
	SetApplicationContext(*core.ApplicationContext)
	ApplicationContext() *core.ApplicationContext
	SetSecurityContext(core.SecurityContext)
	SecurityContext() core.SecurityContext
	NewExecutionContext() *core.ExecutionContext
	TypeConverter() *bindings.TypeConverter
	SetConnectionMetadata(map[string]interface{})
	ConnectionMetadata() map[string]interface{}
}
