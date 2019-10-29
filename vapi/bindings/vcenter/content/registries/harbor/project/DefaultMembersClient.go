
/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Code generated. DO NOT EDIT.

/*
 * Client stubs for service: Members
 * Functions that implement the generated MembersClient interface
 */


package project

import (
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/bindings/vapi/std/errors"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/bindings"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/core"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/data"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/lib"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/log"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/protocol/client"
)

type MembersClientImpl struct {
	interfaceName       string
	interfaceDefinition core.InterfaceDefinition
	methodIdentifiers   []core.MethodIdentifier
	methodNameToDefMap  map[string]*core.MethodDefinition
	errorBindingMap     map[string]bindings.BindingType
	interfaceIdentifier core.InterfaceIdentifier
	connector           client.Connector
}

func NewMembersClientImpl(connector client.Connector) *MembersClientImpl {
	interfaceName := "com.vmware.vcenter.content.registries.harbor.project.members"
	interfaceIdentifier := core.NewInterfaceIdentifier(interfaceName)
	methodIdentifiers := []core.MethodIdentifier{
		core.NewMethodIdentifier(interfaceIdentifier, "create"),
		core.NewMethodIdentifier(interfaceIdentifier, "delete"),
		core.NewMethodIdentifier(interfaceIdentifier, "update"),
		core.NewMethodIdentifier(interfaceIdentifier, "get"),
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
	mIface := MembersClientImpl{interfaceName: interfaceName, methodIdentifiers: methodIdentifiers, interfaceDefinition: interfaceDefinition, errorBindingMap: errorBindingMap, interfaceIdentifier: interfaceIdentifier, connector: connector}
	mIface.methodNameToDefMap = make(map[string]*core.MethodDefinition)
	mIface.methodNameToDefMap["create"] = mIface.createMethodDefinition()
	mIface.methodNameToDefMap["delete"] = mIface.deleteMethodDefinition()
	mIface.methodNameToDefMap["update"] = mIface.updateMethodDefinition()
	mIface.methodNameToDefMap["get"] = mIface.getMethodDefinition()
	mIface.methodNameToDefMap["list"] = mIface.listMethodDefinition()
	return &mIface
}

func (mIface *MembersClientImpl) Create(registryParam string, projectParam string, specParam MembersCreateSpec) (string, error) {
	typeConverter := mIface.connector.TypeConverter()
	methodIdentifier := core.NewMethodIdentifier(mIface.interfaceIdentifier, "create")
	sv := bindings.NewStructValueBuilder(membersCreateInputType(), typeConverter)
	sv.AddStructField("Registry", registryParam)
	sv.AddStructField("Project", projectParam)
	sv.AddStructField("Spec", specParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput string
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := membersCreateRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	mIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := mIface.Invoke(mIface.connector.NewExecutionContext(), methodIdentifier, inputDataValue)
	var emptyOutput string
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), membersCreateOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(string), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), mIface.errorBindingMap[methodResult.Error().Name()])
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (mIface *MembersClientImpl) Delete(registryParam string, projectParam string, memberParam string) error {
	typeConverter := mIface.connector.TypeConverter()
	methodIdentifier := core.NewMethodIdentifier(mIface.interfaceIdentifier, "delete")
	sv := bindings.NewStructValueBuilder(membersDeleteInputType(), typeConverter)
	sv.AddStructField("Registry", registryParam)
	sv.AddStructField("Project", projectParam)
	sv.AddStructField("Member", memberParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		return bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := membersDeleteRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	mIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := mIface.Invoke(mIface.connector.NewExecutionContext(), methodIdentifier, inputDataValue)
	if methodResult.IsSuccess() {
		return nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), mIface.errorBindingMap[methodResult.Error().Name()])
		if errorInError != nil {
			return bindings.VAPIerrorsToError(errorInError)
		}
		return methodError.(error)
	}
}

func (mIface *MembersClientImpl) Update(registryParam string, projectParam string, memberParam string, specParam MembersUpdateSpec) error {
	typeConverter := mIface.connector.TypeConverter()
	methodIdentifier := core.NewMethodIdentifier(mIface.interfaceIdentifier, "update")
	sv := bindings.NewStructValueBuilder(membersUpdateInputType(), typeConverter)
	sv.AddStructField("Registry", registryParam)
	sv.AddStructField("Project", projectParam)
	sv.AddStructField("Member", memberParam)
	sv.AddStructField("Spec", specParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		return bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := membersUpdateRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	mIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := mIface.Invoke(mIface.connector.NewExecutionContext(), methodIdentifier, inputDataValue)
	if methodResult.IsSuccess() {
		return nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), mIface.errorBindingMap[methodResult.Error().Name()])
		if errorInError != nil {
			return bindings.VAPIerrorsToError(errorInError)
		}
		return methodError.(error)
	}
}

