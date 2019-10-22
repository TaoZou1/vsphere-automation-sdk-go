/* Copyright © 2019 VMware, Inc. All Rights Reserved.
     SPDX-License-Identifier: BSD-2-Clause */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Data type definitions file for service: Registries.
 * Includes binding types of a structures and enumerations defined in the service.
 * Shared by client-side stubs and server-side skeletons to ensure type
 * compatibility.
 */

package registries

import (
    "reflect"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/bindings"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/data"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/protocol"
    "net/url"
)

// The resource type for container registry. **Warning:** This constant field is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
const Registries_RESOURCE_TYPE = "com.vmware.vcenter.content.Registry"



// The ``Summary`` class contains basic information about a container registry. **Warning:** This class is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
 type RegistriesSummary struct {
    // Identifier of the container registry. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    Registry string
    // Version of the container registry. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    Version string
    // URL to access the container registry. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    AccessUrl url.URL
}


// TODO some error doesn't have ErrorType. Resolve this later VAPI-3009
//





func registriesListInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func registriesListOutputType() bindings.BindingType {
    return bindings.NewListType(bindings.NewReferenceType(RegistriesSummaryBindingType), reflect.TypeOf([]RegistriesSummary{}))
}

func registriesListRestMetadata() protocol.OperationRestMetadata {
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
       map[string]int{"Error": 500,"Unauthenticated": 401,"Unauthorized": 403})
}



func RegistriesSummaryBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["registry"] = bindings.NewIdType([]string {"com.vmware.vcenter.content.Registry"}, "")
    fieldNameMap["registry"] = "Registry"
    fields["version"] = bindings.NewStringType()
    fieldNameMap["version"] = "Version"
    fields["access_url"] = bindings.NewUriType()
    fieldNameMap["access_url"] = "AccessUrl"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.content.registries.registries.summary",fields, reflect.TypeOf(RegistriesSummary{}), fieldNameMap, validators)
}


