
/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Code generated. DO NOT EDIT.

/*
 * Client stubs for service: KeyOperation
 * Functions that implement the generated KeyOperationClient interface
 */


package kms

import (
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/bindings/vapi/std/errors"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/bindings"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/core"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/data"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/lib"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/log"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/protocol/client"
)

type DefaultKeyOperationClient struct {
	interfaceName       string
	interfaceDefinition core.InterfaceDefinition
	methodIdentifiers   []core.MethodIdentifier
	methodNameToDefMap  map[string]*core.MethodDefinition
	errorBindingMap     map[string]bindings.BindingType
	interfaceIdentifier core.InterfaceIdentifier
	connector           client.Connector
}

func NewDefaultKeyOperationClient(connector client.Connector) *DefaultKeyOperationClient {
	interfaceName := "com.vmware.esx.kms.key_operation"
	interfaceIdentifier := core.NewInterfaceIdentifier(interfaceName)
	methodIdentifiers := []core.MethodIdentifier{
		core.NewMethodIdentifier(interfaceIdentifier, "generateKey"),
		core.NewMethodIdentifier(interfaceIdentifier, "encrypt"),
		core.NewMethodIdentifier(interfaceIdentifier, "decrypt"),
	}
	interfaceDefinition := core.NewInterfaceDefinition(interfaceIdentifier, methodIdentifiers)
	errorBindingMap := make(map[string]bindings.BindingType)
	errorBindingMap[errors.InternalServerError{}.Error()] = errors.InternalServerErrorBindingType()
	errorBindingMap[errors.InvalidArgument{}.Error()] = errors.InvalidArgumentBindingType()
	errorBindingMap[errors.OperationNotFound{}.Error()] = errors.OperationNotFoundBindingType()
	errorBindingMap[errors.UnexpectedInput{}.Error()] = errors.UnexpectedInputBindingType()
	errorBindingMap[errors.ServiceUnavailable{}.Error()] = errors.ServiceUnavailableBindingType()
	errorBindingMap[errors.TimedOut{}.Error()] = errors.TimedOutBindingType()
	kIface := DefaultKeyOperationClient{interfaceName: interfaceName, methodIdentifiers: methodIdentifiers, interfaceDefinition: interfaceDefinition, errorBindingMap: errorBindingMap, interfaceIdentifier: interfaceIdentifier, connector: connector}
	kIface.methodNameToDefMap = make(map[string]*core.MethodDefinition)
	kIface.methodNameToDefMap["generate_key"] = kIface.generateKeyMethodDefinition()
	kIface.methodNameToDefMap["encrypt"] = kIface.encryptMethodDefinition()
	kIface.methodNameToDefMap["decrypt"] = kIface.decryptMethodDefinition()
	return &kIface
}

func (kIface *DefaultKeyOperationClient) GenerateKey(providerParam string, numOfBytesParam int64) (KeyOperationGeneratedKey, error) {
	typeConverter := kIface.connector.TypeConverter()
	methodIdentifier := core.NewMethodIdentifier(kIface.interfaceIdentifier, "generate_key")
	sv := bindings.NewStructValueBuilder(keyOperationGenerateKeyInputType(), typeConverter)
	sv.AddStructField("Provider", providerParam)
	sv.AddStructField("NumOfBytes", numOfBytesParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput KeyOperationGeneratedKey
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := keyOperationGenerateKeyRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	kIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := kIface.Invoke(kIface.connector.NewExecutionContext(), methodIdentifier, inputDataValue)
	var emptyOutput KeyOperationGeneratedKey
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), keyOperationGenerateKeyOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(KeyOperationGeneratedKey), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), kIface.errorBindingMap[methodResult.Error().Name()])
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (kIface *DefaultKeyOperationClient) Encrypt(providerParam string, plaintextParam string) (KeyOperationEncryptResult, error) {
	typeConverter := kIface.connector.TypeConverter()
	methodIdentifier := core.NewMethodIdentifier(kIface.interfaceIdentifier, "encrypt")
	sv := bindings.NewStructValueBuilder(keyOperationEncryptInputType(), typeConverter)
	sv.AddStructField("Provider", providerParam)
	sv.AddStructField("Plaintext", plaintextParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput KeyOperationEncryptResult
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := keyOperationEncryptRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	kIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := kIface.Invoke(kIface.connector.NewExecutionContext(), methodIdentifier, inputDataValue)
	var emptyOutput KeyOperationEncryptResult
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), keyOperationEncryptOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(KeyOperationEncryptResult), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), kIface.errorBindingMap[methodResult.Error().Name()])
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (kIface *DefaultKeyOperationClient) Decrypt(providerParam string, ciphertextParam string) (KeyOperationDecryptResult, error) {
	typeConverter := kIface.connector.TypeConverter()
	methodIdentifier := core.NewMethodIdentifier(kIface.interfaceIdentifier, "decrypt")
	sv := bindings.NewStructValueBuilder(keyOperationDecryptInputType(), typeConverter)
	sv.AddStructField("Provider", providerParam)
	sv.AddStructField("Ciphertext", ciphertextParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput KeyOperationDecryptResult
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := keyOperationDecryptRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	kIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := kIface.Invoke(kIface.connector.NewExecutionContext(), methodIdentifier, inputDataValue)
	var emptyOutput KeyOperationDecryptResult
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), keyOperationDecryptOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(KeyOperationDecryptResult), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), kIface.errorBindingMap[methodResult.Error().Name()])
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}


