
/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Client stubs for service: BaseImages
 * Functions that implement the generated BaseImagesClient interface
 */

package esx

import (
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/bindings/com/vmware/vapi/std/errors"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/bindings"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/core"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/data"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/lib"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/log"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/protocol/client"
)

type BaseImagesClientImpl struct {
	interfaceName       string
	interfaceDefinition core.InterfaceDefinition
	methodIdentifiers   []core.MethodIdentifier
	methodNameToDefMap  map[string]*core.MethodDefinition
	errorBindingMap     map[string]bindings.BindingType
	interfaceIdentifier core.InterfaceIdentifier
	connector           client.Connector
}

func NewBaseImagesClientImpl(connector client.Connector) *BaseImagesClientImpl {
	interfaceName := "com.vmware.esx.attestation.os.esx.base_images"
	interfaceIdentifier := core.NewInterfaceIdentifier(interfaceName)
	methodIdentifiers := []core.MethodIdentifier{
		core.NewMethodIdentifier(interfaceIdentifier, "importFromImgdb"),
		core.NewMethodIdentifier(interfaceIdentifier, "list"),
		core.NewMethodIdentifier(interfaceIdentifier, "delete"),
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
	bIface := BaseImagesClientImpl{interfaceName: interfaceName, methodIdentifiers: methodIdentifiers, interfaceDefinition: interfaceDefinition, errorBindingMap: errorBindingMap, interfaceIdentifier: interfaceIdentifier, connector: connector}
	bIface.methodNameToDefMap = make(map[string]*core.MethodDefinition)
	bIface.methodNameToDefMap["import_from_imgdb"] = bIface.importFromImgdbMethodDefinition()
	bIface.methodNameToDefMap["list"] = bIface.listMethodDefinition()
	bIface.methodNameToDefMap["delete"] = bIface.deleteMethodDefinition()
	bIface.methodNameToDefMap["get"] = bIface.getMethodDefinition()
	return &bIface
}

func (bIface *BaseImagesClientImpl) ImportFromImgdb(imgdbParam []byte) (string, error) {
	typeConverter := bIface.connector.TypeConverter()
	methodIdentifier := core.NewMethodIdentifier(bIface.interfaceIdentifier, "import_from_imgdb")
	sv := bindings.NewStructValueBuilder(baseImagesImportFromImgdbInputType(), typeConverter)
	sv.AddStructField("Imgdb", imgdbParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput string
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := baseImagesImportFromImgdbRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	bIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := bIface.Invoke(bIface.connector.NewExecutionContext(), methodIdentifier, inputDataValue)
	var emptyOutput string
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), baseImagesImportFromImgdbOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(string), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), bIface.errorBindingMap[methodResult.Error().Name()])
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (bIface *BaseImagesClientImpl) List() ([]BaseImagesSummary, error) {
	typeConverter := bIface.connector.TypeConverter()
	methodIdentifier := core.NewMethodIdentifier(bIface.interfaceIdentifier, "list")
	sv := bindings.NewStructValueBuilder(baseImagesListInputType(), typeConverter)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput []BaseImagesSummary
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := baseImagesListRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	bIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := bIface.Invoke(bIface.connector.NewExecutionContext(), methodIdentifier, inputDataValue)
	var emptyOutput []BaseImagesSummary
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), baseImagesListOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.([]BaseImagesSummary), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), bIface.errorBindingMap[methodResult.Error().Name()])
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (bIface *BaseImagesClientImpl) Delete(versionParam string) error {
	typeConverter := bIface.connector.TypeConverter()
	methodIdentifier := core.NewMethodIdentifier(bIface.interfaceIdentifier, "delete")
	sv := bindings.NewStructValueBuilder(baseImagesDeleteInputType(), typeConverter)
	sv.AddStructField("Version", versionParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		return bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := baseImagesDeleteRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	bIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := bIface.Invoke(bIface.connector.NewExecutionContext(), methodIdentifier, inputDataValue)
	if methodResult.IsSuccess() {
		return nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), bIface.errorBindingMap[methodResult.Error().Name()])
		if errorInError != nil {
			return bindings.VAPIerrorsToError(errorInError)
		}
		return methodError.(error)
	}
}

