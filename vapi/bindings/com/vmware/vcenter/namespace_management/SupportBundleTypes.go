/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Data type definitions file for service: SupportBundle.
 * Includes binding types of a structures and enumerations defined in the service.
 * Shared by client-side stubs and server-side skeletons to ensure type
 * compatibility.
 */
package namespace_management

import (
	"reflect"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/bindings"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/data"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/protocol"
)

// The ``Token`` class contains information about the token required in the HTTP GET request to generate the support bundle. **Warning:** This class is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
type SupportBundleToken struct {
    // A one-time, short-lived token required in the HTTP header of the request to the url. This token needs to be passed in as a header with the name "wcp-support-bundle-token". **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
	Token string
    // Expiry time of the token. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
	Expiry string
}

// The ``Location`` class contains the URI location to download the per-cluster support bundle from, as well as a token required (as a header on the HTTP request) to get the bundle. The validity of the token is 5 minutes. After the token expires, any attempt to call the URI with said token will fail. **Warning:** This class is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
type SupportBundleLocation struct {
    // Support Bundle Download URL. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
	Url string
    // Information about the token required in the HTTP GET request to generate the support bundle. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
	WcpSupportBundleToken SupportBundleToken
}


func supportBundleCreateInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["cluster"] = bindings.NewIdType([]string{"ClusterComputeResource"}, "")
	fieldNameMap["cluster"] = "Cluster"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func supportBundleCreateOutputType() bindings.BindingType {
	return bindings.NewReferenceType(SupportBundleLocationBindingType)
}

func supportBundleCreateRestMetadata() protocol.OperationRestMetadata {
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	paramsTypeMap["cluster"] = bindings.NewIdType([]string{"ClusterComputeResource"}, "")
	paramsTypeMap["cluster"] = bindings.NewIdType([]string{"ClusterComputeResource"}, "")
	pathParams["cluster"] = "cluster"
	resultHeaders := map[string]string{}
	errorHeaders := map[string]string{}
	errorHeaders["Unauthenticated.challenge"] = "WWW-Authenticate"
	return protocol.NewOperationRestMetadata(
		paramsTypeMap,
		pathParams,
		queryParams,
		headerParams,
		"",
		"POST",
		"/vcenter/namespace-management/clusters/{cluster}/support-bundle",
		resultHeaders,
		201,
		errorHeaders,
		map[string]int{"NotFound": 404,"Error": 500,"Unauthenticated": 401,"Unauthorized": 403})
}


func SupportBundleTokenBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["token"] = bindings.NewStringType()
	fieldNameMap["token"] = "Token"
	fields["expiry"] = bindings.NewStringType()
	fieldNameMap["expiry"] = "Expiry"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vcenter.namespace_management.support_bundle.token", fields, reflect.TypeOf(SupportBundleToken{}), fieldNameMap, validators)
}

func SupportBundleLocationBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["url"] = bindings.NewStringType()
	fieldNameMap["url"] = "Url"
	fields["wcp_support_bundle_token"] = bindings.NewReferenceType(SupportBundleTokenBindingType)
	fieldNameMap["wcp_support_bundle_token"] = "WcpSupportBundleToken"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vcenter.namespace_management.support_bundle.location", fields, reflect.TypeOf(SupportBundleLocation{}), fieldNameMap, validators)
}

