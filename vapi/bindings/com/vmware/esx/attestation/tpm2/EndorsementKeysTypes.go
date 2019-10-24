/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Data type definitions file for service: EndorsementKeys.
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

// Resource type for TPM 2.0 endorsement key.
const EndorsementKeys_RESOURCE_TYPE = "com.vmware.esx.attestation.tpm2.endorsement_keys"
// The ``Summary`` class contains information that summarizes a TPM endorsement key.
type EndorsementKeysSummary struct {
    // A unique name for the TPM endorsement key.
	Name string
}

// The ``Info`` class contains information that describes a TPM endorsement key.
type EndorsementKeysInfo struct {
    // TPM public endorsement key in PEM format.
	PublicKey string
}

// The ``CreateSpec`` class contains information that describes a TPM endorsement key. 
//
//  Only one of EndorsementKeysCreateSpec#publicKey or EndorsementKeysCreateSpec#certificate must be specified.
type EndorsementKeysCreateSpec struct {
    // A unique name for the TPM endorsement key. 
    //
    //  The unique name should be something that an administrator can use to easily identify the remote system. For example, the hostname, or hardware UUID.
	Name string
    // TPM public endorsement key in PEM format.
	PublicKey *string
    // TPM endorsement key certificate in PEM format. 
    //
    //  When a endorsement key certificate is provided, it will be verified against the CaCertificate list. EK certificates that are not signed by one of the CA certificates will be rejected. 
    //
    //  Using this format allows for failures to be caught during configuration rather than later during attestation.
	Certificate *string
}


func endorsementKeysListInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func endorsementKeysListOutputType() bindings.BindingType {
	return bindings.NewListType(bindings.NewReferenceType(EndorsementKeysSummaryBindingType), reflect.TypeOf([]EndorsementKeysSummary{}))
}

func endorsementKeysListRestMetadata() protocol.OperationRestMetadata {
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
		"/esx/attestation/tpm2/endorsement-keys",
		resultHeaders,
		200,
		errorHeaders,
		map[string]int{"Error": 500,"Unauthenticated": 401,"Unauthorized": 403})
}

func endorsementKeysCreateInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["spec"] = bindings.NewReferenceType(EndorsementKeysCreateSpecBindingType)
	fieldNameMap["spec"] = "Spec"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func endorsementKeysCreateOutputType() bindings.BindingType {
	return bindings.NewVoidType()
}

func endorsementKeysCreateRestMetadata() protocol.OperationRestMetadata {
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	paramsTypeMap["spec"] = bindings.NewReferenceType(EndorsementKeysCreateSpecBindingType)
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
		"/esx/attestation/tpm2/endorsement-keys",
		resultHeaders,
		204,
		errorHeaders,
		map[string]int{"AlreadyExists": 400,"Error": 500,"InvalidArgument": 400,"Unauthenticated": 401,"Unauthorized": 403})
}

func endorsementKeysDeleteInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["name"] = bindings.NewIdType([]string{"com.vmware.esx.attestation.tpm2.endorsement_keys"}, "")
	fieldNameMap["name"] = "Name"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func endorsementKeysDeleteOutputType() bindings.BindingType {
	return bindings.NewVoidType()
}

func endorsementKeysDeleteRestMetadata() protocol.OperationRestMetadata {
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	paramsTypeMap["name"] = bindings.NewIdType([]string{"com.vmware.esx.attestation.tpm2.endorsement_keys"}, "")
	paramsTypeMap["name"] = bindings.NewIdType([]string{"com.vmware.esx.attestation.tpm2.endorsement_keys"}, "")
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
		"/esx/attestation/tpm2/endorsement-keys/{name}",
		resultHeaders,
		204,
		errorHeaders,
		map[string]int{"Error": 500,"InvalidArgument": 400,"NotFound": 404,"Unauthenticated": 401,"Unauthorized": 403})
}

func endorsementKeysGetInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["name"] = bindings.NewIdType([]string{"com.vmware.esx.attestation.tpm2.endorsement_keys"}, "")
	fieldNameMap["name"] = "Name"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func endorsementKeysGetOutputType() bindings.BindingType {
	return bindings.NewReferenceType(EndorsementKeysInfoBindingType)
}

func endorsementKeysGetRestMetadata() protocol.OperationRestMetadata {
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	paramsTypeMap["name"] = bindings.NewIdType([]string{"com.vmware.esx.attestation.tpm2.endorsement_keys"}, "")
	paramsTypeMap["name"] = bindings.NewIdType([]string{"com.vmware.esx.attestation.tpm2.endorsement_keys"}, "")
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
		"/esx/attestation/tpm2/endorsement-keys/{name}",
		resultHeaders,
		200,
		errorHeaders,
		map[string]int{"Error": 500,"InvalidArgument": 400,"NotFound": 404,"Unauthenticated": 401,"Unauthorized": 403})
}


func EndorsementKeysSummaryBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["name"] = bindings.NewIdType([]string{"com.vmware.esx.attestation.tpm2.endorsement_keys"}, "")
	fieldNameMap["name"] = "Name"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.esx.attestation.tpm2.endorsement_keys.summary", fields, reflect.TypeOf(EndorsementKeysSummary{}), fieldNameMap, validators)
}

func EndorsementKeysInfoBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["public_key"] = bindings.NewStringType()
	fieldNameMap["public_key"] = "PublicKey"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.esx.attestation.tpm2.endorsement_keys.info", fields, reflect.TypeOf(EndorsementKeysInfo{}), fieldNameMap, validators)
}

func EndorsementKeysCreateSpecBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["name"] = bindings.NewIdType([]string{"com.vmware.esx.attestation.tpm2.endorsement_keys"}, "")
	fieldNameMap["name"] = "Name"
	fields["public_key"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["public_key"] = "PublicKey"
	fields["certificate"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["certificate"] = "Certificate"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.esx.attestation.tpm2.endorsement_keys.create_spec", fields, reflect.TypeOf(EndorsementKeysCreateSpec{}), fieldNameMap, validators)
}