func (mIface *MembersClientImpl) Get(registryParam string, projectParam string, memberParam string) (MembersInfo, error) {
	typeConverter := mIface.connector.TypeConverter()
	methodIdentifier := core.NewMethodIdentifier(mIface.interfaceIdentifier, "get")
	sv := bindings.NewStructValueBuilder(membersGetInputType(), typeConverter)
	sv.AddStructField("Registry", registryParam)
	sv.AddStructField("Project", projectParam)
	sv.AddStructField("Member", memberParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput MembersInfo
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := membersGetRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	mIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := mIface.Invoke(mIface.connector.NewExecutionContext(), methodIdentifier, inputDataValue)
	var emptyOutput MembersInfo
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), membersGetOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(MembersInfo), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), mIface.errorBindingMap[methodResult.Error().Name()])
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (mIface *MembersClientImpl) List(registryParam string, projectParam string) ([]MembersSummary, error) {
	typeConverter := mIface.connector.TypeConverter()
	methodIdentifier := core.NewMethodIdentifier(mIface.interfaceIdentifier, "list")
	sv := bindings.NewStructValueBuilder(membersListInputType(), typeConverter)
	sv.AddStructField("Registry", registryParam)
	sv.AddStructField("Project", projectParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput []MembersSummary
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := membersListRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	mIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := mIface.Invoke(mIface.connector.NewExecutionContext(), methodIdentifier, inputDataValue)
	var emptyOutput []MembersSummary
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), membersListOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.([]MembersSummary), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), mIface.errorBindingMap[methodResult.Error().Name()])
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}


func (mIface *MembersClientImpl) Invoke(ctx *core.ExecutionContext, methodId core.MethodIdentifier, inputDataValue data.DataValue) core.MethodResult {
	methodResult := mIface.connector.GetApiProvider().Invoke(mIface.interfaceName, methodId.Name(), inputDataValue, ctx)
	return methodResult
}


