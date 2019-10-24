/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Data type definitions file for service: DistributedSwitchCompatibility.
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



// The ``Summary`` class contains information about the compatibility of a vSphere Distributed Switch with the Namespaces feature. **Warning:** This class is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
type DistributedSwitchCompatibilitySummary struct {
    // Identifier of the switch. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
	DistributedSwitch string
    // Compatibility of this switch with vSphere Namespaces. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
	Compatible bool
    // List of reasons for incompatibility. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
	IncompatibilityReasons []std.LocalizableMessage
}

// The ``FilterSpec`` class contains properties used to filter the results when listing Distributed Switches (see DistributedSwitchCompatibility#list) and their compatibility information. **Warning:** This class is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
type DistributedSwitchCompatibilityFilterSpec struct {
    // Compatibility criteria for matching the filter. If true, only Distributed Switches which are compatible with vSphere Namespaces match the filter. If false, only Distributed Switches which are incompatible with vSphere Namespaces match the filter. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
	Compatible *bool
}



func distributedSwitchCompatibilityListInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["cluster"] = bindings.NewIdType([]string{"ClusterComputeResource"}, "")
	fields["filter"] = bindings.NewOptionalType(bindings.NewReferenceType(DistributedSwitchCompatibilityFilterSpecBindingType))
	fieldNameMap["cluster"] = "Cluster"
	fieldNameMap["filter"] = "Filter"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func distributedSwitchCompatibilityListOutputType() bindings.BindingType {
	return bindings.NewListType(bindings.NewReferenceType(DistributedSwitchCompatibilitySummaryBindingType), reflect.TypeOf([]DistributedSwitchCompatibilitySummary{}))
}

func distributedSwitchCompatibilityListRestMetadata() protocol.OperationRestMetadata {
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	paramsTypeMap["filter.compatible"] = bindings.NewOptionalType(bindings.NewBooleanType())
	paramsTypeMap["cluster"] = bindings.NewIdType([]string{"ClusterComputeResource"}, "")
	queryParams["cluster"] = "cluster"
	queryParams["filter.compatible"] = "compatible"
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
		"/vcenter/namespace-management/distributed-switch-compatibility",
		resultHeaders,
		200,
		errorHeaders,
		map[string]int{"Error": 500,"InternalServerError": 500,"NotFound": 404,"Unauthenticated": 401,"Unauthorized": 403})
}


func DistributedSwitchCompatibilitySummaryBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["distributed_switch"] = bindings.NewIdType([]string{"vSphereDistributedSwitch"}, "")
	fieldNameMap["distributed_switch"] = "DistributedSwitch"
	fields["compatible"] = bindings.NewBooleanType()
	fieldNameMap["compatible"] = "Compatible"
	fields["incompatibility_reasons"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewReferenceType(std.LocalizableMessageBindingType), reflect.TypeOf([]std.LocalizableMessage{})))
	fieldNameMap["incompatibility_reasons"] = "IncompatibilityReasons"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vcenter.namespace_management.distributed_switch_compatibility.summary", fields, reflect.TypeOf(DistributedSwitchCompatibilitySummary{}), fieldNameMap, validators)
}

func DistributedSwitchCompatibilityFilterSpecBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["compatible"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fieldNameMap["compatible"] = "Compatible"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vcenter.namespace_management.distributed_switch_compatibility.filter_spec", fields, reflect.TypeOf(DistributedSwitchCompatibilityFilterSpec{}), fieldNameMap, validators)
}

