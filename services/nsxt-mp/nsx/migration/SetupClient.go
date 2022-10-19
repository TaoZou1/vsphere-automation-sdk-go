// Copyright © 2019-2021 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Interface file for service: Setup
// Used by client-side stubs.

package migration

import (
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/lib/vapi/std/errors"
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/runtime/bindings"
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/runtime/core"
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/runtime/lib"
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/runtime/protocol/client"
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/services/nsxt-mp/nsx/model"
)

const _ = core.SupportedByRuntimeVersion1

type SetupClient interface {

	// Add ALB endpoint details for non cross VC migration modes.
	//
	// @param aviEndPointParam (required)
	// @throws Unauthenticated  Unauthorized
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Addalbinfo(aviEndPointParam model.AviEndPoint) error

	// Add NSX-V to NSX-T site mapping.
	//
	// @param v2tSiteMappingSpecParam (required)
	// @throws Unauthenticated  Unauthorized
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Addv2tsitemapping(v2tSiteMappingSpecParam model.V2tSiteMappingSpec) error

	// Get setup details of NSX-V to be migrated.
	// @return com.vmware.nsx.model.MigrationSetupInfo
	// @throws Unauthenticated  Unauthorized
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Get() (model.MigrationSetupInfo, error)

	// This API is to support the add Host workflow once the V2T migration process has started. The high level add host workflow will be as follows. 1. Verify that overall migration status is in PAUSED state in migration co-ordinator. 2. Follow steps mentioned in product documentation to add a new host to cluster. 3. Invoke the migration co-ordinator sync HOST groups API. This will result in migration co-ordinator updating its inventory and accouting for the newly added host transport node. At this point, the list host upgrade units API call to MC will also show the newly added host transport node. But do NOT resume the migration. Since this host transport node is added after the V2T migration has started, we need to perform migration steps on this host through a special API (Accept Host Transport Node) 4. Invoke this API (Accept Host Transport Node) to perform the V2T migration steps on newly added host transport node. Once the API is successful, we should be seeing that migration status has been marked as SUCCESS for the newly added host transport node. Very Importnant Note : Make sure that migration status is PAUSED in migration co-ordinator when performing steps 2 through 4. Also ensure that all the steps 2 through 4 pass without failures. This API should be invoked only when - The overall migration status is in PAUSED status. - The migration mode selected has HOST component in it. - All the migration stages before HOST stage are completed. - The new host has been added to a cluster by following all the steps mentioned in product documentation.
	//
	// @param newHostTransportNodeSpecParam (required)
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Migratenewlyaddedhosttransportnode(newHostTransportNodeSpecParam model.NewHostTransportNodeSpec) error

	// Set the NSX-V ESG to NSX-T Router mapping option.
	//
	// @param mappingOptionParam Mapping option (required)
	// @throws Unauthenticated  Unauthorized
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Setesgtoroutermappingoption(mappingOptionParam string) error

	// Provide setup details of NSX-V to be migrated.
	//
	// @param migrationSetupInfoParam (required)
	// @return com.vmware.nsx.model.MigrationSetupInfo
	// @throws Unauthenticated  Unauthorized
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Update(migrationSetupInfoParam model.MigrationSetupInfo) (model.MigrationSetupInfo, error)
}

type setupClient struct {
	connector           client.Connector
	interfaceDefinition core.InterfaceDefinition
	errorsBindingMap    map[string]bindings.BindingType
}

func NewSetupClient(connector client.Connector) *setupClient {
	interfaceIdentifier := core.NewInterfaceIdentifier("com.vmware.nsx.migration.setup")
	methodIdentifiers := map[string]core.MethodIdentifier{
		"addalbinfo":                         core.NewMethodIdentifier(interfaceIdentifier, "addalbinfo"),
		"addv2tsitemapping":                  core.NewMethodIdentifier(interfaceIdentifier, "addv2tsitemapping"),
		"get":                                core.NewMethodIdentifier(interfaceIdentifier, "get"),
		"migratenewlyaddedhosttransportnode": core.NewMethodIdentifier(interfaceIdentifier, "migratenewlyaddedhosttransportnode"),
		"setesgtoroutermappingoption":        core.NewMethodIdentifier(interfaceIdentifier, "setesgtoroutermappingoption"),
		"update":                             core.NewMethodIdentifier(interfaceIdentifier, "update"),
	}
	interfaceDefinition := core.NewInterfaceDefinition(interfaceIdentifier, methodIdentifiers)
	errorsBindingMap := make(map[string]bindings.BindingType)

	sIface := setupClient{interfaceDefinition: interfaceDefinition, errorsBindingMap: errorsBindingMap, connector: connector}
	return &sIface
}

