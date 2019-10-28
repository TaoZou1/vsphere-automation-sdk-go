/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Interface file for service: Reports
 * Used by client-side stubs.
 */
package hcl


// This interface provides methods to download information generated from the hardware compatibility feature residing on the vCenter Appliance. **Warning:** This interface is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
type ReportsClient interface {

    // Returns the location ReportsLocation information for downloading a compatibility report. **Warning:** This method is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    //
    // @param reportParam identifier of hardware compatiblity report to be downloaded.
    // The parameter must be an identifier for the resource type: ``com.vmware.esx.hcl.resources.CompatibilityReport``.
    // @return ``Location`` class which includes the URI to file, short lived token and expiry of the token in the ReportsLocation object.
    // @throws NotFound if there is no report for the given id.
    // @throws Unauthenticated if the caller is not authenticated.
    // @throws ResourceInaccessible if the vCenter this API is executed on is not part of the Customer Experience Improvement Program (CEIP).
    // @throws Error If there is some unknown error. The accompanying error message will give more details about the failure.
	Get(reportParam string) (ReportsLocation, error)
}
