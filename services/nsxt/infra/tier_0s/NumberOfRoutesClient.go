// Copyright © 2019-2021 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Interface file for service: NumberOfRoutes
// Used by client-side stubs.

package tier_0s

import (
	"github.com/vmware/vsphere-automation-sdk-go/lib/vapi/std/errors"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/bindings"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/core"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/lib"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/protocol/client"
	"github.com/vmware/vsphere-automation-sdk-go/services/nsxt/model"
)

const _ = core.SupportedByRuntimeVersion1

type NumberOfRoutesClient interface {

	// Get total number of IPv4 or IPv6 routes on a Tier0
	//
	// @param tier0IdParam (required)
	// @param edgePathParam Policy path of edge node (required)
	// @param enforcementPointPathParam String Path of the enforcement point (optional)
	// @param includeChildVrfParam Count all the child VRF routes or not. (optional)
	// @return com.vmware.nsx_policy.model.Tier0NumberOfRoutesResult
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Get(tier0IdParam string, edgePathParam string, enforcementPointPathParam *string, includeChildVrfParam *bool) (model.Tier0NumberOfRoutesResult, error)
}

type numberOfRoutesClient struct {
	connector           client.Connector
	interfaceDefinition core.InterfaceDefinition
	errorsBindingMap    map[string]bindings.BindingType
}

func NewNumberOfRoutesClient(connector client.Connector) *numberOfRoutesClient {
	interfaceIdentifier := core.NewInterfaceIdentifier("com.vmware.nsx_policy.infra.tier_0s.number_of_routes")
	methodIdentifiers := map[string]core.MethodIdentifier{
		"get": core.NewMethodIdentifier(interfaceIdentifier, "get"),
	}
	interfaceDefinition := core.NewInterfaceDefinition(interfaceIdentifier, methodIdentifiers)
	errorsBindingMap := make(map[string]bindings.BindingType)

	nIface := numberOfRoutesClient{interfaceDefinition: interfaceDefinition, errorsBindingMap: errorsBindingMap, connector: connector}
	return &nIface
}

func (nIface *numberOfRoutesClient) GetErrorBindingType(errorName string) bindings.BindingType {
	if entry, ok := nIface.errorsBindingMap[errorName]; ok {
		return entry
	}
	return errors.ERROR_BINDINGS_MAP[errorName]
}

func (nIface *numberOfRoutesClient) Get(tier0IdParam string, edgePathParam string, enforcementPointPathParam *string, includeChildVrfParam *bool) (model.Tier0NumberOfRoutesResult, error) {
	typeConverter := nIface.connector.TypeConverter()
	executionContext := nIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(numberOfRoutesGetInputType(), typeConverter)
	sv.AddStructField("Tier0Id", tier0IdParam)
	sv.AddStructField("EdgePath", edgePathParam)
	sv.AddStructField("EnforcementPointPath", enforcementPointPathParam)
	sv.AddStructField("IncludeChildVrf", includeChildVrfParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.Tier0NumberOfRoutesResult
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := numberOfRoutesGetRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	nIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := nIface.connector.GetApiProvider().Invoke("com.vmware.nsx_policy.infra.tier_0s.number_of_routes", "get", inputDataValue, executionContext)
	var emptyOutput model.Tier0NumberOfRoutesResult
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), numberOfRoutesGetOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.Tier0NumberOfRoutesResult), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), nIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}
