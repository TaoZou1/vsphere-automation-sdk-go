// Copyright © 2019-2021 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Interface file for service: Help
// Used by client-side stubs.

package pre_defined_runbooks

import (
	"github.com/vmware/vsphere-automation-sdk-go/lib/vapi/std/errors"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/bindings"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/core"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/lib"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/protocol/client"
	"github.com/vmware/vsphere-automation-sdk-go/services/nsxt/model"
)

const _ = core.SupportedByRuntimeVersion1

type HelpClient interface {

	// Read help information of an Online Diagnostic System predefined runbook.
	//
	// @param runbookIdParam Identifier of a predefined runbook (required)
	// @return com.vmware.nsx_policy.model.OdsRunbookHelpInfo
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Get(runbookIdParam string) (model.OdsRunbookHelpInfo, error)
}

type helpClient struct {
	connector           client.Connector
	interfaceDefinition core.InterfaceDefinition
	errorsBindingMap    map[string]bindings.BindingType
}

func NewHelpClient(connector client.Connector) *helpClient {
	interfaceIdentifier := core.NewInterfaceIdentifier("com.vmware.nsx_policy.infra.sha.pre_defined_runbooks.help")
	methodIdentifiers := map[string]core.MethodIdentifier{
		"get": core.NewMethodIdentifier(interfaceIdentifier, "get"),
	}
	interfaceDefinition := core.NewInterfaceDefinition(interfaceIdentifier, methodIdentifiers)
	errorsBindingMap := make(map[string]bindings.BindingType)

	hIface := helpClient{interfaceDefinition: interfaceDefinition, errorsBindingMap: errorsBindingMap, connector: connector}
	return &hIface
}

func (hIface *helpClient) GetErrorBindingType(errorName string) bindings.BindingType {
	if entry, ok := hIface.errorsBindingMap[errorName]; ok {
		return entry
	}
	return errors.ERROR_BINDINGS_MAP[errorName]
}

func (hIface *helpClient) Get(runbookIdParam string) (model.OdsRunbookHelpInfo, error) {
	typeConverter := hIface.connector.TypeConverter()
	executionContext := hIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(helpGetInputType(), typeConverter)
	sv.AddStructField("RunbookId", runbookIdParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.OdsRunbookHelpInfo
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := helpGetRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	hIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := hIface.connector.GetApiProvider().Invoke("com.vmware.nsx_policy.infra.sha.pre_defined_runbooks.help", "get", inputDataValue, executionContext)
	var emptyOutput model.OdsRunbookHelpInfo
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), helpGetOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.OdsRunbookHelpInfo), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), hIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}
