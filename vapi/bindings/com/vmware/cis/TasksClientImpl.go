
/* Copyright © 2019 VMware, Inc. All Rights Reserved.
     SPDX-License-Identifier: BSD-2-Clause */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Client stubs for service: Tasks
 * Functions that implement the generated TasksClient interface
 */


package cis

import (
    "reflect"
// //     "getClientImplDependenciesOfOps gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/bindings/com/vmware/cis/task"
// //     "getClientImplDependenciesOfOps gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/bindings/com/vmware/vapi/std/errors"
// //     "getClientImplDependenciesOfOps gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/bindings"
// //     "getClientImplDependenciesOfOps gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/core"
// //     "getClientImplDependenciesOfOps gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/data"
// //     "getClientImplDependenciesOfOps gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/lib"
// //     "getClientImplDependenciesOfOps gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/log"
// //     "getClientImplDependenciesOfOps gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/protocol/client"
// //
// //     "getDependenciesOfServiceTypes gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/bindings/com/vmware/cis/task"
// //     "getDependenciesOfServiceTypes gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/bindings/com/vmware/vapi/std"
// //     "getDependenciesOfServiceTypes gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/bindings"
// //     "getDependenciesOfServiceTypes gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/data"
// //     "getDependenciesOfServiceTypes gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/protocol"
// 
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/bindings/com/vmware/cis/task"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/bindings/com/vmware/vapi/std"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/bindings/com/vmware/vapi/std/errors"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/bindings"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/core"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/data"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/lib"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/log"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/protocol"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/protocol/client"
)


type TasksClientImpl struct{
      interfaceName string
      interfaceDefinition core.InterfaceDefinition
      methodIdentifiers[]core.MethodIdentifier
      methodNameToDefMap map[string]*core.MethodDefinition
      errorBindingMap map[string]bindings.BindingType
      interfaceIdentifier core.InterfaceIdentifier
      connector client.Connector
}


func NewTasksClientImpl(connector client.Connector) *TasksClientImpl {
      interfaceName := "com.vmware.cis.tasks"
      interfaceIdentifier := core.NewInterfaceIdentifier(interfaceName)
      methodIdentifiers := []core.MethodIdentifier{
          core.NewMethodIdentifier(interfaceIdentifier, "get"),
          core.NewMethodIdentifier(interfaceIdentifier, "list"),
          core.NewMethodIdentifier(interfaceIdentifier, "cancel"),
      }
      interfaceDefinition := core.NewInterfaceDefinition(interfaceIdentifier, methodIdentifiers)
      errorBindingMap := make(map[string]bindings.BindingType)
      errorBindingMap[errors.InternalServerError{}.Error()] = errors.InternalServerErrorBindingType()
	  errorBindingMap[errors.InvalidArgument{}.Error()] = errors.InvalidArgumentBindingType()
	  errorBindingMap[errors.OperationNotFound{}.Error()] = errors.OperationNotFoundBindingType()
	  errorBindingMap[errors.UnexpectedInput{}.Error()] = errors.UnexpectedInputBindingType()
	  errorBindingMap[errors.ServiceUnavailable{}.Error()] = errors.ServiceUnavailableBindingType()
	  errorBindingMap[errors.TimedOut{}.Error()] = errors.TimedOutBindingType()
      tIface := TasksClientImpl{interfaceName: interfaceName, methodIdentifiers: methodIdentifiers, interfaceDefinition: interfaceDefinition, errorBindingMap: errorBindingMap, interfaceIdentifier:interfaceIdentifier, connector: connector}
      tIface.methodNameToDefMap = make(map[string]*core.MethodDefinition)
      tIface.methodNameToDefMap["get"] = tIface.getMethodDefinition()
      tIface.methodNameToDefMap["list"] = tIface.listMethodDefinition()
      tIface.methodNameToDefMap["cancel"] = tIface.cancelMethodDefinition()
      return &tIface
}

