/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Code generated. DO NOT EDIT.

/*
 * Data type definitions file for service: Storage.
 * Includes binding types of a structures and enumerations defined in the service.
 * Shared by client-side stubs and server-side skeletons to ensure type
 * compatibility.
 */

package model

import (
	"net/url"

	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/services/vsphere/content/library"
)

// The ``Info`` class is the expanded form of FileInfo that includes details about the storage backing for a file in a library item.
type StorageInfo struct {
	// The storage backing on which this object resides. This might not be the same as the default storage backing associated with the library.
	StorageBacking library.StorageBacking
	// URIs that identify the file on the storage backing.
	//
	//  These URIs may be specific to the backing and may need interpretation by the client. A client that understands a URI scheme in this list may use that URI to directly access the file on the storage backing. This can provide high-performance support for file manipulation.
	StorageUris []url.URL
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
