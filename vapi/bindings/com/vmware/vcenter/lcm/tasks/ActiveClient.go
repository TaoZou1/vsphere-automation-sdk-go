/* Copyright © 2019 VMware, Inc. All Rights Reserved.
     SPDX-License-Identifier: BSD-2-Clause */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Interface file for service: Active
 * Used by client-side stubs.
 */

package tasks

import (
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/bindings/com/vmware/vcenter/lcm"
)

type ActiveClient interface {


    // Retrieves information of the tasks that are being executed. **Warning:** This method is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    // @return The information of the tasks that are being executed.
    List() ([]lcm.TasksInfo, error) 

}
