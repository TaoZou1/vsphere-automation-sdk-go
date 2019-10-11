
/* Copyright © 2019 VMware, Inc. All Rights Reserved.
     SPDX-License-Identifier: BSD-2-Clause */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Client stubs for service: Operation
 * Functions that implement the generated OperationClient interface
 */


package service

import (
    "reflect"
// //     "getClientImplDependenciesOfOps gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/bindings/com/vmware/vapi/metadata/authentication"
// //     "getClientImplDependenciesOfOps gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/bindings/com/vmware/vapi/std/errors"
// //     "getClientImplDependenciesOfOps gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/bindings"
// //     "getClientImplDependenciesOfOps gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/core"
// //     "getClientImplDependenciesOfOps gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/data"
// //     "getClientImplDependenciesOfOps gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/lib"
// //     "getClientImplDependenciesOfOps gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/log"
// //     "getClientImplDependenciesOfOps gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/protocol/client"
// //
// //     "getDependenciesOfServiceTypes gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/bindings/com/vmware/vapi/metadata/authentication"
// //     "getDependenciesOfServiceTypes gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/bindings"
// //     "getDependenciesOfServiceTypes gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/data"
// //     "getDependenciesOfServiceTypes gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/protocol"
// 
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/bindings/com/vmware/vapi/metadata/authentication"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/bindings/com/vmware/vapi/std/errors"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/bindings"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/core"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/data"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/lib"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/log"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/protocol"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/protocol/client"
)


type OperationClientImpl struct{
      interfaceName string
      interfaceDefinition core.InterfaceDefinition
      methodIdentifiers[]core.MethodIdentifier
      methodNameToDefMap map[string]*core.MethodDefinition
      errorBindingMap map[string]bindings.BindingType
      interfaceIdentifier core.InterfaceIdentifier
      connector client.Connector
}


func NewOperationClientImpl(connector client.Connector) *OperationClientImpl {
      interfaceName := "com.vmware.vapi.metadata.authentication.service.operation"
      interfaceIdentifier := core.NewInterfaceIdentifier(interfaceName)
      methodIdentifiers := []core.MethodIdentifier{
          core.NewMethodIdentifier(interfaceIdentifier, "list"),
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
      oIface := OperationClientImpl{interfaceName: interfaceName, methodIdentifiers: methodIdentifiers, interfaceDefinition: interfaceDefinition, errorBindingMap: errorBindingMap, interfaceIdentifier:interfaceIdentifier, connector: connector}
      oIface.methodNameToDefMap = make(map[string]*core.MethodDefinition)
      oIface.methodNameToDefMap["list"] = oIface.listMethodDefinition()
      oIface.methodNameToDefMap["get"] = oIface.getMethodDefinition()
      return &oIface
}

func (oIface *OperationClientImpl) List(serviceIdParam string) ([]string, error) {
	typeConverter := oIface.connector.TypeConverter()
	methodIdentifier := core.NewMethodIdentifier(oIface.interfaceIdentifier, "list")
	sv := bindings.NewStructValueBuilder(operationListInputType(), typeConverter)
	sv.AddStructField("ServiceId", serviceIdParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
        var emptyOutput []string
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := operationListRestMetadata
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	oIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult:= oIface.Invoke(oIface.connector.NewExecutionContext(), methodIdentifier, inputDataValue)
	var emptyOutput []string
    if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), operationListOutputType())
		if errorInOutput != nil {
		    return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
	    return output.([]string), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), oIface.errorBindingMap[methodResult.Error().Name()])
		if errorInError != nil {
		    return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}
