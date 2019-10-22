/* Copyright © 2019 VMware, Inc. All Rights Reserved.
     SPDX-License-Identifier: BSD-2-Clause */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Data type definitions file for service: Host.
 * Includes binding types of a structures and enumerations defined in the service.
 * Shared by client-side stubs and server-side skeletons to ensure type
 * compatibility.
 */

package tag

import (
    "reflect"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/bindings"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/data"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/protocol"
)

// The resource type for tags that can be attached to hosts. **Warning:** This constant field is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
const Host_RESOURCE_TYPE = "com.vmware.cis.tagging.Tag:HostSystem"



 type HostSummary struct {
    // Identifier of the tag. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    Tag string
    // Name of the tag. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    Name string
    // Description of the tag. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    Description string
    // Name of the category of which this tag is part. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    CategoryName string
}


// TODO some error doesn't have ErrorType. Resolve this later VAPI-3009
//





func hostListInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func hostListOutputType() bindings.BindingType {
    return bindings.NewListType(bindings.NewReferenceType(HostSummaryBindingType), reflect.TypeOf([]HostSummary{}))
}

func hostListRestMetadata() protocol.OperationRestMetadata {
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



func HostSummaryBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["tag"] = bindings.NewIdType([]string {"com.vmware.cis.tagging.Tag:HostSystem"}, "")
    fieldNameMap["tag"] = "Tag"
    fields["name"] = bindings.NewStringType()
    fieldNameMap["name"] = "Name"
    fields["description"] = bindings.NewStringType()
    fieldNameMap["description"] = "Description"
    fields["category_name"] = bindings.NewStringType()
    fieldNameMap["category_name"] = "CategoryName"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.cis.tagging.tag.host.summary",fields, reflect.TypeOf(HostSummary{}), fieldNameMap, validators)
}


