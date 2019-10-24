/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Data type definitions file for service: Members.
 * Includes binding types of a structures and enumerations defined in the service.
 * Shared by client-side stubs and server-side skeletons to ensure type
 * compatibility.
 */
package project

import (
	"reflect"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/bindings/com/vmware/vapi/std"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/bindings"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/data"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/protocol"
)


// The ``MemberType`` enumeration class describes the type of project member to be created for a Harbor project. **Warning:** This enumeration is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
//
// <p> See {@link com.vmware.vapi.bindings.ApiEnumeration enumerated types description}.
type MembersMemberType string

const (
    // Project member type for an individual user. **Warning:** This constant field is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
	MembersMemberType_USER MembersMemberType = "USER"
    // Project member type for a group of users. **Warning:** This constant field is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
	MembersMemberType_GROUP MembersMemberType = "GROUP"
)

func (m MembersMemberType) MembersMemberType() bool {
	switch m {
	case MembersMemberType_USER:
		return true
	case MembersMemberType_GROUP:
		return true
	default:
		return false
	}
}


// The ``Role`` enumeration class describes roles available in a Harbor project. **Warning:** This enumeration is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
//
// <p> See {@link com.vmware.vapi.bindings.ApiEnumeration enumerated types description}.
type MembersRole string

const (
    // This role allows image pull and push, repository and image deletion on an associated project,. **Warning:** This constant field is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
	MembersRole_MASTER MembersRole = "MASTER"
    // This role allows image pull on an associated project. **Warning:** This constant field is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
	MembersRole_GUEST MembersRole = "GUEST"
)

func (r MembersRole) MembersRole() bool {
	switch r {
	case MembersRole_MASTER:
		return true
	case MembersRole_GUEST:
		return true
	default:
		return false
	}
}


// The ``ConfigStatus`` enumeration class describes the status of reaching the desired configuration state for the Harbor project member. **Warning:** This enumeration is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
//
// <p> See {@link com.vmware.vapi.bindings.ApiEnumeration enumerated types description}.
type MembersConfigStatus string

const (
    // Harbor project member is being created or the configuration is being applied to the project member. **Warning:** This constant field is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
	MembersConfigStatus_PENDING MembersConfigStatus = "PENDING"
    // The configuration is being removed and Harbor project member is being deleted. **Warning:** This constant field is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
	MembersConfigStatus_REMOVING MembersConfigStatus = "REMOVING"
    // Failed to create Harbor project member or apply the configuration to the project member, user intervention needed. **Warning:** This constant field is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
	MembersConfigStatus_ERROR MembersConfigStatus = "ERROR"
    // Harbor project member is created or configured correctly. **Warning:** This constant field is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
	MembersConfigStatus_READY MembersConfigStatus = "READY"
)

func (c MembersConfigStatus) MembersConfigStatus() bool {
	switch c {
	case MembersConfigStatus_PENDING:
		return true
	case MembersConfigStatus_REMOVING:
		return true
	case MembersConfigStatus_ERROR:
		return true
	case MembersConfigStatus_READY:
		return true
	default:
		return false
	}
}

// The ``Summary`` class contains basic Harbor project member information. **Warning:** This class is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
type MembersSummary struct {
    // Identifier of the harbor project member. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
	Member string
    // **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
	MemberName string
}

// The ``Info`` class contains detailed Harbor project member information. **Warning:** This class is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
type MembersInfo struct {
    // **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
	MemberName string
    // Type of the project member. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
	Type_ MembersMemberType
    // Role of the member. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
	Role MembersRole
    // Status of the Harbor project member. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
	ConfigStatus MembersConfigStatus
    // Details about the project member ERROR status. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
	Message *std.LocalizableMessage
}

// The ``CreateSpec`` class contains the specification required to create project member for a specified Harbor project. **Warning:** This class is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
type MembersCreateSpec struct {
    // **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
	MemberName string
    // Type of project member to be created. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
	Type_ MembersMemberType
    // Role to be assigned. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
	Role MembersRole
}

