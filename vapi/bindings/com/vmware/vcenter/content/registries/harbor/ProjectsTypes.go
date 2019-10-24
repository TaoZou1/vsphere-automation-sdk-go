/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Data type definitions file for service: Projects.
 * Includes binding types of a structures and enumerations defined in the service.
 * Shared by client-side stubs and server-side skeletons to ensure type
 * compatibility.
 */
package harbor

import (
	"reflect"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/bindings/com/vmware/vapi/std"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/bindings/com/vmware/vcenter/content/registries/harbor/project"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/bindings"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/data"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/protocol"
	"net/url"
	"time"
)


// The ``Scope`` enumeration class in a project defines access levels of the project. **Warning:** This enumeration is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
//
// <p> See {@link com.vmware.vapi.bindings.ApiEnumeration enumerated types description}.
type ProjectsScope string

const (
    // A Harbor project can be accessed by everyone. **Warning:** This constant field is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
	ProjectsScope_PUBLIC ProjectsScope = "PUBLIC"
    // A Harbor project can only be accessed by assigned users. **Warning:** This constant field is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
	ProjectsScope_PRIVATE ProjectsScope = "PRIVATE"
)

func (s ProjectsScope) ProjectsScope() bool {
	switch s {
	case ProjectsScope_PUBLIC:
		return true
	case ProjectsScope_PRIVATE:
		return true
	default:
		return false
	}
}


// The ``ConfigStatus`` enumeration class describes the status of reaching the desired configuration state for the Harbor project. **Warning:** This enumeration is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
//
// <p> See {@link com.vmware.vapi.bindings.ApiEnumeration enumerated types description}.
type ProjectsConfigStatus string

const (
    // Harbor project is being created or the configuration is being applied to the project. **Warning:** This constant field is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
	ProjectsConfigStatus_PENDING ProjectsConfigStatus = "PENDING"
    // The configuration is being removed and Harbor project is being deleted. **Warning:** This constant field is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
	ProjectsConfigStatus_REMOVING ProjectsConfigStatus = "REMOVING"
    // Failed to create Harbor project or apply the configuration to the project, user intervention needed. **Warning:** This constant field is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
	ProjectsConfigStatus_ERROR ProjectsConfigStatus = "ERROR"
    // Harbor project is created or configured correctly. **Warning:** This constant field is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
	ProjectsConfigStatus_READY ProjectsConfigStatus = "READY"
)

func (c ProjectsConfigStatus) ProjectsConfigStatus() bool {
	switch c {
	case ProjectsConfigStatus_PENDING:
		return true
	case ProjectsConfigStatus_REMOVING:
		return true
	case ProjectsConfigStatus_ERROR:
		return true
	case ProjectsConfigStatus_READY:
		return true
	default:
		return false
	}
}

// The ``CreateSpec`` class defines the information required to create a Harbor project. **Warning:** This class is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
type ProjectsCreateSpec struct {
    // Name of the Harbor project. Should be between 2-63 characters long alphanumeric string and may contain the following characters: a-z,0-9, and '-'. Must be starting with characters or numbers, with the '-' character allowed anywhere except the first or last character. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
	Name string
    // Access type of a Harbor project. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
	Scope ProjectsScope
    // Information for Harbor project members. vSphere SSO user who created the project becomes a project member with master role. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
	Members []project.MembersCreateSpec
}

// The ``Summary`` class contains basic information about a Harbor project. **Warning:** This class is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
type ProjectsSummary struct {
    // Identifier of the harbor project. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
	Project string
    // Name of the Harbor project. Should be between 2-63 characters long alphanumeric string and may contain the following characters: a-z,0-9, and '-'. Must be starting with characters or numbers, with the '-' character allowed anywhere except the first or last character. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
	Name string
    // Access type of a Harbor project. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
	Scope ProjectsScope
}

// The ``Info`` class contains detailed information about a Harbor project. **Warning:** This class is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
type ProjectsInfo struct {
    // Name of the Harbor project. Should be between 2-63 characters long alphanumeric string and may contain the following characters: a-z,0-9, and '-'. Must be starting with characters or numbers, with the '-' character allowed anywhere except the first or last character. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
	Name string
    // The status of the Harbor project. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
	ConfigStatus ProjectsConfigStatus
    // The access type of a Harbor project. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
	Scope ProjectsScope
    // The date and time when the harbor project creation API was triggered and project identifier generated. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
	CreationTime time.Time
    // The date and time when the harbor project purge API was triggered. In case no purge was triggered, ``updateTime`` is same as ``#creationTIme``. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
	UpdateTime *time.Time
    // URL to access the harbor project through docker client. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
	AccessUrl *url.URL
    // Details about the ERROR project status. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
	Message *std.LocalizableMessage
}


func projectsCreateInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["registry"] = bindings.NewIdType([]string{"com.vmware.vcenter.content.Registry"}, "")
	fields["spec"] = bindings.NewReferenceType(ProjectsCreateSpecBindingType)
	fieldNameMap["registry"] = "Registry"
	fieldNameMap["spec"] = "Spec"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func projectsCreateOutputType() bindings.BindingType {
	return bindings.NewIdType([]string{"com.vmware.vcenter.content.Registry.Harbor.Project"}, "")
}

