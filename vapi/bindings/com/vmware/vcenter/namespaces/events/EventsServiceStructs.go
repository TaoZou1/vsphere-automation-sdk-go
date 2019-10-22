/* Copyright © 2019 VMware, Inc. All Rights Reserved.
     SPDX-License-Identifier: BSD-2-Clause */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Data type definitions file for service: Events.
 * Includes binding types of a structures and enumerations defined in the service.
 * Shared by client-side stubs and server-side skeletons to ensure type
 * compatibility.
 */

package events

import (
    "reflect"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/bindings"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/data"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/protocol"
)




// The ``Event`` contains selected fields from the corresponding Kubernetes event. Please refer Kubernetes Events API for more details. **Warning:** This class is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
 type EventsEvent struct {
    // The name of the involved Kubernetes object. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    Name string
    // The kind of the involved Kubernetes object. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    Kind string
    // The type of event. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    Type_ string
    // The message in the event. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    Message string
    // The reason for this event. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    Reason string
    // The Kubernetes component that generated this event. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    Component string
    // Number of times this event occurred. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    Count int64
    // The time at which the most recent occurrence of this event was recorded. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    LastTimeStamp int64
}


// TODO some error doesn't have ErrorType. Resolve this later VAPI-3009
//





func eventsGetInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["namespace"] = bindings.NewIdType([]string {"com.vmware.vcenter.namespaces.Instance"}, "")
    fieldNameMap["namespace"] = "Namespace"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func eventsGetOutputType() bindings.BindingType {
    return bindings.NewListType(bindings.NewReferenceType(EventsEventBindingType), reflect.TypeOf([]EventsEvent{}))
}

func eventsGetRestMetadata() protocol.OperationRestMetadata {
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
       map[string]int{"NotFound": 404,"Unsupported": 400,"Error": 500,"Unauthenticated": 401,"Unauthorized": 403})
}



func EventsEventBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["name"] = bindings.NewStringType()
    fieldNameMap["name"] = "Name"
    fields["kind"] = bindings.NewStringType()
    fieldNameMap["kind"] = "Kind"
    fields["type"] = bindings.NewStringType()
    fieldNameMap["type"] = "Type_"
    fields["message"] = bindings.NewStringType()
    fieldNameMap["message"] = "Message"
    fields["reason"] = bindings.NewStringType()
    fieldNameMap["reason"] = "Reason"
    fields["component"] = bindings.NewStringType()
    fieldNameMap["component"] = "Component"
    fields["count"] = bindings.NewIntegerType()
    fieldNameMap["count"] = "Count"
    fields["last_time_stamp"] = bindings.NewIntegerType()
    fieldNameMap["last_time_stamp"] = "LastTimeStamp"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.namespaces.events.events.event",fields, reflect.TypeOf(EventsEvent{}), fieldNameMap, validators)
}