func (kIface *DefaultKeyOperationClient) Invoke(ctx *core.ExecutionContext, methodId core.MethodIdentifier, inputDataValue data.DataValue) core.MethodResult {
	methodResult := kIface.connector.GetApiProvider().Invoke(kIface.interfaceName, methodId.Name(), inputDataValue, ctx)
	return methodResult
}


func (kIface *DefaultKeyOperationClient) generateKeyMethodDefinition() *core.MethodDefinition {
	interfaceIdentifier := core.NewInterfaceIdentifier(kIface.interfaceName)
	typeConverter := kIface.connector.TypeConverter()

	input, inputError := typeConverter.ConvertToDataDefinition(keyOperationGenerateKeyInputType())
	output, outputError := typeConverter.ConvertToDataDefinition(keyOperationGenerateKeyOutputType())
	if inputError != nil {
		log.Errorf("Error in ConvertToDataDefinition for DefaultKeyOperationClient.generateKey method's input - %s",
			bindings.VAPIerrorsToError(inputError).Error())
		return nil
	}
	if outputError != nil {
		log.Errorf("Error in ConvertToDataDefinition for DefaultKeyOperationClient.generateKey method's output - %s",
			bindings.VAPIerrorsToError(outputError).Error())
		return nil
	}
	methodIdentifier := core.NewMethodIdentifier(interfaceIdentifier, "generateKey")
	errorDefinitions := make([]data.ErrorDefinition, 0)
	kIface.errorBindingMap[errors.InvalidArgument{}.Error()] = errors.InvalidArgumentBindingType()
	errDef1, errError1 := typeConverter.ConvertToDataDefinition(errors.InvalidArgumentBindingType())
	if errError1 != nil {
		log.Errorf("Error in ConvertToDataDefinition for DefaultKeyOperationClient.generateKey method's errors.InvalidArgument error - %s",
			bindings.VAPIerrorsToError(errError1).Error())
		return nil
	}
	errorDefinitions = append(errorDefinitions, errDef1.(data.ErrorDefinition))
	kIface.errorBindingMap[errors.NotFound{}.Error()] = errors.NotFoundBindingType()
	errDef2, errError2 := typeConverter.ConvertToDataDefinition(errors.NotFoundBindingType())
	if errError2 != nil {
		log.Errorf("Error in ConvertToDataDefinition for DefaultKeyOperationClient.generateKey method's errors.NotFound error - %s",
			bindings.VAPIerrorsToError(errError2).Error())
		return nil
	}
	errorDefinitions = append(errorDefinitions, errDef2.(data.ErrorDefinition))
	kIface.errorBindingMap[errors.Unauthenticated{}.Error()] = errors.UnauthenticatedBindingType()
	errDef3, errError3 := typeConverter.ConvertToDataDefinition(errors.UnauthenticatedBindingType())
	if errError3 != nil {
		log.Errorf("Error in ConvertToDataDefinition for DefaultKeyOperationClient.generateKey method's errors.Unauthenticated error - %s",
			bindings.VAPIerrorsToError(errError3).Error())
		return nil
	}
	errorDefinitions = append(errorDefinitions, errDef3.(data.ErrorDefinition))
	kIface.errorBindingMap[errors.Unauthorized{}.Error()] = errors.UnauthorizedBindingType()
	errDef4, errError4 := typeConverter.ConvertToDataDefinition(errors.UnauthorizedBindingType())
	if errError4 != nil {
		log.Errorf("Error in ConvertToDataDefinition for DefaultKeyOperationClient.generateKey method's errors.Unauthorized error - %s",
			bindings.VAPIerrorsToError(errError4).Error())
		return nil
	}
	errorDefinitions = append(errorDefinitions, errDef4.(data.ErrorDefinition))
	kIface.errorBindingMap[errors.Error{}.Error()] = errors.ErrorBindingType()
	errDef5, errError5 := typeConverter.ConvertToDataDefinition(errors.ErrorBindingType())
	if errError5 != nil {
		log.Errorf("Error in ConvertToDataDefinition for DefaultKeyOperationClient.generateKey method's errors.Error error - %s",
			bindings.VAPIerrorsToError(errError5).Error())
		return nil
	}
	errorDefinitions = append(errorDefinitions, errDef5.(data.ErrorDefinition))

	methodDefinition := core.NewMethodDefinition(methodIdentifier, input, output, errorDefinitions)
	return &methodDefinition
}

