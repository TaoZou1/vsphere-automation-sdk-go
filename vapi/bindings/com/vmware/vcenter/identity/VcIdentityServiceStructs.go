/* Copyright © 2019 VMware, Inc. All Rights Reserved.
     SPDX-License-Identifier: BSD-2-Clause */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Data type definitions file for service: VcIdentity.
 * Includes binding types of a structures and enumerations defined in the service.
 * Shared by client-side stubs and server-side skeletons to ensure type
 * compatibility.
 */

package identity

import (
    "reflect"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/bindings/com/vmware/vcenter/certificate_management"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/bindings"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/data"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/protocol"
)




// The ``UpdateSpec`` class contains data to update the local vCenter identity. **Warning:** This class is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
 type VcIdentityUpdateSpec struct {
    // Label of the trust. A non-unique, user-readable label. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    Label *string
    // Type of trust, used to distinguish between different types of domains such as cloud or on-prem. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    Type_ *string
    // Signing certificate chain which will replace existing chain. The existing chain will continue to be used to validate tokens for at least the maximum lifetime of the tokens. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    SigningCertChain *certificate_management.X509CertChain
    // Response URL for web logons. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    Oauth2ResponseEndpoint *string
}


// TODO some error doesn't have ErrorType. Resolve this later VAPI-3009
//


// The ``Info`` class contains data that represents a local vCenter identity. **Warning:** This class is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
 type VcIdentityInfo struct {
    // Domain identifier. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    Domain string
    // label of for the local trust. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    Label string
    // Type of trust. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    Type_ string
    // UPN Suffixes. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    UpnSuffixes []string
    // List of signing certificate chains. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    SigningCertChains []certificate_management.X509CertChain
    // Response URL for web logons. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    Oauth2ResponseEndpoint *string
}


// TODO some error doesn't have ErrorType. Resolve this later VAPI-3009
//





func vcIdentityGetInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func vcIdentityGetOutputType() bindings.BindingType {
    return bindings.NewReferenceType(VcIdentityInfoBindingType)
}

func vcIdentityGetRestMetadata() protocol.OperationRestMetadata {
    paramsTypeMap := map[string]bindings.BindingType{}
    pathParams := map[string]string{}
    queryParams := map[string]string{}
    headerParams := map[string]string{}
    resultHeaders := map[string]string{}
    errorHeaders := map[string]string{}
    return protocol.NewOperationRestMetadata(
      paramsTypeMap,
      pathParams,
      queryParams,
      headerParams,
      "",
      "null",
      "",
       resultHeaders,
       0,
       errorHeaders,
       map[string]int{"Unauthorized": 403})
}


func vcIdentityUpdateInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["spec"] = bindings.NewReferenceType(VcIdentityUpdateSpecBindingType)
    fieldNameMap["spec"] = "Spec"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func vcIdentityUpdateOutputType() bindings.BindingType {
    return bindings.NewVoidType()
}

func vcIdentityUpdateRestMetadata() protocol.OperationRestMetadata {
    paramsTypeMap := map[string]bindings.BindingType{}
    pathParams := map[string]string{}
    queryParams := map[string]string{}
    headerParams := map[string]string{}
    resultHeaders := map[string]string{}
    errorHeaders := map[string]string{}
    return protocol.NewOperationRestMetadata(
      paramsTypeMap,
      pathParams,
      queryParams,
      headerParams,
      "",
      "null",
      "",
       resultHeaders,
       0,
       errorHeaders,
       map[string]int{"Unauthorized": 403,"InvalidArgument": 400})
}



func VcIdentityUpdateSpecBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["label"] = bindings.NewOptionalType(bindings.NewStringType())
    fieldNameMap["label"] = "Label"
    fields["type"] = bindings.NewOptionalType(bindings.NewStringType())
    fieldNameMap["type"] = "Type_"
    fields["signing_cert_chain"] = bindings.NewOptionalType(bindings.NewReferenceType(certificate_management.X509CertChainBindingType))
    fieldNameMap["signing_cert_chain"] = "SigningCertChain"
    fields["oauth2_response_endpoint"] = bindings.NewOptionalType(bindings.NewStringType())
    fieldNameMap["oauth2_response_endpoint"] = "Oauth2ResponseEndpoint"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.identity.vc_identity.update_spec",fields, reflect.TypeOf(VcIdentityUpdateSpec{}), fieldNameMap, validators)
}

func VcIdentityInfoBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["domain"] = bindings.NewStringType()
    fieldNameMap["domain"] = "Domain"
    fields["label"] = bindings.NewStringType()
    fieldNameMap["label"] = "Label"
    fields["type"] = bindings.NewStringType()
    fieldNameMap["type"] = "Type_"
    fields["upn_suffixes"] = bindings.NewListType(bindings.NewStringType(), reflect.TypeOf([]string{}))
    fieldNameMap["upn_suffixes"] = "UpnSuffixes"
    fields["signing_cert_chains"] = bindings.NewListType(bindings.NewReferenceType(certificate_management.X509CertChainBindingType), reflect.TypeOf([]certificate_management.X509CertChain{}))
    fieldNameMap["signing_cert_chains"] = "SigningCertChains"
    fields["oauth2_response_endpoint"] = bindings.NewOptionalType(bindings.NewStringType())
    fieldNameMap["oauth2_response_endpoint"] = "Oauth2ResponseEndpoint"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.identity.vc_identity.info",fields, reflect.TypeOf(VcIdentityInfo{}), fieldNameMap, validators)
}


