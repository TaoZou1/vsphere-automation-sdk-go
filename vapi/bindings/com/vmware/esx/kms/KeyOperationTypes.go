/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Data type definitions file for service: KeyOperation.
 * Includes binding types of a structures and enumerations defined in the service.
 * Shared by client-side stubs and server-side skeletons to ensure type
 * compatibility.
 */
package kms

import (
	"reflect"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/bindings"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/data"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/protocol"
)

// The ``GeneratedKey`` class contains properties that are returned by KeyOperation#generateKey.
type KeyOperationGeneratedKey struct {
    // The encrypted copy of key data. Can be used as key identifier. 
    //
    //  Additional information is encoded in the ciphertext needed for retrieving the key.
	Ciphertext string
    // The plaintext key data. 
	Plaintext *string
}

// The ``EncryptResult`` class contains properties that are returned by KeyOperation#encrypt.
type KeyOperationEncryptResult struct {
    // The encrypted data. 
    //
    //  Additional information is encoded in the ciphertext needed for decryption.
	Ciphertext string
}

// The ``DecryptResult`` class contains properties that are returned by KeyOperation#decrypt.
type KeyOperationDecryptResult struct {
    // The decrypted data.
	Plaintext string
}


func keyOperationGenerateKeyInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["provider"] = bindings.NewIdType([]string{"com.vmware.esx.kms.providers"}, "")
	fields["num_of_bytes"] = bindings.NewIntegerType()
	fieldNameMap["provider"] = "Provider"
	fieldNameMap["num_of_bytes"] = "NumOfBytes"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func keyOperationGenerateKeyOutputType() bindings.BindingType {
	return bindings.NewReferenceType(KeyOperationGeneratedKeyBindingType)
}

func keyOperationGenerateKeyRestMetadata() protocol.OperationRestMetadata {
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
		"null",
		"",
		resultHeaders,
		0,
		errorHeaders,
		map[string]int{"InvalidArgument": 400,"NotFound": 404,"Unauthenticated": 401,"Unauthorized": 403,"Error": 500})
}

func keyOperationEncryptInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["provider"] = bindings.NewIdType([]string{"com.vmware.esx.kms.providers"}, "")
	fields["plaintext"] = bindings.NewSecretType()
	fieldNameMap["provider"] = "Provider"
	fieldNameMap["plaintext"] = "Plaintext"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func keyOperationEncryptOutputType() bindings.BindingType {
	return bindings.NewReferenceType(KeyOperationEncryptResultBindingType)
}

func keyOperationEncryptRestMetadata() protocol.OperationRestMetadata {
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
		"null",
		"",
		resultHeaders,
		0,
		errorHeaders,
		map[string]int{"InvalidArgument": 400,"NotFound": 404,"Unauthenticated": 401,"Unauthorized": 403,"Error": 500})
}

func keyOperationDecryptInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["provider"] = bindings.NewIdType([]string{"com.vmware.esx.kms.providers"}, "")
	fields["ciphertext"] = bindings.NewStringType()
	fieldNameMap["provider"] = "Provider"
	fieldNameMap["ciphertext"] = "Ciphertext"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func keyOperationDecryptOutputType() bindings.BindingType {
	return bindings.NewReferenceType(KeyOperationDecryptResultBindingType)
}

func keyOperationDecryptRestMetadata() protocol.OperationRestMetadata {
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
		"null",
		"",
		resultHeaders,
		0,
		errorHeaders,
		map[string]int{"InvalidArgument": 400,"NotFound": 404,"Unauthenticated": 401,"Unauthorized": 403,"Error": 500})
}


func KeyOperationGeneratedKeyBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["ciphertext"] = bindings.NewStringType()
	fieldNameMap["ciphertext"] = "Ciphertext"
	fields["plaintext"] = bindings.NewOptionalType(bindings.NewSecretType())
	fieldNameMap["plaintext"] = "Plaintext"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.esx.kms.key_operation.generated_key", fields, reflect.TypeOf(KeyOperationGeneratedKey{}), fieldNameMap, validators)
}

func KeyOperationEncryptResultBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["ciphertext"] = bindings.NewStringType()
	fieldNameMap["ciphertext"] = "Ciphertext"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.esx.kms.key_operation.encrypt_result", fields, reflect.TypeOf(KeyOperationEncryptResult{}), fieldNameMap, validators)
}

func KeyOperationDecryptResultBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["plaintext"] = bindings.NewSecretType()
	fieldNameMap["plaintext"] = "Plaintext"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.esx.kms.key_operation.decrypt_result", fields, reflect.TypeOf(KeyOperationDecryptResult{}), fieldNameMap, validators)
}

