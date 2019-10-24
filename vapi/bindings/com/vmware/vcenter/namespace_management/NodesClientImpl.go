
/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Client stubs for service: Nodes
 * Functions that implement the generated NodesClient interface
 */

package namespace_management

import (
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/bindings/com/vmware/vapi/std/errors"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/bindings"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/core"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/data"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/lib"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/log"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/protocol/client"
)

type NodesClientImpl struct {
	interfaceName       string
	interfaceDefinition core.InterfaceDefinition
	methodIdentifiers   []core.MethodIdentifier
	methodNameToDefMap  map[string]*core.MethodDefinition
	errorBindingMap     map[string]bindings.BindingType
	interfaceIdentifier core.InterfaceIdentifier
	connector           client.Connector
}

func NewNodesClientImpl(connector client.Connector) *NodesClientImpl {
	interfaceName := "com.vmware.vcenter.namespace_management.nodes"
	interfaceIdentifier := core.NewInterfaceIdentifier(interfaceName)
	methodIdentifiers := []core.MethodIdentifier{
		core.NewMethodIdentifier(interfaceIdentifier, "enterMaintenanceMode"),
		core.NewMethodIdentifier(interfaceIdentifier, "exitMaintenanceMode"),
		core.NewMethodIdentifier(interfaceIdentifier, "remove"),
	}
	interfaceDefinition := core.NewInterfaceDefinition(interfaceIdentifier, methodIdentifiers)
	errorBindingMap := make(map[string]bindings.BindingType)
	errorBindingMap[errors.InternalServerError{}.Error()] = errors.InternalServerErrorBindingType()
	errorBindingMap[errors.InvalidArgument{}.Error()] = errors.InvalidArgumentBindingType()
	errorBindingMap[errors.OperationNotFound{}.Error()] = errors.OperationNotFoundBindingType()
	errorBindingMap[errors.UnexpectedInput{}.Error()] = errors.UnexpectedInputBindingType()
	errorBindingMap[errors.ServiceUnavailable{}.Error()] = errors.ServiceUnavailableBindingType()
	errorBindingMap[errors.TimedOut{}.Error()] = errors.TimedOutBindingType()
	nIface := NodesClientImpl{interfaceName: interfaceName, methodIdentifiers: methodIdentifiers, interfaceDefinition: interfaceDefinition, errorBindingMap: errorBindingMap, interfaceIdentifier: interfaceIdentifier, connector: connector}
	nIface.methodNameToDefMap = make(map[string]*core.MethodDefinition)
	nIface.methodNameToDefMap["enter_maintenance_mode"] = nIface.enterMaintenanceModeMethodDefinition()
	nIface.methodNameToDefMap["exit_maintenance_mode"] = nIface.exitMaintenanceModeMethodDefinition()
	nIface.methodNameToDefMap["remove"] = nIface.removeMethodDefinition()
	return &nIface
}

func (nIface *NodesClientImpl) EnterMaintenanceMode(clusterParam string, nodeParam NodesNodeIdentity, actionParam *NodesMaintenanceActionType) (string, error) {
	typeConverter := nIface.connector.TypeConverter()
	methodIdentifier := core.NewMethodIdentifier(nIface.interfaceIdentifier, "enter_maintenance_mode")
	sv := bindings.NewStructValueBuilder(nodesEnterMaintenanceModeInputType(), typeConverter)
	sv.AddStructField("Cluster", clusterParam)
	sv.AddStructField("Node", nodeParam)
	sv.AddStructField("Action", actionParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput string
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := nodesEnterMaintenanceModeRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	nIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := nIface.Invoke(nIface.connector.NewExecutionContext(), methodIdentifier, inputDataValue)
	var emptyOutput string
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), nodesEnterMaintenanceModeOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(string), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), nIface.errorBindingMap[methodResult.Error().Name()])
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (nIface *NodesClientImpl) ExitMaintenanceMode(clusterParam string, nodeParam NodesNodeIdentity) (string, error) {
	typeConverter := nIface.connector.TypeConverter()
	methodIdentifier := core.NewMethodIdentifier(nIface.interfaceIdentifier, "exit_maintenance_mode")
	sv := bindings.NewStructValueBuilder(nodesExitMaintenanceModeInputType(), typeConverter)
	sv.AddStructField("Cluster", clusterParam)
	sv.AddStructField("Node", nodeParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput string
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := nodesExitMaintenanceModeRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	nIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := nIface.Invoke(nIface.connector.NewExecutionContext(), methodIdentifier, inputDataValue)
	var emptyOutput string
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), nodesExitMaintenanceModeOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(string), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), nIface.errorBindingMap[methodResult.Error().Name()])
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (nIface *NodesClientImpl) Remove(clusterParam string, nodeParam NodesNodeIdentity) (string, error) {
	typeConverter := nIface.connector.TypeConverter()
	methodIdentifier := core.NewMethodIdentifier(nIface.interfaceIdentifier, "remove")
	sv := bindings.NewStructValueBuilder(nodesRemoveInputType(), typeConverter)
	sv.AddStructField("Cluster", clusterParam)
	sv.AddStructField("Node", nodeParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput string
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := nodesRemoveRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	nIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := nIface.Invoke(nIface.connector.NewExecutionContext(), methodIdentifier, inputDataValue)
	var emptyOutput string
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), nodesRemoveOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(string), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), nIface.errorBindingMap[methodResult.Error().Name()])
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}


