/* Copyright © 2019 VMware, Inc. All Rights Reserved.
     SPDX-License-Identifier: BSD-2-Clause */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Data type definitions file for service: TimeSeries.
 * Includes binding types of a structures and enumerations defined in the service.
 * Shared by client-side stubs and server-side skeletons to ensure type
 * compatibility.
 */

package stats

import (
    "reflect"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/bindings"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/data"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/protocol"
)




// A set of timestamps and statistical values representing a time series. The lengths of TimeSeriesTimeSeries#timeStamps and TimeSeriesTimeSeries#values will always match each other. **Warning:** This class is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
 type TimeSeriesTimeSeries struct {
    // Counter identifier. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    Counter string
    // Sequence of UNIX timestamp values at which statistical values were sampled. https://en.wikipedia.org/wiki/Unix_time. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    TimeStamps []int64
    // Sequence of sampled values corresponding to the timestamps in tss. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    Values []int64
}



//


// Pod identifier. These are the fields required to uniquely identify a pod. **Warning:** This class is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
 type TimeSeriesPodIdentifier struct {
    // The namespace that the pod is running in. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    Namespace string
    // The name of the pod itself. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    PodName string
}



//


// This structure is sent in a request for TimeSeries data and is used to specify what object stats should be returned for. **Warning:** This class is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
 type TimeSeriesSpec struct {
    // Type of statistics object that the request is operating on. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    ObjType TimeSeriesSpec_ObjType
    // Pod Identifier for queries on an individual pod. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    Pod *TimeSeriesPodIdentifier
    // Namespace name for queries for a namespace. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    Namespace *string
    // Cluster identifier for queries for a cluster. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    Cluster *string
    // UNIX timestamp value indicating when the requested series of statistical samples should begin. https://en.wikipedia.org/wiki/Unix_time. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    Start int64
    // UNIX timestamp value indicating when the requested series of statistical samples should end. https://en.wikipedia.org/wiki/Unix_time. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    End int64
}



//
    
    // Type of statistics object that this request is operating on. **Warning:** This enumeration is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    //
    // <p> See {@link com.vmware.vapi.bindings.ApiEnumeration enumerated types description}.
     
    type TimeSeriesSpec_ObjType string

    const (
        // The CLUSTER object type is used when specifying a vSphere cluster. **Warning:** This constant field is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
         TimeSeriesSpec_ObjType_CLUSTER TimeSeriesSpec_ObjType = "CLUSTER"
        // The NAMESPACE object type is used to specify a namespace. **Warning:** This constant field is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
         TimeSeriesSpec_ObjType_NAMESPACE TimeSeriesSpec_ObjType = "NAMESPACE"
        // The POD object type is used to specify an individual pod within a namespace. **Warning:** This constant field is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
         TimeSeriesSpec_ObjType_POD TimeSeriesSpec_ObjType = "POD"
    )

    func (o TimeSeriesSpec_ObjType) TimeSeriesSpec_ObjType() bool {
        switch o {
            case TimeSeriesSpec_ObjType_CLUSTER:
                return true
            case TimeSeriesSpec_ObjType_NAMESPACE:
                return true
            case TimeSeriesSpec_ObjType_POD:
                return true
            default:
                return false
        }
    }






func timeSeriesGetInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["spec"] = bindings.NewReferenceType(TimeSeriesSpecBindingType)
    fieldNameMap["spec"] = "Spec"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func timeSeriesGetOutputType() bindings.BindingType {
    return bindings.NewListType(bindings.NewReferenceType(TimeSeriesTimeSeriesBindingType), reflect.TypeOf([]TimeSeriesTimeSeries{}))
}

