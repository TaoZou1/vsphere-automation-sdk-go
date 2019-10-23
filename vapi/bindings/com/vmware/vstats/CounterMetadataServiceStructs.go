/* Copyright © 2019 VMware, Inc. All Rights Reserved.
     SPDX-License-Identifier: BSD-2-Clause */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Data type definitions file for service: CounterMetadata.
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



// Counter metadata status.
//
// <p> See {@link com.vmware.vapi.bindings.ApiEnumeration enumerated types description}.
 
type CounterMetadata_CounterEditionStatus string

const (
    // The counter edition is current and is the default.
     CounterMetadata_CounterEditionStatus_DEFAULT CounterMetadata_CounterEditionStatus = "DEFAULT"
    // The counter edition is current. This implies a support commitment.
     CounterMetadata_CounterEditionStatus_CURRENT CounterMetadata_CounterEditionStatus = "CURRENT"
    // The counter edition is deprecated. It will be decommissioned rather soon.
     CounterMetadata_CounterEditionStatus_DEPRECATED CounterMetadata_CounterEditionStatus = "DEPRECATED"
    // The counter edition is experimental. Consumers shouldn't rely on it for the long haul.
     CounterMetadata_CounterEditionStatus_EXPERIMENTAL CounterMetadata_CounterEditionStatus = "EXPERIMENTAL"
    // The counter edition was removed.
     CounterMetadata_CounterEditionStatus_REMOVED CounterMetadata_CounterEditionStatus = "REMOVED"
)

func (c CounterMetadata_CounterEditionStatus) CounterMetadata_CounterEditionStatus() bool {
    switch c {
        case CounterMetadata_CounterEditionStatus_DEFAULT:
            return true
        case CounterMetadata_CounterEditionStatus_CURRENT:
            return true
        case CounterMetadata_CounterEditionStatus_DEPRECATED:
            return true
        case CounterMetadata_CounterEditionStatus_EXPERIMENTAL:
            return true
        case CounterMetadata_CounterEditionStatus_REMOVED:
            return true
        default:
            return false
    }
}




// Type of the sampled data.
//
// <p> See {@link com.vmware.vapi.bindings.ApiEnumeration enumerated types description}.
 
type CounterMetadata_SampleType string

const (
    // Raw samples. The value unprocessed as-is sampled.
     CounterMetadata_SampleType_RAW CounterMetadata_SampleType = "RAW"
    // Absolute value samples. Represents an actual value of the counter.
     CounterMetadata_SampleType_ABSOLUTE CounterMetadata_SampleType = "ABSOLUTE"
    // Fraction samples. Implies range from 0.00 to 1.00.
     CounterMetadata_SampleType_FRACTION CounterMetadata_SampleType = "FRACTION"
    // Rate samples. Represents a value that has been normalized over the time period.
     CounterMetadata_SampleType_RATE CounterMetadata_SampleType = "RATE"
    // Delta samples. Represents an amount of change for the counter between the current time-stamp and the last time-stamp when the counter was sampled.
     CounterMetadata_SampleType_DELTA CounterMetadata_SampleType = "DELTA"
    // Log(n) samples. A natural logarithm of the value.
     CounterMetadata_SampleType_LOGN CounterMetadata_SampleType = "LOGN"
)

func (s CounterMetadata_SampleType) CounterMetadata_SampleType() bool {
    switch s {
        case CounterMetadata_SampleType_RAW:
            return true
        case CounterMetadata_SampleType_ABSOLUTE:
            return true
        case CounterMetadata_SampleType_FRACTION:
            return true
        case CounterMetadata_SampleType_RATE:
            return true
        case CounterMetadata_SampleType_DELTA:
            return true
        case CounterMetadata_SampleType_LOGN:
            return true
        default:
            return false
    }
}




// Unit used by a metric.
//
// <p> See {@link com.vmware.vapi.bindings.ApiEnumeration enumerated types description}.
 
