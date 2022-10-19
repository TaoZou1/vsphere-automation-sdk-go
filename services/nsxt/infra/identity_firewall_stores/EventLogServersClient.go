// Copyright © 2019-2021 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Interface file for service: EventLogServers
// Used by client-side stubs.

package identity_firewall_stores

import (
	"github.com/TaoZou1/vsphere-automation-sdk-go/lib/vapi/std/errors"
	"github.com/TaoZou1/vsphere-automation-sdk-go/runtime/bindings"
	"github.com/TaoZou1/vsphere-automation-sdk-go/runtime/core"
	"github.com/TaoZou1/vsphere-automation-sdk-go/runtime/lib"
	"github.com/TaoZou1/vsphere-automation-sdk-go/runtime/protocol/client"
	"github.com/TaoZou1/vsphere-automation-sdk-go/services/nsxt/model"
)

const _ = core.SupportedByRuntimeVersion1

type EventLogServersClient interface {

	// Delete a Event Log server for Firewall Identity store
	//
	// @param identityFirewallStoreIdParam Firewall Identity store identifier (required)
	// @param eventLogServerIdParam Event Log server identifier (required)
	// @param enforcementPointPathParam String Path of the enforcement point (optional)
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Delete(identityFirewallStoreIdParam string, eventLogServerIdParam string, enforcementPointPathParam *string) error

	// Get a specific Event Log server for a given Firewall Identity store
	//
	// @param identityFirewallStoreIdParam Firewall Identity store identifier (required)
	// @param eventLogServerIdParam Event Log server identifier (required)
	// @param enforcementPointPathParam String Path of the enforcement point (optional)
	// @return com.vmware.nsx_policy.model.IdentityFirewallStoreEventLogServer
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Get(identityFirewallStoreIdParam string, eventLogServerIdParam string, enforcementPointPathParam *string) (model.IdentityFirewallStoreEventLogServer, error)

	// More than one Event Log server can be created and only one event log server is used to synchronize directory objects. If more than one Event Log server is configured, NSX will try all the servers until it is able to successfully connect to one.
	//
	// @param identityFirewallStoreIdParam Firewall Identity store identifier (required)
	// @param eventLogServerIdParam Event Log server identifier (required)
	// @param identityFirewallStoreEventLogServerParam (required)
	// @param enforcementPointPathParam String Path of the enforcement point (optional)
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Patch(identityFirewallStoreIdParam string, eventLogServerIdParam string, identityFirewallStoreEventLogServerParam model.IdentityFirewallStoreEventLogServer, enforcementPointPathParam *string) error

	// Update a event log server for Firewall Identity store
	//
	// @param identityFirewallStoreIdParam Firewall Identity store identifier (required)
	// @param eventLogServerIdParam Event Log Server identifier (required)
	// @param identityFirewallStoreEventLogServerParam (required)
	// @param enforcementPointPathParam String Path of the enforcement point (optional)
	// @return com.vmware.nsx_policy.model.IdentityFirewallStoreEventLogServer
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Update(identityFirewallStoreIdParam string, eventLogServerIdParam string, identityFirewallStoreEventLogServerParam model.IdentityFirewallStoreEventLogServer, enforcementPointPathParam *string) (model.IdentityFirewallStoreEventLogServer, error)
}

type eventLogServersClient struct {
	connector           client.Connector
	interfaceDefinition core.InterfaceDefinition
	errorsBindingMap    map[string]bindings.BindingType
}

func NewEventLogServersClient(connector client.Connector) *eventLogServersClient {
	interfaceIdentifier := core.NewInterfaceIdentifier("com.vmware.nsx_policy.infra.identity_firewall_stores.event_log_servers")
	methodIdentifiers := map[string]core.MethodIdentifier{
		"delete": core.NewMethodIdentifier(interfaceIdentifier, "delete"),
		"get":    core.NewMethodIdentifier(interfaceIdentifier, "get"),
		"patch":  core.NewMethodIdentifier(interfaceIdentifier, "patch"),
		"update": core.NewMethodIdentifier(interfaceIdentifier, "update"),
	}
	interfaceDefinition := core.NewInterfaceDefinition(interfaceIdentifier, methodIdentifiers)
	errorsBindingMap := make(map[string]bindings.BindingType)

	eIface := eventLogServersClient{interfaceDefinition: interfaceDefinition, errorsBindingMap: errorsBindingMap, connector: connector}
	return &eIface
}