func (tIface *TasksClientImpl) Get(taskParam string, specParam *GetSpec) (task.Info, error) {
	typeConverter := tIface.connector.TypeConverter()
	methodIdentifier := core.NewMethodIdentifier(tIface.interfaceIdentifier, "get")
	sv := bindings.NewStructValueBuilder(tasksGetInputType(), typeConverter)
	sv.AddStructField("Task", taskParam)
	sv.AddStructField("Spec", specParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
        var emptyOutput task.Info
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := tasksGetRestMetadata
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	tIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult:= tIface.Invoke(tIface.connector.NewExecutionContext(), methodIdentifier, inputDataValue)
	var emptyOutput task.Info
    if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), tasksGetOutputType())
		if errorInOutput != nil {
		    return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
	    return output.(task.Info), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), tIface.errorBindingMap[methodResult.Error().Name()])
		if errorInError != nil {
		    return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}
func (tIface *TasksClientImpl) List(filterSpecParam *FilterSpec, resultSpecParam *GetSpec) (map[string]task.Info, error) {
	typeConverter := tIface.connector.TypeConverter()
	methodIdentifier := core.NewMethodIdentifier(tIface.interfaceIdentifier, "list")
	sv := bindings.NewStructValueBuilder(tasksListInputType(), typeConverter)
	sv.AddStructField("FilterSpec", filterSpecParam)
	sv.AddStructField("ResultSpec", resultSpecParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
        var emptyOutput map[string]task.Info
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := tasksListRestMetadata
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	tIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult:= tIface.Invoke(tIface.connector.NewExecutionContext(), methodIdentifier, inputDataValue)
	var emptyOutput map[string]task.Info
    if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), tasksListOutputType())
		if errorInOutput != nil {
		    return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
	    return output.(map[string]task.Info), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), tIface.errorBindingMap[methodResult.Error().Name()])
		if errorInError != nil {
		    return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}
func (tIface *TasksClientImpl) Cancel(taskParam string) error {
	typeConverter := tIface.connector.TypeConverter()
	methodIdentifier := core.NewMethodIdentifier(tIface.interfaceIdentifier, "cancel")
	sv := bindings.NewStructValueBuilder(tasksCancelInputType(), typeConverter)
	sv.AddStructField("Task", taskParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		return bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := tasksCancelRestMetadata
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	tIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult:= tIface.Invoke(tIface.connector.NewExecutionContext(), methodIdentifier, inputDataValue)
    if methodResult.IsSuccess() {
		return nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), tIface.errorBindingMap[methodResult.Error().Name()])
		if errorInError != nil {
		    return bindings.VAPIerrorsToError(errorInError)
		}
		return methodError.(error)
	}
}

func (tIface *TasksClientImpl) Invoke(ctx *core.ExecutionContext, methodId core.MethodIdentifier, inputDataValue data.DataValue) core.MethodResult {
    methodResult := tIface.connector.GetApiProvider().Invoke(tIface.interfaceName, methodId.Name(), inputDataValue, ctx)
    return methodResult
}

