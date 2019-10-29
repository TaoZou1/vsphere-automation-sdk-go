/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Code generated. DO NOT EDIT.

/*
 * Data type definitions file for service: Forwarding.
 * Includes binding types of a structures and enumerations defined in the service.
 * Shared by client-side stubs and server-side skeletons to ensure type
 * compatibility.
 */

package logging

import (
	"reflect"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/bindings/vapi/std"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/bindings"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/data"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/protocol"
)



// The ``Protocol`` enumeration class defines transport protocols for outbound log messages.
//
// <p> See {@link com.vmware.vapi.bindings.ApiEnumeration enumerated types description}.
type ForwardingProtocol string

const (
    // Log messages will be forwarded to the remote host by using the TLS protocol.
	ForwardingProtocol_TLS ForwardingProtocol = "TLS"
    // Log messages will be forwarded to the remote host using the UDP protocol.
	ForwardingProtocol_UDP ForwardingProtocol = "UDP"
    // Log messages will be forwarded to the remote host using the TCP protocol.
	ForwardingProtocol_TCP ForwardingProtocol = "TCP"
)

func (p ForwardingProtocol) ForwardingProtocol() bool {
	switch p {
	case ForwardingProtocol_TLS:
		return true
	case ForwardingProtocol_UDP:
		return true
	case ForwardingProtocol_TCP:
		return true
	default:
		return false
	}
}


// The ``Config`` class defines the configuration for log message forwarding to remote logging servers.
type ForwardingConfig struct {
    // FQDN or IP address of the logging server to which messages are forwarded.
	Hostname string
    // The port on which the remote logging server is listening for forwarded log messages.
	Port int64
    // Transport protocol used to forward log messages.
	Protocol ForwardingProtocol
}

type ForwardingConnectionStatus struct {
    // FQDN or IP address of the configured remote logging servers.
	Hostname string
    // State of the configured remote logging server.
	State ForwardingConnectionStatusState
    // Message associated with the state of the configured remote logging server.
	Message *std.LocalizableMessage
}

// The ``State`` enumeration class defines the state values that a remote logging server can be in.
//
// <p> See {@link com.vmware.vapi.bindings.ApiEnumeration enumerated types description}.
type ForwardingConnectionStatusState string

const (
    // The remote logging server is reachable.
	ForwardingConnectionStatusState_UP ForwardingConnectionStatusState = "UP"
    // The remote logging server is not reachable.
	ForwardingConnectionStatusState_DOWN ForwardingConnectionStatusState = "DOWN"
    // The status of remote logging server is unknown.
	ForwardingConnectionStatusState_UNKNOWN ForwardingConnectionStatusState = "UNKNOWN"
)

func (s ForwardingConnectionStatusState) ForwardingConnectionStatusState() bool {
	switch s {
	case ForwardingConnectionStatusState_UP:
		return true
	case ForwardingConnectionStatusState_DOWN:
		return true
	case ForwardingConnectionStatusState_UNKNOWN:
		return true
	default:
		return false
	}
}




func forwardingTestInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["send_test_message"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fieldNameMap["send_test_message"] = "SendTestMessage"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func forwardingTestOutputType() bindings.BindingType {
	return bindings.NewListType(bindings.NewReferenceType(ForwardingConnectionStatusBindingType), reflect.TypeOf([]ForwardingConnectionStatus{}))
}

func forwardingTestRestMetadata() protocol.OperationRestMetadata {
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
		map[string]int{})
}

func forwardingSetInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["cfg_list"] = bindings.NewListType(bindings.NewReferenceType(ForwardingConfigBindingType), reflect.TypeOf([]ForwardingConfig{}))
	fieldNameMap["cfg_list"] = "CfgList"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func forwardingSetOutputType() bindings.BindingType {
	return bindings.NewVoidType()
}

func forwardingSetRestMetadata() protocol.OperationRestMetadata {
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
		map[string]int{"InvalidArgument": 400,"UnableToAllocateResource": 500,"Error": 500})
}

func forwardingGetInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func forwardingGetOutputType() bindings.BindingType {
	return bindings.NewListType(bindings.NewReferenceType(ForwardingConfigBindingType), reflect.TypeOf([]ForwardingConfig{}))
}

func forwardingGetRestMetadata() protocol.OperationRestMetadata {
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
		map[string]int{})
}


func ForwardingConfigBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["hostname"] = bindings.NewStringType()
	fieldNameMap["hostname"] = "Hostname"
	fields["port"] = bindings.NewIntegerType()
	fieldNameMap["port"] = "Port"
	fields["protocol"] = bindings.NewEnumType("com.vmware.appliance.logging.forwarding.protocol", reflect.TypeOf(ForwardingProtocol(ForwardingProtocol_TLS)))
	fieldNameMap["protocol"] = "Protocol"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.appliance.logging.forwarding.config", fields, reflect.TypeOf(ForwardingConfig{}), fieldNameMap, validators)
}

func ForwardingConnectionStatusBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["hostname"] = bindings.NewStringType()
	fieldNameMap["hostname"] = "Hostname"
	fields["state"] = bindings.NewEnumType("com.vmware.appliance.logging.forwarding.connection_status.state", reflect.TypeOf(ForwardingConnectionStatusState(ForwardingConnectionStatusState_UP)))
	fieldNameMap["state"] = "State"
	fields["message"] = bindings.NewOptionalType(bindings.NewReferenceType(std.LocalizableMessageBindingType))
	fieldNameMap["message"] = "Message"
	var validators = []bindings.Validator{}
	uv1 := bindings.NewUnionValidator("state",
		map[string][]bindings.FieldData{
			"DOWN": []bindings.FieldData{
				bindings.NewFieldData("message", false),
			},
			"UP": []bindings.FieldData{},
			"UNKNOWN": []bindings.FieldData{},
		},
	)
	validators = append(validators, uv1)
	return bindings.NewStructType("com.vmware.appliance.logging.forwarding.connection_status", fields, reflect.TypeOf(ForwardingConnectionStatus{}), fieldNameMap, validators)
}