func (nIface *NodesClientImpl) Invoke(ctx *core.ExecutionContext, methodId core.MethodIdentifier, inputDataValue data.DataValue) core.MethodResult {
	methodResult := nIface.connector.GetApiProvider().Invoke(nIface.interfaceName, methodId.Name(), inputDataValue, ctx)
	return methodResult
}


func (nIface *NodesClientImpl) enterMaintenanceModeMethodDefinition() *core.MethodDefinition {
	interfaceIdentifier := core.NewInterfaceIdentifier(nIface.interfaceName)
	typeConverter := nIface.connector.TypeConverter()

	input, inputError := typeConverter.ConvertToDataDefinition(nodesEnterMaintenanceModeInputType())
	output, outputError := typeConverter.ConvertToDataDefinition(nodesEnterMaintenanceModeOutputType())
	if inputError != nil {
		log.Errorf("Error in ConvertToDataDefinition for NodesClientImpl.enterMaintenanceMode method's input - %s",
			bindings.VAPIerrorsToError(inputError).Error())
		return nil
	}
	if outputError != nil {
		log.Errorf("Error in ConvertToDataDefinition for NodesClientImpl.enterMaintenanceMode method's output - %s",
			bindings.VAPIerrorsToError(outputError).Error())
		return nil
	}
	methodIdentifier := core.NewMethodIdentifier(interfaceIdentifier, "enterMaintenanceMode")
	errorDefinitions := make([]data.ErrorDefinition, 0)
	nIface.errorBindingMap[errors.Error{}.Error()] = errors.ErrorBindingType()
	errDef1, errError1 := typeConverter.ConvertToDataDefinition(errors.ErrorBindingType())
	if errError1 != nil {
		log.Errorf("Error in ConvertToDataDefinition for NodesClientImpl.enterMaintenanceMode method's errors.Error error - %s",
			bindings.VAPIerrorsToError(errError1).Error())
		return nil
	}
	errorDefinitions = append(errorDefinitions, errDef1.(data.ErrorDefinition))
	nIface.errorBindingMap[errors.AlreadyInDesiredState{}.Error()] = errors.AlreadyInDesiredStateBindingType()
	errDef2, errError2 := typeConverter.ConvertToDataDefinition(errors.AlreadyInDesiredStateBindingType())
	if errError2 != nil {
		log.Errorf("Error in ConvertToDataDefinition for NodesClientImpl.enterMaintenanceMode method's errors.AlreadyInDesiredState error - %s",
			bindings.VAPIerrorsToError(errError2).Error())
		return nil
	}
	errorDefinitions = append(errorDefinitions, errDef2.(data.ErrorDefinition))
	nIface.errorBindingMap[errors.NotFound{}.Error()] = errors.NotFoundBindingType()
	errDef3, errError3 := typeConverter.ConvertToDataDefinition(errors.NotFoundBindingType())
	if errError3 != nil {
		log.Errorf("Error in ConvertToDataDefinition for NodesClientImpl.enterMaintenanceMode method's errors.NotFound error - %s",
			bindings.VAPIerrorsToError(errError3).Error())
		return nil
	}
	errorDefinitions = append(errorDefinitions, errDef3.(data.ErrorDefinition))
	nIface.errorBindingMap[errors.Unsupported{}.Error()] = errors.UnsupportedBindingType()
	errDef4, errError4 := typeConverter.ConvertToDataDefinition(errors.UnsupportedBindingType())
	if errError4 != nil {
		log.Errorf("Error in ConvertToDataDefinition for NodesClientImpl.enterMaintenanceMode method's errors.Unsupported error - %s",
			bindings.VAPIerrorsToError(errError4).Error())
		return nil
	}
	errorDefinitions = append(errorDefinitions, errDef4.(data.ErrorDefinition))
	nIface.errorBindingMap[errors.InvalidElementConfiguration{}.Error()] = errors.InvalidElementConfigurationBindingType()
	errDef5, errError5 := typeConverter.ConvertToDataDefinition(errors.InvalidElementConfigurationBindingType())
	if errError5 != nil {
		log.Errorf("Error in ConvertToDataDefinition for NodesClientImpl.enterMaintenanceMode method's errors.InvalidElementConfiguration error - %s",
			bindings.VAPIerrorsToError(errError5).Error())
		return nil
	}
	errorDefinitions = append(errorDefinitions, errDef5.(data.ErrorDefinition))
	nIface.errorBindingMap[errors.Unauthenticated{}.Error()] = errors.UnauthenticatedBindingType()
	errDef6, errError6 := typeConverter.ConvertToDataDefinition(errors.UnauthenticatedBindingType())
	if errError6 != nil {
		log.Errorf("Error in ConvertToDataDefinition for NodesClientImpl.enterMaintenanceMode method's errors.Unauthenticated error - %s",
			bindings.VAPIerrorsToError(errError6).Error())
		return nil
	}
	errorDefinitions = append(errorDefinitions, errDef6.(data.ErrorDefinition))
	nIface.errorBindingMap[errors.Unauthorized{}.Error()] = errors.UnauthorizedBindingType()
	errDef7, errError7 := typeConverter.ConvertToDataDefinition(errors.UnauthorizedBindingType())
	if errError7 != nil {
		log.Errorf("Error in ConvertToDataDefinition for NodesClientImpl.enterMaintenanceMode method's errors.Unauthorized error - %s",
			bindings.VAPIerrorsToError(errError7).Error())
		return nil
	}
	errorDefinitions = append(errorDefinitions, errDef7.(data.ErrorDefinition))

	methodDefinition := core.NewMethodDefinition(methodIdentifier, input, output, errorDefinitions)
	return &methodDefinition
}

