/*
 * Copyright 2019 VMware, Inc.  All rights reserved.
 */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Interface file for service: Reports
 * Used by client-side stubs.
 */

package reports

import (
)

// The ``Report`` interface provides an method to download the report generated by the interop and precheck operations. To download a report, you must use the ``get()`` method. A ``Report.Location`` class represents the location of the file which has ``Report.Token`` class which represents the token ID (\\\\@name String) and the expiration time of the token ``Report.Token.expiry`` DateTime. ``Report.Location`` class also has the URI for the file which needs to be downloaded.
type ReportsClient interface {


    // Returns the location Location information for downloading the report for the specified file name. 
    //
    //  Retrieving a report involves two steps: 
    //
    // * Step 1: Invoke the Reports#get method to provision a token and a URI.
    // * Step 2: Make an HTTP GET request by using the URI and the token returned in step 1 to retrieve the report.
    //
    //  
    //
    //  The HTTP GET request will: 
    //
    // * Return 401 (Not Authorized) if the download URI is recognized, but the token is invalid, 404 if the URL is not recognized otherwise return 200 (OK)
    // * Provide the CSV contents as the output of the request. The API accepts the file name as input, reads the contents of that CSV file, and returns this text as the result of the API.
    //
    // @param reportParam 
    // The parameter must be an identifier for the resource type: ``com.vmware.vcenter.lcm.report``.
    // @throws Unauthenticated if the user can not be authenticated.
    // @throws NotFound If there is no file associated with ``report`` in the system.
    // @throws Error If there is some unknown internal error. The accompanying error message will give more details about the failure.
    Get(reportParam string) (Location, error) 

}
