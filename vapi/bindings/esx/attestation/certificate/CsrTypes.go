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
package certificate

import (
	"reflect"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/bindings"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/data"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/protocol"
)



// The ``Info`` class contains information for a CSR.
type CsrInfo struct {
    // Certificate Signing Request (CSR) in PEM format.
	Csr string
}

// The ``CreateSpec`` class contains information to generate a private key and CSR.
type CsrCreateSpec struct {
    // Size in bits of the generated private key.
	KeySize *int64
    // Common name for the certificate. 
    //
    //  This is typically the hostname of the associated server.
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
    // Email field in certificate extensions
	EmailAddress string
    // Subject alternative names.
	SubjectAltName []string
}



func csrCreateInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["spec"] = bindings.NewReferenceType(CsrCreateSpecBindingType)
	fieldNameMap["spec"] = "Spec"
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
	paramsTypeMap["spec"] = bindings.NewReferenceType(CsrCreateSpecBindingType)
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
		"/esx/attestation/certificate/csr",
		resultHeaders,
		201,
		errorHeaders,
		map[string]int{"Error": 500,"InvalidArgument": 400,"Unauthenticated": 401,"Unauthorized": 403})
}

func csrDeleteInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func csrDeleteOutputType() bindings.BindingType {
	return bindings.NewVoidType()
}

func csrDeleteRestMetadata() protocol.OperationRestMetadata {
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
		"DELETE",
		"/esx/attestation/certificate/csr",
		resultHeaders,
		204,
		errorHeaders,
		map[string]int{"Error": 500,"NotFound": 404,"Unauthenticated": 401,"Unauthorized": 403})
}

func csrGetInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
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
		"/esx/attestation/certificate/csr",
		resultHeaders,
		200,
		errorHeaders,
		map[string]int{"Error": 500,"NotFound": 404,"Unauthenticated": 401,"Unauthorized": 403})
}


func CsrInfoBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["csr"] = bindings.NewStringType()
	fieldNameMap["csr"] = "Csr"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.esx.attestation.certificate.csr.info", fields, reflect.TypeOf(CsrInfo{}), fieldNameMap, validators)
}

func CsrCreateSpecBindingType() bindings.BindingType {
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
	return bindings.NewStructType("com.vmware.esx.attestation.certificate.csr.create_spec", fields, reflect.TypeOf(CsrCreateSpec{}), fieldNameMap, validators)
}

