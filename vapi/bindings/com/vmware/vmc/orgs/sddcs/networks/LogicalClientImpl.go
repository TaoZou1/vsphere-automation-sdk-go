
/* Copyright © 2019 VMware, Inc. All Rights Reserved.
     SPDX-License-Identifier: BSD-2-Clause */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Client stubs for service: Logical
 * Functions that implement the generated LogicalClient interface
 */


package networks

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


type LogicalClientImpl struct{
      interfaceName string
      interfaceDefinition core.InterfaceDefinition
      methodIdentifiers[]core.MethodIdentifier
      methodNameToDefMap map[string]*core.MethodDefinition
      errorBindingMap map[string]bindings.BindingType
      interfaceIdentifier core.InterfaceIdentifier
      connector client.Connector
}


func NewLogicalClientImpl(connector client.Connector) *LogicalClientImpl {
      interfaceName := "com.vmware.vmc.orgs.sddcs.networks.logical"
      interfaceIdentifier := core.NewInterfaceIdentifier(interfaceName)
      methodIdentifiers := []core.MethodIdentifier{
          core.NewMethodIdentifier(interfaceIdentifier, "create"),
          core.NewMethodIdentifier(interfaceIdentifier, "delete"),
          core.NewMethodIdentifier(interfaceIdentifier, "get"),
          core.NewMethodIdentifier(interfaceIdentifier, "get_0"),
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
      lIface := LogicalClientImpl{interfaceName: interfaceName, methodIdentifiers: methodIdentifiers, interfaceDefinition: interfaceDefinition, errorBindingMap: errorBindingMap, interfaceIdentifier:interfaceIdentifier, connector: connector}
      lIface.methodNameToDefMap = make(map[string]*core.MethodDefinition)
      lIface.methodNameToDefMap["create"] = lIface.createMethodDefinition()
      lIface.methodNameToDefMap["delete"] = lIface.deleteMethodDefinition()
      lIface.methodNameToDefMap["get"] = lIface.getMethodDefinition()
      lIface.methodNameToDefMap["get_0"] = lIface.get_0MethodDefinition()
      lIface.methodNameToDefMap["update"] = lIface.updateMethodDefinition()
      return &lIface
}

func (lIface *LogicalClientImpl) Create(orgParam string, sddcParam string, sddcNetworkParam model.SddcNetwork) error {
	typeConverter := lIface.connector.TypeConverter()
	methodIdentifier := core.NewMethodIdentifier(lIface.interfaceIdentifier, "create")
	sv := bindings.NewStructValueBuilder(logicalCreateInputType(), typeConverter)
	sv.AddStructField("Org", orgParam)
	sv.AddStructField("Sddc", sddcParam)
	sv.AddStructField("SddcNetwork", sddcNetworkParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		return bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := logicalCreateRestMetadata
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	lIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult:= lIface.Invoke(lIface.connector.NewExecutionContext(), methodIdentifier, inputDataValue)
    if methodResult.IsSuccess() {
		return nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), lIface.errorBindingMap[methodResult.Error().Name()])
		if errorInError != nil {
		    return bindings.VAPIerrorsToError(errorInError)
		}
		return methodError.(error)
	}
}
func (lIface *LogicalClientImpl) Delete(orgParam string, sddcParam string, networkIdParam string) error {
	typeConverter := lIface.connector.TypeConverter()
	methodIdentifier := core.NewMethodIdentifier(lIface.interfaceIdentifier, "delete")
	sv := bindings.NewStructValueBuilder(logicalDeleteInputType(), typeConverter)
	sv.AddStructField("Org", orgParam)
	sv.AddStructField("Sddc", sddcParam)
	sv.AddStructField("NetworkId", networkIdParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		return bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := logicalDeleteRestMetadata
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	lIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult:= lIface.Invoke(lIface.connector.NewExecutionContext(), methodIdentifier, inputDataValue)
    if methodResult.IsSuccess() {
		return nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), lIface.errorBindingMap[methodResult.Error().Name()])
		if errorInError != nil {
		    return bindings.VAPIerrorsToError(errorInError)
		}
		return methodError.(error)
	}
}
func (lIface *LogicalClientImpl) Get(orgParam string, sddcParam string, networkIdParam string) (model.SddcNetwork, error) {
	typeConverter := lIface.connector.TypeConverter()
	methodIdentifier := core.NewMethodIdentifier(lIface.interfaceIdentifier, "get")
	sv := bindings.NewStructValueBuilder(logicalGetInputType(), typeConverter)
	sv.AddStructField("Org", orgParam)
	sv.AddStructField("Sddc", sddcParam)
	sv.AddStructField("NetworkId", networkIdParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
        var emptyOutput model.SddcNetwork
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := logicalGetRestMetadata
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	lIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult:= lIface.Invoke(lIface.connector.NewExecutionContext(), methodIdentifier, inputDataValue)
	var emptyOutput model.SddcNetwork
    if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), logicalGetOutputType())
		if errorInOutput != nil {
		    return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
	    return output.(model.SddcNetwork), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), lIface.errorBindingMap[methodResult.Error().Name()])
		if errorInError != nil {
		    return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}
