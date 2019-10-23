/* Copyright © 2019 VMware, Inc. All Rights Reserved.
     SPDX-License-Identifier: BSD-2-Clause */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Data type definitions file for service: Proxy.
 * Includes binding types of a structures and enumerations defined in the service.
 * Shared by client-side stubs and server-side skeletons to ensure type
 * compatibility.
 */

package networking

import (
    "reflect"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/bindings/com/vmware/vapi/std"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/bindings"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/data"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/protocol"
)



// ``Protocol`` enumeration class defines the protocols for which proxying is supported.
//
// <p> See {@link com.vmware.vapi.bindings.ApiEnumeration enumerated types description}.
 
type Proxy_Protocol string

const (
    // Proxy configuration for http.
     Proxy_Protocol_HTTP Proxy_Protocol = "HTTP"
    // Proxy configuration for https.
     Proxy_Protocol_HTTPS Proxy_Protocol = "HTTPS"
    // Proxy configuration for ftp.
     Proxy_Protocol_FTP Proxy_Protocol = "FTP"
)

func (p Proxy_Protocol) Proxy_Protocol() bool {
    switch p {
        case Proxy_Protocol_HTTP:
            return true
        case Proxy_Protocol_HTTPS:
            return true
        case Proxy_Protocol_FTP:
            return true
        default:
            return false
    }
}




// ``ServerStatus`` enumeration class defines the status of the server associated with the test run.
//
// <p> See {@link com.vmware.vapi.bindings.ApiEnumeration enumerated types description}.
 
type Proxy_ServerStatus string

const (
    // Server is reachable.
     Proxy_ServerStatus_SERVER_REACHABLE Proxy_ServerStatus = "SERVER_REACHABLE"
    // Server is unreachable.
     Proxy_ServerStatus_SERVER_UNREACHABLE Proxy_ServerStatus = "SERVER_UNREACHABLE"
)

func (s Proxy_ServerStatus) Proxy_ServerStatus() bool {
    switch s {
        case Proxy_ServerStatus_SERVER_REACHABLE:
            return true
        case Proxy_ServerStatus_SERVER_UNREACHABLE:
            return true
        default:
            return false
    }
}





// The ``Config`` class defines proxy configuration.
 type ProxyConfig struct {
    // Hostname or IP address of the proxy server.
    Server string
    // Port to connect to the proxy server. In a 'get' call, indicates the port connected to the proxy server. In a 'set' call, specifies the port to connect to the proxy server. A value of -1 indicates the default port.
    Port int64
    // Username for proxy server.
    Username *string
    // Password for proxy server.
    Password *string
    // In the result of the ``#get`` and ``#list`` methods this property indicates whether proxying is enabled for a particular protocol. In the input to the ``test`` and ``set`` methods this property specifies whether proxying should be enabled for a particular protocol.
    Enabled bool
}



func (ProxyConfig ProxyConfig) Error() string {
    return "com.vmware.appliance.networking.config"
}



// The ``TestResult`` class contains information about the test operation done on a proxy server.
 type ProxyTestResult struct {
    // Status of the proxy server indicating whether the proxy server is reachable.
    Status Proxy_ServerStatus
    // Message associated with status.
    Message std.LocalizableMessage
}



func (ProxyTestResult ProxyTestResult) Error() string {
    return "com.vmware.appliance.networking.test_result"
}






func proxyTestInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["host"] = bindings.NewStringType()
    fields["protocol"] = bindings.NewStringType()
    fields["config"] = bindings.NewReferenceType(ProxyConfigBindingType)
    fieldNameMap["host"] = "Host"
    fieldNameMap["protocol"] = "Protocol"
    fieldNameMap["config"] = "Config"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func proxyTestOutputType() bindings.BindingType {
    return bindings.NewReferenceType(ProxyTestResultBindingType)
}

func proxyTestRestMetadata() protocol.OperationRestMetadata {
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


func proxySetInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["protocol"] = bindings.NewStringType()
    fields["config"] = bindings.NewReferenceType(ProxyConfigBindingType)
    fieldNameMap["protocol"] = "Protocol"
    fieldNameMap["config"] = "Config"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func proxySetOutputType() bindings.BindingType {
    return bindings.NewVoidType()
}

func proxySetRestMetadata() protocol.OperationRestMetadata {
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


func proxyDeleteInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["protocol"] = bindings.NewStringType()
    fieldNameMap["protocol"] = "Protocol"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func proxyDeleteOutputType() bindings.BindingType {
    return bindings.NewVoidType()
}

func proxyDeleteRestMetadata() protocol.OperationRestMetadata {
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


func proxyListInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func proxyListOutputType() bindings.BindingType {
    return bindings.NewMapType(bindings.NewEnumType("com.vmware.appliance.networking.proxy.protocol", reflect.TypeOf(Proxy_Protocol(Proxy_Protocol_HTTP))), bindings.NewReferenceType(ProxyConfigBindingType),reflect.TypeOf(map[Proxy_Protocol]ProxyConfig{}))
}

func proxyListRestMetadata() protocol.OperationRestMetadata {
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


func proxyGetInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["protocol"] = bindings.NewStringType()
    fieldNameMap["protocol"] = "Protocol"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func proxyGetOutputType() bindings.BindingType {
    return bindings.NewReferenceType(ProxyConfigBindingType)
}

func proxyGetRestMetadata() protocol.OperationRestMetadata {
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



func ProxyConfigBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["server"] = bindings.NewStringType()
    fieldNameMap["server"] = "Server"
    fields["port"] = bindings.NewIntegerType()
    fieldNameMap["port"] = "Port"
    fields["username"] = bindings.NewOptionalType(bindings.NewStringType())
    fieldNameMap["username"] = "Username"
    fields["password"] = bindings.NewOptionalType(bindings.NewSecretType())
    fieldNameMap["password"] = "Password"
    fields["enabled"] = bindings.NewBooleanType()
    fieldNameMap["enabled"] = "Enabled"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.appliance.networking.proxy.config",fields, reflect.TypeOf(ProxyConfig{}), fieldNameMap, validators)
}

func ProxyTestResultBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["status"] = bindings.NewEnumType("com.vmware.appliance.networking.proxy.server_status", reflect.TypeOf(Proxy_ServerStatus(Proxy_ServerStatus_SERVER_REACHABLE)))
    fieldNameMap["status"] = "Status"
    fields["message"] = bindings.NewReferenceType(std.LocalizableMessageBindingType)
    fieldNameMap["message"] = "Message"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.appliance.networking.proxy.test_result",fields, reflect.TypeOf(ProxyTestResult{}), fieldNameMap, validators)
}


