/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Data type definitions file for service: Image.
 * Includes binding types of a structures and enumerations defined in the service.
 * Shared by client-side stubs and server-side skeletons to ensure type
 * compatibility.
 */
package iso

import (
	"reflect"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/bindings"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/data"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/protocol"
)


func imageMountInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["library_item"] = bindings.NewIdType([]string{"com.vmware.content.library.Item"}, "")
	fields["vm"] = bindings.NewIdType([]string{"VirtualMachine"}, "")
	fieldNameMap["library_item"] = "LibraryItem"
	fieldNameMap["vm"] = "Vm"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func imageMountOutputType() bindings.BindingType {
	return bindings.NewIdType([]string{"com.vmware.vcenter.vm.hardware.Cdrom"}, "")
}

func imageMountRestMetadata() protocol.OperationRestMetadata {
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
		map[string]int{"NotFound": 404,"InvalidArgument": 400,"NotAllowedInCurrentState": 400,"ResourceInaccessible": 500})
}

func imageUnmountInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["vm"] = bindings.NewIdType([]string{"VirtualMachine"}, "")
	fields["cdrom"] = bindings.NewIdType([]string{"com.vmware.vcenter.vm.hardware.Cdrom"}, "")
	fieldNameMap["vm"] = "Vm"
	fieldNameMap["cdrom"] = "Cdrom"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func imageUnmountOutputType() bindings.BindingType {
	return bindings.NewVoidType()
}

func imageUnmountRestMetadata() protocol.OperationRestMetadata {
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
		map[string]int{"NotFound": 404,"NotAllowedInCurrentState": 400})
}


