/* Copyright © 2019 VMware, Inc. All Rights Reserved.
     SPDX-License-Identifier: BSD-2-Clause */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Data type definitions file for service: Community.
 * Includes binding types of a structures and enumerations defined in the service.
 * Shared by client-side stubs and server-side skeletons to ensure type
 * compatibility.
 */

package hlm

import (
    "reflect"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/bindings/com/vmware/vapi/std"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/bindings"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/data"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/protocol"
)




// The ``AddSpec`` class is the specification used for connecting to another node and adding it to the local community. **Warning:** This class is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
 type CommunityAddSpec struct {
    // The hostname for the target node. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    Hostname string
    // The HTTPS port of the target node. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    Port *string
    // The base64 encoded authentication token usable in the target node's domain and issued by its Security Token Service service. The token should be prefixed with a type. Currently, only saml tokens are supported. For example: "sign=LKJSDLKDJ..." where prefix is "sign=" for a saml token. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    Token string
    // The SHA-256 thumbprint of the SSL certificate of the target node. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    SslThumbprint *string
    // SSL verification should be enabled or disabled. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    SslVerify *bool
    // The list of groups to be added to which administrator access should be granted. Group names should be un UPN format. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    AdminGroups map[string]bool
}


// TODO some error doesn't have ErrorType. Resolve this later VAPI-3009
//


// The ``AddCheckSpec`` class is the specification used for specifying how to perform network checks within a community before adding a new peer. **Warning:** This class is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
 type CommunityAddCheckSpec struct {
    // The specification used to add a node to the community. This will be used to verify pre-add that all network requirements have been met. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    AddSpec CommunityAddSpec
    // If true, network checks will be done between all nodes in the existing community with the node to be added. Otherwise, the checks will only be done between the local and target nodes. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    CheckAll bool
}


// TODO some error doesn't have ErrorType. Resolve this later VAPI-3009
//


// The ``CertificateInfo`` class contains information about the SSL certificate for a target node. **Warning:** This class is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
 type CommunityCertificateInfo struct {
    // The SHA-256 thumbprint of the SSL certificate for a target node. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    SslThumbprint string
}


// TODO some error doesn't have ErrorType. Resolve this later VAPI-3009
//


// The ``Info`` class has information about the community. **Warning:** This class is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
 type CommunityInfo struct {
    // The members in this community. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    Members []CommunityMember
}


// TODO some error doesn't have ErrorType. Resolve this later VAPI-3009
//


// The ``Member`` class has information about a community member. **Warning:** This class is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
 type CommunityMember struct {
    // The member's hostname. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    Hostname string
}


// TODO some error doesn't have ErrorType. Resolve this later VAPI-3009
//


// The ``CheckInfo`` class contains information on the {\\\\@operation} check. Connectivity checks are validated pairwise. **Warning:** This class is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
 type CommunityCheckInfo struct {
    // The status of the connection from the destination to the source. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    StatusFromDestination CommunityStatus
    // The status of the connection from the source to the destination. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    StatusFromSource CommunityStatus
    // The overall status. This will be true if the status of the connection from the destination to the source or from the source to the destination failed. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    Failed bool
}


// TODO some error doesn't have ErrorType. Resolve this later VAPI-3009
//


// The ``Status`` class contains properties that are used to describe the outcome of a single direction connectivity check. **Warning:** This class is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
 type CommunityStatus struct {
    // The hostname of the node issuing network calls. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    Source string
    // The hostname of the destination node. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    Destination string
    // The latency in milliseconds from source to destination. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    Latency int64
    // This is set to true if there are any error messages in the per path network checks. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    Failed bool
    // The list of any error messages. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    Errors []std.LocalizableMessage
}


// TODO some error doesn't have ErrorType. Resolve this later VAPI-3009
//





func communityGetInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func communityGetOutputType() bindings.BindingType {
    return bindings.NewReferenceType(CommunityInfoBindingType)
}

func communityGetRestMetadata() protocol.OperationRestMetadata {
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


func communityAddInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["spec"] = bindings.NewReferenceType(CommunityAddSpecBindingType)
    fieldNameMap["spec"] = "Spec"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func communityAddOutputType() bindings.BindingType {
    return bindings.NewVoidType()
}

func communityAddRestMetadata() protocol.OperationRestMetadata {
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
       map[string]int{"AlreadyExists": 400,"InvalidArgument": 400,"Unauthenticated": 401,"Unauthorized": 403,"UnverifiedPeer": 400,"Error": 500})
}


func communityRemoveInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["hostname"] = bindings.NewStringType()
    fieldNameMap["hostname"] = "Hostname"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func communityRemoveOutputType() bindings.BindingType {
    return bindings.NewVoidType()
}

func communityRemoveRestMetadata() protocol.OperationRestMetadata {
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


func communityCheckInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["spec"] = bindings.NewOptionalType(bindings.NewReferenceType(CommunityAddCheckSpecBindingType))
    fieldNameMap["spec"] = "Spec"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func communityCheckOutputType() bindings.BindingType {
    return bindings.NewListType(bindings.NewReferenceType(CommunityCheckInfoBindingType), reflect.TypeOf([]CommunityCheckInfo{}))
}

func communityCheckRestMetadata() protocol.OperationRestMetadata {
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
       map[string]int{"Unauthenticated": 401,"Unauthorized": 403,"UnverifiedPeer": 400,"Error": 500})
}



func CommunityAddSpecBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["hostname"] = bindings.NewStringType()
    fieldNameMap["hostname"] = "Hostname"
    fields["port"] = bindings.NewOptionalType(bindings.NewStringType())
    fieldNameMap["port"] = "Port"
    fields["token"] = bindings.NewStringType()
    fieldNameMap["token"] = "Token"
    fields["ssl_thumbprint"] = bindings.NewOptionalType(bindings.NewStringType())
    fieldNameMap["ssl_thumbprint"] = "SslThumbprint"
    fields["ssl_verify"] = bindings.NewOptionalType(bindings.NewBooleanType())
    fieldNameMap["ssl_verify"] = "SslVerify"
    fields["admin_groups"] = bindings.NewOptionalType(bindings.NewSetType(bindings.NewStringType(), reflect.TypeOf(map[string]bool{})))
    fieldNameMap["admin_groups"] = "AdminGroups"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.hlm.community.add_spec",fields, reflect.TypeOf(CommunityAddSpec{}), fieldNameMap, validators)
}

func CommunityAddCheckSpecBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["add_spec"] = bindings.NewReferenceType(CommunityAddSpecBindingType)
    fieldNameMap["add_spec"] = "AddSpec"
    fields["check_all"] = bindings.NewBooleanType()
    fieldNameMap["check_all"] = "CheckAll"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.hlm.community.add_check_spec",fields, reflect.TypeOf(CommunityAddCheckSpec{}), fieldNameMap, validators)
}

func CommunityCertificateInfoBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["ssl_thumbprint"] = bindings.NewStringType()
    fieldNameMap["ssl_thumbprint"] = "SslThumbprint"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.hlm.community.certificate_info",fields, reflect.TypeOf(CommunityCertificateInfo{}), fieldNameMap, validators)
}

func CommunityInfoBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["members"] = bindings.NewListType(bindings.NewReferenceType(CommunityMemberBindingType), reflect.TypeOf([]CommunityMember{}))
    fieldNameMap["members"] = "Members"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.hlm.community.info",fields, reflect.TypeOf(CommunityInfo{}), fieldNameMap, validators)
}

func CommunityMemberBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["hostname"] = bindings.NewStringType()
    fieldNameMap["hostname"] = "Hostname"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.hlm.community.member",fields, reflect.TypeOf(CommunityMember{}), fieldNameMap, validators)
}

func CommunityCheckInfoBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["status_from_destination"] = bindings.NewReferenceType(CommunityStatusBindingType)
    fieldNameMap["status_from_destination"] = "StatusFromDestination"
    fields["status_from_source"] = bindings.NewReferenceType(CommunityStatusBindingType)
    fieldNameMap["status_from_source"] = "StatusFromSource"
    fields["failed"] = bindings.NewBooleanType()
    fieldNameMap["failed"] = "Failed"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.hlm.community.check_info",fields, reflect.TypeOf(CommunityCheckInfo{}), fieldNameMap, validators)
}

func CommunityStatusBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["source"] = bindings.NewStringType()
    fieldNameMap["source"] = "Source"
    fields["destination"] = bindings.NewStringType()
    fieldNameMap["destination"] = "Destination"
    fields["latency"] = bindings.NewIntegerType()
    fieldNameMap["latency"] = "Latency"
    fields["failed"] = bindings.NewBooleanType()
    fieldNameMap["failed"] = "Failed"
    fields["errors"] = bindings.NewListType(bindings.NewReferenceType(std.LocalizableMessageBindingType), reflect.TypeOf([]std.LocalizableMessage{}))
    fieldNameMap["errors"] = "Errors"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.hlm.community.status",fields, reflect.TypeOf(CommunityStatus{}), fieldNameMap, validators)
}


