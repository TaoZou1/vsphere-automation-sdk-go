// Copyright © 2019-2021 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Interface file for service: FirewallIpReputationFeed
// Used by client-side stubs.

package infra

import (
	"github.com/TaoZou1/vsphere-automation-sdk-go/lib/vapi/std/errors"
	"github.com/TaoZou1/vsphere-automation-sdk-go/runtime/bindings"
	"github.com/TaoZou1/vsphere-automation-sdk-go/runtime/core"
	"github.com/TaoZou1/vsphere-automation-sdk-go/runtime/lib"
	"github.com/TaoZou1/vsphere-automation-sdk-go/runtime/protocol/client"
	"github.com/TaoZou1/vsphere-automation-sdk-go/services/nsxt/model"
)

const _ = core.SupportedByRuntimeVersion1

type FirewallIpReputationFeedClient interface {

	// This API can be used to enable or disable auto-download of IP reputation feed, it can also be used to trigger download of IP reputation feed when required. Once auto-download is enabled, IP reputation feed will be downloaded at regular intervals of 720 mins(12 hrs). Auto-download of IP reputation feed can be enabled using the action 'enable_auto_download', to disable use action 'disable_auto_download' and to trigger a download use action 'download'.
	//
	// @param orgIdParam (required)
	// @param projectIdParam (required)
	// @param operationParam action (required)
	// @return com.vmware.nsx_policy.model.PolicyFirewallIpReputationConfig
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Create(orgIdParam string, projectIdParam string, operationParam string) (model.PolicyFirewallIpReputationConfig, error)

	// API to retrieve the current IP Reputation feed configuration.
	//
	// @param orgIdParam (required)
	// @param projectIdParam (required)
	// @return com.vmware.nsx_policy.model.PolicyFirewallIpReputationConfig
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Get(orgIdParam string, projectIdParam string) (model.PolicyFirewallIpReputationConfig, error)
}

type firewallIpReputationFeedClient struct {
	connector           client.Connector
	interfaceDefinition core.InterfaceDefinition
	errorsBindingMap    map[string]bindings.BindingType
}

func NewFirewallIpReputationFeedClient(connector client.Connector) *firewallIpReputationFeedClient {
	interfaceIdentifier := core.NewInterfaceIdentifier("com.vmware.nsx_policy.orgs.projects.infra.firewall_ip_reputation_feed")
	methodIdentifiers := map[string]core.MethodIdentifier{
		"create": core.NewMethodIdentifier(interfaceIdentifier, "create"),
		"get":    core.NewMethodIdentifier(interfaceIdentifier, "get"),
	}
	interfaceDefinition := core.NewInterfaceDefinition(interfaceIdentifier, methodIdentifiers)
	errorsBindingMap := make(map[string]bindings.BindingType)

	fIface := firewallIpReputationFeedClient{interfaceDefinition: interfaceDefinition, errorsBindingMap: errorsBindingMap, connector: connector}
	return &fIface
}

func (fIface *firewallIpReputationFeedClient) GetErrorBindingType(errorName string) bindings.BindingType {
	if entry, ok := fIface.errorsBindingMap[errorName]; ok {
		return entry
	}
	return errors.ERROR_BINDINGS_MAP[errorName]
}

func (fIface *firewallIpReputationFeedClient) Create(orgIdParam string, projectIdParam string, operationParam string) (model.PolicyFirewallIpReputationConfig, error) {
	typeConverter := fIface.connector.TypeConverter()
	executionContext := fIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(firewallIpReputationFeedCreateInputType(), typeConverter)
	sv.AddStructField("OrgId", orgIdParam)
	sv.AddStructField("ProjectId", projectIdParam)
	sv.AddStructField("Operation", operationParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.PolicyFirewallIpReputationConfig
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := firewallIpReputationFeedCreateRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	fIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := fIface.connector.GetApiProvider().Invoke("com.vmware.nsx_policy.orgs.projects.infra.firewall_ip_reputation_feed", "create", inputDataValue, executionContext)
	var emptyOutput model.PolicyFirewallIpReputationConfig
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), firewallIpReputationFeedCreateOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.PolicyFirewallIpReputationConfig), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), fIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (fIface *firewallIpReputationFeedClient) Get(orgIdParam string, projectIdParam string) (model.PolicyFirewallIpReputationConfig, error) {
	typeConverter := fIface.connector.TypeConverter()
	executionContext := fIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(firewallIpReputationFeedGetInputType(), typeConverter)
	sv.AddStructField("OrgId", orgIdParam)
	sv.AddStructField("ProjectId", projectIdParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.PolicyFirewallIpReputationConfig
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := firewallIpReputationFeedGetRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	fIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := fIface.connector.GetApiProvider().Invoke("com.vmware.nsx_policy.orgs.projects.infra.firewall_ip_reputation_feed", "get", inputDataValue, executionContext)
	var emptyOutput model.PolicyFirewallIpReputationConfig
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), firewallIpReputationFeedGetOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.PolicyFirewallIpReputationConfig), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), fIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}
