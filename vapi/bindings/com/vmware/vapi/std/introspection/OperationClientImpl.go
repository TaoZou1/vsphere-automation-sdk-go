
/* Copyright © 2019 VMware, Inc. All Rights Reserved.
     SPDX-License-Identifier: BSD-2-Clause */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Client stubs for service: Operation
 * Functions that implement the generated OperationClient interface
 */


package introspection

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
      interfaceName := "com.vmware.vapi.std.introspection.operation"
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

func (oIface *OperationClientImpl) List(serviceIdParam string) (map[string]bool, error) {
	typeConverter := oIface.connector.TypeConverter()
	methodIdentifier := core.NewMethodIdentifier(oIface.interfaceIdentifier, "list")
	sv := bindings.NewStructValueBuilder(operationListInputType(), typeConverter)
	sv.AddStructField("ServiceId", serviceIdParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
        var emptyOutput map[string]bool
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := operationListRestMetadata
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	oIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult:= oIface.Invoke(oIface.connector.NewExecutionContext(), methodIdentifier, inputDataValue)
	var emptyOutput map[string]bool
    if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), operationListOutputType())
		if errorInOutput != nil {
		    return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
	    return output.(map[string]bool), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), oIface.errorBindingMap[methodResult.Error().Name()])
		if errorInError != nil {
		    return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}
