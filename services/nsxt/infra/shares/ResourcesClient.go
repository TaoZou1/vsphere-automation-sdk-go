// Copyright © 2019-2021 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Interface file for service: Resources
// Used by client-side stubs.

package shares

import (
	"github.com/TaoZou1/vsphere-automation-sdk-go/lib/vapi/std/errors"
	"github.com/TaoZou1/vsphere-automation-sdk-go/runtime/bindings"
	"github.com/TaoZou1/vsphere-automation-sdk-go/runtime/core"
	"github.com/TaoZou1/vsphere-automation-sdk-go/runtime/lib"
	"github.com/TaoZou1/vsphere-automation-sdk-go/runtime/protocol/client"
	"github.com/TaoZou1/vsphere-automation-sdk-go/services/nsxt/model"
)

const _ = core.SupportedByRuntimeVersion1

type ResourcesClient interface {

	// Delete Shared Resource.
	//
	// @param shareIdParam (required)
	// @param sharedResourceIdParam (required)
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Delete(shareIdParam string, sharedResourceIdParam string) error

	// Get the shared resource that represents shared objects.
	//
	// @param shareIdParam Share ID (required)
	// @param sharedResourceIdParam Shared resource id (required)
	// @return com.vmware.nsx_policy.model.SharedResource
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Get(shareIdParam string, sharedResourceIdParam string) (model.SharedResource, error)

	// Get the list of shared resource in a given share.
	//
	// @param shareIdParam Share ID (required)
	// @param resourceTypeParam resource type (optional)
	// @return com.vmware.nsx_policy.model.SharedResourceListResult
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	List(shareIdParam string, resourceTypeParam *string) (model.SharedResourceListResult, error)

	// If Shared resource already exists, update the shared resource entity. User Presence of the shared resource denotes sharing of the parent object.
	//
	// @param shareIdParam Share ID (required)
	// @param sharedResourceIdParam Shared resource id (required)
	// @param sharedResourceParam (required)
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Patch(shareIdParam string, sharedResourceIdParam string, sharedResourceParam model.SharedResource) error

	// If Shared resource already exists, update the shared resource entity. User Presence of the shared resource denotes sharing of the parent object.
	//
	// @param shareIdParam Share ID (required)
	// @param sharedResourceIdParam Shared resource id (required)
	// @param sharedResourceParam (required)
	// @return com.vmware.nsx_policy.model.SharedResource
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Update(shareIdParam string, sharedResourceIdParam string, sharedResourceParam model.SharedResource) (model.SharedResource, error)
}

type resourcesClient struct {
	connector           client.Connector
	interfaceDefinition core.InterfaceDefinition
	errorsBindingMap    map[string]bindings.BindingType
}

func NewResourcesClient(connector client.Connector) *resourcesClient {
	interfaceIdentifier := core.NewInterfaceIdentifier("com.vmware.nsx_policy.infra.shares.resources")
	methodIdentifiers := map[string]core.MethodIdentifier{
		"delete": core.NewMethodIdentifier(interfaceIdentifier, "delete"),
		"get":    core.NewMethodIdentifier(interfaceIdentifier, "get"),
		"list":   core.NewMethodIdentifier(interfaceIdentifier, "list"),
		"patch":  core.NewMethodIdentifier(interfaceIdentifier, "patch"),
		"update": core.NewMethodIdentifier(interfaceIdentifier, "update"),
	}
	interfaceDefinition := core.NewInterfaceDefinition(interfaceIdentifier, methodIdentifiers)
	errorsBindingMap := make(map[string]bindings.BindingType)

	rIface := resourcesClient{interfaceDefinition: interfaceDefinition, errorsBindingMap: errorsBindingMap, connector: connector}
	return &rIface
}

func (rIface *resourcesClient) GetErrorBindingType(errorName string) bindings.BindingType {
	if entry, ok := rIface.errorsBindingMap[errorName]; ok {
		return entry
	}
	return errors.ERROR_BINDINGS_MAP[errorName]
}

func (rIface *resourcesClient) Delete(shareIdParam string, sharedResourceIdParam string) error {
	typeConverter := rIface.connector.TypeConverter()
	executionContext := rIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(resourcesDeleteInputType(), typeConverter)
	sv.AddStructField("ShareId", shareIdParam)
	sv.AddStructField("SharedResourceId", sharedResourceIdParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		return bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := resourcesDeleteRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	rIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := rIface.connector.GetApiProvider().Invoke("com.vmware.nsx_policy.infra.shares.resources", "delete", inputDataValue, executionContext)
	if methodResult.IsSuccess() {
		return nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), rIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return bindings.VAPIerrorsToError(errorInError)
		}
		return methodError.(error)
	}
}

func (rIface *resourcesClient) Get(shareIdParam string, sharedResourceIdParam string) (model.SharedResource, error) {
	typeConverter := rIface.connector.TypeConverter()
	executionContext := rIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(resourcesGetInputType(), typeConverter)
	sv.AddStructField("ShareId", shareIdParam)
	sv.AddStructField("SharedResourceId", sharedResourceIdParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.SharedResource
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := resourcesGetRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	rIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := rIface.connector.GetApiProvider().Invoke("com.vmware.nsx_policy.infra.shares.resources", "get", inputDataValue, executionContext)
	var emptyOutput model.SharedResource
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), resourcesGetOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.SharedResource), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), rIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (rIface *resourcesClient) List(shareIdParam string, resourceTypeParam *string) (model.SharedResourceListResult, error) {
	typeConverter := rIface.connector.TypeConverter()
	executionContext := rIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(resourcesListInputType(), typeConverter)
	sv.AddStructField("ShareId", shareIdParam)
	sv.AddStructField("ResourceType", resourceTypeParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.SharedResourceListResult
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := resourcesListRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	rIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := rIface.connector.GetApiProvider().Invoke("com.vmware.nsx_policy.infra.shares.resources", "list", inputDataValue, executionContext)
	var emptyOutput model.SharedResourceListResult
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), resourcesListOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.SharedResourceListResult), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), rIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (rIface *resourcesClient) Patch(shareIdParam string, sharedResourceIdParam string, sharedResourceParam model.SharedResource) error {
	typeConverter := rIface.connector.TypeConverter()
	executionContext := rIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(resourcesPatchInputType(), typeConverter)
	sv.AddStructField("ShareId", shareIdParam)
	sv.AddStructField("SharedResourceId", sharedResourceIdParam)
	sv.AddStructField("SharedResource", sharedResourceParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		return bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := resourcesPatchRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	rIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := rIface.connector.GetApiProvider().Invoke("com.vmware.nsx_policy.infra.shares.resources", "patch", inputDataValue, executionContext)
	if methodResult.IsSuccess() {
		return nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), rIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return bindings.VAPIerrorsToError(errorInError)
		}
		return methodError.(error)
	}
}

func (rIface *resourcesClient) Update(shareIdParam string, sharedResourceIdParam string, sharedResourceParam model.SharedResource) (model.SharedResource, error) {
	typeConverter := rIface.connector.TypeConverter()
	executionContext := rIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(resourcesUpdateInputType(), typeConverter)
	sv.AddStructField("ShareId", shareIdParam)
	sv.AddStructField("SharedResourceId", sharedResourceIdParam)
	sv.AddStructField("SharedResource", sharedResourceParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.SharedResource
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := resourcesUpdateRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	rIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := rIface.connector.GetApiProvider().Invoke("com.vmware.nsx_policy.infra.shares.resources", "update", inputDataValue, executionContext)
	var emptyOutput model.SharedResource
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), resourcesUpdateOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.SharedResource), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), rIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}
