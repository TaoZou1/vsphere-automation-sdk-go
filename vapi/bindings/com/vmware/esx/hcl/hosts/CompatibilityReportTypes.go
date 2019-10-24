/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Data type definitions file for service: CompatibilityReport.
 * Includes binding types of a structures and enumerations defined in the service.
 * Shared by client-side stubs and server-side skeletons to ensure type
 * compatibility.
 */
package hosts

import (
	"reflect"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/bindings/com/vmware/esx/hcl"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/bindings/com/vmware/vapi/std"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/bindings"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/data"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/protocol"
	"net/url"
	"time"
)

// This ``BiosConstraint`` class contains properties that describe the BIOS that is supported for the given server and ESXi release. **Warning:** This class is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
type CompatibilityReportBiosConstraint struct {
    // The BIOS information about the constraint. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
	Bios hcl.Firmware
    // Any information that should be taken into account when reviewing the BIOS constraint. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
	Notes []std.LocalizableMessage
}

// This ``ServerCompatibility`` class contains properties that provide the compatibility information for a server model, cpu and BIOS. **Warning:** This class is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
type CompatibilityReportServerCompatibility struct {
    // Lists the BIOS constraints that the target ESXi release has for this server. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
	BiosConstraints []CompatibilityReportBiosConstraint
    // The CPU series name. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
	CpuSeries string
    // Provides information about supported releases for this entry. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
	SupportedReleases []string
    // Provides link to the VMware Compatibility Guide for further information on the compatibility. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
	VcgLink url.URL
    // Information that needs to be taken into account when considering this server hardware compatibility. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
	Notes []std.LocalizableMessage
}

// This ``ServerHclInfo`` class contains properties that describe the server of a ESXi host and its compatibility information. **Warning:** This class is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
type CompatibilityReportServerHclInfo struct {
    // Information about the server. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
	Server hcl.Server
    // Provides information about possible compatibility matches for the given server. 
    //
    //  There could be multiple matches returned as there are several possible matches in the Compatibility data.. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
	Matches []CompatibilityReportServerCompatibility
    // Shows if the server model is compatible with given target ESXi release. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
	ModelCompatibility hcl.CompatibilityStatus
}

// This ``DeviceConstraint`` class contains properties that describe pair of driver and firmware that are supported for a given PCI device and ESXi release. **Warning:** This class is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
type CompatibilityReportDeviceConstraint struct {
    // The driver information about the constraint. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
	Driver hcl.Driver
    // The firmware information about the constraint. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
	Firmware *hcl.Firmware
    // Any information that should be taken into account when reviewing the device constraint. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
	Notes []std.LocalizableMessage
}

// This ``DeviceHclInfo`` class contains properties that describe a PCI device of a given ESXi host and its compatibility information. 
//
//  If there are multiple PCI devices of the same type on the host each one will be listed in separate instance of this class.. **Warning:** This class is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
type CompatibilityReportDeviceHclInfo struct {
    // Indicates compatibility status of the PCI device. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
	Compatibility hcl.CompatibilityStatus
    // Information about the PCI device. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
	Device hcl.PCIDevice
    // Lists the constraints the target ESXi release has for this PCI device. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
	DeviceConstraints []CompatibilityReportDeviceConstraint
    // Provides information about supported releases for this device. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
	SupportedReleases []string
    // Provides link to the VMware Compatibility Guide for further information on the compatibility. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
	VcgLink *url.URL
    // Information that needs to be taken into account when considering this device hcl. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
	Notes []std.LocalizableMessage
}

// This ``HclReport`` represents the hardware compatibility report generated for a specific host and target ESXi release. **Warning:** This class is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
type CompatibilityReportHclReport struct {
    // FQDN identifying the ESXi host that the report refers to. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
	Host url.URL
    // Indicates for which ESXi release the report is generated. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
	TargetRelease string
    // Lists compatibility information for the ESXi's server part. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
	ServerHcl CompatibilityReportServerHclInfo
    // Lists compatibility information for discoverable PCI devices of the host. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
	DevicesHcl []CompatibilityReportDeviceHclInfo
    // Specifies the time the report was generated. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
	GeneratedAt time.Time
    // Notifications returned by the operation. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
	Notifications hcl.Notifications
}

