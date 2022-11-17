// Copyright © 2019-2021 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Interface file for service: RoleBindings
// Used by client-side stubs.

package aaa

import (
	"github.com/vmware/vsphere-automation-sdk-go/lib/vapi/std/errors"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/bindings"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/core"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/lib"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/protocol/client"
	"github.com/vmware/vsphere-automation-sdk-go/services/nsxt/model"
)

const _ = core.SupportedByRuntimeVersion1

type RoleBindingsClient interface {

	// Delete the user/group's role assignment on CSP. If the path is provided then deletes only the roles_for_paths that matches the path. If path is provided for the last roles_for_paths then the whole role binding is deleted provided it is not that of a local user.
	//
	// @param orgIdParam (required)
	// @param projectIdParam (required)
	// @param vpcIdParam (required)
	// @param bindingIdParam (required)
	// @param cursorParam Opaque cursor to be used for getting next page of records (supplied by current result page) (optional)
	// @param identitySourceIdParam Identity source ID (optional)
	// @param identitySourceTypeParam Identity source type (optional)
	// @param includedFieldsParam Comma separated list of fields that should be included in query result (optional)
	// @param nameParam User/Group name (optional)
	// @param pageSizeParam Maximum number of results to return in this page (server may return fewer) (optional, default to 1000)
	// @param pathParam Exact path of the context (optional)
	// @param roleParam Role ID (optional)
	// @param rootPathParam Prefix path of the context (optional)
	// @param sortAscendingParam (optional)
	// @param sortByParam Field by which records are sorted (optional)
	// @param type_Param Type (optional)
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Delete(orgIdParam string, projectIdParam string, vpcIdParam string, bindingIdParam string, cursorParam *string, identitySourceIdParam *string, identitySourceTypeParam *string, includedFieldsParam *string, nameParam *string, pageSizeParam *int64, pathParam *string, roleParam *string, rootPathParam *string, sortAscendingParam *bool, sortByParam *string, type_Param *string) error

	// Get user/group's role information from CSP
	//
	// @param orgIdParam (required)
	// @param projectIdParam (required)
	// @param vpcIdParam (required)
	// @param bindingIdParam (required)
	// @param cursorParam Opaque cursor to be used for getting next page of records (supplied by current result page) (optional)
	// @param identitySourceIdParam Identity source ID (optional)
	// @param identitySourceTypeParam Identity source type (optional)
	// @param includedFieldsParam Comma separated list of fields that should be included in query result (optional)
	// @param nameParam User/Group name (optional)
	// @param pageSizeParam Maximum number of results to return in this page (server may return fewer) (optional, default to 1000)
	// @param pathParam Exact path of the context (optional)
	// @param roleParam Role ID (optional)
	// @param rootPathParam Prefix path of the context (optional)
	// @param sortAscendingParam (optional)
	// @param sortByParam Field by which records are sorted (optional)
	// @param type_Param Type (optional)
	// @return com.vmware.nsx_policy.model.RoleBinding
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Get(orgIdParam string, projectIdParam string, vpcIdParam string, bindingIdParam string, cursorParam *string, identitySourceIdParam *string, identitySourceTypeParam *string, includedFieldsParam *string, nameParam *string, pageSizeParam *int64, pathParam *string, roleParam *string, rootPathParam *string, sortAscendingParam *bool, sortByParam *string, type_Param *string) (model.RoleBinding, error)

	// Get all users and groups with their roles from CSP. If the root_path is provided then only return role bindings that start-with or are sub-trees of the provided root path. Also filter the roles_for_paths such that only those roles_for_paths appear that start-with or are sub-tree of the provided root path.
	//
	// @param orgIdParam (required)
	// @param projectIdParam (required)
	// @param vpcIdParam (required)
	// @param cursorParam Opaque cursor to be used for getting next page of records (supplied by current result page) (optional)
	// @param identitySourceIdParam Identity source ID (optional)
	// @param identitySourceTypeParam Identity source type (optional)
	// @param includedFieldsParam Comma separated list of fields that should be included in query result (optional)
	// @param nameParam User/Group name (optional)
	// @param pageSizeParam Maximum number of results to return in this page (server may return fewer) (optional, default to 1000)
	// @param pathParam Exact path of the context (optional)
	// @param roleParam Role ID (optional)
	// @param rootPathParam Prefix path of the context (optional)
	// @param sortAscendingParam (optional)
	// @param sortByParam Field by which records are sorted (optional)
	// @param type_Param Type (optional)
	// @return com.vmware.nsx_policy.model.RoleBindingListResult
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	List(orgIdParam string, projectIdParam string, vpcIdParam string, cursorParam *string, identitySourceIdParam *string, identitySourceTypeParam *string, includedFieldsParam *string, nameParam *string, pageSizeParam *int64, pathParam *string, roleParam *string, rootPathParam *string, sortAscendingParam *bool, sortByParam *string, type_Param *string) (model.RoleBindingListResult, error)
}

