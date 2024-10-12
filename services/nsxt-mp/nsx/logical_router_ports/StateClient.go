// Copyright © 2019-2021 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Interface file for service: State
// Used by client-side stubs.

package logical_router_ports

import (
	"github.com/vmware/vsphere-automation-sdk-go/lib/vapi/std/errors"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/bindings"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/core"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/lib"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/protocol/client"
	"github.com/vmware/vsphere-automation-sdk-go/services/nsxt-mp/nsx/model"
)

const _ = core.SupportedByRuntimeVersion1

type StateClient interface {

	// Return realized state information of a logical router port. Any configuration update that affects the logical router port can use this API to get its realized state by passing a request_id returned by the configuration change operation. e.g. Update configuration of logical router ports, dhcp relays, etc.
	//
	//  To get all realized entities for the intent use below Policy API.
	//  GET /policy/api/v1/infra/realized-state/realized-entities?intent_path=<intent-path>
	//  For realized status of the intent use below Policy API.
	//  GET /policy/api/v1/infra/realized-state/status?intent_path=<intent-path>
	//  <intent-path> can be one of the following.
	//  /infra/tier-0s/<tier-0-id>/locale-services/<locale-service-id>/interfaces/<interface-id>
	//  /infra/tier-1s/<tier-1-id>/locale-services/<locale-service-id>/interfaces/<interface-id>
	//  /infra/tier-1s/<tier-1-id>/segments/<segment-id> for DOWNLINK
	//  /infra/segments/<segment-id> for DOWNLINK
	//  There are specific Policy APIs to get Segment state.
	//  GET /policy/api/v1/infra/segments/<segment-id>/state
	//  GET /policy/api/v1/infra/tier-1s/<tier-1-id>/segments/<segment-id>/state
	//  For DAD status use below Policy APIs.
	//  GET /policy/api/v1/infra/segments/<segment-id>/gateway-interface-dad-state
	//  GET /policy/api/v1/infra/tier-1s/<tier-1-id>/segments/<segment-id>/gateway-interface-dad-state
	//  For DAD status of all interfaces created on Gateway use below Policy APIs.
	//  GET /policy/api/v1/infra/tier-0s/<tier-0-id>/state
	//  GET /policy/api/v1/infra/tier-1s/<tier-1-id>/state
	//
	// @param logicalRouterPortIdParam (required)
	// @param barrierIdParam (optional)
	// @param requestIdParam Realization request ID (optional)
	// @return com.vmware.nsx.model.LogicalRouterPortState
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Get(logicalRouterPortIdParam string, barrierIdParam *int64, requestIdParam *string) (model.LogicalRouterPortState, error)
}

type stateClient struct {
	connector           client.Connector
	interfaceDefinition core.InterfaceDefinition
	errorsBindingMap    map[string]bindings.BindingType
}

func NewStateClient(connector client.Connector) *stateClient {
	interfaceIdentifier := core.NewInterfaceIdentifier("com.vmware.nsx.logical_router_ports.state")
	methodIdentifiers := map[string]core.MethodIdentifier{
		"get": core.NewMethodIdentifier(interfaceIdentifier, "get"),
	}
	interfaceDefinition := core.NewInterfaceDefinition(interfaceIdentifier, methodIdentifiers)
	errorsBindingMap := make(map[string]bindings.BindingType)

	sIface := stateClient{interfaceDefinition: interfaceDefinition, errorsBindingMap: errorsBindingMap, connector: connector}
	return &sIface
}

func (sIface *stateClient) GetErrorBindingType(errorName string) bindings.BindingType {
	if entry, ok := sIface.errorsBindingMap[errorName]; ok {
		return entry
	}
	return errors.ERROR_BINDINGS_MAP[errorName]
}

func (sIface *stateClient) Get(logicalRouterPortIdParam string, barrierIdParam *int64, requestIdParam *string) (model.LogicalRouterPortState, error) {
	typeConverter := sIface.connector.TypeConverter()
	executionContext := sIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(stateGetInputType(), typeConverter)
	sv.AddStructField("LogicalRouterPortId", logicalRouterPortIdParam)
	sv.AddStructField("BarrierId", barrierIdParam)
	sv.AddStructField("RequestId", requestIdParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.LogicalRouterPortState
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := stateGetRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	sIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := sIface.connector.GetApiProvider().Invoke("com.vmware.nsx.logical_router_ports.state", "get", inputDataValue, executionContext)
	var emptyOutput model.LogicalRouterPortState
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), stateGetOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.LogicalRouterPortState), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), sIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}
