
/* Copyright © 2019 VMware, Inc. All Rights Reserved.
     SPDX-License-Identifier: BSD-2-Clause */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Client stubs for service: Peerconfig
 * Functions that implement the generated PeerconfigClient interface
 */


package edges

import (
    "reflect"
// //     "getClientImplDependenciesOfOps gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/bindings/com/vmware/vapi/std/errors"
// //     "getClientImplDependenciesOfOps gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/bindings"
// //     "getClientImplDependenciesOfOps gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/core"
// //     "getClientImplDependenciesOfOps gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/data"
// //     "getClientImplDependenciesOfOps gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/lib"
// //     "getClientImplDependenciesOfOps gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/log"
// //     "getClientImplDependenciesOfOps gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/protocol/client"
// //
// //     "getDependenciesOfServiceTypes gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/bindings"
// //     "getDependenciesOfServiceTypes gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/data"
// //     "getDependenciesOfServiceTypes gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/protocol"
// 
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/bindings/com/vmware/vapi/std/errors"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/bindings"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/core"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/data"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/lib"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/log"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/protocol"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/protocol/client"
)


type PeerconfigClientImpl struct{
      interfaceName string
      interfaceDefinition core.InterfaceDefinition
      methodIdentifiers[]core.MethodIdentifier
      methodNameToDefMap map[string]*core.MethodDefinition
      errorBindingMap map[string]bindings.BindingType
      interfaceIdentifier core.InterfaceIdentifier
      connector client.Connector
}