type roleBindingsClient struct {
	connector           client.Connector
	interfaceDefinition core.InterfaceDefinition
	errorsBindingMap    map[string]bindings.BindingType
}

func NewRoleBindingsClient(connector client.Connector) *roleBindingsClient {
	interfaceIdentifier := core.NewInterfaceIdentifier("com.vmware.nsx_policy.orgs.projects.vpcs.aaa.role_bindings")
	methodIdentifiers := map[string]core.MethodIdentifier{
		"delete": core.NewMethodIdentifier(interfaceIdentifier, "delete"),
		"get":    core.NewMethodIdentifier(interfaceIdentifier, "get"),
		"list":   core.NewMethodIdentifier(interfaceIdentifier, "list"),
	}
	interfaceDefinition := core.NewInterfaceDefinition(interfaceIdentifier, methodIdentifiers)
	errorsBindingMap := make(map[string]bindings.BindingType)

	rIface := roleBindingsClient{interfaceDefinition: interfaceDefinition, errorsBindingMap: errorsBindingMap, connector: connector}
	return &rIface
}

func (rIface *roleBindingsClient) GetErrorBindingType(errorName string) bindings.BindingType {
	if entry, ok := rIface.errorsBindingMap[errorName]; ok {
		return entry
	}
	return errors.ERROR_BINDINGS_MAP[errorName]
}

