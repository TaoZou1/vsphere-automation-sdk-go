/* Copyright © 2019 VMware, Inc. All Rights Reserved.
     SPDX-License-Identifier: BSD-2-Clause */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Data type definitions file for service: Providers.
 * Includes binding types of a structures and enumerations defined in the service.
 * Shared by client-side stubs and server-side skeletons to ensure type
 * compatibility.
 */

package identity

import (
    "reflect"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/bindings/com/vmware/vcenter/certificate_management"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/bindings"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/data"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/protocol"
    "net/url"
)



// The ``ConfigType`` class contains the possible types of vCenter Server identity providers. **Warning:** This enumeration is available as Technology Preview. These are early access APIs provided to test, automate and provide feedback on the feature. Since this can change based on feedback, VMware does not guarantee backwards compatibility and recommends against using them in production environments. Some Technology Preview APIs might only be applicable to specific environments.
//
// <p> See {@link com.vmware.vapi.bindings.ApiEnumeration enumerated types description}.
 
type Providers_ConfigType string

const (
    // Config for OAuth2. **Warning:** This constant field is available as Technology Preview. These are early access APIs provided to test, automate and provide feedback on the feature. Since this can change based on feedback, VMware does not guarantee backwards compatibility and recommends against using them in production environments. Some Technology Preview APIs might only be applicable to specific environments.
     Providers_ConfigType_Oauth2 Providers_ConfigType = "Oauth2"
    // Config for OIDC. **Warning:** This constant field is available as Technology Preview. These are early access APIs provided to test, automate and provide feedback on the feature. Since this can change based on feedback, VMware does not guarantee backwards compatibility and recommends against using them in production environments. Some Technology Preview APIs might only be applicable to specific environments.
     Providers_ConfigType_Oidc Providers_ConfigType = "Oidc"
)

func (c Providers_ConfigType) Providers_ConfigType() bool {
    switch c {
        case Providers_ConfigType_Oauth2:
            return true
        case Providers_ConfigType_Oidc:
            return true
        default:
            return false
    }
}




// The ``IdmProtocol`` class contains the possible types of communication protocols to the identity management endpoints. **Warning:** This enumeration is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
//
// <p> See {@link com.vmware.vapi.bindings.ApiEnumeration enumerated types description}.
 
type Providers_IdmProtocol string

const (
    // REST protocol based identity management endpoints. **Warning:** This constant field is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
     Providers_IdmProtocol_REST Providers_IdmProtocol = "REST"
    // SCIM protocol based identity management endpoints. **Warning:** This constant field is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
     Providers_IdmProtocol_SCIM Providers_IdmProtocol = "SCIM"
    // LDAP protocol based identity management endpoints. **Warning:** This constant field is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
     Providers_IdmProtocol_LDAP Providers_IdmProtocol = "LDAP"
)

func (i Providers_IdmProtocol) Providers_IdmProtocol() bool {
    switch i {
        case Providers_IdmProtocol_REST:
            return true
        case Providers_IdmProtocol_SCIM:
            return true
        case Providers_IdmProtocol_LDAP:
            return true
        default:
            return false
    }
}




// The ``Oauth2AuthenticationMethod`` class contains the possible types of OAuth2 authentication methods. **Warning:** This enumeration is available as Technology Preview. These are early access APIs provided to test, automate and provide feedback on the feature. Since this can change based on feedback, VMware does not guarantee backwards compatibility and recommends against using them in production environments. Some Technology Preview APIs might only be applicable to specific environments.
//
// <p> See {@link com.vmware.vapi.bindings.ApiEnumeration enumerated types description}.
 
type Providers_Oauth2AuthenticationMethod string

const (
    // Clients that have received a client_secret value from the Authorization Server, authenticate with the Authorization Server in accordance with Section 3.2.1 of OAuth 2.0 [RFC6749] using the HTTP Basic authentication scheme. **Warning:** This constant field is available as Technology Preview. These are early access APIs provided to test, automate and provide feedback on the feature. Since this can change based on feedback, VMware does not guarantee backwards compatibility and recommends against using them in production environments. Some Technology Preview APIs might only be applicable to specific environments.
     Providers_Oauth2AuthenticationMethod_CLIENT_SECRET_BASIC Providers_Oauth2AuthenticationMethod = "CLIENT_SECRET_BASIC"
    // Clients that have received a client_secret value from the Authorization Server, authenticate with the Authorization Server in accordance with Section 3.2.1 of OAuth 2.0 [RFC6749] by including the Client Credentials in the request body. **Warning:** This constant field is available as Technology Preview. These are early access APIs provided to test, automate and provide feedback on the feature. Since this can change based on feedback, VMware does not guarantee backwards compatibility and recommends against using them in production environments. Some Technology Preview APIs might only be applicable to specific environments.
     Providers_Oauth2AuthenticationMethod_CLIENT_SECRET_POST Providers_Oauth2AuthenticationMethod = "CLIENT_SECRET_POST"
    // Clients that have received a client_secret value from the Authorization Server, create a JWT using an HMAC SHA algorithm, such as HMAC SHA-256. The HMAC (Hash-based Message Authentication Code) is calculated using the octets of the UTF-8 representation of the client_secret as the shared key. **Warning:** This constant field is available as Technology Preview. These are early access APIs provided to test, automate and provide feedback on the feature. Since this can change based on feedback, VMware does not guarantee backwards compatibility and recommends against using them in production environments. Some Technology Preview APIs might only be applicable to specific environments.
     Providers_Oauth2AuthenticationMethod_CLIENT_SECRET_JWT Providers_Oauth2AuthenticationMethod = "CLIENT_SECRET_JWT"
    // Clients that have registered a public key sign a JWT using that key. The client authenticates in accordance with JSON Web Token (JWT) Profile for OAuth 2.0 Client Authentication and Authorization Grants [OAuth.JWT] and Assertion Framework for OAuth 2.0 Client Authentication and Authorization Grants [OAuth.Assertions]. **Warning:** This constant field is available as Technology Preview. These are early access APIs provided to test, automate and provide feedback on the feature. Since this can change based on feedback, VMware does not guarantee backwards compatibility and recommends against using them in production environments. Some Technology Preview APIs might only be applicable to specific environments.
     Providers_Oauth2AuthenticationMethod_PRIVATE_KEY_JWT Providers_Oauth2AuthenticationMethod = "PRIVATE_KEY_JWT"
)

func (o Providers_Oauth2AuthenticationMethod) Providers_Oauth2AuthenticationMethod() bool {
    switch o {
        case Providers_Oauth2AuthenticationMethod_CLIENT_SECRET_BASIC:
            return true
        case Providers_Oauth2AuthenticationMethod_CLIENT_SECRET_POST:
            return true
        case Providers_Oauth2AuthenticationMethod_CLIENT_SECRET_JWT:
            return true
        case Providers_Oauth2AuthenticationMethod_PRIVATE_KEY_JWT:
            return true
        default:
            return false
    }
}





// The ``Summary`` class contains commonly used information about an identity provider. **Warning:** This class is available as Technology Preview. These are early access APIs provided to test, automate and provide feedback on the feature. Since this can change based on feedback, VMware does not guarantee backwards compatibility and recommends against using them in production environments. Some Technology Preview APIs might only be applicable to specific environments.
 type ProvidersSummary struct {
    // The identifier of the provider. **Warning:** This property is available as Technology Preview. These are early access APIs provided to test, automate and provide feedback on the feature. Since this can change based on feedback, VMware does not guarantee backwards compatibility and recommends against using them in production environments. Some Technology Preview APIs might only be applicable to specific environments.
    Provider string
    // The user friendly name for the provider. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    Name *string
    // The config type of the identity provider. **Warning:** This property is available as Technology Preview. These are early access APIs provided to test, automate and provide feedback on the feature. Since this can change based on feedback, VMware does not guarantee backwards compatibility and recommends against using them in production environments. Some Technology Preview APIs might only be applicable to specific environments.
    ConfigTag Providers_ConfigType
    // OAuth2 Summary. **Warning:** This property is available as Technology Preview. These are early access APIs provided to test, automate and provide feedback on the feature. Since this can change based on feedback, VMware does not guarantee backwards compatibility and recommends against using them in production environments. Some Technology Preview APIs might only be applicable to specific environments.
    Oauth2 *ProvidersOauth2Summary
    // OIDC Summary. **Warning:** This property is available as Technology Preview. These are early access APIs provided to test, automate and provide feedback on the feature. Since this can change based on feedback, VMware does not guarantee backwards compatibility and recommends against using them in production environments. Some Technology Preview APIs might only be applicable to specific environments.
    Oidc *ProvidersOidcSummary
    // Specifies whether the provider is the default provider. **Warning:** This property is available as Technology Preview. These are early access APIs provided to test, automate and provide feedback on the feature. Since this can change based on feedback, VMware does not guarantee backwards compatibility and recommends against using them in production environments. Some Technology Preview APIs might only be applicable to specific environments.
    IsDefault bool
    // Set of fully qualified domain names to trust when federating with this identity provider. Tokens from this identity provider will only be validated if the user belongs to one of these domains, and any domain-qualified groups in the tokens will be filtered to include only those groups that belong to one of these domains. If domainNames is an empty set, domain validation behavior at login with this identity provider will be as follows: the user's domain will be parsed from the User Principal Name (UPN) value that is found in the tokens returned by the identity provider. This domain will then be implicitly trusted and used to filter any groups that are also provided in the tokens. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    DomainNames map[string]bool
    //
//
// key/value pairs that are to be appended to the authEndpoint request. 
//
// How to append to authEndpoint request: If the map is not empty, a "?" is added to the endpoint URL, and combination of each k and each string in the v is added with an "&" delimiter. Details:
//
// * If the value contains only one string, then the key is added with "k=v".
// * If the value is an empty list, then the key is added without a "=v".
// * If the value contains multiple strings, then the key is repeated in the query-string for each string in the value.
//
// . **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    AuthQueryParams map[string][]string
}



