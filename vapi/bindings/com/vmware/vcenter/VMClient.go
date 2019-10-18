/* Copyright © 2019 VMware, Inc. All Rights Reserved.
     SPDX-License-Identifier: BSD-2-Clause */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Interface file for service: VM
 * Used by client-side stubs.
 */

package vcenter

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
    Create(specParam VMCreateSpec) (string, error) 


    // Creates a virtual machine from an existing virtual machine. Each virtual disk of the cloned virtual machines is assigned a storage policy based on the following ordered rules:
    //
    // #. If a virtual disk is specified in VMCloneSpec#disksToUpdate with a VMDiskCloneSpec#storagePolicy, the specified storage policy is assigned to the virtual disk.
    // #. If a storage policy is specified in VMCloneSpec#diskStoragePolicy, the specified storage policy is assigned.
    // #. Otherwise, the system sets a storage policy for the virtual disk as follows:
    //
    //     #. If the corresponding source disk storage policy is compatible with the target datastore, the virtual disk is assigned the same source disk storage policy. If the source disk is using the default storage profile, the storage policy is always compatible so the virtual disk will be assigned the target datastore default storage policy.
    //     #. If the VMCloneSpec#fallbackToDatastoreDefaultPolicy parameter is set to true, the virtual disk is assigned the target datastore default policy, otherwise the server will throw an errors.InvalidArgument exception.
    //
    //  
    //
    //  Similarly, virtual machine home of the cloned virtual machine is assigned a storage policy based on the following ordered rules:
    //
    // #. If a storage policy is specified in VMCloneSpec#vmHomeStoragePolicy, the specified storage policy is assigned to the virtual machine home.
    // #. Otherwise, the system sets a storage policy for the virtual machine home as follows: 
    //
    //     #. If the source home storage policy is compatible with the target datastore, the virtual machine home is assigned the same source home storage policy. If the source home is using the default storage profile, the storage policy is always compatible so the virtual machine home will be assigned the target datastore default storage policy.
    //     #. If the VMCloneSpec#fallbackToDatastoreDefaultPolicy parameter is set to true, the virtual machine home is assigned the target datastore default policy, otherwise the server will throw an errors.InvalidArgument exception.
    //
    //  
    //
    // . **Warning:** This method is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    //
    // @param specParam Virtual machine clone specification.
    // @return ID of newly-created virtual machine.
    // The return value will be an identifier for the resource type: ``VirtualMachine``.
    // @throws AlreadyExists if a virtual machine with the specified name already exists.
    // @throws Error if the system reports an error while responding to the request.
    // @throws InvalidArgument if any of the specified parameters are invalid.
    // @throws NotFound if any of the resources specified in spec could not be found
    // @throws ResourceInaccessible if a specified resource (eg. host) is not accessible.
    // @throws ServiceUnavailable if the system is unable to communicate with a service to complete the request.
    // @throws UnableToAllocateResource if any of the resources needed to clone the virtual machine could not be allocated.
    // @throws Unauthenticated if the user can not be authenticated.
    // @throws Unauthorized if the user doesn't have the required privileges.
    Clone(specParam VMCloneSpec) (string, error) 


    // Relocates a virtual machine based on the specification. The parts of the virtual machine that can move are: FOLDER, RESOURCE_POOL, HOST, CLUSTER and DATASTORE of home of the virtual machine and disks. Each virtual disk of the virtual machine is assigned a storage policy based on the following ordered rules:
    //
    // #. If a virtual disk is specified in VMRelocateSpec#disks with a VMDiskRelocateSpec#storagePolicy, the specified storage policy is assigned to the virtual disk.
    // #. If a storage policy is specified in VMRelocateSpec#diskStoragePolicy, the specified storage policy is assigned.
    // #. Otherwise, the system sets a storage policy for the virtual disk as follows:
    //
    //     #. If the current datastore of the virtual disk is the same as target datastore, the virtual disk keeps its current storage policy.
    //     #. If the current disk storage policy is compatible with the target datastore, the virtual disk is assigned the same storage policy. If the virtual disk is currently associated with the datastore default storage policy, the storage policy is always compatible so the virtual disk will be assigned the target datastore default storage policy.
    //     #. If the VMRelocateSpec#fallbackToDatastoreDefaultPolicy parameter is set to true, the virtual disk is assigned the target datastore default policy, otherwise the server will throw an errors.InvalidArgument exception.
    //
    //  
    //
    //  Similarly, virtual machine home is assigned a storage policy based on the following ordered rules:
    //
    // #. If a storage policy is specified in VMRelocateSpec#vmHomeStoragePolicy, the specified storage policy is assigned to the virtual machine home.
    // #. Otherwise, the system sets a storage policy for the virtual machine home as follows: 
    //
    //     #. If the current datastore of the virtual machine home is the same as target datastore, the virtual machine home keeps its current storage policy.
    //     #. If the current home storage policy is compatible with the target datastore, the virtual machine home is assigned the same home storage policy. If the virtual machine home is currently associated with the datastore default storage profile, the storage policy is always compatible so the virtual machine home will be assigned the target datastore default storage policy.
    //     #. If the VMRelocateSpec#fallbackToDatastoreDefaultPolicy parameter is set to true, the virtual machine home is assigned the target datastore default policy, otherwise the server will throw an errors.InvalidArgument exception.
    //
    //  
    //
    // . **Warning:** This method is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
    //
    // @param vmParam Existing Virtual machine to relocate.
    // The parameter must be an identifier for the resource type: ``VirtualMachine``.
    // @param specParam Relocate specification.
    // @throws Error if the system reports an error while responding to the request.
    // @throws InvalidArgument if any of the specified parameters are invalid.
    // @throws NotFound if any of the resources specified in spec or the given "vm" could not be found
    // @throws ResourceInaccessible if a specified resource (eg. host) is not accessible.
    // @throws ServiceUnavailable if the system is unable to communicate with a service to complete the request.
    // @throws Unauthenticated if the user can not be authenticated.
    // @throws Unauthorized if the user doesn't have the required privileges.
    Relocate(vmParam string, specParam VMRelocateSpec) error 


    // Create an instant clone of an existing virtual machine. This method was added in vSphere API 6.7.1.
    //
    // @param specParam Virtual machine InstantCloneSpec.
    // @return ID of newly-created virtual machine.
    // The return value will be an identifier for the resource type: ``VirtualMachine``.
    // @throws AlreadyExists if a virtual machine with the specified name already exists.
    // @throws Error if the system reports an error while responding to the request.
    // @throws InvalidArgument if any of the specified parameters are invalid.
    // @throws NotFound if any of the resources specified in spec could not be found
    // @throws ResourceInaccessible if a specified resource (eg. host) is not accessible.
    // @throws ServiceUnavailable if the system is unable to communicate with a service to complete the request.
    // @throws UnableToAllocateResource if any of the resources needed to create an instant clone could not be allocated.
    // @throws Unauthenticated if the user can not be authenticated.
    // @throws Unauthorized if the user doesn't have the required privileges.
    InstantClone(specParam VMInstantCloneSpec) (string, error) 


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
    Get(vmParam string) (VMInfo, error) 


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


    // Returns information about at most 4000 visible (subject to permission checks) virtual machines in vCenter matching the VMFilterSpec.
    //
    // @param filterParam Specification of matching virtual machines for which information should be returned.
    // If null, the behavior is equivalent to a VMFilterSpec with all properties null which means all virtual machines match the filter.
    // @return Commonly used information about the virtual machines matching the VMFilterSpec.
    // @throws InvalidArgument if the VMFilterSpec#powerStates property contains a value that is not supported by the server.
    // @throws UnableToAllocateResource if more than 4000 virtual machines match the VMFilterSpec.
    // @throws ServiceUnavailable if the system is unable to communicate with a service to complete the request.
    // @throws Unauthenticated if the user can not be authenticated.
    // @throws Unauthorized if the user doesn't have the required privileges.
    List(filterParam *VMFilterSpec) ([]VMSummary, error) 


    // Creates a virtual machine from existing virtual machine files on storage. This method was added in vSphere API 6.8.7.
    //
    // @param specParam Specification of the location of the virtual machine files and the placement of the new virtual machine.
    // @return Identifier of the newly-created virtual machine.
    // The return value will be an identifier for the resource type: ``VirtualMachine``.
    // @throws AlreadyExists if a virtual machine with the specified name already exists or if a virtual machine using the specified virtual machine files already exists.
    // @throws Error if the system reports an error while responding to the request.
    // @throws InvalidArgument if any of the specified parameters are invalid.
    // @throws NotFound if any of the resources specified in spec could not be found.
    // @throws ResourceInaccessible if a specified resource (eg. host) is not accessible.
    // @throws ServiceUnavailable if the system is unable to communicate with a service to complete the request.
    // @throws UnableToAllocateResource if any of the resources needed to register the virtual machine could not be allocated.
    // @throws Unauthenticated if the user cannot be authenticated.
    // @throws Unauthorized if the user doesn't have the required privileges.
    Register(specParam VMRegisterSpec) (string, error) 


    // Removes the virtual machine corresponding to ``vm`` from the vCenter inventory without removing any of the virtual machine's files from storage. All high-level information stored with the management server (ESXi or vCenter) is removed, including information such as statistics, resource pool association, permissions, and alarms. This method was added in vSphere API 6.8.7.
    //
    // @param vmParam Identifier of the virtual machine to be unregistered.
    // The parameter must be an identifier for the resource type: ``VirtualMachine``.
    // @throws NotAllowedInCurrentState if the virtual machine is running (powered on).
    // @throws NotFound if there is no virtual machine associated with ``vm`` in the system.
    // @throws ResourceBusy if the virtual machine is busy performing another operation.
    // @throws ServiceUnavailable if the system is unable to communicate with a service to complete the request.
    // @throws Unauthenticated if the user can not be authenticated.
    // @throws Unauthorized if the user doesn't have the required privileges.
    Unregister(vmParam string) error 

}