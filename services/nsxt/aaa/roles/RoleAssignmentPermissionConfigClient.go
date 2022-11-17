// Copyright © 2019-2021 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Interface file for service: RoleAssignmentPermissionConfig
// Used by client-side stubs.

package roles

import (
	"github.com/vmware/vsphere-automation-sdk-go/lib/vapi/std/errors"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/bindings"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/core"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/lib"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/protocol/client"
	"github.com/vmware/vsphere-automation-sdk-go/services/nsxt/model"
)

const _ = core.SupportedByRuntimeVersion1

type RoleAssignmentPermissionConfigClient interface {

	// Provides a means to allow or disallow project administrators and VPC administrators to assign roles to other users on projects and VPCs.
	//
	// @param roleParam role id (either \"project_admin\" or \"vpc_admin\") (required)
	// @param roleAssignmentPermissionConfigParam (required)
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Update(roleParam string, roleAssignmentPermissionConfigParam model.RoleAssignmentPermissionConfig) error
}

type roleAssignmentPermissionConfigClient struct {
	connector           client.Connector
	interfaceDefinition core.InterfaceDefinition
	errorsBindingMap    map[string]bindings.BindingType
}

func NewRoleAssignmentPermissionConfigClient(connector client.Connector) *roleAssignmentPermissionConfigClient {
	interfaceIdentifier := core.NewInterfaceIdentifier("com.vmware.nsx_policy.aaa.roles.role_assignment_permission_config")
	methodIdentifiers := map[string]core.MethodIdentifier{
		"update": core.NewMethodIdentifier(interfaceIdentifier, "update"),
	}
	interfaceDefinition := core.NewInterfaceDefinition(interfaceIdentifier, methodIdentifiers)
	errorsBindingMap := make(map[string]bindings.BindingType)

	rIface := roleAssignmentPermissionConfigClient{interfaceDefinition: interfaceDefinition, errorsBindingMap: errorsBindingMap, connector: connector}
	return &rIface
}

func (rIface *roleAssignmentPermissionConfigClient) GetErrorBindingType(errorName string) bindings.BindingType {
	if entry, ok := rIface.errorsBindingMap[errorName]; ok {
		return entry
	}
	return errors.ERROR_BINDINGS_MAP[errorName]
}

func (rIface *roleAssignmentPermissionConfigClient) Update(roleParam string, roleAssignmentPermissionConfigParam model.RoleAssignmentPermissionConfig) error {
	typeConverter := rIface.connector.TypeConverter()
	executionContext := rIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(roleAssignmentPermissionConfigUpdateInputType(), typeConverter)
	sv.AddStructField("Role", roleParam)
	sv.AddStructField("RoleAssignmentPermissionConfig", roleAssignmentPermissionConfigParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		return bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := roleAssignmentPermissionConfigUpdateRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	rIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := rIface.connector.GetApiProvider().Invoke("com.vmware.nsx_policy.aaa.roles.role_assignment_permission_config", "update", inputDataValue, executionContext)
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