type CounterMetadata_MetricUnits string

const (
    // Percent
     CounterMetadata_MetricUnits_PERCENT CounterMetadata_MetricUnits = "PERCENT"
    // Number
     CounterMetadata_MetricUnits_NUMBER CounterMetadata_MetricUnits = "NUMBER"
    // Second
     CounterMetadata_MetricUnits_SECOND CounterMetadata_MetricUnits = "SECOND"
    // Hertz
     CounterMetadata_MetricUnits_HERTZ CounterMetadata_MetricUnits = "HERTZ"
    // Meter
     CounterMetadata_MetricUnits_METER CounterMetadata_MetricUnits = "METER"
    // Meters per second
     CounterMetadata_MetricUnits_METERSPERSECOND CounterMetadata_MetricUnits = "METERSPERSECOND"
    // Meters per second squared
     CounterMetadata_MetricUnits_METERSPERSECONDSQUARED CounterMetadata_MetricUnits = "METERSPERSECONDSQUARED"
    // Byte
     CounterMetadata_MetricUnits_BYTE CounterMetadata_MetricUnits = "BYTE"
    // Bit
     CounterMetadata_MetricUnits_BIT CounterMetadata_MetricUnits = "BIT"
    // Bytes per second
     CounterMetadata_MetricUnits_BYTESPERSECOND CounterMetadata_MetricUnits = "BYTESPERSECOND"
    // Bits per second
     CounterMetadata_MetricUnits_BITSPERSECOND CounterMetadata_MetricUnits = "BITSPERSECOND"
    // Kilogram
     CounterMetadata_MetricUnits_KILOGRAM CounterMetadata_MetricUnits = "KILOGRAM"
    // Gram
     CounterMetadata_MetricUnits_GRAM CounterMetadata_MetricUnits = "GRAM"
    // Celsius
     CounterMetadata_MetricUnits_CELSIUS CounterMetadata_MetricUnits = "CELSIUS"
    // Kelvin
     CounterMetadata_MetricUnits_KELVIN CounterMetadata_MetricUnits = "KELVIN"
    // Joule
     CounterMetadata_MetricUnits_JOULE CounterMetadata_MetricUnits = "JOULE"
    // Watt
     CounterMetadata_MetricUnits_WATT CounterMetadata_MetricUnits = "WATT"
    // Volt
     CounterMetadata_MetricUnits_VOLT CounterMetadata_MetricUnits = "VOLT"
    // Ampere
     CounterMetadata_MetricUnits_AMPERE CounterMetadata_MetricUnits = "AMPERE"
    // Volt Ampere
     CounterMetadata_MetricUnits_VOLTAMPERE CounterMetadata_MetricUnits = "VOLTAMPERE"
    // Candela
     CounterMetadata_MetricUnits_CANDELA CounterMetadata_MetricUnits = "CANDELA"
    // Mole
     CounterMetadata_MetricUnits_MOLE CounterMetadata_MetricUnits = "MOLE"
)

func (m CounterMetadata_MetricUnits) CounterMetadata_MetricUnits() bool {
    switch m {
        case CounterMetadata_MetricUnits_PERCENT:
            return true
        case CounterMetadata_MetricUnits_NUMBER:
            return true
        case CounterMetadata_MetricUnits_SECOND:
            return true
        case CounterMetadata_MetricUnits_HERTZ:
            return true
        case CounterMetadata_MetricUnits_METER:
            return true
        case CounterMetadata_MetricUnits_METERSPERSECOND:
            return true
        case CounterMetadata_MetricUnits_METERSPERSECONDSQUARED:
            return true
        case CounterMetadata_MetricUnits_BYTE:
            return true
        case CounterMetadata_MetricUnits_BIT:
            return true
        case CounterMetadata_MetricUnits_BYTESPERSECOND:
            return true
        case CounterMetadata_MetricUnits_BITSPERSECOND:
            return true
        case CounterMetadata_MetricUnits_KILOGRAM:
            return true
        case CounterMetadata_MetricUnits_GRAM:
            return true
        case CounterMetadata_MetricUnits_CELSIUS:
            return true
        case CounterMetadata_MetricUnits_KELVIN:
            return true
        case CounterMetadata_MetricUnits_JOULE:
            return true
        case CounterMetadata_MetricUnits_WATT:
            return true
        case CounterMetadata_MetricUnits_VOLT:
            return true
        case CounterMetadata_MetricUnits_AMPERE:
            return true
        case CounterMetadata_MetricUnits_VOLTAMPERE:
            return true
        case CounterMetadata_MetricUnits_CANDELA:
            return true
        case CounterMetadata_MetricUnits_MOLE:
            return true
        default:
            return false
    }
}




