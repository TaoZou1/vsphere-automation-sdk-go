/* Copyright © 2019 VMware, Inc. All Rights Reserved.
     SPDX-License-Identifier: BSD-2-Clause */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Data type definitions file for service: Telemetry.
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



// Type of telemetry object.
//
// <p> See {@link com.vmware.vapi.bindings.ApiEnumeration enumerated types description}.
 
type TelemetryTelemetryType string

const (
    // Filter for telemetry counters.
     TelemetryTelemetryType_COUNTERS TelemetryTelemetryType = "COUNTERS"
    // Filter for telemetry timers.
     TelemetryTelemetryType_TIMERS TelemetryTelemetryType = "TIMERS"
    // Filter for telemetry gauges.
     TelemetryTelemetryType_GAUGES TelemetryTelemetryType = "GAUGES"
    // Filter for telemetry meters.
     TelemetryTelemetryType_METERS TelemetryTelemetryType = "METERS"
    // Filter for telemetry counters.
     TelemetryTelemetryType_HISTOGRAMS TelemetryTelemetryType = "HISTOGRAMS"
)

func (t TelemetryTelemetryType) TelemetryTelemetryType() bool {
    switch t {
        case TelemetryTelemetryType_COUNTERS:
            return true
        case TelemetryTelemetryType_TIMERS:
            return true
        case TelemetryTelemetryType_GAUGES:
            return true
        case TelemetryTelemetryType_METERS:
            return true
        case TelemetryTelemetryType_HISTOGRAMS:
            return true
        default:
            return false
    }
}






// The ``TelemetryCounter`` class specifies vStats monotonically increasing integer.
 type TelemetryTelemetryCounter struct {
    // The name of this telemetry-counter.
    Name string
    // The current count of this telemetry-counter.
    Count int64
}







// The ``TelemetryGauge`` class specifies vStats telemetry-gauge, both integer and floating point are presented here.
 type TelemetryTelemetryGauge struct {
    // The name of this telemetry-gauge.
    Name string
    // The current value of this gauge.
    Value float64
}







// The ``TelemetryHistogram`` class specifies vStats telemetry-histogram.
 type TelemetryTelemetryHistogram struct {
    // The name of this telemetry-histogram.
    Name string
    // Current count for this histogram.
    Count int64
    // Current min value.
    Min int64
    // Current max value.
    Max int64
    // Current mean value.
    Mean float64
    // Current standard deviation.
    Stddev float64
    // Median of all recorded samples.
    Median float64
    // 75th percentile of all recorded samples.
    P75 float64
    // 95th percentile of all recorded samples.
    P95 float64
    // 99th percentile of all recoeded samples.
    P99 float64
    // 99.9th percentile of all recorded samples.
    P999 float64
}







// The ``TelemetryMeter`` class specifies vStats telemetry-meter.
 type TelemetryTelemetryMeter struct {
    // The name of this telemetry-meter.
    Name string
    // Current count for this meter.
    Count int64
    // 1-minute rate.
    Min1Rate float64
    // 5-minute rate.
    Min5Rate float64
    // 15-minute rate.
    Min15Rate float64
    // Current mean rate.
    MeanRate float64
}







// The ``TelemetryTimer`` class specifies vStats telemetry-timer. Units are in nanoseconds.
 type TelemetryTelemetryTimer struct {
    // The name of this telemetry-timer.
    Name string
    // Current count for this timer.
    Count int64
    // Current mins value.
    Min float64
    // Current max value.
    Max float64
    // Current mean value.
    Mean float64
    // Current standard deviation.
    Stddev float64
    // Current median.
    Median float64
    // Current 75th percentile.
    P75 float64
    // Current 95th percentile.
    P95 float64
    // Current 99th percentile.
    P99 float64
    // Current 99.9th percentile.
    P999 float64
    // 1-minute rate.
    Min1Rate float64
    // 5-minute rate.
    Min5Rate float64
    // 15-minute rate.
    Min15Rate float64
    // Current mean rate.
    MeanRate float64
}







// The ``Info`` class contains overall telemetry Information.
 type TelemetryInfo struct {
    // List of TelemetryCounter.
    Counters []TelemetryTelemetryCounter
    // List of TelemetryTimer.
    Timers []TelemetryTelemetryTimer
    // List of TelemetryGauge.
    Gauges []TelemetryTelemetryGauge
    // List of TelemetryMeter.
    Meters []TelemetryTelemetryMeter
    // List of TelemetryHistogram.
    Histograms []TelemetryTelemetryHistogram
}







