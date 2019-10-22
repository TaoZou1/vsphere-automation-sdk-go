/* Copyright © 2019 VMware, Inc. All Rights Reserved.
     SPDX-License-Identifier: BSD-2-Clause */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Data type definitions file for service: VM.
 * Includes binding types of a structures and enumerations defined in the service.
 * Shared by client-side stubs and server-side skeletons to ensure type
 * compatibility.
 */

package policies

import (
    "reflect"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/bindings"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/data"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/protocol"
)




// The ``Summary`` describes a policy and its compliance status. **Warning:** This class is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
 type VMSummary struct {
    // The identifier of a policy. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    Policy string
    // The compliance status of the policy. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    Status PolicyCompliance
}


// TODO some error doesn't have ErrorType. Resolve this later VAPI-3009
//


// The ``ListResult`` class contains the returned policies, see VM#list. **Warning:** This class is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
 type VMListResult struct {
    // The list of policies. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    Policies []VMSummary
}


// TODO some error doesn't have ErrorType. Resolve this later VAPI-3009
//





func vMListInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func vMListOutputType() bindings.BindingType {
    return bindings.NewReferenceType(VMListResultBindingType)
}

func vMListRestMetadata() protocol.OperationRestMetadata {
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
      "GET",
      "/vcenter/compute/policies/vm",
       resultHeaders,
       200,
       errorHeaders,
       map[string]int{"Unauthorized": 403})
}



func VMSummaryBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["policy"] = bindings.NewIdType([]string {"com.vmware.vcenter.compute.Policy"}, "")
    fieldNameMap["policy"] = "Policy"
    fields["status"] = bindings.NewEnumType("com.vmware.vcenter.compute.policies.policy_compliance", reflect.TypeOf(PolicyCompliance(PolicyCompliance_UNKNOWN)))
    fieldNameMap["status"] = "Status"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.compute.policies.VM.summary",fields, reflect.TypeOf(VMSummary{}), fieldNameMap, validators)
}

func VMListResultBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["policies"] = bindings.NewListType(bindings.NewReferenceType(VMSummaryBindingType), reflect.TypeOf([]VMSummary{}))
    fieldNameMap["policies"] = "Policies"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.compute.policies.VM.list_result",fields, reflect.TypeOf(VMListResult{}), fieldNameMap, validators)
}


