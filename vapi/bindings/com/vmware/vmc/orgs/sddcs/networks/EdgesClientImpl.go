
/* Copyright © 2019 VMware, Inc. All Rights Reserved.
     SPDX-License-Identifier: BSD-2-Clause */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Client stubs for service: Edges
 * Functions that implement the generated EdgesClient interface
 */


package networks

import (
    "reflect"
// //     "getClientImplDependenciesOfOps gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/bindings/com/vmware/vapi/std/errors"
// //     "getClientImplDependenciesOfOps gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/bindings/com/vmware/vmc/model"
// //     "getClientImplDependenciesOfOps gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/bindings"
// //     "getClientImplDependenciesOfOps gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/core"
// //     "getClientImplDependenciesOfOps gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/data"
// //     "getClientImplDependenciesOfOps gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/lib"
// //     "getClientImplDependenciesOfOps gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/log"
// //     "getClientImplDependenciesOfOps gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/protocol/client"
// //
// //     "getDependenciesOfServiceTypes gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/bindings/com/vmware/vmc/model"
// //     "getDependenciesOfServiceTypes gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/bindings"
// //     "getDependenciesOfServiceTypes gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/data"
// //     "getDependenciesOfServiceTypes gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/protocol"
// 
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/bindings/com/vmware/vapi/std/errors"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/bindings/com/vmware/vmc/model"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/bindings"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/core"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/data"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/lib"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/log"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/protocol"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/protocol/client"
)


type EdgesClientImpl struct{
      interfaceName string
      interfaceDefinition core.InterfaceDefinition
      methodIdentifiers[]core.MethodIdentifier
      methodNameToDefMap map[string]*core.MethodDefinition
      errorBindingMap map[string]bindings.BindingType
      interfaceIdentifier core.InterfaceIdentifier
      connector client.Connector
}


func NewEdgesClientImpl(connector client.Connector) *EdgesClientImpl {
      interfaceName := "com.vmware.vmc.orgs.sddcs.networks.edges"
      interfaceIdentifier := core.NewInterfaceIdentifier(interfaceName)
      methodIdentifiers := []core.MethodIdentifier{
          core.NewMethodIdentifier(interfaceIdentifier, "get"),
      }
      interfaceDefinition := core.NewInterfaceDefinition(interfaceIdentifier, methodIdentifiers)
      errorBindingMap := make(map[string]bindings.BindingType)
      errorBindingMap[errors.InternalServerError{}.Error()] = errors.InternalServerErrorBindingType()
	  errorBindingMap[errors.InvalidArgument{}.Error()] = errors.InvalidArgumentBindingType()
	  errorBindingMap[errors.OperationNotFound{}.Error()] = errors.OperationNotFoundBindingType()
	  errorBindingMap[errors.UnexpectedInput{}.Error()] = errors.UnexpectedInputBindingType()
	  errorBindingMap[errors.ServiceUnavailable{}.Error()] = errors.ServiceUnavailableBindingType()
	  errorBindingMap[errors.TimedOut{}.Error()] = errors.TimedOutBindingType()
      eIface := EdgesClientImpl{interfaceName: interfaceName, methodIdentifiers: methodIdentifiers, interfaceDefinition: interfaceDefinition, errorBindingMap: errorBindingMap, interfaceIdentifier:interfaceIdentifier, connector: connector}
      eIface.methodNameToDefMap = make(map[string]*core.MethodDefinition)
      eIface.methodNameToDefMap["get"] = eIface.getMethodDefinition()
      return &eIface
}

