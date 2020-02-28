/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Code generated. DO NOT EDIT.

/*
 * Interface file for service: DownloadSession
 * Used by client-side stubs.
 */

package item

import (
	. "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/services/vsphere/content/library/item/model"
)

// The ``DownloadSession`` interface manipulates download sessions, which are used to download content from the Content Library Service.
//
//  A download session is an object that tracks the download of content (that is, downloading content from the Content Library Service) and acts as a lease to keep the download links available.
//
//  The File interface provides access to the download links.
type DownloadSessionClient interface {

	// Creates a new download session.
	//
	// @param clientTokenParam  A unique token generated by the client for each creation request. The token should be a universally unique identifier (UUID), for example: ``b8a2a2e3-2314-43cd-a871-6ede0f429751``. This token can be used to guarantee idempotent creation.
	// If not specified creation is not idempotent.
	// @param createSpecParam  Specification for the new download session to be created.
	// @return Identifier of the new download session being created.
	// The return value will be an identifier for the resource type: ``com.vmware.content.library.item.DownloadSession``.
	// @throws InvalidArgument  if the session specification is not valid.
	// @throws InvalidArgument  format.
	// @throws NotFound  if the library item targeted by the download does not exist.
	// @throws com.vmware.vapi.std.errors.Unauthorized if you do not have all of the privileges described as follows:
	//
	// * The resource ``com.vmware.content.library.Item`` referenced by the property DownloadSessionModel#libraryItemId requires ``ContentLibrary.DownloadSession``.
	Create(clientTokenParam *string, createSpecParam DownloadSessionModel) (string, error)

	// Gets the download session with the specified identifier, including the most up-to-date status information for the session.
	//
	// @param downloadSessionIdParam  Identifier of the download session to retrieve.
	// The parameter must be an identifier for the resource type: ``com.vmware.content.library.item.DownloadSession``.
	// @return The DownloadSessionModel instance with the given ``download_session_id``.
	// @throws NotFound  if no download session with the given ``download_session_id`` exists.
	// @throws com.vmware.vapi.std.errors.Unauthorized if you do not have all of the privileges described as follows:
	//
	// * Method execution requires ``System.Anonymous``.
	Get(downloadSessionIdParam string) (DownloadSessionModel, error)

	// Lists the identifiers of the download sessions created by the calling user. Optionally may filter by library item.
	//
	// @param libraryItemIdParam  Library item identifier on which to filter results.
	// The parameter must be an identifier for the resource type: ``com.vmware.content.library.item.DownloadSession``.
	// If not specified all download session identifiers are listed.
	// @return The array of identifiers of all download sessions created by the calling user.
	// The return value will contain identifiers for the resource type: ``com.vmware.content.library.item.DownloadSession``.
	// @throws NotFound  if a library item identifier is given for an item which does not exist.
	// @throws com.vmware.vapi.std.errors.Unauthorized if you do not have all of the privileges described as follows:
	//
	// * The resource ``com.vmware.content.library.item.DownloadSession`` referenced by the parameter ``library_item_id`` requires ``ContentLibrary.DownloadSession``.
	List(libraryItemIdParam *string) ([]string, error)

	// Keeps a download session alive. This operation is allowed only if the session is in the DownloadSessionModelState#State_ACTIVE state.
	//
	//  If there is no activity for a download session for a certain period of time, the download session will expire. The download session expiration timeout is configurable in the Content Library Service system configuration. The default is five minutes. Invoking this method enables a client to specifically extend the lifetime of an active download session.
	//
	// @param downloadSessionIdParam  Identifier of the download session whose lifetime should be extended.
	// The parameter must be an identifier for the resource type: ``com.vmware.content.library.item.DownloadSession``.
	// @param progressParam  Optional update to the progress property of the session. If specified, the new progress should be greater then the current progress. See DownloadSessionModel#clientProgress.
	// If not specified the progress is not updated.
	// @throws NotFound  if no download session with the given identifier exists.
	// @throws NotAllowedInCurrentState  if the download session is not in the DownloadSessionModelState#State_ACTIVE state.
	// @throws com.vmware.vapi.std.errors.Unauthorized if you do not have all of the privileges described as follows:
	//
	// * Method execution requires ``System.Anonymous``.
	KeepAlive(downloadSessionIdParam string, progressParam *int64) error

	// Cancels the download session. This method will abort any ongoing transfers and invalidate transfer urls that the client may be downloading from.
	//
	// @param downloadSessionIdParam  Identifer of the download session that should be canceled.
	// The parameter must be an identifier for the resource type: ``com.vmware.content.library.item.DownloadSession``.
	// @throws NotFound  if no download session with the given identifier exists.
	// @throws NotAllowedInCurrentState  if the download session is not in the DownloadSessionModelState#State_ACTIVE state.
	// @throws com.vmware.vapi.std.errors.Unauthorized if you do not have all of the privileges described as follows:
	//
	// * Method execution requires ``System.Anonymous``.
	Cancel(downloadSessionIdParam string) error

	// Deletes a download session. This removes the session and all information associated with it.
	//
	//  Removing a download session leaves any current transfers for that session in an indeterminate state (there is no guarantee that the transfers will be able to complete). However there will no longer be a means of inspecting the status of those downloads except by seeing the effect on the library item.
	//
	//  Download sessions for which there is no download activity or which are complete will automatically be expired and then deleted after a period of time.
	//
	// @param downloadSessionIdParam  Identifier of the download session to be deleted.
	// The parameter must be an identifier for the resource type: ``com.vmware.content.library.item.DownloadSession``.
	// @throws NotFound  if the download session does not exist.
	// @throws com.vmware.vapi.std.errors.Unauthorized if you do not have all of the privileges described as follows:
	//
	// * Method execution requires ``System.Anonymous``.
	Delete(downloadSessionIdParam string) error

	// Terminates the download session with a client specified error message.
	//
	//  This is useful in transmitting client side failures (for example, not being able to download a file) to the server side.
	//
	// @param downloadSessionIdParam  Identifier of the download session to fail.
	// The parameter must be an identifier for the resource type: ``com.vmware.content.library.item.DownloadSession``.
	// @param clientErrorMessageParam  Client side error message. This can be useful in providing some extra details about the client side failure. Note that the message won't be translated to the user's locale.
	// @throws NotFound  if the download session does not exist.
	// @throws NotAllowedInCurrentState  if the download session is not in the DownloadSessionModelState#State_ACTIVE state.
	// @throws com.vmware.vapi.std.errors.Unauthorized if you do not have all of the privileges described as follows:
	//
	// * Method execution requires ``System.Anonymous``.
	Fail(downloadSessionIdParam string, clientErrorMessageParam string) error
}
