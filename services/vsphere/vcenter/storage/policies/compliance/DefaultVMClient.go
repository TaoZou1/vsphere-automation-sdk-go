
/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Code generated. DO NOT EDIT.

/*
 * Client stubs for service: VM
 * Functions that implement the generated VMClient interface
 */


package compliance

import (
	"github.com/vmware/vsphere-automation-sdk-go/lib/vapi/std/errors"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/bindings"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/core"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/data"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/lib"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/log"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/protocol/client"
)

type DefaultVMClient struct {
	interfaceName       string
	interfaceDefinition core.InterfaceDefinition
	methodIdentifiers   []core.MethodIdentifier
	methodNameToDefMap  map[string]*core.MethodDefinition
	errorBindingMap     map[string]bindings.BindingType
	interfaceIdentifier core.InterfaceIdentifier
	connector           client.Connector
}

func NewDefaultVMClient(connector client.Connector) *DefaultVMClient {
	interfaceName := "com.vmware.vcenter.storage.policies.compliance.VM"
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
	vIface := DefaultVMClient{interfaceName: interfaceName, methodIdentifiers: methodIdentifiers, interfaceDefinition: interfaceDefinition, errorBindingMap: errorBindingMap, interfaceIdentifier: interfaceIdentifier, connector: connector}
	vIface.methodNameToDefMap = make(map[string]*core.MethodDefinition)
	vIface.methodNameToDefMap["list"] = vIface.listMethodDefinition()
	return &vIface
}

func (vIface *DefaultVMClient) List(filterParam VMFilterSpec) (map[string]VMInfo, error) {
	typeConverter := vIface.connector.TypeConverter()
	methodIdentifier := core.NewMethodIdentifier(vIface.interfaceIdentifier, "list")
	sv := bindings.NewStructValueBuilder(vMListInputType(), typeConverter)
	sv.AddStructField("Filter", filterParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput map[string]VMInfo
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := vMListRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	vIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := vIface.Invoke(vIface.connector.NewExecutionContext(), methodIdentifier, inputDataValue)
	var emptyOutput map[string]VMInfo
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), vMListOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(map[string]VMInfo), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), vIface.errorBindingMap[methodResult.Error().Name()])
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}


func (vIface *DefaultVMClient) Invoke(ctx *core.ExecutionContext, methodId core.MethodIdentifier, inputDataValue data.DataValue) core.MethodResult {
	methodResult := vIface.connector.GetApiProvider().Invoke(vIface.interfaceName, methodId.Name(), inputDataValue, ctx)
	return methodResult
}


func (vIface *DefaultVMClient) listMethodDefinition() *core.MethodDefinition {
	interfaceIdentifier := core.NewInterfaceIdentifier(vIface.interfaceName)
	typeConverter := vIface.connector.TypeConverter()

	input, inputError := typeConverter.ConvertToDataDefinition(vMListInputType())
	output, outputError := typeConverter.ConvertToDataDefinition(vMListOutputType())
	if inputError != nil {
		log.Errorf("Error in ConvertToDataDefinition for DefaultVMClient.list method's input - %s",
			bindings.VAPIerrorsToError(inputError).Error())
		return nil
	}
	if outputError != nil {
		log.Errorf("Error in ConvertToDataDefinition for DefaultVMClient.list method's output - %s",
			bindings.VAPIerrorsToError(outputError).Error())
		return nil
	}
	methodIdentifier := core.NewMethodIdentifier(interfaceIdentifier, "list")
	errorDefinitions := make([]data.ErrorDefinition, 0)
	vIface.errorBindingMap[errors.Error{}.Error()] = errors.ErrorBindingType()
	errDef1, errError1 := typeConverter.ConvertToDataDefinition(errors.ErrorBindingType())
	if errError1 != nil {
		log.Errorf("Error in ConvertToDataDefinition for DefaultVMClient.list method's errors.Error error - %s",
			bindings.VAPIerrorsToError(errError1).Error())
		return nil
	}
	errorDefinitions = append(errorDefinitions, errDef1.(data.ErrorDefinition))
	vIface.errorBindingMap[errors.InvalidArgument{}.Error()] = errors.InvalidArgumentBindingType()
	errDef2, errError2 := typeConverter.ConvertToDataDefinition(errors.InvalidArgumentBindingType())
	if errError2 != nil {
		log.Errorf("Error in ConvertToDataDefinition for DefaultVMClient.list method's errors.InvalidArgument error - %s",
			bindings.VAPIerrorsToError(errError2).Error())
		return nil
	}
	errorDefinitions = append(errorDefinitions, errDef2.(data.ErrorDefinition))
	vIface.errorBindingMap[errors.ServiceUnavailable{}.Error()] = errors.ServiceUnavailableBindingType()
	errDef3, errError3 := typeConverter.ConvertToDataDefinition(errors.ServiceUnavailableBindingType())
	if errError3 != nil {
		log.Errorf("Error in ConvertToDataDefinition for DefaultVMClient.list method's errors.ServiceUnavailable error - %s",
			bindings.VAPIerrorsToError(errError3).Error())
		return nil
	}
	errorDefinitions = append(errorDefinitions, errDef3.(data.ErrorDefinition))
	vIface.errorBindingMap[errors.Unauthenticated{}.Error()] = errors.UnauthenticatedBindingType()
	errDef4, errError4 := typeConverter.ConvertToDataDefinition(errors.UnauthenticatedBindingType())
	if errError4 != nil {
		log.Errorf("Error in ConvertToDataDefinition for DefaultVMClient.list method's errors.Unauthenticated error - %s",
			bindings.VAPIerrorsToError(errError4).Error())
		return nil
	}
	errorDefinitions = append(errorDefinitions, errDef4.(data.ErrorDefinition))
	vIface.errorBindingMap[errors.Unauthorized{}.Error()] = errors.UnauthorizedBindingType()
	errDef5, errError5 := typeConverter.ConvertToDataDefinition(errors.UnauthorizedBindingType())
	if errError5 != nil {
		log.Errorf("Error in ConvertToDataDefinition for DefaultVMClient.list method's errors.Unauthorized error - %s",
			bindings.VAPIerrorsToError(errError5).Error())
		return nil
	}
	errorDefinitions = append(errorDefinitions, errDef5.(data.ErrorDefinition))
	vIface.errorBindingMap[errors.Unsupported{}.Error()] = errors.UnsupportedBindingType()
	errDef6, errError6 := typeConverter.ConvertToDataDefinition(errors.UnsupportedBindingType())
	if errError6 != nil {
		log.Errorf("Error in ConvertToDataDefinition for DefaultVMClient.list method's errors.Unsupported error - %s",
			bindings.VAPIerrorsToError(errError6).Error())
		return nil
	}
	errorDefinitions = append(errorDefinitions, errDef6.(data.ErrorDefinition))
	vIface.errorBindingMap[errors.UnableToAllocateResource{}.Error()] = errors.UnableToAllocateResourceBindingType()
	errDef7, errError7 := typeConverter.ConvertToDataDefinition(errors.UnableToAllocateResourceBindingType())
	if errError7 != nil {
		log.Errorf("Error in ConvertToDataDefinition for DefaultVMClient.list method's errors.UnableToAllocateResource error - %s",
			bindings.VAPIerrorsToError(errError7).Error())
		return nil
	}
	errorDefinitions = append(errorDefinitions, errDef7.(data.ErrorDefinition))

	methodDefinition := core.NewMethodDefinition(methodIdentifier, input, output, errorDefinitions)
	return &methodDefinition
}
