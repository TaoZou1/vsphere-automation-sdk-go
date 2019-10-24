/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Data type definitions file for service: ClientProfiles.
 * Includes binding types of a structures and enumerations defined in the service.
 * Shared by client-side stubs and server-side skeletons to ensure type
 * compatibility.
 */
package authentication

import (
	"reflect"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/bindings"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/data"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/protocol"
)

// The resource type for the ClientProfiles instances.
const ClientProfiles_RESOURCE_TYPE = "com.vmware.esx.authentication.clientprofile"

// The ``ResourceType`` enumeration class defines the types of ClientProfilesAccessGrant elements in a client profile. These are permission resource types. There is support for entitlements, but not for groups.
//
// <p> See {@link com.vmware.vapi.bindings.ApiEnumeration enumerated types description}.
type ClientProfilesResourceType string

const (
    // Permission entitlements. 
    //
    //  These are coarse-grained permissions that are not associated with an object, i.e. they are system-wide.
	ClientProfilesResourceType_ENTITLEMENT ClientProfilesResourceType = "ENTITLEMENT"
)

func (r ClientProfilesResourceType) ClientProfilesResourceType() bool {
	switch r {
	case ClientProfilesResourceType_ENTITLEMENT:
		return true
	default:
		return false
	}
}


// The ``Entitlement`` enumeration class defines all permission entitlements supported on the ESX.
//
//  These are coarse-grained permissions that are not associated with an object, i.e. they are system-wide.
//
// <p> See {@link com.vmware.vapi.bindings.ApiEnumeration enumerated types description}.
type ClientProfilesEntitlement string

const (
    // Allows modifying the identity configuration. 
    //
    //  For example: ClientProfiles, SecurityTokenIssuers.
	ClientProfilesEntitlement_IDENTITY_MGMT ClientProfilesEntitlement = "IDENTITY_MGMT"
    // Allows modifying security configuration. 
    //
    //  For example: KMS, Attestation.
	ClientProfilesEntitlement_SECURITY_MGMT ClientProfilesEntitlement = "SECURITY_MGMT"
    // Allows access to some read-only methods. Not all read-only methods are accessible with this entitlement. Check the specific method documentation for the required authorization.
	ClientProfilesEntitlement_READ_ONLY ClientProfilesEntitlement = "READ_ONLY"
)

func (e ClientProfilesEntitlement) ClientProfilesEntitlement() bool {
	switch e {
	case ClientProfilesEntitlement_IDENTITY_MGMT:
		return true
	case ClientProfilesEntitlement_SECURITY_MGMT:
		return true
	case ClientProfilesEntitlement_READ_ONLY:
		return true
	default:
		return false
	}
}


// The ``SubjectType`` enumeration class defines the types of subject matching that a client profile is associated with.
//
// <p> See {@link com.vmware.vapi.bindings.ApiEnumeration enumerated types description}.
type ClientProfilesSubjectType string

const (
    // local user
	ClientProfilesSubjectType_LOCAL_USER ClientProfilesSubjectType = "LOCAL_USER"
    // A group from external source.
	ClientProfilesSubjectType_EXTERNAL_GROUP ClientProfilesSubjectType = "EXTERNAL_GROUP"
    // A user from external source.
	ClientProfilesSubjectType_EXTERNAL_USER ClientProfilesSubjectType = "EXTERNAL_USER"
)

func (s ClientProfilesSubjectType) ClientProfilesSubjectType() bool {
	switch s {
	case ClientProfilesSubjectType_LOCAL_USER:
		return true
	case ClientProfilesSubjectType_EXTERNAL_GROUP:
		return true
	case ClientProfilesSubjectType_EXTERNAL_USER:
		return true
	default:
		return false
	}
}


// The ``SummaryType`` enumeration class defines the types of ClientProfilesSummary members to return from the ClientProfiles#list method.
//
//  The profile information could include the access grants or be a shorter summary.
//
// <p> See {@link com.vmware.vapi.bindings.ApiEnumeration enumerated types description}.
type ClientProfilesSummaryType string

const (
    // The full profile information, including access grants.
	ClientProfilesSummaryType_FULL ClientProfilesSummaryType = "FULL"
    // A summary containing only the profile identifier and the subject information.
	ClientProfilesSummaryType_NORMAL ClientProfilesSummaryType = "NORMAL"
    // A brief summary, containing only the profile identifier.
	ClientProfilesSummaryType_BRIEF ClientProfilesSummaryType = "BRIEF"
)

