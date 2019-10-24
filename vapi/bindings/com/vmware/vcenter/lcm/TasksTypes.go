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
package lcm

import (
	"reflect"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/bindings"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/data"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/protocol"
)



// The type of operation to be executed. **Warning:** This enumeration is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
//
// <p> See {@link com.vmware.vapi.bindings.ApiEnumeration enumerated types description}.
type TasksOperation string

const (
    // The install deployment operation type. **Warning:** This constant field is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
	TasksOperation_INSTALL_DEPLOY TasksOperation = "INSTALL_DEPLOY"
    // The install precheck operation type. **Warning:** This constant field is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
	TasksOperation_INSTALL_PRECHECK TasksOperation = "INSTALL_PRECHECK"
    // The upgrade deployment operation type. **Warning:** This constant field is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
	TasksOperation_UPGRADE_DEPLOY TasksOperation = "UPGRADE_DEPLOY"
    // The upgrade precheck operation type. **Warning:** This constant field is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
	TasksOperation_UPGRADE_PRECHECK TasksOperation = "UPGRADE_PRECHECK"
)

func (o TasksOperation) TasksOperation() bool {
	switch o {
	case TasksOperation_INSTALL_DEPLOY:
		return true
	case TasksOperation_INSTALL_PRECHECK:
		return true
	case TasksOperation_UPGRADE_DEPLOY:
		return true
	case TasksOperation_UPGRADE_PRECHECK:
		return true
	default:
		return false
	}
}


// The data container of the task. **Warning:** This class is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
type TasksInfo struct {
    // Identifier of the LCM task. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
	Id string
    // The identifier of the CIS task. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
	Task *string
    // The type of the task . **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
	Type_ TasksOperation
    // The status of the task. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
	Status TaskInfo
}



func tasksDeleteInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["task"] = bindings.NewIdType([]string{"com.vmware.vcenter.lcm.task"}, "")
	fieldNameMap["task"] = "Task"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func tasksDeleteOutputType() bindings.BindingType {
	return bindings.NewVoidType()
}

func tasksDeleteRestMetadata() protocol.OperationRestMetadata {
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
		map[string]int{"NotFound": 404,"InvalidArgument": 400})
}


func TasksInfoBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["id"] = bindings.NewIdType([]string{"com.vmware.vcenter.lcm.task"}, "")
	fieldNameMap["id"] = "Id"
	fields["task"] = bindings.NewOptionalType(bindings.NewIdType([]string{"com.vmware.cis.task"}, ""))
	fieldNameMap["task"] = "Task"
	fields["type"] = bindings.NewEnumType("com.vmware.vcenter.lcm.tasks.operation", reflect.TypeOf(TasksOperation(TasksOperation_INSTALL_DEPLOY)))
	fieldNameMap["type"] = "Type_"
	fields["status"] = bindings.NewReferenceType(TaskInfoBindingType)
	fieldNameMap["status"] = "Status"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vcenter.lcm.tasks.info", fields, reflect.TypeOf(TasksInfo{}), fieldNameMap, validators)
}

