/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Code generated. DO NOT EDIT.

/*
 * Data type definitions file for service: Tickets.
 * Includes binding types of a structures and enumerations defined in the service.
 * Shared by client-side stubs and server-side skeletons to ensure type
 * compatibility.
 */

package console

import (
	"reflect"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/bindings"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/data"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/protocol"
	"net/url"
)



// The ``Type`` enumeration class defines the types of console tickets. **Warning:** This enumeration is available as Technology Preview. These are early access APIs provided to test, automate and provide feedback on the feature. Since this can change based on feedback, VMware does not guarantee backwards compatibility and recommends against using them in production environments. Some Technology Preview APIs might only be applicable to specific environments.
//
// <p> See {@link com.vmware.vapi.bindings.ApiEnumeration enumerated types description}.
type TicketsType string

const (
    // Virtual machine remote console ticket. **Warning:** This constant field is available as Technology Preview. These are early access APIs provided to test, automate and provide feedback on the feature. Since this can change based on feedback, VMware does not guarantee backwards compatibility and recommends against using them in production environments. Some Technology Preview APIs might only be applicable to specific environments.
	TicketsType_VMRC TicketsType = "VMRC"
    // Web socket console ticket. **Warning:** This constant field is available as Technology Preview. These are early access APIs provided to test, automate and provide feedback on the feature. Since this can change based on feedback, VMware does not guarantee backwards compatibility and recommends against using them in production environments. Some Technology Preview APIs might only be applicable to specific environments.
	TicketsType_WEBMKS TicketsType = "WEBMKS"
)

func (t TicketsType) TicketsType() bool {
	switch t {
	case TicketsType_VMRC:
		return true
	case TicketsType_WEBMKS:
		return true
	default:
		return false
	}
}


// The ``CreateSpec`` class defines the information used to create the virtual machine console ticket. **Warning:** This class is available as Technology Preview. These are early access APIs provided to test, automate and provide feedback on the feature. Since this can change based on feedback, VMware does not guarantee backwards compatibility and recommends against using them in production environments. Some Technology Preview APIs might only be applicable to specific environments.
type TicketsCreateSpec struct {
    // The type of virtual machine console ticket. **Warning:** This property is available as Technology Preview. These are early access APIs provided to test, automate and provide feedback on the feature. Since this can change based on feedback, VMware does not guarantee backwards compatibility and recommends against using them in production environments. Some Technology Preview APIs might only be applicable to specific environments.
	Type_ TicketsType
}

// The ``Summary`` class contains commonly used information about the virtual machine console ticket. **Warning:** This class is available as Technology Preview. These are early access APIs provided to test, automate and provide feedback on the feature. Since this can change based on feedback, VMware does not guarantee backwards compatibility and recommends against using them in production environments. Some Technology Preview APIs might only be applicable to specific environments.
type TicketsSummary struct {
    // Console ticket URI. **Warning:** This property is available as Technology Preview. These are early access APIs provided to test, automate and provide feedback on the feature. Since this can change based on feedback, VMware does not guarantee backwards compatibility and recommends against using them in production environments. Some Technology Preview APIs might only be applicable to specific environments.
	Ticket url.URL
}



func ticketsCreateInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["vm"] = bindings.NewIdType([]string{"VirtualMachine"}, "")
	fields["spec"] = bindings.NewReferenceType(TicketsCreateSpecBindingType)
	fieldNameMap["vm"] = "Vm"
	fieldNameMap["spec"] = "Spec"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func ticketsCreateOutputType() bindings.BindingType {
	return bindings.NewReferenceType(TicketsSummaryBindingType)
}

func ticketsCreateRestMetadata() protocol.OperationRestMetadata {
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
		map[string]int{"Error": 500,"InvalidArgument": 400,"NotAllowedInCurrentState": 400,"NotFound": 404,"ResourceInaccessible": 500,"ServiceUnavailable": 503,"Unauthenticated": 401,"Unauthorized": 403})
}


func TicketsCreateSpecBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["type"] = bindings.NewEnumType("com.vmware.vcenter.vm.console.tickets.type", reflect.TypeOf(TicketsType(TicketsType_VMRC)))
	fieldNameMap["type"] = "Type_"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vcenter.vm.console.tickets.create_spec", fields, reflect.TypeOf(TicketsCreateSpec{}), fieldNameMap, validators)
}

func TicketsSummaryBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["ticket"] = bindings.NewUriType()
	fieldNameMap["ticket"] = "Ticket"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vcenter.vm.console.tickets.summary", fields, reflect.TypeOf(TicketsSummary{}), fieldNameMap, validators)
}

