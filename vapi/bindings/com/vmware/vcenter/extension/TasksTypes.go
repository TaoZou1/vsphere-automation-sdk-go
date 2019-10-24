/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Data type definitions file for service: Tasks.
 * Includes binding types of a structures and enumerations defined in the service.
 * Shared by client-side stubs and server-side skeletons to ensure type
 * compatibility.
 */
package extension

import (
	"reflect"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/bindings/com/vmware/cis/task"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/bindings/com/vmware/vapi/std"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/bindings"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/data"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/protocol"
	"time"
)

// Resource type for task. **Warning:** This constant field is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
const Tasks_RESOURCE_TYPE = "com.vmware.cis.task"
// The ``CreateSpec`` class contains properties used to create a task. **Warning:** This class is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
type TasksCreateSpec struct {
    // Target with which Task will be associated. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
	Target *std.DynamicID
    // Extension registered operation identifier associated the the task. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
	OperationId string
    // The percentage completion of the operation associated with the task. In the current release, the 'LocalizableMessage message' field of Info.progress will be ignored. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
	PercentDone *task.Progress
    // The description of the operation associated with the task. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
	Description *std.LocalizableMessage
    // The status of the operation associated with the task. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
	Status *task.Status
    // The name of the user who performed the operation in the Extension. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
	User *string
    // Flag to indicate whether or not the operation can be canceled. The value may change as the operation progresses. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
	Cancelable bool
    // Parent of the current task. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
	ParentTask *string
    // Description of the error if the operation status is "FAILED". **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
	Error_ *data.ErrorValue
    // Result of the operation. If an operation reports partial results before it completes, this property could be map with bool value before the status has the value "SUCCEEDED". The value could change as the operation progresses. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
	Result data.DataValue
    // Time when the operation is started. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
	StartTime *time.Time
    // Time when the operation is completed. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
	EndTime *time.Time
}

// The ``UpdateSpec`` class contains properties used to update the information of the operation associated with a task. **Warning:** This class is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
type TasksUpdateSpec struct {
    // The percentage completion of the operation associated with the task. In the current release, the 'LocalizableMessage message' field of Info.progress will be ignored. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
	PercentDone *task.Progress
    // The description of the operation associated with the task. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
	Description *std.LocalizableMessage
    // The status of a task on the Targets. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
	Status *task.Status
    // Description of the error if the operation status is "FAILED". **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
	Error_ *data.ErrorValue
    // Result of the operation. If an operation reports partial results before it completes, this property could be map with bool value before the status has the value "SUCCEEDED". The value could change as the operation progresses. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
	Result data.DataValue
    // Time when the operation is started. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
	StartTime *time.Time
    // Time when the operation is completed. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
	EndTime *time.Time
}


func tasksCreateInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["spec"] = bindings.NewReferenceType(TasksCreateSpecBindingType)
	fieldNameMap["spec"] = "Spec"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func tasksCreateOutputType() bindings.BindingType {
	return bindings.NewIdType([]string{"com.vmware.cis.task"}, "")
}

func tasksCreateRestMetadata() protocol.OperationRestMetadata {
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
		map[string]int{"Error": 500,"NotAllowedInCurrentState": 400,"InvalidArgument": 400})
}

func tasksUpdateInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["task_id"] = bindings.NewIdType([]string{"com.vmware.cis.task"}, "")
	fields["spec"] = bindings.NewReferenceType(TasksUpdateSpecBindingType)
	fieldNameMap["task_id"] = "TaskId"
	fieldNameMap["spec"] = "Spec"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func tasksUpdateOutputType() bindings.BindingType {
	return bindings.NewVoidType()
}

func tasksUpdateRestMetadata() protocol.OperationRestMetadata {
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
		map[string]int{"Error": 500,"NotAllowedInCurrentState": 400,"InvalidArgument": 400})
}


func TasksCreateSpecBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["target"] = bindings.NewOptionalType(bindings.NewReferenceType(std.DynamicIDBindingType))
	fieldNameMap["target"] = "Target"
	fields["operation_id"] = bindings.NewIdType([]string{"com.vmware.cis.task"}, "")
	fieldNameMap["operation_id"] = "OperationId"
	fields["percent_done"] = bindings.NewOptionalType(bindings.NewReferenceType(task.ProgressBindingType))
	fieldNameMap["percent_done"] = "PercentDone"
	fields["description"] = bindings.NewOptionalType(bindings.NewReferenceType(std.LocalizableMessageBindingType))
	fieldNameMap["description"] = "Description"
	fields["status"] = bindings.NewOptionalType(bindings.NewEnumType("com.vmware.cis.task.status", reflect.TypeOf(task.Status(task.Status_PENDING))))
	fieldNameMap["status"] = "Status"
	fields["user"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["user"] = "User"
	fields["cancelable"] = bindings.NewBooleanType()
	fieldNameMap["cancelable"] = "Cancelable"
	fields["parent_task"] = bindings.NewOptionalType(bindings.NewIdType([]string{"com.vmware.cis.task"}, ""))
	fieldNameMap["parent_task"] = "ParentTask"
	fields["error"] = bindings.NewOptionalType(bindings.NewAnyErrorType())
	fieldNameMap["error"] = "Error_"
	fields["result"] = bindings.NewOptionalType(bindings.NewOpaqueType())
	fieldNameMap["result"] = "Result"
	fields["start_time"] = bindings.NewOptionalType(bindings.NewDateTimeType())
	fieldNameMap["start_time"] = "StartTime"
	fields["end_time"] = bindings.NewOptionalType(bindings.NewDateTimeType())
	fieldNameMap["end_time"] = "EndTime"
	var validators = []bindings.Validator{}
	uv1 := bindings.NewUnionValidator("status",
	    map[string][]bindings.FieldData{
	        "FAILED": []bindings.FieldData{
	             bindings.NewFieldData("error", false),
	             bindings.NewFieldData("result", false),
	             bindings.NewFieldData("start_time", false),
	             bindings.NewFieldData("end_time", false),
	        },
	        "RUNNING": []bindings.FieldData{
	             bindings.NewFieldData("result", false),
	             bindings.NewFieldData("start_time", false),
	        },
	        "BLOCKED": []bindings.FieldData{
	             bindings.NewFieldData("result", false),
	             bindings.NewFieldData("start_time", false),
	        },
	        "SUCCEEDED": []bindings.FieldData{
	             bindings.NewFieldData("result", false),
	             bindings.NewFieldData("start_time", false),
	             bindings.NewFieldData("end_time", false),
	        },
	        "PENDING": []bindings.FieldData{},
	    },
	)
	validators = append(validators, uv1)
	return bindings.NewStructType("com.vmware.vcenter.extension.tasks.create_spec", fields, reflect.TypeOf(TasksCreateSpec{}), fieldNameMap, validators)
}

func TasksUpdateSpecBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["percent_done"] = bindings.NewOptionalType(bindings.NewReferenceType(task.ProgressBindingType))
	fieldNameMap["percent_done"] = "PercentDone"
	fields["description"] = bindings.NewOptionalType(bindings.NewReferenceType(std.LocalizableMessageBindingType))
	fieldNameMap["description"] = "Description"
	fields["status"] = bindings.NewOptionalType(bindings.NewEnumType("com.vmware.cis.task.status", reflect.TypeOf(task.Status(task.Status_PENDING))))
	fieldNameMap["status"] = "Status"
	fields["error"] = bindings.NewOptionalType(bindings.NewAnyErrorType())
	fieldNameMap["error"] = "Error_"
	fields["result"] = bindings.NewOptionalType(bindings.NewOpaqueType())
	fieldNameMap["result"] = "Result"
	fields["start_time"] = bindings.NewOptionalType(bindings.NewDateTimeType())
	fieldNameMap["start_time"] = "StartTime"
	fields["end_time"] = bindings.NewOptionalType(bindings.NewDateTimeType())
	fieldNameMap["end_time"] = "EndTime"
	var validators = []bindings.Validator{}
	uv1 := bindings.NewUnionValidator("status",
	    map[string][]bindings.FieldData{
	        "FAILED": []bindings.FieldData{
	             bindings.NewFieldData("error", false),
	             bindings.NewFieldData("result", false),
	             bindings.NewFieldData("start_time", false),
	             bindings.NewFieldData("end_time", false),
	        },
	        "RUNNING": []bindings.FieldData{
	             bindings.NewFieldData("result", false),
	             bindings.NewFieldData("start_time", false),
	        },
	        "BLOCKED": []bindings.FieldData{
	             bindings.NewFieldData("result", false),
	             bindings.NewFieldData("start_time", false),
	        },
	        "SUCCEEDED": []bindings.FieldData{
	             bindings.NewFieldData("result", false),
	             bindings.NewFieldData("start_time", false),
	             bindings.NewFieldData("end_time", false),
	        },
	        "PENDING": []bindings.FieldData{},
	    },
	)
	validators = append(validators, uv1)
	return bindings.NewStructType("com.vmware.vcenter.extension.tasks.update_spec", fields, reflect.TypeOf(TasksUpdateSpec{}), fieldNameMap, validators)
}

