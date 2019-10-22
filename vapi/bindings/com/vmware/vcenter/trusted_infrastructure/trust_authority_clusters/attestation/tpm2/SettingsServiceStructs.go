/* Copyright © 2019 VMware, Inc. All Rights Reserved.
     SPDX-License-Identifier: BSD-2-Clause */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Data type definitions file for service: Settings.
 * Includes binding types of a structures and enumerations defined in the service.
 * Shared by client-side stubs and server-side skeletons to ensure type
 * compatibility.
 */

package tpm2

import (
    "reflect"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/bindings/com/vmware/vapi/std"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/bindings"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/data"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/protocol"
)



// The ``Health`` enumeration class is indicator for the consistency of the hosts status in the cluster.
//
// <p> See {@link com.vmware.vapi.bindings.ApiEnumeration enumerated types description}.
 
type Settings_Health string

const (
    // No status available.
     Settings_Health_NONE Settings_Health = "NONE"
    // Each host in the cluster is in consistent state with the rest hosts in the cluster.
     Settings_Health_OK Settings_Health = "OK"
    // Attestation is functioning, however there is an issue that requires attention.
     Settings_Health_WARNING Settings_Health = "WARNING"
    // Not all hosts in the cluster are in consistent state.
     Settings_Health_ERROR Settings_Health = "ERROR"
)

func (h Settings_Health) Settings_Health() bool {
    switch h {
        case Settings_Health_NONE:
            return true
        case Settings_Health_OK:
            return true
        case Settings_Health_WARNING:
            return true
        case Settings_Health_ERROR:
            return true
        default:
            return false
    }
}





// The ``Info`` class contains information that describes the TPM 2.0 protocol settings.
 type SettingsInfo struct {
    // Require registered TPM endorsement keys. 
//
//  During attestation, the attested host will always send its endorsement key to the Attestation Service. With this option is set, the Attestation Service will only proceed with attestation if the endorsement key has been added to the list of configured trusted endorsement keys.
    RequireEndorsementKeys bool
    // Require TPM endorsement key certificate validation. 
//
//  During attestation, the attested host will send its endorsement key certificate if one is available. With this option set, the Attestation Service will validate the endorsement key certificate against the list of configured trusted TPM CA certificates. Only endorsement key certificates that are signed by a trusted TPM CA certificate will be able to successfully attest.
    RequireCertificateValidation bool
    // A health indicator which indicates whether each host in the cluster has the same attestation settings.
    Health Settings_Health
    // Details regarding the health. 
//
//  When the ``Health`` is not Settings_Health#SettingsHealth_OK or Settings_Health#SettingsHealth_NONE, this member will provide an actionable description of the issues present.
    Details []std.LocalizableMessage
}


// TODO some error doesn't have ErrorType. Resolve this later VAPI-3009
//


// The ``UpdateSpec`` class contains information that describes changes to the TPM 2.0 protocol settings.
 type SettingsUpdateSpec struct {
    // Require registered TPM endorsement keys.
    RequireEndorsementKeys *bool
    // Require TPM endorsement key certificate validation.
    RequireCertificateValidation *bool
}


// TODO some error doesn't have ErrorType. Resolve this later VAPI-3009
//





func settingsGetInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["cluster"] = bindings.NewIdType([]string {"ClusterComputeResource"}, "")
    fieldNameMap["cluster"] = "Cluster"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func settingsGetOutputType() bindings.BindingType {
    return bindings.NewReferenceType(SettingsInfoBindingType)
}

func settingsGetRestMetadata() protocol.OperationRestMetadata {
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
      "/vcenter/trusted-infrastructure/trust-authority-clusters/{cluster}/attestation/tpm2/settings",
       resultHeaders,
       202,
       errorHeaders,
       map[string]int{"Error": 500,"InvalidArgument": 400,"NotFound": 404,"Unauthenticated": 401})
}


func settingsUpdateInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["cluster"] = bindings.NewIdType([]string {"ClusterComputeResource"}, "")
    fields["spec"] = bindings.NewReferenceType(SettingsUpdateSpecBindingType)
    fieldNameMap["cluster"] = "Cluster"
    fieldNameMap["spec"] = "Spec"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func settingsUpdateOutputType() bindings.BindingType {
    return bindings.NewVoidType()
}

func settingsUpdateRestMetadata() protocol.OperationRestMetadata {
    paramsTypeMap := map[string]bindings.BindingType{}
    pathParams := map[string]string{}
    queryParams := map[string]string{}
    headerParams := map[string]string{}
    paramsTypeMap["cluster"] = bindings.NewIdType([]string {"ClusterComputeResource"}, "")
    paramsTypeMap["spec"] = bindings.NewReferenceType(SettingsUpdateSpecBindingType)
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
      "/vcenter/trusted-infrastructure/trust-authority-clusters/{cluster}/attestation/tpm2/settings",
       resultHeaders,
       202,
       errorHeaders,
       map[string]int{"Error": 500,"InvalidArgument": 400,"NotFound": 404,"Unauthenticated": 401})
}



func SettingsInfoBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["require_endorsement_keys"] = bindings.NewBooleanType()
    fieldNameMap["require_endorsement_keys"] = "RequireEndorsementKeys"
    fields["require_certificate_validation"] = bindings.NewBooleanType()
    fieldNameMap["require_certificate_validation"] = "RequireCertificateValidation"
    fields["health"] = bindings.NewEnumType("com.vmware.vcenter.trusted_infrastructure.trust_authority_clusters.attestation.tpm2.settings.health", reflect.TypeOf(Settings_Health(Settings_Health_NONE)))
    fieldNameMap["health"] = "Health"
    fields["details"] = bindings.NewListType(bindings.NewReferenceType(std.LocalizableMessageBindingType), reflect.TypeOf([]std.LocalizableMessage{}))
    fieldNameMap["details"] = "Details"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.trusted_infrastructure.trust_authority_clusters.attestation.tpm2.settings.info",fields, reflect.TypeOf(SettingsInfo{}), fieldNameMap, validators)
}

func SettingsUpdateSpecBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["require_endorsement_keys"] = bindings.NewOptionalType(bindings.NewBooleanType())
    fieldNameMap["require_endorsement_keys"] = "RequireEndorsementKeys"
    fields["require_certificate_validation"] = bindings.NewOptionalType(bindings.NewBooleanType())
    fieldNameMap["require_certificate_validation"] = "RequireCertificateValidation"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.trusted_infrastructure.trust_authority_clusters.attestation.tpm2.settings.update_spec",fields, reflect.TypeOf(SettingsUpdateSpec{}), fieldNameMap, validators)
}


