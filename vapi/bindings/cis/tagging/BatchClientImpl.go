
/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Client stubs for service: Batch
 * Functions that implement the generated BatchClient interface
 */

package tagging

import (
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/bindings/vapi/std"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/bindings/vapi/std/errors"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/bindings"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/core"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/data"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/lib"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/log"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/protocol/client"
)

type BatchClientImpl struct {
	interfaceName       string
	interfaceDefinition core.InterfaceDefinition
	methodIdentifiers   []core.MethodIdentifier
	methodNameToDefMap  map[string]*core.MethodDefinition
	errorBindingMap     map[string]bindings.BindingType
	interfaceIdentifier core.InterfaceIdentifier
	connector           client.Connector
}

func NewBatchClientImpl(connector client.Connector) *BatchClientImpl {
	interfaceName := "com.vmware.cis.tagging.batch"
	interfaceIdentifier := core.NewInterfaceIdentifier(interfaceName)
	methodIdentifiers := []core.MethodIdentifier{
		core.NewMethodIdentifier(interfaceIdentifier, "getCategories"),
		core.NewMethodIdentifier(interfaceIdentifier, "getAllCategories"),
		core.NewMethodIdentifier(interfaceIdentifier, "getTags"),
		core.NewMethodIdentifier(interfaceIdentifier, "getAllTags"),
		core.NewMethodIdentifier(interfaceIdentifier, "listTagsForCategories"),
		core.NewMethodIdentifier(interfaceIdentifier, "findTagsByName"),
		core.NewMethodIdentifier(interfaceIdentifier, "listAttachedObjects"),
		core.NewMethodIdentifier(interfaceIdentifier, "listAttachedObjectsOnTags"),
		core.NewMethodIdentifier(interfaceIdentifier, "listAllAttachedObjectsOnTags"),
		core.NewMethodIdentifier(interfaceIdentifier, "listAttachedTags"),
		core.NewMethodIdentifier(interfaceIdentifier, "listAttachedTagsOnObjects"),
	}
	interfaceDefinition := core.NewInterfaceDefinition(interfaceIdentifier, methodIdentifiers)
	errorBindingMap := make(map[string]bindings.BindingType)
	errorBindingMap[errors.InternalServerError{}.Error()] = errors.InternalServerErrorBindingType()
	errorBindingMap[errors.InvalidArgument{}.Error()] = errors.InvalidArgumentBindingType()
	errorBindingMap[errors.OperationNotFound{}.Error()] = errors.OperationNotFoundBindingType()
	errorBindingMap[errors.UnexpectedInput{}.Error()] = errors.UnexpectedInputBindingType()
	errorBindingMap[errors.ServiceUnavailable{}.Error()] = errors.ServiceUnavailableBindingType()
	errorBindingMap[errors.TimedOut{}.Error()] = errors.TimedOutBindingType()
	bIface := BatchClientImpl{interfaceName: interfaceName, methodIdentifiers: methodIdentifiers, interfaceDefinition: interfaceDefinition, errorBindingMap: errorBindingMap, interfaceIdentifier: interfaceIdentifier, connector: connector}
	bIface.methodNameToDefMap = make(map[string]*core.MethodDefinition)
	bIface.methodNameToDefMap["get_categories"] = bIface.getCategoriesMethodDefinition()
	bIface.methodNameToDefMap["get_all_categories"] = bIface.getAllCategoriesMethodDefinition()
	bIface.methodNameToDefMap["get_tags"] = bIface.getTagsMethodDefinition()
	bIface.methodNameToDefMap["get_all_tags"] = bIface.getAllTagsMethodDefinition()
	bIface.methodNameToDefMap["list_tags_for_categories"] = bIface.listTagsForCategoriesMethodDefinition()
	bIface.methodNameToDefMap["find_tags_by_name"] = bIface.findTagsByNameMethodDefinition()
	bIface.methodNameToDefMap["list_attached_objects"] = bIface.listAttachedObjectsMethodDefinition()
	bIface.methodNameToDefMap["list_attached_objects_on_tags"] = bIface.listAttachedObjectsOnTagsMethodDefinition()
	bIface.methodNameToDefMap["list_all_attached_objects_on_tags"] = bIface.listAllAttachedObjectsOnTagsMethodDefinition()
	bIface.methodNameToDefMap["list_attached_tags"] = bIface.listAttachedTagsMethodDefinition()
	bIface.methodNameToDefMap["list_attached_tags_on_objects"] = bIface.listAttachedTagsOnObjectsMethodDefinition()
	return &bIface
}

