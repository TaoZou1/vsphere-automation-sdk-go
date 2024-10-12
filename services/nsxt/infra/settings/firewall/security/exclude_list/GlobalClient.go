// Copyright © 2019-2021 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Interface file for service: Global
// Used by client-side stubs.

package exclude_list

import (
	"github.com/vmware/vsphere-automation-sdk-go/lib/vapi/std/errors"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/bindings"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/core"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/lib"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/protocol/client"
	"github.com/vmware/vsphere-automation-sdk-go/services/nsxt/model"
)

const _ = core.SupportedByRuntimeVersion1

type GlobalClient interface {

	// Read exclude list for firewall
	// @return com.vmware.nsx_policy.model.PolicyExcludeList
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Get() (model.PolicyExcludeList, error)

	// Patch exclusion list.
	//
	// @param policyExcludeListParam (required)
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Patch(policyExcludeListParam model.PolicyExcludeList) error

	// Update the exclusion list
	//
	// @param policyExcludeListParam (required)
	// @return com.vmware.nsx_policy.model.PolicyExcludeList
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Update(policyExcludeListParam model.PolicyExcludeList) (model.PolicyExcludeList, error)
}

type globalClient struct {
	connector           client.Connector
	interfaceDefinition core.InterfaceDefinition
	errorsBindingMap    map[string]bindings.BindingType
}

func NewGlobalClient(connector client.Connector) *globalClient {
	interfaceIdentifier := core.NewInterfaceIdentifier("com.vmware.nsx_policy.infra.settings.firewall.security.exclude_list.global")
	methodIdentifiers := map[string]core.MethodIdentifier{
		"get":    core.NewMethodIdentifier(interfaceIdentifier, "get"),
		"patch":  core.NewMethodIdentifier(interfaceIdentifier, "patch"),
		"update": core.NewMethodIdentifier(interfaceIdentifier, "update"),
	}
	interfaceDefinition := core.NewInterfaceDefinition(interfaceIdentifier, methodIdentifiers)
	errorsBindingMap := make(map[string]bindings.BindingType)

	gIface := globalClient{interfaceDefinition: interfaceDefinition, errorsBindingMap: errorsBindingMap, connector: connector}
	return &gIface
}

func (gIface *globalClient) GetErrorBindingType(errorName string) bindings.BindingType {
	if entry, ok := gIface.errorsBindingMap[errorName]; ok {
		return entry
	}
	return errors.ERROR_BINDINGS_MAP[errorName]
}

func (gIface *globalClient) Get() (model.PolicyExcludeList, error) {
	typeConverter := gIface.connector.TypeConverter()
	executionContext := gIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(globalGetInputType(), typeConverter)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.PolicyExcludeList
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := globalGetRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	gIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := gIface.connector.GetApiProvider().Invoke("com.vmware.nsx_policy.infra.settings.firewall.security.exclude_list.global", "get", inputDataValue, executionContext)
	var emptyOutput model.PolicyExcludeList
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), globalGetOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.PolicyExcludeList), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), gIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (gIface *globalClient) Patch(policyExcludeListParam model.PolicyExcludeList) error {
	typeConverter := gIface.connector.TypeConverter()
	executionContext := gIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(globalPatchInputType(), typeConverter)
	sv.AddStructField("PolicyExcludeList", policyExcludeListParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		return bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := globalPatchRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	gIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := gIface.connector.GetApiProvider().Invoke("com.vmware.nsx_policy.infra.settings.firewall.security.exclude_list.global", "patch", inputDataValue, executionContext)
	if methodResult.IsSuccess() {
		return nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), gIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return bindings.VAPIerrorsToError(errorInError)
		}
		return methodError.(error)
	}
}

func (gIface *globalClient) Update(policyExcludeListParam model.PolicyExcludeList) (model.PolicyExcludeList, error) {
	typeConverter := gIface.connector.TypeConverter()
	executionContext := gIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(globalUpdateInputType(), typeConverter)
	sv.AddStructField("PolicyExcludeList", policyExcludeListParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.PolicyExcludeList
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := globalUpdateRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	gIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := gIface.connector.GetApiProvider().Invoke("com.vmware.nsx_policy.infra.settings.firewall.security.exclude_list.global", "update", inputDataValue, executionContext)
	var emptyOutput model.PolicyExcludeList
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), globalUpdateOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.PolicyExcludeList), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), gIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}