func NewPeerconfigClientImpl(connector client.Connector) *PeerconfigClientImpl {
      interfaceName := "com.vmware.vmc.orgs.sddcs.networks.edges.peerconfig"
      interfaceIdentifier := core.NewInterfaceIdentifier(interfaceName)
      methodIdentifiers := []core.MethodIdentifier{
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
      pIface := PeerconfigClientImpl{interfaceName: interfaceName, methodIdentifiers: methodIdentifiers, interfaceDefinition: interfaceDefinition, errorBindingMap: errorBindingMap, interfaceIdentifier:interfaceIdentifier, connector: connector}
      pIface.methodNameToDefMap = make(map[string]*core.MethodDefinition)
      pIface.methodNameToDefMap["get"] = pIface.getMethodDefinition()
      return &pIface
}

func (pIface *PeerconfigClientImpl) Get(orgParam string, sddcParam string, edgeIdParam string, objecttypeParam string, objectidParam string, templateidParam *string) (*data.StructValue, error) {
	typeConverter := pIface.connector.TypeConverter()
	methodIdentifier := core.NewMethodIdentifier(pIface.interfaceIdentifier, "get")
	sv := bindings.NewStructValueBuilder(peerconfigGetInputType(), typeConverter)
	sv.AddStructField("Org", orgParam)
	sv.AddStructField("Sddc", sddcParam)
	sv.AddStructField("EdgeId", edgeIdParam)
	sv.AddStructField("Objecttype", objecttypeParam)
	sv.AddStructField("Objectid", objectidParam)
	sv.AddStructField("Templateid", templateidParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
        var emptyOutput *data.StructValue
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := peerconfigGetRestMetadata
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	pIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult:= pIface.Invoke(pIface.connector.NewExecutionContext(), methodIdentifier, inputDataValue)
	var emptyOutput *data.StructValue
    if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), peerconfigGetOutputType())
		if errorInOutput != nil {
		    return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
	    return output.(*data.StructValue), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), pIface.errorBindingMap[methodResult.Error().Name()])
		if errorInError != nil {
		    return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (pIface *PeerconfigClientImpl) Invoke(ctx *core.ExecutionContext, methodId core.MethodIdentifier, inputDataValue data.DataValue) core.MethodResult {
    methodResult := pIface.connector.GetApiProvider().Invoke(pIface.interfaceName, methodId.Name(), inputDataValue, ctx)
    return methodResult
}

func (pIface *PeerconfigClientImpl) getMethodDefinition() *core.MethodDefinition {
      interfaceIdentifier := core.NewInterfaceIdentifier(pIface.interfaceName)
      typeConverter := pIface.connector.TypeConverter()

      input, inputError := typeConverter.ConvertToDataDefinition(peerconfigGetInputType())
      output, outputError := typeConverter.ConvertToDataDefinition(peerconfigGetOutputType())
      if(inputError != nil) {
          log.Errorf("Error in ConvertToDataDefinition for PeerconfigClientImpl.get method's input - %s",
              bindings.VAPIerrorsToError(inputError).Error())
          return nil
      }
      if(outputError != nil) {
          log.Errorf("Error in ConvertToDataDefinition for PeerconfigClientImpl.get method's output - %s",
              bindings.VAPIerrorsToError(outputError).Error())
          return nil
      }
      methodIdentifier := core.NewMethodIdentifier(interfaceIdentifier, "get")
      errorDefinitions := make([]data.ErrorDefinition,0)
      pIface.errorBindingMap[errors.InvalidRequest{}.Error()] = errors.InvalidRequestBindingType()
	  errDef1, errError1 := typeConverter.ConvertToDataDefinition(errors.InvalidRequestBindingType())
	  if(errError1 != nil) {
	    log.Errorf("Error in ConvertToDataDefinition for PeerconfigClientImpl.get method's errors.InvalidRequest error - %s",
	        bindings.VAPIerrorsToError(errError1).Error())
        return nil
	  }
	  errorDefinitions = append(errorDefinitions, errDef1.(data.ErrorDefinition))
      pIface.errorBindingMap[errors.Unauthorized{}.Error()] = errors.UnauthorizedBindingType()
	  errDef2, errError2 := typeConverter.ConvertToDataDefinition(errors.UnauthorizedBindingType())
	  if(errError2 != nil) {
	    log.Errorf("Error in ConvertToDataDefinition for PeerconfigClientImpl.get method's errors.Unauthorized error - %s",
	        bindings.VAPIerrorsToError(errError2).Error())
        return nil
	  }
	  errorDefinitions = append(errorDefinitions, errDef2.(data.ErrorDefinition))
      pIface.errorBindingMap[errors.NotFound{}.Error()] = errors.NotFoundBindingType()
	  errDef3, errError3 := typeConverter.ConvertToDataDefinition(errors.NotFoundBindingType())
	  if(errError3 != nil) {
	    log.Errorf("Error in ConvertToDataDefinition for PeerconfigClientImpl.get method's errors.NotFound error - %s",
	        bindings.VAPIerrorsToError(errError3).Error())
        return nil
	  }
	  errorDefinitions = append(errorDefinitions, errDef3.(data.ErrorDefinition))

      methodDefinition := core.NewMethodDefinition(methodIdentifier, input, output, errorDefinitions)
      return &methodDefinition
}







func peerconfigGetInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["org"] = bindings.NewStringType()
    fields["sddc"] = bindings.NewStringType()
    fields["edge_id"] = bindings.NewStringType()
    fields["objecttype"] = bindings.NewStringType()
    fields["objectid"] = bindings.NewStringType()
    fields["templateid"] = bindings.NewOptionalType(bindings.NewStringType())
    fieldNameMap["org"] = "Org"
    fieldNameMap["sddc"] = "Sddc"
    fieldNameMap["edge_id"] = "EdgeId"
    fieldNameMap["objecttype"] = "Objecttype"
    fieldNameMap["objectid"] = "Objectid"
    fieldNameMap["templateid"] = "Templateid"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func peerconfigGetOutputType() bindings.BindingType {
    return bindings.NewDynamicStructType(nil, bindings.REST)
}

func peerconfigGetRestMetadata() protocol.OperationRestMetadata {
    paramsTypeMap := map[string]bindings.BindingType{}
    pathParams := map[string]string{}
    queryParams := map[string]string{}
    headerParams := map[string]string{}
    paramsTypeMap["templateid"] = bindings.NewOptionalType(bindings.NewStringType())
    paramsTypeMap["org"] = bindings.NewStringType()
    paramsTypeMap["sddc"] = bindings.NewStringType()
    paramsTypeMap["edge_id"] = bindings.NewStringType()
    paramsTypeMap["objecttype"] = bindings.NewStringType()
    paramsTypeMap["objectid"] = bindings.NewStringType()
    paramsTypeMap["org"] = bindings.NewStringType()
    paramsTypeMap["sddc"] = bindings.NewStringType()
    paramsTypeMap["edgeId"] = bindings.NewStringType()
    pathParams["edge_id"] = "edgeId"
    pathParams["org"] = "org"
    pathParams["sddc"] = "sddc"
    queryParams["objecttype"] = "objecttype"
    queryParams["templateid"] = "templateid"
    queryParams["objectid"] = "objectid"
    resultHeaders := map[string]string{}
    errorHeaders := map[string]string{}
    return protocol.NewOperationRestMetadata(
      paramsTypeMap,
      pathParams,
      queryParams,
      headerParams,
      "",
      "GET",
      "/vmc/api/orgs/{org}/sddcs/{sddc}/networks/4.0/edges/{edgeId}/peerconfig",
       resultHeaders,
       200,
       errorHeaders,
       map[string]int{"InvalidRequest": 400,"Unauthorized": 403,"NotFound": 404})
}




