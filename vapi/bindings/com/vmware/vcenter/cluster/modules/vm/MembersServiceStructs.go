/* Copyright © 2019 VMware, Inc. All Rights Reserved.
     SPDX-License-Identifier: BSD-2-Clause */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Data type definitions file for service: Members.
 * Includes binding types of a structures and enumerations defined in the service.
 * Shared by client-side stubs and server-side skeletons to ensure type
 * compatibility.
 */

package vm

import (
    "reflect"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/bindings"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/data"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/protocol"
)




// The ``Status`` class contains properties used to report the status of updating module membership, see Members#add and Members#remove). **Warning:** This class is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
 type MembersStatus struct {
    // This is true if the complete method succeeded without any errors. Otherwise it is false and all or some methods have failed. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    Success bool
}


// TODO some error doesn't have ErrorType. Resolve this later VAPI-3009
//


// The ``GetResult`` class contains properties used to return the virtual machines that are members of the module, see Members#get). **Warning:** This class is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
 type MembersGetResult struct {
    // The set of virtual machines. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    Vms map[string]bool
}


// TODO some error doesn't have ErrorType. Resolve this later VAPI-3009
//





func membersAddInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["module"] = bindings.NewIdType([]string {"com.vmware.vcenter.cluster.modules"}, "")
    fields["vms"] = bindings.NewSetType(bindings.NewIdType([]string {"VirtualMachine"}, ""), reflect.TypeOf(map[string]bool{}))
    fieldNameMap["module"] = "Module"
    fieldNameMap["vms"] = "Vms"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func membersAddOutputType() bindings.BindingType {
    return bindings.NewReferenceType(MembersStatusBindingType)
}

func membersAddRestMetadata() protocol.OperationRestMetadata {
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
       map[string]int{"NotFound": 404})
}


func membersRemoveInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["module"] = bindings.NewIdType([]string {"com.vmware.vcenter.cluster.modules"}, "")
    fields["vms"] = bindings.NewSetType(bindings.NewIdType([]string {"VirtualMachine"}, ""), reflect.TypeOf(map[string]bool{}))
    fieldNameMap["module"] = "Module"
    fieldNameMap["vms"] = "Vms"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func membersRemoveOutputType() bindings.BindingType {
    return bindings.NewReferenceType(MembersStatusBindingType)
}

func membersRemoveRestMetadata() protocol.OperationRestMetadata {
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
       map[string]int{"NotFound": 404})
}


func membersGetInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["module"] = bindings.NewIdType([]string {"com.vmware.vcenter.cluster.modules"}, "")
    fieldNameMap["module"] = "Module"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func membersGetOutputType() bindings.BindingType {
    return bindings.NewReferenceType(MembersGetResultBindingType)
}

func membersGetRestMetadata() protocol.OperationRestMetadata {
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
       map[string]int{"NotFound": 404})
}



func MembersStatusBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["success"] = bindings.NewBooleanType()
    fieldNameMap["success"] = "Success"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.cluster.modules.vm.members.status",fields, reflect.TypeOf(MembersStatus{}), fieldNameMap, validators)
}

func MembersGetResultBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["vms"] = bindings.NewSetType(bindings.NewIdType([]string {"VirtualMachine"}, ""), reflect.TypeOf(map[string]bool{}))
    fieldNameMap["vms"] = "Vms"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.cluster.modules.vm.members.get_result",fields, reflect.TypeOf(MembersGetResult{}), fieldNameMap, validators)
}