func (eIface *eventLogServersClient) GetErrorBindingType(errorName string) bindings.BindingType {
	if entry, ok := eIface.errorsBindingMap[errorName]; ok {
		return entry
	}
	return errors.ERROR_BINDINGS_MAP[errorName]
}

func (eIface *eventLogServersClient) Delete(identityFirewallStoreIdParam string, eventLogServerIdParam string, enforcementPointPathParam *string) error {
	typeConverter := eIface.connector.TypeConverter()
	executionContext := eIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(eventLogServersDeleteInputType(), typeConverter)
	sv.AddStructField("IdentityFirewallStoreId", identityFirewallStoreIdParam)
	sv.AddStructField("EventLogServerId", eventLogServerIdParam)
	sv.AddStructField("EnforcementPointPath", enforcementPointPathParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		return bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := eventLogServersDeleteRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	eIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := eIface.connector.GetApiProvider().Invoke("com.vmware.nsx_policy.infra.identity_firewall_stores.event_log_servers", "delete", inputDataValue, executionContext)
	if methodResult.IsSuccess() {
		return nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), eIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return bindings.VAPIerrorsToError(errorInError)
		}
		return methodError.(error)
	}
}

func (eIface *eventLogServersClient) Get(identityFirewallStoreIdParam string, eventLogServerIdParam string, enforcementPointPathParam *string) (model.IdentityFirewallStoreEventLogServer, error) {
	typeConverter := eIface.connector.TypeConverter()
	executionContext := eIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(eventLogServersGetInputType(), typeConverter)
	sv.AddStructField("IdentityFirewallStoreId", identityFirewallStoreIdParam)
	sv.AddStructField("EventLogServerId", eventLogServerIdParam)
	sv.AddStructField("EnforcementPointPath", enforcementPointPathParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.IdentityFirewallStoreEventLogServer
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := eventLogServersGetRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	eIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := eIface.connector.GetApiProvider().Invoke("com.vmware.nsx_policy.infra.identity_firewall_stores.event_log_servers", "get", inputDataValue, executionContext)
	var emptyOutput model.IdentityFirewallStoreEventLogServer
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), eventLogServersGetOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.IdentityFirewallStoreEventLogServer), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), eIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (eIface *eventLogServersClient) Patch(identityFirewallStoreIdParam string, eventLogServerIdParam string, identityFirewallStoreEventLogServerParam model.IdentityFirewallStoreEventLogServer, enforcementPointPathParam *string) error {
	typeConverter := eIface.connector.TypeConverter()
	executionContext := eIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(eventLogServersPatchInputType(), typeConverter)
	sv.AddStructField("IdentityFirewallStoreId", identityFirewallStoreIdParam)
	sv.AddStructField("EventLogServerId", eventLogServerIdParam)
	sv.AddStructField("IdentityFirewallStoreEventLogServer", identityFirewallStoreEventLogServerParam)
	sv.AddStructField("EnforcementPointPath", enforcementPointPathParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		return bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := eventLogServersPatchRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	eIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := eIface.connector.GetApiProvider().Invoke("com.vmware.nsx_policy.infra.identity_firewall_stores.event_log_servers", "patch", inputDataValue, executionContext)
	if methodResult.IsSuccess() {
		return nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), eIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return bindings.VAPIerrorsToError(errorInError)
		}
		return methodError.(error)
	}
}

func (eIface *eventLogServersClient) Update(identityFirewallStoreIdParam string, eventLogServerIdParam string, identityFirewallStoreEventLogServerParam model.IdentityFirewallStoreEventLogServer, enforcementPointPathParam *string) (model.IdentityFirewallStoreEventLogServer, error) {
	typeConverter := eIface.connector.TypeConverter()
	executionContext := eIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(eventLogServersUpdateInputType(), typeConverter)
	sv.AddStructField("IdentityFirewallStoreId", identityFirewallStoreIdParam)
	sv.AddStructField("EventLogServerId", eventLogServerIdParam)
	sv.AddStructField("IdentityFirewallStoreEventLogServer", identityFirewallStoreEventLogServerParam)
	sv.AddStructField("EnforcementPointPath", enforcementPointPathParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.IdentityFirewallStoreEventLogServer
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := eventLogServersUpdateRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	eIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := eIface.connector.GetApiProvider().Invoke("com.vmware.nsx_policy.infra.identity_firewall_stores.event_log_servers", "update", inputDataValue, executionContext)
	var emptyOutput model.IdentityFirewallStoreEventLogServer
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), eventLogServersUpdateOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.IdentityFirewallStoreEventLogServer), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), eIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}
