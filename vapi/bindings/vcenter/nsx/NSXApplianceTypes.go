/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Code generated. DO NOT EDIT.

/*
 * Data type definitions file for service: NSXAppliance.
 * Includes binding types of a structures and enumerations defined in the service.
 * Shared by client-side stubs and server-side skeletons to ensure type
 * compatibility.
 */

package nsx

import (
	"reflect"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/bindings/vapi/std"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/bindings"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/data"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/protocol"
)



// The ``Operation`` enumeration class describes the operation in progress. **Warning:** This enumeration is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
//
// <p> See {@link com.vmware.vapi.bindings.ApiEnumeration enumerated types description}.
type NSXApplianceOperation string

const (
    // Install the NSX appliance on the vCenter. **Warning:** This constant field is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
	NSXApplianceOperation_ENABLE NSXApplianceOperation = "ENABLE"
    // Delete the NSX appliance on the vCenter. **Warning:** This constant field is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
	NSXApplianceOperation_DISABLE NSXApplianceOperation = "DISABLE"
    // No ongoing operations. **Warning:** This constant field is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
	NSXApplianceOperation_NONE NSXApplianceOperation = "NONE"
)

func (o NSXApplianceOperation) NSXApplianceOperation() bool {
	switch o {
	case NSXApplianceOperation_ENABLE:
		return true
	case NSXApplianceOperation_DISABLE:
		return true
	case NSXApplianceOperation_NONE:
		return true
	default:
		return false
	}
}


// The ``Phase`` enumeration class represents the current phase of the object with respect to realizing the deployment. **Warning:** This enumeration is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
//
// <p> See {@link com.vmware.vapi.bindings.ApiEnumeration enumerated types description}.
type NSXAppliancePhase string

const (
    // Precheck phase validates if the NSX deployment can complete successfully. **Warning:** This constant field is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
	NSXAppliancePhase_PRECHECK NSXAppliancePhase = "PRECHECK"
    // NSX ovf deployment phase. **Warning:** This constant field is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
	NSXAppliancePhase_DEPLOY NSXAppliancePhase = "DEPLOY"
    // Configuration phase of the deployed NSX appliance. **Warning:** This constant field is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
	NSXAppliancePhase_CONFIGURE NSXAppliancePhase = "CONFIGURE"
)

func (p NSXAppliancePhase) NSXAppliancePhase() bool {
	switch p {
	case NSXAppliancePhase_PRECHECK:
		return true
	case NSXAppliancePhase_DEPLOY:
		return true
	case NSXAppliancePhase_CONFIGURE:
		return true
	default:
		return false
	}
}


// The ``ConfigStatus`` enumeration class describes the status of NSX deployment. **Warning:** This enumeration is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
//
// <p> See {@link com.vmware.vapi.bindings.ApiEnumeration enumerated types description}.
type NSXApplianceConfigStatus string

const (
    // NSX is not installed on the vCenter Server. **Warning:** This constant field is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
	NSXApplianceConfigStatus_UNCONFIGURED NSXApplianceConfigStatus = "UNCONFIGURED"
    // NSX is successfully configured on the vCenter Server. **Warning:** This constant field is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
	NSXApplianceConfigStatus_CONFIGURED NSXApplianceConfigStatus = "CONFIGURED"
    // NSX deployment is in progress. **Warning:** This constant field is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
	NSXApplianceConfigStatus_INPROGRESS NSXApplianceConfigStatus = "INPROGRESS"
    // NSX deployment is failed. Retry the deployment after resolving the errors. **Warning:** This constant field is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
	NSXApplianceConfigStatus_ERROR NSXApplianceConfigStatus = "ERROR"
)

func (c NSXApplianceConfigStatus) NSXApplianceConfigStatus() bool {
	switch c {
	case NSXApplianceConfigStatus_UNCONFIGURED:
		return true
	case NSXApplianceConfigStatus_CONFIGURED:
		return true
	case NSXApplianceConfigStatus_INPROGRESS:
		return true
	case NSXApplianceConfigStatus_ERROR:
		return true
	default:
		return false
	}
}