func (bIface *BatchClientImpl) GetCategories(categoryIdsParam []string) ([]CategoryModel, error) {
	typeConverter := bIface.connector.TypeConverter()
	methodIdentifier := core.NewMethodIdentifier(bIface.interfaceIdentifier, "get_categories")
	sv := bindings.NewStructValueBuilder(batchGetCategoriesInputType(), typeConverter)
	sv.AddStructField("CategoryIds", categoryIdsParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput []CategoryModel
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := batchGetCategoriesRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	bIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := bIface.Invoke(bIface.connector.NewExecutionContext(), methodIdentifier, inputDataValue)
	var emptyOutput []CategoryModel
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), batchGetCategoriesOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.([]CategoryModel), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), bIface.errorBindingMap[methodResult.Error().Name()])
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (bIface *BatchClientImpl) GetAllCategories() ([]CategoryModel, error) {
	typeConverter := bIface.connector.TypeConverter()
	methodIdentifier := core.NewMethodIdentifier(bIface.interfaceIdentifier, "get_all_categories")
	sv := bindings.NewStructValueBuilder(batchGetAllCategoriesInputType(), typeConverter)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput []CategoryModel
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := batchGetAllCategoriesRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	bIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := bIface.Invoke(bIface.connector.NewExecutionContext(), methodIdentifier, inputDataValue)
	var emptyOutput []CategoryModel
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), batchGetAllCategoriesOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.([]CategoryModel), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), bIface.errorBindingMap[methodResult.Error().Name()])
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (bIface *BatchClientImpl) GetTags(tagIdsParam []string) ([]TagModel, error) {
	typeConverter := bIface.connector.TypeConverter()
	methodIdentifier := core.NewMethodIdentifier(bIface.interfaceIdentifier, "get_tags")
	sv := bindings.NewStructValueBuilder(batchGetTagsInputType(), typeConverter)
	sv.AddStructField("TagIds", tagIdsParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput []TagModel
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := batchGetTagsRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	bIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := bIface.Invoke(bIface.connector.NewExecutionContext(), methodIdentifier, inputDataValue)
	var emptyOutput []TagModel
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), batchGetTagsOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.([]TagModel), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), bIface.errorBindingMap[methodResult.Error().Name()])
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (bIface *BatchClientImpl) GetAllTags() ([]TagModel, error) {
	typeConverter := bIface.connector.TypeConverter()
	methodIdentifier := core.NewMethodIdentifier(bIface.interfaceIdentifier, "get_all_tags")
	sv := bindings.NewStructValueBuilder(batchGetAllTagsInputType(), typeConverter)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput []TagModel
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := batchGetAllTagsRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	bIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := bIface.Invoke(bIface.connector.NewExecutionContext(), methodIdentifier, inputDataValue)
	var emptyOutput []TagModel
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), batchGetAllTagsOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.([]TagModel), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), bIface.errorBindingMap[methodResult.Error().Name()])
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (bIface *BatchClientImpl) ListTagsForCategories(categoryIdsParam []string) ([]string, error) {
	typeConverter := bIface.connector.TypeConverter()
	methodIdentifier := core.NewMethodIdentifier(bIface.interfaceIdentifier, "list_tags_for_categories")
	sv := bindings.NewStructValueBuilder(batchListTagsForCategoriesInputType(), typeConverter)
	sv.AddStructField("CategoryIds", categoryIdsParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput []string
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := batchListTagsForCategoriesRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	bIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := bIface.Invoke(bIface.connector.NewExecutionContext(), methodIdentifier, inputDataValue)
	var emptyOutput []string
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), batchListTagsForCategoriesOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.([]string), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), bIface.errorBindingMap[methodResult.Error().Name()])
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (bIface *BatchClientImpl) FindTagsByName(tagNameParam string) ([]string, error) {
	typeConverter := bIface.connector.TypeConverter()
	methodIdentifier := core.NewMethodIdentifier(bIface.interfaceIdentifier, "find_tags_by_name")
	sv := bindings.NewStructValueBuilder(batchFindTagsByNameInputType(), typeConverter)
	sv.AddStructField("TagName", tagNameParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput []string
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := batchFindTagsByNameRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	bIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := bIface.Invoke(bIface.connector.NewExecutionContext(), methodIdentifier, inputDataValue)
	var emptyOutput []string
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), batchFindTagsByNameOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.([]string), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), bIface.errorBindingMap[methodResult.Error().Name()])
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (bIface *BatchClientImpl) ListAttachedObjects(tagIdsParam []string) ([]std.DynamicID, error) {
	typeConverter := bIface.connector.TypeConverter()
	methodIdentifier := core.NewMethodIdentifier(bIface.interfaceIdentifier, "list_attached_objects")
	sv := bindings.NewStructValueBuilder(batchListAttachedObjectsInputType(), typeConverter)
	sv.AddStructField("TagIds", tagIdsParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput []std.DynamicID
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := batchListAttachedObjectsRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	bIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := bIface.Invoke(bIface.connector.NewExecutionContext(), methodIdentifier, inputDataValue)
	var emptyOutput []std.DynamicID
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), batchListAttachedObjectsOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.([]std.DynamicID), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), bIface.errorBindingMap[methodResult.Error().Name()])
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (bIface *BatchClientImpl) ListAttachedObjectsOnTags(tagIdsParam []string) ([]BatchTagToObjects, error) {
	typeConverter := bIface.connector.TypeConverter()
	methodIdentifier := core.NewMethodIdentifier(bIface.interfaceIdentifier, "list_attached_objects_on_tags")
	sv := bindings.NewStructValueBuilder(batchListAttachedObjectsOnTagsInputType(), typeConverter)
	sv.AddStructField("TagIds", tagIdsParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput []BatchTagToObjects
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := batchListAttachedObjectsOnTagsRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	bIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := bIface.Invoke(bIface.connector.NewExecutionContext(), methodIdentifier, inputDataValue)
	var emptyOutput []BatchTagToObjects
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), batchListAttachedObjectsOnTagsOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.([]BatchTagToObjects), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), bIface.errorBindingMap[methodResult.Error().Name()])
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (bIface *BatchClientImpl) ListAllAttachedObjectsOnTags() ([]BatchTagToObjects, error) {
	typeConverter := bIface.connector.TypeConverter()
	methodIdentifier := core.NewMethodIdentifier(bIface.interfaceIdentifier, "list_all_attached_objects_on_tags")
	sv := bindings.NewStructValueBuilder(batchListAllAttachedObjectsOnTagsInputType(), typeConverter)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput []BatchTagToObjects
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := batchListAllAttachedObjectsOnTagsRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	bIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := bIface.Invoke(bIface.connector.NewExecutionContext(), methodIdentifier, inputDataValue)
	var emptyOutput []BatchTagToObjects
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), batchListAllAttachedObjectsOnTagsOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.([]BatchTagToObjects), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), bIface.errorBindingMap[methodResult.Error().Name()])
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (bIface *BatchClientImpl) ListAttachedTags(objectIdsParam []std.DynamicID) ([]string, error) {
	typeConverter := bIface.connector.TypeConverter()
	methodIdentifier := core.NewMethodIdentifier(bIface.interfaceIdentifier, "list_attached_tags")
	sv := bindings.NewStructValueBuilder(batchListAttachedTagsInputType(), typeConverter)
	sv.AddStructField("ObjectIds", objectIdsParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput []string
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := batchListAttachedTagsRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	bIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := bIface.Invoke(bIface.connector.NewExecutionContext(), methodIdentifier, inputDataValue)
	var emptyOutput []string
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), batchListAttachedTagsOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.([]string), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), bIface.errorBindingMap[methodResult.Error().Name()])
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (bIface *BatchClientImpl) ListAttachedTagsOnObjects(objectIdsParam []std.DynamicID) ([]BatchObjectToTags, error) {
	typeConverter := bIface.connector.TypeConverter()
	methodIdentifier := core.NewMethodIdentifier(bIface.interfaceIdentifier, "list_attached_tags_on_objects")
	sv := bindings.NewStructValueBuilder(batchListAttachedTagsOnObjectsInputType(), typeConverter)
	sv.AddStructField("ObjectIds", objectIdsParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput []BatchObjectToTags
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := batchListAttachedTagsOnObjectsRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	bIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := bIface.Invoke(bIface.connector.NewExecutionContext(), methodIdentifier, inputDataValue)
	var emptyOutput []BatchObjectToTags
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), batchListAttachedTagsOnObjectsOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.([]BatchObjectToTags), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), bIface.errorBindingMap[methodResult.Error().Name()])
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}


