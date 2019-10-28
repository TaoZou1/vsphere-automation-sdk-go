/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Interface file for service: Management
 * Used by client-side stubs.
 */
package vstats


// The ``Management`` interface provides methods to perform various management related operations. The service provides access to the database settings of the system.
type ManagementClient interface {

    // Returns the configuration information of databases.
    // @return Configuration information of databases.
    // @throws Error if the system reports an error while responding to the request.
    // @throws Unauthenticated if the user can not be authenticated.
    // @throws Unauthorized if the user does not have sufficient privileges.
	Get() (ManagementInfo, error)
}
