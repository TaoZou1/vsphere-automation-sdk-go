/* Copyright © 2019 VMware, Inc. All Rights Reserved.
     SPDX-License-Identifier: BSD-2-Clause */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Data type definitions file for service: Tls.
 * Includes binding types of a structures and enumerations defined in the service.
 * Shared by client-side stubs and server-side skeletons to ensure type
 * compatibility.
 */

package vcenter

import (
    "reflect"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/bindings"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/data"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/protocol"
    "time"
)




// The ``Info`` class contains information from a TLS certificate.
 type TlsInfo struct {
    // Version (version number) value from the certificate.
    Version int64
    // SerialNumber value from the certificate.
    SerialNumber string
    // Signature algorithm name from the certificate.
    SignatureAlgorithm string
    // Issuer (issuer distinguished name) value from the certificate.
    IssuerDn string
    // validFrom specify the start date of the certificate.
    ValidFrom time.Time
    // validTo specify the end date of the certificate.
    ValidTo time.Time
    // Subject (subject distinguished name) value from the certificate.
    SubjectDn string
    // Thumbprint value from the certificate.
    Thumbprint string
    // Certificate constraints isCA from the critical BasicConstraints extension, (OID = 2.5.29.19).
    IsCA bool
    // Certificate constraints path length from the critical BasicConstraints extension, (OID = 2.5.29.19).
    PathLengthConstraint int64
    // Collection of keyusage contained in the certificate.
    KeyUsage []string
    // Collection of extended keyusage that contains details for which the certificate can be used for.
    ExtendedKeyUsage []string
    // Collection of subject alternative names.
    SubjectAlternativeName []string
    // Collection of authority information access URI.
    AuthorityInformationAccessUri []string
    // TLS certificate in PEM format.
    Cert string
}



func (TlsInfo TlsInfo) Error() string {
    return "com.vmware.vcenter.certificate_management.vcenter.info"
}



// The ``Spec`` class contains information for a Certificate and Private Key.
 type TlsSpec struct {
    // Certificate string in PEM format.
    Cert string
    // Private key string in PEM format.
    Key *string
    // Third party Root CA certificate in PEM format.
    RootCert *string
}



func (TlsSpec TlsSpec) Error() string {
    return "com.vmware.vcenter.certificate_management.vcenter.spec"
}



// The ``ReplaceSpec`` class contains information to generate a Private Key , CSR and hence VMCA signed machine SSL.
 type TlsReplaceSpec struct {
    // The size of the key to be used for public and private key generation.
    KeySize *int64
    // The common name of the host for which certificate is generated
    CommonName *string
    // Organization field in certificate subject
    Organization string
    // Organization unit field in certificate subject
    OrganizationUnit string
    // Locality field in certificate subject
    Locality string
    // State field in certificate subject
    StateOrProvince string
    // Country field in certificate subject
    Country string
    // Email field in Certificate extensions
    EmailAddress string
    // SubjectAltName is list of Dns Names and Ip addresses
    SubjectAltName []string
}



func (TlsReplaceSpec TlsReplaceSpec) Error() string {
    return "com.vmware.vcenter.certificate_management.vcenter.replace_spec"
}






func tlsSetInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["spec"] = bindings.NewReferenceType(TlsSpecBindingType)
    fieldNameMap["spec"] = "Spec"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func tlsSetOutputType() bindings.BindingType {
    return bindings.NewVoidType()
}

func tlsSetRestMetadata() protocol.OperationRestMetadata {
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
       map[string]int{"NotFound": 404,"AlreadyExists": 400,"Error": 500})
}


func tlsGetInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func tlsGetOutputType() bindings.BindingType {
    return bindings.NewReferenceType(TlsInfoBindingType)
}

func tlsGetRestMetadata() protocol.OperationRestMetadata {
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
       map[string]int{"NotFound": 404,"Error": 500})
}


func tlsRenewInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["duration"] = bindings.NewOptionalType(bindings.NewIntegerType())
    fieldNameMap["duration"] = "Duration"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func tlsRenewOutputType() bindings.BindingType {
    return bindings.NewVoidType()
}

func tlsRenewRestMetadata() protocol.OperationRestMetadata {
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
       map[string]int{"Unsupported": 400,"InvalidArgument": 400,"Error": 500})
}


func tlsReplaceVmcaSignedInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["spec"] = bindings.NewReferenceType(TlsReplaceSpecBindingType)
    fieldNameMap["spec"] = "Spec"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func tlsReplaceVmcaSignedOutputType() bindings.BindingType {
    return bindings.NewVoidType()
}

