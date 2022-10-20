// Copyright © 2019-2021 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Data type definitions file for service: Get_vm_group_execution_details.
// Includes binding types of a structures and enumerations defined in the service.
// Shared by client-side stubs and server-side skeletons to ensure type
// compatibility.

package actions

import (
	"github.com/TaoZou1/vsphere-automation-sdk-go/runtime/bindings"
	"github.com/TaoZou1/vsphere-automation-sdk-go/runtime/data"
	"github.com/TaoZou1/vsphere-automation-sdk-go/runtime/protocol"
	"github.com/TaoZou1/vsphere-automation-sdk-go/services/nsxt-mp/nsx/model"
	"reflect"
)

func getVmGroupExecutionDetailsListInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["group_id"] = bindings.NewStringType()
	fields["federation_site_id"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["group_id"] = "GroupId"
	fieldNameMap["federation_site_id"] = "FederationSiteId"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func getVmGroupExecutionDetailsListOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.VmGroupExecutionDetailsBindingType)
}

func getVmGroupExecutionDetailsListRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["group_id"] = bindings.NewStringType()
	fields["federation_site_id"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["group_id"] = "GroupId"
	fieldNameMap["federation_site_id"] = "FederationSiteId"
	paramsTypeMap["federation_site_id"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["group_id"] = bindings.NewStringType()
	queryParams["group_id"] = "group_id"
	queryParams["federation_site_id"] = "federation_site_id"
	resultHeaders := map[string]string{}
	errorHeaders := map[string]map[string]string{}
	return protocol.NewOperationRestMetadata(
		fields,
		fieldNameMap,
		paramsTypeMap,
		pathParams,
		queryParams,
		headerParams,
		dispatchHeaderParams,
		bodyFieldsMap,
		"",
		"",
		"GET",
		"/api/v1/migration/vmgroup/actions/get_vm_group_execution_details",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}
