// Copyright © 2019-2021 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Interface file for service: Crls
// Used by client-side stubs.

package trust_management

import (
	"github.com/vmware/vsphere-automation-sdk-go/lib/vapi/std/errors"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/bindings"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/core"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/lib"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/protocol/client"
	"github.com/vmware/vsphere-automation-sdk-go/services/nsxt-mp/nsx/model"
)

const _ = core.SupportedByRuntimeVersion1

type CrlsClient interface {

	// Deletes an existing CRL.
	//
	// @param crlIdParam ID of CRL to delete (required)
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Delete(crlIdParam string) error

	// Returns information about the specified CRL. For additional information, include the ?details=true modifier at the end of the request URI.
	//
	// @param crlIdParam ID of CRL to read (required)
	// @param detailsParam whether to expand the pem data and show all its details (optional, default to false)
	// @return com.vmware.nsx.model.Crl
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Get(crlIdParam string, detailsParam *bool) (model.Crl, error)

	// Adds a new certificate revocation list (CRL). The CRL is used to verify the client certificate status against the revocation lists published by the CA. For this reason, the administrator needs to add the CRL in certificate repository as well. A CRL can be in the PEM X.509 format (crl_type=X509) or JSON OneCRL (crl_type=OneCRL). If crl_type is not specified, it is auto-detected based on the presence of fields pem_encoded or one_crl.
	//
	// @param crlObjectDataParam (required)
	// @return com.vmware.nsx.model.CrlList
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Importcrl(crlObjectDataParam model.CrlObjectData) (model.CrlList, error)

	// Returns information about all CRLs. For additional information, include the ?details=true modifier at the end of the request URI.
	//
	// @param cursorParam Opaque cursor to be used for getting next page of records (supplied by current result page) (optional)
	// @param detailsParam whether to expand the pem data and show all its details (optional, default to false)
	// @param includedFieldsParam Comma separated list of fields that should be included in query result (optional)
	// @param nodeIdParam Node ID of certificate to return (optional)
	// @param pageSizeParam Maximum number of results to return in this page (server may return fewer) (optional, default to 1000)
	// @param sortAscendingParam (optional)
	// @param sortByParam Field by which records are sorted (optional)
	// @param type_Param Type of certificate to return (optional)
	// @return com.vmware.nsx.model.CrlList
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	List(cursorParam *string, detailsParam *bool, includedFieldsParam *string, nodeIdParam *string, pageSizeParam *int64, sortAscendingParam *bool, sortByParam *string, type_Param *string) (model.CrlList, error)

	// Updates an existing CRL.
	//
	// @param crlIdParam ID of CRL to update (required)
	// @param crlParam (required)
	// @return com.vmware.nsx.model.Crl
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Update(crlIdParam string, crlParam model.Crl) (model.Crl, error)
}

type crlsClient struct {
	connector           client.Connector
	interfaceDefinition core.InterfaceDefinition
	errorsBindingMap    map[string]bindings.BindingType
}

func NewCrlsClient(connector client.Connector) *crlsClient {
	interfaceIdentifier := core.NewInterfaceIdentifier("com.vmware.nsx.trust_management.crls")
	methodIdentifiers := map[string]core.MethodIdentifier{
		"delete":    core.NewMethodIdentifier(interfaceIdentifier, "delete"),
		"get":       core.NewMethodIdentifier(interfaceIdentifier, "get"),
		"importcrl": core.NewMethodIdentifier(interfaceIdentifier, "importcrl"),
		"list":      core.NewMethodIdentifier(interfaceIdentifier, "list"),
		"update":    core.NewMethodIdentifier(interfaceIdentifier, "update"),
	}
	interfaceDefinition := core.NewInterfaceDefinition(interfaceIdentifier, methodIdentifiers)
	errorsBindingMap := make(map[string]bindings.BindingType)

	cIface := crlsClient{interfaceDefinition: interfaceDefinition, errorsBindingMap: errorsBindingMap, connector: connector}
	return &cIface
}

