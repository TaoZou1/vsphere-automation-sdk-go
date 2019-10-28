/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Data type definitions file for service: CompatibilityReleases.
 * Includes binding types of a structures and enumerations defined in the service.
 * Shared by client-side stubs and server-side skeletons to ensure type
 * compatibility.
 */
package hosts

import (
	"reflect"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/bindings/esx/hcl"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/bindings"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/data"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/protocol"
)



// This ``EsxiCompatibilityReleases`` class contains properties that describe available releases for generating compatibility report for a specific ESXi host. **Warning:** This class is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
type CompatibilityReleasesEsxiCompatibilityReleases struct {
    // The current release of the ESXi, which also can be checked for compatibility. 
    //
    //  The information for the release does not include patch information.. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
	CurrentCompatibilityRelease string
    // The available ESXi releases, greater than the current one, than can be checked for compatibility. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
	NewerCompatibilityReleases []string
    // Notifications returned by the operation. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
	Notifications hcl.Notifications
}



func compatibilityReleasesListInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["host"] = bindings.NewIdType([]string{"HostSystem"}, "")
	fieldNameMap["host"] = "Host"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func compatibilityReleasesListOutputType() bindings.BindingType {
	return bindings.NewReferenceType(CompatibilityReleasesEsxiCompatibilityReleasesBindingType)
}

func compatibilityReleasesListRestMetadata() protocol.OperationRestMetadata {
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	paramsTypeMap["host"] = bindings.NewIdType([]string{"HostSystem"}, "")
	paramsTypeMap["host"] = bindings.NewIdType([]string{"HostSystem"}, "")
	pathParams["host"] = "host"
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
		"/esx/hcl/hosts/{host}/compatibility-releases",
		resultHeaders,
		200,
		errorHeaders,
		map[string]int{"InternalServerError": 500,"NotAllowedInCurrentState": 400,"NotFound": 404,"Unauthenticated": 401,"Unsupported": 400,"ResourceInaccessible": 500,"Error": 500})
}


func CompatibilityReleasesEsxiCompatibilityReleasesBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["current_compatibility_release"] = bindings.NewStringType()
	fieldNameMap["current_compatibility_release"] = "CurrentCompatibilityRelease"
	fields["newer_compatibility_releases"] = bindings.NewListType(bindings.NewStringType(), reflect.TypeOf([]string{}))
	fieldNameMap["newer_compatibility_releases"] = "NewerCompatibilityReleases"
	fields["notifications"] = bindings.NewReferenceType(hcl.NotificationsBindingType)
	fieldNameMap["notifications"] = "Notifications"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.esx.hcl.hosts.compatibility_releases.esxi_compatibility_releases", fields, reflect.TypeOf(CompatibilityReleasesEsxiCompatibilityReleases{}), fieldNameMap, validators)
}