//


// The ``Info`` class contains the information about an identity provider. **Warning:** This class is available as Technology Preview. These are early access APIs provided to test, automate and provide feedback on the feature. Since this can change based on feedback, VMware does not guarantee backwards compatibility and recommends against using them in production environments. Some Technology Preview APIs might only be applicable to specific environments.
 type ProvidersInfo struct {
    // The user friendly name for the provider. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    Name *string
    // The set of orgIds as part of SDDC creation which provides the basis for tenancy. **Warning:** This property is available as Technology Preview. These are early access APIs provided to test, automate and provide feedback on the feature. Since this can change based on feedback, VMware does not guarantee backwards compatibility and recommends against using them in production environments. Some Technology Preview APIs might only be applicable to specific environments.
    OrgIds map[string]bool
    // The config type of the identity provider. **Warning:** This property is available as Technology Preview. These are early access APIs provided to test, automate and provide feedback on the feature. Since this can change based on feedback, VMware does not guarantee backwards compatibility and recommends against using them in production environments. Some Technology Preview APIs might only be applicable to specific environments.
    ConfigTag Providers_ConfigType
    // OAuth2 Info. **Warning:** This property is available as Technology Preview. These are early access APIs provided to test, automate and provide feedback on the feature. Since this can change based on feedback, VMware does not guarantee backwards compatibility and recommends against using them in production environments. Some Technology Preview APIs might only be applicable to specific environments.
    Oauth2 *ProvidersOauth2Info
    // OIDC Info. **Warning:** This property is available as Technology Preview. These are early access APIs provided to test, automate and provide feedback on the feature. Since this can change based on feedback, VMware does not guarantee backwards compatibility and recommends against using them in production environments. Some Technology Preview APIs might only be applicable to specific environments.
    Oidc *ProvidersOidcInfo
    // Specifies whether the provider is the default provider. **Warning:** This property is available as Technology Preview. These are early access APIs provided to test, automate and provide feedback on the feature. Since this can change based on feedback, VMware does not guarantee backwards compatibility and recommends against using them in production environments. Some Technology Preview APIs might only be applicable to specific environments.
    IsDefault bool
    // Set of fully qualified domain names to trust when federating with this identity provider. Tokens from this identity provider will only be validated if the user belongs to one of these domains, and any domain-qualified groups in the tokens will be filtered to include only those groups that belong to one of these domains. If domainNames is an empty set, domain validation behavior at login with this identity provider will be as follows: the user's domain will be parsed from the User Principal Name (UPN) value that is found in the tokens returned by the identity provider. This domain will then be implicitly trusted and used to filter any groups that are also provided in the tokens. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    DomainNames map[string]bool
    //
//
// key/value pairs that are to be appended to the authEndpoint request. 
//
// How to append to authEndpoint request: If the map is not empty, a "?" is added to the endpoint URL, and combination of each k and each string in the v is added with an "&" delimiter. Details:
//
// * If the value contains only one string, then the key is added with "k=v".
// * If the value is an empty list, then the key is added without a "=v".
// * If the value contains multiple strings, then the key is repeated in the query-string for each string in the value.
//
// . **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    AuthQueryParams map[string][]string
    // Communication protocol to the identity management endpoints. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    IdmProtocol *Providers_IdmProtocol
    // Identity management endpoints. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    IdmEndpoints []url.URL
    // Identity management configuration. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    ActiveDirectoryOverLdap *ProvidersActiveDirectoryOverLdap
    // Specifies which claim provides the user principal name (UPN) for the user. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    UpnClaim *string
    // Specifies which claim provides the group membership for the token subject. If empty, the default behavior for CSP is used. In this case, the groups for the subject will be comprised of the groups in 'group_names' and 'group_ids' claims. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    GroupsClaim *string
}



//


// The ``CreateSpec`` class contains the information used to create an identity provider. **Warning:** This class is available as Technology Preview. These are early access APIs provided to test, automate and provide feedback on the feature. Since this can change based on feedback, VMware does not guarantee backwards compatibility and recommends against using them in production environments. Some Technology Preview APIs might only be applicable to specific environments.
 type ProvidersCreateSpec struct {
    // The config type of the identity provider. **Warning:** This property is available as Technology Preview. These are early access APIs provided to test, automate and provide feedback on the feature. Since this can change based on feedback, VMware does not guarantee backwards compatibility and recommends against using them in production environments. Some Technology Preview APIs might only be applicable to specific environments.
    ConfigTag Providers_ConfigType
    // OAuth2 CreateSpec. **Warning:** This property is available as Technology Preview. These are early access APIs provided to test, automate and provide feedback on the feature. Since this can change based on feedback, VMware does not guarantee backwards compatibility and recommends against using them in production environments. Some Technology Preview APIs might only be applicable to specific environments.
    Oauth2 *ProvidersOauth2CreateSpec
    // OIDC CreateSpec. **Warning:** This property is available as Technology Preview. These are early access APIs provided to test, automate and provide feedback on the feature. Since this can change based on feedback, VMware does not guarantee backwards compatibility and recommends against using them in production environments. Some Technology Preview APIs might only be applicable to specific environments.
    Oidc *ProvidersOidcCreateSpec
    // The set of orgIds as part of SDDC creation which provides the basis for tenancy. **Warning:** This property is available as Technology Preview. These are early access APIs provided to test, automate and provide feedback on the feature. Since this can change based on feedback, VMware does not guarantee backwards compatibility and recommends against using them in production environments. Some Technology Preview APIs might only be applicable to specific environments.
    OrgIds map[string]bool
    // Specifies whether the provider is the default provider. Setting ``isDefault`` of current provider to True makes all other providers non-default. If no other providers created in this vCenter Server before, this parameter will be disregarded, and the provider will always be set to the default. **Warning:** This property is available as Technology Preview. These are early access APIs provided to test, automate and provide feedback on the feature. Since this can change based on feedback, VMware does not guarantee backwards compatibility and recommends against using them in production environments. Some Technology Preview APIs might only be applicable to specific environments.
    IsDefault *bool
    // The user friendly name for the provider. This name can be used for human-readable identification purposes, but it does not have to be unique, as the system will use internal UUIDs to differentiate providers. **Warning:** This property is available as Technology Preview. These are early access APIs provided to test, automate and provide feedback on the feature. Since this can change based on feedback, VMware does not guarantee backwards compatibility and recommends against using them in production environments. Some Technology Preview APIs might only be applicable to specific environments.
    Name *string
    // Set of fully qualified domain names to trust when federating with this identity provider. Tokens from this identity provider will only be validated if the user belongs to one of these domains, and any domain-qualified groups in the tokens will be filtered to include only those groups that belong to one of these domains. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    DomainNames map[string]bool
    //
//
// key/value pairs that are to be appended to the authEndpoint request. 
//
// How to append to authEndpoint request: If the map is not empty, a "?" is added to the endpoint URL, and combination of each k and each string in the v is added with an "&" delimiter. Details:
//
// * If the value contains only one string, then the key is added with "k=v".
// * If the value is an empty list, then the key is added without a "=v".
// * If the value contains multiple strings, then the key is repeated in the query-string for each string in the value.
//
// . **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    AuthQueryParams map[string][]string
    // Communication protocol to the identity management endpoints. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    IdmProtocol *Providers_IdmProtocol
    // Identity management endpoints. When specified, at least one endpoint must be provided. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    IdmEndpoints []url.URL
    // Identity management configuration. If the protocol is LDAP, the configuration must be set, else InvalidArgument is thrown. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    ActiveDirectoryOverLdap *ProvidersActiveDirectoryOverLdap
    // Specifies which claim provides the user principal name (UPN) for the user. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    UpnClaim *string
    // Specifies which claim provides the group membership for the token subject. These groups will be used for mapping to local groups per the claim map. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    GroupsClaim *string
}



//


