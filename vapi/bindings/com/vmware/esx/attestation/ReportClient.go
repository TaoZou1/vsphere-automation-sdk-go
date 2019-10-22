/* Copyright © 2019 VMware, Inc. All Rights Reserved.
     SPDX-License-Identifier: BSD-2-Clause */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Interface file for service: Report
 * Used by client-side stubs.
 */

package attestation

import (
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/data"
)

// The ``Report`` interface provides methods to get attestation reports.
type ReportClient interface {


    // Request a report using remote attestation.
    //
    // @param requestParam The report request.
    // The parameter must contain all the properties defined in report.AttestRequest.
    // @return The report result.
    // The return value will contain all the properties defined in report.AttestResult.
    // @throws Error If a generic error.
    // @throws InvalidArgument If the request is invalid.
    // @throws Unauthenticated If the caller is not authenticated.
    // @throws Unauthorized If the caller is not authorized.
    Attest(requestParam *data.StructValue) (*data.StructValue, error) 

}
