/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Code generated. DO NOT EDIT.

/*
 * Data type definitions file for service: Settings.
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



// The ``Info`` class contains information that describes the TPM 2.0 protocol settings.
type SettingsInfo struct {
    // Require registered TPM endorsement keys. 
    //
    //  During attestation, the attested host will always send its endorsement key to the Attestation Service. With this option is set, the Attestation Service will only proceed with attestation if the endorsement key has been added to the list of configured trusted endorsement keys.
	RequireEndorsementKeys bool
    // Require TPM endorsement key certificate validation. 
    //
    //  During attestation, the attested host will send its endorsement key certificate if one is available. With this option set, the Attestation Service will validate the endorsement key certificate against the list of configured trusted TPM CA certificates. Only endorsement key certificates that are signed by a trusted TPM CA certificate will be able to successfully attest.
	RequireCertificateValidation bool
}

// The ``UpdateSpec`` class contains information that describes changes to the TPM 2.0 protocol settings.
type SettingsUpdateSpec struct {
    // Require registered TPM endorsement keys.
	RequireEndorsementKeys *bool
    // Require TPM endorsement key certificate validation.
	RequireCertificateValidation *bool
}



func settingsGetInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func settingsGetOutputType() bindings.BindingType {
	return bindings.NewReferenceType(SettingsInfoBindingType)
}

func settingsGetRestMetadata() protocol.OperationRestMetadata {
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
		"/esx/attestation/tpm2/settings",
		resultHeaders,
		200,
		errorHeaders,
		map[string]int{"Error": 500,"Unauthenticated": 401,"Unauthorized": 403})
}

func settingsUpdateInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["spec"] = bindings.NewReferenceType(SettingsUpdateSpecBindingType)
	fieldNameMap["spec"] = "Spec"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func settingsUpdateOutputType() bindings.BindingType {
	return bindings.NewVoidType()
}

func settingsUpdateRestMetadata() protocol.OperationRestMetadata {
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	paramsTypeMap["spec"] = bindings.NewReferenceType(SettingsUpdateSpecBindingType)
	resultHeaders := map[string]string{}
	errorHeaders := map[string]string{}
	errorHeaders["Unauthenticated.challenge"] = "WWW-Authenticate"
	return protocol.NewOperationRestMetadata(
		paramsTypeMap,
		pathParams,
		queryParams,
		headerParams,
		"spec",
		"PATCH",
		"/esx/attestation/tpm2/settings",
		resultHeaders,
		204,
		errorHeaders,
		map[string]int{"Error": 500,"InvalidArgument": 400,"Unauthenticated": 401,"Unauthorized": 403})
}


func SettingsInfoBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["require_endorsement_keys"] = bindings.NewBooleanType()
	fieldNameMap["require_endorsement_keys"] = "RequireEndorsementKeys"
	fields["require_certificate_validation"] = bindings.NewBooleanType()
	fieldNameMap["require_certificate_validation"] = "RequireCertificateValidation"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.esx.attestation.tpm2.settings.info", fields, reflect.TypeOf(SettingsInfo{}), fieldNameMap, validators)
}

func SettingsUpdateSpecBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["require_endorsement_keys"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fieldNameMap["require_endorsement_keys"] = "RequireEndorsementKeys"
	fields["require_certificate_validation"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fieldNameMap["require_certificate_validation"] = "RequireCertificateValidation"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.esx.attestation.tpm2.settings.update_spec", fields, reflect.TypeOf(SettingsUpdateSpec{}), fieldNameMap, validators)
}