// The ``UpdateSpec`` class contains the information used to update the identity provider. **Warning:** This class is available as Technology Preview. These are early access APIs provided to test, automate and provide feedback on the feature. Since this can change based on feedback, VMware does not guarantee backwards compatibility and recommends against using them in production environments. Some Technology Preview APIs might only be applicable to specific environments.
 type ProvidersUpdateSpec struct {
    // The config type of the identity provider. **Warning:** This property is available as Technology Preview. These are early access APIs provided to test, automate and provide feedback on the feature. Since this can change based on feedback, VMware does not guarantee backwards compatibility and recommends against using them in production environments. Some Technology Preview APIs might only be applicable to specific environments.
    ConfigTag Providers_ConfigType
    // OAuth2 UpdateSpec. **Warning:** This property is available as Technology Preview. These are early access APIs provided to test, automate and provide feedback on the feature. Since this can change based on feedback, VMware does not guarantee backwards compatibility and recommends against using them in production environments. Some Technology Preview APIs might only be applicable to specific environments.
    Oauth2 *ProvidersOauth2UpdateSpec
    // OIDC UpdateSpec. **Warning:** This property is available as Technology Preview. These are early access APIs provided to test, automate and provide feedback on the feature. Since this can change based on feedback, VMware does not guarantee backwards compatibility and recommends against using them in production environments. Some Technology Preview APIs might only be applicable to specific environments.
    Oidc *ProvidersOidcUpdateSpec
    // The set orgIds as part of SDDC creation which provides the basis for tenancy. **Warning:** This property is available as Technology Preview. These are early access APIs provided to test, automate and provide feedback on the feature. Since this can change based on feedback, VMware does not guarantee backwards compatibility and recommends against using them in production environments. Some Technology Preview APIs might only be applicable to specific environments.
    OrgIds map[string]bool
    // Specifies whether to make this the default provider. If ``makeDefault`` is set to true, this provider will be flagged as the default provider and any other providers that had previously been flagged as the default will be made non-default. If ``makeDefault`` is set to false, this provider's default flag will not be modified. **Warning:** This property is available as Technology Preview. These are early access APIs provided to test, automate and provide feedback on the feature. Since this can change based on feedback, VMware does not guarantee backwards compatibility and recommends against using them in production environments. Some Technology Preview APIs might only be applicable to specific environments.
    MakeDefault *bool
    // The user friendly name for the provider. This name can be used for human-readable identification purposes, but it does not have to be unique, as the system will use internal UUIDs to differentiate providers. **Warning:** This property is available as Technology Preview. These are early access APIs provided to test, automate and provide feedback on the feature. Since this can change based on feedback, VMware does not guarantee backwards compatibility and recommends against using them in production environments. Some Technology Preview APIs might only be applicable to specific environments.
    Name *string
    // Set of fully qualified domain names to trust when federating with this identity provider. Tokens from this identity provider will only be validated if the user belongs to one of these domains, and any domain-qualified groups in the tokens will be filtered to include only those groups that belong to one of these domains. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    DomainNames map[string]bool
    // key/value pairs that are to be appended to the authEndpoint request. How to append to authEndpoint request: If the map is not empty, a "?" is added to the endpoint URL, and combination of each k and each string in the v is added with an "&" delimiter. Details: If the value contains only one string, then the key is added with "k=v". If the value is an empty list, then the key is added without a "=v". If the value contains multiple strings, then the key is repeated in the query-string for each string in the value. If the map is empty, deletes all params. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    AuthQueryParams map[string][]string
    // The protocol to communicate to the identity management endpoints. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    IdmProtocol *Providers_IdmProtocol
    // Identity management endpoints. When specified, at least one endpoint must be provided. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    IdmEndpoints []url.URL
    // Identity management configuration. If the protocol is LDAP, the configuration must be set, else InvalidArgument is thrown. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    ActiveDirectoryOverLdap *ProvidersActiveDirectoryOverLdap
    // Specifies which claim provides the user principal name (UPN) for the subject of the token. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    UpnClaim *string
    // Flag indicating whether the user principal name (UPN) claim should be set back to its default value. If this field is set to ``true``, the user principal name (UPN) claim will be set to 'acct', which is used for backwards compatibility with CSP. If this field is set to ``false``, the existing user principal name (UPN) claim will be changed to the value specified in ProvidersUpdateSpec#upnClaim, if any. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    ResetUpnClaim *bool
    // Specifies which claim provides the group membership for the token subject. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    GroupsClaim *string
    // Flag indicating whether any existing groups claim value should be removed. If this field is set to ``true``, the existing groups claim value is removed which defaults to backwards compatibility with CSP. In this case, the groups for the subject will be comprised of the groups in 'group_names' and 'group_ids' claims. If this field is set to ``false``, the existing groups claim will be changed to the value specified in ProvidersUpdateSpec#groupsClaim, if any. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    ResetGroupsClaim *bool
}



//


// The ``Oauth2Summary`` class contains commonly used information about an OAuth2 identity provider. **Warning:** This class is available as Technology Preview. These are early access APIs provided to test, automate and provide feedback on the feature. Since this can change based on feedback, VMware does not guarantee backwards compatibility and recommends against using them in production environments. Some Technology Preview APIs might only be applicable to specific environments.
 type ProvidersOauth2Summary struct {
    // Authentication/authorization endpoint of the provider. **Warning:** This property is available as Technology Preview. These are early access APIs provided to test, automate and provide feedback on the feature. Since this can change based on feedback, VMware does not guarantee backwards compatibility and recommends against using them in production environments. Some Technology Preview APIs might only be applicable to specific environments.
    AuthEndpoint url.URL
    // Token endpoint of the provider. **Warning:** This property is available as Technology Preview. These are early access APIs provided to test, automate and provide feedback on the feature. Since this can change based on feedback, VMware does not guarantee backwards compatibility and recommends against using them in production environments. Some Technology Preview APIs might only be applicable to specific environments.
    TokenEndpoint url.URL
    // Client identifier to connect to the provider. **Warning:** This property is available as Technology Preview. These are early access APIs provided to test, automate and provide feedback on the feature. Since this can change based on feedback, VMware does not guarantee backwards compatibility and recommends against using them in production environments. Some Technology Preview APIs might only be applicable to specific environments.
    ClientId string
    // The authentication data used as part of request header to acquire or refresh an OAuth2 token. The data format depends on the authentication method used. Example of basic authentication format: Authorization: Basic [base64Encode(clientId + ":" + secret)]. **Warning:** This property is available as Technology Preview. These are early access APIs provided to test, automate and provide feedback on the feature. Since this can change based on feedback, VMware does not guarantee backwards compatibility and recommends against using them in production environments. Some Technology Preview APIs might only be applicable to specific environments.
    AuthenticationHeader string
    //
//
// key/value pairs that are to be appended to the authEndpoint request. 
//
// How to append to authEndpoint request: If the map is not empty, a "?" is added to the endpoint URL, and combination of each k and each string in the v is added with an "&" delimiter. Details:
//
// * If the value contains only one string, then the key is added with "k=v".
// * If the value is an empty list, then the key is added without a "=v".
// * If the value contains multiple strings, then the key is repeated in the query-string for each string in the value.
//
// . **Warning:** This property is available as Technology Preview. These are early access APIs provided to test, automate and provide feedback on the feature. Since this can change based on feedback, VMware does not guarantee backwards compatibility and recommends against using them in production environments. Some Technology Preview APIs might only be applicable to specific environments.
    AuthQueryParams map[string][]string
}



//


// The ``Oauth2Info`` class contains the information about an OAuth2 identity provider. **Warning:** This class is available as Technology Preview. These are early access APIs provided to test, automate and provide feedback on the feature. Since this can change based on feedback, VMware does not guarantee backwards compatibility and recommends against using them in production environments. Some Technology Preview APIs might only be applicable to specific environments.
 type ProvidersOauth2Info struct {
    // Authentication/authorization endpoint of the provider. **Warning:** This property is available as Technology Preview. These are early access APIs provided to test, automate and provide feedback on the feature. Since this can change based on feedback, VMware does not guarantee backwards compatibility and recommends against using them in production environments. Some Technology Preview APIs might only be applicable to specific environments.
    AuthEndpoint url.URL
    // Token endpoint of the provider. **Warning:** This property is available as Technology Preview. These are early access APIs provided to test, automate and provide feedback on the feature. Since this can change based on feedback, VMware does not guarantee backwards compatibility and recommends against using them in production environments. Some Technology Preview APIs might only be applicable to specific environments.
    TokenEndpoint url.URL
    // Endpoint to retrieve the provider public key for validation. **Warning:** This property is available as Technology Preview. These are early access APIs provided to test, automate and provide feedback on the feature. Since this can change based on feedback, VMware does not guarantee backwards compatibility and recommends against using them in production environments. Some Technology Preview APIs might only be applicable to specific environments.
    PublicKeyUri url.URL
    // Client identifier to connect to the provider. **Warning:** This property is available as Technology Preview. These are early access APIs provided to test, automate and provide feedback on the feature. Since this can change based on feedback, VMware does not guarantee backwards compatibility and recommends against using them in production environments. Some Technology Preview APIs might only be applicable to specific environments.
    ClientId string
    // The secret shared between the client and the provider. **Warning:** This property is available as Technology Preview. These are early access APIs provided to test, automate and provide feedback on the feature. Since this can change based on feedback, VMware does not guarantee backwards compatibility and recommends against using them in production environments. Some Technology Preview APIs might only be applicable to specific environments.
    ClientSecret string
    // The map used to transform an OAuth2 claim to a corresponding claim that vCenter Server understands. Currently only the key "perms" is supported. The key "perms" is used for mapping the "perms" claim of incoming JWT. The value is another map with an external group as the key and a vCenter Server group as value. **Warning:** This property is available as Technology Preview. These are early access APIs provided to test, automate and provide feedback on the feature. Since this can change based on feedback, VMware does not guarantee backwards compatibility and recommends against using them in production environments. Some Technology Preview APIs might only be applicable to specific environments.
    ClaimMap map[string]map[string][]string
    // The identity provider namespace. It is used to validate the issuer in the acquired OAuth2 token. **Warning:** This property is available as Technology Preview. These are early access APIs provided to test, automate and provide feedback on the feature. Since this can change based on feedback, VMware does not guarantee backwards compatibility and recommends against using them in production environments. Some Technology Preview APIs might only be applicable to specific environments.
    Issuer string
    // Authentication method used by the provider. **Warning:** This property is available as Technology Preview. These are early access APIs provided to test, automate and provide feedback on the feature. Since this can change based on feedback, VMware does not guarantee backwards compatibility and recommends against using them in production environments. Some Technology Preview APIs might only be applicable to specific environments.
    AuthenticationMethod Providers_Oauth2AuthenticationMethod
    //
//
// key/value pairs that are to be appended to the authEndpoint request. 
//
// How to append to authEndpoint request: If the map is not empty, a "?" is added to the endpoint URL, and combination of each k and each string in the v is added with an "&" delimiter. Details:
//
// * If the value contains only one string, then the key is added with "k=v".
// * If the value is an empty list, then the key is added without a "=v".
// * If the value contains multiple strings, then the key is repeated in the query-string for each string in the value.
//
// . **Warning:** This property is available as Technology Preview. These are early access APIs provided to test, automate and provide feedback on the feature. Since this can change based on feedback, VMware does not guarantee backwards compatibility and recommends against using them in production environments. Some Technology Preview APIs might only be applicable to specific environments.
    AuthQueryParams map[string][]string
}



//


