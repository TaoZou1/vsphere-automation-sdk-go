
/* Copyright © 2019 VMware, Inc. All Rights Reserved.
     SPDX-License-Identifier: BSD-2-Clause */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Client stubs for service: AssociatedProducts
 * Functions that implement the generated AssociatedProductsClient interface
 */


package discovery
import (
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/bindings/com/vmware/vapi/std/errors"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/bindings"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/core"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/data"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/lib"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/log"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/protocol/client"
)


type AssociatedProductsClientImpl struct{
      interfaceName string
      interfaceDefinition core.InterfaceDefinition
      methodIdentifiers[]core.MethodIdentifier
      methodNameToDefMap map[string]*core.MethodDefinition
      errorBindingMap map[string]bindings.BindingType
      interfaceIdentifier core.InterfaceIdentifier
      connector client.Connector
}


func NewAssociatedProductsClientImpl(connector client.Connector) *AssociatedProductsClientImpl {
      interfaceName := "com.vmware.vcenter.lcm.discovery.associated_products"
      interfaceIdentifier := core.NewInterfaceIdentifier(interfaceName)
      methodIdentifiers := []core.MethodIdentifier{
          core.NewMethodIdentifier(interfaceIdentifier, "list"),
          core.NewMethodIdentifier(interfaceIdentifier, "create"),
          core.NewMethodIdentifier(interfaceIdentifier, "get"),
          core.NewMethodIdentifier(interfaceIdentifier, "update"),
          core.NewMethodIdentifier(interfaceIdentifier, "delete"),
      }
      interfaceDefinition := core.NewInterfaceDefinition(interfaceIdentifier, methodIdentifiers)
      errorBindingMap := make(map[string]bindings.BindingType)
      errorBindingMap[errors.InternalServerError{}.Error()] = errors.InternalServerErrorBindingType()
	  errorBindingMap[errors.InvalidArgument{}.Error()] = errors.InvalidArgumentBindingType()
	  errorBindingMap[errors.OperationNotFound{}.Error()] = errors.OperationNotFoundBindingType()
	  errorBindingMap[errors.UnexpectedInput{}.Error()] = errors.UnexpectedInputBindingType()
	  errorBindingMap[errors.ServiceUnavailable{}.Error()] = errors.ServiceUnavailableBindingType()
	  errorBindingMap[errors.TimedOut{}.Error()] = errors.TimedOutBindingType()
      aIface := AssociatedProductsClientImpl{interfaceName: interfaceName, methodIdentifiers: methodIdentifiers, interfaceDefinition: interfaceDefinition, errorBindingMap: errorBindingMap, interfaceIdentifier:interfaceIdentifier, connector: connector}
      aIface.methodNameToDefMap = make(map[string]*core.MethodDefinition)
      aIface.methodNameToDefMap["list"] = aIface.listMethodDefinition()
      aIface.methodNameToDefMap["create"] = aIface.createMethodDefinition()
      aIface.methodNameToDefMap["get"] = aIface.getMethodDefinition()
      aIface.methodNameToDefMap["update"] = aIface.updateMethodDefinition()
      aIface.methodNameToDefMap["delete"] = aIface.deleteMethodDefinition()
      return &aIface
}

func (aIface *AssociatedProductsClientImpl) List() ([]Product, error) {
	typeConverter := aIface.connector.TypeConverter()
	methodIdentifier := core.NewMethodIdentifier(aIface.interfaceIdentifier, "list")
	sv := bindings.NewStructValueBuilder(associatedProductsListInputType(), typeConverter)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
        var emptyOutput []Product
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := associatedProductsListRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	aIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult:= aIface.Invoke(aIface.connector.NewExecutionContext(), methodIdentifier, inputDataValue)
	var emptyOutput []Product
    if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), associatedProductsListOutputType())
		if errorInOutput != nil {
		    return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
	    return output.([]Product), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), aIface.errorBindingMap[methodResult.Error().Name()])
		if errorInError != nil {
		    return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}
func (aIface *AssociatedProductsClientImpl) Create(specParam AssociatedProductsCreateSpec) (string, error) {
	typeConverter := aIface.connector.TypeConverter()
	methodIdentifier := core.NewMethodIdentifier(aIface.interfaceIdentifier, "create")
	sv := bindings.NewStructValueBuilder(associatedProductsCreateInputType(), typeConverter)
	sv.AddStructField("Spec", specParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
        var emptyOutput string
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := associatedProductsCreateRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	aIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult:= aIface.Invoke(aIface.connector.NewExecutionContext(), methodIdentifier, inputDataValue)
	var emptyOutput string
    if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), associatedProductsCreateOutputType())
		if errorInOutput != nil {
		    return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
	    return output.(string), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), aIface.errorBindingMap[methodResult.Error().Name()])
		if errorInError != nil {
		    return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}
