/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Data type definitions file for service: License.
 * Includes binding types of a structures and enumerations defined in the service.
 * Shared by client-side stubs and server-side skeletons to ensure type
 * compatibility.
 */
package vcha

import (
	"reflect"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/bindings"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/data"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/protocol"
)



// The ``Info`` class contains information about the license of the active node of a VCHA configuration. **Warning:** This class is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
type LicenseInfo struct {
    // Does the vCenter have the appropriate license for VCHA. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
	IsAvailable bool
}



func licenseGetInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func licenseGetOutputType() bindings.BindingType {
	return bindings.NewReferenceType(LicenseInfoBindingType)
}

func licenseGetRestMetadata() protocol.OperationRestMetadata {
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
		map[string]int{"Unauthorized": 403,"Error": 500})
}


func LicenseInfoBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["is_available"] = bindings.NewBooleanType()
	fieldNameMap["is_available"] = "IsAvailable"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vcenter.vcha.license.info", fields, reflect.TypeOf(LicenseInfo{}), fieldNameMap, validators)
}