func (bIface *BatchClientImpl) Invoke(ctx *core.ExecutionContext, methodId core.MethodIdentifier, inputDataValue data.DataValue) core.MethodResult {
	methodResult := bIface.connector.GetApiProvider().Invoke(bIface.interfaceName, methodId.Name(), inputDataValue, ctx)
	return methodResult
}


func (bIface *BatchClientImpl) getCategoriesMethodDefinition() *core.MethodDefinition {
	interfaceIdentifier := core.NewInterfaceIdentifier(bIface.interfaceName)
	typeConverter := bIface.connector.TypeConverter()

	input, inputError := typeConverter.ConvertToDataDefinition(batchGetCategoriesInputType())
	output, outputError := typeConverter.ConvertToDataDefinition(batchGetCategoriesOutputType())
	if inputError != nil {
		log.Errorf("Error in ConvertToDataDefinition for BatchClientImpl.getCategories method's input - %s",
			bindings.VAPIerrorsToError(inputError).Error())
		return nil
	}
	if outputError != nil {
		log.Errorf("Error in ConvertToDataDefinition for BatchClientImpl.getCategories method's output - %s",
			bindings.VAPIerrorsToError(outputError).Error())
		return nil
	}
	methodIdentifier := core.NewMethodIdentifier(interfaceIdentifier, "getCategories")
	errorDefinitions := make([]data.ErrorDefinition, 0)

	methodDefinition := core.NewMethodDefinition(methodIdentifier, input, output, errorDefinitions)
	return &methodDefinition
}