func (aIface *AssociatedProductsClientImpl) Get(productParam string) (Product, error) {
	typeConverter := aIface.connector.TypeConverter()
	methodIdentifier := core.NewMethodIdentifier(aIface.interfaceIdentifier, "get")
	sv := bindings.NewStructValueBuilder(associatedProductsGetInputType(), typeConverter)
	sv.AddStructField("Product", productParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
        var emptyOutput Product
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := associatedProductsGetRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	aIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult:= aIface.Invoke(aIface.connector.NewExecutionContext(), methodIdentifier, inputDataValue)
	var emptyOutput Product
    if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), associatedProductsGetOutputType())
		if errorInOutput != nil {
		    return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
	    return output.(Product), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), aIface.errorBindingMap[methodResult.Error().Name()])
		if errorInError != nil {
		    return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}
func (aIface *AssociatedProductsClientImpl) Update(productParam string, specParam AssociatedProductsUpdateSpec) error {
	typeConverter := aIface.connector.TypeConverter()
	methodIdentifier := core.NewMethodIdentifier(aIface.interfaceIdentifier, "update")
	sv := bindings.NewStructValueBuilder(associatedProductsUpdateInputType(), typeConverter)
	sv.AddStructField("Product", productParam)
	sv.AddStructField("Spec", specParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		return bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := associatedProductsUpdateRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	aIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult:= aIface.Invoke(aIface.connector.NewExecutionContext(), methodIdentifier, inputDataValue)
    if methodResult.IsSuccess() {
		return nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), aIface.errorBindingMap[methodResult.Error().Name()])
		if errorInError != nil {
		    return bindings.VAPIerrorsToError(errorInError)
		}
		return methodError.(error)
	}
}
func (aIface *AssociatedProductsClientImpl) Delete(productParam string) error {
	typeConverter := aIface.connector.TypeConverter()
	methodIdentifier := core.NewMethodIdentifier(aIface.interfaceIdentifier, "delete")
	sv := bindings.NewStructValueBuilder(associatedProductsDeleteInputType(), typeConverter)
	sv.AddStructField("Product", productParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		return bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := associatedProductsDeleteRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	aIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult:= aIface.Invoke(aIface.connector.NewExecutionContext(), methodIdentifier, inputDataValue)
    if methodResult.IsSuccess() {
		return nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), aIface.errorBindingMap[methodResult.Error().Name()])
		if errorInError != nil {
		    return bindings.VAPIerrorsToError(errorInError)
		}
		return methodError.(error)
	}
}

func (aIface *AssociatedProductsClientImpl) Invoke(ctx *core.ExecutionContext, methodId core.MethodIdentifier, inputDataValue data.DataValue) core.MethodResult {
    methodResult := aIface.connector.GetApiProvider().Invoke(aIface.interfaceName, methodId.Name(), inputDataValue, ctx)
    return methodResult
}

