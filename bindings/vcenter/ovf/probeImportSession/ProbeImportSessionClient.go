/*
 * Copyright 2019 VMware, Inc.  All rights reserved.
 */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Interface file for service: ProbeImportSession
 * Used by client-side stubs.
 */

package probeImportSession

import (
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/bindings/vcenter/ovf/importSession"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/data"
)

// The ``ProbeImportSession`` interface provides methods to dry-run an import of an OVF package without having selected a target. This can be used for querying instantiation parameters such as name and EULA that are not dependent on the target. Parameters that depend on the target such as networks are omitted. 
//
//  Once a target has been chosen, a new session must be created with ImportSession. The probe session cannot be instantiated or altered to refer to a target.
type ProbeImportSessionClient interface {


    // Creates a probe import session. This can be used for inspecting various instantiation parameters using ProbeImportSession#tryInstantiate. 
    //
    //  This operation is idempotent. Two calls with the same client token will result in the same session.
    //
    // @param clientTokenParam  The client token that ensures idempotency. The client token must be a valid UUID.
    // If not specified, then it will be generated by the system.
    // @param createSpecParam  Creation spec containing a set of parameters to control the behavior of the import.
    // The parameter must contain all the properties defined in importSession.CreateSpec.
    // @return the identifier of the probe import session.
    // The return value will be an identifier for the resource type: ``com.vmware.vcenter.OvfProbeImportSession``.
    // @throws InvalidArgument  if the provided createSpec contain an invalid argument.
    // @throws NotFound  if a specified source or destinaton for the OVF package does not exist.
    // @throws ResourceInaccessible  if there was an error accessing the OVF package at the specified source.
    // @throws com.vmware.vapi.std.errors.Unauthorized if you do not have all of the privileges described as follows: 
    //
    // * Method execution requires ``System.Read``.
    // * The resource ``com.vmware.content.library.Item`` referenced by the property importSession.CreateSpec#contentLibraryItem requires ``System.Read``.
    CreateProbeImportSession(clientTokenParam *string, createSpecParam *data.StructValue) (string, error) 


    // Retrieves a probe import session.
    //
    // @param idParam  the probe import session ID.
    // The parameter must be an identifier for the resource type: ``com.vmware.vcenter.OvfProbeImportSession``.
    // @return Returns information about a probe session.
    // @throws NotFound  if the specified session could not be found. It may have been deleted or timed out.
    // @throws com.vmware.vapi.std.errors.Unauthorized if you do not have all of the privileges described as follows: 
    //
    // * Method execution requires ``System.Read``.
    // * The resource ``com.vmware.vcenter.OvfProbeImportSession`` referenced by the parameter ``id`` requires ``System.Read``.
    Get(idParam string) (Info, error) 


    // When a probe import session is in the PROBE_IMPORT_SELECTING_OVF_PARAMS state, this method returns information about instantiation parameters that can be specified for the probe session. The instantiation parameters returned depend on the choices specified in the instantiation parameters provided as input arguments. 
    //
    //  This method can be iteratively called with progressively more detailed instantiation parameters. 
    //
    //  The operation does not change the state of the probe session. The call must supply all instantiation parameters each time. They are not remembered from one invocation to the next. This also means that the operation is naturally idempotent.
    //
    // @param idParam  the probe import session ID.
    // The parameter must be an identifier for the resource type: ``com.vmware.vcenter.OvfProbeImportSession``.
    // @param instantiationParametersParam  a list of OVF parameter structures that specifies a set of deployment specific parameters.
    // The parameter must contain all the properties defined in ovf.OvfParams.
    // @return A list of OVF parameters structures that provides information about the OVF package, along with any warnings or errors.
    // @throws NotFound  if the specified session could not be found. It may have been deleted or timed out.
    // @throws NotAllowedInCurrentState  if the specified session is not in the IMPORT_SELECTING_OVF_PARAMS state.
    // @throws NotAllowedInCurrentState  if the specified session is not in the IMPORT_SELECTING_OVF_PARAMS state.
    // @throws Unauthorized  If you do not have all of the privileges described in the following list: 
    //
    // * Operation execution requires VirtualMachine.Config.AddNewDisk if ovf has disk drive (type 17) section.
    // * Operation execution requires VirtualMachine.Config.AdvancedConfig if ovf has ExtraConfig section.
    // * Operation execution requires Extension.Register for specified resource group if ovf has vServiceDependency section.
    // * Operation execution requires Network.Assign for target network.
    // * Operation execution requires Datastore.AllocateSpace for target datastore.
    // @throws com.vmware.vapi.std.errors.Unauthorized if you do not have all of the privileges described as follows: 
    //
    // * Method execution requires ``System.Read``.
    // * The resource ``com.vmware.vcenter.OvfProbeImportSession`` referenced by the parameter ``id`` requires ``System.Read``.
    TryInstantiate(idParam string, instantiationParametersParam []*data.StructValue) (importSession.OvfValidationResult, error) 


    // Deletes (or cancels) a probe import session. 
    //
    //  A client should terminate a session to free up resources on the server. A probe session will automatically be reclaimed after a period of inactivity. For example it will only remain in PROBE_IMPORT_ERROR state for a period of time. 
    //
    //  Deleting an already deleted session has no effect.
    //
    // @param idParam  the probe import session ID.
    // The parameter must be an identifier for the resource type: ``com.vmware.vcenter.OvfProbeImportSession``.
    // @throws com.vmware.vapi.std.errors.Unauthorized if you do not have all of the privileges described as follows: 
    //
    // * Method execution requires ``System.Read``.
    // * The resource ``com.vmware.vcenter.OvfProbeImportSession`` referenced by the parameter ``id`` requires ``System.Read``.
    Delete(idParam string) error 

}
