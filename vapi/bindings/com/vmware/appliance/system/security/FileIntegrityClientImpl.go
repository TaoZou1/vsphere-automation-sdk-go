
/* Copyright © 2019 VMware, Inc. All Rights Reserved.
     SPDX-License-Identifier: BSD-2-Clause */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Client stubs for service: FileIntegrity
 * Functions that implement the generated FileIntegrityClient interface
 */


package security
import (
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/bindings/com/vmware/vapi/std/errors"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/bindings"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/core"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/data"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/lib"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/log"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/protocol/client"
)


type FileIntegrityClientImpl struct{
      interfaceName string
      interfaceDefinition core.InterfaceDefinition
      methodIdentifiers[]core.MethodIdentifier
      methodNameToDefMap map[string]*core.MethodDefinition
      errorBindingMap map[string]bindings.BindingType
      interfaceIdentifier core.InterfaceIdentifier
      connector client.Connector
}


func NewFileIntegrityClientImpl(connector client.Connector) *FileIntegrityClientImpl {
      interfaceName := "com.vmware.appliance.system.security.file_integrity"
      interfaceIdentifier := core.NewInterfaceIdentifier(interfaceName)
      methodIdentifiers := []core.MethodIdentifier{
          core.NewMethodIdentifier(interfaceIdentifier, "baseline"),
          core.NewMethodIdentifier(interfaceIdentifier, "get"),
          core.NewMethodIdentifier(interfaceIdentifier, "set"),
      }
      interfaceDefinition := core.NewInterfaceDefinition(interfaceIdentifier, methodIdentifiers)
      errorBindingMap := make(map[string]bindings.BindingType)
      errorBindingMap[errors.InternalServerError{}.Error()] = errors.InternalServerErrorBindingType()
	  errorBindingMap[errors.InvalidArgument{}.Error()] = errors.InvalidArgumentBindingType()
	  errorBindingMap[errors.OperationNotFound{}.Error()] = errors.OperationNotFoundBindingType()
	  errorBindingMap[errors.UnexpectedInput{}.Error()] = errors.UnexpectedInputBindingType()
	  errorBindingMap[errors.ServiceUnavailable{}.Error()] = errors.ServiceUnavailableBindingType()
	  errorBindingMap[errors.TimedOut{}.Error()] = errors.TimedOutBindingType()
      fIface := FileIntegrityClientImpl{interfaceName: interfaceName, methodIdentifiers: methodIdentifiers, interfaceDefinition: interfaceDefinition, errorBindingMap: errorBindingMap, interfaceIdentifier:interfaceIdentifier, connector: connector}
      fIface.methodNameToDefMap = make(map[string]*core.MethodDefinition)
      fIface.methodNameToDefMap["baseline"] = fIface.baselineMethodDefinition()
      fIface.methodNameToDefMap["get"] = fIface.getMethodDefinition()
      fIface.methodNameToDefMap["set"] = fIface.setMethodDefinition()
      return &fIface
}

func (fIface *FileIntegrityClientImpl) Baseline() error {
	typeConverter := fIface.connector.TypeConverter()
	methodIdentifier := core.NewMethodIdentifier(fIface.interfaceIdentifier, "baseline")
	sv := bindings.NewStructValueBuilder(fileIntegrityBaselineInputType(), typeConverter)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		return bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := fileIntegrityBaselineRestMetadata
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	fIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult:= fIface.Invoke(fIface.connector.NewExecutionContext(), methodIdentifier, inputDataValue)
    if methodResult.IsSuccess() {
		return nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), fIface.errorBindingMap[methodResult.Error().Name()])
		if errorInError != nil {
		    return bindings.VAPIerrorsToError(errorInError)
		}
		return methodError.(error)
	}
}
func (fIface *FileIntegrityClientImpl) Get() (bool, error) {
	typeConverter := fIface.connector.TypeConverter()
	methodIdentifier := core.NewMethodIdentifier(fIface.interfaceIdentifier, "get")
	sv := bindings.NewStructValueBuilder(fileIntegrityGetInputType(), typeConverter)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
        var emptyOutput bool
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := fileIntegrityGetRestMetadata
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	fIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult:= fIface.Invoke(fIface.connector.NewExecutionContext(), methodIdentifier, inputDataValue)
	var emptyOutput bool
    if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), fileIntegrityGetOutputType())
		if errorInOutput != nil {
		    return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
	    return output.(bool), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), fIface.errorBindingMap[methodResult.Error().Name()])
		if errorInError != nil {
		    return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}
func (fIface *FileIntegrityClientImpl) Set(enabledParam bool) error {
	typeConverter := fIface.connector.TypeConverter()
	methodIdentifier := core.NewMethodIdentifier(fIface.interfaceIdentifier, "set")
	sv := bindings.NewStructValueBuilder(fileIntegritySetInputType(), typeConverter)
	sv.AddStructField("Enabled", enabledParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		return bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := fileIntegritySetRestMetadata
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	fIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult:= fIface.Invoke(fIface.connector.NewExecutionContext(), methodIdentifier, inputDataValue)
    if methodResult.IsSuccess() {
		return nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), fIface.errorBindingMap[methodResult.Error().Name()])
		if errorInError != nil {
		    return bindings.VAPIerrorsToError(errorInError)
		}
		return methodError.(error)
	}
}

