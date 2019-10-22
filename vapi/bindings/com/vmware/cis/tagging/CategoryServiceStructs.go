/* Copyright © 2019 VMware, Inc. All Rights Reserved.
     SPDX-License-Identifier: BSD-2-Clause */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Data type definitions file for service: Category.
 * Includes binding types of a structures and enumerations defined in the service.
 * Shared by client-side stubs and server-side skeletons to ensure type
 * compatibility.
 */

package tagging

import (
    "reflect"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/bindings"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/data"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/protocol"
)




// The ``CreateSpec`` class is used to create a category. 
//
//  Use the Category#create method to create a category defined by the create specification.
 type CategoryCreateSpec struct {
    // The display name of the category.
    Name string
    // The description of the category.
    Description string
    // The associated cardinality (SINGLE, MULTIPLE) of the category.
    Cardinality CategoryModel_Cardinality
    // Object types to which this category's tags can be attached.
    AssociableTypes map[string]bool
    //
    CategoryId *string
}


// TODO some error doesn't have ErrorType. Resolve this later VAPI-3009
//


// The ``UpdateSpec`` class describes the updates to be made to an existing category. 
//
//  Use the Category#update method to modify a category. When you call the method, specify the category identifier. You obtain the category identifier when you call the Category#create method. You can also retrieve an identifier by using the Category#list method.
 type CategoryUpdateSpec struct {
    // The display name of the category.
    Name *string
    // The description of the category.
    Description *string
    // The associated cardinality (SINGLE, MULTIPLE) of the category.
    Cardinality *CategoryModel_Cardinality
    // Object types to which this category's tags can be attached. 
//
//  The map with bool value of associable types cannot be updated incrementally. For example, if CategoryUpdateSpec#associableTypes originally contains {A,B,C} and you want to add D, then you need to pass {A,B,C,D} in your update specification. You also cannot remove any item from this map with bool value. For example, if you have {A,B,C}, then you cannot remove say {A} from it. Similarly, if you start with an empty map with bool value, then that implies that you can tag any object and hence you cannot later pass say {A}, because that would be restricting the type of objects you want to tag. Thus, associable types can only grow and not shrink.
    AssociableTypes map[string]bool
}


// TODO some error doesn't have ErrorType. Resolve this later VAPI-3009
//





func categoryCreateInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["create_spec"] = bindings.NewReferenceType(CategoryCreateSpecBindingType)
    fieldNameMap["create_spec"] = "CreateSpec"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func categoryCreateOutputType() bindings.BindingType {
    return bindings.NewIdType([]string {"com.vmware.cis.tagging.Category"}, "")
}

func categoryCreateRestMetadata() protocol.OperationRestMetadata {
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
       map[string]int{"AlreadyExists": 400,"InvalidArgument": 400,"Unauthorized": 403})
}


func categoryGetInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["category_id"] = bindings.NewIdType([]string {"com.vmware.cis.tagging.Category"}, "")
    fieldNameMap["category_id"] = "CategoryId"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func categoryGetOutputType() bindings.BindingType {
    return bindings.NewReferenceType(CategoryModelBindingType)
}

func categoryGetRestMetadata() protocol.OperationRestMetadata {
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
       map[string]int{"NotFound": 404,"Unauthorized": 403})
}


func categoryUpdateInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["category_id"] = bindings.NewIdType([]string {"com.vmware.cis.tagging.Category"}, "")
    fields["update_spec"] = bindings.NewReferenceType(CategoryUpdateSpecBindingType)
    fieldNameMap["category_id"] = "CategoryId"
    fieldNameMap["update_spec"] = "UpdateSpec"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func categoryUpdateOutputType() bindings.BindingType {
    return bindings.NewVoidType()
}

func categoryUpdateRestMetadata() protocol.OperationRestMetadata {
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
       map[string]int{"AlreadyExists": 400,"InvalidArgument": 400,"NotFound": 404,"Unauthorized": 403})
}


func categoryDeleteInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["category_id"] = bindings.NewIdType([]string {"com.vmware.cis.tagging.Category"}, "")
    fieldNameMap["category_id"] = "CategoryId"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func categoryDeleteOutputType() bindings.BindingType {
    return bindings.NewVoidType()
}

func categoryDeleteRestMetadata() protocol.OperationRestMetadata {
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
       map[string]int{"NotFound": 404,"Unauthorized": 403,"Unauthenticated": 401})
}


func categoryListInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func categoryListOutputType() bindings.BindingType {
    return bindings.NewListType(bindings.NewIdType([]string {"com.vmware.cis.tagging.Category"}, ""), reflect.TypeOf([]string{}))
}

func categoryListRestMetadata() protocol.OperationRestMetadata {
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
       map[string]int{})
}


func categoryListUsedCategoriesInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["used_by_entity"] = bindings.NewStringType()
    fieldNameMap["used_by_entity"] = "UsedByEntity"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func categoryListUsedCategoriesOutputType() bindings.BindingType {
    return bindings.NewListType(bindings.NewIdType([]string {"com.vmware.cis.tagging.Category"}, ""), reflect.TypeOf([]string{}))
}

func categoryListUsedCategoriesRestMetadata() protocol.OperationRestMetadata {
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
       map[string]int{})
}


func categoryAddToUsedByInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["category_id"] = bindings.NewIdType([]string {"com.vmware.cis.tagging.Category"}, "")
    fields["used_by_entity"] = bindings.NewStringType()
    fieldNameMap["category_id"] = "CategoryId"
    fieldNameMap["used_by_entity"] = "UsedByEntity"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func categoryAddToUsedByOutputType() bindings.BindingType {
    return bindings.NewVoidType()
}

func categoryAddToUsedByRestMetadata() protocol.OperationRestMetadata {
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
       map[string]int{"NotFound": 404,"Unauthorized": 403})
}


func categoryRemoveFromUsedByInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["category_id"] = bindings.NewIdType([]string {"com.vmware.cis.tagging.Category"}, "")
    fields["used_by_entity"] = bindings.NewStringType()
    fieldNameMap["category_id"] = "CategoryId"
    fieldNameMap["used_by_entity"] = "UsedByEntity"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func categoryRemoveFromUsedByOutputType() bindings.BindingType {
    return bindings.NewVoidType()
}

func categoryRemoveFromUsedByRestMetadata() protocol.OperationRestMetadata {
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
       map[string]int{"NotFound": 404,"Unauthorized": 403})
}


func categoryRevokePropagatingPermissionsInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["category_id"] = bindings.NewIdType([]string {"com.vmware.cis.tagging.Category"}, "")
    fieldNameMap["category_id"] = "CategoryId"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func categoryRevokePropagatingPermissionsOutputType() bindings.BindingType {
    return bindings.NewVoidType()
}

func categoryRevokePropagatingPermissionsRestMetadata() protocol.OperationRestMetadata {
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
       map[string]int{"NotFound": 404,"Unauthorized": 403})
}



func CategoryCreateSpecBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["name"] = bindings.NewStringType()
    fieldNameMap["name"] = "Name"
    fields["description"] = bindings.NewStringType()
    fieldNameMap["description"] = "Description"
    fields["cardinality"] = bindings.NewEnumType("com.vmware.cis.tagging.category_model.cardinality", reflect.TypeOf(CategoryModel_Cardinality(CategoryModel_Cardinality_SINGLE)))
    fieldNameMap["cardinality"] = "Cardinality"
    fields["associable_types"] = bindings.NewSetType(bindings.NewStringType(), reflect.TypeOf(map[string]bool{}))
    fieldNameMap["associable_types"] = "AssociableTypes"
    fields["category_id"] = bindings.NewOptionalType(bindings.NewIdType([]string {"com.vmware.cis.tagging.Category"}, ""))
    fieldNameMap["category_id"] = "CategoryId"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.cis.tagging.category.create_spec",fields, reflect.TypeOf(CategoryCreateSpec{}), fieldNameMap, validators)
}

func CategoryUpdateSpecBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["name"] = bindings.NewOptionalType(bindings.NewStringType())
    fieldNameMap["name"] = "Name"
    fields["description"] = bindings.NewOptionalType(bindings.NewStringType())
    fieldNameMap["description"] = "Description"
    fields["cardinality"] = bindings.NewOptionalType(bindings.NewEnumType("com.vmware.cis.tagging.category_model.cardinality", reflect.TypeOf(CategoryModel_Cardinality(CategoryModel_Cardinality_SINGLE))))
    fieldNameMap["cardinality"] = "Cardinality"
    fields["associable_types"] = bindings.NewOptionalType(bindings.NewSetType(bindings.NewStringType(), reflect.TypeOf(map[string]bool{})))
    fieldNameMap["associable_types"] = "AssociableTypes"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.cis.tagging.category.update_spec",fields, reflect.TypeOf(CategoryUpdateSpec{}), fieldNameMap, validators)
}


