/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Code generated. DO NOT EDIT.

/*
 * Data type definitions file for service: Announcements.
 * Includes binding types of a structures and enumerations defined in the service.
 * Shared by client-side stubs and server-side skeletons to ensure type
 * compatibility.
 */

package system

import (
	"reflect"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/bindings/vapi/std"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/bindings"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/data"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/protocol"
	"time"
)

// Resource type for the announcements. **Warning:** This constant field is available as Technology Preview. These are early access APIs provided to test, automate and provide feedback on the feature. Since this can change based on feedback, VMware does not guarantee backwards compatibility and recommends against using them in production environments. Some Technology Preview APIs might only be applicable to specific environments.
const Announcements_RESOURCE_TYPE = "com.vmware.vcenter.system.announcement"


// The ``Severity`` enumeration class defines the severity of the announcements. **Warning:** This enumeration is available as Technology Preview. These are early access APIs provided to test, automate and provide feedback on the feature. Since this can change based on feedback, VMware does not guarantee backwards compatibility and recommends against using them in production environments. Some Technology Preview APIs might only be applicable to specific environments.
//
// <p> See {@link com.vmware.vapi.bindings.ApiEnumeration enumerated types description}.
type AnnouncementsSeverity string

const (
    // Critical problem. **Warning:** This constant field is available as Technology Preview. These are early access APIs provided to test, automate and provide feedback on the feature. Since this can change based on feedback, VMware does not guarantee backwards compatibility and recommends against using them in production environments. Some Technology Preview APIs might only be applicable to specific environments.
	AnnouncementsSeverity_CRITICAL AnnouncementsSeverity = "CRITICAL"
    // Warning. **Warning:** This constant field is available as Technology Preview. These are early access APIs provided to test, automate and provide feedback on the feature. Since this can change based on feedback, VMware does not guarantee backwards compatibility and recommends against using them in production environments. Some Technology Preview APIs might only be applicable to specific environments.
	AnnouncementsSeverity_WARNING AnnouncementsSeverity = "WARNING"
    // Information. **Warning:** This constant field is available as Technology Preview. These are early access APIs provided to test, automate and provide feedback on the feature. Since this can change based on feedback, VMware does not guarantee backwards compatibility and recommends against using them in production environments. Some Technology Preview APIs might only be applicable to specific environments.
	AnnouncementsSeverity_INFO AnnouncementsSeverity = "INFO"
)

func (s AnnouncementsSeverity) AnnouncementsSeverity() bool {
	switch s {
	case AnnouncementsSeverity_CRITICAL:
		return true
	case AnnouncementsSeverity_WARNING:
		return true
	case AnnouncementsSeverity_INFO:
		return true
	default:
		return false
	}
}


// The ``Info`` class defines the announcement properties. **Warning:** This class is available as Technology Preview. These are early access APIs provided to test, automate and provide feedback on the feature. Since this can change based on feedback, VMware does not guarantee backwards compatibility and recommends against using them in production environments. Some Technology Preview APIs might only be applicable to specific environments.
type AnnouncementsInfo struct {
    // Message to display. **Warning:** This property is available as Technology Preview. These are early access APIs provided to test, automate and provide feedback on the feature. Since this can change based on feedback, VMware does not guarantee backwards compatibility and recommends against using them in production environments. Some Technology Preview APIs might only be applicable to specific environments.
	Message std.LocalizableMessage
    // Message severity. **Warning:** This property is available as Technology Preview. These are early access APIs provided to test, automate and provide feedback on the feature. Since this can change based on feedback, VMware does not guarantee backwards compatibility and recommends against using them in production environments. Some Technology Preview APIs might only be applicable to specific environments.
	Severity AnnouncementsSeverity
    // Date and time when the announcement will expire. **Warning:** This property is available as Technology Preview. These are early access APIs provided to test, automate and provide feedback on the feature. Since this can change based on feedback, VMware does not guarantee backwards compatibility and recommends against using them in production environments. Some Technology Preview APIs might only be applicable to specific environments.
	ExpiresAt time.Time
}