// Unit factor of measurement.
//
// <p> See {@link com.vmware.vapi.bindings.ApiEnumeration enumerated types description}.
 
type CounterMetadata_UnitsFactor string

const (
    // Yotta 10^24
     CounterMetadata_UnitsFactor_YOTTA CounterMetadata_UnitsFactor = "YOTTA"
    // Zetta 10^21
     CounterMetadata_UnitsFactor_ZETTA CounterMetadata_UnitsFactor = "ZETTA"
    // Exa 10^18
     CounterMetadata_UnitsFactor_EXA CounterMetadata_UnitsFactor = "EXA"
    // Peta 10^15
     CounterMetadata_UnitsFactor_PETA CounterMetadata_UnitsFactor = "PETA"
    // Tera 10^12
     CounterMetadata_UnitsFactor_TERA CounterMetadata_UnitsFactor = "TERA"
    // Giga 10^9
     CounterMetadata_UnitsFactor_GIGA CounterMetadata_UnitsFactor = "GIGA"
    // Mega 10^6
     CounterMetadata_UnitsFactor_MEGA CounterMetadata_UnitsFactor = "MEGA"
    // Kilo 10^3
     CounterMetadata_UnitsFactor_KILO CounterMetadata_UnitsFactor = "KILO"
    // Hecto 10^2
     CounterMetadata_UnitsFactor_HECTO CounterMetadata_UnitsFactor = "HECTO"
    // Deca 10
     CounterMetadata_UnitsFactor_DECA CounterMetadata_UnitsFactor = "DECA"
    // One
     CounterMetadata_UnitsFactor_ONE CounterMetadata_UnitsFactor = "ONE"
    // Deci 10^-1
     CounterMetadata_UnitsFactor_DECI CounterMetadata_UnitsFactor = "DECI"
    // Centi 10^-2
     CounterMetadata_UnitsFactor_CENTI CounterMetadata_UnitsFactor = "CENTI"
    // Milli 10^-3
     CounterMetadata_UnitsFactor_MILLI CounterMetadata_UnitsFactor = "MILLI"
    // Micro 10^-6
     CounterMetadata_UnitsFactor_MICRO CounterMetadata_UnitsFactor = "MICRO"
    // Nano 10^-9
     CounterMetadata_UnitsFactor_NANO CounterMetadata_UnitsFactor = "NANO"
    // Pico 10^-12
     CounterMetadata_UnitsFactor_PIKO CounterMetadata_UnitsFactor = "PIKO"
    // Femto 10^-15
     CounterMetadata_UnitsFactor_FEMTO CounterMetadata_UnitsFactor = "FEMTO"
    // Atto 10^-18
     CounterMetadata_UnitsFactor_ATTO CounterMetadata_UnitsFactor = "ATTO"
    // Zepto 10^-21
     CounterMetadata_UnitsFactor_ZEPTO CounterMetadata_UnitsFactor = "ZEPTO"
    // Yocto 10^-24
     CounterMetadata_UnitsFactor_YOCTO CounterMetadata_UnitsFactor = "YOCTO"
    // Yobi 2^80, 1024^8
     CounterMetadata_UnitsFactor_YOBI CounterMetadata_UnitsFactor = "YOBI"
    // Zebi 2^70, 1024^7
     CounterMetadata_UnitsFactor_ZEBI CounterMetadata_UnitsFactor = "ZEBI"
    // Exbi 2^60, 1024^6
     CounterMetadata_UnitsFactor_EXBI CounterMetadata_UnitsFactor = "EXBI"
    // Pebi 2^50, 1024^5
     CounterMetadata_UnitsFactor_PEBI CounterMetadata_UnitsFactor = "PEBI"
    // Tebi 2^40, 1024^4
     CounterMetadata_UnitsFactor_TEBI CounterMetadata_UnitsFactor = "TEBI"
    // Gibi 2^30, 1024^3
     CounterMetadata_UnitsFactor_GIBI CounterMetadata_UnitsFactor = "GIBI"
    // Mebi 2^20, 1024^2
     CounterMetadata_UnitsFactor_MEBI CounterMetadata_UnitsFactor = "MEBI"
    // Kibi 2^10, 1024
     CounterMetadata_UnitsFactor_KIBI CounterMetadata_UnitsFactor = "KIBI"
)

