/* Copyright © 2019 VMware, Inc. All Rights Reserved.
     SPDX-License-Identifier: BSD-2-Clause */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Data type definitions file for service: Capabilities.
 * Includes binding types of a structures and enumerations defined in the service.
 * Shared by client-side stubs and server-side skeletons to ensure type
 * compatibility.
 */

package vcha

import (
    "reflect"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/bindings"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/data"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/protocol"
)





// The ``Info`` class contains information about the capabilities of the active node of a VCHA configuration. This class was added in vSphere API 6.7.1.
 type CapabilitiesInfo struct {
    // Are the active's management vCenter server credentials required. This property was added in vSphere API 6.7.1.
    ManagementVcenterCredentialsRequired bool
    // Does the vCenter have the appropriate license for VCHA. This property was added in vSphere API 6.7.1.
    VchaLicensed bool
    // Is the assignment mode of failover IP of the active node in the VCHA cluster valid. This property was added in vSphere API 6.7.1.
    FailoverIpModeValid bool
}









func capabilitiesGetInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func capabilitiesGetOutputType() bindings.BindingType {
    return bindings.NewReferenceType(CapabilitiesInfoBindingType)
}

func capabilitiesGetRestMetadata() protocol.OperationRestMetadata {
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



func CapabilitiesInfoBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["management_vcenter_credentials_required"] = bindings.NewBooleanType()
    fieldNameMap["management_vcenter_credentials_required"] = "ManagementVcenterCredentialsRequired"
    fields["vcha_licensed"] = bindings.NewBooleanType()
    fieldNameMap["vcha_licensed"] = "VchaLicensed"
    fields["failover_ip_mode_valid"] = bindings.NewBooleanType()
    fieldNameMap["failover_ip_mode_valid"] = "FailoverIpModeValid"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.vcha.capabilities.info",fields, reflect.TypeOf(CapabilitiesInfo{}), fieldNameMap, validators)
}