func (kIface *DefaultKeyOperationClient) encryptMethodDefinition() *core.MethodDefinition {
	interfaceIdentifier := core.NewInterfaceIdentifier(kIface.interfaceName)
	typeConverter := kIface.connector.TypeConverter()

	input, inputError := typeConverter.ConvertToDataDefinition(keyOperationEncryptInputType())
	output, outputError := typeConverter.ConvertToDataDefinition(keyOperationEncryptOutputType())
	if inputError != nil {
		log.Errorf("Error in ConvertToDataDefinition for DefaultKeyOperationClient.encrypt method's input - %s",
			bindings.VAPIerrorsToError(inputError).Error())
		return nil
	}
	if outputError != nil {
		log.Errorf("Error in ConvertToDataDefinition for DefaultKeyOperationClient.encrypt method's output - %s",
			bindings.VAPIerrorsToError(outputError).Error())
		return nil
	}
	methodIdentifier := core.NewMethodIdentifier(interfaceIdentifier, "encrypt")
	errorDefinitions := make([]data.ErrorDefinition, 0)
	kIface.errorBindingMap[errors.InvalidArgument{}.Error()] = errors.InvalidArgumentBindingType()
	errDef1, errError1 := typeConverter.ConvertToDataDefinition(errors.InvalidArgumentBindingType())
	if errError1 != nil {
		log.Errorf("Error in ConvertToDataDefinition for DefaultKeyOperationClient.encrypt method's errors.InvalidArgument error - %s",
			bindings.VAPIerrorsToError(errError1).Error())
		return nil
	}
	errorDefinitions = append(errorDefinitions, errDef1.(data.ErrorDefinition))
	kIface.errorBindingMap[errors.NotFound{}.Error()] = errors.NotFoundBindingType()
	errDef2, errError2 := typeConverter.ConvertToDataDefinition(errors.NotFoundBindingType())
	if errError2 != nil {
		log.Errorf("Error in ConvertToDataDefinition for DefaultKeyOperationClient.encrypt method's errors.NotFound error - %s",
			bindings.VAPIerrorsToError(errError2).Error())
		return nil
	}
	errorDefinitions = append(errorDefinitions, errDef2.(data.ErrorDefinition))
	kIface.errorBindingMap[errors.Unauthenticated{}.Error()] = errors.UnauthenticatedBindingType()
	errDef3, errError3 := typeConverter.ConvertToDataDefinition(errors.UnauthenticatedBindingType())
	if errError3 != nil {
		log.Errorf("Error in ConvertToDataDefinition for DefaultKeyOperationClient.encrypt method's errors.Unauthenticated error - %s",
			bindings.VAPIerrorsToError(errError3).Error())
		return nil
	}
	errorDefinitions = append(errorDefinitions, errDef3.(data.ErrorDefinition))
	kIface.errorBindingMap[errors.Unauthorized{}.Error()] = errors.UnauthorizedBindingType()
	errDef4, errError4 := typeConverter.ConvertToDataDefinition(errors.UnauthorizedBindingType())
	if errError4 != nil {
		log.Errorf("Error in ConvertToDataDefinition for DefaultKeyOperationClient.encrypt method's errors.Unauthorized error - %s",
			bindings.VAPIerrorsToError(errError4).Error())
		return nil
	}
	errorDefinitions = append(errorDefinitions, errDef4.(data.ErrorDefinition))
	kIface.errorBindingMap[errors.Error{}.Error()] = errors.ErrorBindingType()
	errDef5, errError5 := typeConverter.ConvertToDataDefinition(errors.ErrorBindingType())
	if errError5 != nil {
		log.Errorf("Error in ConvertToDataDefinition for DefaultKeyOperationClient.encrypt method's errors.Error error - %s",
			bindings.VAPIerrorsToError(errError5).Error())
		return nil
	}
	errorDefinitions = append(errorDefinitions, errDef5.(data.ErrorDefinition))

	methodDefinition := core.NewMethodDefinition(methodIdentifier, input, output, errorDefinitions)
	return &methodDefinition
}