func (aIface *AssociatedProductsClientImpl) listMethodDefinition() *core.MethodDefinition {
      interfaceIdentifier := core.NewInterfaceIdentifier(aIface.interfaceName)
      typeConverter := aIface.connector.TypeConverter()

      input, inputError := typeConverter.ConvertToDataDefinition(associatedProductsListInputType())
      output, outputError := typeConverter.ConvertToDataDefinition(associatedProductsListOutputType())
      if(inputError != nil) {
          log.Errorf("Error in ConvertToDataDefinition for AssociatedProductsClientImpl.list method's input - %s",
              bindings.VAPIerrorsToError(inputError).Error())
          return nil
      }
      if(outputError != nil) {
          log.Errorf("Error in ConvertToDataDefinition for AssociatedProductsClientImpl.list method's output - %s",
              bindings.VAPIerrorsToError(outputError).Error())
          return nil
      }
      methodIdentifier := core.NewMethodIdentifier(interfaceIdentifier, "list")
      errorDefinitions := make([]data.ErrorDefinition,0)
      aIface.errorBindingMap[errors.Unauthenticated{}.Error()] = errors.UnauthenticatedBindingType()
	  errDef1, errError1 := typeConverter.ConvertToDataDefinition(errors.UnauthenticatedBindingType())
	  if(errError1 != nil) {
	    log.Errorf("Error in ConvertToDataDefinition for AssociatedProductsClientImpl.list method's errors.Unauthenticated error - %s",
	        bindings.VAPIerrorsToError(errError1).Error())
        return nil
	  }
	  errorDefinitions = append(errorDefinitions, errDef1.(data.ErrorDefinition))
      aIface.errorBindingMap[errors.Error{}.Error()] = errors.ErrorBindingType()
	  errDef2, errError2 := typeConverter.ConvertToDataDefinition(errors.ErrorBindingType())
	  if(errError2 != nil) {
	    log.Errorf("Error in ConvertToDataDefinition for AssociatedProductsClientImpl.list method's errors.Error error - %s",
	        bindings.VAPIerrorsToError(errError2).Error())
        return nil
	  }
	  errorDefinitions = append(errorDefinitions, errDef2.(data.ErrorDefinition))

      methodDefinition := core.NewMethodDefinition(methodIdentifier, input, output, errorDefinitions)
      return &methodDefinition
}
func (aIface *AssociatedProductsClientImpl) createMethodDefinition() *core.MethodDefinition {
      interfaceIdentifier := core.NewInterfaceIdentifier(aIface.interfaceName)
      typeConverter := aIface.connector.TypeConverter()

      input, inputError := typeConverter.ConvertToDataDefinition(associatedProductsCreateInputType())
      output, outputError := typeConverter.ConvertToDataDefinition(associatedProductsCreateOutputType())
      if(inputError != nil) {
          log.Errorf("Error in ConvertToDataDefinition for AssociatedProductsClientImpl.create method's input - %s",
              bindings.VAPIerrorsToError(inputError).Error())
          return nil
      }
      if(outputError != nil) {
          log.Errorf("Error in ConvertToDataDefinition for AssociatedProductsClientImpl.create method's output - %s",
              bindings.VAPIerrorsToError(outputError).Error())
          return nil
      }
      methodIdentifier := core.NewMethodIdentifier(interfaceIdentifier, "create")
      errorDefinitions := make([]data.ErrorDefinition,0)
      aIface.errorBindingMap[errors.Unauthenticated{}.Error()] = errors.UnauthenticatedBindingType()
	  errDef1, errError1 := typeConverter.ConvertToDataDefinition(errors.UnauthenticatedBindingType())
	  if(errError1 != nil) {
	    log.Errorf("Error in ConvertToDataDefinition for AssociatedProductsClientImpl.create method's errors.Unauthenticated error - %s",
	        bindings.VAPIerrorsToError(errError1).Error())
        return nil
	  }
	  errorDefinitions = append(errorDefinitions, errDef1.(data.ErrorDefinition))
      aIface.errorBindingMap[errors.AlreadyExists{}.Error()] = errors.AlreadyExistsBindingType()
	  errDef2, errError2 := typeConverter.ConvertToDataDefinition(errors.AlreadyExistsBindingType())
	  if(errError2 != nil) {
	    log.Errorf("Error in ConvertToDataDefinition for AssociatedProductsClientImpl.create method's errors.AlreadyExists error - %s",
	        bindings.VAPIerrorsToError(errError2).Error())
        return nil
	  }
	  errorDefinitions = append(errorDefinitions, errDef2.(data.ErrorDefinition))
      aIface.errorBindingMap[errors.InvalidArgument{}.Error()] = errors.InvalidArgumentBindingType()
	  errDef3, errError3 := typeConverter.ConvertToDataDefinition(errors.InvalidArgumentBindingType())
	  if(errError3 != nil) {
	    log.Errorf("Error in ConvertToDataDefinition for AssociatedProductsClientImpl.create method's errors.InvalidArgument error - %s",
	        bindings.VAPIerrorsToError(errError3).Error())
        return nil
	  }
	  errorDefinitions = append(errorDefinitions, errDef3.(data.ErrorDefinition))
      aIface.errorBindingMap[errors.Error{}.Error()] = errors.ErrorBindingType()
	  errDef4, errError4 := typeConverter.ConvertToDataDefinition(errors.ErrorBindingType())
	  if(errError4 != nil) {
	    log.Errorf("Error in ConvertToDataDefinition for AssociatedProductsClientImpl.create method's errors.Error error - %s",
	        bindings.VAPIerrorsToError(errError4).Error())
        return nil
	  }
	  errorDefinitions = append(errorDefinitions, errDef4.(data.ErrorDefinition))

      methodDefinition := core.NewMethodDefinition(methodIdentifier, input, output, errorDefinitions)
      return &methodDefinition
}
func (aIface *AssociatedProductsClientImpl) getMethodDefinition() *core.MethodDefinition {
      interfaceIdentifier := core.NewInterfaceIdentifier(aIface.interfaceName)
      typeConverter := aIface.connector.TypeConverter()

      input, inputError := typeConverter.ConvertToDataDefinition(associatedProductsGetInputType())
      output, outputError := typeConverter.ConvertToDataDefinition(associatedProductsGetOutputType())
      if(inputError != nil) {
          log.Errorf("Error in ConvertToDataDefinition for AssociatedProductsClientImpl.get method's input - %s",
              bindings.VAPIerrorsToError(inputError).Error())
          return nil
      }
      if(outputError != nil) {
          log.Errorf("Error in ConvertToDataDefinition for AssociatedProductsClientImpl.get method's output - %s",
              bindings.VAPIerrorsToError(outputError).Error())
          return nil
      }
      methodIdentifier := core.NewMethodIdentifier(interfaceIdentifier, "get")
      errorDefinitions := make([]data.ErrorDefinition,0)
      aIface.errorBindingMap[errors.Unauthenticated{}.Error()] = errors.UnauthenticatedBindingType()
	  errDef1, errError1 := typeConverter.ConvertToDataDefinition(errors.UnauthenticatedBindingType())
	  if(errError1 != nil) {
	    log.Errorf("Error in ConvertToDataDefinition for AssociatedProductsClientImpl.get method's errors.Unauthenticated error - %s",
	        bindings.VAPIerrorsToError(errError1).Error())
        return nil
	  }
	  errorDefinitions = append(errorDefinitions, errDef1.(data.ErrorDefinition))
      aIface.errorBindingMap[errors.NotFound{}.Error()] = errors.NotFoundBindingType()
	  errDef2, errError2 := typeConverter.ConvertToDataDefinition(errors.NotFoundBindingType())
	  if(errError2 != nil) {
	    log.Errorf("Error in ConvertToDataDefinition for AssociatedProductsClientImpl.get method's errors.NotFound error - %s",
	        bindings.VAPIerrorsToError(errError2).Error())
        return nil
	  }
	  errorDefinitions = append(errorDefinitions, errDef2.(data.ErrorDefinition))
      aIface.errorBindingMap[errors.Error{}.Error()] = errors.ErrorBindingType()
	  errDef3, errError3 := typeConverter.ConvertToDataDefinition(errors.ErrorBindingType())
	  if(errError3 != nil) {
	    log.Errorf("Error in ConvertToDataDefinition for AssociatedProductsClientImpl.get method's errors.Error error - %s",
	        bindings.VAPIerrorsToError(errError3).Error())
        return nil
	  }
	  errorDefinitions = append(errorDefinitions, errDef3.(data.ErrorDefinition))

      methodDefinition := core.NewMethodDefinition(methodIdentifier, input, output, errorDefinitions)
      return &methodDefinition
}
func (aIface *AssociatedProductsClientImpl) updateMethodDefinition() *core.MethodDefinition {
      interfaceIdentifier := core.NewInterfaceIdentifier(aIface.interfaceName)
      typeConverter := aIface.connector.TypeConverter()

      input, inputError := typeConverter.ConvertToDataDefinition(associatedProductsUpdateInputType())
      output, outputError := typeConverter.ConvertToDataDefinition(associatedProductsUpdateOutputType())
      if(inputError != nil) {
          log.Errorf("Error in ConvertToDataDefinition for AssociatedProductsClientImpl.update method's input - %s",
              bindings.VAPIerrorsToError(inputError).Error())
          return nil
      }
      if(outputError != nil) {
          log.Errorf("Error in ConvertToDataDefinition for AssociatedProductsClientImpl.update method's output - %s",
              bindings.VAPIerrorsToError(outputError).Error())
          return nil
      }
      methodIdentifier := core.NewMethodIdentifier(interfaceIdentifier, "update")
      errorDefinitions := make([]data.ErrorDefinition,0)
      aIface.errorBindingMap[errors.Unauthenticated{}.Error()] = errors.UnauthenticatedBindingType()
	  errDef1, errError1 := typeConverter.ConvertToDataDefinition(errors.UnauthenticatedBindingType())
	  if(errError1 != nil) {
	    log.Errorf("Error in ConvertToDataDefinition for AssociatedProductsClientImpl.update method's errors.Unauthenticated error - %s",
	        bindings.VAPIerrorsToError(errError1).Error())
        return nil
	  }
	  errorDefinitions = append(errorDefinitions, errDef1.(data.ErrorDefinition))
      aIface.errorBindingMap[errors.InvalidArgument{}.Error()] = errors.InvalidArgumentBindingType()
	  errDef2, errError2 := typeConverter.ConvertToDataDefinition(errors.InvalidArgumentBindingType())
	  if(errError2 != nil) {
	    log.Errorf("Error in ConvertToDataDefinition for AssociatedProductsClientImpl.update method's errors.InvalidArgument error - %s",
	        bindings.VAPIerrorsToError(errError2).Error())
        return nil
	  }
	  errorDefinitions = append(errorDefinitions, errDef2.(data.ErrorDefinition))
      aIface.errorBindingMap[errors.NotFound{}.Error()] = errors.NotFoundBindingType()
	  errDef3, errError3 := typeConverter.ConvertToDataDefinition(errors.NotFoundBindingType())
	  if(errError3 != nil) {
	    log.Errorf("Error in ConvertToDataDefinition for AssociatedProductsClientImpl.update method's errors.NotFound error - %s",
	        bindings.VAPIerrorsToError(errError3).Error())
        return nil
	  }
	  errorDefinitions = append(errorDefinitions, errDef3.(data.ErrorDefinition))
      aIface.errorBindingMap[errors.Error{}.Error()] = errors.ErrorBindingType()
	  errDef4, errError4 := typeConverter.ConvertToDataDefinition(errors.ErrorBindingType())
	  if(errError4 != nil) {
	    log.Errorf("Error in ConvertToDataDefinition for AssociatedProductsClientImpl.update method's errors.Error error - %s",
	        bindings.VAPIerrorsToError(errError4).Error())
        return nil
	  }
	  errorDefinitions = append(errorDefinitions, errDef4.(data.ErrorDefinition))

      methodDefinition := core.NewMethodDefinition(methodIdentifier, input, output, errorDefinitions)
      return &methodDefinition
}
func (aIface *AssociatedProductsClientImpl) deleteMethodDefinition() *core.MethodDefinition {
      interfaceIdentifier := core.NewInterfaceIdentifier(aIface.interfaceName)
      typeConverter := aIface.connector.TypeConverter()

      input, inputError := typeConverter.ConvertToDataDefinition(associatedProductsDeleteInputType())
      output, outputError := typeConverter.ConvertToDataDefinition(associatedProductsDeleteOutputType())
      if(inputError != nil) {
          log.Errorf("Error in ConvertToDataDefinition for AssociatedProductsClientImpl.delete method's input - %s",
              bindings.VAPIerrorsToError(inputError).Error())
          return nil
      }
      if(outputError != nil) {
          log.Errorf("Error in ConvertToDataDefinition for AssociatedProductsClientImpl.delete method's output - %s",
              bindings.VAPIerrorsToError(outputError).Error())
          return nil
      }
      methodIdentifier := core.NewMethodIdentifier(interfaceIdentifier, "delete")
      errorDefinitions := make([]data.ErrorDefinition,0)
      aIface.errorBindingMap[errors.Unauthenticated{}.Error()] = errors.UnauthenticatedBindingType()
	  errDef1, errError1 := typeConverter.ConvertToDataDefinition(errors.UnauthenticatedBindingType())
	  if(errError1 != nil) {
	    log.Errorf("Error in ConvertToDataDefinition for AssociatedProductsClientImpl.delete method's errors.Unauthenticated error - %s",
	        bindings.VAPIerrorsToError(errError1).Error())
        return nil
	  }
	  errorDefinitions = append(errorDefinitions, errDef1.(data.ErrorDefinition))
      aIface.errorBindingMap[errors.NotFound{}.Error()] = errors.NotFoundBindingType()
	  errDef2, errError2 := typeConverter.ConvertToDataDefinition(errors.NotFoundBindingType())
	  if(errError2 != nil) {
	    log.Errorf("Error in ConvertToDataDefinition for AssociatedProductsClientImpl.delete method's errors.NotFound error - %s",
	        bindings.VAPIerrorsToError(errError2).Error())
        return nil
	  }
	  errorDefinitions = append(errorDefinitions, errDef2.(data.ErrorDefinition))
      aIface.errorBindingMap[errors.Error{}.Error()] = errors.ErrorBindingType()
	  errDef3, errError3 := typeConverter.ConvertToDataDefinition(errors.ErrorBindingType())
	  if(errError3 != nil) {
	    log.Errorf("Error in ConvertToDataDefinition for AssociatedProductsClientImpl.delete method's errors.Error error - %s",
	        bindings.VAPIerrorsToError(errError3).Error())
        return nil
	  }
	  errorDefinitions = append(errorDefinitions, errDef3.(data.ErrorDefinition))

      methodDefinition := core.NewMethodDefinition(methodIdentifier, input, output, errorDefinitions)
      return &methodDefinition
}


