/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Code generated. DO NOT EDIT.

/*
 * Data type definitions file for service: Configuration.
 * Includes binding types of a structures and enumerations defined in the service.
 * Shared by client-side stubs and server-side skeletons to ensure type
 * compatibility.
 */

package vstats

import (
	"reflect"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/bindings"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/data"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/protocol"
)



// The amount of information captured in the log files varies, depending on the level setting.
//
// <p> See {@link com.vmware.vapi.bindings.ApiEnumeration enumerated types description}.
type ConfigurationLogLevel string

const (
    // Disables logging.
	ConfigurationLogLevel_NONE ConfigurationLogLevel = "NONE"
    // Logging limited to error messages.
	ConfigurationLogLevel_ERROR ConfigurationLogLevel = "ERROR"
    // Error messages plus warning messages are logged.
	ConfigurationLogLevel_WARNING ConfigurationLogLevel = "WARNING"
    // Default setting. Errors, warnings, plus informational messages about normal operations are logged. Acceptable for production environments.
	ConfigurationLogLevel_INFO ConfigurationLogLevel = "INFO"
    // Can facilitate troubleshooting and debugging. Not recommended for production environments.
	ConfigurationLogLevel_VERBOSE ConfigurationLogLevel = "VERBOSE"
    // Extended verbose logging. Provides complete detail. Use for debugging and to facilitate client application development only. Not recommended for production environments.
	ConfigurationLogLevel_TRIVIA ConfigurationLogLevel = "TRIVIA"
)

func (l ConfigurationLogLevel) ConfigurationLogLevel() bool {
	switch l {
	case ConfigurationLogLevel_NONE:
		return true
	case ConfigurationLogLevel_ERROR:
		return true
	case ConfigurationLogLevel_WARNING:
		return true
	case ConfigurationLogLevel_INFO:
		return true
	case ConfigurationLogLevel_VERBOSE:
		return true
	case ConfigurationLogLevel_TRIVIA:
		return true
	default:
		return false
	}
}


// The ``Info`` class contains vStats configuration.
type ConfigurationInfo struct {
    // vStats log level.
	LogLevel ConfigurationLogLevel
}

// The ``UpdateSpec`` class contains modifiable properties from vStats configuration.
type ConfigurationUpdateSpec struct {
    // vStats log level.
	LogLevel *ConfigurationLogLevel
}



func configurationUpdateInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["update_spec"] = bindings.NewReferenceType(ConfigurationUpdateSpecBindingType)
	fieldNameMap["update_spec"] = "UpdateSpec"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func configurationUpdateOutputType() bindings.BindingType {
	return bindings.NewVoidType()
}

func configurationUpdateRestMetadata() protocol.OperationRestMetadata {
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	paramsTypeMap["update_spec"] = bindings.NewReferenceType(ConfigurationUpdateSpecBindingType)
	resultHeaders := map[string]string{}
	errorHeaders := map[string]string{}
	errorHeaders["Unauthenticated.challenge"] = "WWW-Authenticate"
	return protocol.NewOperationRestMetadata(
		paramsTypeMap,
		pathParams,
		queryParams,
		headerParams,
		"update_spec",
		"PATCH",
		"/stats/configuration",
		resultHeaders,
		204,
		errorHeaders,
		map[string]int{"Error": 500,"InvalidArgument": 400,"Unauthenticated": 401,"Unauthorized": 403})
}

func configurationGetInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func configurationGetOutputType() bindings.BindingType {
	return bindings.NewReferenceType(ConfigurationInfoBindingType)
}

func configurationGetRestMetadata() protocol.OperationRestMetadata {
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
		"/stats/configuration",
		resultHeaders,
		200,
		errorHeaders,
		map[string]int{"Error": 500,"Unauthenticated": 401,"Unauthorized": 403})
}


func ConfigurationInfoBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["log_level"] = bindings.NewEnumType("com.vmware.vstats.configuration.log_level", reflect.TypeOf(ConfigurationLogLevel(ConfigurationLogLevel_NONE)))
	fieldNameMap["log_level"] = "LogLevel"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vstats.configuration.info", fields, reflect.TypeOf(ConfigurationInfo{}), fieldNameMap, validators)
}

func ConfigurationUpdateSpecBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["log_level"] = bindings.NewOptionalType(bindings.NewEnumType("com.vmware.vstats.configuration.log_level", reflect.TypeOf(ConfigurationLogLevel(ConfigurationLogLevel_NONE))))
	fieldNameMap["log_level"] = "LogLevel"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vstats.configuration.update_spec", fields, reflect.TypeOf(ConfigurationUpdateSpec{}), fieldNameMap, validators)
}

