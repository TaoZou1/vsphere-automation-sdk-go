/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Interface file for service: Validation
 * Used by client-side stubs.
 */
package lcm


// The service that provides validation of a section of full deployment specification. **Warning:** This interface is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
type ValidationClient interface {

    // Validate the name of the appliance to be deployed. 
    //
    // #. 1. Return False if the there is already an appliance that has the same name as given in the spec exist in the path.
    //
    // . **Warning:** This method is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    //
    // @param specParam  The configuration needed to validate the name of the appliance to be deployed.
    // @return False if the name of the appliance already exists. True otherwise.
	CheckApplianceName(specParam ValidationApplianceNameRequest) (bool, error)
}