// The ``Oauth2CreateSpec`` class contains the information used to create an OAuth2 identity provider. **Warning:** This class is available as Technology Preview. These are early access APIs provided to test, automate and provide feedback on the feature. Since this can change based on feedback, VMware does not guarantee backwards compatibility and recommends against using them in production environments. Some Technology Preview APIs might only be applicable to specific environments.
 type ProvidersOauth2CreateSpec struct {
    // Authentication/authorization endpoint of the provider. **Warning:** This property is available as Technology Preview. These are early access APIs provided to test, automate and provide feedback on the feature. Since this can change based on feedback, VMware does not guarantee backwards compatibility and recommends against using them in production environments. Some Technology Preview APIs might only be applicable to specific environments.
    AuthEndpoint url.URL
    // Token endpoint of the provider. **Warning:** This property is available as Technology Preview. These are early access APIs provided to test, automate and provide feedback on the feature. Since this can change based on feedback, VMware does not guarantee backwards compatibility and recommends against using them in production environments. Some Technology Preview APIs might only be applicable to specific environments.
    TokenEndpoint url.URL
    // Endpoint to retrieve the provider public key for validation. **Warning:** This property is available as Technology Preview. These are early access APIs provided to test, automate and provide feedback on the feature. Since this can change based on feedback, VMware does not guarantee backwards compatibility and recommends against using them in production environments. Some Technology Preview APIs might only be applicable to specific environments.
    PublicKeyUri url.URL
    // Client identifier to connect to the provider. **Warning:** This property is available as Technology Preview. These are early access APIs provided to test, automate and provide feedback on the feature. Since this can change based on feedback, VMware does not guarantee backwards compatibility and recommends against using them in production environments. Some Technology Preview APIs might only be applicable to specific environments.
    ClientId string
    // The secret shared between the client and the provider. **Warning:** This property is available as Technology Preview. These are early access APIs provided to test, automate and provide feedback on the feature. Since this can change based on feedback, VMware does not guarantee backwards compatibility and recommends against using them in production environments. Some Technology Preview APIs might only be applicable to specific environments.
    ClientSecret string
    // The map used to transform an OAuth2 claim to a corresponding claim that vCenter Server understands. Currently only the key "perms" is supported. The key "perms" is used for mapping the "perms" claim of incoming JWT. The value is another map with an external group as the key and a vCenter Server group as value. **Warning:** This property is available as Technology Preview. These are early access APIs provided to test, automate and provide feedback on the feature. Since this can change based on feedback, VMware does not guarantee backwards compatibility and recommends against using them in production environments. Some Technology Preview APIs might only be applicable to specific environments.
    ClaimMap map[string]map[string][]string
    // The identity provider namespace. It is used to validate the issuer in the acquired OAuth2 token. **Warning:** This property is available as Technology Preview. These are early access APIs provided to test, automate and provide feedback on the feature. Since this can change based on feedback, VMware does not guarantee backwards compatibility and recommends against using them in production environments. Some Technology Preview APIs might only be applicable to specific environments.
    Issuer string
    // Authentication method used by the provider. **Warning:** This property is available as Technology Preview. These are early access APIs provided to test, automate and provide feedback on the feature. Since this can change based on feedback, VMware does not guarantee backwards compatibility and recommends against using them in production environments. Some Technology Preview APIs might only be applicable to specific environments.
    AuthenticationMethod Providers_Oauth2AuthenticationMethod
    //
//
// key/value pairs that are to be appended to the authEndpoint request. 
//
// How to append to authEndpoint request: If the map is not empty, a "?" is added to the endpoint URL, and combination of each k and each string in the v is added with an "&" delimiter. Details:
//
// * If the value contains only one string, then the key is added with "k=v".
// * If the value is an empty list, then the key is added without a "=v".
// * If the value contains multiple strings, then the key is repeated in the query-string for each string in the value.
//
// . **Warning:** This property is available as Technology Preview. These are early access APIs provided to test, automate and provide feedback on the feature. Since this can change based on feedback, VMware does not guarantee backwards compatibility and recommends against using them in production environments. Some Technology Preview APIs might only be applicable to specific environments.
    AuthQueryParams map[string][]string
}



//


// The ``Oauth2UpdateSpec`` class contains the information used to update the OAuth2 identity provider. **Warning:** This class is available as Technology Preview. These are early access APIs provided to test, automate and provide feedback on the feature. Since this can change based on feedback, VMware does not guarantee backwards compatibility and recommends against using them in production environments. Some Technology Preview APIs might only be applicable to specific environments.
 type ProvidersOauth2UpdateSpec struct {
    // Authentication/authorization endpoint of the provider. **Warning:** This property is available as Technology Preview. These are early access APIs provided to test, automate and provide feedback on the feature. Since this can change based on feedback, VMware does not guarantee backwards compatibility and recommends against using them in production environments. Some Technology Preview APIs might only be applicable to specific environments.
    AuthEndpoint *url.URL
    // Token endpoint of the provider. **Warning:** This property is available as Technology Preview. These are early access APIs provided to test, automate and provide feedback on the feature. Since this can change based on feedback, VMware does not guarantee backwards compatibility and recommends against using them in production environments. Some Technology Preview APIs might only be applicable to specific environments.
    TokenEndpoint *url.URL
    // Endpoint to retrieve the provider public key for validation. **Warning:** This property is available as Technology Preview. These are early access APIs provided to test, automate and provide feedback on the feature. Since this can change based on feedback, VMware does not guarantee backwards compatibility and recommends against using them in production environments. Some Technology Preview APIs might only be applicable to specific environments.
    PublicKeyUri *url.URL
    // Client identifier to connect to the provider. **Warning:** This property is available as Technology Preview. These are early access APIs provided to test, automate and provide feedback on the feature. Since this can change based on feedback, VMware does not guarantee backwards compatibility and recommends against using them in production environments. Some Technology Preview APIs might only be applicable to specific environments.
    ClientId *string
    // Shared secret between identity provider and client. **Warning:** This property is available as Technology Preview. These are early access APIs provided to test, automate and provide feedback on the feature. Since this can change based on feedback, VMware does not guarantee backwards compatibility and recommends against using them in production environments. Some Technology Preview APIs might only be applicable to specific environments.
    ClientSecret *string
    // The map used to transform an OAuth2 claim to a corresponding claim that vCenter Server understands. Currently only the key "perms" is supported. The key "perms" is used for mapping the "perms" claim of incoming JWT. The value is another map with an external group as the key and a vCenter Server group as value. **Warning:** This property is available as Technology Preview. These are early access APIs provided to test, automate and provide feedback on the feature. Since this can change based on feedback, VMware does not guarantee backwards compatibility and recommends against using them in production environments. Some Technology Preview APIs might only be applicable to specific environments.
    ClaimMap map[string]map[string][]string
    // The identity provider namespace. It is used to validate the issuer in the acquired OAuth2 token. **Warning:** This property is available as Technology Preview. These are early access APIs provided to test, automate and provide feedback on the feature. Since this can change based on feedback, VMware does not guarantee backwards compatibility and recommends against using them in production environments. Some Technology Preview APIs might only be applicable to specific environments.
    Issuer *string
    // Authentication method used by the provider. **Warning:** This property is available as Technology Preview. These are early access APIs provided to test, automate and provide feedback on the feature. Since this can change based on feedback, VMware does not guarantee backwards compatibility and recommends against using them in production environments. Some Technology Preview APIs might only be applicable to specific environments.
    AuthenticationMethod *Providers_Oauth2AuthenticationMethod
    // key/value pairs that are to be appended to the authEndpoint request. How to append to authEndpoint request: If the map is not empty, a "?" is added to the endpoint URL, and combination of each k and each string in the v is added with an "&" delimiter. Details: If the value contains only one string, then the key is added with "k=v". If the value is an empty list, then the key is added without a "=v". If the value contains multiple strings, then the key is repeated in the query-string for each string in the value. If the map is empty, deletes all params. **Warning:** This property is available as Technology Preview. These are early access APIs provided to test, automate and provide feedback on the feature. Since this can change based on feedback, VMware does not guarantee backwards compatibility and recommends against using them in production environments. Some Technology Preview APIs might only be applicable to specific environments.
    AuthQueryParams map[string][]string
}



//


// The ``OidcSummary`` class contains commonly used information about an OIDC identity provider. OIDC is a discovery protocol for OAuth2 configuration metadata, so ``OidcSummary`` contains discovered OAuth2 metadata. **Warning:** This class is available as Technology Preview. These are early access APIs provided to test, automate and provide feedback on the feature. Since this can change based on feedback, VMware does not guarantee backwards compatibility and recommends against using them in production environments. Some Technology Preview APIs might only be applicable to specific environments.
 type ProvidersOidcSummary struct {
    // Endpoint to retrieve the provider metadata. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    DiscoveryEndpoint *url.URL
    // The endpoint to use for terminating the user's session at the identity provider. This value is automatically derived from the metadata information provided by the OIDC discovery endpoint. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    LogoutEndpoint *url.URL
    // Authentication/authorization endpoint of the provider. **Warning:** This property is available as Technology Preview. These are early access APIs provided to test, automate and provide feedback on the feature. Since this can change based on feedback, VMware does not guarantee backwards compatibility and recommends against using them in production environments. Some Technology Preview APIs might only be applicable to specific environments.
    AuthEndpoint url.URL
    // Token endpoint of the provider. **Warning:** This property is available as Technology Preview. These are early access APIs provided to test, automate and provide feedback on the feature. Since this can change based on feedback, VMware does not guarantee backwards compatibility and recommends against using them in production environments. Some Technology Preview APIs might only be applicable to specific environments.
    TokenEndpoint url.URL
    // Client identifier to connect to the provider. **Warning:** This property is available as Technology Preview. These are early access APIs provided to test, automate and provide feedback on the feature. Since this can change based on feedback, VMware does not guarantee backwards compatibility and recommends against using them in production environments. Some Technology Preview APIs might only be applicable to specific environments.
    ClientId string
    // The authentication data used as part of request header to acquire or refresh an OAuth2 token. The data format depends on the authentication method used. Example of basic authentication format: Authorization: Basic [base64Encode(clientId + ":" + secret)]. **Warning:** This property is available as Technology Preview. These are early access APIs provided to test, automate and provide feedback on the feature. Since this can change based on feedback, VMware does not guarantee backwards compatibility and recommends against using them in production environments. Some Technology Preview APIs might only be applicable to specific environments.
    AuthenticationHeader string
    //
//
// key/value pairs that are to be appended to the authEndpoint request. 
//
// How to append to authEndpoint request: If the map is not empty, a "?" is added to the endpoint URL, and combination of each k and each string in the v is added with an "&" delimiter. Details:
//
// * If the value contains only one string, then the key is added with "k=v".
// * If the value is an empty list, then the key is added without a "=v".
// * If the value contains multiple strings, then the key is repeated in the query-string for each string in the value.
//
// . **Warning:** This property is available as Technology Preview. These are early access APIs provided to test, automate and provide feedback on the feature. Since this can change based on feedback, VMware does not guarantee backwards compatibility and recommends against using them in production environments. Some Technology Preview APIs might only be applicable to specific environments.
    AuthQueryParams map[string][]string
}



