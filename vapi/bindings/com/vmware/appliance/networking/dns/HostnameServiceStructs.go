/* Copyright © 2019 VMware, Inc. All Rights Reserved.
     SPDX-License-Identifier: BSD-2-Clause */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Data type definitions file for service: Hostname.
 * Includes binding types of a structures and enumerations defined in the service.
 * Shared by client-side stubs and server-side skeletons to ensure type
 * compatibility.
 */

package dns

import (
    "reflect"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/bindings"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/data"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/protocol"
)



// ``TestStatus`` enumeration class Health indicator
//
// <p> See {@link com.vmware.vapi.bindings.ApiEnumeration enumerated types description}.
 
type Hostname_TestStatus string

const (
    // In case data has more than one test, this indicates not all tests were successful
     Hostname_TestStatus_orange Hostname_TestStatus = "orange"
    // All tests were successful for given data
     Hostname_TestStatus_green Hostname_TestStatus = "green"
    // All tests failed for given data
     Hostname_TestStatus_red Hostname_TestStatus = "red"
)

func (t Hostname_TestStatus) Hostname_TestStatus() bool {
    switch t {
        case Hostname_TestStatus_orange:
            return true
        case Hostname_TestStatus_green:
            return true
        case Hostname_TestStatus_red:
            return true
        default:
            return false
    }
}




// ``MessageStatus`` enumeration class Individual test result
//
// <p> See {@link com.vmware.vapi.bindings.ApiEnumeration enumerated types description}.
 
type Hostname_MessageStatus string

const (
    // message indicates the test failed.
     Hostname_MessageStatus_failure Hostname_MessageStatus = "failure"
    // message indicates that the test was successful.
     Hostname_MessageStatus_success Hostname_MessageStatus = "success"
)

func (m Hostname_MessageStatus) Hostname_MessageStatus() bool {
    switch m {
        case Hostname_MessageStatus_failure:
            return true
        case Hostname_MessageStatus_success:
            return true
        default:
            return false
    }
}





// ``Message`` class Test result and message
 type HostnameMessage struct {
    // message
    Message string
    // result of the test
    Result Hostname_MessageStatus
}



//


// ``TestStatusInfo`` class Overall test result
 type HostnameTestStatusInfo struct {
    // Overall status of tests run.
    Status Hostname_TestStatus
    // messages
    Messages []HostnameMessage
}



//





func hostnameTestInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["name"] = bindings.NewStringType()
    fieldNameMap["name"] = "Name"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func hostnameTestOutputType() bindings.BindingType {
    return bindings.NewReferenceType(HostnameTestStatusInfoBindingType)
}

func hostnameTestRestMetadata() protocol.OperationRestMetadata {
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


func hostnameSetInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["name"] = bindings.NewStringType()
    fieldNameMap["name"] = "Name"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func hostnameSetOutputType() bindings.BindingType {
    return bindings.NewVoidType()
}

func hostnameSetRestMetadata() protocol.OperationRestMetadata {
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


func hostnameGetInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func hostnameGetOutputType() bindings.BindingType {
    return bindings.NewStringType()
}

func hostnameGetRestMetadata() protocol.OperationRestMetadata {
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



func HostnameMessageBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["message"] = bindings.NewStringType()
    fieldNameMap["message"] = "Message"
    fields["result"] = bindings.NewEnumType("com.vmware.appliance.networking.dns.hostname.message_status", reflect.TypeOf(Hostname_MessageStatus(Hostname_MessageStatus_failure)))
    fieldNameMap["result"] = "Result"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.appliance.networking.dns.hostname.message",fields, reflect.TypeOf(HostnameMessage{}), fieldNameMap, validators)
}

func HostnameTestStatusInfoBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["status"] = bindings.NewEnumType("com.vmware.appliance.networking.dns.hostname.test_status", reflect.TypeOf(Hostname_TestStatus(Hostname_TestStatus_orange)))
    fieldNameMap["status"] = "Status"
    fields["messages"] = bindings.NewListType(bindings.NewReferenceType(HostnameMessageBindingType), reflect.TypeOf([]HostnameMessage{}))
    fieldNameMap["messages"] = "Messages"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.appliance.networking.dns.hostname.test_status_info",fields, reflect.TypeOf(HostnameTestStatusInfo{}), fieldNameMap, validators)
}


