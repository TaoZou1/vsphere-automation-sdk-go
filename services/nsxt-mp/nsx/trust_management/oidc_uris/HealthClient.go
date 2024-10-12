// Copyright © 2019-2021 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Interface file for service: Health
// Used by client-side stubs.

package oidc_uris

import (
	"github.com/vmware/vsphere-automation-sdk-go/lib/vapi/std/errors"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/bindings"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/core"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/lib"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/protocol/client"
	"github.com/vmware/vsphere-automation-sdk-go/services/nsxt-mp/nsx/model"
)

const _ = core.SupportedByRuntimeVersion1

type HealthClient interface {

	// Connect to the OpenID Connect end-point and verify that it appears to be functioning properly.
	//
	// @param idParam (required)
	// @return com.vmware.nsx.model.OidcEndPointHealthStatus
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Get(idParam string) (model.OidcEndPointHealthStatus, error)
}

type healthClient struct {
	connector           client.Connector
	interfaceDefinition core.InterfaceDefinition
	errorsBindingMap    map[string]bindings.BindingType
}

func NewHealthClient(connector client.Connector) *healthClient {
	interfaceIdentifier := core.NewInterfaceIdentifier("com.vmware.nsx.trust_management.oidc_uris.health")
	methodIdentifiers := map[string]core.MethodIdentifier{
		"get": core.NewMethodIdentifier(interfaceIdentifier, "get"),
	}
	interfaceDefinition := core.NewInterfaceDefinition(interfaceIdentifier, methodIdentifiers)
	errorsBindingMap := make(map[string]bindings.BindingType)

	hIface := healthClient{interfaceDefinition: interfaceDefinition, errorsBindingMap: errorsBindingMap, connector: connector}
	return &hIface
}

func (hIface *healthClient) GetErrorBindingType(errorName string) bindings.BindingType {
	if entry, ok := hIface.errorsBindingMap[errorName]; ok {
		return entry
	}
	return errors.ERROR_BINDINGS_MAP[errorName]
}

func (hIface *healthClient) Get(idParam string) (model.OidcEndPointHealthStatus, error) {
	typeConverter := hIface.connector.TypeConverter()
	executionContext := hIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(healthGetInputType(), typeConverter)
	sv.AddStructField("Id", idParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.OidcEndPointHealthStatus
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := healthGetRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	hIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := hIface.connector.GetApiProvider().Invoke("com.vmware.nsx.trust_management.oidc_uris.health", "get", inputDataValue, executionContext)
	var emptyOutput model.OidcEndPointHealthStatus
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), healthGetOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.OidcEndPointHealthStatus), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), hIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}
