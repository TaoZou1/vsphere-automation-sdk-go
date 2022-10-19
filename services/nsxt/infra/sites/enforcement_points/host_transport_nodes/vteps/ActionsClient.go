// Copyright © 2019-2021 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Interface file for service: Actions
// Used by client-side stubs.

package vteps

import (
	"github.com/TaoZou1/vsphere-automation-sdk-go/lib/vapi/std/errors"
	"github.com/TaoZou1/vsphere-automation-sdk-go/runtime/bindings"
	"github.com/TaoZou1/vsphere-automation-sdk-go/runtime/core"
	"github.com/TaoZou1/vsphere-automation-sdk-go/runtime/data"
	"github.com/TaoZou1/vsphere-automation-sdk-go/runtime/lib"
	"github.com/TaoZou1/vsphere-automation-sdk-go/runtime/protocol/client"
	"github.com/TaoZou1/vsphere-automation-sdk-go/services/nsxt/model"
)

const _ = core.SupportedByRuntimeVersion1

type ActionsClient interface {

	// Submit a new VTEP action for a particular TransportNode. The status of submitted actions could be retrieved using the ListTransportNodeVtepActionsStatus API.
	//
	// @param siteIdParam (required)
	// @param enforcementpointIdParam (required)
	// @param hostTransportNodeIdParam (required)
	// @param transportNodeVtepActionParam (required)
	// The parameter must contain all the properties defined in model.TransportNodeVtepAction.
	// @return com.vmware.nsx_policy.model.TransportNodeVtepActionReference
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Create(siteIdParam string, enforcementpointIdParam string, hostTransportNodeIdParam string, transportNodeVtepActionParam *data.StructValue) (model.TransportNodeVtepActionReference, error)
}

type actionsClient struct {
	connector           client.Connector
	interfaceDefinition core.InterfaceDefinition
	errorsBindingMap    map[string]bindings.BindingType
}

func NewActionsClient(connector client.Connector) *actionsClient {
	interfaceIdentifier := core.NewInterfaceIdentifier("com.vmware.nsx_policy.infra.sites.enforcement_points.host_transport_nodes.vteps.actions")
	methodIdentifiers := map[string]core.MethodIdentifier{
		"create": core.NewMethodIdentifier(interfaceIdentifier, "create"),
	}
	interfaceDefinition := core.NewInterfaceDefinition(interfaceIdentifier, methodIdentifiers)
	errorsBindingMap := make(map[string]bindings.BindingType)

	aIface := actionsClient{interfaceDefinition: interfaceDefinition, errorsBindingMap: errorsBindingMap, connector: connector}
	return &aIface
}

func (aIface *actionsClient) GetErrorBindingType(errorName string) bindings.BindingType {
	if entry, ok := aIface.errorsBindingMap[errorName]; ok {
		return entry
	}
	return errors.ERROR_BINDINGS_MAP[errorName]
}

func (aIface *actionsClient) Create(siteIdParam string, enforcementpointIdParam string, hostTransportNodeIdParam string, transportNodeVtepActionParam *data.StructValue) (model.TransportNodeVtepActionReference, error) {
	typeConverter := aIface.connector.TypeConverter()
	executionContext := aIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(actionsCreateInputType(), typeConverter)
	sv.AddStructField("SiteId", siteIdParam)
	sv.AddStructField("EnforcementpointId", enforcementpointIdParam)
	sv.AddStructField("HostTransportNodeId", hostTransportNodeIdParam)
	sv.AddStructField("TransportNodeVtepAction", transportNodeVtepActionParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.TransportNodeVtepActionReference
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := actionsCreateRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	aIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := aIface.connector.GetApiProvider().Invoke("com.vmware.nsx_policy.infra.sites.enforcement_points.host_transport_nodes.vteps.actions", "create", inputDataValue, executionContext)
	var emptyOutput model.TransportNodeVtepActionReference
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), actionsCreateOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.TransportNodeVtepActionReference), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), aIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}
