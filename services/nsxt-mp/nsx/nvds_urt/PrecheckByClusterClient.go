// Copyright © 2019-2021 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Interface file for service: PrecheckByCluster
// Used by client-side stubs.

package nvds_urt

import (
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/lib/vapi/std/errors"
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/runtime/bindings"
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/runtime/core"
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/runtime/lib"
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/runtime/protocol/client"
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/services/nsxt-mp/nsx/model"
)

const _ = core.SupportedByRuntimeVersion1

type PrecheckByClusterClient interface {

	// Start precheck for N-VDS to VDS migration by cluster
	//
	// @param clusterIdParam (required)
	// @param tolerateDifferentConfigurationsParam tolerate differnet configurations (optional, default to true)
	// @return com.vmware.nsx.model.NvdsUpgradePrecheckId
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Create(clusterIdParam string, tolerateDifferentConfigurationsParam *bool) (model.NvdsUpgradePrecheckId, error)

	// Retrieve latest precheck ID of the N-VDS to VDS migration for the cluster
	//
	// @param clusterIdParam (required)
	// @return com.vmware.nsx.model.NvdsUpgradePrecheckId
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Get(clusterIdParam string) (model.NvdsUpgradePrecheckId, error)
}

type precheckByClusterClient struct {
	connector           client.Connector
	interfaceDefinition core.InterfaceDefinition
	errorsBindingMap    map[string]bindings.BindingType
}

func NewPrecheckByClusterClient(connector client.Connector) *precheckByClusterClient {
	interfaceIdentifier := core.NewInterfaceIdentifier("com.vmware.nsx.nvds_urt.precheck_by_cluster")
	methodIdentifiers := map[string]core.MethodIdentifier{
		"create": core.NewMethodIdentifier(interfaceIdentifier, "create"),
		"get":    core.NewMethodIdentifier(interfaceIdentifier, "get"),
	}
	interfaceDefinition := core.NewInterfaceDefinition(interfaceIdentifier, methodIdentifiers)
	errorsBindingMap := make(map[string]bindings.BindingType)

	pIface := precheckByClusterClient{interfaceDefinition: interfaceDefinition, errorsBindingMap: errorsBindingMap, connector: connector}
	return &pIface
}

func (pIface *precheckByClusterClient) GetErrorBindingType(errorName string) bindings.BindingType {
	if entry, ok := pIface.errorsBindingMap[errorName]; ok {
		return entry
	}
	return errors.ERROR_BINDINGS_MAP[errorName]
}

func (pIface *precheckByClusterClient) Create(clusterIdParam string, tolerateDifferentConfigurationsParam *bool) (model.NvdsUpgradePrecheckId, error) {
	typeConverter := pIface.connector.TypeConverter()
	executionContext := pIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(precheckByClusterCreateInputType(), typeConverter)
	sv.AddStructField("ClusterId", clusterIdParam)
	sv.AddStructField("TolerateDifferentConfigurations", tolerateDifferentConfigurationsParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.NvdsUpgradePrecheckId
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := precheckByClusterCreateRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	pIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := pIface.connector.GetApiProvider().Invoke("com.vmware.nsx.nvds_urt.precheck_by_cluster", "create", inputDataValue, executionContext)
	var emptyOutput model.NvdsUpgradePrecheckId
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), precheckByClusterCreateOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.NvdsUpgradePrecheckId), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), pIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (pIface *precheckByClusterClient) Get(clusterIdParam string) (model.NvdsUpgradePrecheckId, error) {
	typeConverter := pIface.connector.TypeConverter()
	executionContext := pIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(precheckByClusterGetInputType(), typeConverter)
	sv.AddStructField("ClusterId", clusterIdParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.NvdsUpgradePrecheckId
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := precheckByClusterGetRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	pIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := pIface.connector.GetApiProvider().Invoke("com.vmware.nsx.nvds_urt.precheck_by_cluster", "get", inputDataValue, executionContext)
	var emptyOutput model.NvdsUpgradePrecheckId
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), precheckByClusterGetOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.NvdsUpgradePrecheckId), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), pIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}
