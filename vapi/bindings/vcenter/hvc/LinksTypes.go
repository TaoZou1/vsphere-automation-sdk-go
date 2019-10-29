/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Code generated. DO NOT EDIT.

/*
 * Data type definitions file for service: Links.
 * Includes binding types of a structures and enumerations defined in the service.
 * Shared by client-side stubs and server-side skeletons to ensure type
 * compatibility.
 */

package hvc

import (
	"reflect"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/bindings/vapi/std"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/bindings"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/data"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/protocol"
)



// The ``Summary`` class contains information about the hybrid link. **Warning:** This class is available as Technology Preview. These are early access APIs provided to test, automate and provide feedback on the feature. Since this can change based on feedback, VMware does not guarantee backwards compatibility and recommends against using them in production environments. Some Technology Preview APIs might only be applicable to specific environments.
type LinksSummary struct {
    // Unique identifier for the link. **Warning:** This property is available as Technology Preview. These are early access APIs provided to test, automate and provide feedback on the feature. Since this can change based on feedback, VMware does not guarantee backwards compatibility and recommends against using them in production environments. Some Technology Preview APIs might only be applicable to specific environments.
	Link string
    // The display name is set to the domain name which was set during create. **Warning:** This property is available as Technology Preview. These are early access APIs provided to test, automate and provide feedback on the feature. Since this can change based on feedback, VMware does not guarantee backwards compatibility and recommends against using them in production environments. Some Technology Preview APIs might only be applicable to specific environments.
	DisplayName string
}

// The ``CreateSpec`` class is the specification used for the hybrid link creation. **Warning:** This class is available as Technology Preview. These are early access APIs provided to test, automate and provide feedback on the feature. Since this can change based on feedback, VMware does not guarantee backwards compatibility and recommends against using them in production environments. Some Technology Preview APIs might only be applicable to specific environments.
type LinksCreateSpec struct {
    // The PSC hostname for the domain to be linked. **Warning:** This property is available as Technology Preview. These are early access APIs provided to test, automate and provide feedback on the feature. Since this can change based on feedback, VMware does not guarantee backwards compatibility and recommends against using them in production environments. Some Technology Preview APIs might only be applicable to specific environments.
	PscHostname string
    // The HTTPS port of the PSC to be linked. **Warning:** This property is available as Technology Preview. These are early access APIs provided to test, automate and provide feedback on the feature. Since this can change based on feedback, VMware does not guarantee backwards compatibility and recommends against using them in production environments. Some Technology Preview APIs might only be applicable to specific environments.
	Port *string
    // The domain to which the PSC belongs. **Warning:** This property is available as Technology Preview. These are early access APIs provided to test, automate and provide feedback on the feature. Since this can change based on feedback, VMware does not guarantee backwards compatibility and recommends against using them in production environments. Some Technology Preview APIs might only be applicable to specific environments.
	DomainName string
    // The administrator username of the PSC. **Warning:** This property is available as Technology Preview. These are early access APIs provided to test, automate and provide feedback on the feature. Since this can change based on feedback, VMware does not guarantee backwards compatibility and recommends against using them in production environments. Some Technology Preview APIs might only be applicable to specific environments.
	Username string
    // The administrator password of the PSC. **Warning:** This property is available as Technology Preview. These are early access APIs provided to test, automate and provide feedback on the feature. Since this can change based on feedback, VMware does not guarantee backwards compatibility and recommends against using them in production environments. Some Technology Preview APIs might only be applicable to specific environments.
	Password string
    // The ssl thumbprint of the server. **Warning:** This property is available as Technology Preview. These are early access APIs provided to test, automate and provide feedback on the feature. Since this can change based on feedback, VMware does not guarantee backwards compatibility and recommends against using them in production environments. Some Technology Preview APIs might only be applicable to specific environments.
	SslThumbprint *string
    // List of groups to be added to enable administrator access to. **Warning:** This property is available as Technology Preview. These are early access APIs provided to test, automate and provide feedback on the feature. Since this can change based on feedback, VMware does not guarantee backwards compatibility and recommends against using them in production environments. Some Technology Preview APIs might only be applicable to specific environments.
	AdminGroups map[string]bool
}