func (bIface *BatchClientImpl) getAllCategoriesMethodDefinition() *core.MethodDefinition {
	interfaceIdentifier := core.NewInterfaceIdentifier(bIface.interfaceName)
	typeConverter := bIface.connector.TypeConverter()

	input, inputError := typeConverter.ConvertToDataDefinition(batchGetAllCategoriesInputType())
	output, outputError := typeConverter.ConvertToDataDefinition(batchGetAllCategoriesOutputType())
	if inputError != nil {
		log.Errorf("Error in ConvertToDataDefinition for BatchClientImpl.getAllCategories method's input - %s",
			bindings.VAPIerrorsToError(inputError).Error())
		return nil
	}
	if outputError != nil {
		log.Errorf("Error in ConvertToDataDefinition for BatchClientImpl.getAllCategories method's output - %s",
			bindings.VAPIerrorsToError(outputError).Error())
		return nil
	}
	methodIdentifier := core.NewMethodIdentifier(interfaceIdentifier, "getAllCategories")
	errorDefinitions := make([]data.ErrorDefinition, 0)

	methodDefinition := core.NewMethodDefinition(methodIdentifier, input, output, errorDefinitions)
	return &methodDefinition
}

func (bIface *BatchClientImpl) getTagsMethodDefinition() *core.MethodDefinition {
	interfaceIdentifier := core.NewInterfaceIdentifier(bIface.interfaceName)
	typeConverter := bIface.connector.TypeConverter()

	input, inputError := typeConverter.ConvertToDataDefinition(batchGetTagsInputType())
	output, outputError := typeConverter.ConvertToDataDefinition(batchGetTagsOutputType())
	if inputError != nil {
		log.Errorf("Error in ConvertToDataDefinition for BatchClientImpl.getTags method's input - %s",
			bindings.VAPIerrorsToError(inputError).Error())
		return nil
	}
	if outputError != nil {
		log.Errorf("Error in ConvertToDataDefinition for BatchClientImpl.getTags method's output - %s",
			bindings.VAPIerrorsToError(outputError).Error())
		return nil
	}
	methodIdentifier := core.NewMethodIdentifier(interfaceIdentifier, "getTags")
	errorDefinitions := make([]data.ErrorDefinition, 0)

	methodDefinition := core.NewMethodDefinition(methodIdentifier, input, output, errorDefinitions)
	return &methodDefinition
}

