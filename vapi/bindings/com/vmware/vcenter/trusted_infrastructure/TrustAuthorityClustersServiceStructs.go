/* Copyright © 2019 VMware, Inc. All Rights Reserved.
     SPDX-License-Identifier: BSD-2-Clause */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Data type definitions file for service: TrustAuthorityClusters.
 * Includes binding types of a structures and enumerations defined in the service.
 * Shared by client-side stubs and server-side skeletons to ensure type
 * compatibility.
 */

package trusted_infrastructure

import (
    "reflect"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/bindings"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/data"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/protocol"
)



// The ``State`` enumeration class defines the states of the TrustAuthorityClusters.
//
// <p> See {@link com.vmware.vapi.bindings.ApiEnumeration enumerated types description}.
 
type TrustAuthorityClusters_State string

const (
    // The TrustAuthorityClusters is enabled.
     TrustAuthorityClusters_State_ENABLE TrustAuthorityClusters_State = "ENABLE"
    // The TrustAuthorityClusters is disabled.
     TrustAuthorityClusters_State_DISABLE TrustAuthorityClusters_State = "DISABLE"
)

func (s TrustAuthorityClusters_State) TrustAuthorityClusters_State() bool {
    switch s {
        case TrustAuthorityClusters_State_ENABLE:
            return true
        case TrustAuthorityClusters_State_DISABLE:
            return true
        default:
            return false
    }
}





// The ``Summary`` class contains information about TrustAuthorityClusters id and state.
 type TrustAuthorityClustersSummary struct {
    // Identifies the cluster.
    Cluster string
    // The state of the TrustAuthorityClusters.
    State TrustAuthorityClusters_State
}



func (TrustAuthorityClustersSummary TrustAuthorityClustersSummary) Error() string {
    return "com.vmware.vcenter.trusted_infrastructure.summary"
}



// The ``FilterSpec`` class contains the data necessary for identifying a TrustAuthorityClusters.
 type TrustAuthorityClustersFilterSpec struct {
    // Identifies the cluster.
    Cluster map[string]bool
    // The state of the TrustAuthorityClusters.
    State map[TrustAuthorityClusters_State]bool
}



func (TrustAuthorityClustersFilterSpec TrustAuthorityClustersFilterSpec) Error() string {
    return "com.vmware.vcenter.trusted_infrastructure.filter_spec"
}



// The ``UpdateSpec`` class contains the data necessary for update of a TrustAuthorityClusters.
 type TrustAuthorityClustersUpdateSpec struct {
    // The state of the TrustAuthorityClusters.
    State *TrustAuthorityClusters_State
}



func (TrustAuthorityClustersUpdateSpec TrustAuthorityClustersUpdateSpec) Error() string {
    return "com.vmware.vcenter.trusted_infrastructure.update_spec"
}



// The ``Info`` class contains the data necessary for retrieving the TrustAuthorityClusters info.
 type TrustAuthorityClustersInfo struct {
    // Identifies the cluster.
    Cluster string
    // The state of the cluster.
    State TrustAuthorityClusters_State
}



func (TrustAuthorityClustersInfo TrustAuthorityClustersInfo) Error() string {
    return "com.vmware.vcenter.trusted_infrastructure.info"
}






func trustAuthorityClustersUpdateInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["cluster"] = bindings.NewIdType([]string {"ClusterComputeResource"}, "")
    fields["spec"] = bindings.NewReferenceType(TrustAuthorityClustersUpdateSpecBindingType)
    fieldNameMap["cluster"] = "Cluster"
    fieldNameMap["spec"] = "Spec"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func trustAuthorityClustersUpdateOutputType() bindings.BindingType {
    return bindings.NewVoidType()
}

func trustAuthorityClustersUpdateRestMetadata() protocol.OperationRestMetadata {
    paramsTypeMap := map[string]bindings.BindingType{}
    pathParams := map[string]string{}
    queryParams := map[string]string{}
    headerParams := map[string]string{}
    paramsTypeMap["cluster"] = bindings.NewIdType([]string {"ClusterComputeResource"}, "")
    paramsTypeMap["spec"] = bindings.NewReferenceType(TrustAuthorityClustersUpdateSpecBindingType)
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
      "spec",
      "PATCH",
      "/vcenter/trusted-infrastructure/trust-authority-clusters/{cluster}",
       resultHeaders,
       202,
       errorHeaders,
       map[string]int{"Error": 500,"NotFound": 404,"Unauthenticated": 401})
}


func trustAuthorityClustersGetInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["cluster"] = bindings.NewIdType([]string {"ClusterComputeResource"}, "")
    fieldNameMap["cluster"] = "Cluster"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func trustAuthorityClustersGetOutputType() bindings.BindingType {
    return bindings.NewReferenceType(TrustAuthorityClustersInfoBindingType)
}