func (u CounterMetadata_UnitsFactor) CounterMetadata_UnitsFactor() bool {
    switch u {
        case CounterMetadata_UnitsFactor_YOTTA:
            return true
        case CounterMetadata_UnitsFactor_ZETTA:
            return true
        case CounterMetadata_UnitsFactor_EXA:
            return true
        case CounterMetadata_UnitsFactor_PETA:
            return true
        case CounterMetadata_UnitsFactor_TERA:
            return true
        case CounterMetadata_UnitsFactor_GIGA:
            return true
        case CounterMetadata_UnitsFactor_MEGA:
            return true
        case CounterMetadata_UnitsFactor_KILO:
            return true
        case CounterMetadata_UnitsFactor_HECTO:
            return true
        case CounterMetadata_UnitsFactor_DECA:
            return true
        case CounterMetadata_UnitsFactor_ONE:
            return true
        case CounterMetadata_UnitsFactor_DECI:
            return true
        case CounterMetadata_UnitsFactor_CENTI:
            return true
        case CounterMetadata_UnitsFactor_MILLI:
            return true
        case CounterMetadata_UnitsFactor_MICRO:
            return true
        case CounterMetadata_UnitsFactor_NANO:
            return true
        case CounterMetadata_UnitsFactor_PIKO:
            return true
        case CounterMetadata_UnitsFactor_FEMTO:
            return true
        case CounterMetadata_UnitsFactor_ATTO:
            return true
        case CounterMetadata_UnitsFactor_ZEPTO:
            return true
        case CounterMetadata_UnitsFactor_YOCTO:
            return true
        case CounterMetadata_UnitsFactor_YOBI:
            return true
        case CounterMetadata_UnitsFactor_ZEBI:
            return true
        case CounterMetadata_UnitsFactor_EXBI:
            return true
        case CounterMetadata_UnitsFactor_PEBI:
            return true
        case CounterMetadata_UnitsFactor_TEBI:
            return true
        case CounterMetadata_UnitsFactor_GIBI:
            return true
        case CounterMetadata_UnitsFactor_MEBI:
            return true
        case CounterMetadata_UnitsFactor_KIBI:
            return true
        default:
            return false
    }
}





// The ``Info`` class contains information about CounterMetadata. It represents edition of the Counter.
 type CounterMetadataInfo struct {
    // Counter Id.
    Cid string
    // Metadata Id.
    Mid string
    // Counter Edition status.
    Status CounterMetadata_CounterEditionStatus
    // Numeric properties of the sampled data.
    Type_ CounterMetadata_SampleType
    // The units of the measurement.
    Units CounterMetadata_MetricUnits
    // Additional multiplier factors to be used with units.
    Scale *CounterMetadata_UnitsFactor
    // Human legible localizable text about the counter.
    UserInfo *UserInfo
    // ID of the respective provider.
    Pid *string
}



