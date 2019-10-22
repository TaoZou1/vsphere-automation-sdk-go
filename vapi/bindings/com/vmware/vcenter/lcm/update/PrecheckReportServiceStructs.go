/* Copyright © 2019 VMware, Inc. All Rights Reserved.
     SPDX-License-Identifier: BSD-2-Clause */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Data type definitions file for service: PrecheckReport.
 * Includes binding types of a structures and enumerations defined in the service.
 * Shared by client-side stubs and server-side skeletons to ensure type
 * compatibility.
 */

package update

import (
    "reflect"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/bindings/com/vmware/vcenter/lcm"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/bindings"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/data"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/protocol"
    "time"
)

// Resource type for precheck report
const PrecheckReport_RESOURCE_TYPE = "com.vmware.vcenter.lcm.report"



// The ``Summary`` Class contains the summary of precheck report.
 type PrecheckReportReportSummary struct {
    // Number of errors detected by precheck process
    ErrorCount int64
    // Number of warnings detected by precheck process
    WarningCount int64
}


// TODO some error doesn't have ErrorType. Resolve this later VAPI-3009
//


// The ``Report`` class contains estimates of how long it will take an update as well as a list of possible warnings and errors with applying the update.
 type PrecheckReportReport struct {
    // Time when this precheck report was generated
    DateCreated time.Time
    // Rough estimate of time to update vCenter Server in minutes.
    EstimatedTimeToUpdate *int64
    // Lists of the issues and warnings
    Issues *lcm.Notifications
    // A summary of the report consist of count of warnings and errors returned by running the precheck.
    Summary PrecheckReportReportSummary
}


// TODO some error doesn't have ErrorType. Resolve this later VAPI-3009
//


// The ``Result`` class contains the precheck report and a link to download the CSV report.
 type PrecheckReportResult struct {
    // The report generated by running the precheck.
    Report PrecheckReportReport
    // The identifier of CSV formatted precheck report. null provides location where the CSV report can be downloaded from based on the ``csvReport``.
    CsvReport *string
}


// TODO some error doesn't have ErrorType. Resolve this later VAPI-3009
//





func precheckReportCreateInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["version"] = bindings.NewIdType([]string {"com.vmware.vcenter.lcm.update.pending"}, "")
    fieldNameMap["version"] = "Version"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func precheckReportCreateOutputType() bindings.BindingType {
    return bindings.NewReferenceType(PrecheckReportResultBindingType)
}

func precheckReportCreateRestMetadata() protocol.OperationRestMetadata {
    paramsTypeMap := map[string]bindings.BindingType{}
    pathParams := map[string]string{}
    queryParams := map[string]string{}
    headerParams := map[string]string{}
    paramsTypeMap["version"] = bindings.NewIdType([]string {"com.vmware.vcenter.lcm.update.pending"}, "")
    paramsTypeMap["version"] = bindings.NewIdType([]string {"com.vmware.vcenter.lcm.update.pending"}, "")
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
      "POST",
      "/vcenter/lcm/update/pending/{version}/precheck-report",
       resultHeaders,
       202,
       errorHeaders,
       map[string]int{"Unauthenticated": 401,"NotFound": 404,"NotAllowedInCurrentState": 400,"Error": 500})
}



func PrecheckReportReportSummaryBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["error_count"] = bindings.NewIntegerType()
    fieldNameMap["error_count"] = "ErrorCount"
    fields["warning_count"] = bindings.NewIntegerType()
    fieldNameMap["warning_count"] = "WarningCount"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.lcm.update.precheck_report.report_summary",fields, reflect.TypeOf(PrecheckReportReportSummary{}), fieldNameMap, validators)
}

func PrecheckReportReportBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["date_created"] = bindings.NewDateTimeType()
    fieldNameMap["date_created"] = "DateCreated"
    fields["estimated_time_to_update"] = bindings.NewOptionalType(bindings.NewIntegerType())
    fieldNameMap["estimated_time_to_update"] = "EstimatedTimeToUpdate"
    fields["issues"] = bindings.NewOptionalType(bindings.NewReferenceType(lcm.NotificationsBindingType))
    fieldNameMap["issues"] = "Issues"
    fields["summary"] = bindings.NewReferenceType(PrecheckReportReportSummaryBindingType)
    fieldNameMap["summary"] = "Summary"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.lcm.update.precheck_report.report",fields, reflect.TypeOf(PrecheckReportReport{}), fieldNameMap, validators)
}

func PrecheckReportResultBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["report"] = bindings.NewReferenceType(PrecheckReportReportBindingType)
    fieldNameMap["report"] = "Report"
    fields["csv_report"] = bindings.NewOptionalType(bindings.NewIdType([]string {"com.vmware.vcenter.lcm.report"}, ""))
    fieldNameMap["csv_report"] = "CsvReport"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.lcm.update.precheck_report.result",fields, reflect.TypeOf(PrecheckReportResult{}), fieldNameMap, validators)
}


