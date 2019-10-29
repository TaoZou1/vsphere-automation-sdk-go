/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Code generated. DO NOT EDIT.

/*
 * Data type definitions file for service: Identity.
 * Includes binding types of a structures and enumerations defined in the service.
 * Shared by client-side stubs and server-side skeletons to ensure type
 * compatibility.
 */

package guest

import (
	"reflect"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/bindings/vapi/std"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/bindings/vcenter/vm"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/bindings"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/data"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/protocol"
)



// The ``Info`` class contains information describing the guest operating system identification.
type IdentityInfo struct {
    // Guest operating system identifier (short name).
	Name vm.GuestOS
    // Guest operating system family.
	Family vm.GuestOSFamily
    // Guest operating system full name.
	FullName std.LocalizableMessage
    // Hostname of the guest operating system.
	HostName string
    // IP address assigned by the guest operating system.
	IpAddress *string
}



func identityGetInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["vm"] = bindings.NewIdType([]string{"VirtualMachine"}, "")
	fieldNameMap["vm"] = "Vm"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func identityGetOutputType() bindings.BindingType {
	return bindings.NewReferenceType(IdentityInfoBindingType)
}

func identityGetRestMetadata() protocol.OperationRestMetadata {
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	paramsTypeMap["vm"] = bindings.NewIdType([]string{"VirtualMachine"}, "")
	paramsTypeMap["vm"] = bindings.NewIdType([]string{"VirtualMachine"}, "")
	pathParams["vm"] = "vm"
	resultHeaders := map[string]string{}
	errorHeaders := map[string]string{}
	return protocol.NewOperationRestMetadata(
		paramsTypeMap,
		pathParams,
		queryParams,
		headerParams,
		"",
		"GET",
		"/vcenter/vm/{vm}/guest/identity",
		resultHeaders,
		200,
		errorHeaders,
		map[string]int{"Error": 500,"NotFound": 404,"ServiceUnavailable": 503})
}


func IdentityInfoBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["name"] = bindings.NewEnumType("com.vmware.vcenter.vm.guest_OS", reflect.TypeOf(vm.GuestOS(vm.GuestOS_DOS)))
	fieldNameMap["name"] = "Name"
	fields["family"] = bindings.NewEnumType("com.vmware.vcenter.vm.guest_OS_family", reflect.TypeOf(vm.GuestOSFamily(vm.GuestOSFamily_WINDOWS)))
	fieldNameMap["family"] = "Family"
	fields["full_name"] = bindings.NewReferenceType(std.LocalizableMessageBindingType)
	fieldNameMap["full_name"] = "FullName"
	fields["host_name"] = bindings.NewStringType()
	fieldNameMap["host_name"] = "HostName"
	fields["ip_address"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["ip_address"] = "IpAddress"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vcenter.vm.guest.identity.info", fields, reflect.TypeOf(IdentityInfo{}), fieldNameMap, validators)
}

