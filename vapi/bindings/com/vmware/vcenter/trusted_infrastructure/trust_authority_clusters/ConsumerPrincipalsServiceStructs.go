/* Copyright © 2019 VMware, Inc. All Rights Reserved.
     SPDX-License-Identifier: BSD-2-Clause */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Data type definitions file for service: ConsumerPrincipals.
 * Includes binding types of a structures and enumerations defined in the service.
 * Shared by client-side stubs and server-side skeletons to ensure type
 * compatibility.
 */

package trust_authority_clusters

import (
    "reflect"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/bindings/com/vmware/vapi/std"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/bindings/com/vmware/vcenter/trusted_infrastructure"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/bindings"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/data"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/protocol"
)



// The ``Health`` enumeration class defines the possible health states. This enumeration was added in vSphere API 7.0.
//
// <p> See {@link com.vmware.vapi.bindings.ApiEnumeration enumerated types description}.
 
type ConsumerPrincipals_Health string

const (
    // None. No status available. This constant field was added in vSphere API 7.0.
     ConsumerPrincipals_Health_NONE ConsumerPrincipals_Health = "NONE"
    // OK. Health is normal. This constant field was added in vSphere API 7.0.
     ConsumerPrincipals_Health_OK ConsumerPrincipals_Health = "OK"
    // Warning. Health is normal, however there is an issue that requires attention. This constant field was added in vSphere API 7.0.
     ConsumerPrincipals_Health_WARNING ConsumerPrincipals_Health = "WARNING"
    // Error. Not healthy. This constant field was added in vSphere API 7.0.
     ConsumerPrincipals_Health_ERROR ConsumerPrincipals_Health = "ERROR"
)

func (h ConsumerPrincipals_Health) ConsumerPrincipals_Health() bool {
    switch h {
        case ConsumerPrincipals_Health_NONE:
            return true
        case ConsumerPrincipals_Health_OK:
            return true
        case ConsumerPrincipals_Health_WARNING:
            return true
        case ConsumerPrincipals_Health_ERROR:
            return true
        default:
            return false
    }
}





// The ``CreateSpec`` class contains the information necessary to establish trust between a workload vCenter and a Trust Authority Host. This class was added in vSphere API 7.0.
 type ConsumerPrincipalsCreateSpec struct {
    // The certificates used by the vCenter STS to sign tokens. This property was added in vSphere API 7.0.
    Certificates []trusted_infrastructure.X509CertChain
    // A user-friendly alias of the service which created and signed the security token. This property was added in vSphere API 7.0.
    IssuerAlias string
    // The service which created and signed the security token. This property was added in vSphere API 7.0.
    Issuer string
    // The principal used by the vCenter to retrieve tokens. This property was added in vSphere API 7.0.
    Principal trusted_infrastructure.StsPrincipal
}



//


// The ``FilterSpec`` class contains data which identifies a connection profile on the trusted vCenter. This class was added in vSphere API 7.0.
 type ConsumerPrincipalsFilterSpec struct {
    // The unqiue identifier of a connection profile. This property was added in vSphere API 7.0.
    Id map[string]bool
    // The principal used by the vCenter to retrieve tokens. This property was added in vSphere API 7.0.
    Principals []trusted_infrastructure.StsPrincipal
    // The service which created and signed the security token. This property was added in vSphere API 7.0.
    Issuer map[string]bool
}



//


// The ``Info`` class contains the information necessary to establish trust between a workload vCenter and a Trust Authority Host. This class was added in vSphere API 7.0.
 type ConsumerPrincipalsInfo struct {
    // The unqiue identifier of a connection profile. This property was added in vSphere API 7.0.
    Id string
    // The principal used by the vCenter to retrieve tokens. Currently this is the vCenter solution user. This property was added in vSphere API 7.0.
    Principal trusted_infrastructure.StsPrincipal
    // A user-friendly alias of the service which created and signed the security token. This property was added in vSphere API 7.0.
    IssuerAlias string
    // The service which created and signed the security token. This property was added in vSphere API 7.0.
    Issuer string
    // The certificates used by the vCenter STS to sign tokens. This property was added in vSphere API 7.0.
    Certificates []trusted_infrastructure.X509CertChain
    // The consistency of the profile across the hosts in the cluster. This property was added in vSphere API 7.0.
    Health ConsumerPrincipals_Health
    // A localizable message describing the health of the profile. This property was added in vSphere API 7.0.
    Message *std.LocalizableMessage
}