func (rIface *roleBindingsClient) Delete(orgIdParam string, projectIdParam string, vpcIdParam string, bindingIdParam string, cursorParam *string, identitySourceIdParam *string, identitySourceTypeParam *string, includedFieldsParam *string, nameParam *string, pageSizeParam *int64, pathParam *string, roleParam *string, rootPathParam *string, sortAscendingParam *bool, sortByParam *string, type_Param *string) error {
	typeConverter := rIface.connector.TypeConverter()
	executionContext := rIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(roleBindingsDeleteInputType(), typeConverter)
	sv.AddStructField("OrgId", orgIdParam)
	sv.AddStructField("ProjectId", projectIdParam)
	sv.AddStructField("VpcId", vpcIdParam)
	sv.AddStructField("BindingId", bindingIdParam)
	sv.AddStructField("Cursor", cursorParam)
	sv.AddStructField("IdentitySourceId", identitySourceIdParam)
	sv.AddStructField("IdentitySourceType", identitySourceTypeParam)
	sv.AddStructField("IncludedFields", includedFieldsParam)
	sv.AddStructField("Name", nameParam)
	sv.AddStructField("PageSize", pageSizeParam)
	sv.AddStructField("Path", pathParam)
	sv.AddStructField("Role", roleParam)
	sv.AddStructField("RootPath", rootPathParam)
	sv.AddStructField("SortAscending", sortAscendingParam)
	sv.AddStructField("SortBy", sortByParam)
	sv.AddStructField("Type_", type_Param)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		return bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := roleBindingsDeleteRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	rIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := rIface.connector.GetApiProvider().Invoke("com.vmware.nsx_policy.orgs.projects.vpcs.aaa.role_bindings", "delete", inputDataValue, executionContext)
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

func (rIface *roleBindingsClient) Get(orgIdParam string, projectIdParam string, vpcIdParam string, bindingIdParam string, cursorParam *string, identitySourceIdParam *string, identitySourceTypeParam *string, includedFieldsParam *string, nameParam *string, pageSizeParam *int64, pathParam *string, roleParam *string, rootPathParam *string, sortAscendingParam *bool, sortByParam *string, type_Param *string) (model.RoleBinding, error) {
	typeConverter := rIface.connector.TypeConverter()
	executionContext := rIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(roleBindingsGetInputType(), typeConverter)
	sv.AddStructField("OrgId", orgIdParam)
	sv.AddStructField("ProjectId", projectIdParam)
	sv.AddStructField("VpcId", vpcIdParam)
	sv.AddStructField("BindingId", bindingIdParam)
	sv.AddStructField("Cursor", cursorParam)
	sv.AddStructField("IdentitySourceId", identitySourceIdParam)
	sv.AddStructField("IdentitySourceType", identitySourceTypeParam)
	sv.AddStructField("IncludedFields", includedFieldsParam)
	sv.AddStructField("Name", nameParam)
	sv.AddStructField("PageSize", pageSizeParam)
	sv.AddStructField("Path", pathParam)
	sv.AddStructField("Role", roleParam)
	sv.AddStructField("RootPath", rootPathParam)
	sv.AddStructField("SortAscending", sortAscendingParam)
	sv.AddStructField("SortBy", sortByParam)
	sv.AddStructField("Type_", type_Param)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.RoleBinding
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := roleBindingsGetRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	rIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := rIface.connector.GetApiProvider().Invoke("com.vmware.nsx_policy.orgs.projects.vpcs.aaa.role_bindings", "get", inputDataValue, executionContext)
	var emptyOutput model.RoleBinding
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), roleBindingsGetOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.RoleBinding), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), rIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (rIface *roleBindingsClient) List(orgIdParam string, projectIdParam string, vpcIdParam string, cursorParam *string, identitySourceIdParam *string, identitySourceTypeParam *string, includedFieldsParam *string, nameParam *string, pageSizeParam *int64, pathParam *string, roleParam *string, rootPathParam *string, sortAscendingParam *bool, sortByParam *string, type_Param *string) (model.RoleBindingListResult, error) {
	typeConverter := rIface.connector.TypeConverter()
	executionContext := rIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(roleBindingsListInputType(), typeConverter)
	sv.AddStructField("OrgId", orgIdParam)
	sv.AddStructField("ProjectId", projectIdParam)
	sv.AddStructField("VpcId", vpcIdParam)
	sv.AddStructField("Cursor", cursorParam)
	sv.AddStructField("IdentitySourceId", identitySourceIdParam)
	sv.AddStructField("IdentitySourceType", identitySourceTypeParam)
	sv.AddStructField("IncludedFields", includedFieldsParam)
	sv.AddStructField("Name", nameParam)
	sv.AddStructField("PageSize", pageSizeParam)
	sv.AddStructField("Path", pathParam)
	sv.AddStructField("Role", roleParam)
	sv.AddStructField("RootPath", rootPathParam)
	sv.AddStructField("SortAscending", sortAscendingParam)
	sv.AddStructField("SortBy", sortByParam)
	sv.AddStructField("Type_", type_Param)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.RoleBindingListResult
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := roleBindingsListRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	rIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := rIface.connector.GetApiProvider().Invoke("com.vmware.nsx_policy.orgs.projects.vpcs.aaa.role_bindings", "list", inputDataValue, executionContext)
	var emptyOutput model.RoleBindingListResult
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), roleBindingsListOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.RoleBindingListResult), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), rIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}
