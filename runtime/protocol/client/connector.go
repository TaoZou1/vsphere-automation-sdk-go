/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

package client

import (
	"github.com/TaoZou1/vsphere-automation-sdk-go/runtime/bindings"
	"github.com/TaoZou1/vsphere-automation-sdk-go/runtime/core"
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
