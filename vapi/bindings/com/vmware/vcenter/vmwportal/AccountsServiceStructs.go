/* Copyright © 2019 VMware, Inc. All Rights Reserved.
     SPDX-License-Identifier: BSD-2-Clause */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Data type definitions file for service: Accounts.
 * Includes binding types of a structures and enumerations defined in the service.
 * Shared by client-side stubs and server-side skeletons to ensure type
 * compatibility.
 */

package vmwportal

import (
    "reflect"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/bindings"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/data"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/protocol"
)





// The ``CheckSpec`` class contains the product name for which we are checking the entitlement. **Warning:** This class is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
 type AccountsCheckSpec struct {
    // Product Name for which user wants to check the entitlement. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    ProductName string
}







// The ``Summary`` class contains the information about the entitlement account. **Warning:** This class is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
 type AccountsSummary struct {
    // Account identifier. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    Account string
    // Account Name. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    Name string
}









func accountsListInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func accountsListOutputType() bindings.BindingType {
    return bindings.NewListType(bindings.NewReferenceType(AccountsSummaryBindingType), reflect.TypeOf([]AccountsSummary{}))
}

func accountsListRestMetadata() protocol.OperationRestMetadata {
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
      "/vcenter/vmwportal/accounts",
       resultHeaders,
       200,
       errorHeaders,
       map[string]int{"ResourceInaccessible": 500,"Unauthenticated": 401,"Error": 500})
}


func accountsCheckInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["account"] = bindings.NewIdType([]string {"com.vmware.vcenter.vmwportal.Account"}, "")
    fields["product"] = bindings.NewReferenceType(AccountsCheckSpecBindingType)
    fieldNameMap["account"] = "Account"
    fieldNameMap["product"] = "Product"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func accountsCheckOutputType() bindings.BindingType {
    return bindings.NewVoidType()
}

func accountsCheckRestMetadata() protocol.OperationRestMetadata {
    paramsTypeMap := map[string]bindings.BindingType{}
    pathParams := map[string]string{}
    queryParams := map[string]string{}
    headerParams := map[string]string{}
    paramsTypeMap["product"] = bindings.NewReferenceType(AccountsCheckSpecBindingType)
    paramsTypeMap["account"] = bindings.NewIdType([]string {"com.vmware.vcenter.vmwportal.Account"}, "")
    paramsTypeMap["account"] = bindings.NewIdType([]string {"com.vmware.vcenter.vmwportal.Account"}, "")
    pathParams["account"] = "account"
    resultHeaders := map[string]string{}
    errorHeaders := map[string]string{}
    errorHeaders["Unauthenticated.challenge"] = "WWW-Authenticate"
    return protocol.NewOperationRestMetadata(
      paramsTypeMap,
      pathParams,
      queryParams,
      headerParams,
      "product",
      "POST",
      "/vcenter/vmwportal/accounts/{account}",
       resultHeaders,
       204,
       errorHeaders,
       map[string]int{"NotFound": 404,"Unauthenticated": 401,"ResourceInaccessible": 500,"Error": 500})
}



func AccountsCheckSpecBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["product_name"] = bindings.NewStringType()
    fieldNameMap["product_name"] = "ProductName"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.vmwportal.accounts.check_spec",fields, reflect.TypeOf(AccountsCheckSpec{}), fieldNameMap, validators)
}

func AccountsSummaryBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["account"] = bindings.NewIdType([]string {"com.vmware.vcenter.vmwportal.Account"}, "")
    fieldNameMap["account"] = "Account"
    fields["name"] = bindings.NewStringType()
    fieldNameMap["name"] = "Name"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.vmwportal.accounts.summary",fields, reflect.TypeOf(AccountsSummary{}), fieldNameMap, validators)
}