func (mIface *MembersClientImpl) createMethodDefinition() *core.MethodDefinition {
	interfaceIdentifier := core.NewInterfaceIdentifier(mIface.interfaceName)
	typeConverter := mIface.connector.TypeConverter()

	input, inputError := typeConverter.ConvertToDataDefinition(membersCreateInputType())
	output, outputError := typeConverter.ConvertToDataDefinition(membersCreateOutputType())
	if inputError != nil {
		log.Errorf("Error in ConvertToDataDefinition for MembersClientImpl.create method's input - %s",
			bindings.VAPIerrorsToError(inputError).Error())
		return nil
	}
	if outputError != nil {
		log.Errorf("Error in ConvertToDataDefinition for MembersClientImpl.create method's output - %s",
			bindings.VAPIerrorsToError(outputError).Error())
		return nil
	}
	methodIdentifier := core.NewMethodIdentifier(interfaceIdentifier, "create")
	errorDefinitions := make([]data.ErrorDefinition, 0)
	mIface.errorBindingMap[errors.InvalidArgument{}.Error()] = errors.InvalidArgumentBindingType()
	errDef1, errError1 := typeConverter.ConvertToDataDefinition(errors.InvalidArgumentBindingType())
	if errError1 != nil {
		log.Errorf("Error in ConvertToDataDefinition for MembersClientImpl.create method's errors.InvalidArgument error - %s",
			bindings.VAPIerrorsToError(errError1).Error())
		return nil
	}
	errorDefinitions = append(errorDefinitions, errDef1.(data.ErrorDefinition))
	mIface.errorBindingMap[errors.NotAllowedInCurrentState{}.Error()] = errors.NotAllowedInCurrentStateBindingType()
	errDef2, errError2 := typeConverter.ConvertToDataDefinition(errors.NotAllowedInCurrentStateBindingType())
	if errError2 != nil {
		log.Errorf("Error in ConvertToDataDefinition for MembersClientImpl.create method's errors.NotAllowedInCurrentState error - %s",
			bindings.VAPIerrorsToError(errError2).Error())
		return nil
	}
	errorDefinitions = append(errorDefinitions, errDef2.(data.ErrorDefinition))
	mIface.errorBindingMap[errors.NotFound{}.Error()] = errors.NotFoundBindingType()
	errDef3, errError3 := typeConverter.ConvertToDataDefinition(errors.NotFoundBindingType())
	if errError3 != nil {
		log.Errorf("Error in ConvertToDataDefinition for MembersClientImpl.create method's errors.NotFound error - %s",
			bindings.VAPIerrorsToError(errError3).Error())
		return nil
	}
	errorDefinitions = append(errorDefinitions, errDef3.(data.ErrorDefinition))
	mIface.errorBindingMap[errors.AlreadyExists{}.Error()] = errors.AlreadyExistsBindingType()
	errDef4, errError4 := typeConverter.ConvertToDataDefinition(errors.AlreadyExistsBindingType())
	if errError4 != nil {
		log.Errorf("Error in ConvertToDataDefinition for MembersClientImpl.create method's errors.AlreadyExists error - %s",
			bindings.VAPIerrorsToError(errError4).Error())
		return nil
	}
	errorDefinitions = append(errorDefinitions, errDef4.(data.ErrorDefinition))
	mIface.errorBindingMap[errors.Error{}.Error()] = errors.ErrorBindingType()
	errDef5, errError5 := typeConverter.ConvertToDataDefinition(errors.ErrorBindingType())
	if errError5 != nil {
		log.Errorf("Error in ConvertToDataDefinition for MembersClientImpl.create method's errors.Error error - %s",
			bindings.VAPIerrorsToError(errError5).Error())
		return nil
	}
	errorDefinitions = append(errorDefinitions, errDef5.(data.ErrorDefinition))
	mIface.errorBindingMap[errors.Unauthenticated{}.Error()] = errors.UnauthenticatedBindingType()
	errDef6, errError6 := typeConverter.ConvertToDataDefinition(errors.UnauthenticatedBindingType())
	if errError6 != nil {
		log.Errorf("Error in ConvertToDataDefinition for MembersClientImpl.create method's errors.Unauthenticated error - %s",
			bindings.VAPIerrorsToError(errError6).Error())
		return nil
	}
	errorDefinitions = append(errorDefinitions, errDef6.(data.ErrorDefinition))
	mIface.errorBindingMap[errors.Unauthorized{}.Error()] = errors.UnauthorizedBindingType()
	errDef7, errError7 := typeConverter.ConvertToDataDefinition(errors.UnauthorizedBindingType())
	if errError7 != nil {
		log.Errorf("Error in ConvertToDataDefinition for MembersClientImpl.create method's errors.Unauthorized error - %s",
			bindings.VAPIerrorsToError(errError7).Error())
		return nil
	}
	errorDefinitions = append(errorDefinitions, errDef7.(data.ErrorDefinition))

	methodDefinition := core.NewMethodDefinition(methodIdentifier, input, output, errorDefinitions)
	return &methodDefinition
}