func (kIface *DefaultKeyOperationClient) decryptMethodDefinition() *core.MethodDefinition {
	interfaceIdentifier := core.NewInterfaceIdentifier(kIface.interfaceName)
	typeConverter := kIface.connector.TypeConverter()

	input, inputError := typeConverter.ConvertToDataDefinition(keyOperationDecryptInputType())
	output, outputError := typeConverter.ConvertToDataDefinition(keyOperationDecryptOutputType())
	if inputError != nil {
		log.Errorf("Error in ConvertToDataDefinition for DefaultKeyOperationClient.decrypt method's input - %s",
			bindings.VAPIerrorsToError(inputError).Error())
		return nil
	}
	if outputError != nil {
		log.Errorf("Error in ConvertToDataDefinition for DefaultKeyOperationClient.decrypt method's output - %s",
			bindings.VAPIerrorsToError(outputError).Error())
		return nil
	}
	methodIdentifier := core.NewMethodIdentifier(interfaceIdentifier, "decrypt")
	errorDefinitions := make([]data.ErrorDefinition, 0)
	kIface.errorBindingMap[errors.InvalidArgument{}.Error()] = errors.InvalidArgumentBindingType()
	errDef1, errError1 := typeConverter.ConvertToDataDefinition(errors.InvalidArgumentBindingType())
	if errError1 != nil {
		log.Errorf("Error in ConvertToDataDefinition for DefaultKeyOperationClient.decrypt method's errors.InvalidArgument error - %s",
			bindings.VAPIerrorsToError(errError1).Error())
		return nil
	}
	errorDefinitions = append(errorDefinitions, errDef1.(data.ErrorDefinition))
	kIface.errorBindingMap[errors.NotFound{}.Error()] = errors.NotFoundBindingType()
	errDef2, errError2 := typeConverter.ConvertToDataDefinition(errors.NotFoundBindingType())
	if errError2 != nil {
		log.Errorf("Error in ConvertToDataDefinition for DefaultKeyOperationClient.decrypt method's errors.NotFound error - %s",
			bindings.VAPIerrorsToError(errError2).Error())
		return nil
	}
	errorDefinitions = append(errorDefinitions, errDef2.(data.ErrorDefinition))
	kIface.errorBindingMap[errors.Unauthenticated{}.Error()] = errors.UnauthenticatedBindingType()
	errDef3, errError3 := typeConverter.ConvertToDataDefinition(errors.UnauthenticatedBindingType())
	if errError3 != nil {
		log.Errorf("Error in ConvertToDataDefinition for DefaultKeyOperationClient.decrypt method's errors.Unauthenticated error - %s",
			bindings.VAPIerrorsToError(errError3).Error())
		return nil
	}
	errorDefinitions = append(errorDefinitions, errDef3.(data.ErrorDefinition))
	kIface.errorBindingMap[errors.Unauthorized{}.Error()] = errors.UnauthorizedBindingType()
	errDef4, errError4 := typeConverter.ConvertToDataDefinition(errors.UnauthorizedBindingType())
	if errError4 != nil {
		log.Errorf("Error in ConvertToDataDefinition for DefaultKeyOperationClient.decrypt method's errors.Unauthorized error - %s",
			bindings.VAPIerrorsToError(errError4).Error())
		return nil
	}
	errorDefinitions = append(errorDefinitions, errDef4.(data.ErrorDefinition))
	kIface.errorBindingMap[errors.Error{}.Error()] = errors.ErrorBindingType()
	errDef5, errError5 := typeConverter.ConvertToDataDefinition(errors.ErrorBindingType())
	if errError5 != nil {
		log.Errorf("Error in ConvertToDataDefinition for DefaultKeyOperationClient.decrypt method's errors.Error error - %s",
			bindings.VAPIerrorsToError(errError5).Error())
		return nil
	}
	errorDefinitions = append(errorDefinitions, errDef5.(data.ErrorDefinition))

	methodDefinition := core.NewMethodDefinition(methodIdentifier, input, output, errorDefinitions)
	return &methodDefinition
}
