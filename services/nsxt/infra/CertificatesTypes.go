/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Code generated. DO NOT EDIT.

/*
 * Data type definitions file for service: Certificates.
 * Includes binding types of a structures and enumerations defined in the service.
 * Shared by client-side stubs and server-side skeletons to ensure type
 * compatibility.
 */

package infra

import (
	"reflect"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/services/nsxt/model"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/bindings"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/data"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/protocol"
)

// Possible value for ``type`` of method Certificates#list.
const Certificates_LIST_TYPE_CERTIFICATE = "cluster_api_certificate"




func certificatesDeleteInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["certificate_id"] = bindings.NewStringType()
	fieldNameMap["certificate_id"] = "CertificateId"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func certificatesDeleteOutputType() bindings.BindingType {
	return bindings.NewVoidType()
}

func certificatesDeleteRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	fields["certificate_id"] = bindings.NewStringType()
	fieldNameMap["certificate_id"] = "CertificateId"
	paramsTypeMap["certificate_id"] = bindings.NewStringType()
	paramsTypeMap["certificateId"] = bindings.NewStringType()
	pathParams["certificate_id"] = "certificateId"
	resultHeaders := map[string]string{}
	errorHeaders := map[string]string{}
	return protocol.NewOperationRestMetadata(
		fields,
		fieldNameMap,
		paramsTypeMap,
		pathParams,
		queryParams,
		headerParams,
		"",
		"",
		"DELETE",
		"/policy/api/v1/infra/certificates/{certificateId}",
		resultHeaders,
		204,
		errorHeaders,
		map[string]int{"InvalidRequest": 400,"Unauthorized": 403,"ServiceUnavailable": 503,"InternalServerError": 500,"NotFound": 404})
}

func certificatesGetInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["certificate_id"] = bindings.NewStringType()
	fields["details"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fieldNameMap["certificate_id"] = "CertificateId"
	fieldNameMap["details"] = "Details"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func certificatesGetOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.TlsCertificateBindingType)
}

func certificatesGetRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	fields["certificate_id"] = bindings.NewStringType()
	fields["details"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fieldNameMap["certificate_id"] = "CertificateId"
	fieldNameMap["details"] = "Details"
	paramsTypeMap["certificate_id"] = bindings.NewStringType()
	paramsTypeMap["details"] = bindings.NewOptionalType(bindings.NewBooleanType())
	paramsTypeMap["certificateId"] = bindings.NewStringType()
	pathParams["certificate_id"] = "certificateId"
	queryParams["details"] = "details"
	resultHeaders := map[string]string{}
	errorHeaders := map[string]string{}
	return protocol.NewOperationRestMetadata(
		fields,
		fieldNameMap,
		paramsTypeMap,
		pathParams,
		queryParams,
		headerParams,
		"",
		"",
		"GET",
		"/policy/api/v1/infra/certificates/{certificateId}",
		resultHeaders,
		200,
		errorHeaders,
		map[string]int{"InvalidRequest": 400,"Unauthorized": 403,"ServiceUnavailable": 503,"InternalServerError": 500,"NotFound": 404})
}

func certificatesListInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["cursor"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["details"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fields["included_fields"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["page_size"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fields["sort_ascending"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fields["sort_by"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["type"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["cursor"] = "Cursor"
	fieldNameMap["details"] = "Details"
	fieldNameMap["included_fields"] = "IncludedFields"
	fieldNameMap["page_size"] = "PageSize"
	fieldNameMap["sort_ascending"] = "SortAscending"
	fieldNameMap["sort_by"] = "SortBy"
	fieldNameMap["type"] = "Type_"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func certificatesListOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.TlsCertificateListBindingType)
}

func certificatesListRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	fields["cursor"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["details"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fields["included_fields"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["page_size"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fields["sort_ascending"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fields["sort_by"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["type"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["cursor"] = "Cursor"
	fieldNameMap["details"] = "Details"
	fieldNameMap["included_fields"] = "IncludedFields"
	fieldNameMap["page_size"] = "PageSize"
	fieldNameMap["sort_ascending"] = "SortAscending"
	fieldNameMap["sort_by"] = "SortBy"
	fieldNameMap["type"] = "Type_"
	paramsTypeMap["included_fields"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["page_size"] = bindings.NewOptionalType(bindings.NewIntegerType())
	paramsTypeMap["type"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["cursor"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["sort_by"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["details"] = bindings.NewOptionalType(bindings.NewBooleanType())
	paramsTypeMap["sort_ascending"] = bindings.NewOptionalType(bindings.NewBooleanType())
	queryParams["cursor"] = "cursor"
	queryParams["sort_ascending"] = "sort_ascending"
	queryParams["included_fields"] = "included_fields"
	queryParams["details"] = "details"
	queryParams["sort_by"] = "sort_by"
	queryParams["type"] = "type"
	queryParams["page_size"] = "page_size"
	resultHeaders := map[string]string{}
	errorHeaders := map[string]string{}
	return protocol.NewOperationRestMetadata(
		fields,
		fieldNameMap,
		paramsTypeMap,
		pathParams,
		queryParams,
		headerParams,
		"",
		"",
		"GET",
		"/policy/api/v1/infra/certificates",
		resultHeaders,
		200,
		errorHeaders,
		map[string]int{"InvalidRequest": 400,"Unauthorized": 403,"ServiceUnavailable": 503,"InternalServerError": 500,"NotFound": 404})
}

func certificatesPatchInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["certificate_id"] = bindings.NewStringType()
	fields["tls_trust_data"] = bindings.NewReferenceType(model.TlsTrustDataBindingType)
	fieldNameMap["certificate_id"] = "CertificateId"
	fieldNameMap["tls_trust_data"] = "TlsTrustData"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func certificatesPatchOutputType() bindings.BindingType {
	return bindings.NewVoidType()
}

func certificatesPatchRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	fields["certificate_id"] = bindings.NewStringType()
	fields["tls_trust_data"] = bindings.NewReferenceType(model.TlsTrustDataBindingType)
	fieldNameMap["certificate_id"] = "CertificateId"
	fieldNameMap["tls_trust_data"] = "TlsTrustData"
	paramsTypeMap["certificate_id"] = bindings.NewStringType()
	paramsTypeMap["tls_trust_data"] = bindings.NewReferenceType(model.TlsTrustDataBindingType)
	paramsTypeMap["certificateId"] = bindings.NewStringType()
	pathParams["certificate_id"] = "certificateId"
	resultHeaders := map[string]string{}
	errorHeaders := map[string]string{}
	return protocol.NewOperationRestMetadata(
		fields,
		fieldNameMap,
		paramsTypeMap,
		pathParams,
		queryParams,
		headerParams,
		"",
		"tls_trust_data",
		"PATCH",
		"/policy/api/v1/infra/certificates/{certificateId}",
		resultHeaders,
		204,
		errorHeaders,
		map[string]int{"InvalidRequest": 400,"Unauthorized": 403,"ServiceUnavailable": 503,"InternalServerError": 500,"NotFound": 404})
}

func certificatesUpdateInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["certificate_id"] = bindings.NewStringType()
	fields["tls_trust_data"] = bindings.NewReferenceType(model.TlsTrustDataBindingType)
	fieldNameMap["certificate_id"] = "CertificateId"
	fieldNameMap["tls_trust_data"] = "TlsTrustData"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func certificatesUpdateOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.TlsCertificateBindingType)
}

func certificatesUpdateRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	fields["certificate_id"] = bindings.NewStringType()
	fields["tls_trust_data"] = bindings.NewReferenceType(model.TlsTrustDataBindingType)
	fieldNameMap["certificate_id"] = "CertificateId"
	fieldNameMap["tls_trust_data"] = "TlsTrustData"
	paramsTypeMap["certificate_id"] = bindings.NewStringType()
	paramsTypeMap["tls_trust_data"] = bindings.NewReferenceType(model.TlsTrustDataBindingType)
	paramsTypeMap["certificateId"] = bindings.NewStringType()
	pathParams["certificate_id"] = "certificateId"
	resultHeaders := map[string]string{}
	errorHeaders := map[string]string{}
	return protocol.NewOperationRestMetadata(
		fields,
		fieldNameMap,
		paramsTypeMap,
		pathParams,
		queryParams,
		headerParams,
		"",
		"tls_trust_data",
		"PUT",
		"/policy/api/v1/infra/certificates/{certificateId}",
		resultHeaders,
		200,
		errorHeaders,
		map[string]int{"InvalidRequest": 400,"Unauthorized": 403,"ServiceUnavailable": 503,"InternalServerError": 500,"NotFound": 404})
}

