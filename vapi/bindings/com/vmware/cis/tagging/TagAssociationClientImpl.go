
/* Copyright © 2019 VMware, Inc. All Rights Reserved.
     SPDX-License-Identifier: BSD-2-Clause */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Client stubs for service: TagAssociation
 * Functions that implement the generated TagAssociationClient interface
 */


package tagging
import (
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/bindings/com/vmware/vapi/std"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/bindings/com/vmware/vapi/std/errors"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/bindings"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/core"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/data"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/lib"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/log"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/protocol/client"
)


type TagAssociationClientImpl struct{
      interfaceName string
      interfaceDefinition core.InterfaceDefinition
      methodIdentifiers[]core.MethodIdentifier
      methodNameToDefMap map[string]*core.MethodDefinition
      errorBindingMap map[string]bindings.BindingType
      interfaceIdentifier core.InterfaceIdentifier
      connector client.Connector
}


func NewTagAssociationClientImpl(connector client.Connector) *TagAssociationClientImpl {
      interfaceName := "com.vmware.cis.tagging.tag_association"
      interfaceIdentifier := core.NewInterfaceIdentifier(interfaceName)
      methodIdentifiers := []core.MethodIdentifier{
          core.NewMethodIdentifier(interfaceIdentifier, "attach"),
          core.NewMethodIdentifier(interfaceIdentifier, "attachMultipleTagsToObject"),
          core.NewMethodIdentifier(interfaceIdentifier, "attachTagToMultipleObjects"),
          core.NewMethodIdentifier(interfaceIdentifier, "detach"),
          core.NewMethodIdentifier(interfaceIdentifier, "detachMultipleTagsFromObject"),
          core.NewMethodIdentifier(interfaceIdentifier, "detachTagFromMultipleObjects"),
          core.NewMethodIdentifier(interfaceIdentifier, "listAttachedObjects"),
          core.NewMethodIdentifier(interfaceIdentifier, "listAttachedObjectsOnTags"),
          core.NewMethodIdentifier(interfaceIdentifier, "listAttachedTags"),
          core.NewMethodIdentifier(interfaceIdentifier, "listAttachedTagsOnObjects"),
          core.NewMethodIdentifier(interfaceIdentifier, "listAttachableTags"),
      }
      interfaceDefinition := core.NewInterfaceDefinition(interfaceIdentifier, methodIdentifiers)
      errorBindingMap := make(map[string]bindings.BindingType)
      errorBindingMap[errors.InternalServerError{}.Error()] = errors.InternalServerErrorBindingType()
	  errorBindingMap[errors.InvalidArgument{}.Error()] = errors.InvalidArgumentBindingType()
	  errorBindingMap[errors.OperationNotFound{}.Error()] = errors.OperationNotFoundBindingType()
	  errorBindingMap[errors.UnexpectedInput{}.Error()] = errors.UnexpectedInputBindingType()
	  errorBindingMap[errors.ServiceUnavailable{}.Error()] = errors.ServiceUnavailableBindingType()
	  errorBindingMap[errors.TimedOut{}.Error()] = errors.TimedOutBindingType()
      tIface := TagAssociationClientImpl{interfaceName: interfaceName, methodIdentifiers: methodIdentifiers, interfaceDefinition: interfaceDefinition, errorBindingMap: errorBindingMap, interfaceIdentifier:interfaceIdentifier, connector: connector}
      tIface.methodNameToDefMap = make(map[string]*core.MethodDefinition)
      tIface.methodNameToDefMap["attach"] = tIface.attachMethodDefinition()
      tIface.methodNameToDefMap["attach_multiple_tags_to_object"] = tIface.attachMultipleTagsToObjectMethodDefinition()
      tIface.methodNameToDefMap["attach_tag_to_multiple_objects"] = tIface.attachTagToMultipleObjectsMethodDefinition()
      tIface.methodNameToDefMap["detach"] = tIface.detachMethodDefinition()
      tIface.methodNameToDefMap["detach_multiple_tags_from_object"] = tIface.detachMultipleTagsFromObjectMethodDefinition()
      tIface.methodNameToDefMap["detach_tag_from_multiple_objects"] = tIface.detachTagFromMultipleObjectsMethodDefinition()
      tIface.methodNameToDefMap["list_attached_objects"] = tIface.listAttachedObjectsMethodDefinition()
      tIface.methodNameToDefMap["list_attached_objects_on_tags"] = tIface.listAttachedObjectsOnTagsMethodDefinition()
      tIface.methodNameToDefMap["list_attached_tags"] = tIface.listAttachedTagsMethodDefinition()
      tIface.methodNameToDefMap["list_attached_tags_on_objects"] = tIface.listAttachedTagsOnObjectsMethodDefinition()
      tIface.methodNameToDefMap["list_attachable_tags"] = tIface.listAttachableTagsMethodDefinition()
      return &tIface
}