func (mIface *MembersClientImpl) deleteMethodDefinition() *core.MethodDefinition {
	interfaceIdentifier := core.NewInterfaceIdentifier(mIface.interfaceName)
	typeConverter := mIface.connector.TypeConverter()

	input, inputError := typeConverter.ConvertToDataDefinition(membersDeleteInputType())
	output, outputError := typeConverter.ConvertToDataDefinition(membersDeleteOutputType())
	if inputError != nil {
		log.Errorf("Error in ConvertToDataDefinition for MembersClientImpl.delete method's input - %s",
			bindings.VAPIerrorsToError(inputError).Error())
		return nil
	}
	if outputError != nil {
		log.Errorf("Error in ConvertToDataDefinition for MembersClientImpl.delete method's output - %s",
			bindings.VAPIerrorsToError(outputError).Error())
		return nil
	}
	methodIdentifier := core.NewMethodIdentifier(interfaceIdentifier, "delete")
	errorDefinitions := make([]data.ErrorDefinition, 0)
	mIface.errorBindingMap[errors.NotAllowedInCurrentState{}.Error()] = errors.NotAllowedInCurrentStateBindingType()
	errDef1, errError1 := typeConverter.ConvertToDataDefinition(errors.NotAllowedInCurrentStateBindingType())
	if errError1 != nil {
		log.Errorf("Error in ConvertToDataDefinition for MembersClientImpl.delete method's errors.NotAllowedInCurrentState error - %s",
			bindings.VAPIerrorsToError(errError1).Error())
		return nil
	}
	errorDefinitions = append(errorDefinitions, errDef1.(data.ErrorDefinition))
	mIface.errorBindingMap[errors.NotFound{}.Error()] = errors.NotFoundBindingType()
	errDef2, errError2 := typeConverter.ConvertToDataDefinition(errors.NotFoundBindingType())
	if errError2 != nil {
		log.Errorf("Error in ConvertToDataDefinition for MembersClientImpl.delete method's errors.NotFound error - %s",
			bindings.VAPIerrorsToError(errError2).Error())
		return nil
	}
	errorDefinitions = append(errorDefinitions, errDef2.(data.ErrorDefinition))
	mIface.errorBindingMap[errors.Error{}.Error()] = errors.ErrorBindingType()
	errDef3, errError3 := typeConverter.ConvertToDataDefinition(errors.ErrorBindingType())
	if errError3 != nil {
		log.Errorf("Error in ConvertToDataDefinition for MembersClientImpl.delete method's errors.Error error - %s",
			bindings.VAPIerrorsToError(errError3).Error())
		return nil
	}
	errorDefinitions = append(errorDefinitions, errDef3.(data.ErrorDefinition))
	mIface.errorBindingMap[errors.Unauthenticated{}.Error()] = errors.UnauthenticatedBindingType()
	errDef4, errError4 := typeConverter.ConvertToDataDefinition(errors.UnauthenticatedBindingType())
	if errError4 != nil {
		log.Errorf("Error in ConvertToDataDefinition for MembersClientImpl.delete method's errors.Unauthenticated error - %s",
			bindings.VAPIerrorsToError(errError4).Error())
		return nil
	}
	errorDefinitions = append(errorDefinitions, errDef4.(data.ErrorDefinition))
	mIface.errorBindingMap[errors.Unauthorized{}.Error()] = errors.UnauthorizedBindingType()
	errDef5, errError5 := typeConverter.ConvertToDataDefinition(errors.UnauthorizedBindingType())
	if errError5 != nil {
		log.Errorf("Error in ConvertToDataDefinition for MembersClientImpl.delete method's errors.Unauthorized error - %s",
			bindings.VAPIerrorsToError(errError5).Error())
		return nil
	}
	errorDefinitions = append(errorDefinitions, errDef5.(data.ErrorDefinition))

	methodDefinition := core.NewMethodDefinition(methodIdentifier, input, output, errorDefinitions)
	return &methodDefinition
}

