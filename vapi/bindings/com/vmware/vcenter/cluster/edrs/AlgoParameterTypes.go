/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Data type definitions file for service: AlgoParameter.
 * Includes binding types of a structures and enumerations defined in the service.
 * Shared by client-side stubs and server-side skeletons to ensure type
 * compatibility.
 */
package edrs

import (
	"reflect"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/bindings"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/data"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/protocol"
)

// The ``Param`` class contains EDRS algorithm parameters. **Warning:** This class is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
type AlgoParameterParam struct {
    // Generate periodical EDRS recommendation event. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
	GenerateRecommendationEvent *bool
    // Maximum number of hosts allowed in the cluster. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
	MaxHosts *int64
    // Minimum number of hosts allowed in the cluster. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
	MinHosts *int64
    // CPU low threshold to trigger scale in. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
	CpuScaleInThreshold *int64
    // CPU high threshold to trigger scale out. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
	CpuScaleOutThreshold *int64
    // Weight to compute CPU usage EWMA to trigger scale in. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
	CpuScaleInEwmaWeightPercent *int64
    // Weight to compute CPU usage EWMA to trigger scale out. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
	CpuScaleOutEwmaWeightPercent *int64
    // Memory low threshold to trigger scale in. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
	MemoryScaleInThreshold *int64
    // Memory high threshold to trigger scale out. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
	MemoryScaleOutThreshold *int64
    // Weight to compute memory usage EWMA to trigger scale in. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
	MemoryScaleInEwmaWeightPercent *int64
    // Weight to compute memory usage EWMA to trigger scale out. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
	MemoryScaleOutEwmaWeightPercent *int64
    // vSAN storage low threshold to trigger scale in. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
	VsanScaleInThreshold *int64
    // vSAN storage high threshold to trigger scale out. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
	VsanScaleOutThreshold *int64
    // Weight to compute vSAN storage usage EWMA to trigger scale in. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
	VsanScaleInEwmaWeightPercent *int64
    // Weight to compute vSAN storage usage EWMA to trigger scale out. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
	VsanScaleOutEwmaWeightPercent *int64
    // Only vSAN storage scale out recommendation event can be generated. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
	StorageScaleOutOnly *bool
}


func algoParameterSetInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["cluster"] = bindings.NewIdType([]string{"ClusterComputeResource"}, "")
	fields["params"] = bindings.NewReferenceType(AlgoParameterParamBindingType)
	fieldNameMap["cluster"] = "Cluster"
	fieldNameMap["params"] = "Params"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func algoParameterSetOutputType() bindings.BindingType {
	return bindings.NewVoidType()
}

func algoParameterSetRestMetadata() protocol.OperationRestMetadata {
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
		map[string]int{"Error": 500,"NotFound": 404,"Unauthorized": 403})
}

func algoParameterGetInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["cluster"] = bindings.NewIdType([]string{"ClusterComputeResource"}, "")
	fieldNameMap["cluster"] = "Cluster"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func algoParameterGetOutputType() bindings.BindingType {
	return bindings.NewReferenceType(AlgoParameterParamBindingType)
}

func algoParameterGetRestMetadata() protocol.OperationRestMetadata {
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
		map[string]int{"NotFound": 404,"Unauthorized": 403})
}


func AlgoParameterParamBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["generate_recommendation_event"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fieldNameMap["generate_recommendation_event"] = "GenerateRecommendationEvent"
	fields["max_hosts"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fieldNameMap["max_hosts"] = "MaxHosts"
	fields["min_hosts"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fieldNameMap["min_hosts"] = "MinHosts"
	fields["cpu_scale_in_threshold"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fieldNameMap["cpu_scale_in_threshold"] = "CpuScaleInThreshold"
	fields["cpu_scale_out_threshold"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fieldNameMap["cpu_scale_out_threshold"] = "CpuScaleOutThreshold"
	fields["cpu_scale_in_ewma_weight_percent"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fieldNameMap["cpu_scale_in_ewma_weight_percent"] = "CpuScaleInEwmaWeightPercent"
	fields["cpu_scale_out_ewma_weight_percent"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fieldNameMap["cpu_scale_out_ewma_weight_percent"] = "CpuScaleOutEwmaWeightPercent"
	fields["memory_scale_in_threshold"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fieldNameMap["memory_scale_in_threshold"] = "MemoryScaleInThreshold"
	fields["memory_scale_out_threshold"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fieldNameMap["memory_scale_out_threshold"] = "MemoryScaleOutThreshold"
	fields["memory_scale_in_ewma_weight_percent"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fieldNameMap["memory_scale_in_ewma_weight_percent"] = "MemoryScaleInEwmaWeightPercent"
	fields["memory_scale_out_ewma_weight_percent"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fieldNameMap["memory_scale_out_ewma_weight_percent"] = "MemoryScaleOutEwmaWeightPercent"
	fields["vsan_scale_in_threshold"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fieldNameMap["vsan_scale_in_threshold"] = "VsanScaleInThreshold"
	fields["vsan_scale_out_threshold"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fieldNameMap["vsan_scale_out_threshold"] = "VsanScaleOutThreshold"
	fields["vsan_scale_in_ewma_weight_percent"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fieldNameMap["vsan_scale_in_ewma_weight_percent"] = "VsanScaleInEwmaWeightPercent"
	fields["vsan_scale_out_ewma_weight_percent"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fieldNameMap["vsan_scale_out_ewma_weight_percent"] = "VsanScaleOutEwmaWeightPercent"
	fields["storage_scale_out_only"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fieldNameMap["storage_scale_out_only"] = "StorageScaleOutOnly"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vcenter.cluster.edrs.algo_parameter.param", fields, reflect.TypeOf(AlgoParameterParam{}), fieldNameMap, validators)
}

