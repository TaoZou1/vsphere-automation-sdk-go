// Copyright © 2019-2021 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Interface file for service: IpSubnets
// Used by client-side stubs.

package pools

import (
	"github.com/TaoZou1/vsphere-automation-sdk-go/lib/vapi/std/errors"
	"github.com/TaoZou1/vsphere-automation-sdk-go/runtime/bindings"
	"github.com/TaoZou1/vsphere-automation-sdk-go/runtime/core"
	"github.com/TaoZou1/vsphere-automation-sdk-go/runtime/lib"
	"github.com/TaoZou1/vsphere-automation-sdk-go/runtime/protocol/client"
	"github.com/TaoZou1/vsphere-automation-sdk-go/services/nsxt-mp/nsx/model"
)

const _ = core.SupportedByRuntimeVersion1

type IpSubnetsClient interface {

	//
	//
	// @param subnetIdParam IP subnet id (required)
	// @param allocationIpAddressParam (required)
	// @param actionParam Specifies allocate or release action (required)
	// @return com.vmware.nsx.model.AllocationIpAddress
	// @throws ConcurrentChange  Conflict
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Allocateorreleasefromipblocksubnet(subnetIdParam string, allocationIpAddressParam model.AllocationIpAddress, actionParam string) (model.AllocationIpAddress, error)

	//
	//
	// @param ipBlockSubnetParam (required)
	// @return com.vmware.nsx.model.IpBlockSubnet
	// @throws ConcurrentChange  Conflict
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Create(ipBlockSubnetParam model.IpBlockSubnet) (model.IpBlockSubnet, error)

	//
	//
	// @param subnetIdParam Subnet id (required)
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Delete(subnetIdParam string) error

	//
	//
	// @param subnetIdParam Subnet id (required)
	// @return com.vmware.nsx.model.IpBlockSubnet
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	Get(subnetIdParam string) (model.IpBlockSubnet, error)

	// Returns information about all subnets present within an IP address block. Information includes subnet's id, display_name, description, cidr and allocation ranges.
	//
	// @param blockIdParam (optional)
	// @param cursorParam Opaque cursor to be used for getting next page of records (supplied by current result page) (optional)
	// @param includedFieldsParam Comma separated list of fields that should be included in query result (optional)
	// @param pageSizeParam Maximum number of results to return in this page (server may return fewer) (optional, default to 1000)
	// @param sortAscendingParam (optional)
	// @param sortByParam Field by which records are sorted (optional)
	// @return com.vmware.nsx.model.IpBlockSubnetListResult
	// @throws InvalidRequest  Bad Request, Precondition Failed
	// @throws Unauthorized  Forbidden
	// @throws ServiceUnavailable  Service Unavailable
	// @throws InternalServerError  Internal Server Error
	// @throws NotFound  Not Found
	List(blockIdParam *string, cursorParam *string, includedFieldsParam *string, pageSizeParam *int64, sortAscendingParam *bool, sortByParam *string) (model.IpBlockSubnetListResult, error)
}

type ipSubnetsClient struct {
	connector           client.Connector
	interfaceDefinition core.InterfaceDefinition
	errorsBindingMap    map[string]bindings.BindingType
}

func NewIpSubnetsClient(connector client.Connector) *ipSubnetsClient {
	interfaceIdentifier := core.NewInterfaceIdentifier("com.vmware.nsx.pools.ip_subnets")
	methodIdentifiers := map[string]core.MethodIdentifier{
		"allocateorreleasefromipblocksubnet": core.NewMethodIdentifier(interfaceIdentifier, "allocateorreleasefromipblocksubnet"),
		"create":                             core.NewMethodIdentifier(interfaceIdentifier, "create"),
		"delete":                             core.NewMethodIdentifier(interfaceIdentifier, "delete"),
		"get":                                core.NewMethodIdentifier(interfaceIdentifier, "get"),
		"list":                               core.NewMethodIdentifier(interfaceIdentifier, "list"),
	}
	interfaceDefinition := core.NewInterfaceDefinition(interfaceIdentifier, methodIdentifiers)
	errorsBindingMap := make(map[string]bindings.BindingType)

	iIface := ipSubnetsClient{interfaceDefinition: interfaceDefinition, errorsBindingMap: errorsBindingMap, connector: connector}
	return &iIface
}