func (bIface *BaseImagesClientImpl) Get(versionParam string) (BaseImagesInfo, error) {
	typeConverter := bIface.connector.TypeConverter()
	methodIdentifier := core.NewMethodIdentifier(bIface.interfaceIdentifier, "get")
	sv := bindings.NewStructValueBuilder(baseImagesGetInputType(), typeConverter)
	sv.AddStructField("Version", versionParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput BaseImagesInfo
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := baseImagesGetRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	bIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := bIface.Invoke(bIface.connector.NewExecutionContext(), methodIdentifier, inputDataValue)
	var emptyOutput BaseImagesInfo
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), baseImagesGetOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(BaseImagesInfo), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), bIface.errorBindingMap[methodResult.Error().Name()])
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}


func (bIface *BaseImagesClientImpl) Invoke(ctx *core.ExecutionContext, methodId core.MethodIdentifier, inputDataValue data.DataValue) core.MethodResult {
	methodResult := bIface.connector.GetApiProvider().Invoke(bIface.interfaceName, methodId.Name(), inputDataValue, ctx)
	return methodResult
}


func (bIface *BaseImagesClientImpl) importFromImgdbMethodDefinition() *core.MethodDefinition {
	interfaceIdentifier := core.NewInterfaceIdentifier(bIface.interfaceName)
	typeConverter := bIface.connector.TypeConverter()

	input, inputError := typeConverter.ConvertToDataDefinition(baseImagesImportFromImgdbInputType())
	output, outputError := typeConverter.ConvertToDataDefinition(baseImagesImportFromImgdbOutputType())
	if inputError != nil {
		log.Errorf("Error in ConvertToDataDefinition for BaseImagesClientImpl.importFromImgdb method's input - %s",
			bindings.VAPIerrorsToError(inputError).Error())
		return nil
	}
	if outputError != nil {
		log.Errorf("Error in ConvertToDataDefinition for BaseImagesClientImpl.importFromImgdb method's output - %s",
			bindings.VAPIerrorsToError(outputError).Error())
		return nil
	}
	methodIdentifier := core.NewMethodIdentifier(interfaceIdentifier, "importFromImgdb")
	errorDefinitions := make([]data.ErrorDefinition, 0)
	bIface.errorBindingMap[errors.AlreadyExists{}.Error()] = errors.AlreadyExistsBindingType()
	errDef1, errError1 := typeConverter.ConvertToDataDefinition(errors.AlreadyExistsBindingType())
	if errError1 != nil {
		log.Errorf("Error in ConvertToDataDefinition for BaseImagesClientImpl.importFromImgdb method's errors.AlreadyExists error - %s",
			bindings.VAPIerrorsToError(errError1).Error())
		return nil
	}
	errorDefinitions = append(errorDefinitions, errDef1.(data.ErrorDefinition))
	bIface.errorBindingMap[errors.Error{}.Error()] = errors.ErrorBindingType()
	errDef2, errError2 := typeConverter.ConvertToDataDefinition(errors.ErrorBindingType())
	if errError2 != nil {
		log.Errorf("Error in ConvertToDataDefinition for BaseImagesClientImpl.importFromImgdb method's errors.Error error - %s",
			bindings.VAPIerrorsToError(errError2).Error())
		return nil
	}
	errorDefinitions = append(errorDefinitions, errDef2.(data.ErrorDefinition))
	bIface.errorBindingMap[errors.InvalidArgument{}.Error()] = errors.InvalidArgumentBindingType()
	errDef3, errError3 := typeConverter.ConvertToDataDefinition(errors.InvalidArgumentBindingType())
	if errError3 != nil {
		log.Errorf("Error in ConvertToDataDefinition for BaseImagesClientImpl.importFromImgdb method's errors.InvalidArgument error - %s",
			bindings.VAPIerrorsToError(errError3).Error())
		return nil
	}
	errorDefinitions = append(errorDefinitions, errDef3.(data.ErrorDefinition))
	bIface.errorBindingMap[errors.Unauthenticated{}.Error()] = errors.UnauthenticatedBindingType()
	errDef4, errError4 := typeConverter.ConvertToDataDefinition(errors.UnauthenticatedBindingType())
	if errError4 != nil {
		log.Errorf("Error in ConvertToDataDefinition for BaseImagesClientImpl.importFromImgdb method's errors.Unauthenticated error - %s",
			bindings.VAPIerrorsToError(errError4).Error())
		return nil
	}
	errorDefinitions = append(errorDefinitions, errDef4.(data.ErrorDefinition))
	bIface.errorBindingMap[errors.Unauthorized{}.Error()] = errors.UnauthorizedBindingType()
	errDef5, errError5 := typeConverter.ConvertToDataDefinition(errors.UnauthorizedBindingType())
	if errError5 != nil {
		log.Errorf("Error in ConvertToDataDefinition for BaseImagesClientImpl.importFromImgdb method's errors.Unauthorized error - %s",
			bindings.VAPIerrorsToError(errError5).Error())
		return nil
	}
	errorDefinitions = append(errorDefinitions, errDef5.(data.ErrorDefinition))

	methodDefinition := core.NewMethodDefinition(methodIdentifier, input, output, errorDefinitions)
	return &methodDefinition
}

