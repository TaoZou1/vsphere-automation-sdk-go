/* Copyright © 2019 VMware, Inc. All Rights Reserved.
     SPDX-License-Identifier: BSD-2-Clause */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Data type definitions file for service: Networking.
 * Includes binding types of a structures and enumerations defined in the service.
 * Shared by client-side stubs and server-side skeletons to ensure type
 * compatibility.
 */

package guest

import (
    "reflect"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/bindings"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/data"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/protocol"
)




// The ``Info`` class contains information about networking as configured in the guest operating system.
 type NetworkingInfo struct {
    // Client DNS values. Data assigned by DNS.
    DnsValues *DnsAssignedValues
    // Client DNS configuration. How DNS queries are resolved.
    Dns *DnsConfigInfo
}


// TODO some error doesn't have ErrorType. Resolve this later VAPI-3009
//





func networkingGetInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["vm"] = bindings.NewIdType([]string {"VirtualMachine"}, "")
    fieldNameMap["vm"] = "Vm"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func networkingGetOutputType() bindings.BindingType {
    return bindings.NewReferenceType(NetworkingInfoBindingType)
}

func networkingGetRestMetadata() protocol.OperationRestMetadata {
    paramsTypeMap := map[string]bindings.BindingType{}
    pathParams := map[string]string{}
    queryParams := map[string]string{}
    headerParams := map[string]string{}
    paramsTypeMap["vm"] = bindings.NewIdType([]string {"VirtualMachine"}, "")
    paramsTypeMap["vm"] = bindings.NewIdType([]string {"VirtualMachine"}, "")
    pathParams["vm"] = "vm"
    resultHeaders := map[string]string{}
    errorHeaders := map[string]string{}
    return protocol.NewOperationRestMetadata(
      paramsTypeMap,
      pathParams,
      queryParams,
      headerParams,
      "",
      "GET",
      "/vcenter/vm/{vm}/guest/networking",
       resultHeaders,
       200,
       errorHeaders,
       map[string]int{"Error": 500,"NotFound": 404,"ServiceUnavailable": 503})
}



func NetworkingInfoBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["dns_values"] = bindings.NewOptionalType(bindings.NewReferenceType(DnsAssignedValuesBindingType))
    fieldNameMap["dns_values"] = "DnsValues"
    fields["dns"] = bindings.NewOptionalType(bindings.NewReferenceType(DnsConfigInfoBindingType))
    fieldNameMap["dns"] = "Dns"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.vm.guest.networking.info",fields, reflect.TypeOf(NetworkingInfo{}), fieldNameMap, validators)
}


