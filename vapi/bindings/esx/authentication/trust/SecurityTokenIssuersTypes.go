/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Code generated. DO NOT EDIT.

/*
 * Data type definitions file for service: SecurityTokenIssuers.
 * Includes binding types of a structures and enumerations defined in the service.
 * Shared by client-side stubs and server-side skeletons to ensure type
 * compatibility.
 */

package trust

import (
	"reflect"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/bindings"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/data"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/protocol"
)

// The resource type for the SecurityTokenIssuers instances.
const SecurityTokenIssuers_RESOURCE_TYPE = "com.vmware.esx.authentication.trust.security-token-issuer"


// The ``SummaryType`` enumeration class defines the types of SecurityTokenIssuersSummary members to return from the SecurityTokenIssuers#list method.
//
//  The profile information could include the access grants or be a shorter summary.
//
// <p> See {@link com.vmware.vapi.bindings.ApiEnumeration enumerated types description}.
type SecurityTokenIssuersSummaryType string

const (
    // The full security token isuer information, including certificates.
	SecurityTokenIssuersSummaryType_FULL SecurityTokenIssuersSummaryType = "FULL"
    // A summary containing only the security token issuer alias and the issuer string.
	SecurityTokenIssuersSummaryType_NORMAL SecurityTokenIssuersSummaryType = "NORMAL"
    // A brief summary, containing only the security token issuer alias.
	SecurityTokenIssuersSummaryType_BRIEF SecurityTokenIssuersSummaryType = "BRIEF"
)

func (s SecurityTokenIssuersSummaryType) SecurityTokenIssuersSummaryType() bool {
	switch s {
	case SecurityTokenIssuersSummaryType_FULL:
		return true
	case SecurityTokenIssuersSummaryType_NORMAL:
		return true
	case SecurityTokenIssuersSummaryType_BRIEF:
		return true
	default:
		return false
	}
}


// The ``X509CertChain`` class contains information about a chain of X.509 certificates.
//
//  The structure includes a list of strings, each string is a PEM encoded certificate.
type SecurityTokenIssuersX509CertChain struct {
    // Certificate chain in base64 format
	CertChain []string
}

// The ``IssuerAlreadyExistsInfo`` class contains the information under which alias the issuer is.
type SecurityTokenIssuersIssuerAlreadyExistsInfo struct {
    // Existing issuer
	Issuer string
    // The alias for the issuer
	IssuerAlias string
}

// The ``Info`` class contains information about an existing security token issuer trust.
//
//  The structure includes an issuer and a list of token signing certificate chains.
type SecurityTokenIssuersInfo struct {
    // Who created and signed the security token. 
    //
    //  A unique string used by the security token issuer and added to all tokens. It must match the issuer field in the security token. 
    //
    //  For example: 
    //
    // * vCenter SAML issuer: https://vsphere1.example.com/websso/SAML2/Metadata/vsphere.local
    // * vCenter JWT issuer: https://vsphere2.example.com/openidconnect/vsphere.local
	Issuer string
    // Token signing certificate chains. 
    //
    //  These certificate chains are used only for verifying existing tokens and **not** for issuing new tokens on behalf of the external security token issuer.
	SigningCertChains []SecurityTokenIssuersX509CertChain
}

// The ``Summary`` class contains summary from the list of existing security token issuer trusts.
//
//  The structure includes the alias identifier and the issuer string.
type SecurityTokenIssuersSummary struct {
    // Defines the verbosity of the sumary.
	SummaryType SecurityTokenIssuersSummaryType
    // The security token issuer trust identifier.
	IssuerAlias *string
    // Who created and signed the security token.
	Issuer *string
    // Token signing certificate chains. 
    //
    //  These certificate chains are used only for verifying existing tokens and **not** for issuing new tokens on behalf of the external security token issuer.
	SigningCertChains []SecurityTokenIssuersX509CertChain
}

// The ``CreateSpec`` class contains fields to be specified for creating a new security token issuer trust. The structure includes an alias identifier, an issuer and a list of certificate chains.
type SecurityTokenIssuersCreateSpec struct {
    // A unique string selected by the user.
	IssuerAlias string
    // Who created and signed the security token. 
    //
    //  A unique string used by the security token issuer and added to all tokens. It must match the issuer field in the security token. 
    //
    //  For example: 
    //
    // * vCenter SAML issuer: https://vsphere1.example.com/websso/SAML2/Metadata/vsphere.local
    // * vCenter JWT issuer: https://vsphere2.example.com/openidconnect/vsphere.local
	Issuer string
    // List of token signing certificate chains. Private key is not included. JWT and SAML token issuer certificates are supported with public-private keys. Symmetric keys are not supported.
	SigningCertChains []SecurityTokenIssuersX509CertChain
}

// The ``UpdateSpec`` class contains the fields of the existing security token issuer trust which can be updated.
//
//  The structure includes an issuer and a list of token signing certificate chains.
type SecurityTokenIssuersUpdateSpec struct {
    // Who created and signed the security token.
	Issuer *string
    // Token signing certificate chains.
	SigningCertChains []SecurityTokenIssuersX509CertChain
}



func securityTokenIssuersListInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["projection"] = bindings.NewOptionalType(bindings.NewEnumType("com.vmware.esx.authentication.trust.security_token_issuers.summary_type", reflect.TypeOf(SecurityTokenIssuersSummaryType(SecurityTokenIssuersSummaryType_FULL))))
	fieldNameMap["projection"] = "Projection"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func securityTokenIssuersListOutputType() bindings.BindingType {
	return bindings.NewListType(bindings.NewReferenceType(SecurityTokenIssuersSummaryBindingType), reflect.TypeOf([]SecurityTokenIssuersSummary{}))
}

func securityTokenIssuersListRestMetadata() protocol.OperationRestMetadata {
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	paramsTypeMap["projection"] = bindings.NewOptionalType(bindings.NewEnumType("com.vmware.esx.authentication.trust.security_token_issuers.summary_type", reflect.TypeOf(SecurityTokenIssuersSummaryType(SecurityTokenIssuersSummaryType_FULL))))
	queryParams["projection"] = "projection"
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
		"/esx/authentication/trust/security-token-issuers",
		resultHeaders,
		200,
		errorHeaders,
		map[string]int{"Error": 500,"InvalidArgument": 400,"Unauthenticated": 401})
}

func securityTokenIssuersCreateInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["spec"] = bindings.NewReferenceType(SecurityTokenIssuersCreateSpecBindingType)
	fieldNameMap["spec"] = "Spec"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func securityTokenIssuersCreateOutputType() bindings.BindingType {
	return bindings.NewVoidType()
}

func securityTokenIssuersCreateRestMetadata() protocol.OperationRestMetadata {
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	paramsTypeMap["spec"] = bindings.NewReferenceType(SecurityTokenIssuersCreateSpecBindingType)
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
		"/esx/authentication/trust/security-token-issuers",
		resultHeaders,
		204,
		errorHeaders,
		map[string]int{"AlreadyExists": 400,"InvalidArgument": 400,"Error": 500,"Unauthenticated": 401})
}

func securityTokenIssuersGetInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["issuer_alias"] = bindings.NewIdType([]string{"com.vmware.esx.authentication.trust.security-token-issuer"}, "")
	fieldNameMap["issuer_alias"] = "IssuerAlias"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func securityTokenIssuersGetOutputType() bindings.BindingType {
	return bindings.NewReferenceType(SecurityTokenIssuersInfoBindingType)
}

func securityTokenIssuersGetRestMetadata() protocol.OperationRestMetadata {
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	paramsTypeMap["issuer_alias"] = bindings.NewIdType([]string{"com.vmware.esx.authentication.trust.security-token-issuer"}, "")
	paramsTypeMap["issuerAlias"] = bindings.NewIdType([]string{"com.vmware.esx.authentication.trust.security-token-issuer"}, "")
	pathParams["issuer_alias"] = "issuerAlias"
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
		"/esx/authentication/trust/security-token-issuers/{issuerAlias}",
		resultHeaders,
		200,
		errorHeaders,
		map[string]int{"NotFound": 404,"Error": 500,"Unauthenticated": 401})
}

func securityTokenIssuersUpdateInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["issuer_alias"] = bindings.NewIdType([]string{"com.vmware.esx.authentication.trust.security-token-issuer"}, "")
	fields["spec"] = bindings.NewReferenceType(SecurityTokenIssuersUpdateSpecBindingType)
	fieldNameMap["issuer_alias"] = "IssuerAlias"
	fieldNameMap["spec"] = "Spec"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func securityTokenIssuersUpdateOutputType() bindings.BindingType {
	return bindings.NewVoidType()
}

func securityTokenIssuersUpdateRestMetadata() protocol.OperationRestMetadata {
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	paramsTypeMap["spec"] = bindings.NewReferenceType(SecurityTokenIssuersUpdateSpecBindingType)
	paramsTypeMap["issuer_alias"] = bindings.NewIdType([]string{"com.vmware.esx.authentication.trust.security-token-issuer"}, "")
	paramsTypeMap["issuerAlias"] = bindings.NewIdType([]string{"com.vmware.esx.authentication.trust.security-token-issuer"}, "")
	pathParams["issuer_alias"] = "issuerAlias"
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
		"/esx/authentication/trust/security-token-issuers/{issuerAlias}",
		resultHeaders,
		204,
		errorHeaders,
		map[string]int{"NotFound": 404,"Error": 500,"InvalidArgument": 400,"Unauthenticated": 401})
}

func securityTokenIssuersDeleteInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["issuer_alias"] = bindings.NewIdType([]string{"com.vmware.esx.authentication.trust.security-token-issuer"}, "")
	fieldNameMap["issuer_alias"] = "IssuerAlias"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func securityTokenIssuersDeleteOutputType() bindings.BindingType {
	return bindings.NewVoidType()
}

