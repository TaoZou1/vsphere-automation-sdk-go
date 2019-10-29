/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Code generated. DO NOT EDIT.

/*
 * Data type definitions file for service: Host.
 * Includes binding types of a structures and enumerations defined in the service.
 * Shared by client-side stubs and server-side skeletons to ensure type
 * compatibility.
 */

package tags

import (
	"reflect"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/bindings"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/data"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/protocol"
)



// The ``FilterSpec`` class contains properties used to filter the results when listing tags associated to hosts, see Host#list. **Warning:** This class is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
type HostFilterSpec struct {
    // Tag associations need to have one of the tag identifiers in this set to match the filter. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
	Tags map[string]bool
}

// The ``Summary`` class contains commonly used information about a tag that is assigned to a host. **Warning:** This class is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
type HostSummary struct {
    // The identifier of a tag. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
	Tag string
    // Number of hosts associated with the tag. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
	Count int64
}

// The ``ListResult`` class contains commonly used information about tags that are assigned to hosts. See Host#list. **Warning:** This class is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
type HostListResult struct {
    // The list of tag summaries. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
	Tags []HostSummary
}



func hostListInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["filter"] = bindings.NewOptionalType(bindings.NewReferenceType(HostFilterSpecBindingType))
	fieldNameMap["filter"] = "Filter"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func hostListOutputType() bindings.BindingType {
	return bindings.NewReferenceType(HostListResultBindingType)
}

func hostListRestMetadata() protocol.OperationRestMetadata {
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	resultHeaders := map[string]string{}
	errorHeaders := map[string]string{}
	return protocol.NewOperationRestMetadata(
		paramsTypeMap,
		pathParams,
		queryParams,
		headerParams,
		"",
		"null",
		"",
		resultHeaders,
		0,
		errorHeaders,
		map[string]int{"Unauthorized": 403})
}

func hostDeleteInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["tag"] = bindings.NewIdType([]string{"com.vmware.cis.tagging.Tag:HostSystem"}, "")
	fieldNameMap["tag"] = "Tag"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func hostDeleteOutputType() bindings.BindingType {
	return bindings.NewVoidType()
}

func hostDeleteRestMetadata() protocol.OperationRestMetadata {
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	resultHeaders := map[string]string{}
	errorHeaders := map[string]string{}
	return protocol.NewOperationRestMetadata(
		paramsTypeMap,
		pathParams,
		queryParams,
		headerParams,
		"",
		"null",
		"",
		resultHeaders,
		0,
		errorHeaders,
		map[string]int{"NotFound": 404,"Unauthorized": 403})
}


func HostFilterSpecBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["tags"] = bindings.NewOptionalType(bindings.NewSetType(bindings.NewIdType([]string{"com.vmware.cis.tagging.Tag:HostSystem"}, ""), reflect.TypeOf(map[string]bool{})))
	fieldNameMap["tags"] = "Tags"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vcenter.tags.host.filter_spec", fields, reflect.TypeOf(HostFilterSpec{}), fieldNameMap, validators)
}

func HostSummaryBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["tag"] = bindings.NewIdType([]string{"com.vmware.cis.tagging.Tag:HostSystem"}, "")
	fieldNameMap["tag"] = "Tag"
	fields["count"] = bindings.NewIntegerType()
	fieldNameMap["count"] = "Count"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vcenter.tags.host.summary", fields, reflect.TypeOf(HostSummary{}), fieldNameMap, validators)
}

func HostListResultBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["tags"] = bindings.NewListType(bindings.NewReferenceType(HostSummaryBindingType), reflect.TypeOf([]HostSummary{}))
	fieldNameMap["tags"] = "Tags"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vcenter.tags.host.list_result", fields, reflect.TypeOf(HostListResult{}), fieldNameMap, validators)
}

