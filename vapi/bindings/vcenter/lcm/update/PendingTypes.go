/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Code generated. DO NOT EDIT.

/*
 * Data type definitions file for service: Pending.
 * Includes binding types of a structures and enumerations defined in the service.
 * Shared by client-side stubs and server-side skeletons to ensure type
 * compatibility.
 */

package update

import (
	"reflect"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/bindings"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/data"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/protocol"
	"net/url"
	"time"
)

// Resource type for pending update
const Pending_RESOURCE_TYPE = "com.vmware.vcenter.lcm.update.pending"


// Level of severity for applying a given patch or update.
//
// <p> See {@link com.vmware.vapi.bindings.ApiEnumeration enumerated types description}.
type PendingSeverityType string

const (
    // Vulnerabilities that can be exploited by an unauthenticated attacker from the Internet or those that break the guest/host Operating System isolation.
	PendingSeverityType_CRITICAL PendingSeverityType = "CRITICAL"
    // Vulnerabilities that are not rated critical but whose exploitation results in the complete compromise of confidentiality and/or integrity of user data and/or processing resources through user assistance or by authenticated attackers.
	PendingSeverityType_IMPORTANT PendingSeverityType = "IMPORTANT"
    // Vulnerabilities where the ability to exploit is mitigated to a significant degree by configuration or difficulty of exploitation, but in certain deployment scenarios could still lead to the compromise of confidentiality, integrity, or availability of user data and/or processing resources.
	PendingSeverityType_MODERATE PendingSeverityType = "MODERATE"
    // All other issues that may or maynot have a security impact. Vulnerabilities where exploitation is believed to be extremely difficult, or where successful exploitation would have minimal impact.
	PendingSeverityType_LOW PendingSeverityType = "LOW"
)

func (s PendingSeverityType) PendingSeverityType() bool {
	switch s {
	case PendingSeverityType_CRITICAL:
		return true
	case PendingSeverityType_IMPORTANT:
		return true
	case PendingSeverityType_MODERATE:
		return true
	case PendingSeverityType_LOW:
		return true
	default:
		return false
	}
}


// The ``Category`` enumeration class defines the type of payload this release has on top of previous release
//
// <p> See {@link com.vmware.vapi.bindings.ApiEnumeration enumerated types description}.
type PendingCategory string

const (
    // Fixes vulnerabilities, doesn't change functionality
	PendingCategory_SECURITY PendingCategory = "SECURITY"
    // Fixes bugs/vulnerabilities, doesn't change functionality
	PendingCategory_FIX PendingCategory = "FIX"
    // Changes product functionality
	PendingCategory_UPDATE PendingCategory = "UPDATE"
    // Introduces new features, significantly changes product functionality
	PendingCategory_UPGRADE PendingCategory = "UPGRADE"
)

func (c PendingCategory) PendingCategory() bool {
	switch c {
	case PendingCategory_SECURITY:
		return true
	case PendingCategory_FIX:
		return true
	case PendingCategory_UPDATE:
		return true
	case PendingCategory_UPGRADE:
		return true
	default:
		return false
	}
}


// The ``UpdateType`` enumeration class defines update type
//
// <p> See {@link com.vmware.vapi.bindings.ApiEnumeration enumerated types description}.
type PendingUpdateType string

const (
    // Fixes bugs/vulnerabilities, doesn't change functionality
	PendingUpdateType_PATCH PendingUpdateType = "PATCH"
    // Changes product functionality
	PendingUpdateType_UPDATE PendingUpdateType = "UPDATE"
    // Introduces new features, significantly changes product functionality
	PendingUpdateType_UPGRADE PendingUpdateType = "UPGRADE"
)

func (u PendingUpdateType) PendingUpdateType() bool {
	switch u {
	case PendingUpdateType_PATCH:
		return true
	case PendingUpdateType_UPDATE:
		return true
	case PendingUpdateType_UPGRADE:
		return true
	default:
		return false
	}
}


