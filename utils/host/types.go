/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

package host

import (
	"net/http"

	"github.com/vmware/vsphere-automation-sdk-go/runtime/protocol/client"
	"github.com/vmware/vsphere-automation-sdk-go/utils/args"
	"github.com/vmware/vsphere-automation-sdk-go/utils/auth"
	"github.com/vmware/vsphere-automation-sdk-go/utils/host/kind"
	"github.com/vmware/vsphere-automation-sdk-go/utils/session"
	"github.com/vmware/vsphere-automation-sdk-go/utils/task"
)

// Info represents Host details.
type Info interface {
	// Name returns name of the Host.
	Name() string
	// Server returns address of the Host.
	Server() string
	// Kind returns kind of the Host.
	Kind() kind.Info
	// AuthScheme returns authentication scheme of the Host.
	AuthScheme() auth.Scheme
	// AuthDetails returns Authentication Details of the Host.
	AuthDetails() auth.Info
	// ExecuteTask executes task 't' on the Host.
	ExecuteTask(t task.Function) error
	// Connector returns connector for the Host.
	Connector() client.Connector
	// IsSet returns 'true' if properties are set for the Host, else returns 'false'.
	IsSet() bool
	// getHTTPClient returns HTTP Client for the Host.
	getHTTPClient() http.Client
	// getSessionManager returns Session Manager of the Host.
	getSessionManager() (session.Manager, error)
	args.Properties
}