func (s ClientProfilesSummaryType) ClientProfilesSummaryType() bool {
	switch s {
	case ClientProfilesSummaryType_FULL:
		return true
	case ClientProfilesSummaryType_NORMAL:
		return true
	case ClientProfilesSummaryType_BRIEF:
		return true
	default:
		return false
	}
}

// The ``FilterSpec`` class contains information about filtering the list of existing ClientProfiles.
//
//  The structure includes local user, external group or external user specification.
//
//  Only one of the subject types can be specified. The filter is appied using exact match of all fields.
type ClientProfilesFilterSpec struct {
    // Filter for local user by name. 
    //
    //  Exactly one of ClientProfilesFilterSpec#localUserName, ClientProfilesFilterSpec#externalGroupName or ClientProfilesFilterSpec#externalUserName must be map with bool value.
	LocalUserName *string
    // The external group name. 
    //
    //  Exactly one of ClientProfilesFilterSpec#localUserName, ClientProfilesFilterSpec#externalGroupName or ClientProfilesFilterSpec#externalUserName must be map with bool value.
	ExternalGroupName *string
    // The external user name. 
    //
    //  Exactly one of ClientProfilesFilterSpec#localUserName, ClientProfilesFilterSpec#externalGroupName or ClientProfilesFilterSpec#externalUserName must be map with bool value.
	ExternalUserName *string
    // The security token issuer alias, who created and signed the security token.
	IssuerAlias *string
    // Domain of the external principal.
	Domain *string
}

// The ``AccessGrant`` class contains information about access permissions.
//
//  The structure includes resource type and the specific resource information - entitlement.
type ClientProfilesAccessGrant struct {
    // Type of permission entity.
	ResourceType ClientProfilesResourceType
    // The entitlement in the access grant.
	Entitlement *ClientProfilesEntitlement
}

// The ``Subject`` class contains information about the subject that a client profile is associated with.
//
//  The structure includes a name, issuer alias, and domain.
type ClientProfilesSubject struct {
    // subject type
	Type_ ClientProfilesSubjectType
    // The user or group name
	Name *string
    // The security token issuer alias, who created and signed the security token.
	IssuerAlias *string
    // Domain of the principal.
	Domain *string
}

// The ``Summary`` class contains summary from the list of existing ClientProfiles.
//
//  The structure includes a client profile identifier, subject type, the details of the subject - local user, external user or external group.
type ClientProfilesSummary struct {
    // Defines the verbosity of the sumary.
	SummaryType ClientProfilesSummaryType
    // Client profile identifier.
	ClientProfile *string
    // The subject of the profile.
	Subject *ClientProfilesSubject
    // Access grants. When the list is empty, the matching subjects don't receive grants from this client profile. but will receive grants from other matching ClientProfiles.
	Grants []ClientProfilesAccessGrant
}

// The ``Info`` class contains information about an existing client profile.
//
//  The structure includes a subject type, the details of the subject - local user, external user or external group, and a list of access grants. 
//
// * A local user is a user account configured on the ESX system.
// * An external user is a user account configured in an external for the ESX identity provider.
// * An external group is a group account configured in an external for the ESX identity provider.
type ClientProfilesInfo struct {
    // The subject of the profile.
	Subject ClientProfilesSubject
    // Access grants. When the list is empty, the matching subjects don't receive grants from this client profile. but will receive grants from other matching ClientProfiles.
	Grants []ClientProfilesAccessGrant
}

// The ``CreateSpec`` class contains fields to be specified for creating a new client profile.
//
//  The structure includes a subject - local user, external user or external group, and a list of access grants.
type ClientProfilesCreateSpec struct {
    // The local user name. 
    //
    //  Exactly one of ClientProfilesCreateSpec#localUserName, ClientProfilesCreateSpec#externalGroupName or ClientProfilesCreateSpec#externalUserName must be map with bool value.
	LocalUserName *string
    // The external group name. 
    //
    //  Exactly one of ClientProfilesCreateSpec#localUserName, ClientProfilesCreateSpec#externalGroupName or ClientProfilesCreateSpec#externalUserName must be map with bool value.
	ExternalGroupName *string
    // The external user name. 
    //
    //  Exactly one of ClientProfilesCreateSpec#localUserName, ClientProfilesCreateSpec#externalGroupName or ClientProfilesCreateSpec#externalUserName must be map with bool value.
	ExternalUserName *string
    // The security token issuer alias, who created and signed the security token.
	IssuerAlias *string
    // Domain of the external principal.
	Domain *string
    // Access grants. When the list is empty, the matching subjects don't receive grants from this client profile. but will receive grants from other matching ClientProfiles.
	Grants []ClientProfilesAccessGrant
}

