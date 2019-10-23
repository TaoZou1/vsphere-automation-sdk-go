/* Copyright © 2019 VMware, Inc. All Rights Reserved.
     SPDX-License-Identifier: BSD-2-Clause */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Data type definitions file for service: ClientCertificate.
 * Includes binding types of a structures and enumerations defined in the service.
 * Shared by client-side stubs and server-side skeletons to ensure type
 * compatibility.
 */

package providers

import (
    "reflect"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/bindings"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/data"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/protocol"
)




// The ``Info`` class contains information for a certificate.
 type ClientCertificateInfo struct {
    // Certificate. 
    //
    //  PEM format.
    Certificate string
    // Private key. 
    //
    //  PEM format.
    PrivateKey *string
}



func (ClientCertificateInfo ClientCertificateInfo) Error() string {
    return "com.vmware.esx.kms.providers.info"
}



// The ``GetSpec`` class contains properties that describe the specifications for getting private key along with client certificate.
 type ClientCertificateGetSpec struct {
    // Include private key. 
    IncludePrivateKey *bool
}



func (ClientCertificateGetSpec ClientCertificateGetSpec) Error() string {
    return "com.vmware.esx.kms.providers.get_spec"
}



// The ``SetSpec`` class contains properties that describe the specifications for setting certificate and private key.
 type ClientCertificateSetSpec struct {
    // Certificate. 
    //
    //  PEM format.
    Certificate string
    // Private key. 
    PrivateKey *string
}



func (ClientCertificateSetSpec ClientCertificateSetSpec) Error() string {
    return "com.vmware.esx.kms.providers.set_spec"
}






func clientCertificateCreateInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["provider"] = bindings.NewIdType([]string {"com.vmware.esx.kms.providers"}, "")
    fieldNameMap["provider"] = "Provider"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func clientCertificateCreateOutputType() bindings.BindingType {
    return bindings.NewVoidType()
}

func clientCertificateCreateRestMetadata() protocol.OperationRestMetadata {
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
      "/esx/kms/providers/{provider}/client-certificate",
       resultHeaders,
       204,
       errorHeaders,
       map[string]int{"InvalidArgument": 400,"NotFound": 404,"Unauthenticated": 401,"Unauthorized": 403,"Error": 500})
}


func clientCertificateGetInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["provider"] = bindings.NewIdType([]string {"com.vmware.esx.kms.providers"}, "")
    fields["spec"] = bindings.NewOptionalType(bindings.NewReferenceType(ClientCertificateGetSpecBindingType))
    fieldNameMap["provider"] = "Provider"
    fieldNameMap["spec"] = "Spec"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func clientCertificateGetOutputType() bindings.BindingType {
    return bindings.NewReferenceType(ClientCertificateInfoBindingType)
}

func clientCertificateGetRestMetadata() protocol.OperationRestMetadata {
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
      "/esx/kms/providers/{provider}/client-certificate",
       resultHeaders,
       200,
       errorHeaders,
       map[string]int{"InvalidArgument": 400,"NotFound": 404,"Unauthenticated": 401,"Unauthorized": 403,"Error": 500})
}


func clientCertificateSetInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["provider"] = bindings.NewIdType([]string {"com.vmware.esx.kms.providers"}, "")
    fields["spec"] = bindings.NewReferenceType(ClientCertificateSetSpecBindingType)
    fieldNameMap["provider"] = "Provider"
    fieldNameMap["spec"] = "Spec"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func clientCertificateSetOutputType() bindings.BindingType {
    return bindings.NewVoidType()
}

func clientCertificateSetRestMetadata() protocol.OperationRestMetadata {
    paramsTypeMap := map[string]bindings.BindingType{}
    pathParams := map[string]string{}
    queryParams := map[string]string{}
    headerParams := map[string]string{}
    paramsTypeMap["provider"] = bindings.NewIdType([]string {"com.vmware.esx.kms.providers"}, "")
    paramsTypeMap["spec"] = bindings.NewReferenceType(ClientCertificateSetSpecBindingType)
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
      "/esx/kms/providers/{provider}/client-certificate",
       resultHeaders,
       204,
       errorHeaders,
       map[string]int{"InvalidArgument": 400,"NotFound": 404,"Unauthenticated": 401,"Unauthorized": 403,"Error": 500})
}



func ClientCertificateInfoBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["certificate"] = bindings.NewStringType()
    fieldNameMap["certificate"] = "Certificate"
    fields["private_key"] = bindings.NewOptionalType(bindings.NewSecretType())
    fieldNameMap["private_key"] = "PrivateKey"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.esx.kms.providers.client_certificate.info",fields, reflect.TypeOf(ClientCertificateInfo{}), fieldNameMap, validators)
}

func ClientCertificateGetSpecBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["include_private_key"] = bindings.NewOptionalType(bindings.NewBooleanType())
    fieldNameMap["include_private_key"] = "IncludePrivateKey"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.esx.kms.providers.client_certificate.get_spec",fields, reflect.TypeOf(ClientCertificateGetSpec{}), fieldNameMap, validators)
}

func ClientCertificateSetSpecBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["certificate"] = bindings.NewStringType()
    fieldNameMap["certificate"] = "Certificate"
    fields["private_key"] = bindings.NewOptionalType(bindings.NewSecretType())
    fieldNameMap["private_key"] = "PrivateKey"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.esx.kms.providers.client_certificate.set_spec",fields, reflect.TypeOf(ClientCertificateSetSpec{}), fieldNameMap, validators)
}


