/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Code generated. DO NOT EDIT.

/*
 * Data type definitions file for service: Clusters.
 * Includes binding types of a structures and enumerations defined in the service.
 * Shared by client-side stubs and server-side skeletons to ensure type
 * compatibility.
 */

package software

import (
	"reflect"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/bindings/vapi/std"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/bindings"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/data"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/protocol"
	"time"
)



// The ``State`` enumeration class describes the state of the upgrade. **Warning:** This enumeration is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
//
// <p> See {@link com.vmware.vapi.bindings.ApiEnumeration enumerated types description}.
type ClustersState string

const (
    // Upgrade is in progress. **Warning:** This constant field is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
	ClustersState_PENDING ClustersState = "PENDING"
    // Cluster is ready when there is no upgrade or upgrade is completed. **Warning:** This constant field is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
	ClustersState_READY ClustersState = "READY"
    // Upgrade failed and need user intervention. **Warning:** This constant field is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
	ClustersState_ERROR ClustersState = "ERROR"
)

func (s ClustersState) ClustersState() bool {
	switch s {
	case ClustersState_PENDING:
		return true
	case ClustersState_READY:
		return true
	case ClustersState_ERROR:
		return true
	default:
		return false
	}
}


// The ``Result`` class contains the result of batch upgrade method. **Warning:** This class is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
type ClustersResult struct {
    // The result of batch upgrade method. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
	Res ClustersResultRes
    // Exception when cluster pre-check failed during upgrade invocation. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
	Exception *data.ErrorValue
}

// The ``Res`` enumeration class represents the upgrade invocation result for each cluster. **Warning:** This enumeration is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
//
// <p> See {@link com.vmware.vapi.bindings.ApiEnumeration enumerated types description}.
type ClustersResultRes string

const (
    // Upgrade is started. **Warning:** This constant field is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
	ClustersResultRes_STARTED ClustersResultRes = "STARTED"
    // Upgrade is rejected. This implies pre-check failed when invoking upgrade of the cluster. **Warning:** This constant field is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
	ClustersResultRes_REJECTED ClustersResultRes = "REJECTED"
)

func (r ClustersResultRes) ClustersResultRes() bool {
	switch r {
	case ClustersResultRes_STARTED:
		return true
	case ClustersResultRes_REJECTED:
		return true
	default:
		return false
	}
}


// The ``UpgradeSpec`` class contains the specification required to upgrade a cluster. **Warning:** This class is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
type ClustersUpgradeSpec struct {
    // Version number the cluster is going to be upgraded to. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
	DesiredVersion string
}

// The ``Summary`` class contains basic information about the cluster upgrade related information. **Warning:** This class is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
type ClustersSummary struct {
    // Identifier for the cluster. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
	Cluster string
    // Name of the cluster. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
	ClusterName string
    // Current version of the cluster. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
	CurrentVersion string
    // Set of versions available for upgrade. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
	AvailableVersions []string
    // Date of last successful upgrade. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
	LastUpgradedDate *time.Time
    // Desired version the cluster will be upgraded to. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
	DesiredVersion *string
    // Current state of the upgrade. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
	State ClustersState
}

// The ``Info`` class contains detailed information about the cluster upgrade status and related information. **Warning:** This class is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
type ClustersInfo struct {
    // Current version of the cluster. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
	CurrentVersion string
    // Set of available versions can be upgraded to. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
	AvailableVersions []string
    // Date of last successful upgrade. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
	LastUpgradedDate *time.Time
    // Current set of messages associated with the cluster version. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
	Messages []ClustersMessage
    // Current state of the upgrade. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
	State ClustersState
    // Information about upgrade in progress. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
	UpgradeStatus *ClustersUpgradeStatus
}

// The ``UpgradeStatus`` class contains detailed information about the cluster when upgraded is in progress. **Warning:** This class is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
type ClustersUpgradeStatus struct {
    // Desired version the cluster will be upgraded to. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
	DesiredVersion *string
    // Current set of messages associated with the upgrade state. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
	Messages []ClustersMessage
    // Information about upgrade progess. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
	Progress *ClustersUpgradeProgress
}

// The ``UpgradeProgress`` class contains detailed information about the cluster upgrade progess. **Warning:** This class is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
type ClustersUpgradeProgress struct {
    // Total amount of the work for the operation. The work here represents the number of master nodes in the cluster need to be upgraded. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
	Total int64
    // The amount of work completed for the operation. The value can only be incremented. The number or master nodes which upgrade completed. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
	Completed int64
    // Message about the work progress. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
	Message std.LocalizableMessage
}

// The ``Message`` class contains the information about the object configuration. **Warning:** This class is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
type ClustersMessage struct {
    // Type of the message. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
	Severity ClustersMessageSeverity
    // Details about the message. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
	Details std.LocalizableMessage
}

// The ``Severity`` enumeration class represents the severity of the message. **Warning:** This enumeration is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
//
// <p> See {@link com.vmware.vapi.bindings.ApiEnumeration enumerated types description}.
type ClustersMessageSeverity string

