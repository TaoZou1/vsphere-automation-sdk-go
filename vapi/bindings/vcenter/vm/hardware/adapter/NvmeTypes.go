/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Code generated. DO NOT EDIT.

/*
 * Data type definitions file for service: Nvme.
 * Includes binding types of a structures and enumerations defined in the service.
 * Shared by client-side stubs and server-side skeletons to ensure type
 * compatibility.
 */

package adapter

import (
	"reflect"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/bindings"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/data"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/protocol"
)

// Resource type for the virtual NVMe adapter device. **Warning:** This constant field is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
const Nvme_RESOURCE_TYPE = "com.vmware.vcenter.vm.hardware.NvmeAdapter"


// The ``Info`` class contains information about a virtual NVMe adapter. **Warning:** This class is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
type NvmeInfo struct {
    // Device label. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
	Label string
    // NVMe bus number. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
	Bus int64
    // Address of the NVMe adapter on the PCI bus. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
	PciSlotNumber *int64
}

// The ``CreateSpec`` class provides a specification for the configuration of a newly-created virtual NVMe adapter. **Warning:** This class is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
type NvmeCreateSpec struct {
    // NVMe bus number. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
	Bus *int64
    // Address of the NVMe adapter on the PCI bus. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
	PciSlotNumber *int64
}

// The ``Summary`` class contains commonly used information about a Virtual NVMe adapter. **Warning:** This class is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
type NvmeSummary struct {
    // Identifier of the virtual NVMe adapter. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
	Adapter string
}



func nvmeListInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["vm"] = bindings.NewIdType([]string{"VirtualMachine"}, "")
	fieldNameMap["vm"] = "Vm"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func nvmeListOutputType() bindings.BindingType {
	return bindings.NewListType(bindings.NewReferenceType(NvmeSummaryBindingType), reflect.TypeOf([]NvmeSummary{}))
}

func nvmeListRestMetadata() protocol.OperationRestMetadata {
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
		map[string]int{"Error": 500,"NotFound": 404,"ResourceInaccessible": 500,"ServiceUnavailable": 503,"Unauthenticated": 401,"Unauthorized": 403})
}

func nvmeGetInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["vm"] = bindings.NewIdType([]string{"VirtualMachine"}, "")
	fields["adapter"] = bindings.NewIdType([]string{"com.vmware.vcenter.vm.hardware.NvmeAdapter"}, "")
	fieldNameMap["vm"] = "Vm"
	fieldNameMap["adapter"] = "Adapter"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func nvmeGetOutputType() bindings.BindingType {
	return bindings.NewReferenceType(NvmeInfoBindingType)
}

func nvmeGetRestMetadata() protocol.OperationRestMetadata {
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
		map[string]int{"Error": 500,"NotFound": 404,"ResourceInaccessible": 500,"ServiceUnavailable": 503,"Unauthenticated": 401,"Unauthorized": 403})
}

func nvmeCreateInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["vm"] = bindings.NewIdType([]string{"VirtualMachine"}, "")
	fields["spec"] = bindings.NewReferenceType(NvmeCreateSpecBindingType)
	fieldNameMap["vm"] = "Vm"
	fieldNameMap["spec"] = "Spec"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func nvmeCreateOutputType() bindings.BindingType {
	return bindings.NewIdType([]string{"com.vmware.vcenter.vm.hardware.NvmeAdapter"}, "")
}

func nvmeCreateRestMetadata() protocol.OperationRestMetadata {
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
		map[string]int{"Error": 500,"NotAllowedInCurrentState": 400,"NotFound": 404,"UnableToAllocateResource": 500,"ResourceInUse": 400,"InvalidArgument": 400,"ResourceBusy": 500,"ResourceInaccessible": 500,"ServiceUnavailable": 503,"Unauthenticated": 401,"Unauthorized": 403,"Unsupported": 400})
}

func nvmeDeleteInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["vm"] = bindings.NewIdType([]string{"VirtualMachine"}, "")
	fields["adapter"] = bindings.NewIdType([]string{"com.vmware.vcenter.vm.hardware.NvmeAdapter"}, "")
	fieldNameMap["vm"] = "Vm"
	fieldNameMap["adapter"] = "Adapter"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func nvmeDeleteOutputType() bindings.BindingType {
	return bindings.NewVoidType()
}

func nvmeDeleteRestMetadata() protocol.OperationRestMetadata {
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
		map[string]int{"Error": 500,"NotAllowedInCurrentState": 400,"NotFound": 404,"ResourceBusy": 500,"ResourceInaccessible": 500,"ServiceUnavailable": 503,"Unauthenticated": 401,"Unauthorized": 403})
}


func NvmeInfoBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["label"] = bindings.NewStringType()
	fieldNameMap["label"] = "Label"
	fields["bus"] = bindings.NewIntegerType()
	fieldNameMap["bus"] = "Bus"
	fields["pci_slot_number"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fieldNameMap["pci_slot_number"] = "PciSlotNumber"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vcenter.vm.hardware.adapter.nvme.info", fields, reflect.TypeOf(NvmeInfo{}), fieldNameMap, validators)
}

func NvmeCreateSpecBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["bus"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fieldNameMap["bus"] = "Bus"
	fields["pci_slot_number"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fieldNameMap["pci_slot_number"] = "PciSlotNumber"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vcenter.vm.hardware.adapter.nvme.create_spec", fields, reflect.TypeOf(NvmeCreateSpec{}), fieldNameMap, validators)
}

func NvmeSummaryBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["adapter"] = bindings.NewIdType([]string{"com.vmware.vcenter.vm.hardware.NvmeAdapter"}, "")
	fieldNameMap["adapter"] = "Adapter"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vcenter.vm.hardware.adapter.nvme.summary", fields, reflect.TypeOf(NvmeSummary{}), fieldNameMap, validators)
}