//


// The ``OidcInfo`` class contains information about an OIDC identity provider. OIDC is a discovery protocol for OAuth2 configuration metadata, so ``OidcInfo`` contains additional discovered OAuth2 metadata. **Warning:** This class is available as Technology Preview. These are early access APIs provided to test, automate and provide feedback on the feature. Since this can change based on feedback, VMware does not guarantee backwards compatibility and recommends against using them in production environments. Some Technology Preview APIs might only be applicable to specific environments.
 type ProvidersOidcInfo struct {
    // Endpoint to retrieve the provider metadata. **Warning:** This property is available as Technology Preview. These are early access APIs provided to test, automate and provide feedback on the feature. Since this can change based on feedback, VMware does not guarantee backwards compatibility and recommends against using them in production environments. Some Technology Preview APIs might only be applicable to specific environments.
    DiscoveryEndpoint url.URL
    // The endpoint to use for terminating the user's session at the identity provider. This value is automatically derived from the metadata information provided by the OIDC discovery endpoint. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    LogoutEndpoint *url.URL
    // Authentication/authorization endpoint of the provider. **Warning:** This property is available as Technology Preview. These are early access APIs provided to test, automate and provide feedback on the feature. Since this can change based on feedback, VMware does not guarantee backwards compatibility and recommends against using them in production environments. Some Technology Preview APIs might only be applicable to specific environments.
    AuthEndpoint url.URL
    // Token endpoint of the provider. **Warning:** This property is available as Technology Preview. These are early access APIs provided to test, automate and provide feedback on the feature. Since this can change based on feedback, VMware does not guarantee backwards compatibility and recommends against using them in production environments. Some Technology Preview APIs might only be applicable to specific environments.
    TokenEndpoint url.URL
    // Endpoint to retrieve the provider public key for validation. **Warning:** This property is available as Technology Preview. These are early access APIs provided to test, automate and provide feedback on the feature. Since this can change based on feedback, VMware does not guarantee backwards compatibility and recommends against using them in production environments. Some Technology Preview APIs might only be applicable to specific environments.
    PublicKeyUri url.URL
    // Client identifier to connect to the provider. **Warning:** This property is available as Technology Preview. These are early access APIs provided to test, automate and provide feedback on the feature. Since this can change based on feedback, VMware does not guarantee backwards compatibility and recommends against using them in production environments. Some Technology Preview APIs might only be applicable to specific environments.
    ClientId string
    // The secret shared between the client and the provider. **Warning:** This property is available as Technology Preview. These are early access APIs provided to test, automate and provide feedback on the feature. Since this can change based on feedback, VMware does not guarantee backwards compatibility and recommends against using them in production environments. Some Technology Preview APIs might only be applicable to specific environments.
    ClientSecret string
    // The map used to transform an OAuth2 claim to a corresponding claim that vCenter Server understands. Currently only the key "perms" is supported. The key "perms" is used for mapping the "perms" claim of incoming JWT. The value is another map with an external group as the key and a vCenter Server group as value. **Warning:** This property is available as Technology Preview. These are early access APIs provided to test, automate and provide feedback on the feature. Since this can change based on feedback, VMware does not guarantee backwards compatibility and recommends against using them in production environments. Some Technology Preview APIs might only be applicable to specific environments.
    ClaimMap map[string]map[string][]string
    // The identity provider namespace. It is used to validate the issuer in the acquired OAuth2 token. **Warning:** This property is available as Technology Preview. These are early access APIs provided to test, automate and provide feedback on the feature. Since this can change based on feedback, VMware does not guarantee backwards compatibility and recommends against using them in production environments. Some Technology Preview APIs might only be applicable to specific environments.
    Issuer string
    // Authentication method used by the provider. **Warning:** This property is available as Technology Preview. These are early access APIs provided to test, automate and provide feedback on the feature. Since this can change based on feedback, VMware does not guarantee backwards compatibility and recommends against using them in production environments. Some Technology Preview APIs might only be applicable to specific environments.
    AuthenticationMethod Providers_Oauth2AuthenticationMethod
    //
//
// key/value pairs that are to be appended to the authEndpoint request. 
//
// How to append to authEndpoint request: If the map is not empty, a "?" is added to the endpoint URL, and combination of each k and each string in the v is added with an "&" delimiter. Details:
//
// * If the value contains only one string, then the key is added with "k=v".
// * If the value is an empty list, then the key is added without a "=v".
// * If the value contains multiple strings, then the key is repeated in the query-string for each string in the value.
//
// . **Warning:** This property is available as Technology Preview. These are early access APIs provided to test, automate and provide feedback on the feature. Since this can change based on feedback, VMware does not guarantee backwards compatibility and recommends against using them in production environments. Some Technology Preview APIs might only be applicable to specific environments.
    AuthQueryParams map[string][]string
}



//


// The ``OidcCreateSpec`` class contains the information used to create an OIDC identity provider. **Warning:** This class is available as Technology Preview. These are early access APIs provided to test, automate and provide feedback on the feature. Since this can change based on feedback, VMware does not guarantee backwards compatibility and recommends against using them in production environments. Some Technology Preview APIs might only be applicable to specific environments.
 type ProvidersOidcCreateSpec struct {
    // Endpoint to retrieve the provider metadata. **Warning:** This property is available as Technology Preview. These are early access APIs provided to test, automate and provide feedback on the feature. Since this can change based on feedback, VMware does not guarantee backwards compatibility and recommends against using them in production environments. Some Technology Preview APIs might only be applicable to specific environments.
    DiscoveryEndpoint url.URL
    // Client identifier to connect to the provider. **Warning:** This property is available as Technology Preview. These are early access APIs provided to test, automate and provide feedback on the feature. Since this can change based on feedback, VMware does not guarantee backwards compatibility and recommends against using them in production environments. Some Technology Preview APIs might only be applicable to specific environments.
    ClientId string
    // The secret shared between the client and the provider. **Warning:** This property is available as Technology Preview. These are early access APIs provided to test, automate and provide feedback on the feature. Since this can change based on feedback, VMware does not guarantee backwards compatibility and recommends against using them in production environments. Some Technology Preview APIs might only be applicable to specific environments.
    ClientSecret string
    // The map used to transform an OAuth2 claim to a corresponding claim that vCenter Server understands. Currently only the key "perms" is supported. The key "perms" is used for mapping the "perms" claim of incoming JWT. The value is another map with an external group as the key and a vCenter Server group as value. **Warning:** This property is available as Technology Preview. These are early access APIs provided to test, automate and provide feedback on the feature. Since this can change based on feedback, VMware does not guarantee backwards compatibility and recommends against using them in production environments. Some Technology Preview APIs might only be applicable to specific environments.
    ClaimMap map[string]map[string][]string
}



//


// The ``OidcUpdateSpec`` class contains the information used to update the OIDC identity provider. **Warning:** This class is available as Technology Preview. These are early access APIs provided to test, automate and provide feedback on the feature. Since this can change based on feedback, VMware does not guarantee backwards compatibility and recommends against using them in production environments. Some Technology Preview APIs might only be applicable to specific environments.
 type ProvidersOidcUpdateSpec struct {
    // Endpoint to retrieve the provider metadata. **Warning:** This property is available as Technology Preview. These are early access APIs provided to test, automate and provide feedback on the feature. Since this can change based on feedback, VMware does not guarantee backwards compatibility and recommends against using them in production environments. Some Technology Preview APIs might only be applicable to specific environments.
    DiscoveryEndpoint *url.URL
    // Client identifier to connect to the provider. **Warning:** This property is available as Technology Preview. These are early access APIs provided to test, automate and provide feedback on the feature. Since this can change based on feedback, VMware does not guarantee backwards compatibility and recommends against using them in production environments. Some Technology Preview APIs might only be applicable to specific environments.
    ClientId *string
    // The secret shared between the client and the provider. **Warning:** This property is available as Technology Preview. These are early access APIs provided to test, automate and provide feedback on the feature. Since this can change based on feedback, VMware does not guarantee backwards compatibility and recommends against using them in production environments. Some Technology Preview APIs might only be applicable to specific environments.
    ClientSecret *string
    // The map used to transform an OAuth2 claim to a corresponding claim that vCenter Server understands. Currently only the key "perms" is supported. The key "perms" is used for mapping the "perms" claim of incoming JWT. The value is another map with an external group as the key and a vCenter Server group as value. **Warning:** This property is available as Technology Preview. These are early access APIs provided to test, automate and provide feedback on the feature. Since this can change based on feedback, VMware does not guarantee backwards compatibility and recommends against using them in production environments. Some Technology Preview APIs might only be applicable to specific environments.
    ClaimMap map[string]map[string][]string
}



//


// The ``ActiveDirectoryOverLdap`` class contains the information about to how to use an Active Directory over LDAP connection to allow searching for users and groups if the identity provider is an On-Prem service. **Warning:** This class is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
 type ProvidersActiveDirectoryOverLdap struct {
    // User name to connect to the active directory server. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    UserName string
    // Password to connect to the active directory server. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    Password string
    // Base distinguished name for users. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    UsersBaseDn string
    // Base distinguished name for groups. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    GroupsBaseDn string
    // Active directory server endpoints. At least one active directory server endpoint must be set. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    ServerEndpoints []url.URL
    // SSL certificate chain in base64 encoding. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    CertChain *certificate_management.X509CertChain
}



