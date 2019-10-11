
/* Copyright © 2019 VMware, Inc. All Rights Reserved.
     SPDX-License-Identifier: BSD-2-Clause */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Client stubs for service: Vnics
 * Functions that implement the generated VnicsClient interface
 */


package edges

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


type VnicsClientImpl struct{
      interfaceName string
      interfaceDefinition core.InterfaceDefinition
      methodIdentifiers[]core.MethodIdentifier
      methodNameToDefMap map[string]*core.MethodDefinition
      errorBindingMap map[string]bindings.BindingType
      interfaceIdentifier core.InterfaceIdentifier
      connector client.Connector
}


func NewVnicsClientImpl(connector client.Connector) *VnicsClientImpl {
      interfaceName := "com.vmware.vmc.orgs.sddcs.networks.edges.vnics"
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
      vIface := VnicsClientImpl{interfaceName: interfaceName, methodIdentifiers: methodIdentifiers, interfaceDefinition: interfaceDefinition, errorBindingMap: errorBindingMap, interfaceIdentifier:interfaceIdentifier, connector: connector}
      vIface.methodNameToDefMap = make(map[string]*core.MethodDefinition)
      vIface.methodNameToDefMap["get"] = vIface.getMethodDefinition()
      return &vIface
}

func (vIface *VnicsClientImpl) Get(orgParam string, sddcParam string, edgeIdParam string) (model.Vnics, error) {
	typeConverter := vIface.connector.TypeConverter()
	methodIdentifier := core.NewMethodIdentifier(vIface.interfaceIdentifier, "get")
	sv := bindings.NewStructValueBuilder(vnicsGetInputType(), typeConverter)
	sv.AddStructField("Org", orgParam)
	sv.AddStructField("Sddc", sddcParam)
	sv.AddStructField("EdgeId", edgeIdParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
        var emptyOutput model.Vnics
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := vnicsGetRestMetadata
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	vIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult:= vIface.Invoke(vIface.connector.NewExecutionContext(), methodIdentifier, inputDataValue)
	var emptyOutput model.Vnics
    if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), vnicsGetOutputType())
		if errorInOutput != nil {
		    return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
	    return output.(model.Vnics), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), vIface.errorBindingMap[methodResult.Error().Name()])
		if errorInError != nil {
		    return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (vIface *VnicsClientImpl) Invoke(ctx *core.ExecutionContext, methodId core.MethodIdentifier, inputDataValue data.DataValue) core.MethodResult {
    methodResult := vIface.connector.GetApiProvider().Invoke(vIface.interfaceName, methodId.Name(), inputDataValue, ctx)
    return methodResult
}

func (vIface *VnicsClientImpl) getMethodDefinition() *core.MethodDefinition {
      interfaceIdentifier := core.NewInterfaceIdentifier(vIface.interfaceName)
      typeConverter := vIface.connector.TypeConverter()

      input, inputError := typeConverter.ConvertToDataDefinition(vnicsGetInputType())
      output, outputError := typeConverter.ConvertToDataDefinition(vnicsGetOutputType())
      if(inputError != nil) {
          log.Errorf("Error in ConvertToDataDefinition for VnicsClientImpl.get method's input - %s",
              bindings.VAPIerrorsToError(inputError).Error())
          return nil
      }
      if(outputError != nil) {
          log.Errorf("Error in ConvertToDataDefinition for VnicsClientImpl.get method's output - %s",
              bindings.VAPIerrorsToError(outputError).Error())
          return nil
      }
      methodIdentifier := core.NewMethodIdentifier(interfaceIdentifier, "get")
      errorDefinitions := make([]data.ErrorDefinition,0)
      vIface.errorBindingMap[errors.InvalidRequest{}.Error()] = errors.InvalidRequestBindingType()
	  errDef1, errError1 := typeConverter.ConvertToDataDefinition(errors.InvalidRequestBindingType())
	  if(errError1 != nil) {
	    log.Errorf("Error in ConvertToDataDefinition for VnicsClientImpl.get method's errors.InvalidRequest error - %s",
	        bindings.VAPIerrorsToError(errError1).Error())
        return nil
	  }
	  errorDefinitions = append(errorDefinitions, errDef1.(data.ErrorDefinition))
      vIface.errorBindingMap[errors.Unauthorized{}.Error()] = errors.UnauthorizedBindingType()
	  errDef2, errError2 := typeConverter.ConvertToDataDefinition(errors.UnauthorizedBindingType())
	  if(errError2 != nil) {
	    log.Errorf("Error in ConvertToDataDefinition for VnicsClientImpl.get method's errors.Unauthorized error - %s",
	        bindings.VAPIerrorsToError(errError2).Error())
        return nil
	  }
	  errorDefinitions = append(errorDefinitions, errDef2.(data.ErrorDefinition))
      vIface.errorBindingMap[errors.NotFound{}.Error()] = errors.NotFoundBindingType()
	  errDef3, errError3 := typeConverter.ConvertToDataDefinition(errors.NotFoundBindingType())
	  if(errError3 != nil) {
	    log.Errorf("Error in ConvertToDataDefinition for VnicsClientImpl.get method's errors.NotFound error - %s",
	        bindings.VAPIerrorsToError(errError3).Error())
        return nil
	  }
	  errorDefinitions = append(errorDefinitions, errDef3.(data.ErrorDefinition))

      methodDefinition := core.NewMethodDefinition(methodIdentifier, input, output, errorDefinitions)
      return &methodDefinition
}







func vnicsGetInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["org"] = bindings.NewStringType()
    fields["sddc"] = bindings.NewStringType()
    fields["edge_id"] = bindings.NewStringType()
    fieldNameMap["org"] = "Org"
    fieldNameMap["sddc"] = "Sddc"
    fieldNameMap["edge_id"] = "EdgeId"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func vnicsGetOutputType() bindings.BindingType {
    return bindings.NewReferenceType(model.VnicsBindingType)
}

func vnicsGetRestMetadata() protocol.OperationRestMetadata {
    paramsTypeMap := map[string]bindings.BindingType{}
    pathParams := map[string]string{}
    queryParams := map[string]string{}
    headerParams := map[string]string{}
    paramsTypeMap["org"] = bindings.NewStringType()
    paramsTypeMap["sddc"] = bindings.NewStringType()
    paramsTypeMap["edge_id"] = bindings.NewStringType()
    paramsTypeMap["org"] = bindings.NewStringType()
    paramsTypeMap["sddc"] = bindings.NewStringType()
    paramsTypeMap["edgeId"] = bindings.NewStringType()
    pathParams["edge_id"] = "edgeId"
    pathParams["org"] = "org"
    pathParams["sddc"] = "sddc"
    resultHeaders := map[string]string{}
    errorHeaders := map[string]string{}
    return protocol.NewOperationRestMetadata(
      paramsTypeMap,
      pathParams,
      queryParams,
      headerParams,
      "",
      "GET",
      "/vmc/api/orgs/{org}/sddcs/{sddc}/networks/4.0/edges/{edgeId}/vnics",
       resultHeaders,
       200,
       errorHeaders,
       map[string]int{"InvalidRequest": 400,"Unauthorized": 403,"NotFound": 404})
}




