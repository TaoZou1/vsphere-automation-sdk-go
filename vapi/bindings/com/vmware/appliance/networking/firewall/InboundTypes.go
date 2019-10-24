/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Data type definitions file for service: Inbound.
 * Includes binding types of a structures and enumerations defined in the service.
 * Shared by client-side stubs and server-side skeletons to ensure type
 * compatibility.
 */
package firewall

import (
	"reflect"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/bindings"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/data"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/protocol"
)


// ``Policy`` enumeration class Defines firewall rule policies.
//
// <p> See {@link com.vmware.vapi.bindings.ApiEnumeration enumerated types description}.
type InboundPolicy string

const (
    // Drop packet with correpsonding address.
	InboundPolicy_IGNORE InboundPolicy = "IGNORE"
    // Allow packet with corresponding address.
	InboundPolicy_ACCEPT InboundPolicy = "ACCEPT"
    // Drop packet with corresponding address sending destination is not reachable.
	InboundPolicy_REJECT InboundPolicy = "REJECT"
    // Apply default or port-specific rules to packet with corresponding address.
	InboundPolicy_RETURN InboundPolicy = "RETURN"
)

func (p InboundPolicy) InboundPolicy() bool {
	switch p {
	case InboundPolicy_IGNORE:
		return true
	case InboundPolicy_ACCEPT:
		return true
	case InboundPolicy_REJECT:
		return true
	case InboundPolicy_RETURN:
		return true
	default:
		return false
	}
}

// ``Rule`` class Structure that defines a single address-based firewall rule.
type InboundRule struct {
    // IPv4 or IPv6 address.
	Address string
    // CIDR prefix used to mask address. For example, an IPv4 prefix of 24 ignores the low-order 8 bits of address.
	Prefix int64
    // The allow or deny policy of this rule.
	Policy InboundPolicy
    // The interface to which this rule applies. An empty string indicates that the rule applies to all interfaces.
	InterfaceName *string
}


func inboundSetInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["rules"] = bindings.NewListType(bindings.NewReferenceType(InboundRuleBindingType), reflect.TypeOf([]InboundRule{}))
	fieldNameMap["rules"] = "Rules"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func inboundSetOutputType() bindings.BindingType {
	return bindings.NewVoidType()
}

func inboundSetRestMetadata() protocol.OperationRestMetadata {
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
		map[string]int{"Error": 500})
}

func inboundGetInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func inboundGetOutputType() bindings.BindingType {
	return bindings.NewListType(bindings.NewReferenceType(InboundRuleBindingType), reflect.TypeOf([]InboundRule{}))
}

func inboundGetRestMetadata() protocol.OperationRestMetadata {
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
		map[string]int{"Error": 500})
}


func InboundRuleBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["address"] = bindings.NewStringType()
	fieldNameMap["address"] = "Address"
	fields["prefix"] = bindings.NewIntegerType()
	fieldNameMap["prefix"] = "Prefix"
	fields["policy"] = bindings.NewEnumType("com.vmware.appliance.networking.firewall.inbound.policy", reflect.TypeOf(InboundPolicy(InboundPolicy_IGNORE)))
	fieldNameMap["policy"] = "Policy"
	fields["interface_name"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["interface_name"] = "InterfaceName"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.appliance.networking.firewall.inbound.rule", fields, reflect.TypeOf(InboundRule{}), fieldNameMap, validators)
}
