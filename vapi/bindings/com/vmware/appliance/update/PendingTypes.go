/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Data type definitions file for service: Pending.
 * Includes binding types of a structures and enumerations defined in the service.
 * Shared by client-side stubs and server-side skeletons to ensure type
 * compatibility.
 */
package update

import (
	"reflect"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/bindings/com/vmware/appliance"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/bindings/com/vmware/vapi/std"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/bindings"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/data"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/protocol"
	"net/url"
	"time"
)



// The ``SourceType`` enumeration class defines the supported types of sources of updates.
//
// <p> See {@link com.vmware.vapi.bindings.ApiEnumeration enumerated types description}.
type PendingSourceType string

const (
    // Do not perform a new check, return the previous result
	PendingSourceType_LAST_CHECK PendingSourceType = "LAST_CHECK"
    // Check the local sources, ISO devices, staged area
	PendingSourceType_LOCAL PendingSourceType = "LOCAL"
    // Check the local sources, ISO devices, staged area, then online repository as stated in update policy
	PendingSourceType_LOCAL_AND_ONLINE PendingSourceType = "LOCAL_AND_ONLINE"
)

func (s PendingSourceType) PendingSourceType() bool {
	switch s {
	case PendingSourceType_LAST_CHECK:
		return true
	case PendingSourceType_LOCAL:
		return true
	case PendingSourceType_LOCAL_AND_ONLINE:
		return true
	default:
		return false
	}
}


// The ``Info`` class contains the extended information about the update
type PendingInfo struct {
    // Name of the update. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
	Name *string
    // List of the 1. issues addressed since previous/current version 2. new features/improvements
	Contents []std.LocalizableMessage
    // List of the services that will be stopped and restarted during the update installation.
	ServicesWillBeStopped []ServiceInfo
    // List of EULAs. This list has multiple entries and can be dynamic based on what we are actually installing.
	Eulas []std.LocalizableMessage
    // Is the update staged
	Staged bool
    // URL for the knowledge base article describing this update. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
	KnowledgeBase *url.URL
    // Description of the update. The short information what this update is. E.g. "Update2 for vCenter Server Appliance 6.5"
	Description std.LocalizableMessage
    // Update priority
	Priority CommonInfoPriority
    // Update severity
	Severity CommonInfoSeverity
    // Update category
	UpdateType CommonInfoCategory
    // Update release date.
	ReleaseDate time.Time
    // Flag indicating whether reboot is required after update.
	RebootRequired bool
    // Download Size of update in Megabytes.
	Size int64
}

// The ``Question`` class describes a item of information that must be provided by the user in order to install the update.
type PendingQuestion struct {
    // ID of the data item
	DataItem string
    // Label for the item to be used in GUI/CLI
	Text std.LocalizableMessage
    // Description of the item
	Description std.LocalizableMessage
    // How this field shoudl be represented in GUI or CLI
	Type_ PendingQuestionInputType
    // List of allowed values
	AllowedValues []string
    // Regexp to validate the input
	Regexp *string
    // Default answer
	DefaultAnswer *string
}

// The ``InputType`` enumeration class defines representation of field fields in GUI or CLI
//
// <p> See {@link com.vmware.vapi.bindings.ApiEnumeration enumerated types description}.
type PendingQuestionInputType string

const (
    // plain text answer
	PendingQuestionInputType_PLAIN_TEXT PendingQuestionInputType = "PLAIN_TEXT"
    // Yes/No,On/Off,Checkbox answer
	PendingQuestionInputType_BOOLEAN PendingQuestionInputType = "BOOLEAN"
    // Password (masked) answer
	PendingQuestionInputType_PASSWORD PendingQuestionInputType = "PASSWORD"
)

func (i PendingQuestionInputType) PendingQuestionInputType() bool {
	switch i {
	case PendingQuestionInputType_PLAIN_TEXT:
		return true
	case PendingQuestionInputType_BOOLEAN:
		return true
	case PendingQuestionInputType_PASSWORD:
		return true
	default:
		return false
	}
}


// The ``PrecheckResult`` class contains estimates of how long it will take install and rollback an update as well as a list of possible warnings and problems with installing the update.
type PendingPrecheckResult struct {
    // Time when this precheck was run
	CheckTime time.Time
    // Rough estimate of time to install the update (minutes).
	EstimatedTimeToInstall *int64
    // Rough estimate of time to rollback the update (minutes).
	EstimatedTimeToRollback *int64
    // Is reboot required to install the update.
	RebootRequired bool
    // Lists of the issues and warnings
	Issues *appliance.Notifications
    // List of questions that must be answered to install the update.
	Questions []PendingQuestion
}



func pendingListInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["source_type"] = bindings.NewEnumType("com.vmware.appliance.update.pending.source_type", reflect.TypeOf(PendingSourceType(PendingSourceType_LAST_CHECK)))
	fields["url"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["source_type"] = "SourceType"
	fieldNameMap["url"] = "Url"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func pendingListOutputType() bindings.BindingType {
	return bindings.NewListType(bindings.NewReferenceType(SummaryBindingType), reflect.TypeOf([]Summary{}))
}

func pendingListRestMetadata() protocol.OperationRestMetadata {
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
		map[string]int{"Error": 500,"NotFound": 404,"Unauthenticated": 401,"Unauthorized": 403})
}

func pendingGetInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["version"] = bindings.NewIdType([]string{"com.vmware.appliance.update.pending"}, "")
	fieldNameMap["version"] = "Version"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func pendingGetOutputType() bindings.BindingType {
	return bindings.NewReferenceType(PendingInfoBindingType)
}

func pendingGetRestMetadata() protocol.OperationRestMetadata {
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
		map[string]int{"Error": 500,"Unauthenticated": 401,"Unauthorized": 403,"NotFound": 404,"AlreadyInDesiredState": 400})
}

func pendingPrecheckInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["version"] = bindings.NewIdType([]string{"com.vmware.appliance.update.pending"}, "")
	fieldNameMap["version"] = "Version"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func pendingPrecheckOutputType() bindings.BindingType {
	return bindings.NewReferenceType(PendingPrecheckResultBindingType)
}

func pendingPrecheckRestMetadata() protocol.OperationRestMetadata {
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
		map[string]int{"Error": 500,"Unauthenticated": 401,"Unauthorized": 403,"NotFound": 404,"AlreadyInDesiredState": 400,"NotAllowedInCurrentState": 400})
}

func pendingStageInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["version"] = bindings.NewIdType([]string{"com.vmware.appliance.update.pending"}, "")
	fieldNameMap["version"] = "Version"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func pendingStageOutputType() bindings.BindingType {
	return bindings.NewVoidType()
}

func pendingStageRestMetadata() protocol.OperationRestMetadata {
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
		map[string]int{"Error": 500,"Unauthenticated": 401,"Unauthorized": 403,"NotFound": 404,"AlreadyInDesiredState": 400,"AlreadyExists": 400,"NotAllowedInCurrentState": 400})
}

func pendingValidateInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["version"] = bindings.NewIdType([]string{"com.vmware.appliance.update.pending"}, "")
	fields["user_data"] = bindings.NewMapType(bindings.NewIdType([]string{"com.vmware.applicance.update.pending.dataitem"}, ""), bindings.NewStringType(),reflect.TypeOf(map[string]string{}))
	fieldNameMap["version"] = "Version"
	fieldNameMap["user_data"] = "UserData"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func pendingValidateOutputType() bindings.BindingType {
	return bindings.NewReferenceType(appliance.NotificationsBindingType)
}

func pendingValidateRestMetadata() protocol.OperationRestMetadata {
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
		map[string]int{"Error": 500,"Unauthenticated": 401,"Unauthorized": 403,"NotFound": 404,"AlreadyInDesiredState": 400,"NotAllowedInCurrentState": 400})
}

func pendingInstallInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["version"] = bindings.NewIdType([]string{"com.vmware.appliance.update.pending"}, "")
	fields["user_data"] = bindings.NewMapType(bindings.NewIdType([]string{"com.vmware.applicance.update.pending.dataitem"}, ""), bindings.NewStringType(),reflect.TypeOf(map[string]string{}))
	fieldNameMap["version"] = "Version"
	fieldNameMap["user_data"] = "UserData"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func pendingInstallOutputType() bindings.BindingType {
	return bindings.NewVoidType()
}

func pendingInstallRestMetadata() protocol.OperationRestMetadata {
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
		map[string]int{"Error": 500,"Unauthenticated": 401,"Unauthorized": 403,"NotFound": 404,"AlreadyInDesiredState": 400,"NotAllowedInCurrentState": 400})
}

func pendingStageAndInstallInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["version"] = bindings.NewIdType([]string{"com.vmware.appliance.update.pending"}, "")
	fields["user_data"] = bindings.NewMapType(bindings.NewIdType([]string{"com.vmware.applicance.update.pending.dataitem"}, ""), bindings.NewStringType(),reflect.TypeOf(map[string]string{}))
	fieldNameMap["version"] = "Version"
	fieldNameMap["user_data"] = "UserData"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func pendingStageAndInstallOutputType() bindings.BindingType {
	return bindings.NewVoidType()
}

func pendingStageAndInstallRestMetadata() protocol.OperationRestMetadata {
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
		map[string]int{"Error": 500,"Unauthenticated": 401,"Unauthorized": 403,"NotFound": 404,"AlreadyInDesiredState": 400,"NotAllowedInCurrentState": 400})
}


func PendingInfoBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["name"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["name"] = "Name"
	fields["contents"] = bindings.NewListType(bindings.NewReferenceType(std.LocalizableMessageBindingType), reflect.TypeOf([]std.LocalizableMessage{}))
	fieldNameMap["contents"] = "Contents"
	fields["services_will_be_stopped"] = bindings.NewListType(bindings.NewReferenceType(ServiceInfoBindingType), reflect.TypeOf([]ServiceInfo{}))
	fieldNameMap["services_will_be_stopped"] = "ServicesWillBeStopped"
	fields["eulas"] = bindings.NewListType(bindings.NewReferenceType(std.LocalizableMessageBindingType), reflect.TypeOf([]std.LocalizableMessage{}))
	fieldNameMap["eulas"] = "Eulas"
	fields["staged"] = bindings.NewBooleanType()
	fieldNameMap["staged"] = "Staged"
	fields["knowledge_base"] = bindings.NewOptionalType(bindings.NewUriType())
	fieldNameMap["knowledge_base"] = "KnowledgeBase"
	fields["description"] = bindings.NewReferenceType(std.LocalizableMessageBindingType)
	fieldNameMap["description"] = "Description"
	fields["priority"] = bindings.NewEnumType("com.vmware.appliance.update.common_info.priority", reflect.TypeOf(CommonInfoPriority(CommonInfoPriority_HIGH)))
	fieldNameMap["priority"] = "Priority"
	fields["severity"] = bindings.NewEnumType("com.vmware.appliance.update.common_info.severity", reflect.TypeOf(CommonInfoSeverity(CommonInfoSeverity_CRITICAL)))
	fieldNameMap["severity"] = "Severity"
	fields["update_type"] = bindings.NewEnumType("com.vmware.appliance.update.common_info.category", reflect.TypeOf(CommonInfoCategory(CommonInfoCategory_SECURITY)))
	fieldNameMap["update_type"] = "UpdateType"
	fields["release_date"] = bindings.NewDateTimeType()
	fieldNameMap["release_date"] = "ReleaseDate"
	fields["reboot_required"] = bindings.NewBooleanType()
	fieldNameMap["reboot_required"] = "RebootRequired"
	fields["size"] = bindings.NewIntegerType()
	fieldNameMap["size"] = "Size"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.appliance.update.pending.info", fields, reflect.TypeOf(PendingInfo{}), fieldNameMap, validators)
}

func PendingQuestionBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["data_item"] = bindings.NewIdType([]string{"com.vmware.applicance.update.pending.dataitem"}, "")
	fieldNameMap["data_item"] = "DataItem"
	fields["text"] = bindings.NewReferenceType(std.LocalizableMessageBindingType)
	fieldNameMap["text"] = "Text"
	fields["description"] = bindings.NewReferenceType(std.LocalizableMessageBindingType)
	fieldNameMap["description"] = "Description"
	fields["type"] = bindings.NewEnumType("com.vmware.appliance.update.pending.question.input_type", reflect.TypeOf(PendingQuestionInputType(PendingQuestionInputType_PLAIN_TEXT)))
	fieldNameMap["type"] = "Type_"
	fields["allowed_values"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewStringType(), reflect.TypeOf([]string{})))
	fieldNameMap["allowed_values"] = "AllowedValues"
	fields["regexp"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["regexp"] = "Regexp"
	fields["default_answer"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["default_answer"] = "DefaultAnswer"
	var validators = []bindings.Validator{}
	uv1 := bindings.NewUnionValidator("type",
		map[string][]bindings.FieldData{
			"PLAIN_TEXT": []bindings.FieldData{
				bindings.NewFieldData("allowed_values", false),
				bindings.NewFieldData("regexp", false),
				bindings.NewFieldData("default_answer", false),
			},
			"BOOLEAN": []bindings.FieldData{},
			"PASSWORD": []bindings.FieldData{},
		},
	)
	validators = append(validators, uv1)
	return bindings.NewStructType("com.vmware.appliance.update.pending.question", fields, reflect.TypeOf(PendingQuestion{}), fieldNameMap, validators)
}

func PendingPrecheckResultBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["check_time"] = bindings.NewDateTimeType()
	fieldNameMap["check_time"] = "CheckTime"
	fields["estimated_time_to_install"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fieldNameMap["estimated_time_to_install"] = "EstimatedTimeToInstall"
	fields["estimated_time_to_rollback"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fieldNameMap["estimated_time_to_rollback"] = "EstimatedTimeToRollback"
	fields["reboot_required"] = bindings.NewBooleanType()
	fieldNameMap["reboot_required"] = "RebootRequired"
	fields["issues"] = bindings.NewOptionalType(bindings.NewReferenceType(appliance.NotificationsBindingType))
	fieldNameMap["issues"] = "Issues"
	fields["questions"] = bindings.NewListType(bindings.NewReferenceType(PendingQuestionBindingType), reflect.TypeOf([]PendingQuestion{}))
	fieldNameMap["questions"] = "Questions"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.appliance.update.pending.precheck_result", fields, reflect.TypeOf(PendingPrecheckResult{}), fieldNameMap, validators)
}

