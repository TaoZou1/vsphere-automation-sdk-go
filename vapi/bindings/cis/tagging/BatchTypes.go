/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Code generated. DO NOT EDIT.

/*
 * Data type definitions file for service: Batch.
 * Includes binding types of a structures and enumerations defined in the service.
 * Shared by client-side stubs and server-side skeletons to ensure type
 * compatibility.
 */

package tagging

import (
	"reflect"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/bindings/vapi/std"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/bindings"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/data"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/protocol"
)



// The ``TagToObjects`` class describes a tag and its related objects. Use the Batch#listAttachedObjectsOnTags method or the Batch#listAllAttachedObjectsOnTags method to retrieve a array with each element containing a tag and objects its attached to.
type BatchTagToObjects struct {
    // The identifier of the tag.
	TagId string
    // The identifiers of the related objects.
	ObjectIds []std.DynamicID
}

// The ``ObjectToTags`` class describes an object and its related tags. Use the Batch#listAttachedTagsOnObjects method to retrieve a array with each element containing an object and tags attached to it.
type BatchObjectToTags struct {
    // The identifier of the object.
	ObjectId std.DynamicID
    // The identifiers of the related tags.
	TagIds []string
}



func batchGetCategoriesInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["category_ids"] = bindings.NewListType(bindings.NewIdType([]string{"com.vmware.cis.tagging.Category"}, ""), reflect.TypeOf([]string{}))
	fieldNameMap["category_ids"] = "CategoryIds"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func batchGetCategoriesOutputType() bindings.BindingType {
	return bindings.NewListType(bindings.NewReferenceType(CategoryModelBindingType), reflect.TypeOf([]CategoryModel{}))
}