func (tIface *TasksClientImpl) getMethodDefinition() *core.MethodDefinition {
      interfaceIdentifier := core.NewInterfaceIdentifier(tIface.interfaceName)
      typeConverter := tIface.connector.TypeConverter()

      input, inputError := typeConverter.ConvertToDataDefinition(tasksGetInputType())
      output, outputError := typeConverter.ConvertToDataDefinition(tasksGetOutputType())
      if(inputError != nil) {
          log.Errorf("Error in ConvertToDataDefinition for TasksClientImpl.get method's input - %s",
              bindings.VAPIerrorsToError(inputError).Error())
          return nil
      }
      if(outputError != nil) {
          log.Errorf("Error in ConvertToDataDefinition for TasksClientImpl.get method's output - %s",
              bindings.VAPIerrorsToError(outputError).Error())
          return nil
      }
      methodIdentifier := core.NewMethodIdentifier(interfaceIdentifier, "get")
      errorDefinitions := make([]data.ErrorDefinition,0)
      tIface.errorBindingMap[errors.Error{}.Error()] = errors.ErrorBindingType()
	  errDef1, errError1 := typeConverter.ConvertToDataDefinition(errors.ErrorBindingType())
	  if(errError1 != nil) {
	    log.Errorf("Error in ConvertToDataDefinition for TasksClientImpl.get method's errors.Error error - %s",
	        bindings.VAPIerrorsToError(errError1).Error())
        return nil
	  }
	  errorDefinitions = append(errorDefinitions, errDef1.(data.ErrorDefinition))
      tIface.errorBindingMap[errors.NotFound{}.Error()] = errors.NotFoundBindingType()
	  errDef2, errError2 := typeConverter.ConvertToDataDefinition(errors.NotFoundBindingType())
	  if(errError2 != nil) {
	    log.Errorf("Error in ConvertToDataDefinition for TasksClientImpl.get method's errors.NotFound error - %s",
	        bindings.VAPIerrorsToError(errError2).Error())
        return nil
	  }
	  errorDefinitions = append(errorDefinitions, errDef2.(data.ErrorDefinition))
      tIface.errorBindingMap[errors.ResourceInaccessible{}.Error()] = errors.ResourceInaccessibleBindingType()
	  errDef3, errError3 := typeConverter.ConvertToDataDefinition(errors.ResourceInaccessibleBindingType())
	  if(errError3 != nil) {
	    log.Errorf("Error in ConvertToDataDefinition for TasksClientImpl.get method's errors.ResourceInaccessible error - %s",
	        bindings.VAPIerrorsToError(errError3).Error())
        return nil
	  }
	  errorDefinitions = append(errorDefinitions, errDef3.(data.ErrorDefinition))
      tIface.errorBindingMap[errors.ServiceUnavailable{}.Error()] = errors.ServiceUnavailableBindingType()
	  errDef4, errError4 := typeConverter.ConvertToDataDefinition(errors.ServiceUnavailableBindingType())
	  if(errError4 != nil) {
	    log.Errorf("Error in ConvertToDataDefinition for TasksClientImpl.get method's errors.ServiceUnavailable error - %s",
	        bindings.VAPIerrorsToError(errError4).Error())
        return nil
	  }
	  errorDefinitions = append(errorDefinitions, errDef4.(data.ErrorDefinition))
      tIface.errorBindingMap[errors.Unauthenticated{}.Error()] = errors.UnauthenticatedBindingType()
	  errDef5, errError5 := typeConverter.ConvertToDataDefinition(errors.UnauthenticatedBindingType())
	  if(errError5 != nil) {
	    log.Errorf("Error in ConvertToDataDefinition for TasksClientImpl.get method's errors.Unauthenticated error - %s",
	        bindings.VAPIerrorsToError(errError5).Error())
        return nil
	  }
	  errorDefinitions = append(errorDefinitions, errDef5.(data.ErrorDefinition))
      tIface.errorBindingMap[errors.Unauthorized{}.Error()] = errors.UnauthorizedBindingType()
	  errDef6, errError6 := typeConverter.ConvertToDataDefinition(errors.UnauthorizedBindingType())
	  if(errError6 != nil) {
	    log.Errorf("Error in ConvertToDataDefinition for TasksClientImpl.get method's errors.Unauthorized error - %s",
	        bindings.VAPIerrorsToError(errError6).Error())
        return nil
	  }
	  errorDefinitions = append(errorDefinitions, errDef6.(data.ErrorDefinition))

      methodDefinition := core.NewMethodDefinition(methodIdentifier, input, output, errorDefinitions)
      return &methodDefinition
}
func (tIface *TasksClientImpl) listMethodDefinition() *core.MethodDefinition {
      interfaceIdentifier := core.NewInterfaceIdentifier(tIface.interfaceName)
      typeConverter := tIface.connector.TypeConverter()

      input, inputError := typeConverter.ConvertToDataDefinition(tasksListInputType())
      output, outputError := typeConverter.ConvertToDataDefinition(tasksListOutputType())
      if(inputError != nil) {
          log.Errorf("Error in ConvertToDataDefinition for TasksClientImpl.list method's input - %s",
              bindings.VAPIerrorsToError(inputError).Error())
          return nil
      }
      if(outputError != nil) {
          log.Errorf("Error in ConvertToDataDefinition for TasksClientImpl.list method's output - %s",
              bindings.VAPIerrorsToError(outputError).Error())
          return nil
      }
      methodIdentifier := core.NewMethodIdentifier(interfaceIdentifier, "list")
      errorDefinitions := make([]data.ErrorDefinition,0)
      tIface.errorBindingMap[errors.InvalidArgument{}.Error()] = errors.InvalidArgumentBindingType()
	  errDef1, errError1 := typeConverter.ConvertToDataDefinition(errors.InvalidArgumentBindingType())
	  if(errError1 != nil) {
	    log.Errorf("Error in ConvertToDataDefinition for TasksClientImpl.list method's errors.InvalidArgument error - %s",
	        bindings.VAPIerrorsToError(errError1).Error())
        return nil
	  }
	  errorDefinitions = append(errorDefinitions, errDef1.(data.ErrorDefinition))
      tIface.errorBindingMap[errors.ResourceInaccessible{}.Error()] = errors.ResourceInaccessibleBindingType()
	  errDef2, errError2 := typeConverter.ConvertToDataDefinition(errors.ResourceInaccessibleBindingType())
	  if(errError2 != nil) {
	    log.Errorf("Error in ConvertToDataDefinition for TasksClientImpl.list method's errors.ResourceInaccessible error - %s",
	        bindings.VAPIerrorsToError(errError2).Error())
        return nil
	  }
	  errorDefinitions = append(errorDefinitions, errDef2.(data.ErrorDefinition))
      tIface.errorBindingMap[errors.ServiceUnavailable{}.Error()] = errors.ServiceUnavailableBindingType()
	  errDef3, errError3 := typeConverter.ConvertToDataDefinition(errors.ServiceUnavailableBindingType())
	  if(errError3 != nil) {
	    log.Errorf("Error in ConvertToDataDefinition for TasksClientImpl.list method's errors.ServiceUnavailable error - %s",
	        bindings.VAPIerrorsToError(errError3).Error())
        return nil
	  }
	  errorDefinitions = append(errorDefinitions, errDef3.(data.ErrorDefinition))
      tIface.errorBindingMap[errors.Unauthenticated{}.Error()] = errors.UnauthenticatedBindingType()
	  errDef4, errError4 := typeConverter.ConvertToDataDefinition(errors.UnauthenticatedBindingType())
	  if(errError4 != nil) {
	    log.Errorf("Error in ConvertToDataDefinition for TasksClientImpl.list method's errors.Unauthenticated error - %s",
	        bindings.VAPIerrorsToError(errError4).Error())
        return nil
	  }
	  errorDefinitions = append(errorDefinitions, errDef4.(data.ErrorDefinition))
      tIface.errorBindingMap[errors.Unauthorized{}.Error()] = errors.UnauthorizedBindingType()
	  errDef5, errError5 := typeConverter.ConvertToDataDefinition(errors.UnauthorizedBindingType())
	  if(errError5 != nil) {
	    log.Errorf("Error in ConvertToDataDefinition for TasksClientImpl.list method's errors.Unauthorized error - %s",
	        bindings.VAPIerrorsToError(errError5).Error())
        return nil
	  }
	  errorDefinitions = append(errorDefinitions, errDef5.(data.ErrorDefinition))

      methodDefinition := core.NewMethodDefinition(methodIdentifier, input, output, errorDefinitions)
      return &methodDefinition
}
func (tIface *TasksClientImpl) cancelMethodDefinition() *core.MethodDefinition {
      interfaceIdentifier := core.NewInterfaceIdentifier(tIface.interfaceName)
      typeConverter := tIface.connector.TypeConverter()

      input, inputError := typeConverter.ConvertToDataDefinition(tasksCancelInputType())
      output, outputError := typeConverter.ConvertToDataDefinition(tasksCancelOutputType())
      if(inputError != nil) {
          log.Errorf("Error in ConvertToDataDefinition for TasksClientImpl.cancel method's input - %s",
              bindings.VAPIerrorsToError(inputError).Error())
          return nil
      }
      if(outputError != nil) {
          log.Errorf("Error in ConvertToDataDefinition for TasksClientImpl.cancel method's output - %s",
              bindings.VAPIerrorsToError(outputError).Error())
          return nil
      }
      methodIdentifier := core.NewMethodIdentifier(interfaceIdentifier, "cancel")
      errorDefinitions := make([]data.ErrorDefinition,0)
      tIface.errorBindingMap[errors.Error{}.Error()] = errors.ErrorBindingType()
	  errDef1, errError1 := typeConverter.ConvertToDataDefinition(errors.ErrorBindingType())
	  if(errError1 != nil) {
	    log.Errorf("Error in ConvertToDataDefinition for TasksClientImpl.cancel method's errors.Error error - %s",
	        bindings.VAPIerrorsToError(errError1).Error())
        return nil
	  }
	  errorDefinitions = append(errorDefinitions, errDef1.(data.ErrorDefinition))
      tIface.errorBindingMap[errors.NotAllowedInCurrentState{}.Error()] = errors.NotAllowedInCurrentStateBindingType()
	  errDef2, errError2 := typeConverter.ConvertToDataDefinition(errors.NotAllowedInCurrentStateBindingType())
	  if(errError2 != nil) {
	    log.Errorf("Error in ConvertToDataDefinition for TasksClientImpl.cancel method's errors.NotAllowedInCurrentState error - %s",
	        bindings.VAPIerrorsToError(errError2).Error())
        return nil
	  }
	  errorDefinitions = append(errorDefinitions, errDef2.(data.ErrorDefinition))
      tIface.errorBindingMap[errors.NotFound{}.Error()] = errors.NotFoundBindingType()
	  errDef3, errError3 := typeConverter.ConvertToDataDefinition(errors.NotFoundBindingType())
	  if(errError3 != nil) {
	    log.Errorf("Error in ConvertToDataDefinition for TasksClientImpl.cancel method's errors.NotFound error - %s",
	        bindings.VAPIerrorsToError(errError3).Error())
        return nil
	  }
	  errorDefinitions = append(errorDefinitions, errDef3.(data.ErrorDefinition))
      tIface.errorBindingMap[errors.ResourceInaccessible{}.Error()] = errors.ResourceInaccessibleBindingType()
	  errDef4, errError4 := typeConverter.ConvertToDataDefinition(errors.ResourceInaccessibleBindingType())
	  if(errError4 != nil) {
	    log.Errorf("Error in ConvertToDataDefinition for TasksClientImpl.cancel method's errors.ResourceInaccessible error - %s",
	        bindings.VAPIerrorsToError(errError4).Error())
        return nil
	  }
	  errorDefinitions = append(errorDefinitions, errDef4.(data.ErrorDefinition))
      tIface.errorBindingMap[errors.ServiceUnavailable{}.Error()] = errors.ServiceUnavailableBindingType()
	  errDef5, errError5 := typeConverter.ConvertToDataDefinition(errors.ServiceUnavailableBindingType())
	  if(errError5 != nil) {
	    log.Errorf("Error in ConvertToDataDefinition for TasksClientImpl.cancel method's errors.ServiceUnavailable error - %s",
	        bindings.VAPIerrorsToError(errError5).Error())
        return nil
	  }
	  errorDefinitions = append(errorDefinitions, errDef5.(data.ErrorDefinition))
      tIface.errorBindingMap[errors.Unauthenticated{}.Error()] = errors.UnauthenticatedBindingType()
	  errDef6, errError6 := typeConverter.ConvertToDataDefinition(errors.UnauthenticatedBindingType())
	  if(errError6 != nil) {
	    log.Errorf("Error in ConvertToDataDefinition for TasksClientImpl.cancel method's errors.Unauthenticated error - %s",
	        bindings.VAPIerrorsToError(errError6).Error())
        return nil
	  }
	  errorDefinitions = append(errorDefinitions, errDef6.(data.ErrorDefinition))
      tIface.errorBindingMap[errors.Unauthorized{}.Error()] = errors.UnauthorizedBindingType()
	  errDef7, errError7 := typeConverter.ConvertToDataDefinition(errors.UnauthorizedBindingType())
	  if(errError7 != nil) {
	    log.Errorf("Error in ConvertToDataDefinition for TasksClientImpl.cancel method's errors.Unauthorized error - %s",
	        bindings.VAPIerrorsToError(errError7).Error())
        return nil
	  }
	  errorDefinitions = append(errorDefinitions, errDef7.(data.ErrorDefinition))
      tIface.errorBindingMap[errors.Unsupported{}.Error()] = errors.UnsupportedBindingType()
	  errDef8, errError8 := typeConverter.ConvertToDataDefinition(errors.UnsupportedBindingType())
	  if(errError8 != nil) {
	    log.Errorf("Error in ConvertToDataDefinition for TasksClientImpl.cancel method's errors.Unsupported error - %s",
	        bindings.VAPIerrorsToError(errError8).Error())
        return nil
	  }
	  errorDefinitions = append(errorDefinitions, errDef8.(data.ErrorDefinition))

      methodDefinition := core.NewMethodDefinition(methodIdentifier, input, output, errorDefinitions)
      return &methodDefinition
}