// The ``Spec`` class defines the announcement properties for set and create operations. **Warning:** This class is available as Technology Preview. These are early access APIs provided to test, automate and provide feedback on the feature. Since this can change based on feedback, VMware does not guarantee backwards compatibility and recommends against using them in production environments. Some Technology Preview APIs might only be applicable to specific environments.
type AnnouncementsSpec struct {
    // Message to display. **Warning:** This property is available as Technology Preview. These are early access APIs provided to test, automate and provide feedback on the feature. Since this can change based on feedback, VMware does not guarantee backwards compatibility and recommends against using them in production environments. Some Technology Preview APIs might only be applicable to specific environments.
	Message std.LocalizableMessage
    // Message severity. **Warning:** This property is available as Technology Preview. These are early access APIs provided to test, automate and provide feedback on the feature. Since this can change based on feedback, VMware does not guarantee backwards compatibility and recommends against using them in production environments. Some Technology Preview APIs might only be applicable to specific environments.
	Severity AnnouncementsSeverity
    // Date and time when the announcement will expire. **Warning:** This property is available as Technology Preview. These are early access APIs provided to test, automate and provide feedback on the feature. Since this can change based on feedback, VMware does not guarantee backwards compatibility and recommends against using them in production environments. Some Technology Preview APIs might only be applicable to specific environments.
	ExpiresAt time.Time
}



func announcementsListInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func announcementsListOutputType() bindings.BindingType {
	return bindings.NewMapType(bindings.NewIdType([]string{"com.vmware.vcenter.system.announcements"}, ""), bindings.NewReferenceType(AnnouncementsInfoBindingType),reflect.TypeOf(map[string]AnnouncementsInfo{}))
}

func announcementsListRestMetadata() protocol.OperationRestMetadata {
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
		map[string]int{"Error": 500})
}

func announcementsGetInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["announcement"] = bindings.NewIdType([]string{"com.vmware.vcenter.system.announcements"}, "")
	fieldNameMap["announcement"] = "Announcement"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func announcementsGetOutputType() bindings.BindingType {
	return bindings.NewReferenceType(AnnouncementsInfoBindingType)
}

func announcementsGetRestMetadata() protocol.OperationRestMetadata {
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
		map[string]int{"Error": 500})
}

func announcementsCreateInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["spec"] = bindings.NewReferenceType(AnnouncementsSpecBindingType)
	fieldNameMap["spec"] = "Spec"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func announcementsCreateOutputType() bindings.BindingType {
	return bindings.NewIdType([]string{"com.vmware.vcenter.system.announcements"}, "")
}

func announcementsCreateRestMetadata() protocol.OperationRestMetadata {
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
		map[string]int{"Error": 500})
}

func announcementsSetInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["announcement"] = bindings.NewIdType([]string{"com.vmware.vcenter.system.announcements"}, "")
	fields["spec"] = bindings.NewReferenceType(AnnouncementsSpecBindingType)
	fieldNameMap["announcement"] = "Announcement"
	fieldNameMap["spec"] = "Spec"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func announcementsSetOutputType() bindings.BindingType {
	return bindings.NewVoidType()
}

func announcementsSetRestMetadata() protocol.OperationRestMetadata {
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
		map[string]int{"Error": 500})
}

func announcementsDeleteInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["announcement"] = bindings.NewIdType([]string{"com.vmware.vcenter.system.announcements"}, "")
	fieldNameMap["announcement"] = "Announcement"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func announcementsDeleteOutputType() bindings.BindingType {
	return bindings.NewVoidType()
}

func announcementsDeleteRestMetadata() protocol.OperationRestMetadata {
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
		map[string]int{"NotFound": 404,"Error": 500})
}


func AnnouncementsInfoBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["message"] = bindings.NewReferenceType(std.LocalizableMessageBindingType)
	fieldNameMap["message"] = "Message"
	fields["severity"] = bindings.NewEnumType("com.vmware.vcenter.system.announcements.severity", reflect.TypeOf(AnnouncementsSeverity(AnnouncementsSeverity_CRITICAL)))
	fieldNameMap["severity"] = "Severity"
	fields["expires_at"] = bindings.NewDateTimeType()
	fieldNameMap["expires_at"] = "ExpiresAt"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vcenter.system.announcements.info", fields, reflect.TypeOf(AnnouncementsInfo{}), fieldNameMap, validators)
}

func AnnouncementsSpecBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["message"] = bindings.NewReferenceType(std.LocalizableMessageBindingType)
	fieldNameMap["message"] = "Message"
	fields["severity"] = bindings.NewEnumType("com.vmware.vcenter.system.announcements.severity", reflect.TypeOf(AnnouncementsSeverity(AnnouncementsSeverity_CRITICAL)))
	fieldNameMap["severity"] = "Severity"
	fields["expires_at"] = bindings.NewDateTimeType()
	fieldNameMap["expires_at"] = "ExpiresAt"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vcenter.system.announcements.spec", fields, reflect.TypeOf(AnnouncementsSpec{}), fieldNameMap, validators)
}

