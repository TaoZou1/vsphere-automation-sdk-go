/* Copyright © 2019 VMware, Inc. All Rights Reserved.
     SPDX-License-Identifier: BSD-2-Clause */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Data type definitions file for service: Csr.
 * Includes binding types of a structures and enumerations defined in the service.
 * Shared by client-side stubs and server-side skeletons to ensure type
 * compatibility.
 */

package client_certificate

import (
    "reflect"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/bindings"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/data"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/protocol"
)




// The ``Info`` class contains the certificate signing request.
 type CsrInfo struct {
    // Certificate signing request.
    Csr string
}



func (CsrInfo CsrInfo) Error() string {
    return "com.vmware.vcenter.trusted_infrastructure.trust_authority_clusters.kms.providers.client_certificate.info"
}






func csrCreateInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["cluster"] = bindings.NewIdType([]string {"ClusterComputeResource"}, "")
    fields["provider"] = bindings.NewIdType([]string {"com.vmware.vcenter.trusted_platform.trusted_clusters.kms.Provider"}, "")
    fieldNameMap["cluster"] = "Cluster"
    fieldNameMap["provider"] = "Provider"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func csrCreateOutputType() bindings.BindingType {
    return bindings.NewReferenceType(CsrInfoBindingType)
}

func csrCreateRestMetadata() protocol.OperationRestMetadata {
    paramsTypeMap := map[string]bindings.BindingType{}
    pathParams := map[string]string{}
    queryParams := map[string]string{}
    headerParams := map[string]string{}
    paramsTypeMap["cluster"] = bindings.NewIdType([]string {"ClusterComputeResource"}, "")
    paramsTypeMap["provider"] = bindings.NewIdType([]string {"com.vmware.vcenter.trusted_platform.trusted_clusters.kms.Provider"}, "")
    paramsTypeMap["cluster"] = bindings.NewIdType([]string {"ClusterComputeResource"}, "")
    paramsTypeMap["provider"] = bindings.NewIdType([]string {"com.vmware.vcenter.trusted_platform.trusted_clusters.kms.Provider"}, "")
    pathParams["cluster"] = "cluster"
    pathParams["provider"] = "provider"
    resultHeaders := map[string]string{}
    errorHeaders := map[string]string{}
    errorHeaders["Unauthenticated.challenge"] = "WWW-Authenticate"
    return protocol.NewOperationRestMetadata(
      paramsTypeMap,
      pathParams,
      queryParams,
      headerParams,
      "",
      "POST",
      "/vcenter/trusted-infrastructure/trust-authority-clusters/{cluster}/kms/providers/{provider}/client-certificate/csr",
       resultHeaders,
       202,
       errorHeaders,
       map[string]int{"InvalidArgument": 400,"NotFound": 404,"Unauthenticated": 401,"Error": 500})
}


func csrGetInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["cluster"] = bindings.NewIdType([]string {"ClusterComputeResource"}, "")
    fields["provider"] = bindings.NewIdType([]string {"com.vmware.vcenter.trusted_platform.trusted_clusters.kms.Provider"}, "")
    fieldNameMap["cluster"] = "Cluster"
    fieldNameMap["provider"] = "Provider"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func csrGetOutputType() bindings.BindingType {
    return bindings.NewReferenceType(CsrInfoBindingType)
}

func csrGetRestMetadata() protocol.OperationRestMetadata {
    paramsTypeMap := map[string]bindings.BindingType{}
    pathParams := map[string]string{}
    queryParams := map[string]string{}
    headerParams := map[string]string{}
    paramsTypeMap["cluster"] = bindings.NewIdType([]string {"ClusterComputeResource"}, "")
    paramsTypeMap["provider"] = bindings.NewIdType([]string {"com.vmware.vcenter.trusted_platform.trusted_clusters.kms.Provider"}, "")
    paramsTypeMap["cluster"] = bindings.NewIdType([]string {"ClusterComputeResource"}, "")
    paramsTypeMap["provider"] = bindings.NewIdType([]string {"com.vmware.vcenter.trusted_platform.trusted_clusters.kms.Provider"}, "")
    pathParams["cluster"] = "cluster"
    pathParams["provider"] = "provider"
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
      "/vcenter/trusted-infrastructure/trust-authority-clusters/{cluster}/kms/providers/{provider}/client-certificate/csr",
       resultHeaders,
       202,
       errorHeaders,
       map[string]int{"InvalidArgument": 400,"NotFound": 404,"Unauthenticated": 401,"Error": 500})
}



func CsrInfoBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["csr"] = bindings.NewStringType()
    fieldNameMap["csr"] = "Csr"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.trusted_infrastructure.trust_authority_clusters.kms.providers.client_certificate.csr.info",fields, reflect.TypeOf(CsrInfo{}), fieldNameMap, validators)
}


