/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Data type definitions file for service: LibraryItems.
 * Includes binding types of a structures and enumerations defined in the service.
 * Shared by client-side stubs and server-side skeletons to ensure type
 * compatibility.
 */
package vm_template

import (
	"reflect"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/bindings/com/vmware/vcenter/vm"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/bindings"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/data"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/protocol"
)



// The ``CreateSpec`` class defines the information required to create a library item containing a virtual machine template.
type LibraryItemsCreateSpec struct {
    // Identifier of the source virtual machine to create the library item from.
	SourceVm string
    // Name of the library item.
	Name string
    // Description of the library item.
	Description *string
    // Identifier of the library in which the new library item should be created.
	Library string
    // Storage location for the virtual machine template's configuration and log files.
	VmHomeStorage *LibraryItemsCreateSpecVmHomeStorage
    // Storage specification for the virtual machine template's disks.
	DiskStorage *LibraryItemsCreateSpecDiskStorage
    // Storage specification for individual disks in the virtual machine template. This is specified as a mapping between disk identifiers in the source virtual machine and their respective storage specifications.
	DiskStorageOverrides map[string]LibraryItemsCreateSpecDiskStorage
    // Information used to place the virtual machine template.
	Placement *LibraryItemsCreatePlacementSpec
}

// The ``CreatePlacementSpec`` class contains information used to place a virtual machine template onto resources within the vCenter inventory.
type LibraryItemsCreatePlacementSpec struct {
    // Virtual machine folder into which the virtual machine template should be placed.
	Folder *string
    // Resource pool into which the virtual machine template should be placed.
	ResourcePool *string
    // Host onto which the virtual machine template should be placed. If ``host`` and ``resourcePool`` are both specified, ``resourcePool`` must belong to ``host``. If ``host`` and ``cluster`` are both specified, ``host`` must be a member of ``cluster``.
	Host *string
    // Cluster onto which the virtual machine template should be placed. If ``cluster`` and ``resourcePool`` are both specified, ``resourcePool`` must belong to ``cluster``. If ``cluster`` and ``host`` are both specified, ``host`` must be a member of ``cluster``.
	Cluster *string
}

// The ``CreateSpecVmHomeStoragePolicy`` class defines the storage policy specification for a virtual machine template's configuration and log files.
type LibraryItemsCreateSpecVmHomeStoragePolicy struct {
    // Policy type to be used when creating the virtual machine template's configuration and log files.
	Type_ LibraryItemsCreateSpecVmHomeStoragePolicyType
    // Identifier for the storage policy to use.
	Policy *string
}

// Policy type for the virtual machine template's configuration and log files.
//
// <p> See {@link com.vmware.vapi.bindings.ApiEnumeration enumerated types description}.
type LibraryItemsCreateSpecVmHomeStoragePolicyType string

const (
    // Use the specified policy.
	LibraryItemsCreateSpecVmHomeStoragePolicyType_USE_SPECIFIED_POLICY LibraryItemsCreateSpecVmHomeStoragePolicyType = "USE_SPECIFIED_POLICY"
)

func (t LibraryItemsCreateSpecVmHomeStoragePolicyType) LibraryItemsCreateSpecVmHomeStoragePolicyType() bool {
	switch t {
	case LibraryItemsCreateSpecVmHomeStoragePolicyType_USE_SPECIFIED_POLICY:
		return true
	default:
		return false
	}
}


// The ``CreateSpecVmHomeStorage`` class defines the storage specification for a virtual machine template's configuration and log files.
type LibraryItemsCreateSpecVmHomeStorage struct {
    // Identifier of the datastore for the virtual machine template's configuration and log files.
	Datastore *string
    // Storage policy for the virtual machine template's configuration and log files.
	StoragePolicy *LibraryItemsCreateSpecVmHomeStoragePolicy
}

// The ``CreateSpecDiskStoragePolicy`` class defines the storage policy specification for a virtual machine template's disks.
type LibraryItemsCreateSpecDiskStoragePolicy struct {
    // Policy type to be used when creating a virtual machine template's disk.
	Type_ LibraryItemsCreateSpecDiskStoragePolicyType
    // Identifier for the storage policy to use.
	Policy *string
}

// Policy type for a virtual machine template's disk.
//
// <p> See {@link com.vmware.vapi.bindings.ApiEnumeration enumerated types description}.
type LibraryItemsCreateSpecDiskStoragePolicyType string

const (
    // Use the specified policy.
	LibraryItemsCreateSpecDiskStoragePolicyType_USE_SPECIFIED_POLICY LibraryItemsCreateSpecDiskStoragePolicyType = "USE_SPECIFIED_POLICY"
)

func (t LibraryItemsCreateSpecDiskStoragePolicyType) LibraryItemsCreateSpecDiskStoragePolicyType() bool {
	switch t {
	case LibraryItemsCreateSpecDiskStoragePolicyType_USE_SPECIFIED_POLICY:
		return true
	default:
		return false
	}
}


// The ``CreateSpecDiskStorage`` class defines the storage specification for a virtual machine template's disks.
type LibraryItemsCreateSpecDiskStorage struct {
    // Identifier for the datastore associated with a virtual machine template's disk.
	Datastore *string
    // Storage policy for a virtual machine template's disk.
	StoragePolicy *LibraryItemsCreateSpecDiskStoragePolicy
}

