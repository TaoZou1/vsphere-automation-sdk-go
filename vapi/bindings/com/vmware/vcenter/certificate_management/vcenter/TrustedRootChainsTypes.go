/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Data type definitions file for service: TrustedRootChains.
 * Includes binding types of a structures and enumerations defined in the service.
 * Shared by client-side stubs and server-side skeletons to ensure type
 * compatibility.
 */
package vcenter

import (
	"reflect"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/bindings/com/vmware/vcenter/certificate_management"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/bindings"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/data"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/protocol"
)

// The ``Info`` class contains information for a trusted root certificate chain.
type TrustedRootChainsInfo struct {
    // A certificate chain in base64 encoding
	CertChain certificate_management.X509CertChain
}

// The ``Summary`` class contains a trusted root certificate chain summary suitable for UI presentation.
type TrustedRootChainsSummary struct {
    // Unique identifier for chain.
	Chain string
}

// The ``CreateSpec`` class contains information to create a trusted root certificate chain.
type TrustedRootChainsCreateSpec struct {
    // Certificate chain in base64 encoding
	CertChain certificate_management.X509CertChain
    // Unique identifier for this trusted root. Client can specify at creation as long as it is unique, otherwise one will be generated. An example of a client providing the identifier would be if this trusted root is associated with a VC trust. In this case the identifier would be the domain id.
	Chain *string
}


func trustedRootChainsListInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func trustedRootChainsListOutputType() bindings.BindingType {
	return bindings.NewListType(bindings.NewReferenceType(TrustedRootChainsSummaryBindingType), reflect.TypeOf([]TrustedRootChainsSummary{}))
}

func trustedRootChainsListRestMetadata() protocol.OperationRestMetadata {
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

func trustedRootChainsCreateInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["spec"] = bindings.NewReferenceType(TrustedRootChainsCreateSpecBindingType)
	fieldNameMap["spec"] = "Spec"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func trustedRootChainsCreateOutputType() bindings.BindingType {
	return bindings.NewStringType()
}

func trustedRootChainsCreateRestMetadata() protocol.OperationRestMetadata {
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
		map[string]int{"Unauthorized": 403,"AlreadyExists": 400})
}

func trustedRootChainsGetInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["chain"] = bindings.NewStringType()
	fieldNameMap["chain"] = "Chain"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func trustedRootChainsGetOutputType() bindings.BindingType {
	return bindings.NewReferenceType(TrustedRootChainsInfoBindingType)
}

func trustedRootChainsGetRestMetadata() protocol.OperationRestMetadata {
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

func trustedRootChainsDeleteInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["chain"] = bindings.NewStringType()
	fieldNameMap["chain"] = "Chain"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func trustedRootChainsDeleteOutputType() bindings.BindingType {
	return bindings.NewVoidType()
}

func trustedRootChainsDeleteRestMetadata() protocol.OperationRestMetadata {
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


func TrustedRootChainsInfoBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["cert_chain"] = bindings.NewReferenceType(certificate_management.X509CertChainBindingType)
	fieldNameMap["cert_chain"] = "CertChain"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vcenter.certificate_management.vcenter.trusted_root_chains.info", fields, reflect.TypeOf(TrustedRootChainsInfo{}), fieldNameMap, validators)
}

func TrustedRootChainsSummaryBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["chain"] = bindings.NewStringType()
	fieldNameMap["chain"] = "Chain"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vcenter.certificate_management.vcenter.trusted_root_chains.summary", fields, reflect.TypeOf(TrustedRootChainsSummary{}), fieldNameMap, validators)
}

func TrustedRootChainsCreateSpecBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["cert_chain"] = bindings.NewReferenceType(certificate_management.X509CertChainBindingType)
	fieldNameMap["cert_chain"] = "CertChain"
	fields["chain"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["chain"] = "Chain"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vcenter.certificate_management.vcenter.trusted_root_chains.create_spec", fields, reflect.TypeOf(TrustedRootChainsCreateSpec{}), fieldNameMap, validators)
}
