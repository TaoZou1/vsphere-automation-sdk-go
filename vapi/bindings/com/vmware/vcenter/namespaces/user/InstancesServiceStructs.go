/* Copyright © 2019 VMware, Inc. All Rights Reserved.
     SPDX-License-Identifier: BSD-2-Clause */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Data type definitions file for service: Instances.
 * Includes binding types of a structures and enumerations defined in the service.
 * Shared by client-side stubs and server-side skeletons to ensure type
 * compatibility.
 */

package user

import (
    "reflect"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/bindings"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/data"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/protocol"
)





// The ``Summary`` class contains information about a namespace that user is authorized to access. **Warning:** This class is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
 type InstancesSummary struct {
    // Identifier of the namespace. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    Namespace string
    // IP address or FQDN of the API endpoint for the given namespace. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    MasterHost string
}









func instancesListInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func instancesListOutputType() bindings.BindingType {
    return bindings.NewListType(bindings.NewReferenceType(InstancesSummaryBindingType), reflect.TypeOf([]InstancesSummary{}))
}

func instancesListRestMetadata() protocol.OperationRestMetadata {
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
      "/vcenter/namespaces-user/namespaces",
       resultHeaders,
       200,
       errorHeaders,
       map[string]int{"Error": 500,"Unauthenticated": 401})
}



func InstancesSummaryBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["namespace"] = bindings.NewIdType([]string {"com.vmware.vcenter.namespaces.Instance"}, "")
    fieldNameMap["namespace"] = "Namespace"
    fields["master_host"] = bindings.NewStringType()
    fieldNameMap["master_host"] = "MasterHost"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.namespaces.user.instances.summary",fields, reflect.TypeOf(InstancesSummary{}), fieldNameMap, validators)
}


