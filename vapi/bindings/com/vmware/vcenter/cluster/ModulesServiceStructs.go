/* Copyright © 2019 VMware, Inc. All Rights Reserved.
     SPDX-License-Identifier: BSD-2-Clause */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Data type definitions file for service: Modules.
 * Includes binding types of a structures and enumerations defined in the service.
 * Shared by client-side stubs and server-side skeletons to ensure type
 * compatibility.
 */

package cluster

import (
    "reflect"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/bindings"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/data"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/protocol"
)

// The resource type for the module. **Warning:** This constant field is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
const Modules_RESOURCE_TYPE = "com.vmware.vcenter.cluster.modules"



// The ``CreateSpec`` class contains properties used to create modules, see Modules#create). **Warning:** This class is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
 type ModulesCreateSpec struct {
    // The identifier of the cluster in which this module is created. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    Cluster string
}



func (ModulesCreateSpec ModulesCreateSpec) Error() string {
    return "com.vmware.vcenter.cluster.create_spec"
}



// The ``Summary`` class contains commonly used information about a module in a vCenter cluster. **Warning:** This class is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
 type ModulesSummary struct {
    // Identifier of the cluster in which this module resides. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    Cluster string
    // Identifier of the module. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    Module string
}



func (ModulesSummary ModulesSummary) Error() string {
    return "com.vmware.vcenter.cluster.summary"
}



// The ``ListResult`` class contains a list of commonly used information about each module in this vCenter server. **Warning:** This class is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
 type ModulesListResult struct {
    // List of commonly used information for each module in this vCenter server. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    Summaries []ModulesSummary
}



func (ModulesListResult ModulesListResult) Error() string {
    return "com.vmware.vcenter.cluster.list_result"
}






func modulesCreateInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["spec"] = bindings.NewReferenceType(ModulesCreateSpecBindingType)
    fieldNameMap["spec"] = "Spec"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func modulesCreateOutputType() bindings.BindingType {
    return bindings.NewIdType([]string {"com.vmware.vcenter.cluster.modules"}, "")
}

func modulesCreateRestMetadata() protocol.OperationRestMetadata {
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
       map[string]int{"InvalidArgument": 400})
}


func modulesListInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func modulesListOutputType() bindings.BindingType {
    return bindings.NewReferenceType(ModulesListResultBindingType)
}

func modulesListRestMetadata() protocol.OperationRestMetadata {
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


func modulesDeleteInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["module"] = bindings.NewIdType([]string {"com.vmware.vcenter.cluster.modules"}, "")
    fieldNameMap["module"] = "Module"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func modulesDeleteOutputType() bindings.BindingType {
    return bindings.NewVoidType()
}

func modulesDeleteRestMetadata() protocol.OperationRestMetadata {
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



func ModulesCreateSpecBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["cluster"] = bindings.NewIdType([]string {"ClusterComputeResource"}, "")
    fieldNameMap["cluster"] = "Cluster"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.cluster.modules.create_spec",fields, reflect.TypeOf(ModulesCreateSpec{}), fieldNameMap, validators)
}

func ModulesSummaryBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["cluster"] = bindings.NewIdType([]string {"ClusterComputeResource"}, "")
    fieldNameMap["cluster"] = "Cluster"
    fields["module"] = bindings.NewIdType([]string {"com.vmware.vcenter.cluster.modules"}, "")
    fieldNameMap["module"] = "Module"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.cluster.modules.summary",fields, reflect.TypeOf(ModulesSummary{}), fieldNameMap, validators)
}

func ModulesListResultBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["summaries"] = bindings.NewListType(bindings.NewReferenceType(ModulesSummaryBindingType), reflect.TypeOf([]ModulesSummary{}))
    fieldNameMap["summaries"] = "Summaries"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.cluster.modules.list_result",fields, reflect.TypeOf(ModulesListResult{}), fieldNameMap, validators)
}


