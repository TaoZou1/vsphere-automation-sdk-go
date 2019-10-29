/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Code generated. DO NOT EDIT.

/*
 * Data type definitions file for service: Convergence.
 * Includes binding types of a structures and enumerations defined in the service.
 * Shared by client-side stubs and server-side skeletons to ensure type
 * compatibility.
 */

package hlm

import (
	"reflect"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/bindings"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/data"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/protocol"
)



// The ``Status`` is the convergence status of a node in a community. Or the community status, which will pick the high priority status from all nodes. Priority: ERROR > NETWORK_ERROR > NOT_CONVERGED > RUNNING > CONVERGED. **Warning:** This enumeration is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
//
// <p> See {@link com.vmware.vapi.bindings.ApiEnumeration enumerated types description}.
type ConvergenceStatus string

const (
	ConvergenceStatus_CONVERGED ConvergenceStatus = "CONVERGED"
	ConvergenceStatus_RUNNING ConvergenceStatus = "RUNNING"
	ConvergenceStatus_NOT_CONVERGED ConvergenceStatus = "NOT_CONVERGED"
	ConvergenceStatus_NETWORK_ERROR ConvergenceStatus = "NETWORK_ERROR"
	ConvergenceStatus_ERROR ConvergenceStatus = "ERROR"
)

func (s ConvergenceStatus) ConvergenceStatus() bool {
	switch s {
	case ConvergenceStatus_CONVERGED:
		return true
	case ConvergenceStatus_RUNNING:
		return true
	case ConvergenceStatus_NOT_CONVERGED:
		return true
	case ConvergenceStatus_NETWORK_ERROR:
		return true
	case ConvergenceStatus_ERROR:
		return true
	default:
		return false
	}
}


// The ``ServiceConvergenceInfo`` class contains properties the information that are used to describe if data of the service has converged for each syncable service. Converged means all data has been replicated to other partners in the community. **Warning:** This class is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
type ConvergenceServiceConvergenceInfo struct {
    // The result of convergence check for community. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
	ServiceStatus ConvergenceStatus
    // The ``nodeConvergenceStatus`` is a lang.Map collection that key is a null hostname of the replica and value is ConvergenceStatus convergence status of the host. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
	NodeConvergenceStatus map[string]ConvergenceStatus
}

// The ``ConvergenceInfo`` class contains properties that are used to describe the outcome of convergence check. **Warning:** This class is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
type ConvergenceConvergenceInfo struct {
    // The result of convergence check for community. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
	CommunityStatus ConvergenceStatus
    // The detail information for each service, the information includes the current status of the service, and whether it has been converged for each node in the community. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
	ServiceStatus map[string]ConvergenceServiceConvergenceInfo
}



func convergenceConvergeInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func convergenceConvergeOutputType() bindings.BindingType {
	return bindings.NewReferenceType(ConvergenceConvergenceInfoBindingType)
}

func convergenceConvergeRestMetadata() protocol.OperationRestMetadata {
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
		map[string]int{"Unauthenticated": 401,"Unauthorized": 403,"Error": 500})
}


func ConvergenceServiceConvergenceInfoBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["service_status"] = bindings.NewEnumType("com.vmware.vcenter.hlm.convergence.status", reflect.TypeOf(ConvergenceStatus(ConvergenceStatus_CONVERGED)))
	fieldNameMap["service_status"] = "ServiceStatus"
	fields["node_convergence_status"] = bindings.NewMapType(bindings.NewStringType(), bindings.NewEnumType("com.vmware.vcenter.hlm.convergence.status", reflect.TypeOf(ConvergenceStatus(ConvergenceStatus_CONVERGED))),reflect.TypeOf(map[string]ConvergenceStatus{}))
	fieldNameMap["node_convergence_status"] = "NodeConvergenceStatus"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vcenter.hlm.convergence.service_convergence_info", fields, reflect.TypeOf(ConvergenceServiceConvergenceInfo{}), fieldNameMap, validators)
}

func ConvergenceConvergenceInfoBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["community_status"] = bindings.NewEnumType("com.vmware.vcenter.hlm.convergence.status", reflect.TypeOf(ConvergenceStatus(ConvergenceStatus_CONVERGED)))
	fieldNameMap["community_status"] = "CommunityStatus"
	fields["service_status"] = bindings.NewMapType(bindings.NewStringType(), bindings.NewReferenceType(ConvergenceServiceConvergenceInfoBindingType),reflect.TypeOf(map[string]ConvergenceServiceConvergenceInfo{}))
	fieldNameMap["service_status"] = "ServiceStatus"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vcenter.hlm.convergence.convergence_info", fields, reflect.TypeOf(ConvergenceConvergenceInfo{}), fieldNameMap, validators)
}

