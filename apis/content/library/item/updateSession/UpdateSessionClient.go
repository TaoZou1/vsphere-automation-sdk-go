/*
 * Copyright 2019 VMware, Inc.  All rights reserved.
 */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Interface file for service: UpdateSession
 * Used by client-side stubs.
 */

package updatesession

import (
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/apis/content/library/item"
)

// The ``UpdateSession`` interface manipulates sessions that are used to upload content into the Content Library Service, and/or to remove files from a library item. 
//
//  An update session is a resource which tracks changes to content. An update session is created with a set of files that are intended to be uploaded to a specific library.ItemModel, or removed from an item. The session object can be used to track the uploads and inspect the changes that are being made to the item by that upload. It can also serve as a channel to check on the result of the upload, and status messages such as errors and warnings for the upload. 
//
//  Modifications are not visible to other clients unless the session is completed and all necessary files have been received. 
//
//  The management of the files within the session is done through the File interface.
type UpdateSessionClient interface {


    // Creates a new update session. An update session is used to make modifications to a library item. Modifications are not visible to other clients unless the session is completed and all necessary files have been received. 
    //
    //  Content Library Service allows only one single update session to be active for a specific library item.
    //
    // @param clientTokenParam  Unique token generated by the client for each creation request. The token should be a universally unique identifier (UUID), for example: ``b8a2a2e3-2314-43cd-a871-6ede0f429751``. This token can be used to guarantee idempotent creation.
    // If not specified creation is not idempotent.
    // @param createSpecParam  Specification for the new update session to be created.
    // @return Identifier of the new update session being created.
    // The return value will be an identifier for the resource type: ``com.vmware.content.library.item.UpdateSession``.
    // @throws InvalidArgument  if the session specification is not valid.
    // @throws InvalidArgument  if the ``client_token`` does not conform to the UUID format.
    // @throws InvalidElementType  if the update session is being created on a subscribed library item.
    // @throws NotFound  if the item targeted for update does not exist.
    // @throws ResourceBusy  if there is another update session on the same library item.
    // @throws com.vmware.vapi.std.errors.Unauthorized if you do not have all of the privileges described as follows: 
    //
    // * The resource ``com.vmware.content.library.Item`` referenced by the property item.UpdateSessionModel#libraryItemId requires ``ContentLibrary.UpdateSession``.
    Create(clientTokenParam *string, createSpecParam item.UpdateSessionModel) (string, error) 


    // Gets the update session with the specified identifier, including the most up-to-date status information for the session.
    //
    // @param updateSessionIdParam  Identifier of the update session to retrieve.
    // The parameter must be an identifier for the resource type: ``com.vmware.content.library.item.UpdateSession``.
    // @return The item.UpdateSessionModel instance with the given ``update_session_id``.
    // @throws NotFound  if no update session with the given identifier exists.
    // @throws com.vmware.vapi.std.errors.Unauthorized if you do not have all of the privileges described as follows: 
    //
    // * Method execution requires ``System.Anonymous``.
    Get(updateSessionIdParam string) (item.UpdateSessionModel, error) 


    // Lists the identifiers of the update session created by the calling user. Optionally may filter by library item.
    //
    // @param libraryItemIdParam  Optional library item identifier on which to filter results.
    // The parameter must be an identifier for the resource type: ``com.vmware.content.library.item.UpdateSession``.
    // If not specified the results are not filtered.
    // @return The array of identifiers of all update sessions created by the calling user.
    // The return value will contain identifiers for the resource type: ``com.vmware.content.library.item.UpdateSession``.
    // @throws NotFound  if a library item identifier is given for an item which does not exist.
    // @throws com.vmware.vapi.std.errors.Unauthorized if you do not have all of the privileges described as follows: 
    //
    // * The resource ``com.vmware.content.library.item.UpdateSession`` referenced by the parameter ``library_item_id`` requires ``ContentLibrary.UpdateSession``.
    List(libraryItemIdParam *string) ([]string, error) 


    // Completes the update session. This indicates that the client has finished making all the changes required to the underlying library item. If the client is pushing the content to the server, the library item will be updated once this call returns. If the server is pulling the content, the call may return before the changes become visible. In that case, the client can track the session to know when the server is done. 
    //
    //  This method requires the session to be in the item.UpdateSessionModel_State#State_ACTIVE state. 
    //
    //  Depending on the type of the library item associated with this session, a type adapter may be invoked to verify the validity of the files uploaded. The user can explicitly validate the session before completing the session by using the File#validate method. 
    //
    //  Modifications are not visible to other clients unless the session is completed and all necessary files have been received.
    //
    // @param updateSessionIdParam  Identifier of the update session that should be completed.
    // The parameter must be an identifier for the resource type: ``com.vmware.content.library.item.UpdateSession``.
    // @throws NotFound  if no update session with the given identifier exists.
    // @throws NotAllowedInCurrentState  if the update session is not in the item.UpdateSessionModel_State#State_ACTIVE state, or if some of the files that will be uploaded by the client aren't received correctly.
    // @throws com.vmware.vapi.std.errors.Unauthorized if you do not have all of the privileges described as follows: 
    //
    // * Method execution requires ``System.Anonymous``.
    Complete(updateSessionIdParam string) error 


    // Keeps an update session alive. 
    //
    //  If there is no activity for an update session after a period of time, the update session will expire, then be deleted. The update session expiration timeout is configurable in the Content Library Service system configuration. The default is five minutes. Invoking this method enables a client to specifically extend the lifetime of the update session.
    //
    // @param updateSessionIdParam  Identifier of the update session whose lifetime should be extended.
    // The parameter must be an identifier for the resource type: ``com.vmware.content.library.item.UpdateSession``.
    // @param clientProgressParam  Optional update to the progress property of the session. If specified, the new progress should be greater then the current progress. See item.UpdateSessionModel#clientProgress.
    // If not specified the progress is not updated.
    // @throws NotFound  if no update session with the given identifier exists.
    // @throws NotAllowedInCurrentState  if the update session is not in the item.UpdateSessionModel_State#State_ACTIVE state.
    // @throws com.vmware.vapi.std.errors.Unauthorized if you do not have all of the privileges described as follows: 
    //
    // * Method execution requires ``System.Anonymous``.
    KeepAlive(updateSessionIdParam string, clientProgressParam *int64) error 


    // Cancels the update session and sets its state to item.UpdateSessionModel_State#State_CANCELED. This method will free up any temporary resources currently associated with the session. 
    //
    //  This method is not allowed if the session has been already completed. 
    //
    //  Cancelling an update session will cancel any in progress transfers (either uploaded by the client or pulled by the server). Any content that has been already received will be scheduled for deletion.
    //
    // @param updateSessionIdParam  Identifier of the update session that should be canceled.
    // The parameter must be an identifier for the resource type: ``com.vmware.content.library.item.UpdateSession``.
    // @throws NotFound  if no update session with the given identifier exists.
    // @throws NotAllowedInCurrentState  if the update session is not in the item.UpdateSessionModel_State#State_ACTIVE state.
    // @throws com.vmware.vapi.std.errors.Unauthorized if you do not have all of the privileges described as follows: 
    //
    // * Method execution requires ``System.Anonymous``.
    Cancel(updateSessionIdParam string) error 


    // Terminates the update session with a client specified error message. 
    //
    //  This is useful in transmitting client side failures (for example, not being able to access a file) to the server side.
    //
    // @param updateSessionIdParam  Identifier of the update session to fail.
    // The parameter must be an identifier for the resource type: ``com.vmware.content.library.item.UpdateSession``.
    // @param clientErrorMessageParam  Client side error message. This can be useful in providing some extra details about the client side failure. Note that the message won't be translated to the user's locale.
    // @throws NotFound  if the update session does not exist.
    // @throws NotAllowedInCurrentState  if the update session is not in the item.UpdateSessionModel_State#State_ACTIVE state.
    // @throws com.vmware.vapi.std.errors.Unauthorized if you do not have all of the privileges described as follows: 
    //
    // * Method execution requires ``System.Anonymous``.
    Fail(updateSessionIdParam string, clientErrorMessageParam string) error 


    // Deletes an update session. This removes the session and all information associated with it. 
    //
    //  Removing an update session leaves any current transfers for that session in an indeterminate state (there is no guarantee that the server will terminate the transfers, or that the transfers can be completed). However there will no longer be a means of inspecting the status of those uploads except by seeing the effect on the library item. 
    //
    //  Update sessions for which there is no upload activity or which are complete will automatically be deleted after a period of time.
    //
    // @param updateSessionIdParam  Identifer of the update session to delete.
    // The parameter must be an identifier for the resource type: ``com.vmware.content.library.item.UpdateSession``.
    // @throws NotFound  if the update session does not exist.
    // @throws NotAllowedInCurrentState  if the update session is in the item.UpdateSessionModel_State#State_ACTIVE state.
    // @throws com.vmware.vapi.std.errors.Unauthorized if you do not have all of the privileges described as follows: 
    //
    // * Method execution requires ``System.Anonymous``.
    Delete(updateSessionIdParam string) error 


    // Updates the properties of an update session. 
    //
    //  This is an incremental update to the update session. Any property in the item.UpdateSessionModel class that is null will not be modified. 
    //
    //  This method will only update the property item.UpdateSessionModel#warningBehavior of the update session. This will not, for example, update the item.UpdateSessionModel#libraryItemId or item.UpdateSessionModel#state of an update session. 
    //
    //  This method requires the session to be in the item.UpdateSessionModel_State#State_ACTIVE state.
    //
    // @param updateSessionIdParam  Identifer of the update session that should be updated.
    // The parameter must be an identifier for the resource type: ``com.vmware.content.library.item.UpdateSession``.
    // @param updateSpecParam  Specification for the new property values to be set on the update session.
    // @throws NotFound  if the update session does not exist.
    // @throws NotAllowedInCurrentState  if the update session is not in the item.UpdateSessionModel_State#State_ACTIVE state.
    // @throws InvalidArgument  if the update session specification is not valid.
    // @throws com.vmware.vapi.std.errors.Unauthorized if you do not have all of the privileges described as follows: 
    //
    // * Method execution requires ``System.Anonymous``.
    Update(updateSessionIdParam string, updateSpecParam item.UpdateSessionModel) error 

}