func (lIface *LogicalClientImpl) Get0(orgParam string, sddcParam string, pageSizeParam *int64, startIndexParam *int64, prevSddcNetworkIdParam *string, sortOrderAscendingParam *bool) (model.DataPageSddcNetwork, error) {
	typeConverter := lIface.connector.TypeConverter()
	methodIdentifier := core.NewMethodIdentifier(lIface.interfaceIdentifier, "get_0")
	sv := bindings.NewStructValueBuilder(logicalGet0InputType(), typeConverter)
	sv.AddStructField("Org", orgParam)
	sv.AddStructField("Sddc", sddcParam)
	sv.AddStructField("PageSize", pageSizeParam)
	sv.AddStructField("StartIndex", startIndexParam)
	sv.AddStructField("PrevSddcNetworkId", prevSddcNetworkIdParam)
	sv.AddStructField("SortOrderAscending", sortOrderAscendingParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
        var emptyOutput model.DataPageSddcNetwork
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := logicalGet0RestMetadata
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	lIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult:= lIface.Invoke(lIface.connector.NewExecutionContext(), methodIdentifier, inputDataValue)
	var emptyOutput model.DataPageSddcNetwork
    if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), logicalGet0OutputType())
		if errorInOutput != nil {
		    return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
	    return output.(model.DataPageSddcNetwork), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), lIface.errorBindingMap[methodResult.Error().Name()])
		if errorInError != nil {
		    return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}
func (lIface *LogicalClientImpl) Update(orgParam string, sddcParam string, networkIdParam string, sddcNetworkParam model.SddcNetwork) error {
	typeConverter := lIface.connector.TypeConverter()
	methodIdentifier := core.NewMethodIdentifier(lIface.interfaceIdentifier, "update")
	sv := bindings.NewStructValueBuilder(logicalUpdateInputType(), typeConverter)
	sv.AddStructField("Org", orgParam)
	sv.AddStructField("Sddc", sddcParam)
	sv.AddStructField("NetworkId", networkIdParam)
	sv.AddStructField("SddcNetwork", sddcNetworkParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		return bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := logicalUpdateRestMetadata
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	lIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult:= lIface.Invoke(lIface.connector.NewExecutionContext(), methodIdentifier, inputDataValue)
    if methodResult.IsSuccess() {
		return nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), lIface.errorBindingMap[methodResult.Error().Name()])
		if errorInError != nil {
		    return bindings.VAPIerrorsToError(errorInError)
		}
		return methodError.(error)
	}
}

func (lIface *LogicalClientImpl) Invoke(ctx *core.ExecutionContext, methodId core.MethodIdentifier, inputDataValue data.DataValue) core.MethodResult {
    methodResult := lIface.connector.GetApiProvider().Invoke(lIface.interfaceName, methodId.Name(), inputDataValue, ctx)
    return methodResult
}