func (eIface *EdgesClientImpl) Get(orgParam string, sddcParam string, edgeTypeParam string, prevEdgeIdParam *string, startIndexParam *int64, pageSizeParam *int64, sortOrderAscendingParam *bool, sortByParam *string, filterParam *string, ldRnameParam *string) (model.PagedEdgeList, error) {
	typeConverter := eIface.connector.TypeConverter()
	methodIdentifier := core.NewMethodIdentifier(eIface.interfaceIdentifier, "get")
	sv := bindings.NewStructValueBuilder(edgesGetInputType(), typeConverter)
	sv.AddStructField("Org", orgParam)
	sv.AddStructField("Sddc", sddcParam)
	sv.AddStructField("EdgeType", edgeTypeParam)
	sv.AddStructField("PrevEdgeId", prevEdgeIdParam)
	sv.AddStructField("StartIndex", startIndexParam)
	sv.AddStructField("PageSize", pageSizeParam)
	sv.AddStructField("SortOrderAscending", sortOrderAscendingParam)
	sv.AddStructField("SortBy", sortByParam)
	sv.AddStructField("Filter", filterParam)
	sv.AddStructField("LdRname", ldRnameParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
        var emptyOutput model.PagedEdgeList
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := edgesGetRestMetadata
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	eIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult:= eIface.Invoke(eIface.connector.NewExecutionContext(), methodIdentifier, inputDataValue)
	var emptyOutput model.PagedEdgeList
    if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), edgesGetOutputType())
		if errorInOutput != nil {
		    return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
	    return output.(model.PagedEdgeList), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), eIface.errorBindingMap[methodResult.Error().Name()])
		if errorInError != nil {
		    return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (eIface *EdgesClientImpl) Invoke(ctx *core.ExecutionContext, methodId core.MethodIdentifier, inputDataValue data.DataValue) core.MethodResult {
    methodResult := eIface.connector.GetApiProvider().Invoke(eIface.interfaceName, methodId.Name(), inputDataValue, ctx)
    return methodResult
}

func (eIface *EdgesClientImpl) getMethodDefinition() *core.MethodDefinition {
      interfaceIdentifier := core.NewInterfaceIdentifier(eIface.interfaceName)
      typeConverter := eIface.connector.TypeConverter()

      input, inputError := typeConverter.ConvertToDataDefinition(edgesGetInputType())
      output, outputError := typeConverter.ConvertToDataDefinition(edgesGetOutputType())
      if(inputError != nil) {
          log.Errorf("Error in ConvertToDataDefinition for EdgesClientImpl.get method's input - %s",
              bindings.VAPIerrorsToError(inputError).Error())
          return nil
      }
      if(outputError != nil) {
          log.Errorf("Error in ConvertToDataDefinition for EdgesClientImpl.get method's output - %s",
              bindings.VAPIerrorsToError(outputError).Error())
          return nil
      }
      methodIdentifier := core.NewMethodIdentifier(interfaceIdentifier, "get")
      errorDefinitions := make([]data.ErrorDefinition,0)
      eIface.errorBindingMap[errors.InvalidRequest{}.Error()] = errors.InvalidRequestBindingType()
	  errDef1, errError1 := typeConverter.ConvertToDataDefinition(errors.InvalidRequestBindingType())
	  if(errError1 != nil) {
	    log.Errorf("Error in ConvertToDataDefinition for EdgesClientImpl.get method's errors.InvalidRequest error - %s",
	        bindings.VAPIerrorsToError(errError1).Error())
        return nil
	  }
	  errorDefinitions = append(errorDefinitions, errDef1.(data.ErrorDefinition))
      eIface.errorBindingMap[errors.Unauthorized{}.Error()] = errors.UnauthorizedBindingType()
	  errDef2, errError2 := typeConverter.ConvertToDataDefinition(errors.UnauthorizedBindingType())
	  if(errError2 != nil) {
	    log.Errorf("Error in ConvertToDataDefinition for EdgesClientImpl.get method's errors.Unauthorized error - %s",
	        bindings.VAPIerrorsToError(errError2).Error())
        return nil
	  }
	  errorDefinitions = append(errorDefinitions, errDef2.(data.ErrorDefinition))
      eIface.errorBindingMap[errors.NotFound{}.Error()] = errors.NotFoundBindingType()
	  errDef3, errError3 := typeConverter.ConvertToDataDefinition(errors.NotFoundBindingType())
	  if(errError3 != nil) {
	    log.Errorf("Error in ConvertToDataDefinition for EdgesClientImpl.get method's errors.NotFound error - %s",
	        bindings.VAPIerrorsToError(errError3).Error())
        return nil
	  }
	  errorDefinitions = append(errorDefinitions, errDef3.(data.ErrorDefinition))

      methodDefinition := core.NewMethodDefinition(methodIdentifier, input, output, errorDefinitions)
      return &methodDefinition
}







func edgesGetInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["org"] = bindings.NewStringType()
    fields["sddc"] = bindings.NewStringType()
    fields["edge_type"] = bindings.NewStringType()
    fields["prev_edge_id"] = bindings.NewOptionalType(bindings.NewStringType())
    fields["start_index"] = bindings.NewOptionalType(bindings.NewIntegerType())
    fields["page_size"] = bindings.NewOptionalType(bindings.NewIntegerType())
    fields["sort_order_ascending"] = bindings.NewOptionalType(bindings.NewBooleanType())
    fields["sort_by"] = bindings.NewOptionalType(bindings.NewStringType())
    fields["filter"] = bindings.NewOptionalType(bindings.NewStringType())
    fields["ld_rname"] = bindings.NewOptionalType(bindings.NewStringType())
    fieldNameMap["org"] = "Org"
    fieldNameMap["sddc"] = "Sddc"
    fieldNameMap["edge_type"] = "EdgeType"
    fieldNameMap["prev_edge_id"] = "PrevEdgeId"
    fieldNameMap["start_index"] = "StartIndex"
    fieldNameMap["page_size"] = "PageSize"
    fieldNameMap["sort_order_ascending"] = "SortOrderAscending"
    fieldNameMap["sort_by"] = "SortBy"
    fieldNameMap["filter"] = "Filter"
    fieldNameMap["ld_rname"] = "LdRname"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func edgesGetOutputType() bindings.BindingType {
    return bindings.NewReferenceType(model.PagedEdgeListBindingType)
}

func edgesGetRestMetadata() protocol.OperationRestMetadata {
    paramsTypeMap := map[string]bindings.BindingType{}
    pathParams := map[string]string{}
    queryParams := map[string]string{}
    headerParams := map[string]string{}
    paramsTypeMap["edge_type"] = bindings.NewStringType()
    paramsTypeMap["org"] = bindings.NewStringType()
    paramsTypeMap["sddc"] = bindings.NewStringType()
    paramsTypeMap["filter"] = bindings.NewOptionalType(bindings.NewStringType())
    paramsTypeMap["sort_order_ascending"] = bindings.NewOptionalType(bindings.NewBooleanType())
    paramsTypeMap["page_size"] = bindings.NewOptionalType(bindings.NewIntegerType())
    paramsTypeMap["start_index"] = bindings.NewOptionalType(bindings.NewIntegerType())
    paramsTypeMap["sort_by"] = bindings.NewOptionalType(bindings.NewStringType())
    paramsTypeMap["prev_edge_id"] = bindings.NewOptionalType(bindings.NewStringType())
    paramsTypeMap["ld_rname"] = bindings.NewOptionalType(bindings.NewStringType())
    paramsTypeMap["org"] = bindings.NewStringType()
    paramsTypeMap["sddc"] = bindings.NewStringType()
    pathParams["org"] = "org"
    pathParams["sddc"] = "sddc"
    queryParams["filter"] = "filter"
    queryParams["start_index"] = "startIndex"
    queryParams["edge_type"] = "edgeType"
    queryParams["page_size"] = "pageSize"
    queryParams["sort_by"] = "sortBy"
    queryParams["prev_edge_id"] = "prevEdgeId"
    queryParams["sort_order_ascending"] = "sortOrderAscending"
    queryParams["ld_rname"] = "LDRname"
    resultHeaders := map[string]string{}
    errorHeaders := map[string]string{}
    return protocol.NewOperationRestMetadata(
      paramsTypeMap,
      pathParams,
      queryParams,
      headerParams,
      "",
      "GET",
      "/vmc/api/orgs/{org}/sddcs/{sddc}/networks/4.0/edges",
       resultHeaders,
       200,
       errorHeaders,
       map[string]int{"InvalidRequest": 400,"Unauthorized": 403,"NotFound": 404})
}




