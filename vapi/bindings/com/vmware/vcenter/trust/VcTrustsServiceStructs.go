/* Copyright © 2019 VMware, Inc. All Rights Reserved.
     SPDX-License-Identifier: BSD-2-Clause */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Data type definitions file for service: VcTrusts.
 * Includes binding types of a structures and enumerations defined in the service.
 * Shared by client-side stubs and server-side skeletons to ensure type
 * compatibility.
 */

package trust

import (
    "reflect"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/bindings/com/vmware/vcenter/certificate_management"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/bindings"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/data"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/protocol"
)




// The ``CreateSpec`` class contains information to create a trust. **Warning:** This class is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
 type VcTrustsCreateSpec struct {
    // Domain identifier. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    Domain string
    // Label of the trust. A non-unique, user-readable label. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    Label string
    // Type of trust, used to distinguish between different types of domains such as cloud or on-prem. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    Type_ *string
    // List of signing certificate chains. At least one cert chain is required. Private key is not included. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    SigningCertChains []certificate_management.X509CertChain
    // Maps a group in an existing token to a list of local SSO groups in the local domain. Map key is expected to be in UPN/email address format. Example is group name and domain name separated by '\\\\@': <group-name>\\\\@<domain-name> Map value is the name of an existing group. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    GroupMap map[string][]string
    // List of the UPN suffixes used in this trust relationship. Must contain at least one entry for the trusted domain. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    UpnSuffixes []string
}


// TODO some error doesn't have ErrorType. Resolve this later VAPI-3009
//


// The ``UpdateSpec`` class contains data to update a VcTrusts spec. **Warning:** This class is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
 type VcTrustsUpdateSpec struct {
    // Label of the trust. A non-unique, user-readable label. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    Label *string
    // Type of trust, used to distinguish between different types of domains such as cloud or on-prem. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    Type_ *string
    // List of signing certificate chains which will replace existing chain. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    SigningCertChains []certificate_management.X509CertChain
    // Mapping of group in existing token to a list of local SSO groups in the local domain for new token. If empty, deletes all existing mappings. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    GroupMap map[string][]string
    // Lists the UPN suffixes used in this trust relationship. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    UpnSuffixes []string
}


// TODO some error doesn't have ErrorType. Resolve this later VAPI-3009
//


// The ``Info`` class contains data that represents a trust. **Warning:** This class is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
 type VcTrustsInfo struct {
    // Label of the trust. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    Label string
    // Type of trust. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    Type_ string
    // List of signing certificate chains. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    SigningCertChains []certificate_management.X509CertChain
    // Mapping of group in existing token to a list of local SSO groups in the local domain for new token. If empty, no group mappings have been configured. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    GroupMap map[string][]string
    // Lists the UPN suffixes used in this trust relationship. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    UpnSuffixes []string
}


// TODO some error doesn't have ErrorType. Resolve this later VAPI-3009
//


// The ``Summary`` class contains a trust summary suitable for UI presentation. **Warning:** This class is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
 type VcTrustsSummary struct {
    // Domain identifier. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    Domain string
    // Label of the trust. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    Label string
    // Type of trust. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    Type_ string
}


// TODO some error doesn't have ErrorType. Resolve this later VAPI-3009
//





func vcTrustsListInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func vcTrustsListOutputType() bindings.BindingType {
    return bindings.NewListType(bindings.NewReferenceType(VcTrustsSummaryBindingType), reflect.TypeOf([]VcTrustsSummary{}))
}

func vcTrustsListRestMetadata() protocol.OperationRestMetadata {
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


func vcTrustsCreateInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["spec"] = bindings.NewReferenceType(VcTrustsCreateSpecBindingType)
    fieldNameMap["spec"] = "Spec"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func vcTrustsCreateOutputType() bindings.BindingType {
    return bindings.NewVoidType()
}

func vcTrustsCreateRestMetadata() protocol.OperationRestMetadata {
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
       map[string]int{"Unauthorized": 403,"AlreadyExists": 400,"InvalidArgument": 400})
}


func vcTrustsGetInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["domain"] = bindings.NewStringType()
    fieldNameMap["domain"] = "Domain"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func vcTrustsGetOutputType() bindings.BindingType {
    return bindings.NewReferenceType(VcTrustsInfoBindingType)
}

func vcTrustsGetRestMetadata() protocol.OperationRestMetadata {
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
       map[string]int{"Unauthorized": 403,"NotFound": 404})
}


func vcTrustsDeleteInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["domain"] = bindings.NewStringType()
    fieldNameMap["domain"] = "Domain"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func vcTrustsDeleteOutputType() bindings.BindingType {
    return bindings.NewVoidType()
}

