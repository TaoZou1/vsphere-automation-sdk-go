/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Code generated. DO NOT EDIT.

/*
 * Data type definitions file for service: Reports.
 * Includes binding types of a structures and enumerations defined in the service.
 * Shared by client-side stubs and server-side skeletons to ensure type
 * compatibility.
 */

package hcl

import (
	"reflect"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/bindings"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/data"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/protocol"
	"net/url"
	"time"
)



// The ``Token`` class contains information about the token required to be passed in the HTTP header in the HTTP GET request to generate the report. **Warning:** This class is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
type ReportsToken struct {
    // A one-time, short-lived token required in the HTTP header of the request to the url. This token needs to be passed in as a header with the name "session-id". **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
	Token string
    // Expiry time of the token. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
	Expiry time.Time
}

// The ``Location`` class contains the URI location to download generated compatibility report, as well as a token required (as a header on the HTTP GET request) to get the report. The validity of the token is 5 minutes. After the token expires, any attempt to call the URI with said token will fail. **Warning:** This class is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
type ReportsLocation struct {
    // Compatibility report download URL. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
	Url url.URL
    // Information about the token required in the HTTP GET request to download the compatibility report. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
	ReportToken ReportsToken
}



func reportsGetInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["report"] = bindings.NewIdType([]string{"com.vmware.esx.hcl.resources.CompatibilityReport"}, "")
	fieldNameMap["report"] = "Report"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func reportsGetOutputType() bindings.BindingType {
	return bindings.NewReferenceType(ReportsLocationBindingType)
}

func reportsGetRestMetadata() protocol.OperationRestMetadata {
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	paramsTypeMap["report"] = bindings.NewIdType([]string{"com.vmware.esx.hcl.resources.CompatibilityReport"}, "")
	paramsTypeMap["report"] = bindings.NewIdType([]string{"com.vmware.esx.hcl.resources.CompatibilityReport"}, "")
	pathParams["report"] = "report"
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
		"/esx/hcl/reports/{report}",
		resultHeaders,
		200,
		errorHeaders,
		map[string]int{"NotFound": 404,"Unauthenticated": 401,"ResourceInaccessible": 500,"Error": 500})
}


func ReportsTokenBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["token"] = bindings.NewSecretType()
	fieldNameMap["token"] = "Token"
	fields["expiry"] = bindings.NewDateTimeType()
	fieldNameMap["expiry"] = "Expiry"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.esx.hcl.reports.token", fields, reflect.TypeOf(ReportsToken{}), fieldNameMap, validators)
}

func ReportsLocationBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["url"] = bindings.NewUriType()
	fieldNameMap["url"] = "Url"
	fields["report_token"] = bindings.NewReferenceType(ReportsTokenBindingType)
	fieldNameMap["report_token"] = "ReportToken"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.esx.hcl.reports.location", fields, reflect.TypeOf(ReportsLocation{}), fieldNameMap, validators)
}

