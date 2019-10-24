/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Data type definitions file for service: RetentionPlans.
 * Includes binding types of a structures and enumerations defined in the service.
 * Shared by client-side stubs and server-side skeletons to ensure type
 * compatibility.
 */
package vstats

import (
	"reflect"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/bindings"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/data"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/protocol"
)

// The ``MaxAgesInfo`` class specifies retention times for roll up intervals. The intervals are specified in seconds.
type RetentionPlansMaxAgesInfo struct {
    // Retention time for raw data. Default is one week. Unit is seconds.
	Raw int64
    // Retention time for data rolled up into 1-minute intervals. Unit is seconds.
	Minute1 int64
    // Retention time for data rolled up into 1-hour intervals. Unit is seconds.
	Hour1 int64
    // Retention time for data rolled up into 1-week intervals. Unit is seconds.
	Week1 int64
}

// ``Info`` class describes a data retention plan.
type RetentionPlansInfo struct {
    // Retention plan identifier.
	Id string
    // Maximum retention times for the different roll up levels.
	MaxAges RetentionPlansMaxAgesInfo
    // Consumer provided text about this Retention Plan.
	Memo_ *string
}


func retentionPlansGetDefaultInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func retentionPlansGetDefaultOutputType() bindings.BindingType {
	return bindings.NewReferenceType(RetentionPlansInfoBindingType)
}

func retentionPlansGetDefaultRestMetadata() protocol.OperationRestMetadata {
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
		"/stats/retentions/default",
		resultHeaders,
		200,
		errorHeaders,
		map[string]int{"Error": 500,"NotFound": 404,"Unauthenticated": 401,"Unauthorized": 403})
}


func RetentionPlansMaxAgesInfoBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["raw"] = bindings.NewIntegerType()
	fieldNameMap["raw"] = "Raw"
	fields["minute1"] = bindings.NewIntegerType()
	fieldNameMap["minute1"] = "Minute1"
	fields["hour1"] = bindings.NewIntegerType()
	fieldNameMap["hour1"] = "Hour1"
	fields["week1"] = bindings.NewIntegerType()
	fieldNameMap["week1"] = "Week1"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vstats.retention_plans.max_ages_info", fields, reflect.TypeOf(RetentionPlansMaxAgesInfo{}), fieldNameMap, validators)
}

func RetentionPlansInfoBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["id"] = bindings.NewIdType([]string{"com.vmware.vstats.model.RetentionPlan"}, "")
	fieldNameMap["id"] = "Id"
	fields["max_ages"] = bindings.NewReferenceType(RetentionPlansMaxAgesInfoBindingType)
	fieldNameMap["max_ages"] = "MaxAges"
	fields["memo_"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["memo_"] = "Memo_"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vstats.retention_plans.info", fields, reflect.TypeOf(RetentionPlansInfo{}), fieldNameMap, validators)
}

