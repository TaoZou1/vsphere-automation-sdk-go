/* Copyright © 2019 VMware, Inc. All Rights Reserved.
     SPDX-License-Identifier: BSD-2-Clause */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Data type definitions file for service: CounterSets.
 * Includes binding types of a structures and enumerations defined in the service.
 * Shared by client-side stubs and server-side skeletons to ensure type
 * compatibility.
 */

package vstats

import (
    "reflect"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/bindings"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/data"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/protocol"
)

// Resource type for counter sets.
const CounterSets_RESOURCE_TYPE = "com.vmware.vstats.model.CounterSet"



// The ``Info`` class contains information about a set of counters.
 type CounterSetsInfo struct {
    // Counter set identifier.
    Id string
    // List of Counter CidMids.
    Counters []CidMid
    // Human legible localizable conter set description.
    UserInfo *UserInfo
}


// TODO some error doesn't have ErrorType. Resolve this later VAPI-3009
//





func counterSetsListInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func counterSetsListOutputType() bindings.BindingType {
    return bindings.NewListType(bindings.NewReferenceType(CounterSetsInfoBindingType), reflect.TypeOf([]CounterSetsInfo{}))
}

func counterSetsListRestMetadata() protocol.OperationRestMetadata {
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
      "/stats/counter-sets",
       resultHeaders,
       200,
       errorHeaders,
       map[string]int{"Error": 500,"Unauthenticated": 401,"Unauthorized": 403})
}


func counterSetsGetInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["counter_set"] = bindings.NewIdType([]string {"com.vmware.vstats.model.CounterSet"}, "")
    fieldNameMap["counter_set"] = "CounterSet"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func counterSetsGetOutputType() bindings.BindingType {
    return bindings.NewReferenceType(CounterSetsInfoBindingType)
}

func counterSetsGetRestMetadata() protocol.OperationRestMetadata {
    paramsTypeMap := map[string]bindings.BindingType{}
    pathParams := map[string]string{}
    queryParams := map[string]string{}
    headerParams := map[string]string{}
    paramsTypeMap["counter_set"] = bindings.NewIdType([]string {"com.vmware.vstats.model.CounterSet"}, "")
    paramsTypeMap["counterSet"] = bindings.NewIdType([]string {"com.vmware.vstats.model.CounterSet"}, "")
    pathParams["counter_set"] = "counterSet"
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
      "/stats/counter-sets/{counterSet}",
       resultHeaders,
       200,
       errorHeaders,
       map[string]int{"Error": 500,"InvalidArgument": 400,"NotFound": 404,"Unauthenticated": 401,"Unauthorized": 403})
}



func CounterSetsInfoBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["id"] = bindings.NewStringType()
    fieldNameMap["id"] = "Id"
    fields["counters"] = bindings.NewListType(bindings.NewReferenceType(CidMidBindingType), reflect.TypeOf([]CidMid{}))
    fieldNameMap["counters"] = "Counters"
    fields["user_info"] = bindings.NewOptionalType(bindings.NewReferenceType(UserInfoBindingType))
    fieldNameMap["user_info"] = "UserInfo"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vstats.counter_sets.info",fields, reflect.TypeOf(CounterSetsInfo{}), fieldNameMap, validators)
}


