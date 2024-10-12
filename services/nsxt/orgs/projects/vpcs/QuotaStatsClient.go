// Copyright © 2019-2021 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Interface file for service: QuotaStats
// Used by client-side stubs.

package vpcs

import (
	"github.com/vmware/vsphere-automation-sdk-go/lib/vapi/std/errors"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/bindings"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/core"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/lib"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/protocol/client"
	"github.com/vmware/vsphere-automation-sdk-go/services/nsxt/model"
)

const _ = core.SupportedByRuntimeVersion1

type QuotaStatsClient interface {

	// Get quota details
	//
	// @param orgIdParam (required)
	// @param projectIdParam (required)
	// @param vpcIdParam (required)
	// @param pathPrefixParam Path prefix for retriving the quota details. (required)
	// @param constraintPathParam Constraint path to retrive the quota details. (optional)
	// @return com.vmware.nsx_policy.model.QuotaStatsListResult
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Get(orgIdParam string, projectIdParam string, vpcIdParam string, pathPrefixParam string, constraintPathParam *string) (model.QuotaStatsListResult, error)
}

type quotaStatsClient struct {
	connector           client.Connector
	interfaceDefinition core.InterfaceDefinition
	errorsBindingMap    map[string]bindings.BindingType
}

func NewQuotaStatsClient(connector client.Connector) *quotaStatsClient {
	interfaceIdentifier := core.NewInterfaceIdentifier("com.vmware.nsx_policy.orgs.projects.vpcs.quota_stats")
	methodIdentifiers := map[string]core.MethodIdentifier{
		"get": core.NewMethodIdentifier(interfaceIdentifier, "get"),
	}
	interfaceDefinition := core.NewInterfaceDefinition(interfaceIdentifier, methodIdentifiers)
	errorsBindingMap := make(map[string]bindings.BindingType)

	qIface := quotaStatsClient{interfaceDefinition: interfaceDefinition, errorsBindingMap: errorsBindingMap, connector: connector}
	return &qIface
}

func (qIface *quotaStatsClient) GetErrorBindingType(errorName string) bindings.BindingType {
	if entry, ok := qIface.errorsBindingMap[errorName]; ok {
		return entry
	}
	return errors.ERROR_BINDINGS_MAP[errorName]
}

func (qIface *quotaStatsClient) Get(orgIdParam string, projectIdParam string, vpcIdParam string, pathPrefixParam string, constraintPathParam *string) (model.QuotaStatsListResult, error) {
	typeConverter := qIface.connector.TypeConverter()
	executionContext := qIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(quotaStatsGetInputType(), typeConverter)
	sv.AddStructField("OrgId", orgIdParam)
	sv.AddStructField("ProjectId", projectIdParam)
	sv.AddStructField("VpcId", vpcIdParam)
	sv.AddStructField("PathPrefix", pathPrefixParam)
	sv.AddStructField("ConstraintPath", constraintPathParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.QuotaStatsListResult
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := quotaStatsGetRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	qIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := qIface.connector.GetApiProvider().Invoke("com.vmware.nsx_policy.orgs.projects.vpcs.quota_stats", "get", inputDataValue, executionContext)
	var emptyOutput model.QuotaStatsListResult
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), quotaStatsGetOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.QuotaStatsListResult), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), qIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}
