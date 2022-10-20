// Copyright © 2019-2021 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Interface file for service: Source
// Used by client-side stubs.

package privilege

import (
	"github.com/vmware/vsphere-automation-sdk-go/lib/vapi/std/errors"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/bindings"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/core"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/lib"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/protocol/client"
)

const _ = core.SupportedByRuntimeVersion1

// The ``Source`` interface provides methods to manage the sources of privilege metadata information.
//
//  The interface definition language infrastructure provides tools to generate various kinds of metadata in JSON format from the interface definition files and additional properties files. One of the generated files contains privilege information. The generated file can be registered as a source of metadata.
//
//  The privilege file contains all the data present in the interface definition files. Each privilege file contains data about one component element. When a privilege file is added as a source, each source contributes only one component element's metadata.
//
//  Privilege metadata can also be discovered from a remote server that supports the privilege metadata interfaces (see com.vmware.vapi.metadata.privilege). Since multiple components can be registered with a single metadata server, when a remote server is registered as a source, that source can contribute more than one component.
type SourceClient interface {

	// Creates a new metadata source. Once the server validates the registration information of the metadata source, the privilege metadata is retrieved from the source. This populates elements in all the interfaces defined in com.vmware.vapi.metadata.privilege package.
	//
	// @param sourceIdParam metadata source identifier.
	// The parameter must be an identifier for the resource type: ``com.vmware.vapi.metadata.privilege.source``.
	// @param specParam create specification.
	// @throws AlreadyExists if the metadata source identifier is already registered with the infrastructure.
	// @throws InvalidArgument if the type of the source specified in metadata.SourceCreateSpec#type is invalid.
	// @throws InvalidArgument if the file specified in metadata.SourceCreateSpec#filepath is not a valid JSON file or if the format of the privilege metadata in the JSON file is invalid.
	// @throws InvalidArgument if the URI specified in metadata.SourceCreateSpec#address is unreachable or if there is a transport protocol or message protocol mismatch between the client and the server or if the remote server do not have interfaces present in com.vmware.vapi.metadata.privilege package.
	// @throws NotFound if the file specified in metadata.SourceCreateSpec#filepath does not exist.
	Create(sourceIdParam string, specParam SourceCreateSpec) error

	// Deletes an existing privilege metadata source from the infrastructure.
	//
	// @param sourceIdParam Identifier of the metadata source.
	// The parameter must be an identifier for the resource type: ``com.vmware.vapi.metadata.privilege.source``.
	// @throws NotFound if the metadata source associated with ``source_id`` is not found.
	Delete(sourceIdParam string) error

	// Retrieves information about the metadata source corresponding to ``source_id``.
	//
	// @param sourceIdParam Identifier of the metadata source.
	// The parameter must be an identifier for the resource type: ``com.vmware.vapi.metadata.privilege.source``.
	// @return The SourceInfo instance that corresponds to ``source_id``
	// @throws NotFound if the metadata source associated with ``source_id`` is not found.
	Get(sourceIdParam string) (SourceInfo, error)

	// Returns the identifiers of the metadata sources currently registered with the infrastructure.
	// @return The list of identifiers for metadata sources currently registered.
	// The return value will contain identifiers for the resource type: ``com.vmware.vapi.metadata.privilege.source``.
	List() ([]string, error)

	// Reloads the privilege metadata from all the metadata sources or of a particular metadata source if ``source_id`` is specified.
	//
	// @param sourceIdParam Identifier of the metadata source.
	// The parameter must be an identifier for the resource type: ``com.vmware.vapi.metadata.privilege.source``.
	// If unspecified, all the metadata sources are reloaded.
	// @throws NotFound if the metadata source associated with ``source_id`` is not found.
	Reload(sourceIdParam *string) error

	// Returns the aggregate fingerprint of metadata from all the metadata sources or from a particular metadata source if ``source_id`` is specified.
	//
	// @param sourceIdParam Identifier of the metadata source.
	// The parameter must be an identifier for the resource type: ``com.vmware.vapi.metadata.privilege.source``.
	// If unspecified, the fingerprint of all the metadata sources is returned.
	// @return Aggregate fingerprint of all the metadata sources or of a particular metadata source.
	// @throws NotFound if the metadata source associated with ``source_id`` is not found.
	Fingerprint(sourceIdParam *string) (string, error)
}

type sourceClient struct {
	connector           client.Connector
	interfaceDefinition core.InterfaceDefinition
	errorsBindingMap    map[string]bindings.BindingType
}