func (bIface *BaseImagesClientImpl) listMethodDefinition() *core.MethodDefinition {
	interfaceIdentifier := core.NewInterfaceIdentifier(bIface.interfaceName)
	typeConverter := bIface.connector.TypeConverter()

	input, inputError := typeConverter.ConvertToDataDefinition(baseImagesListInputType())
	output, outputError := typeConverter.ConvertToDataDefinition(baseImagesListOutputType())
	if inputError != nil {
		log.Errorf("Error in ConvertToDataDefinition for BaseImagesClientImpl.list method's input - %s",
			bindings.VAPIerrorsToError(inputError).Error())
		return nil
	}
	if outputError != nil {
		log.Errorf("Error in ConvertToDataDefinition for BaseImagesClientImpl.list method's output - %s",
			bindings.VAPIerrorsToError(outputError).Error())
		return nil
	}
	methodIdentifier := core.NewMethodIdentifier(interfaceIdentifier, "list")
	errorDefinitions := make([]data.ErrorDefinition, 0)
	bIface.errorBindingMap[errors.Error{}.Error()] = errors.ErrorBindingType()
	errDef1, errError1 := typeConverter.ConvertToDataDefinition(errors.ErrorBindingType())
	if errError1 != nil {
		log.Errorf("Error in ConvertToDataDefinition for BaseImagesClientImpl.list method's errors.Error error - %s",
			bindings.VAPIerrorsToError(errError1).Error())
		return nil
	}
	errorDefinitions = append(errorDefinitions, errDef1.(data.ErrorDefinition))
	bIface.errorBindingMap[errors.Unauthenticated{}.Error()] = errors.UnauthenticatedBindingType()
	errDef2, errError2 := typeConverter.ConvertToDataDefinition(errors.UnauthenticatedBindingType())
	if errError2 != nil {
		log.Errorf("Error in ConvertToDataDefinition for BaseImagesClientImpl.list method's errors.Unauthenticated error - %s",
			bindings.VAPIerrorsToError(errError2).Error())
		return nil
	}
	errorDefinitions = append(errorDefinitions, errDef2.(data.ErrorDefinition))
	bIface.errorBindingMap[errors.Unauthorized{}.Error()] = errors.UnauthorizedBindingType()
	errDef3, errError3 := typeConverter.ConvertToDataDefinition(errors.UnauthorizedBindingType())
	if errError3 != nil {
		log.Errorf("Error in ConvertToDataDefinition for BaseImagesClientImpl.list method's errors.Unauthorized error - %s",
			bindings.VAPIerrorsToError(errError3).Error())
		return nil
	}
	errorDefinitions = append(errorDefinitions, errDef3.(data.ErrorDefinition))

	methodDefinition := core.NewMethodDefinition(methodIdentifier, input, output, errorDefinitions)
	return &methodDefinition
}