func (oIface *OperationClientImpl) Get(serviceIdParam string, operationIdParam string) (Info, error) {
	typeConverter := oIface.connector.TypeConverter()
	methodIdentifier := core.NewMethodIdentifier(oIface.interfaceIdentifier, "get")
	sv := bindings.NewStructValueBuilder(operationGetInputType(), typeConverter)
	sv.AddStructField("ServiceId", serviceIdParam)
	sv.AddStructField("OperationId", operationIdParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
        var emptyOutput Info
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := operationGetRestMetadata
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	oIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult:= oIface.Invoke(oIface.connector.NewExecutionContext(), methodIdentifier, inputDataValue)
	var emptyOutput Info
    if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), operationGetOutputType())
		if errorInOutput != nil {
		    return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
	    return output.(Info), nil
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




// The DataDefinition structure describes a vAPI data type.
 type DataDefinition struct {
    // Data type of the value.
    Type_ DataDefinition_DataType
    // Contains the element definition for generic data types like List and Optional.
    ElementDefinition *DataDefinition
    // Fully qualified name of the structure.
    Name *string
    // Fields of the structure type. The key of the map is the canonical name of the field and the value is the DataDefinition for the field. The order of the structure fields defined in IDL is not maintained by the Operation service.
    Fields map[string]DataDefinition
}




    
    // The DataDefinition_DataType enumeration provides values representing the data types supported by the vAPI infrastructure.
    //
    // <p> See {@link com.vmware.vapi.bindings.ApiEnumeration enumerated types description}.
     
    type DataDefinition_DataType string

    const (
        // Indicates the value is a binary type.
         DataDefinition_DataType_BINARY DataDefinition_DataType = "BINARY"
        // Indicates the value is a boolean type. The possible values are True and False equivalent of the language used to invoke this operation.
         DataDefinition_DataType_BOOLEAN DataDefinition_DataType = "BOOLEAN"
        // Indicates the value is a double type. It is a 64 bit floating point number.
         DataDefinition_DataType_DOUBLE DataDefinition_DataType = "DOUBLE"
        // Indicates the value is a dynamic structure. This means, any data of type DataDefinition_DataType#DataType_STRUCTURE can be used.
         DataDefinition_DataType_DYNAMIC_STRUCTURE DataDefinition_DataType = "DYNAMIC_STRUCTURE"
        // Indicates the value is a specific error type.
         DataDefinition_DataType_ERROR DataDefinition_DataType = "ERROR"
        // Indicates the value is arbitrary error type. This means, any data of type DataDefinition_DataType#DataType_ERROR can be used.
         DataDefinition_DataType_ANY_ERROR DataDefinition_DataType = "ANY_ERROR"
        // Indicates the value is a list data type. Any value of this type can have zero or more elements in the list.
         DataDefinition_DataType_LIST DataDefinition_DataType = "LIST"
        // Indicates the value is a long data type. It is a 64 bit signed integer number.
         DataDefinition_DataType_LONG DataDefinition_DataType = "LONG"
        // Indicates the value is an opaque type. This means, data of any DataDefinition_DataType can be used.
         DataDefinition_DataType_OPAQUE DataDefinition_DataType = "OPAQUE"
        // Indicates the value is an optional data type. Any value of this type can be null.
         DataDefinition_DataType_OPTIONAL DataDefinition_DataType = "OPTIONAL"
        // Indicates the value is a secret data type. This is used for sensitive information. The server will not log any data of this type and if possible wipe the data from the memory after usage.
         DataDefinition_DataType_SECRET DataDefinition_DataType = "SECRET"
        // Indicates the value is a string data type. This is a unicode string.
         DataDefinition_DataType_STRING DataDefinition_DataType = "STRING"
        // Indicates the value is a structure data type. A structure has string identifier and a set of fields with corresponding values.
         DataDefinition_DataType_STRUCTURE DataDefinition_DataType = "STRUCTURE"
        // Indicates the value is a structure reference. This is used to break circular dependencies in the type references. This just has a string identifier of the structure. Clients have to maintain a list of structures already visited and use that to resolve this reference.
         DataDefinition_DataType_STRUCTURE_REF DataDefinition_DataType = "STRUCTURE_REF"
        // Indicates the value is a void data type.
         DataDefinition_DataType_VOID DataDefinition_DataType = "VOID"
    )

    func (d DataDefinition_DataType) DataDefinition_DataType() bool {
        switch d {
            case DataDefinition_DataType_BINARY:
                return true
            case DataDefinition_DataType_BOOLEAN:
                return true
            case DataDefinition_DataType_DOUBLE:
                return true
            case DataDefinition_DataType_DYNAMIC_STRUCTURE:
                return true
            case DataDefinition_DataType_ERROR:
                return true
            case DataDefinition_DataType_ANY_ERROR:
                return true
            case DataDefinition_DataType_LIST:
                return true
            case DataDefinition_DataType_LONG:
                return true
            case DataDefinition_DataType_OPAQUE:
                return true
            case DataDefinition_DataType_OPTIONAL:
                return true
            case DataDefinition_DataType_SECRET:
                return true
            case DataDefinition_DataType_STRING:
                return true
            case DataDefinition_DataType_STRUCTURE:
                return true
            case DataDefinition_DataType_STRUCTURE_REF:
                return true
            case DataDefinition_DataType_VOID:
                return true
            default:
                return false
        }
    }



// Information about a vAPI operation.
 type Info struct {
    // DataDefinition describing the operation input. 
//
//  The DataDefinition#type of this field will be DataDefinition_DataType#DataType_STRUCTURE. The keys of DataDefinition#fields are the names of the operation parameters, and the values of DataDefinition#fields describe the type of the operation parameters.
    InputDefinition DataDefinition
    // DataDefinition describing the operation output.
    OutputDefinition DataDefinition
    // List of DataDefinition describing the errors that the operation might report. 
//
//  The DataDefinition#type of every element in this list will be DataDefinition_DataType#DataType_ERROR.
    ErrorDefinitions []DataDefinition
}









func operationListInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["service_id"] = bindings.NewIdType([]string {"com.vmware.vapi.service"}, "")
    fieldNameMap["service_id"] = "ServiceId"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func operationListOutputType() bindings.BindingType {
    return bindings.NewSetType(bindings.NewIdType([]string {"com.vmware.vapi.operation"}, ""), reflect.TypeOf(map[string]bool{}))
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
    return bindings.NewReferenceType(InfoBindingType)
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



func DataDefinitionBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["type"] = bindings.NewEnumType("com.vmware.vapi.std.introspection.operation.data_definition.data_type", reflect.TypeOf(DataDefinition_DataType(DataDefinition_DataType_BINARY)))
    fieldNameMap["type"] = "Type_"
    fields["element_definition"] = bindings.NewOptionalType(bindings.NewReferenceType(DataDefinitionBindingType))
    fieldNameMap["element_definition"] = "ElementDefinition"
    fields["name"] = bindings.NewOptionalType(bindings.NewStringType())
    fieldNameMap["name"] = "Name"
    fields["fields"] = bindings.NewOptionalType(bindings.NewMapType(bindings.NewStringType(), bindings.NewReferenceType(DataDefinitionBindingType),reflect.TypeOf(map[string]DataDefinition{})))
    fieldNameMap["fields"] = "Fields"
    var validators = []bindings.Validator{}
    uv1 := bindings.NewUnionValidator("type",
        map[string][]bindings.FieldData {
            "OPTIONAL": []bindings.FieldData {
                 bindings.NewFieldData("element_definition", true),
            },
            "LIST": []bindings.FieldData {
                 bindings.NewFieldData("element_definition", true),
            },
            "STRUCTURE": []bindings.FieldData {
                 bindings.NewFieldData("name", true),
                 bindings.NewFieldData("fields", true),
            },
            "STRUCTURE_REF": []bindings.FieldData {
                 bindings.NewFieldData("name", true),
            },
            "ERROR": []bindings.FieldData {
                 bindings.NewFieldData("name", true),
                 bindings.NewFieldData("fields", true),
            },
            "BINARY": []bindings.FieldData {},
            "BOOLEAN": []bindings.FieldData {},
            "DOUBLE": []bindings.FieldData {},
            "DYNAMIC_STRUCTURE": []bindings.FieldData {},
            "ANY_ERROR": []bindings.FieldData {},
            "LONG": []bindings.FieldData {},
            "OPAQUE": []bindings.FieldData {},
            "SECRET": []bindings.FieldData {},
            "STRING": []bindings.FieldData {},
            "VOID": []bindings.FieldData {},
        },
    )
    validators = append(validators, uv1)
    return bindings.NewStructType("com.vmware.vapi.std.introspection.operation.data_definition",fields, reflect.TypeOf(DataDefinition{}), fieldNameMap, validators)
}

func InfoBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["input_definition"] = bindings.NewReferenceType(DataDefinitionBindingType)
    fieldNameMap["input_definition"] = "InputDefinition"
    fields["output_definition"] = bindings.NewReferenceType(DataDefinitionBindingType)
    fieldNameMap["output_definition"] = "OutputDefinition"
    fields["error_definitions"] = bindings.NewListType(bindings.NewReferenceType(DataDefinitionBindingType), reflect.TypeOf([]DataDefinition{}))
    fieldNameMap["error_definitions"] = "ErrorDefinitions"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vapi.std.introspection.operation.info",fields, reflect.TypeOf(Info{}), fieldNameMap, validators)
}


