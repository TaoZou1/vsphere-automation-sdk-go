// Copyright © 2019-2021 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Interface file for service: StandaloneHostSwitchSetting
// Used by client-side stubs.

package idfw

import (
	"github.com/vmware/vsphere-automation-sdk-go/lib/vapi/std/errors"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/bindings"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/core"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/lib"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/protocol/client"
	"github.com/vmware/vsphere-automation-sdk-go/services/nsxt-mp/nsx/model"
)

const _ = core.SupportedByRuntimeVersion1

type StandaloneHostSwitchSettingClient interface {

	// Fetches IDFW standalone hosts switch setting to check whether standalone hosts is enabled or disabled
	//
	//  Use the following Policy API -
	//  GET /policy/api/v1/infra/settings/firewall/idfw/standalone-host-switch-setting
	// @return com.vmware.nsx.model.IdfwStandaloneHostsSwitchSetting
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Get() (model.IdfwStandaloneHostsSwitchSetting, error)

	// Update Identity Firewall standalone hosts switch setting (true=enabled / false=disabled).
	//
	//  Use the following Policy API -
	//  PUT/PATCH /policy/api/v1/infra/settings/firewall/idfw/standalone-host-switch-setting
	//
	// @param idfwStandaloneHostsSwitchSettingParam (required)
	// @return com.vmware.nsx.model.IdfwStandaloneHostsSwitchSetting
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Update(idfwStandaloneHostsSwitchSettingParam model.IdfwStandaloneHostsSwitchSetting) (model.IdfwStandaloneHostsSwitchSetting, error)
}

type standaloneHostSwitchSettingClient struct {
	connector           client.Connector
	interfaceDefinition core.InterfaceDefinition
	errorsBindingMap    map[string]bindings.BindingType
}

func NewStandaloneHostSwitchSettingClient(connector client.Connector) *standaloneHostSwitchSettingClient {
	interfaceIdentifier := core.NewInterfaceIdentifier("com.vmware.nsx.idfw.standalone_host_switch_setting")
	methodIdentifiers := map[string]core.MethodIdentifier{
		"get":    core.NewMethodIdentifier(interfaceIdentifier, "get"),
		"update": core.NewMethodIdentifier(interfaceIdentifier, "update"),
	}
	interfaceDefinition := core.NewInterfaceDefinition(interfaceIdentifier, methodIdentifiers)
	errorsBindingMap := make(map[string]bindings.BindingType)

	sIface := standaloneHostSwitchSettingClient{interfaceDefinition: interfaceDefinition, errorsBindingMap: errorsBindingMap, connector: connector}
	return &sIface
}

func (sIface *standaloneHostSwitchSettingClient) GetErrorBindingType(errorName string) bindings.BindingType {
	if entry, ok := sIface.errorsBindingMap[errorName]; ok {
		return entry
	}
	return errors.ERROR_BINDINGS_MAP[errorName]
}

func (sIface *standaloneHostSwitchSettingClient) Get() (model.IdfwStandaloneHostsSwitchSetting, error) {
	typeConverter := sIface.connector.TypeConverter()
	executionContext := sIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(standaloneHostSwitchSettingGetInputType(), typeConverter)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.IdfwStandaloneHostsSwitchSetting
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := standaloneHostSwitchSettingGetRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	sIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := sIface.connector.GetApiProvider().Invoke("com.vmware.nsx.idfw.standalone_host_switch_setting", "get", inputDataValue, executionContext)
	var emptyOutput model.IdfwStandaloneHostsSwitchSetting
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), standaloneHostSwitchSettingGetOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.IdfwStandaloneHostsSwitchSetting), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), sIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (sIface *standaloneHostSwitchSettingClient) Update(idfwStandaloneHostsSwitchSettingParam model.IdfwStandaloneHostsSwitchSetting) (model.IdfwStandaloneHostsSwitchSetting, error) {
	typeConverter := sIface.connector.TypeConverter()
	executionContext := sIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(standaloneHostSwitchSettingUpdateInputType(), typeConverter)
	sv.AddStructField("IdfwStandaloneHostsSwitchSetting", idfwStandaloneHostsSwitchSettingParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.IdfwStandaloneHostsSwitchSetting
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := standaloneHostSwitchSettingUpdateRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	sIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := sIface.connector.GetApiProvider().Invoke("com.vmware.nsx.idfw.standalone_host_switch_setting", "update", inputDataValue, executionContext)
	var emptyOutput model.IdfwStandaloneHostsSwitchSetting
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), standaloneHostSwitchSettingUpdateOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.IdfwStandaloneHostsSwitchSetting), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), sIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}
