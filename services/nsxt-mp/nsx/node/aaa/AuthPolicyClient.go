// Copyright © 2019-2021 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Interface file for service: AuthPolicy
// Used by client-side stubs.

package aaa

import (
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/lib/vapi/std/errors"
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/runtime/bindings"
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/runtime/core"
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/runtime/lib"
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/runtime/protocol/client"
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/services/nsxt-mp/nsx/model"
)

const _ = core.SupportedByRuntimeVersion1

type AuthPolicyClient interface {

	// Returns information about the currently configured authentication policies and password complexity on the node.
	// @return com.vmware.nsx.model.AuthenticationPolicyProperties
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Get() (model.AuthenticationPolicyProperties, error)

	// Resets to default, currently configured authentication policy and password complexity on the node. Administrators need to enforce password change for existing user accounts in order to match newly configured complexity requirements in system.
	//
	// **reset-all**: resets configured Authentication policy and Password complexity
	//
	// **reset-auth-policies**: resets only configured Authentication policy
	//  *includes - {api_failed_auth_lockout_period, api_failed_auth_reset_period, api_max_auth_failures, cli_failed_auth_lockout_period, cli_max_auth_failures}*
	//
	// **reset-pwd-complexity**: resets only configured Password complexity
	// @return com.vmware.nsx.model.AuthenticationPolicyProperties
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Resetall() (model.AuthenticationPolicyProperties, error)

	// Resets to default, currently configured authentication policy and password complexity on the node. Administrators need to enforce password change for existing user accounts in order to match newly configured complexity requirements in system.
	//
	// **reset-all**: resets configured Authentication policy and Password complexity
	//
	// **reset-auth-policies**: resets only configured Authentication policy
	//  *includes - {api_failed_auth_lockout_period, api_failed_auth_reset_period, api_max_auth_failures, cli_failed_auth_lockout_period, cli_max_auth_failures}*
	//
	// **reset-pwd-complexity**: resets only configured Password complexity
	// @return com.vmware.nsx.model.AuthenticationPolicyProperties
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Resetauthpolicies() (model.AuthenticationPolicyProperties, error)

	// Resets to default, currently configured authentication policy and password complexity on the node. Administrators need to enforce password change for existing user accounts in order to match newly configured complexity requirements in system.
	//
	// **reset-all**: resets configured Authentication policy and Password complexity
	//
	// **reset-auth-policies**: resets only configured Authentication policy
	//  *includes - {api_failed_auth_lockout_period, api_failed_auth_reset_period, api_max_auth_failures, cli_failed_auth_lockout_period, cli_max_auth_failures}*
	//
	// **reset-pwd-complexity**: resets only configured Password complexity
	// @return com.vmware.nsx.model.AuthenticationPolicyProperties
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Resetpwdcomplexity() (model.AuthenticationPolicyProperties, error)

	// Update the currently configured authentication policy and password complexity on the node. If any of api_max_auth_failures, api_failed_auth_reset_period, or api_failed_auth_lockout_period are modified, the http service is automatically restarted. Whereas change in any password complexity will not be applicable on already configured user passwords. Administrators need to enforce password change for existing user accounts in order to match newly configured complexity requirements enforced in system. All values from AuthenticationPolicyProperties are in sync among the management cluster nodes.
	//
	// @param authenticationPolicyPropertiesParam (required)
	// @return com.vmware.nsx.model.AuthenticationPolicyProperties
	// @throws ConcurrentChange  Conflict
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Update(authenticationPolicyPropertiesParam model.AuthenticationPolicyProperties) (model.AuthenticationPolicyProperties, error)
}

type authPolicyClient struct {
	connector           client.Connector
	interfaceDefinition core.InterfaceDefinition
	errorsBindingMap    map[string]bindings.BindingType
}

