/* Copyright © 2019 VMware, Inc. All Rights Reserved.
     SPDX-License-Identifier: BSD-2-Clause */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Interface file for service: Reservations
 * Used by client-side stubs.
 */

package orgs

import (
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/bindings/com/vmware/vmc/model"
)

type ReservationsClient interface {


    // Get all reservations for this org
    //
    // @param orgParam Organization identifier. (required)
    // @throws Unauthenticated  Unauthorized
    // @throws Unauthorized  Access not allowed to the operation for the current user
    List(orgParam string) ([]model.MaintenanceWindowEntry, error) 

}
