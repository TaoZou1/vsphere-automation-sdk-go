/* Copyright © 2019 VMware, Inc. All Rights Reserved.
     SPDX-License-Identifier: BSD-2-Clause */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Data type definitions file for service: Snmp.
 * Includes binding types of a structures and enumerations defined in the service.
 * Shared by client-side stubs and server-side skeletons to ensure type
 * compatibility.
 */

package monitoring

import (
    "reflect"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/bindings"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/data"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/protocol"
)



// ``SNMPAuthProto`` enumeration class Defines SNMP authentication protocols
//
// <p> See {@link com.vmware.vapi.bindings.ApiEnumeration enumerated types description}.
 
type Snmp_SNMPAuthProto string

const (
    // NONE
     Snmp_SNMPAuthProto_none Snmp_SNMPAuthProto = "none"
    // SHA1
     Snmp_SNMPAuthProto_SHA1 Snmp_SNMPAuthProto = "SHA1"
    // MD5
     Snmp_SNMPAuthProto_MD5 Snmp_SNMPAuthProto = "MD5"
)

func (s Snmp_SNMPAuthProto) Snmp_SNMPAuthProto() bool {
    switch s {
        case Snmp_SNMPAuthProto_none:
            return true
        case Snmp_SNMPAuthProto_SHA1:
            return true
        case Snmp_SNMPAuthProto_MD5:
            return true
        default:
            return false
    }
}




// ``SNMPPrivProto`` enumeration class Defines SNMP privacy protocols
//
// <p> See {@link com.vmware.vapi.bindings.ApiEnumeration enumerated types description}.
 
type Snmp_SNMPPrivProto string

const (
    // AES128
     Snmp_SNMPPrivProto_AES128 Snmp_SNMPPrivProto = "AES128"
    // NONE
     Snmp_SNMPPrivProto_none Snmp_SNMPPrivProto = "none"
)

func (s Snmp_SNMPPrivProto) Snmp_SNMPPrivProto() bool {
    switch s {
        case Snmp_SNMPPrivProto_AES128:
            return true
        case Snmp_SNMPPrivProto_none:
            return true
        default:
            return false
    }
}




// ``SNMPSecLevel`` enumeration class Defines SNMP decurity levels
//
// <p> See {@link com.vmware.vapi.bindings.ApiEnumeration enumerated types description}.
 
type Snmp_SNMPSecLevel string

const (
    // none
     Snmp_SNMPSecLevel_none Snmp_SNMPSecLevel = "none"
    // auth
     Snmp_SNMPSecLevel_auth Snmp_SNMPSecLevel = "auth"
    // priv
     Snmp_SNMPSecLevel_priv Snmp_SNMPSecLevel = "priv"
)

func (s Snmp_SNMPSecLevel) Snmp_SNMPSecLevel() bool {
    switch s {
        case Snmp_SNMPSecLevel_none:
            return true
        case Snmp_SNMPSecLevel_auth:
            return true
        case Snmp_SNMPSecLevel_priv:
            return true
        default:
            return false
    }
}




// ``SNMPv3Notfication`` enumeration class Defines SNMP v3 notification types
//
// <p> See {@link com.vmware.vapi.bindings.ApiEnumeration enumerated types description}.
 
type Snmp_SNMPv3Notfication string

const (
    // inform
     Snmp_SNMPv3Notfication_inform Snmp_SNMPv3Notfication = "inform"
    // trap
     Snmp_SNMPv3Notfication_trap Snmp_SNMPv3Notfication = "trap"
)

func (s Snmp_SNMPv3Notfication) Snmp_SNMPv3Notfication() bool {
    switch s {
        case Snmp_SNMPv3Notfication_inform:
            return true
        case Snmp_SNMPv3Notfication_trap:
            return true
        default:
            return false
    }
}





