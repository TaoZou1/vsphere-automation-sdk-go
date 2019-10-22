/* Copyright © 2019 VMware, Inc. All Rights Reserved.
     SPDX-License-Identifier: BSD-2-Clause */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Data type definitions file for service: AssociatedProducts.
 * Includes binding types of a structures and enumerations defined in the service.
 * Shared by client-side stubs and server-side skeletons to ensure type
 * compatibility.
 */

package discovery

import (
    "reflect"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/bindings"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/data"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/protocol"
)

// The resource type for the products interface.
const AssociatedProducts_RESOURCE_TYPE = "com.vmware.vcenter.lcm.product"



// The ``CreateSpec`` class is the specification used for the product creation.
 type AssociatedProductsCreateSpec struct {
    // The name of the product.
    ProductName string
    // Current product version.
    Version string
    // The list of hostname/IPs of the instances of the VMware products deployed in the environment.
    Deployments []string
}


// TODO some error doesn't have ErrorType. Resolve this later VAPI-3009
//


// The ``UpdateSpec`` class is the specification for the product update.
 type AssociatedProductsUpdateSpec struct {
    // The list of hostname/IPs of the instances of the VMware products deployed in the environment.
    Deployments []string
}


// TODO some error doesn't have ErrorType. Resolve this later VAPI-3009
//





func associatedProductsListInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func associatedProductsListOutputType() bindings.BindingType {
    return bindings.NewListType(bindings.NewReferenceType(ProductBindingType), reflect.TypeOf([]Product{}))
}

func associatedProductsListRestMetadata() protocol.OperationRestMetadata {
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
      "/vcenter/lcm/discovery/associated-products",
       resultHeaders,
       200,
       errorHeaders,
       map[string]int{"Unauthenticated": 401,"Error": 500})
}


func associatedProductsCreateInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["spec"] = bindings.NewReferenceType(AssociatedProductsCreateSpecBindingType)
    fieldNameMap["spec"] = "Spec"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func associatedProductsCreateOutputType() bindings.BindingType {
    return bindings.NewIdType([]string {"com.vmware.vcenter.lcm.product"}, "")
}

func associatedProductsCreateRestMetadata() protocol.OperationRestMetadata {
    paramsTypeMap := map[string]bindings.BindingType{}
    pathParams := map[string]string{}
    queryParams := map[string]string{}
    headerParams := map[string]string{}
    paramsTypeMap["spec"] = bindings.NewReferenceType(AssociatedProductsCreateSpecBindingType)
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
      "/vcenter/lcm/discovery/associated-products",
       resultHeaders,
       201,
       errorHeaders,
       map[string]int{"Unauthenticated": 401,"AlreadyExists": 400,"InvalidArgument": 400,"Error": 500})
}


func associatedProductsGetInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["product"] = bindings.NewIdType([]string {"com.vmware.vcenter.lcm.product"}, "")
    fieldNameMap["product"] = "Product"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func associatedProductsGetOutputType() bindings.BindingType {
    return bindings.NewReferenceType(ProductBindingType)
}

func associatedProductsGetRestMetadata() protocol.OperationRestMetadata {
    paramsTypeMap := map[string]bindings.BindingType{}
    pathParams := map[string]string{}
    queryParams := map[string]string{}
    headerParams := map[string]string{}
    paramsTypeMap["product"] = bindings.NewIdType([]string {"com.vmware.vcenter.lcm.product"}, "")
    paramsTypeMap["product"] = bindings.NewIdType([]string {"com.vmware.vcenter.lcm.product"}, "")
    pathParams["product"] = "product"
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
      "/vcenter/lcm/discovery/associated-products/{product}",
       resultHeaders,
       200,
       errorHeaders,
       map[string]int{"Unauthenticated": 401,"NotFound": 404,"Error": 500})
}


func associatedProductsUpdateInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["product"] = bindings.NewIdType([]string {"com.vmware.vcenter.lcm.product"}, "")
    fields["spec"] = bindings.NewReferenceType(AssociatedProductsUpdateSpecBindingType)
    fieldNameMap["product"] = "Product"
    fieldNameMap["spec"] = "Spec"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func associatedProductsUpdateOutputType() bindings.BindingType {
    return bindings.NewVoidType()
}

func associatedProductsUpdateRestMetadata() protocol.OperationRestMetadata {
    paramsTypeMap := map[string]bindings.BindingType{}
    pathParams := map[string]string{}
    queryParams := map[string]string{}
    headerParams := map[string]string{}
    paramsTypeMap["product"] = bindings.NewIdType([]string {"com.vmware.vcenter.lcm.product"}, "")
    paramsTypeMap["spec"] = bindings.NewReferenceType(AssociatedProductsUpdateSpecBindingType)
    paramsTypeMap["product"] = bindings.NewIdType([]string {"com.vmware.vcenter.lcm.product"}, "")
    pathParams["product"] = "product"
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
      "/vcenter/lcm/discovery/associated-products/{product}",
       resultHeaders,
       204,
       errorHeaders,
       map[string]int{"Unauthenticated": 401,"InvalidArgument": 400,"NotFound": 404,"Error": 500})
}


func associatedProductsDeleteInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["product"] = bindings.NewIdType([]string {"com.vmware.vcenter.lcm.product"}, "")
    fieldNameMap["product"] = "Product"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func associatedProductsDeleteOutputType() bindings.BindingType {
    return bindings.NewVoidType()
}

func associatedProductsDeleteRestMetadata() protocol.OperationRestMetadata {
    paramsTypeMap := map[string]bindings.BindingType{}
    pathParams := map[string]string{}
    queryParams := map[string]string{}
    headerParams := map[string]string{}
    paramsTypeMap["product"] = bindings.NewIdType([]string {"com.vmware.vcenter.lcm.product"}, "")
    paramsTypeMap["product"] = bindings.NewIdType([]string {"com.vmware.vcenter.lcm.product"}, "")
    pathParams["product"] = "product"
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
      "/vcenter/lcm/discovery/associated-products/{product}",
       resultHeaders,
       204,
       errorHeaders,
       map[string]int{"Unauthenticated": 401,"NotFound": 404,"Error": 500})
}



func AssociatedProductsCreateSpecBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["product_name"] = bindings.NewStringType()
    fieldNameMap["product_name"] = "ProductName"
    fields["version"] = bindings.NewStringType()
    fieldNameMap["version"] = "Version"
    fields["deployments"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewStringType(), reflect.TypeOf([]string{})))
    fieldNameMap["deployments"] = "Deployments"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.lcm.discovery.associated_products.create_spec",fields, reflect.TypeOf(AssociatedProductsCreateSpec{}), fieldNameMap, validators)
}

func AssociatedProductsUpdateSpecBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["deployments"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewStringType(), reflect.TypeOf([]string{})))
    fieldNameMap["deployments"] = "Deployments"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.lcm.discovery.associated_products.update_spec",fields, reflect.TypeOf(AssociatedProductsUpdateSpec{}), fieldNameMap, validators)
}


