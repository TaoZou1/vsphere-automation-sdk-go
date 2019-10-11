
/* Copyright © 2019 VMware, Inc. All Rights Reserved.
     SPDX-License-Identifier: BSD-2-Clause */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Client stubs for service: Credentials
 * Functions that implement the generated CredentialsClient interface
 */


package addons

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


type CredentialsClientImpl struct{
      interfaceName string
      interfaceDefinition core.InterfaceDefinition
      methodIdentifiers[]core.MethodIdentifier
      methodNameToDefMap map[string]*core.MethodDefinition
      errorBindingMap map[string]bindings.BindingType
      interfaceIdentifier core.InterfaceIdentifier
      connector client.Connector
}


func NewCredentialsClientImpl(connector client.Connector) *CredentialsClientImpl {
      interfaceName := "com.vmware.vmc.orgs.sddcs.addons.credentials"
      interfaceIdentifier := core.NewInterfaceIdentifier(interfaceName)
      methodIdentifiers := []core.MethodIdentifier{
          core.NewMethodIdentifier(interfaceIdentifier, "create"),
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
      cIface := CredentialsClientImpl{interfaceName: interfaceName, methodIdentifiers: methodIdentifiers, interfaceDefinition: interfaceDefinition, errorBindingMap: errorBindingMap, interfaceIdentifier:interfaceIdentifier, connector: connector}
      cIface.methodNameToDefMap = make(map[string]*core.MethodDefinition)
      cIface.methodNameToDefMap["create"] = cIface.createMethodDefinition()
      cIface.methodNameToDefMap["get"] = cIface.getMethodDefinition()
      cIface.methodNameToDefMap["list"] = cIface.listMethodDefinition()
      cIface.methodNameToDefMap["update"] = cIface.updateMethodDefinition()
      return &cIface
}

func (cIface *CredentialsClientImpl) Create(orgParam string, sddcIdParam string, addonTypeParam string, credentialsParam model.NewCredentials) (model.NewCredentials, error) {
	typeConverter := cIface.connector.TypeConverter()
	methodIdentifier := core.NewMethodIdentifier(cIface.interfaceIdentifier, "create")
	sv := bindings.NewStructValueBuilder(credentialsCreateInputType(), typeConverter)
	sv.AddStructField("Org", orgParam)
	sv.AddStructField("SddcId", sddcIdParam)
	sv.AddStructField("AddonType", addonTypeParam)
	sv.AddStructField("Credentials", credentialsParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
        var emptyOutput model.NewCredentials
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := credentialsCreateRestMetadata
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	cIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult:= cIface.Invoke(cIface.connector.NewExecutionContext(), methodIdentifier, inputDataValue)
	var emptyOutput model.NewCredentials
    if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), credentialsCreateOutputType())
		if errorInOutput != nil {
		    return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
	    return output.(model.NewCredentials), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), cIface.errorBindingMap[methodResult.Error().Name()])
		if errorInError != nil {
		    return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}
func (cIface *CredentialsClientImpl) Get(orgParam string, sddcIdParam string, addonTypeParam string, nameParam string) (model.NewCredentials, error) {
	typeConverter := cIface.connector.TypeConverter()
	methodIdentifier := core.NewMethodIdentifier(cIface.interfaceIdentifier, "get")
	sv := bindings.NewStructValueBuilder(credentialsGetInputType(), typeConverter)
	sv.AddStructField("Org", orgParam)
	sv.AddStructField("SddcId", sddcIdParam)
	sv.AddStructField("AddonType", addonTypeParam)
	sv.AddStructField("Name", nameParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
        var emptyOutput model.NewCredentials
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := credentialsGetRestMetadata
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	cIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult:= cIface.Invoke(cIface.connector.NewExecutionContext(), methodIdentifier, inputDataValue)
	var emptyOutput model.NewCredentials
    if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), credentialsGetOutputType())
		if errorInOutput != nil {
		    return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
	    return output.(model.NewCredentials), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), cIface.errorBindingMap[methodResult.Error().Name()])
		if errorInError != nil {
		    return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}
