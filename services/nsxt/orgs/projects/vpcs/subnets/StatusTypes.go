// Copyright © 2019-2021 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Data type definitions file for service: Status.
// Includes binding types of a structures and enumerations defined in the service.
// Shared by client-side stubs and server-side skeletons to ensure type
// compatibility.

package subnets

import (
	"github.com/vmware/vsphere-automation-sdk-go/runtime/bindings"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/data"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/protocol"
	"github.com/vmware/vsphere-automation-sdk-go/services/nsxt/model"
	"reflect"
)

func statusListInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["org_id"] = bindings.NewStringType()
	fields["projectid"] = bindings.NewStringType()
	fields["vpc_id"] = bindings.NewStringType()
	fields["subnet_id"] = bindings.NewStringType()
	fieldNameMap["org_id"] = "OrgId"
	fieldNameMap["projectid"] = "Projectid"
	fieldNameMap["vpc_id"] = "VpcId"
	fieldNameMap["subnet_id"] = "SubnetId"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func statusListOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.VpcSubnetStatusListResultBindingType)
}

func statusListRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["org_id"] = bindings.NewStringType()
	fields["projectid"] = bindings.NewStringType()
	fields["vpc_id"] = bindings.NewStringType()
	fields["subnet_id"] = bindings.NewStringType()
	fieldNameMap["org_id"] = "OrgId"
	fieldNameMap["projectid"] = "Projectid"
	fieldNameMap["vpc_id"] = "VpcId"
	fieldNameMap["subnet_id"] = "SubnetId"
	paramsTypeMap["vpc_id"] = bindings.NewStringType()
	paramsTypeMap["org_id"] = bindings.NewStringType()
	paramsTypeMap["subnet_id"] = bindings.NewStringType()
	paramsTypeMap["projectid"] = bindings.NewStringType()
	paramsTypeMap["orgId"] = bindings.NewStringType()
	paramsTypeMap["projectid"] = bindings.NewStringType()
	paramsTypeMap["vpcId"] = bindings.NewStringType()
	paramsTypeMap["subnetId"] = bindings.NewStringType()
	pathParams["subnet_id"] = "subnetId"
	pathParams["vpc_id"] = "vpcId"
	pathParams["projectid"] = "projectid"
	pathParams["org_id"] = "orgId"
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
		"/policy/api/v1/orgs/{orgId}/projects/{projectid}/vpcs/{vpcId}/subnets/{subnetId}/status",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}
