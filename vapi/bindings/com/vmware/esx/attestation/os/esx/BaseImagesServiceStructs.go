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
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/bindings"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/data"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/protocol"
)

// Resource type for ESX base image.
const BaseImages_RESOURCE_TYPE = "com.vmware.esx.attestation.os.esx.base_images"



// The ``Summary`` class contains information that summarizes an ESX base image.
 type BaseImagesSummary struct {
    // A unique ESX version number.
    Version string
    // A unique ESX version formatted for display.
    DisplayName string
}


// TODO some error doesn't have ErrorType. Resolve this later VAPI-3009
//


// The ``Info`` class contains information that describes an ESX base image.
 type BaseImagesInfo struct {
    // A unique ESX version formatted for display.
    DisplayName string
}


// TODO some error doesn't have ErrorType. Resolve this later VAPI-3009
//





func baseImagesImportFromImgdbInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["imgdb"] = bindings.NewBlobType()
    fieldNameMap["imgdb"] = "Imgdb"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func baseImagesImportFromImgdbOutputType() bindings.BindingType {
    return bindings.NewIdType([]string {"com.vmware.esx.attestation.os.esx.base_images"}, "")
}

func baseImagesImportFromImgdbRestMetadata() protocol.OperationRestMetadata {
    paramsTypeMap := map[string]bindings.BindingType{}
    pathParams := map[string]string{}
    queryParams := map[string]string{}
    headerParams := map[string]string{}
    paramsTypeMap["imgdb"] = bindings.NewBlobType()
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
      "/esx/attestation/os/esx/base-images",
       resultHeaders,
       200,
       errorHeaders,
       map[string]int{"AlreadyExists": 400,"Error": 500,"InvalidArgument": 400,"Unauthenticated": 401,"Unauthorized": 403})
}


func baseImagesListInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
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
      "/esx/attestation/os/esx/base-images",
       resultHeaders,
       200,
       errorHeaders,
       map[string]int{"Error": 500,"Unauthenticated": 401,"Unauthorized": 403})
}


func baseImagesDeleteInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["version"] = bindings.NewIdType([]string {"com.vmware.esx.attestation.os.esx.base_images"}, "")
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
    paramsTypeMap["version"] = bindings.NewIdType([]string {"com.vmware.esx.attestation.os.esx.base_images"}, "")
    paramsTypeMap["version"] = bindings.NewIdType([]string {"com.vmware.esx.attestation.os.esx.base_images"}, "")
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
      "/esx/attestation/os/esx/base-images/{version}",
       resultHeaders,
       204,
       errorHeaders,
       map[string]int{"Error": 500,"InvalidArgument": 400,"NotFound": 404,"Unauthenticated": 401,"Unauthorized": 403})
}


func baseImagesGetInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["version"] = bindings.NewIdType([]string {"com.vmware.esx.attestation.os.esx.base_images"}, "")
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
    paramsTypeMap["version"] = bindings.NewIdType([]string {"com.vmware.esx.attestation.os.esx.base_images"}, "")
    paramsTypeMap["version"] = bindings.NewIdType([]string {"com.vmware.esx.attestation.os.esx.base_images"}, "")
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
      "/esx/attestation/os/esx/base-images/{version}",
       resultHeaders,
       200,
       errorHeaders,
       map[string]int{"Error": 500,"InvalidArgument": 400,"NotFound": 404,"Unauthenticated": 401,"Unauthorized": 403})
}



func BaseImagesSummaryBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["version"] = bindings.NewIdType([]string {"com.vmware.esx.attestation.os.esx.base_images"}, "")
    fieldNameMap["version"] = "Version"
    fields["display_name"] = bindings.NewStringType()
    fieldNameMap["display_name"] = "DisplayName"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.esx.attestation.os.esx.base_images.summary",fields, reflect.TypeOf(BaseImagesSummary{}), fieldNameMap, validators)
}

func BaseImagesInfoBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["display_name"] = bindings.NewStringType()
    fieldNameMap["display_name"] = "DisplayName"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.esx.attestation.os.esx.base_images.info",fields, reflect.TypeOf(BaseImagesInfo{}), fieldNameMap, validators)
}


