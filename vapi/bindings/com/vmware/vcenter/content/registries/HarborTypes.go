/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Data type definitions file for service: Harbor.
 * Includes binding types of a structures and enumerations defined in the service.
 * Shared by client-side stubs and server-side skeletons to ensure type
 * compatibility.
 */
package registries

import (
	"reflect"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/bindings"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/data"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/protocol"
	"net/url"
	"time"
)



// The ``StorageSpec`` class contains the specification required to configure storage associated with a Harbor registry. In this version, Harbor registry is created in Kubernetes environment, information in this class will result in storage quotas on a Kubernetes namespace. **Warning:** This class is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
type HarborStorageSpec struct {
    // Identifier of the storage policy. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
	Policy string
    // The maximum amount of storage (in mebibytes) which can be utilized by the registry for this specification. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
	Limit *int64
}

// The ``StorageInfo`` class contains the detailed information about storage used by the Harbor registry. **Warning:** This class is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
type HarborStorageInfo struct {
    // Identifier of the storage policy. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
	Policy string
    // Total capacity for the registry storage (in mebibytes). This is the storage limit set on the Harbor registry. If a storage limit was not set on the registry, the default registry capacity - 204800 mebibytes is used. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
	Capacity int64
    // Overall storage used by the registry (in mebibytes). This is the sum of used storage associated with storage policies configured for the registry. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
	Used int64
}

// The ``GarbageCollection`` class contains garbage collection configuration for the Harbor registry. **Warning:** This class is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
type HarborGarbageCollection struct {
    // Frequency of garbage collection. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
	Type_ Recurrence
    // Day of the week when garbage collection should run. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
	DayOfWeek *DayOfWeek
    // Hour at which garbage collection should run. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
	Hour *int64
    // Minute at which garbage collection should run. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
	Minute *int64
}

// The ``CreateSpec`` class contains the specification required to create a Harbor registry. **Warning:** This class is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
type HarborCreateSpec struct {
    // Identifier of the cluster hosting the registry. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
	Cluster *string
    // Garbage collection configuration for the Harbor registry. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
	GarbageCollection *HarborGarbageCollection
    // Storage associated with the Harbor registry. The list contains only one storage backing in this version. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
	Storage []HarborStorageSpec
}

// The ``Summary`` class contains basic information about a running Harbor registry. **Warning:** This class is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
type HarborSummary struct {
    // Identifier of the cluster. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
	Cluster *string
    // Identifier of the registry. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
	Registry string
    // Version of the registry. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
	Version string
    // URL to access the UI of the registry. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
	UiAccessUrl url.URL
}

// The ``Info`` class contains detailed information about a running Harbor registry. **Warning:** This class is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
type HarborInfo struct {
    // Identifier of the cluster. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
	Cluster *string
    // Identifier of the Harbor namespace in case it is created in a Kubernetes environment. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
	Namespace *string
    // Version of the registry. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
	Version string
    // The date and time when the harbor registry was created. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
	CreationTime time.Time
    // URL to access the UI of the registry. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
	UiAccessUrl url.URL
    // Harbor certificate chain in base64 format. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
	CertChain []string
    // Garbage collection information for the registry. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
	GarbageCollection HarborGarbageCollection
    // Storage information associated with the registry. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
	Storage []HarborStorageInfo
    // Health status of the container registry. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
	Health HealthInfo
}



func harborCreateInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["client_token"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["spec"] = bindings.NewReferenceType(HarborCreateSpecBindingType)
	fieldNameMap["client_token"] = "ClientToken"
	fieldNameMap["spec"] = "Spec"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func harborCreateOutputType() bindings.BindingType {
	return bindings.NewIdType([]string{"com.vmware.vcenter.content.Registry"}, "")
}

func harborCreateRestMetadata() protocol.OperationRestMetadata {
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
		map[string]int{"AlreadyExists": 400,"NotFound": 404,"Unsupported": 400,"InvalidArgument": 400,"Error": 500,"Unauthenticated": 401,"Unauthorized": 403})
}

func harborDeleteInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["registry"] = bindings.NewIdType([]string{"com.vmware.vcenter.content.Registry"}, "")
	fieldNameMap["registry"] = "Registry"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func harborDeleteOutputType() bindings.BindingType {
	return bindings.NewVoidType()
}

func harborDeleteRestMetadata() protocol.OperationRestMetadata {
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
		map[string]int{"NotFound": 404,"Error": 500,"Unauthenticated": 401,"Unauthorized": 403})
}

func harborGetInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["registry"] = bindings.NewIdType([]string{"com.vmware.vcenter.content.Registry"}, "")
	fieldNameMap["registry"] = "Registry"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func harborGetOutputType() bindings.BindingType {
	return bindings.NewReferenceType(HarborInfoBindingType)
}

func harborGetRestMetadata() protocol.OperationRestMetadata {
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
		map[string]int{"NotFound": 404,"Error": 500,"Unauthenticated": 401,"Unauthorized": 403})
}

func harborListInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func harborListOutputType() bindings.BindingType {
	return bindings.NewListType(bindings.NewReferenceType(HarborSummaryBindingType), reflect.TypeOf([]HarborSummary{}))
}

func harborListRestMetadata() protocol.OperationRestMetadata {
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
		map[string]int{"Error": 500,"Unauthenticated": 401,"Unauthorized": 403})
}


func HarborStorageSpecBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["policy"] = bindings.NewIdType([]string{"SpsStorageProfile"}, "")
	fieldNameMap["policy"] = "Policy"
	fields["limit"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fieldNameMap["limit"] = "Limit"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vcenter.content.registries.harbor.storage_spec", fields, reflect.TypeOf(HarborStorageSpec{}), fieldNameMap, validators)
}

func HarborStorageInfoBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["policy"] = bindings.NewIdType([]string{"SpsStorageProfile"}, "")
	fieldNameMap["policy"] = "Policy"
	fields["capacity"] = bindings.NewIntegerType()
	fieldNameMap["capacity"] = "Capacity"
	fields["used"] = bindings.NewIntegerType()
	fieldNameMap["used"] = "Used"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vcenter.content.registries.harbor.storage_info", fields, reflect.TypeOf(HarborStorageInfo{}), fieldNameMap, validators)
}

func HarborGarbageCollectionBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["type"] = bindings.NewEnumType("com.vmware.vcenter.content.registries.recurrence", reflect.TypeOf(Recurrence(Recurrence_NONE)))
	fieldNameMap["type"] = "Type_"
	fields["day_of_week"] = bindings.NewOptionalType(bindings.NewEnumType("com.vmware.vcenter.content.registries.day_of_week", reflect.TypeOf(DayOfWeek(DayOfWeek_SUNDAY))))
	fieldNameMap["day_of_week"] = "DayOfWeek"
	fields["hour"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fieldNameMap["hour"] = "Hour"
	fields["minute"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fieldNameMap["minute"] = "Minute"
	var validators = []bindings.Validator{}
	uv1 := bindings.NewUnionValidator("type",
		map[string][]bindings.FieldData{
			"WEEKLY": []bindings.FieldData{
				bindings.NewFieldData("day_of_week", true),
				bindings.NewFieldData("hour", true),
				bindings.NewFieldData("minute", true),
			},
			"DAILY": []bindings.FieldData{
				bindings.NewFieldData("hour", true),
				bindings.NewFieldData("minute", true),
			},
			"NONE": []bindings.FieldData{},
		},
	)
	validators = append(validators, uv1)
	return bindings.NewStructType("com.vmware.vcenter.content.registries.harbor.garbage_collection", fields, reflect.TypeOf(HarborGarbageCollection{}), fieldNameMap, validators)
}

func HarborCreateSpecBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["cluster"] = bindings.NewOptionalType(bindings.NewIdType([]string{"ClusterComputeResource"}, ""))
	fieldNameMap["cluster"] = "Cluster"
	fields["garbage_collection"] = bindings.NewOptionalType(bindings.NewReferenceType(HarborGarbageCollectionBindingType))
	fieldNameMap["garbage_collection"] = "GarbageCollection"
	fields["storage"] = bindings.NewListType(bindings.NewReferenceType(HarborStorageSpecBindingType), reflect.TypeOf([]HarborStorageSpec{}))
	fieldNameMap["storage"] = "Storage"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vcenter.content.registries.harbor.create_spec", fields, reflect.TypeOf(HarborCreateSpec{}), fieldNameMap, validators)
}

func HarborSummaryBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["cluster"] = bindings.NewOptionalType(bindings.NewIdType([]string{"ClusterComputeResource"}, ""))
	fieldNameMap["cluster"] = "Cluster"
	fields["registry"] = bindings.NewIdType([]string{"com.vmware.vcenter.content.Registry"}, "")
	fieldNameMap["registry"] = "Registry"
	fields["version"] = bindings.NewStringType()
	fieldNameMap["version"] = "Version"
	fields["ui_access_url"] = bindings.NewUriType()
	fieldNameMap["ui_access_url"] = "UiAccessUrl"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vcenter.content.registries.harbor.summary", fields, reflect.TypeOf(HarborSummary{}), fieldNameMap, validators)
}

func HarborInfoBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["cluster"] = bindings.NewOptionalType(bindings.NewIdType([]string{"ClusterComputeResource"}, ""))
	fieldNameMap["cluster"] = "Cluster"
	fields["namespace"] = bindings.NewOptionalType(bindings.NewIdType([]string{"NamespaceInstance"}, ""))
	fieldNameMap["namespace"] = "Namespace"
	fields["version"] = bindings.NewStringType()
	fieldNameMap["version"] = "Version"
	fields["creation_time"] = bindings.NewDateTimeType()
	fieldNameMap["creation_time"] = "CreationTime"
	fields["ui_access_url"] = bindings.NewUriType()
	fieldNameMap["ui_access_url"] = "UiAccessUrl"
	fields["cert_chain"] = bindings.NewListType(bindings.NewStringType(), reflect.TypeOf([]string{}))
	fieldNameMap["cert_chain"] = "CertChain"
	fields["garbage_collection"] = bindings.NewReferenceType(HarborGarbageCollectionBindingType)
	fieldNameMap["garbage_collection"] = "GarbageCollection"
	fields["storage"] = bindings.NewListType(bindings.NewReferenceType(HarborStorageInfoBindingType), reflect.TypeOf([]HarborStorageInfo{}))
	fieldNameMap["storage"] = "Storage"
	fields["health"] = bindings.NewReferenceType(HealthInfoBindingType)
	fieldNameMap["health"] = "Health"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vcenter.content.registries.harbor.info", fields, reflect.TypeOf(HarborInfo{}), fieldNameMap, validators)
}

