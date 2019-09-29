/*
 * Copyright 2019 VMware, Inc.  All rights reserved.
 */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Interface file for service: VM
 * Used by client-side stubs.
 */

package VM

import (
)

// The ``VM`` interface provides methods for managing the lifecycle of a virtual machine.
type VMClient interface {


    // Creates a virtual machine.
    //
    // @param specParam Virtual machine specification.
    // @return ID of newly-created virtual machine.
    // The return value will be an identifier for the resource type: ``VirtualMachine``.
    // @throws AlreadyExists if a virtual machine with the specified name already exists.
    // @throws Error if the system reports an error while responding to the request.
    // @throws InvalidArgument if any of the specified parameters are invalid.
    // @throws NotFound if any of the resources specified in spec could not be found
    // @throws ResourceInaccessible if a specified resource (eg. host) is not accessible.
    // @throws ResourceInUse if any of the specified storage addresses (eg. IDE, SATA, SCSI, NVMe) result in a storage address conflict.
    // @throws ServiceUnavailable if the system is unable to communicate with a service to complete the request.
    // @throws UnableToAllocateResource if any of the resources needed to create the virtual machine could not be allocated.
    // @throws Unauthenticated if the user can not be authenticated.
    // @throws Unauthorized if the user doesn't have the required privileges.
    // @throws Unsupported if ``guestOS`` is not supported for the requested virtual hardware version and spec includes null properties that default to guest-specific values.
    Create(specParam CreateSpec) (string, error) 


    // Returns information about a virtual machine.
    //
    // @param vmParam Virtual machine identifier.
    // The parameter must be an identifier for the resource type: ``VirtualMachine``.
    // @return Information about the specified virtual machine.
    // @throws Error if the system reports an error while responding to the request.
    // @throws NotFound if the virtual machine is not found.
    // @throws ResourceInaccessible if the virtual machine's configuration state cannot be accessed.
    // @throws ServiceUnavailable if the system is unable to communicate with a service to complete the request.
    // @throws Unauthenticated if the user can not be authenticated.
    // @throws Unauthorized if the user doesn't have the required privileges.
    Get(vmParam string) (Info, error) 


    // Deletes a virtual machine.
    //
    // @param vmParam Virtual machine identifier.
    // The parameter must be an identifier for the resource type: ``VirtualMachine``.
    // @throws Error if the system reports an error while responding to the request.
    // @throws NotAllowedInCurrentState if the virtual machine is running (powered on).
    // @throws NotFound if the virtual machine is not found.
    // @throws ResourceBusy if the virtual machine is busy performing another operation.
    // @throws ResourceInaccessible if the virtual machine's configuration state cannot be accessed.
    // @throws ServiceUnavailable if the system is unable to communicate with a service to complete the request.
    // @throws Unauthenticated if the user can not be authenticated.
    // @throws Unauthorized if the user doesn't have the required privileges.
    Delete(vmParam string) error 


    // Returns information about at most 4000 visible (subject to permission checks) virtual machines in vCenter matching the FilterSpec.
    //
    // @param filterParam Specification of matching virtual machines for which information should be returned.
    // If null, the behavior is equivalent to a FilterSpec with all properties null which means all virtual machines match the filter.
    // @return Commonly used information about the virtual machines matching the FilterSpec.
    // @throws InvalidArgument if the FilterSpec#powerStates property contains a value that is not supported by the server.
    // @throws UnableToAllocateResource if more than 4000 virtual machines match the FilterSpec.
    // @throws ServiceUnavailable if the system is unable to communicate with a service to complete the request.
    // @throws Unauthenticated if the user can not be authenticated.
    // @throws Unauthorized if the user doesn't have the required privileges.
    List(filterParam *FilterSpec) ([]Summary, error) 

}