// The ``UpdateSpec`` class contains the fields of the existing client profile which can be updated.
//
//  The structure includes a list of access grants.
type ClientProfilesUpdateSpec struct {
    // Access grants.
	Grants []ClientProfilesAccessGrant
}


func clientProfilesListInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["filter"] = bindings.NewOptionalType(bindings.NewReferenceType(ClientProfilesFilterSpecBindingType))
	fields["projection"] = bindings.NewOptionalType(bindings.NewEnumType("com.vmware.esx.authentication.client_profiles.summary_type", reflect.TypeOf(ClientProfilesSummaryType(ClientProfilesSummaryType_FULL))))
	fieldNameMap["filter"] = "Filter"
	fieldNameMap["projection"] = "Projection"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func clientProfilesListOutputType() bindings.BindingType {
	return bindings.NewListType(bindings.NewReferenceType(ClientProfilesSummaryBindingType), reflect.TypeOf([]ClientProfilesSummary{}))
}

func clientProfilesListRestMetadata() protocol.OperationRestMetadata {
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	paramsTypeMap["projection"] = bindings.NewOptionalType(bindings.NewEnumType("com.vmware.esx.authentication.client_profiles.summary_type", reflect.TypeOf(ClientProfilesSummaryType(ClientProfilesSummaryType_FULL))))
	paramsTypeMap["filter.domain"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["filter.external_user_name"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["filter.external_group_name"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["filter.issuer_alias"] = bindings.NewOptionalType(bindings.NewIdType([]string{"com.vmware.esx.authentication.trust.security-token-issuer"}, ""))
	paramsTypeMap["filter.local_user_name"] = bindings.NewOptionalType(bindings.NewStringType())
	queryParams["filter.external_user_name"] = "external_user_name"
	queryParams["filter.local_user_name"] = "local_user_name"
	queryParams["filter.domain"] = "domain"
	queryParams["filter.external_group_name"] = "external_group_name"
	queryParams["filter.issuer_alias"] = "issuer_alias"
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
		"/esx/authentication/client-profiles",
		resultHeaders,
		200,
		errorHeaders,
		map[string]int{"Error": 500,"InvalidArgument": 400,"Unauthenticated": 401})
}

func clientProfilesCreateInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["spec"] = bindings.NewReferenceType(ClientProfilesCreateSpecBindingType)
	fieldNameMap["spec"] = "Spec"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func clientProfilesCreateOutputType() bindings.BindingType {
	return bindings.NewIdType([]string{"com.vmware.esx.authentication.clientprofile"}, "")
}

func clientProfilesCreateRestMetadata() protocol.OperationRestMetadata {
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	paramsTypeMap["spec"] = bindings.NewReferenceType(ClientProfilesCreateSpecBindingType)
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
		"/esx/authentication/client-profiles",
		resultHeaders,
		201,
		errorHeaders,
		map[string]int{"AlreadyExists": 400,"InvalidArgument": 400,"Error": 500,"Unauthenticated": 401})
}

func clientProfilesGetInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["profile"] = bindings.NewIdType([]string{"com.vmware.esx.authentication.clientprofile"}, "")
	fieldNameMap["profile"] = "Profile"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func clientProfilesGetOutputType() bindings.BindingType {
	return bindings.NewReferenceType(ClientProfilesInfoBindingType)
}

func clientProfilesGetRestMetadata() protocol.OperationRestMetadata {
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	paramsTypeMap["profile"] = bindings.NewIdType([]string{"com.vmware.esx.authentication.clientprofile"}, "")
	paramsTypeMap["profile"] = bindings.NewIdType([]string{"com.vmware.esx.authentication.clientprofile"}, "")
	pathParams["profile"] = "profile"
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
		"/esx/authentication/client-profiles/{profile}",
		resultHeaders,
		200,
		errorHeaders,
		map[string]int{"NotFound": 404,"Error": 500,"Unauthenticated": 401})
}

func clientProfilesUpdateInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["profile"] = bindings.NewIdType([]string{"com.vmware.esx.authentication.clientprofile"}, "")
	fields["spec"] = bindings.NewReferenceType(ClientProfilesUpdateSpecBindingType)
	fieldNameMap["profile"] = "Profile"
	fieldNameMap["spec"] = "Spec"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func clientProfilesUpdateOutputType() bindings.BindingType {
	return bindings.NewVoidType()
}