// The ``DeploySpec`` class defines the deployment parameters that can be specified for the ``deploy()`` method.
type LibraryItemsDeploySpec struct {
    // Name of the deployed virtual machine.
	Name string
    // Description of the deployed virtual machine.
	Description *string
    // Storage location for the deployed virtual machine's configuration and log files.
	VmHomeStorage *LibraryItemsDeploySpecVmHomeStorage
    // Storage specification for the deployed virtual machine's disks.
	DiskStorage *LibraryItemsDeploySpecDiskStorage
    // Storage specification for individual disks in the deployed virtual machine. This is specified as a mapping between disk identifiers in the source virtual machine template contained in the library item and their storage specifications.
	DiskStorageOverrides map[string]LibraryItemsDeploySpecDiskStorage
    // Information used to place the deployed virtual machine.
	Placement *LibraryItemsDeployPlacementSpec
    // Specifies whether the deployed virtual machine should be powered on after deployment.
	PoweredOn *bool
    // Guest customization spec to apply to the deployed virtual machine.
	GuestCustomization *LibraryItemsGuestCustomizationSpec
    // Hardware customization spec which specifies updates to the deployed virtual machine.
	HardwareCustomization *LibraryItemsHardwareCustomizationSpec
}

// The ``HardwareCustomizationSpec`` class defines the hardware customization options that are applied to the deployed virtual machine.
type LibraryItemsHardwareCustomizationSpec struct {
    // Map of Ethernet network adapters to update.
	Nics map[string]LibraryItemsEthernetUpdateSpec
    // Idenfiers of disks to remove from the deployed virtual machine.
	DisksToRemove map[string]bool
    // Disk update specification for individual disks in the deployed virtual machine.
	DisksToUpdate map[string]LibraryItemsDiskUpdateSpec
    // CPU update specification for the deployed virtual machine.
	CpuUpdate *LibraryItemsCpuUpdateSpec
    // Memory update specification for the deployed virtual machine.
	MemoryUpdate *LibraryItemsMemoryUpdateSpec
}

// The ``DiskUpdateSpec`` class describes updates to the configuration of a virtual disk in the deployed virtual machine.
type LibraryItemsDiskUpdateSpec struct {
    // Updated capacity of the virtual disk backing in bytes. This value has to be larger than the original capacity of the disk.
	Capacity int64
}

// The ``CpuUpdateSpec`` class describes updates to the CPU configuration of the deployed virtual machine.
type LibraryItemsCpuUpdateSpec struct {
    // Number of virtual processors in the deployed virtual machine.
	NumCpus *int64
    // Number of cores among which to distribute CPUs in the deployed virtual machine.
	NumCoresPerSocket *int64
}

// The ``MemoryUpdateSpec`` class describes updates to the memory configuration of the deployed virtual machine.
type LibraryItemsMemoryUpdateSpec struct {
    // Size of a virtual machine's memory in MB.
	Memory *int64
}

// The ``EthernetUpdateSpec`` class describes the network that the ethernet adapter of the deployed virtual machine should be connected to.
type LibraryItemsEthernetUpdateSpec struct {
    // Identifier of the network backing the virtual Ethernet adapter.
	Network *string
}

// The ``DeployPlacementSpec`` class contains information used to place a virtual machine onto resources within the vCenter inventory.
type LibraryItemsDeployPlacementSpec struct {
    // Virtual machine folder into which the deployed virtual machine should be placed.
	Folder *string
    // Resource pool into which the deployed virtual machine should be placed.
	ResourcePool *string
    // Host onto which the virtual machine should be placed. If ``host`` and ``resourcePool`` are both specified, ``resourcePool`` must belong to ``host``. If ``host`` and ``cluster`` are both specified, ``host`` must be a member of ``cluster``.
	Host *string
    // Cluster onto which the deployed virtual machine should be placed. If ``cluster`` and ``resourcePool`` are both specified, ``resourcePool`` must belong to ``cluster``. If ``cluster`` and ``host`` are both specified, ``host`` must be a member of ``cluster``.
	Cluster *string
}

// The ``DeploySpecVmHomeStoragePolicy`` class defines the storage policy specification for the deployed virtual machine's configuration and log files.
type LibraryItemsDeploySpecVmHomeStoragePolicy struct {
    // Policy type to be used when creating the deployed virtual machine's configuration and log files.
	Type_ LibraryItemsDeploySpecVmHomeStoragePolicyType
    // Identifier for the storage policy to use.
	Policy *string
}

// Policy type for the deployed virtual machine's configuration and log files.
//
// <p> See {@link com.vmware.vapi.bindings.ApiEnumeration enumerated types description}.
type LibraryItemsDeploySpecVmHomeStoragePolicyType string

const (
    // Use the specified policy.
	LibraryItemsDeploySpecVmHomeStoragePolicyType_USE_SPECIFIED_POLICY LibraryItemsDeploySpecVmHomeStoragePolicyType = "USE_SPECIFIED_POLICY"
    // Use the storage policy that is associated with the source virtual machine template's configuration and log files.
	LibraryItemsDeploySpecVmHomeStoragePolicyType_USE_SOURCE_POLICY LibraryItemsDeploySpecVmHomeStoragePolicyType = "USE_SOURCE_POLICY"
)

func (t LibraryItemsDeploySpecVmHomeStoragePolicyType) LibraryItemsDeploySpecVmHomeStoragePolicyType() bool {
	switch t {
	case LibraryItemsDeploySpecVmHomeStoragePolicyType_USE_SPECIFIED_POLICY:
		return true
	case LibraryItemsDeploySpecVmHomeStoragePolicyType_USE_SOURCE_POLICY:
		return true
	default:
		return false
	}
}