func timeSeriesGetRestMetadata() protocol.OperationRestMetadata {
    paramsTypeMap := map[string]bindings.BindingType{}
    pathParams := map[string]string{}
    queryParams := map[string]string{}
    headerParams := map[string]string{}
    paramsTypeMap["spec.cluster"] = bindings.NewOptionalType(bindings.NewIdType([]string {"ClusterComputeResource"}, ""))
    paramsTypeMap["spec.obj_type"] = bindings.NewEnumType("com.vmware.vcenter.namespace_management.stats.time_series.spec.obj_type", reflect.TypeOf(TimeSeriesSpec_ObjType(TimeSeriesSpec_ObjType_CLUSTER)))
    paramsTypeMap["spec.start"] = bindings.NewIntegerType()
    paramsTypeMap["spec.end"] = bindings.NewIntegerType()
    paramsTypeMap["spec.pod.namespace"] = bindings.NewIdType([]string {"com.vmware.vcenter.namespaces.Instance"}, "")
    paramsTypeMap["spec.namespace"] = bindings.NewOptionalType(bindings.NewIdType([]string {"com.vmware.vcenter.namespaces.Instance"}, ""))
    paramsTypeMap["spec.pod.pod_name"] = bindings.NewStringType()
    queryParams["spec.cluster"] = "cluster"
    queryParams["spec.pod.namespace"] = "pod-namespace"
    queryParams["spec.namespace"] = "namespace"
    queryParams["spec.start"] = "start"
    queryParams["spec.obj_type"] = "obj_type"
    queryParams["spec.end"] = "end"
    queryParams["spec.pod.pod_name"] = "pod_name"
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
      "/vcenter/namespace-management/stats/time-series",
       resultHeaders,
       200,
       errorHeaders,
       map[string]int{"Error": 500,"InvalidArgument": 400,"NotFound": 404,"Unsupported": 400,"Unauthenticated": 401,"Unauthorized": 403})
}



func TimeSeriesTimeSeriesBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["counter"] = bindings.NewStringType()
    fieldNameMap["counter"] = "Counter"
    fields["time_stamps"] = bindings.NewListType(bindings.NewIntegerType(), reflect.TypeOf([]int64{}))
    fieldNameMap["time_stamps"] = "TimeStamps"
    fields["values"] = bindings.NewListType(bindings.NewIntegerType(), reflect.TypeOf([]int64{}))
    fieldNameMap["values"] = "Values"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.namespace_management.stats.time_series.time_series",fields, reflect.TypeOf(TimeSeriesTimeSeries{}), fieldNameMap, validators)
}

func TimeSeriesPodIdentifierBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["namespace"] = bindings.NewIdType([]string {"com.vmware.vcenter.namespaces.Instance"}, "")
    fieldNameMap["namespace"] = "Namespace"
    fields["pod_name"] = bindings.NewStringType()
    fieldNameMap["pod_name"] = "PodName"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.namespace_management.stats.time_series.pod_identifier",fields, reflect.TypeOf(TimeSeriesPodIdentifier{}), fieldNameMap, validators)
}

func TimeSeriesSpecBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["obj_type"] = bindings.NewEnumType("com.vmware.vcenter.namespace_management.stats.time_series.spec.obj_type", reflect.TypeOf(TimeSeriesSpec_ObjType(TimeSeriesSpec_ObjType_CLUSTER)))
    fieldNameMap["obj_type"] = "ObjType"
    fields["pod"] = bindings.NewOptionalType(bindings.NewReferenceType(TimeSeriesPodIdentifierBindingType))
    fieldNameMap["pod"] = "Pod"
    fields["namespace"] = bindings.NewOptionalType(bindings.NewIdType([]string {"com.vmware.vcenter.namespaces.Instance"}, ""))
    fieldNameMap["namespace"] = "Namespace"
    fields["cluster"] = bindings.NewOptionalType(bindings.NewIdType([]string {"ClusterComputeResource"}, ""))
    fieldNameMap["cluster"] = "Cluster"
    fields["start"] = bindings.NewIntegerType()
    fieldNameMap["start"] = "Start"
    fields["end"] = bindings.NewIntegerType()
    fieldNameMap["end"] = "End"
    var validators = []bindings.Validator{}
    uv1 := bindings.NewUnionValidator("obj_type",
        map[string][]bindings.FieldData {
            "POD": []bindings.FieldData {
                 bindings.NewFieldData("pod", true),
            },
            "NAMESPACE": []bindings.FieldData {
                 bindings.NewFieldData("namespace", true),
            },
            "CLUSTER": []bindings.FieldData {
                 bindings.NewFieldData("cluster", true),
            },
        },
    )
    validators = append(validators, uv1)
    return bindings.NewStructType("com.vmware.vcenter.namespace_management.stats.time_series.spec",fields, reflect.TypeOf(TimeSeriesSpec{}), fieldNameMap, validators)
}