// ``SNMPLimits`` class Structure that provides various limits of the SNMP agent.
 type SnmpSNMPLimits struct {
    // Set up maximum communities limit
    MaxCommunities int64
    // Set up max trap destinations limit
    MaxTrapDestinationsV1 int64
    // Set up max destinations limit
    MaxDestinationsV3 int64
    // Set up max notification Filters
    MaxNotificationFilters int64
    // Set up max community length
    MaxCommunityLength int64
    // Set up max buffer size
    MaxBufferSize int64
}


// TODO some error doesn't have ErrorType. Resolve this later VAPI-3009
//


// ``SNMPTestResults`` class Structure to provide operators diagnostics test results.
 type SnmpSNMPTestResults struct {
    // Set success to true/false
    Success bool
    // message
    Message string
}


// TODO some error doesn't have ErrorType. Resolve this later VAPI-3009
//


// ``SNMPStats`` class Structure to provide operators diagnostics on snmp agent itself.
 type SnmpSNMPStats struct {
    // System uptime
    Sysuptime string
    // Last updated time
    Worstrtimelast string
    // Average response time
    Avgresponsetime string
    // Response time
    Worstresponsetime string
    // No of input packets
    Inpkts int64
    // No of output packets
    Outpkts int64
    // No of stats not in time window
    Usmstatsnotintimewindows int64
    // No of usm stats unknown
    Usmstatsunknownusernames int64
    // No of usm stats unknown engine ids
    Usmstatsunknownengineids int64
    // No of wrogn digests
    Usmstatswrongdigests int64
    // No. of decryption errors
    Usmstatsdecryptionerrors int64
    // No of bad versions
    Inbadversions int64
    // No of bad community names
    Inbadcommunitynames int64
    // No of bad community uses
    Inbadcommunityuses int64
    // No of parse errors
    Inasnparseerrs int64
    // No of too bigs
    Intoobigs int64
    // No of no such names
    Innosuchnames int64
    // No of bad values
    Inbadvalues int64
    // No of gen errors
    Ingenerrs int64
    // No out output too bigs
    Outtoobigs int64
    // No of no such names
    Outnosuchnames int64
    // No of bad values
    Outbadvalues int64
    // No of gen errors
    Outgenerrs int64
    // No of output traps
    Outtraps int64
    // No of silent drops
    Silentdrops int64
    // No of ave:rage var binds
    Avgvarbinds int64
    // No of max var binds
    Maxvarbinds int64
}


// TODO some error doesn't have ErrorType. Resolve this later VAPI-3009
//


// ``SNMPConfig`` class Structure that defines the SNMP configuration, provided as input to set(), and never the result of get(). See SNMPConfigReadOnly. This structure is used to configure SNMP v1, v2c, and v3.
 type SnmpSNMPConfig struct {
    // Set the default authentication protocol. Values can be none, MD5, or SHA1.
    Authentication Snmp_SNMPAuthProto
    // Set up to ten communities, each of no more than 64 characters long. The format is: community1[,community2,...]. This setting overwrites any previous settings.
    Communities []string
    // Set SNMPv3 engine ID. The engine ID must contain 5 to 32 hexadecimal characters. "0x" and colon (:) are removed from the ID.
    Engineid string
    // System Agent syslog logging level: debug|info|warning|error.
    Loglevel string
    // Comma-separated list of trap OIDs (object identifiers) for traps not to be sent by the agent. Use 'reset' to clear the setting.
    Notraps []string
    // Set up a UDP port which the SNMP agent uses to listen on for polling requests. The default UDP port is 161.
    Port int64
    // Set the default privacy protocol. Values: none or AES128.
    Privacy Snmp_SNMPPrivProto
    // Set up to five inform user IDs. The format is: user/auth-proto/-|auth-hash/priv-proto/-|priv-hash/engine-id[,...]. Here, user must be maximum 32 characters long; auth-proto is none, MD5 or SHA1; priv-proto is none or AES; '-' indicates no hash; engine-id is a hexadecimal string '0x0-9a-f' and must be up to 32 characters long.
    Remoteusers []string
    // System contact string as presented in sysContact.0. Up to 255 characters long.
    Syscontact string
    // System location string as presented in sysLocation.0. Up to 255 characters long.
    Syslocation string
    // Set up to three targets to which to send SNMPv1 traps. The format is: ip-or-hostname[\\\\@port]/community[,...]. The default port is UDP 162. This setting overwrites any previous settings.
    Targets []string
    // Set up to five local users. The format is: user/-|auth-hash/-|priv-hash/model[,...]. Here user is maximum 32 characters long; '-' indicates no hash; model is one of none, auth or priv.
    Users []string
    // Set up to three SNMPv3 notification targets. Format is: ip-or-hostname[\\\\@port]/remote-user/security-level/trap|inform[,...].
    V3targets []string
}


