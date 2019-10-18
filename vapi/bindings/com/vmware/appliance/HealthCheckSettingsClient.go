/* Copyright © 2019 VMware, Inc. All Rights Reserved.
     SPDX-License-Identifier: BSD-2-Clause */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Interface file for service: HealthCheckSettings
 * Used by client-side stubs.
 */

package appliance

import (
)

// The ``HealthCheckSettings`` interface provides methods to enable/disable health check settings in vCenter Server. **Warning:** This interface is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
type HealthCheckSettingsClient interface {


    // Get health settings for vCenter Server. **Warning:** This method is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    // @return health settings.
    // @throws Error Generic error
    Get() (HealthCheckSettingsSettingSpec, error) 


    // Set health settings based on provided healthSpec. **Warning:** This method is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    //
    // @param healthSettingsParam 
    // @throws Error Generic error.
    Update(healthSettingsParam HealthCheckSettingsUpdateSpec) error 

}