func (bIface *BatchClientImpl) getAllTagsMethodDefinition() *core.MethodDefinition {
	interfaceIdentifier := core.NewInterfaceIdentifier(bIface.interfaceName)
	typeConverter := bIface.connector.TypeConverter()

	input, inputError := typeConverter.ConvertToDataDefinition(batchGetAllTagsInputType())
	output, outputError := typeConverter.ConvertToDataDefinition(batchGetAllTagsOutputType())
	if inputError != nil {
		log.Errorf("Error in ConvertToDataDefinition for BatchClientImpl.getAllTags method's input - %s",
			bindings.VAPIerrorsToError(inputError).Error())
		return nil
	}
	if outputError != nil {
		log.Errorf("Error in ConvertToDataDefinition for BatchClientImpl.getAllTags method's output - %s",
			bindings.VAPIerrorsToError(outputError).Error())
		return nil
	}
	methodIdentifier := core.NewMethodIdentifier(interfaceIdentifier, "getAllTags")
	errorDefinitions := make([]data.ErrorDefinition, 0)

	methodDefinition := core.NewMethodDefinition(methodIdentifier, input, output, errorDefinitions)
	return &methodDefinition
}

func (bIface *BatchClientImpl) listTagsForCategoriesMethodDefinition() *core.MethodDefinition {
	interfaceIdentifier := core.NewInterfaceIdentifier(bIface.interfaceName)
	typeConverter := bIface.connector.TypeConverter()

	input, inputError := typeConverter.ConvertToDataDefinition(batchListTagsForCategoriesInputType())
	output, outputError := typeConverter.ConvertToDataDefinition(batchListTagsForCategoriesOutputType())
	if inputError != nil {
		log.Errorf("Error in ConvertToDataDefinition for BatchClientImpl.listTagsForCategories method's input - %s",
			bindings.VAPIerrorsToError(inputError).Error())
		return nil
	}
	if outputError != nil {
		log.Errorf("Error in ConvertToDataDefinition for BatchClientImpl.listTagsForCategories method's output - %s",
			bindings.VAPIerrorsToError(outputError).Error())
		return nil
	}
	methodIdentifier := core.NewMethodIdentifier(interfaceIdentifier, "listTagsForCategories")
	errorDefinitions := make([]data.ErrorDefinition, 0)

	methodDefinition := core.NewMethodDefinition(methodIdentifier, input, output, errorDefinitions)
	return &methodDefinition
}

func (bIface *BatchClientImpl) findTagsByNameMethodDefinition() *core.MethodDefinition {
	interfaceIdentifier := core.NewInterfaceIdentifier(bIface.interfaceName)
	typeConverter := bIface.connector.TypeConverter()

	input, inputError := typeConverter.ConvertToDataDefinition(batchFindTagsByNameInputType())
	output, outputError := typeConverter.ConvertToDataDefinition(batchFindTagsByNameOutputType())
	if inputError != nil {
		log.Errorf("Error in ConvertToDataDefinition for BatchClientImpl.findTagsByName method's input - %s",
			bindings.VAPIerrorsToError(inputError).Error())
		return nil
	}
	if outputError != nil {
		log.Errorf("Error in ConvertToDataDefinition for BatchClientImpl.findTagsByName method's output - %s",
			bindings.VAPIerrorsToError(outputError).Error())
		return nil
	}
	methodIdentifier := core.NewMethodIdentifier(interfaceIdentifier, "findTagsByName")
	errorDefinitions := make([]data.ErrorDefinition, 0)

	methodDefinition := core.NewMethodDefinition(methodIdentifier, input, output, errorDefinitions)
	return &methodDefinition
}