func (lIface *LogicalClientImpl) createMethodDefinition() *core.MethodDefinition {
      interfaceIdentifier := core.NewInterfaceIdentifier(lIface.interfaceName)
      typeConverter := lIface.connector.TypeConverter()

      input, inputError := typeConverter.ConvertToDataDefinition(logicalCreateInputType())
      output, outputError := typeConverter.ConvertToDataDefinition(logicalCreateOutputType())
      if(inputError != nil) {
          log.Errorf("Error in ConvertToDataDefinition for LogicalClientImpl.create method's input - %s",
              bindings.VAPIerrorsToError(inputError).Error())
          return nil
      }
      if(outputError != nil) {
          log.Errorf("Error in ConvertToDataDefinition for LogicalClientImpl.create method's output - %s",
              bindings.VAPIerrorsToError(outputError).Error())
          return nil
      }
      methodIdentifier := core.NewMethodIdentifier(interfaceIdentifier, "create")
      errorDefinitions := make([]data.ErrorDefinition,0)
      lIface.errorBindingMap[errors.InvalidRequest{}.Error()] = errors.InvalidRequestBindingType()
	  errDef1, errError1 := typeConverter.ConvertToDataDefinition(errors.InvalidRequestBindingType())
	  if(errError1 != nil) {
	    log.Errorf("Error in ConvertToDataDefinition for LogicalClientImpl.create method's errors.InvalidRequest error - %s",
	        bindings.VAPIerrorsToError(errError1).Error())
        return nil
	  }
	  errorDefinitions = append(errorDefinitions, errDef1.(data.ErrorDefinition))
      lIface.errorBindingMap[errors.Unauthorized{}.Error()] = errors.UnauthorizedBindingType()
	  errDef2, errError2 := typeConverter.ConvertToDataDefinition(errors.UnauthorizedBindingType())
	  if(errError2 != nil) {
	    log.Errorf("Error in ConvertToDataDefinition for LogicalClientImpl.create method's errors.Unauthorized error - %s",
	        bindings.VAPIerrorsToError(errError2).Error())
        return nil
	  }
	  errorDefinitions = append(errorDefinitions, errDef2.(data.ErrorDefinition))
      lIface.errorBindingMap[errors.NotFound{}.Error()] = errors.NotFoundBindingType()
	  errDef3, errError3 := typeConverter.ConvertToDataDefinition(errors.NotFoundBindingType())
	  if(errError3 != nil) {
	    log.Errorf("Error in ConvertToDataDefinition for LogicalClientImpl.create method's errors.NotFound error - %s",
	        bindings.VAPIerrorsToError(errError3).Error())
        return nil
	  }
	  errorDefinitions = append(errorDefinitions, errDef3.(data.ErrorDefinition))

      methodDefinition := core.NewMethodDefinition(methodIdentifier, input, output, errorDefinitions)
      return &methodDefinition
}
func (lIface *LogicalClientImpl) deleteMethodDefinition() *core.MethodDefinition {
      interfaceIdentifier := core.NewInterfaceIdentifier(lIface.interfaceName)
      typeConverter := lIface.connector.TypeConverter()

      input, inputError := typeConverter.ConvertToDataDefinition(logicalDeleteInputType())
      output, outputError := typeConverter.ConvertToDataDefinition(logicalDeleteOutputType())
      if(inputError != nil) {
          log.Errorf("Error in ConvertToDataDefinition for LogicalClientImpl.delete method's input - %s",
              bindings.VAPIerrorsToError(inputError).Error())
          return nil
      }
      if(outputError != nil) {
          log.Errorf("Error in ConvertToDataDefinition for LogicalClientImpl.delete method's output - %s",
              bindings.VAPIerrorsToError(outputError).Error())
          return nil
      }
      methodIdentifier := core.NewMethodIdentifier(interfaceIdentifier, "delete")
      errorDefinitions := make([]data.ErrorDefinition,0)
      lIface.errorBindingMap[errors.InvalidRequest{}.Error()] = errors.InvalidRequestBindingType()
	  errDef1, errError1 := typeConverter.ConvertToDataDefinition(errors.InvalidRequestBindingType())
	  if(errError1 != nil) {
	    log.Errorf("Error in ConvertToDataDefinition for LogicalClientImpl.delete method's errors.InvalidRequest error - %s",
	        bindings.VAPIerrorsToError(errError1).Error())
        return nil
	  }
	  errorDefinitions = append(errorDefinitions, errDef1.(data.ErrorDefinition))
      lIface.errorBindingMap[errors.Unauthorized{}.Error()] = errors.UnauthorizedBindingType()
	  errDef2, errError2 := typeConverter.ConvertToDataDefinition(errors.UnauthorizedBindingType())
	  if(errError2 != nil) {
	    log.Errorf("Error in ConvertToDataDefinition for LogicalClientImpl.delete method's errors.Unauthorized error - %s",
	        bindings.VAPIerrorsToError(errError2).Error())
        return nil
	  }
	  errorDefinitions = append(errorDefinitions, errDef2.(data.ErrorDefinition))
      lIface.errorBindingMap[errors.NotFound{}.Error()] = errors.NotFoundBindingType()
	  errDef3, errError3 := typeConverter.ConvertToDataDefinition(errors.NotFoundBindingType())
	  if(errError3 != nil) {
	    log.Errorf("Error in ConvertToDataDefinition for LogicalClientImpl.delete method's errors.NotFound error - %s",
	        bindings.VAPIerrorsToError(errError3).Error())
        return nil
	  }
	  errorDefinitions = append(errorDefinitions, errDef3.(data.ErrorDefinition))

      methodDefinition := core.NewMethodDefinition(methodIdentifier, input, output, errorDefinitions)
      return &methodDefinition
}
func (lIface *LogicalClientImpl) getMethodDefinition() *core.MethodDefinition {
      interfaceIdentifier := core.NewInterfaceIdentifier(lIface.interfaceName)
      typeConverter := lIface.connector.TypeConverter()

      input, inputError := typeConverter.ConvertToDataDefinition(logicalGetInputType())
      output, outputError := typeConverter.ConvertToDataDefinition(logicalGetOutputType())
      if(inputError != nil) {
          log.Errorf("Error in ConvertToDataDefinition for LogicalClientImpl.get method's input - %s",
              bindings.VAPIerrorsToError(inputError).Error())
          return nil
      }
      if(outputError != nil) {
          log.Errorf("Error in ConvertToDataDefinition for LogicalClientImpl.get method's output - %s",
              bindings.VAPIerrorsToError(outputError).Error())
          return nil
      }
      methodIdentifier := core.NewMethodIdentifier(interfaceIdentifier, "get")
      errorDefinitions := make([]data.ErrorDefinition,0)
      lIface.errorBindingMap[errors.InvalidRequest{}.Error()] = errors.InvalidRequestBindingType()
	  errDef1, errError1 := typeConverter.ConvertToDataDefinition(errors.InvalidRequestBindingType())
	  if(errError1 != nil) {
	    log.Errorf("Error in ConvertToDataDefinition for LogicalClientImpl.get method's errors.InvalidRequest error - %s",
	        bindings.VAPIerrorsToError(errError1).Error())
        return nil
	  }
	  errorDefinitions = append(errorDefinitions, errDef1.(data.ErrorDefinition))
      lIface.errorBindingMap[errors.Unauthorized{}.Error()] = errors.UnauthorizedBindingType()
	  errDef2, errError2 := typeConverter.ConvertToDataDefinition(errors.UnauthorizedBindingType())
	  if(errError2 != nil) {
	    log.Errorf("Error in ConvertToDataDefinition for LogicalClientImpl.get method's errors.Unauthorized error - %s",
	        bindings.VAPIerrorsToError(errError2).Error())
        return nil
	  }
	  errorDefinitions = append(errorDefinitions, errDef2.(data.ErrorDefinition))
      lIface.errorBindingMap[errors.NotFound{}.Error()] = errors.NotFoundBindingType()
	  errDef3, errError3 := typeConverter.ConvertToDataDefinition(errors.NotFoundBindingType())
	  if(errError3 != nil) {
	    log.Errorf("Error in ConvertToDataDefinition for LogicalClientImpl.get method's errors.NotFound error - %s",
	        bindings.VAPIerrorsToError(errError3).Error())
        return nil
	  }
	  errorDefinitions = append(errorDefinitions, errDef3.(data.ErrorDefinition))

      methodDefinition := core.NewMethodDefinition(methodIdentifier, input, output, errorDefinitions)
      return &methodDefinition
}
func (lIface *LogicalClientImpl) get_0MethodDefinition() *core.MethodDefinition {
      interfaceIdentifier := core.NewInterfaceIdentifier(lIface.interfaceName)
      typeConverter := lIface.connector.TypeConverter()

      input, inputError := typeConverter.ConvertToDataDefinition(logicalGet0InputType())
      output, outputError := typeConverter.ConvertToDataDefinition(logicalGet0OutputType())
      if(inputError != nil) {
          log.Errorf("Error in ConvertToDataDefinition for LogicalClientImpl.get_0 method's input - %s",
              bindings.VAPIerrorsToError(inputError).Error())
          return nil
      }
      if(outputError != nil) {
          log.Errorf("Error in ConvertToDataDefinition for LogicalClientImpl.get_0 method's output - %s",
              bindings.VAPIerrorsToError(outputError).Error())
          return nil
      }
      methodIdentifier := core.NewMethodIdentifier(interfaceIdentifier, "get_0")
      errorDefinitions := make([]data.ErrorDefinition,0)
      lIface.errorBindingMap[errors.InvalidRequest{}.Error()] = errors.InvalidRequestBindingType()
	  errDef1, errError1 := typeConverter.ConvertToDataDefinition(errors.InvalidRequestBindingType())
	  if(errError1 != nil) {
	    log.Errorf("Error in ConvertToDataDefinition for LogicalClientImpl.get_0 method's errors.InvalidRequest error - %s",
	        bindings.VAPIerrorsToError(errError1).Error())
        return nil
	  }
	  errorDefinitions = append(errorDefinitions, errDef1.(data.ErrorDefinition))
      lIface.errorBindingMap[errors.Unauthorized{}.Error()] = errors.UnauthorizedBindingType()
	  errDef2, errError2 := typeConverter.ConvertToDataDefinition(errors.UnauthorizedBindingType())
	  if(errError2 != nil) {
	    log.Errorf("Error in ConvertToDataDefinition for LogicalClientImpl.get_0 method's errors.Unauthorized error - %s",
	        bindings.VAPIerrorsToError(errError2).Error())
        return nil
	  }
	  errorDefinitions = append(errorDefinitions, errDef2.(data.ErrorDefinition))
      lIface.errorBindingMap[errors.NotFound{}.Error()] = errors.NotFoundBindingType()
	  errDef3, errError3 := typeConverter.ConvertToDataDefinition(errors.NotFoundBindingType())
	  if(errError3 != nil) {
	    log.Errorf("Error in ConvertToDataDefinition for LogicalClientImpl.get_0 method's errors.NotFound error - %s",
	        bindings.VAPIerrorsToError(errError3).Error())
        return nil
	  }
	  errorDefinitions = append(errorDefinitions, errDef3.(data.ErrorDefinition))

      methodDefinition := core.NewMethodDefinition(methodIdentifier, input, output, errorDefinitions)
      return &methodDefinition
}
func (lIface *LogicalClientImpl) updateMethodDefinition() *core.MethodDefinition {
      interfaceIdentifier := core.NewInterfaceIdentifier(lIface.interfaceName)
      typeConverter := lIface.connector.TypeConverter()

      input, inputError := typeConverter.ConvertToDataDefinition(logicalUpdateInputType())
      output, outputError := typeConverter.ConvertToDataDefinition(logicalUpdateOutputType())
      if(inputError != nil) {
          log.Errorf("Error in ConvertToDataDefinition for LogicalClientImpl.update method's input - %s",
              bindings.VAPIerrorsToError(inputError).Error())
          return nil
      }
      if(outputError != nil) {
          log.Errorf("Error in ConvertToDataDefinition for LogicalClientImpl.update method's output - %s",
              bindings.VAPIerrorsToError(outputError).Error())
          return nil
      }
      methodIdentifier := core.NewMethodIdentifier(interfaceIdentifier, "update")
      errorDefinitions := make([]data.ErrorDefinition,0)
      lIface.errorBindingMap[errors.InvalidRequest{}.Error()] = errors.InvalidRequestBindingType()
	  errDef1, errError1 := typeConverter.ConvertToDataDefinition(errors.InvalidRequestBindingType())
	  if(errError1 != nil) {
	    log.Errorf("Error in ConvertToDataDefinition for LogicalClientImpl.update method's errors.InvalidRequest error - %s",
	        bindings.VAPIerrorsToError(errError1).Error())
        return nil
	  }
	  errorDefinitions = append(errorDefinitions, errDef1.(data.ErrorDefinition))
      lIface.errorBindingMap[errors.Unauthorized{}.Error()] = errors.UnauthorizedBindingType()
	  errDef2, errError2 := typeConverter.ConvertToDataDefinition(errors.UnauthorizedBindingType())
	  if(errError2 != nil) {
	    log.Errorf("Error in ConvertToDataDefinition for LogicalClientImpl.update method's errors.Unauthorized error - %s",
	        bindings.VAPIerrorsToError(errError2).Error())
        return nil
	  }
	  errorDefinitions = append(errorDefinitions, errDef2.(data.ErrorDefinition))
      lIface.errorBindingMap[errors.NotFound{}.Error()] = errors.NotFoundBindingType()
	  errDef3, errError3 := typeConverter.ConvertToDataDefinition(errors.NotFoundBindingType())
	  if(errError3 != nil) {
	    log.Errorf("Error in ConvertToDataDefinition for LogicalClientImpl.update method's errors.NotFound error - %s",
	        bindings.VAPIerrorsToError(errError3).Error())
        return nil
	  }
	  errorDefinitions = append(errorDefinitions, errDef3.(data.ErrorDefinition))

      methodDefinition := core.NewMethodDefinition(methodIdentifier, input, output, errorDefinitions)
      return &methodDefinition
}







func logicalCreateInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["org"] = bindings.NewStringType()
    fields["sddc"] = bindings.NewStringType()
    fields["sddc_network"] = bindings.NewReferenceType(model.SddcNetworkBindingType)
    fieldNameMap["org"] = "Org"
    fieldNameMap["sddc"] = "Sddc"
    fieldNameMap["sddc_network"] = "SddcNetwork"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func logicalCreateOutputType() bindings.BindingType {
    return bindings.NewVoidType()
}

func logicalCreateRestMetadata() protocol.OperationRestMetadata {
    paramsTypeMap := map[string]bindings.BindingType{}
    pathParams := map[string]string{}
    queryParams := map[string]string{}
    headerParams := map[string]string{}
    paramsTypeMap["org"] = bindings.NewStringType()
    paramsTypeMap["sddc"] = bindings.NewStringType()
    paramsTypeMap["sddc_network"] = bindings.NewReferenceType(model.SddcNetworkBindingType)
    paramsTypeMap["org"] = bindings.NewStringType()
    paramsTypeMap["sddc"] = bindings.NewStringType()
    pathParams["org"] = "org"
    pathParams["sddc"] = "sddc"
    resultHeaders := map[string]string{}
    errorHeaders := map[string]string{}
    return protocol.NewOperationRestMetadata(
      paramsTypeMap,
      pathParams,
      queryParams,
      headerParams,
      "sddc_network",
      "POST",
      "/vmc/api/orgs/{org}/sddcs/{sddc}/networks/4.0/sddc/networks",
       resultHeaders,
       204,
       errorHeaders,
       map[string]int{"InvalidRequest": 400,"Unauthorized": 403,"NotFound": 404})
}


func logicalDeleteInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["org"] = bindings.NewStringType()
    fields["sddc"] = bindings.NewStringType()
    fields["network_id"] = bindings.NewStringType()
    fieldNameMap["org"] = "Org"
    fieldNameMap["sddc"] = "Sddc"
    fieldNameMap["network_id"] = "NetworkId"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func logicalDeleteOutputType() bindings.BindingType {
    return bindings.NewVoidType()
}

func logicalDeleteRestMetadata() protocol.OperationRestMetadata {
    paramsTypeMap := map[string]bindings.BindingType{}
    pathParams := map[string]string{}
    queryParams := map[string]string{}
    headerParams := map[string]string{}
    paramsTypeMap["network_id"] = bindings.NewStringType()
    paramsTypeMap["org"] = bindings.NewStringType()
    paramsTypeMap["sddc"] = bindings.NewStringType()
    paramsTypeMap["org"] = bindings.NewStringType()
    paramsTypeMap["sddc"] = bindings.NewStringType()
    paramsTypeMap["networkId"] = bindings.NewStringType()
    pathParams["org"] = "org"
    pathParams["sddc"] = "sddc"
    pathParams["network_id"] = "networkId"
    resultHeaders := map[string]string{}
    errorHeaders := map[string]string{}
    return protocol.NewOperationRestMetadata(
      paramsTypeMap,
      pathParams,
      queryParams,
      headerParams,
      "",
      "DELETE",
      "/vmc/api/orgs/{org}/sddcs/{sddc}/networks/4.0/sddc/networks/{networkId}",
       resultHeaders,
       204,
       errorHeaders,
       map[string]int{"InvalidRequest": 400,"Unauthorized": 403,"NotFound": 404})
}


func logicalGetInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["org"] = bindings.NewStringType()
    fields["sddc"] = bindings.NewStringType()
    fields["network_id"] = bindings.NewStringType()
    fieldNameMap["org"] = "Org"
    fieldNameMap["sddc"] = "Sddc"
    fieldNameMap["network_id"] = "NetworkId"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func logicalGetOutputType() bindings.BindingType {
    return bindings.NewReferenceType(model.SddcNetworkBindingType)
}

func logicalGetRestMetadata() protocol.OperationRestMetadata {
    paramsTypeMap := map[string]bindings.BindingType{}
    pathParams := map[string]string{}
    queryParams := map[string]string{}
    headerParams := map[string]string{}
    paramsTypeMap["network_id"] = bindings.NewStringType()
    paramsTypeMap["org"] = bindings.NewStringType()
    paramsTypeMap["sddc"] = bindings.NewStringType()
    paramsTypeMap["org"] = bindings.NewStringType()
    paramsTypeMap["sddc"] = bindings.NewStringType()
    paramsTypeMap["networkId"] = bindings.NewStringType()
    pathParams["org"] = "org"
    pathParams["sddc"] = "sddc"
    pathParams["network_id"] = "networkId"
    resultHeaders := map[string]string{}
    errorHeaders := map[string]string{}
    return protocol.NewOperationRestMetadata(
      paramsTypeMap,
      pathParams,
      queryParams,
      headerParams,
      "",
      "GET",
      "/vmc/api/orgs/{org}/sddcs/{sddc}/networks/4.0/sddc/networks/{networkId}",
       resultHeaders,
       200,
       errorHeaders,
       map[string]int{"InvalidRequest": 400,"Unauthorized": 403,"NotFound": 404})
}


