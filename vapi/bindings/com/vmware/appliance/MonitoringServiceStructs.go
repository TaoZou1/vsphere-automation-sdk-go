/* Copyright © 2019 VMware, Inc. All Rights Reserved.
     SPDX-License-Identifier: BSD-2-Clause */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Data type definitions file for service: Monitoring.
 * Includes binding types of a structures and enumerations defined in the service.
 * Shared by client-side stubs and server-side skeletons to ensure type
 * compatibility.
 */

package appliance

import (
    "reflect"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/bindings"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/data"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/protocol"
    "time"
)



// ``FunctionType`` enumeration class Defines aggregation function
//
// <p> See {@link com.vmware.vapi.bindings.ApiEnumeration enumerated types description}.
 
type Monitoring_FunctionType string

const (
    // Aggregation takes count per period (sum)
     Monitoring_FunctionType_COUNT Monitoring_FunctionType = "COUNT"
    // Aggregation takes maximums per period
     Monitoring_FunctionType_MAX Monitoring_FunctionType = "MAX"
    // Aggregation takes average per period
     Monitoring_FunctionType_AVG Monitoring_FunctionType = "AVG"
    // Aggregation takes minimums per period
     Monitoring_FunctionType_MIN Monitoring_FunctionType = "MIN"
)

func (f Monitoring_FunctionType) Monitoring_FunctionType() bool {
    switch f {
        case Monitoring_FunctionType_COUNT:
            return true
        case Monitoring_FunctionType_MAX:
            return true
        case Monitoring_FunctionType_AVG:
            return true
        case Monitoring_FunctionType_MIN:
            return true
        default:
            return false
    }
}




// ``IntervalType`` enumeration class Defines interval between the values in hours and mins, for which aggregation will apply
//
// <p> See {@link com.vmware.vapi.bindings.ApiEnumeration enumerated types description}.
 
type Monitoring_IntervalType string

const (
    // Thirty minutes interval between values. One week is 336 values.
     Monitoring_IntervalType_MINUTES30 Monitoring_IntervalType = "MINUTES30"
    // Two hours interval between values. One month has 360 values.
     Monitoring_IntervalType_HOURS2 Monitoring_IntervalType = "HOURS2"
    // Five minutes interval between values (finest). One day would have 288 values, one week is 2016.
     Monitoring_IntervalType_MINUTES5 Monitoring_IntervalType = "MINUTES5"
    // 24 hours interval between values. One year has 365 values.
     Monitoring_IntervalType_DAY1 Monitoring_IntervalType = "DAY1"
    // Six hour interval between values. One quarter is 360 values.
     Monitoring_IntervalType_HOURS6 Monitoring_IntervalType = "HOURS6"
)

func (i Monitoring_IntervalType) Monitoring_IntervalType() bool {
    switch i {
        case Monitoring_IntervalType_MINUTES30:
            return true
        case Monitoring_IntervalType_HOURS2:
            return true
        case Monitoring_IntervalType_MINUTES5:
            return true
        case Monitoring_IntervalType_DAY1:
            return true
        case Monitoring_IntervalType_HOURS6:
            return true
        default:
            return false
    }
}





// ``MonitoredItemData`` class Structure representing monitored item data.
 type MonitoringMonitoredItemData struct {
    // Monitored item IDs Ex: CPU, MEMORY, STORAGE_TOTAL
    Name string
    // interval between values in hours, minutes
    Interval Monitoring_IntervalType
    // aggregation function
    Function Monitoring_FunctionType
    // Start time in UTC
    StartTime time.Time
    // End time in UTC
    EndTime time.Time
    // list of values
    Data []string
}


// TODO some error doesn't have ErrorType. Resolve this later VAPI-3009
//


// ``MonitoredItemDataRequest`` class Structure representing requested monitored item data.
 type MonitoringMonitoredItemDataRequest struct {
    // monitored item IDs Ex: CPU, MEMORY
    Names []string
    // interval between values in hours, minutes
    Interval Monitoring_IntervalType
    // aggregation function
    Function Monitoring_FunctionType
    // Start time in UTC
    StartTime time.Time
    // End time in UTC
    EndTime time.Time
}


// TODO some error doesn't have ErrorType. Resolve this later VAPI-3009
//


// ``MonitoredItem`` class Structure representing requested monitored item data.
 type MonitoringMonitoredItem struct {
    // monitored item ID Ex: CPU, MEMORY
    Id string
    // monitored item name Ex: "Network write speed"
    Name string
    // Y-axis label EX: "Mbps", "%"
    Units string
    // category Ex: network, storage etc
    Category string
    // instance name Ex: eth0
    Instance string
    // monitored item description Ex: com.vmware.applmgmt.mon.descr.net.rx.packetRate.eth0
    Description string
}


