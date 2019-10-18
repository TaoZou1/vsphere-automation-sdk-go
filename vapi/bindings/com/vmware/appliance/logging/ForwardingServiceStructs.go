/* Copyright © 2019 VMware, Inc. All Rights Reserved.
     SPDX-License-Identifier: BSD-2-Clause */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Data type definitions file for service: Forwarding.
 * Includes binding types of a structures and enumerations defined in the service.
 * Shared by client-side stubs and server-side skeletons to ensure type
 * compatibility.
 */

package logging

import (
    "reflect"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/bindings/com/vmware/vapi/std"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/bindings"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/data"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/protocol"
)



// The ``Protocol`` enumeration class defines transport protocols for outbound log messages. This enumeration was added in vSphere API 6.7.
//
// <p> See {@link com.vmware.vapi.bindings.ApiEnumeration enumerated types description}.
 
type ForwardingProtocol string

const (
    // Log messages will be forwarded to the remote host by using the TLS protocol. This constant field was added in vSphere API 6.7.
     ForwardingProtocol_TLS ForwardingProtocol = "TLS"
    // Log messages will be forwarded to the remote host using the UDP protocol. This constant field was added in vSphere API 6.7.
     ForwardingProtocol_UDP ForwardingProtocol = "UDP"
    // Log messages will be forwarded to the remote host using the TCP protocol. This constant field was added in vSphere API 6.7.
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






// The ``Config`` class defines the configuration for log message forwarding to remote logging servers. This class was added in vSphere API 6.7.
 type ForwardingConfig struct {
    // FQDN or IP address of the logging server to which messages are forwarded. This property was added in vSphere API 6.7.
    Hostname string
    // The port on which the remote logging server is listening for forwarded log messages. This property was added in vSphere API 6.7.
    Port int64
    // Transport protocol used to forward log messages. This property was added in vSphere API 6.7.
    Protocol ForwardingProtocol
}







 type ForwardingConnectionStatus struct {
    // FQDN or IP address of the configured remote logging servers. This property was added in vSphere API 6.7.
    Hostname string
    // State of the configured remote logging server. This property was added in vSphere API 6.7.
    State ConnectionStatus_State
    // Message associated with the state of the configured remote logging server. This property was added in vSphere API 6.7.
    Message *std.LocalizableMessage
}




    
    // The ``State`` enumeration class defines the state values that a remote logging server can be in. This enumeration was added in vSphere API 6.7.
    //
    // <p> See {@link com.vmware.vapi.bindings.ApiEnumeration enumerated types description}.
     
    type ConnectionStatus_State string

    const (
        // The remote logging server is reachable. This constant field was added in vSphere API 6.7.
         ConnectionStatus_State_UP ConnectionStatus_State = "UP"
        // The remote logging server is not reachable. This constant field was added in vSphere API 6.7.
         ConnectionStatus_State_DOWN ConnectionStatus_State = "DOWN"
        // The status of remote logging server is unknown. This constant field was added in vSphere API 6.7.
         ConnectionStatus_State_UNKNOWN ConnectionStatus_State = "UNKNOWN"
    )

    func (s ConnectionStatus_State) ConnectionStatus_State() bool {
        switch s {
            case ConnectionStatus_State_UP:
                return true
            case ConnectionStatus_State_DOWN:
                return true
            case ConnectionStatus_State_UNKNOWN:
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
    return bindings.NewStructType("com.vmware.appliance.logging.forwarding.config",fields, reflect.TypeOf(ForwardingConfig{}), fieldNameMap, validators)
}

func ForwardingConnectionStatusBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["hostname"] = bindings.NewStringType()
    fieldNameMap["hostname"] = "Hostname"
    fields["state"] = bindings.NewEnumType("com.vmware.appliance.logging.forwarding.connection_status.state", reflect.TypeOf(ConnectionStatus_State(ConnectionStatus_State_UP)))
    fieldNameMap["state"] = "State"
    fields["message"] = bindings.NewOptionalType(bindings.NewReferenceType(std.LocalizableMessageBindingType))
    fieldNameMap["message"] = "Message"
    var validators = []bindings.Validator{}
    uv1 := bindings.NewUnionValidator("state",
        map[string][]bindings.FieldData {
            "DOWN": []bindings.FieldData {
                 bindings.NewFieldData("message", false),
            },
            "UP": []bindings.FieldData {},
            "UNKNOWN": []bindings.FieldData {},
        },
    )
    validators = append(validators, uv1)
    return bindings.NewStructType("com.vmware.appliance.logging.forwarding.connection_status",fields, reflect.TypeOf(ForwardingConnectionStatus{}), fieldNameMap, validators)
}


