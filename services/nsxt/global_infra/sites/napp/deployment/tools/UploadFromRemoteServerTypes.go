// Copyright © 2019-2021 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Data type definitions file for service: Upload_from_remote_server.
// Includes binding types of a structures and enumerations defined in the service.
// Shared by client-side stubs and server-side skeletons to ensure type
// compatibility.

package tools

import (
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/runtime/bindings"
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/runtime/data"
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/runtime/protocol"
	"gitlab.eng.vmware.com/vapi-sdk/vsphere-automation-sdk-go/services/nsxt/model"
	"reflect"
)

func uploadFromRemoteServerCreateInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["site_id"] = bindings.NewStringType()
	fields["upload_bundle_remote_server"] = bindings.NewReferenceType(model.UploadBundleRemoteServerBindingType)
	fieldNameMap["site_id"] = "SiteId"
	fieldNameMap["upload_bundle_remote_server"] = "UploadBundleRemoteServer"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func uploadFromRemoteServerCreateOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.UploadBundleIdBindingType)
}

func uploadFromRemoteServerCreateRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["site_id"] = bindings.NewStringType()
	fields["upload_bundle_remote_server"] = bindings.NewReferenceType(model.UploadBundleRemoteServerBindingType)
	fieldNameMap["site_id"] = "SiteId"
	fieldNameMap["upload_bundle_remote_server"] = "UploadBundleRemoteServer"
	paramsTypeMap["upload_bundle_remote_server"] = bindings.NewReferenceType(model.UploadBundleRemoteServerBindingType)
	paramsTypeMap["site_id"] = bindings.NewStringType()
	paramsTypeMap["siteId"] = bindings.NewStringType()
	pathParams["site_id"] = "siteId"
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
		"upload_bundle_remote_server",
		"POST",
		"/policy/api/v1/global-infra/sites/{siteId}/napp/deployment/tools/upload_from_remote_server",
		"",
		resultHeaders,
		201,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}
