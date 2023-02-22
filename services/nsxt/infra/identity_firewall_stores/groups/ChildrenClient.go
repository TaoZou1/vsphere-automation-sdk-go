// Copyright © 2019-2021 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Interface file for service: Children
// Used by client-side stubs.

package groups

import (
	"github.com/vmware/vsphere-automation-sdk-go/lib/vapi/std/errors"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/bindings"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/core"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/lib"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/protocol/client"
	"github.com/vmware/vsphere-automation-sdk-go/services/nsxt/model"
)

const _ = core.SupportedByRuntimeVersion1

type ChildrenClient interface {

	// A member group could be either direct member of the group specified by group_id or nested member of it. Both direct member groups and nested member groups are returned.
	//
	// @param identityFirewallStoreIdParam Firewall Identity store identifier (required)
	// @param groupIdParam Directory group identifier (required)
	// @param includeUsersParam include users filter (required)
	// @return com.vmware.nsx_policy.model.IdentityFirewallStoreUserOrGrouptListResults
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	List(identityFirewallStoreIdParam string, groupIdParam string, includeUsersParam bool) (model.IdentityFirewallStoreUserOrGrouptListResults, error)
}

type childrenClient struct {
	connector           client.Connector
	interfaceDefinition core.InterfaceDefinition
	errorsBindingMap    map[string]bindings.BindingType
}

func NewChildrenClient(connector client.Connector) *childrenClient {
	interfaceIdentifier := core.NewInterfaceIdentifier("com.vmware.nsx_policy.infra.identity_firewall_stores.groups.children")
	methodIdentifiers := map[string]core.MethodIdentifier{
		"list": core.NewMethodIdentifier(interfaceIdentifier, "list"),
	}
	interfaceDefinition := core.NewInterfaceDefinition(interfaceIdentifier, methodIdentifiers)
	errorsBindingMap := make(map[string]bindings.BindingType)

	cIface := childrenClient{interfaceDefinition: interfaceDefinition, errorsBindingMap: errorsBindingMap, connector: connector}
	return &cIface
}

func (cIface *childrenClient) GetErrorBindingType(errorName string) bindings.BindingType {
	if entry, ok := cIface.errorsBindingMap[errorName]; ok {
		return entry
	}
	return errors.ERROR_BINDINGS_MAP[errorName]
}

func (cIface *childrenClient) List(identityFirewallStoreIdParam string, groupIdParam string, includeUsersParam bool) (model.IdentityFirewallStoreUserOrGrouptListResults, error) {
	typeConverter := cIface.connector.TypeConverter()
	executionContext := cIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(childrenListInputType(), typeConverter)
	sv.AddStructField("IdentityFirewallStoreId", identityFirewallStoreIdParam)
	sv.AddStructField("GroupId", groupIdParam)
	sv.AddStructField("IncludeUsers", includeUsersParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.IdentityFirewallStoreUserOrGrouptListResults
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := childrenListRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	cIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := cIface.connector.GetApiProvider().Invoke("com.vmware.nsx_policy.infra.identity_firewall_stores.groups.children", "list", inputDataValue, executionContext)
	var emptyOutput model.IdentityFirewallStoreUserOrGrouptListResults
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), childrenListOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.IdentityFirewallStoreUserOrGrouptListResults), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), cIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}
