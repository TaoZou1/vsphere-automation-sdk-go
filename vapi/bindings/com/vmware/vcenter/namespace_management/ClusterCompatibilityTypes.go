/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Data type definitions file for service: ClusterCompatibility.
 * Includes binding types of a structures and enumerations defined in the service.
 * Shared by client-side stubs and server-side skeletons to ensure type
 * compatibility.
 */
package namespace_management

import (
	"reflect"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/bindings/com/vmware/vapi/std"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/bindings"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/data"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/protocol"
)



// The ``Summary`` class contains the information about the compatibility of a cluster. **Warning:** This class is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
type ClusterCompatibilitySummary struct {
    // Identifier of the cluster. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
	Cluster string
    // Compatibility of this cluster. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
	Compatible bool
    // Reason for incompatibility. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
	IncompatibilityReason *std.LocalizableMessage
}

// The ``FilterSpec`` class contains properties used to filter the results when listing clusters (see ClusterCompatibility#list) and their compatibility information. **Warning:** This class is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
type ClusterCompatibilityFilterSpec struct {
    // Compatibility criteria for matching the filter. If true, only clusters which are compatible for Namespaces match the filter. If false, all clusters match the filter. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
	Compatible *bool
}



func clusterCompatibilityListInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["filter"] = bindings.NewOptionalType(bindings.NewReferenceType(ClusterCompatibilityFilterSpecBindingType))
	fieldNameMap["filter"] = "Filter"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func clusterCompatibilityListOutputType() bindings.BindingType {
	return bindings.NewListType(bindings.NewReferenceType(ClusterCompatibilitySummaryBindingType), reflect.TypeOf([]ClusterCompatibilitySummary{}))
}

func clusterCompatibilityListRestMetadata() protocol.OperationRestMetadata {
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	paramsTypeMap["filter.compatible"] = bindings.NewOptionalType(bindings.NewBooleanType())
	queryParams["filter.compatible"] = "compatible"
	resultHeaders := map[string]string{}
	errorHeaders := map[string]string{}
	return protocol.NewOperationRestMetadata(
		paramsTypeMap,
		pathParams,
		queryParams,
		headerParams,
		"",
		"GET",
		"/vcenter/namespace-management/cluster-compatibility",
		resultHeaders,
		200,
		errorHeaders,
		map[string]int{})
}


func ClusterCompatibilitySummaryBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["cluster"] = bindings.NewIdType([]string{"ClusterComputeResource"}, "")
	fieldNameMap["cluster"] = "Cluster"
	fields["compatible"] = bindings.NewBooleanType()
	fieldNameMap["compatible"] = "Compatible"
	fields["incompatibility_reason"] = bindings.NewOptionalType(bindings.NewReferenceType(std.LocalizableMessageBindingType))
	fieldNameMap["incompatibility_reason"] = "IncompatibilityReason"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vcenter.namespace_management.cluster_compatibility.summary", fields, reflect.TypeOf(ClusterCompatibilitySummary{}), fieldNameMap, validators)
}

func ClusterCompatibilityFilterSpecBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["compatible"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fieldNameMap["compatible"] = "Compatible"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vcenter.namespace_management.cluster_compatibility.filter_spec", fields, reflect.TypeOf(ClusterCompatibilityFilterSpec{}), fieldNameMap, validators)
}