func batchGetCategoriesRestMetadata() protocol.OperationRestMetadata {
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

func batchGetAllCategoriesInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func batchGetAllCategoriesOutputType() bindings.BindingType {
	return bindings.NewListType(bindings.NewReferenceType(CategoryModelBindingType), reflect.TypeOf([]CategoryModel{}))
}

func batchGetAllCategoriesRestMetadata() protocol.OperationRestMetadata {
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

func batchGetTagsInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["tag_ids"] = bindings.NewListType(bindings.NewIdType([]string{"com.vmware.cis.tagging.Tag"}, ""), reflect.TypeOf([]string{}))
	fieldNameMap["tag_ids"] = "TagIds"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func batchGetTagsOutputType() bindings.BindingType {
	return bindings.NewListType(bindings.NewReferenceType(TagModelBindingType), reflect.TypeOf([]TagModel{}))
}

func batchGetTagsRestMetadata() protocol.OperationRestMetadata {
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

func batchGetAllTagsInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func batchGetAllTagsOutputType() bindings.BindingType {
	return bindings.NewListType(bindings.NewReferenceType(TagModelBindingType), reflect.TypeOf([]TagModel{}))
}

func batchGetAllTagsRestMetadata() protocol.OperationRestMetadata {
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

func batchListTagsForCategoriesInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["category_ids"] = bindings.NewListType(bindings.NewIdType([]string{"com.vmware.cis.tagging.Category"}, ""), reflect.TypeOf([]string{}))
	fieldNameMap["category_ids"] = "CategoryIds"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func batchListTagsForCategoriesOutputType() bindings.BindingType {
	return bindings.NewListType(bindings.NewIdType([]string{"com.vmware.cis.tagging.Tag"}, ""), reflect.TypeOf([]string{}))
}

func batchListTagsForCategoriesRestMetadata() protocol.OperationRestMetadata {
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

func batchFindTagsByNameInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["tag_name"] = bindings.NewStringType()
	fieldNameMap["tag_name"] = "TagName"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func batchFindTagsByNameOutputType() bindings.BindingType {
	return bindings.NewListType(bindings.NewIdType([]string{"com.vmware.cis.tagging.Tag"}, ""), reflect.TypeOf([]string{}))
}

func batchFindTagsByNameRestMetadata() protocol.OperationRestMetadata {
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

func batchListAttachedObjectsInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["tag_ids"] = bindings.NewListType(bindings.NewIdType([]string{"com.vmware.cis.tagging.Tag"}, ""), reflect.TypeOf([]string{}))
	fieldNameMap["tag_ids"] = "TagIds"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func batchListAttachedObjectsOutputType() bindings.BindingType {
	return bindings.NewListType(bindings.NewReferenceType(std.DynamicIDBindingType), reflect.TypeOf([]std.DynamicID{}))
}

func batchListAttachedObjectsRestMetadata() protocol.OperationRestMetadata {
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
		map[string]int{"Unauthenticated": 401})
}

func batchListAttachedObjectsOnTagsInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["tag_ids"] = bindings.NewListType(bindings.NewIdType([]string{"com.vmware.cis.tagging.Tag"}, ""), reflect.TypeOf([]string{}))
	fieldNameMap["tag_ids"] = "TagIds"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func batchListAttachedObjectsOnTagsOutputType() bindings.BindingType {
	return bindings.NewListType(bindings.NewReferenceType(BatchTagToObjectsBindingType), reflect.TypeOf([]BatchTagToObjects{}))
}

func batchListAttachedObjectsOnTagsRestMetadata() protocol.OperationRestMetadata {
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
		map[string]int{"Unauthenticated": 401})
}

func batchListAllAttachedObjectsOnTagsInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func batchListAllAttachedObjectsOnTagsOutputType() bindings.BindingType {
	return bindings.NewListType(bindings.NewReferenceType(BatchTagToObjectsBindingType), reflect.TypeOf([]BatchTagToObjects{}))
}

func batchListAllAttachedObjectsOnTagsRestMetadata() protocol.OperationRestMetadata {
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
		map[string]int{"Unauthenticated": 401})
}

func batchListAttachedTagsInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["object_ids"] = bindings.NewListType(bindings.NewReferenceType(std.DynamicIDBindingType), reflect.TypeOf([]std.DynamicID{}))
	fieldNameMap["object_ids"] = "ObjectIds"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func batchListAttachedTagsOutputType() bindings.BindingType {
	return bindings.NewListType(bindings.NewIdType([]string{"com.vmware.cis.tagging.Tag"}, ""), reflect.TypeOf([]string{}))
}

func batchListAttachedTagsRestMetadata() protocol.OperationRestMetadata {
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
		map[string]int{"Unauthenticated": 401})
}

func batchListAttachedTagsOnObjectsInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["object_ids"] = bindings.NewListType(bindings.NewReferenceType(std.DynamicIDBindingType), reflect.TypeOf([]std.DynamicID{}))
	fieldNameMap["object_ids"] = "ObjectIds"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func batchListAttachedTagsOnObjectsOutputType() bindings.BindingType {
	return bindings.NewListType(bindings.NewReferenceType(BatchObjectToTagsBindingType), reflect.TypeOf([]BatchObjectToTags{}))
}

func batchListAttachedTagsOnObjectsRestMetadata() protocol.OperationRestMetadata {
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
		map[string]int{"Unauthenticated": 401})
}


func BatchTagToObjectsBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["tag_id"] = bindings.NewIdType([]string{"com.vmware.cis.tagging.Tag"}, "")
	fieldNameMap["tag_id"] = "TagId"
	fields["object_ids"] = bindings.NewListType(bindings.NewReferenceType(std.DynamicIDBindingType), reflect.TypeOf([]std.DynamicID{}))
	fieldNameMap["object_ids"] = "ObjectIds"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.cis.tagging.batch.tag_to_objects", fields, reflect.TypeOf(BatchTagToObjects{}), fieldNameMap, validators)
}

func BatchObjectToTagsBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["object_id"] = bindings.NewReferenceType(std.DynamicIDBindingType)
	fieldNameMap["object_id"] = "ObjectId"
	fields["tag_ids"] = bindings.NewListType(bindings.NewIdType([]string{"com.vmware.cis.tagging.Tag"}, ""), reflect.TypeOf([]string{}))
	fieldNameMap["tag_ids"] = "TagIds"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.cis.tagging.batch.object_to_tags", fields, reflect.TypeOf(BatchObjectToTags{}), fieldNameMap, validators)
}

