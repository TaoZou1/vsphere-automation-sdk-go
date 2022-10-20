// Copyright © 2019-2021 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Interface file for service: Search
// Used by client-side stubs.

package oidc_uris

import (
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/lib/vapi/std/errors"
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/runtime/bindings"
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/runtime/core"
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/runtime/lib"
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/runtime/protocol/client"
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/services/nsxt-mp/nsx/model"
)

const _ = core.SupportedByRuntimeVersion1

type SearchClient interface {

	// Search the System for Cross-domain Identity Management (SCIM) source for users and groups whose names match the search string.
	//
	// @param idParam (required)
	// @param searchStringParam Search filter (required)
	// @param cursorParam Opaque cursor to be used for getting next page of records (supplied by current result page) (optional)
	// @param includedFieldsParam Comma separated list of fields that should be included in query result (optional)
	// @param pageSizeParam Maximum number of results to return in this page (server may return fewer) (optional, default to 1000)
	// @param sortAscendingParam (optional)
	// @param sortByParam Field by which records are sorted (optional)
	// @return com.vmware.nsx.model.ScimSearchListResult
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Create(idParam string, searchStringParam string, cursorParam *string, includedFieldsParam *string, pageSizeParam *int64, sortAscendingParam *bool, sortByParam *string) (model.ScimSearchListResult, error)
}

type searchClient struct {
	connector           client.Connector
	interfaceDefinition core.InterfaceDefinition
	errorsBindingMap    map[string]bindings.BindingType
}

func NewSearchClient(connector client.Connector) *searchClient {
	interfaceIdentifier := core.NewInterfaceIdentifier("com.vmware.nsx.trust_management.oidc_uris.search")
	methodIdentifiers := map[string]core.MethodIdentifier{
		"create": core.NewMethodIdentifier(interfaceIdentifier, "create"),
	}
	interfaceDefinition := core.NewInterfaceDefinition(interfaceIdentifier, methodIdentifiers)
	errorsBindingMap := make(map[string]bindings.BindingType)

	sIface := searchClient{interfaceDefinition: interfaceDefinition, errorsBindingMap: errorsBindingMap, connector: connector}
	return &sIface
}

func (sIface *searchClient) GetErrorBindingType(errorName string) bindings.BindingType {
	if entry, ok := sIface.errorsBindingMap[errorName]; ok {
		return entry
	}
	return errors.ERROR_BINDINGS_MAP[errorName]
}

func (sIface *searchClient) Create(idParam string, searchStringParam string, cursorParam *string, includedFieldsParam *string, pageSizeParam *int64, sortAscendingParam *bool, sortByParam *string) (model.ScimSearchListResult, error) {
	typeConverter := sIface.connector.TypeConverter()
	executionContext := sIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(searchCreateInputType(), typeConverter)
	sv.AddStructField("Id", idParam)
	sv.AddStructField("SearchString", searchStringParam)
	sv.AddStructField("Cursor", cursorParam)
	sv.AddStructField("IncludedFields", includedFieldsParam)
	sv.AddStructField("PageSize", pageSizeParam)
	sv.AddStructField("SortAscending", sortAscendingParam)
	sv.AddStructField("SortBy", sortByParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.ScimSearchListResult
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := searchCreateRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	sIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := sIface.connector.GetApiProvider().Invoke("com.vmware.nsx.trust_management.oidc_uris.search", "create", inputDataValue, executionContext)
	var emptyOutput model.ScimSearchListResult
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), searchCreateOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.ScimSearchListResult), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), sIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}
