/* Copyright © 2019 VMware, Inc. All Rights Reserved.
     SPDX-License-Identifier: BSD-2-Clause */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Data type definitions file for package: com.vmware.appliance.recovery.backup.
 * Includes binding types of a top level structures and enumerations.
 * Shared by client-side stubs and server-side skeletons to ensure type
 * compatibility.
 */

package backup

import (
    "reflect"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/bindings"
    "net/url"
)



// The ``LocationSpec`` class has fields to represent a location on the backup server.
type LocationSpec struct {
    // Backup location URL.
    Location url.URL
    // Username for the given location.
    LocationUser *string
    // Password for the given location.
    LocationPassword *string
}


// TODO some error doesn't have ErrorType. Resolve this later VAPI-3009
//






func LocationSpecBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["location"] = bindings.NewUriType()
    fieldNameMap["location"] = "Location"
    fields["location_user"] = bindings.NewOptionalType(bindings.NewStringType())
    fieldNameMap["location_user"] = "LocationUser"
    fields["location_password"] = bindings.NewOptionalType(bindings.NewSecretType())
    fieldNameMap["location_password"] = "LocationPassword"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.appliance.recovery.backup.location_spec",fields, reflect.TypeOf(LocationSpec{}), fieldNameMap, validators)
}


