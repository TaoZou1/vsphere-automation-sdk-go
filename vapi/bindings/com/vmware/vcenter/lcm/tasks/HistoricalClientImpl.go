/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Client stubs for service: Historical
 * Functions that implement the generated HistoricalClient interface
 */

package tasks

import (
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/bindings/com/vmware/vapi/std/errors"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/bindings/com/vmware/vcenter/lcm"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/bindings"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/core"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/data"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/lib"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/log"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/protocol/client"
)

type HistoricalClientImpl struct {
	interfaceName       string
	interfaceDefinition core.InterfaceDefinition
	methodIdentifiers   []core.MethodIdentifier
	methodNameToDefMap  map[string]*core.MethodDefinition
	errorBindingMap     map[string]bindings.BindingType
	interfaceIdentifier core.InterfaceIdentifier
	connector           client.Connector
}

func NewHistoricalClientImpl(connector client.Connector) *HistoricalClientImpl {
	interfaceName := "com.vmware.vcenter.lcm.tasks.historical"
	interfaceIdentifier := core.NewInterfaceIdentifier(interfaceName)
	methodIdentifiers := []core.MethodIdentifier{
		core.NewMethodIdentifier(interfaceIdentifier, "list"),
	}
	interfaceDefinition := core.NewInterfaceDefinition(interfaceIdentifier, methodIdentifiers)
	errorBindingMap := make(map[string]bindings.BindingType)
	errorBindingMap[errors.InternalServerError{}.Error()] = errors.InternalServerErrorBindingType()
	errorBindingMap[errors.InvalidArgument{}.Error()] = errors.InvalidArgumentBindingType()
	errorBindingMap[errors.OperationNotFound{}.Error()] = errors.OperationNotFoundBindingType()
	errorBindingMap[errors.UnexpectedInput{}.Error()] = errors.UnexpectedInputBindingType()
	errorBindingMap[errors.ServiceUnavailable{}.Error()] = errors.ServiceUnavailableBindingType()
	errorBindingMap[errors.TimedOut{}.Error()] = errors.TimedOutBindingType()
	hIface := HistoricalClientImpl{interfaceName: interfaceName, methodIdentifiers: methodIdentifiers, interfaceDefinition: interfaceDefinition, errorBindingMap: errorBindingMap, interfaceIdentifier: interfaceIdentifier, connector: connector}
	hIface.methodNameToDefMap = make(map[string]*core.MethodDefinition)
	hIface.methodNameToDefMap["list"] = hIface.listMethodDefinition()
	return &hIface
}


func (hIface *HistoricalClientImpl) List() ([]lcm.TasksInfo, error) {
	typeConverter := hIface.connector.TypeConverter()
	methodIdentifier := core.NewMethodIdentifier(hIface.interfaceIdentifier, "list")
	sv := bindings.NewStructValueBuilder(historicalListInputType(), typeConverter)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput []lcm.TasksInfo
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := historicalListRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	hIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := hIface.Invoke(hIface.connector.NewExecutionContext(), methodIdentifier, inputDataValue)
	var emptyOutput []lcm.TasksInfo
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), historicalListOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.([]lcm.TasksInfo), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), hIface.errorBindingMap[methodResult.Error().Name()])
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (hIface *HistoricalClientImpl) Invoke(ctx *core.ExecutionContext, methodId core.MethodIdentifier, inputDataValue data.DataValue) core.MethodResult {
	methodResult := hIface.connector.GetApiProvider().Invoke(hIface.interfaceName, methodId.Name(), inputDataValue, ctx)
	return methodResult
}


func (hIface *HistoricalClientImpl) listMethodDefinition() *core.MethodDefinition {
	interfaceIdentifier := core.NewInterfaceIdentifier(hIface.interfaceName)
	typeConverter := hIface.connector.TypeConverter()

	input, inputError := typeConverter.ConvertToDataDefinition(historicalListInputType())
	output, outputError := typeConverter.ConvertToDataDefinition(historicalListOutputType())
	if inputError != nil {
		log.Errorf("Error in ConvertToDataDefinition for HistoricalClientImpl.list method's input - %s",
			bindings.VAPIerrorsToError(inputError).Error())
		return nil
	}
	if outputError != nil {
		log.Errorf("Error in ConvertToDataDefinition for HistoricalClientImpl.list method's output - %s",
			bindings.VAPIerrorsToError(outputError).Error())
		return nil
	}
	methodIdentifier := core.NewMethodIdentifier(interfaceIdentifier, "list")
	errorDefinitions := make([]data.ErrorDefinition, 0)

	methodDefinition := core.NewMethodDefinition(methodIdentifier, input, output, errorDefinitions)
	return &methodDefinition
}
