/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Code generated. DO NOT EDIT.

/*
 * Data type definitions file for service: File.
 * Includes binding types of a structures and enumerations defined in the service.
 * Shared by client-side stubs and server-side skeletons to ensure type
 * compatibility.
 */

package item

import (
	"reflect"

	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/bindings"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/data"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/protocol"
	. "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/services/vsphere/content/library/item/model"
)

// The ``Info`` class provides information about a file in Content Library Service storage.
//
//  A file is an actual stored object for a library item. An item will have zero files initially, but one or more can be uploaded to the item.
type FileInfo struct {
	// A checksum for validating the content of the file.
	//
	//  This value can be used to verify that a transfer was completed without errors.
	ChecksumInfo *FileChecksumInfo
	// The name of the file.
	//
	//  This value will be unique within the library item for each file. It cannot be an empty string.
	Name string
	// The file size, in bytes. The file size is the storage used and not the uploaded or provisioned size. For example, when uploading a disk to a datastore, the amount of storage that the disk consumes may be different from the disk file size. When the file is not cached, the size is 0.
	Size int64
	// Indicates whether the file is on disk or not.
	Cached bool
	// The version of this file; incremented when a new copy of the file is uploaded.
	Version string
}

func fileListInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["download_session_id"] = bindings.NewIdType([]string{"com.vmware.content.library.item.DownloadSession"}, "")
	fieldNameMap["download_session_id"] = "DownloadSessionId"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func fileListOutputType() bindings.BindingType {
	return bindings.NewListType(bindings.NewReferenceType(FileInfoBindingType), reflect.TypeOf([]FileInfo{}))
}

func fileListRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	fields["download_session_id"] = bindings.NewIdType([]string{"com.vmware.content.library.item.DownloadSession"}, "")
	fieldNameMap["download_session_id"] = "DownloadSessionId"
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

func filePrepareInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["download_session_id"] = bindings.NewIdType([]string{"com.vmware.content.library.item.DownloadSession"}, "")
	fields["file_name"] = bindings.NewStringType()
	fields["endpoint_type"] = bindings.NewOptionalType(bindings.NewEnumType("com.vmware.content.library.item.downloadsession.file.endpoint_type", reflect.TypeOf(FileEndpointType(FileEndpointType_HTTPS))))
	fieldNameMap["download_session_id"] = "DownloadSessionId"
	fieldNameMap["file_name"] = "FileName"
	fieldNameMap["endpoint_type"] = "EndpointType"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func filePrepareOutputType() bindings.BindingType {
	return bindings.NewReferenceType(FileInfoBindingType)
}

func filePrepareRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	fields["download_session_id"] = bindings.NewIdType([]string{"com.vmware.content.library.item.DownloadSession"}, "")
	fields["file_name"] = bindings.NewStringType()
	fields["endpoint_type"] = bindings.NewOptionalType(bindings.NewEnumType("com.vmware.content.library.item.downloadsession.file.endpoint_type", reflect.TypeOf(FileEndpointType(FileEndpointType_HTTPS))))
	fieldNameMap["download_session_id"] = "DownloadSessionId"
	fieldNameMap["file_name"] = "FileName"
	fieldNameMap["endpoint_type"] = "EndpointType"
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
		map[string]int{"NotFound": 404, "InvalidArgument": 400, "Unauthorized": 403})
}

func fileGetInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["download_session_id"] = bindings.NewIdType([]string{"com.vmware.content.library.item.DownloadSession"}, "")
	fields["file_name"] = bindings.NewStringType()
	fieldNameMap["download_session_id"] = "DownloadSessionId"
	fieldNameMap["file_name"] = "FileName"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func fileGetOutputType() bindings.BindingType {
	return bindings.NewReferenceType(FileInfoBindingType)
}

func fileGetRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	fields["download_session_id"] = bindings.NewIdType([]string{"com.vmware.content.library.item.DownloadSession"}, "")
	fields["file_name"] = bindings.NewStringType()
	fieldNameMap["download_session_id"] = "DownloadSessionId"
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
		map[string]int{"NotFound": 404, "InvalidArgument": 400})
}

func FileInfoBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
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
	return bindings.NewStructType("com.vmware.content.library.item.file.info", fields, reflect.TypeOf(FileInfo{}), fieldNameMap, validators)
}