func (cIface *crlsClient) GetErrorBindingType(errorName string) bindings.BindingType {
	if entry, ok := cIface.errorsBindingMap[errorName]; ok {
		return entry
	}
	return errors.ERROR_BINDINGS_MAP[errorName]
}

func (cIface *crlsClient) Delete(crlIdParam string) error {
	typeConverter := cIface.connector.TypeConverter()
	executionContext := cIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(crlsDeleteInputType(), typeConverter)
	sv.AddStructField("CrlId", crlIdParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		return bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := crlsDeleteRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	cIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := cIface.connector.GetApiProvider().Invoke("com.vmware.nsx.trust_management.crls", "delete", inputDataValue, executionContext)
	if methodResult.IsSuccess() {
		return nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), cIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return bindings.VAPIerrorsToError(errorInError)
		}
		return methodError.(error)
	}
}

func (cIface *crlsClient) Get(crlIdParam string, detailsParam *bool) (model.Crl, error) {
	typeConverter := cIface.connector.TypeConverter()
	executionContext := cIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(crlsGetInputType(), typeConverter)
	sv.AddStructField("CrlId", crlIdParam)
	sv.AddStructField("Details", detailsParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.Crl
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := crlsGetRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	cIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := cIface.connector.GetApiProvider().Invoke("com.vmware.nsx.trust_management.crls", "get", inputDataValue, executionContext)
	var emptyOutput model.Crl
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), crlsGetOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.Crl), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), cIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (cIface *crlsClient) Importcrl(crlObjectDataParam model.CrlObjectData) (model.CrlList, error) {
	typeConverter := cIface.connector.TypeConverter()
	executionContext := cIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(crlsImportcrlInputType(), typeConverter)
	sv.AddStructField("CrlObjectData", crlObjectDataParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.CrlList
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := crlsImportcrlRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	cIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := cIface.connector.GetApiProvider().Invoke("com.vmware.nsx.trust_management.crls", "importcrl", inputDataValue, executionContext)
	var emptyOutput model.CrlList
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), crlsImportcrlOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.CrlList), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), cIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (cIface *crlsClient) List(cursorParam *string, detailsParam *bool, includedFieldsParam *string, nodeIdParam *string, pageSizeParam *int64, sortAscendingParam *bool, sortByParam *string, type_Param *string) (model.CrlList, error) {
	typeConverter := cIface.connector.TypeConverter()
	executionContext := cIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(crlsListInputType(), typeConverter)
	sv.AddStructField("Cursor", cursorParam)
	sv.AddStructField("Details", detailsParam)
	sv.AddStructField("IncludedFields", includedFieldsParam)
	sv.AddStructField("NodeId", nodeIdParam)
	sv.AddStructField("PageSize", pageSizeParam)
	sv.AddStructField("SortAscending", sortAscendingParam)
	sv.AddStructField("SortBy", sortByParam)
	sv.AddStructField("Type_", type_Param)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.CrlList
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := crlsListRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	cIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := cIface.connector.GetApiProvider().Invoke("com.vmware.nsx.trust_management.crls", "list", inputDataValue, executionContext)
	var emptyOutput model.CrlList
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), crlsListOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.CrlList), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), cIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (cIface *crlsClient) Update(crlIdParam string, crlParam model.Crl) (model.Crl, error) {
	typeConverter := cIface.connector.TypeConverter()
	executionContext := cIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(crlsUpdateInputType(), typeConverter)
	sv.AddStructField("CrlId", crlIdParam)
	sv.AddStructField("Crl", crlParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.Crl
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := crlsUpdateRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	cIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := cIface.connector.GetApiProvider().Invoke("com.vmware.nsx.trust_management.crls", "update", inputDataValue, executionContext)
	var emptyOutput model.Crl
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), crlsUpdateOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.Crl), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), cIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}