// The ``UpdateSpec`` class contains the specification required to set new project role for an existing project member. More fields could be added in future release. **Warning:** This class is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
type MembersUpdateSpec struct {
    // Role to be assigned. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
	Role *MembersRole
}


func membersCreateInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["registry"] = bindings.NewIdType([]string{"com.vmware.vcenter.content.Registry"}, "")
	fields["project"] = bindings.NewIdType([]string{"com.vmware.vcenter.content.Registry.Harbor.Project"}, "")
	fields["spec"] = bindings.NewReferenceType(MembersCreateSpecBindingType)
	fieldNameMap["registry"] = "Registry"
	fieldNameMap["project"] = "Project"
	fieldNameMap["spec"] = "Spec"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func membersCreateOutputType() bindings.BindingType {
	return bindings.NewIdType([]string{"com.vmware.vcenter.content.Registry.Harbor.Project.Member"}, "")
}

func membersCreateRestMetadata() protocol.OperationRestMetadata {
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
		map[string]int{"InvalidArgument": 400,"NotAllowedInCurrentState": 400,"NotFound": 404,"AlreadyExists": 400,"Error": 500,"Unauthenticated": 401,"Unauthorized": 403})
}

func membersDeleteInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["registry"] = bindings.NewIdType([]string{"com.vmware.vcenter.content.Registry"}, "")
	fields["project"] = bindings.NewIdType([]string{"com.vmware.vcenter.content.Registry.Harbor.Project"}, "")
	fields["member"] = bindings.NewIdType([]string{"com.vmware.vcenter.content.Registry.Harbor.Project.Member"}, "")
	fieldNameMap["registry"] = "Registry"
	fieldNameMap["project"] = "Project"
	fieldNameMap["member"] = "Member"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func membersDeleteOutputType() bindings.BindingType {
	return bindings.NewVoidType()
}

func membersDeleteRestMetadata() protocol.OperationRestMetadata {
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
		map[string]int{"NotAllowedInCurrentState": 400,"NotFound": 404,"Error": 500,"Unauthenticated": 401,"Unauthorized": 403})
}

func membersUpdateInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["registry"] = bindings.NewIdType([]string{"com.vmware.vcenter.content.Registry"}, "")
	fields["project"] = bindings.NewIdType([]string{"com.vmware.vcenter.content.Registry.Harbor.Project"}, "")
	fields["member"] = bindings.NewIdType([]string{"com.vmware.vcenter.content.Registry.Harbor.Project.Member"}, "")
	fields["spec"] = bindings.NewReferenceType(MembersUpdateSpecBindingType)
	fieldNameMap["registry"] = "Registry"
	fieldNameMap["project"] = "Project"
	fieldNameMap["member"] = "Member"
	fieldNameMap["spec"] = "Spec"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func membersUpdateOutputType() bindings.BindingType {
	return bindings.NewVoidType()
}

func membersUpdateRestMetadata() protocol.OperationRestMetadata {
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
		map[string]int{"InvalidArgument": 400,"NotAllowedInCurrentState": 400,"NotFound": 404,"Error": 500,"Unauthenticated": 401,"Unauthorized": 403})
}

func membersGetInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["registry"] = bindings.NewIdType([]string{"com.vmware.vcenter.content.Registry"}, "")
	fields["project"] = bindings.NewIdType([]string{"com.vmware.vcenter.content.Registry.Harbor.Project"}, "")
	fields["member"] = bindings.NewIdType([]string{"com.vmware.vcenter.content.Registry.Harbor.Project.Member"}, "")
	fieldNameMap["registry"] = "Registry"
	fieldNameMap["project"] = "Project"
	fieldNameMap["member"] = "Member"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func membersGetOutputType() bindings.BindingType {
	return bindings.NewReferenceType(MembersInfoBindingType)
}

func membersGetRestMetadata() protocol.OperationRestMetadata {
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
		map[string]int{"NotFound": 404,"Error": 500,"Unauthenticated": 401,"Unauthorized": 403})
}

func membersListInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["registry"] = bindings.NewIdType([]string{"com.vmware.vcenter.content.Registry"}, "")
	fields["project"] = bindings.NewIdType([]string{"com.vmware.vcenter.content.Registry.Harbor.Project"}, "")
	fieldNameMap["registry"] = "Registry"
	fieldNameMap["project"] = "Project"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func membersListOutputType() bindings.BindingType {
	return bindings.NewListType(bindings.NewReferenceType(MembersSummaryBindingType), reflect.TypeOf([]MembersSummary{}))
}

func membersListRestMetadata() protocol.OperationRestMetadata {
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
		map[string]int{"Error": 500,"NotFound": 404,"Unauthenticated": 401,"Unauthorized": 403})
}


func MembersSummaryBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["member"] = bindings.NewIdType([]string{"com.vmware.vcenter.content.Registry.Harbor.Project.Member"}, "")
	fieldNameMap["member"] = "Member"
	fields["member_name"] = bindings.NewStringType()
	fieldNameMap["member_name"] = "MemberName"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vcenter.content.registries.harbor.project.members.summary", fields, reflect.TypeOf(MembersSummary{}), fieldNameMap, validators)
}

func MembersInfoBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["member_name"] = bindings.NewStringType()
	fieldNameMap["member_name"] = "MemberName"
	fields["type"] = bindings.NewEnumType("com.vmware.vcenter.content.registries.harbor.project.members.member_type", reflect.TypeOf(MembersMemberType(MembersMemberType_USER)))
	fieldNameMap["type"] = "Type_"
	fields["role"] = bindings.NewEnumType("com.vmware.vcenter.content.registries.harbor.project.members.role", reflect.TypeOf(MembersRole(MembersRole_MASTER)))
	fieldNameMap["role"] = "Role"
	fields["config_status"] = bindings.NewEnumType("com.vmware.vcenter.content.registries.harbor.project.members.config_status", reflect.TypeOf(MembersConfigStatus(MembersConfigStatus_PENDING)))
	fieldNameMap["config_status"] = "ConfigStatus"
	fields["message"] = bindings.NewOptionalType(bindings.NewReferenceType(std.LocalizableMessageBindingType))
	fieldNameMap["message"] = "Message"
	var validators = []bindings.Validator{}
	uv1 := bindings.NewUnionValidator("config_status",
	    map[string][]bindings.FieldData{
	        "ERROR": []bindings.FieldData{
	             bindings.NewFieldData("message", true),
	        },
	        "PENDING": []bindings.FieldData{},
	        "REMOVING": []bindings.FieldData{},
	        "READY": []bindings.FieldData{},
	    },
	)
	validators = append(validators, uv1)
	return bindings.NewStructType("com.vmware.vcenter.content.registries.harbor.project.members.info", fields, reflect.TypeOf(MembersInfo{}), fieldNameMap, validators)
}

func MembersCreateSpecBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["member_name"] = bindings.NewStringType()
	fieldNameMap["member_name"] = "MemberName"
	fields["type"] = bindings.NewEnumType("com.vmware.vcenter.content.registries.harbor.project.members.member_type", reflect.TypeOf(MembersMemberType(MembersMemberType_USER)))
	fieldNameMap["type"] = "Type_"
	fields["role"] = bindings.NewEnumType("com.vmware.vcenter.content.registries.harbor.project.members.role", reflect.TypeOf(MembersRole(MembersRole_MASTER)))
	fieldNameMap["role"] = "Role"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vcenter.content.registries.harbor.project.members.create_spec", fields, reflect.TypeOf(MembersCreateSpec{}), fieldNameMap, validators)
}

func MembersUpdateSpecBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["role"] = bindings.NewOptionalType(bindings.NewEnumType("com.vmware.vcenter.content.registries.harbor.project.members.role", reflect.TypeOf(MembersRole(MembersRole_MASTER))))
	fieldNameMap["role"] = "Role"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vcenter.content.registries.harbor.project.members.update_spec", fields, reflect.TypeOf(MembersUpdateSpec{}), fieldNameMap, validators)
}