func trustAuthorityClustersGetRestMetadata() protocol.OperationRestMetadata {
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
      "/vcenter/trusted-infrastructure/trust-authority-clusters/{cluster}",
       resultHeaders,
       200,
       errorHeaders,
       map[string]int{"Error": 500,"NotFound": 404,"Unauthenticated": 401})
}


func trustAuthorityClustersListInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["spec"] = bindings.NewOptionalType(bindings.NewReferenceType(TrustAuthorityClustersFilterSpecBindingType))
    fieldNameMap["spec"] = "Spec"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func trustAuthorityClustersListOutputType() bindings.BindingType {
    return bindings.NewListType(bindings.NewReferenceType(TrustAuthorityClustersSummaryBindingType), reflect.TypeOf([]TrustAuthorityClustersSummary{}))
}

func trustAuthorityClustersListRestMetadata() protocol.OperationRestMetadata {
    paramsTypeMap := map[string]bindings.BindingType{}
    pathParams := map[string]string{}
    queryParams := map[string]string{}
    headerParams := map[string]string{}
    paramsTypeMap["spec.cluster"] = bindings.NewOptionalType(bindings.NewSetType(bindings.NewIdType([]string {"ClusterComputeResource"}, ""), reflect.TypeOf(map[string]bool{})))
    paramsTypeMap["spec.state"] = bindings.NewOptionalType(bindings.NewSetType(bindings.NewEnumType("com.vmware.vcenter.trusted_infrastructure.trust_authority_clusters.state", reflect.TypeOf(TrustAuthorityClusters_State(TrustAuthorityClusters_State_ENABLE))), reflect.TypeOf(map[TrustAuthorityClusters_State]bool{})))
    queryParams["spec.cluster"] = "cluster"
    queryParams["spec.state"] = "state"
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
      "/vcenter/trusted-infrastructure/trust-authority-clusters",
       resultHeaders,
       200,
       errorHeaders,
       map[string]int{"Error": 500,"Unauthenticated": 401})
}



func TrustAuthorityClustersSummaryBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["cluster"] = bindings.NewIdType([]string {"ClusterComputeResource"}, "")
    fieldNameMap["cluster"] = "Cluster"
    fields["state"] = bindings.NewEnumType("com.vmware.vcenter.trusted_infrastructure.trust_authority_clusters.state", reflect.TypeOf(TrustAuthorityClusters_State(TrustAuthorityClusters_State_ENABLE)))
    fieldNameMap["state"] = "State"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.trusted_infrastructure.trust_authority_clusters.summary",fields, reflect.TypeOf(TrustAuthorityClustersSummary{}), fieldNameMap, validators)
}

func TrustAuthorityClustersFilterSpecBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["cluster"] = bindings.NewOptionalType(bindings.NewSetType(bindings.NewIdType([]string {"ClusterComputeResource"}, ""), reflect.TypeOf(map[string]bool{})))
    fieldNameMap["cluster"] = "Cluster"
    fields["state"] = bindings.NewOptionalType(bindings.NewSetType(bindings.NewEnumType("com.vmware.vcenter.trusted_infrastructure.trust_authority_clusters.state", reflect.TypeOf(TrustAuthorityClusters_State(TrustAuthorityClusters_State_ENABLE))), reflect.TypeOf(map[TrustAuthorityClusters_State]bool{})))
    fieldNameMap["state"] = "State"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.trusted_infrastructure.trust_authority_clusters.filter_spec",fields, reflect.TypeOf(TrustAuthorityClustersFilterSpec{}), fieldNameMap, validators)
}

func TrustAuthorityClustersUpdateSpecBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["state"] = bindings.NewOptionalType(bindings.NewEnumType("com.vmware.vcenter.trusted_infrastructure.trust_authority_clusters.state", reflect.TypeOf(TrustAuthorityClusters_State(TrustAuthorityClusters_State_ENABLE))))
    fieldNameMap["state"] = "State"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.trusted_infrastructure.trust_authority_clusters.update_spec",fields, reflect.TypeOf(TrustAuthorityClustersUpdateSpec{}), fieldNameMap, validators)
}

func TrustAuthorityClustersInfoBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["cluster"] = bindings.NewIdType([]string {"ClusterComputeResource"}, "")
    fieldNameMap["cluster"] = "Cluster"
    fields["state"] = bindings.NewEnumType("com.vmware.vcenter.trusted_infrastructure.trust_authority_clusters.state", reflect.TypeOf(TrustAuthorityClusters_State(TrustAuthorityClusters_State_ENABLE)))
    fieldNameMap["state"] = "State"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.trusted_infrastructure.trust_authority_clusters.info",fields, reflect.TypeOf(TrustAuthorityClustersInfo{}), fieldNameMap, validators)
}


