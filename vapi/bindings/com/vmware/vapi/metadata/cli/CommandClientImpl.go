
/* Copyright © 2019 VMware, Inc. All Rights Reserved.
     SPDX-License-Identifier: BSD-2-Clause */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Client stubs for service: Command
 * Functions that implement the generated CommandClient interface
 */


package cli

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


type CommandClientImpl struct{
      interfaceName string
      interfaceDefinition core.InterfaceDefinition
      methodIdentifiers[]core.MethodIdentifier
      methodNameToDefMap map[string]*core.MethodDefinition
      errorBindingMap map[string]bindings.BindingType
      interfaceIdentifier core.InterfaceIdentifier
      connector client.Connector
}


func NewCommandClientImpl(connector client.Connector) *CommandClientImpl {
      interfaceName := "com.vmware.vapi.metadata.cli.command"
      interfaceIdentifier := core.NewInterfaceIdentifier(interfaceName)
      methodIdentifiers := []core.MethodIdentifier{
          core.NewMethodIdentifier(interfaceIdentifier, "list"),
          core.NewMethodIdentifier(interfaceIdentifier, "get"),
          core.NewMethodIdentifier(interfaceIdentifier, "fingerprint"),
      }
      interfaceDefinition := core.NewInterfaceDefinition(interfaceIdentifier, methodIdentifiers)
      errorBindingMap := make(map[string]bindings.BindingType)
      errorBindingMap[errors.InternalServerError{}.Error()] = errors.InternalServerErrorBindingType()
	  errorBindingMap[errors.InvalidArgument{}.Error()] = errors.InvalidArgumentBindingType()
	  errorBindingMap[errors.OperationNotFound{}.Error()] = errors.OperationNotFoundBindingType()
	  errorBindingMap[errors.UnexpectedInput{}.Error()] = errors.UnexpectedInputBindingType()
	  errorBindingMap[errors.ServiceUnavailable{}.Error()] = errors.ServiceUnavailableBindingType()
	  errorBindingMap[errors.TimedOut{}.Error()] = errors.TimedOutBindingType()
      cIface := CommandClientImpl{interfaceName: interfaceName, methodIdentifiers: methodIdentifiers, interfaceDefinition: interfaceDefinition, errorBindingMap: errorBindingMap, interfaceIdentifier:interfaceIdentifier, connector: connector}
      cIface.methodNameToDefMap = make(map[string]*core.MethodDefinition)
      cIface.methodNameToDefMap["list"] = cIface.listMethodDefinition()
      cIface.methodNameToDefMap["get"] = cIface.getMethodDefinition()
      cIface.methodNameToDefMap["fingerprint"] = cIface.fingerprintMethodDefinition()
      return &cIface
}

func (cIface *CommandClientImpl) List(pathParam *string) ([]Identity, error) {
	typeConverter := cIface.connector.TypeConverter()
	methodIdentifier := core.NewMethodIdentifier(cIface.interfaceIdentifier, "list")
	sv := bindings.NewStructValueBuilder(commandListInputType(), typeConverter)
	sv.AddStructField("Path", pathParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
        var emptyOutput []Identity
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := commandListRestMetadata
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	cIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult:= cIface.Invoke(cIface.connector.NewExecutionContext(), methodIdentifier, inputDataValue)
	var emptyOutput []Identity
    if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), commandListOutputType())
		if errorInOutput != nil {
		    return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
	    return output.([]Identity), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), cIface.errorBindingMap[methodResult.Error().Name()])
		if errorInError != nil {
		    return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}
func (cIface *CommandClientImpl) Get(identityParam Identity) (Info, error) {
	typeConverter := cIface.connector.TypeConverter()
	methodIdentifier := core.NewMethodIdentifier(cIface.interfaceIdentifier, "get")
	sv := bindings.NewStructValueBuilder(commandGetInputType(), typeConverter)
	sv.AddStructField("Identity", identityParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
        var emptyOutput Info
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := commandGetRestMetadata
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	cIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult:= cIface.Invoke(cIface.connector.NewExecutionContext(), methodIdentifier, inputDataValue)
	var emptyOutput Info
    if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), commandGetOutputType())
		if errorInOutput != nil {
		    return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
	    return output.(Info), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), cIface.errorBindingMap[methodResult.Error().Name()])
		if errorInError != nil {
		    return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}
