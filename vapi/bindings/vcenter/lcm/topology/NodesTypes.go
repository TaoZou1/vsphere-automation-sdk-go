/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Code generated. DO NOT EDIT.

/*
 * Data type definitions file for service: Nodes.
 * Includes binding types of a structures and enumerations defined in the service.
 * Shared by client-side stubs and server-side skeletons to ensure type
 * compatibility.
 */

package topology

import (
	"reflect"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/bindings"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/data"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/protocol"
)



// The ``VcNode`` class contains vCenter or Platform Services Controller node details.
type NodesVcNode struct {
    // Address of a particular VC node in the federation.
	Address string
    // Version of a given VC node in the federation.
	Version string
}



func nodesListInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func nodesListOutputType() bindings.BindingType {
	return bindings.NewListType(bindings.NewReferenceType(NodesVcNodeBindingType), reflect.TypeOf([]NodesVcNode{}))
}

func nodesListRestMetadata() protocol.OperationRestMetadata {
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
		"/vcenter/lcm/topology/nodes",
		resultHeaders,
		200,
		errorHeaders,
		map[string]int{"Unauthenticated": 401,"Error": 500})
}


func NodesVcNodeBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["address"] = bindings.NewStringType()
	fieldNameMap["address"] = "Address"
	fields["version"] = bindings.NewStringType()
	fieldNameMap["version"] = "Version"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vcenter.lcm.topology.nodes.vc_node", fields, reflect.TypeOf(NodesVcNode{}), fieldNameMap, validators)
}