// TODO some error doesn't have ErrorType. Resolve this later VAPI-3009
//


// ``SNMPUser`` class Structure that defines information associated with an SNMP user. authKey and privKey are localized keys defined in http://tools.ietf.org/html/rfc3826#section-1.2.
 type SnmpSNMPUser struct {
    // SNMP Username
    Username string
    // SNMP security level
    SecLevel Snmp_SNMPSecLevel
    // SNMP authorization key
    AuthKey string
    // SNMP privacy key
    PrivKey string
}


// TODO some error doesn't have ErrorType. Resolve this later VAPI-3009
//


// ``SNMPv3Target`` class Structure that defines an SNMP v3 inform or trap target.
 type SnmpSNMPv3Target struct {
    // SNMP target type
    Type_ Snmp_SNMPv3Notfication
    // SNMP security level
    SecLevel Snmp_SNMPSecLevel
    // SNMP target ip
    Ip string
    // SNMP target port
    Port int64
    // SNMP User
    User string
}


// TODO some error doesn't have ErrorType. Resolve this later VAPI-3009
//


// ``SNMPv1TrapTarget`` class Structure that defines an SNMP v1/v2c trap target.
 type SnmpSNMPv1TrapTarget struct {
    // SNMP target ip
    Ip string
    // SNMP target port
    Port int64
    // SNMP target community
    Community string
}


// TODO some error doesn't have ErrorType. Resolve this later VAPI-3009
//


// ``SNMPRemoteUser`` class Structure that defines a user at particular remote SNMPv3 entity needed when using informs. auth_key and priv_key contained localized keys as defined in http://tools.ietf.org/html/rfc3826#section-1.2.
 type SnmpSNMPRemoteUser struct {
    // SNMP Username
    Username string
    // SNMP security level
    SecLevel Snmp_SNMPSecLevel
    // SNMP authorization protocol
    Authentication Snmp_SNMPAuthProto
    // SNMP authorization key
    AuthKey string
    // SNMP privacy protocol
    Privacy Snmp_SNMPPrivProto
    // SNMP privacy key
    PrivKey string
    // SNMP v3 engine id
    Engineid string
}


// TODO some error doesn't have ErrorType. Resolve this later VAPI-3009
//


// ``SNMPConfigReadOnly`` class Structure that defines the SNMP configuration, the result of get(), and never provided as input to set(). This structure differs from SNMPConfig because it contains localized keys (as defined in http://tools.ietf.org/html/rfc3826#section-1.2), instead of raw secret strings. This structure can be used to configure SNMP v1, v2c, and v3. Keep this structure in sync with vmw_snmp.py:_default_config(). Note that if a field if left empty, it is considered unset and will be ignored. Existing array elements below can be unset by sending an element with the string 'reset'.
 type SnmpSNMPConfigReadOnly struct {
    // Set the default authentication protocol. Values can be none, MD5, or SHA1.
    Authentication Snmp_SNMPAuthProto
    // Set up to ten communities, each of no more than 64 characters long. The format is: community1[,community2,...]. This setting overwrites any previous settings.
    Communities []string
    // Set enable to true/false
    Enable bool
    // Set SNMPv3 engine ID.
    Engineid string
    // System Agent syslog logging level: debug|info|warning|error.
    Loglevel string
    // Comma-separated list of trap OIDs (object identifiers) for traps not to be sent by the agent. Use 'reset' to clear the setting.
    Notraps []string
    // Set up a UDP port which the SNMP agent uses to listen on for polling requests. The default UDP port is 161.
    Port int64
    // Set the default privacy protocol.
    Privacy Snmp_SNMPPrivProto
    // System contact string as presented in sysContact.0. Up to 255 characters long.
    Syscontact string
    // System location string as presented in sysLocation.0. Up to 255 characters long.
    Syslocation string
    // Set up to three targets to which to send SNMPv1 traps.
    Targets []SnmpSNMPv1TrapTarget
    // Set up to five local users.
    Users []SnmpSNMPUser
    // Set up remote users.
    Remoteusers []SnmpSNMPRemoteUser
    // Set up to three SNMPv3 notification targets. Format is: ip-or-hostname[\\\\@port]/remote-user/security-level/trap|inform[,...].
    V3targets []SnmpSNMPv3Target
    // Set up pid
    Pid string
}


