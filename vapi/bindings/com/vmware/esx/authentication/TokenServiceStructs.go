/* Copyright © 2019 VMware, Inc. All Rights Reserved.
     SPDX-License-Identifier: BSD-2-Clause */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Data type definitions file for service: Token.
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




// The ``TokenInfo`` class contains API access token.
 type TokenTokenInfo struct {
    // The ``accessToken`` contains the token.
    AccessToken string
}



func (TokenTokenInfo TokenTokenInfo) Error() string {
    return "com.vmware.esx.authentication.token_info"
}






func tokenCreateInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func tokenCreateOutputType() bindings.BindingType {
    return bindings.NewReferenceType(TokenTokenInfoBindingType)
}

func tokenCreateRestMetadata() protocol.OperationRestMetadata {
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
      "POST",
      "/esx/authentication/token",
       resultHeaders,
       201,
       errorHeaders,
       map[string]int{"Unauthenticated": 401})
}



func TokenTokenInfoBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["access_token"] = bindings.NewSecretType()
    fieldNameMap["access_token"] = "AccessToken"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.esx.authentication.token.token_info",fields, reflect.TypeOf(TokenTokenInfo{}), fieldNameMap, validators)
}


