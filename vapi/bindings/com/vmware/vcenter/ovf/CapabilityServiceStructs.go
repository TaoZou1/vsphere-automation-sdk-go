/* Copyright © 2019 VMware, Inc. All Rights Reserved.
     SPDX-License-Identifier: BSD-2-Clause */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Data type definitions file for service: Capability.
 * Includes binding types of a structures and enumerations defined in the service.
 * Shared by client-side stubs and server-side skeletons to ensure type
 * compatibility.
 */

package ovf

import (
    "reflect"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/bindings"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/data"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/protocol"
)




// The ``Info`` class describes OVF capabilities supported by the server.
 type CapabilityCapabilityInfo struct {
    // Boolean flag to show whether or not import OVA template is supported.
    ImportOva bool
    // Boolean flag to show whether or not export OVA template is supported.
    ExportOva bool
}



//





func capabilityGetInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["server_guid"] = bindings.NewIdType([]string {"com.vmware.vcenter.VCenter"}, "")
    fieldNameMap["server_guid"] = "ServerGuid"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func capabilityGetOutputType() bindings.BindingType {
    return bindings.NewReferenceType(CapabilityCapabilityInfoBindingType)
}

func capabilityGetRestMetadata() protocol.OperationRestMetadata {
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



func CapabilityCapabilityInfoBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["import_ova"] = bindings.NewBooleanType()
    fieldNameMap["import_ova"] = "ImportOva"
    fields["export_ova"] = bindings.NewBooleanType()
    fieldNameMap["export_ova"] = "ExportOva"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.ovf.capability.capability_info",fields, reflect.TypeOf(CapabilityCapabilityInfo{}), fieldNameMap, validators)
}


