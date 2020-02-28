/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Code generated. DO NOT EDIT.

/*
 * Data type definitions file for service: File.
 * Includes binding types of a structures and enumerations defined in the service.
 * Shared by client-side stubs and server-side skeletons to ensure type
 * compatibility.
 */

package model

import (
	"reflect"

	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/bindings"
)

// The ``ChecksumAlgorithm`` enumeration class defines the valid checksum algorithms.
//
// <p> See {@link com.vmware.vapi.bindings.ApiEnumeration enumerated types description}.
type FileChecksumAlgorithm string

const (
	// Checksum algorithm: SHA-1
	FileChecksumAlgorithm_SHA1 FileChecksumAlgorithm = "SHA1"
	// Checksum algorithm: MD5
	FileChecksumAlgorithm_MD5 FileChecksumAlgorithm = "MD5"
	// Checksum algorithm: SHA-256. This constant field was added in vSphere API 6.8.
	FileChecksumAlgorithm_SHA256 FileChecksumAlgorithm = "SHA256"
	// Checksum algorithm: SHA-512. This constant field was added in vSphere API 6.8.
	FileChecksumAlgorithm_SHA512 FileChecksumAlgorithm = "SHA512"
)

func (c FileChecksumAlgorithm) FileChecksumAlgorithm() bool {
	switch c {
	case FileChecksumAlgorithm_SHA1:
		return true
	case FileChecksumAlgorithm_MD5:
		return true
	case FileChecksumAlgorithm_SHA256:
		return true
	case FileChecksumAlgorithm_SHA512:
		return true
	default:
		return false
	}
}

// Provides checksums for a FileInfo object.
type FileChecksumInfo struct {
	// The checksum algorithm (SHA1, MD5, SHA256, SHA512) used to calculate the checksum.
	Algorithm *FileChecksumAlgorithm
	// The checksum value calculated with FileChecksumInfo#algorithm.
	Checksum string
}

func FileChecksumInfoBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["algorithm"] = bindings.NewOptionalType(bindings.NewEnumType("com.vmware.content.library.item.file.checksum_algorithm", reflect.TypeOf(FileChecksumAlgorithm(FileChecksumAlgorithm_SHA1))))
	fieldNameMap["algorithm"] = "Algorithm"
	fields["checksum"] = bindings.NewStringType()
	fieldNameMap["checksum"] = "Checksum"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.content.library.item.file.checksum_info", fields, reflect.TypeOf(FileChecksumInfo{}), fieldNameMap, validators)
}

// The ``EndpointType`` enumeration class defines the types of endpoints used to download the file.
//
// <p> See {@link com.vmware.vapi.bindings.ApiEnumeration enumerated types description}.
type FileEndpointType string

const (
	// An https download endpoint.
	FileEndpointType_HTTPS FileEndpointType = "HTTPS"
	// A direct download endpoint indicating the location of the file on storage. The caller is responsible for retrieving the file from the storage location directly.
	FileEndpointType_DIRECT FileEndpointType = "DIRECT"
)

func (e FileEndpointType) FileEndpointType() bool {
	switch e {
	case FileEndpointType_HTTPS:
		return true
	case FileEndpointType_DIRECT:
		return true
	default:
		return false
	}
}