func (oIface *OperationClientImpl) Get(serviceIdParam string, operationIdParam string) (authentication.OperationInfo, error) {
	typeConverter := oIface.connector.TypeConverter()
	methodIdentifier := core.NewMethodIdentifier(oIface.interfaceIdentifier, "get")
	sv := bindings.NewStructValueBuilder(operationGetInputType(), typeConverter)
	sv.AddStructField("ServiceId", serviceIdParam)
	sv.AddStructField("OperationId", operationIdParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
        var emptyOutput authentication.OperationInfo
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := operationGetRestMetadata
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	oIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult:= oIface.Invoke(oIface.connector.NewExecutionContext(), methodIdentifier, inputDataValue)
	var emptyOutput authentication.OperationInfo
    if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), operationGetOutputType())
		if errorInOutput != nil {
		    return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
	    return output.(authentication.OperationInfo), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), oIface.errorBindingMap[methodResult.Error().Name()])
		if errorInError != nil {
		    return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (oIface *OperationClientImpl) Invoke(ctx *core.ExecutionContext, methodId core.MethodIdentifier, inputDataValue data.DataValue) core.MethodResult {
    methodResult := oIface.connector.GetApiProvider().Invoke(oIface.interfaceName, methodId.Name(), inputDataValue, ctx)
    return methodResult
}

func (oIface *OperationClientImpl) listMethodDefinition() *core.MethodDefinition {
      interfaceIdentifier := core.NewInterfaceIdentifier(oIface.interfaceName)
      typeConverter := oIface.connector.TypeConverter()

      input, inputError := typeConverter.ConvertToDataDefinition(operationListInputType())
      output, outputError := typeConverter.ConvertToDataDefinition(operationListOutputType())
      if(inputError != nil) {
          log.Errorf("Error in ConvertToDataDefinition for OperationClientImpl.list method's input - %s",
              bindings.VAPIerrorsToError(inputError).Error())
          return nil
      }
      if(outputError != nil) {
          log.Errorf("Error in ConvertToDataDefinition for OperationClientImpl.list method's output - %s",
              bindings.VAPIerrorsToError(outputError).Error())
          return nil
      }
      methodIdentifier := core.NewMethodIdentifier(interfaceIdentifier, "list")
      errorDefinitions := make([]data.ErrorDefinition,0)
      oIface.errorBindingMap[errors.NotFound{}.Error()] = errors.NotFoundBindingType()
	  errDef1, errError1 := typeConverter.ConvertToDataDefinition(errors.NotFoundBindingType())
	  if(errError1 != nil) {
	    log.Errorf("Error in ConvertToDataDefinition for OperationClientImpl.list method's errors.NotFound error - %s",
	        bindings.VAPIerrorsToError(errError1).Error())
        return nil
	  }
	  errorDefinitions = append(errorDefinitions, errDef1.(data.ErrorDefinition))

      methodDefinition := core.NewMethodDefinition(methodIdentifier, input, output, errorDefinitions)
      return &methodDefinition
}
func (oIface *OperationClientImpl) getMethodDefinition() *core.MethodDefinition {
      interfaceIdentifier := core.NewInterfaceIdentifier(oIface.interfaceName)
      typeConverter := oIface.connector.TypeConverter()

      input, inputError := typeConverter.ConvertToDataDefinition(operationGetInputType())
      output, outputError := typeConverter.ConvertToDataDefinition(operationGetOutputType())
      if(inputError != nil) {
          log.Errorf("Error in ConvertToDataDefinition for OperationClientImpl.get method's input - %s",
              bindings.VAPIerrorsToError(inputError).Error())
          return nil
      }
      if(outputError != nil) {
          log.Errorf("Error in ConvertToDataDefinition for OperationClientImpl.get method's output - %s",
              bindings.VAPIerrorsToError(outputError).Error())
          return nil
      }
      methodIdentifier := core.NewMethodIdentifier(interfaceIdentifier, "get")
      errorDefinitions := make([]data.ErrorDefinition,0)
      oIface.errorBindingMap[errors.NotFound{}.Error()] = errors.NotFoundBindingType()
	  errDef1, errError1 := typeConverter.ConvertToDataDefinition(errors.NotFoundBindingType())
	  if(errError1 != nil) {
	    log.Errorf("Error in ConvertToDataDefinition for OperationClientImpl.get method's errors.NotFound error - %s",
	        bindings.VAPIerrorsToError(errError1).Error())
        return nil
	  }
	  errorDefinitions = append(errorDefinitions, errDef1.(data.ErrorDefinition))

      methodDefinition := core.NewMethodDefinition(methodIdentifier, input, output, errorDefinitions)
      return &methodDefinition
}

// Resource type for operation.
const Operation_RESOURCE_TYPE = "com.vmware.vapi.operation"






func operationListInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["service_id"] = bindings.NewIdType([]string {"com.vmware.vapi.service"}, "")
    fieldNameMap["service_id"] = "ServiceId"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func operationListOutputType() bindings.BindingType {
    return bindings.NewListType(bindings.NewIdType([]string {"com.vmware.vapi.operation"}, ""), reflect.TypeOf([]string{}))
}

func operationListRestMetadata() protocol.OperationRestMetadata {
    paramsTypeMap := map[string]bindings.BindingType{}
    pathParams := map[string]string{}
    queryParams := map[string]string{}
    headerParams := map[string]string{}
    resultHeaders := map[string]string{}
    errorHeaders := map[string]string{}
    return protocol.NewOperationRestMetadata(
      paramsTypeMap,
      pathParams,
      queryParams,
      headerParams,
      "",
      "null",
      "",
       resultHeaders,
       0,
       errorHeaders,
       map[string]int{"NotFound": 404})
}


func operationGetInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["service_id"] = bindings.NewIdType([]string {"com.vmware.vapi.service"}, "")
    fields["operation_id"] = bindings.NewIdType([]string {"com.vmware.vapi.operation"}, "")
    fieldNameMap["service_id"] = "ServiceId"
    fieldNameMap["operation_id"] = "OperationId"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func operationGetOutputType() bindings.BindingType {
    return bindings.NewReferenceType(authentication.OperationInfoBindingType)
}

func operationGetRestMetadata() protocol.OperationRestMetadata {
    paramsTypeMap := map[string]bindings.BindingType{}
    pathParams := map[string]string{}
    queryParams := map[string]string{}
    headerParams := map[string]string{}
    resultHeaders := map[string]string{}
    errorHeaders := map[string]string{}
    return protocol.NewOperationRestMetadata(
      paramsTypeMap,
      pathParams,
      queryParams,
      headerParams,
      "",
      "null",
      "",
       resultHeaders,
       0,
       errorHeaders,
       map[string]int{"NotFound": 404})
}