// The ``Result`` class contains the result of hardware compatibility report creation operation. **Warning:** This class is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
type CompatibilityReportResult struct {
    // The hardware compatibility report. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
	Report CompatibilityReportHclReport
    // The identifier of the compatibility report. 
    //
    //  Reports#get provides location where a file based report based on the ``HclReport`` can be downloaded using this identifier.. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
	Identifier *string
}

// The ``Spec`` class contains properties to describe the input configuration for an ESXi's compatibility report generation. **Warning:** This class is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
type CompatibilityReportSpec struct {
    // A target ESXi release which will be used to generate a compatibility report. Releases that can be used to generate report can be found using CompatibilityReleases#list. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
	Release string
}


func compatibilityReportCreateInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["host"] = bindings.NewIdType([]string{"HostSystem"}, "")
	fields["spec"] = bindings.NewOptionalType(bindings.NewReferenceType(CompatibilityReportSpecBindingType))
	fieldNameMap["host"] = "Host"
	fieldNameMap["spec"] = "Spec"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func compatibilityReportCreateOutputType() bindings.BindingType {
	return bindings.NewReferenceType(CompatibilityReportResultBindingType)
}

func compatibilityReportCreateRestMetadata() protocol.OperationRestMetadata {
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	paramsTypeMap["spec"] = bindings.NewOptionalType(bindings.NewReferenceType(CompatibilityReportSpecBindingType))
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
		"spec",
		"POST",
		"/esx/hcl/hosts/{host}/compatibility-report",
		resultHeaders,
		202,
		errorHeaders,
		map[string]int{"NotFound": 404,"Unauthenticated": 401,"Unsupported": 400,"ResourceInaccessible": 500,"NotAllowedInCurrentState": 400,"Error": 500})
}


func CompatibilityReportBiosConstraintBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["bios"] = bindings.NewReferenceType(hcl.FirmwareBindingType)
	fieldNameMap["bios"] = "Bios"
	fields["notes"] = bindings.NewListType(bindings.NewReferenceType(std.LocalizableMessageBindingType), reflect.TypeOf([]std.LocalizableMessage{}))
	fieldNameMap["notes"] = "Notes"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.esx.hcl.hosts.compatibility_report.bios_constraint", fields, reflect.TypeOf(CompatibilityReportBiosConstraint{}), fieldNameMap, validators)
}

func CompatibilityReportServerCompatibilityBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["bios_constraints"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewReferenceType(CompatibilityReportBiosConstraintBindingType), reflect.TypeOf([]CompatibilityReportBiosConstraint{})))
	fieldNameMap["bios_constraints"] = "BiosConstraints"
	fields["cpu_series"] = bindings.NewStringType()
	fieldNameMap["cpu_series"] = "CpuSeries"
	fields["supported_releases"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewStringType(), reflect.TypeOf([]string{})))
	fieldNameMap["supported_releases"] = "SupportedReleases"
	fields["vcg_link"] = bindings.NewUriType()
	fieldNameMap["vcg_link"] = "VcgLink"
	fields["notes"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewReferenceType(std.LocalizableMessageBindingType), reflect.TypeOf([]std.LocalizableMessage{})))
	fieldNameMap["notes"] = "Notes"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.esx.hcl.hosts.compatibility_report.server_compatibility", fields, reflect.TypeOf(CompatibilityReportServerCompatibility{}), fieldNameMap, validators)
}

func CompatibilityReportServerHclInfoBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["server"] = bindings.NewReferenceType(hcl.ServerBindingType)
	fieldNameMap["server"] = "Server"
	fields["matches"] = bindings.NewListType(bindings.NewReferenceType(CompatibilityReportServerCompatibilityBindingType), reflect.TypeOf([]CompatibilityReportServerCompatibility{}))
	fieldNameMap["matches"] = "Matches"
	fields["model_compatibility"] = bindings.NewEnumType("com.vmware.esx.hcl.compatibility_status", reflect.TypeOf(hcl.CompatibilityStatus(hcl.CompatibilityStatus_COMPATIBLE)))
	fieldNameMap["model_compatibility"] = "ModelCompatibility"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.esx.hcl.hosts.compatibility_report.server_hcl_info", fields, reflect.TypeOf(CompatibilityReportServerHclInfo{}), fieldNameMap, validators)
}

func CompatibilityReportDeviceConstraintBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["driver"] = bindings.NewReferenceType(hcl.DriverBindingType)
	fieldNameMap["driver"] = "Driver"
	fields["firmware"] = bindings.NewOptionalType(bindings.NewReferenceType(hcl.FirmwareBindingType))
	fieldNameMap["firmware"] = "Firmware"
	fields["notes"] = bindings.NewListType(bindings.NewReferenceType(std.LocalizableMessageBindingType), reflect.TypeOf([]std.LocalizableMessage{}))
	fieldNameMap["notes"] = "Notes"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.esx.hcl.hosts.compatibility_report.device_constraint", fields, reflect.TypeOf(CompatibilityReportDeviceConstraint{}), fieldNameMap, validators)
}

func CompatibilityReportDeviceHclInfoBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["compatibility"] = bindings.NewEnumType("com.vmware.esx.hcl.compatibility_status", reflect.TypeOf(hcl.CompatibilityStatus(hcl.CompatibilityStatus_COMPATIBLE)))
	fieldNameMap["compatibility"] = "Compatibility"
	fields["device"] = bindings.NewReferenceType(hcl.PCIDeviceBindingType)
	fieldNameMap["device"] = "Device"
	fields["device_constraints"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewReferenceType(CompatibilityReportDeviceConstraintBindingType), reflect.TypeOf([]CompatibilityReportDeviceConstraint{})))
	fieldNameMap["device_constraints"] = "DeviceConstraints"
	fields["supported_releases"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewStringType(), reflect.TypeOf([]string{})))
	fieldNameMap["supported_releases"] = "SupportedReleases"
	fields["vcg_link"] = bindings.NewOptionalType(bindings.NewUriType())
	fieldNameMap["vcg_link"] = "VcgLink"
	fields["notes"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewReferenceType(std.LocalizableMessageBindingType), reflect.TypeOf([]std.LocalizableMessage{})))
	fieldNameMap["notes"] = "Notes"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.esx.hcl.hosts.compatibility_report.device_hcl_info", fields, reflect.TypeOf(CompatibilityReportDeviceHclInfo{}), fieldNameMap, validators)
}

func CompatibilityReportHclReportBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["host"] = bindings.NewUriType()
	fieldNameMap["host"] = "Host"
	fields["target_release"] = bindings.NewStringType()
	fieldNameMap["target_release"] = "TargetRelease"
	fields["server_hcl"] = bindings.NewReferenceType(CompatibilityReportServerHclInfoBindingType)
	fieldNameMap["server_hcl"] = "ServerHcl"
	fields["devices_hcl"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewReferenceType(CompatibilityReportDeviceHclInfoBindingType), reflect.TypeOf([]CompatibilityReportDeviceHclInfo{})))
	fieldNameMap["devices_hcl"] = "DevicesHcl"
	fields["generated_at"] = bindings.NewDateTimeType()
	fieldNameMap["generated_at"] = "GeneratedAt"
	fields["notifications"] = bindings.NewReferenceType(hcl.NotificationsBindingType)
	fieldNameMap["notifications"] = "Notifications"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.esx.hcl.hosts.compatibility_report.hcl_report", fields, reflect.TypeOf(CompatibilityReportHclReport{}), fieldNameMap, validators)
}

func CompatibilityReportResultBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["report"] = bindings.NewReferenceType(CompatibilityReportHclReportBindingType)
	fieldNameMap["report"] = "Report"
	fields["identifier"] = bindings.NewOptionalType(bindings.NewIdType([]string{"com.vmware.esx.hcl.resources.CompatibilityReport"}, ""))
	fieldNameMap["identifier"] = "Identifier"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.esx.hcl.hosts.compatibility_report.result", fields, reflect.TypeOf(CompatibilityReportResult{}), fieldNameMap, validators)
}

func CompatibilityReportSpecBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["release"] = bindings.NewStringType()
	fieldNameMap["release"] = "Release"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.esx.hcl.hosts.compatibility_report.spec", fields, reflect.TypeOf(CompatibilityReportSpec{}), fieldNameMap, validators)
}