// The ``Summary`` class contains basic information about the vCenter patch/update/upgrade
type PendingSummary struct {
    // Identifier of the given vSphere update
	PendingUpdate string
    // Version of the vSphere update or patch
	Version string
    // Release date of the vSphere update or patch
	ReleaseDate time.Time
    // Severity of the issues fixed in the vSphere update or patch
	Severity PendingSeverityType
    // Build number of the vCenter Release
	Build string
    // Type of the Release based on the current vCenter version
	UpdateType PendingUpdateType
    // Category of the release based on features bundled on top of previous release
	Category PendingCategory
    // Flag to suggest a reboot after the release is applied
	RebootRequired bool
    // VAMI or ISO URL for update or upgrade execute phase redirection
	ExecuteURL url.URL
    // List of URI pointing to patch or update release notes
	ReleaseNotes []url.URL
}

// The ``ListResult`` class contains information about the pending patch/updates for the given vCenter server.
type PendingListResult struct {
    // Time when the software depo was last checked.
	LastCheckTime time.Time
    // Number of pending updates
	UpdateCount *int64
    // Number of pending upgrades
	UpgradeCount *int64
    // List of pending update details
	Updates []PendingSummary
}

// The ``Info`` class contains detailed information about the vCenter patch/update.
type PendingInfo struct {
    // Description of the vSphere update
	Description string
    // Identifier of the given vSphere update
	PendingUpdate string
    // Version of the vSphere update or patch
	Version string
    // Release date of the vSphere update or patch
	ReleaseDate time.Time
    // Severity of the issues fixed in the vSphere update or patch
	Severity PendingSeverityType
    // Build number of the vCenter Release
	Build string
    // Type of the Release based on the current vCenter version
	UpdateType PendingUpdateType
    // Category of the release based on features bundled on top of previous release
	Category PendingCategory
    // Flag to suggest a reboot after the release is applied
	RebootRequired bool
    // VAMI or ISO URL for update or upgrade execute phase redirection
	ExecuteURL url.URL
    // List of URI pointing to patch or update release notes
	ReleaseNotes []url.URL
}



func pendingListInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func pendingListOutputType() bindings.BindingType {
	return bindings.NewReferenceType(PendingListResultBindingType)
}

func pendingListRestMetadata() protocol.OperationRestMetadata {
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
		"/vcenter/lcm/update/pending",
		resultHeaders,
		200,
		errorHeaders,
		map[string]int{"Unauthenticated": 401,"Error": 500})
}

func pendingGetInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["version"] = bindings.NewIdType([]string{"com.vmware.vcenter.lcm.update.pending"}, "")
	fieldNameMap["version"] = "Version"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func pendingGetOutputType() bindings.BindingType {
	return bindings.NewReferenceType(PendingInfoBindingType)
}

func pendingGetRestMetadata() protocol.OperationRestMetadata {
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	paramsTypeMap["version"] = bindings.NewIdType([]string{"com.vmware.vcenter.lcm.update.pending"}, "")
	paramsTypeMap["version"] = bindings.NewIdType([]string{"com.vmware.vcenter.lcm.update.pending"}, "")
	pathParams["version"] = "version"
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
		"/vcenter/lcm/update/pending/{version}",
		resultHeaders,
		200,
		errorHeaders,
		map[string]int{"Unauthenticated": 401,"NotFound": 404,"Error": 500})
}


func PendingSummaryBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["pending_update"] = bindings.NewIdType([]string{"com.vmware.vcenter.lcm.update.pending"}, "")
	fieldNameMap["pending_update"] = "PendingUpdate"
	fields["version"] = bindings.NewStringType()
	fieldNameMap["version"] = "Version"
	fields["release_date"] = bindings.NewDateTimeType()
	fieldNameMap["release_date"] = "ReleaseDate"
	fields["severity"] = bindings.NewEnumType("com.vmware.vcenter.lcm.update.pending.severity_type", reflect.TypeOf(PendingSeverityType(PendingSeverityType_CRITICAL)))
	fieldNameMap["severity"] = "Severity"
	fields["build"] = bindings.NewStringType()
	fieldNameMap["build"] = "Build"
	fields["update_type"] = bindings.NewEnumType("com.vmware.vcenter.lcm.update.pending.update_type", reflect.TypeOf(PendingUpdateType(PendingUpdateType_PATCH)))
	fieldNameMap["update_type"] = "UpdateType"
	fields["category"] = bindings.NewEnumType("com.vmware.vcenter.lcm.update.pending.category", reflect.TypeOf(PendingCategory(PendingCategory_SECURITY)))
	fieldNameMap["category"] = "Category"
	fields["reboot_required"] = bindings.NewBooleanType()
	fieldNameMap["reboot_required"] = "RebootRequired"
	fields["execute_URL"] = bindings.NewUriType()
	fieldNameMap["execute_URL"] = "ExecuteURL"
	fields["release_notes"] = bindings.NewListType(bindings.NewUriType(), reflect.TypeOf([]url.URL{}))
	fieldNameMap["release_notes"] = "ReleaseNotes"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vcenter.lcm.update.pending.summary", fields, reflect.TypeOf(PendingSummary{}), fieldNameMap, validators)
}

func PendingListResultBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["last_check_time"] = bindings.NewDateTimeType()
	fieldNameMap["last_check_time"] = "LastCheckTime"
	fields["update_count"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fieldNameMap["update_count"] = "UpdateCount"
	fields["upgrade_count"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fieldNameMap["upgrade_count"] = "UpgradeCount"
	fields["updates"] = bindings.NewListType(bindings.NewReferenceType(PendingSummaryBindingType), reflect.TypeOf([]PendingSummary{}))
	fieldNameMap["updates"] = "Updates"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vcenter.lcm.update.pending.list_result", fields, reflect.TypeOf(PendingListResult{}), fieldNameMap, validators)
}

func PendingInfoBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["description"] = bindings.NewStringType()
	fieldNameMap["description"] = "Description"
	fields["pending_update"] = bindings.NewIdType([]string{"com.vmware.vcenter.lcm.update.pending"}, "")
	fieldNameMap["pending_update"] = "PendingUpdate"
	fields["version"] = bindings.NewStringType()
	fieldNameMap["version"] = "Version"
	fields["release_date"] = bindings.NewDateTimeType()
	fieldNameMap["release_date"] = "ReleaseDate"
	fields["severity"] = bindings.NewEnumType("com.vmware.vcenter.lcm.update.pending.severity_type", reflect.TypeOf(PendingSeverityType(PendingSeverityType_CRITICAL)))
	fieldNameMap["severity"] = "Severity"
	fields["build"] = bindings.NewStringType()
	fieldNameMap["build"] = "Build"
	fields["update_type"] = bindings.NewEnumType("com.vmware.vcenter.lcm.update.pending.update_type", reflect.TypeOf(PendingUpdateType(PendingUpdateType_PATCH)))
	fieldNameMap["update_type"] = "UpdateType"
	fields["category"] = bindings.NewEnumType("com.vmware.vcenter.lcm.update.pending.category", reflect.TypeOf(PendingCategory(PendingCategory_SECURITY)))
	fieldNameMap["category"] = "Category"
	fields["reboot_required"] = bindings.NewBooleanType()
	fieldNameMap["reboot_required"] = "RebootRequired"
	fields["execute_URL"] = bindings.NewUriType()
	fieldNameMap["execute_URL"] = "ExecuteURL"
	fields["release_notes"] = bindings.NewListType(bindings.NewUriType(), reflect.TypeOf([]url.URL{}))
	fieldNameMap["release_notes"] = "ReleaseNotes"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vcenter.lcm.update.pending.info", fields, reflect.TypeOf(PendingInfo{}), fieldNameMap, validators)
}