// The ``DeploySpecVmHomeStorage`` class defines the storage specification for a deployed virtual machine's configuration and log files.
type LibraryItemsDeploySpecVmHomeStorage struct {
    // Identifier of the datastore for the deployed virtual machine's configuration and log files.
	Datastore *string
    // Storage policy for the deployed virtual machine's configuration and log files.
	StoragePolicy *LibraryItemsDeploySpecVmHomeStoragePolicy
}

// The ``DeploySpecDiskStoragePolicy`` class describes the storage policy specification for the deployed virtual machine's disks.
type LibraryItemsDeploySpecDiskStoragePolicy struct {
    // Policy type to be used when creating the deployed virtual machine's disk.
	Type_ LibraryItemsDeploySpecDiskStoragePolicyType
    // Identifier of the storage policy to use.
	Policy *string
}

// Policy type for the deployed virtual machine's disk.
//
// <p> See {@link com.vmware.vapi.bindings.ApiEnumeration enumerated types description}.
type LibraryItemsDeploySpecDiskStoragePolicyType string

const (
    // Use the specified policy.
	LibraryItemsDeploySpecDiskStoragePolicyType_USE_SPECIFIED_POLICY LibraryItemsDeploySpecDiskStoragePolicyType = "USE_SPECIFIED_POLICY"
    // Use the storage policy that is associated with the corresponding disk in the source virtual machine template.
	LibraryItemsDeploySpecDiskStoragePolicyType_USE_SOURCE_POLICY LibraryItemsDeploySpecDiskStoragePolicyType = "USE_SOURCE_POLICY"
)

func (t LibraryItemsDeploySpecDiskStoragePolicyType) LibraryItemsDeploySpecDiskStoragePolicyType() bool {
	switch t {
	case LibraryItemsDeploySpecDiskStoragePolicyType_USE_SPECIFIED_POLICY:
		return true
	case LibraryItemsDeploySpecDiskStoragePolicyType_USE_SOURCE_POLICY:
		return true
	default:
		return false
	}
}


// The ``DeploySpecDiskStorage`` class contains the storage specification for disks in the virtual machine.
type LibraryItemsDeploySpecDiskStorage struct {
    // Identifier for the datastore associated the deployed virtual machine's disk.
	Datastore *string
    // Storage policy for the deployed virtual machine's disk.
	StoragePolicy *LibraryItemsDeploySpecDiskStoragePolicy
}

// The ``GuestCustomizationSpec`` class contains information required to customize the deployed virtual machine.
type LibraryItemsGuestCustomizationSpec struct {
    // Name of the customization specification.
	Name *string
}

// The ``Info`` class contains information about a virtual machine template item in content library.
type LibraryItemsInfo struct {
    // Configured guest operating system of the virtual machine template.
	GuestOS vm.GuestOS
    // CPU configuration of the virtual machine template.
	Cpu LibraryItemsCpuInfo
    // Memory configuration of the virtual machine template.
	Memory LibraryItemsMemoryInfo
    // Storage information about the virtual machine template's configuration and log files.
	VmHomeStorage LibraryItemsVmHomeStorageInfo
    // Storage information about the virtual machine template's virtual disks.
	Disks map[string]LibraryItemsDiskInfo
    // Information about the virtual machine template's virtual ethernet adapters.
	Nics map[string]LibraryItemsEthernetInfo
    // Identifier of the latest virtual machine template contained in the library item. This property is the managed object identifier used to identify the virtual machine template in the vSphere Management (SOAP) API.
	VmTemplate string
}

// The ``CpuInfo`` class contains CPU related information about the virtual machine template.
type LibraryItemsCpuInfo struct {
    // Number of CPU cores.
	Count int64
    // Number of CPU cores per socket.
	CoresPerSocket int64
}

// The ``MemoryInfo`` class contains memory related information about the virtual machine template.
type LibraryItemsMemoryInfo struct {
    // Memory size in mebibytes.
	SizeMiB int64
}

// The ``VmHomeStorageInfo`` class contains storage information about the virtual machine template's configuration and log files.
type LibraryItemsVmHomeStorageInfo struct {
    // Identifier of the datastore where the virtual machine template's configuration and log files are stored.
	Datastore string
    // Identifier of the storage policy associated with the virtual machine template's configuration and log files.
	StoragePolicy *string
}

// The ``DiskInfo`` class contains information about the virtual machine template's virtual disk.
type LibraryItemsDiskInfo struct {
    // Capacity of the virtual disk in bytes.
	Capacity *int64
    // Disk storage related information.
	DiskStorage LibraryItemsDiskStorageInfo
}

// The ``DiskStorageInfo`` class contains storage related information about a virtual machine template's virtual disk.
type LibraryItemsDiskStorageInfo struct {
    // Identifier of the datastore where the disk is stored.
	Datastore string
    // Identifier of the storage policy associated with the virtual disk.
	StoragePolicy *string
}

// The ``EthernetInfo`` class contains information about a virtual machine template's virtual Ethernet adapter.
type LibraryItemsEthernetInfo struct {
    // Network backing type for the virtual Ethernet adapter.
	BackingType LibraryItemsEthernetInfoNetworkBackingType
    // MAC address type of the ethernet adapter.
	MacType LibraryItemsEthernetInfoMacAddressType
    // Identifier of the network backing the virtual Ethernet adapter.
	Network *string
}

// The ``NetworkBackingType`` enumeration class defines valid network backing types for a virtual Ethernet adapter.
//
// <p> See {@link com.vmware.vapi.bindings.ApiEnumeration enumerated types description}.
type LibraryItemsEthernetInfoNetworkBackingType string

