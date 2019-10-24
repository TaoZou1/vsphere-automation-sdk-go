/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Data type definitions file for service: Operations.
 * Includes binding types of a structures and enumerations defined in the service.
 * Shared by client-side stubs and server-side skeletons to ensure type
 * compatibility.
 */
package vcha

import (
	"reflect"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/bindings"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/data"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/protocol"
)

// Identifier of the cluster deploy operation.
const Operations_CLUSTER_DEPLOY_OP = "com.vmware.vcenter.vcha.cluster.deploy"
// Identifier of the cluster failover operation.
const Operations_CLUSTER_FAILOVER_OP = "com.vmware.vcenter.vcha.cluster.failover"
// Identifier of the cluster get operation.
const Operations_CLUSTER_GET_OP = "com.vmware.vcenter.vcha.cluster.get"
// Identifier of the passive redeploy operation.
const Operations_PASSIVE_REDEPLOY_OP = "com.vmware.vcenter.vcha.cluster.passive.redeploy"
// Identifier of the witness redeploy operation.
const Operations_WITNESS_REDEPLOY_OP = "com.vmware.vcenter.vcha.cluster.witness.redeploy"
// Identifier of the get mode operation.
const Operations_MODE_GET_OP = "com.vmware.vcenter.vcha.cluster.mode.get"
// Identifier of the set mode operation.
const Operations_MODE_SET_OP = "com.vmware.vcenter.vcha.cluster.mode.set"
// Identifier of the cluster undeploy operation.
const Operations_CLUSTER_UNDEPLOY_OP = "com.vmware.vcenter.vcha.cluster.undeploy"


// The ``Info`` class contains information about which the VCHA operations cannot be invoked in the current state of the system and the list of currently active VCHA operations.
type OperationsInfo struct {
    // Identifiers of the operations that are current disabled. These operation strings are one of Operations#Operations_CLUSTER_DEPLOY_OP, Operations#Operations_CLUSTER_FAILOVER_OP, Operations#Operations_PASSIVE_REDEPLOY_OP, Operations#Operations_WITNESS_REDEPLOY_OP, Operations#Operations_MODE_SET_OP, Operations#Operations_CLUSTER_UNDEPLOY_OP and Operations#Operations_CLUSTER_GET_OP.
	Disabled map[string]bool
    // Identifiers of the operations that are currently running. These operation strings are one of Operations#Operations_CLUSTER_DEPLOY_OP, Operations#Operations_CLUSTER_FAILOVER_OP, Operations#Operations_PASSIVE_REDEPLOY_OP, Operations#Operations_WITNESS_REDEPLOY_OP, Operations#Operations_MODE_SET_OP, and Operations#Operations_CLUSTER_UNDEPLOY_OP.
	Active map[string]bool
}



func operationsGetInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func operationsGetOutputType() bindings.BindingType {
	return bindings.NewReferenceType(OperationsInfoBindingType)
}

func operationsGetRestMetadata() protocol.OperationRestMetadata {
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
		map[string]int{"Unauthorized": 403,"Error": 500})
}


func OperationsInfoBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["disabled"] = bindings.NewSetType(bindings.NewIdType([]string{"com.vmware.vapi.operation"}, ""), reflect.TypeOf(map[string]bool{}))
	fieldNameMap["disabled"] = "Disabled"
	fields["active"] = bindings.NewSetType(bindings.NewIdType([]string{"com.vmware.vapi.operation"}, ""), reflect.TypeOf(map[string]bool{}))
	fieldNameMap["active"] = "Active"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vcenter.vcha.operations.info", fields, reflect.TypeOf(OperationsInfo{}), fieldNameMap, validators)
}

