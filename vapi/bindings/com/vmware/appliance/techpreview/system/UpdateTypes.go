/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Data type definitions file for service: Update.
 * Includes binding types of a structures and enumerations defined in the service.
 * Shared by client-side stubs and server-side skeletons to ensure type
 * compatibility.
 */
package system

import (
	"reflect"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/bindings"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/data"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/protocol"
)


// ``AutoUpdateNotification`` enumeration class Defines state for automatic update notification.
//
// <p> See {@link com.vmware.vapi.bindings.ApiEnumeration enumerated types description}.
type UpdateAutoUpdateNotification string

const (
    // Automatic update notification is disabled. Disable periodically query the configured url for updates.
	UpdateAutoUpdateNotification_disabled UpdateAutoUpdateNotification = "disabled"
    // Automatic update notification is enabled. Enable periodically query the configured url for updates.
	UpdateAutoUpdateNotification_enabled UpdateAutoUpdateNotification = "enabled"
)

func (a UpdateAutoUpdateNotification) UpdateAutoUpdateNotification() bool {
	switch a {
	case UpdateAutoUpdateNotification_disabled:
		return true
	case UpdateAutoUpdateNotification_enabled:
		return true
	default:
		return false
	}
}


// ``UpdateDay`` enumeration class Defines days to query for updates.
//
// <p> See {@link com.vmware.vapi.bindings.ApiEnumeration enumerated types description}.
type UpdateUpdateDay string

const (
    // query for updates on Monday
	UpdateUpdateDay_Monday UpdateUpdateDay = "Monday"
    // query for updates on Tuesday
	UpdateUpdateDay_Tuesday UpdateUpdateDay = "Tuesday"
    // query for updates on Friday
	UpdateUpdateDay_Friday UpdateUpdateDay = "Friday"
    // query for updates on Wednesday
	UpdateUpdateDay_Wednesday UpdateUpdateDay = "Wednesday"
    // query for updates on Thursday
	UpdateUpdateDay_Thursday UpdateUpdateDay = "Thursday"
    // query for updates on Saturday
	UpdateUpdateDay_Saturday UpdateUpdateDay = "Saturday"
    // query for updates on Sunday
	UpdateUpdateDay_Sunday UpdateUpdateDay = "Sunday"
    // query for updates everyday
	UpdateUpdateDay_Everyday UpdateUpdateDay = "Everyday"
)

func (u UpdateUpdateDay) UpdateUpdateDay() bool {
	switch u {
	case UpdateUpdateDay_Monday:
		return true
	case UpdateUpdateDay_Tuesday:
		return true
	case UpdateUpdateDay_Friday:
		return true
	case UpdateUpdateDay_Wednesday:
		return true
	case UpdateUpdateDay_Thursday:
		return true
	case UpdateUpdateDay_Saturday:
		return true
	case UpdateUpdateDay_Sunday:
		return true
	case UpdateUpdateDay_Everyday:
		return true
	default:
		return false
	}
}

// ``UpdateStructSet`` class Structure to set url update repository.
type UpdateUpdateStructSet struct {
    // Current appliance update repository URL. Enter "default" to reset the url to the default url.
	CurrentURL string
    // Check for update at the pre-configured repository URL.
	CheckUpdates UpdateAutoUpdateNotification
    // time to query for updates Format: HH:MM:SS Military (24 hour) Time Format
	Time string
    // day to query for updates
	Day UpdateUpdateDay
    // username for the url update repository
	Username string
    // password for the url update repository
	Password string
}

// ``UpdateStructGet`` class Structure to get url update repository.
type UpdateUpdateStructGet struct {
    // Current appliance update repository URL.
	CurrentURL string
    // Default appliance update repository URL.
	DefaultURL string
    // Check for update at the pre-configured repository URL.
	CheckUpdates UpdateAutoUpdateNotification
    // time to query for updates Format: HH:MM:SS Military (24 hour) Time Format
	Time string
    // day to query for updates
	Day UpdateUpdateDay
    // timestamp of latest update installation
	LatestUpdateInstallTime string
    // timestamp of latest query to update repository
	LatestUpdateQueryTime string
    // username for the url update repository
	Username string
    // password for the url update repository
	Password string
}


func updateSetInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["config"] = bindings.NewReferenceType(UpdateUpdateStructSetBindingType)
	fieldNameMap["config"] = "Config"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func updateSetOutputType() bindings.BindingType {
	return bindings.NewVoidType()
}

func updateSetRestMetadata() protocol.OperationRestMetadata {
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

func updateGetInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func updateGetOutputType() bindings.BindingType {
	return bindings.NewReferenceType(UpdateUpdateStructGetBindingType)
}

func updateGetRestMetadata() protocol.OperationRestMetadata {
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


func UpdateUpdateStructSetBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["current_URL"] = bindings.NewStringType()
	fieldNameMap["current_URL"] = "CurrentURL"
	fields["check_updates"] = bindings.NewEnumType("com.vmware.appliance.techpreview.system.update.auto_update_notification", reflect.TypeOf(UpdateAutoUpdateNotification(UpdateAutoUpdateNotification_disabled)))
	fieldNameMap["check_updates"] = "CheckUpdates"
	fields["time"] = bindings.NewStringType()
	fieldNameMap["time"] = "Time"
	fields["day"] = bindings.NewEnumType("com.vmware.appliance.techpreview.system.update.update_day", reflect.TypeOf(UpdateUpdateDay(UpdateUpdateDay_Monday)))
	fieldNameMap["day"] = "Day"
	fields["username"] = bindings.NewStringType()
	fieldNameMap["username"] = "Username"
	fields["password"] = bindings.NewSecretType()
	fieldNameMap["password"] = "Password"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.appliance.techpreview.system.update.update_struct_set", fields, reflect.TypeOf(UpdateUpdateStructSet{}), fieldNameMap, validators)
}

func UpdateUpdateStructGetBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["current_URL"] = bindings.NewStringType()
	fieldNameMap["current_URL"] = "CurrentURL"
	fields["default_URL"] = bindings.NewStringType()
	fieldNameMap["default_URL"] = "DefaultURL"
	fields["check_updates"] = bindings.NewEnumType("com.vmware.appliance.techpreview.system.update.auto_update_notification", reflect.TypeOf(UpdateAutoUpdateNotification(UpdateAutoUpdateNotification_disabled)))
	fieldNameMap["check_updates"] = "CheckUpdates"
	fields["time"] = bindings.NewStringType()
	fieldNameMap["time"] = "Time"
	fields["day"] = bindings.NewEnumType("com.vmware.appliance.techpreview.system.update.update_day", reflect.TypeOf(UpdateUpdateDay(UpdateUpdateDay_Monday)))
	fieldNameMap["day"] = "Day"
	fields["latest_update_install_time"] = bindings.NewStringType()
	fieldNameMap["latest_update_install_time"] = "LatestUpdateInstallTime"
	fields["latest_update_query_time"] = bindings.NewStringType()
	fieldNameMap["latest_update_query_time"] = "LatestUpdateQueryTime"
	fields["username"] = bindings.NewStringType()
	fieldNameMap["username"] = "Username"
	fields["password"] = bindings.NewStringType()
	fieldNameMap["password"] = "Password"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.appliance.techpreview.system.update.update_struct_get", fields, reflect.TypeOf(UpdateUpdateStructGet{}), fieldNameMap, validators)
}