const (
    // vSphere standard portgroup network backing.
	LibraryItemsEthernetInfoNetworkBackingType_STANDARD_PORTGROUP LibraryItemsEthernetInfoNetworkBackingType = "STANDARD_PORTGROUP"
    // Legacy host device network backing. Imported VMs may have virtual Ethernet adapters with this type of backing, but this type of backing cannot be used to create or to update a virtual Ethernet adapter.
	LibraryItemsEthernetInfoNetworkBackingType_HOST_DEVICE LibraryItemsEthernetInfoNetworkBackingType = "HOST_DEVICE"
    // Distributed virtual switch backing.
	LibraryItemsEthernetInfoNetworkBackingType_DISTRIBUTED_PORTGROUP LibraryItemsEthernetInfoNetworkBackingType = "DISTRIBUTED_PORTGROUP"
    // Opaque network backing.
	LibraryItemsEthernetInfoNetworkBackingType_OPAQUE_NETWORK LibraryItemsEthernetInfoNetworkBackingType = "OPAQUE_NETWORK"
)

func (n LibraryItemsEthernetInfoNetworkBackingType) LibraryItemsEthernetInfoNetworkBackingType() bool {
	switch n {
	case LibraryItemsEthernetInfoNetworkBackingType_STANDARD_PORTGROUP:
		return true
	case LibraryItemsEthernetInfoNetworkBackingType_HOST_DEVICE:
		return true
	case LibraryItemsEthernetInfoNetworkBackingType_DISTRIBUTED_PORTGROUP:
		return true
	case LibraryItemsEthernetInfoNetworkBackingType_OPAQUE_NETWORK:
		return true
	default:
		return false
	}
}


// The ``MacAddressType`` enumeration class defines the valid MAC address origins for a virtual Ethernet adapter.
//
// <p> See {@link com.vmware.vapi.bindings.ApiEnumeration enumerated types description}.
type LibraryItemsEthernetInfoMacAddressType string

const (
    // MAC address is assigned statically.
	LibraryItemsEthernetInfoMacAddressType_MANUAL LibraryItemsEthernetInfoMacAddressType = "MANUAL"
    // MAC address is generated automatically.
	LibraryItemsEthernetInfoMacAddressType_GENERATED LibraryItemsEthernetInfoMacAddressType = "GENERATED"
    // MAC address is assigned by vCenter Server.
	LibraryItemsEthernetInfoMacAddressType_ASSIGNED LibraryItemsEthernetInfoMacAddressType = "ASSIGNED"
)

func (m LibraryItemsEthernetInfoMacAddressType) LibraryItemsEthernetInfoMacAddressType() bool {
	switch m {
	case LibraryItemsEthernetInfoMacAddressType_MANUAL:
		return true
	case LibraryItemsEthernetInfoMacAddressType_GENERATED:
		return true
	case LibraryItemsEthernetInfoMacAddressType_ASSIGNED:
		return true
	default:
		return false
	}
}




func libraryItemsCreateInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["spec"] = bindings.NewReferenceType(LibraryItemsCreateSpecBindingType)
	fieldNameMap["spec"] = "Spec"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func libraryItemsCreateOutputType() bindings.BindingType {
	return bindings.NewIdType([]string{"com.vmware.content.library.Item"}, "")
}

func libraryItemsCreateRestMetadata() protocol.OperationRestMetadata {
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	resultHeaders := map[string]string{}
	errorHeaders := map[string]string{}
	errorHeaders["Unauthenticated.challenge"] = "WWW-Authenticate"
	return protocol.NewOperationRestMetadata(
		paramsTypeMap,
		pathParams,
		queryParams,
		headerParams,
		"",
		"null",
		"",
		resultHeaders,
		0,
		errorHeaders,
		map[string]int{"AlreadyExists": 400,"Error": 500,"InvalidArgument": 400,"NotAllowedInCurrentState": 400,"NotFound": 404,"ResourceInaccessible": 500,"ResourceInUse": 400,"ServiceUnavailable": 503,"UnableToAllocateResource": 500,"Unauthenticated": 401,"Unauthorized": 403})
}

func libraryItemsDeployInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["template_library_item"] = bindings.NewIdType([]string{"com.vmware.content.library.Item"}, "")
	fields["spec"] = bindings.NewReferenceType(LibraryItemsDeploySpecBindingType)
	fieldNameMap["template_library_item"] = "TemplateLibraryItem"
	fieldNameMap["spec"] = "Spec"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func libraryItemsDeployOutputType() bindings.BindingType {
	return bindings.NewIdType([]string{"VirtualMachine"}, "")
}

func libraryItemsDeployRestMetadata() protocol.OperationRestMetadata {
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	resultHeaders := map[string]string{}
	errorHeaders := map[string]string{}
	errorHeaders["Unauthenticated.challenge"] = "WWW-Authenticate"
	return protocol.NewOperationRestMetadata(
		paramsTypeMap,
		pathParams,
		queryParams,
		headerParams,
		"",
		"null",
		"",
		resultHeaders,
		0,
		errorHeaders,
		map[string]int{"AlreadyExists": 400,"Error": 500,"InvalidArgument": 400,"NotAllowedInCurrentState": 400,"NotFound": 404,"ResourceInaccessible": 500,"ServiceUnavailable": 503,"UnableToAllocateResource": 500,"Unauthenticated": 401,"Unauthorized": 403})
}

func libraryItemsGetInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["template_library_item"] = bindings.NewIdType([]string{"com.vmware.content.library.Item"}, "")
	fieldNameMap["template_library_item"] = "TemplateLibraryItem"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func libraryItemsGetOutputType() bindings.BindingType {
	return bindings.NewOptionalType(bindings.NewReferenceType(LibraryItemsInfoBindingType))
}