func (tIface *TagAssociationClientImpl) Attach(tagIdParam string, objectIdParam std.DynamicID) error {
	typeConverter := tIface.connector.TypeConverter()
	methodIdentifier := core.NewMethodIdentifier(tIface.interfaceIdentifier, "attach")
	sv := bindings.NewStructValueBuilder(tagAssociationAttachInputType(), typeConverter)
	sv.AddStructField("TagId", tagIdParam)
	sv.AddStructField("ObjectId", objectIdParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		return bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := tagAssociationAttachRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	tIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult:= tIface.Invoke(tIface.connector.NewExecutionContext(), methodIdentifier, inputDataValue)
    if methodResult.IsSuccess() {
		return nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), tIface.errorBindingMap[methodResult.Error().Name()])
		if errorInError != nil {
		    return bindings.VAPIerrorsToError(errorInError)
		}
		return methodError.(error)
	}
}
func (tIface *TagAssociationClientImpl) AttachMultipleTagsToObject(objectIdParam std.DynamicID, tagIdsParam []string) (TagAssociationBatchResult, error) {
	typeConverter := tIface.connector.TypeConverter()
	methodIdentifier := core.NewMethodIdentifier(tIface.interfaceIdentifier, "attach_multiple_tags_to_object")
	sv := bindings.NewStructValueBuilder(tagAssociationAttachMultipleTagsToObjectInputType(), typeConverter)
	sv.AddStructField("ObjectId", objectIdParam)
	sv.AddStructField("TagIds", tagIdsParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
        var emptyOutput TagAssociationBatchResult
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := tagAssociationAttachMultipleTagsToObjectRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	tIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult:= tIface.Invoke(tIface.connector.NewExecutionContext(), methodIdentifier, inputDataValue)
	var emptyOutput TagAssociationBatchResult
    if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), tagAssociationAttachMultipleTagsToObjectOutputType())
		if errorInOutput != nil {
		    return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
	    return output.(TagAssociationBatchResult), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), tIface.errorBindingMap[methodResult.Error().Name()])
		if errorInError != nil {
		    return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}
func (tIface *TagAssociationClientImpl) AttachTagToMultipleObjects(tagIdParam string, objectIdsParam []std.DynamicID) (TagAssociationBatchResult, error) {
	typeConverter := tIface.connector.TypeConverter()
	methodIdentifier := core.NewMethodIdentifier(tIface.interfaceIdentifier, "attach_tag_to_multiple_objects")
	sv := bindings.NewStructValueBuilder(tagAssociationAttachTagToMultipleObjectsInputType(), typeConverter)
	sv.AddStructField("TagId", tagIdParam)
	sv.AddStructField("ObjectIds", objectIdsParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
        var emptyOutput TagAssociationBatchResult
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := tagAssociationAttachTagToMultipleObjectsRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	tIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult:= tIface.Invoke(tIface.connector.NewExecutionContext(), methodIdentifier, inputDataValue)
	var emptyOutput TagAssociationBatchResult
    if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), tagAssociationAttachTagToMultipleObjectsOutputType())
		if errorInOutput != nil {
		    return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
	    return output.(TagAssociationBatchResult), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), tIface.errorBindingMap[methodResult.Error().Name()])
		if errorInError != nil {
		    return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}
func (tIface *TagAssociationClientImpl) Detach(tagIdParam string, objectIdParam std.DynamicID) error {
	typeConverter := tIface.connector.TypeConverter()
	methodIdentifier := core.NewMethodIdentifier(tIface.interfaceIdentifier, "detach")
	sv := bindings.NewStructValueBuilder(tagAssociationDetachInputType(), typeConverter)
	sv.AddStructField("TagId", tagIdParam)
	sv.AddStructField("ObjectId", objectIdParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		return bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := tagAssociationDetachRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	tIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult:= tIface.Invoke(tIface.connector.NewExecutionContext(), methodIdentifier, inputDataValue)
    if methodResult.IsSuccess() {
		return nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), tIface.errorBindingMap[methodResult.Error().Name()])
		if errorInError != nil {
		    return bindings.VAPIerrorsToError(errorInError)
		}
		return methodError.(error)
	}
}
func (tIface *TagAssociationClientImpl) DetachMultipleTagsFromObject(objectIdParam std.DynamicID, tagIdsParam []string) (TagAssociationBatchResult, error) {
	typeConverter := tIface.connector.TypeConverter()
	methodIdentifier := core.NewMethodIdentifier(tIface.interfaceIdentifier, "detach_multiple_tags_from_object")
	sv := bindings.NewStructValueBuilder(tagAssociationDetachMultipleTagsFromObjectInputType(), typeConverter)
	sv.AddStructField("ObjectId", objectIdParam)
	sv.AddStructField("TagIds", tagIdsParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
        var emptyOutput TagAssociationBatchResult
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := tagAssociationDetachMultipleTagsFromObjectRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	tIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult:= tIface.Invoke(tIface.connector.NewExecutionContext(), methodIdentifier, inputDataValue)
	var emptyOutput TagAssociationBatchResult
    if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), tagAssociationDetachMultipleTagsFromObjectOutputType())
		if errorInOutput != nil {
		    return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
	    return output.(TagAssociationBatchResult), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), tIface.errorBindingMap[methodResult.Error().Name()])
		if errorInError != nil {
		    return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}
func (tIface *TagAssociationClientImpl) DetachTagFromMultipleObjects(tagIdParam string, objectIdsParam []std.DynamicID) (TagAssociationBatchResult, error) {
	typeConverter := tIface.connector.TypeConverter()
	methodIdentifier := core.NewMethodIdentifier(tIface.interfaceIdentifier, "detach_tag_from_multiple_objects")
	sv := bindings.NewStructValueBuilder(tagAssociationDetachTagFromMultipleObjectsInputType(), typeConverter)
	sv.AddStructField("TagId", tagIdParam)
	sv.AddStructField("ObjectIds", objectIdsParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
        var emptyOutput TagAssociationBatchResult
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := tagAssociationDetachTagFromMultipleObjectsRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	tIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult:= tIface.Invoke(tIface.connector.NewExecutionContext(), methodIdentifier, inputDataValue)
	var emptyOutput TagAssociationBatchResult
    if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), tagAssociationDetachTagFromMultipleObjectsOutputType())
		if errorInOutput != nil {
		    return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
	    return output.(TagAssociationBatchResult), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), tIface.errorBindingMap[methodResult.Error().Name()])
		if errorInError != nil {
		    return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}