// The ``ProductType`` enumeration class describes the type of NSX deployment. **Warning:** This enumeration is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
//
// <p> See {@link com.vmware.vapi.bindings.ApiEnumeration enumerated types description}.
type NSXApplianceProductType string

const (
    // Integrated NSX shipped with vSphere. **Warning:** This constant field is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
	NSXApplianceProductType_NSXI NSXApplianceProductType = "NSXI"
    // Licensed version of NSX. **Warning:** This constant field is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
	NSXApplianceProductType_NSXT NSXApplianceProductType = "NSXT"
)

func (p NSXApplianceProductType) NSXApplianceProductType() bool {
	switch p {
	case NSXApplianceProductType_NSXI:
		return true
	case NSXApplianceProductType_NSXT:
		return true
	default:
		return false
	}
}


// The ``Message`` class contains the information about the ongoing deployment. **Warning:** This class is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
type NSXApplianceMessage struct {
    // Type of the message. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
	Severity NSXApplianceMessageSeverity
    // Details about the message. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
	Details *std.LocalizableMessage
}

// The ``Severity`` enumeration class represents the severity of the message. **Warning:** This enumeration is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
//
// <p> See {@link com.vmware.vapi.bindings.ApiEnumeration enumerated types description}.
type NSXApplianceMessageSeverity string

const (
    // Informational message. This may be accompanied by vCenter event. **Warning:** This constant field is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
	NSXApplianceMessageSeverity_INFO NSXApplianceMessageSeverity = "INFO"
    // Warning message. This may be accompanied by vCenter event. **Warning:** This constant field is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
	NSXApplianceMessageSeverity_WARNING NSXApplianceMessageSeverity = "WARNING"
    // Error message. This is accompanied by vCenter event and/or alarm. **Warning:** This constant field is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
	NSXApplianceMessageSeverity_ERROR NSXApplianceMessageSeverity = "ERROR"
)

func (s NSXApplianceMessageSeverity) NSXApplianceMessageSeverity() bool {
	switch s {
	case NSXApplianceMessageSeverity_INFO:
		return true
	case NSXApplianceMessageSeverity_WARNING:
		return true
	case NSXApplianceMessageSeverity_ERROR:
		return true
	default:
		return false
	}
}


// The ``Info`` class contains detailed information about the status of NSX deployment. **Warning:** This class is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
type NSXApplianceInfo struct {
    // Current setting for ``ConfigStatus``. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
	ConfigStatus NSXApplianceConfigStatus
    // The ongoing operation for which we are querying the status value is None if no operation is in progress. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
	Operation NSXApplianceOperation
    // Current phase in realizing ``Operation``. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
	Phase *NSXAppliancePhase
    // NSX integrated or Licensed version. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
	ProductType *NSXApplianceProductType
    // Current set of messages associated with the ongoing deployment. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
	Messages []NSXApplianceMessage
}

// The ``InstallSpec`` contains the inputs related to appliance deployment operation. **Warning:** This class is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
type NSXApplianceInstallSpec struct {
    // This subsection describes the VC inventory on which to deploy the appliance. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
	DestinationLocation PlacementDetails
    // Spec to describe configuration the new NSX appliance. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
	DestinationAppliance ApplianceConfig
	ManagementVcenter *Connection
}



func nSXApplianceGetInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func nSXApplianceGetOutputType() bindings.BindingType {
	return bindings.NewReferenceType(NSXApplianceInfoBindingType)
}

func nSXApplianceGetRestMetadata() protocol.OperationRestMetadata {
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
		"/vcenter/nsx/nsx-appliance",
		resultHeaders,
		200,
		errorHeaders,
		map[string]int{"Error": 500,"NotFound": 404,"Unauthenticated": 401,"Unauthorized": 403})
}

func nSXApplianceCreateInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["spec"] = bindings.NewReferenceType(NSXApplianceInstallSpecBindingType)
	fieldNameMap["spec"] = "Spec"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func nSXApplianceCreateOutputType() bindings.BindingType {
	return bindings.NewVoidType()
}