func logicalGet0InputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["org"] = bindings.NewStringType()
    fields["sddc"] = bindings.NewStringType()
    fields["page_size"] = bindings.NewOptionalType(bindings.NewIntegerType())
    fields["start_index"] = bindings.NewOptionalType(bindings.NewIntegerType())
    fields["prev_sddc_network_id"] = bindings.NewOptionalType(bindings.NewStringType())
    fields["sort_order_ascending"] = bindings.NewOptionalType(bindings.NewBooleanType())
    fieldNameMap["org"] = "Org"
    fieldNameMap["sddc"] = "Sddc"
    fieldNameMap["page_size"] = "PageSize"
    fieldNameMap["start_index"] = "StartIndex"
    fieldNameMap["prev_sddc_network_id"] = "PrevSddcNetworkId"
    fieldNameMap["sort_order_ascending"] = "SortOrderAscending"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func logicalGet0OutputType() bindings.BindingType {
    return bindings.NewReferenceType(model.DataPageSddcNetworkBindingType)
}

func logicalGet0RestMetadata() protocol.OperationRestMetadata {
    paramsTypeMap := map[string]bindings.BindingType{}
    pathParams := map[string]string{}
    queryParams := map[string]string{}
    headerParams := map[string]string{}
    paramsTypeMap["org"] = bindings.NewStringType()
    paramsTypeMap["sddc"] = bindings.NewStringType()
    paramsTypeMap["sort_order_ascending"] = bindings.NewOptionalType(bindings.NewBooleanType())
    paramsTypeMap["page_size"] = bindings.NewOptionalType(bindings.NewIntegerType())
    paramsTypeMap["start_index"] = bindings.NewOptionalType(bindings.NewIntegerType())
    paramsTypeMap["prev_sddc_network_id"] = bindings.NewOptionalType(bindings.NewStringType())
    paramsTypeMap["org"] = bindings.NewStringType()
    paramsTypeMap["sddc"] = bindings.NewStringType()
    pathParams["org"] = "org"
    pathParams["sddc"] = "sddc"
    queryParams["start_index"] = "startIndex"
    queryParams["prev_sddc_network_id"] = "prevSddcNetworkId"
    queryParams["page_size"] = "pageSize"
    queryParams["sort_order_ascending"] = "sortOrderAscending"
    resultHeaders := map[string]string{}
    errorHeaders := map[string]string{}
    return protocol.NewOperationRestMetadata(
      paramsTypeMap,
      pathParams,
      queryParams,
      headerParams,
      "",
      "GET",
      "/vmc/api/orgs/{org}/sddcs/{sddc}/networks/4.0/sddc/networks",
       resultHeaders,
       200,
       errorHeaders,
       map[string]int{"InvalidRequest": 400,"Unauthorized": 403,"NotFound": 404})
}


func logicalUpdateInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["org"] = bindings.NewStringType()
    fields["sddc"] = bindings.NewStringType()
    fields["network_id"] = bindings.NewStringType()
    fields["sddc_network"] = bindings.NewReferenceType(model.SddcNetworkBindingType)
    fieldNameMap["org"] = "Org"
    fieldNameMap["sddc"] = "Sddc"
    fieldNameMap["network_id"] = "NetworkId"
    fieldNameMap["sddc_network"] = "SddcNetwork"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func logicalUpdateOutputType() bindings.BindingType {
    return bindings.NewVoidType()
}

func logicalUpdateRestMetadata() protocol.OperationRestMetadata {
    paramsTypeMap := map[string]bindings.BindingType{}
    pathParams := map[string]string{}
    queryParams := map[string]string{}
    headerParams := map[string]string{}
    paramsTypeMap["network_id"] = bindings.NewStringType()
    paramsTypeMap["org"] = bindings.NewStringType()
    paramsTypeMap["sddc"] = bindings.NewStringType()
    paramsTypeMap["sddc_network"] = bindings.NewReferenceType(model.SddcNetworkBindingType)
    paramsTypeMap["org"] = bindings.NewStringType()
    paramsTypeMap["sddc"] = bindings.NewStringType()
    paramsTypeMap["networkId"] = bindings.NewStringType()
    pathParams["org"] = "org"
    pathParams["sddc"] = "sddc"
    pathParams["network_id"] = "networkId"
    resultHeaders := map[string]string{}
    errorHeaders := map[string]string{}
    return protocol.NewOperationRestMetadata(
      paramsTypeMap,
      pathParams,
      queryParams,
      headerParams,
      "sddc_network",
      "PUT",
      "/vmc/api/orgs/{org}/sddcs/{sddc}/networks/4.0/sddc/networks/{networkId}",
       resultHeaders,
       204,
       errorHeaders,
       map[string]int{"InvalidRequest": 400,"Unauthorized": 403,"NotFound": 404})
}




