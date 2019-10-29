/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Code generated. DO NOT EDIT.

/*
 * Interface file for service: KeyOperation
 * Used by client-side stubs.
 */

package kms


// The ``KeyOperation`` interface provides methods to encrypt and decrypt data using a key that is provisioned for a Key Provider.
type KeyOperationClient interface {

    // Generate a new data encryption key.
    //
    // @param providerParam Identifier of the Key Provider.
    // The parameter must be an identifier for the resource type: ``com.vmware.esx.kms.providers``.
    // @param numOfBytesParam Key length.
    // @return A new data encryption key.
    // @throws InvalidArgument if the arguments are invalid.
    // @throws NotFound if the provider is not found.
    // @throws Unauthenticated if the caller is not authenticated.
    // @throws Unauthorized if the caller is not authorized.
    // @throws Error if any other error occurs.
	GenerateKey(providerParam string, numOfBytesParam int64) (KeyOperationGeneratedKey, error)

    // Encrypt plaintext.
    //
    // @param providerParam Identifier of the Key Provider.
    // The parameter must be an identifier for the resource type: ``com.vmware.esx.kms.providers``.
    // @param plaintextParam Plaintext to encrypt.
    // @return Encrypted plaintext.
    // @throws InvalidArgument if the arguments are invalid.
    // @throws NotFound if the provider is not found.
    // @throws Unauthenticated if the caller is not authenticated.
    // @throws Unauthorized if the caller is not authorized.
    // @throws Error if any other error occurs.
	Encrypt(providerParam string, plaintextParam string) (KeyOperationEncryptResult, error)

    // Decrypt ciphertext.
    //
    // @param providerParam Identifier of the Key Provider.
    // The parameter must be an identifier for the resource type: ``com.vmware.esx.kms.providers``.
    // @param ciphertextParam Ciphertext to decrypt.
    // @return Decrypted ciphertext.
    // @throws InvalidArgument if the arguments are invalid.
    // @throws NotFound if the provider is not found.
    // @throws Unauthenticated if the caller is not authenticated.
    // @throws Unauthorized if the caller is not authorized.
    // @throws Error if any other error occurs.
	Decrypt(providerParam string, ciphertextParam string) (KeyOperationDecryptResult, error)
}
