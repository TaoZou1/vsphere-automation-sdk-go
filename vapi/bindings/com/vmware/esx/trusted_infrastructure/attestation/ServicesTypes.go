/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Data type definitions file for service: Services.
 * Includes binding types of a structures and enumerations defined in the service.
 * Shared by client-side stubs and server-side skeletons to ensure type
 * compatibility.
 */
package attestation

import (
	"reflect"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/bindings/com/vmware/esx/trusted_infrastructure"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/bindings"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/data"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/protocol"
)

// The resource type for the Services interface.
const Services_RESOURCE_TYPE = "com.vmware.esx.trusted_infrastructure.attestation.service"
// The ``Summary`` class contains a summary of an attestation service instance.
type ServicesSummary struct {
    // The service's unique identifier.
	Service string
    // The service's address.
	Address trusted_infrastructure.NetworkAddress
    // The group identifier determines which KMS service instances this attestation service can communicate with.
    //
    //  If the group attribute values of attestation and KMS records match then those services are associated with each other.
	Group string
}

// The ``Info`` class contains all the stored information about an attestation service instance.
type ServicesInfo struct {
    // The service's address.
	Address trusted_infrastructure.NetworkAddress
    // The service's TLS certificate chain.
	TrustedCA trusted_infrastructure.X509CertChain
    // The group identifier determines which KMS service instances this attestation service can communicate with.
    //
    //  If the group attribute values of attestation and KMS records match then those services are associated with each other.
	Group string
}

// The ``CreateSpec`` class contains the data necessary for adding an attestation service instance to the environment.
type ServicesCreateSpec struct {
    // The service's address.
	Address trusted_infrastructure.NetworkAddress
    // The service's TLS certificate chain.
	TrustedCA trusted_infrastructure.X509CertChain
    // The group identifier determines which KMS service instances this attestation service can communicate with.
    //
    //  If the group attribute values of attestation and KMS records match then those services are associated with each other.
	Group string
}

// The ``FilterSpec`` class contains the data necessary for identifying an attestation service instance.
type ServicesFilterSpec struct {
    // A set of identifiers by which to filter the services.
	Services map[string]bool
    // A set of addresses by which to filter.
	Address []trusted_infrastructure.NetworkAddress
    // The group identifier determines which KMS service instances this attestation service can communicate with.
    //
    //  If the group attribute values of attestation and KMS records match then those services are associated with each other.
	Group map[string]bool
}


func servicesListInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["filter"] = bindings.NewOptionalType(bindings.NewReferenceType(ServicesFilterSpecBindingType))
	fieldNameMap["filter"] = "Filter"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func servicesListOutputType() bindings.BindingType {
	return bindings.NewListType(bindings.NewReferenceType(ServicesSummaryBindingType), reflect.TypeOf([]ServicesSummary{}))
}

func servicesListRestMetadata() protocol.OperationRestMetadata {
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	paramsTypeMap["filter"] = bindings.NewOptionalType(bindings.NewReferenceType(ServicesFilterSpecBindingType))
	resultHeaders := map[string]string{}
	errorHeaders := map[string]string{}
	errorHeaders["Unauthenticated.challenge"] = "WWW-Authenticate"
	return protocol.NewOperationRestMetadata(
		paramsTypeMap,
		pathParams,
		queryParams,
		headerParams,
		"filter",
		"POST",
		"/esx/trusted-infrastructure/attestation/services",
		resultHeaders,
		200,
		errorHeaders,
		map[string]int{"Error": 500,"Unauthenticated": 401})
}

func servicesGetInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["service"] = bindings.NewIdType([]string{"com.vmware.esx.trusted_infrastructure.attestation.service"}, "")
	fieldNameMap["service"] = "Service"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func servicesGetOutputType() bindings.BindingType {
	return bindings.NewReferenceType(ServicesInfoBindingType)
}

func servicesGetRestMetadata() protocol.OperationRestMetadata {
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	paramsTypeMap["service"] = bindings.NewIdType([]string{"com.vmware.esx.trusted_infrastructure.attestation.service"}, "")
	paramsTypeMap["service"] = bindings.NewIdType([]string{"com.vmware.esx.trusted_infrastructure.attestation.service"}, "")
	pathParams["service"] = "service"
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
		"/esx/trusted-infrastructure/attestation/services/{service}",
		resultHeaders,
		200,
		errorHeaders,
		map[string]int{"NotFound": 404,"Error": 500,"Unauthenticated": 401})
}

func servicesCreateInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["spec"] = bindings.NewReferenceType(ServicesCreateSpecBindingType)
	fieldNameMap["spec"] = "Spec"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func servicesCreateOutputType() bindings.BindingType {
	return bindings.NewIdType([]string{"com.vmware.esx.trusted_infrastructure.attestation.service"}, "")
}

func servicesCreateRestMetadata() protocol.OperationRestMetadata {
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	paramsTypeMap["spec"] = bindings.NewReferenceType(ServicesCreateSpecBindingType)
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
		"/esx/trusted-infrastructure/attestation/services",
		resultHeaders,
		201,
		errorHeaders,
		map[string]int{"AlreadyExists": 400,"InvalidArgument": 400,"UnableToAllocateResource": 500,"Error": 500,"Unauthenticated": 401})
}

func servicesDeleteInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["service"] = bindings.NewIdType([]string{"com.vmware.esx.trusted_infrastructure.attestation.service"}, "")
	fieldNameMap["service"] = "Service"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func servicesDeleteOutputType() bindings.BindingType {
	return bindings.NewVoidType()
}

func servicesDeleteRestMetadata() protocol.OperationRestMetadata {
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	paramsTypeMap["service"] = bindings.NewIdType([]string{"com.vmware.esx.trusted_infrastructure.attestation.service"}, "")
	paramsTypeMap["service"] = bindings.NewIdType([]string{"com.vmware.esx.trusted_infrastructure.attestation.service"}, "")
	pathParams["service"] = "service"
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
		"/esx/trusted-infrastructure/attestation/services/{service}",
		resultHeaders,
		204,
		errorHeaders,
		map[string]int{"NotFound": 404,"Error": 500,"Unauthenticated": 401})
}


func ServicesSummaryBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["service"] = bindings.NewIdType([]string{"com.vmware.esx.trusted_infrastructure.attestation.service"}, "")
	fieldNameMap["service"] = "Service"
	fields["address"] = bindings.NewReferenceType(trusted_infrastructure.NetworkAddressBindingType)
	fieldNameMap["address"] = "Address"
	fields["group"] = bindings.NewStringType()
	fieldNameMap["group"] = "Group"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.esx.trusted_infrastructure.attestation.services.summary", fields, reflect.TypeOf(ServicesSummary{}), fieldNameMap, validators)
}

func ServicesInfoBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["address"] = bindings.NewReferenceType(trusted_infrastructure.NetworkAddressBindingType)
	fieldNameMap["address"] = "Address"
	fields["trusted_CA"] = bindings.NewReferenceType(trusted_infrastructure.X509CertChainBindingType)
	fieldNameMap["trusted_CA"] = "TrustedCA"
	fields["group"] = bindings.NewStringType()
	fieldNameMap["group"] = "Group"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.esx.trusted_infrastructure.attestation.services.info", fields, reflect.TypeOf(ServicesInfo{}), fieldNameMap, validators)
}

func ServicesCreateSpecBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["address"] = bindings.NewReferenceType(trusted_infrastructure.NetworkAddressBindingType)
	fieldNameMap["address"] = "Address"
	fields["trusted_CA"] = bindings.NewReferenceType(trusted_infrastructure.X509CertChainBindingType)
	fieldNameMap["trusted_CA"] = "TrustedCA"
	fields["group"] = bindings.NewStringType()
	fieldNameMap["group"] = "Group"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.esx.trusted_infrastructure.attestation.services.create_spec", fields, reflect.TypeOf(ServicesCreateSpec{}), fieldNameMap, validators)
}

func ServicesFilterSpecBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["services"] = bindings.NewOptionalType(bindings.NewSetType(bindings.NewIdType([]string{"com.vmware.esx.trusted_infrastructure.attestation.service"}, ""), reflect.TypeOf(map[string]bool{})))
	fieldNameMap["services"] = "Services"
	fields["address"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewReferenceType(trusted_infrastructure.NetworkAddressBindingType), reflect.TypeOf([]trusted_infrastructure.NetworkAddress{})))
	fieldNameMap["address"] = "Address"
	fields["group"] = bindings.NewOptionalType(bindings.NewSetType(bindings.NewStringType(), reflect.TypeOf(map[string]bool{})))
	fieldNameMap["group"] = "Group"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.esx.trusted_infrastructure.attestation.services.filter_spec", fields, reflect.TypeOf(ServicesFilterSpec{}), fieldNameMap, validators)
}

