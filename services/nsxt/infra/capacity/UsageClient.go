// Copyright © 2019-2021 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Interface file for service: Usage
// Used by client-side stubs.

package capacity

import (
	"github.com/TaoZou1/vsphere-automation-sdk-go/lib/vapi/std/errors"
	"github.com/TaoZou1/vsphere-automation-sdk-go/runtime/bindings"
	"github.com/TaoZou1/vsphere-automation-sdk-go/runtime/core"
	"github.com/TaoZou1/vsphere-automation-sdk-go/runtime/lib"
	"github.com/TaoZou1/vsphere-automation-sdk-go/runtime/protocol/client"
	"github.com/TaoZou1/vsphere-automation-sdk-go/services/nsxt/model"
)

const _ = core.SupportedByRuntimeVersion1

type UsageClient interface {

	// Returns capacity usage data for NSX objects
	//
	//  This API is no longer been used by UI, please use below policy api instead.
	//  GET https://<nsx-mgr>/policy/api/v1/infra/capacity/dashboard/usage
	//
	// @param categoryParam (optional)
	// @param cursorParam Opaque cursor to be used for getting next page of records (supplied by current result page) (optional)
	// @param includedFieldsParam Comma separated list of fields that should be included in query result (optional)
	// @param pageSizeParam Maximum number of results to return in this page (server may return fewer) (optional, default to 1000)
	// @param sortAscendingParam (optional)
	// @param sortByParam Field by which records are sorted (optional)
	// @return com.vmware.nsx_policy.model.PolicyCapacityUsageResponse
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Get(categoryParam *string, cursorParam *string, includedFieldsParam *string, pageSizeParam *int64, sortAscendingParam *bool, sortByParam *string) (model.PolicyCapacityUsageResponse, error)
}

type usageClient struct {
	connector           client.Connector
	interfaceDefinition core.InterfaceDefinition
	errorsBindingMap    map[string]bindings.BindingType
}

func NewUsageClient(connector client.Connector) *usageClient {
	interfaceIdentifier := core.NewInterfaceIdentifier("com.vmware.nsx_policy.infra.capacity.usage")
	methodIdentifiers := map[string]core.MethodIdentifier{
		"get": core.NewMethodIdentifier(interfaceIdentifier, "get"),
	}
	interfaceDefinition := core.NewInterfaceDefinition(interfaceIdentifier, methodIdentifiers)
	errorsBindingMap := make(map[string]bindings.BindingType)

	uIface := usageClient{interfaceDefinition: interfaceDefinition, errorsBindingMap: errorsBindingMap, connector: connector}
	return &uIface
}

func (uIface *usageClient) GetErrorBindingType(errorName string) bindings.BindingType {
	if entry, ok := uIface.errorsBindingMap[errorName]; ok {
		return entry
	}
	return errors.ERROR_BINDINGS_MAP[errorName]
}

func (uIface *usageClient) Get(categoryParam *string, cursorParam *string, includedFieldsParam *string, pageSizeParam *int64, sortAscendingParam *bool, sortByParam *string) (model.PolicyCapacityUsageResponse, error) {
	typeConverter := uIface.connector.TypeConverter()
	executionContext := uIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(usageGetInputType(), typeConverter)
	sv.AddStructField("Category", categoryParam)
	sv.AddStructField("Cursor", cursorParam)
	sv.AddStructField("IncludedFields", includedFieldsParam)
	sv.AddStructField("PageSize", pageSizeParam)
	sv.AddStructField("SortAscending", sortAscendingParam)
	sv.AddStructField("SortBy", sortByParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.PolicyCapacityUsageResponse
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := usageGetRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	uIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := uIface.connector.GetApiProvider().Invoke("com.vmware.nsx_policy.infra.capacity.usage", "get", inputDataValue, executionContext)
	var emptyOutput model.PolicyCapacityUsageResponse
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), usageGetOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.PolicyCapacityUsageResponse), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), uIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}
