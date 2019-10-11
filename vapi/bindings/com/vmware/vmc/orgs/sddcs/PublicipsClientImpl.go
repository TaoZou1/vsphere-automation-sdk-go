
/* Copyright © 2019 VMware, Inc. All Rights Reserved.
     SPDX-License-Identifier: BSD-2-Clause */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Client stubs for service: Publicips
 * Functions that implement the generated PublicipsClient interface
 */


package sddcs

import (
    "reflect"
// //     "getClientImplDependenciesOfOps gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/bindings/com/vmware/vapi/std/errors"
// //     "getClientImplDependenciesOfOps gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/bindings/com/vmware/vmc/model"
// //     "getClientImplDependenciesOfOps gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/bindings"
// //     "getClientImplDependenciesOfOps gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/core"
// //     "getClientImplDependenciesOfOps gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/data"
// //     "getClientImplDependenciesOfOps gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/lib"
// //     "getClientImplDependenciesOfOps gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/log"
// //     "getClientImplDependenciesOfOps gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/protocol/client"
// //
// //     "getDependenciesOfServiceTypes gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/bindings/com/vmware/vmc/model"
// //     "getDependenciesOfServiceTypes gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/bindings"
// //     "getDependenciesOfServiceTypes gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/data"
// //     "getDependenciesOfServiceTypes gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/protocol"
// 
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/bindings/com/vmware/vapi/std/errors"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/bindings/com/vmware/vmc/model"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/bindings"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/core"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/data"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/lib"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/log"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/protocol"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/protocol/client"
)


type PublicipsClientImpl struct{
      interfaceName string
      interfaceDefinition core.InterfaceDefinition
      methodIdentifiers[]core.MethodIdentifier
      methodNameToDefMap map[string]*core.MethodDefinition
      errorBindingMap map[string]bindings.BindingType
      interfaceIdentifier core.InterfaceIdentifier
      connector client.Connector
}


func NewPublicipsClientImpl(connector client.Connector) *PublicipsClientImpl {
      interfaceName := "com.vmware.vmc.orgs.sddcs.publicips"
      interfaceIdentifier := core.NewInterfaceIdentifier(interfaceName)
      methodIdentifiers := []core.MethodIdentifier{
          core.NewMethodIdentifier(interfaceIdentifier, "create"),
          core.NewMethodIdentifier(interfaceIdentifier, "delete"),
          core.NewMethodIdentifier(interfaceIdentifier, "get"),
          core.NewMethodIdentifier(interfaceIdentifier, "list"),
          core.NewMethodIdentifier(interfaceIdentifier, "update"),
      }
      interfaceDefinition := core.NewInterfaceDefinition(interfaceIdentifier, methodIdentifiers)
      errorBindingMap := make(map[string]bindings.BindingType)
      errorBindingMap[errors.InternalServerError{}.Error()] = errors.InternalServerErrorBindingType()
	  errorBindingMap[errors.InvalidArgument{}.Error()] = errors.InvalidArgumentBindingType()
	  errorBindingMap[errors.OperationNotFound{}.Error()] = errors.OperationNotFoundBindingType()
	  errorBindingMap[errors.UnexpectedInput{}.Error()] = errors.UnexpectedInputBindingType()
	  errorBindingMap[errors.ServiceUnavailable{}.Error()] = errors.ServiceUnavailableBindingType()
	  errorBindingMap[errors.TimedOut{}.Error()] = errors.TimedOutBindingType()
      pIface := PublicipsClientImpl{interfaceName: interfaceName, methodIdentifiers: methodIdentifiers, interfaceDefinition: interfaceDefinition, errorBindingMap: errorBindingMap, interfaceIdentifier:interfaceIdentifier, connector: connector}
      pIface.methodNameToDefMap = make(map[string]*core.MethodDefinition)
      pIface.methodNameToDefMap["create"] = pIface.createMethodDefinition()
      pIface.methodNameToDefMap["delete"] = pIface.deleteMethodDefinition()
      pIface.methodNameToDefMap["get"] = pIface.getMethodDefinition()
      pIface.methodNameToDefMap["list"] = pIface.listMethodDefinition()
      pIface.methodNameToDefMap["update"] = pIface.updateMethodDefinition()
      return &pIface
}