func (CounterMetadataInfo CounterMetadataInfo) Error() string {
    return "com.vmware.vstats.info"
}



// The ``FilterSpec`` class is used to filter the counter metadata list.
 type CounterMetadataFilterSpec struct {
    // Counter edition status.
    Status *CounterMetadata_CounterEditionStatus
}



func (CounterMetadataFilterSpec CounterMetadataFilterSpec) Error() string {
    return "com.vmware.vstats.filter_spec"
}






func counterMetadataListInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["cid"] = bindings.NewIdType([]string {"com.vmware.vstats.model.Counter"}, "")
    fields["filter"] = bindings.NewOptionalType(bindings.NewReferenceType(CounterMetadataFilterSpecBindingType))
    fieldNameMap["cid"] = "Cid"
    fieldNameMap["filter"] = "Filter"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func counterMetadataListOutputType() bindings.BindingType {
    return bindings.NewListType(bindings.NewReferenceType(CounterMetadataInfoBindingType), reflect.TypeOf([]CounterMetadataInfo{}))
}

func counterMetadataListRestMetadata() protocol.OperationRestMetadata {
    paramsTypeMap := map[string]bindings.BindingType{}
    pathParams := map[string]string{}
    queryParams := map[string]string{}
    headerParams := map[string]string{}
    paramsTypeMap["filter.status"] = bindings.NewOptionalType(bindings.NewEnumType("com.vmware.vstats.counter_metadata.counter_edition_status", reflect.TypeOf(CounterMetadata_CounterEditionStatus(CounterMetadata_CounterEditionStatus_DEFAULT))))
    paramsTypeMap["cid"] = bindings.NewIdType([]string {"com.vmware.vstats.model.Counter"}, "")
    paramsTypeMap["cid"] = bindings.NewIdType([]string {"com.vmware.vstats.model.Counter"}, "")
    pathParams["cid"] = "cid"
    queryParams["filter.status"] = "status"
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
      "/stats/counters/{cid}/metadata",
       resultHeaders,
       200,
       errorHeaders,
       map[string]int{"Error": 500,"InvalidArgument": 400,"NotFound": 404,"Unauthenticated": 401,"Unauthorized": 403})
}


func counterMetadataGetDefaultInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["cid"] = bindings.NewIdType([]string {"com.vmware.vstats.model.Counter"}, "")
    fieldNameMap["cid"] = "Cid"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func counterMetadataGetDefaultOutputType() bindings.BindingType {
    return bindings.NewListType(bindings.NewReferenceType(CounterMetadataInfoBindingType), reflect.TypeOf([]CounterMetadataInfo{}))
}

func counterMetadataGetDefaultRestMetadata() protocol.OperationRestMetadata {
    paramsTypeMap := map[string]bindings.BindingType{}
    pathParams := map[string]string{}
    queryParams := map[string]string{}
    headerParams := map[string]string{}
    paramsTypeMap["cid"] = bindings.NewIdType([]string {"com.vmware.vstats.model.Counter"}, "")
    paramsTypeMap["cid"] = bindings.NewIdType([]string {"com.vmware.vstats.model.Counter"}, "")
    pathParams["cid"] = "cid"
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
      "/stats/counters/{cid}/metadata/default",
       resultHeaders,
       200,
       errorHeaders,
       map[string]int{"Error": 500,"InvalidArgument": 400,"NotFound": 404,"Unauthenticated": 401,"Unauthorized": 403})
}


func counterMetadataGetInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["cid"] = bindings.NewIdType([]string {"com.vmware.vstats.model.Counter"}, "")
    fields["mid"] = bindings.NewIdType([]string {"com.vmware.vstats.model.CounterMetadata"}, "")
    fieldNameMap["cid"] = "Cid"
    fieldNameMap["mid"] = "Mid"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func counterMetadataGetOutputType() bindings.BindingType {
    return bindings.NewReferenceType(CounterMetadataInfoBindingType)
}

func counterMetadataGetRestMetadata() protocol.OperationRestMetadata {
    paramsTypeMap := map[string]bindings.BindingType{}
    pathParams := map[string]string{}
    queryParams := map[string]string{}
    headerParams := map[string]string{}
    paramsTypeMap["mid"] = bindings.NewIdType([]string {"com.vmware.vstats.model.CounterMetadata"}, "")
    paramsTypeMap["cid"] = bindings.NewIdType([]string {"com.vmware.vstats.model.Counter"}, "")
    paramsTypeMap["cid"] = bindings.NewIdType([]string {"com.vmware.vstats.model.Counter"}, "")
    paramsTypeMap["mid"] = bindings.NewIdType([]string {"com.vmware.vstats.model.CounterMetadata"}, "")
    pathParams["mid"] = "mid"
    pathParams["cid"] = "cid"
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
      "/stats/counters/{cid}/metadata/{mid}",
       resultHeaders,
       200,
       errorHeaders,
       map[string]int{"Error": 500,"InvalidArgument": 400,"NotFound": 404,"Unauthenticated": 401,"Unauthorized": 403})
}



func CounterMetadataInfoBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["cid"] = bindings.NewIdType([]string {"com.vmware.vstats.model.Counter"}, "")
    fieldNameMap["cid"] = "Cid"
    fields["mid"] = bindings.NewIdType([]string {"com.vmware.vstats.model.CounterMetadata"}, "")
    fieldNameMap["mid"] = "Mid"
    fields["status"] = bindings.NewEnumType("com.vmware.vstats.counter_metadata.counter_edition_status", reflect.TypeOf(CounterMetadata_CounterEditionStatus(CounterMetadata_CounterEditionStatus_DEFAULT)))
    fieldNameMap["status"] = "Status"
    fields["type"] = bindings.NewEnumType("com.vmware.vstats.counter_metadata.sample_type", reflect.TypeOf(CounterMetadata_SampleType(CounterMetadata_SampleType_RAW)))
    fieldNameMap["type"] = "Type_"
    fields["units"] = bindings.NewEnumType("com.vmware.vstats.counter_metadata.metric_units", reflect.TypeOf(CounterMetadata_MetricUnits(CounterMetadata_MetricUnits_PERCENT)))
    fieldNameMap["units"] = "Units"
    fields["scale"] = bindings.NewOptionalType(bindings.NewEnumType("com.vmware.vstats.counter_metadata.units_factor", reflect.TypeOf(CounterMetadata_UnitsFactor(CounterMetadata_UnitsFactor_YOTTA))))
    fieldNameMap["scale"] = "Scale"
    fields["user_info"] = bindings.NewOptionalType(bindings.NewReferenceType(UserInfoBindingType))
    fieldNameMap["user_info"] = "UserInfo"
    fields["pid"] = bindings.NewOptionalType(bindings.NewIdType([]string {"com.vmware.vstats.model.Provider"}, ""))
    fieldNameMap["pid"] = "Pid"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vstats.counter_metadata.info",fields, reflect.TypeOf(CounterMetadataInfo{}), fieldNameMap, validators)
}

func CounterMetadataFilterSpecBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["status"] = bindings.NewOptionalType(bindings.NewEnumType("com.vmware.vstats.counter_metadata.counter_edition_status", reflect.TypeOf(CounterMetadata_CounterEditionStatus(CounterMetadata_CounterEditionStatus_DEFAULT))))
    fieldNameMap["status"] = "Status"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vstats.counter_metadata.filter_spec",fields, reflect.TypeOf(CounterMetadataFilterSpec{}), fieldNameMap, validators)
}


