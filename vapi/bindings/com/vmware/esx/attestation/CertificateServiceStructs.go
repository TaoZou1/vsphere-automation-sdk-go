/* Copyright © 2019 VMware, Inc. All Rights Reserved.
     SPDX-License-Identifier: BSD-2-Clause */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Data type definitions file for service: Certificate.
 * Includes binding types of a structures and enumerations defined in the service.
 * Shared by client-side stubs and server-side skeletons to ensure type
 * compatibility.
 */

package attestation

import (
    "reflect"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/bindings"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/data"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/protocol"
)




// The ``Info`` class contains information describing the certificate that is used for attestation report signing.
 type CertificateInfo struct {
    // The attestation report signing certificate in PEM format.
    Certificate string
}


// TODO some error doesn't have ErrorType. Resolve this later VAPI-3009
//


// The ``SetSpec`` class contains information that describes a certificate that will replace the existing attestation report signing certificate.
 type CertificateSetSpec struct {
    // The attestation report signing certificate in PEM format. 
//
//  This certificate must be the result of processing the CSR that is returned by Csr#create, or Csr#get.
    Certificate string
}


// TODO some error doesn't have ErrorType. Resolve this later VAPI-3009
//





func certificateGetInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func certificateGetOutputType() bindings.BindingType {
    return bindings.NewReferenceType(CertificateInfoBindingType)
}

func certificateGetRestMetadata() protocol.OperationRestMetadata {
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
      "/esx/attestation/certificate",
       resultHeaders,
       200,
       errorHeaders,
       map[string]int{"Error": 500,"Unauthenticated": 401,"Unauthorized": 403})
}


func certificateSetInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["spec"] = bindings.NewReferenceType(CertificateSetSpecBindingType)
    fieldNameMap["spec"] = "Spec"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func certificateSetOutputType() bindings.BindingType {
    return bindings.NewVoidType()
}

func certificateSetRestMetadata() protocol.OperationRestMetadata {
    paramsTypeMap := map[string]bindings.BindingType{}
    pathParams := map[string]string{}
    queryParams := map[string]string{}
    headerParams := map[string]string{}
    paramsTypeMap["spec"] = bindings.NewReferenceType(CertificateSetSpecBindingType)
    resultHeaders := map[string]string{}
    errorHeaders := map[string]string{}
    errorHeaders["Unauthenticated.challenge"] = "WWW-Authenticate"
    return protocol.NewOperationRestMetadata(
      paramsTypeMap,
      pathParams,
      queryParams,
      headerParams,
      "spec",
      "PUT",
      "/esx/attestation/certificate",
       resultHeaders,
       204,
       errorHeaders,
       map[string]int{"Error": 500,"InvalidArgument": 400,"Unauthenticated": 401,"Unauthorized": 403})
}



func CertificateInfoBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["certificate"] = bindings.NewStringType()
    fieldNameMap["certificate"] = "Certificate"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.esx.attestation.certificate.info",fields, reflect.TypeOf(CertificateInfo{}), fieldNameMap, validators)
}

func CertificateSetSpecBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["certificate"] = bindings.NewStringType()
    fieldNameMap["certificate"] = "Certificate"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.esx.attestation.certificate.set_spec",fields, reflect.TypeOf(CertificateSetSpec{}), fieldNameMap, validators)
}


