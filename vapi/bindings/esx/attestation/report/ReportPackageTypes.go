/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Code generated. DO NOT EDIT.

/*
 * Data type definitions file for package: com.vmware.esx.attestation.report.
 * Includes binding types of a top level structures and enumerations.
 * Shared by client-side stubs and server-side skeletons to ensure type
 * compatibility.
 */

package report

import (
	"reflect"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/bindings"
)


// The ``DocumentType`` enumeration class defines valid attestation report document types.
//
// <p> See {@link com.vmware.vapi.bindings.ApiEnumeration enumerated types description}.
type DocumentType string

const (
    // A JSON Web Token for use in KMS authentication.
	DocumentType_KMS_AUTH_JWT DocumentType = "KMS_AUTH_JWT"
)

func (d DocumentType) DocumentType() bool {
	switch d {
	case DocumentType_KMS_AUTH_JWT:
		return true
	default:
		return false
	}
}


// The ``AttestResult`` class contains the result of an attestation report request.
type AttestResult struct {
    // Indicates if the requester has been successfully attested. 
    //
    //  When true, all requested documents are available.
	Attested bool
    // The documents that were requested. 
    //
    //  This list is only populated when AttestResult#attested reports true.
	Documents []Document
}

// The ``AttestRequest`` class contains information used to specify an attestation report request.
type AttestRequest struct {
    // The types of document being requested.
	Types []DocumentType
}

// The ``Document`` class contains an attestation report document.
type Document struct {
    // The type of attestation report document.
	Type_ DocumentType
    // A JSON Web Token for use in KMS authentication.
	KmsAuthJwt *string
}

// The ``Tpm2Result`` class contains the result of an attestation report request using the TPM 2.0 protocol.
type Tpm2Result struct {
    // A set of contexts that must be processed by the client.
	Contexts [][]byte
    // Indicates if the requester has been successfully attested. 
    //
    //  When true, all requested documents are available.
	Attested bool
    // The documents that were requested. 
    //
    //  This list is only populated when AttestResult#attested reports true.
	Documents []Document
}

// The ``Tpm2Identity`` class contains a unique TPM 2.0 identifier for the remote host.
type Tpm2Identity struct {
    // TPM 2.0 endorsement key (EK) certificate 
    //
    //  When an endorsement key certificate is available, all attestation requests should begin by sending it as the remote host's identity. The server will validate the certificate against a configured set of trusted TPM certificate authorities (CAs).
	EkCert *string
    // TPM 2.0 public endorsement key (EK) 
    //
    //  When an endorsement key certificate is not available, an attestation request may be made using the public part of the endorsement key instead. This identification method will only work if the server is configured to accept unsigned EKs, and the public key has been added to the trusted set. 
    //
    //  Not required when Tpm2Identity#ekCert is specified.
	EkPub *string
}

// The ``Tpm2Request`` class contains information used to specify an attestation report request using the TPM 2.0 protocol.
type Tpm2Request struct {
    // The remote TPM identity.
	Identity Tpm2Identity
    // A set of contexts that will be processed by the server.
	Contexts [][]byte
    // The types of document being requested.
	Types []DocumentType
}




func AttestResultBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["attested"] = bindings.NewBooleanType()
	fieldNameMap["attested"] = "Attested"
	fields["documents"] = bindings.NewListType(bindings.NewReferenceType(DocumentBindingType), reflect.TypeOf([]Document{}))
	fieldNameMap["documents"] = "Documents"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.esx.attestation.report.attest_result", fields, reflect.TypeOf(AttestResult{}), fieldNameMap, validators)
}

func AttestRequestBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["types"] = bindings.NewListType(bindings.NewEnumType("com.vmware.esx.attestation.report.document_type", reflect.TypeOf(DocumentType(DocumentType_KMS_AUTH_JWT))), reflect.TypeOf([]DocumentType{}))
	fieldNameMap["types"] = "Types"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.esx.attestation.report.attest_request", fields, reflect.TypeOf(AttestRequest{}), fieldNameMap, validators)
}

func DocumentBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["type"] = bindings.NewEnumType("com.vmware.esx.attestation.report.document_type", reflect.TypeOf(DocumentType(DocumentType_KMS_AUTH_JWT)))
	fieldNameMap["type"] = "Type_"
	fields["kms_auth_jwt"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["kms_auth_jwt"] = "KmsAuthJwt"
	var validators = []bindings.Validator{}
	uv1 := bindings.NewUnionValidator("type",
		map[string][]bindings.FieldData{
			"KMS_AUTH_JWT": []bindings.FieldData{
				bindings.NewFieldData("kms_auth_jwt", true),
			},
		},
	)
	validators = append(validators, uv1)
	return bindings.NewStructType("com.vmware.esx.attestation.report.document", fields, reflect.TypeOf(Document{}), fieldNameMap, validators)
}

func Tpm2ResultBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["contexts"] = bindings.NewListType(bindings.NewBlobType(), reflect.TypeOf([][]byte{}))
	fieldNameMap["contexts"] = "Contexts"
	fields["attested"] = bindings.NewBooleanType()
	fieldNameMap["attested"] = "Attested"
	fields["documents"] = bindings.NewListType(bindings.NewReferenceType(DocumentBindingType), reflect.TypeOf([]Document{}))
	fieldNameMap["documents"] = "Documents"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.esx.attestation.report.tpm2_result", fields, reflect.TypeOf(Tpm2Result{}), fieldNameMap, validators)
}

func Tpm2IdentityBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["ek_cert"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["ek_cert"] = "EkCert"
	fields["ek_pub"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["ek_pub"] = "EkPub"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.esx.attestation.report.tpm2_identity", fields, reflect.TypeOf(Tpm2Identity{}), fieldNameMap, validators)
}

func Tpm2RequestBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["identity"] = bindings.NewReferenceType(Tpm2IdentityBindingType)
	fieldNameMap["identity"] = "Identity"
	fields["contexts"] = bindings.NewListType(bindings.NewBlobType(), reflect.TypeOf([][]byte{}))
	fieldNameMap["contexts"] = "Contexts"
	fields["types"] = bindings.NewListType(bindings.NewEnumType("com.vmware.esx.attestation.report.document_type", reflect.TypeOf(DocumentType(DocumentType_KMS_AUTH_JWT))), reflect.TypeOf([]DocumentType{}))
	fieldNameMap["types"] = "Types"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.esx.attestation.report.tpm2_request", fields, reflect.TypeOf(Tpm2Request{}), fieldNameMap, validators)
}