//


// The ``Summary`` class contains a summary of the information necessary to establish trust between a workload vCenter and a Trust Authority Host. This class was added in vSphere API 7.0.
 type ConsumerPrincipalsSummary struct {
    // The unqiue identifier of a connection profile. This property was added in vSphere API 7.0.
    Id string
    // The principal used by the vCenter to retrieve tokens. Currently this is the vCenter solution user. This property was added in vSphere API 7.0.
    Principal trusted_infrastructure.StsPrincipal
    // A user-friendly alias of the service which created and signed the security token. This property was added in vSphere API 7.0.
    IssuerAlias string
    // The service which created and signed the security token. This property was added in vSphere API 7.0.
    Issuer string
}



//





func consumerPrincipalsCreateInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["cluster"] = bindings.NewIdType([]string {"ClusterComputeResource"}, "")
    fields["spec"] = bindings.NewReferenceType(ConsumerPrincipalsCreateSpecBindingType)
    fieldNameMap["cluster"] = "Cluster"
    fieldNameMap["spec"] = "Spec"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func consumerPrincipalsCreateOutputType() bindings.BindingType {
    return bindings.NewIdType([]string {"com.vmware.esx.authentication.clientprofile"}, "")
}

func consumerPrincipalsCreateRestMetadata() protocol.OperationRestMetadata {
    paramsTypeMap := map[string]bindings.BindingType{}
    pathParams := map[string]string{}
    queryParams := map[string]string{}
    headerParams := map[string]string{}
    paramsTypeMap["cluster"] = bindings.NewIdType([]string {"ClusterComputeResource"}, "")
    paramsTypeMap["spec"] = bindings.NewReferenceType(ConsumerPrincipalsCreateSpecBindingType)
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
      "POST",
      "/vcenter/trusted-infrastructure/trust-authority-clusters/{cluster}/consumer-principals",
       resultHeaders,
       202,
       errorHeaders,
       map[string]int{"AlreadyExists": 400,"Error": 500,"NotFound": 404,"Unauthenticated": 401})
}


func consumerPrincipalsDeleteInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["cluster"] = bindings.NewIdType([]string {"ClusterComputeResource"}, "")
    fields["profile"] = bindings.NewIdType([]string {"com.vmware.esx.authentication.clientprofile"}, "")
    fieldNameMap["cluster"] = "Cluster"
    fieldNameMap["profile"] = "Profile"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func consumerPrincipalsDeleteOutputType() bindings.BindingType {
    return bindings.NewVoidType()
}

func consumerPrincipalsDeleteRestMetadata() protocol.OperationRestMetadata {
    paramsTypeMap := map[string]bindings.BindingType{}
    pathParams := map[string]string{}
    queryParams := map[string]string{}
    headerParams := map[string]string{}
    paramsTypeMap["cluster"] = bindings.NewIdType([]string {"ClusterComputeResource"}, "")
    paramsTypeMap["profile"] = bindings.NewIdType([]string {"com.vmware.esx.authentication.clientprofile"}, "")
    paramsTypeMap["cluster"] = bindings.NewIdType([]string {"ClusterComputeResource"}, "")
    paramsTypeMap["profile"] = bindings.NewIdType([]string {"com.vmware.esx.authentication.clientprofile"}, "")
    pathParams["cluster"] = "cluster"
    pathParams["profile"] = "profile"
    resultHeaders := map[string]string{}
    errorHeaders := map[string]string{}
    errorHeaders["Unauthenticated.challenge"] = "WWW-Authenticate"
    return protocol.NewOperationRestMetadata(
      paramsTypeMap,
      pathParams,
      queryParams,
      headerParams,
      "",
      "DELETE",
      "/vcenter/trusted-infrastructure/trust-authority-clusters/{cluster}/consumer-principals/{profile}",
       resultHeaders,
       202,
       errorHeaders,
       map[string]int{"Error": 500,"NotFound": 404,"Unauthenticated": 401})
}


func consumerPrincipalsGetInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["cluster"] = bindings.NewIdType([]string {"ClusterComputeResource"}, "")
    fields["profile"] = bindings.NewIdType([]string {"com.vmware.esx.authentication.clientprofile"}, "")
    fieldNameMap["cluster"] = "Cluster"
    fieldNameMap["profile"] = "Profile"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func consumerPrincipalsGetOutputType() bindings.BindingType {
    return bindings.NewReferenceType(ConsumerPrincipalsInfoBindingType)
}

func consumerPrincipalsGetRestMetadata() protocol.OperationRestMetadata {
    paramsTypeMap := map[string]bindings.BindingType{}
    pathParams := map[string]string{}
    queryParams := map[string]string{}
    headerParams := map[string]string{}
    paramsTypeMap["cluster"] = bindings.NewIdType([]string {"ClusterComputeResource"}, "")
    paramsTypeMap["profile"] = bindings.NewIdType([]string {"com.vmware.esx.authentication.clientprofile"}, "")
    paramsTypeMap["cluster"] = bindings.NewIdType([]string {"ClusterComputeResource"}, "")
    paramsTypeMap["profile"] = bindings.NewIdType([]string {"com.vmware.esx.authentication.clientprofile"}, "")
    pathParams["cluster"] = "cluster"
    pathParams["profile"] = "profile"
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
      "/vcenter/trusted-infrastructure/trust-authority-clusters/{cluster}/consumer-principals/{profile}",
       resultHeaders,
       202,
       errorHeaders,
       map[string]int{"Error": 500,"NotFound": 404,"Unauthenticated": 401})
}


func consumerPrincipalsListInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["cluster"] = bindings.NewIdType([]string {"ClusterComputeResource"}, "")
    fields["spec"] = bindings.NewOptionalType(bindings.NewReferenceType(ConsumerPrincipalsFilterSpecBindingType))
    fieldNameMap["cluster"] = "Cluster"
    fieldNameMap["spec"] = "Spec"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func consumerPrincipalsListOutputType() bindings.BindingType {
    return bindings.NewListType(bindings.NewReferenceType(ConsumerPrincipalsSummaryBindingType), reflect.TypeOf([]ConsumerPrincipalsSummary{}))
}

func consumerPrincipalsListRestMetadata() protocol.OperationRestMetadata {
    paramsTypeMap := map[string]bindings.BindingType{}
    pathParams := map[string]string{}
    queryParams := map[string]string{}
    headerParams := map[string]string{}
    paramsTypeMap["cluster"] = bindings.NewIdType([]string {"ClusterComputeResource"}, "")
    paramsTypeMap["spec"] = bindings.NewOptionalType(bindings.NewReferenceType(ConsumerPrincipalsFilterSpecBindingType))
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
      "POST",
      "/vcenter/trusted-infrastructure/trust-authority-clusters/{cluster}/consumer-principals",
       resultHeaders,
       202,
       errorHeaders,
       map[string]int{"Error": 500,"NotFound": 404,"Unauthenticated": 401})
}



func ConsumerPrincipalsCreateSpecBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["certificates"] = bindings.NewListType(bindings.NewReferenceType(trusted_infrastructure.X509CertChainBindingType), reflect.TypeOf([]trusted_infrastructure.X509CertChain{}))
    fieldNameMap["certificates"] = "Certificates"
    fields["issuer_alias"] = bindings.NewIdType([]string {"com.vmware.esx.authentication.trust.security-token-issuer"}, "")
    fieldNameMap["issuer_alias"] = "IssuerAlias"
    fields["issuer"] = bindings.NewStringType()
    fieldNameMap["issuer"] = "Issuer"
    fields["principal"] = bindings.NewReferenceType(trusted_infrastructure.StsPrincipalBindingType)
    fieldNameMap["principal"] = "Principal"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.trusted_infrastructure.trust_authority_clusters.consumer_principals.create_spec",fields, reflect.TypeOf(ConsumerPrincipalsCreateSpec{}), fieldNameMap, validators)
}

func ConsumerPrincipalsFilterSpecBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["id"] = bindings.NewOptionalType(bindings.NewSetType(bindings.NewIdType([]string {"com.vmware.esx.authentication.clientprofile"}, ""), reflect.TypeOf(map[string]bool{})))
    fieldNameMap["id"] = "Id"
    fields["principals"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewReferenceType(trusted_infrastructure.StsPrincipalBindingType), reflect.TypeOf([]trusted_infrastructure.StsPrincipal{})))
    fieldNameMap["principals"] = "Principals"
    fields["issuer"] = bindings.NewOptionalType(bindings.NewSetType(bindings.NewIdType([]string {"com.vmware.esx.authentication.trust.security-token-issuer"}, ""), reflect.TypeOf(map[string]bool{})))
    fieldNameMap["issuer"] = "Issuer"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.trusted_infrastructure.trust_authority_clusters.consumer_principals.filter_spec",fields, reflect.TypeOf(ConsumerPrincipalsFilterSpec{}), fieldNameMap, validators)
}

func ConsumerPrincipalsInfoBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["id"] = bindings.NewIdType([]string {"com.vmware.esx.authentication.clientprofile"}, "")
    fieldNameMap["id"] = "Id"
    fields["principal"] = bindings.NewReferenceType(trusted_infrastructure.StsPrincipalBindingType)
    fieldNameMap["principal"] = "Principal"
    fields["issuer_alias"] = bindings.NewIdType([]string {"com.vmware.esx.authentication.trust.security-token-issuer"}, "")
    fieldNameMap["issuer_alias"] = "IssuerAlias"
    fields["issuer"] = bindings.NewStringType()
    fieldNameMap["issuer"] = "Issuer"
    fields["certificates"] = bindings.NewListType(bindings.NewReferenceType(trusted_infrastructure.X509CertChainBindingType), reflect.TypeOf([]trusted_infrastructure.X509CertChain{}))
    fieldNameMap["certificates"] = "Certificates"
    fields["health"] = bindings.NewEnumType("com.vmware.vcenter.trusted_infrastructure.trust_authority_clusters.consumer_principals.health", reflect.TypeOf(ConsumerPrincipals_Health(ConsumerPrincipals_Health_NONE)))
    fieldNameMap["health"] = "Health"
    fields["message"] = bindings.NewOptionalType(bindings.NewReferenceType(std.LocalizableMessageBindingType))
    fieldNameMap["message"] = "Message"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.trusted_infrastructure.trust_authority_clusters.consumer_principals.info",fields, reflect.TypeOf(ConsumerPrincipalsInfo{}), fieldNameMap, validators)
}

func ConsumerPrincipalsSummaryBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["id"] = bindings.NewIdType([]string {"com.vmware.esx.authentication.clientprofile"}, "")
    fieldNameMap["id"] = "Id"
    fields["principal"] = bindings.NewReferenceType(trusted_infrastructure.StsPrincipalBindingType)
    fieldNameMap["principal"] = "Principal"
    fields["issuer_alias"] = bindings.NewIdType([]string {"com.vmware.esx.authentication.trust.security-token-issuer"}, "")
    fieldNameMap["issuer_alias"] = "IssuerAlias"
    fields["issuer"] = bindings.NewStringType()
    fieldNameMap["issuer"] = "Issuer"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.trusted_infrastructure.trust_authority_clusters.consumer_principals.summary",fields, reflect.TypeOf(ConsumerPrincipalsSummary{}), fieldNameMap, validators)
}