func (mIface *MembersClientImpl) updateMethodDefinition() *core.MethodDefinition {
	interfaceIdentifier := core.NewInterfaceIdentifier(mIface.interfaceName)
	typeConverter := mIface.connector.TypeConverter()

	input, inputError := typeConverter.ConvertToDataDefinition(membersUpdateInputType())
	output, outputError := typeConverter.ConvertToDataDefinition(membersUpdateOutputType())
	if inputError != nil {
		log.Errorf("Error in ConvertToDataDefinition for MembersClientImpl.update method's input - %s",
			bindings.VAPIerrorsToError(inputError).Error())
		return nil
	}
	if outputError != nil {
		log.Errorf("Error in ConvertToDataDefinition for MembersClientImpl.update method's output - %s",
			bindings.VAPIerrorsToError(outputError).Error())
		return nil
	}
	methodIdentifier := core.NewMethodIdentifier(interfaceIdentifier, "update")
	errorDefinitions := make([]data.ErrorDefinition, 0)
	mIface.errorBindingMap[errors.InvalidArgument{}.Error()] = errors.InvalidArgumentBindingType()
	errDef1, errError1 := typeConverter.ConvertToDataDefinition(errors.InvalidArgumentBindingType())
	if errError1 != nil {
		log.Errorf("Error in ConvertToDataDefinition for MembersClientImpl.update method's errors.InvalidArgument error - %s",
			bindings.VAPIerrorsToError(errError1).Error())
		return nil
	}
	errorDefinitions = append(errorDefinitions, errDef1.(data.ErrorDefinition))
	mIface.errorBindingMap[errors.NotAllowedInCurrentState{}.Error()] = errors.NotAllowedInCurrentStateBindingType()
	errDef2, errError2 := typeConverter.ConvertToDataDefinition(errors.NotAllowedInCurrentStateBindingType())
	if errError2 != nil {
		log.Errorf("Error in ConvertToDataDefinition for MembersClientImpl.update method's errors.NotAllowedInCurrentState error - %s",
			bindings.VAPIerrorsToError(errError2).Error())
		return nil
	}
	errorDefinitions = append(errorDefinitions, errDef2.(data.ErrorDefinition))
	mIface.errorBindingMap[errors.NotFound{}.Error()] = errors.NotFoundBindingType()
	errDef3, errError3 := typeConverter.ConvertToDataDefinition(errors.NotFoundBindingType())
	if errError3 != nil {
		log.Errorf("Error in ConvertToDataDefinition for MembersClientImpl.update method's errors.NotFound error - %s",
			bindings.VAPIerrorsToError(errError3).Error())
		return nil
	}
	errorDefinitions = append(errorDefinitions, errDef3.(data.ErrorDefinition))
	mIface.errorBindingMap[errors.Error{}.Error()] = errors.ErrorBindingType()
	errDef4, errError4 := typeConverter.ConvertToDataDefinition(errors.ErrorBindingType())
	if errError4 != nil {
		log.Errorf("Error in ConvertToDataDefinition for MembersClientImpl.update method's errors.Error error - %s",
			bindings.VAPIerrorsToError(errError4).Error())
		return nil
	}
	errorDefinitions = append(errorDefinitions, errDef4.(data.ErrorDefinition))
	mIface.errorBindingMap[errors.Unauthenticated{}.Error()] = errors.UnauthenticatedBindingType()
	errDef5, errError5 := typeConverter.ConvertToDataDefinition(errors.UnauthenticatedBindingType())
	if errError5 != nil {
		log.Errorf("Error in ConvertToDataDefinition for MembersClientImpl.update method's errors.Unauthenticated error - %s",
			bindings.VAPIerrorsToError(errError5).Error())
		return nil
	}
	errorDefinitions = append(errorDefinitions, errDef5.(data.ErrorDefinition))
	mIface.errorBindingMap[errors.Unauthorized{}.Error()] = errors.UnauthorizedBindingType()
	errDef6, errError6 := typeConverter.ConvertToDataDefinition(errors.UnauthorizedBindingType())
	if errError6 != nil {
		log.Errorf("Error in ConvertToDataDefinition for MembersClientImpl.update method's errors.Unauthorized error - %s",
			bindings.VAPIerrorsToError(errError6).Error())
		return nil
	}
	errorDefinitions = append(errorDefinitions, errDef6.(data.ErrorDefinition))

	methodDefinition := core.NewMethodDefinition(methodIdentifier, input, output, errorDefinitions)
	return &methodDefinition
}

func (mIface *MembersClientImpl) getMethodDefinition() *core.MethodDefinition {
	interfaceIdentifier := core.NewInterfaceIdentifier(mIface.interfaceName)
	typeConverter := mIface.connector.TypeConverter()

	input, inputError := typeConverter.ConvertToDataDefinition(membersGetInputType())
	output, outputError := typeConverter.ConvertToDataDefinition(membersGetOutputType())
	if inputError != nil {
		log.Errorf("Error in ConvertToDataDefinition for MembersClientImpl.get method's input - %s",
			bindings.VAPIerrorsToError(inputError).Error())
		return nil
	}
	if outputError != nil {
		log.Errorf("Error in ConvertToDataDefinition for MembersClientImpl.get method's output - %s",
			bindings.VAPIerrorsToError(outputError).Error())
		return nil
	}
	methodIdentifier := core.NewMethodIdentifier(interfaceIdentifier, "get")
	errorDefinitions := make([]data.ErrorDefinition, 0)
	mIface.errorBindingMap[errors.NotFound{}.Error()] = errors.NotFoundBindingType()
	errDef1, errError1 := typeConverter.ConvertToDataDefinition(errors.NotFoundBindingType())
	if errError1 != nil {
		log.Errorf("Error in ConvertToDataDefinition for MembersClientImpl.get method's errors.NotFound error - %s",
			bindings.VAPIerrorsToError(errError1).Error())
		return nil
	}
	errorDefinitions = append(errorDefinitions, errDef1.(data.ErrorDefinition))
	mIface.errorBindingMap[errors.Error{}.Error()] = errors.ErrorBindingType()
	errDef2, errError2 := typeConverter.ConvertToDataDefinition(errors.ErrorBindingType())
	if errError2 != nil {
		log.Errorf("Error in ConvertToDataDefinition for MembersClientImpl.get method's errors.Error error - %s",
			bindings.VAPIerrorsToError(errError2).Error())
		return nil
	}
	errorDefinitions = append(errorDefinitions, errDef2.(data.ErrorDefinition))
	mIface.errorBindingMap[errors.Unauthenticated{}.Error()] = errors.UnauthenticatedBindingType()
	errDef3, errError3 := typeConverter.ConvertToDataDefinition(errors.UnauthenticatedBindingType())
	if errError3 != nil {
		log.Errorf("Error in ConvertToDataDefinition for MembersClientImpl.get method's errors.Unauthenticated error - %s",
			bindings.VAPIerrorsToError(errError3).Error())
		return nil
	}
	errorDefinitions = append(errorDefinitions, errDef3.(data.ErrorDefinition))
	mIface.errorBindingMap[errors.Unauthorized{}.Error()] = errors.UnauthorizedBindingType()
	errDef4, errError4 := typeConverter.ConvertToDataDefinition(errors.UnauthorizedBindingType())
	if errError4 != nil {
		log.Errorf("Error in ConvertToDataDefinition for MembersClientImpl.get method's errors.Unauthorized error - %s",
			bindings.VAPIerrorsToError(errError4).Error())
		return nil
	}
	errorDefinitions = append(errorDefinitions, errDef4.(data.ErrorDefinition))

	methodDefinition := core.NewMethodDefinition(methodIdentifier, input, output, errorDefinitions)
	return &methodDefinition
}

