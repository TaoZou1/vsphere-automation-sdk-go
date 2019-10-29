/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Code generated. DO NOT EDIT.

/*
 * Data type definitions file for service: Recovery.
 * Includes binding types of a structures and enumerations defined in the service.
 * Shared by client-side stubs and server-side skeletons to ensure type
 * compatibility.
 */

package nsx

import (
	"reflect"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/bindings/vapi/std"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/bindings"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/data"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/protocol"
	"time"
)



// The ``LossType`` enumeration class defines the type of loss to recover from. **Warning:** This enumeration is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
//
// <p> See {@link com.vmware.vapi.bindings.ApiEnumeration enumerated types description}.
type RecoveryLossType string

const (
    // Indicates loss of storage and servers. **Warning:** This constant field is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
	RecoveryLossType_ALL RecoveryLossType = "ALL"
)

func (l RecoveryLossType) RecoveryLossType() bool {
	switch l {
	case RecoveryLossType_ALL:
		return true
	default:
		return false
	}
}


// The ``RecoveryStage`` enumeration class defines the stage of the NSX infrastructure recovery operation. **Warning:** This enumeration is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
//
// <p> See {@link com.vmware.vapi.bindings.ApiEnumeration enumerated types description}.
type RecoveryRecoveryStage string

const (
    // Recovery operation is not in progress. **Warning:** This constant field is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
	RecoveryRecoveryStage_NONE RecoveryRecoveryStage = "NONE"
    // Recovery operation is resetting the service account password. **Warning:** This constant field is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
	RecoveryRecoveryStage_SERVICE_ACCOUNT_PASSWD_RESET RecoveryRecoveryStage = "SERVICE_ACCOUNT_PASSWD_RESET"
    // Recovery operation is locating the NSX-I VM. **Warning:** This constant field is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
	RecoveryRecoveryStage_NSX_LOCATE RecoveryRecoveryStage = "NSX_LOCATE"
    // Recovery operation is creating NSX infrastructure for WCP clusters. **Warning:** This constant field is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
	RecoveryRecoveryStage_NSX_INSTALL RecoveryRecoveryStage = "NSX_INSTALL"
    // Recovery operation is complete. **Warning:** This constant field is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
	RecoveryRecoveryStage_DONE RecoveryRecoveryStage = "DONE"
)

func (r RecoveryRecoveryStage) RecoveryRecoveryStage() bool {
	switch r {
	case RecoveryRecoveryStage_NONE:
		return true
	case RecoveryRecoveryStage_SERVICE_ACCOUNT_PASSWD_RESET:
		return true
	case RecoveryRecoveryStage_NSX_LOCATE:
		return true
	case RecoveryRecoveryStage_NSX_INSTALL:
		return true
	case RecoveryRecoveryStage_DONE:
		return true
	default:
		return false
	}
}


// The ``Info`` class contains information related to recovery needed. **Warning:** This class is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
type RecoveryInfo struct {
    // Indicates whether recovery needs to be performed. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
	Needed bool
    // Indicates the type of loss that needs recovery. If recovery is needed, then loss is present. If recovery is not needed, then loss is not present. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
	Loss *RecoveryLossType
}

// The ``ExecuteSpec`` class contains information related to recovery of NSX infrastructure for WCP clusters. **Warning:** This class is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
type RecoveryExecuteSpec struct {
    // Password of the NSX-I VM and NSX Edge nodes that would be created. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
	RootPassword string
}

// The ``ExecutionStatus`` class contains information related to the status of recovery of NSX infrastructure for WCP clusters. **Warning:** This class is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
type RecoveryExecutionStatus struct {
    // Stage of the recovery operation. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
	Stage RecoveryRecoveryStage
    // Error that occurred during the recovery operation. This error does not represent any error that occurs while recovering the NSX infrastructure. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
	Errors []std.LocalizableMessage
    // Start time of the recovery operation. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
	StartTime time.Time
    // End time of the recovery operation. It can be unspecified if the recovery operation is in progress. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
	EndTime time.Time
}



func recoveryGetInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func recoveryGetOutputType() bindings.BindingType {
	return bindings.NewReferenceType(RecoveryInfoBindingType)
}

func recoveryGetRestMetadata() protocol.OperationRestMetadata {
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
		"POST",
		"/vcenter/nsx/recovery",
		resultHeaders,
		200,
		errorHeaders,
		map[string]int{"Unauthorized": 403,"Unauthenticated": 401,"InternalServerError": 500,"Error": 500})
}

func recoveryExecuteInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["spec"] = bindings.NewReferenceType(RecoveryExecuteSpecBindingType)
	fieldNameMap["spec"] = "Spec"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func recoveryExecuteOutputType() bindings.BindingType {
	return bindings.NewIdType([]string{"Task"}, "")
}

func recoveryExecuteRestMetadata() protocol.OperationRestMetadata {
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	paramsTypeMap["spec"] = bindings.NewReferenceType(RecoveryExecuteSpecBindingType)
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
		"/vcenter/nsx/recovery",
		resultHeaders,
		200,
		errorHeaders,
		map[string]int{"Unauthenticated": 401,"Unauthorized": 403,"NotAllowedInCurrentState": 400,"AlreadyInDesiredState": 400,"Unsupported": 400,"InternalServerError": 500,"Error": 500})
}

func recoveryExecuteStatusInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func recoveryExecuteStatusOutputType() bindings.BindingType {
	return bindings.NewReferenceType(RecoveryExecutionStatusBindingType)
}

func recoveryExecuteStatusRestMetadata() protocol.OperationRestMetadata {
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
		"/vcenter/nsx/recovery/execute/status",
		resultHeaders,
		200,
		errorHeaders,
		map[string]int{"Unauthenticated": 401,"Unauthorized": 403,"Error": 500})
}


func RecoveryInfoBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["needed"] = bindings.NewBooleanType()
	fieldNameMap["needed"] = "Needed"
	fields["loss"] = bindings.NewOptionalType(bindings.NewEnumType("com.vmware.vcenter.nsx.recovery.loss_type", reflect.TypeOf(RecoveryLossType(RecoveryLossType_ALL))))
	fieldNameMap["loss"] = "Loss"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vcenter.nsx.recovery.info", fields, reflect.TypeOf(RecoveryInfo{}), fieldNameMap, validators)
}

func RecoveryExecuteSpecBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["root_password"] = bindings.NewSecretType()
	fieldNameMap["root_password"] = "RootPassword"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vcenter.nsx.recovery.execute_spec", fields, reflect.TypeOf(RecoveryExecuteSpec{}), fieldNameMap, validators)
}

func RecoveryExecutionStatusBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["stage"] = bindings.NewEnumType("com.vmware.vcenter.nsx.recovery.recovery_stage", reflect.TypeOf(RecoveryRecoveryStage(RecoveryRecoveryStage_NONE)))
	fieldNameMap["stage"] = "Stage"
	fields["errors"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewReferenceType(std.LocalizableMessageBindingType), reflect.TypeOf([]std.LocalizableMessage{})))
	fieldNameMap["errors"] = "Errors"
	fields["start_time"] = bindings.NewDateTimeType()
	fieldNameMap["start_time"] = "StartTime"
	fields["end_time"] = bindings.NewDateTimeType()
	fieldNameMap["end_time"] = "EndTime"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vcenter.nsx.recovery.execution_status", fields, reflect.TypeOf(RecoveryExecutionStatus{}), fieldNameMap, validators)
}

