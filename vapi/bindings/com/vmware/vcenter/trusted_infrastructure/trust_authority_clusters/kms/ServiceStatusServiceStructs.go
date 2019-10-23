/* Copyright © 2019 VMware, Inc. All Rights Reserved.
     SPDX-License-Identifier: BSD-2-Clause */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Data type definitions file for service: ServiceStatus.
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



// The ``Health`` enumeration class defines the possible service health states.
//
// <p> See {@link com.vmware.vapi.bindings.ApiEnumeration enumerated types description}.
 
type ServiceStatus_Health string

const (
    // No status available.
     ServiceStatus_Health_NONE ServiceStatus_Health = "NONE"
    // Service is functioning normally.
     ServiceStatus_Health_OK ServiceStatus_Health = "OK"
    // Service is functioning, however there is an issue that requires attention.
     ServiceStatus_Health_WARNING ServiceStatus_Health = "WARNING"
    // Service is not functioning.
     ServiceStatus_Health_ERROR ServiceStatus_Health = "ERROR"
)

func (h ServiceStatus_Health) ServiceStatus_Health() bool {
    switch h {
        case ServiceStatus_Health_NONE:
            return true
        case ServiceStatus_Health_OK:
            return true
        case ServiceStatus_Health_WARNING:
            return true
        case ServiceStatus_Health_ERROR:
            return true
        default:
            return false
    }
}





// The ``Info`` class contains information that describes the status of the service.
 type ServiceStatusInfo struct {
    // The service health status.
    Health ServiceStatus_Health
    // Details regarding the health of the service. 
    //
    //  When the service ``Health`` is not ServiceStatus_Health#ServiceStatusHealth_OK or ServiceStatus_Health#ServiceStatusHealth_NONE, this member will provide an actionable description of the issues present.
    Details []std.LocalizableMessage
}



func (ServiceStatusInfo ServiceStatusInfo) Error() string {
    return "com.vmware.vcenter.trusted_infrastructure.trust_authority_clusters.kms.info"
}






func serviceStatusGetInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["cluster"] = bindings.NewIdType([]string {"ClusterComputeResource"}, "")
    fieldNameMap["cluster"] = "Cluster"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func serviceStatusGetOutputType() bindings.BindingType {
    return bindings.NewReferenceType(ServiceStatusInfoBindingType)
}

func serviceStatusGetRestMetadata() protocol.OperationRestMetadata {
    paramsTypeMap := map[string]bindings.BindingType{}
    pathParams := map[string]string{}
    queryParams := map[string]string{}
    headerParams := map[string]string{}
    paramsTypeMap["cluster"] = bindings.NewIdType([]string {"ClusterComputeResource"}, "")
    paramsTypeMap["cluster"] = bindings.NewIdType([]string {"ClusterComputeResource"}, "")
    pathParams["cluster"] = "cluster"
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
      "/vcenter/trusted-infrastructure/trust-authority-clusters/{cluster}/kms/service-status",
       resultHeaders,
       202,
       errorHeaders,
       map[string]int{"Error": 500,"InvalidArgument": 400,"NotFound": 404,"Unauthenticated": 401})
}



func ServiceStatusInfoBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["health"] = bindings.NewEnumType("com.vmware.vcenter.trusted_infrastructure.trust_authority_clusters.kms.service_status.health", reflect.TypeOf(ServiceStatus_Health(ServiceStatus_Health_NONE)))
    fieldNameMap["health"] = "Health"
    fields["details"] = bindings.NewListType(bindings.NewReferenceType(std.LocalizableMessageBindingType), reflect.TypeOf([]std.LocalizableMessage{}))
    fieldNameMap["details"] = "Details"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.trusted_infrastructure.trust_authority_clusters.kms.service_status.info",fields, reflect.TypeOf(ServiceStatusInfo{}), fieldNameMap, validators)
}


