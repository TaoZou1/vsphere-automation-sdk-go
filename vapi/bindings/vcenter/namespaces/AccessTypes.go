/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Code generated. DO NOT EDIT.

/*
 * Data type definitions file for service: Access.
 * Includes binding types of a structures and enumerations defined in the service.
 * Shared by client-side stubs and server-side skeletons to ensure type
 * compatibility.
 */

package namespaces

import (
	"reflect"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/bindings"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/data"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/protocol"
)



// The ``Role`` enumeration class lists the default roles which can be associated with a subject on a domain on the namespace. **Warning:** This enumeration is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
//
// <p> See {@link com.vmware.vapi.bindings.ApiEnumeration enumerated types description}.
type AccessRole string

const (
    // This role allows modification of the namespace. **Warning:** This constant field is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
	AccessRole_EDIT AccessRole = "EDIT"
    // This is a read-only role on the namespace. **Warning:** This constant field is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
	AccessRole_VIEW AccessRole = "VIEW"
)

func (r AccessRole) AccessRole() bool {
	switch r {
	case AccessRole_EDIT:
		return true
	case AccessRole_VIEW:
		return true
	default:
		return false
	}
}


// The ``SubjectType`` enumeration class lists the types of subjects who can be associated with a ``Role`` on the namespace. **Warning:** This enumeration is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
//
// <p> See {@link com.vmware.vapi.bindings.ApiEnumeration enumerated types description}.
type AccessSubjectType string

const (
    // Single user. **Warning:** This constant field is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
	AccessSubjectType_USER AccessSubjectType = "USER"
    // Group of users. **Warning:** This constant field is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
	AccessSubjectType_GROUP AccessSubjectType = "GROUP"
)

func (s AccessSubjectType) AccessSubjectType() bool {
	switch s {
	case AccessSubjectType_USER:
		return true
	case AccessSubjectType_GROUP:
		return true
	default:
		return false
	}
}


// The ``Info`` class contains the information about the access control of the subject on given domain on the namespace. **Warning:** This class is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
type AccessInfo struct {
    // Type of the subject. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
	SubjectType AccessSubjectType
    // Role of the subject on the namespace. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
	Role AccessRole
}

// The ``CreateSpec`` class contains the specification required to create access control on the namespace. **Warning:** This class is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
type AccessCreateSpec struct {
    // Type of the principal. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
	Type_ AccessSubjectType
    // Role to be assigned. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
	Role AccessRole
}

// The ``SetSpec`` class contains the specification required to set new access control on the namespace. **Warning:** This class is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
type AccessSetSpec struct {
    // Role to be assigned. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
	Role AccessRole
}



func accessCreateInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["namespace"] = bindings.NewIdType([]string{"com.vmware.vcenter.namespaces.Instance"}, "")
	fields["domain"] = bindings.NewStringType()
	fields["subject"] = bindings.NewStringType()
	fields["spec"] = bindings.NewReferenceType(AccessCreateSpecBindingType)
	fieldNameMap["namespace"] = "Namespace"
	fieldNameMap["domain"] = "Domain"
	fieldNameMap["subject"] = "Subject"
	fieldNameMap["spec"] = "Spec"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func accessCreateOutputType() bindings.BindingType {
	return bindings.NewVoidType()
}

func accessCreateRestMetadata() protocol.OperationRestMetadata {
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	paramsTypeMap["spec"] = bindings.NewReferenceType(AccessCreateSpecBindingType)
	paramsTypeMap["subject"] = bindings.NewStringType()
	paramsTypeMap["domain"] = bindings.NewStringType()
	paramsTypeMap["namespace"] = bindings.NewIdType([]string{"com.vmware.vcenter.namespaces.Instance"}, "")
	paramsTypeMap["namespace"] = bindings.NewIdType([]string{"com.vmware.vcenter.namespaces.Instance"}, "")
	paramsTypeMap["domain"] = bindings.NewStringType()
	paramsTypeMap["subject"] = bindings.NewStringType()
	pathParams["subject"] = "subject"
	pathParams["domain"] = "domain"
	pathParams["namespace"] = "namespace"
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
		"/vcenter/namespaces/instances/{namespace}/access/{domain}/{subject}",
		resultHeaders,
		204,
		errorHeaders,
		map[string]int{"AlreadyExists": 400,"Error": 500,"NotAllowedInCurrentState": 400,"NotFound": 404,"Unauthenticated": 401,"Unauthorized": 403})
}

func accessDeleteInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["namespace"] = bindings.NewIdType([]string{"com.vmware.vcenter.namespaces.Instance"}, "")
	fields["domain"] = bindings.NewStringType()
	fields["subject"] = bindings.NewStringType()
	fieldNameMap["namespace"] = "Namespace"
	fieldNameMap["domain"] = "Domain"
	fieldNameMap["subject"] = "Subject"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func accessDeleteOutputType() bindings.BindingType {
	return bindings.NewVoidType()
}

func accessDeleteRestMetadata() protocol.OperationRestMetadata {
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	paramsTypeMap["subject"] = bindings.NewStringType()
	paramsTypeMap["domain"] = bindings.NewStringType()
	paramsTypeMap["namespace"] = bindings.NewIdType([]string{"com.vmware.vcenter.namespaces.Instance"}, "")
	paramsTypeMap["namespace"] = bindings.NewIdType([]string{"com.vmware.vcenter.namespaces.Instance"}, "")
	paramsTypeMap["domain"] = bindings.NewStringType()
	paramsTypeMap["subject"] = bindings.NewStringType()
	pathParams["subject"] = "subject"
	pathParams["domain"] = "domain"
	pathParams["namespace"] = "namespace"
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
		"/vcenter/namespaces/instances/{namespace}/access/{domain}/{subject}",
		resultHeaders,
		204,
		errorHeaders,
		map[string]int{"Error": 500,"NotAllowedInCurrentState": 400,"NotFound": 404,"Unsupported": 400,"Unauthenticated": 401,"Unauthorized": 403})
}

func accessSetInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["namespace"] = bindings.NewIdType([]string{"com.vmware.vcenter.namespaces.Instance"}, "")
	fields["domain"] = bindings.NewStringType()
	fields["subject"] = bindings.NewStringType()
	fields["spec"] = bindings.NewReferenceType(AccessSetSpecBindingType)
	fieldNameMap["namespace"] = "Namespace"
	fieldNameMap["domain"] = "Domain"
	fieldNameMap["subject"] = "Subject"
	fieldNameMap["spec"] = "Spec"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func accessSetOutputType() bindings.BindingType {
	return bindings.NewVoidType()
}

func accessSetRestMetadata() protocol.OperationRestMetadata {
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	paramsTypeMap["spec"] = bindings.NewReferenceType(AccessSetSpecBindingType)
	paramsTypeMap["subject"] = bindings.NewStringType()
	paramsTypeMap["domain"] = bindings.NewStringType()
	paramsTypeMap["namespace"] = bindings.NewIdType([]string{"com.vmware.vcenter.namespaces.Instance"}, "")
	paramsTypeMap["namespace"] = bindings.NewIdType([]string{"com.vmware.vcenter.namespaces.Instance"}, "")
	paramsTypeMap["domain"] = bindings.NewStringType()
	paramsTypeMap["subject"] = bindings.NewStringType()
	pathParams["subject"] = "subject"
	pathParams["domain"] = "domain"
	pathParams["namespace"] = "namespace"
	resultHeaders := map[string]string{}
	errorHeaders := map[string]string{}
	errorHeaders["Unauthenticated.challenge"] = "WWW-Authenticate"
	return protocol.NewOperationRestMetadata(
		paramsTypeMap,
		pathParams,
		queryParams,
		headerParams,
		"spec",
		"PUT",
		"/vcenter/namespaces/instances/{namespace}/access/{domain}/{subject}",
		resultHeaders,
		204,
		errorHeaders,
		map[string]int{"Error": 500,"NotAllowedInCurrentState": 400,"NotFound": 404,"Unsupported": 400,"Unauthenticated": 401,"Unauthorized": 403})
}

func accessGetInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["namespace"] = bindings.NewIdType([]string{"com.vmware.vcenter.namespaces.Instance"}, "")
	fields["domain"] = bindings.NewStringType()
	fields["subject"] = bindings.NewStringType()
	fieldNameMap["namespace"] = "Namespace"
	fieldNameMap["domain"] = "Domain"
	fieldNameMap["subject"] = "Subject"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func accessGetOutputType() bindings.BindingType {
	return bindings.NewReferenceType(AccessInfoBindingType)
}

func accessGetRestMetadata() protocol.OperationRestMetadata {
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	paramsTypeMap["subject"] = bindings.NewStringType()
	paramsTypeMap["domain"] = bindings.NewStringType()
	paramsTypeMap["namespace"] = bindings.NewIdType([]string{"com.vmware.vcenter.namespaces.Instance"}, "")
	paramsTypeMap["namespace"] = bindings.NewIdType([]string{"com.vmware.vcenter.namespaces.Instance"}, "")
	paramsTypeMap["domain"] = bindings.NewStringType()
	paramsTypeMap["subject"] = bindings.NewStringType()
	pathParams["subject"] = "subject"
	pathParams["domain"] = "domain"
	pathParams["namespace"] = "namespace"
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
		"/vcenter/namespaces/instances/{namespace}/access/{domain}/{subject}",
		resultHeaders,
		200,
		errorHeaders,
		map[string]int{"Error": 500,"NotFound": 404,"Unsupported": 400,"Unauthenticated": 401,"Unauthorized": 403})
}


func AccessInfoBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["subject_type"] = bindings.NewEnumType("com.vmware.vcenter.namespaces.access.subject_type", reflect.TypeOf(AccessSubjectType(AccessSubjectType_USER)))
	fieldNameMap["subject_type"] = "SubjectType"
	fields["role"] = bindings.NewEnumType("com.vmware.vcenter.namespaces.access.role", reflect.TypeOf(AccessRole(AccessRole_EDIT)))
	fieldNameMap["role"] = "Role"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vcenter.namespaces.access.info", fields, reflect.TypeOf(AccessInfo{}), fieldNameMap, validators)
}

func AccessCreateSpecBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["type"] = bindings.NewEnumType("com.vmware.vcenter.namespaces.access.subject_type", reflect.TypeOf(AccessSubjectType(AccessSubjectType_USER)))
	fieldNameMap["type"] = "Type_"
	fields["role"] = bindings.NewEnumType("com.vmware.vcenter.namespaces.access.role", reflect.TypeOf(AccessRole(AccessRole_EDIT)))
	fieldNameMap["role"] = "Role"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vcenter.namespaces.access.create_spec", fields, reflect.TypeOf(AccessCreateSpec{}), fieldNameMap, validators)
}

func AccessSetSpecBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["role"] = bindings.NewEnumType("com.vmware.vcenter.namespaces.access.role", reflect.TypeOf(AccessRole(AccessRole_EDIT)))
	fieldNameMap["role"] = "Role"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vcenter.namespaces.access.set_spec", fields, reflect.TypeOf(AccessSetSpec{}), fieldNameMap, validators)
}

