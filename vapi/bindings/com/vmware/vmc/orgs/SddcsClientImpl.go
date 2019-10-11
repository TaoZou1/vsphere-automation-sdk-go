
/* Copyright © 2019 VMware, Inc. All Rights Reserved.
     SPDX-License-Identifier: BSD-2-Clause */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Client stubs for service: Sddcs
 * Functions that implement the generated SddcsClient interface
 */


package orgs

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


type SddcsClientImpl struct{
      interfaceName string
      interfaceDefinition core.InterfaceDefinition
      methodIdentifiers[]core.MethodIdentifier
      methodNameToDefMap map[string]*core.MethodDefinition
      errorBindingMap map[string]bindings.BindingType
      interfaceIdentifier core.InterfaceIdentifier
      connector client.Connector
}


func NewSddcsClientImpl(connector client.Connector) *SddcsClientImpl {
      interfaceName := "com.vmware.vmc.orgs.sddcs"
      interfaceIdentifier := core.NewInterfaceIdentifier(interfaceName)
      methodIdentifiers := []core.MethodIdentifier{
          core.NewMethodIdentifier(interfaceIdentifier, "create"),
          core.NewMethodIdentifier(interfaceIdentifier, "delete"),
          core.NewMethodIdentifier(interfaceIdentifier, "get"),
          core.NewMethodIdentifier(interfaceIdentifier, "list"),
          core.NewMethodIdentifier(interfaceIdentifier, "patch"),
      }
      interfaceDefinition := core.NewInterfaceDefinition(interfaceIdentifier, methodIdentifiers)
      errorBindingMap := make(map[string]bindings.BindingType)
      errorBindingMap[errors.InternalServerError{}.Error()] = errors.InternalServerErrorBindingType()
	  errorBindingMap[errors.InvalidArgument{}.Error()] = errors.InvalidArgumentBindingType()
	  errorBindingMap[errors.OperationNotFound{}.Error()] = errors.OperationNotFoundBindingType()
	  errorBindingMap[errors.UnexpectedInput{}.Error()] = errors.UnexpectedInputBindingType()
	  errorBindingMap[errors.ServiceUnavailable{}.Error()] = errors.ServiceUnavailableBindingType()
	  errorBindingMap[errors.TimedOut{}.Error()] = errors.TimedOutBindingType()
      sIface := SddcsClientImpl{interfaceName: interfaceName, methodIdentifiers: methodIdentifiers, interfaceDefinition: interfaceDefinition, errorBindingMap: errorBindingMap, interfaceIdentifier:interfaceIdentifier, connector: connector}
      sIface.methodNameToDefMap = make(map[string]*core.MethodDefinition)
      sIface.methodNameToDefMap["create"] = sIface.createMethodDefinition()
      sIface.methodNameToDefMap["delete"] = sIface.deleteMethodDefinition()
      sIface.methodNameToDefMap["get"] = sIface.getMethodDefinition()
      sIface.methodNameToDefMap["list"] = sIface.listMethodDefinition()
      sIface.methodNameToDefMap["patch"] = sIface.patchMethodDefinition()
      return &sIface
}

func (sIface *SddcsClientImpl) Create(orgParam string, sddcConfigParam model.AwsSddcConfig) (model.Task, error) {
	typeConverter := sIface.connector.TypeConverter()
	methodIdentifier := core.NewMethodIdentifier(sIface.interfaceIdentifier, "create")
	sv := bindings.NewStructValueBuilder(sddcsCreateInputType(), typeConverter)
	sv.AddStructField("Org", orgParam)
	sv.AddStructField("SddcConfig", sddcConfigParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
        var emptyOutput model.Task
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := sddcsCreateRestMetadata
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	sIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult:= sIface.Invoke(sIface.connector.NewExecutionContext(), methodIdentifier, inputDataValue)
	var emptyOutput model.Task
    if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), sddcsCreateOutputType())
		if errorInOutput != nil {
		    return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
	    return output.(model.Task), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), sIface.errorBindingMap[methodResult.Error().Name()])
		if errorInError != nil {
		    return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}
