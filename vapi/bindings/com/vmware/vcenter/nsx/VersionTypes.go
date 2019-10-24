/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Data type definitions file for service: Version.
 * Includes binding types of a structures and enumerations defined in the service.
 * Shared by client-side stubs and server-side skeletons to ensure type
 * compatibility.
 */
package nsx

import (
	"reflect"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/bindings"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/data"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/protocol"
)

// The ``VersionInfo`` class contains version information about supported NSX version for vCenter Server. **Warning:** This class is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
type VersionVersionInfo struct {
    // Supported NSX version on vCenter Server. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
	SupportedVersion string
}


func versionGetVersionInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func versionGetVersionOutputType() bindings.BindingType {
	return bindings.NewReferenceType(VersionVersionInfoBindingType)
}

func versionGetVersionRestMetadata() protocol.OperationRestMetadata {
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	resultHeaders := map[string]string{}
	errorHeaders := map[string]string{}
	errorHeaders["Unauthenticated.challenge"] = "WWW-Authenticate"
	return protocol.NewOperationRestMetadata(
		paramsTypeMap,
		pathParams,
		queryParams,
		headerParams,
		"",
		"GET",
		"/vcenter/nsx/supported-version",
		resultHeaders,
		200,
		errorHeaders,
		map[string]int{"Unauthenticated": 401,"Unauthorized": 403,"Error": 500})
}


func VersionVersionInfoBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["supported_version"] = bindings.NewStringType()
	fieldNameMap["supported_version"] = "SupportedVersion"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vcenter.nsx.version.version_info", fields, reflect.TypeOf(VersionVersionInfo{}), fieldNameMap, validators)
}

