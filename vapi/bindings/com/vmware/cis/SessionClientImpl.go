
/* Copyright © 2019 VMware, Inc. All Rights Reserved.
     SPDX-License-Identifier: BSD-2-Clause */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Client stubs for service: Session
 * Functions that implement the generated SessionClient interface
 */


package cis

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
// //     "getDependenciesOfServiceTypes time"
// 
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/bindings/com/vmware/vapi/std/errors"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/bindings"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/core"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/data"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/lib"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/log"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/protocol"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/protocol/client"
    "time"
)


type SessionClientImpl struct{
      interfaceName string
      interfaceDefinition core.InterfaceDefinition
      methodIdentifiers[]core.MethodIdentifier
      methodNameToDefMap map[string]*core.MethodDefinition
      errorBindingMap map[string]bindings.BindingType
      interfaceIdentifier core.InterfaceIdentifier
      connector client.Connector
}


func NewSessionClientImpl(connector client.Connector) *SessionClientImpl {
      interfaceName := "com.vmware.cis.session"
      interfaceIdentifier := core.NewInterfaceIdentifier(interfaceName)
      methodIdentifiers := []core.MethodIdentifier{
          core.NewMethodIdentifier(interfaceIdentifier, "create"),
          core.NewMethodIdentifier(interfaceIdentifier, "delete"),
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
      sIface := SessionClientImpl{interfaceName: interfaceName, methodIdentifiers: methodIdentifiers, interfaceDefinition: interfaceDefinition, errorBindingMap: errorBindingMap, interfaceIdentifier:interfaceIdentifier, connector: connector}
      sIface.methodNameToDefMap = make(map[string]*core.MethodDefinition)
      sIface.methodNameToDefMap["create"] = sIface.createMethodDefinition()
      sIface.methodNameToDefMap["delete"] = sIface.deleteMethodDefinition()
      sIface.methodNameToDefMap["get"] = sIface.getMethodDefinition()
      return &sIface
}

func (sIface *SessionClientImpl) Create() (string, error) {
	typeConverter := sIface.connector.TypeConverter()
	methodIdentifier := core.NewMethodIdentifier(sIface.interfaceIdentifier, "create")
	sv := bindings.NewStructValueBuilder(sessionCreateInputType(), typeConverter)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
        var emptyOutput string
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := sessionCreateRestMetadata
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	sIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult:= sIface.Invoke(sIface.connector.NewExecutionContext(), methodIdentifier, inputDataValue)
	var emptyOutput string
    if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), sessionCreateOutputType())
		if errorInOutput != nil {
		    return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
	    return output.(string), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), sIface.errorBindingMap[methodResult.Error().Name()])
		if errorInError != nil {
		    return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}