func (iIface *ipSubnetsClient) GetErrorBindingType(errorName string) bindings.BindingType {
	if entry, ok := iIface.errorsBindingMap[errorName]; ok {
		return entry
	}
	return errors.ERROR_BINDINGS_MAP[errorName]
}

func (iIface *ipSubnetsClient) Allocateorreleasefromipblocksubnet(subnetIdParam string, allocationIpAddressParam model.AllocationIpAddress, actionParam string) (model.AllocationIpAddress, error) {
	typeConverter := iIface.connector.TypeConverter()
	executionContext := iIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(ipSubnetsAllocateorreleasefromipblocksubnetInputType(), typeConverter)
	sv.AddStructField("SubnetId", subnetIdParam)
	sv.AddStructField("AllocationIpAddress", allocationIpAddressParam)
	sv.AddStructField("Action", actionParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.AllocationIpAddress
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := ipSubnetsAllocateorreleasefromipblocksubnetRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	iIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := iIface.connector.GetApiProvider().Invoke("com.vmware.nsx.pools.ip_subnets", "allocateorreleasefromipblocksubnet", inputDataValue, executionContext)
	var emptyOutput model.AllocationIpAddress
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), ipSubnetsAllocateorreleasefromipblocksubnetOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.AllocationIpAddress), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), iIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (iIface *ipSubnetsClient) Create(ipBlockSubnetParam model.IpBlockSubnet) (model.IpBlockSubnet, error) {
	typeConverter := iIface.connector.TypeConverter()
	executionContext := iIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(ipSubnetsCreateInputType(), typeConverter)
	sv.AddStructField("IpBlockSubnet", ipBlockSubnetParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.IpBlockSubnet
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := ipSubnetsCreateRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	iIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := iIface.connector.GetApiProvider().Invoke("com.vmware.nsx.pools.ip_subnets", "create", inputDataValue, executionContext)
	var emptyOutput model.IpBlockSubnet
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), ipSubnetsCreateOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.IpBlockSubnet), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), iIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (iIface *ipSubnetsClient) Delete(subnetIdParam string) error {
	typeConverter := iIface.connector.TypeConverter()
	executionContext := iIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(ipSubnetsDeleteInputType(), typeConverter)
	sv.AddStructField("SubnetId", subnetIdParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		return bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := ipSubnetsDeleteRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	iIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := iIface.connector.GetApiProvider().Invoke("com.vmware.nsx.pools.ip_subnets", "delete", inputDataValue, executionContext)
	if methodResult.IsSuccess() {
		return nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), iIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return bindings.VAPIerrorsToError(errorInError)
		}
		return methodError.(error)
	}
}

func (iIface *ipSubnetsClient) Get(subnetIdParam string) (model.IpBlockSubnet, error) {
	typeConverter := iIface.connector.TypeConverter()
	executionContext := iIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(ipSubnetsGetInputType(), typeConverter)
	sv.AddStructField("SubnetId", subnetIdParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.IpBlockSubnet
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := ipSubnetsGetRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	iIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := iIface.connector.GetApiProvider().Invoke("com.vmware.nsx.pools.ip_subnets", "get", inputDataValue, executionContext)
	var emptyOutput model.IpBlockSubnet
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), ipSubnetsGetOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.IpBlockSubnet), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), iIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (iIface *ipSubnetsClient) List(blockIdParam *string, cursorParam *string, includedFieldsParam *string, pageSizeParam *int64, sortAscendingParam *bool, sortByParam *string) (model.IpBlockSubnetListResult, error) {
	typeConverter := iIface.connector.TypeConverter()
	executionContext := iIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(ipSubnetsListInputType(), typeConverter)
	sv.AddStructField("BlockId", blockIdParam)
	sv.AddStructField("Cursor", cursorParam)
	sv.AddStructField("IncludedFields", includedFieldsParam)
	sv.AddStructField("PageSize", pageSizeParam)
	sv.AddStructField("SortAscending", sortAscendingParam)
	sv.AddStructField("SortBy", sortByParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.IpBlockSubnetListResult
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := ipSubnetsListRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	iIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := iIface.connector.GetApiProvider().Invoke("com.vmware.nsx.pools.ip_subnets", "list", inputDataValue, executionContext)
	var emptyOutput model.IpBlockSubnetListResult
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), ipSubnetsListOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.IpBlockSubnetListResult), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), iIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}