func (fIface *FileIntegrityClientImpl) Invoke(ctx *core.ExecutionContext, methodId core.MethodIdentifier, inputDataValue data.DataValue) core.MethodResult {
    methodResult := fIface.connector.GetApiProvider().Invoke(fIface.interfaceName, methodId.Name(), inputDataValue, ctx)
    return methodResult
}

func (fIface *FileIntegrityClientImpl) baselineMethodDefinition() *core.MethodDefinition {
      interfaceIdentifier := core.NewInterfaceIdentifier(fIface.interfaceName)
      typeConverter := fIface.connector.TypeConverter()

      input, inputError := typeConverter.ConvertToDataDefinition(fileIntegrityBaselineInputType())
      output, outputError := typeConverter.ConvertToDataDefinition(fileIntegrityBaselineOutputType())
      if(inputError != nil) {
          log.Errorf("Error in ConvertToDataDefinition for FileIntegrityClientImpl.baseline method's input - %s",
              bindings.VAPIerrorsToError(inputError).Error())
          return nil
      }
      if(outputError != nil) {
          log.Errorf("Error in ConvertToDataDefinition for FileIntegrityClientImpl.baseline method's output - %s",
              bindings.VAPIerrorsToError(outputError).Error())
          return nil
      }
      methodIdentifier := core.NewMethodIdentifier(interfaceIdentifier, "baseline")
      errorDefinitions := make([]data.ErrorDefinition,0)
      fIface.errorBindingMap[errors.Error{}.Error()] = errors.ErrorBindingType()
	  errDef1, errError1 := typeConverter.ConvertToDataDefinition(errors.ErrorBindingType())
	  if(errError1 != nil) {
	    log.Errorf("Error in ConvertToDataDefinition for FileIntegrityClientImpl.baseline method's errors.Error error - %s",
	        bindings.VAPIerrorsToError(errError1).Error())
        return nil
	  }
	  errorDefinitions = append(errorDefinitions, errDef1.(data.ErrorDefinition))

      methodDefinition := core.NewMethodDefinition(methodIdentifier, input, output, errorDefinitions)
      return &methodDefinition
}
func (fIface *FileIntegrityClientImpl) getMethodDefinition() *core.MethodDefinition {
      interfaceIdentifier := core.NewInterfaceIdentifier(fIface.interfaceName)
      typeConverter := fIface.connector.TypeConverter()

      input, inputError := typeConverter.ConvertToDataDefinition(fileIntegrityGetInputType())
      output, outputError := typeConverter.ConvertToDataDefinition(fileIntegrityGetOutputType())
      if(inputError != nil) {
          log.Errorf("Error in ConvertToDataDefinition for FileIntegrityClientImpl.get method's input - %s",
              bindings.VAPIerrorsToError(inputError).Error())
          return nil
      }
      if(outputError != nil) {
          log.Errorf("Error in ConvertToDataDefinition for FileIntegrityClientImpl.get method's output - %s",
              bindings.VAPIerrorsToError(outputError).Error())
          return nil
      }
      methodIdentifier := core.NewMethodIdentifier(interfaceIdentifier, "get")
      errorDefinitions := make([]data.ErrorDefinition,0)
      fIface.errorBindingMap[errors.Error{}.Error()] = errors.ErrorBindingType()
	  errDef1, errError1 := typeConverter.ConvertToDataDefinition(errors.ErrorBindingType())
	  if(errError1 != nil) {
	    log.Errorf("Error in ConvertToDataDefinition for FileIntegrityClientImpl.get method's errors.Error error - %s",
	        bindings.VAPIerrorsToError(errError1).Error())
        return nil
	  }
	  errorDefinitions = append(errorDefinitions, errDef1.(data.ErrorDefinition))

      methodDefinition := core.NewMethodDefinition(methodIdentifier, input, output, errorDefinitions)
      return &methodDefinition
}
func (fIface *FileIntegrityClientImpl) setMethodDefinition() *core.MethodDefinition {
      interfaceIdentifier := core.NewInterfaceIdentifier(fIface.interfaceName)
      typeConverter := fIface.connector.TypeConverter()

      input, inputError := typeConverter.ConvertToDataDefinition(fileIntegritySetInputType())
      output, outputError := typeConverter.ConvertToDataDefinition(fileIntegritySetOutputType())
      if(inputError != nil) {
          log.Errorf("Error in ConvertToDataDefinition for FileIntegrityClientImpl.set method's input - %s",
              bindings.VAPIerrorsToError(inputError).Error())
          return nil
      }
      if(outputError != nil) {
          log.Errorf("Error in ConvertToDataDefinition for FileIntegrityClientImpl.set method's output - %s",
              bindings.VAPIerrorsToError(outputError).Error())
          return nil
      }
      methodIdentifier := core.NewMethodIdentifier(interfaceIdentifier, "set")
      errorDefinitions := make([]data.ErrorDefinition,0)
      fIface.errorBindingMap[errors.Error{}.Error()] = errors.ErrorBindingType()
	  errDef1, errError1 := typeConverter.ConvertToDataDefinition(errors.ErrorBindingType())
	  if(errError1 != nil) {
	    log.Errorf("Error in ConvertToDataDefinition for FileIntegrityClientImpl.set method's errors.Error error - %s",
	        bindings.VAPIerrorsToError(errError1).Error())
        return nil
	  }
	  errorDefinitions = append(errorDefinitions, errDef1.(data.ErrorDefinition))

      methodDefinition := core.NewMethodDefinition(methodIdentifier, input, output, errorDefinitions)
      return &methodDefinition
}