// TODO some error doesn't have ErrorType. Resolve this later VAPI-3009
//


// ``SNMPHashConfig`` class Structure to provide up to two secrets to combine with the SNMPv3 engine ID and authentication or privacy protocol to form a localized hash. auth_hash is always required, priv_hash can be empty. By default arguments are paths on the local filesystem, raw_secret takes path to be the actual raw secret. First implementation was in ESXi: esxcli system snmp hash --help
 type SnmpSNMPHashConfig struct {
    // Provide filename to secret for authentication hash, use in set --users (required secret)
    AuthHash string
    // Provide filename to secret for privacy hash, use in set --users (secret)
    PrivHash string
    // Make --auth_path and --priv_path flags read raw secret from command line instead of file.
    RawSecret bool
}


// TODO some error doesn't have ErrorType. Resolve this later VAPI-3009
//


// ``SNMPHashResults`` class Structure to provide operators diagnostics test results.
 type SnmpSNMPHashResults struct {
    // SNMP authentication key
    AuthKey string
    // SNMP privacy key
    PrivKey string
}


// TODO some error doesn't have ErrorType. Resolve this later VAPI-3009
//





func snmpResetInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func snmpResetOutputType() bindings.BindingType {
    return bindings.NewVoidType()
}

func snmpResetRestMetadata() protocol.OperationRestMetadata {
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


func snmpEnableInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func snmpEnableOutputType() bindings.BindingType {
    return bindings.NewVoidType()
}

func snmpEnableRestMetadata() protocol.OperationRestMetadata {
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


func snmpHashInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["config"] = bindings.NewReferenceType(SnmpSNMPHashConfigBindingType)
    fieldNameMap["config"] = "Config"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func snmpHashOutputType() bindings.BindingType {
    return bindings.NewReferenceType(SnmpSNMPHashResultsBindingType)
}

func snmpHashRestMetadata() protocol.OperationRestMetadata {
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


func snmpLimitsInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func snmpLimitsOutputType() bindings.BindingType {
    return bindings.NewReferenceType(SnmpSNMPLimitsBindingType)
}

func snmpLimitsRestMetadata() protocol.OperationRestMetadata {
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


func snmpGetInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func snmpGetOutputType() bindings.BindingType {
    return bindings.NewReferenceType(SnmpSNMPConfigReadOnlyBindingType)
}

func snmpGetRestMetadata() protocol.OperationRestMetadata {
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


func snmpDisableInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func snmpDisableOutputType() bindings.BindingType {
    return bindings.NewVoidType()
}

func snmpDisableRestMetadata() protocol.OperationRestMetadata {
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


func snmpSetInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["config"] = bindings.NewReferenceType(SnmpSNMPConfigBindingType)
    fieldNameMap["config"] = "Config"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func snmpSetOutputType() bindings.BindingType {
    return bindings.NewVoidType()
}

func snmpSetRestMetadata() protocol.OperationRestMetadata {
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


func snmpTestInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func snmpTestOutputType() bindings.BindingType {
    return bindings.NewReferenceType(SnmpSNMPTestResultsBindingType)
}

func snmpTestRestMetadata() protocol.OperationRestMetadata {
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


func snmpStatsInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func snmpStatsOutputType() bindings.BindingType {
    return bindings.NewReferenceType(SnmpSNMPStatsBindingType)
}

func snmpStatsRestMetadata() protocol.OperationRestMetadata {
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



func SnmpSNMPLimitsBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["max_communities"] = bindings.NewIntegerType()
    fieldNameMap["max_communities"] = "MaxCommunities"
    fields["max_trap_destinations_v1"] = bindings.NewIntegerType()
    fieldNameMap["max_trap_destinations_v1"] = "MaxTrapDestinationsV1"
    fields["max_destinations_v3"] = bindings.NewIntegerType()
    fieldNameMap["max_destinations_v3"] = "MaxDestinationsV3"
    fields["max_notification_filters"] = bindings.NewIntegerType()
    fieldNameMap["max_notification_filters"] = "MaxNotificationFilters"
    fields["max_community_length"] = bindings.NewIntegerType()
    fieldNameMap["max_community_length"] = "MaxCommunityLength"
    fields["max_buffer_size"] = bindings.NewIntegerType()
    fieldNameMap["max_buffer_size"] = "MaxBufferSize"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.appliance.techpreview.monitoring.snmp.SNMP_limits",fields, reflect.TypeOf(SnmpSNMPLimits{}), fieldNameMap, validators)
}

func SnmpSNMPTestResultsBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["success"] = bindings.NewBooleanType()
    fieldNameMap["success"] = "Success"
    fields["message"] = bindings.NewStringType()
    fieldNameMap["message"] = "Message"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.appliance.techpreview.monitoring.snmp.SNMP_test_results",fields, reflect.TypeOf(SnmpSNMPTestResults{}), fieldNameMap, validators)
}

func SnmpSNMPStatsBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["sysuptime"] = bindings.NewStringType()
    fieldNameMap["sysuptime"] = "Sysuptime"
    fields["worstrtimelast"] = bindings.NewStringType()
    fieldNameMap["worstrtimelast"] = "Worstrtimelast"
    fields["avgresponsetime"] = bindings.NewStringType()
    fieldNameMap["avgresponsetime"] = "Avgresponsetime"
    fields["worstresponsetime"] = bindings.NewStringType()
    fieldNameMap["worstresponsetime"] = "Worstresponsetime"
    fields["inpkts"] = bindings.NewIntegerType()
    fieldNameMap["inpkts"] = "Inpkts"
    fields["outpkts"] = bindings.NewIntegerType()
    fieldNameMap["outpkts"] = "Outpkts"
    fields["usmstatsnotintimewindows"] = bindings.NewIntegerType()
    fieldNameMap["usmstatsnotintimewindows"] = "Usmstatsnotintimewindows"
    fields["usmstatsunknownusernames"] = bindings.NewIntegerType()
    fieldNameMap["usmstatsunknownusernames"] = "Usmstatsunknownusernames"
    fields["usmstatsunknownengineids"] = bindings.NewIntegerType()
    fieldNameMap["usmstatsunknownengineids"] = "Usmstatsunknownengineids"
    fields["usmstatswrongdigests"] = bindings.NewIntegerType()
    fieldNameMap["usmstatswrongdigests"] = "Usmstatswrongdigests"
    fields["usmstatsdecryptionerrors"] = bindings.NewIntegerType()
    fieldNameMap["usmstatsdecryptionerrors"] = "Usmstatsdecryptionerrors"
    fields["inbadversions"] = bindings.NewIntegerType()
    fieldNameMap["inbadversions"] = "Inbadversions"
    fields["inbadcommunitynames"] = bindings.NewIntegerType()
    fieldNameMap["inbadcommunitynames"] = "Inbadcommunitynames"
    fields["inbadcommunityuses"] = bindings.NewIntegerType()
    fieldNameMap["inbadcommunityuses"] = "Inbadcommunityuses"
    fields["inasnparseerrs"] = bindings.NewIntegerType()
    fieldNameMap["inasnparseerrs"] = "Inasnparseerrs"
    fields["intoobigs"] = bindings.NewIntegerType()
    fieldNameMap["intoobigs"] = "Intoobigs"
    fields["innosuchnames"] = bindings.NewIntegerType()
    fieldNameMap["innosuchnames"] = "Innosuchnames"
    fields["inbadvalues"] = bindings.NewIntegerType()
    fieldNameMap["inbadvalues"] = "Inbadvalues"
    fields["ingenerrs"] = bindings.NewIntegerType()
    fieldNameMap["ingenerrs"] = "Ingenerrs"
    fields["outtoobigs"] = bindings.NewIntegerType()
    fieldNameMap["outtoobigs"] = "Outtoobigs"
    fields["outnosuchnames"] = bindings.NewIntegerType()
    fieldNameMap["outnosuchnames"] = "Outnosuchnames"
    fields["outbadvalues"] = bindings.NewIntegerType()
    fieldNameMap["outbadvalues"] = "Outbadvalues"
    fields["outgenerrs"] = bindings.NewIntegerType()
    fieldNameMap["outgenerrs"] = "Outgenerrs"
    fields["outtraps"] = bindings.NewIntegerType()
    fieldNameMap["outtraps"] = "Outtraps"
    fields["silentdrops"] = bindings.NewIntegerType()
    fieldNameMap["silentdrops"] = "Silentdrops"
    fields["avgvarbinds"] = bindings.NewIntegerType()
    fieldNameMap["avgvarbinds"] = "Avgvarbinds"
    fields["maxvarbinds"] = bindings.NewIntegerType()
    fieldNameMap["maxvarbinds"] = "Maxvarbinds"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.appliance.techpreview.monitoring.snmp.SNMP_stats",fields, reflect.TypeOf(SnmpSNMPStats{}), fieldNameMap, validators)
}

func SnmpSNMPConfigBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["authentication"] = bindings.NewEnumType("com.vmware.appliance.techpreview.monitoring.snmp.SNMP_auth_proto", reflect.TypeOf(Snmp_SNMPAuthProto(Snmp_SNMPAuthProto_none)))
    fieldNameMap["authentication"] = "Authentication"
    fields["communities"] = bindings.NewListType(bindings.NewStringType(), reflect.TypeOf([]string{}))
    fieldNameMap["communities"] = "Communities"
    fields["engineid"] = bindings.NewStringType()
    fieldNameMap["engineid"] = "Engineid"
    fields["loglevel"] = bindings.NewStringType()
    fieldNameMap["loglevel"] = "Loglevel"
    fields["notraps"] = bindings.NewListType(bindings.NewStringType(), reflect.TypeOf([]string{}))
    fieldNameMap["notraps"] = "Notraps"
    fields["port"] = bindings.NewIntegerType()
    fieldNameMap["port"] = "Port"
    fields["privacy"] = bindings.NewEnumType("com.vmware.appliance.techpreview.monitoring.snmp.SNMP_priv_proto", reflect.TypeOf(Snmp_SNMPPrivProto(Snmp_SNMPPrivProto_AES128)))
    fieldNameMap["privacy"] = "Privacy"
    fields["remoteusers"] = bindings.NewListType(bindings.NewStringType(), reflect.TypeOf([]string{}))
    fieldNameMap["remoteusers"] = "Remoteusers"
    fields["syscontact"] = bindings.NewStringType()
    fieldNameMap["syscontact"] = "Syscontact"
    fields["syslocation"] = bindings.NewStringType()
    fieldNameMap["syslocation"] = "Syslocation"
    fields["targets"] = bindings.NewListType(bindings.NewStringType(), reflect.TypeOf([]string{}))
    fieldNameMap["targets"] = "Targets"
    fields["users"] = bindings.NewListType(bindings.NewStringType(), reflect.TypeOf([]string{}))
    fieldNameMap["users"] = "Users"
    fields["v3targets"] = bindings.NewListType(bindings.NewStringType(), reflect.TypeOf([]string{}))
    fieldNameMap["v3targets"] = "V3targets"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.appliance.techpreview.monitoring.snmp.SNMP_config",fields, reflect.TypeOf(SnmpSNMPConfig{}), fieldNameMap, validators)
}

func SnmpSNMPUserBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["username"] = bindings.NewStringType()
    fieldNameMap["username"] = "Username"
    fields["sec_level"] = bindings.NewEnumType("com.vmware.appliance.techpreview.monitoring.snmp.SNMP_sec_level", reflect.TypeOf(Snmp_SNMPSecLevel(Snmp_SNMPSecLevel_none)))
    fieldNameMap["sec_level"] = "SecLevel"
    fields["auth_key"] = bindings.NewStringType()
    fieldNameMap["auth_key"] = "AuthKey"
    fields["priv_key"] = bindings.NewStringType()
    fieldNameMap["priv_key"] = "PrivKey"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.appliance.techpreview.monitoring.snmp.SNMP_user",fields, reflect.TypeOf(SnmpSNMPUser{}), fieldNameMap, validators)
}

func SnmpSNMPv3TargetBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["type"] = bindings.NewEnumType("com.vmware.appliance.techpreview.monitoring.snmp.SNM_pv3_notfication", reflect.TypeOf(Snmp_SNMPv3Notfication(Snmp_SNMPv3Notfication_inform)))
    fieldNameMap["type"] = "Type_"
    fields["sec_level"] = bindings.NewEnumType("com.vmware.appliance.techpreview.monitoring.snmp.SNMP_sec_level", reflect.TypeOf(Snmp_SNMPSecLevel(Snmp_SNMPSecLevel_none)))
    fieldNameMap["sec_level"] = "SecLevel"
    fields["ip"] = bindings.NewStringType()
    fieldNameMap["ip"] = "Ip"
    fields["port"] = bindings.NewIntegerType()
    fieldNameMap["port"] = "Port"
    fields["user"] = bindings.NewStringType()
    fieldNameMap["user"] = "User"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.appliance.techpreview.monitoring.snmp.SNM_pv3_target",fields, reflect.TypeOf(SnmpSNMPv3Target{}), fieldNameMap, validators)
}

func SnmpSNMPv1TrapTargetBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["ip"] = bindings.NewStringType()
    fieldNameMap["ip"] = "Ip"
    fields["port"] = bindings.NewIntegerType()
    fieldNameMap["port"] = "Port"
    fields["community"] = bindings.NewStringType()
    fieldNameMap["community"] = "Community"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.appliance.techpreview.monitoring.snmp.SNM_pv1_trap_target",fields, reflect.TypeOf(SnmpSNMPv1TrapTarget{}), fieldNameMap, validators)
}

func SnmpSNMPRemoteUserBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["username"] = bindings.NewStringType()
    fieldNameMap["username"] = "Username"
    fields["sec_level"] = bindings.NewEnumType("com.vmware.appliance.techpreview.monitoring.snmp.SNMP_sec_level", reflect.TypeOf(Snmp_SNMPSecLevel(Snmp_SNMPSecLevel_none)))
    fieldNameMap["sec_level"] = "SecLevel"
    fields["authentication"] = bindings.NewEnumType("com.vmware.appliance.techpreview.monitoring.snmp.SNMP_auth_proto", reflect.TypeOf(Snmp_SNMPAuthProto(Snmp_SNMPAuthProto_none)))
    fieldNameMap["authentication"] = "Authentication"
    fields["auth_key"] = bindings.NewStringType()
    fieldNameMap["auth_key"] = "AuthKey"
    fields["privacy"] = bindings.NewEnumType("com.vmware.appliance.techpreview.monitoring.snmp.SNMP_priv_proto", reflect.TypeOf(Snmp_SNMPPrivProto(Snmp_SNMPPrivProto_AES128)))
    fieldNameMap["privacy"] = "Privacy"
    fields["priv_key"] = bindings.NewStringType()
    fieldNameMap["priv_key"] = "PrivKey"
    fields["engineid"] = bindings.NewStringType()
    fieldNameMap["engineid"] = "Engineid"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.appliance.techpreview.monitoring.snmp.SNMP_remote_user",fields, reflect.TypeOf(SnmpSNMPRemoteUser{}), fieldNameMap, validators)
}

func SnmpSNMPConfigReadOnlyBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["authentication"] = bindings.NewEnumType("com.vmware.appliance.techpreview.monitoring.snmp.SNMP_auth_proto", reflect.TypeOf(Snmp_SNMPAuthProto(Snmp_SNMPAuthProto_none)))
    fieldNameMap["authentication"] = "Authentication"
    fields["communities"] = bindings.NewListType(bindings.NewStringType(), reflect.TypeOf([]string{}))
    fieldNameMap["communities"] = "Communities"
    fields["enable"] = bindings.NewBooleanType()
    fieldNameMap["enable"] = "Enable"
    fields["engineid"] = bindings.NewStringType()
    fieldNameMap["engineid"] = "Engineid"
    fields["loglevel"] = bindings.NewStringType()
    fieldNameMap["loglevel"] = "Loglevel"
    fields["notraps"] = bindings.NewListType(bindings.NewStringType(), reflect.TypeOf([]string{}))
    fieldNameMap["notraps"] = "Notraps"
    fields["port"] = bindings.NewIntegerType()
    fieldNameMap["port"] = "Port"
    fields["privacy"] = bindings.NewEnumType("com.vmware.appliance.techpreview.monitoring.snmp.SNMP_priv_proto", reflect.TypeOf(Snmp_SNMPPrivProto(Snmp_SNMPPrivProto_AES128)))
    fieldNameMap["privacy"] = "Privacy"
    fields["syscontact"] = bindings.NewStringType()
    fieldNameMap["syscontact"] = "Syscontact"
    fields["syslocation"] = bindings.NewStringType()
    fieldNameMap["syslocation"] = "Syslocation"
    fields["targets"] = bindings.NewListType(bindings.NewReferenceType(SnmpSNMPv1TrapTargetBindingType), reflect.TypeOf([]SnmpSNMPv1TrapTarget{}))
    fieldNameMap["targets"] = "Targets"
    fields["users"] = bindings.NewListType(bindings.NewReferenceType(SnmpSNMPUserBindingType), reflect.TypeOf([]SnmpSNMPUser{}))
    fieldNameMap["users"] = "Users"
    fields["remoteusers"] = bindings.NewListType(bindings.NewReferenceType(SnmpSNMPRemoteUserBindingType), reflect.TypeOf([]SnmpSNMPRemoteUser{}))
    fieldNameMap["remoteusers"] = "Remoteusers"
    fields["v3targets"] = bindings.NewListType(bindings.NewReferenceType(SnmpSNMPv3TargetBindingType), reflect.TypeOf([]SnmpSNMPv3Target{}))
    fieldNameMap["v3targets"] = "V3targets"
    fields["pid"] = bindings.NewStringType()
    fieldNameMap["pid"] = "Pid"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.appliance.techpreview.monitoring.snmp.SNMP_config_read_only",fields, reflect.TypeOf(SnmpSNMPConfigReadOnly{}), fieldNameMap, validators)
}

func SnmpSNMPHashConfigBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["auth_hash"] = bindings.NewStringType()
    fieldNameMap["auth_hash"] = "AuthHash"
    fields["priv_hash"] = bindings.NewStringType()
    fieldNameMap["priv_hash"] = "PrivHash"
    fields["raw_secret"] = bindings.NewBooleanType()
    fieldNameMap["raw_secret"] = "RawSecret"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.appliance.techpreview.monitoring.snmp.SNMP_hash_config",fields, reflect.TypeOf(SnmpSNMPHashConfig{}), fieldNameMap, validators)
}

func SnmpSNMPHashResultsBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["auth_key"] = bindings.NewStringType()
    fieldNameMap["auth_key"] = "AuthKey"
    fields["priv_key"] = bindings.NewStringType()
    fieldNameMap["priv_key"] = "PrivKey"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.appliance.techpreview.monitoring.snmp.SNMP_hash_results",fields, reflect.TypeOf(SnmpSNMPHashResults{}), fieldNameMap, validators)
}


