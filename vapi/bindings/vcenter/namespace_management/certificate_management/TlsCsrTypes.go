/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Data type definitions file for service: TlsCsr.
 * Includes binding types of a structures and enumerations defined in the service.
 * Shared by client-side stubs and server-side skeletons to ensure type
 * compatibility.
 */
package certificate_management

import (
	"reflect"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/bindings"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/data"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/protocol"
)



// The ``Spec`` class contains the specification required to create new Certificate Signing Request for the Cluster. **Warning:** This class is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
type TlsCsrSpec struct {
    // Size of key associated with the new Certificate Signing Request. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
	KeySize *int64
    // Name identifying the machine. Its value will be set in Certificate Signing Request Subject's CN field. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
	CommonName string
    // Two-letter country code (as defined in ISO-3166) to be set in Certificate Signing Request Subject's C field. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
	Country string
    // State or province to be set in Certificate Signing Request Subject's ST field. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
	StateOrProvince string
    // Locality (city) to be set in Certificate Signing Request Subject's L field. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
	Locality string
    // Organization to be set in Certificate Signing Request Subject's O field. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
	OrganizationName string
    // Organizational Unit to be set in Certificate Signing Request Subject's OU field. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
	OrganizationUnitName string
    // Email to be set in Certificate Signing Request Subject Alternative Names. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
	EmailAddress string
}



func tlsCsrCreateInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["cluster"] = bindings.NewIdType([]string{"ClusterComputeResource"}, "")
	fields["spec"] = bindings.NewReferenceType(TlsCsrSpecBindingType)
	fieldNameMap["cluster"] = "Cluster"
	fieldNameMap["spec"] = "Spec"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func tlsCsrCreateOutputType() bindings.BindingType {
	return bindings.NewStringType()
}

func tlsCsrCreateRestMetadata() protocol.OperationRestMetadata {
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	paramsTypeMap["cluster"] = bindings.NewIdType([]string{"ClusterComputeResource"}, "")
	paramsTypeMap["spec"] = bindings.NewReferenceType(TlsCsrSpecBindingType)
	paramsTypeMap["cluster"] = bindings.NewIdType([]string{"ClusterComputeResource"}, "")
	pathParams["cluster"] = "cluster"
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
		"/{cluster}/csr/tls-endpoint",
		resultHeaders,
		201,
		errorHeaders,
		map[string]int{"Error": 500,"InvalidArgument": 400,"NotFound": 404,"Unauthenticated": 401,"Unauthorized": 403})
}

func tlsCsrCreateNcpDefaultIngressTlsInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["cluster"] = bindings.NewIdType([]string{"ClusterComputeResource"}, "")
	fields["spec"] = bindings.NewReferenceType(TlsCsrSpecBindingType)
	fieldNameMap["cluster"] = "Cluster"
	fieldNameMap["spec"] = "Spec"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func tlsCsrCreateNcpDefaultIngressTlsOutputType() bindings.BindingType {
	return bindings.NewStringType()
}

func tlsCsrCreateNcpDefaultIngressTlsRestMetadata() protocol.OperationRestMetadata {
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	paramsTypeMap["cluster"] = bindings.NewIdType([]string{"ClusterComputeResource"}, "")
	paramsTypeMap["spec"] = bindings.NewReferenceType(TlsCsrSpecBindingType)
	paramsTypeMap["cluster"] = bindings.NewIdType([]string{"ClusterComputeResource"}, "")
	pathParams["cluster"] = "cluster"
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
		"/{cluster}/csr/ncp-default-ingress-tls",
		resultHeaders,
		200,
		errorHeaders,
		map[string]int{"Error": 500,"InvalidArgument": 400,"NotFound": 404,"Unauthenticated": 401,"Unauthorized": 403})
}


func TlsCsrSpecBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["key_size"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fieldNameMap["key_size"] = "KeySize"
	fields["common_name"] = bindings.NewStringType()
	fieldNameMap["common_name"] = "CommonName"
	fields["country"] = bindings.NewStringType()
	fieldNameMap["country"] = "Country"
	fields["state_or_province"] = bindings.NewStringType()
	fieldNameMap["state_or_province"] = "StateOrProvince"
	fields["locality"] = bindings.NewStringType()
	fieldNameMap["locality"] = "Locality"
	fields["organization_name"] = bindings.NewStringType()
	fieldNameMap["organization_name"] = "OrganizationName"
	fields["organization_unit_name"] = bindings.NewStringType()
	fieldNameMap["organization_unit_name"] = "OrganizationUnitName"
	fields["email_address"] = bindings.NewStringType()
	fieldNameMap["email_address"] = "EmailAddress"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vcenter.namespace_management.certificate_management.tls_csr.spec", fields, reflect.TypeOf(TlsCsrSpec{}), fieldNameMap, validators)
}

