/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Interface file for service: AssociatedProducts
 * Used by client-side stubs.
 */
package discovery


// The ``AssociatedProducts`` interface provides options to list, add, modify, and delete VMware products associated with vCenter Server. Some products can be auto-detected by the system while others can be added manually.
type AssociatedProductsClient interface {

    // Retrieves a list of all associated VMware product deployments with vCenter Server in the environment. The list contains both product deployments discovered automatically and deployments registered manually through the API.
    // @return List of all the registered products with vCenter.
    // @throws Unauthenticated if the user can not be authenticated.
    // @throws Error If there is some unknown internal error. The accompanying error message will give more details about the failure.
	List() ([]Product, error)

    // Associates a VMware product with vCenter Server manually. The product must be taken from the product catalog API.
    //
    // @param specParam Info creation specification.
    // @return Identifier of the newly-added product.
    // The return value will be an identifier for the resource type: ``com.vmware.vcenter.lcm.product``.
    // @throws Unauthenticated if the user can not be authenticated.
    // @throws AlreadyExists if this version is already added
    // @throws InvalidArgument if the spec argument is not allowed
    // @throws Error If there is some unknown internal error. The accompanying error message will give more details about the failure.
	Create(specParam AssociatedProductsCreateSpec) (string, error)

    // Returns the detailed information of a product associated with vCenter Server.
    //
    // @param productParam  An identifier of the product to be modified.
    // The parameter must be an identifier for the resource type: ``com.vmware.vcenter.lcm.product``.
    // @return Product details.
    // @throws Unauthenticated if the user can not be authenticated.
    // @throws NotFound If there is no record associated with ``product`` in the system.
    // @throws Error If there is some unknown internal error. The accompanying error message will give more details about the failure.
	Get(productParam string) (Product, error)

    // Modifies a VMware product associated with vCenter Server which was added manually. Automatically discovered VMware products cannot be modified.
    //
    // @param productParam  An id of the product to be modified.
    // The parameter must be an identifier for the resource type: ``com.vmware.vcenter.lcm.product``.
    // @param specParam 
    // @throws Unauthenticated if the user can not be authenticated.
    // @throws InvalidArgument if the spec argument is not allowed
    // @throws NotFound If there is no record associated with ``product`` in the system.
    // @throws Error If there is some unknown internal error. The accompanying error message will give more details about the failure.
	Update(productParam string, specParam AssociatedProductsUpdateSpec) error

    // Deletes or dissociates a VMware product associated with vCenter Server which was added manually. Automatically discovered VMware products cannot be deleted or dissociated.
    //
    // @param productParam  An id of the product to be removed.
    // The parameter must be an identifier for the resource type: ``com.vmware.vcenter.lcm.product``.
    // @throws Unauthenticated if the user can not be authenticated.
    // @throws NotFound If there is no record associated with ``product`` in the system database.
    // @throws Error If there is some unknown internal error. The accompanying error message will give more details about the failure.
	Delete(productParam string) error
}
