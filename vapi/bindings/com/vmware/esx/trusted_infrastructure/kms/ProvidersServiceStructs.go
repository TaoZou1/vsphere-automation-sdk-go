/* Copyright © 2019 VMware, Inc. All Rights Reserved.
     SPDX-License-Identifier: BSD-2-Clause */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Data type definitions file for service: Providers.
 * Includes binding types of a structures and enumerations defined in the service.
 * Shared by client-side stubs and server-side skeletons to ensure type
 * compatibility.
 */

package kms

import (
    "reflect"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/bindings/com/vmware/vapi/std"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/bindings"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/data"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/protocol"
)

// Resource type for a Key Provider.
const Providers_RESOURCE_TYPE = "com.vmware.esx.trusted_infrastructure.kms.provider"



// The ``Summary`` class contains a summary of a key provider.
 type ProvidersSummary struct {
    // Identifier of the provider.
    Provider string
    // Whether the provider is available for key operation.
    Available bool
    // Details regarding the availability of the provider.
    Details []std.LocalizableMessage
}



func (ProvidersSummary ProvidersSummary) Error() string {
    return "com.vmware.esx.trusted_infrastructure.kms.summary"
}



// The ``FilterSpec`` class contains properties used to filter the results when listing the providers.
 type ProvidersFilterSpec struct {
    // Identifier of the provider.
    Providers map[string]bool
    // Whether the provider is available for key operations.
    Available *bool
}



func (ProvidersFilterSpec ProvidersFilterSpec) Error() string {
    return "com.vmware.esx.trusted_infrastructure.kms.filter_spec"
}






func providersListInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["filter_spec"] = bindings.NewOptionalType(bindings.NewReferenceType(ProvidersFilterSpecBindingType))
    fieldNameMap["filter_spec"] = "FilterSpec"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func providersListOutputType() bindings.BindingType {
    return bindings.NewListType(bindings.NewReferenceType(ProvidersSummaryBindingType), reflect.TypeOf([]ProvidersSummary{}))
}

func providersListRestMetadata() protocol.OperationRestMetadata {
    paramsTypeMap := map[string]bindings.BindingType{}
    pathParams := map[string]string{}
    queryParams := map[string]string{}
    headerParams := map[string]string{}
    paramsTypeMap["filter_spec.available"] = bindings.NewOptionalType(bindings.NewBooleanType())
    paramsTypeMap["filter_spec.providers"] = bindings.NewOptionalType(bindings.NewSetType(bindings.NewIdType([]string {"com.vmware.esx.trusted_infrastructure.kms.provider"}, ""), reflect.TypeOf(map[string]bool{})))
    queryParams["filter_spec.available"] = "available"
    queryParams["filter_spec.providers"] = "providers"
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
      "/esx/trusted-infrastructure/kms/providers",
       resultHeaders,
       200,
       errorHeaders,
       map[string]int{"Unauthenticated": 401,"Unauthorized": 403,"Error": 500})
}



func ProvidersSummaryBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["provider"] = bindings.NewIdType([]string {"com.vmware.esx.trusted_infrastructure.kms.provider"}, "")
    fieldNameMap["provider"] = "Provider"
    fields["available"] = bindings.NewBooleanType()
    fieldNameMap["available"] = "Available"
    fields["details"] = bindings.NewListType(bindings.NewReferenceType(std.LocalizableMessageBindingType), reflect.TypeOf([]std.LocalizableMessage{}))
    fieldNameMap["details"] = "Details"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.esx.trusted_infrastructure.kms.providers.summary",fields, reflect.TypeOf(ProvidersSummary{}), fieldNameMap, validators)
}

func ProvidersFilterSpecBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["providers"] = bindings.NewOptionalType(bindings.NewSetType(bindings.NewIdType([]string {"com.vmware.esx.trusted_infrastructure.kms.provider"}, ""), reflect.TypeOf(map[string]bool{})))
    fieldNameMap["providers"] = "Providers"
    fields["available"] = bindings.NewOptionalType(bindings.NewBooleanType())
    fieldNameMap["available"] = "Available"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.esx.trusted_infrastructure.kms.providers.filter_spec",fields, reflect.TypeOf(ProvidersFilterSpec{}), fieldNameMap, validators)
}