func (cIface *CommandClientImpl) Fingerprint() (string, error) {
	typeConverter := cIface.connector.TypeConverter()
	methodIdentifier := core.NewMethodIdentifier(cIface.interfaceIdentifier, "fingerprint")
	sv := bindings.NewStructValueBuilder(commandFingerprintInputType(), typeConverter)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
        var emptyOutput string
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := commandFingerprintRestMetadata
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	cIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult:= cIface.Invoke(cIface.connector.NewExecutionContext(), methodIdentifier, inputDataValue)
	var emptyOutput string
    if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), commandFingerprintOutputType())
		if errorInOutput != nil {
		    return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
	    return output.(string), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), cIface.errorBindingMap[methodResult.Error().Name()])
		if errorInError != nil {
		    return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (cIface *CommandClientImpl) Invoke(ctx *core.ExecutionContext, methodId core.MethodIdentifier, inputDataValue data.DataValue) core.MethodResult {
    methodResult := cIface.connector.GetApiProvider().Invoke(cIface.interfaceName, methodId.Name(), inputDataValue, ctx)
    return methodResult
}

func (cIface *CommandClientImpl) listMethodDefinition() *core.MethodDefinition {
      interfaceIdentifier := core.NewInterfaceIdentifier(cIface.interfaceName)
      typeConverter := cIface.connector.TypeConverter()

      input, inputError := typeConverter.ConvertToDataDefinition(commandListInputType())
      output, outputError := typeConverter.ConvertToDataDefinition(commandListOutputType())
      if(inputError != nil) {
          log.Errorf("Error in ConvertToDataDefinition for CommandClientImpl.list method's input - %s",
              bindings.VAPIerrorsToError(inputError).Error())
          return nil
      }
      if(outputError != nil) {
          log.Errorf("Error in ConvertToDataDefinition for CommandClientImpl.list method's output - %s",
              bindings.VAPIerrorsToError(outputError).Error())
          return nil
      }
      methodIdentifier := core.NewMethodIdentifier(interfaceIdentifier, "list")
      errorDefinitions := make([]data.ErrorDefinition,0)
      cIface.errorBindingMap[errors.NotFound{}.Error()] = errors.NotFoundBindingType()
	  errDef1, errError1 := typeConverter.ConvertToDataDefinition(errors.NotFoundBindingType())
	  if(errError1 != nil) {
	    log.Errorf("Error in ConvertToDataDefinition for CommandClientImpl.list method's errors.NotFound error - %s",
	        bindings.VAPIerrorsToError(errError1).Error())
        return nil
	  }
	  errorDefinitions = append(errorDefinitions, errDef1.(data.ErrorDefinition))

      methodDefinition := core.NewMethodDefinition(methodIdentifier, input, output, errorDefinitions)
      return &methodDefinition
}
func (cIface *CommandClientImpl) getMethodDefinition() *core.MethodDefinition {
      interfaceIdentifier := core.NewInterfaceIdentifier(cIface.interfaceName)
      typeConverter := cIface.connector.TypeConverter()

      input, inputError := typeConverter.ConvertToDataDefinition(commandGetInputType())
      output, outputError := typeConverter.ConvertToDataDefinition(commandGetOutputType())
      if(inputError != nil) {
          log.Errorf("Error in ConvertToDataDefinition for CommandClientImpl.get method's input - %s",
              bindings.VAPIerrorsToError(inputError).Error())
          return nil
      }
      if(outputError != nil) {
          log.Errorf("Error in ConvertToDataDefinition for CommandClientImpl.get method's output - %s",
              bindings.VAPIerrorsToError(outputError).Error())
          return nil
      }
      methodIdentifier := core.NewMethodIdentifier(interfaceIdentifier, "get")
      errorDefinitions := make([]data.ErrorDefinition,0)
      cIface.errorBindingMap[errors.NotFound{}.Error()] = errors.NotFoundBindingType()
	  errDef1, errError1 := typeConverter.ConvertToDataDefinition(errors.NotFoundBindingType())
	  if(errError1 != nil) {
	    log.Errorf("Error in ConvertToDataDefinition for CommandClientImpl.get method's errors.NotFound error - %s",
	        bindings.VAPIerrorsToError(errError1).Error())
        return nil
	  }
	  errorDefinitions = append(errorDefinitions, errDef1.(data.ErrorDefinition))

      methodDefinition := core.NewMethodDefinition(methodIdentifier, input, output, errorDefinitions)
      return &methodDefinition
}
func (cIface *CommandClientImpl) fingerprintMethodDefinition() *core.MethodDefinition {
      interfaceIdentifier := core.NewInterfaceIdentifier(cIface.interfaceName)
      typeConverter := cIface.connector.TypeConverter()

      input, inputError := typeConverter.ConvertToDataDefinition(commandFingerprintInputType())
      output, outputError := typeConverter.ConvertToDataDefinition(commandFingerprintOutputType())
      if(inputError != nil) {
          log.Errorf("Error in ConvertToDataDefinition for CommandClientImpl.fingerprint method's input - %s",
              bindings.VAPIerrorsToError(inputError).Error())
          return nil
      }
      if(outputError != nil) {
          log.Errorf("Error in ConvertToDataDefinition for CommandClientImpl.fingerprint method's output - %s",
              bindings.VAPIerrorsToError(outputError).Error())
          return nil
      }
      methodIdentifier := core.NewMethodIdentifier(interfaceIdentifier, "fingerprint")
      errorDefinitions := make([]data.ErrorDefinition,0)

      methodDefinition := core.NewMethodDefinition(methodIdentifier, input, output, errorDefinitions)
      return &methodDefinition
}



// The ``FormatterType`` enumeration class defines supported CLI output formatter types. See Info#formatter.
//
// <p> See {@link com.vmware.vapi.bindings.ApiEnumeration enumerated types description}.
 
type FormatterType string

const (
    // Displays command output as it is.
     FormatterType_SIMPLE FormatterType = "SIMPLE"
    // Displays command output in table format.
     FormatterType_TABLE FormatterType = "TABLE"
    // Displays command output in JSON format.
     FormatterType_JSON FormatterType = "JSON"
    // Displays command output in XML format.
     FormatterType_XML FormatterType = "XML"
    // Displays command output in CSV format.
     FormatterType_CSV FormatterType = "CSV"
    // Displays command output in HTML format.
     FormatterType_HTML FormatterType = "HTML"
)

func (f FormatterType) FormatterType() bool {
    switch f {
        case FormatterType_SIMPLE:
            return true
        case FormatterType_TABLE:
            return true
        case FormatterType_JSON:
            return true
        case FormatterType_XML:
            return true
        case FormatterType_CSV:
            return true
        case FormatterType_HTML:
            return true
        default:
            return false
    }
}




// The ``GenericType`` enumeration class defines generic types supported by ``Command`` interface. See OptionInfo#generic.
//
// <p> See {@link com.vmware.vapi.bindings.ApiEnumeration enumerated types description}.
 
type GenericType string

const (
    // Default case.
     GenericType_NONE GenericType = "NONE"
    // Input parameter is an optional.
     GenericType_OPTIONAL GenericType = "OPTIONAL"
    // Input parameter is a list.
     GenericType_LIST GenericType = "LIST"
    // Input parameter is an optional of type list.
     GenericType_OPTIONAL_LIST GenericType = "OPTIONAL_LIST"
    // Input parameter is a list of optionals.
     GenericType_LIST_OPTIONAL GenericType = "LIST_OPTIONAL"
)

func (g GenericType) GenericType() bool {
    switch g {
        case GenericType_NONE:
            return true
        case GenericType_OPTIONAL:
            return true
        case GenericType_LIST:
            return true
        case GenericType_OPTIONAL_LIST:
            return true
        case GenericType_LIST_OPTIONAL:
            return true
        default:
            return false
    }
}





// The ``OutputFieldInfo`` class describes the name used by the CLI to display a single property of a class element in the interface definition language.
 type OutputFieldInfo struct {
    // Name of the property.
    FieldName string
    // Name used by the CLI to display the property.
    DisplayName string
}






// The ``OutputInfo`` class describes the names used by the CLI to display the properties of a class element in the interface definition language as well as the order in which the properties will be displayed.
 type OutputInfo struct {
    // Name of the class.
    StructureId string
    // The order in which the properties of the class will be displayed by the CLI as well as the names used to display the properties.
    OutputFields []OutputFieldInfo
}






// The ``OptionInfo`` class describes information about a specific input option of a command.
 type OptionInfo struct {
    // The long option name of the parameter as used by the user.
    LongOption string
    // The single character value option name.
    ShortOption *string
    // The fully qualified name of the option referred to by the operation element in Info#operationId.
    FieldName string
    // The description of the option to be displayed to the user when they request usage information for a CLI command.
    Description string
    // The type of option. This is used to display information about what kind of data is expected (string, number, boolean, etc.) for the option when they request usage information for a CLI command. For enumeration class this stores the fully qualified enumeration class id.
    Type_ string
    // This is used to tell the user whether the option is required or optional, or whether they can specify the option multiple times.
    Generic GenericType
}






// The ``Identity`` class uniquely identifies a command in the CLI commands tree.
 type Identity struct {
    // The dot-separated path of the namespace containing the command in the CLI command tree.
    Path string
    // Name of the command.
    Name string
}






// The ``Info`` class contains information about a command. It includes the identity of the command, a description, information about the interface and method that implement the command, and CLI-specific information for the command.
 type Info struct {
    // Basic command identity.
    Identity Identity
    // The text description displayed to the user in help output.
    Description string
    // The service identifier that contains the operations for this CLI command.
    ServiceId string
    // The operation identifier corresponding to this CLI command.
    OperationId string
    // The input for this command.
    Options []OptionInfo
    // The formatter to use when displaying the output of this command.
    Formatter *FormatterType
    // List of output structure name and output field info.
    OutputFieldList []OutputInfo
}









func commandListInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["path"] = bindings.NewOptionalType(bindings.NewStringType())
    fieldNameMap["path"] = "Path"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func commandListOutputType() bindings.BindingType {
    return bindings.NewListType(bindings.NewReferenceType(IdentityBindingType), reflect.TypeOf([]Identity{}))
}

func commandListRestMetadata() protocol.OperationRestMetadata {
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


func commandGetInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["identity"] = bindings.NewReferenceType(IdentityBindingType)
    fieldNameMap["identity"] = "Identity"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func commandGetOutputType() bindings.BindingType {
    return bindings.NewReferenceType(InfoBindingType)
}

func commandGetRestMetadata() protocol.OperationRestMetadata {
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


func commandFingerprintInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func commandFingerprintOutputType() bindings.BindingType {
    return bindings.NewStringType()
}

func commandFingerprintRestMetadata() protocol.OperationRestMetadata {
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
       map[string]int{})
}



func OutputFieldInfoBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["field_name"] = bindings.NewStringType()
    fieldNameMap["field_name"] = "FieldName"
    fields["display_name"] = bindings.NewStringType()
    fieldNameMap["display_name"] = "DisplayName"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vapi.metadata.cli.command.output_field_info",fields, reflect.TypeOf(OutputFieldInfo{}), fieldNameMap, validators)
}

func OutputInfoBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["structure_id"] = bindings.NewIdType([]string {"com.vmware.vapi.structure"}, "")
    fieldNameMap["structure_id"] = "StructureId"
    fields["output_fields"] = bindings.NewListType(bindings.NewReferenceType(OutputFieldInfoBindingType), reflect.TypeOf([]OutputFieldInfo{}))
    fieldNameMap["output_fields"] = "OutputFields"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vapi.metadata.cli.command.output_info",fields, reflect.TypeOf(OutputInfo{}), fieldNameMap, validators)
}

func OptionInfoBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["long_option"] = bindings.NewStringType()
    fieldNameMap["long_option"] = "LongOption"
    fields["short_option"] = bindings.NewOptionalType(bindings.NewStringType())
    fieldNameMap["short_option"] = "ShortOption"
    fields["field_name"] = bindings.NewStringType()
    fieldNameMap["field_name"] = "FieldName"
    fields["description"] = bindings.NewStringType()
    fieldNameMap["description"] = "Description"
    fields["type"] = bindings.NewStringType()
    fieldNameMap["type"] = "Type_"
    fields["generic"] = bindings.NewEnumType("com.vmware.vapi.metadata.cli.command.generic_type", reflect.TypeOf(GenericType(GenericType_NONE)))
    fieldNameMap["generic"] = "Generic"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vapi.metadata.cli.command.option_info",fields, reflect.TypeOf(OptionInfo{}), fieldNameMap, validators)
}

func IdentityBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["path"] = bindings.NewStringType()
    fieldNameMap["path"] = "Path"
    fields["name"] = bindings.NewStringType()
    fieldNameMap["name"] = "Name"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vapi.metadata.cli.command.identity",fields, reflect.TypeOf(Identity{}), fieldNameMap, validators)
}

func InfoBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["identity"] = bindings.NewReferenceType(IdentityBindingType)
    fieldNameMap["identity"] = "Identity"
    fields["description"] = bindings.NewStringType()
    fieldNameMap["description"] = "Description"
    fields["service_id"] = bindings.NewIdType([]string {"com.vmware.vapi.service"}, "")
    fieldNameMap["service_id"] = "ServiceId"
    fields["operation_id"] = bindings.NewIdType([]string {"com.vmware.vapi.operation"}, "")
    fieldNameMap["operation_id"] = "OperationId"
    fields["options"] = bindings.NewListType(bindings.NewReferenceType(OptionInfoBindingType), reflect.TypeOf([]OptionInfo{}))
    fieldNameMap["options"] = "Options"
    fields["formatter"] = bindings.NewOptionalType(bindings.NewEnumType("com.vmware.vapi.metadata.cli.command.formatter_type", reflect.TypeOf(FormatterType(FormatterType_SIMPLE))))
    fieldNameMap["formatter"] = "Formatter"
    fields["output_field_list"] = bindings.NewListType(bindings.NewReferenceType(OutputInfoBindingType), reflect.TypeOf([]OutputInfo{}))
    fieldNameMap["output_field_list"] = "OutputFieldList"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vapi.metadata.cli.command.info",fields, reflect.TypeOf(Info{}), fieldNameMap, validators)
}