const (
    // Informational message. This may be accompanied by vCenter event. **Warning:** This constant field is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
	ClustersMessageSeverity_INFO ClustersMessageSeverity = "INFO"
    // Warning message. This may be accompanied by vCenter event. **Warning:** This constant field is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
	ClustersMessageSeverity_WARNING ClustersMessageSeverity = "WARNING"
    // Error message. This is accompanied by vCenter event and/or alarm. **Warning:** This constant field is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
	ClustersMessageSeverity_ERROR ClustersMessageSeverity = "ERROR"
)

func (s ClustersMessageSeverity) ClustersMessageSeverity() bool {
	switch s {
	case ClustersMessageSeverity_INFO:
		return true
	case ClustersMessageSeverity_WARNING:
		return true
	case ClustersMessageSeverity_ERROR:
		return true
	default:
		return false
	}
}




func clustersUpgradeInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["cluster"] = bindings.NewIdType([]string{"ClusterComputeResource"}, "")
	fields["spec"] = bindings.NewReferenceType(ClustersUpgradeSpecBindingType)
	fieldNameMap["cluster"] = "Cluster"
	fieldNameMap["spec"] = "Spec"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func clustersUpgradeOutputType() bindings.BindingType {
	return bindings.NewVoidType()
}

func clustersUpgradeRestMetadata() protocol.OperationRestMetadata {
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	paramsTypeMap["cluster"] = bindings.NewIdType([]string{"ClusterComputeResource"}, "")
	paramsTypeMap["spec"] = bindings.NewReferenceType(ClustersUpgradeSpecBindingType)
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
		"spec",
		"POST",
		"/vcenter/namespace-management/software/clusters/{cluster}",
		resultHeaders,
		204,
		errorHeaders,
		map[string]int{"Error": 500,"Unsupported": 400,"NotFound": 404,"Unauthenticated": 401,"Unauthorized": 403})
}

func clustersUpgradeMultipleInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["specs"] = bindings.NewMapType(bindings.NewIdType([]string{"ClusterComputeResource"}, ""), bindings.NewReferenceType(ClustersUpgradeSpecBindingType),reflect.TypeOf(map[string]ClustersUpgradeSpec{}))
	fieldNameMap["specs"] = "Specs"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func clustersUpgradeMultipleOutputType() bindings.BindingType {
	return bindings.NewMapType(bindings.NewIdType([]string{"ClusterComputeResource"}, ""), bindings.NewReferenceType(ClustersResultBindingType),reflect.TypeOf(map[string]ClustersResult{}))
}

func clustersUpgradeMultipleRestMetadata() protocol.OperationRestMetadata {
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	paramsTypeMap["specs"] = bindings.NewMapType(bindings.NewIdType([]string{"ClusterComputeResource"}, ""), bindings.NewReferenceType(ClustersUpgradeSpecBindingType),reflect.TypeOf(map[string]ClustersUpgradeSpec{}))
	resultHeaders := map[string]string{}
	errorHeaders := map[string]string{}
	errorHeaders["Unauthenticated.challenge"] = "WWW-Authenticate"
	return protocol.NewOperationRestMetadata(
		paramsTypeMap,
		pathParams,
		queryParams,
		headerParams,
		"specs",
		"POST",
		"/vcenter/namespace-management/software/clusters",
		resultHeaders,
		200,
		errorHeaders,
		map[string]int{"Error": 500,"Unauthenticated": 401,"Unauthorized": 403})
}

func clustersGetInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["cluster"] = bindings.NewIdType([]string{"ClusterComputeResource"}, "")
	fieldNameMap["cluster"] = "Cluster"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func clustersGetOutputType() bindings.BindingType {
	return bindings.NewReferenceType(ClustersInfoBindingType)
}

func clustersGetRestMetadata() protocol.OperationRestMetadata {
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
		"GET",
		"/vcenter/namespace-management/software/clusters/{cluster}",
		resultHeaders,
		200,
		errorHeaders,
		map[string]int{"NotFound": 404,"Error": 500,"Unauthenticated": 401,"Unauthorized": 403})
}

func clustersListInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func clustersListOutputType() bindings.BindingType {
	return bindings.NewListType(bindings.NewReferenceType(ClustersSummaryBindingType), reflect.TypeOf([]ClustersSummary{}))
}

func clustersListRestMetadata() protocol.OperationRestMetadata {
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
		"/vcenter/namespace-management/software/clusters",
		resultHeaders,
		200,
		errorHeaders,
		map[string]int{"Error": 500,"Unauthenticated": 401,"Unauthorized": 403})
}


func ClustersResultBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["res"] = bindings.NewEnumType("com.vmware.vcenter.namespace_management.software.clusters.result.res", reflect.TypeOf(ClustersResultRes(ClustersResultRes_STARTED)))
	fieldNameMap["res"] = "Res"
	fields["exception"] = bindings.NewOptionalType(bindings.NewAnyErrorType())
	fieldNameMap["exception"] = "Exception"
	var validators = []bindings.Validator{}
	uv1 := bindings.NewUnionValidator("res",
		map[string][]bindings.FieldData{
			"REJECTED": []bindings.FieldData{
				bindings.NewFieldData("exception", true),
			},
			"STARTED": []bindings.FieldData{},
		},
	)
	validators = append(validators, uv1)
	return bindings.NewStructType("com.vmware.vcenter.namespace_management.software.clusters.result", fields, reflect.TypeOf(ClustersResult{}), fieldNameMap, validators)
}

func ClustersUpgradeSpecBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["desired_version"] = bindings.NewStringType()
	fieldNameMap["desired_version"] = "DesiredVersion"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vcenter.namespace_management.software.clusters.upgrade_spec", fields, reflect.TypeOf(ClustersUpgradeSpec{}), fieldNameMap, validators)
}

func ClustersSummaryBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["cluster"] = bindings.NewIdType([]string{"ClusterComputeResource"}, "")
	fieldNameMap["cluster"] = "Cluster"
	fields["cluster_name"] = bindings.NewIdType([]string{"ClusterComputeResource.name"}, "")
	fieldNameMap["cluster_name"] = "ClusterName"
	fields["current_version"] = bindings.NewStringType()
	fieldNameMap["current_version"] = "CurrentVersion"
	fields["available_versions"] = bindings.NewListType(bindings.NewStringType(), reflect.TypeOf([]string{}))
	fieldNameMap["available_versions"] = "AvailableVersions"
	fields["last_upgraded_date"] = bindings.NewOptionalType(bindings.NewDateTimeType())
	fieldNameMap["last_upgraded_date"] = "LastUpgradedDate"
	fields["desired_version"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["desired_version"] = "DesiredVersion"
	fields["state"] = bindings.NewEnumType("com.vmware.vcenter.namespace_management.software.clusters.state", reflect.TypeOf(ClustersState(ClustersState_PENDING)))
	fieldNameMap["state"] = "State"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vcenter.namespace_management.software.clusters.summary", fields, reflect.TypeOf(ClustersSummary{}), fieldNameMap, validators)
}

func ClustersInfoBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["current_version"] = bindings.NewStringType()
	fieldNameMap["current_version"] = "CurrentVersion"
	fields["available_versions"] = bindings.NewListType(bindings.NewStringType(), reflect.TypeOf([]string{}))
	fieldNameMap["available_versions"] = "AvailableVersions"
	fields["last_upgraded_date"] = bindings.NewOptionalType(bindings.NewDateTimeType())
	fieldNameMap["last_upgraded_date"] = "LastUpgradedDate"
	fields["messages"] = bindings.NewListType(bindings.NewReferenceType(ClustersMessageBindingType), reflect.TypeOf([]ClustersMessage{}))
	fieldNameMap["messages"] = "Messages"
	fields["state"] = bindings.NewEnumType("com.vmware.vcenter.namespace_management.software.clusters.state", reflect.TypeOf(ClustersState(ClustersState_PENDING)))
	fieldNameMap["state"] = "State"
	fields["upgrade_status"] = bindings.NewOptionalType(bindings.NewReferenceType(ClustersUpgradeStatusBindingType))
	fieldNameMap["upgrade_status"] = "UpgradeStatus"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vcenter.namespace_management.software.clusters.info", fields, reflect.TypeOf(ClustersInfo{}), fieldNameMap, validators)
}

func ClustersUpgradeStatusBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["desired_version"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["desired_version"] = "DesiredVersion"
	fields["messages"] = bindings.NewListType(bindings.NewReferenceType(ClustersMessageBindingType), reflect.TypeOf([]ClustersMessage{}))
	fieldNameMap["messages"] = "Messages"
	fields["progress"] = bindings.NewOptionalType(bindings.NewReferenceType(ClustersUpgradeProgressBindingType))
	fieldNameMap["progress"] = "Progress"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vcenter.namespace_management.software.clusters.upgrade_status", fields, reflect.TypeOf(ClustersUpgradeStatus{}), fieldNameMap, validators)
}

func ClustersUpgradeProgressBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["total"] = bindings.NewIntegerType()
	fieldNameMap["total"] = "Total"
	fields["completed"] = bindings.NewIntegerType()
	fieldNameMap["completed"] = "Completed"
	fields["message"] = bindings.NewReferenceType(std.LocalizableMessageBindingType)
	fieldNameMap["message"] = "Message"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vcenter.namespace_management.software.clusters.upgrade_progress", fields, reflect.TypeOf(ClustersUpgradeProgress{}), fieldNameMap, validators)
}

func ClustersMessageBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["severity"] = bindings.NewEnumType("com.vmware.vcenter.namespace_management.software.clusters.message.severity", reflect.TypeOf(ClustersMessageSeverity(ClustersMessageSeverity_INFO)))
	fieldNameMap["severity"] = "Severity"
	fields["details"] = bindings.NewReferenceType(std.LocalizableMessageBindingType)
	fieldNameMap["details"] = "Details"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vcenter.namespace_management.software.clusters.message", fields, reflect.TypeOf(ClustersMessage{}), fieldNameMap, validators)
}

