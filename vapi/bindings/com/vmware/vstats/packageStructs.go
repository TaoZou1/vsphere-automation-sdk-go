/* Copyright © 2019 VMware, Inc. All Rights Reserved.
     SPDX-License-Identifier: BSD-2-Clause */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Data type definitions file for package: com.vmware.vstats.
 * Includes binding types of a top level structures and enumerations.
 * Shared by client-side stubs and server-side skeletons to ensure type
 * compatibility.
 */

package vstats

import (
    "reflect"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/bindings/com/vmware/vapi/std"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/bindings"
)


// The ``QueryPredicate`` enum describes how to query an id present in a ``RsrcId``.
//
// <p> See {@link com.vmware.vapi.bindings.ApiEnumeration enumerated types description}.

type QueryPredicate string

const (
    // Matching id-s by equality.
     QueryPredicate_EQUAL QueryPredicate = "EQUAL"
    // Matching all available id-s.
     QueryPredicate_ALL QueryPredicate = "ALL"
)

func (q QueryPredicate) QueryPredicate() bool {
    switch q {
        case QueryPredicate_EQUAL:
            return true
        case QueryPredicate_ALL:
            return true
        default:
            return false
    }
}





// The ``CidMid`` class is used to designate a counter. It contains a counter id that identifies the semantical counter. There is optional metadata identifier that identifies the particular generation of the counter. When counter metadata is not designated vStats will use a default for the counter metadata.
type CidMid struct {
    // Counter Id. CID is a string with rather strong semantic consistency across deployments and versions. If two CIDs are identical it implies the corresponding counters are the same.
    Cid string
    // MID is a 64-bit integer with strong consistency. Given a particular CID=cid, if two MIDs are the same, then the corresponding counter-metadata objects are same.
    Mid *string
}


// TODO some error doesn't have ErrorType. Resolve this later VAPI-3009
//


// The ``RsrcId`` class specifies identification of a resource to be monitored by an acquisition specification record.
type RsrcId struct {
    // Key relates to the corresponding ResourceIdDefinition of the related resource address schema.
    Key *string
    // Type of the resource identified by the Resource Id.
    Type_ *string
    // The id value binding the related resource id definition.
    IdValue string
    // Predicate to use to match resource Ids.
    Predicate *QueryPredicate
    // An optional designation of the scheme.
    Scheme *string
}


// TODO some error doesn't have ErrorType. Resolve this later VAPI-3009
//


// The ``UserInfo`` class contains human legible, localizable description, used for VMware provided objects.
type UserInfo struct {
    // Short label.
    Label std.LocalizableMessage
    // Detailed description of the object.
    Description std.LocalizableMessage
}


// TODO some error doesn't have ErrorType. Resolve this later VAPI-3009
//






func CidMidBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["cid"] = bindings.NewIdType([]string {"com.vmware.vstats.model.Counter"}, "")
    fieldNameMap["cid"] = "Cid"
    fields["mid"] = bindings.NewOptionalType(bindings.NewIdType([]string {"com.vmware.vstats.model.CounterMetadata"}, ""))
    fieldNameMap["mid"] = "Mid"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vstats.cid_mid",fields, reflect.TypeOf(CidMid{}), fieldNameMap, validators)
}

func RsrcIdBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["key"] = bindings.NewOptionalType(bindings.NewStringType())
    fieldNameMap["key"] = "Key"
    fields["type"] = bindings.NewOptionalType(bindings.NewStringType())
    fieldNameMap["type"] = "Type_"
    fields["id_value"] = bindings.NewStringType()
    fieldNameMap["id_value"] = "IdValue"
    fields["predicate"] = bindings.NewOptionalType(bindings.NewEnumType("com.vmware.vstats.query_predicate", reflect.TypeOf(QueryPredicate(QueryPredicate_EQUAL))))
    fieldNameMap["predicate"] = "Predicate"
    fields["scheme"] = bindings.NewOptionalType(bindings.NewStringType())
    fieldNameMap["scheme"] = "Scheme"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vstats.rsrc_id",fields, reflect.TypeOf(RsrcId{}), fieldNameMap, validators)
}

func UserInfoBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["label"] = bindings.NewReferenceType(std.LocalizableMessageBindingType)
    fieldNameMap["label"] = "Label"
    fields["description"] = bindings.NewReferenceType(std.LocalizableMessageBindingType)
    fieldNameMap["description"] = "Description"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vstats.user_info",fields, reflect.TypeOf(UserInfo{}), fieldNameMap, validators)
}


