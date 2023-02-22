// Copyright © 2019-2021 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Interface file for service: EffectiveProfiles
// Used by client-side stubs.

package ports

import (
	"github.com/vmware/vsphere-automation-sdk-go/lib/vapi/std/errors"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/bindings"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/core"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/lib"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/protocol/client"
	"github.com/vmware/vsphere-automation-sdk-go/services/nsxt/model"
)

const _ = core.SupportedByRuntimeVersion1

type EffectiveProfilesClient interface {

	// List all effective profiles for this segment port on given enforcement point.
	//
	// @param segmentIdParam (required)
	// @param portIdParam (required)
	// @param enforcementPointPathParam Enforcement point path (optional)
	// @return com.vmware.nsx_policy.model.EffectiveProfilesResponse
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Get(segmentIdParam string, portIdParam string, enforcementPointPathParam *string) (model.EffectiveProfilesResponse, error)
}

type effectiveProfilesClient struct {
	connector           client.Connector
	interfaceDefinition core.InterfaceDefinition
	errorsBindingMap    map[string]bindings.BindingType
}

func NewEffectiveProfilesClient(connector client.Connector) *effectiveProfilesClient {
	interfaceIdentifier := core.NewInterfaceIdentifier("com.vmware.nsx_policy.infra.segments.ports.effective_profiles")
	methodIdentifiers := map[string]core.MethodIdentifier{
		"get": core.NewMethodIdentifier(interfaceIdentifier, "get"),
	}
	interfaceDefinition := core.NewInterfaceDefinition(interfaceIdentifier, methodIdentifiers)
	errorsBindingMap := make(map[string]bindings.BindingType)

	eIface := effectiveProfilesClient{interfaceDefinition: interfaceDefinition, errorsBindingMap: errorsBindingMap, connector: connector}
	return &eIface
}

func (eIface *effectiveProfilesClient) GetErrorBindingType(errorName string) bindings.BindingType {
	if entry, ok := eIface.errorsBindingMap[errorName]; ok {
		return entry
	}
	return errors.ERROR_BINDINGS_MAP[errorName]
}

func (eIface *effectiveProfilesClient) Get(segmentIdParam string, portIdParam string, enforcementPointPathParam *string) (model.EffectiveProfilesResponse, error) {
	typeConverter := eIface.connector.TypeConverter()
	executionContext := eIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(effectiveProfilesGetInputType(), typeConverter)
	sv.AddStructField("SegmentId", segmentIdParam)
	sv.AddStructField("PortId", portIdParam)
	sv.AddStructField("EnforcementPointPath", enforcementPointPathParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.EffectiveProfilesResponse
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := effectiveProfilesGetRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	eIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := eIface.connector.GetApiProvider().Invoke("com.vmware.nsx_policy.infra.segments.ports.effective_profiles", "get", inputDataValue, executionContext)
	var emptyOutput model.EffectiveProfilesResponse
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), effectiveProfilesGetOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.EffectiveProfilesResponse), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), eIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}
