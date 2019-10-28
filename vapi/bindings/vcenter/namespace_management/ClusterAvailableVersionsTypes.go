/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Data type definitions file for service: ClusterAvailableVersions.
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
	"time"
)



// The ``Summary`` class contains the information about each available upgrade version. **Warning:** This class is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
type ClusterAvailableVersionsSummary struct {
    // Version of the upgrade. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
	Version string
    // Name of the upgrade. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
	Name string
    // Description of the upgrade. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
	Description string
    // Release date of the upgrade. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
	ReleaseDate time.Time
    // Release details of the upgrade. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
	ReleaseNotes string
}



func clusterAvailableVersionsListInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func clusterAvailableVersionsListOutputType() bindings.BindingType {
	return bindings.NewListType(bindings.NewReferenceType(ClusterAvailableVersionsSummaryBindingType), reflect.TypeOf([]ClusterAvailableVersionsSummary{}))
}

func clusterAvailableVersionsListRestMetadata() protocol.OperationRestMetadata {
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
		"/vcenter/namespace-management/software/cluster-available-versions",
		resultHeaders,
		200,
		errorHeaders,
		map[string]int{"Error": 500,"Unauthenticated": 401,"Unauthorized": 403})
}


func ClusterAvailableVersionsSummaryBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["version"] = bindings.NewStringType()
	fieldNameMap["version"] = "Version"
	fields["name"] = bindings.NewStringType()
	fieldNameMap["name"] = "Name"
	fields["description"] = bindings.NewStringType()
	fieldNameMap["description"] = "Description"
	fields["release_date"] = bindings.NewDateTimeType()
	fieldNameMap["release_date"] = "ReleaseDate"
	fields["release_notes"] = bindings.NewStringType()
	fieldNameMap["release_notes"] = "ReleaseNotes"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vcenter.namespace_management.cluster_available_versions.summary", fields, reflect.TypeOf(ClusterAvailableVersionsSummary{}), fieldNameMap, validators)
}