// The ``CertificateInfo`` class contains information about the SSL certificate for a destination PSC endpoint. **Warning:** This class is available as Technology Preview. These are early access APIs provided to test, automate and provide feedback on the feature. Since this can change based on feedback, VMware does not guarantee backwards compatibility and recommends against using them in production environments. Some Technology Preview APIs might only be applicable to specific environments.
type LinksCertificateInfo struct {
    // The SHA-256 thumbprint of the SSL certificate for the destination PSC endpoint. **Warning:** This property is available as Technology Preview. These are early access APIs provided to test, automate and provide feedback on the feature. Since this can change based on feedback, VMware does not guarantee backwards compatibility and recommends against using them in production environments. Some Technology Preview APIs might only be applicable to specific environments.
	SslThumbprint string
}

// The ``Info`` class contains information about link. **Warning:** This class is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
type LinksInfo struct {
    // Health status of the connection. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
	ConnectionHealthStatus *LinksInfoHealthStatus
    // Localizable messages associated with health status. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
	HealthStatusMessage *std.LocalizableMessage
}

// The ``HealthStatus`` enumeration class defines the possible states for health of a link. **Warning:** This enumeration is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
//
// <p> See {@link com.vmware.vapi.bindings.ApiEnumeration enumerated types description}.
type LinksInfoHealthStatus string

const (
    // Connection is healthy. **Warning:** This constant field is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
	LinksInfoHealthStatus_HEALTHY LinksInfoHealthStatus = "HEALTHY"
    // Connection issues will need to be remediated. **Warning:** This constant field is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
	LinksInfoHealthStatus_UNHEALTHY LinksInfoHealthStatus = "UNHEALTHY"
)

func (h LinksInfoHealthStatus) LinksInfoHealthStatus() bool {
	switch h {
	case LinksInfoHealthStatus_HEALTHY:
		return true
	case LinksInfoHealthStatus_UNHEALTHY:
		return true
	default:
		return false
	}
}


// The ``Credentials`` interface specifies user credentials to make a successful connection to remote endpoint. **Warning:** This class is available as Technology Preview. These are early access APIs provided to test, automate and provide feedback on the feature. Since this can change based on feedback, VMware does not guarantee backwards compatibility and recommends against using them in production environments. Some Technology Preview APIs might only be applicable to specific environments.
type LinksCredentials struct {
    // Name of the user to authenticate. **Warning:** This property is available as Technology Preview. These are early access APIs provided to test, automate and provide feedback on the feature. Since this can change based on feedback, VMware does not guarantee backwards compatibility and recommends against using them in production environments. Some Technology Preview APIs might only be applicable to specific environments.
	UserName string
    // Password for the user. **Warning:** This property is available as Technology Preview. These are early access APIs provided to test, automate and provide feedback on the feature. Since this can change based on feedback, VMware does not guarantee backwards compatibility and recommends against using them in production environments. Some Technology Preview APIs might only be applicable to specific environments.
	Password string
}



func linksCreateInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["spec"] = bindings.NewReferenceType(LinksCreateSpecBindingType)
	fieldNameMap["spec"] = "Spec"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func linksCreateOutputType() bindings.BindingType {
	return bindings.NewIdType([]string{"com.vmware.vcenter.hvc.Links"}, "")
}

func linksCreateRestMetadata() protocol.OperationRestMetadata {
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
		map[string]int{"AlreadyExists": 400,"InvalidArgument": 400,"Unsupported": 400,"Unauthorized": 403,"UnverifiedPeer": 400,"Error": 500})
}

func linksDeleteInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["link"] = bindings.NewIdType([]string{"com.vmware.vcenter.hvc.Links"}, "")
	fieldNameMap["link"] = "Link"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func linksDeleteOutputType() bindings.BindingType {
	return bindings.NewVoidType()
}

func linksDeleteRestMetadata() protocol.OperationRestMetadata {
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
		map[string]int{"NotFound": 404,"Unauthorized": 403,"Error": 500})
}

func linksDeleteWithCredentialsInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["link"] = bindings.NewIdType([]string{"com.vmware.vcenter.hvc.Links"}, "")
	fields["credentials"] = bindings.NewOptionalType(bindings.NewReferenceType(LinksCredentialsBindingType))
	fieldNameMap["link"] = "Link"
	fieldNameMap["credentials"] = "Credentials"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func linksDeleteWithCredentialsOutputType() bindings.BindingType {
	return bindings.NewVoidType()
}

