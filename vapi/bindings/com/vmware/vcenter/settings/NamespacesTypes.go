/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Data type definitions file for service: Namespaces.
 * Includes binding types of a structures and enumerations defined in the service.
 * Shared by client-side stubs and server-side skeletons to ensure type
 * compatibility.
 */
package settings

import (
	"reflect"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/bindings"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/data"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/protocol"
)

// Resource type for Settings Namespaces. **Warning:** This constant field is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
const Namespaces_RESOURCE_TYPE = "com.vmware.vcenter.settings.namespaces"
// Info class describing a namespace in the Settings Store. **Warning:** This class is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
type NamespacesInfo struct {
    // The name given to the namespace. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
	Name string
    // A description of this namespace. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
	Description string
    // The privilege that is required for reading into this namespace. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
	ReadPrivilege string
    // The privilege that is required for writing into this namespace. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
	WritePrivilege string
    // The flag for enable sync. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
	IsSyncable bool
}

// The ``CreateSpec`` class contains specifications used to create a namespace. **Warning:** This class is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
type NamespacesCreateSpec struct {
    // The namespace name. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
	Name string
    // Description of the namespace. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
	Description *string
    // The privilege required for a user to read this namespace. The user must be assigned a global permission containing this privilege. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
	ReadPrivilege string
    // The privilege required for a user to update values in this namespace. The user must be assigned a global permission containing this privilege. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
	WritePrivilege string
    // The boolean flag to indicate if the namespace is syncable. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
	Syncable bool
}

// Spec used to update the namespace itself, not its associated data items. **Warning:** This class is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
type NamespacesUpdateSpec struct {
    // Description for the namespace. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
	Description *string
    // Privilege required for reading this namespace. The user must be assigned a global permission containing this privilege to take effect. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
	ReadPrivilege *string
    // Privilege required for updating values in this namespace. The user must be assigned a global permission containing this privilege to take effect. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
	WritePrivilege *string
    // Syncable flag this flag enables / disables sync of the namespace when appliances are linked. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
	Syncable *bool
}


func namespacesListInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func namespacesListOutputType() bindings.BindingType {
	return bindings.NewListType(bindings.NewReferenceType(NamespacesInfoBindingType), reflect.TypeOf([]NamespacesInfo{}))
}

func namespacesListRestMetadata() protocol.OperationRestMetadata {
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
		map[string]int{"Unauthorized": 403,"Error": 500})
}

func namespacesGetInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["namespace"] = bindings.NewStringType()
	fieldNameMap["namespace"] = "Namespace"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func namespacesGetOutputType() bindings.BindingType {
	return bindings.NewReferenceType(NamespacesInfoBindingType)
}

func namespacesGetRestMetadata() protocol.OperationRestMetadata {
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
		map[string]int{"Unauthorized": 403,"Error": 500})
}

func namespacesCreateInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["spec"] = bindings.NewReferenceType(NamespacesCreateSpecBindingType)
	fieldNameMap["spec"] = "Spec"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func namespacesCreateOutputType() bindings.BindingType {
	return bindings.NewVoidType()
}

func namespacesCreateRestMetadata() protocol.OperationRestMetadata {
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
		map[string]int{"AlreadyExists": 400,"Unauthorized": 403,"InvalidArgument": 400,"Error": 500})
}

func namespacesDeleteInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["namespace"] = bindings.NewIdType([]string{"com.vmware.vcenter.settings.namespaces"}, "")
	fieldNameMap["namespace"] = "Namespace"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func namespacesDeleteOutputType() bindings.BindingType {
	return bindings.NewVoidType()
}

func namespacesDeleteRestMetadata() protocol.OperationRestMetadata {
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
		map[string]int{"Unauthorized": 403,"Error": 500})
}

func namespacesUpdateInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["namespace"] = bindings.NewIdType([]string{"com.vmware.vcenter.settings.namespaces"}, "")
	fields["spec"] = bindings.NewReferenceType(NamespacesUpdateSpecBindingType)
	fieldNameMap["namespace"] = "Namespace"
	fieldNameMap["spec"] = "Spec"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func namespacesUpdateOutputType() bindings.BindingType {
	return bindings.NewVoidType()
}

func namespacesUpdateRestMetadata() protocol.OperationRestMetadata {
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
		map[string]int{"NotFound": 404,"Unauthorized": 403,"InvalidArgument": 400,"Error": 500})
}


func NamespacesInfoBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["name"] = bindings.NewStringType()
	fieldNameMap["name"] = "Name"
	fields["description"] = bindings.NewStringType()
	fieldNameMap["description"] = "Description"
	fields["read_privilege"] = bindings.NewStringType()
	fieldNameMap["read_privilege"] = "ReadPrivilege"
	fields["write_privilege"] = bindings.NewStringType()
	fieldNameMap["write_privilege"] = "WritePrivilege"
	fields["is_syncable"] = bindings.NewBooleanType()
	fieldNameMap["is_syncable"] = "IsSyncable"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vcenter.settings.namespaces.info", fields, reflect.TypeOf(NamespacesInfo{}), fieldNameMap, validators)
}

func NamespacesCreateSpecBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["name"] = bindings.NewStringType()
	fieldNameMap["name"] = "Name"
	fields["description"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["description"] = "Description"
	fields["read_privilege"] = bindings.NewStringType()
	fieldNameMap["read_privilege"] = "ReadPrivilege"
	fields["write_privilege"] = bindings.NewStringType()
	fieldNameMap["write_privilege"] = "WritePrivilege"
	fields["syncable"] = bindings.NewBooleanType()
	fieldNameMap["syncable"] = "Syncable"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vcenter.settings.namespaces.create_spec", fields, reflect.TypeOf(NamespacesCreateSpec{}), fieldNameMap, validators)
}

func NamespacesUpdateSpecBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["description"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["description"] = "Description"
	fields["read_privilege"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["read_privilege"] = "ReadPrivilege"
	fields["write_privilege"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["write_privilege"] = "WritePrivilege"
	fields["syncable"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fieldNameMap["syncable"] = "Syncable"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vcenter.settings.namespaces.update_spec", fields, reflect.TypeOf(NamespacesUpdateSpec{}), fieldNameMap, validators)
}

