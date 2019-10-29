/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Code generated. DO NOT EDIT.

/*
 * Data type definitions file for service: Management.
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



// The ``TsdbSizeInfo`` class contains Time-series-database (TSDB) size information.
type ManagementTsdbSizeInfo struct {
    // Maximum allowed database size(in bytes).
	Maximum int64
    // Actual database size(in bytes). Read only.
	Actual *int64
}

// The ``DbsConfig`` class specifies configuration of database connection.
type ManagementDbsConfig struct {
    // Database name.
	Name string
    // Port of the database.
	Port int64
    // Database version.
	Version string
}

// The ``DbsConfigs`` class specifies configuration information of the databases.
type ManagementDbsConfigs struct {
    // Configuration information of core Database.
	CoredbConfig ManagementDbsConfig
    // Configuration information of state Database.
	SttdbConfig ManagementDbsConfig
    // Configuration information of Time Series Database.
	TsdbConfig ManagementDbsConfig
}

// The ``Info`` class contains information about the database configuration of the vstats service.
type ManagementInfo struct {
    // Database connection information.
	DbsConfigs ManagementDbsConfigs
    // Time series database size information.
	TsDbSize ManagementTsdbSizeInfo
}



func managementGetInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func managementGetOutputType() bindings.BindingType {
	return bindings.NewReferenceType(ManagementInfoBindingType)
}

func managementGetRestMetadata() protocol.OperationRestMetadata {
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
		"/stats/management",
		resultHeaders,
		200,
		errorHeaders,
		map[string]int{"Error": 500,"Unauthenticated": 401,"Unauthorized": 403})
}


func ManagementTsdbSizeInfoBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["maximum"] = bindings.NewIntegerType()
	fieldNameMap["maximum"] = "Maximum"
	fields["actual"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fieldNameMap["actual"] = "Actual"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vstats.management.tsdb_size_info", fields, reflect.TypeOf(ManagementTsdbSizeInfo{}), fieldNameMap, validators)
}

func ManagementDbsConfigBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["name"] = bindings.NewStringType()
	fieldNameMap["name"] = "Name"
	fields["port"] = bindings.NewIntegerType()
	fieldNameMap["port"] = "Port"
	fields["version"] = bindings.NewStringType()
	fieldNameMap["version"] = "Version"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vstats.management.dbs_config", fields, reflect.TypeOf(ManagementDbsConfig{}), fieldNameMap, validators)
}

func ManagementDbsConfigsBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["coredb_config"] = bindings.NewReferenceType(ManagementDbsConfigBindingType)
	fieldNameMap["coredb_config"] = "CoredbConfig"
	fields["sttdb_config"] = bindings.NewReferenceType(ManagementDbsConfigBindingType)
	fieldNameMap["sttdb_config"] = "SttdbConfig"
	fields["tsdb_config"] = bindings.NewReferenceType(ManagementDbsConfigBindingType)
	fieldNameMap["tsdb_config"] = "TsdbConfig"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vstats.management.dbs_configs", fields, reflect.TypeOf(ManagementDbsConfigs{}), fieldNameMap, validators)
}

func ManagementInfoBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["dbs_configs"] = bindings.NewReferenceType(ManagementDbsConfigsBindingType)
	fieldNameMap["dbs_configs"] = "DbsConfigs"
	fields["ts_db_size"] = bindings.NewReferenceType(ManagementTsdbSizeInfoBindingType)
	fieldNameMap["ts_db_size"] = "TsDbSize"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vstats.management.info", fields, reflect.TypeOf(ManagementInfo{}), fieldNameMap, validators)
}