// ``FilterSpec`` class describes filter criteria for telemetry.
 type TelemetryFilterSpec struct {
    // To query telemetry by name prefix.
    Prefix *string
    // Specifies metric types to return.
    Types map[TelemetryTelemetryType]bool
}









func telemetryGetInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["filter"] = bindings.NewOptionalType(bindings.NewReferenceType(TelemetryFilterSpecBindingType))
    fieldNameMap["filter"] = "Filter"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func telemetryGetOutputType() bindings.BindingType {
    return bindings.NewReferenceType(TelemetryInfoBindingType)
}

func telemetryGetRestMetadata() protocol.OperationRestMetadata {
    paramsTypeMap := map[string]bindings.BindingType{}
    pathParams := map[string]string{}
    queryParams := map[string]string{}
    headerParams := map[string]string{}
    paramsTypeMap["filter.prefix"] = bindings.NewOptionalType(bindings.NewStringType())
    paramsTypeMap["filter.types"] = bindings.NewOptionalType(bindings.NewSetType(bindings.NewEnumType("com.vmware.vstats.telemetry.telemetry_type", reflect.TypeOf(TelemetryTelemetryType(TelemetryTelemetryType_COUNTERS))), reflect.TypeOf(map[TelemetryTelemetryType]bool{})))
    queryParams["filter.types"] = "types"
    queryParams["filter.prefix"] = "prefix"
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
      "/stats/telemetry",
       resultHeaders,
       200,
       errorHeaders,
       map[string]int{"Error": 500,"NotFound": 404,"Unauthenticated": 401,"Unauthorized": 403})
}



func TelemetryTelemetryCounterBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["name"] = bindings.NewStringType()
    fieldNameMap["name"] = "Name"
    fields["count"] = bindings.NewIntegerType()
    fieldNameMap["count"] = "Count"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vstats.telemetry.telemetry_counter",fields, reflect.TypeOf(TelemetryTelemetryCounter{}), fieldNameMap, validators)
}

func TelemetryTelemetryGaugeBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["name"] = bindings.NewStringType()
    fieldNameMap["name"] = "Name"
    fields["value"] = bindings.NewDoubleType()
    fieldNameMap["value"] = "Value"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vstats.telemetry.telemetry_gauge",fields, reflect.TypeOf(TelemetryTelemetryGauge{}), fieldNameMap, validators)
}

func TelemetryTelemetryHistogramBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["name"] = bindings.NewStringType()
    fieldNameMap["name"] = "Name"
    fields["count"] = bindings.NewIntegerType()
    fieldNameMap["count"] = "Count"
    fields["min"] = bindings.NewIntegerType()
    fieldNameMap["min"] = "Min"
    fields["max"] = bindings.NewIntegerType()
    fieldNameMap["max"] = "Max"
    fields["mean"] = bindings.NewDoubleType()
    fieldNameMap["mean"] = "Mean"
    fields["stddev"] = bindings.NewDoubleType()
    fieldNameMap["stddev"] = "Stddev"
    fields["median"] = bindings.NewDoubleType()
    fieldNameMap["median"] = "Median"
    fields["p75"] = bindings.NewDoubleType()
    fieldNameMap["p75"] = "P75"
    fields["p95"] = bindings.NewDoubleType()
    fieldNameMap["p95"] = "P95"
    fields["p99"] = bindings.NewDoubleType()
    fieldNameMap["p99"] = "P99"
    fields["p999"] = bindings.NewDoubleType()
    fieldNameMap["p999"] = "P999"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vstats.telemetry.telemetry_histogram",fields, reflect.TypeOf(TelemetryTelemetryHistogram{}), fieldNameMap, validators)
}

func TelemetryTelemetryMeterBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["name"] = bindings.NewStringType()
    fieldNameMap["name"] = "Name"
    fields["count"] = bindings.NewIntegerType()
    fieldNameMap["count"] = "Count"
    fields["min1_rate"] = bindings.NewDoubleType()
    fieldNameMap["min1_rate"] = "Min1Rate"
    fields["min5_rate"] = bindings.NewDoubleType()
    fieldNameMap["min5_rate"] = "Min5Rate"
    fields["min15_rate"] = bindings.NewDoubleType()
    fieldNameMap["min15_rate"] = "Min15Rate"
    fields["mean_rate"] = bindings.NewDoubleType()
    fieldNameMap["mean_rate"] = "MeanRate"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vstats.telemetry.telemetry_meter",fields, reflect.TypeOf(TelemetryTelemetryMeter{}), fieldNameMap, validators)
}

func TelemetryTelemetryTimerBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["name"] = bindings.NewStringType()
    fieldNameMap["name"] = "Name"
    fields["count"] = bindings.NewIntegerType()
    fieldNameMap["count"] = "Count"
    fields["min"] = bindings.NewDoubleType()
    fieldNameMap["min"] = "Min"
    fields["max"] = bindings.NewDoubleType()
    fieldNameMap["max"] = "Max"
    fields["mean"] = bindings.NewDoubleType()
    fieldNameMap["mean"] = "Mean"
    fields["stddev"] = bindings.NewDoubleType()
    fieldNameMap["stddev"] = "Stddev"
    fields["median"] = bindings.NewDoubleType()
    fieldNameMap["median"] = "Median"
    fields["p75"] = bindings.NewDoubleType()
    fieldNameMap["p75"] = "P75"
    fields["p95"] = bindings.NewDoubleType()
    fieldNameMap["p95"] = "P95"
    fields["p99"] = bindings.NewDoubleType()
    fieldNameMap["p99"] = "P99"
    fields["p999"] = bindings.NewDoubleType()
    fieldNameMap["p999"] = "P999"
    fields["min1_rate"] = bindings.NewDoubleType()
    fieldNameMap["min1_rate"] = "Min1Rate"
    fields["min5_rate"] = bindings.NewDoubleType()
    fieldNameMap["min5_rate"] = "Min5Rate"
    fields["min15_rate"] = bindings.NewDoubleType()
    fieldNameMap["min15_rate"] = "Min15Rate"
    fields["mean_rate"] = bindings.NewDoubleType()
    fieldNameMap["mean_rate"] = "MeanRate"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vstats.telemetry.telemetry_timer",fields, reflect.TypeOf(TelemetryTelemetryTimer{}), fieldNameMap, validators)
}

func TelemetryInfoBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["counters"] = bindings.NewListType(bindings.NewReferenceType(TelemetryTelemetryCounterBindingType), reflect.TypeOf([]TelemetryTelemetryCounter{}))
    fieldNameMap["counters"] = "Counters"
    fields["timers"] = bindings.NewListType(bindings.NewReferenceType(TelemetryTelemetryTimerBindingType), reflect.TypeOf([]TelemetryTelemetryTimer{}))
    fieldNameMap["timers"] = "Timers"
    fields["gauges"] = bindings.NewListType(bindings.NewReferenceType(TelemetryTelemetryGaugeBindingType), reflect.TypeOf([]TelemetryTelemetryGauge{}))
    fieldNameMap["gauges"] = "Gauges"
    fields["meters"] = bindings.NewListType(bindings.NewReferenceType(TelemetryTelemetryMeterBindingType), reflect.TypeOf([]TelemetryTelemetryMeter{}))
    fieldNameMap["meters"] = "Meters"
    fields["histograms"] = bindings.NewListType(bindings.NewReferenceType(TelemetryTelemetryHistogramBindingType), reflect.TypeOf([]TelemetryTelemetryHistogram{}))
    fieldNameMap["histograms"] = "Histograms"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vstats.telemetry.info",fields, reflect.TypeOf(TelemetryInfo{}), fieldNameMap, validators)
}

func TelemetryFilterSpecBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["prefix"] = bindings.NewOptionalType(bindings.NewStringType())
    fieldNameMap["prefix"] = "Prefix"
    fields["types"] = bindings.NewOptionalType(bindings.NewSetType(bindings.NewEnumType("com.vmware.vstats.telemetry.telemetry_type", reflect.TypeOf(TelemetryTelemetryType(TelemetryTelemetryType_COUNTERS))), reflect.TypeOf(map[TelemetryTelemetryType]bool{})))
    fieldNameMap["types"] = "Types"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vstats.telemetry.filter_spec",fields, reflect.TypeOf(TelemetryFilterSpec{}), fieldNameMap, validators)
}