//





func providersListInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func providersListOutputType() bindings.BindingType {
    return bindings.NewListType(bindings.NewReferenceType(ProvidersSummaryBindingType), reflect.TypeOf([]ProvidersSummary{}))
}

func providersListRestMetadata() protocol.OperationRestMetadata {
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


func providersGetInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["provider"] = bindings.NewIdType([]string {"com.vmware.vcenter.identity.Providers"}, "")
    fieldNameMap["provider"] = "Provider"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func providersGetOutputType() bindings.BindingType {
    return bindings.NewReferenceType(ProvidersInfoBindingType)
}

func providersGetRestMetadata() protocol.OperationRestMetadata {
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
       map[string]int{"Unauthorized": 403,"NotFound": 404})
}


func providersCreateInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["spec"] = bindings.NewReferenceType(ProvidersCreateSpecBindingType)
    fieldNameMap["spec"] = "Spec"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func providersCreateOutputType() bindings.BindingType {
    return bindings.NewIdType([]string {"com.vmware.vcenter.identity.Providers"}, "")
}

func providersCreateRestMetadata() protocol.OperationRestMetadata {
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


func providersUpdateInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["provider"] = bindings.NewIdType([]string {"com.vmware.vcenter.identity.Providers"}, "")
    fields["spec"] = bindings.NewReferenceType(ProvidersUpdateSpecBindingType)
    fieldNameMap["provider"] = "Provider"
    fieldNameMap["spec"] = "Spec"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func providersUpdateOutputType() bindings.BindingType {
    return bindings.NewVoidType()
}

func providersUpdateRestMetadata() protocol.OperationRestMetadata {
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
       map[string]int{"Unauthorized": 403,"InvalidArgument": 400,"NotFound": 404})
}


func providersDeleteInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["provider"] = bindings.NewIdType([]string {"com.vmware.vcenter.identity.Providers"}, "")
    fieldNameMap["provider"] = "Provider"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func providersDeleteOutputType() bindings.BindingType {
    return bindings.NewVoidType()
}

func providersDeleteRestMetadata() protocol.OperationRestMetadata {
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
       map[string]int{"Unauthorized": 403,"NotFound": 404})
}



func ProvidersSummaryBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["provider"] = bindings.NewIdType([]string {"com.vmware.vcenter.identity.Providers"}, "")
    fieldNameMap["provider"] = "Provider"
    fields["name"] = bindings.NewOptionalType(bindings.NewStringType())
    fieldNameMap["name"] = "Name"
    fields["config_tag"] = bindings.NewEnumType("com.vmware.vcenter.identity.providers.config_type", reflect.TypeOf(Providers_ConfigType(Providers_ConfigType_Oauth2)))
    fieldNameMap["config_tag"] = "ConfigTag"
    fields["oauth2"] = bindings.NewOptionalType(bindings.NewReferenceType(ProvidersOauth2SummaryBindingType))
    fieldNameMap["oauth2"] = "Oauth2"
    fields["oidc"] = bindings.NewOptionalType(bindings.NewReferenceType(ProvidersOidcSummaryBindingType))
    fieldNameMap["oidc"] = "Oidc"
    fields["is_default"] = bindings.NewBooleanType()
    fieldNameMap["is_default"] = "IsDefault"
    fields["domain_names"] = bindings.NewOptionalType(bindings.NewSetType(bindings.NewStringType(), reflect.TypeOf(map[string]bool{})))
    fieldNameMap["domain_names"] = "DomainNames"
    fields["auth_query_params"] = bindings.NewOptionalType(bindings.NewMapType(bindings.NewStringType(), bindings.NewListType(bindings.NewStringType(), reflect.TypeOf([]string{})),reflect.TypeOf(map[string][]string{})))
    fieldNameMap["auth_query_params"] = "AuthQueryParams"
    var validators = []bindings.Validator{}
    uv1 := bindings.NewUnionValidator("config_tag",
        map[string][]bindings.FieldData {
            "Oauth2": []bindings.FieldData {
                 bindings.NewFieldData("oauth2", true),
            },
            "Oidc": []bindings.FieldData {
                 bindings.NewFieldData("oidc", true),
            },
        },
    )
    validators = append(validators, uv1)
    return bindings.NewStructType("com.vmware.vcenter.identity.providers.summary",fields, reflect.TypeOf(ProvidersSummary{}), fieldNameMap, validators)
}

func ProvidersInfoBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["name"] = bindings.NewOptionalType(bindings.NewStringType())
    fieldNameMap["name"] = "Name"
    fields["org_ids"] = bindings.NewSetType(bindings.NewStringType(), reflect.TypeOf(map[string]bool{}))
    fieldNameMap["org_ids"] = "OrgIds"
    fields["config_tag"] = bindings.NewEnumType("com.vmware.vcenter.identity.providers.config_type", reflect.TypeOf(Providers_ConfigType(Providers_ConfigType_Oauth2)))
    fieldNameMap["config_tag"] = "ConfigTag"
    fields["oauth2"] = bindings.NewOptionalType(bindings.NewReferenceType(ProvidersOauth2InfoBindingType))
    fieldNameMap["oauth2"] = "Oauth2"
    fields["oidc"] = bindings.NewOptionalType(bindings.NewReferenceType(ProvidersOidcInfoBindingType))
    fieldNameMap["oidc"] = "Oidc"
    fields["is_default"] = bindings.NewBooleanType()
    fieldNameMap["is_default"] = "IsDefault"
    fields["domain_names"] = bindings.NewOptionalType(bindings.NewSetType(bindings.NewStringType(), reflect.TypeOf(map[string]bool{})))
    fieldNameMap["domain_names"] = "DomainNames"
    fields["auth_query_params"] = bindings.NewOptionalType(bindings.NewMapType(bindings.NewStringType(), bindings.NewListType(bindings.NewStringType(), reflect.TypeOf([]string{})),reflect.TypeOf(map[string][]string{})))
    fieldNameMap["auth_query_params"] = "AuthQueryParams"
    fields["idm_protocol"] = bindings.NewOptionalType(bindings.NewEnumType("com.vmware.vcenter.identity.providers.idm_protocol", reflect.TypeOf(Providers_IdmProtocol(Providers_IdmProtocol_REST))))
    fieldNameMap["idm_protocol"] = "IdmProtocol"
    fields["idm_endpoints"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewUriType(), reflect.TypeOf([]url.URL{})))
    fieldNameMap["idm_endpoints"] = "IdmEndpoints"
    fields["active_directory_over_ldap"] = bindings.NewOptionalType(bindings.NewReferenceType(ProvidersActiveDirectoryOverLdapBindingType))
    fieldNameMap["active_directory_over_ldap"] = "ActiveDirectoryOverLdap"
    fields["upn_claim"] = bindings.NewOptionalType(bindings.NewStringType())
    fieldNameMap["upn_claim"] = "UpnClaim"
    fields["groups_claim"] = bindings.NewOptionalType(bindings.NewStringType())
    fieldNameMap["groups_claim"] = "GroupsClaim"
    var validators = []bindings.Validator{}
    uv1 := bindings.NewUnionValidator("config_tag",
        map[string][]bindings.FieldData {
            "Oauth2": []bindings.FieldData {
                 bindings.NewFieldData("oauth2", true),
            },
            "Oidc": []bindings.FieldData {
                 bindings.NewFieldData("oidc", true),
            },
        },
    )
    validators = append(validators, uv1)
    uv2 := bindings.NewUnionValidator("idm_protocol",
        map[string][]bindings.FieldData {
            "REST": []bindings.FieldData {
                 bindings.NewFieldData("idm_endpoints", true),
            },
            "SCIM": []bindings.FieldData {
                 bindings.NewFieldData("idm_endpoints", true),
            },
            "LDAP": []bindings.FieldData {
                 bindings.NewFieldData("active_directory_over_ldap", true),
            },
        },
    )
    validators = append(validators, uv2)
    return bindings.NewStructType("com.vmware.vcenter.identity.providers.info",fields, reflect.TypeOf(ProvidersInfo{}), fieldNameMap, validators)
}

func ProvidersCreateSpecBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["config_tag"] = bindings.NewEnumType("com.vmware.vcenter.identity.providers.config_type", reflect.TypeOf(Providers_ConfigType(Providers_ConfigType_Oauth2)))
    fieldNameMap["config_tag"] = "ConfigTag"
    fields["oauth2"] = bindings.NewOptionalType(bindings.NewReferenceType(ProvidersOauth2CreateSpecBindingType))
    fieldNameMap["oauth2"] = "Oauth2"
    fields["oidc"] = bindings.NewOptionalType(bindings.NewReferenceType(ProvidersOidcCreateSpecBindingType))
    fieldNameMap["oidc"] = "Oidc"
    fields["org_ids"] = bindings.NewOptionalType(bindings.NewSetType(bindings.NewStringType(), reflect.TypeOf(map[string]bool{})))
    fieldNameMap["org_ids"] = "OrgIds"
    fields["is_default"] = bindings.NewOptionalType(bindings.NewBooleanType())
    fieldNameMap["is_default"] = "IsDefault"
    fields["name"] = bindings.NewOptionalType(bindings.NewStringType())
    fieldNameMap["name"] = "Name"
    fields["domain_names"] = bindings.NewOptionalType(bindings.NewSetType(bindings.NewStringType(), reflect.TypeOf(map[string]bool{})))
    fieldNameMap["domain_names"] = "DomainNames"
    fields["auth_query_params"] = bindings.NewOptionalType(bindings.NewMapType(bindings.NewStringType(), bindings.NewListType(bindings.NewStringType(), reflect.TypeOf([]string{})),reflect.TypeOf(map[string][]string{})))
    fieldNameMap["auth_query_params"] = "AuthQueryParams"
    fields["idm_protocol"] = bindings.NewOptionalType(bindings.NewEnumType("com.vmware.vcenter.identity.providers.idm_protocol", reflect.TypeOf(Providers_IdmProtocol(Providers_IdmProtocol_REST))))
    fieldNameMap["idm_protocol"] = "IdmProtocol"
    fields["idm_endpoints"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewUriType(), reflect.TypeOf([]url.URL{})))
    fieldNameMap["idm_endpoints"] = "IdmEndpoints"
    fields["active_directory_over_ldap"] = bindings.NewOptionalType(bindings.NewReferenceType(ProvidersActiveDirectoryOverLdapBindingType))
    fieldNameMap["active_directory_over_ldap"] = "ActiveDirectoryOverLdap"
    fields["upn_claim"] = bindings.NewOptionalType(bindings.NewStringType())
    fieldNameMap["upn_claim"] = "UpnClaim"
    fields["groups_claim"] = bindings.NewOptionalType(bindings.NewStringType())
    fieldNameMap["groups_claim"] = "GroupsClaim"
    var validators = []bindings.Validator{}
    uv1 := bindings.NewUnionValidator("config_tag",
        map[string][]bindings.FieldData {
            "Oauth2": []bindings.FieldData {
                 bindings.NewFieldData("oauth2", true),
            },
            "Oidc": []bindings.FieldData {
                 bindings.NewFieldData("oidc", true),
            },
        },
    )
    validators = append(validators, uv1)
    uv2 := bindings.NewUnionValidator("idm_protocol",
        map[string][]bindings.FieldData {
            "REST": []bindings.FieldData {
                 bindings.NewFieldData("idm_endpoints", true),
            },
            "SCIM": []bindings.FieldData {
                 bindings.NewFieldData("idm_endpoints", true),
            },
            "LDAP": []bindings.FieldData {
                 bindings.NewFieldData("active_directory_over_ldap", true),
            },
        },
    )
    validators = append(validators, uv2)
    return bindings.NewStructType("com.vmware.vcenter.identity.providers.create_spec",fields, reflect.TypeOf(ProvidersCreateSpec{}), fieldNameMap, validators)
}

func ProvidersUpdateSpecBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["config_tag"] = bindings.NewEnumType("com.vmware.vcenter.identity.providers.config_type", reflect.TypeOf(Providers_ConfigType(Providers_ConfigType_Oauth2)))
    fieldNameMap["config_tag"] = "ConfigTag"
    fields["oauth2"] = bindings.NewOptionalType(bindings.NewReferenceType(ProvidersOauth2UpdateSpecBindingType))
    fieldNameMap["oauth2"] = "Oauth2"
    fields["oidc"] = bindings.NewOptionalType(bindings.NewReferenceType(ProvidersOidcUpdateSpecBindingType))
    fieldNameMap["oidc"] = "Oidc"
    fields["org_ids"] = bindings.NewOptionalType(bindings.NewSetType(bindings.NewStringType(), reflect.TypeOf(map[string]bool{})))
    fieldNameMap["org_ids"] = "OrgIds"
    fields["make_default"] = bindings.NewOptionalType(bindings.NewBooleanType())
    fieldNameMap["make_default"] = "MakeDefault"
    fields["name"] = bindings.NewOptionalType(bindings.NewStringType())
    fieldNameMap["name"] = "Name"
    fields["domain_names"] = bindings.NewOptionalType(bindings.NewSetType(bindings.NewStringType(), reflect.TypeOf(map[string]bool{})))
    fieldNameMap["domain_names"] = "DomainNames"
    fields["auth_query_params"] = bindings.NewOptionalType(bindings.NewMapType(bindings.NewStringType(), bindings.NewListType(bindings.NewStringType(), reflect.TypeOf([]string{})),reflect.TypeOf(map[string][]string{})))
    fieldNameMap["auth_query_params"] = "AuthQueryParams"
    fields["idm_protocol"] = bindings.NewOptionalType(bindings.NewEnumType("com.vmware.vcenter.identity.providers.idm_protocol", reflect.TypeOf(Providers_IdmProtocol(Providers_IdmProtocol_REST))))
    fieldNameMap["idm_protocol"] = "IdmProtocol"
    fields["idm_endpoints"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewUriType(), reflect.TypeOf([]url.URL{})))
    fieldNameMap["idm_endpoints"] = "IdmEndpoints"
    fields["active_directory_over_ldap"] = bindings.NewOptionalType(bindings.NewReferenceType(ProvidersActiveDirectoryOverLdapBindingType))
    fieldNameMap["active_directory_over_ldap"] = "ActiveDirectoryOverLdap"
    fields["upn_claim"] = bindings.NewOptionalType(bindings.NewStringType())
    fieldNameMap["upn_claim"] = "UpnClaim"
    fields["reset_upn_claim"] = bindings.NewOptionalType(bindings.NewBooleanType())
    fieldNameMap["reset_upn_claim"] = "ResetUpnClaim"
    fields["groups_claim"] = bindings.NewOptionalType(bindings.NewStringType())
    fieldNameMap["groups_claim"] = "GroupsClaim"
    fields["reset_groups_claim"] = bindings.NewOptionalType(bindings.NewBooleanType())
    fieldNameMap["reset_groups_claim"] = "ResetGroupsClaim"
    var validators = []bindings.Validator{}
    uv1 := bindings.NewUnionValidator("config_tag",
        map[string][]bindings.FieldData {
            "Oauth2": []bindings.FieldData {
                 bindings.NewFieldData("oauth2", true),
            },
            "Oidc": []bindings.FieldData {
                 bindings.NewFieldData("oidc", true),
            },
        },
    )
    validators = append(validators, uv1)
    uv2 := bindings.NewUnionValidator("idm_protocol",
        map[string][]bindings.FieldData {
            "REST": []bindings.FieldData {
                 bindings.NewFieldData("idm_endpoints", true),
            },
            "SCIM": []bindings.FieldData {
                 bindings.NewFieldData("idm_endpoints", true),
            },
            "LDAP": []bindings.FieldData {
                 bindings.NewFieldData("active_directory_over_ldap", true),
            },
        },
    )
    validators = append(validators, uv2)
    return bindings.NewStructType("com.vmware.vcenter.identity.providers.update_spec",fields, reflect.TypeOf(ProvidersUpdateSpec{}), fieldNameMap, validators)
}

func ProvidersOauth2SummaryBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["auth_endpoint"] = bindings.NewUriType()
    fieldNameMap["auth_endpoint"] = "AuthEndpoint"
    fields["token_endpoint"] = bindings.NewUriType()
    fieldNameMap["token_endpoint"] = "TokenEndpoint"
    fields["client_id"] = bindings.NewStringType()
    fieldNameMap["client_id"] = "ClientId"
    fields["authentication_header"] = bindings.NewStringType()
    fieldNameMap["authentication_header"] = "AuthenticationHeader"
    fields["auth_query_params"] = bindings.NewMapType(bindings.NewStringType(), bindings.NewListType(bindings.NewStringType(), reflect.TypeOf([]string{})),reflect.TypeOf(map[string][]string{}))
    fieldNameMap["auth_query_params"] = "AuthQueryParams"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.identity.providers.oauth2_summary",fields, reflect.TypeOf(ProvidersOauth2Summary{}), fieldNameMap, validators)
}

func ProvidersOauth2InfoBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["auth_endpoint"] = bindings.NewUriType()
    fieldNameMap["auth_endpoint"] = "AuthEndpoint"
    fields["token_endpoint"] = bindings.NewUriType()
    fieldNameMap["token_endpoint"] = "TokenEndpoint"
    fields["public_key_uri"] = bindings.NewUriType()
    fieldNameMap["public_key_uri"] = "PublicKeyUri"
    fields["client_id"] = bindings.NewStringType()
    fieldNameMap["client_id"] = "ClientId"
    fields["client_secret"] = bindings.NewStringType()
    fieldNameMap["client_secret"] = "ClientSecret"
    fields["claim_map"] = bindings.NewMapType(bindings.NewStringType(), bindings.NewMapType(bindings.NewStringType(), bindings.NewListType(bindings.NewStringType(), reflect.TypeOf([]string{})),reflect.TypeOf(map[string][]string{})),reflect.TypeOf(map[string]map[string][]string{}))
    fieldNameMap["claim_map"] = "ClaimMap"
    fields["issuer"] = bindings.NewStringType()
    fieldNameMap["issuer"] = "Issuer"
    fields["authentication_method"] = bindings.NewEnumType("com.vmware.vcenter.identity.providers.oauth2_authentication_method", reflect.TypeOf(Providers_Oauth2AuthenticationMethod(Providers_Oauth2AuthenticationMethod_CLIENT_SECRET_BASIC)))
    fieldNameMap["authentication_method"] = "AuthenticationMethod"
    fields["auth_query_params"] = bindings.NewMapType(bindings.NewStringType(), bindings.NewListType(bindings.NewStringType(), reflect.TypeOf([]string{})),reflect.TypeOf(map[string][]string{}))
    fieldNameMap["auth_query_params"] = "AuthQueryParams"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.identity.providers.oauth2_info",fields, reflect.TypeOf(ProvidersOauth2Info{}), fieldNameMap, validators)
}