func (pIface *PublicipsClientImpl) Create(orgParam string, sddcParam string, specParam model.SddcAllocatePublicIpSpec) (model.Task, error) {
	typeConverter := pIface.connector.TypeConverter()
	methodIdentifier := core.NewMethodIdentifier(pIface.interfaceIdentifier, "create")
	sv := bindings.NewStructValueBuilder(publicipsCreateInputType(), typeConverter)
	sv.AddStructField("Org", orgParam)
	sv.AddStructField("Sddc", sddcParam)
	sv.AddStructField("Spec", specParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
        var emptyOutput model.Task
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := publicipsCreateRestMetadata
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	pIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult:= pIface.Invoke(pIface.connector.NewExecutionContext(), methodIdentifier, inputDataValue)
	var emptyOutput model.Task
    if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), publicipsCreateOutputType())
		if errorInOutput != nil {
		    return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
	    return output.(model.Task), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), pIface.errorBindingMap[methodResult.Error().Name()])
		if errorInError != nil {
		    return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}
func (pIface *PublicipsClientImpl) Delete(orgParam string, sddcParam string, idParam string) (model.Task, error) {
	typeConverter := pIface.connector.TypeConverter()
	methodIdentifier := core.NewMethodIdentifier(pIface.interfaceIdentifier, "delete")
	sv := bindings.NewStructValueBuilder(publicipsDeleteInputType(), typeConverter)
	sv.AddStructField("Org", orgParam)
	sv.AddStructField("Sddc", sddcParam)
	sv.AddStructField("Id", idParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
        var emptyOutput model.Task
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := publicipsDeleteRestMetadata
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	pIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult:= pIface.Invoke(pIface.connector.NewExecutionContext(), methodIdentifier, inputDataValue)
	var emptyOutput model.Task
    if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), publicipsDeleteOutputType())
		if errorInOutput != nil {
		    return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
	    return output.(model.Task), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), pIface.errorBindingMap[methodResult.Error().Name()])
		if errorInError != nil {
		    return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}
func (pIface *PublicipsClientImpl) Get(orgParam string, sddcParam string, idParam string) (model.SddcPublicIp, error) {
	typeConverter := pIface.connector.TypeConverter()
	methodIdentifier := core.NewMethodIdentifier(pIface.interfaceIdentifier, "get")
	sv := bindings.NewStructValueBuilder(publicipsGetInputType(), typeConverter)
	sv.AddStructField("Org", orgParam)
	sv.AddStructField("Sddc", sddcParam)
	sv.AddStructField("Id", idParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
        var emptyOutput model.SddcPublicIp
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := publicipsGetRestMetadata
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	pIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult:= pIface.Invoke(pIface.connector.NewExecutionContext(), methodIdentifier, inputDataValue)
	var emptyOutput model.SddcPublicIp
    if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), publicipsGetOutputType())
		if errorInOutput != nil {
		    return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
	    return output.(model.SddcPublicIp), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), pIface.errorBindingMap[methodResult.Error().Name()])
		if errorInError != nil {
		    return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}
func (pIface *PublicipsClientImpl) List(orgParam string, sddcParam string) ([]model.SddcPublicIp, error) {
	typeConverter := pIface.connector.TypeConverter()
	methodIdentifier := core.NewMethodIdentifier(pIface.interfaceIdentifier, "list")
	sv := bindings.NewStructValueBuilder(publicipsListInputType(), typeConverter)
	sv.AddStructField("Org", orgParam)
	sv.AddStructField("Sddc", sddcParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
        var emptyOutput []model.SddcPublicIp
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := publicipsListRestMetadata
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	pIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult:= pIface.Invoke(pIface.connector.NewExecutionContext(), methodIdentifier, inputDataValue)
	var emptyOutput []model.SddcPublicIp
    if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), publicipsListOutputType())
		if errorInOutput != nil {
		    return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
	    return output.([]model.SddcPublicIp), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), pIface.errorBindingMap[methodResult.Error().Name()])
		if errorInError != nil {
		    return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}