func NewSourceClient(connector client.Connector) *sourceClient {
	interfaceIdentifier := core.NewInterfaceIdentifier("com.vmware.vapi.metadata.privilege.source")
	methodIdentifiers := map[string]core.MethodIdentifier{
		"create":      core.NewMethodIdentifier(interfaceIdentifier, "create"),
		"delete":      core.NewMethodIdentifier(interfaceIdentifier, "delete"),
		"get":         core.NewMethodIdentifier(interfaceIdentifier, "get"),
		"list":        core.NewMethodIdentifier(interfaceIdentifier, "list"),
		"reload":      core.NewMethodIdentifier(interfaceIdentifier, "reload"),
		"fingerprint": core.NewMethodIdentifier(interfaceIdentifier, "fingerprint"),
	}
	interfaceDefinition := core.NewInterfaceDefinition(interfaceIdentifier, methodIdentifiers)
	errorsBindingMap := make(map[string]bindings.BindingType)

	sIface := sourceClient{interfaceDefinition: interfaceDefinition, errorsBindingMap: errorsBindingMap, connector: connector}
	return &sIface
}

func (sIface *sourceClient) GetErrorBindingType(errorName string) bindings.BindingType {
	if entry, ok := sIface.errorsBindingMap[errorName]; ok {
		return entry
	}
	return errors.ERROR_BINDINGS_MAP[errorName]
}

func (sIface *sourceClient) Create(sourceIdParam string, specParam SourceCreateSpec) error {
	typeConverter := sIface.connector.TypeConverter()
	executionContext := sIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(sourceCreateInputType(), typeConverter)
	sv.AddStructField("SourceId", sourceIdParam)
	sv.AddStructField("Spec", specParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		return bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := sourceCreateRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	sIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := sIface.connector.GetApiProvider().Invoke("com.vmware.vapi.metadata.privilege.source", "create", inputDataValue, executionContext)
	if methodResult.IsSuccess() {
		return nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), sIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return bindings.VAPIerrorsToError(errorInError)
		}
		return methodError.(error)
	}
}

func (sIface *sourceClient) Delete(sourceIdParam string) error {
	typeConverter := sIface.connector.TypeConverter()
	executionContext := sIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(sourceDeleteInputType(), typeConverter)
	sv.AddStructField("SourceId", sourceIdParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		return bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := sourceDeleteRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	sIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := sIface.connector.GetApiProvider().Invoke("com.vmware.vapi.metadata.privilege.source", "delete", inputDataValue, executionContext)
	if methodResult.IsSuccess() {
		return nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), sIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return bindings.VAPIerrorsToError(errorInError)
		}
		return methodError.(error)
	}
}

func (sIface *sourceClient) Get(sourceIdParam string) (SourceInfo, error) {
	typeConverter := sIface.connector.TypeConverter()
	executionContext := sIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(sourceGetInputType(), typeConverter)
	sv.AddStructField("SourceId", sourceIdParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput SourceInfo
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := sourceGetRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	sIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := sIface.connector.GetApiProvider().Invoke("com.vmware.vapi.metadata.privilege.source", "get", inputDataValue, executionContext)
	var emptyOutput SourceInfo
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), sourceGetOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(SourceInfo), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), sIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (sIface *sourceClient) List() ([]string, error) {
	typeConverter := sIface.connector.TypeConverter()
	executionContext := sIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(sourceListInputType(), typeConverter)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput []string
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := sourceListRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	sIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := sIface.connector.GetApiProvider().Invoke("com.vmware.vapi.metadata.privilege.source", "list", inputDataValue, executionContext)
	var emptyOutput []string
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), sourceListOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.([]string), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), sIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (sIface *sourceClient) Reload(sourceIdParam *string) error {
	typeConverter := sIface.connector.TypeConverter()
	executionContext := sIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(sourceReloadInputType(), typeConverter)
	sv.AddStructField("SourceId", sourceIdParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		return bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := sourceReloadRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	sIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := sIface.connector.GetApiProvider().Invoke("com.vmware.vapi.metadata.privilege.source", "reload", inputDataValue, executionContext)
	if methodResult.IsSuccess() {
		return nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), sIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return bindings.VAPIerrorsToError(errorInError)
		}
		return methodError.(error)
	}
}

func (sIface *sourceClient) Fingerprint(sourceIdParam *string) (string, error) {
	typeConverter := sIface.connector.TypeConverter()
	executionContext := sIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(sourceFingerprintInputType(), typeConverter)
	sv.AddStructField("SourceId", sourceIdParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput string
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := sourceFingerprintRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	sIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := sIface.connector.GetApiProvider().Invoke("com.vmware.vapi.metadata.privilege.source", "fingerprint", inputDataValue, executionContext)
	var emptyOutput string
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), sourceFingerprintOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(string), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), sIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}