func (bIface *BatchClientImpl) listAttachedObjectsMethodDefinition() *core.MethodDefinition {
	interfaceIdentifier := core.NewInterfaceIdentifier(bIface.interfaceName)
	typeConverter := bIface.connector.TypeConverter()

	input, inputError := typeConverter.ConvertToDataDefinition(batchListAttachedObjectsInputType())
	output, outputError := typeConverter.ConvertToDataDefinition(batchListAttachedObjectsOutputType())
	if inputError != nil {
		log.Errorf("Error in ConvertToDataDefinition for BatchClientImpl.listAttachedObjects method's input - %s",
			bindings.VAPIerrorsToError(inputError).Error())
		return nil
	}
	if outputError != nil {
		log.Errorf("Error in ConvertToDataDefinition for BatchClientImpl.listAttachedObjects method's output - %s",
			bindings.VAPIerrorsToError(outputError).Error())
		return nil
	}
	methodIdentifier := core.NewMethodIdentifier(interfaceIdentifier, "listAttachedObjects")
	errorDefinitions := make([]data.ErrorDefinition, 0)
	bIface.errorBindingMap[errors.Unauthenticated{}.Error()] = errors.UnauthenticatedBindingType()
	errDef1, errError1 := typeConverter.ConvertToDataDefinition(errors.UnauthenticatedBindingType())
	if errError1 != nil {
		log.Errorf("Error in ConvertToDataDefinition for BatchClientImpl.listAttachedObjects method's errors.Unauthenticated error - %s",
			bindings.VAPIerrorsToError(errError1).Error())
		return nil
	}
	errorDefinitions = append(errorDefinitions, errDef1.(data.ErrorDefinition))

	methodDefinition := core.NewMethodDefinition(methodIdentifier, input, output, errorDefinitions)
	return &methodDefinition
}

func (bIface *BatchClientImpl) listAttachedObjectsOnTagsMethodDefinition() *core.MethodDefinition {
	interfaceIdentifier := core.NewInterfaceIdentifier(bIface.interfaceName)
	typeConverter := bIface.connector.TypeConverter()

	input, inputError := typeConverter.ConvertToDataDefinition(batchListAttachedObjectsOnTagsInputType())
	output, outputError := typeConverter.ConvertToDataDefinition(batchListAttachedObjectsOnTagsOutputType())
	if inputError != nil {
		log.Errorf("Error in ConvertToDataDefinition for BatchClientImpl.listAttachedObjectsOnTags method's input - %s",
			bindings.VAPIerrorsToError(inputError).Error())
		return nil
	}
	if outputError != nil {
		log.Errorf("Error in ConvertToDataDefinition for BatchClientImpl.listAttachedObjectsOnTags method's output - %s",
			bindings.VAPIerrorsToError(outputError).Error())
		return nil
	}
	methodIdentifier := core.NewMethodIdentifier(interfaceIdentifier, "listAttachedObjectsOnTags")
	errorDefinitions := make([]data.ErrorDefinition, 0)
	bIface.errorBindingMap[errors.Unauthenticated{}.Error()] = errors.UnauthenticatedBindingType()
	errDef1, errError1 := typeConverter.ConvertToDataDefinition(errors.UnauthenticatedBindingType())
	if errError1 != nil {
		log.Errorf("Error in ConvertToDataDefinition for BatchClientImpl.listAttachedObjectsOnTags method's errors.Unauthenticated error - %s",
			bindings.VAPIerrorsToError(errError1).Error())
		return nil
	}
	errorDefinitions = append(errorDefinitions, errDef1.(data.ErrorDefinition))

	methodDefinition := core.NewMethodDefinition(methodIdentifier, input, output, errorDefinitions)
	return &methodDefinition
}

func (bIface *BatchClientImpl) listAllAttachedObjectsOnTagsMethodDefinition() *core.MethodDefinition {
	interfaceIdentifier := core.NewInterfaceIdentifier(bIface.interfaceName)
	typeConverter := bIface.connector.TypeConverter()

	input, inputError := typeConverter.ConvertToDataDefinition(batchListAllAttachedObjectsOnTagsInputType())
	output, outputError := typeConverter.ConvertToDataDefinition(batchListAllAttachedObjectsOnTagsOutputType())
	if inputError != nil {
		log.Errorf("Error in ConvertToDataDefinition for BatchClientImpl.listAllAttachedObjectsOnTags method's input - %s",
			bindings.VAPIerrorsToError(inputError).Error())
		return nil
	}
	if outputError != nil {
		log.Errorf("Error in ConvertToDataDefinition for BatchClientImpl.listAllAttachedObjectsOnTags method's output - %s",
			bindings.VAPIerrorsToError(outputError).Error())
		return nil
	}
	methodIdentifier := core.NewMethodIdentifier(interfaceIdentifier, "listAllAttachedObjectsOnTags")
	errorDefinitions := make([]data.ErrorDefinition, 0)
	bIface.errorBindingMap[errors.Unauthenticated{}.Error()] = errors.UnauthenticatedBindingType()
	errDef1, errError1 := typeConverter.ConvertToDataDefinition(errors.UnauthenticatedBindingType())
	if errError1 != nil {
		log.Errorf("Error in ConvertToDataDefinition for BatchClientImpl.listAllAttachedObjectsOnTags method's errors.Unauthenticated error - %s",
			bindings.VAPIerrorsToError(errError1).Error())
		return nil
	}
	errorDefinitions = append(errorDefinitions, errDef1.(data.ErrorDefinition))

	methodDefinition := core.NewMethodDefinition(methodIdentifier, input, output, errorDefinitions)
	return &methodDefinition
}