func (cIface *CredentialsClientImpl) List(orgParam string, sddcIdParam string, addonTypeParam string) ([]model.NewCredentials, error) {
	typeConverter := cIface.connector.TypeConverter()
	methodIdentifier := core.NewMethodIdentifier(cIface.interfaceIdentifier, "list")
	sv := bindings.NewStructValueBuilder(credentialsListInputType(), typeConverter)
	sv.AddStructField("Org", orgParam)
	sv.AddStructField("SddcId", sddcIdParam)
	sv.AddStructField("AddonType", addonTypeParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
        var emptyOutput []model.NewCredentials
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := credentialsListRestMetadata
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	cIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult:= cIface.Invoke(cIface.connector.NewExecutionContext(), methodIdentifier, inputDataValue)
	var emptyOutput []model.NewCredentials
    if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), credentialsListOutputType())
		if errorInOutput != nil {
		    return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
	    return output.([]model.NewCredentials), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), cIface.errorBindingMap[methodResult.Error().Name()])
		if errorInError != nil {
		    return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}
func (cIface *CredentialsClientImpl) Update(orgParam string, sddcIdParam string, addonTypeParam string, nameParam string, credentialsParam model.UpdateCredentials) (model.NewCredentials, error) {
	typeConverter := cIface.connector.TypeConverter()
	methodIdentifier := core.NewMethodIdentifier(cIface.interfaceIdentifier, "update")
	sv := bindings.NewStructValueBuilder(credentialsUpdateInputType(), typeConverter)
	sv.AddStructField("Org", orgParam)
	sv.AddStructField("SddcId", sddcIdParam)
	sv.AddStructField("AddonType", addonTypeParam)
	sv.AddStructField("Name", nameParam)
	sv.AddStructField("Credentials", credentialsParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
        var emptyOutput model.NewCredentials
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := credentialsUpdateRestMetadata
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	cIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult:= cIface.Invoke(cIface.connector.NewExecutionContext(), methodIdentifier, inputDataValue)
	var emptyOutput model.NewCredentials
    if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), credentialsUpdateOutputType())
		if errorInOutput != nil {
		    return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
	    return output.(model.NewCredentials), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), cIface.errorBindingMap[methodResult.Error().Name()])
		if errorInError != nil {
		    return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (cIface *CredentialsClientImpl) Invoke(ctx *core.ExecutionContext, methodId core.MethodIdentifier, inputDataValue data.DataValue) core.MethodResult {
    methodResult := cIface.connector.GetApiProvider().Invoke(cIface.interfaceName, methodId.Name(), inputDataValue, ctx)
    return methodResult
}

func (cIface *CredentialsClientImpl) createMethodDefinition() *core.MethodDefinition {
      interfaceIdentifier := core.NewInterfaceIdentifier(cIface.interfaceName)
      typeConverter := cIface.connector.TypeConverter()

      input, inputError := typeConverter.ConvertToDataDefinition(credentialsCreateInputType())
      output, outputError := typeConverter.ConvertToDataDefinition(credentialsCreateOutputType())
      if(inputError != nil) {
          log.Errorf("Error in ConvertToDataDefinition for CredentialsClientImpl.create method's input - %s",
              bindings.VAPIerrorsToError(inputError).Error())
          return nil
      }
      if(outputError != nil) {
          log.Errorf("Error in ConvertToDataDefinition for CredentialsClientImpl.create method's output - %s",
              bindings.VAPIerrorsToError(outputError).Error())
          return nil
      }
      methodIdentifier := core.NewMethodIdentifier(interfaceIdentifier, "create")
      errorDefinitions := make([]data.ErrorDefinition,0)
      cIface.errorBindingMap[errors.ConcurrentChange{}.Error()] = errors.ConcurrentChangeBindingType()
	  errDef1, errError1 := typeConverter.ConvertToDataDefinition(errors.ConcurrentChangeBindingType())
	  if(errError1 != nil) {
	    log.Errorf("Error in ConvertToDataDefinition for CredentialsClientImpl.create method's errors.ConcurrentChange error - %s",
	        bindings.VAPIerrorsToError(errError1).Error())
        return nil
	  }
	  errorDefinitions = append(errorDefinitions, errDef1.(data.ErrorDefinition))
      cIface.errorBindingMap[errors.InvalidRequest{}.Error()] = errors.InvalidRequestBindingType()
	  errDef2, errError2 := typeConverter.ConvertToDataDefinition(errors.InvalidRequestBindingType())
	  if(errError2 != nil) {
	    log.Errorf("Error in ConvertToDataDefinition for CredentialsClientImpl.create method's errors.InvalidRequest error - %s",
	        bindings.VAPIerrorsToError(errError2).Error())
        return nil
	  }
	  errorDefinitions = append(errorDefinitions, errDef2.(data.ErrorDefinition))
      cIface.errorBindingMap[errors.Unauthorized{}.Error()] = errors.UnauthorizedBindingType()
	  errDef3, errError3 := typeConverter.ConvertToDataDefinition(errors.UnauthorizedBindingType())
	  if(errError3 != nil) {
	    log.Errorf("Error in ConvertToDataDefinition for CredentialsClientImpl.create method's errors.Unauthorized error - %s",
	        bindings.VAPIerrorsToError(errError3).Error())
        return nil
	  }
	  errorDefinitions = append(errorDefinitions, errDef3.(data.ErrorDefinition))

      methodDefinition := core.NewMethodDefinition(methodIdentifier, input, output, errorDefinitions)
      return &methodDefinition
}
func (cIface *CredentialsClientImpl) getMethodDefinition() *core.MethodDefinition {
      interfaceIdentifier := core.NewInterfaceIdentifier(cIface.interfaceName)
      typeConverter := cIface.connector.TypeConverter()

      input, inputError := typeConverter.ConvertToDataDefinition(credentialsGetInputType())
      output, outputError := typeConverter.ConvertToDataDefinition(credentialsGetOutputType())
      if(inputError != nil) {
          log.Errorf("Error in ConvertToDataDefinition for CredentialsClientImpl.get method's input - %s",
              bindings.VAPIerrorsToError(inputError).Error())
          return nil
      }
      if(outputError != nil) {
          log.Errorf("Error in ConvertToDataDefinition for CredentialsClientImpl.get method's output - %s",
              bindings.VAPIerrorsToError(outputError).Error())
          return nil
      }
      methodIdentifier := core.NewMethodIdentifier(interfaceIdentifier, "get")
      errorDefinitions := make([]data.ErrorDefinition,0)
      cIface.errorBindingMap[errors.Unauthorized{}.Error()] = errors.UnauthorizedBindingType()
	  errDef1, errError1 := typeConverter.ConvertToDataDefinition(errors.UnauthorizedBindingType())
	  if(errError1 != nil) {
	    log.Errorf("Error in ConvertToDataDefinition for CredentialsClientImpl.get method's errors.Unauthorized error - %s",
	        bindings.VAPIerrorsToError(errError1).Error())
        return nil
	  }
	  errorDefinitions = append(errorDefinitions, errDef1.(data.ErrorDefinition))

      methodDefinition := core.NewMethodDefinition(methodIdentifier, input, output, errorDefinitions)
      return &methodDefinition
}
func (cIface *CredentialsClientImpl) listMethodDefinition() *core.MethodDefinition {
      interfaceIdentifier := core.NewInterfaceIdentifier(cIface.interfaceName)
      typeConverter := cIface.connector.TypeConverter()

      input, inputError := typeConverter.ConvertToDataDefinition(credentialsListInputType())
      output, outputError := typeConverter.ConvertToDataDefinition(credentialsListOutputType())
      if(inputError != nil) {
          log.Errorf("Error in ConvertToDataDefinition for CredentialsClientImpl.list method's input - %s",
              bindings.VAPIerrorsToError(inputError).Error())
          return nil
      }
      if(outputError != nil) {
          log.Errorf("Error in ConvertToDataDefinition for CredentialsClientImpl.list method's output - %s",
              bindings.VAPIerrorsToError(outputError).Error())
          return nil
      }
      methodIdentifier := core.NewMethodIdentifier(interfaceIdentifier, "list")
      errorDefinitions := make([]data.ErrorDefinition,0)
      cIface.errorBindingMap[errors.Unauthorized{}.Error()] = errors.UnauthorizedBindingType()
	  errDef1, errError1 := typeConverter.ConvertToDataDefinition(errors.UnauthorizedBindingType())
	  if(errError1 != nil) {
	    log.Errorf("Error in ConvertToDataDefinition for CredentialsClientImpl.list method's errors.Unauthorized error - %s",
	        bindings.VAPIerrorsToError(errError1).Error())
        return nil
	  }
	  errorDefinitions = append(errorDefinitions, errDef1.(data.ErrorDefinition))

      methodDefinition := core.NewMethodDefinition(methodIdentifier, input, output, errorDefinitions)
      return &methodDefinition
}
func (cIface *CredentialsClientImpl) updateMethodDefinition() *core.MethodDefinition {
      interfaceIdentifier := core.NewInterfaceIdentifier(cIface.interfaceName)
      typeConverter := cIface.connector.TypeConverter()

      input, inputError := typeConverter.ConvertToDataDefinition(credentialsUpdateInputType())
      output, outputError := typeConverter.ConvertToDataDefinition(credentialsUpdateOutputType())
      if(inputError != nil) {
          log.Errorf("Error in ConvertToDataDefinition for CredentialsClientImpl.update method's input - %s",
              bindings.VAPIerrorsToError(inputError).Error())
          return nil
      }
      if(outputError != nil) {
          log.Errorf("Error in ConvertToDataDefinition for CredentialsClientImpl.update method's output - %s",
              bindings.VAPIerrorsToError(outputError).Error())
          return nil
      }
      methodIdentifier := core.NewMethodIdentifier(interfaceIdentifier, "update")
      errorDefinitions := make([]data.ErrorDefinition,0)
      cIface.errorBindingMap[errors.InvalidRequest{}.Error()] = errors.InvalidRequestBindingType()
	  errDef1, errError1 := typeConverter.ConvertToDataDefinition(errors.InvalidRequestBindingType())
	  if(errError1 != nil) {
	    log.Errorf("Error in ConvertToDataDefinition for CredentialsClientImpl.update method's errors.InvalidRequest error - %s",
	        bindings.VAPIerrorsToError(errError1).Error())
        return nil
	  }
	  errorDefinitions = append(errorDefinitions, errDef1.(data.ErrorDefinition))
      cIface.errorBindingMap[errors.Unauthorized{}.Error()] = errors.UnauthorizedBindingType()
	  errDef2, errError2 := typeConverter.ConvertToDataDefinition(errors.UnauthorizedBindingType())
	  if(errError2 != nil) {
	    log.Errorf("Error in ConvertToDataDefinition for CredentialsClientImpl.update method's errors.Unauthorized error - %s",
	        bindings.VAPIerrorsToError(errError2).Error())
        return nil
	  }
	  errorDefinitions = append(errorDefinitions, errDef2.(data.ErrorDefinition))

      methodDefinition := core.NewMethodDefinition(methodIdentifier, input, output, errorDefinitions)
      return &methodDefinition
}

// Possible value for ``addonType`` of method Credentials#create.
const Credentials_CREATE_ADDON_TYPE_HCX = "HCX"
// Possible value for ``addonType`` of method Credentials#get.
const Credentials_GET_ADDON_TYPE_HCX = "HCX"
// Possible value for ``addonType`` of method Credentials#list.
const Credentials_LIST_ADDON_TYPE_HCX = "HCX"
// Possible value for ``addonType`` of method Credentials#update.
const Credentials_UPDATE_ADDON_TYPE_HCX = "HCX"






func credentialsCreateInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["org"] = bindings.NewStringType()
    fields["sddc_id"] = bindings.NewStringType()
    fields["addon_type"] = bindings.NewStringType()
    fields["credentials"] = bindings.NewReferenceType(model.NewCredentialsBindingType)
    fieldNameMap["org"] = "Org"
    fieldNameMap["sddc_id"] = "SddcId"
    fieldNameMap["addon_type"] = "AddonType"
    fieldNameMap["credentials"] = "Credentials"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func credentialsCreateOutputType() bindings.BindingType {
    return bindings.NewReferenceType(model.NewCredentialsBindingType)
}

func credentialsCreateRestMetadata() protocol.OperationRestMetadata {
    paramsTypeMap := map[string]bindings.BindingType{}
    pathParams := map[string]string{}
    queryParams := map[string]string{}
    headerParams := map[string]string{}
    paramsTypeMap["org"] = bindings.NewStringType()
    paramsTypeMap["addon_type"] = bindings.NewStringType()
    paramsTypeMap["credentials"] = bindings.NewReferenceType(model.NewCredentialsBindingType)
    paramsTypeMap["sddc_id"] = bindings.NewStringType()
    paramsTypeMap["org"] = bindings.NewStringType()
    paramsTypeMap["sddcId"] = bindings.NewStringType()
    paramsTypeMap["addonType"] = bindings.NewStringType()
    pathParams["org"] = "org"
    pathParams["sddc_id"] = "sddcId"
    pathParams["addon_type"] = "addonType"
    resultHeaders := map[string]string{}
    errorHeaders := map[string]string{}
    return protocol.NewOperationRestMetadata(
      paramsTypeMap,
      pathParams,
      queryParams,
      headerParams,
      "credentials",
      "POST",
      "/vmc/api/orgs/{org}/sddcs/{sddcId}/addons/{addonType}/credentials",
       resultHeaders,
       201,
       errorHeaders,
       map[string]int{"ConcurrentChange": 409,"InvalidRequest": 400,"Unauthorized": 403})
}


func credentialsGetInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["org"] = bindings.NewStringType()
    fields["sddc_id"] = bindings.NewStringType()
    fields["addon_type"] = bindings.NewStringType()
    fields["name"] = bindings.NewStringType()
    fieldNameMap["org"] = "Org"
    fieldNameMap["sddc_id"] = "SddcId"
    fieldNameMap["addon_type"] = "AddonType"
    fieldNameMap["name"] = "Name"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func credentialsGetOutputType() bindings.BindingType {
    return bindings.NewReferenceType(model.NewCredentialsBindingType)
}

func credentialsGetRestMetadata() protocol.OperationRestMetadata {
    paramsTypeMap := map[string]bindings.BindingType{}
    pathParams := map[string]string{}
    queryParams := map[string]string{}
    headerParams := map[string]string{}
    paramsTypeMap["org"] = bindings.NewStringType()
    paramsTypeMap["addon_type"] = bindings.NewStringType()
    paramsTypeMap["name"] = bindings.NewStringType()
    paramsTypeMap["sddc_id"] = bindings.NewStringType()
    paramsTypeMap["org"] = bindings.NewStringType()
    paramsTypeMap["sddcId"] = bindings.NewStringType()
    paramsTypeMap["addonType"] = bindings.NewStringType()
    paramsTypeMap["name"] = bindings.NewStringType()
    pathParams["org"] = "org"
    pathParams["sddc_id"] = "sddcId"
    pathParams["addon_type"] = "addonType"
    pathParams["name"] = "name"
    resultHeaders := map[string]string{}
    errorHeaders := map[string]string{}
    return protocol.NewOperationRestMetadata(
      paramsTypeMap,
      pathParams,
      queryParams,
      headerParams,
      "",
      "GET",
      "/vmc/api/orgs/{org}/sddcs/{sddcId}/addons/{addonType}/credentials/{name}",
       resultHeaders,
       200,
       errorHeaders,
       map[string]int{"Unauthorized": 403})
}


func credentialsListInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["org"] = bindings.NewStringType()
    fields["sddc_id"] = bindings.NewStringType()
    fields["addon_type"] = bindings.NewStringType()
    fieldNameMap["org"] = "Org"
    fieldNameMap["sddc_id"] = "SddcId"
    fieldNameMap["addon_type"] = "AddonType"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func credentialsListOutputType() bindings.BindingType {
    return bindings.NewListType(bindings.NewReferenceType(model.NewCredentialsBindingType), reflect.TypeOf([]model.NewCredentials{}))
}

func credentialsListRestMetadata() protocol.OperationRestMetadata {
    paramsTypeMap := map[string]bindings.BindingType{}
    pathParams := map[string]string{}
    queryParams := map[string]string{}
    headerParams := map[string]string{}
    paramsTypeMap["org"] = bindings.NewStringType()
    paramsTypeMap["addon_type"] = bindings.NewStringType()
    paramsTypeMap["sddc_id"] = bindings.NewStringType()
    paramsTypeMap["org"] = bindings.NewStringType()
    paramsTypeMap["sddcId"] = bindings.NewStringType()
    paramsTypeMap["addonType"] = bindings.NewStringType()
    pathParams["org"] = "org"
    pathParams["sddc_id"] = "sddcId"
    pathParams["addon_type"] = "addonType"
    resultHeaders := map[string]string{}
    errorHeaders := map[string]string{}
    return protocol.NewOperationRestMetadata(
      paramsTypeMap,
      pathParams,
      queryParams,
      headerParams,
      "",
      "GET",
      "/vmc/api/orgs/{org}/sddcs/{sddcId}/addons/{addonType}/credentials",
       resultHeaders,
       200,
       errorHeaders,
       map[string]int{"Unauthorized": 403})
}


func credentialsUpdateInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["org"] = bindings.NewStringType()
    fields["sddc_id"] = bindings.NewStringType()
    fields["addon_type"] = bindings.NewStringType()
    fields["name"] = bindings.NewStringType()
    fields["credentials"] = bindings.NewReferenceType(model.UpdateCredentialsBindingType)
    fieldNameMap["org"] = "Org"
    fieldNameMap["sddc_id"] = "SddcId"
    fieldNameMap["addon_type"] = "AddonType"
    fieldNameMap["name"] = "Name"
    fieldNameMap["credentials"] = "Credentials"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func credentialsUpdateOutputType() bindings.BindingType {
    return bindings.NewReferenceType(model.NewCredentialsBindingType)
}

func credentialsUpdateRestMetadata() protocol.OperationRestMetadata {
    paramsTypeMap := map[string]bindings.BindingType{}
    pathParams := map[string]string{}
    queryParams := map[string]string{}
    headerParams := map[string]string{}
    paramsTypeMap["org"] = bindings.NewStringType()
    paramsTypeMap["addon_type"] = bindings.NewStringType()
    paramsTypeMap["name"] = bindings.NewStringType()
    paramsTypeMap["credentials"] = bindings.NewReferenceType(model.UpdateCredentialsBindingType)
    paramsTypeMap["sddc_id"] = bindings.NewStringType()
    paramsTypeMap["org"] = bindings.NewStringType()
    paramsTypeMap["sddcId"] = bindings.NewStringType()
    paramsTypeMap["addonType"] = bindings.NewStringType()
    paramsTypeMap["name"] = bindings.NewStringType()
    pathParams["org"] = "org"
    pathParams["sddc_id"] = "sddcId"
    pathParams["addon_type"] = "addonType"
    pathParams["name"] = "name"
    resultHeaders := map[string]string{}
    errorHeaders := map[string]string{}
    return protocol.NewOperationRestMetadata(
      paramsTypeMap,
      pathParams,
      queryParams,
      headerParams,
      "credentials",
      "PUT",
      "/vmc/api/orgs/{org}/sddcs/{sddcId}/addons/{addonType}/credentials/{name}",
       resultHeaders,
       200,
       errorHeaders,
       map[string]int{"InvalidRequest": 400,"Unauthorized": 403})
}




