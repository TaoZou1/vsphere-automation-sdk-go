/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Data type definitions file for service: Policies.
 * Includes binding types of a structures and enumerations defined in the service.
 * Shared by client-side stubs and server-side skeletons to ensure type
 * compatibility.
 */
package storage

import (
	"reflect"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/bindings"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/data"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/protocol"
)

// Resource type for vAPI metadata policy
const Policies_RESOURCE_TYPE = "com.vmware.vcenter.StoragePolicy"


// The ``FilterSpec`` class contains properties used to filter the results when listing the storage policies (see Policies#list)
type PoliciesFilterSpec struct {
    // Identifiers of storage policies that can match the filter.
	Policies map[string]bool
}

// The ``Summary`` class contains commonly used information about a storage policy.
type PoliciesSummary struct {
    // Identifier of the storage policy.
	Policy string
    // Name of the storage policy.
	Name string
    // Description of the storage policy.
	Description string
}

// The ``CompatibleDatastoreInfo`` class contains compatible datastore's information.
type PoliciesCompatibleDatastoreInfo struct {
    // Identifier of the datastore.
	Datastore string
}

// The ``CompatibilityInfo`` class contains info about a list of datastores compatible with a specific storage policy.
type PoliciesCompatibilityInfo struct {
    // Info about a list of datastores compatible with a specific storage policy.
	CompatibleDatastores []PoliciesCompatibleDatastoreInfo
}



func policiesListInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["filter"] = bindings.NewOptionalType(bindings.NewReferenceType(PoliciesFilterSpecBindingType))
	fieldNameMap["filter"] = "Filter"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func policiesListOutputType() bindings.BindingType {
	return bindings.NewListType(bindings.NewReferenceType(PoliciesSummaryBindingType), reflect.TypeOf([]PoliciesSummary{}))
}

func policiesListRestMetadata() protocol.OperationRestMetadata {
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
		"null",
		"",
		resultHeaders,
		0,
		errorHeaders,
		map[string]int{"InvalidArgument": 400,"ServiceUnavailable": 503,"Unauthenticated": 401,"Unauthorized": 403,"UnableToAllocateResource": 500})
}

func policiesCheckCompatibilityInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["policy"] = bindings.NewIdType([]string{"com.vmware.vcenter.StoragePolicy"}, "")
	fields["datastores"] = bindings.NewSetType(bindings.NewIdType([]string{"Datastore"}, ""), reflect.TypeOf(map[string]bool{}))
	fieldNameMap["policy"] = "Policy"
	fieldNameMap["datastores"] = "Datastores"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func policiesCheckCompatibilityOutputType() bindings.BindingType {
	return bindings.NewReferenceType(PoliciesCompatibilityInfoBindingType)
}

func policiesCheckCompatibilityRestMetadata() protocol.OperationRestMetadata {
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
		"null",
		"",
		resultHeaders,
		0,
		errorHeaders,
		map[string]int{"Error": 500,"NotFound": 404,"ServiceUnavailable": 503,"Unauthenticated": 401,"Unauthorized": 403,"UnableToAllocateResource": 500})
}


func PoliciesFilterSpecBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["policies"] = bindings.NewOptionalType(bindings.NewSetType(bindings.NewIdType([]string{"com.vmware.vcenter.StoragePolicy"}, ""), reflect.TypeOf(map[string]bool{})))
	fieldNameMap["policies"] = "Policies"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vcenter.storage.policies.filter_spec", fields, reflect.TypeOf(PoliciesFilterSpec{}), fieldNameMap, validators)
}

func PoliciesSummaryBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["policy"] = bindings.NewIdType([]string{"com.vmware.vcenter.StoragePolicy"}, "")
	fieldNameMap["policy"] = "Policy"
	fields["name"] = bindings.NewStringType()
	fieldNameMap["name"] = "Name"
	fields["description"] = bindings.NewStringType()
	fieldNameMap["description"] = "Description"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vcenter.storage.policies.summary", fields, reflect.TypeOf(PoliciesSummary{}), fieldNameMap, validators)
}

func PoliciesCompatibleDatastoreInfoBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["datastore"] = bindings.NewIdType([]string{"Datastore"}, "")
	fieldNameMap["datastore"] = "Datastore"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vcenter.storage.policies.compatible_datastore_info", fields, reflect.TypeOf(PoliciesCompatibleDatastoreInfo{}), fieldNameMap, validators)
}

func PoliciesCompatibilityInfoBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["compatible_datastores"] = bindings.NewListType(bindings.NewReferenceType(PoliciesCompatibleDatastoreInfoBindingType), reflect.TypeOf([]PoliciesCompatibleDatastoreInfo{}))
	fieldNameMap["compatible_datastores"] = "CompatibleDatastores"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vcenter.storage.policies.compatibility_info", fields, reflect.TypeOf(PoliciesCompatibilityInfo{}), fieldNameMap, validators)
}

