/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Data type definitions file for service: InteropReport.
 * Includes binding types of a structures and enumerations defined in the service.
 * Shared by client-side stubs and server-side skeletons to ensure type
 * compatibility.
 */
package discovery

import (
	"reflect"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/bindings/com/vmware/vcenter/lcm"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/bindings"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/data"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/protocol"
	"net/url"
	"time"
)

// The ``ReleaseInfo`` class contains a product release information.
type InteropReportReleaseInfo struct {
    // The version of the release.
	Version string
    // A link to the release notes of the release.
	Note *url.URL
}

// The ``ReportRow`` class contains the interoperability between a given product and the target product.
type InteropReportReportRow struct {
    // The product to compare to the target product.
	Product Product
    // Defines whether the product is compatible against the target product.
	Compatible bool
    // A list of compatible releases of the product with the target product.
	CompatibleReleases []InteropReportReleaseInfo
}

// The ``ReportSummary`` class contains a summary of the InteropReportReport#products. It consists of the count of compatible and incompatible products to the target product.
type InteropReportReportSummary struct {
    // Number of compatible products.
	CompatibleCount int64
    // Number of incompatible products.
	IncompatibleCount int64
}

// The ``Report`` class contains the interoperability report between the target product and the other registered products in the vCenter Server instance.
type InteropReportReport struct {
    // Time when the report is created.
	DateCreated time.Time
    // A product against the other products are compared to. Only vCenter Server is supported.
	TargetProduct Product
    // Interoperability matrix for the supplied target product and the other registered products.
	Products []InteropReportReportRow
    // Lists of issues encountered during report creation.
	Issues *lcm.Notifications
    // A summary of the interoperability matrix.
	Summary InteropReportReportSummary
}

// The ``Result`` class contains the result of interoperability report creation operation.
type InteropReportResult struct {
    // The interoperability report.
	Report InteropReportReport
    // The identifier of CSV formatted interopability report. 
    //
    //  null provides location where the CSV report can be downloaded from based on the ``csvReport``.
	CsvReport *string
}

// Configuration of report generation.
type InteropReportSpec struct {
    // The vCenter Server version. 
    //
    //  It is used for checking against the other products registered with that instance of vCenter Server.
	TargetVersion string
}


func interopReportCreateInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["spec"] = bindings.NewOptionalType(bindings.NewReferenceType(InteropReportSpecBindingType))
	fieldNameMap["spec"] = "Spec"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func interopReportCreateOutputType() bindings.BindingType {
	return bindings.NewReferenceType(InteropReportResultBindingType)
}

func interopReportCreateRestMetadata() protocol.OperationRestMetadata {
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	paramsTypeMap["spec"] = bindings.NewOptionalType(bindings.NewReferenceType(InteropReportSpecBindingType))
	resultHeaders := map[string]string{}
	errorHeaders := map[string]string{}
	errorHeaders["Unauthenticated.challenge"] = "WWW-Authenticate"
	return protocol.NewOperationRestMetadata(
		paramsTypeMap,
		pathParams,
		queryParams,
		headerParams,
		"spec",
		"POST",
		"/vcenter/lcm/discovery/interop-report",
		resultHeaders,
		202,
		errorHeaders,
		map[string]int{"Unauthenticated": 401,"Error": 500})
}


func InteropReportReleaseInfoBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["version"] = bindings.NewStringType()
	fieldNameMap["version"] = "Version"
	fields["note"] = bindings.NewOptionalType(bindings.NewUriType())
	fieldNameMap["note"] = "Note"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vcenter.lcm.discovery.interop_report.release_info", fields, reflect.TypeOf(InteropReportReleaseInfo{}), fieldNameMap, validators)
}

func InteropReportReportRowBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["product"] = bindings.NewReferenceType(ProductBindingType)
	fieldNameMap["product"] = "Product"
	fields["compatible"] = bindings.NewBooleanType()
	fieldNameMap["compatible"] = "Compatible"
	fields["compatible_releases"] = bindings.NewListType(bindings.NewReferenceType(InteropReportReleaseInfoBindingType), reflect.TypeOf([]InteropReportReleaseInfo{}))
	fieldNameMap["compatible_releases"] = "CompatibleReleases"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vcenter.lcm.discovery.interop_report.report_row", fields, reflect.TypeOf(InteropReportReportRow{}), fieldNameMap, validators)
}

func InteropReportReportSummaryBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["compatible_count"] = bindings.NewIntegerType()
	fieldNameMap["compatible_count"] = "CompatibleCount"
	fields["incompatible_count"] = bindings.NewIntegerType()
	fieldNameMap["incompatible_count"] = "IncompatibleCount"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vcenter.lcm.discovery.interop_report.report_summary", fields, reflect.TypeOf(InteropReportReportSummary{}), fieldNameMap, validators)
}

func InteropReportReportBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["date_created"] = bindings.NewDateTimeType()
	fieldNameMap["date_created"] = "DateCreated"
	fields["target_product"] = bindings.NewReferenceType(ProductBindingType)
	fieldNameMap["target_product"] = "TargetProduct"
	fields["products"] = bindings.NewListType(bindings.NewReferenceType(InteropReportReportRowBindingType), reflect.TypeOf([]InteropReportReportRow{}))
	fieldNameMap["products"] = "Products"
	fields["issues"] = bindings.NewOptionalType(bindings.NewReferenceType(lcm.NotificationsBindingType))
	fieldNameMap["issues"] = "Issues"
	fields["summary"] = bindings.NewReferenceType(InteropReportReportSummaryBindingType)
	fieldNameMap["summary"] = "Summary"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vcenter.lcm.discovery.interop_report.report", fields, reflect.TypeOf(InteropReportReport{}), fieldNameMap, validators)
}

func InteropReportResultBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["report"] = bindings.NewReferenceType(InteropReportReportBindingType)
	fieldNameMap["report"] = "Report"
	fields["csv_report"] = bindings.NewOptionalType(bindings.NewIdType([]string{"com.vmware.vcenter.lcm.report"}, ""))
	fieldNameMap["csv_report"] = "CsvReport"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vcenter.lcm.discovery.interop_report.result", fields, reflect.TypeOf(InteropReportResult{}), fieldNameMap, validators)
}

func InteropReportSpecBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["target_version"] = bindings.NewStringType()
	fieldNameMap["target_version"] = "TargetVersion"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vcenter.lcm.discovery.interop_report.spec", fields, reflect.TypeOf(InteropReportSpec{}), fieldNameMap, validators)
}