func linksDeleteWithCredentialsRestMetadata() protocol.OperationRestMetadata {
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
		map[string]int{"NotFound": 404,"Unauthorized": 403,"Error": 500})
}

func linksListInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func linksListOutputType() bindings.BindingType {
	return bindings.NewListType(bindings.NewReferenceType(LinksSummaryBindingType), reflect.TypeOf([]LinksSummary{}))
}

func linksListRestMetadata() protocol.OperationRestMetadata {
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
		map[string]int{"Error": 500})
}

func linksGetInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["link"] = bindings.NewIdType([]string{"com.vmware.vcenter.hvc.Links"}, "")
	fieldNameMap["link"] = "Link"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func linksGetOutputType() bindings.BindingType {
	return bindings.NewReferenceType(LinksInfoBindingType)
}

func linksGetRestMetadata() protocol.OperationRestMetadata {
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
		map[string]int{"Error": 500,"NotFound": 404,"Unauthorized": 403,"ResourceBusy": 500})
}


func LinksSummaryBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["link"] = bindings.NewIdType([]string{"com.vmware.vcenter.hvc.Links"}, "")
	fieldNameMap["link"] = "Link"
	fields["display_name"] = bindings.NewStringType()
	fieldNameMap["display_name"] = "DisplayName"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vcenter.hvc.links.summary", fields, reflect.TypeOf(LinksSummary{}), fieldNameMap, validators)
}

func LinksCreateSpecBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["psc_hostname"] = bindings.NewStringType()
	fieldNameMap["psc_hostname"] = "PscHostname"
	fields["port"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["port"] = "Port"
	fields["domain_name"] = bindings.NewStringType()
	fieldNameMap["domain_name"] = "DomainName"
	fields["username"] = bindings.NewStringType()
	fieldNameMap["username"] = "Username"
	fields["password"] = bindings.NewSecretType()
	fieldNameMap["password"] = "Password"
	fields["ssl_thumbprint"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["ssl_thumbprint"] = "SslThumbprint"
	fields["admin_groups"] = bindings.NewOptionalType(bindings.NewSetType(bindings.NewStringType(), reflect.TypeOf(map[string]bool{})))
	fieldNameMap["admin_groups"] = "AdminGroups"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vcenter.hvc.links.create_spec", fields, reflect.TypeOf(LinksCreateSpec{}), fieldNameMap, validators)
}

func LinksCertificateInfoBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["ssl_thumbprint"] = bindings.NewStringType()
	fieldNameMap["ssl_thumbprint"] = "SslThumbprint"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vcenter.hvc.links.certificate_info", fields, reflect.TypeOf(LinksCertificateInfo{}), fieldNameMap, validators)
}

func LinksInfoBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["connection_health_status"] = bindings.NewOptionalType(bindings.NewEnumType("com.vmware.vcenter.hvc.links.info.health_status", reflect.TypeOf(LinksInfoHealthStatus(LinksInfoHealthStatus_HEALTHY))))
	fieldNameMap["connection_health_status"] = "ConnectionHealthStatus"
	fields["health_status_message"] = bindings.NewOptionalType(bindings.NewReferenceType(std.LocalizableMessageBindingType))
	fieldNameMap["health_status_message"] = "HealthStatusMessage"
	var validators = []bindings.Validator{}
	uv1 := bindings.NewUnionValidator("connection_health_status",
		map[string][]bindings.FieldData{
			"UNHEALTHY": []bindings.FieldData{
				bindings.NewFieldData("health_status_message", true),
			},
			"HEALTHY": []bindings.FieldData{},
		},
	)
	validators = append(validators, uv1)
	return bindings.NewStructType("com.vmware.vcenter.hvc.links.info", fields, reflect.TypeOf(LinksInfo{}), fieldNameMap, validators)
}

func LinksCredentialsBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["user_name"] = bindings.NewStringType()
	fieldNameMap["user_name"] = "UserName"
	fields["password"] = bindings.NewSecretType()
	fieldNameMap["password"] = "Password"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vcenter.hvc.links.credentials", fields, reflect.TypeOf(LinksCredentials{}), fieldNameMap, validators)
}