func NewAuthPolicyClient(connector client.Connector) *authPolicyClient {
	interfaceIdentifier := core.NewInterfaceIdentifier("com.vmware.nsx.node.aaa.auth_policy")
	methodIdentifiers := map[string]core.MethodIdentifier{
		"get":                core.NewMethodIdentifier(interfaceIdentifier, "get"),
		"resetall":           core.NewMethodIdentifier(interfaceIdentifier, "resetall"),
		"resetauthpolicies":  core.NewMethodIdentifier(interfaceIdentifier, "resetauthpolicies"),
		"resetpwdcomplexity": core.NewMethodIdentifier(interfaceIdentifier, "resetpwdcomplexity"),
		"update":             core.NewMethodIdentifier(interfaceIdentifier, "update"),
	}
	interfaceDefinition := core.NewInterfaceDefinition(interfaceIdentifier, methodIdentifiers)
	errorsBindingMap := make(map[string]bindings.BindingType)

	aIface := authPolicyClient{interfaceDefinition: interfaceDefinition, errorsBindingMap: errorsBindingMap, connector: connector}
	return &aIface
}

func (aIface *authPolicyClient) GetErrorBindingType(errorName string) bindings.BindingType {
	if entry, ok := aIface.errorsBindingMap[errorName]; ok {
		return entry
	}
	return errors.ERROR_BINDINGS_MAP[errorName]
}

func (aIface *authPolicyClient) Get() (model.AuthenticationPolicyProperties, error) {
	typeConverter := aIface.connector.TypeConverter()
	executionContext := aIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(authPolicyGetInputType(), typeConverter)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.AuthenticationPolicyProperties
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := authPolicyGetRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	aIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := aIface.connector.GetApiProvider().Invoke("com.vmware.nsx.node.aaa.auth_policy", "get", inputDataValue, executionContext)
	var emptyOutput model.AuthenticationPolicyProperties
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), authPolicyGetOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.AuthenticationPolicyProperties), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), aIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (aIface *authPolicyClient) Resetall() (model.AuthenticationPolicyProperties, error) {
	typeConverter := aIface.connector.TypeConverter()
	executionContext := aIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(authPolicyResetallInputType(), typeConverter)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.AuthenticationPolicyProperties
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := authPolicyResetallRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	aIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := aIface.connector.GetApiProvider().Invoke("com.vmware.nsx.node.aaa.auth_policy", "resetall", inputDataValue, executionContext)
	var emptyOutput model.AuthenticationPolicyProperties
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), authPolicyResetallOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.AuthenticationPolicyProperties), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), aIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (aIface *authPolicyClient) Resetauthpolicies() (model.AuthenticationPolicyProperties, error) {
	typeConverter := aIface.connector.TypeConverter()
	executionContext := aIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(authPolicyResetauthpoliciesInputType(), typeConverter)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.AuthenticationPolicyProperties
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := authPolicyResetauthpoliciesRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	aIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := aIface.connector.GetApiProvider().Invoke("com.vmware.nsx.node.aaa.auth_policy", "resetauthpolicies", inputDataValue, executionContext)
	var emptyOutput model.AuthenticationPolicyProperties
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), authPolicyResetauthpoliciesOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.AuthenticationPolicyProperties), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), aIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (aIface *authPolicyClient) Resetpwdcomplexity() (model.AuthenticationPolicyProperties, error) {
	typeConverter := aIface.connector.TypeConverter()
	executionContext := aIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(authPolicyResetpwdcomplexityInputType(), typeConverter)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.AuthenticationPolicyProperties
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := authPolicyResetpwdcomplexityRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	aIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := aIface.connector.GetApiProvider().Invoke("com.vmware.nsx.node.aaa.auth_policy", "resetpwdcomplexity", inputDataValue, executionContext)
	var emptyOutput model.AuthenticationPolicyProperties
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), authPolicyResetpwdcomplexityOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.AuthenticationPolicyProperties), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), aIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (aIface *authPolicyClient) Update(authenticationPolicyPropertiesParam model.AuthenticationPolicyProperties) (model.AuthenticationPolicyProperties, error) {
	typeConverter := aIface.connector.TypeConverter()
	executionContext := aIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(authPolicyUpdateInputType(), typeConverter)
	sv.AddStructField("AuthenticationPolicyProperties", authenticationPolicyPropertiesParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.AuthenticationPolicyProperties
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := authPolicyUpdateRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	aIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := aIface.connector.GetApiProvider().Invoke("com.vmware.nsx.node.aaa.auth_policy", "update", inputDataValue, executionContext)
	var emptyOutput model.AuthenticationPolicyProperties
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), authPolicyUpdateOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.AuthenticationPolicyProperties), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), aIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}