func (sIface *setupClient) GetErrorBindingType(errorName string) bindings.BindingType {
	if entry, ok := sIface.errorsBindingMap[errorName]; ok {
		return entry
	}
	return errors.ERROR_BINDINGS_MAP[errorName]
}

func (sIface *setupClient) Addalbinfo(aviEndPointParam model.AviEndPoint) error {
	typeConverter := sIface.connector.TypeConverter()
	executionContext := sIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(setupAddalbinfoInputType(), typeConverter)
	sv.AddStructField("AviEndPoint", aviEndPointParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		return bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := setupAddalbinfoRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	sIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := sIface.connector.GetApiProvider().Invoke("com.vmware.nsx.migration.setup", "addalbinfo", inputDataValue, executionContext)
	if methodResult.IsSuccess() {
		return nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), sIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return bindings.VAPIerrorsToError(errorInError)
		}
		return methodError.(error)
	}
}

func (sIface *setupClient) Addv2tsitemapping(v2tSiteMappingSpecParam model.V2tSiteMappingSpec) error {
	typeConverter := sIface.connector.TypeConverter()
	executionContext := sIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(setupAddv2tsitemappingInputType(), typeConverter)
	sv.AddStructField("V2tSiteMappingSpec", v2tSiteMappingSpecParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		return bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := setupAddv2tsitemappingRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	sIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := sIface.connector.GetApiProvider().Invoke("com.vmware.nsx.migration.setup", "addv2tsitemapping", inputDataValue, executionContext)
	if methodResult.IsSuccess() {
		return nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), sIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return bindings.VAPIerrorsToError(errorInError)
		}
		return methodError.(error)
	}
}

func (sIface *setupClient) Get() (model.MigrationSetupInfo, error) {
	typeConverter := sIface.connector.TypeConverter()
	executionContext := sIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(setupGetInputType(), typeConverter)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.MigrationSetupInfo
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := setupGetRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	sIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := sIface.connector.GetApiProvider().Invoke("com.vmware.nsx.migration.setup", "get", inputDataValue, executionContext)
	var emptyOutput model.MigrationSetupInfo
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), setupGetOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.MigrationSetupInfo), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), sIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (sIface *setupClient) Migratenewlyaddedhosttransportnode(newHostTransportNodeSpecParam model.NewHostTransportNodeSpec) error {
	typeConverter := sIface.connector.TypeConverter()
	executionContext := sIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(setupMigratenewlyaddedhosttransportnodeInputType(), typeConverter)
	sv.AddStructField("NewHostTransportNodeSpec", newHostTransportNodeSpecParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		return bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := setupMigratenewlyaddedhosttransportnodeRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	sIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := sIface.connector.GetApiProvider().Invoke("com.vmware.nsx.migration.setup", "migratenewlyaddedhosttransportnode", inputDataValue, executionContext)
	if methodResult.IsSuccess() {
		return nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), sIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return bindings.VAPIerrorsToError(errorInError)
		}
		return methodError.(error)
	}
}

func (sIface *setupClient) Setesgtoroutermappingoption(mappingOptionParam string) error {
	typeConverter := sIface.connector.TypeConverter()
	executionContext := sIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(setupSetesgtoroutermappingoptionInputType(), typeConverter)
	sv.AddStructField("MappingOption", mappingOptionParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		return bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := setupSetesgtoroutermappingoptionRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	sIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := sIface.connector.GetApiProvider().Invoke("com.vmware.nsx.migration.setup", "setesgtoroutermappingoption", inputDataValue, executionContext)
	if methodResult.IsSuccess() {
		return nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), sIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return bindings.VAPIerrorsToError(errorInError)
		}
		return methodError.(error)
	}
}

func (sIface *setupClient) Update(migrationSetupInfoParam model.MigrationSetupInfo) (model.MigrationSetupInfo, error) {
	typeConverter := sIface.connector.TypeConverter()
	executionContext := sIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(setupUpdateInputType(), typeConverter)
	sv.AddStructField("MigrationSetupInfo", migrationSetupInfoParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.MigrationSetupInfo
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := setupUpdateRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	sIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := sIface.connector.GetApiProvider().Invoke("com.vmware.nsx.migration.setup", "update", inputDataValue, executionContext)
	var emptyOutput model.MigrationSetupInfo
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), setupUpdateOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.MigrationSetupInfo), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), sIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}