// Resource type for task.
const Tasks_RESOURCE_TYPE = "com.vmware.cis.task"



// The ``GetSpec`` class describes what data should be included when retrieving information about a task.
 type GetSpec struct {
    // If true, all data, including operation-specific data, will be returned, otherwise only the data described in task.Info will be returned.
    ReturnAll *bool
    // If true, the result will not be included in the task information, otherwise it will be included.
    ExcludeResult *bool
}






// The ``FilterSpec`` class contains properties used to filter the results when listing tasks (see Tasks#list). If multiple properties are specified, only tasks matching all of the properties match the filter. 
//
//  Currently at least one of FilterSpec#tasks or FilterSpec#services must be specified and not empty.
 type FilterSpec struct {
    // Identifiers of tasks that can match the filter.
    Tasks map[string]bool
    // Identifiers of services. Tasks created by operations in these services match the filter (see task.CommonInfo#service).
    Services map[string]bool
    // Identifiers of operations. Tasks created by these operations match the filter (see task.CommonInfo#operation). 
//
//  Note that an operation identifier by itself is not globally unique. To filter on an operation, the identifier of the service interface containing the operation should also be specified in ``services``.. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    Operations map[string]bool
    // Status that a task must have to match the filter (see task.CommonInfo#status).
    Status map[task.Status]bool
    // Identifiers of the targets the operation for the associated task created or was performed on (see task.CommonInfo#target).
    Targets []std.DynamicID
    // Users who must have initiated the operation for the associated task to match the filter (see task.CommonInfo#user).
    Users map[string]bool
}









func tasksGetInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["task"] = bindings.NewIdType([]string {"com.vmware.cis.task"}, "")
    fields["spec"] = bindings.NewOptionalType(bindings.NewReferenceType(GetSpecBindingType))
    fieldNameMap["task"] = "Task"
    fieldNameMap["spec"] = "Spec"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func tasksGetOutputType() bindings.BindingType {
    return bindings.NewReferenceType(task.InfoBindingType)
}

func tasksGetRestMetadata() protocol.OperationRestMetadata {
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
       map[string]int{"Error": 500,"NotFound": 404,"ResourceInaccessible": 500,"ServiceUnavailable": 503,"Unauthenticated": 401,"Unauthorized": 403})
}


func tasksListInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["filter_spec"] = bindings.NewOptionalType(bindings.NewReferenceType(FilterSpecBindingType))
    fields["result_spec"] = bindings.NewOptionalType(bindings.NewReferenceType(GetSpecBindingType))
    fieldNameMap["filter_spec"] = "FilterSpec"
    fieldNameMap["result_spec"] = "ResultSpec"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func tasksListOutputType() bindings.BindingType {
    return bindings.NewMapType(bindings.NewIdType([]string {"com.vmware.cis.task"}, ""), bindings.NewReferenceType(task.InfoBindingType),reflect.TypeOf(map[string]task.Info{}))
}

