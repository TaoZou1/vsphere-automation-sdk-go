/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Code generated. DO NOT EDIT.

/*
 * Interface file for service: InteropReport
 * Used by client-side stubs.
 */

package discovery


// The ``InteropReport`` interface provides methods to report the interoperability between a vCenter Server release version and the other installed VMware products registered in the vCenter Server instance.
type InteropReportClient interface {

    // Creates interoperability report between a vCenter Server release version and all registered products with the vCenter Server instance. 
    //
    //  The result of this operation can be queried by calling the null method where ``task`` is the response of this operation.
    //
    // @param specParam 
    // Specifies the target version against this interoperability check report will be generated. If null the report will be generated for the currently installed version of the vCenter server.
    // @return The interoperability report.
    // @throws Unauthenticated if the user can not be authenticated.
    // @throws Error If there is some unknown internal error. The accompanying error message will give more details about the failure.
	Create(specParam *InteropReportSpec) (InteropReportResult, error)
}