func securityTokenIssuersDeleteRestMetadata() protocol.OperationRestMetadata {
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	paramsTypeMap["issuer_alias"] = bindings.NewIdType([]string{"com.vmware.esx.authentication.trust.security-token-issuer"}, "")
	paramsTypeMap["issuerAlias"] = bindings.NewIdType([]string{"com.vmware.esx.authentication.trust.security-token-issuer"}, "")
	pathParams["issuer_alias"] = "issuerAlias"
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
		"/esx/authentication/trust/security-token-issuers/{issuerAlias}",
		resultHeaders,
		204,
		errorHeaders,
		map[string]int{"NotFound": 404,"Error": 500,"Unauthenticated": 401})
}


func SecurityTokenIssuersX509CertChainBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["cert_chain"] = bindings.NewListType(bindings.NewStringType(), reflect.TypeOf([]string{}))
	fieldNameMap["cert_chain"] = "CertChain"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.esx.authentication.trust.security_token_issuers.x509_cert_chain", fields, reflect.TypeOf(SecurityTokenIssuersX509CertChain{}), fieldNameMap, validators)
}

func SecurityTokenIssuersIssuerAlreadyExistsInfoBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["issuer"] = bindings.NewStringType()
	fieldNameMap["issuer"] = "Issuer"
	fields["issuer_alias"] = bindings.NewStringType()
	fieldNameMap["issuer_alias"] = "IssuerAlias"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.esx.authentication.trust.security_token_issuers.issuer_already_exists_info", fields, reflect.TypeOf(SecurityTokenIssuersIssuerAlreadyExistsInfo{}), fieldNameMap, validators)
}

func SecurityTokenIssuersInfoBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["issuer"] = bindings.NewStringType()
	fieldNameMap["issuer"] = "Issuer"
	fields["signing_cert_chains"] = bindings.NewListType(bindings.NewReferenceType(SecurityTokenIssuersX509CertChainBindingType), reflect.TypeOf([]SecurityTokenIssuersX509CertChain{}))
	fieldNameMap["signing_cert_chains"] = "SigningCertChains"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.esx.authentication.trust.security_token_issuers.info", fields, reflect.TypeOf(SecurityTokenIssuersInfo{}), fieldNameMap, validators)
}

func SecurityTokenIssuersSummaryBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["summary_type"] = bindings.NewEnumType("com.vmware.esx.authentication.trust.security_token_issuers.summary_type", reflect.TypeOf(SecurityTokenIssuersSummaryType(SecurityTokenIssuersSummaryType_FULL)))
	fieldNameMap["summary_type"] = "SummaryType"
	fields["issuer_alias"] = bindings.NewOptionalType(bindings.NewIdType([]string{"com.vmware.esx.authentication.trust.security-token-issuer"}, ""))
	fieldNameMap["issuer_alias"] = "IssuerAlias"
	fields["issuer"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["issuer"] = "Issuer"
	fields["signing_cert_chains"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewReferenceType(SecurityTokenIssuersX509CertChainBindingType), reflect.TypeOf([]SecurityTokenIssuersX509CertChain{})))
	fieldNameMap["signing_cert_chains"] = "SigningCertChains"
	var validators = []bindings.Validator{}
	uv1 := bindings.NewUnionValidator("summary_type",
		map[string][]bindings.FieldData{
			"BRIEF": []bindings.FieldData{
				bindings.NewFieldData("issuer_alias", true),
			},
			"NORMAL": []bindings.FieldData{
				bindings.NewFieldData("issuer_alias", true),
				bindings.NewFieldData("issuer", true),
			},
			"FULL": []bindings.FieldData{
				bindings.NewFieldData("issuer_alias", true),
				bindings.NewFieldData("issuer", true),
				bindings.NewFieldData("signing_cert_chains", true),
			},
		},
	)
	validators = append(validators, uv1)
	return bindings.NewStructType("com.vmware.esx.authentication.trust.security_token_issuers.summary", fields, reflect.TypeOf(SecurityTokenIssuersSummary{}), fieldNameMap, validators)
}

func SecurityTokenIssuersCreateSpecBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["issuer_alias"] = bindings.NewIdType([]string{"com.vmware.esx.authentication.trust.security-token-issuer"}, "")
	fieldNameMap["issuer_alias"] = "IssuerAlias"
	fields["issuer"] = bindings.NewStringType()
	fieldNameMap["issuer"] = "Issuer"
	fields["signing_cert_chains"] = bindings.NewListType(bindings.NewReferenceType(SecurityTokenIssuersX509CertChainBindingType), reflect.TypeOf([]SecurityTokenIssuersX509CertChain{}))
	fieldNameMap["signing_cert_chains"] = "SigningCertChains"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.esx.authentication.trust.security_token_issuers.create_spec", fields, reflect.TypeOf(SecurityTokenIssuersCreateSpec{}), fieldNameMap, validators)
}

func SecurityTokenIssuersUpdateSpecBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["issuer"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["issuer"] = "Issuer"
	fields["signing_cert_chains"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewReferenceType(SecurityTokenIssuersX509CertChainBindingType), reflect.TypeOf([]SecurityTokenIssuersX509CertChain{})))
	fieldNameMap["signing_cert_chains"] = "SigningCertChains"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.esx.authentication.trust.security_token_issuers.update_spec", fields, reflect.TypeOf(SecurityTokenIssuersUpdateSpec{}), fieldNameMap, validators)
}