func (mIface *MembersClientImpl) listMethodDefinition() *core.MethodDefinition {
	interfaceIdentifier := core.NewInterfaceIdentifier(mIface.interfaceName)
	typeConverter := mIface.connector.TypeConverter()

	input, inputError := typeConverter.ConvertToDataDefinition(membersListInputType())
	output, outputError := typeConverter.ConvertToDataDefinition(membersListOutputType())
	if inputError != nil {
		log.Errorf("Error in ConvertToDataDefinition for MembersClientImpl.list method's input - %s",
			bindings.VAPIerrorsToError(inputError).Error())
		return nil
	}
	if outputError != nil {
		log.Errorf("Error in ConvertToDataDefinition for MembersClientImpl.list method's output - %s",
			bindings.VAPIerrorsToError(outputError).Error())
		return nil
	}
	methodIdentifier := core.NewMethodIdentifier(interfaceIdentifier, "list")
	errorDefinitions := make([]data.ErrorDefinition, 0)
	mIface.errorBindingMap[errors.Error{}.Error()] = errors.ErrorBindingType()
	errDef1, errError1 := typeConverter.ConvertToDataDefinition(errors.ErrorBindingType())
	if errError1 != nil {
		log.Errorf("Error in ConvertToDataDefinition for MembersClientImpl.list method's errors.Error error - %s",
			bindings.VAPIerrorsToError(errError1).Error())
		return nil
	}
	errorDefinitions = append(errorDefinitions, errDef1.(data.ErrorDefinition))
	mIface.errorBindingMap[errors.NotFound{}.Error()] = errors.NotFoundBindingType()
	errDef2, errError2 := typeConverter.ConvertToDataDefinition(errors.NotFoundBindingType())
	if errError2 != nil {
		log.Errorf("Error in ConvertToDataDefinition for MembersClientImpl.list method's errors.NotFound error - %s",
			bindings.VAPIerrorsToError(errError2).Error())
		return nil
	}
	errorDefinitions = append(errorDefinitions, errDef2.(data.ErrorDefinition))
	mIface.errorBindingMap[errors.Unauthenticated{}.Error()] = errors.UnauthenticatedBindingType()
	errDef3, errError3 := typeConverter.ConvertToDataDefinition(errors.UnauthenticatedBindingType())
	if errError3 != nil {
		log.Errorf("Error in ConvertToDataDefinition for MembersClientImpl.list method's errors.Unauthenticated error - %s",
			bindings.VAPIerrorsToError(errError3).Error())
		return nil
	}
	errorDefinitions = append(errorDefinitions, errDef3.(data.ErrorDefinition))
	mIface.errorBindingMap[errors.Unauthorized{}.Error()] = errors.UnauthorizedBindingType()
	errDef4, errError4 := typeConverter.ConvertToDataDefinition(errors.UnauthorizedBindingType())
	if errError4 != nil {
		log.Errorf("Error in ConvertToDataDefinition for MembersClientImpl.list method's errors.Unauthorized error - %s",
			bindings.VAPIerrorsToError(errError4).Error())
		return nil
	}
	errorDefinitions = append(errorDefinitions, errDef4.(data.ErrorDefinition))

	methodDefinition := core.NewMethodDefinition(methodIdentifier, input, output, errorDefinitions)
	return &methodDefinition
}