func (bIface *BaseImagesClientImpl) deleteMethodDefinition() *core.MethodDefinition {
	interfaceIdentifier := core.NewInterfaceIdentifier(bIface.interfaceName)
	typeConverter := bIface.connector.TypeConverter()

	input, inputError := typeConverter.ConvertToDataDefinition(baseImagesDeleteInputType())
	output, outputError := typeConverter.ConvertToDataDefinition(baseImagesDeleteOutputType())
	if inputError != nil {
		log.Errorf("Error in ConvertToDataDefinition for BaseImagesClientImpl.delete method's input - %s",
			bindings.VAPIerrorsToError(inputError).Error())
		return nil
	}
	if outputError != nil {
		log.Errorf("Error in ConvertToDataDefinition for BaseImagesClientImpl.delete method's output - %s",
			bindings.VAPIerrorsToError(outputError).Error())
		return nil
	}
	methodIdentifier := core.NewMethodIdentifier(interfaceIdentifier, "delete")
	errorDefinitions := make([]data.ErrorDefinition, 0)
	bIface.errorBindingMap[errors.Error{}.Error()] = errors.ErrorBindingType()
	errDef1, errError1 := typeConverter.ConvertToDataDefinition(errors.ErrorBindingType())
	if errError1 != nil {
		log.Errorf("Error in ConvertToDataDefinition for BaseImagesClientImpl.delete method's errors.Error error - %s",
			bindings.VAPIerrorsToError(errError1).Error())
		return nil
	}
	errorDefinitions = append(errorDefinitions, errDef1.(data.ErrorDefinition))
	bIface.errorBindingMap[errors.InvalidArgument{}.Error()] = errors.InvalidArgumentBindingType()
	errDef2, errError2 := typeConverter.ConvertToDataDefinition(errors.InvalidArgumentBindingType())
	if errError2 != nil {
		log.Errorf("Error in ConvertToDataDefinition for BaseImagesClientImpl.delete method's errors.InvalidArgument error - %s",
			bindings.VAPIerrorsToError(errError2).Error())
		return nil
	}
	errorDefinitions = append(errorDefinitions, errDef2.(data.ErrorDefinition))
	bIface.errorBindingMap[errors.NotFound{}.Error()] = errors.NotFoundBindingType()
	errDef3, errError3 := typeConverter.ConvertToDataDefinition(errors.NotFoundBindingType())
	if errError3 != nil {
		log.Errorf("Error in ConvertToDataDefinition for BaseImagesClientImpl.delete method's errors.NotFound error - %s",
			bindings.VAPIerrorsToError(errError3).Error())
		return nil
	}
	errorDefinitions = append(errorDefinitions, errDef3.(data.ErrorDefinition))
	bIface.errorBindingMap[errors.Unauthenticated{}.Error()] = errors.UnauthenticatedBindingType()
	errDef4, errError4 := typeConverter.ConvertToDataDefinition(errors.UnauthenticatedBindingType())
	if errError4 != nil {
		log.Errorf("Error in ConvertToDataDefinition for BaseImagesClientImpl.delete method's errors.Unauthenticated error - %s",
			bindings.VAPIerrorsToError(errError4).Error())
		return nil
	}
	errorDefinitions = append(errorDefinitions, errDef4.(data.ErrorDefinition))
	bIface.errorBindingMap[errors.Unauthorized{}.Error()] = errors.UnauthorizedBindingType()
	errDef5, errError5 := typeConverter.ConvertToDataDefinition(errors.UnauthorizedBindingType())
	if errError5 != nil {
		log.Errorf("Error in ConvertToDataDefinition for BaseImagesClientImpl.delete method's errors.Unauthorized error - %s",
			bindings.VAPIerrorsToError(errError5).Error())
		return nil
	}
	errorDefinitions = append(errorDefinitions, errDef5.(data.ErrorDefinition))

	methodDefinition := core.NewMethodDefinition(methodIdentifier, input, output, errorDefinitions)
	return &methodDefinition
}

