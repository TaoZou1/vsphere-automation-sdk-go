/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Code generated. DO NOT EDIT.

/*
 * Data type definitions file for service: ForeignSecurityPrincipals.
 * Includes binding types of a structures and enumerations defined in the service.
 * Shared by client-side stubs and server-side skeletons to ensure type
 * compatibility.
 */

package identity

import (
	"reflect"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/bindings"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/data"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/protocol"
)



// The ``Info`` class contains information about a foreign security principal. **Warning:** This class is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
type ForeignSecurityPrincipalsInfo struct {
    // The identifier of the principal. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
	Id string
    // The identity provider that this principal belongs to. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
	IdentityProviderId string
    // The name of the principal. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
	Name string
    // The domain that the principal belongs to. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
	Domain string
    // The names of groups that the principal is a member of. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
	GroupNames []string
    // The ids of groups that the principal is a member of. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
	GroupIds []string
}

// The ``CreateSpec`` class contains information to create a foreign security principal. **Warning:** This class is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
type ForeignSecurityPrincipalsCreateSpec struct {
    // The identity provider that this principal belongs to. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
	IdentityProviderId string
    // The name of the principal. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
	Name string
    // The domain that the principal belongs to. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
	Domain string
    // The names of groups that the principal is a member of. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
	GroupNames []string
    // The ids of groups that the principal is a member of. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
	GroupIds []string
}

// The ``UpdateSpec`` class contains information to update a foreign security principal. **Warning:** This class is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
type ForeignSecurityPrincipalsUpdateSpec struct {
    // The name of the principal. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
	Name *string
    // The domain that the principal belongs to. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
	Domain *string
    // The names of groups that the principal is a member of. If the list is empty, deletes all groups. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
	GroupNames []string
    // The ids of groups that the principal is a member of. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
	GroupIds []string
}



func foreignSecurityPrincipalsGetInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["principal"] = bindings.NewStringType()
	fieldNameMap["principal"] = "Principal"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func foreignSecurityPrincipalsGetOutputType() bindings.BindingType {
	return bindings.NewReferenceType(ForeignSecurityPrincipalsInfoBindingType)
}

func foreignSecurityPrincipalsGetRestMetadata() protocol.OperationRestMetadata {
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
		map[string]int{"Unauthorized": 403,"NotFound": 404,"InvalidArgument": 400})
}

func foreignSecurityPrincipalsCreateInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["id"] = bindings.NewIdType([]string{"com.vmware.vcenter.identity.ForeignSecurityPrincipals"}, "")
	fields["spec"] = bindings.NewReferenceType(ForeignSecurityPrincipalsCreateSpecBindingType)
	fieldNameMap["id"] = "Id"
	fieldNameMap["spec"] = "Spec"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func foreignSecurityPrincipalsCreateOutputType() bindings.BindingType {
	return bindings.NewVoidType()
}

func foreignSecurityPrincipalsCreateRestMetadata() protocol.OperationRestMetadata {
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
		map[string]int{"Unauthorized": 403,"InvalidArgument": 400})
}

func foreignSecurityPrincipalsUpdateInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["id"] = bindings.NewIdType([]string{"com.vmware.vcenter.identity.ForeignSecurityPrincipals"}, "")
	fields["spec"] = bindings.NewReferenceType(ForeignSecurityPrincipalsUpdateSpecBindingType)
	fieldNameMap["id"] = "Id"
	fieldNameMap["spec"] = "Spec"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func foreignSecurityPrincipalsUpdateOutputType() bindings.BindingType {
	return bindings.NewVoidType()
}

func foreignSecurityPrincipalsUpdateRestMetadata() protocol.OperationRestMetadata {
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
		map[string]int{"Unauthorized": 403,"NotFound": 404,"InvalidArgument": 400})
}

func foreignSecurityPrincipalsListDomainsInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func foreignSecurityPrincipalsListDomainsOutputType() bindings.BindingType {
	return bindings.NewSetType(bindings.NewStringType(), reflect.TypeOf(map[string]bool{}))
}

func foreignSecurityPrincipalsListDomainsRestMetadata() protocol.OperationRestMetadata {
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


func ForeignSecurityPrincipalsInfoBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["id"] = bindings.NewIdType([]string{"com.vmware.vcenter.identity.ForeignSecurityPrincipals"}, "")
	fieldNameMap["id"] = "Id"
	fields["identity_provider_id"] = bindings.NewStringType()
	fieldNameMap["identity_provider_id"] = "IdentityProviderId"
	fields["name"] = bindings.NewStringType()
	fieldNameMap["name"] = "Name"
	fields["domain"] = bindings.NewStringType()
	fieldNameMap["domain"] = "Domain"
	fields["group_names"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewStringType(), reflect.TypeOf([]string{})))
	fieldNameMap["group_names"] = "GroupNames"
	fields["group_ids"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewStringType(), reflect.TypeOf([]string{})))
	fieldNameMap["group_ids"] = "GroupIds"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vcenter.identity.foreign_security_principals.info", fields, reflect.TypeOf(ForeignSecurityPrincipalsInfo{}), fieldNameMap, validators)
}

func ForeignSecurityPrincipalsCreateSpecBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["identity_provider_id"] = bindings.NewStringType()
	fieldNameMap["identity_provider_id"] = "IdentityProviderId"
	fields["name"] = bindings.NewStringType()
	fieldNameMap["name"] = "Name"
	fields["domain"] = bindings.NewStringType()
	fieldNameMap["domain"] = "Domain"
	fields["group_names"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewStringType(), reflect.TypeOf([]string{})))
	fieldNameMap["group_names"] = "GroupNames"
	fields["group_ids"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewStringType(), reflect.TypeOf([]string{})))
	fieldNameMap["group_ids"] = "GroupIds"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vcenter.identity.foreign_security_principals.create_spec", fields, reflect.TypeOf(ForeignSecurityPrincipalsCreateSpec{}), fieldNameMap, validators)
}

func ForeignSecurityPrincipalsUpdateSpecBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["name"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["name"] = "Name"
	fields["domain"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["domain"] = "Domain"
	fields["group_names"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewStringType(), reflect.TypeOf([]string{})))
	fieldNameMap["group_names"] = "GroupNames"
	fields["group_ids"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewStringType(), reflect.TypeOf([]string{})))
	fieldNameMap["group_ids"] = "GroupIds"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vcenter.identity.foreign_security_principals.update_spec", fields, reflect.TypeOf(ForeignSecurityPrincipalsUpdateSpec{}), fieldNameMap, validators)
}

