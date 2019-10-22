/* Copyright © 2019 VMware, Inc. All Rights Reserved.
     SPDX-License-Identifier: BSD-2-Clause */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Data type definitions file for package: com.vmware.esx.trusted_infrastructure.
 * Includes binding types of a top level structures and enumerations.
 * Shared by client-side stubs and server-side skeletons to ensure type
 * compatibility.
 */

package trusted_infrastructure

import (
    "reflect"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/bindings"
)



// The ``NetworkAddress`` class contains an IP address or DNS resolvable name and a port on which a connection can be established.
type NetworkAddress struct {
    // The IP address or DNS resolvable name of the service.
    Hostname string
    // The port of the service.
    Port *int64
}


// TODO some error doesn't have ErrorType. Resolve this later VAPI-3009
//


// The ``X509CertChain`` class contains x509 certificate chain.
type X509CertChain struct {
    // Certificate chain in base64 format
    CertChain []string
}


// TODO some error doesn't have ErrorType. Resolve this later VAPI-3009
//






func NetworkAddressBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["hostname"] = bindings.NewStringType()
    fieldNameMap["hostname"] = "Hostname"
    fields["port"] = bindings.NewOptionalType(bindings.NewIntegerType())
    fieldNameMap["port"] = "Port"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.esx.trusted_infrastructure.network_address",fields, reflect.TypeOf(NetworkAddress{}), fieldNameMap, validators)
}

func X509CertChainBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["cert_chain"] = bindings.NewListType(bindings.NewStringType(), reflect.TypeOf([]string{}))
    fieldNameMap["cert_chain"] = "CertChain"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.esx.trusted_infrastructure.x509_cert_chain",fields, reflect.TypeOf(X509CertChain{}), fieldNameMap, validators)
}