func (sIface *SessionClientImpl) Delete() error {
	typeConverter := sIface.connector.TypeConverter()
	methodIdentifier := core.NewMethodIdentifier(sIface.interfaceIdentifier, "delete")
	sv := bindings.NewStructValueBuilder(sessionDeleteInputType(), typeConverter)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		return bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := sessionDeleteRestMetadata
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	sIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult:= sIface.Invoke(sIface.connector.NewExecutionContext(), methodIdentifier, inputDataValue)
    if methodResult.IsSuccess() {
		return nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), sIface.errorBindingMap[methodResult.Error().Name()])
		if errorInError != nil {
		    return bindings.VAPIerrorsToError(errorInError)
		}
		return methodError.(error)
	}
}
func (sIface *SessionClientImpl) Get() (Info, error) {
	typeConverter := sIface.connector.TypeConverter()
	methodIdentifier := core.NewMethodIdentifier(sIface.interfaceIdentifier, "get")
	sv := bindings.NewStructValueBuilder(sessionGetInputType(), typeConverter)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
        var emptyOutput Info
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := sessionGetRestMetadata
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	sIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult:= sIface.Invoke(sIface.connector.NewExecutionContext(), methodIdentifier, inputDataValue)
	var emptyOutput Info
    if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), sessionGetOutputType())
		if errorInOutput != nil {
		    return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
	    return output.(Info), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), sIface.errorBindingMap[methodResult.Error().Name()])
		if errorInError != nil {
		    return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (sIface *SessionClientImpl) Invoke(ctx *core.ExecutionContext, methodId core.MethodIdentifier, inputDataValue data.DataValue) core.MethodResult {
    methodResult := sIface.connector.GetApiProvider().Invoke(sIface.interfaceName, methodId.Name(), inputDataValue, ctx)
    return methodResult
}

func (sIface *SessionClientImpl) createMethodDefinition() *core.MethodDefinition {
      interfaceIdentifier := core.NewInterfaceIdentifier(sIface.interfaceName)
      typeConverter := sIface.connector.TypeConverter()

      input, inputError := typeConverter.ConvertToDataDefinition(sessionCreateInputType())
      output, outputError := typeConverter.ConvertToDataDefinition(sessionCreateOutputType())
      if(inputError != nil) {
          log.Errorf("Error in ConvertToDataDefinition for SessionClientImpl.create method's input - %s",
              bindings.VAPIerrorsToError(inputError).Error())
          return nil
      }
      if(outputError != nil) {
          log.Errorf("Error in ConvertToDataDefinition for SessionClientImpl.create method's output - %s",
              bindings.VAPIerrorsToError(outputError).Error())
          return nil
      }
      methodIdentifier := core.NewMethodIdentifier(interfaceIdentifier, "create")
      errorDefinitions := make([]data.ErrorDefinition,0)
      sIface.errorBindingMap[errors.Unauthenticated{}.Error()] = errors.UnauthenticatedBindingType()
	  errDef1, errError1 := typeConverter.ConvertToDataDefinition(errors.UnauthenticatedBindingType())
	  if(errError1 != nil) {
	    log.Errorf("Error in ConvertToDataDefinition for SessionClientImpl.create method's errors.Unauthenticated error - %s",
	        bindings.VAPIerrorsToError(errError1).Error())
        return nil
	  }
	  errorDefinitions = append(errorDefinitions, errDef1.(data.ErrorDefinition))
      sIface.errorBindingMap[errors.ServiceUnavailable{}.Error()] = errors.ServiceUnavailableBindingType()
	  errDef2, errError2 := typeConverter.ConvertToDataDefinition(errors.ServiceUnavailableBindingType())
	  if(errError2 != nil) {
	    log.Errorf("Error in ConvertToDataDefinition for SessionClientImpl.create method's errors.ServiceUnavailable error - %s",
	        bindings.VAPIerrorsToError(errError2).Error())
        return nil
	  }
	  errorDefinitions = append(errorDefinitions, errDef2.(data.ErrorDefinition))

      methodDefinition := core.NewMethodDefinition(methodIdentifier, input, output, errorDefinitions)
      return &methodDefinition
}
func (sIface *SessionClientImpl) deleteMethodDefinition() *core.MethodDefinition {
      interfaceIdentifier := core.NewInterfaceIdentifier(sIface.interfaceName)
      typeConverter := sIface.connector.TypeConverter()

      input, inputError := typeConverter.ConvertToDataDefinition(sessionDeleteInputType())
      output, outputError := typeConverter.ConvertToDataDefinition(sessionDeleteOutputType())
      if(inputError != nil) {
          log.Errorf("Error in ConvertToDataDefinition for SessionClientImpl.delete method's input - %s",
              bindings.VAPIerrorsToError(inputError).Error())
          return nil
      }
      if(outputError != nil) {
          log.Errorf("Error in ConvertToDataDefinition for SessionClientImpl.delete method's output - %s",
              bindings.VAPIerrorsToError(outputError).Error())
          return nil
      }
      methodIdentifier := core.NewMethodIdentifier(interfaceIdentifier, "delete")
      errorDefinitions := make([]data.ErrorDefinition,0)
      sIface.errorBindingMap[errors.Unauthenticated{}.Error()] = errors.UnauthenticatedBindingType()
	  errDef1, errError1 := typeConverter.ConvertToDataDefinition(errors.UnauthenticatedBindingType())
	  if(errError1 != nil) {
	    log.Errorf("Error in ConvertToDataDefinition for SessionClientImpl.delete method's errors.Unauthenticated error - %s",
	        bindings.VAPIerrorsToError(errError1).Error())
        return nil
	  }
	  errorDefinitions = append(errorDefinitions, errDef1.(data.ErrorDefinition))
      sIface.errorBindingMap[errors.ServiceUnavailable{}.Error()] = errors.ServiceUnavailableBindingType()
	  errDef2, errError2 := typeConverter.ConvertToDataDefinition(errors.ServiceUnavailableBindingType())
	  if(errError2 != nil) {
	    log.Errorf("Error in ConvertToDataDefinition for SessionClientImpl.delete method's errors.ServiceUnavailable error - %s",
	        bindings.VAPIerrorsToError(errError2).Error())
        return nil
	  }
	  errorDefinitions = append(errorDefinitions, errDef2.(data.ErrorDefinition))

      methodDefinition := core.NewMethodDefinition(methodIdentifier, input, output, errorDefinitions)
      return &methodDefinition
}
func (sIface *SessionClientImpl) getMethodDefinition() *core.MethodDefinition {
      interfaceIdentifier := core.NewInterfaceIdentifier(sIface.interfaceName)
      typeConverter := sIface.connector.TypeConverter()

      input, inputError := typeConverter.ConvertToDataDefinition(sessionGetInputType())
      output, outputError := typeConverter.ConvertToDataDefinition(sessionGetOutputType())
      if(inputError != nil) {
          log.Errorf("Error in ConvertToDataDefinition for SessionClientImpl.get method's input - %s",
              bindings.VAPIerrorsToError(inputError).Error())
          return nil
      }
      if(outputError != nil) {
          log.Errorf("Error in ConvertToDataDefinition for SessionClientImpl.get method's output - %s",
              bindings.VAPIerrorsToError(outputError).Error())
          return nil
      }
      methodIdentifier := core.NewMethodIdentifier(interfaceIdentifier, "get")
      errorDefinitions := make([]data.ErrorDefinition,0)
      sIface.errorBindingMap[errors.Unauthenticated{}.Error()] = errors.UnauthenticatedBindingType()
	  errDef1, errError1 := typeConverter.ConvertToDataDefinition(errors.UnauthenticatedBindingType())
	  if(errError1 != nil) {
	    log.Errorf("Error in ConvertToDataDefinition for SessionClientImpl.get method's errors.Unauthenticated error - %s",
	        bindings.VAPIerrorsToError(errError1).Error())
        return nil
	  }
	  errorDefinitions = append(errorDefinitions, errDef1.(data.ErrorDefinition))
      sIface.errorBindingMap[errors.ServiceUnavailable{}.Error()] = errors.ServiceUnavailableBindingType()
	  errDef2, errError2 := typeConverter.ConvertToDataDefinition(errors.ServiceUnavailableBindingType())
	  if(errError2 != nil) {
	    log.Errorf("Error in ConvertToDataDefinition for SessionClientImpl.get method's errors.ServiceUnavailable error - %s",
	        bindings.VAPIerrorsToError(errError2).Error())
        return nil
	  }
	  errorDefinitions = append(errorDefinitions, errDef2.(data.ErrorDefinition))

      methodDefinition := core.NewMethodDefinition(methodIdentifier, input, output, errorDefinitions)
      return &methodDefinition
}




// Represents data associated with an API session.
 type Info struct {
    // Fully qualified name of the end user that created the session, for example Administrator\\\\@vsphere.local. A typical use case for this information is in Graphical User Interfaces (GUI) or logging systems to visualize the identity of the current user.
    User string
    // Time when the session was created.
    CreatedTime time.Time
    // Last time this session was used by passing the session key for invoking an API.
    LastAccessedTime time.Time
}









func sessionCreateInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func sessionCreateOutputType() bindings.BindingType {
    return bindings.NewSecretType()
}

func sessionCreateRestMetadata() protocol.OperationRestMetadata {
    paramsTypeMap := map[string]bindings.BindingType{}
    pathParams := map[string]string{}
    queryParams := map[string]string{}
    headerParams := map[string]string{}
    resultHeaders := map[string]string{}
    errorHeaders := map[string]string{}
    errorHeaders["Unauthenticated.challenge"] = "WWW-Authenticate"
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
       map[string]int{"Unauthenticated": 401,"ServiceUnavailable": 503})
}


func sessionDeleteInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func sessionDeleteOutputType() bindings.BindingType {
    return bindings.NewVoidType()
}

func sessionDeleteRestMetadata() protocol.OperationRestMetadata {
    paramsTypeMap := map[string]bindings.BindingType{}
    pathParams := map[string]string{}
    queryParams := map[string]string{}
    headerParams := map[string]string{}
    resultHeaders := map[string]string{}
    errorHeaders := map[string]string{}
    errorHeaders["Unauthenticated.challenge"] = "WWW-Authenticate"
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
       map[string]int{"Unauthenticated": 401,"ServiceUnavailable": 503})
}


func sessionGetInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func sessionGetOutputType() bindings.BindingType {
    return bindings.NewReferenceType(InfoBindingType)
}

func sessionGetRestMetadata() protocol.OperationRestMetadata {
    paramsTypeMap := map[string]bindings.BindingType{}
    pathParams := map[string]string{}
    queryParams := map[string]string{}
    headerParams := map[string]string{}
    resultHeaders := map[string]string{}
    errorHeaders := map[string]string{}
    errorHeaders["Unauthenticated.challenge"] = "WWW-Authenticate"
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
       map[string]int{"Unauthenticated": 401,"ServiceUnavailable": 503})
}



func InfoBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["user"] = bindings.NewStringType()
    fieldNameMap["user"] = "User"
    fields["created_time"] = bindings.NewDateTimeType()
    fieldNameMap["created_time"] = "CreatedTime"
    fields["last_accessed_time"] = bindings.NewDateTimeType()
    fieldNameMap["last_accessed_time"] = "LastAccessedTime"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.cis.session.info",fields, reflect.TypeOf(Info{}), fieldNameMap, validators)
}