func (nIface *NodesClientImpl) exitMaintenanceModeMethodDefinition() *core.MethodDefinition {
	interfaceIdentifier := core.NewInterfaceIdentifier(nIface.interfaceName)
	typeConverter := nIface.connector.TypeConverter()

	input, inputError := typeConverter.ConvertToDataDefinition(nodesExitMaintenanceModeInputType())
	output, outputError := typeConverter.ConvertToDataDefinition(nodesExitMaintenanceModeOutputType())
	if inputError != nil {
		log.Errorf("Error in ConvertToDataDefinition for NodesClientImpl.exitMaintenanceMode method's input - %s",
			bindings.VAPIerrorsToError(inputError).Error())
		return nil
	}
	if outputError != nil {
		log.Errorf("Error in ConvertToDataDefinition for NodesClientImpl.exitMaintenanceMode method's output - %s",
			bindings.VAPIerrorsToError(outputError).Error())
		return nil
	}
	methodIdentifier := core.NewMethodIdentifier(interfaceIdentifier, "exitMaintenanceMode")
	errorDefinitions := make([]data.ErrorDefinition, 0)
	nIface.errorBindingMap[errors.Error{}.Error()] = errors.ErrorBindingType()
	errDef1, errError1 := typeConverter.ConvertToDataDefinition(errors.ErrorBindingType())
	if errError1 != nil {
		log.Errorf("Error in ConvertToDataDefinition for NodesClientImpl.exitMaintenanceMode method's errors.Error error - %s",
			bindings.VAPIerrorsToError(errError1).Error())
		return nil
	}
	errorDefinitions = append(errorDefinitions, errDef1.(data.ErrorDefinition))
	nIface.errorBindingMap[errors.AlreadyInDesiredState{}.Error()] = errors.AlreadyInDesiredStateBindingType()
	errDef2, errError2 := typeConverter.ConvertToDataDefinition(errors.AlreadyInDesiredStateBindingType())
	if errError2 != nil {
		log.Errorf("Error in ConvertToDataDefinition for NodesClientImpl.exitMaintenanceMode method's errors.AlreadyInDesiredState error - %s",
			bindings.VAPIerrorsToError(errError2).Error())
		return nil
	}
	errorDefinitions = append(errorDefinitions, errDef2.(data.ErrorDefinition))
	nIface.errorBindingMap[errors.NotFound{}.Error()] = errors.NotFoundBindingType()
	errDef3, errError3 := typeConverter.ConvertToDataDefinition(errors.NotFoundBindingType())
	if errError3 != nil {
		log.Errorf("Error in ConvertToDataDefinition for NodesClientImpl.exitMaintenanceMode method's errors.NotFound error - %s",
			bindings.VAPIerrorsToError(errError3).Error())
		return nil
	}
	errorDefinitions = append(errorDefinitions, errDef3.(data.ErrorDefinition))
	nIface.errorBindingMap[errors.Unsupported{}.Error()] = errors.UnsupportedBindingType()
	errDef4, errError4 := typeConverter.ConvertToDataDefinition(errors.UnsupportedBindingType())
	if errError4 != nil {
		log.Errorf("Error in ConvertToDataDefinition for NodesClientImpl.exitMaintenanceMode method's errors.Unsupported error - %s",
			bindings.VAPIerrorsToError(errError4).Error())
		return nil
	}
	errorDefinitions = append(errorDefinitions, errDef4.(data.ErrorDefinition))
	nIface.errorBindingMap[errors.InvalidElementConfiguration{}.Error()] = errors.InvalidElementConfigurationBindingType()
	errDef5, errError5 := typeConverter.ConvertToDataDefinition(errors.InvalidElementConfigurationBindingType())
	if errError5 != nil {
		log.Errorf("Error in ConvertToDataDefinition for NodesClientImpl.exitMaintenanceMode method's errors.InvalidElementConfiguration error - %s",
			bindings.VAPIerrorsToError(errError5).Error())
		return nil
	}
	errorDefinitions = append(errorDefinitions, errDef5.(data.ErrorDefinition))
	nIface.errorBindingMap[errors.Unauthenticated{}.Error()] = errors.UnauthenticatedBindingType()
	errDef6, errError6 := typeConverter.ConvertToDataDefinition(errors.UnauthenticatedBindingType())
	if errError6 != nil {
		log.Errorf("Error in ConvertToDataDefinition for NodesClientImpl.exitMaintenanceMode method's errors.Unauthenticated error - %s",
			bindings.VAPIerrorsToError(errError6).Error())
		return nil
	}
	errorDefinitions = append(errorDefinitions, errDef6.(data.ErrorDefinition))
	nIface.errorBindingMap[errors.Unauthorized{}.Error()] = errors.UnauthorizedBindingType()
	errDef7, errError7 := typeConverter.ConvertToDataDefinition(errors.UnauthorizedBindingType())
	if errError7 != nil {
		log.Errorf("Error in ConvertToDataDefinition for NodesClientImpl.exitMaintenanceMode method's errors.Unauthorized error - %s",
			bindings.VAPIerrorsToError(errError7).Error())
		return nil
	}
	errorDefinitions = append(errorDefinitions, errDef7.(data.ErrorDefinition))

	methodDefinition := core.NewMethodDefinition(methodIdentifier, input, output, errorDefinitions)
	return &methodDefinition
}

