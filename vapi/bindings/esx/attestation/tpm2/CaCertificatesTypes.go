/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Code generated. DO NOT EDIT.

/*
 * Data type definitions file for service: CaCertificates.
 * Includes binding types of a structures and enumerations defined in the service.
 * Shared by client-side stubs and server-side skeletons to ensure type
 * compatibility.
 */

package tpm2

import (
	"reflect"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/bindings"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/data"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/protocol"
)

// Resource type for TPM 2.0 CA certificate.
const CaCertificates_RESOURCE_TYPE = "com.vmware.esx.attestation.tpm2.ca_certificates"


// The ``Summary`` class contains information that summarizes a TPM CA certificate.
type CaCertificatesSummary struct {
    // A unique name for the TPM CA certificate.
	Name string
}

// The ``X509CertChain`` class contains information that fully describes a certificate chain.
type CaCertificatesX509CertChain struct {
    // A list of certificates each in PEM format. 
    //
    //  Certificates must be ordered from the root first.
	Certs []string
}

// The ``Info`` class contains information that describes a TPM CA certificate.
type CaCertificatesInfo struct {
    // The CA certificate chain.
	CertChain CaCertificatesX509CertChain
}

// The ``CreateSpec`` class contains information that describes a TPM CA certificate.
type CaCertificatesCreateSpec struct {
    // A unique name for the TPM CA certificate.
	Name string
    // The CA certificate chain. 
    //
    //  Certificates may either be added one at a time, or as a chain. Adding the certificates as a chain allows the group to be managed as a whole. For example, an entire chain can be deleted in one CaCertificates#delete operation. 
    //
    //  When certificates are added one at a time, the order must be root first, followed by any intermediates. The intermediates certificates must also be ordered in the direction from root to leaf. 
    //
    //  Similarly, when added as a chain the list must be ordered in the direction from root to leaf.
	CertChain *CaCertificatesX509CertChain
}



func caCertificatesListInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func caCertificatesListOutputType() bindings.BindingType {
	return bindings.NewListType(bindings.NewReferenceType(CaCertificatesSummaryBindingType), reflect.TypeOf([]CaCertificatesSummary{}))
}

func caCertificatesListRestMetadata() protocol.OperationRestMetadata {
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
		"/esx/attestation/tpm2/ca-certificates",
		resultHeaders,
		200,
		errorHeaders,
		map[string]int{"Error": 500,"Unauthenticated": 401,"Unauthorized": 403})
}

func caCertificatesCreateInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["spec"] = bindings.NewReferenceType(CaCertificatesCreateSpecBindingType)
	fieldNameMap["spec"] = "Spec"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func caCertificatesCreateOutputType() bindings.BindingType {
	return bindings.NewVoidType()
}

func caCertificatesCreateRestMetadata() protocol.OperationRestMetadata {
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	paramsTypeMap["spec"] = bindings.NewReferenceType(CaCertificatesCreateSpecBindingType)
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
		"/esx/attestation/tpm2/ca-certificates",
		resultHeaders,
		204,
		errorHeaders,
		map[string]int{"AlreadyExists": 400,"Error": 500,"InvalidArgument": 400,"Unauthenticated": 401,"Unauthorized": 403})
}

func caCertificatesDeleteInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["name"] = bindings.NewIdType([]string{"com.vmware.esx.attestation.tpm2.ca_certificates"}, "")
	fieldNameMap["name"] = "Name"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func caCertificatesDeleteOutputType() bindings.BindingType {
	return bindings.NewVoidType()
}

func caCertificatesDeleteRestMetadata() protocol.OperationRestMetadata {
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	paramsTypeMap["name"] = bindings.NewIdType([]string{"com.vmware.esx.attestation.tpm2.ca_certificates"}, "")
	paramsTypeMap["name"] = bindings.NewIdType([]string{"com.vmware.esx.attestation.tpm2.ca_certificates"}, "")
	pathParams["name"] = "name"
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
		"/esx/attestation/tpm2/ca-certificates/{name}",
		resultHeaders,
		204,
		errorHeaders,
		map[string]int{"Error": 500,"InvalidArgument": 400,"NotFound": 404,"Unauthenticated": 401,"Unauthorized": 403})
}

func caCertificatesGetInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["name"] = bindings.NewIdType([]string{"com.vmware.esx.attestation.tpm2.ca_certificates"}, "")
	fieldNameMap["name"] = "Name"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func caCertificatesGetOutputType() bindings.BindingType {
	return bindings.NewReferenceType(CaCertificatesInfoBindingType)
}

func caCertificatesGetRestMetadata() protocol.OperationRestMetadata {
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	paramsTypeMap["name"] = bindings.NewIdType([]string{"com.vmware.esx.attestation.tpm2.ca_certificates"}, "")
	paramsTypeMap["name"] = bindings.NewIdType([]string{"com.vmware.esx.attestation.tpm2.ca_certificates"}, "")
	pathParams["name"] = "name"
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
		"/esx/attestation/tpm2/ca-certificates/{name}",
		resultHeaders,
		200,
		errorHeaders,
		map[string]int{"Error": 500,"InvalidArgument": 400,"NotFound": 404,"Unauthenticated": 401,"Unauthorized": 403})
}


func CaCertificatesSummaryBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["name"] = bindings.NewIdType([]string{"com.vmware.esx.attestation.tpm2.ca_certificates"}, "")
	fieldNameMap["name"] = "Name"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.esx.attestation.tpm2.ca_certificates.summary", fields, reflect.TypeOf(CaCertificatesSummary{}), fieldNameMap, validators)
}

func CaCertificatesX509CertChainBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["certs"] = bindings.NewListType(bindings.NewStringType(), reflect.TypeOf([]string{}))
	fieldNameMap["certs"] = "Certs"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.esx.attestation.tpm2.ca_certificates.x509_cert_chain", fields, reflect.TypeOf(CaCertificatesX509CertChain{}), fieldNameMap, validators)
}

func CaCertificatesInfoBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["cert_chain"] = bindings.NewReferenceType(CaCertificatesX509CertChainBindingType)
	fieldNameMap["cert_chain"] = "CertChain"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.esx.attestation.tpm2.ca_certificates.info", fields, reflect.TypeOf(CaCertificatesInfo{}), fieldNameMap, validators)
}

func CaCertificatesCreateSpecBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["name"] = bindings.NewIdType([]string{"com.vmware.esx.attestation.tpm2.ca_certificates"}, "")
	fieldNameMap["name"] = "Name"
	fields["cert_chain"] = bindings.NewOptionalType(bindings.NewReferenceType(CaCertificatesX509CertChainBindingType))
	fieldNameMap["cert_chain"] = "CertChain"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.esx.attestation.tpm2.ca_certificates.create_spec", fields, reflect.TypeOf(CaCertificatesCreateSpec{}), fieldNameMap, validators)
}

