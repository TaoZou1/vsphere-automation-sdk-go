/* Copyright © 2019 VMware, Inc. All Rights Reserved.
     SPDX-License-Identifier: BSD-2-Clause */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Data type definitions file for service: Sddcs.
 * Includes binding types of a structures and enumerations defined in the service.
 * Shared by client-side stubs and server-side skeletons to ensure type
 * compatibility.
 */

package hlm

import (
    "reflect"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/bindings"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/data"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/protocol"
)




// The ``Info`` class gives information about an Sddc. **Warning:** This class is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
 type SddcsSummary struct {
    // The vCenter hostname. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    Hostname string
    // The Sddc display name. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    Name string
    // The AWS region of the Sddc. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    Region string
}


// TODO some error doesn't have ErrorType. Resolve this later VAPI-3009
//


// The ``FilterSpec`` class contains properties used to list Sddcs. **Warning:** This class is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
 type SddcsFilterSpec struct {
    // Whether or not to filter results based on if the sddc is available to add to the community. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    OnlyAvailableToAdd bool
}


// TODO some error doesn't have ErrorType. Resolve this later VAPI-3009
//





func sddcsListInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["filter"] = bindings.NewOptionalType(bindings.NewReferenceType(SddcsFilterSpecBindingType))
    fieldNameMap["filter"] = "Filter"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func sddcsListOutputType() bindings.BindingType {
    return bindings.NewListType(bindings.NewReferenceType(SddcsSummaryBindingType), reflect.TypeOf([]SddcsSummary{}))
}

func sddcsListRestMetadata() protocol.OperationRestMetadata {
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
       map[string]int{"Unauthorized": 403,"Error": 500})
}



func SddcsSummaryBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["hostname"] = bindings.NewStringType()
    fieldNameMap["hostname"] = "Hostname"
    fields["name"] = bindings.NewStringType()
    fieldNameMap["name"] = "Name"
    fields["region"] = bindings.NewStringType()
    fieldNameMap["region"] = "Region"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.hlm.sddcs.summary",fields, reflect.TypeOf(SddcsSummary{}), fieldNameMap, validators)
}

func SddcsFilterSpecBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["only_available_to_add"] = bindings.NewBooleanType()
    fieldNameMap["only_available_to_add"] = "OnlyAvailableToAdd"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.hlm.sddcs.filter_spec",fields, reflect.TypeOf(SddcsFilterSpec{}), fieldNameMap, validators)
}