func (bIface *BatchClientImpl) listAttachedTagsMethodDefinition() *core.MethodDefinition {
	interfaceIdentifier := core.NewInterfaceIdentifier(bIface.interfaceName)
	typeConverter := bIface.connector.TypeConverter()

	input, inputError := typeConverter.ConvertToDataDefinition(batchListAttachedTagsInputType())
	output, outputError := typeConverter.ConvertToDataDefinition(batchListAttachedTagsOutputType())
	if inputError != nil {
		log.Errorf("Error in ConvertToDataDefinition for BatchClientImpl.listAttachedTags method's input - %s",
			bindings.VAPIerrorsToError(inputError).Error())
		return nil
	}
	if outputError != nil {
		log.Errorf("Error in ConvertToDataDefinition for BatchClientImpl.listAttachedTags method's output - %s",
			bindings.VAPIerrorsToError(outputError).Error())
		return nil
	}
	methodIdentifier := core.NewMethodIdentifier(interfaceIdentifier, "listAttachedTags")
	errorDefinitions := make([]data.ErrorDefinition, 0)
	bIface.errorBindingMap[errors.Unauthenticated{}.Error()] = errors.UnauthenticatedBindingType()
	errDef1, errError1 := typeConverter.ConvertToDataDefinition(errors.UnauthenticatedBindingType())
	if errError1 != nil {
		log.Errorf("Error in ConvertToDataDefinition for BatchClientImpl.listAttachedTags method's errors.Unauthenticated error - %s",
			bindings.VAPIerrorsToError(errError1).Error())
		return nil
	}
	errorDefinitions = append(errorDefinitions, errDef1.(data.ErrorDefinition))

	methodDefinition := core.NewMethodDefinition(methodIdentifier, input, output, errorDefinitions)
	return &methodDefinition
}

func (bIface *BatchClientImpl) listAttachedTagsOnObjectsMethodDefinition() *core.MethodDefinition {
	interfaceIdentifier := core.NewInterfaceIdentifier(bIface.interfaceName)
	typeConverter := bIface.connector.TypeConverter()

	input, inputError := typeConverter.ConvertToDataDefinition(batchListAttachedTagsOnObjectsInputType())
	output, outputError := typeConverter.ConvertToDataDefinition(batchListAttachedTagsOnObjectsOutputType())
	if inputError != nil {
		log.Errorf("Error in ConvertToDataDefinition for BatchClientImpl.listAttachedTagsOnObjects method's input - %s",
			bindings.VAPIerrorsToError(inputError).Error())
		return nil
	}
	if outputError != nil {
		log.Errorf("Error in ConvertToDataDefinition for BatchClientImpl.listAttachedTagsOnObjects method's output - %s",
			bindings.VAPIerrorsToError(outputError).Error())
		return nil
	}
	methodIdentifier := core.NewMethodIdentifier(interfaceIdentifier, "listAttachedTagsOnObjects")
	errorDefinitions := make([]data.ErrorDefinition, 0)
	bIface.errorBindingMap[errors.Unauthenticated{}.Error()] = errors.UnauthenticatedBindingType()
	errDef1, errError1 := typeConverter.ConvertToDataDefinition(errors.UnauthenticatedBindingType())
	if errError1 != nil {
		log.Errorf("Error in ConvertToDataDefinition for BatchClientImpl.listAttachedTagsOnObjects method's errors.Unauthenticated error - %s",
			bindings.VAPIerrorsToError(errError1).Error())
		return nil
	}
	errorDefinitions = append(errorDefinitions, errDef1.(data.ErrorDefinition))

	methodDefinition := core.NewMethodDefinition(methodIdentifier, input, output, errorDefinitions)
	return &methodDefinition
}