// TODO some error doesn't have ErrorType. Resolve this later VAPI-3009
//





func monitoringQueryInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["item"] = bindings.NewReferenceType(MonitoringMonitoredItemDataRequestBindingType)
    fieldNameMap["item"] = "Item"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func monitoringQueryOutputType() bindings.BindingType {
    return bindings.NewListType(bindings.NewReferenceType(MonitoringMonitoredItemDataBindingType), reflect.TypeOf([]MonitoringMonitoredItemData{}))
}

func monitoringQueryRestMetadata() protocol.OperationRestMetadata {
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
       map[string]int{"Error": 500})
}


func monitoringListInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func monitoringListOutputType() bindings.BindingType {
    return bindings.NewListType(bindings.NewReferenceType(MonitoringMonitoredItemBindingType), reflect.TypeOf([]MonitoringMonitoredItem{}))
}

func monitoringListRestMetadata() protocol.OperationRestMetadata {
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
       map[string]int{"Error": 500})
}


func monitoringGetInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["stat_id"] = bindings.NewIdType([]string {"com.vmware.appliance.monitoring"}, "")
    fieldNameMap["stat_id"] = "StatId"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func monitoringGetOutputType() bindings.BindingType {
    return bindings.NewReferenceType(MonitoringMonitoredItemBindingType)
}

func monitoringGetRestMetadata() protocol.OperationRestMetadata {
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
       map[string]int{"Error": 500})
}



func MonitoringMonitoredItemDataBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["name"] = bindings.NewIdType([]string {"com.vmware.appliance.monitoring"}, "")
    fieldNameMap["name"] = "Name"
    fields["interval"] = bindings.NewEnumType("com.vmware.appliance.monitoring.interval_type", reflect.TypeOf(Monitoring_IntervalType(Monitoring_IntervalType_MINUTES30)))
    fieldNameMap["interval"] = "Interval"
    fields["function"] = bindings.NewEnumType("com.vmware.appliance.monitoring.function_type", reflect.TypeOf(Monitoring_FunctionType(Monitoring_FunctionType_COUNT)))
    fieldNameMap["function"] = "Function"
    fields["start_time"] = bindings.NewDateTimeType()
    fieldNameMap["start_time"] = "StartTime"
    fields["end_time"] = bindings.NewDateTimeType()
    fieldNameMap["end_time"] = "EndTime"
    fields["data"] = bindings.NewListType(bindings.NewStringType(), reflect.TypeOf([]string{}))
    fieldNameMap["data"] = "Data"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.appliance.monitoring.monitored_item_data",fields, reflect.TypeOf(MonitoringMonitoredItemData{}), fieldNameMap, validators)
}

func MonitoringMonitoredItemDataRequestBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["names"] = bindings.NewListType(bindings.NewIdType([]string {"com.vmware.appliance.monitoring"}, ""), reflect.TypeOf([]string{}))
    fieldNameMap["names"] = "Names"
    fields["interval"] = bindings.NewEnumType("com.vmware.appliance.monitoring.interval_type", reflect.TypeOf(Monitoring_IntervalType(Monitoring_IntervalType_MINUTES30)))
    fieldNameMap["interval"] = "Interval"
    fields["function"] = bindings.NewEnumType("com.vmware.appliance.monitoring.function_type", reflect.TypeOf(Monitoring_FunctionType(Monitoring_FunctionType_COUNT)))
    fieldNameMap["function"] = "Function"
    fields["start_time"] = bindings.NewDateTimeType()
    fieldNameMap["start_time"] = "StartTime"
    fields["end_time"] = bindings.NewDateTimeType()
    fieldNameMap["end_time"] = "EndTime"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.appliance.monitoring.monitored_item_data_request",fields, reflect.TypeOf(MonitoringMonitoredItemDataRequest{}), fieldNameMap, validators)
}

func MonitoringMonitoredItemBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["id"] = bindings.NewIdType([]string {"com.vmware.appliance.monitoring"}, "")
    fieldNameMap["id"] = "Id"
    fields["name"] = bindings.NewStringType()
    fieldNameMap["name"] = "Name"
    fields["units"] = bindings.NewStringType()
    fieldNameMap["units"] = "Units"
    fields["category"] = bindings.NewStringType()
    fieldNameMap["category"] = "Category"
    fields["instance"] = bindings.NewStringType()
    fieldNameMap["instance"] = "Instance"
    fields["description"] = bindings.NewStringType()
    fieldNameMap["description"] = "Description"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.appliance.monitoring.monitored_item",fields, reflect.TypeOf(MonitoringMonitoredItem{}), fieldNameMap, validators)
}


