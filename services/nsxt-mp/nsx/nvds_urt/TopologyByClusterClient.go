// Copyright © 2019-2021 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Interface file for service: TopologyByCluster
// Used by client-side stubs.

package nvds_urt

import (
	"github.com/TaoZou1/vsphere-automation-sdk-go/lib/vapi/std/errors"
	"github.com/TaoZou1/vsphere-automation-sdk-go/runtime/bindings"
	"github.com/TaoZou1/vsphere-automation-sdk-go/runtime/core"
	"github.com/TaoZou1/vsphere-automation-sdk-go/runtime/lib"
	"github.com/TaoZou1/vsphere-automation-sdk-go/runtime/protocol/client"
	"github.com/TaoZou1/vsphere-automation-sdk-go/services/nsxt-mp/nsx/model"
)

const _ = core.SupportedByRuntimeVersion1

type TopologyByClusterClient interface {

	// Set VDS configuration and create it in vCenter
	//
	// @param precheckIdParam (required)
	// @param upgradeTopologyParam (required)
	// @param clusterIdParam cluster identifier (optional)
	// @param useRecommendedTopologyConfigParam Flag to indicate if use recommended topology got from the latest precheck (optional)
	// @return com.vmware.nsx.model.UpgradeTopology
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Apply(precheckIdParam string, upgradeTopologyParam model.UpgradeTopology, clusterIdParam *string, useRecommendedTopologyConfigParam *bool) (model.UpgradeTopology, error)

	// Recommmended topology
	//
	// @param precheckIdParam (required)
	// @param clusterIdParam cluster identifier (optional)
	// @param computeManagerIdParam vCenter identifier (optional)
	// @param showVdsConfigParam Flag to indicate if VdsTopology should contain VDS configuration (optional)
	// @return com.vmware.nsx.model.UpgradeTopology
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Get(precheckIdParam string, clusterIdParam *string, computeManagerIdParam *string, showVdsConfigParam *bool) (model.UpgradeTopology, error)
}

type topologyByClusterClient struct {
	connector           client.Connector
	interfaceDefinition core.InterfaceDefinition
	errorsBindingMap    map[string]bindings.BindingType
}

func NewTopologyByClusterClient(connector client.Connector) *topologyByClusterClient {
	interfaceIdentifier := core.NewInterfaceIdentifier("com.vmware.nsx.nvds_urt.topology_by_cluster")
	methodIdentifiers := map[string]core.MethodIdentifier{
		"apply": core.NewMethodIdentifier(interfaceIdentifier, "apply"),
		"get":   core.NewMethodIdentifier(interfaceIdentifier, "get"),
	}
	interfaceDefinition := core.NewInterfaceDefinition(interfaceIdentifier, methodIdentifiers)
	errorsBindingMap := make(map[string]bindings.BindingType)

	tIface := topologyByClusterClient{interfaceDefinition: interfaceDefinition, errorsBindingMap: errorsBindingMap, connector: connector}
	return &tIface
}

func (tIface *topologyByClusterClient) GetErrorBindingType(errorName string) bindings.BindingType {
	if entry, ok := tIface.errorsBindingMap[errorName]; ok {
		return entry
	}
	return errors.ERROR_BINDINGS_MAP[errorName]
}

func (tIface *topologyByClusterClient) Apply(precheckIdParam string, upgradeTopologyParam model.UpgradeTopology, clusterIdParam *string, useRecommendedTopologyConfigParam *bool) (model.UpgradeTopology, error) {
	typeConverter := tIface.connector.TypeConverter()
	executionContext := tIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(topologyByClusterApplyInputType(), typeConverter)
	sv.AddStructField("PrecheckId", precheckIdParam)
	sv.AddStructField("UpgradeTopology", upgradeTopologyParam)
	sv.AddStructField("ClusterId", clusterIdParam)
	sv.AddStructField("UseRecommendedTopologyConfig", useRecommendedTopologyConfigParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.UpgradeTopology
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := topologyByClusterApplyRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	tIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := tIface.connector.GetApiProvider().Invoke("com.vmware.nsx.nvds_urt.topology_by_cluster", "apply", inputDataValue, executionContext)
	var emptyOutput model.UpgradeTopology
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), topologyByClusterApplyOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.UpgradeTopology), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), tIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (tIface *topologyByClusterClient) Get(precheckIdParam string, clusterIdParam *string, computeManagerIdParam *string, showVdsConfigParam *bool) (model.UpgradeTopology, error) {
	typeConverter := tIface.connector.TypeConverter()
	executionContext := tIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(topologyByClusterGetInputType(), typeConverter)
	sv.AddStructField("PrecheckId", precheckIdParam)
	sv.AddStructField("ClusterId", clusterIdParam)
	sv.AddStructField("ComputeManagerId", computeManagerIdParam)
	sv.AddStructField("ShowVdsConfig", showVdsConfigParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.UpgradeTopology
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := topologyByClusterGetRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	tIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := tIface.connector.GetApiProvider().Invoke("com.vmware.nsx.nvds_urt.topology_by_cluster", "get", inputDataValue, executionContext)
	var emptyOutput model.UpgradeTopology
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), topologyByClusterGetOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.UpgradeTopology), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), tIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}