func (sIface *SddcsClientImpl) Delete(orgParam string, sddcParam string, retainConfigurationParam *bool, templateNameParam *string, forceParam *string) (model.Task, error) {
	typeConverter := sIface.connector.TypeConverter()
	methodIdentifier := core.NewMethodIdentifier(sIface.interfaceIdentifier, "delete")
	sv := bindings.NewStructValueBuilder(sddcsDeleteInputType(), typeConverter)
	sv.AddStructField("Org", orgParam)
	sv.AddStructField("Sddc", sddcParam)
	sv.AddStructField("RetainConfiguration", retainConfigurationParam)
	sv.AddStructField("TemplateName", templateNameParam)
	sv.AddStructField("Force", forceParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
        var emptyOutput model.Task
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := sddcsDeleteRestMetadata
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	sIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult:= sIface.Invoke(sIface.connector.NewExecutionContext(), methodIdentifier, inputDataValue)
	var emptyOutput model.Task
    if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), sddcsDeleteOutputType())
		if errorInOutput != nil {
		    return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
	    return output.(model.Task), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), sIface.errorBindingMap[methodResult.Error().Name()])
		if errorInError != nil {
		    return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}
func (sIface *SddcsClientImpl) Get(orgParam string, sddcParam string) (model.Sddc, error) {
	typeConverter := sIface.connector.TypeConverter()
	methodIdentifier := core.NewMethodIdentifier(sIface.interfaceIdentifier, "get")
	sv := bindings.NewStructValueBuilder(sddcsGetInputType(), typeConverter)
	sv.AddStructField("Org", orgParam)
	sv.AddStructField("Sddc", sddcParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
        var emptyOutput model.Sddc
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := sddcsGetRestMetadata
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	sIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult:= sIface.Invoke(sIface.connector.NewExecutionContext(), methodIdentifier, inputDataValue)
	var emptyOutput model.Sddc
    if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), sddcsGetOutputType())
		if errorInOutput != nil {
		    return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
	    return output.(model.Sddc), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), sIface.errorBindingMap[methodResult.Error().Name()])
		if errorInError != nil {
		    return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}
func (sIface *SddcsClientImpl) List(orgParam string) ([]model.Sddc, error) {
	typeConverter := sIface.connector.TypeConverter()
	methodIdentifier := core.NewMethodIdentifier(sIface.interfaceIdentifier, "list")
	sv := bindings.NewStructValueBuilder(sddcsListInputType(), typeConverter)
	sv.AddStructField("Org", orgParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
        var emptyOutput []model.Sddc
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := sddcsListRestMetadata
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	sIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult:= sIface.Invoke(sIface.connector.NewExecutionContext(), methodIdentifier, inputDataValue)
	var emptyOutput []model.Sddc
    if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), sddcsListOutputType())
		if errorInOutput != nil {
		    return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
	    return output.([]model.Sddc), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), sIface.errorBindingMap[methodResult.Error().Name()])
		if errorInError != nil {
		    return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}