func (tIface *TagAssociationClientImpl) ListAttachedObjects(tagIdParam string) ([]std.DynamicID, error) {
	typeConverter := tIface.connector.TypeConverter()
	methodIdentifier := core.NewMethodIdentifier(tIface.interfaceIdentifier, "list_attached_objects")
	sv := bindings.NewStructValueBuilder(tagAssociationListAttachedObjectsInputType(), typeConverter)
	sv.AddStructField("TagId", tagIdParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
        var emptyOutput []std.DynamicID
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := tagAssociationListAttachedObjectsRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	tIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult:= tIface.Invoke(tIface.connector.NewExecutionContext(), methodIdentifier, inputDataValue)
	var emptyOutput []std.DynamicID
    if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), tagAssociationListAttachedObjectsOutputType())
		if errorInOutput != nil {
		    return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
	    return output.([]std.DynamicID), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), tIface.errorBindingMap[methodResult.Error().Name()])
		if errorInError != nil {
		    return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}
func (tIface *TagAssociationClientImpl) ListAttachedObjectsOnTags(tagIdsParam []string) ([]TagAssociationTagToObjects, error) {
	typeConverter := tIface.connector.TypeConverter()
	methodIdentifier := core.NewMethodIdentifier(tIface.interfaceIdentifier, "list_attached_objects_on_tags")
	sv := bindings.NewStructValueBuilder(tagAssociationListAttachedObjectsOnTagsInputType(), typeConverter)
	sv.AddStructField("TagIds", tagIdsParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
        var emptyOutput []TagAssociationTagToObjects
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := tagAssociationListAttachedObjectsOnTagsRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	tIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult:= tIface.Invoke(tIface.connector.NewExecutionContext(), methodIdentifier, inputDataValue)
	var emptyOutput []TagAssociationTagToObjects
    if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), tagAssociationListAttachedObjectsOnTagsOutputType())
		if errorInOutput != nil {
		    return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
	    return output.([]TagAssociationTagToObjects), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), tIface.errorBindingMap[methodResult.Error().Name()])
		if errorInError != nil {
		    return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}
func (tIface *TagAssociationClientImpl) ListAttachedTags(objectIdParam std.DynamicID) ([]string, error) {
	typeConverter := tIface.connector.TypeConverter()
	methodIdentifier := core.NewMethodIdentifier(tIface.interfaceIdentifier, "list_attached_tags")
	sv := bindings.NewStructValueBuilder(tagAssociationListAttachedTagsInputType(), typeConverter)
	sv.AddStructField("ObjectId", objectIdParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
        var emptyOutput []string
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := tagAssociationListAttachedTagsRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	tIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult:= tIface.Invoke(tIface.connector.NewExecutionContext(), methodIdentifier, inputDataValue)
	var emptyOutput []string
    if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), tagAssociationListAttachedTagsOutputType())
		if errorInOutput != nil {
		    return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
	    return output.([]string), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), tIface.errorBindingMap[methodResult.Error().Name()])
		if errorInError != nil {
		    return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}
func (tIface *TagAssociationClientImpl) ListAttachedTagsOnObjects(objectIdsParam []std.DynamicID) ([]TagAssociationObjectToTags, error) {
	typeConverter := tIface.connector.TypeConverter()
	methodIdentifier := core.NewMethodIdentifier(tIface.interfaceIdentifier, "list_attached_tags_on_objects")
	sv := bindings.NewStructValueBuilder(tagAssociationListAttachedTagsOnObjectsInputType(), typeConverter)
	sv.AddStructField("ObjectIds", objectIdsParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
        var emptyOutput []TagAssociationObjectToTags
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := tagAssociationListAttachedTagsOnObjectsRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	tIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult:= tIface.Invoke(tIface.connector.NewExecutionContext(), methodIdentifier, inputDataValue)
	var emptyOutput []TagAssociationObjectToTags
    if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), tagAssociationListAttachedTagsOnObjectsOutputType())
		if errorInOutput != nil {
		    return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
	    return output.([]TagAssociationObjectToTags), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), tIface.errorBindingMap[methodResult.Error().Name()])
		if errorInError != nil {
		    return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}
