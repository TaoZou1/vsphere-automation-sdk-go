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

package kms

import (
    "reflect"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/bindings"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/data"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/protocol"
)

// Resource type for a Key Provider.
const Providers_RESOURCE_TYPE = "com.vmware.esx.kms.providers"



// The ``NetworkAddress`` class contains an IP address or hostname and a port on which a connection can be established.
 type ProvidersNetworkAddress struct {
    // The IP address or the hostname.
    Hostname string
    // The port for the connection.
    Port *int64
}


// TODO some error doesn't have ErrorType. Resolve this later VAPI-3009
//


// The ``Server`` class contains properties that describe a connection endpoint.
 type ProvidersServer struct {
    // Name of the server. 
//
//  A unique string chosen by the client.
    Name string
    // The server's address.
    Address ProvidersNetworkAddress
}


// TODO some error doesn't have ErrorType. Resolve this later VAPI-3009
//


// The ``KmipServerCreateSpec`` class contains properties that describe the desired configuration for KMIP based key server.
 type ProvidersKmipServerCreateSpec struct {
    // List of KMIP compliant key servers. 
//
//  Key Servers must be configured for active-active replication. If the server port is null, a default value of 5696 will be used.
    Servers []ProvidersServer
    // Username for authentication. 
    Username *string
}


// TODO some error doesn't have ErrorType. Resolve this later VAPI-3009
//


// The ``KeyServerCreateSpec`` class contains properties that describe the desired configuration for key server.
 type ProvidersKeyServerCreateSpec struct {
    // Type of the key server.
    Type_ ProvidersKeyServerCreateSpec_Type
    // Description of the key server. 
    Description *string
    // Proxy server configuration. 
    ProxyServer *ProvidersNetworkAddress
    // Connection timeout in seconds. 
    ConnectionTimeout *int64
    // Key Server configuration detail for KMIP based KMS. 
    KmipServer *ProvidersKmipServerCreateSpec
}


// TODO some error doesn't have ErrorType. Resolve this later VAPI-3009
//
    
    // The ``Type`` enumeration class list the key server types.
    //
    // <p> See {@link com.vmware.vapi.bindings.ApiEnumeration enumerated types description}.
     
    type ProvidersKeyServerCreateSpec_Type string

    const (
        // Key Management Interoperability Protocol (KMIP) based key management server.
         ProvidersKeyServerCreateSpec_Type_KMIP ProvidersKeyServerCreateSpec_Type = "KMIP"
    )

    func (t ProvidersKeyServerCreateSpec_Type) ProvidersKeyServerCreateSpec_Type() bool {
        switch t {
            case ProvidersKeyServerCreateSpec_Type_KMIP:
                return true
            default:
                return false
        }
    }



// The ``CreateSpec`` class contains properties that describe the desired configuration for a new provider.
 type ProvidersCreateSpec struct {
    // Name of the provider. 
//
//  A unique string chosen by the client.
    Provider string
    // Master key identifier created for the provider. 
//
//  A unique Key identifier.
    MasterKeyId string
    // Key server associated with this provider.
    KeyServer ProvidersKeyServerCreateSpec
}


// TODO some error doesn't have ErrorType. Resolve this later VAPI-3009
//


// The ``KmipServerUpdateSpec`` class contains properties that describe new configuration for KMIP based key server.
 type ProvidersKmipServerUpdateSpec struct {
    // List of KMIP compliant key servers. 
//
//  Key Servers must be configured for active-active replication. If the server port is null, a default value of 5696 will be used. 
    Servers []ProvidersServer
    // Username for authentication. 
    Username *string
}


// TODO some error doesn't have ErrorType. Resolve this later VAPI-3009
//


// The ``KeyServerUpdateSpec`` class contains properties that describe new configuration for an existing key server.
 type ProvidersKeyServerUpdateSpec struct {
    // Type of the key server. 
    Type_ *ProvidersKeyServerUpdateSpec_Type
    // Description of the key server. 
    Description *string
    // Proxy server configuration. 
    ProxyServer *ProvidersNetworkAddress
    // Connection timeout in seconds. 
    ConnectionTimeout *int64
    // Key Server configuration detail for KMIP based KMS. 
    KmipServer *ProvidersKmipServerUpdateSpec
}