func projectsCreateRestMetadata() protocol.OperationRestMetadata {
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

func projectsDeleteInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["registry"] = bindings.NewIdType([]string{"com.vmware.vcenter.content.Registry"}, "")
	fields["project"] = bindings.NewIdType([]string{"com.vmware.vcenter.content.Registry.Harbor.Project"}, "")
	fieldNameMap["registry"] = "Registry"
	fieldNameMap["project"] = "Project"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func projectsDeleteOutputType() bindings.BindingType {
	return bindings.NewVoidType()
}

func projectsDeleteRestMetadata() protocol.OperationRestMetadata {
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

func projectsGetInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["registry"] = bindings.NewIdType([]string{"com.vmware.vcenter.content.Registry"}, "")
	fields["project"] = bindings.NewIdType([]string{"com.vmware.vcenter.content.Registry.Harbor.Project"}, "")
	fieldNameMap["registry"] = "Registry"
	fieldNameMap["project"] = "Project"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func projectsGetOutputType() bindings.BindingType {
	return bindings.NewReferenceType(ProjectsInfoBindingType)
}

func projectsGetRestMetadata() protocol.OperationRestMetadata {
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

func projectsListInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["registry"] = bindings.NewIdType([]string{"com.vmware.vcenter.content.Registry"}, "")
	fieldNameMap["registry"] = "Registry"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func projectsListOutputType() bindings.BindingType {
	return bindings.NewListType(bindings.NewReferenceType(ProjectsSummaryBindingType), reflect.TypeOf([]ProjectsSummary{}))
}

func projectsListRestMetadata() protocol.OperationRestMetadata {
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

func projectsPurgeInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["registry"] = bindings.NewIdType([]string{"com.vmware.vcenter.content.Registry"}, "")
	fields["project"] = bindings.NewIdType([]string{"com.vmware.vcenter.content.Registry.Harbor.Project"}, "")
	fieldNameMap["registry"] = "Registry"
	fieldNameMap["project"] = "Project"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func projectsPurgeOutputType() bindings.BindingType {
	return bindings.NewVoidType()
}

func projectsPurgeRestMetadata() protocol.OperationRestMetadata {
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
		map[string]int{"NotFound": 404,"NotAllowedInCurrentState": 400,"Error": 500,"Unauthenticated": 401,"Unauthorized": 403})
}


func ProjectsCreateSpecBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["name"] = bindings.NewStringType()
	fieldNameMap["name"] = "Name"
	fields["scope"] = bindings.NewEnumType("com.vmware.vcenter.content.registries.harbor.projects.scope", reflect.TypeOf(ProjectsScope(ProjectsScope_PUBLIC)))
	fieldNameMap["scope"] = "Scope"
	fields["members"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewReferenceType(project.MembersCreateSpecBindingType), reflect.TypeOf([]project.MembersCreateSpec{})))
	fieldNameMap["members"] = "Members"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vcenter.content.registries.harbor.projects.create_spec", fields, reflect.TypeOf(ProjectsCreateSpec{}), fieldNameMap, validators)
}

func ProjectsSummaryBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["project"] = bindings.NewIdType([]string{"com.vmware.vcenter.content.Registry.Harbor.Project"}, "")
	fieldNameMap["project"] = "Project"
	fields["name"] = bindings.NewStringType()
	fieldNameMap["name"] = "Name"
	fields["scope"] = bindings.NewEnumType("com.vmware.vcenter.content.registries.harbor.projects.scope", reflect.TypeOf(ProjectsScope(ProjectsScope_PUBLIC)))
	fieldNameMap["scope"] = "Scope"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vcenter.content.registries.harbor.projects.summary", fields, reflect.TypeOf(ProjectsSummary{}), fieldNameMap, validators)
}

func ProjectsInfoBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["name"] = bindings.NewStringType()
	fieldNameMap["name"] = "Name"
	fields["config_status"] = bindings.NewEnumType("com.vmware.vcenter.content.registries.harbor.projects.config_status", reflect.TypeOf(ProjectsConfigStatus(ProjectsConfigStatus_PENDING)))
	fieldNameMap["config_status"] = "ConfigStatus"
	fields["scope"] = bindings.NewEnumType("com.vmware.vcenter.content.registries.harbor.projects.scope", reflect.TypeOf(ProjectsScope(ProjectsScope_PUBLIC)))
	fieldNameMap["scope"] = "Scope"
	fields["creation_time"] = bindings.NewDateTimeType()
	fieldNameMap["creation_time"] = "CreationTime"
	fields["update_time"] = bindings.NewOptionalType(bindings.NewDateTimeType())
	fieldNameMap["update_time"] = "UpdateTime"
	fields["access_url"] = bindings.NewOptionalType(bindings.NewUriType())
	fieldNameMap["access_url"] = "AccessUrl"
	fields["message"] = bindings.NewOptionalType(bindings.NewReferenceType(std.LocalizableMessageBindingType))
	fieldNameMap["message"] = "Message"
	var validators = []bindings.Validator{}
	uv1 := bindings.NewUnionValidator("config_status",
	    map[string][]bindings.FieldData{
	        "READY": []bindings.FieldData{
	             bindings.NewFieldData("update_time", true),
	             bindings.NewFieldData("access_url", true),
	        },
	        "ERROR": []bindings.FieldData{
	             bindings.NewFieldData("message", true),
	        },
	        "PENDING": []bindings.FieldData{},
	        "REMOVING": []bindings.FieldData{},
	    },
	)
	validators = append(validators, uv1)
	return bindings.NewStructType("com.vmware.vcenter.content.registries.harbor.projects.info", fields, reflect.TypeOf(ProjectsInfo{}), fieldNameMap, validators)
}

