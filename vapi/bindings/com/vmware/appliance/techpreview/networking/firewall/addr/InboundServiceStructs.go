/* Copyright © 2019 VMware, Inc. All Rights Reserved.
     SPDX-License-Identifier: BSD-2-Clause */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Data type definitions file for service: Inbound.
 * Includes binding types of a structures and enumerations defined in the service.
 * Shared by client-side stubs and server-side skeletons to ensure type
 * compatibility.
 */

package addr

import (
    "reflect"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/bindings"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/data"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/protocol"
)



// ``FirewallRulePolicy`` enumeration class Defines firewall rule policies
//
// <p> See {@link com.vmware.vapi.bindings.ApiEnumeration enumerated types description}.
 
type Inbound_FirewallRulePolicy string

const (
    // Deny packet with correpsonding address.
     Inbound_FirewallRulePolicy_deny Inbound_FirewallRulePolicy = "deny"
    // Allow packet with corresponding address.
     Inbound_FirewallRulePolicy_allow Inbound_FirewallRulePolicy = "allow"
)

func (f Inbound_FirewallRulePolicy) Inbound_FirewallRulePolicy() bool {
    switch f {
        case Inbound_FirewallRulePolicy_deny:
            return true
        case Inbound_FirewallRulePolicy_allow:
            return true
        default:
            return false
    }
}





// ``FirewallAddressRule`` class Structure that defines a single address-based firewall rule.
 type InboundFirewallAddressRule struct {
    // IPv4 or IPv6 address.
    Address string
    // CIDR prefix used to mask address. For example, an IPv4 prefix of 24 ignores the low-order 8 bits of address.
    Prefix int64
    // The allow or deny policy of this rule.
    Policy Inbound_FirewallRulePolicy
    // The interface to which this rule applies. An empty string or "\*" indicates that the rule applies to all interfaces.
    InterfaceName string
}



func (InboundFirewallAddressRule InboundFirewallAddressRule) Error() string {
    return "com.vmware.appliance.techpreview.networking.firewall.addr.firewall_address_rule"
}



// ``DeleteFirewallRule`` class Structure that defines networking.firewall.addr.inbound.delete api input argument
 type InboundDeleteFirewallRule struct {
    // Position before which to insert the rule (zero-based). If you try to insert the rule in a position whose number is greater than the number of rules, the firewall rule is inserted at the end of the list.
    Position int64
    // Delete all firewall rules. Set all argument to "true" to delete all rules or set the all argument to "false" to delete a single rule.
    All bool
}



func (InboundDeleteFirewallRule InboundDeleteFirewallRule) Error() string {
    return "com.vmware.appliance.techpreview.networking.firewall.addr.delete_firewall_rule"
}






func inboundAddInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["pos"] = bindings.NewIntegerType()
    fields["rule"] = bindings.NewReferenceType(InboundFirewallAddressRuleBindingType)
    fieldNameMap["pos"] = "Pos"
    fieldNameMap["rule"] = "Rule"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func inboundAddOutputType() bindings.BindingType {
    return bindings.NewVoidType()
}

func inboundAddRestMetadata() protocol.OperationRestMetadata {
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


func inboundSetInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["rules"] = bindings.NewListType(bindings.NewReferenceType(InboundFirewallAddressRuleBindingType), reflect.TypeOf([]InboundFirewallAddressRule{}))
    fieldNameMap["rules"] = "Rules"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func inboundSetOutputType() bindings.BindingType {
    return bindings.NewVoidType()
}

func inboundSetRestMetadata() protocol.OperationRestMetadata {
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


func inboundListInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func inboundListOutputType() bindings.BindingType {
    return bindings.NewListType(bindings.NewReferenceType(InboundFirewallAddressRuleBindingType), reflect.TypeOf([]InboundFirewallAddressRule{}))
}

func inboundListRestMetadata() protocol.OperationRestMetadata {
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


func inboundDeleteInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["config"] = bindings.NewReferenceType(InboundDeleteFirewallRuleBindingType)
    fieldNameMap["config"] = "Config"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func inboundDeleteOutputType() bindings.BindingType {
    return bindings.NewVoidType()
}

func inboundDeleteRestMetadata() protocol.OperationRestMetadata {
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



func InboundFirewallAddressRuleBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["address"] = bindings.NewStringType()
    fieldNameMap["address"] = "Address"
    fields["prefix"] = bindings.NewIntegerType()
    fieldNameMap["prefix"] = "Prefix"
    fields["policy"] = bindings.NewEnumType("com.vmware.appliance.techpreview.networking.firewall.addr.inbound.firewall_rule_policy", reflect.TypeOf(Inbound_FirewallRulePolicy(Inbound_FirewallRulePolicy_deny)))
    fieldNameMap["policy"] = "Policy"
    fields["interface_name"] = bindings.NewStringType()
    fieldNameMap["interface_name"] = "InterfaceName"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.appliance.techpreview.networking.firewall.addr.inbound.firewall_address_rule",fields, reflect.TypeOf(InboundFirewallAddressRule{}), fieldNameMap, validators)
}

func InboundDeleteFirewallRuleBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["position"] = bindings.NewIntegerType()
    fieldNameMap["position"] = "Position"
    fields["all"] = bindings.NewBooleanType()
    fieldNameMap["all"] = "All"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.appliance.techpreview.networking.firewall.addr.inbound.delete_firewall_rule",fields, reflect.TypeOf(InboundDeleteFirewallRule{}), fieldNameMap, validators)
}


