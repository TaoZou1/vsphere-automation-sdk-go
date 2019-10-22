/* Copyright © 2019 VMware, Inc. All Rights Reserved.
     SPDX-License-Identifier: BSD-2-Clause */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Data type definitions file for service: Reports.
 * Includes binding types of a structures and enumerations defined in the service.
 * Shared by client-side stubs and server-side skeletons to ensure type
 * compatibility.
 */

package lcm

import (
    "reflect"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/bindings"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/data"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/protocol"
    "net/url"
    "time"
)




// The ``Token`` class contains information about the token required in the HTTP GET request to generate the report.
 type ReportsToken struct {
    // A one-time, short-lived token required in the HTTP header of the request to the url. This token needs to be passed in as a header with the name "session-id".
    Token string
    // Expiry time of the token
    Expiry time.Time
}


// TODO some error doesn't have ErrorType. Resolve this later VAPI-3009
//


// The ``Location`` class contains the URI location to download the report from, as well as a token required (as a header on the HTTP request) to get the bundle. The validity of the token is 5 minutes as best attempt. After the token expires, any attempt to call the URI with said token will fail.
 type ReportsLocation struct {
    // Report Download URI.
    Uri url.URL
    // Information about the token required in the HTTP GET request to retrieve the report
    DownloadFileToken ReportsToken
}


// TODO some error doesn't have ErrorType. Resolve this later VAPI-3009
//





func reportsGetInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["report"] = bindings.NewIdType([]string {"com.vmware.vcenter.lcm.report"}, "")
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
    paramsTypeMap["report"] = bindings.NewIdType([]string {"com.vmware.vcenter.lcm.report"}, "")
    paramsTypeMap["report"] = bindings.NewIdType([]string {"com.vmware.vcenter.lcm.report"}, "")
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
      "/vcenter/lcm/reports/{report}",
       resultHeaders,
       200,
       errorHeaders,
       map[string]int{"Unauthenticated": 401,"NotFound": 404,"Error": 500})
}



func ReportsTokenBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["token"] = bindings.NewStringType()
    fieldNameMap["token"] = "Token"
    fields["expiry"] = bindings.NewDateTimeType()
    fieldNameMap["expiry"] = "Expiry"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.lcm.reports.token",fields, reflect.TypeOf(ReportsToken{}), fieldNameMap, validators)
}

func ReportsLocationBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["uri"] = bindings.NewUriType()
    fieldNameMap["uri"] = "Uri"
    fields["download_file_token"] = bindings.NewReferenceType(ReportsTokenBindingType)
    fieldNameMap["download_file_token"] = "DownloadFileToken"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.lcm.reports.location",fields, reflect.TypeOf(ReportsLocation{}), fieldNameMap, validators)
}