// TODO some error doesn't have ErrorType. Resolve this later VAPI-3009
//
    
    // The ``Type`` enumeration class list the {\\\\@name KeyServer) types.
    //
    // <p> See {@link com.vmware.vapi.bindings.ApiEnumeration enumerated types description}.
     
    type ProvidersKeyServerUpdateSpec_Type string

    const (
        // Key Management Interoperability Protocol (KMIP) based key management server.
         ProvidersKeyServerUpdateSpec_Type_KMIP ProvidersKeyServerUpdateSpec_Type = "KMIP"
    )

    func (t ProvidersKeyServerUpdateSpec_Type) ProvidersKeyServerUpdateSpec_Type() bool {
        switch t {
            case ProvidersKeyServerUpdateSpec_Type_KMIP:
                return true
            default:
                return false
        }
    }



// The ``UpdateSpec`` class contains properties that describe the new configuration for an existing provider.
 type ProvidersUpdateSpec struct {
    // Master key identifier created for the provider. 
//
//  A unique Key identifier. 
    MasterKeyId *string
    // Key server associated with this provider. 
    KeyServer *ProvidersKeyServerUpdateSpec
}


// TODO some error doesn't have ErrorType. Resolve this later VAPI-3009
//


// The ``Summary`` class contains properties that summarize a provider.
 type ProvidersSummary struct {
    // Name of the provider. 
//
//  A unique string chosen by the client.
    Provider string
}


// TODO some error doesn't have ErrorType. Resolve this later VAPI-3009
//


// The ``KmipServerInfo`` class contains properties that describe the current configuration of a KMIP based key server.
 type ProvidersKmipServerInfo struct {
    // List of KMIP compliant key servers.
    Servers []ProvidersServer
    // Username for authentication. 
    Username *string
}


// TODO some error doesn't have ErrorType. Resolve this later VAPI-3009
//


// The ``KeyServerInfo`` class contains properties that describe the current configuration of a key server.
 type ProvidersKeyServerInfo struct {
    // Type of the key server.
    Type_ ProvidersKeyServerInfo_Type
    // Description of the key server.
    Description string
    // Proxy server configuration. 
    ProxyServer *ProvidersNetworkAddress
    // Connection timeout in seconds. 
    ConnectionTimeout *int64
    // Key Server configuration detail for KMIP based KMS.
    KmipServer *ProvidersKmipServerInfo
}


// TODO some error doesn't have ErrorType. Resolve this later VAPI-3009
//
    
    // The ``Type`` enumeration class list the {\\\\@name KeyServer) types.
    //
    // <p> See {@link com.vmware.vapi.bindings.ApiEnumeration enumerated types description}.
     
    type ProvidersKeyServerInfo_Type string

    const (
        // Key Management Interoperability Protocol (KMIP) based key management server.
         ProvidersKeyServerInfo_Type_KMIP ProvidersKeyServerInfo_Type = "KMIP"
    )

    func (t ProvidersKeyServerInfo_Type) ProvidersKeyServerInfo_Type() bool {
        switch t {
            case ProvidersKeyServerInfo_Type_KMIP:
                return true
            default:
                return false
        }
    }



// The ``Info`` class contains properties that describe the current configuration of a provider.
 type ProvidersInfo struct {
    // Master key identifier created for the provider. 
//
//  A unique Key identifier.
    MasterKeyId string
    // Key server associated with this provider.
    KeyServer ProvidersKeyServerInfo
}


