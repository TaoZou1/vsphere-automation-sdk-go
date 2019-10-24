/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Interface file for service: Providers
 * Used by client-side stubs.
 */
package identity


// The ``Providers`` interface provides methods to list, read and modify vCenter Server identity providers. **Warning:** This interface is available as Technology Preview. These are early access APIs provided to test, automate and provide feedback on the feature. Since this can change based on feedback, VMware does not guarantee backwards compatibility and recommends against using them in production environments. Some Technology Preview APIs might only be applicable to specific environments.
type ProvidersClient interface {

    // Retrieve all identity providers. **Warning:** This method is available as Technology Preview. These are early access APIs provided to test, automate and provide feedback on the feature. Since this can change based on feedback, VMware does not guarantee backwards compatibility and recommends against using them in production environments. Some Technology Preview APIs might only be applicable to specific environments.
    // @return Commonly used information about the identity providers.
    // @throws Unauthorized if authorization is not given to caller.
	List() ([]ProvidersSummary, error)

    // Retrieve detailed information of the specified identity provider. **Warning:** This method is available as Technology Preview. These are early access APIs provided to test, automate and provide feedback on the feature. Since this can change based on feedback, VMware does not guarantee backwards compatibility and recommends against using them in production environments. Some Technology Preview APIs might only be applicable to specific environments.
    //
    // @param providerParam the identifier of the provider
    // The parameter must be an identifier for the resource type: ``com.vmware.vcenter.identity.Providers``.
    // @return Detailed information of the specified identity provider.
    // @throws Unauthorized if authorization is not given to caller.
    // @throws NotFound if no provider found with the given provider identifier.
	Get(providerParam string) (ProvidersInfo, error)

    // Create a vCenter Server identity provider. **Warning:** This method is available as Technology Preview. These are early access APIs provided to test, automate and provide feedback on the feature. Since this can change based on feedback, VMware does not guarantee backwards compatibility and recommends against using them in production environments. Some Technology Preview APIs might only be applicable to specific environments.
    //
    // @param specParam the CreateSpec contains the information used to create the provider
    // @return The identifier of the created identity provider.
    // The return value will be an identifier for the resource type: ``com.vmware.vcenter.identity.Providers``.
    // @throws Unauthorized if authorization is not given to caller.
    // @throws InvalidArgument if invalid arguments are provided in createSpec.
	Create(specParam ProvidersCreateSpec) (string, error)

    // Update a vCenter Server identity provider. **Warning:** This method is available as Technology Preview. These are early access APIs provided to test, automate and provide feedback on the feature. Since this can change based on feedback, VMware does not guarantee backwards compatibility and recommends against using them in production environments. Some Technology Preview APIs might only be applicable to specific environments.
    //
    // @param providerParam the identifier of the provider to update
    // The parameter must be an identifier for the resource type: ``com.vmware.vcenter.identity.Providers``.
    // @param specParam the UpdateSpec contains the information used to update the provider
    // @throws Unauthorized if authorization is not given to caller.
    // @throws InvalidArgument if invalid arguments are provided in updateSpec.
    // @throws NotFound if no provider found with the given provider identifier.
	Update(providerParam string, specParam ProvidersUpdateSpec) error

    // Delete a vCenter Server identity provider. **Warning:** This method is available as Technology Preview. These are early access APIs provided to test, automate and provide feedback on the feature. Since this can change based on feedback, VMware does not guarantee backwards compatibility and recommends against using them in production environments. Some Technology Preview APIs might only be applicable to specific environments.
    //
    // @param providerParam the identifier of the provider to delete
    // The parameter must be an identifier for the resource type: ``com.vmware.vcenter.identity.Providers``.
    // @throws Unauthorized if authorization is not given to caller.
    // @throws NotFound if no provider found with the given provider identifier.
	Delete(providerParam string) error
}