func ProvidersOauth2CreateSpecBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["auth_endpoint"] = bindings.NewUriType()
    fieldNameMap["auth_endpoint"] = "AuthEndpoint"
    fields["token_endpoint"] = bindings.NewUriType()
    fieldNameMap["token_endpoint"] = "TokenEndpoint"
    fields["public_key_uri"] = bindings.NewUriType()
    fieldNameMap["public_key_uri"] = "PublicKeyUri"
    fields["client_id"] = bindings.NewStringType()
    fieldNameMap["client_id"] = "ClientId"
    fields["client_secret"] = bindings.NewStringType()
    fieldNameMap["client_secret"] = "ClientSecret"
    fields["claim_map"] = bindings.NewMapType(bindings.NewStringType(), bindings.NewMapType(bindings.NewStringType(), bindings.NewListType(bindings.NewStringType(), reflect.TypeOf([]string{})),reflect.TypeOf(map[string][]string{})),reflect.TypeOf(map[string]map[string][]string{}))
    fieldNameMap["claim_map"] = "ClaimMap"
    fields["issuer"] = bindings.NewStringType()
    fieldNameMap["issuer"] = "Issuer"
    fields["authentication_method"] = bindings.NewEnumType("com.vmware.vcenter.identity.providers.oauth2_authentication_method", reflect.TypeOf(Providers_Oauth2AuthenticationMethod(Providers_Oauth2AuthenticationMethod_CLIENT_SECRET_BASIC)))
    fieldNameMap["authentication_method"] = "AuthenticationMethod"
    fields["auth_query_params"] = bindings.NewOptionalType(bindings.NewMapType(bindings.NewStringType(), bindings.NewListType(bindings.NewStringType(), reflect.TypeOf([]string{})),reflect.TypeOf(map[string][]string{})))
    fieldNameMap["auth_query_params"] = "AuthQueryParams"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.identity.providers.oauth2_create_spec",fields, reflect.TypeOf(ProvidersOauth2CreateSpec{}), fieldNameMap, validators)
}

func ProvidersOauth2UpdateSpecBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["auth_endpoint"] = bindings.NewOptionalType(bindings.NewUriType())
    fieldNameMap["auth_endpoint"] = "AuthEndpoint"
    fields["token_endpoint"] = bindings.NewOptionalType(bindings.NewUriType())
    fieldNameMap["token_endpoint"] = "TokenEndpoint"
    fields["public_key_uri"] = bindings.NewOptionalType(bindings.NewUriType())
    fieldNameMap["public_key_uri"] = "PublicKeyUri"
    fields["client_id"] = bindings.NewOptionalType(bindings.NewStringType())
    fieldNameMap["client_id"] = "ClientId"
    fields["client_secret"] = bindings.NewOptionalType(bindings.NewStringType())
    fieldNameMap["client_secret"] = "ClientSecret"
    fields["claim_map"] = bindings.NewOptionalType(bindings.NewMapType(bindings.NewStringType(), bindings.NewMapType(bindings.NewStringType(), bindings.NewListType(bindings.NewStringType(), reflect.TypeOf([]string{})),reflect.TypeOf(map[string][]string{})),reflect.TypeOf(map[string]map[string][]string{})))
    fieldNameMap["claim_map"] = "ClaimMap"
    fields["issuer"] = bindings.NewOptionalType(bindings.NewStringType())
    fieldNameMap["issuer"] = "Issuer"
    fields["authentication_method"] = bindings.NewOptionalType(bindings.NewEnumType("com.vmware.vcenter.identity.providers.oauth2_authentication_method", reflect.TypeOf(Providers_Oauth2AuthenticationMethod(Providers_Oauth2AuthenticationMethod_CLIENT_SECRET_BASIC))))
    fieldNameMap["authentication_method"] = "AuthenticationMethod"
    fields["auth_query_params"] = bindings.NewOptionalType(bindings.NewMapType(bindings.NewStringType(), bindings.NewListType(bindings.NewStringType(), reflect.TypeOf([]string{})),reflect.TypeOf(map[string][]string{})))
    fieldNameMap["auth_query_params"] = "AuthQueryParams"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.identity.providers.oauth2_update_spec",fields, reflect.TypeOf(ProvidersOauth2UpdateSpec{}), fieldNameMap, validators)
}

func ProvidersOidcSummaryBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["discovery_endpoint"] = bindings.NewOptionalType(bindings.NewUriType())
    fieldNameMap["discovery_endpoint"] = "DiscoveryEndpoint"
    fields["logout_endpoint"] = bindings.NewOptionalType(bindings.NewUriType())
    fieldNameMap["logout_endpoint"] = "LogoutEndpoint"
    fields["auth_endpoint"] = bindings.NewUriType()
    fieldNameMap["auth_endpoint"] = "AuthEndpoint"
    fields["token_endpoint"] = bindings.NewUriType()
    fieldNameMap["token_endpoint"] = "TokenEndpoint"
    fields["client_id"] = bindings.NewStringType()
    fieldNameMap["client_id"] = "ClientId"
    fields["authentication_header"] = bindings.NewStringType()
    fieldNameMap["authentication_header"] = "AuthenticationHeader"
    fields["auth_query_params"] = bindings.NewMapType(bindings.NewStringType(), bindings.NewListType(bindings.NewStringType(), reflect.TypeOf([]string{})),reflect.TypeOf(map[string][]string{}))
    fieldNameMap["auth_query_params"] = "AuthQueryParams"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.identity.providers.oidc_summary",fields, reflect.TypeOf(ProvidersOidcSummary{}), fieldNameMap, validators)
}

func ProvidersOidcInfoBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["discovery_endpoint"] = bindings.NewUriType()
    fieldNameMap["discovery_endpoint"] = "DiscoveryEndpoint"
    fields["logout_endpoint"] = bindings.NewOptionalType(bindings.NewUriType())
    fieldNameMap["logout_endpoint"] = "LogoutEndpoint"
    fields["auth_endpoint"] = bindings.NewUriType()
    fieldNameMap["auth_endpoint"] = "AuthEndpoint"
    fields["token_endpoint"] = bindings.NewUriType()
    fieldNameMap["token_endpoint"] = "TokenEndpoint"
    fields["public_key_uri"] = bindings.NewUriType()
    fieldNameMap["public_key_uri"] = "PublicKeyUri"
    fields["client_id"] = bindings.NewStringType()
    fieldNameMap["client_id"] = "ClientId"
    fields["client_secret"] = bindings.NewStringType()
    fieldNameMap["client_secret"] = "ClientSecret"
    fields["claim_map"] = bindings.NewMapType(bindings.NewStringType(), bindings.NewMapType(bindings.NewStringType(), bindings.NewListType(bindings.NewStringType(), reflect.TypeOf([]string{})),reflect.TypeOf(map[string][]string{})),reflect.TypeOf(map[string]map[string][]string{}))
    fieldNameMap["claim_map"] = "ClaimMap"
    fields["issuer"] = bindings.NewStringType()
    fieldNameMap["issuer"] = "Issuer"
    fields["authentication_method"] = bindings.NewEnumType("com.vmware.vcenter.identity.providers.oauth2_authentication_method", reflect.TypeOf(Providers_Oauth2AuthenticationMethod(Providers_Oauth2AuthenticationMethod_CLIENT_SECRET_BASIC)))
    fieldNameMap["authentication_method"] = "AuthenticationMethod"
    fields["auth_query_params"] = bindings.NewMapType(bindings.NewStringType(), bindings.NewListType(bindings.NewStringType(), reflect.TypeOf([]string{})),reflect.TypeOf(map[string][]string{}))
    fieldNameMap["auth_query_params"] = "AuthQueryParams"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.identity.providers.oidc_info",fields, reflect.TypeOf(ProvidersOidcInfo{}), fieldNameMap, validators)
}

func ProvidersOidcCreateSpecBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["discovery_endpoint"] = bindings.NewUriType()
    fieldNameMap["discovery_endpoint"] = "DiscoveryEndpoint"
    fields["client_id"] = bindings.NewStringType()
    fieldNameMap["client_id"] = "ClientId"
    fields["client_secret"] = bindings.NewStringType()
    fieldNameMap["client_secret"] = "ClientSecret"
    fields["claim_map"] = bindings.NewMapType(bindings.NewStringType(), bindings.NewMapType(bindings.NewStringType(), bindings.NewListType(bindings.NewStringType(), reflect.TypeOf([]string{})),reflect.TypeOf(map[string][]string{})),reflect.TypeOf(map[string]map[string][]string{}))
    fieldNameMap["claim_map"] = "ClaimMap"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.identity.providers.oidc_create_spec",fields, reflect.TypeOf(ProvidersOidcCreateSpec{}), fieldNameMap, validators)
}

func ProvidersOidcUpdateSpecBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["discovery_endpoint"] = bindings.NewOptionalType(bindings.NewUriType())
    fieldNameMap["discovery_endpoint"] = "DiscoveryEndpoint"
    fields["client_id"] = bindings.NewOptionalType(bindings.NewStringType())
    fieldNameMap["client_id"] = "ClientId"
    fields["client_secret"] = bindings.NewOptionalType(bindings.NewStringType())
    fieldNameMap["client_secret"] = "ClientSecret"
    fields["claim_map"] = bindings.NewOptionalType(bindings.NewMapType(bindings.NewStringType(), bindings.NewMapType(bindings.NewStringType(), bindings.NewListType(bindings.NewStringType(), reflect.TypeOf([]string{})),reflect.TypeOf(map[string][]string{})),reflect.TypeOf(map[string]map[string][]string{})))
    fieldNameMap["claim_map"] = "ClaimMap"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.identity.providers.oidc_update_spec",fields, reflect.TypeOf(ProvidersOidcUpdateSpec{}), fieldNameMap, validators)
}

func ProvidersActiveDirectoryOverLdapBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["user_name"] = bindings.NewStringType()
    fieldNameMap["user_name"] = "UserName"
    fields["password"] = bindings.NewSecretType()
    fieldNameMap["password"] = "Password"
    fields["users_base_dn"] = bindings.NewStringType()
    fieldNameMap["users_base_dn"] = "UsersBaseDn"
    fields["groups_base_dn"] = bindings.NewStringType()
    fieldNameMap["groups_base_dn"] = "GroupsBaseDn"
    fields["server_endpoints"] = bindings.NewListType(bindings.NewUriType(), reflect.TypeOf([]url.URL{}))
    fieldNameMap["server_endpoints"] = "ServerEndpoints"
    fields["cert_chain"] = bindings.NewOptionalType(bindings.NewReferenceType(certificate_management.X509CertChainBindingType))
    fieldNameMap["cert_chain"] = "CertChain"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.identity.providers.active_directory_over_ldap",fields, reflect.TypeOf(ProvidersActiveDirectoryOverLdap{}), fieldNameMap, validators)
}