func (sIface *SddcsClientImpl) Patch(orgParam string, sddcParam string, sddcPatchRequestParam model.SddcPatchRequest) (model.Sddc, error) {
	typeConverter := sIface.connector.TypeConverter()
	methodIdentifier := core.NewMethodIdentifier(sIface.interfaceIdentifier, "patch")
	sv := bindings.NewStructValueBuilder(sddcsPatchInputType(), typeConverter)
	sv.AddStructField("Org", orgParam)
	sv.AddStructField("Sddc", sddcParam)
	sv.AddStructField("SddcPatchRequest", sddcPatchRequestParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
        var emptyOutput model.Sddc
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := sddcsPatchRestMetadata
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	sIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult:= sIface.Invoke(sIface.connector.NewExecutionContext(), methodIdentifier, inputDataValue)
	var emptyOutput model.Sddc
    if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), sddcsPatchOutputType())
		if errorInOutput != nil {
		    return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
	    return output.(model.Sddc), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), sIface.errorBindingMap[methodResult.Error().Name()])
		if errorInError != nil {
		    return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (sIface *SddcsClientImpl) Invoke(ctx *core.ExecutionContext, methodId core.MethodIdentifier, inputDataValue data.DataValue) core.MethodResult {
    methodResult := sIface.connector.GetApiProvider().Invoke(sIface.interfaceName, methodId.Name(), inputDataValue, ctx)
    return methodResult
}

func (sIface *SddcsClientImpl) createMethodDefinition() *core.MethodDefinition {
      interfaceIdentifier := core.NewInterfaceIdentifier(sIface.interfaceName)
      typeConverter := sIface.connector.TypeConverter()

      input, inputError := typeConverter.ConvertToDataDefinition(sddcsCreateInputType())
      output, outputError := typeConverter.ConvertToDataDefinition(sddcsCreateOutputType())
      if(inputError != nil) {
          log.Errorf("Error in ConvertToDataDefinition for SddcsClientImpl.create method's input - %s",
              bindings.VAPIerrorsToError(inputError).Error())
          return nil
      }
      if(outputError != nil) {
          log.Errorf("Error in ConvertToDataDefinition for SddcsClientImpl.create method's output - %s",
              bindings.VAPIerrorsToError(outputError).Error())
          return nil
      }
      methodIdentifier := core.NewMethodIdentifier(interfaceIdentifier, "create")
      errorDefinitions := make([]data.ErrorDefinition,0)
      sIface.errorBindingMap[errors.Unauthenticated{}.Error()] = errors.UnauthenticatedBindingType()
	  errDef1, errError1 := typeConverter.ConvertToDataDefinition(errors.UnauthenticatedBindingType())
	  if(errError1 != nil) {
	    log.Errorf("Error in ConvertToDataDefinition for SddcsClientImpl.create method's errors.Unauthenticated error - %s",
	        bindings.VAPIerrorsToError(errError1).Error())
        return nil
	  }
	  errorDefinitions = append(errorDefinitions, errDef1.(data.ErrorDefinition))
      sIface.errorBindingMap[errors.Unauthorized{}.Error()] = errors.UnauthorizedBindingType()
	  errDef2, errError2 := typeConverter.ConvertToDataDefinition(errors.UnauthorizedBindingType())
	  if(errError2 != nil) {
	    log.Errorf("Error in ConvertToDataDefinition for SddcsClientImpl.create method's errors.Unauthorized error - %s",
	        bindings.VAPIerrorsToError(errError2).Error())
        return nil
	  }
	  errorDefinitions = append(errorDefinitions, errDef2.(data.ErrorDefinition))

      methodDefinition := core.NewMethodDefinition(methodIdentifier, input, output, errorDefinitions)
      return &methodDefinition
}
func (sIface *SddcsClientImpl) deleteMethodDefinition() *core.MethodDefinition {
      interfaceIdentifier := core.NewInterfaceIdentifier(sIface.interfaceName)
      typeConverter := sIface.connector.TypeConverter()

      input, inputError := typeConverter.ConvertToDataDefinition(sddcsDeleteInputType())
      output, outputError := typeConverter.ConvertToDataDefinition(sddcsDeleteOutputType())
      if(inputError != nil) {
          log.Errorf("Error in ConvertToDataDefinition for SddcsClientImpl.delete method's input - %s",
              bindings.VAPIerrorsToError(inputError).Error())
          return nil
      }
      if(outputError != nil) {
          log.Errorf("Error in ConvertToDataDefinition for SddcsClientImpl.delete method's output - %s",
              bindings.VAPIerrorsToError(outputError).Error())
          return nil
      }
      methodIdentifier := core.NewMethodIdentifier(interfaceIdentifier, "delete")
      errorDefinitions := make([]data.ErrorDefinition,0)
      sIface.errorBindingMap[errors.Unauthenticated{}.Error()] = errors.UnauthenticatedBindingType()
	  errDef1, errError1 := typeConverter.ConvertToDataDefinition(errors.UnauthenticatedBindingType())
	  if(errError1 != nil) {
	    log.Errorf("Error in ConvertToDataDefinition for SddcsClientImpl.delete method's errors.Unauthenticated error - %s",
	        bindings.VAPIerrorsToError(errError1).Error())
        return nil
	  }
	  errorDefinitions = append(errorDefinitions, errDef1.(data.ErrorDefinition))
      sIface.errorBindingMap[errors.InvalidRequest{}.Error()] = errors.InvalidRequestBindingType()
	  errDef2, errError2 := typeConverter.ConvertToDataDefinition(errors.InvalidRequestBindingType())
	  if(errError2 != nil) {
	    log.Errorf("Error in ConvertToDataDefinition for SddcsClientImpl.delete method's errors.InvalidRequest error - %s",
	        bindings.VAPIerrorsToError(errError2).Error())
        return nil
	  }
	  errorDefinitions = append(errorDefinitions, errDef2.(data.ErrorDefinition))
      sIface.errorBindingMap[errors.Unauthorized{}.Error()] = errors.UnauthorizedBindingType()
	  errDef3, errError3 := typeConverter.ConvertToDataDefinition(errors.UnauthorizedBindingType())
	  if(errError3 != nil) {
	    log.Errorf("Error in ConvertToDataDefinition for SddcsClientImpl.delete method's errors.Unauthorized error - %s",
	        bindings.VAPIerrorsToError(errError3).Error())
        return nil
	  }
	  errorDefinitions = append(errorDefinitions, errDef3.(data.ErrorDefinition))
      sIface.errorBindingMap[errors.NotFound{}.Error()] = errors.NotFoundBindingType()
	  errDef4, errError4 := typeConverter.ConvertToDataDefinition(errors.NotFoundBindingType())
	  if(errError4 != nil) {
	    log.Errorf("Error in ConvertToDataDefinition for SddcsClientImpl.delete method's errors.NotFound error - %s",
	        bindings.VAPIerrorsToError(errError4).Error())
        return nil
	  }
	  errorDefinitions = append(errorDefinitions, errDef4.(data.ErrorDefinition))

      methodDefinition := core.NewMethodDefinition(methodIdentifier, input, output, errorDefinitions)
      return &methodDefinition
}
func (sIface *SddcsClientImpl) getMethodDefinition() *core.MethodDefinition {
      interfaceIdentifier := core.NewInterfaceIdentifier(sIface.interfaceName)
      typeConverter := sIface.connector.TypeConverter()

      input, inputError := typeConverter.ConvertToDataDefinition(sddcsGetInputType())
      output, outputError := typeConverter.ConvertToDataDefinition(sddcsGetOutputType())
      if(inputError != nil) {
          log.Errorf("Error in ConvertToDataDefinition for SddcsClientImpl.get method's input - %s",
              bindings.VAPIerrorsToError(inputError).Error())
          return nil
      }
      if(outputError != nil) {
          log.Errorf("Error in ConvertToDataDefinition for SddcsClientImpl.get method's output - %s",
              bindings.VAPIerrorsToError(outputError).Error())
          return nil
      }
      methodIdentifier := core.NewMethodIdentifier(interfaceIdentifier, "get")
      errorDefinitions := make([]data.ErrorDefinition,0)
      sIface.errorBindingMap[errors.Unauthenticated{}.Error()] = errors.UnauthenticatedBindingType()
	  errDef1, errError1 := typeConverter.ConvertToDataDefinition(errors.UnauthenticatedBindingType())
	  if(errError1 != nil) {
	    log.Errorf("Error in ConvertToDataDefinition for SddcsClientImpl.get method's errors.Unauthenticated error - %s",
	        bindings.VAPIerrorsToError(errError1).Error())
        return nil
	  }
	  errorDefinitions = append(errorDefinitions, errDef1.(data.ErrorDefinition))
      sIface.errorBindingMap[errors.Unauthorized{}.Error()] = errors.UnauthorizedBindingType()
	  errDef2, errError2 := typeConverter.ConvertToDataDefinition(errors.UnauthorizedBindingType())
	  if(errError2 != nil) {
	    log.Errorf("Error in ConvertToDataDefinition for SddcsClientImpl.get method's errors.Unauthorized error - %s",
	        bindings.VAPIerrorsToError(errError2).Error())
        return nil
	  }
	  errorDefinitions = append(errorDefinitions, errDef2.(data.ErrorDefinition))
      sIface.errorBindingMap[errors.NotFound{}.Error()] = errors.NotFoundBindingType()
	  errDef3, errError3 := typeConverter.ConvertToDataDefinition(errors.NotFoundBindingType())
	  if(errError3 != nil) {
	    log.Errorf("Error in ConvertToDataDefinition for SddcsClientImpl.get method's errors.NotFound error - %s",
	        bindings.VAPIerrorsToError(errError3).Error())
        return nil
	  }
	  errorDefinitions = append(errorDefinitions, errDef3.(data.ErrorDefinition))

      methodDefinition := core.NewMethodDefinition(methodIdentifier, input, output, errorDefinitions)
      return &methodDefinition
}
func (sIface *SddcsClientImpl) listMethodDefinition() *core.MethodDefinition {
      interfaceIdentifier := core.NewInterfaceIdentifier(sIface.interfaceName)
      typeConverter := sIface.connector.TypeConverter()

      input, inputError := typeConverter.ConvertToDataDefinition(sddcsListInputType())
      output, outputError := typeConverter.ConvertToDataDefinition(sddcsListOutputType())
      if(inputError != nil) {
          log.Errorf("Error in ConvertToDataDefinition for SddcsClientImpl.list method's input - %s",
              bindings.VAPIerrorsToError(inputError).Error())
          return nil
      }
      if(outputError != nil) {
          log.Errorf("Error in ConvertToDataDefinition for SddcsClientImpl.list method's output - %s",
              bindings.VAPIerrorsToError(outputError).Error())
          return nil
      }
      methodIdentifier := core.NewMethodIdentifier(interfaceIdentifier, "list")
      errorDefinitions := make([]data.ErrorDefinition,0)
      sIface.errorBindingMap[errors.Unauthenticated{}.Error()] = errors.UnauthenticatedBindingType()
	  errDef1, errError1 := typeConverter.ConvertToDataDefinition(errors.UnauthenticatedBindingType())
	  if(errError1 != nil) {
	    log.Errorf("Error in ConvertToDataDefinition for SddcsClientImpl.list method's errors.Unauthenticated error - %s",
	        bindings.VAPIerrorsToError(errError1).Error())
        return nil
	  }
	  errorDefinitions = append(errorDefinitions, errDef1.(data.ErrorDefinition))
      sIface.errorBindingMap[errors.Unauthorized{}.Error()] = errors.UnauthorizedBindingType()
	  errDef2, errError2 := typeConverter.ConvertToDataDefinition(errors.UnauthorizedBindingType())
	  if(errError2 != nil) {
	    log.Errorf("Error in ConvertToDataDefinition for SddcsClientImpl.list method's errors.Unauthorized error - %s",
	        bindings.VAPIerrorsToError(errError2).Error())
        return nil
	  }
	  errorDefinitions = append(errorDefinitions, errDef2.(data.ErrorDefinition))

      methodDefinition := core.NewMethodDefinition(methodIdentifier, input, output, errorDefinitions)
      return &methodDefinition
}
func (sIface *SddcsClientImpl) patchMethodDefinition() *core.MethodDefinition {
      interfaceIdentifier := core.NewInterfaceIdentifier(sIface.interfaceName)
      typeConverter := sIface.connector.TypeConverter()

      input, inputError := typeConverter.ConvertToDataDefinition(sddcsPatchInputType())
      output, outputError := typeConverter.ConvertToDataDefinition(sddcsPatchOutputType())
      if(inputError != nil) {
          log.Errorf("Error in ConvertToDataDefinition for SddcsClientImpl.patch method's input - %s",
              bindings.VAPIerrorsToError(inputError).Error())
          return nil
      }
      if(outputError != nil) {
          log.Errorf("Error in ConvertToDataDefinition for SddcsClientImpl.patch method's output - %s",
              bindings.VAPIerrorsToError(outputError).Error())
          return nil
      }
      methodIdentifier := core.NewMethodIdentifier(interfaceIdentifier, "patch")
      errorDefinitions := make([]data.ErrorDefinition,0)
      sIface.errorBindingMap[errors.Unauthenticated{}.Error()] = errors.UnauthenticatedBindingType()
	  errDef1, errError1 := typeConverter.ConvertToDataDefinition(errors.UnauthenticatedBindingType())
	  if(errError1 != nil) {
	    log.Errorf("Error in ConvertToDataDefinition for SddcsClientImpl.patch method's errors.Unauthenticated error - %s",
	        bindings.VAPIerrorsToError(errError1).Error())
        return nil
	  }
	  errorDefinitions = append(errorDefinitions, errDef1.(data.ErrorDefinition))
      sIface.errorBindingMap[errors.InvalidRequest{}.Error()] = errors.InvalidRequestBindingType()
	  errDef2, errError2 := typeConverter.ConvertToDataDefinition(errors.InvalidRequestBindingType())
	  if(errError2 != nil) {
	    log.Errorf("Error in ConvertToDataDefinition for SddcsClientImpl.patch method's errors.InvalidRequest error - %s",
	        bindings.VAPIerrorsToError(errError2).Error())
        return nil
	  }
	  errorDefinitions = append(errorDefinitions, errDef2.(data.ErrorDefinition))
      sIface.errorBindingMap[errors.Unauthorized{}.Error()] = errors.UnauthorizedBindingType()
	  errDef3, errError3 := typeConverter.ConvertToDataDefinition(errors.UnauthorizedBindingType())
	  if(errError3 != nil) {
	    log.Errorf("Error in ConvertToDataDefinition for SddcsClientImpl.patch method's errors.Unauthorized error - %s",
	        bindings.VAPIerrorsToError(errError3).Error())
        return nil
	  }
	  errorDefinitions = append(errorDefinitions, errDef3.(data.ErrorDefinition))
      sIface.errorBindingMap[errors.NotFound{}.Error()] = errors.NotFoundBindingType()
	  errDef4, errError4 := typeConverter.ConvertToDataDefinition(errors.NotFoundBindingType())
	  if(errError4 != nil) {
	    log.Errorf("Error in ConvertToDataDefinition for SddcsClientImpl.patch method's errors.NotFound error - %s",
	        bindings.VAPIerrorsToError(errError4).Error())
        return nil
	  }
	  errorDefinitions = append(errorDefinitions, errDef4.(data.ErrorDefinition))

      methodDefinition := core.NewMethodDefinition(methodIdentifier, input, output, errorDefinitions)
      return &methodDefinition
}







func sddcsCreateInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["org"] = bindings.NewStringType()
    fields["sddc_config"] = bindings.NewReferenceType(model.AwsSddcConfigBindingType)
    fieldNameMap["org"] = "Org"
    fieldNameMap["sddc_config"] = "SddcConfig"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func sddcsCreateOutputType() bindings.BindingType {
    return bindings.NewReferenceType(model.TaskBindingType)
}

func sddcsCreateRestMetadata() protocol.OperationRestMetadata {
    paramsTypeMap := map[string]bindings.BindingType{}
    pathParams := map[string]string{}
    queryParams := map[string]string{}
    headerParams := map[string]string{}
    paramsTypeMap["org"] = bindings.NewStringType()
    paramsTypeMap["sddc_config"] = bindings.NewReferenceType(model.AwsSddcConfigBindingType)
    paramsTypeMap["org"] = bindings.NewStringType()
    pathParams["org"] = "org"
    resultHeaders := map[string]string{}
    errorHeaders := map[string]string{}
    errorHeaders["Unauthenticated.challenge"] = "WWW-Authenticate"
    return protocol.NewOperationRestMetadata(
      paramsTypeMap,
      pathParams,
      queryParams,
      headerParams,
      "sddc_config",
      "POST",
      "/vmc/api/orgs/{org}/sddcs",
       resultHeaders,
       201,
       errorHeaders,
       map[string]int{"Unauthenticated": 401,"Unauthorized": 403})
}


func sddcsDeleteInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["org"] = bindings.NewStringType()
    fields["sddc"] = bindings.NewStringType()
    fields["retain_configuration"] = bindings.NewOptionalType(bindings.NewBooleanType())
    fields["template_name"] = bindings.NewOptionalType(bindings.NewStringType())
    fields["force"] = bindings.NewOptionalType(bindings.NewStringType())
    fieldNameMap["org"] = "Org"
    fieldNameMap["sddc"] = "Sddc"
    fieldNameMap["retain_configuration"] = "RetainConfiguration"
    fieldNameMap["template_name"] = "TemplateName"
    fieldNameMap["force"] = "Force"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func sddcsDeleteOutputType() bindings.BindingType {
    return bindings.NewReferenceType(model.TaskBindingType)
}

func sddcsDeleteRestMetadata() protocol.OperationRestMetadata {
    paramsTypeMap := map[string]bindings.BindingType{}
    pathParams := map[string]string{}
    queryParams := map[string]string{}
    headerParams := map[string]string{}
    paramsTypeMap["force"] = bindings.NewOptionalType(bindings.NewStringType())
    paramsTypeMap["org"] = bindings.NewStringType()
    paramsTypeMap["sddc"] = bindings.NewStringType()
    paramsTypeMap["template_name"] = bindings.NewOptionalType(bindings.NewStringType())
    paramsTypeMap["retain_configuration"] = bindings.NewOptionalType(bindings.NewBooleanType())
    paramsTypeMap["org"] = bindings.NewStringType()
    paramsTypeMap["sddc"] = bindings.NewStringType()
    pathParams["org"] = "org"
    pathParams["sddc"] = "sddc"
    queryParams["retain_configuration"] = "retain_configuration"
    queryParams["template_name"] = "template_name"
    queryParams["force"] = "force"
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
      "/vmc/api/orgs/{org}/sddcs/{sddc}",
       resultHeaders,
       200,
       errorHeaders,
       map[string]int{"Unauthenticated": 401,"InvalidRequest": 400,"Unauthorized": 403,"NotFound": 404})
}


func sddcsGetInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["org"] = bindings.NewStringType()
    fields["sddc"] = bindings.NewStringType()
    fieldNameMap["org"] = "Org"
    fieldNameMap["sddc"] = "Sddc"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func sddcsGetOutputType() bindings.BindingType {
    return bindings.NewReferenceType(model.SddcBindingType)
}

func sddcsGetRestMetadata() protocol.OperationRestMetadata {
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
      "/vmc/api/orgs/{org}/sddcs/{sddc}",
       resultHeaders,
       200,
       errorHeaders,
       map[string]int{"Unauthenticated": 401,"Unauthorized": 403,"NotFound": 404})
}


func sddcsListInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["org"] = bindings.NewStringType()
    fieldNameMap["org"] = "Org"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func sddcsListOutputType() bindings.BindingType {
    return bindings.NewListType(bindings.NewReferenceType(model.SddcBindingType), reflect.TypeOf([]model.Sddc{}))
}

func sddcsListRestMetadata() protocol.OperationRestMetadata {
    paramsTypeMap := map[string]bindings.BindingType{}
    pathParams := map[string]string{}
    queryParams := map[string]string{}
    headerParams := map[string]string{}
    paramsTypeMap["org"] = bindings.NewStringType()
    paramsTypeMap["org"] = bindings.NewStringType()
    pathParams["org"] = "org"
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
      "/vmc/api/orgs/{org}/sddcs",
       resultHeaders,
       200,
       errorHeaders,
       map[string]int{"Unauthenticated": 401,"Unauthorized": 403})
}


func sddcsPatchInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["org"] = bindings.NewStringType()
    fields["sddc"] = bindings.NewStringType()
    fields["sddc_patch_request"] = bindings.NewReferenceType(model.SddcPatchRequestBindingType)
    fieldNameMap["org"] = "Org"
    fieldNameMap["sddc"] = "Sddc"
    fieldNameMap["sddc_patch_request"] = "SddcPatchRequest"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func sddcsPatchOutputType() bindings.BindingType {
    return bindings.NewReferenceType(model.SddcBindingType)
}

func sddcsPatchRestMetadata() protocol.OperationRestMetadata {
    paramsTypeMap := map[string]bindings.BindingType{}
    pathParams := map[string]string{}
    queryParams := map[string]string{}
    headerParams := map[string]string{}
    paramsTypeMap["org"] = bindings.NewStringType()
    paramsTypeMap["sddc"] = bindings.NewStringType()
    paramsTypeMap["sddc_patch_request"] = bindings.NewReferenceType(model.SddcPatchRequestBindingType)
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
      "sddc_patch_request",
      "PATCH",
      "/vmc/api/orgs/{org}/sddcs/{sddc}",
       resultHeaders,
       200,
       errorHeaders,
       map[string]int{"Unauthenticated": 401,"InvalidRequest": 400,"Unauthorized": 403,"NotFound": 404})
}




