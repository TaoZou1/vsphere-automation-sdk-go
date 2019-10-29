/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Code generated. DO NOT EDIT.

/*
 * Data type definitions file for service: Routes.
 * Includes binding types of a structures and enumerations defined in the service.
 * Shared by client-side stubs and server-side skeletons to ensure type
 * compatibility.
 */

package networking

import (
	"reflect"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/bindings"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/data"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/protocol"
)



// The ``Info`` class describes an individual host, network or default destination network reachable through a gateway.
type RoutesInfo struct {
    // IP Address of the destination IP network. IPv4 address is specified using dotted decimal notation. For example, "192.0.2.1". IPv6 addresses are 128-bit specified using as eight fields of up to four hexadecimal digits. A colon separates each field (:). For example, 2001:DB8:101::230:6eff:fe04:d9ff. The address can also consist of symbol '::' to represent multiple 16-bit groups of contiguous 0's only once in an address as described in RFC 2373.
	Network string
    // The prefix length. For IPv4 the value range is 0-32. For IPv6 prefixLength is a decimal value range 0-128. The property represents the number of contiguous, higher-order bits of the address that make up the network portion of the IP address.
	PrefixLength int64
    // Where to send the packets for this route. Unicast IP Address of the next hop router. IPv4 address is specified using dotted decimal notation. For example, "192.0.2.1". IPv6 addresses are 128-bit specified using as eight fields of up to four hexadecimal digits. A colon separates each field (:). For example, 2001:DB8:101::230:6eff:fe04:d9ff. The address can also consist of symbol '::' to represent multiple 16-bit groups of contiguous 0's only once in an address as described in RFC 2373.
	GatewayAddress *string
    // The network interface associated with this route. This is an index into the result of Interfaces#list The index refers to the relative position of an element in a array. For example, an index of 0 refers to the first element in the array while an index of 1 refers to the second element.
	InterfaceIndex *int64
}



func routesListInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["vm"] = bindings.NewIdType([]string{"VirtualMachine"}, "")
	fieldNameMap["vm"] = "Vm"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func routesListOutputType() bindings.BindingType {
	return bindings.NewListType(bindings.NewReferenceType(RoutesInfoBindingType), reflect.TypeOf([]RoutesInfo{}))
}

func routesListRestMetadata() protocol.OperationRestMetadata {
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
		"/vcenter/vm/{vm}/guest/networking/routes",
		resultHeaders,
		200,
		errorHeaders,
		map[string]int{"Error": 500,"NotFound": 404,"ServiceUnavailable": 503})
}


func RoutesInfoBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["network"] = bindings.NewStringType()
	fieldNameMap["network"] = "Network"
	fields["prefix_length"] = bindings.NewIntegerType()
	fieldNameMap["prefix_length"] = "PrefixLength"
	fields["gateway_address"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["gateway_address"] = "GatewayAddress"
	fields["interface_index"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fieldNameMap["interface_index"] = "InterfaceIndex"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vcenter.vm.guest.networking.routes.info", fields, reflect.TypeOf(RoutesInfo{}), fieldNameMap, validators)
}

