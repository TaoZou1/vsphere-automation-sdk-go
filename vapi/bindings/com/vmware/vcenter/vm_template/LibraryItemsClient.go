/* Copyright © 2019 VMware, Inc. All Rights Reserved.
     SPDX-License-Identifier: BSD-2-Clause */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Interface file for service: LibraryItems
 * Used by client-side stubs.
 */

package vm_template

import (
)

// The ``LibraryItems`` interface provides methods to deploy virtual machines from library items containing virtual machine templates, as well as methods to create library items containing virtual machine templates. The ``LibraryItems`` interface also provides an operation to retrieve information about the template contained in the library item.
type LibraryItemsClient interface {


    // Creates a library item in content library from a virtual machine. This method creates a library item in content library whose content is a virtual machine template created from the source virtual machine, using the supplied create specification. The virtual machine template is stored in a newly created library item.
    //
    // @param specParam  information used to create the library item from the source virtual machine.
    // @return Identifier of the newly created library item.
    // The return value will be an identifier for the resource type: ``com.vmware.content.library.Item``.
    // @throws AlreadyExists  if an entity with the name specified by LibraryItemsCreateSpec#name already exists in the folder specified by LibraryItemsCreatePlacementSpec#folder.
    // @throws Error  if the system reports an error while responding to the request.
    // @throws InvalidArgument  if LibraryItemsCreateSpec contains invalid arguments.
    // @throws NotAllowedInCurrentState  if the operation cannot be performed because of the source virtual machine's current state.
    // @throws NotFound  if the source virtual machine specified by LibraryItemsCreateSpec#sourceVm does not exist.
    // @throws NotFound  if the library specified by LibraryItemsCreateSpec#library does not exist.
    // @throws ResourceInaccessible  if there was an error accessing a file from the source virtual machine.
    // @throws ResourceInUse  if the source virtual machine is busy.
    // @throws ServiceUnavailable  if any of the services involved in the method are unavailable.
    // @throws UnableToAllocateResource  if any of the resources needed to create the virtual machine template could not be allocated.
    // @throws Unauthenticated  if the user that requested the method cannot be authenticated.
    // @throws Unauthorized  if the user that requested the method is not authorized to perform the method.
    Create(specParam LibraryItemsCreateSpec) (string, error) 


    // Deploys a virtual machine as a copy of the source virtual machine template contained in the library item specified by ``template_library_item``. It uses the deployment specification in ``spec``. If LibraryItemsDeploySpec#poweredOn and/or LibraryItemsDeploySpec#guestCustomization are specified, the server triggers the power on and/or guest customization operations, which are executed asynchronously.
    //
    // @param templateLibraryItemParam  identifier of the content library item containing the source virtual machine template to be deployed.
    // The parameter must be an identifier for the resource type: ``com.vmware.content.library.Item``.
    // @param specParam  specification of how the virtual machine should be deployed.
    // @return Identifier of the deployed virtual machine.
    // The return value will be an identifier for the resource type: ``VirtualMachine``.
    // @throws AlreadyExists  if a virtual machine with the name specified by LibraryItemsDeploySpec#name already exists.
    // @throws Error  if the system reports an error while responding to the request.
    // @throws InvalidArgument  if ``spec`` contains invalid arguments.
    // @throws NotAllowedInCurrentState  if either a specified host or a specified datastore is in an invalid state for the deployment, such as maintenance mode.
    // @throws NotFound  if the library item specified by ``template_library_item`` cannot be found.
    // @throws NotFound  if any resource specified by a property of the LibraryItemsDeploySpec class, specified by ``spec`` cannot be found.
    // @throws ResourceInaccessible  if there was an error accessing the source virtual machine template contained in the library item specified by ``template_library_item``.
    // @throws ResourceInaccessible  if there an error accessing any of the resources specified in the ``spec``.
    // @throws ServiceUnavailable  if any of the services involved in the method are unavailable.
    // @throws UnableToAllocateResource  if there was an error in allocating any of the resources required by the method.
    // @throws Unauthenticated  if the user that requested the method cannot be authenticated.
    // @throws Unauthorized  if the user that requested the method is not authorized to perform the method.
    Deploy(templateLibraryItemParam string, specParam LibraryItemsDeploySpec) (string, error) 


    // Returns information about a virtual machine template contained in the library item specified by ``template_library_item``
    //
    // @param templateLibraryItemParam  identifier of the library item containing the virtual machine template.
    // The parameter must be an identifier for the resource type: ``com.vmware.content.library.Item``.
    // @return Information about the virtual machine template item contained in the library item.
    // If null, the library item specified by ``template_library_item`` does not contain a virtual machine template.
    // @throws Error  if the system reports an error while responding to the request.
    // @throws NotFound  if the library item could not be found.
    // @throws ResourceInaccessible  if the virtual machine template's configuration state cannot be accessed.
    // @throws ServiceUnavailable  if any of the services involved in the method are unavailable.
    // @throws Unauthenticated  if the user that requested the method cannot be authenticated.
    // @throws Unauthorized  if the user that requested the method is not authorized to perform the method.
    Get(templateLibraryItemParam string) (*LibraryItemsInfo, error) 

}