// TODO some error doesn't have ErrorType. Resolve this later VAPI-3009
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
    errorHeaders["Unauthenticated.challenge"] = "WWW-Authenticate"
    return protocol.NewOperationRestMetadata(
      paramsTypeMap,
      pathParams,
      queryParams,
      headerParams,
      "",
      "GET",
      "/esx/kms/providers",
       resultHeaders,
       200,
       errorHeaders,
       map[string]int{"Unauthenticated": 401,"Unauthorized": 403,"Error": 500})
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
    return bindings.NewVoidType()
}

func providersCreateRestMetadata() protocol.OperationRestMetadata {
    paramsTypeMap := map[string]bindings.BindingType{}
    pathParams := map[string]string{}
    queryParams := map[string]string{}
    headerParams := map[string]string{}
    paramsTypeMap["spec"] = bindings.NewReferenceType(ProvidersCreateSpecBindingType)
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
      "/esx/kms/providers",
       resultHeaders,
       204,
       errorHeaders,
       map[string]int{"AlreadyExists": 400,"InvalidArgument": 400,"Unauthenticated": 401,"Unauthorized": 403,"Error": 500})
}


func providersUpdateInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["provider"] = bindings.NewIdType([]string {"com.vmware.esx.kms.providers"}, "")
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
    paramsTypeMap["provider"] = bindings.NewIdType([]string {"com.vmware.esx.kms.providers"}, "")
    paramsTypeMap["spec"] = bindings.NewReferenceType(ProvidersUpdateSpecBindingType)
    paramsTypeMap["provider"] = bindings.NewIdType([]string {"com.vmware.esx.kms.providers"}, "")
    pathParams["provider"] = "provider"
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
      "/esx/kms/providers/{provider}",
       resultHeaders,
       204,
       errorHeaders,
       map[string]int{"InvalidArgument": 400,"NotFound": 404,"Unauthenticated": 401,"Unauthorized": 403,"Error": 500})
}


func providersDeleteInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["provider"] = bindings.NewIdType([]string {"com.vmware.esx.kms.providers"}, "")
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
    paramsTypeMap["provider"] = bindings.NewIdType([]string {"com.vmware.esx.kms.providers"}, "")
    paramsTypeMap["provider"] = bindings.NewIdType([]string {"com.vmware.esx.kms.providers"}, "")
    pathParams["provider"] = "provider"
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
      "/esx/kms/providers/{provider}",
       resultHeaders,
       204,
       errorHeaders,
       map[string]int{"InvalidArgument": 400,"NotFound": 404,"Unauthenticated": 401,"Unauthorized": 403,"Error": 500})
}


func providersGetInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["provider"] = bindings.NewIdType([]string {"com.vmware.esx.kms.providers"}, "")
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
    paramsTypeMap["provider"] = bindings.NewIdType([]string {"com.vmware.esx.kms.providers"}, "")
    paramsTypeMap["provider"] = bindings.NewIdType([]string {"com.vmware.esx.kms.providers"}, "")
    pathParams["provider"] = "provider"
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
      "/esx/kms/providers/{provider}",
       resultHeaders,
       200,
       errorHeaders,
       map[string]int{"InvalidArgument": 400,"NotFound": 404,"Unauthenticated": 401,"Unauthorized": 403,"Error": 500})
}



func ProvidersNetworkAddressBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["hostname"] = bindings.NewStringType()
    fieldNameMap["hostname"] = "Hostname"
    fields["port"] = bindings.NewOptionalType(bindings.NewIntegerType())
    fieldNameMap["port"] = "Port"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.esx.kms.providers.network_address",fields, reflect.TypeOf(ProvidersNetworkAddress{}), fieldNameMap, validators)
}

func ProvidersServerBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["name"] = bindings.NewStringType()
    fieldNameMap["name"] = "Name"
    fields["address"] = bindings.NewReferenceType(ProvidersNetworkAddressBindingType)
    fieldNameMap["address"] = "Address"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.esx.kms.providers.server",fields, reflect.TypeOf(ProvidersServer{}), fieldNameMap, validators)
}

func ProvidersKmipServerCreateSpecBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["servers"] = bindings.NewListType(bindings.NewReferenceType(ProvidersServerBindingType), reflect.TypeOf([]ProvidersServer{}))
    fieldNameMap["servers"] = "Servers"
    fields["username"] = bindings.NewOptionalType(bindings.NewStringType())
    fieldNameMap["username"] = "Username"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.esx.kms.providers.kmip_server_create_spec",fields, reflect.TypeOf(ProvidersKmipServerCreateSpec{}), fieldNameMap, validators)
}

func ProvidersKeyServerCreateSpecBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["type"] = bindings.NewEnumType("com.vmware.esx.kms.providers.key_server_create_spec.type", reflect.TypeOf(ProvidersKeyServerCreateSpec_Type(ProvidersKeyServerCreateSpec_Type_KMIP)))
    fieldNameMap["type"] = "Type_"
    fields["description"] = bindings.NewOptionalType(bindings.NewStringType())
    fieldNameMap["description"] = "Description"
    fields["proxy_server"] = bindings.NewOptionalType(bindings.NewReferenceType(ProvidersNetworkAddressBindingType))
    fieldNameMap["proxy_server"] = "ProxyServer"
    fields["connection_timeout"] = bindings.NewOptionalType(bindings.NewIntegerType())
    fieldNameMap["connection_timeout"] = "ConnectionTimeout"
    fields["kmip_server"] = bindings.NewOptionalType(bindings.NewReferenceType(ProvidersKmipServerCreateSpecBindingType))
    fieldNameMap["kmip_server"] = "KmipServer"
    var validators = []bindings.Validator{}
    uv1 := bindings.NewUnionValidator("type",
        map[string][]bindings.FieldData {
            "KMIP": []bindings.FieldData {
                 bindings.NewFieldData("kmip_server", true),
            },
        },
    )
    validators = append(validators, uv1)
    return bindings.NewStructType("com.vmware.esx.kms.providers.key_server_create_spec",fields, reflect.TypeOf(ProvidersKeyServerCreateSpec{}), fieldNameMap, validators)
}

func ProvidersCreateSpecBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["provider"] = bindings.NewIdType([]string {"com.vmware.esx.kms.providers"}, "")
    fieldNameMap["provider"] = "Provider"
    fields["master_key_id"] = bindings.NewStringType()
    fieldNameMap["master_key_id"] = "MasterKeyId"
    fields["key_server"] = bindings.NewReferenceType(ProvidersKeyServerCreateSpecBindingType)
    fieldNameMap["key_server"] = "KeyServer"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.esx.kms.providers.create_spec",fields, reflect.TypeOf(ProvidersCreateSpec{}), fieldNameMap, validators)
}

func ProvidersKmipServerUpdateSpecBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["servers"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewReferenceType(ProvidersServerBindingType), reflect.TypeOf([]ProvidersServer{})))
    fieldNameMap["servers"] = "Servers"
    fields["username"] = bindings.NewOptionalType(bindings.NewStringType())
    fieldNameMap["username"] = "Username"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.esx.kms.providers.kmip_server_update_spec",fields, reflect.TypeOf(ProvidersKmipServerUpdateSpec{}), fieldNameMap, validators)
}

func ProvidersKeyServerUpdateSpecBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["type"] = bindings.NewOptionalType(bindings.NewEnumType("com.vmware.esx.kms.providers.key_server_update_spec.type", reflect.TypeOf(ProvidersKeyServerUpdateSpec_Type(ProvidersKeyServerUpdateSpec_Type_KMIP))))
    fieldNameMap["type"] = "Type_"
    fields["description"] = bindings.NewOptionalType(bindings.NewStringType())
    fieldNameMap["description"] = "Description"
    fields["proxy_server"] = bindings.NewOptionalType(bindings.NewReferenceType(ProvidersNetworkAddressBindingType))
    fieldNameMap["proxy_server"] = "ProxyServer"
    fields["connection_timeout"] = bindings.NewOptionalType(bindings.NewIntegerType())
    fieldNameMap["connection_timeout"] = "ConnectionTimeout"
    fields["kmip_server"] = bindings.NewOptionalType(bindings.NewReferenceType(ProvidersKmipServerUpdateSpecBindingType))
    fieldNameMap["kmip_server"] = "KmipServer"
    var validators = []bindings.Validator{}
    uv1 := bindings.NewUnionValidator("type",
        map[string][]bindings.FieldData {
            "KMIP": []bindings.FieldData {
                 bindings.NewFieldData("kmip_server", false),
            },
        },
    )
    validators = append(validators, uv1)
    return bindings.NewStructType("com.vmware.esx.kms.providers.key_server_update_spec",fields, reflect.TypeOf(ProvidersKeyServerUpdateSpec{}), fieldNameMap, validators)
}

func ProvidersUpdateSpecBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["master_key_id"] = bindings.NewOptionalType(bindings.NewStringType())
    fieldNameMap["master_key_id"] = "MasterKeyId"
    fields["key_server"] = bindings.NewOptionalType(bindings.NewReferenceType(ProvidersKeyServerUpdateSpecBindingType))
    fieldNameMap["key_server"] = "KeyServer"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.esx.kms.providers.update_spec",fields, reflect.TypeOf(ProvidersUpdateSpec{}), fieldNameMap, validators)
}

func ProvidersSummaryBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["provider"] = bindings.NewIdType([]string {"com.vmware.esx.kms.providers"}, "")
    fieldNameMap["provider"] = "Provider"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.esx.kms.providers.summary",fields, reflect.TypeOf(ProvidersSummary{}), fieldNameMap, validators)
}

func ProvidersKmipServerInfoBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["servers"] = bindings.NewListType(bindings.NewReferenceType(ProvidersServerBindingType), reflect.TypeOf([]ProvidersServer{}))
    fieldNameMap["servers"] = "Servers"
    fields["username"] = bindings.NewOptionalType(bindings.NewStringType())
    fieldNameMap["username"] = "Username"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.esx.kms.providers.kmip_server_info",fields, reflect.TypeOf(ProvidersKmipServerInfo{}), fieldNameMap, validators)
}

func ProvidersKeyServerInfoBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["type"] = bindings.NewEnumType("com.vmware.esx.kms.providers.key_server_info.type", reflect.TypeOf(ProvidersKeyServerInfo_Type(ProvidersKeyServerInfo_Type_KMIP)))
    fieldNameMap["type"] = "Type_"
    fields["description"] = bindings.NewStringType()
    fieldNameMap["description"] = "Description"
    fields["proxy_server"] = bindings.NewOptionalType(bindings.NewReferenceType(ProvidersNetworkAddressBindingType))
    fieldNameMap["proxy_server"] = "ProxyServer"
    fields["connection_timeout"] = bindings.NewOptionalType(bindings.NewIntegerType())
    fieldNameMap["connection_timeout"] = "ConnectionTimeout"
    fields["kmip_server"] = bindings.NewOptionalType(bindings.NewReferenceType(ProvidersKmipServerInfoBindingType))
    fieldNameMap["kmip_server"] = "KmipServer"
    var validators = []bindings.Validator{}
    uv1 := bindings.NewUnionValidator("type",
        map[string][]bindings.FieldData {
            "KMIP": []bindings.FieldData {
                 bindings.NewFieldData("kmip_server", true),
            },
        },
    )
    validators = append(validators, uv1)
    return bindings.NewStructType("com.vmware.esx.kms.providers.key_server_info",fields, reflect.TypeOf(ProvidersKeyServerInfo{}), fieldNameMap, validators)
}

func ProvidersInfoBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["master_key_id"] = bindings.NewStringType()
    fieldNameMap["master_key_id"] = "MasterKeyId"
    fields["key_server"] = bindings.NewReferenceType(ProvidersKeyServerInfoBindingType)
    fieldNameMap["key_server"] = "KeyServer"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.esx.kms.providers.info",fields, reflect.TypeOf(ProvidersInfo{}), fieldNameMap, validators)
}


