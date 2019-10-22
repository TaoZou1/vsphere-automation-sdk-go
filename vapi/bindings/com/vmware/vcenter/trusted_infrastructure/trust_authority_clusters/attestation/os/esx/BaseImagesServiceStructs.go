/* Copyright © 2019 VMware, Inc. All Rights Reserved.
     SPDX-License-Identifier: BSD-2-Clause */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Data type definitions file for service: BaseImages.
 * Includes binding types of a structures and enumerations defined in the service.
 * Shared by client-side stubs and server-side skeletons to ensure type
 * compatibility.
 */

package esx

import (
    "reflect"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/bindings/com/vmware/vapi/std"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/bindings"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/data"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/protocol"
)

// Resource type for ESX base image.
const BaseImages_RESOURCE_TYPE = "com.vmware.vcenter.trusted_platform.trusted_clusters.attestation.os.esx.BaseImage"


// The ``Health`` enumeration class is indicator for the consistency of the hosts status in the cluster.
//
// <p> See {@link com.vmware.vapi.bindings.ApiEnumeration enumerated types description}.
 
type BaseImages_Health string

const (
    // No status available.
     BaseImages_Health_NONE BaseImages_Health = "NONE"
    // Each host in the cluster is in consistent state with the rest hosts in the cluster.
     BaseImages_Health_OK BaseImages_Health = "OK"
    // Attestation is funtioning, however there is an issue that requires attention.
     BaseImages_Health_WARNING BaseImages_Health = "WARNING"
    // Not all hosts in the cluster are in consistent state.
     BaseImages_Health_ERROR BaseImages_Health = "ERROR"
)

func (h BaseImages_Health) BaseImages_Health() bool {
    switch h {
        case BaseImages_Health_NONE:
            return true
        case BaseImages_Health_OK:
            return true
        case BaseImages_Health_WARNING:
            return true
        case BaseImages_Health_ERROR:
            return true
        default:
            return false
    }
}





// The ``Summary`` class contains information that summarizes an ESX base image.
 type BaseImagesSummary struct {
    // A unique ESX version number.
    Version string
    // A unique ESX version formatted for display.
    DisplayName string
    // A health indicator which indicates whether each host in the cluster has this version of the ESX base image.
    Health BaseImages_Health
}


// TODO some error doesn't have ErrorType. Resolve this later VAPI-3009
//


// The ``Info`` class contains information that describes an ESX base image.
 type BaseImagesInfo struct {
    // A unique ESX version formatted for display.
    DisplayName string
    // A health indicator which indicates whether each host in the cluster has this version of the ESX base image.
    Health BaseImages_Health
    // Details regarding the health. 
//
//  When the ``Health`` is not BaseImages_Health#BaseImagesHealth_OK or BaseImages_Health#BaseImagesHealth_NONE, this member will provide an actionable description of the issues present.
    Details []std.LocalizableMessage
}


// TODO some error doesn't have ErrorType. Resolve this later VAPI-3009
//


// The ``FilterSpec`` class contains the data necessary for identifying a Trust Authority Host in a cluster.
 type BaseImagesFilterSpec struct {
    // Search criteria by ESX base image version numbers.
    Version map[string]bool
    // Search criteria by ESX base image version version numbers.
    DisplayName map[string]bool
    // Search criteria by health indicator.
    Health map[BaseImages_Health]bool
}


// TODO some error doesn't have ErrorType. Resolve this later VAPI-3009
//





func baseImagesImportFromImgdbInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["cluster"] = bindings.NewIdType([]string {"ClusterComputeResource"}, "")
    fields["imgdb"] = bindings.NewBlobType()
    fieldNameMap["cluster"] = "Cluster"
    fieldNameMap["imgdb"] = "Imgdb"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func baseImagesImportFromImgdbOutputType() bindings.BindingType {
    return bindings.NewIdType([]string {"com.vmware.vcenter.trusted_platform.trusted_clusters.attestation.os.esx.BaseImage"}, "")
}

func baseImagesImportFromImgdbRestMetadata() protocol.OperationRestMetadata {
    paramsTypeMap := map[string]bindings.BindingType{}
    pathParams := map[string]string{}
    queryParams := map[string]string{}
    headerParams := map[string]string{}
    paramsTypeMap["cluster"] = bindings.NewIdType([]string {"ClusterComputeResource"}, "")
    paramsTypeMap["imgdb"] = bindings.NewBlobType()
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
      "imgdb",
      "POST",
      "/vcenter/trusted-infrastructure/trust-authority-clusters/{cluster}/attestation/os/esx/base-images",
       resultHeaders,
       202,
       errorHeaders,
       map[string]int{"AlreadyExists": 400,"Error": 500,"InvalidArgument": 400,"NotFound": 404,"Unauthenticated": 401})
}


func baseImagesListInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["cluster"] = bindings.NewIdType([]string {"ClusterComputeResource"}, "")
    fields["spec"] = bindings.NewOptionalType(bindings.NewReferenceType(BaseImagesFilterSpecBindingType))
    fieldNameMap["cluster"] = "Cluster"
    fieldNameMap["spec"] = "Spec"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func baseImagesListOutputType() bindings.BindingType {
    return bindings.NewListType(bindings.NewReferenceType(BaseImagesSummaryBindingType), reflect.TypeOf([]BaseImagesSummary{}))
}

func baseImagesListRestMetadata() protocol.OperationRestMetadata {
    paramsTypeMap := map[string]bindings.BindingType{}
    pathParams := map[string]string{}
    queryParams := map[string]string{}
    headerParams := map[string]string{}
    paramsTypeMap["cluster"] = bindings.NewIdType([]string {"ClusterComputeResource"}, "")
    paramsTypeMap["spec.health"] = bindings.NewOptionalType(bindings.NewSetType(bindings.NewEnumType("com.vmware.vcenter.trusted_infrastructure.trust_authority_clusters.attestation.os.esx.base_images.health", reflect.TypeOf(BaseImages_Health(BaseImages_Health_NONE))), reflect.TypeOf(map[BaseImages_Health]bool{})))
    paramsTypeMap["spec.display_name"] = bindings.NewOptionalType(bindings.NewSetType(bindings.NewStringType(), reflect.TypeOf(map[string]bool{})))
    paramsTypeMap["spec.version"] = bindings.NewOptionalType(bindings.NewSetType(bindings.NewIdType([]string {"com.vmware.vcenter.trusted_platform.trusted_clusters.attestation.os.esx.BaseImage"}, ""), reflect.TypeOf(map[string]bool{})))
    paramsTypeMap["cluster"] = bindings.NewIdType([]string {"ClusterComputeResource"}, "")
    pathParams["cluster"] = "cluster"
    queryParams["spec.health"] = "health"
    queryParams["spec.display_name"] = "display_name"
    queryParams["spec.version"] = "version"
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
      "/vcenter/trusted-infrastructure/trust-authority-clusters/{cluster}/attestation/os/esx/base-images",
       resultHeaders,
       202,
       errorHeaders,
       map[string]int{"Error": 500,"InvalidArgument": 400,"NotFound": 404,"Unauthenticated": 401})
}


func baseImagesDeleteInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["cluster"] = bindings.NewIdType([]string {"ClusterComputeResource"}, "")
    fields["version"] = bindings.NewIdType([]string {"com.vmware.vcenter.trusted_platform.trusted_clusters.attestation.os.esx.BaseImage"}, "")
    fieldNameMap["cluster"] = "Cluster"
    fieldNameMap["version"] = "Version"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func baseImagesDeleteOutputType() bindings.BindingType {
    return bindings.NewVoidType()
}

func baseImagesDeleteRestMetadata() protocol.OperationRestMetadata {
    paramsTypeMap := map[string]bindings.BindingType{}
    pathParams := map[string]string{}
    queryParams := map[string]string{}
    headerParams := map[string]string{}
    paramsTypeMap["cluster"] = bindings.NewIdType([]string {"ClusterComputeResource"}, "")
    paramsTypeMap["version"] = bindings.NewIdType([]string {"com.vmware.vcenter.trusted_platform.trusted_clusters.attestation.os.esx.BaseImage"}, "")
    paramsTypeMap["cluster"] = bindings.NewIdType([]string {"ClusterComputeResource"}, "")
    paramsTypeMap["version"] = bindings.NewIdType([]string {"com.vmware.vcenter.trusted_platform.trusted_clusters.attestation.os.esx.BaseImage"}, "")
    pathParams["cluster"] = "cluster"
    pathParams["version"] = "version"
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
      "/vcenter/trusted-infrastructure/trust-authority-clusters/{cluster}/attestation/os/esx/base-images/{version}",
       resultHeaders,
       202,
       errorHeaders,
       map[string]int{"Error": 500,"InvalidArgument": 400,"NotFound": 404,"Unauthenticated": 401})
}


func baseImagesGetInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["cluster"] = bindings.NewIdType([]string {"ClusterComputeResource"}, "")
    fields["version"] = bindings.NewIdType([]string {"com.vmware.vcenter.trusted_platform.trusted_clusters.attestation.os.esx.BaseImage"}, "")
    fieldNameMap["cluster"] = "Cluster"
    fieldNameMap["version"] = "Version"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func baseImagesGetOutputType() bindings.BindingType {
    return bindings.NewReferenceType(BaseImagesInfoBindingType)
}

func baseImagesGetRestMetadata() protocol.OperationRestMetadata {
    paramsTypeMap := map[string]bindings.BindingType{}
    pathParams := map[string]string{}
    queryParams := map[string]string{}
    headerParams := map[string]string{}
    paramsTypeMap["cluster"] = bindings.NewIdType([]string {"ClusterComputeResource"}, "")
    paramsTypeMap["version"] = bindings.NewIdType([]string {"com.vmware.vcenter.trusted_platform.trusted_clusters.attestation.os.esx.BaseImage"}, "")
    paramsTypeMap["cluster"] = bindings.NewIdType([]string {"ClusterComputeResource"}, "")
    paramsTypeMap["version"] = bindings.NewIdType([]string {"com.vmware.vcenter.trusted_platform.trusted_clusters.attestation.os.esx.BaseImage"}, "")
    pathParams["cluster"] = "cluster"
    pathParams["version"] = "version"
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
      "/vcenter/trusted-infrastructure/trust-authority-clusters/{cluster}/attestation/os/esx/base-images/{version}",
       resultHeaders,
       202,
       errorHeaders,
       map[string]int{"Error": 500,"InvalidArgument": 400,"NotFound": 404,"Unauthenticated": 401})
}



func BaseImagesSummaryBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["version"] = bindings.NewIdType([]string {"com.vmware.vcenter.trusted_platform.trusted_clusters.attestation.os.esx.BaseImage"}, "")
    fieldNameMap["version"] = "Version"
    fields["display_name"] = bindings.NewStringType()
    fieldNameMap["display_name"] = "DisplayName"
    fields["health"] = bindings.NewEnumType("com.vmware.vcenter.trusted_infrastructure.trust_authority_clusters.attestation.os.esx.base_images.health", reflect.TypeOf(BaseImages_Health(BaseImages_Health_NONE)))
    fieldNameMap["health"] = "Health"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.trusted_infrastructure.trust_authority_clusters.attestation.os.esx.base_images.summary",fields, reflect.TypeOf(BaseImagesSummary{}), fieldNameMap, validators)
}

func BaseImagesInfoBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["display_name"] = bindings.NewStringType()
    fieldNameMap["display_name"] = "DisplayName"
    fields["health"] = bindings.NewEnumType("com.vmware.vcenter.trusted_infrastructure.trust_authority_clusters.attestation.os.esx.base_images.health", reflect.TypeOf(BaseImages_Health(BaseImages_Health_NONE)))
    fieldNameMap["health"] = "Health"
    fields["details"] = bindings.NewListType(bindings.NewReferenceType(std.LocalizableMessageBindingType), reflect.TypeOf([]std.LocalizableMessage{}))
    fieldNameMap["details"] = "Details"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.trusted_infrastructure.trust_authority_clusters.attestation.os.esx.base_images.info",fields, reflect.TypeOf(BaseImagesInfo{}), fieldNameMap, validators)
}

func BaseImagesFilterSpecBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["version"] = bindings.NewOptionalType(bindings.NewSetType(bindings.NewIdType([]string {"com.vmware.vcenter.trusted_platform.trusted_clusters.attestation.os.esx.BaseImage"}, ""), reflect.TypeOf(map[string]bool{})))
    fieldNameMap["version"] = "Version"
    fields["display_name"] = bindings.NewOptionalType(bindings.NewSetType(bindings.NewStringType(), reflect.TypeOf(map[string]bool{})))
    fieldNameMap["display_name"] = "DisplayName"
    fields["health"] = bindings.NewOptionalType(bindings.NewSetType(bindings.NewEnumType("com.vmware.vcenter.trusted_infrastructure.trust_authority_clusters.attestation.os.esx.base_images.health", reflect.TypeOf(BaseImages_Health(BaseImages_Health_NONE))), reflect.TypeOf(map[BaseImages_Health]bool{})))
    fieldNameMap["health"] = "Health"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.trusted_infrastructure.trust_authority_clusters.attestation.os.esx.base_images.filter_spec",fields, reflect.TypeOf(BaseImagesFilterSpec{}), fieldNameMap, validators)
}


