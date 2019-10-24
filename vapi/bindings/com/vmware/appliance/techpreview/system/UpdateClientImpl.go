/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Client stubs for service: Update
 * Functions that implement the generated UpdateClient interface
 */

package system

import (
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/bindings/com/vmware/vapi/std/errors"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/bindings"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/core"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/data"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/lib"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/log"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/protocol/client"
)

type UpdateClientImpl struct {
	interfaceName       string
	interfaceDefinition core.InterfaceDefinition
	methodIdentifiers   []core.MethodIdentifier
	methodNameToDefMap  map[string]*core.MethodDefinition
	errorBindingMap     map[string]bindings.BindingType
	interfaceIdentifier core.InterfaceIdentifier
	connector           client.Connector
}

func NewUpdateClientImpl(connector client.Connector) *UpdateClientImpl {
	interfaceName := "com.vmware.appliance.techpreview.system.update"
	interfaceIdentifier := core.NewInterfaceIdentifier(interfaceName)
	methodIdentifiers := []core.MethodIdentifier{
		core.NewMethodIdentifier(interfaceIdentifier, "set"),
		core.NewMethodIdentifier(interfaceIdentifier, "get"),
	}
	interfaceDefinition := core.NewInterfaceDefinition(interfaceIdentifier, methodIdentifiers)
	errorBindingMap := make(map[string]bindings.BindingType)
	errorBindingMap[errors.InternalServerError{}.Error()] = errors.InternalServerErrorBindingType()
	errorBindingMap[errors.InvalidArgument{}.Error()] = errors.InvalidArgumentBindingType()
	errorBindingMap[errors.OperationNotFound{}.Error()] = errors.OperationNotFoundBindingType()
	errorBindingMap[errors.UnexpectedInput{}.Error()] = errors.UnexpectedInputBindingType()
	errorBindingMap[errors.ServiceUnavailable{}.Error()] = errors.ServiceUnavailableBindingType()
	errorBindingMap[errors.TimedOut{}.Error()] = errors.TimedOutBindingType()
	uIface := UpdateClientImpl{interfaceName: interfaceName, methodIdentifiers: methodIdentifiers, interfaceDefinition: interfaceDefinition, errorBindingMap: errorBindingMap, interfaceIdentifier: interfaceIdentifier, connector: connector}
	uIface.methodNameToDefMap = make(map[string]*core.MethodDefinition)
	uIface.methodNameToDefMap["set"] = uIface.setMethodDefinition()
	uIface.methodNameToDefMap["get"] = uIface.getMethodDefinition()
	return &uIface
}


func (uIface *UpdateClientImpl) Set(configParam UpdateUpdateStructSet) error {
	typeConverter := uIface.connector.TypeConverter()
	methodIdentifier := core.NewMethodIdentifier(uIface.interfaceIdentifier, "set")
	sv := bindings.NewStructValueBuilder(updateSetInputType(), typeConverter)
	sv.AddStructField("Config", configParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		return bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := updateSetRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	uIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := uIface.Invoke(uIface.connector.NewExecutionContext(), methodIdentifier, inputDataValue)
	if methodResult.IsSuccess() {
		return nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), uIface.errorBindingMap[methodResult.Error().Name()])
		if errorInError != nil {
			return bindings.VAPIerrorsToError(errorInError)
		}
		return methodError.(error)
	}
}

func (uIface *UpdateClientImpl) Get() (UpdateUpdateStructGet, error) {
	typeConverter := uIface.connector.TypeConverter()
	methodIdentifier := core.NewMethodIdentifier(uIface.interfaceIdentifier, "get")
	sv := bindings.NewStructValueBuilder(updateGetInputType(), typeConverter)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput UpdateUpdateStructGet
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := updateGetRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	uIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := uIface.Invoke(uIface.connector.NewExecutionContext(), methodIdentifier, inputDataValue)
	var emptyOutput UpdateUpdateStructGet
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), updateGetOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(UpdateUpdateStructGet), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), uIface.errorBindingMap[methodResult.Error().Name()])
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (uIface *UpdateClientImpl) Invoke(ctx *core.ExecutionContext, methodId core.MethodIdentifier, inputDataValue data.DataValue) core.MethodResult {
	methodResult := uIface.connector.GetApiProvider().Invoke(uIface.interfaceName, methodId.Name(), inputDataValue, ctx)
	return methodResult
}


func (uIface *UpdateClientImpl) setMethodDefinition() *core.MethodDefinition {
	interfaceIdentifier := core.NewInterfaceIdentifier(uIface.interfaceName)
	typeConverter := uIface.connector.TypeConverter()

	input, inputError := typeConverter.ConvertToDataDefinition(updateSetInputType())
	output, outputError := typeConverter.ConvertToDataDefinition(updateSetOutputType())
	if inputError != nil {
		log.Errorf("Error in ConvertToDataDefinition for UpdateClientImpl.set method's input - %s",
			bindings.VAPIerrorsToError(inputError).Error())
		return nil
	}
	if outputError != nil {
		log.Errorf("Error in ConvertToDataDefinition for UpdateClientImpl.set method's output - %s",
			bindings.VAPIerrorsToError(outputError).Error())
		return nil
	}
	methodIdentifier := core.NewMethodIdentifier(interfaceIdentifier, "set")
	errorDefinitions := make([]data.ErrorDefinition, 0)
	uIface.errorBindingMap[errors.Error{}.Error()] = errors.ErrorBindingType()
	errDef1, errError1 := typeConverter.ConvertToDataDefinition(errors.ErrorBindingType())
	if errError1 != nil {
		log.Errorf("Error in ConvertToDataDefinition for UpdateClientImpl.set method's errors.Error error - %s",
			bindings.VAPIerrorsToError(errError1).Error())
		return nil
	}
	errorDefinitions = append(errorDefinitions, errDef1.(data.ErrorDefinition))

	methodDefinition := core.NewMethodDefinition(methodIdentifier, input, output, errorDefinitions)
	return &methodDefinition
}

func (uIface *UpdateClientImpl) getMethodDefinition() *core.MethodDefinition {
	interfaceIdentifier := core.NewInterfaceIdentifier(uIface.interfaceName)
	typeConverter := uIface.connector.TypeConverter()

	input, inputError := typeConverter.ConvertToDataDefinition(updateGetInputType())
	output, outputError := typeConverter.ConvertToDataDefinition(updateGetOutputType())
	if inputError != nil {
		log.Errorf("Error in ConvertToDataDefinition for UpdateClientImpl.get method's input - %s",
			bindings.VAPIerrorsToError(inputError).Error())
		return nil
	}
	if outputError != nil {
		log.Errorf("Error in ConvertToDataDefinition for UpdateClientImpl.get method's output - %s",
			bindings.VAPIerrorsToError(outputError).Error())
		return nil
	}
	methodIdentifier := core.NewMethodIdentifier(interfaceIdentifier, "get")
	errorDefinitions := make([]data.ErrorDefinition, 0)
	uIface.errorBindingMap[errors.Error{}.Error()] = errors.ErrorBindingType()
	errDef1, errError1 := typeConverter.ConvertToDataDefinition(errors.ErrorBindingType())
	if errError1 != nil {
		log.Errorf("Error in ConvertToDataDefinition for UpdateClientImpl.get method's errors.Error error - %s",
			bindings.VAPIerrorsToError(errError1).Error())
		return nil
	}
	errorDefinitions = append(errorDefinitions, errDef1.(data.ErrorDefinition))

	methodDefinition := core.NewMethodDefinition(methodIdentifier, input, output, errorDefinitions)
	return &methodDefinition
}