func (tIface *TagAssociationClientImpl) ListAttachableTags(objectIdParam std.DynamicID) ([]string, error) {
	typeConverter := tIface.connector.TypeConverter()
	methodIdentifier := core.NewMethodIdentifier(tIface.interfaceIdentifier, "list_attachable_tags")
	sv := bindings.NewStructValueBuilder(tagAssociationListAttachableTagsInputType(), typeConverter)
	sv.AddStructField("ObjectId", objectIdParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
        var emptyOutput []string
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := tagAssociationListAttachableTagsRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	tIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult:= tIface.Invoke(tIface.connector.NewExecutionContext(), methodIdentifier, inputDataValue)
	var emptyOutput []string
    if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), tagAssociationListAttachableTagsOutputType())
		if errorInOutput != nil {
		    return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
	    return output.([]string), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), tIface.errorBindingMap[methodResult.Error().Name()])
		if errorInError != nil {
		    return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (tIface *TagAssociationClientImpl) Invoke(ctx *core.ExecutionContext, methodId core.MethodIdentifier, inputDataValue data.DataValue) core.MethodResult {
    methodResult := tIface.connector.GetApiProvider().Invoke(tIface.interfaceName, methodId.Name(), inputDataValue, ctx)
    return methodResult
}

func (tIface *TagAssociationClientImpl) attachMethodDefinition() *core.MethodDefinition {
      interfaceIdentifier := core.NewInterfaceIdentifier(tIface.interfaceName)
      typeConverter := tIface.connector.TypeConverter()

      input, inputError := typeConverter.ConvertToDataDefinition(tagAssociationAttachInputType())
      output, outputError := typeConverter.ConvertToDataDefinition(tagAssociationAttachOutputType())
      if(inputError != nil) {
          log.Errorf("Error in ConvertToDataDefinition for TagAssociationClientImpl.attach method's input - %s",
              bindings.VAPIerrorsToError(inputError).Error())
          return nil
      }
      if(outputError != nil) {
          log.Errorf("Error in ConvertToDataDefinition for TagAssociationClientImpl.attach method's output - %s",
              bindings.VAPIerrorsToError(outputError).Error())
          return nil
      }
      methodIdentifier := core.NewMethodIdentifier(interfaceIdentifier, "attach")
      errorDefinitions := make([]data.ErrorDefinition,0)
      tIface.errorBindingMap[errors.NotFound{}.Error()] = errors.NotFoundBindingType()
	  errDef1, errError1 := typeConverter.ConvertToDataDefinition(errors.NotFoundBindingType())
	  if(errError1 != nil) {
	    log.Errorf("Error in ConvertToDataDefinition for TagAssociationClientImpl.attach method's errors.NotFound error - %s",
	        bindings.VAPIerrorsToError(errError1).Error())
        return nil
	  }
	  errorDefinitions = append(errorDefinitions, errDef1.(data.ErrorDefinition))
      tIface.errorBindingMap[errors.InvalidArgument{}.Error()] = errors.InvalidArgumentBindingType()
	  errDef2, errError2 := typeConverter.ConvertToDataDefinition(errors.InvalidArgumentBindingType())
	  if(errError2 != nil) {
	    log.Errorf("Error in ConvertToDataDefinition for TagAssociationClientImpl.attach method's errors.InvalidArgument error - %s",
	        bindings.VAPIerrorsToError(errError2).Error())
        return nil
	  }
	  errorDefinitions = append(errorDefinitions, errDef2.(data.ErrorDefinition))
      tIface.errorBindingMap[errors.Unauthorized{}.Error()] = errors.UnauthorizedBindingType()
	  errDef3, errError3 := typeConverter.ConvertToDataDefinition(errors.UnauthorizedBindingType())
	  if(errError3 != nil) {
	    log.Errorf("Error in ConvertToDataDefinition for TagAssociationClientImpl.attach method's errors.Unauthorized error - %s",
	        bindings.VAPIerrorsToError(errError3).Error())
        return nil
	  }
	  errorDefinitions = append(errorDefinitions, errDef3.(data.ErrorDefinition))
      tIface.errorBindingMap[errors.Unauthenticated{}.Error()] = errors.UnauthenticatedBindingType()
	  errDef4, errError4 := typeConverter.ConvertToDataDefinition(errors.UnauthenticatedBindingType())
	  if(errError4 != nil) {
	    log.Errorf("Error in ConvertToDataDefinition for TagAssociationClientImpl.attach method's errors.Unauthenticated error - %s",
	        bindings.VAPIerrorsToError(errError4).Error())
        return nil
	  }
	  errorDefinitions = append(errorDefinitions, errDef4.(data.ErrorDefinition))

      methodDefinition := core.NewMethodDefinition(methodIdentifier, input, output, errorDefinitions)
      return &methodDefinition
}
func (tIface *TagAssociationClientImpl) attachMultipleTagsToObjectMethodDefinition() *core.MethodDefinition {
      interfaceIdentifier := core.NewInterfaceIdentifier(tIface.interfaceName)
      typeConverter := tIface.connector.TypeConverter()

      input, inputError := typeConverter.ConvertToDataDefinition(tagAssociationAttachMultipleTagsToObjectInputType())
      output, outputError := typeConverter.ConvertToDataDefinition(tagAssociationAttachMultipleTagsToObjectOutputType())
      if(inputError != nil) {
          log.Errorf("Error in ConvertToDataDefinition for TagAssociationClientImpl.attachMultipleTagsToObject method's input - %s",
              bindings.VAPIerrorsToError(inputError).Error())
          return nil
      }
      if(outputError != nil) {
          log.Errorf("Error in ConvertToDataDefinition for TagAssociationClientImpl.attachMultipleTagsToObject method's output - %s",
              bindings.VAPIerrorsToError(outputError).Error())
          return nil
      }
      methodIdentifier := core.NewMethodIdentifier(interfaceIdentifier, "attachMultipleTagsToObject")
      errorDefinitions := make([]data.ErrorDefinition,0)
      tIface.errorBindingMap[errors.Unauthorized{}.Error()] = errors.UnauthorizedBindingType()
	  errDef1, errError1 := typeConverter.ConvertToDataDefinition(errors.UnauthorizedBindingType())
	  if(errError1 != nil) {
	    log.Errorf("Error in ConvertToDataDefinition for TagAssociationClientImpl.attachMultipleTagsToObject method's errors.Unauthorized error - %s",
	        bindings.VAPIerrorsToError(errError1).Error())
        return nil
	  }
	  errorDefinitions = append(errorDefinitions, errDef1.(data.ErrorDefinition))
      tIface.errorBindingMap[errors.Unauthenticated{}.Error()] = errors.UnauthenticatedBindingType()
	  errDef2, errError2 := typeConverter.ConvertToDataDefinition(errors.UnauthenticatedBindingType())
	  if(errError2 != nil) {
	    log.Errorf("Error in ConvertToDataDefinition for TagAssociationClientImpl.attachMultipleTagsToObject method's errors.Unauthenticated error - %s",
	        bindings.VAPIerrorsToError(errError2).Error())
        return nil
	  }
	  errorDefinitions = append(errorDefinitions, errDef2.(data.ErrorDefinition))

      methodDefinition := core.NewMethodDefinition(methodIdentifier, input, output, errorDefinitions)
      return &methodDefinition
}
func (tIface *TagAssociationClientImpl) attachTagToMultipleObjectsMethodDefinition() *core.MethodDefinition {
      interfaceIdentifier := core.NewInterfaceIdentifier(tIface.interfaceName)
      typeConverter := tIface.connector.TypeConverter()

      input, inputError := typeConverter.ConvertToDataDefinition(tagAssociationAttachTagToMultipleObjectsInputType())
      output, outputError := typeConverter.ConvertToDataDefinition(tagAssociationAttachTagToMultipleObjectsOutputType())
      if(inputError != nil) {
          log.Errorf("Error in ConvertToDataDefinition for TagAssociationClientImpl.attachTagToMultipleObjects method's input - %s",
              bindings.VAPIerrorsToError(inputError).Error())
          return nil
      }
      if(outputError != nil) {
          log.Errorf("Error in ConvertToDataDefinition for TagAssociationClientImpl.attachTagToMultipleObjects method's output - %s",
              bindings.VAPIerrorsToError(outputError).Error())
          return nil
      }
      methodIdentifier := core.NewMethodIdentifier(interfaceIdentifier, "attachTagToMultipleObjects")
      errorDefinitions := make([]data.ErrorDefinition,0)
      tIface.errorBindingMap[errors.NotFound{}.Error()] = errors.NotFoundBindingType()
	  errDef1, errError1 := typeConverter.ConvertToDataDefinition(errors.NotFoundBindingType())
	  if(errError1 != nil) {
	    log.Errorf("Error in ConvertToDataDefinition for TagAssociationClientImpl.attachTagToMultipleObjects method's errors.NotFound error - %s",
	        bindings.VAPIerrorsToError(errError1).Error())
        return nil
	  }
	  errorDefinitions = append(errorDefinitions, errDef1.(data.ErrorDefinition))
      tIface.errorBindingMap[errors.Unauthorized{}.Error()] = errors.UnauthorizedBindingType()
	  errDef2, errError2 := typeConverter.ConvertToDataDefinition(errors.UnauthorizedBindingType())
	  if(errError2 != nil) {
	    log.Errorf("Error in ConvertToDataDefinition for TagAssociationClientImpl.attachTagToMultipleObjects method's errors.Unauthorized error - %s",
	        bindings.VAPIerrorsToError(errError2).Error())
        return nil
	  }
	  errorDefinitions = append(errorDefinitions, errDef2.(data.ErrorDefinition))
      tIface.errorBindingMap[errors.Unauthenticated{}.Error()] = errors.UnauthenticatedBindingType()
	  errDef3, errError3 := typeConverter.ConvertToDataDefinition(errors.UnauthenticatedBindingType())
	  if(errError3 != nil) {
	    log.Errorf("Error in ConvertToDataDefinition for TagAssociationClientImpl.attachTagToMultipleObjects method's errors.Unauthenticated error - %s",
	        bindings.VAPIerrorsToError(errError3).Error())
        return nil
	  }
	  errorDefinitions = append(errorDefinitions, errDef3.(data.ErrorDefinition))

      methodDefinition := core.NewMethodDefinition(methodIdentifier, input, output, errorDefinitions)
      return &methodDefinition
}
func (tIface *TagAssociationClientImpl) detachMethodDefinition() *core.MethodDefinition {
      interfaceIdentifier := core.NewInterfaceIdentifier(tIface.interfaceName)
      typeConverter := tIface.connector.TypeConverter()

      input, inputError := typeConverter.ConvertToDataDefinition(tagAssociationDetachInputType())
      output, outputError := typeConverter.ConvertToDataDefinition(tagAssociationDetachOutputType())
      if(inputError != nil) {
          log.Errorf("Error in ConvertToDataDefinition for TagAssociationClientImpl.detach method's input - %s",
              bindings.VAPIerrorsToError(inputError).Error())
          return nil
      }
      if(outputError != nil) {
          log.Errorf("Error in ConvertToDataDefinition for TagAssociationClientImpl.detach method's output - %s",
              bindings.VAPIerrorsToError(outputError).Error())
          return nil
      }
      methodIdentifier := core.NewMethodIdentifier(interfaceIdentifier, "detach")
      errorDefinitions := make([]data.ErrorDefinition,0)
      tIface.errorBindingMap[errors.NotFound{}.Error()] = errors.NotFoundBindingType()
	  errDef1, errError1 := typeConverter.ConvertToDataDefinition(errors.NotFoundBindingType())
	  if(errError1 != nil) {
	    log.Errorf("Error in ConvertToDataDefinition for TagAssociationClientImpl.detach method's errors.NotFound error - %s",
	        bindings.VAPIerrorsToError(errError1).Error())
        return nil
	  }
	  errorDefinitions = append(errorDefinitions, errDef1.(data.ErrorDefinition))
      tIface.errorBindingMap[errors.Unauthorized{}.Error()] = errors.UnauthorizedBindingType()
	  errDef2, errError2 := typeConverter.ConvertToDataDefinition(errors.UnauthorizedBindingType())
	  if(errError2 != nil) {
	    log.Errorf("Error in ConvertToDataDefinition for TagAssociationClientImpl.detach method's errors.Unauthorized error - %s",
	        bindings.VAPIerrorsToError(errError2).Error())
        return nil
	  }
	  errorDefinitions = append(errorDefinitions, errDef2.(data.ErrorDefinition))
      tIface.errorBindingMap[errors.Unauthenticated{}.Error()] = errors.UnauthenticatedBindingType()
	  errDef3, errError3 := typeConverter.ConvertToDataDefinition(errors.UnauthenticatedBindingType())
	  if(errError3 != nil) {
	    log.Errorf("Error in ConvertToDataDefinition for TagAssociationClientImpl.detach method's errors.Unauthenticated error - %s",
	        bindings.VAPIerrorsToError(errError3).Error())
        return nil
	  }
	  errorDefinitions = append(errorDefinitions, errDef3.(data.ErrorDefinition))

      methodDefinition := core.NewMethodDefinition(methodIdentifier, input, output, errorDefinitions)
      return &methodDefinition
}
func (tIface *TagAssociationClientImpl) detachMultipleTagsFromObjectMethodDefinition() *core.MethodDefinition {
      interfaceIdentifier := core.NewInterfaceIdentifier(tIface.interfaceName)
      typeConverter := tIface.connector.TypeConverter()

      input, inputError := typeConverter.ConvertToDataDefinition(tagAssociationDetachMultipleTagsFromObjectInputType())
      output, outputError := typeConverter.ConvertToDataDefinition(tagAssociationDetachMultipleTagsFromObjectOutputType())
      if(inputError != nil) {
          log.Errorf("Error in ConvertToDataDefinition for TagAssociationClientImpl.detachMultipleTagsFromObject method's input - %s",
              bindings.VAPIerrorsToError(inputError).Error())
          return nil
      }
      if(outputError != nil) {
          log.Errorf("Error in ConvertToDataDefinition for TagAssociationClientImpl.detachMultipleTagsFromObject method's output - %s",
              bindings.VAPIerrorsToError(outputError).Error())
          return nil
      }
      methodIdentifier := core.NewMethodIdentifier(interfaceIdentifier, "detachMultipleTagsFromObject")
      errorDefinitions := make([]data.ErrorDefinition,0)
      tIface.errorBindingMap[errors.Unauthorized{}.Error()] = errors.UnauthorizedBindingType()
	  errDef1, errError1 := typeConverter.ConvertToDataDefinition(errors.UnauthorizedBindingType())
	  if(errError1 != nil) {
	    log.Errorf("Error in ConvertToDataDefinition for TagAssociationClientImpl.detachMultipleTagsFromObject method's errors.Unauthorized error - %s",
	        bindings.VAPIerrorsToError(errError1).Error())
        return nil
	  }
	  errorDefinitions = append(errorDefinitions, errDef1.(data.ErrorDefinition))
      tIface.errorBindingMap[errors.Unauthenticated{}.Error()] = errors.UnauthenticatedBindingType()
	  errDef2, errError2 := typeConverter.ConvertToDataDefinition(errors.UnauthenticatedBindingType())
	  if(errError2 != nil) {
	    log.Errorf("Error in ConvertToDataDefinition for TagAssociationClientImpl.detachMultipleTagsFromObject method's errors.Unauthenticated error - %s",
	        bindings.VAPIerrorsToError(errError2).Error())
        return nil
	  }
	  errorDefinitions = append(errorDefinitions, errDef2.(data.ErrorDefinition))

      methodDefinition := core.NewMethodDefinition(methodIdentifier, input, output, errorDefinitions)
      return &methodDefinition
}
func (tIface *TagAssociationClientImpl) detachTagFromMultipleObjectsMethodDefinition() *core.MethodDefinition {
      interfaceIdentifier := core.NewInterfaceIdentifier(tIface.interfaceName)
      typeConverter := tIface.connector.TypeConverter()

      input, inputError := typeConverter.ConvertToDataDefinition(tagAssociationDetachTagFromMultipleObjectsInputType())
      output, outputError := typeConverter.ConvertToDataDefinition(tagAssociationDetachTagFromMultipleObjectsOutputType())
      if(inputError != nil) {
          log.Errorf("Error in ConvertToDataDefinition for TagAssociationClientImpl.detachTagFromMultipleObjects method's input - %s",
              bindings.VAPIerrorsToError(inputError).Error())
          return nil
      }
      if(outputError != nil) {
          log.Errorf("Error in ConvertToDataDefinition for TagAssociationClientImpl.detachTagFromMultipleObjects method's output - %s",
              bindings.VAPIerrorsToError(outputError).Error())
          return nil
      }
      methodIdentifier := core.NewMethodIdentifier(interfaceIdentifier, "detachTagFromMultipleObjects")
      errorDefinitions := make([]data.ErrorDefinition,0)
      tIface.errorBindingMap[errors.NotFound{}.Error()] = errors.NotFoundBindingType()
	  errDef1, errError1 := typeConverter.ConvertToDataDefinition(errors.NotFoundBindingType())
	  if(errError1 != nil) {
	    log.Errorf("Error in ConvertToDataDefinition for TagAssociationClientImpl.detachTagFromMultipleObjects method's errors.NotFound error - %s",
	        bindings.VAPIerrorsToError(errError1).Error())
        return nil
	  }
	  errorDefinitions = append(errorDefinitions, errDef1.(data.ErrorDefinition))
      tIface.errorBindingMap[errors.Unauthorized{}.Error()] = errors.UnauthorizedBindingType()
	  errDef2, errError2 := typeConverter.ConvertToDataDefinition(errors.UnauthorizedBindingType())
	  if(errError2 != nil) {
	    log.Errorf("Error in ConvertToDataDefinition for TagAssociationClientImpl.detachTagFromMultipleObjects method's errors.Unauthorized error - %s",
	        bindings.VAPIerrorsToError(errError2).Error())
        return nil
	  }
	  errorDefinitions = append(errorDefinitions, errDef2.(data.ErrorDefinition))
      tIface.errorBindingMap[errors.Unauthenticated{}.Error()] = errors.UnauthenticatedBindingType()
	  errDef3, errError3 := typeConverter.ConvertToDataDefinition(errors.UnauthenticatedBindingType())
	  if(errError3 != nil) {
	    log.Errorf("Error in ConvertToDataDefinition for TagAssociationClientImpl.detachTagFromMultipleObjects method's errors.Unauthenticated error - %s",
	        bindings.VAPIerrorsToError(errError3).Error())
        return nil
	  }
	  errorDefinitions = append(errorDefinitions, errDef3.(data.ErrorDefinition))

      methodDefinition := core.NewMethodDefinition(methodIdentifier, input, output, errorDefinitions)
      return &methodDefinition
}
func (tIface *TagAssociationClientImpl) listAttachedObjectsMethodDefinition() *core.MethodDefinition {
      interfaceIdentifier := core.NewInterfaceIdentifier(tIface.interfaceName)
      typeConverter := tIface.connector.TypeConverter()

      input, inputError := typeConverter.ConvertToDataDefinition(tagAssociationListAttachedObjectsInputType())
      output, outputError := typeConverter.ConvertToDataDefinition(tagAssociationListAttachedObjectsOutputType())
      if(inputError != nil) {
          log.Errorf("Error in ConvertToDataDefinition for TagAssociationClientImpl.listAttachedObjects method's input - %s",
              bindings.VAPIerrorsToError(inputError).Error())
          return nil
      }
      if(outputError != nil) {
          log.Errorf("Error in ConvertToDataDefinition for TagAssociationClientImpl.listAttachedObjects method's output - %s",
              bindings.VAPIerrorsToError(outputError).Error())
          return nil
      }
      methodIdentifier := core.NewMethodIdentifier(interfaceIdentifier, "listAttachedObjects")
      errorDefinitions := make([]data.ErrorDefinition,0)
      tIface.errorBindingMap[errors.NotFound{}.Error()] = errors.NotFoundBindingType()
	  errDef1, errError1 := typeConverter.ConvertToDataDefinition(errors.NotFoundBindingType())
	  if(errError1 != nil) {
	    log.Errorf("Error in ConvertToDataDefinition for TagAssociationClientImpl.listAttachedObjects method's errors.NotFound error - %s",
	        bindings.VAPIerrorsToError(errError1).Error())
        return nil
	  }
	  errorDefinitions = append(errorDefinitions, errDef1.(data.ErrorDefinition))
      tIface.errorBindingMap[errors.Unauthorized{}.Error()] = errors.UnauthorizedBindingType()
	  errDef2, errError2 := typeConverter.ConvertToDataDefinition(errors.UnauthorizedBindingType())
	  if(errError2 != nil) {
	    log.Errorf("Error in ConvertToDataDefinition for TagAssociationClientImpl.listAttachedObjects method's errors.Unauthorized error - %s",
	        bindings.VAPIerrorsToError(errError2).Error())
        return nil
	  }
	  errorDefinitions = append(errorDefinitions, errDef2.(data.ErrorDefinition))
      tIface.errorBindingMap[errors.Unauthenticated{}.Error()] = errors.UnauthenticatedBindingType()
	  errDef3, errError3 := typeConverter.ConvertToDataDefinition(errors.UnauthenticatedBindingType())
	  if(errError3 != nil) {
	    log.Errorf("Error in ConvertToDataDefinition for TagAssociationClientImpl.listAttachedObjects method's errors.Unauthenticated error - %s",
	        bindings.VAPIerrorsToError(errError3).Error())
        return nil
	  }
	  errorDefinitions = append(errorDefinitions, errDef3.(data.ErrorDefinition))

      methodDefinition := core.NewMethodDefinition(methodIdentifier, input, output, errorDefinitions)
      return &methodDefinition
}
func (tIface *TagAssociationClientImpl) listAttachedObjectsOnTagsMethodDefinition() *core.MethodDefinition {
      interfaceIdentifier := core.NewInterfaceIdentifier(tIface.interfaceName)
      typeConverter := tIface.connector.TypeConverter()

      input, inputError := typeConverter.ConvertToDataDefinition(tagAssociationListAttachedObjectsOnTagsInputType())
      output, outputError := typeConverter.ConvertToDataDefinition(tagAssociationListAttachedObjectsOnTagsOutputType())
      if(inputError != nil) {
          log.Errorf("Error in ConvertToDataDefinition for TagAssociationClientImpl.listAttachedObjectsOnTags method's input - %s",
              bindings.VAPIerrorsToError(inputError).Error())
          return nil
      }
      if(outputError != nil) {
          log.Errorf("Error in ConvertToDataDefinition for TagAssociationClientImpl.listAttachedObjectsOnTags method's output - %s",
              bindings.VAPIerrorsToError(outputError).Error())
          return nil
      }
      methodIdentifier := core.NewMethodIdentifier(interfaceIdentifier, "listAttachedObjectsOnTags")
      errorDefinitions := make([]data.ErrorDefinition,0)
      tIface.errorBindingMap[errors.Unauthenticated{}.Error()] = errors.UnauthenticatedBindingType()
	  errDef1, errError1 := typeConverter.ConvertToDataDefinition(errors.UnauthenticatedBindingType())
	  if(errError1 != nil) {
	    log.Errorf("Error in ConvertToDataDefinition for TagAssociationClientImpl.listAttachedObjectsOnTags method's errors.Unauthenticated error - %s",
	        bindings.VAPIerrorsToError(errError1).Error())
        return nil
	  }
	  errorDefinitions = append(errorDefinitions, errDef1.(data.ErrorDefinition))

      methodDefinition := core.NewMethodDefinition(methodIdentifier, input, output, errorDefinitions)
      return &methodDefinition
}
func (tIface *TagAssociationClientImpl) listAttachedTagsMethodDefinition() *core.MethodDefinition {
      interfaceIdentifier := core.NewInterfaceIdentifier(tIface.interfaceName)
      typeConverter := tIface.connector.TypeConverter()

      input, inputError := typeConverter.ConvertToDataDefinition(tagAssociationListAttachedTagsInputType())
      output, outputError := typeConverter.ConvertToDataDefinition(tagAssociationListAttachedTagsOutputType())
      if(inputError != nil) {
          log.Errorf("Error in ConvertToDataDefinition for TagAssociationClientImpl.listAttachedTags method's input - %s",
              bindings.VAPIerrorsToError(inputError).Error())
          return nil
      }
      if(outputError != nil) {
          log.Errorf("Error in ConvertToDataDefinition for TagAssociationClientImpl.listAttachedTags method's output - %s",
              bindings.VAPIerrorsToError(outputError).Error())
          return nil
      }
      methodIdentifier := core.NewMethodIdentifier(interfaceIdentifier, "listAttachedTags")
      errorDefinitions := make([]data.ErrorDefinition,0)
      tIface.errorBindingMap[errors.Unauthorized{}.Error()] = errors.UnauthorizedBindingType()
	  errDef1, errError1 := typeConverter.ConvertToDataDefinition(errors.UnauthorizedBindingType())
	  if(errError1 != nil) {
	    log.Errorf("Error in ConvertToDataDefinition for TagAssociationClientImpl.listAttachedTags method's errors.Unauthorized error - %s",
	        bindings.VAPIerrorsToError(errError1).Error())
        return nil
	  }
	  errorDefinitions = append(errorDefinitions, errDef1.(data.ErrorDefinition))
      tIface.errorBindingMap[errors.Unauthenticated{}.Error()] = errors.UnauthenticatedBindingType()
	  errDef2, errError2 := typeConverter.ConvertToDataDefinition(errors.UnauthenticatedBindingType())
	  if(errError2 != nil) {
	    log.Errorf("Error in ConvertToDataDefinition for TagAssociationClientImpl.listAttachedTags method's errors.Unauthenticated error - %s",
	        bindings.VAPIerrorsToError(errError2).Error())
        return nil
	  }
	  errorDefinitions = append(errorDefinitions, errDef2.(data.ErrorDefinition))

      methodDefinition := core.NewMethodDefinition(methodIdentifier, input, output, errorDefinitions)
      return &methodDefinition
}
func (tIface *TagAssociationClientImpl) listAttachedTagsOnObjectsMethodDefinition() *core.MethodDefinition {
      interfaceIdentifier := core.NewInterfaceIdentifier(tIface.interfaceName)
      typeConverter := tIface.connector.TypeConverter()

      input, inputError := typeConverter.ConvertToDataDefinition(tagAssociationListAttachedTagsOnObjectsInputType())
      output, outputError := typeConverter.ConvertToDataDefinition(tagAssociationListAttachedTagsOnObjectsOutputType())
      if(inputError != nil) {
          log.Errorf("Error in ConvertToDataDefinition for TagAssociationClientImpl.listAttachedTagsOnObjects method's input - %s",
              bindings.VAPIerrorsToError(inputError).Error())
          return nil
      }
      if(outputError != nil) {
          log.Errorf("Error in ConvertToDataDefinition for TagAssociationClientImpl.listAttachedTagsOnObjects method's output - %s",
              bindings.VAPIerrorsToError(outputError).Error())
          return nil
      }
      methodIdentifier := core.NewMethodIdentifier(interfaceIdentifier, "listAttachedTagsOnObjects")
      errorDefinitions := make([]data.ErrorDefinition,0)
      tIface.errorBindingMap[errors.Unauthenticated{}.Error()] = errors.UnauthenticatedBindingType()
	  errDef1, errError1 := typeConverter.ConvertToDataDefinition(errors.UnauthenticatedBindingType())
	  if(errError1 != nil) {
	    log.Errorf("Error in ConvertToDataDefinition for TagAssociationClientImpl.listAttachedTagsOnObjects method's errors.Unauthenticated error - %s",
	        bindings.VAPIerrorsToError(errError1).Error())
        return nil
	  }
	  errorDefinitions = append(errorDefinitions, errDef1.(data.ErrorDefinition))

      methodDefinition := core.NewMethodDefinition(methodIdentifier, input, output, errorDefinitions)
      return &methodDefinition
}
func (tIface *TagAssociationClientImpl) listAttachableTagsMethodDefinition() *core.MethodDefinition {
      interfaceIdentifier := core.NewInterfaceIdentifier(tIface.interfaceName)
      typeConverter := tIface.connector.TypeConverter()

      input, inputError := typeConverter.ConvertToDataDefinition(tagAssociationListAttachableTagsInputType())
      output, outputError := typeConverter.ConvertToDataDefinition(tagAssociationListAttachableTagsOutputType())
      if(inputError != nil) {
          log.Errorf("Error in ConvertToDataDefinition for TagAssociationClientImpl.listAttachableTags method's input - %s",
              bindings.VAPIerrorsToError(inputError).Error())
          return nil
      }
      if(outputError != nil) {
          log.Errorf("Error in ConvertToDataDefinition for TagAssociationClientImpl.listAttachableTags method's output - %s",
              bindings.VAPIerrorsToError(outputError).Error())
          return nil
      }
      methodIdentifier := core.NewMethodIdentifier(interfaceIdentifier, "listAttachableTags")
      errorDefinitions := make([]data.ErrorDefinition,0)
      tIface.errorBindingMap[errors.Unauthorized{}.Error()] = errors.UnauthorizedBindingType()
	  errDef1, errError1 := typeConverter.ConvertToDataDefinition(errors.UnauthorizedBindingType())
	  if(errError1 != nil) {
	    log.Errorf("Error in ConvertToDataDefinition for TagAssociationClientImpl.listAttachableTags method's errors.Unauthorized error - %s",
	        bindings.VAPIerrorsToError(errError1).Error())
        return nil
	  }
	  errorDefinitions = append(errorDefinitions, errDef1.(data.ErrorDefinition))
      tIface.errorBindingMap[errors.Unauthenticated{}.Error()] = errors.UnauthenticatedBindingType()
	  errDef2, errError2 := typeConverter.ConvertToDataDefinition(errors.UnauthenticatedBindingType())
	  if(errError2 != nil) {
	    log.Errorf("Error in ConvertToDataDefinition for TagAssociationClientImpl.listAttachableTags method's errors.Unauthenticated error - %s",
	        bindings.VAPIerrorsToError(errError2).Error())
        return nil
	  }
	  errorDefinitions = append(errorDefinitions, errDef2.(data.ErrorDefinition))

      methodDefinition := core.NewMethodDefinition(methodIdentifier, input, output, errorDefinitions)
      return &methodDefinition
}


