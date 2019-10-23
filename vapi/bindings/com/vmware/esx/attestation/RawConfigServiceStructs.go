/* Copyright © 2019 VMware, Inc. All Rights Reserved.
     SPDX-License-Identifier: BSD-2-Clause */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Data type definitions file for service: RawConfig.
 * Includes binding types of a structures and enumerations defined in the service.
 * Shared by client-side stubs and server-side skeletons to ensure type
 * compatibility.
 */

package attestation

import (
    "reflect"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/bindings"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/data"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/protocol"
    "time"
)



// The ``InfoType`` enumeration class describes the type of information that is populated in RawConfigInfo when returned by RawConfig#get.
//
// <p> See {@link com.vmware.vapi.bindings.ApiEnumeration enumerated types description}.
 
type RawConfig_InfoType string

const (
    // All properties in RawConfigInfo are populated.
     RawConfig_InfoType_FULL RawConfig_InfoType = "FULL"
    // The RawConfigInfo#configuration property is not populated. 
    //
    //  The RawConfigInfo#configuration property may be a large document. Specifying a RawConfig_InfoType#RawConfigInfoType_BRIEF request will reduce the size of the RawConfigInfo class on common operations.
     RawConfig_InfoType_BRIEF RawConfig_InfoType = "BRIEF"
)

func (i RawConfig_InfoType) RawConfig_InfoType() bool {
    switch i {
        case RawConfig_InfoType_FULL:
            return true
        case RawConfig_InfoType_BRIEF:
            return true
        default:
            return false
    }
}





// The ``Info`` class contains information that describes the persistent user configuration.
 type RawConfigInfo struct {
    // An opaque configuration blob that represents the entire service persistent user configuration. 
    //
    //  This configuration blob can be provided to RawConfig#set in order to either restore a previous configuration, or to replicate configuration across instances in a cluster. 
    //
    //  Note that the configuration does not include private signing keys that are used to sign attestation reports. The signing keys cannot be replicated.
    Configuration []byte
    // A configuration digest which can be used to uniquely identify the service configuration. 
    //
    //  Two service instances can be assumed to have the same configuration if their RawConfigInfo#fingerprint fields match.
    Fingerprint string
    // The configuration last update time.
    LastUpdateTime time.Time
}



func (RawConfigInfo RawConfigInfo) Error() string {
    return "com.vmware.esx.attestation.info"
}



// The ``SetSpec`` class contains information that describes a new configuration to applied using RawConfig#set.
 type RawConfigSetSpec struct {
    // An opaque configuration blob that represents the entire service persistent user configuration.
    Configuration []byte
}



func (RawConfigSetSpec RawConfigSetSpec) Error() string {
    return "com.vmware.esx.attestation.set_spec"
}






func rawConfigGetInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["type"] = bindings.NewOptionalType(bindings.NewEnumType("com.vmware.esx.attestation.raw_config.info_type", reflect.TypeOf(RawConfig_InfoType(RawConfig_InfoType_FULL))))
    fieldNameMap["type"] = "Type_"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func rawConfigGetOutputType() bindings.BindingType {
    return bindings.NewReferenceType(RawConfigInfoBindingType)
}

func rawConfigGetRestMetadata() protocol.OperationRestMetadata {
    paramsTypeMap := map[string]bindings.BindingType{}
    pathParams := map[string]string{}
    queryParams := map[string]string{}
    headerParams := map[string]string{}
    paramsTypeMap["type"] = bindings.NewOptionalType(bindings.NewEnumType("com.vmware.esx.attestation.raw_config.info_type", reflect.TypeOf(RawConfig_InfoType(RawConfig_InfoType_FULL))))
    queryParams["type"] = "type"
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
      "/esx/attestation/raw-config",
       resultHeaders,
       200,
       errorHeaders,
       map[string]int{"Error": 500,"InvalidArgument": 400,"Unauthenticated": 401,"Unauthorized": 403})
}


func rawConfigSetInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["spec"] = bindings.NewReferenceType(RawConfigSetSpecBindingType)
    fieldNameMap["spec"] = "Spec"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func rawConfigSetOutputType() bindings.BindingType {
    return bindings.NewVoidType()
}

func rawConfigSetRestMetadata() protocol.OperationRestMetadata {
    paramsTypeMap := map[string]bindings.BindingType{}
    pathParams := map[string]string{}
    queryParams := map[string]string{}
    headerParams := map[string]string{}
    paramsTypeMap["spec"] = bindings.NewReferenceType(RawConfigSetSpecBindingType)
    resultHeaders := map[string]string{}
    errorHeaders := map[string]string{}
    errorHeaders["Unauthenticated.challenge"] = "WWW-Authenticate"
    return protocol.NewOperationRestMetadata(
      paramsTypeMap,
      pathParams,
      queryParams,
      headerParams,
      "spec",
      "PUT",
      "/esx/attestation/raw-config",
       resultHeaders,
       204,
       errorHeaders,
       map[string]int{"Error": 500,"InvalidArgument": 400,"Unauthenticated": 401,"Unauthorized": 403})
}



func RawConfigInfoBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["configuration"] = bindings.NewOptionalType(bindings.NewBlobType())
    fieldNameMap["configuration"] = "Configuration"
    fields["fingerprint"] = bindings.NewStringType()
    fieldNameMap["fingerprint"] = "Fingerprint"
    fields["last_update_time"] = bindings.NewDateTimeType()
    fieldNameMap["last_update_time"] = "LastUpdateTime"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.esx.attestation.raw_config.info",fields, reflect.TypeOf(RawConfigInfo{}), fieldNameMap, validators)
}

func RawConfigSetSpecBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["configuration"] = bindings.NewBlobType()
    fieldNameMap["configuration"] = "Configuration"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.esx.attestation.raw_config.set_spec",fields, reflect.TypeOf(RawConfigSetSpec{}), fieldNameMap, validators)
}