func libraryItemsGetRestMetadata() protocol.OperationRestMetadata {
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	resultHeaders := map[string]string{}
	errorHeaders := map[string]string{}
	errorHeaders["Unauthenticated.challenge"] = "WWW-Authenticate"
	return protocol.NewOperationRestMetadata(
		paramsTypeMap,
		pathParams,
		queryParams,
		headerParams,
		"",
		"null",
		"",
		resultHeaders,
		0,
		errorHeaders,
		map[string]int{"Error": 500,"NotFound": 404,"ResourceInaccessible": 500,"ServiceUnavailable": 503,"Unauthenticated": 401,"Unauthorized": 403})
}


func LibraryItemsCreateSpecBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["source_vm"] = bindings.NewIdType([]string{"VirtualMachine"}, "")
	fieldNameMap["source_vm"] = "SourceVm"
	fields["name"] = bindings.NewStringType()
	fieldNameMap["name"] = "Name"
	fields["description"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["description"] = "Description"
	fields["library"] = bindings.NewIdType([]string{"com.vmware.content.Library"}, "")
	fieldNameMap["library"] = "Library"
	fields["vm_home_storage"] = bindings.NewOptionalType(bindings.NewReferenceType(LibraryItemsCreateSpecVmHomeStorageBindingType))
	fieldNameMap["vm_home_storage"] = "VmHomeStorage"
	fields["disk_storage"] = bindings.NewOptionalType(bindings.NewReferenceType(LibraryItemsCreateSpecDiskStorageBindingType))
	fieldNameMap["disk_storage"] = "DiskStorage"
	fields["disk_storage_overrides"] = bindings.NewOptionalType(bindings.NewMapType(bindings.NewIdType([]string{"com.vmware.vcenter.vm.hardware.Disk"}, ""), bindings.NewReferenceType(LibraryItemsCreateSpecDiskStorageBindingType),reflect.TypeOf(map[string]LibraryItemsCreateSpecDiskStorage{})))
	fieldNameMap["disk_storage_overrides"] = "DiskStorageOverrides"
	fields["placement"] = bindings.NewOptionalType(bindings.NewReferenceType(LibraryItemsCreatePlacementSpecBindingType))
	fieldNameMap["placement"] = "Placement"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vcenter.vm_template.library_items.create_spec", fields, reflect.TypeOf(LibraryItemsCreateSpec{}), fieldNameMap, validators)
}

func LibraryItemsCreatePlacementSpecBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["folder"] = bindings.NewOptionalType(bindings.NewIdType([]string{"Folder"}, ""))
	fieldNameMap["folder"] = "Folder"
	fields["resource_pool"] = bindings.NewOptionalType(bindings.NewIdType([]string{"ResourcePool"}, ""))
	fieldNameMap["resource_pool"] = "ResourcePool"
	fields["host"] = bindings.NewOptionalType(bindings.NewIdType([]string{"HostSystem"}, ""))
	fieldNameMap["host"] = "Host"
	fields["cluster"] = bindings.NewOptionalType(bindings.NewIdType([]string{"ClusterComputeResource"}, ""))
	fieldNameMap["cluster"] = "Cluster"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vcenter.vm_template.library_items.create_placement_spec", fields, reflect.TypeOf(LibraryItemsCreatePlacementSpec{}), fieldNameMap, validators)
}

func LibraryItemsCreateSpecVmHomeStoragePolicyBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["type"] = bindings.NewEnumType("com.vmware.vcenter.vm_template.library_items.create_spec_vm_home_storage_policy.type", reflect.TypeOf(LibraryItemsCreateSpecVmHomeStoragePolicyType(LibraryItemsCreateSpecVmHomeStoragePolicyType_USE_SPECIFIED_POLICY)))
	fieldNameMap["type"] = "Type_"
	fields["policy"] = bindings.NewOptionalType(bindings.NewIdType([]string{"com.vmware.spbm.StorageProfile"}, ""))
	fieldNameMap["policy"] = "Policy"
	var validators = []bindings.Validator{}
	uv1 := bindings.NewUnionValidator("type",
		map[string][]bindings.FieldData{
			"USE_SPECIFIED_POLICY": []bindings.FieldData{
				bindings.NewFieldData("policy", true),
			},
		},
	)
	validators = append(validators, uv1)
	return bindings.NewStructType("com.vmware.vcenter.vm_template.library_items.create_spec_vm_home_storage_policy", fields, reflect.TypeOf(LibraryItemsCreateSpecVmHomeStoragePolicy{}), fieldNameMap, validators)
}

func LibraryItemsCreateSpecVmHomeStorageBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["datastore"] = bindings.NewOptionalType(bindings.NewIdType([]string{"Datastore"}, ""))
	fieldNameMap["datastore"] = "Datastore"
	fields["storage_policy"] = bindings.NewOptionalType(bindings.NewReferenceType(LibraryItemsCreateSpecVmHomeStoragePolicyBindingType))
	fieldNameMap["storage_policy"] = "StoragePolicy"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vcenter.vm_template.library_items.create_spec_vm_home_storage", fields, reflect.TypeOf(LibraryItemsCreateSpecVmHomeStorage{}), fieldNameMap, validators)
}

