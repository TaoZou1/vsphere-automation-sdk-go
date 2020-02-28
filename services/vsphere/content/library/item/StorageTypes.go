/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Code generated. DO NOT EDIT.

/*
 * Data type definitions file for service: Storage.
 * Includes binding types of a structures and enumerations defined in the service.
 * Shared by client-side stubs and server-side skeletons to ensure type
 * compatibility.
 */

package item

import (
	"net/url"
	"reflect"

	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/bindings"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/data"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/protocol"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/services/vsphere/content/library"
	. "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/services/vsphere/content/library/item/model"
)

func storageGetInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["library_item_id"] = bindings.NewIdType([]string{"com.vmware.content.library.Item"}, "")
	fields["file_name"] = bindings.NewStringType()
	fieldNameMap["library_item_id"] = "LibraryItemId"
	fieldNameMap["file_name"] = "FileName"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func storageGetOutputType() bindings.BindingType {
	return bindings.NewListType(bindings.NewReferenceType(StorageInfoBindingType), reflect.TypeOf([]StorageInfo{}))
}

func storageGetRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	fields["library_item_id"] = bindings.NewIdType([]string{"com.vmware.content.library.Item"}, "")
	fields["file_name"] = bindings.NewStringType()
	fieldNameMap["library_item_id"] = "LibraryItemId"
	fieldNameMap["file_name"] = "FileName"
	resultHeaders := map[string]string{}
	errorHeaders := map[string]string{}
	return protocol.NewOperationRestMetadata(
		fields,
		fieldNameMap,
		paramsTypeMap,
		pathParams,
		queryParams,
		headerParams,
		"",
		"",
		"null",
		"",
		resultHeaders,
		0,
		errorHeaders,
		map[string]int{"NotFound": 404})
}

func storageListInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["library_item_id"] = bindings.NewIdType([]string{"com.vmware.content.library.Item"}, "")
	fieldNameMap["library_item_id"] = "LibraryItemId"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func storageListOutputType() bindings.BindingType {
	return bindings.NewListType(bindings.NewReferenceType(StorageInfoBindingType), reflect.TypeOf([]StorageInfo{}))
}

func storageListRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	fields["library_item_id"] = bindings.NewIdType([]string{"com.vmware.content.library.Item"}, "")
	fieldNameMap["library_item_id"] = "LibraryItemId"
	resultHeaders := map[string]string{}
	errorHeaders := map[string]string{}
	return protocol.NewOperationRestMetadata(
		fields,
		fieldNameMap,
		paramsTypeMap,
		pathParams,
		queryParams,
		headerParams,
		"",
		"",
		"null",
		"",
		resultHeaders,
		0,
		errorHeaders,
		map[string]int{"NotFound": 404})
}

func StorageInfoBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["storage_backing"] = bindings.NewReferenceType(library.StorageBackingBindingType)
	fieldNameMap["storage_backing"] = "StorageBacking"
	fields["storage_uris"] = bindings.NewListType(bindings.NewUriType(), reflect.TypeOf([]url.URL{}))
	fieldNameMap["storage_uris"] = "StorageUris"
	fields["checksum_info"] = bindings.NewOptionalType(bindings.NewReferenceType(FileChecksumInfoBindingType))
	fieldNameMap["checksum_info"] = "ChecksumInfo"
	fields["name"] = bindings.NewStringType()
	fieldNameMap["name"] = "Name"
	fields["size"] = bindings.NewIntegerType()
	fieldNameMap["size"] = "Size"
	fields["cached"] = bindings.NewBooleanType()
	fieldNameMap["cached"] = "Cached"
	fields["version"] = bindings.NewStringType()
	fieldNameMap["version"] = "Version"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.content.library.item.storage.info", fields, reflect.TypeOf(StorageInfo{}), fieldNameMap, validators)
}