func tlsReplaceVmcaSignedRestMetadata() protocol.OperationRestMetadata {
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
       map[string]int{"InvalidArgument": 400,"Error": 500})
}



func TlsInfoBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["version"] = bindings.NewIntegerType()
    fieldNameMap["version"] = "Version"
    fields["serial_number"] = bindings.NewStringType()
    fieldNameMap["serial_number"] = "SerialNumber"
    fields["signature_algorithm"] = bindings.NewStringType()
    fieldNameMap["signature_algorithm"] = "SignatureAlgorithm"
    fields["issuer_dn"] = bindings.NewStringType()
    fieldNameMap["issuer_dn"] = "IssuerDn"
    fields["valid_from"] = bindings.NewDateTimeType()
    fieldNameMap["valid_from"] = "ValidFrom"
    fields["valid_to"] = bindings.NewDateTimeType()
    fieldNameMap["valid_to"] = "ValidTo"
    fields["subject_dn"] = bindings.NewStringType()
    fieldNameMap["subject_dn"] = "SubjectDn"
    fields["thumbprint"] = bindings.NewStringType()
    fieldNameMap["thumbprint"] = "Thumbprint"
    fields["is_CA"] = bindings.NewBooleanType()
    fieldNameMap["is_CA"] = "IsCA"
    fields["path_length_constraint"] = bindings.NewIntegerType()
    fieldNameMap["path_length_constraint"] = "PathLengthConstraint"
    fields["key_usage"] = bindings.NewListType(bindings.NewStringType(), reflect.TypeOf([]string{}))
    fieldNameMap["key_usage"] = "KeyUsage"
    fields["extended_key_usage"] = bindings.NewListType(bindings.NewStringType(), reflect.TypeOf([]string{}))
    fieldNameMap["extended_key_usage"] = "ExtendedKeyUsage"
    fields["subject_alternative_name"] = bindings.NewListType(bindings.NewStringType(), reflect.TypeOf([]string{}))
    fieldNameMap["subject_alternative_name"] = "SubjectAlternativeName"
    fields["authority_information_access_uri"] = bindings.NewListType(bindings.NewStringType(), reflect.TypeOf([]string{}))
    fieldNameMap["authority_information_access_uri"] = "AuthorityInformationAccessUri"
    fields["cert"] = bindings.NewStringType()
    fieldNameMap["cert"] = "Cert"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.certificate_management.vcenter.tls.info",fields, reflect.TypeOf(TlsInfo{}), fieldNameMap, validators)
}

func TlsSpecBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["cert"] = bindings.NewStringType()
    fieldNameMap["cert"] = "Cert"
    fields["key"] = bindings.NewOptionalType(bindings.NewSecretType())
    fieldNameMap["key"] = "Key"
    fields["root_cert"] = bindings.NewOptionalType(bindings.NewStringType())
    fieldNameMap["root_cert"] = "RootCert"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.certificate_management.vcenter.tls.spec",fields, reflect.TypeOf(TlsSpec{}), fieldNameMap, validators)
}

func TlsReplaceSpecBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["key_size"] = bindings.NewOptionalType(bindings.NewIntegerType())
    fieldNameMap["key_size"] = "KeySize"
    fields["common_name"] = bindings.NewOptionalType(bindings.NewStringType())
    fieldNameMap["common_name"] = "CommonName"
    fields["organization"] = bindings.NewStringType()
    fieldNameMap["organization"] = "Organization"
    fields["organization_unit"] = bindings.NewStringType()
    fieldNameMap["organization_unit"] = "OrganizationUnit"
    fields["locality"] = bindings.NewStringType()
    fieldNameMap["locality"] = "Locality"
    fields["state_or_province"] = bindings.NewStringType()
    fieldNameMap["state_or_province"] = "StateOrProvince"
    fields["country"] = bindings.NewStringType()
    fieldNameMap["country"] = "Country"
    fields["email_address"] = bindings.NewStringType()
    fieldNameMap["email_address"] = "EmailAddress"
    fields["subject_alt_name"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewStringType(), reflect.TypeOf([]string{})))
    fieldNameMap["subject_alt_name"] = "SubjectAltName"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.certificate_management.vcenter.tls.replace_spec",fields, reflect.TypeOf(TlsReplaceSpec{}), fieldNameMap, validators)
}