func (pIface *PublicipsClientImpl) Update(orgParam string, sddcParam string, idParam string, actionParam string, sddcPublicIpObjectParam model.SddcPublicIp) (model.Task, error) {
	typeConverter := pIface.connector.TypeConverter()
	methodIdentifier := core.NewMethodIdentifier(pIface.interfaceIdentifier, "update")
	sv := bindings.NewStructValueBuilder(publicipsUpdateInputType(), typeConverter)
	sv.AddStructField("Org", orgParam)
	sv.AddStructField("Sddc", sddcParam)
	sv.AddStructField("Id", idParam)
	sv.AddStructField("Action", actionParam)
	sv.AddStructField("SddcPublicIpObject", sddcPublicIpObjectParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
        var emptyOutput model.Task
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := publicipsUpdateRestMetadata
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	pIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult:= pIface.Invoke(pIface.connector.NewExecutionContext(), methodIdentifier, inputDataValue)
	var emptyOutput model.Task
    if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), publicipsUpdateOutputType())
		if errorInOutput != nil {
		    return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
	    return output.(model.Task), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), pIface.errorBindingMap[methodResult.Error().Name()])
		if errorInError != nil {
		    return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (pIface *PublicipsClientImpl) Invoke(ctx *core.ExecutionContext, methodId core.MethodIdentifier, inputDataValue data.DataValue) core.MethodResult {
    methodResult := pIface.connector.GetApiProvider().Invoke(pIface.interfaceName, methodId.Name(), inputDataValue, ctx)
    return methodResult
}

func (pIface *PublicipsClientImpl) createMethodDefinition() *core.MethodDefinition {
      interfaceIdentifier := core.NewInterfaceIdentifier(pIface.interfaceName)
      typeConverter := pIface.connector.TypeConverter()

      input, inputError := typeConverter.ConvertToDataDefinition(publicipsCreateInputType())
      output, outputError := typeConverter.ConvertToDataDefinition(publicipsCreateOutputType())
      if(inputError != nil) {
          log.Errorf("Error in ConvertToDataDefinition for PublicipsClientImpl.create method's input - %s",
              bindings.VAPIerrorsToError(inputError).Error())
          return nil
      }
      if(outputError != nil) {
          log.Errorf("Error in ConvertToDataDefinition for PublicipsClientImpl.create method's output - %s",
              bindings.VAPIerrorsToError(outputError).Error())
          return nil
      }
      methodIdentifier := core.NewMethodIdentifier(interfaceIdentifier, "create")
      errorDefinitions := make([]data.ErrorDefinition,0)
      pIface.errorBindingMap[errors.Unauthenticated{}.Error()] = errors.UnauthenticatedBindingType()
	  errDef1, errError1 := typeConverter.ConvertToDataDefinition(errors.UnauthenticatedBindingType())
	  if(errError1 != nil) {
	    log.Errorf("Error in ConvertToDataDefinition for PublicipsClientImpl.create method's errors.Unauthenticated error - %s",
	        bindings.VAPIerrorsToError(errError1).Error())
        return nil
	  }
	  errorDefinitions = append(errorDefinitions, errDef1.(data.ErrorDefinition))
      pIface.errorBindingMap[errors.InvalidRequest{}.Error()] = errors.InvalidRequestBindingType()
	  errDef2, errError2 := typeConverter.ConvertToDataDefinition(errors.InvalidRequestBindingType())
	  if(errError2 != nil) {
	    log.Errorf("Error in ConvertToDataDefinition for PublicipsClientImpl.create method's errors.InvalidRequest error - %s",
	        bindings.VAPIerrorsToError(errError2).Error())
        return nil
	  }
	  errorDefinitions = append(errorDefinitions, errDef2.(data.ErrorDefinition))
      pIface.errorBindingMap[errors.Unauthorized{}.Error()] = errors.UnauthorizedBindingType()
	  errDef3, errError3 := typeConverter.ConvertToDataDefinition(errors.UnauthorizedBindingType())
	  if(errError3 != nil) {
	    log.Errorf("Error in ConvertToDataDefinition for PublicipsClientImpl.create method's errors.Unauthorized error - %s",
	        bindings.VAPIerrorsToError(errError3).Error())
        return nil
	  }
	  errorDefinitions = append(errorDefinitions, errDef3.(data.ErrorDefinition))
      pIface.errorBindingMap[errors.NotFound{}.Error()] = errors.NotFoundBindingType()
	  errDef4, errError4 := typeConverter.ConvertToDataDefinition(errors.NotFoundBindingType())
	  if(errError4 != nil) {
	    log.Errorf("Error in ConvertToDataDefinition for PublicipsClientImpl.create method's errors.NotFound error - %s",
	        bindings.VAPIerrorsToError(errError4).Error())
        return nil
	  }
	  errorDefinitions = append(errorDefinitions, errDef4.(data.ErrorDefinition))

      methodDefinition := core.NewMethodDefinition(methodIdentifier, input, output, errorDefinitions)
      return &methodDefinition
}
func (pIface *PublicipsClientImpl) deleteMethodDefinition() *core.MethodDefinition {
      interfaceIdentifier := core.NewInterfaceIdentifier(pIface.interfaceName)
      typeConverter := pIface.connector.TypeConverter()

      input, inputError := typeConverter.ConvertToDataDefinition(publicipsDeleteInputType())
      output, outputError := typeConverter.ConvertToDataDefinition(publicipsDeleteOutputType())
      if(inputError != nil) {
          log.Errorf("Error in ConvertToDataDefinition for PublicipsClientImpl.delete method's input - %s",
              bindings.VAPIerrorsToError(inputError).Error())
          return nil
      }
      if(outputError != nil) {
          log.Errorf("Error in ConvertToDataDefinition for PublicipsClientImpl.delete method's output - %s",
              bindings.VAPIerrorsToError(outputError).Error())
          return nil
      }
      methodIdentifier := core.NewMethodIdentifier(interfaceIdentifier, "delete")
      errorDefinitions := make([]data.ErrorDefinition,0)
      pIface.errorBindingMap[errors.Unauthenticated{}.Error()] = errors.UnauthenticatedBindingType()
	  errDef1, errError1 := typeConverter.ConvertToDataDefinition(errors.UnauthenticatedBindingType())
	  if(errError1 != nil) {
	    log.Errorf("Error in ConvertToDataDefinition for PublicipsClientImpl.delete method's errors.Unauthenticated error - %s",
	        bindings.VAPIerrorsToError(errError1).Error())
        return nil
	  }
	  errorDefinitions = append(errorDefinitions, errDef1.(data.ErrorDefinition))
      pIface.errorBindingMap[errors.InvalidRequest{}.Error()] = errors.InvalidRequestBindingType()
	  errDef2, errError2 := typeConverter.ConvertToDataDefinition(errors.InvalidRequestBindingType())
	  if(errError2 != nil) {
	    log.Errorf("Error in ConvertToDataDefinition for PublicipsClientImpl.delete method's errors.InvalidRequest error - %s",
	        bindings.VAPIerrorsToError(errError2).Error())
        return nil
	  }
	  errorDefinitions = append(errorDefinitions, errDef2.(data.ErrorDefinition))
      pIface.errorBindingMap[errors.Unauthorized{}.Error()] = errors.UnauthorizedBindingType()
	  errDef3, errError3 := typeConverter.ConvertToDataDefinition(errors.UnauthorizedBindingType())
	  if(errError3 != nil) {
	    log.Errorf("Error in ConvertToDataDefinition for PublicipsClientImpl.delete method's errors.Unauthorized error - %s",
	        bindings.VAPIerrorsToError(errError3).Error())
        return nil
	  }
	  errorDefinitions = append(errorDefinitions, errDef3.(data.ErrorDefinition))
      pIface.errorBindingMap[errors.NotFound{}.Error()] = errors.NotFoundBindingType()
	  errDef4, errError4 := typeConverter.ConvertToDataDefinition(errors.NotFoundBindingType())
	  if(errError4 != nil) {
	    log.Errorf("Error in ConvertToDataDefinition for PublicipsClientImpl.delete method's errors.NotFound error - %s",
	        bindings.VAPIerrorsToError(errError4).Error())
        return nil
	  }
	  errorDefinitions = append(errorDefinitions, errDef4.(data.ErrorDefinition))

      methodDefinition := core.NewMethodDefinition(methodIdentifier, input, output, errorDefinitions)
      return &methodDefinition
}
func (pIface *PublicipsClientImpl) getMethodDefinition() *core.MethodDefinition {
      interfaceIdentifier := core.NewInterfaceIdentifier(pIface.interfaceName)
      typeConverter := pIface.connector.TypeConverter()

      input, inputError := typeConverter.ConvertToDataDefinition(publicipsGetInputType())
      output, outputError := typeConverter.ConvertToDataDefinition(publicipsGetOutputType())
      if(inputError != nil) {
          log.Errorf("Error in ConvertToDataDefinition for PublicipsClientImpl.get method's input - %s",
              bindings.VAPIerrorsToError(inputError).Error())
          return nil
      }
      if(outputError != nil) {
          log.Errorf("Error in ConvertToDataDefinition for PublicipsClientImpl.get method's output - %s",
              bindings.VAPIerrorsToError(outputError).Error())
          return nil
      }
      methodIdentifier := core.NewMethodIdentifier(interfaceIdentifier, "get")
      errorDefinitions := make([]data.ErrorDefinition,0)
      pIface.errorBindingMap[errors.Unauthenticated{}.Error()] = errors.UnauthenticatedBindingType()
	  errDef1, errError1 := typeConverter.ConvertToDataDefinition(errors.UnauthenticatedBindingType())
	  if(errError1 != nil) {
	    log.Errorf("Error in ConvertToDataDefinition for PublicipsClientImpl.get method's errors.Unauthenticated error - %s",
	        bindings.VAPIerrorsToError(errError1).Error())
        return nil
	  }
	  errorDefinitions = append(errorDefinitions, errDef1.(data.ErrorDefinition))
      pIface.errorBindingMap[errors.Unauthorized{}.Error()] = errors.UnauthorizedBindingType()
	  errDef2, errError2 := typeConverter.ConvertToDataDefinition(errors.UnauthorizedBindingType())
	  if(errError2 != nil) {
	    log.Errorf("Error in ConvertToDataDefinition for PublicipsClientImpl.get method's errors.Unauthorized error - %s",
	        bindings.VAPIerrorsToError(errError2).Error())
        return nil
	  }
	  errorDefinitions = append(errorDefinitions, errDef2.(data.ErrorDefinition))
      pIface.errorBindingMap[errors.NotFound{}.Error()] = errors.NotFoundBindingType()
	  errDef3, errError3 := typeConverter.ConvertToDataDefinition(errors.NotFoundBindingType())
	  if(errError3 != nil) {
	    log.Errorf("Error in ConvertToDataDefinition for PublicipsClientImpl.get method's errors.NotFound error - %s",
	        bindings.VAPIerrorsToError(errError3).Error())
        return nil
	  }
	  errorDefinitions = append(errorDefinitions, errDef3.(data.ErrorDefinition))

      methodDefinition := core.NewMethodDefinition(methodIdentifier, input, output, errorDefinitions)
      return &methodDefinition
}
func (pIface *PublicipsClientImpl) listMethodDefinition() *core.MethodDefinition {
      interfaceIdentifier := core.NewInterfaceIdentifier(pIface.interfaceName)
      typeConverter := pIface.connector.TypeConverter()

      input, inputError := typeConverter.ConvertToDataDefinition(publicipsListInputType())
      output, outputError := typeConverter.ConvertToDataDefinition(publicipsListOutputType())
      if(inputError != nil) {
          log.Errorf("Error in ConvertToDataDefinition for PublicipsClientImpl.list method's input - %s",
              bindings.VAPIerrorsToError(inputError).Error())
          return nil
      }
      if(outputError != nil) {
          log.Errorf("Error in ConvertToDataDefinition for PublicipsClientImpl.list method's output - %s",
              bindings.VAPIerrorsToError(outputError).Error())
          return nil
      }
      methodIdentifier := core.NewMethodIdentifier(interfaceIdentifier, "list")
      errorDefinitions := make([]data.ErrorDefinition,0)
      pIface.errorBindingMap[errors.Unauthenticated{}.Error()] = errors.UnauthenticatedBindingType()
	  errDef1, errError1 := typeConverter.ConvertToDataDefinition(errors.UnauthenticatedBindingType())
	  if(errError1 != nil) {
	    log.Errorf("Error in ConvertToDataDefinition for PublicipsClientImpl.list method's errors.Unauthenticated error - %s",
	        bindings.VAPIerrorsToError(errError1).Error())
        return nil
	  }
	  errorDefinitions = append(errorDefinitions, errDef1.(data.ErrorDefinition))
      pIface.errorBindingMap[errors.Unauthorized{}.Error()] = errors.UnauthorizedBindingType()
	  errDef2, errError2 := typeConverter.ConvertToDataDefinition(errors.UnauthorizedBindingType())
	  if(errError2 != nil) {
	    log.Errorf("Error in ConvertToDataDefinition for PublicipsClientImpl.list method's errors.Unauthorized error - %s",
	        bindings.VAPIerrorsToError(errError2).Error())
        return nil
	  }
	  errorDefinitions = append(errorDefinitions, errDef2.(data.ErrorDefinition))
      pIface.errorBindingMap[errors.NotFound{}.Error()] = errors.NotFoundBindingType()
	  errDef3, errError3 := typeConverter.ConvertToDataDefinition(errors.NotFoundBindingType())
	  if(errError3 != nil) {
	    log.Errorf("Error in ConvertToDataDefinition for PublicipsClientImpl.list method's errors.NotFound error - %s",
	        bindings.VAPIerrorsToError(errError3).Error())
        return nil
	  }
	  errorDefinitions = append(errorDefinitions, errDef3.(data.ErrorDefinition))

      methodDefinition := core.NewMethodDefinition(methodIdentifier, input, output, errorDefinitions)
      return &methodDefinition
}
func (pIface *PublicipsClientImpl) updateMethodDefinition() *core.MethodDefinition {
      interfaceIdentifier := core.NewInterfaceIdentifier(pIface.interfaceName)
      typeConverter := pIface.connector.TypeConverter()

      input, inputError := typeConverter.ConvertToDataDefinition(publicipsUpdateInputType())
      output, outputError := typeConverter.ConvertToDataDefinition(publicipsUpdateOutputType())
      if(inputError != nil) {
          log.Errorf("Error in ConvertToDataDefinition for PublicipsClientImpl.update method's input - %s",
              bindings.VAPIerrorsToError(inputError).Error())
          return nil
      }
      if(outputError != nil) {
          log.Errorf("Error in ConvertToDataDefinition for PublicipsClientImpl.update method's output - %s",
              bindings.VAPIerrorsToError(outputError).Error())
          return nil
      }
      methodIdentifier := core.NewMethodIdentifier(interfaceIdentifier, "update")
      errorDefinitions := make([]data.ErrorDefinition,0)
      pIface.errorBindingMap[errors.Unauthenticated{}.Error()] = errors.UnauthenticatedBindingType()
	  errDef1, errError1 := typeConverter.ConvertToDataDefinition(errors.UnauthenticatedBindingType())
	  if(errError1 != nil) {
	    log.Errorf("Error in ConvertToDataDefinition for PublicipsClientImpl.update method's errors.Unauthenticated error - %s",
	        bindings.VAPIerrorsToError(errError1).Error())
        return nil
	  }
	  errorDefinitions = append(errorDefinitions, errDef1.(data.ErrorDefinition))
      pIface.errorBindingMap[errors.InvalidRequest{}.Error()] = errors.InvalidRequestBindingType()
	  errDef2, errError2 := typeConverter.ConvertToDataDefinition(errors.InvalidRequestBindingType())
	  if(errError2 != nil) {
	    log.Errorf("Error in ConvertToDataDefinition for PublicipsClientImpl.update method's errors.InvalidRequest error - %s",
	        bindings.VAPIerrorsToError(errError2).Error())
        return nil
	  }
	  errorDefinitions = append(errorDefinitions, errDef2.(data.ErrorDefinition))
      pIface.errorBindingMap[errors.Unauthorized{}.Error()] = errors.UnauthorizedBindingType()
	  errDef3, errError3 := typeConverter.ConvertToDataDefinition(errors.UnauthorizedBindingType())
	  if(errError3 != nil) {
	    log.Errorf("Error in ConvertToDataDefinition for PublicipsClientImpl.update method's errors.Unauthorized error - %s",
	        bindings.VAPIerrorsToError(errError3).Error())
        return nil
	  }
	  errorDefinitions = append(errorDefinitions, errDef3.(data.ErrorDefinition))
      pIface.errorBindingMap[errors.NotFound{}.Error()] = errors.NotFoundBindingType()
	  errDef4, errError4 := typeConverter.ConvertToDataDefinition(errors.NotFoundBindingType())
	  if(errError4 != nil) {
	    log.Errorf("Error in ConvertToDataDefinition for PublicipsClientImpl.update method's errors.NotFound error - %s",
	        bindings.VAPIerrorsToError(errError4).Error())
        return nil
	  }
	  errorDefinitions = append(errorDefinitions, errDef4.(data.ErrorDefinition))

      methodDefinition := core.NewMethodDefinition(methodIdentifier, input, output, errorDefinitions)
      return &methodDefinition
}







func publicipsCreateInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["org"] = bindings.NewStringType()
    fields["sddc"] = bindings.NewStringType()
    fields["spec"] = bindings.NewReferenceType(model.SddcAllocatePublicIpSpecBindingType)
    fieldNameMap["org"] = "Org"
    fieldNameMap["sddc"] = "Sddc"
    fieldNameMap["spec"] = "Spec"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func publicipsCreateOutputType() bindings.BindingType {
    return bindings.NewReferenceType(model.TaskBindingType)
}

func publicipsCreateRestMetadata() protocol.OperationRestMetadata {
    paramsTypeMap := map[string]bindings.BindingType{}
    pathParams := map[string]string{}
    queryParams := map[string]string{}
    headerParams := map[string]string{}
    paramsTypeMap["org"] = bindings.NewStringType()
    paramsTypeMap["sddc"] = bindings.NewStringType()
    paramsTypeMap["spec"] = bindings.NewReferenceType(model.SddcAllocatePublicIpSpecBindingType)
    paramsTypeMap["org"] = bindings.NewStringType()
    paramsTypeMap["sddc"] = bindings.NewStringType()
    pathParams["org"] = "org"
    pathParams["sddc"] = "sddc"
    resultHeaders := map[string]string{}
    errorHeaders := map[string]string{}
    errorHeaders["Unauthenticated.challenge"] = "WWW-Authenticate"
    return protocol.NewOperationRestMetadata(
      paramsTypeMap,
      pathParams,
      queryParams,
      headerParams,
      "spec",
      "POST",
      "/vmc/api/orgs/{org}/sddcs/{sddc}/publicips",
       resultHeaders,
       201,
       errorHeaders,
       map[string]int{"Unauthenticated": 401,"InvalidRequest": 400,"Unauthorized": 403,"NotFound": 404})
}


func publicipsDeleteInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["org"] = bindings.NewStringType()
    fields["sddc"] = bindings.NewStringType()
    fields["id"] = bindings.NewStringType()
    fieldNameMap["org"] = "Org"
    fieldNameMap["sddc"] = "Sddc"
    fieldNameMap["id"] = "Id"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func publicipsDeleteOutputType() bindings.BindingType {
    return bindings.NewReferenceType(model.TaskBindingType)
}

func publicipsDeleteRestMetadata() protocol.OperationRestMetadata {
    paramsTypeMap := map[string]bindings.BindingType{}
    pathParams := map[string]string{}
    queryParams := map[string]string{}
    headerParams := map[string]string{}
    paramsTypeMap["org"] = bindings.NewStringType()
    paramsTypeMap["sddc"] = bindings.NewStringType()
    paramsTypeMap["id"] = bindings.NewStringType()
    paramsTypeMap["org"] = bindings.NewStringType()
    paramsTypeMap["sddc"] = bindings.NewStringType()
    paramsTypeMap["id"] = bindings.NewStringType()
    pathParams["org"] = "org"
    pathParams["sddc"] = "sddc"
    pathParams["id"] = "id"
    resultHeaders := map[string]string{}
    errorHeaders := map[string]string{}
    errorHeaders["Unauthenticated.challenge"] = "WWW-Authenticate"
    return protocol.NewOperationRestMetadata(
      paramsTypeMap,
      pathParams,
      queryParams,
      headerParams,
      "",
      "DELETE",
      "/vmc/api/orgs/{org}/sddcs/{sddc}/publicips/{id}",
       resultHeaders,
       200,
       errorHeaders,
       map[string]int{"Unauthenticated": 401,"InvalidRequest": 400,"Unauthorized": 403,"NotFound": 404})
}


func publicipsGetInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["org"] = bindings.NewStringType()
    fields["sddc"] = bindings.NewStringType()
    fields["id"] = bindings.NewStringType()
    fieldNameMap["org"] = "Org"
    fieldNameMap["sddc"] = "Sddc"
    fieldNameMap["id"] = "Id"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func publicipsGetOutputType() bindings.BindingType {
    return bindings.NewReferenceType(model.SddcPublicIpBindingType)
}

func publicipsGetRestMetadata() protocol.OperationRestMetadata {
    paramsTypeMap := map[string]bindings.BindingType{}
    pathParams := map[string]string{}
    queryParams := map[string]string{}
    headerParams := map[string]string{}
    paramsTypeMap["org"] = bindings.NewStringType()
    paramsTypeMap["sddc"] = bindings.NewStringType()
    paramsTypeMap["id"] = bindings.NewStringType()
    paramsTypeMap["org"] = bindings.NewStringType()
    paramsTypeMap["sddc"] = bindings.NewStringType()
    paramsTypeMap["id"] = bindings.NewStringType()
    pathParams["org"] = "org"
    pathParams["sddc"] = "sddc"
    pathParams["id"] = "id"
    resultHeaders := map[string]string{}
    errorHeaders := map[string]string{}
    errorHeaders["Unauthenticated.challenge"] = "WWW-Authenticate"
    return protocol.NewOperationRestMetadata(
      paramsTypeMap,
      pathParams,
      queryParams,
      headerParams,
      "",
      "GET",
      "/vmc/api/orgs/{org}/sddcs/{sddc}/publicips/{id}",
       resultHeaders,
       200,
       errorHeaders,
       map[string]int{"Unauthenticated": 401,"Unauthorized": 403,"NotFound": 404})
}


func publicipsListInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["org"] = bindings.NewStringType()
    fields["sddc"] = bindings.NewStringType()
    fieldNameMap["org"] = "Org"
    fieldNameMap["sddc"] = "Sddc"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func publicipsListOutputType() bindings.BindingType {
    return bindings.NewListType(bindings.NewReferenceType(model.SddcPublicIpBindingType), reflect.TypeOf([]model.SddcPublicIp{}))
}

func publicipsListRestMetadata() protocol.OperationRestMetadata {
    paramsTypeMap := map[string]bindings.BindingType{}
    pathParams := map[string]string{}
    queryParams := map[string]string{}
    headerParams := map[string]string{}
    paramsTypeMap["org"] = bindings.NewStringType()
    paramsTypeMap["sddc"] = bindings.NewStringType()
    paramsTypeMap["org"] = bindings.NewStringType()
    paramsTypeMap["sddc"] = bindings.NewStringType()
    pathParams["org"] = "org"
    pathParams["sddc"] = "sddc"
    resultHeaders := map[string]string{}
    errorHeaders := map[string]string{}
    errorHeaders["Unauthenticated.challenge"] = "WWW-Authenticate"
    return protocol.NewOperationRestMetadata(
      paramsTypeMap,
      pathParams,
      queryParams,
      headerParams,
      "",
      "GET",
      "/vmc/api/orgs/{org}/sddcs/{sddc}/publicips",
       resultHeaders,
       200,
       errorHeaders,
       map[string]int{"Unauthenticated": 401,"Unauthorized": 403,"NotFound": 404})
}


func publicipsUpdateInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["org"] = bindings.NewStringType()
    fields["sddc"] = bindings.NewStringType()
    fields["id"] = bindings.NewStringType()
    fields["action"] = bindings.NewStringType()
    fields["sddc_public_ip_object"] = bindings.NewReferenceType(model.SddcPublicIpBindingType)
    fieldNameMap["org"] = "Org"
    fieldNameMap["sddc"] = "Sddc"
    fieldNameMap["id"] = "Id"
    fieldNameMap["action"] = "Action"
    fieldNameMap["sddc_public_ip_object"] = "SddcPublicIpObject"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func publicipsUpdateOutputType() bindings.BindingType {
    return bindings.NewReferenceType(model.TaskBindingType)
}

func publicipsUpdateRestMetadata() protocol.OperationRestMetadata {
    paramsTypeMap := map[string]bindings.BindingType{}
    pathParams := map[string]string{}
    queryParams := map[string]string{}
    headerParams := map[string]string{}
    paramsTypeMap["action"] = bindings.NewStringType()
    paramsTypeMap["org"] = bindings.NewStringType()
    paramsTypeMap["sddc"] = bindings.NewStringType()
    paramsTypeMap["id"] = bindings.NewStringType()
    paramsTypeMap["sddc_public_ip_object"] = bindings.NewReferenceType(model.SddcPublicIpBindingType)
    paramsTypeMap["org"] = bindings.NewStringType()
    paramsTypeMap["sddc"] = bindings.NewStringType()
    paramsTypeMap["id"] = bindings.NewStringType()
    pathParams["org"] = "org"
    pathParams["sddc"] = "sddc"
    pathParams["id"] = "id"
    queryParams["action"] = "action"
    resultHeaders := map[string]string{}
    errorHeaders := map[string]string{}
    errorHeaders["Unauthenticated.challenge"] = "WWW-Authenticate"
    return protocol.NewOperationRestMetadata(
      paramsTypeMap,
      pathParams,
      queryParams,
      headerParams,
      "sddc_public_ip_object",
      "PATCH",
      "/vmc/api/orgs/{org}/sddcs/{sddc}/publicips/{id}",
       resultHeaders,
       200,
       errorHeaders,
       map[string]int{"Unauthenticated": 401,"InvalidRequest": 400,"Unauthorized": 403,"NotFound": 404})
}