func LibraryItemsCreateSpecDiskStoragePolicyBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["type"] = bindings.NewEnumType("com.vmware.vcenter.vm_template.library_items.create_spec_disk_storage_policy.type", reflect.TypeOf(LibraryItemsCreateSpecDiskStoragePolicyType(LibraryItemsCreateSpecDiskStoragePolicyType_USE_SPECIFIED_POLICY)))
	fieldNameMap["type"] = "Type_"
	fields["policy"] = bindings.NewOptionalType(bindings.NewIdType([]string{"com.vmware.spbm.StorageProfile"}, ""))
	fieldNameMap["policy"] = "Policy"
	var validators = []bindings.Validator{}
	uv1 := bindings.NewUnionValidator("type",
		map[string][]bindings.FieldData{
			"USE_SPECIFIED_POLICY": []bindings.FieldData{
				bindings.NewFieldData("policy", true),
			},
		},
	)
	validators = append(validators, uv1)
	return bindings.NewStructType("com.vmware.vcenter.vm_template.library_items.create_spec_disk_storage_policy", fields, reflect.TypeOf(LibraryItemsCreateSpecDiskStoragePolicy{}), fieldNameMap, validators)
}

func LibraryItemsCreateSpecDiskStorageBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["datastore"] = bindings.NewOptionalType(bindings.NewIdType([]string{"Datastore"}, ""))
	fieldNameMap["datastore"] = "Datastore"
	fields["storage_policy"] = bindings.NewOptionalType(bindings.NewReferenceType(LibraryItemsCreateSpecDiskStoragePolicyBindingType))
	fieldNameMap["storage_policy"] = "StoragePolicy"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vcenter.vm_template.library_items.create_spec_disk_storage", fields, reflect.TypeOf(LibraryItemsCreateSpecDiskStorage{}), fieldNameMap, validators)
}

func LibraryItemsDeploySpecBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["name"] = bindings.NewStringType()
	fieldNameMap["name"] = "Name"
	fields["description"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["description"] = "Description"
	fields["vm_home_storage"] = bindings.NewOptionalType(bindings.NewReferenceType(LibraryItemsDeploySpecVmHomeStorageBindingType))
	fieldNameMap["vm_home_storage"] = "VmHomeStorage"
	fields["disk_storage"] = bindings.NewOptionalType(bindings.NewReferenceType(LibraryItemsDeploySpecDiskStorageBindingType))
	fieldNameMap["disk_storage"] = "DiskStorage"
	fields["disk_storage_overrides"] = bindings.NewOptionalType(bindings.NewMapType(bindings.NewIdType([]string{"com.vmware.vcenter.vm.hardware.Disk"}, ""), bindings.NewReferenceType(LibraryItemsDeploySpecDiskStorageBindingType),reflect.TypeOf(map[string]LibraryItemsDeploySpecDiskStorage{})))
	fieldNameMap["disk_storage_overrides"] = "DiskStorageOverrides"
	fields["placement"] = bindings.NewOptionalType(bindings.NewReferenceType(LibraryItemsDeployPlacementSpecBindingType))
	fieldNameMap["placement"] = "Placement"
	fields["powered_on"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fieldNameMap["powered_on"] = "PoweredOn"
	fields["guest_customization"] = bindings.NewOptionalType(bindings.NewReferenceType(LibraryItemsGuestCustomizationSpecBindingType))
	fieldNameMap["guest_customization"] = "GuestCustomization"
	fields["hardware_customization"] = bindings.NewOptionalType(bindings.NewReferenceType(LibraryItemsHardwareCustomizationSpecBindingType))
	fieldNameMap["hardware_customization"] = "HardwareCustomization"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vcenter.vm_template.library_items.deploy_spec", fields, reflect.TypeOf(LibraryItemsDeploySpec{}), fieldNameMap, validators)
}

func LibraryItemsHardwareCustomizationSpecBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["nics"] = bindings.NewOptionalType(bindings.NewMapType(bindings.NewIdType([]string{"com.vmware.vcenter.vm.hardware.Ethernet"}, ""), bindings.NewReferenceType(LibraryItemsEthernetUpdateSpecBindingType),reflect.TypeOf(map[string]LibraryItemsEthernetUpdateSpec{})))
	fieldNameMap["nics"] = "Nics"
	fields["disks_to_remove"] = bindings.NewOptionalType(bindings.NewSetType(bindings.NewIdType([]string{"com.vmware.vcenter.vm.hardware.Disk"}, ""), reflect.TypeOf(map[string]bool{})))
	fieldNameMap["disks_to_remove"] = "DisksToRemove"
	fields["disks_to_update"] = bindings.NewOptionalType(bindings.NewMapType(bindings.NewIdType([]string{"com.vmware.vcenter.vm.hardware.Disk"}, ""), bindings.NewReferenceType(LibraryItemsDiskUpdateSpecBindingType),reflect.TypeOf(map[string]LibraryItemsDiskUpdateSpec{})))
	fieldNameMap["disks_to_update"] = "DisksToUpdate"
	fields["cpu_update"] = bindings.NewOptionalType(bindings.NewReferenceType(LibraryItemsCpuUpdateSpecBindingType))
	fieldNameMap["cpu_update"] = "CpuUpdate"
	fields["memory_update"] = bindings.NewOptionalType(bindings.NewReferenceType(LibraryItemsMemoryUpdateSpecBindingType))
	fieldNameMap["memory_update"] = "MemoryUpdate"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vcenter.vm_template.library_items.hardware_customization_spec", fields, reflect.TypeOf(LibraryItemsHardwareCustomizationSpec{}), fieldNameMap, validators)
}

func LibraryItemsDiskUpdateSpecBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["capacity"] = bindings.NewIntegerType()
	fieldNameMap["capacity"] = "Capacity"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vcenter.vm_template.library_items.disk_update_spec", fields, reflect.TypeOf(LibraryItemsDiskUpdateSpec{}), fieldNameMap, validators)
}

func LibraryItemsCpuUpdateSpecBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["num_cpus"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fieldNameMap["num_cpus"] = "NumCpus"
	fields["num_cores_per_socket"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fieldNameMap["num_cores_per_socket"] = "NumCoresPerSocket"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vcenter.vm_template.library_items.cpu_update_spec", fields, reflect.TypeOf(LibraryItemsCpuUpdateSpec{}), fieldNameMap, validators)
}

func LibraryItemsMemoryUpdateSpecBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["memory"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fieldNameMap["memory"] = "Memory"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vcenter.vm_template.library_items.memory_update_spec", fields, reflect.TypeOf(LibraryItemsMemoryUpdateSpec{}), fieldNameMap, validators)
}

func LibraryItemsEthernetUpdateSpecBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["network"] = bindings.NewOptionalType(bindings.NewIdType([]string{"Network"}, ""))
	fieldNameMap["network"] = "Network"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vcenter.vm_template.library_items.ethernet_update_spec", fields, reflect.TypeOf(LibraryItemsEthernetUpdateSpec{}), fieldNameMap, validators)
}

func LibraryItemsDeployPlacementSpecBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["folder"] = bindings.NewOptionalType(bindings.NewIdType([]string{"Folder"}, ""))
	fieldNameMap["folder"] = "Folder"
	fields["resource_pool"] = bindings.NewOptionalType(bindings.NewIdType([]string{"ResourcePool"}, ""))
	fieldNameMap["resource_pool"] = "ResourcePool"
	fields["host"] = bindings.NewOptionalType(bindings.NewIdType([]string{"HostSystem"}, ""))
	fieldNameMap["host"] = "Host"
	fields["cluster"] = bindings.NewOptionalType(bindings.NewIdType([]string{"ClusterComputeResource"}, ""))
	fieldNameMap["cluster"] = "Cluster"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vcenter.vm_template.library_items.deploy_placement_spec", fields, reflect.TypeOf(LibraryItemsDeployPlacementSpec{}), fieldNameMap, validators)
}

func LibraryItemsDeploySpecVmHomeStoragePolicyBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["type"] = bindings.NewEnumType("com.vmware.vcenter.vm_template.library_items.deploy_spec_vm_home_storage_policy.type", reflect.TypeOf(LibraryItemsDeploySpecVmHomeStoragePolicyType(LibraryItemsDeploySpecVmHomeStoragePolicyType_USE_SPECIFIED_POLICY)))
	fieldNameMap["type"] = "Type_"
	fields["policy"] = bindings.NewOptionalType(bindings.NewIdType([]string{"com.vmware.spbm.StorageProfile"}, ""))
	fieldNameMap["policy"] = "Policy"
	var validators = []bindings.Validator{}
	uv1 := bindings.NewUnionValidator("type",
		map[string][]bindings.FieldData{
			"USE_SPECIFIED_POLICY": []bindings.FieldData{
				bindings.NewFieldData("policy", true),
			},
			"USE_SOURCE_POLICY": []bindings.FieldData{},
		},
	)
	validators = append(validators, uv1)
	return bindings.NewStructType("com.vmware.vcenter.vm_template.library_items.deploy_spec_vm_home_storage_policy", fields, reflect.TypeOf(LibraryItemsDeploySpecVmHomeStoragePolicy{}), fieldNameMap, validators)
}

func LibraryItemsDeploySpecVmHomeStorageBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["datastore"] = bindings.NewOptionalType(bindings.NewIdType([]string{"Datastore"}, ""))
	fieldNameMap["datastore"] = "Datastore"
	fields["storage_policy"] = bindings.NewOptionalType(bindings.NewReferenceType(LibraryItemsDeploySpecVmHomeStoragePolicyBindingType))
	fieldNameMap["storage_policy"] = "StoragePolicy"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vcenter.vm_template.library_items.deploy_spec_vm_home_storage", fields, reflect.TypeOf(LibraryItemsDeploySpecVmHomeStorage{}), fieldNameMap, validators)
}

func LibraryItemsDeploySpecDiskStoragePolicyBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["type"] = bindings.NewEnumType("com.vmware.vcenter.vm_template.library_items.deploy_spec_disk_storage_policy.type", reflect.TypeOf(LibraryItemsDeploySpecDiskStoragePolicyType(LibraryItemsDeploySpecDiskStoragePolicyType_USE_SPECIFIED_POLICY)))
	fieldNameMap["type"] = "Type_"
	fields["policy"] = bindings.NewOptionalType(bindings.NewIdType([]string{"com.vmware.spbm.StorageProfile"}, ""))
	fieldNameMap["policy"] = "Policy"
	var validators = []bindings.Validator{}
	uv1 := bindings.NewUnionValidator("type",
		map[string][]bindings.FieldData{
			"USE_SPECIFIED_POLICY": []bindings.FieldData{
				bindings.NewFieldData("policy", true),
			},
			"USE_SOURCE_POLICY": []bindings.FieldData{},
		},
	)
	validators = append(validators, uv1)
	return bindings.NewStructType("com.vmware.vcenter.vm_template.library_items.deploy_spec_disk_storage_policy", fields, reflect.TypeOf(LibraryItemsDeploySpecDiskStoragePolicy{}), fieldNameMap, validators)
}

func LibraryItemsDeploySpecDiskStorageBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["datastore"] = bindings.NewOptionalType(bindings.NewIdType([]string{"Datastore"}, ""))
	fieldNameMap["datastore"] = "Datastore"
	fields["storage_policy"] = bindings.NewOptionalType(bindings.NewReferenceType(LibraryItemsDeploySpecDiskStoragePolicyBindingType))
	fieldNameMap["storage_policy"] = "StoragePolicy"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vcenter.vm_template.library_items.deploy_spec_disk_storage", fields, reflect.TypeOf(LibraryItemsDeploySpecDiskStorage{}), fieldNameMap, validators)
}

func LibraryItemsGuestCustomizationSpecBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["name"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["name"] = "Name"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vcenter.vm_template.library_items.guest_customization_spec", fields, reflect.TypeOf(LibraryItemsGuestCustomizationSpec{}), fieldNameMap, validators)
}

func LibraryItemsInfoBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["guest_OS"] = bindings.NewEnumType("com.vmware.vcenter.vm.guest_OS", reflect.TypeOf(vm.GuestOS(vm.GuestOS_DOS)))
	fieldNameMap["guest_OS"] = "GuestOS"
	fields["cpu"] = bindings.NewReferenceType(LibraryItemsCpuInfoBindingType)
	fieldNameMap["cpu"] = "Cpu"
	fields["memory"] = bindings.NewReferenceType(LibraryItemsMemoryInfoBindingType)
	fieldNameMap["memory"] = "Memory"
	fields["vm_home_storage"] = bindings.NewReferenceType(LibraryItemsVmHomeStorageInfoBindingType)
	fieldNameMap["vm_home_storage"] = "VmHomeStorage"
	fields["disks"] = bindings.NewMapType(bindings.NewIdType([]string{"com.vmware.vcenter.vm.hardware.Disk"}, ""), bindings.NewReferenceType(LibraryItemsDiskInfoBindingType),reflect.TypeOf(map[string]LibraryItemsDiskInfo{}))
	fieldNameMap["disks"] = "Disks"
	fields["nics"] = bindings.NewMapType(bindings.NewIdType([]string{"com.vmware.vcenter.vm.hardware.Ethernet"}, ""), bindings.NewReferenceType(LibraryItemsEthernetInfoBindingType),reflect.TypeOf(map[string]LibraryItemsEthernetInfo{}))
	fieldNameMap["nics"] = "Nics"
	fields["vm_template"] = bindings.NewStringType()
	fieldNameMap["vm_template"] = "VmTemplate"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vcenter.vm_template.library_items.info", fields, reflect.TypeOf(LibraryItemsInfo{}), fieldNameMap, validators)
}

func LibraryItemsCpuInfoBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["count"] = bindings.NewIntegerType()
	fieldNameMap["count"] = "Count"
	fields["cores_per_socket"] = bindings.NewIntegerType()
	fieldNameMap["cores_per_socket"] = "CoresPerSocket"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vcenter.vm_template.library_items.cpu_info", fields, reflect.TypeOf(LibraryItemsCpuInfo{}), fieldNameMap, validators)
}

func LibraryItemsMemoryInfoBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["size_MiB"] = bindings.NewIntegerType()
	fieldNameMap["size_MiB"] = "SizeMiB"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vcenter.vm_template.library_items.memory_info", fields, reflect.TypeOf(LibraryItemsMemoryInfo{}), fieldNameMap, validators)
}

func LibraryItemsVmHomeStorageInfoBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["datastore"] = bindings.NewIdType([]string{"Datastore"}, "")
	fieldNameMap["datastore"] = "Datastore"
	fields["storage_policy"] = bindings.NewOptionalType(bindings.NewIdType([]string{"com.vmware.spbm.StorageProfile"}, ""))
	fieldNameMap["storage_policy"] = "StoragePolicy"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vcenter.vm_template.library_items.vm_home_storage_info", fields, reflect.TypeOf(LibraryItemsVmHomeStorageInfo{}), fieldNameMap, validators)
}

func LibraryItemsDiskInfoBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["capacity"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fieldNameMap["capacity"] = "Capacity"
	fields["disk_storage"] = bindings.NewReferenceType(LibraryItemsDiskStorageInfoBindingType)
	fieldNameMap["disk_storage"] = "DiskStorage"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vcenter.vm_template.library_items.disk_info", fields, reflect.TypeOf(LibraryItemsDiskInfo{}), fieldNameMap, validators)
}

func LibraryItemsDiskStorageInfoBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["datastore"] = bindings.NewIdType([]string{"Datastore"}, "")
	fieldNameMap["datastore"] = "Datastore"
	fields["storage_policy"] = bindings.NewOptionalType(bindings.NewIdType([]string{"com.vmware.spbm.StorageProfile"}, ""))
	fieldNameMap["storage_policy"] = "StoragePolicy"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vcenter.vm_template.library_items.disk_storage_info", fields, reflect.TypeOf(LibraryItemsDiskStorageInfo{}), fieldNameMap, validators)
}

func LibraryItemsEthernetInfoBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["backing_type"] = bindings.NewEnumType("com.vmware.vcenter.vm_template.library_items.ethernet_info.network_backing_type", reflect.TypeOf(LibraryItemsEthernetInfoNetworkBackingType(LibraryItemsEthernetInfoNetworkBackingType_STANDARD_PORTGROUP)))
	fieldNameMap["backing_type"] = "BackingType"
	fields["mac_type"] = bindings.NewEnumType("com.vmware.vcenter.vm_template.library_items.ethernet_info.mac_address_type", reflect.TypeOf(LibraryItemsEthernetInfoMacAddressType(LibraryItemsEthernetInfoMacAddressType_MANUAL)))
	fieldNameMap["mac_type"] = "MacType"
	fields["network"] = bindings.NewOptionalType(bindings.NewIdType([]string{"Network"}, ""))
	fieldNameMap["network"] = "Network"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vcenter.vm_template.library_items.ethernet_info", fields, reflect.TypeOf(LibraryItemsEthernetInfo{}), fieldNameMap, validators)
}

