// Copyright © 2019-2021 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Interface file for service: Metrics
// Used by client-side stubs.

package agents

import (
	"github.com/vmware/vsphere-automation-sdk-go/lib/vapi/std/errors"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/bindings"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/core"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/lib"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/protocol/client"
	"github.com/vmware/vsphere-automation-sdk-go/services/nsxt/model"
)

const _ = core.SupportedByRuntimeVersion1

type MetricsClient interface {

	// Fetch current Metric agent config
	//
	// @param orgIdParam The organization ID (required)
	// @param projectIdParam The project ID (required)
	// @return com.vmware.nsx_policy.model.MetricsAgent
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Get(orgIdParam string, projectIdParam string) (model.MetricsAgent, error)

	// Create Metrics agent intent resource
	//
	// @param orgIdParam The organization ID (required)
	// @param projectIdParam The project ID (required)
	// @param metricsAgentParam (required)
	// @return com.vmware.nsx_policy.model.MetricsAgent
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Update(orgIdParam string, projectIdParam string, metricsAgentParam model.MetricsAgent) (model.MetricsAgent, error)
}

type metricsClient struct {
	connector           client.Connector
	interfaceDefinition core.InterfaceDefinition
	errorsBindingMap    map[string]bindings.BindingType
}

func NewMetricsClient(connector client.Connector) *metricsClient {
	interfaceIdentifier := core.NewInterfaceIdentifier("com.vmware.nsx_policy.orgs.projects.infra.sites.agents.metrics")
	methodIdentifiers := map[string]core.MethodIdentifier{
		"get":    core.NewMethodIdentifier(interfaceIdentifier, "get"),
		"update": core.NewMethodIdentifier(interfaceIdentifier, "update"),
	}
	interfaceDefinition := core.NewInterfaceDefinition(interfaceIdentifier, methodIdentifiers)
	errorsBindingMap := make(map[string]bindings.BindingType)

	mIface := metricsClient{interfaceDefinition: interfaceDefinition, errorsBindingMap: errorsBindingMap, connector: connector}
	return &mIface
}

func (mIface *metricsClient) GetErrorBindingType(errorName string) bindings.BindingType {
	if entry, ok := mIface.errorsBindingMap[errorName]; ok {
		return entry
	}
	return errors.ERROR_BINDINGS_MAP[errorName]
}

func (mIface *metricsClient) Get(orgIdParam string, projectIdParam string) (model.MetricsAgent, error) {
	typeConverter := mIface.connector.TypeConverter()
	executionContext := mIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(metricsGetInputType(), typeConverter)
	sv.AddStructField("OrgId", orgIdParam)
	sv.AddStructField("ProjectId", projectIdParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.MetricsAgent
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := metricsGetRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	mIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := mIface.connector.GetApiProvider().Invoke("com.vmware.nsx_policy.orgs.projects.infra.sites.agents.metrics", "get", inputDataValue, executionContext)
	var emptyOutput model.MetricsAgent
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), metricsGetOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.MetricsAgent), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), mIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (mIface *metricsClient) Update(orgIdParam string, projectIdParam string, metricsAgentParam model.MetricsAgent) (model.MetricsAgent, error) {
	typeConverter := mIface.connector.TypeConverter()
	executionContext := mIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(metricsUpdateInputType(), typeConverter)
	sv.AddStructField("OrgId", orgIdParam)
	sv.AddStructField("ProjectId", projectIdParam)
	sv.AddStructField("MetricsAgent", metricsAgentParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.MetricsAgent
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := metricsUpdateRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	mIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := mIface.connector.GetApiProvider().Invoke("com.vmware.nsx_policy.orgs.projects.infra.sites.agents.metrics", "update", inputDataValue, executionContext)
	var emptyOutput model.MetricsAgent
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), metricsUpdateOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.MetricsAgent), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), mIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}