func (bIface *BaseImagesClientImpl) getMethodDefinition() *core.MethodDefinition {
	interfaceIdentifier := core.NewInterfaceIdentifier(bIface.interfaceName)
	typeConverter := bIface.connector.TypeConverter()

	input, inputError := typeConverter.ConvertToDataDefinition(baseImagesGetInputType())
	output, outputError := typeConverter.ConvertToDataDefinition(baseImagesGetOutputType())
	if inputError != nil {
		log.Errorf("Error in ConvertToDataDefinition for BaseImagesClientImpl.get method's input - %s",
			bindings.VAPIerrorsToError(inputError).Error())
		return nil
	}
	if outputError != nil {
		log.Errorf("Error in ConvertToDataDefinition for BaseImagesClientImpl.get method's output - %s",
			bindings.VAPIerrorsToError(outputError).Error())
		return nil
	}
	methodIdentifier := core.NewMethodIdentifier(interfaceIdentifier, "get")
	errorDefinitions := make([]data.ErrorDefinition, 0)
	bIface.errorBindingMap[errors.Error{}.Error()] = errors.ErrorBindingType()
	errDef1, errError1 := typeConverter.ConvertToDataDefinition(errors.ErrorBindingType())
	if errError1 != nil {
		log.Errorf("Error in ConvertToDataDefinition for BaseImagesClientImpl.get method's errors.Error error - %s",
			bindings.VAPIerrorsToError(errError1).Error())
		return nil
	}
	errorDefinitions = append(errorDefinitions, errDef1.(data.ErrorDefinition))
	bIface.errorBindingMap[errors.InvalidArgument{}.Error()] = errors.InvalidArgumentBindingType()
	errDef2, errError2 := typeConverter.ConvertToDataDefinition(errors.InvalidArgumentBindingType())
	if errError2 != nil {
		log.Errorf("Error in ConvertToDataDefinition for BaseImagesClientImpl.get method's errors.InvalidArgument error - %s",
			bindings.VAPIerrorsToError(errError2).Error())
		return nil
	}
	errorDefinitions = append(errorDefinitions, errDef2.(data.ErrorDefinition))
	bIface.errorBindingMap[errors.NotFound{}.Error()] = errors.NotFoundBindingType()
	errDef3, errError3 := typeConverter.ConvertToDataDefinition(errors.NotFoundBindingType())
	if errError3 != nil {
		log.Errorf("Error in ConvertToDataDefinition for BaseImagesClientImpl.get method's errors.NotFound error - %s",
			bindings.VAPIerrorsToError(errError3).Error())
		return nil
	}
	errorDefinitions = append(errorDefinitions, errDef3.(data.ErrorDefinition))
	bIface.errorBindingMap[errors.Unauthenticated{}.Error()] = errors.UnauthenticatedBindingType()
	errDef4, errError4 := typeConverter.ConvertToDataDefinition(errors.UnauthenticatedBindingType())
	if errError4 != nil {
		log.Errorf("Error in ConvertToDataDefinition for BaseImagesClientImpl.get method's errors.Unauthenticated error - %s",
			bindings.VAPIerrorsToError(errError4).Error())
		return nil
	}
	errorDefinitions = append(errorDefinitions, errDef4.(data.ErrorDefinition))
	bIface.errorBindingMap[errors.Unauthorized{}.Error()] = errors.UnauthorizedBindingType()
	errDef5, errError5 := typeConverter.ConvertToDataDefinition(errors.UnauthorizedBindingType())
	if errError5 != nil {
		log.Errorf("Error in ConvertToDataDefinition for BaseImagesClientImpl.get method's errors.Unauthorized error - %s",
			bindings.VAPIerrorsToError(errError5).Error())
		return nil
	}
	errorDefinitions = append(errorDefinitions, errDef5.(data.ErrorDefinition))

	methodDefinition := core.NewMethodDefinition(methodIdentifier, input, output, errorDefinitions)
	return &methodDefinition
}
