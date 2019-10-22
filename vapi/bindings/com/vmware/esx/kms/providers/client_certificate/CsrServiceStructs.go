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




// The ``Info`` class contains information for a CSR.
 type CsrInfo struct {
    // Certificate Signing Request. 
//
//  PEM format.
    Csr string
    // Private key. 
//
//  PEM format.
    PrivateKey *string
}


// TODO some error doesn't have ErrorType. Resolve this later VAPI-3009
//


// The ``GetSpec`` class contains properties that describe the specifications for getting private key along with CSR.
 type CsrGetSpec struct {
    // Include private key. 
    IncludePrivateKey *bool
}


// TODO some error doesn't have ErrorType. Resolve this later VAPI-3009
//


// The ``SetSpec`` class contains properties that describe the specifications for setting CSR and private key.
 type CsrSetSpec struct {
    // Certificate Signing Request. 
//
//  PEM format.
    Csr string
    // Private key. 
//
//  PEM format.
    PrivateKey string
}


// TODO some error doesn't have ErrorType. Resolve this later VAPI-3009
//





func csrCreateInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["provider"] = bindings.NewIdType([]string {"com.vmware.esx.kms.providers"}, "")
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
    paramsTypeMap["provider"] = bindings.NewIdType([]string {"com.vmware.esx.kms.providers"}, "")
    paramsTypeMap["provider"] = bindings.NewIdType([]string {"com.vmware.esx.kms.providers"}, "")
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
      "/esx/kms/providers/{provider}/client-certificate/csr",
       resultHeaders,
       201,
       errorHeaders,
       map[string]int{"InvalidArgument": 400,"NotFound": 404,"Unauthenticated": 401,"Unauthorized": 403,"Error": 500})
}


func csrGetInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["provider"] = bindings.NewIdType([]string {"com.vmware.esx.kms.providers"}, "")
    fields["spec"] = bindings.NewOptionalType(bindings.NewReferenceType(CsrGetSpecBindingType))
    fieldNameMap["provider"] = "Provider"
    fieldNameMap["spec"] = "Spec"
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
    paramsTypeMap["provider"] = bindings.NewIdType([]string {"com.vmware.esx.kms.providers"}, "")
    paramsTypeMap["spec.include_private_key"] = bindings.NewOptionalType(bindings.NewBooleanType())
    paramsTypeMap["provider"] = bindings.NewIdType([]string {"com.vmware.esx.kms.providers"}, "")
    pathParams["provider"] = "provider"
    queryParams["spec.include_private_key"] = "include_private_key"
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
      "/esx/kms/providers/{provider}/client-certificate/csr",
       resultHeaders,
       200,
       errorHeaders,
       map[string]int{"InvalidArgument": 400,"NotFound": 404,"Unauthenticated": 401,"Unauthorized": 403,"Error": 500})
}


func csrSetInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["provider"] = bindings.NewIdType([]string {"com.vmware.esx.kms.providers"}, "")
    fields["spec"] = bindings.NewReferenceType(CsrSetSpecBindingType)
    fieldNameMap["provider"] = "Provider"
    fieldNameMap["spec"] = "Spec"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func csrSetOutputType() bindings.BindingType {
    return bindings.NewVoidType()
}

func csrSetRestMetadata() protocol.OperationRestMetadata {
    paramsTypeMap := map[string]bindings.BindingType{}
    pathParams := map[string]string{}
    queryParams := map[string]string{}
    headerParams := map[string]string{}
    paramsTypeMap["provider"] = bindings.NewIdType([]string {"com.vmware.esx.kms.providers"}, "")
    paramsTypeMap["spec"] = bindings.NewReferenceType(CsrSetSpecBindingType)
    paramsTypeMap["provider"] = bindings.NewIdType([]string {"com.vmware.esx.kms.providers"}, "")
    pathParams["provider"] = "provider"
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
      "/esx/kms/providers/{provider}/client-certificate/csr",
       resultHeaders,
       204,
       errorHeaders,
       map[string]int{"InvalidArgument": 400,"NotFound": 404,"Unauthenticated": 401,"Unauthorized": 403,"Error": 500})
}



func CsrInfoBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["csr"] = bindings.NewStringType()
    fieldNameMap["csr"] = "Csr"
    fields["private_key"] = bindings.NewOptionalType(bindings.NewSecretType())
    fieldNameMap["private_key"] = "PrivateKey"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.esx.kms.providers.client_certificate.csr.info",fields, reflect.TypeOf(CsrInfo{}), fieldNameMap, validators)
}

func CsrGetSpecBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["include_private_key"] = bindings.NewOptionalType(bindings.NewBooleanType())
    fieldNameMap["include_private_key"] = "IncludePrivateKey"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.esx.kms.providers.client_certificate.csr.get_spec",fields, reflect.TypeOf(CsrGetSpec{}), fieldNameMap, validators)
}

func CsrSetSpecBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["csr"] = bindings.NewStringType()
    fieldNameMap["csr"] = "Csr"
    fields["private_key"] = bindings.NewSecretType()
    fieldNameMap["private_key"] = "PrivateKey"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.esx.kms.providers.client_certificate.csr.set_spec",fields, reflect.TypeOf(CsrSetSpec{}), fieldNameMap, validators)
}