func clientProfilesUpdateRestMetadata() protocol.OperationRestMetadata {
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	paramsTypeMap["spec"] = bindings.NewReferenceType(ClientProfilesUpdateSpecBindingType)
	paramsTypeMap["profile"] = bindings.NewIdType([]string{"com.vmware.esx.authentication.clientprofile"}, "")
	paramsTypeMap["profile"] = bindings.NewIdType([]string{"com.vmware.esx.authentication.clientprofile"}, "")
	pathParams["profile"] = "profile"
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
		"/esx/authentication/client-profiles/{profile}",
		resultHeaders,
		204,
		errorHeaders,
		map[string]int{"NotFound": 404,"InvalidArgument": 400,"Error": 500,"Unauthenticated": 401})
}

func clientProfilesDeleteInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["profile"] = bindings.NewIdType([]string{"com.vmware.esx.authentication.clientprofile"}, "")
	fieldNameMap["profile"] = "Profile"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func clientProfilesDeleteOutputType() bindings.BindingType {
	return bindings.NewVoidType()
}

func clientProfilesDeleteRestMetadata() protocol.OperationRestMetadata {
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	paramsTypeMap["profile"] = bindings.NewIdType([]string{"com.vmware.esx.authentication.clientprofile"}, "")
	paramsTypeMap["profile"] = bindings.NewIdType([]string{"com.vmware.esx.authentication.clientprofile"}, "")
	pathParams["profile"] = "profile"
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
		"/esx/authentication/client-profiles/{profile}",
		resultHeaders,
		204,
		errorHeaders,
		map[string]int{"NotFound": 404,"Error": 500,"Unauthenticated": 401})
}


func ClientProfilesFilterSpecBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["local_user_name"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["local_user_name"] = "LocalUserName"
	fields["external_group_name"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["external_group_name"] = "ExternalGroupName"
	fields["external_user_name"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["external_user_name"] = "ExternalUserName"
	fields["issuer_alias"] = bindings.NewOptionalType(bindings.NewIdType([]string{"com.vmware.esx.authentication.trust.security-token-issuer"}, ""))
	fieldNameMap["issuer_alias"] = "IssuerAlias"
	fields["domain"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["domain"] = "Domain"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.esx.authentication.client_profiles.filter_spec", fields, reflect.TypeOf(ClientProfilesFilterSpec{}), fieldNameMap, validators)
}

func ClientProfilesAccessGrantBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["resource_type"] = bindings.NewEnumType("com.vmware.esx.authentication.client_profiles.resource_type", reflect.TypeOf(ClientProfilesResourceType(ClientProfilesResourceType_ENTITLEMENT)))
	fieldNameMap["resource_type"] = "ResourceType"
	fields["entitlement"] = bindings.NewOptionalType(bindings.NewEnumType("com.vmware.esx.authentication.client_profiles.entitlement", reflect.TypeOf(ClientProfilesEntitlement(ClientProfilesEntitlement_IDENTITY_MGMT))))
	fieldNameMap["entitlement"] = "Entitlement"
	var validators = []bindings.Validator{}
	uv1 := bindings.NewUnionValidator("resource_type",
	    map[string][]bindings.FieldData{
	        "ENTITLEMENT": []bindings.FieldData{
	             bindings.NewFieldData("entitlement", true),
	        },
	    },
	)
	validators = append(validators, uv1)
	return bindings.NewStructType("com.vmware.esx.authentication.client_profiles.access_grant", fields, reflect.TypeOf(ClientProfilesAccessGrant{}), fieldNameMap, validators)
}

func ClientProfilesSubjectBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["type"] = bindings.NewEnumType("com.vmware.esx.authentication.client_profiles.subject_type", reflect.TypeOf(ClientProfilesSubjectType(ClientProfilesSubjectType_LOCAL_USER)))
	fieldNameMap["type"] = "Type_"
	fields["name"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["name"] = "Name"
	fields["issuer_alias"] = bindings.NewOptionalType(bindings.NewIdType([]string{"com.vmware.esx.authentication.trust.security-token-issuer"}, ""))
	fieldNameMap["issuer_alias"] = "IssuerAlias"
	fields["domain"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["domain"] = "Domain"
	var validators = []bindings.Validator{}
	uv1 := bindings.NewUnionValidator("type",
	    map[string][]bindings.FieldData{
	        "LOCAL_USER": []bindings.FieldData{
	             bindings.NewFieldData("name", true),
	        },
	        "EXTERNAL_GROUP": []bindings.FieldData{
	             bindings.NewFieldData("name", true),
	             bindings.NewFieldData("issuer_alias", true),
	             bindings.NewFieldData("domain", true),
	        },
	        "EXTERNAL_USER": []bindings.FieldData{
	             bindings.NewFieldData("name", true),
	             bindings.NewFieldData("issuer_alias", true),
	             bindings.NewFieldData("domain", true),
	        },
	    },
	)
	validators = append(validators, uv1)
	return bindings.NewStructType("com.vmware.esx.authentication.client_profiles.subject", fields, reflect.TypeOf(ClientProfilesSubject{}), fieldNameMap, validators)
}

func ClientProfilesSummaryBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["summary_type"] = bindings.NewEnumType("com.vmware.esx.authentication.client_profiles.summary_type", reflect.TypeOf(ClientProfilesSummaryType(ClientProfilesSummaryType_FULL)))
	fieldNameMap["summary_type"] = "SummaryType"
	fields["client_profile"] = bindings.NewOptionalType(bindings.NewIdType([]string{"com.vmware.esx.authentication.clientprofile"}, ""))
	fieldNameMap["client_profile"] = "ClientProfile"
	fields["subject"] = bindings.NewOptionalType(bindings.NewReferenceType(ClientProfilesSubjectBindingType))
	fieldNameMap["subject"] = "Subject"
	fields["grants"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewReferenceType(ClientProfilesAccessGrantBindingType), reflect.TypeOf([]ClientProfilesAccessGrant{})))
	fieldNameMap["grants"] = "Grants"
	var validators = []bindings.Validator{}
	uv1 := bindings.NewUnionValidator("summary_type",
	    map[string][]bindings.FieldData{
	        "BRIEF": []bindings.FieldData{
	             bindings.NewFieldData("client_profile", true),
	        },
	        "NORMAL": []bindings.FieldData{
	             bindings.NewFieldData("client_profile", true),
	             bindings.NewFieldData("subject", true),
	        },
	        "FULL": []bindings.FieldData{
	             bindings.NewFieldData("client_profile", true),
	             bindings.NewFieldData("subject", true),
	             bindings.NewFieldData("grants", true),
	        },
	    },
	)
	validators = append(validators, uv1)
	return bindings.NewStructType("com.vmware.esx.authentication.client_profiles.summary", fields, reflect.TypeOf(ClientProfilesSummary{}), fieldNameMap, validators)
}

func ClientProfilesInfoBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["subject"] = bindings.NewReferenceType(ClientProfilesSubjectBindingType)
	fieldNameMap["subject"] = "Subject"
	fields["grants"] = bindings.NewListType(bindings.NewReferenceType(ClientProfilesAccessGrantBindingType), reflect.TypeOf([]ClientProfilesAccessGrant{}))
	fieldNameMap["grants"] = "Grants"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.esx.authentication.client_profiles.info", fields, reflect.TypeOf(ClientProfilesInfo{}), fieldNameMap, validators)
}

func ClientProfilesCreateSpecBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["local_user_name"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["local_user_name"] = "LocalUserName"
	fields["external_group_name"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["external_group_name"] = "ExternalGroupName"
	fields["external_user_name"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["external_user_name"] = "ExternalUserName"
	fields["issuer_alias"] = bindings.NewOptionalType(bindings.NewIdType([]string{"com.vmware.esx.authentication.trust.security-token-issuer"}, ""))
	fieldNameMap["issuer_alias"] = "IssuerAlias"
	fields["domain"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["domain"] = "Domain"
	fields["grants"] = bindings.NewListType(bindings.NewReferenceType(ClientProfilesAccessGrantBindingType), reflect.TypeOf([]ClientProfilesAccessGrant{}))
	fieldNameMap["grants"] = "Grants"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.esx.authentication.client_profiles.create_spec", fields, reflect.TypeOf(ClientProfilesCreateSpec{}), fieldNameMap, validators)
}

func ClientProfilesUpdateSpecBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["grants"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewReferenceType(ClientProfilesAccessGrantBindingType), reflect.TypeOf([]ClientProfilesAccessGrant{})))
	fieldNameMap["grants"] = "Grants"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.esx.authentication.client_profiles.update_spec", fields, reflect.TypeOf(ClientProfilesUpdateSpec{}), fieldNameMap, validators)
}