func nSXApplianceCreateRestMetadata() protocol.OperationRestMetadata {
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	paramsTypeMap["spec"] = bindings.NewReferenceType(NSXApplianceInstallSpecBindingType)
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
		"/vcenter/nsx/nsx-appliance",
		resultHeaders,
		201,
		errorHeaders,
		map[string]int{"Error": 500,"InvalidArgument": 400,"Unauthenticated": 401,"Unauthorized": 403,"NotAllowedInCurrentState": 400,"AlreadyExists": 400})
}

func nSXApplianceDeleteInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func nSXApplianceDeleteOutputType() bindings.BindingType {
	return bindings.NewVoidType()
}

func nSXApplianceDeleteRestMetadata() protocol.OperationRestMetadata {
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
		"DELETE",
		"/vcenter/nsx/nsx-appliance",
		resultHeaders,
		204,
		errorHeaders,
		map[string]int{"Error": 500,"NotFound": 404,"Unauthenticated": 401,"Unauthorized": 403,"NotAllowedInCurrentState": 400})
}


func NSXApplianceMessageBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["severity"] = bindings.NewEnumType("com.vmware.vcenter.nsx.NSX_appliance.message.severity", reflect.TypeOf(NSXApplianceMessageSeverity(NSXApplianceMessageSeverity_INFO)))
	fieldNameMap["severity"] = "Severity"
	fields["details"] = bindings.NewOptionalType(bindings.NewReferenceType(std.LocalizableMessageBindingType))
	fieldNameMap["details"] = "Details"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vcenter.nsx.NSX_appliance.message", fields, reflect.TypeOf(NSXApplianceMessage{}), fieldNameMap, validators)
}

func NSXApplianceInfoBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["config_status"] = bindings.NewEnumType("com.vmware.vcenter.nsx.NSX_appliance.config_status", reflect.TypeOf(NSXApplianceConfigStatus(NSXApplianceConfigStatus_UNCONFIGURED)))
	fieldNameMap["config_status"] = "ConfigStatus"
	fields["operation"] = bindings.NewEnumType("com.vmware.vcenter.nsx.NSX_appliance.operation", reflect.TypeOf(NSXApplianceOperation(NSXApplianceOperation_ENABLE)))
	fieldNameMap["operation"] = "Operation"
	fields["phase"] = bindings.NewOptionalType(bindings.NewEnumType("com.vmware.vcenter.nsx.NSX_appliance.phase", reflect.TypeOf(NSXAppliancePhase(NSXAppliancePhase_PRECHECK))))
	fieldNameMap["phase"] = "Phase"
	fields["product_type"] = bindings.NewOptionalType(bindings.NewEnumType("com.vmware.vcenter.nsx.NSX_appliance.product_type", reflect.TypeOf(NSXApplianceProductType(NSXApplianceProductType_NSXI))))
	fieldNameMap["product_type"] = "ProductType"
	fields["messages"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewReferenceType(NSXApplianceMessageBindingType), reflect.TypeOf([]NSXApplianceMessage{})))
	fieldNameMap["messages"] = "Messages"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vcenter.nsx.NSX_appliance.info", fields, reflect.TypeOf(NSXApplianceInfo{}), fieldNameMap, validators)
}

func NSXApplianceInstallSpecBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["destination_location"] = bindings.NewReferenceType(PlacementDetailsBindingType)
	fieldNameMap["destination_location"] = "DestinationLocation"
	fields["destination_appliance"] = bindings.NewReferenceType(ApplianceConfigBindingType)
	fieldNameMap["destination_appliance"] = "DestinationAppliance"
	fields["management_vcenter"] = bindings.NewOptionalType(bindings.NewReferenceType(ConnectionBindingType))
	fieldNameMap["management_vcenter"] = "ManagementVcenter"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vcenter.nsx.NSX_appliance.install_spec", fields, reflect.TypeOf(NSXApplianceInstallSpec{}), fieldNameMap, validators)
}

