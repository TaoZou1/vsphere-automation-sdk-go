/* Copyright © 2019 VMware, Inc. All Rights Reserved.
     SPDX-License-Identifier: BSD-2-Clause */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Interface file for service: Tickets
 * Used by client-side stubs.
 */

package console

import (
)

// The ``Tickets`` interface provides methods for managing the virtual machine console tickets. **Warning:** This interface is available as Technology Preview. These are early access APIs provided to test, automate and provide feedback on the feature. Since this can change based on feedback, VMware does not guarantee backwards compatibility and recommends against using them in production environments. Some Technology Preview APIs might only be applicable to specific environments.
type TicketsClient interface {


    // Creates a virtual machine console ticket of a given ticket type. The created ticket is a one time use URI. The validity of the ticket is 30 minutes, if not used with in the time frame the ticket expires. 
    //
    //  The TicketsType#Type_VMRC ticket contains the IP address or the DNS resolvable name of the vCenter server. This ticket requires installation of VMware Workstation, VMware Fusion or VMRC to be installed on the machine where the ticket has to be opened. This ticket can be acquired even when the VM is turned off. 
    //
    //  The TicketsType#Type_WEBMKS ticket contains the IP address of the DNS resolvable name of the ESX server. This ticket requires user to embed this ticket in a HTML page using VMware HTML Console SDK - https://www.vmware.com/support/developer/html-console This ticket can be acquired only when the VM is turned on.. **Warning:** This method is available as Technology Preview. These are early access APIs provided to test, automate and provide feedback on the feature. Since this can change based on feedback, VMware does not guarantee backwards compatibility and recommends against using them in production environments. Some Technology Preview APIs might only be applicable to specific environments.
    //
    // @param vmParam Virtual machine identifier.
    // The parameter must be an identifier for the resource type: ``VirtualMachine``.
    // @param specParam Specification for the console ticket to be created.
    // @return Commonly used information about the virtual machine console ticket.
    // @throws Error if the system reports an error while responding to the request.
    // @throws InvalidArgument if the {\\\\@link CreateSpec#type) property contains a value that is not supported by the server.
    // @throws NotAllowedInCurrentState if the virtual machine is powered off and requested ticket type is TicketsType#Type_WEBMKS
    // @throws NotFound if the virtual machine is not found.
    // @throws ResourceInaccessible if the virtual machine's configuration or execution state cannot be accessed.
    // @throws ServiceUnavailable if the system is unable to communicate with a service to complete the request.
    // @throws Unauthenticated if the user can not be authenticated.
    // @throws Unauthorized if the user doesn't have the required privileges.
    Create(vmParam string, specParam TicketsCreateSpec) (TicketsSummary, error) 

}