func (nIface *NodesClientImpl) removeMethodDefinition() *core.MethodDefinition {
	interfaceIdentifier := core.NewInterfaceIdentifier(nIface.interfaceName)
	typeConverter := nIface.connector.TypeConverter()

	input, inputError := typeConverter.ConvertToDataDefinition(nodesRemoveInputType())
	output, outputError := typeConverter.ConvertToDataDefinition(nodesRemoveOutputType())
	if inputError != nil {
		log.Errorf("Error in ConvertToDataDefinition for NodesClientImpl.remove method's input - %s",
			bindings.VAPIerrorsToError(inputError).Error())
		return nil
	}
	if outputError != nil {
		log.Errorf("Error in ConvertToDataDefinition for NodesClientImpl.remove method's output - %s",
			bindings.VAPIerrorsToError(outputError).Error())
		return nil
	}
	methodIdentifier := core.NewMethodIdentifier(interfaceIdentifier, "remove")
	errorDefinitions := make([]data.ErrorDefinition, 0)
	nIface.errorBindingMap[errors.Error{}.Error()] = errors.ErrorBindingType()
	errDef1, errError1 := typeConverter.ConvertToDataDefinition(errors.ErrorBindingType())
	if errError1 != nil {
		log.Errorf("Error in ConvertToDataDefinition for NodesClientImpl.remove method's errors.Error error - %s",
			bindings.VAPIerrorsToError(errError1).Error())
		return nil
	}
	errorDefinitions = append(errorDefinitions, errDef1.(data.ErrorDefinition))
	nIface.errorBindingMap[errors.AlreadyInDesiredState{}.Error()] = errors.AlreadyInDesiredStateBindingType()
	errDef2, errError2 := typeConverter.ConvertToDataDefinition(errors.AlreadyInDesiredStateBindingType())
	if errError2 != nil {
		log.Errorf("Error in ConvertToDataDefinition for NodesClientImpl.remove method's errors.AlreadyInDesiredState error - %s",
			bindings.VAPIerrorsToError(errError2).Error())
		return nil
	}
	errorDefinitions = append(errorDefinitions, errDef2.(data.ErrorDefinition))
	nIface.errorBindingMap[errors.NotFound{}.Error()] = errors.NotFoundBindingType()
	errDef3, errError3 := typeConverter.ConvertToDataDefinition(errors.NotFoundBindingType())
	if errError3 != nil {
		log.Errorf("Error in ConvertToDataDefinition for NodesClientImpl.remove method's errors.NotFound error - %s",
			bindings.VAPIerrorsToError(errError3).Error())
		return nil
	}
	errorDefinitions = append(errorDefinitions, errDef3.(data.ErrorDefinition))
	nIface.errorBindingMap[errors.Unsupported{}.Error()] = errors.UnsupportedBindingType()
	errDef4, errError4 := typeConverter.ConvertToDataDefinition(errors.UnsupportedBindingType())
	if errError4 != nil {
		log.Errorf("Error in ConvertToDataDefinition for NodesClientImpl.remove method's errors.Unsupported error - %s",
			bindings.VAPIerrorsToError(errError4).Error())
		return nil
	}
	errorDefinitions = append(errorDefinitions, errDef4.(data.ErrorDefinition))
	nIface.errorBindingMap[errors.InvalidElementConfiguration{}.Error()] = errors.InvalidElementConfigurationBindingType()
	errDef5, errError5 := typeConverter.ConvertToDataDefinition(errors.InvalidElementConfigurationBindingType())
	if errError5 != nil {
		log.Errorf("Error in ConvertToDataDefinition for NodesClientImpl.remove method's errors.InvalidElementConfiguration error - %s",
			bindings.VAPIerrorsToError(errError5).Error())
		return nil
	}
	errorDefinitions = append(errorDefinitions, errDef5.(data.ErrorDefinition))
	nIface.errorBindingMap[errors.Unauthenticated{}.Error()] = errors.UnauthenticatedBindingType()
	errDef6, errError6 := typeConverter.ConvertToDataDefinition(errors.UnauthenticatedBindingType())
	if errError6 != nil {
		log.Errorf("Error in ConvertToDataDefinition for NodesClientImpl.remove method's errors.Unauthenticated error - %s",
			bindings.VAPIerrorsToError(errError6).Error())
		return nil
	}
	errorDefinitions = append(errorDefinitions, errDef6.(data.ErrorDefinition))
	nIface.errorBindingMap[errors.Unauthorized{}.Error()] = errors.UnauthorizedBindingType()
	errDef7, errError7 := typeConverter.ConvertToDataDefinition(errors.UnauthorizedBindingType())
	if errError7 != nil {
		log.Errorf("Error in ConvertToDataDefinition for NodesClientImpl.remove method's errors.Unauthorized error - %s",
			bindings.VAPIerrorsToError(errError7).Error())
		return nil
	}
	errorDefinitions = append(errorDefinitions, errDef7.(data.ErrorDefinition))

	methodDefinition := core.NewMethodDefinition(methodIdentifier, input, output, errorDefinitions)
	return &methodDefinition
}