func vcTrustsDeleteRestMetadata() protocol.OperationRestMetadata {
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
       map[string]int{"Unauthorized": 403,"NotFound": 404,"InvalidArgument": 400})
}


func vcTrustsUpdateInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["domain"] = bindings.NewStringType()
    fields["spec"] = bindings.NewReferenceType(VcTrustsUpdateSpecBindingType)
    fieldNameMap["domain"] = "Domain"
    fieldNameMap["spec"] = "Spec"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func vcTrustsUpdateOutputType() bindings.BindingType {
    return bindings.NewVoidType()
}

func vcTrustsUpdateRestMetadata() protocol.OperationRestMetadata {
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
       map[string]int{"Unauthorized": 403,"NotFound": 404,"InvalidArgument": 400})
}



func VcTrustsCreateSpecBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["domain"] = bindings.NewStringType()
    fieldNameMap["domain"] = "Domain"
    fields["label"] = bindings.NewStringType()
    fieldNameMap["label"] = "Label"
    fields["type"] = bindings.NewOptionalType(bindings.NewStringType())
    fieldNameMap["type"] = "Type_"
    fields["signing_cert_chains"] = bindings.NewListType(bindings.NewReferenceType(certificate_management.X509CertChainBindingType), reflect.TypeOf([]certificate_management.X509CertChain{}))
    fieldNameMap["signing_cert_chains"] = "SigningCertChains"
    fields["group_map"] = bindings.NewOptionalType(bindings.NewMapType(bindings.NewStringType(), bindings.NewListType(bindings.NewStringType(), reflect.TypeOf([]string{})),reflect.TypeOf(map[string][]string{})))
    fieldNameMap["group_map"] = "GroupMap"
    fields["upn_suffixes"] = bindings.NewListType(bindings.NewStringType(), reflect.TypeOf([]string{}))
    fieldNameMap["upn_suffixes"] = "UpnSuffixes"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.trust.vc_trusts.create_spec",fields, reflect.TypeOf(VcTrustsCreateSpec{}), fieldNameMap, validators)
}

func VcTrustsUpdateSpecBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["label"] = bindings.NewOptionalType(bindings.NewStringType())
    fieldNameMap["label"] = "Label"
    fields["type"] = bindings.NewOptionalType(bindings.NewStringType())
    fieldNameMap["type"] = "Type_"
    fields["signing_cert_chains"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewReferenceType(certificate_management.X509CertChainBindingType), reflect.TypeOf([]certificate_management.X509CertChain{})))
    fieldNameMap["signing_cert_chains"] = "SigningCertChains"
    fields["group_map"] = bindings.NewOptionalType(bindings.NewMapType(bindings.NewStringType(), bindings.NewListType(bindings.NewStringType(), reflect.TypeOf([]string{})),reflect.TypeOf(map[string][]string{})))
    fieldNameMap["group_map"] = "GroupMap"
    fields["upn_suffixes"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewStringType(), reflect.TypeOf([]string{})))
    fieldNameMap["upn_suffixes"] = "UpnSuffixes"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.trust.vc_trusts.update_spec",fields, reflect.TypeOf(VcTrustsUpdateSpec{}), fieldNameMap, validators)
}

func VcTrustsInfoBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["label"] = bindings.NewStringType()
    fieldNameMap["label"] = "Label"
    fields["type"] = bindings.NewStringType()
    fieldNameMap["type"] = "Type_"
    fields["signing_cert_chains"] = bindings.NewListType(bindings.NewReferenceType(certificate_management.X509CertChainBindingType), reflect.TypeOf([]certificate_management.X509CertChain{}))
    fieldNameMap["signing_cert_chains"] = "SigningCertChains"
    fields["group_map"] = bindings.NewMapType(bindings.NewStringType(), bindings.NewListType(bindings.NewStringType(), reflect.TypeOf([]string{})),reflect.TypeOf(map[string][]string{}))
    fieldNameMap["group_map"] = "GroupMap"
    fields["upn_suffixes"] = bindings.NewListType(bindings.NewStringType(), reflect.TypeOf([]string{}))
    fieldNameMap["upn_suffixes"] = "UpnSuffixes"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.trust.vc_trusts.info",fields, reflect.TypeOf(VcTrustsInfo{}), fieldNameMap, validators)
}

func VcTrustsSummaryBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["domain"] = bindings.NewStringType()
    fieldNameMap["domain"] = "Domain"
    fields["label"] = bindings.NewStringType()
    fieldNameMap["label"] = "Label"
    fields["type"] = bindings.NewStringType()
    fieldNameMap["type"] = "Type_"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.trust.vc_trusts.summary",fields, reflect.TypeOf(VcTrustsSummary{}), fieldNameMap, validators)
}