func tasksListRestMetadata() protocol.OperationRestMetadata {
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
       map[string]int{"InvalidArgument": 400,"ResourceInaccessible": 500,"ServiceUnavailable": 503,"Unauthenticated": 401,"Unauthorized": 403})
}


func tasksCancelInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["task"] = bindings.NewIdType([]string {"com.vmware.cis.task"}, "")
    fieldNameMap["task"] = "Task"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func tasksCancelOutputType() bindings.BindingType {
    return bindings.NewVoidType()
}

func tasksCancelRestMetadata() protocol.OperationRestMetadata {
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
       map[string]int{"Error": 500,"NotAllowedInCurrentState": 400,"NotFound": 404,"ResourceInaccessible": 500,"ServiceUnavailable": 503,"Unauthenticated": 401,"Unauthorized": 403,"Unsupported": 400})
}



func GetSpecBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["return_all"] = bindings.NewOptionalType(bindings.NewBooleanType())
    fieldNameMap["return_all"] = "ReturnAll"
    fields["exclude_result"] = bindings.NewOptionalType(bindings.NewBooleanType())
    fieldNameMap["exclude_result"] = "ExcludeResult"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.cis.tasks.get_spec",fields, reflect.TypeOf(GetSpec{}), fieldNameMap, validators)
}

func FilterSpecBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["tasks"] = bindings.NewOptionalType(bindings.NewSetType(bindings.NewIdType([]string {"com.vmware.cis.task"}, ""), reflect.TypeOf(map[string]bool{})))
    fieldNameMap["tasks"] = "Tasks"
    fields["services"] = bindings.NewOptionalType(bindings.NewSetType(bindings.NewIdType([]string {"com.vmware.vapi.service"}, ""), reflect.TypeOf(map[string]bool{})))
    fieldNameMap["services"] = "Services"
    fields["operations"] = bindings.NewOptionalType(bindings.NewSetType(bindings.NewIdType([]string {"com.vmware.vapi.operation"}, ""), reflect.TypeOf(map[string]bool{})))
    fieldNameMap["operations"] = "Operations"
    fields["status"] = bindings.NewOptionalType(bindings.NewSetType(bindings.NewEnumType("com.vmware.cis.task.status", reflect.TypeOf(task.Status(task.Status_PENDING))), reflect.TypeOf(map[task.Status]bool{})))
    fieldNameMap["status"] = "Status"
    fields["targets"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewReferenceType(std.DynamicIDBindingType), reflect.TypeOf([]std.DynamicID{})))
    fieldNameMap["targets"] = "Targets"
    fields["users"] = bindings.NewOptionalType(bindings.NewSetType(bindings.NewStringType(), reflect.TypeOf(map[string]bool{})))
    fieldNameMap["users"] = "Users"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.cis.tasks.filter_spec",fields, reflect.TypeOf(FilterSpec{}), fieldNameMap, validators)
}


