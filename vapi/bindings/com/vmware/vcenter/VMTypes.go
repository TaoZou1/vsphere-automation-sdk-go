/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Data type definitions file for service: VM.
 * Includes binding types of a structures and enumerations defined in the service.
 * Shared by client-side stubs and server-side skeletons to ensure type
 * compatibility.
 */
package vcenter

import (
	"reflect"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/bindings/com/vmware/vcenter/vm"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/bindings/com/vmware/vcenter/vm/hardware"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/bindings/com/vmware/vcenter/vm/hardware/adapter"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/bindings/com/vmware/vcenter/vm/hardware/boot"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/bindings"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/data"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/protocol"
)

// Resource type for virtual machines.
const VM_RESOURCE_TYPE = "VirtualMachine"


// The ``InventoryPlacementSpec`` class contains information used to place a virtual machine in the vCenter inventory.
type VMInventoryPlacementSpec struct {
    // Virtual machine folder into which the virtual machine should be placed.
	Folder *string
}

// The ``ComputePlacementSpec`` class contains information used to place a virtual machine on compute resources.
type VMComputePlacementSpec struct {
    // Resource pool into which the virtual machine should be placed.
	ResourcePool *string
    // Host onto which the virtual machine should be placed. 
    //
    //  If ``host`` and ``resourcePool`` are both specified, ``resourcePool`` must belong to ``host``. 
    //
    //  If ``host`` and ``cluster`` are both specified, ``host`` must be a member of ``cluster``.
	Host *string
    // Cluster into which the virtual machine should be placed. 
    //
    //  If ``cluster`` and ``resourcePool`` are both specified, ``resourcePool`` must belong to ``cluster``. 
    //
    //  If ``cluster`` and ``host`` are both specified, ``host`` must be a member of ``cluster``.
	Cluster *string
}

// The ``StoragePlacementSpec`` class contains information used to store a virtual machine's files.
type VMStoragePlacementSpec struct {
    // Datastore on which the virtual machine's configuration state should be stored. This datastore will also be used for any virtual disks that are created as part of the virtual machine creation operation.
	Datastore *string
}

// The ``PlacementSpec`` class contains information used to place a virtual machine onto resources within the vCenter inventory.
type VMPlacementSpec struct {
    // Virtual machine folder into which the virtual machine should be placed.
	Folder *string
    // Resource pool into which the virtual machine should be placed.
	ResourcePool *string
    // Host onto which the virtual machine should be placed. 
    //
    //  If ``host`` and ``resourcePool`` are both specified, ``resourcePool`` must belong to ``host``. 
    //
    //  If ``host`` and ``cluster`` are both specified, ``host`` must be a member of ``cluster``.
	Host *string
    // Cluster into which the virtual machine should be placed. 
    //
    //  If ``cluster`` and ``resourcePool`` are both specified, ``resourcePool`` must belong to ``cluster``. 
    //
    //  If ``cluster`` and ``host`` are both specified, ``host`` must be a member of ``cluster``.
	Cluster *string
    // Datastore on which the virtual machine's configuration state should be stored. This datastore will also be used for any virtual disks that are created as part of the virtual machine creation operation.
	Datastore *string
}

// The ``StoragePolicySpec`` class contains information about the storage policy to be associated with a virtual machine object.
type VMStoragePolicySpec struct {
    // Identifier of the storage policy which should be associated with the virtual machine.
	Policy string
}

// Document-based creation spec.
type VMCreateSpec struct {
    // Guest OS.
	GuestOS vm.GuestOS
    // Virtual machine name.
	Name *string
    // Virtual machine placement information.
	Placement *VMPlacementSpec
    // Virtual hardware version.
	HardwareVersion *vm.HardwareVersion
    // Boot configuration.
	Boot *hardware.BootCreateSpec
    // Boot device configuration.
	BootDevices []boot.DeviceEntryCreateSpec
    // CPU configuration.
	Cpu *hardware.CpuUpdateSpec
    // Memory configuration.
	Memory *hardware.MemoryUpdateSpec
    // List of disks.
	Disks []hardware.DiskCreateSpec
    // List of Ethernet adapters.
	Nics []hardware.EthernetCreateSpec
    // List of CD-ROMs.
	Cdroms []hardware.CdromCreateSpec
    // List of floppy drives.
	Floppies []hardware.FloppyCreateSpec
    // List of parallel ports.
	ParallelPorts []hardware.ParallelCreateSpec
    // List of serial ports.
	SerialPorts []hardware.SerialCreateSpec
    // List of SATA adapters.
	SataAdapters []adapter.SataCreateSpec
    // List of SCSI adapters.
	ScsiAdapters []adapter.ScsiCreateSpec
    // List of NVMe adapters. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
	NvmeAdapters []adapter.NvmeCreateSpec
    // The ``StoragePolicySpec`` class contains information about the storage policy that is to be associated with the virtual machine home (which contains the configuration and log files).
	StoragePolicy *VMStoragePolicySpec
}

// Document-based info.
type VMInfo struct {
    // Guest OS.
	GuestOS vm.GuestOS
    // Virtual machine name.
	Name string
    // Identity of the virtual machine.
	Identity *vm.IdentityInfo
    // Power state of the virtual machine.
	PowerState vm.PowerState
    // Indicates whether the virtual machine is frozen for instant clone, or not.
	InstantCloneFrozen *bool
    // Virtual hardware version information.
	Hardware vm.HardwareInfo
    // Boot configuration.
	Boot hardware.BootInfo
    // Boot device configuration. If the array has no entries, a server-specific default boot sequence is used.
	BootDevices []boot.DeviceEntry
    // CPU configuration.
	Cpu hardware.CpuInfo
    // Memory configuration.
	Memory hardware.MemoryInfo
    // List of disks.
	Disks map[string]hardware.DiskInfo
    // List of Ethernet adapters.
	Nics map[string]hardware.EthernetInfo
    // List of CD-ROMs.
	Cdroms map[string]hardware.CdromInfo
    // List of floppy drives.
	Floppies map[string]hardware.FloppyInfo
    // List of parallel ports.
	ParallelPorts map[string]hardware.ParallelInfo
    // List of serial ports.
	SerialPorts map[string]hardware.SerialInfo
    // List of SATA adapters.
	SataAdapters map[string]adapter.SataInfo
    // List of SCSI adapters.
	ScsiAdapters map[string]adapter.ScsiInfo
    // List of NVMe adapters. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
	NvmeAdapters map[string]adapter.NvmeInfo
}

// The ``GuestCustomizationSpec`` class contains information required to customize a virtual machine when deploying it. **Warning:** This class is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
type VMGuestCustomizationSpec struct {
    // Name of the customization specification. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
	Name *string
}

// Document-based disk clone spec. **Warning:** This class is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
type VMDiskCloneSpec struct {
    // Destination datastore to clone disk. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
	Datastore *string
    // Storage policy to be associated with cloned disk. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
	StoragePolicy *hardware.DiskStoragePolicySpec
}

// The ``ClonePlacementSpec`` class contains information used to place a clone of a virtual machine onto resources within the vCenter inventory. **Warning:** This class is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
type VMClonePlacementSpec struct {
    // Virtual machine folder into which the cloned virtual machine should be placed. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
	Folder *string
    // Resource pool into which the cloned virtual machine should be placed. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
	ResourcePool *string
    // Host onto which the cloned virtual machine should be placed. 
    //
    //  If ``host`` and ``resourcePool`` are both specified, ``resourcePool`` must belong to ``host``. 
    //
    //  If ``host`` and ``cluster`` are both specified, ``host`` must be a member of ``cluster``.. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
	Host *string
    // Cluster into which the cloned virtual machine should be placed. 
    //
    //  If ``cluster`` and ``resourcePool`` are both specified, ``resourcePool`` must belong to ``cluster``. 
    //
    //  If ``cluster`` and ``host`` are both specified, ``host`` must be a member of ``cluster``.. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
	Cluster *string
    // Datastore on which the cloned virtual machine's configuration state should be stored. This datastore will also be used for any virtual disks that are created as part of the virtual machine clone operation unless individually overridden. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
	Datastore *string
}

// Document-based clone spec. **Warning:** This class is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
type VMCloneSpec struct {
    // Virtual machine to clone from. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
	Source string
    // Virtual machine name. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
	Name string
    // Virtual machine placement information. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
	Placement *VMClonePlacementSpec
    // Set of Disks to Remove. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
	DisksToRemove map[string]bool
    // Map of Disks to Update. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
	DisksToUpdate map[string]VMDiskCloneSpec
    // Storage policy to be associated with the virtual machine home (which contains the configuration and log files). **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
	VmHomeStoragePolicy *VMStoragePolicySpec
    // Storage policy to be associated with the virtual machine disks. If set, each of the cloned virtual disks is assigned the specified storage policy, unless it is overwritten in VMCloneSpec#disksToUpdate. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
	DiskStoragePolicy *VMStoragePolicySpec
    // Flag used by the server to auto assign storage policy when a target storage policy is not explicitly specified for the method and the source virtual machine storage policy is not compatible with target datastore. If flag is set to true, the server will use the default storage policy of the target datastore. Otherwise, the server will throw an error for the method. See also the storage policy assignment rules for the VM#clone method. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
	FallbackToDatastoreDefaultPolicy *bool
    // Attempt to perform a VMCloneSpec#powerOn after clone. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
	PowerOn *bool
    // Guest customization spec to apply to the virtual machine after the virtual machine is deployed. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
	GuestCustomizationSpec *VMGuestCustomizationSpec
}

// Document-based disk relocate spec. **Warning:** This class is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
type VMDiskRelocateSpec struct {
    // Destination datastore to relocate disk. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
	Datastore *string
    // Storage policy to be associated with destination disk. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
	StoragePolicy *hardware.DiskStoragePolicySpec
}

// The ``RelocatePlacementSpec`` class contains information used to change the placement of an existing virtual machine within the vCenter inventory. **Warning:** This class is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
type VMRelocatePlacementSpec struct {
    // Virtual machine folder into which the virtual machine should be placed. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
	Folder *string
    // Resource pool into which the virtual machine should be placed. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
	ResourcePool *string
    // Host onto which the virtual machine should be placed. 
    //
    //  If ``host`` and ``resourcePool`` are both specified, ``resourcePool`` must belong to ``host``. 
    //
    //  If ``host`` and ``cluster`` are both specified, ``host`` must be a member of ``cluster``.. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
	Host *string
    // Cluster into which the virtual machine should be placed. 
    //
    //  If ``cluster`` and ``resourcePool`` are both specified, ``resourcePool`` must belong to ``cluster``. 
    //
    //  If ``cluster`` and ``host`` are both specified, ``host`` must be a member of ``cluster``.. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
	Cluster *string
    // Datastore on which the virtual machine's configuration state should be stored. This datastore will also be used for any virtual disks that are associated with the virtual machine, unless individually overridden. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
	Datastore *string
}

// Document-based relocate spec. **Warning:** This class is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
type VMRelocateSpec struct {
    // Virtual machine placement information. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
	Placement *VMRelocatePlacementSpec
    // Individual disk relocation map. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
	Disks map[string]VMDiskRelocateSpec
    // Storage policy to be associated with the virtual machine home (which contains the configuration and log files). **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
	VmHomeStoragePolicy *VMStoragePolicySpec
    // Storage policy to be associated with the virtual machine disks. If set, each of the destination disks is assigned the specified storage policy, unless it is overwritten in VMRelocateSpec#disks. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
	DiskStoragePolicy *VMStoragePolicySpec
    // Flag used by the server to auto assign storage policy when a storage policy is not explicitly specified for the method and the original virtual machine storage policy is not compatible with the target datastore. If flag is set to true, the server will use the default storage policy of the target datastore. Otherwise, the server will throw an error for the method. See also the storage policy assignment rules for the VM#relocate method. **Warning:** This property is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
	FallbackToDatastoreDefaultPolicy *bool
}

// The ``InstantClonePlacementSpec`` class contains information used to place an InstantClone of a virtual machine onto resources within the vCenter inventory.
type VMInstantClonePlacementSpec struct {
    // Virtual machine folder into which the InstantCloned virtual machine should be placed.
	Folder *string
    // Resource pool into which the InstantCloned virtual machine should be placed.
	ResourcePool *string
    // Datastore on which the InstantCloned virtual machine's configuration state should be stored. This datastore will also be used for any virtual disks that are created as part of the virtual machine InstantClone operation.
	Datastore *string
}

// Document-based InstantClone spec.
type VMInstantCloneSpec struct {
    // Virtual machine to InstantClone from.
	Source string
    // Name of the new virtual machine.
	Name string
    // Virtual machine placement information.
	Placement *VMInstantClonePlacementSpec
    // Map of NICs to update.
	NicsToUpdate map[string]hardware.EthernetUpdateSpec
    // Indicates whether all NICs on the destination virtual machine should be disconnected from the newtwork
	DisconnectAllNics *bool
    // Map of parallel ports to Update.
	ParallelPortsToUpdate map[string]hardware.ParallelUpdateSpec
    // Map of serial ports to Update.
	SerialPortsToUpdate map[string]hardware.SerialUpdateSpec
    // 128-bit SMBIOS UUID of a virtual machine represented as a hexadecimal string in "12345678-abcd-1234-cdef-123456789abc" format.
	BiosUuid *string
}

// The ``FilterSpec`` class contains properties used to filter the results when listing virtual machines (see VM#list). If multiple properties are specified, only virtual machines matching all of the properties match the filter.
type VMFilterSpec struct {
    // Identifiers of virtual machines that can match the filter.
	Vms map[string]bool
    // Names that virtual machines must have to match the filter (see VMInfo#name).
	Names map[string]bool
    // Folders that must contain the virtual machine for the virtual machine to match the filter.
	Folders map[string]bool
    // Datacenters that must contain the virtual machine for the virtual machine to match the filter.
	Datacenters map[string]bool
    // Hosts that must contain the virtual machine for the virtual machine to match the filter.
	Hosts map[string]bool
    // Clusters that must contain the virtual machine for the virtual machine to match the filter.
	Clusters map[string]bool
    // Resource pools that must contain the virtual machine for the virtual machine to match the filter.
	ResourcePools map[string]bool
    // Power states that a virtual machine must be in to match the filter (see vm.PowerInfo#state.
	PowerStates map[vm.PowerState]bool
}

// The ``Summary`` class contains commonly used information about a virtual machine.
type VMSummary struct {
    // Identifier of the virtual machine.
	Vm string
    // Name of the Virtual machine.
	Name string
    // Power state of the virtual machine.
	PowerState vm.PowerState
    // Number of CPU cores.
	CpuCount *int64
    // Memory size in mebibytes.
	MemorySizeMiB *int64
}

// The ``RegisterPlacementSpec`` class contains information used to place a virtual machine, created from existing virtual machine files on storage, onto resources within the vCenter inventory.
type VMRegisterPlacementSpec struct {
    // Virtual machine folder into which the virtual machine should be placed.
	Folder *string
    // Resource pool into which the virtual machine should be placed.
	ResourcePool *string
    // Host onto which the virtual machine should be placed. 
    //
    //  If ``host`` and ``resourcePool`` are both specified, ``resourcePool`` must belong to ``host``. 
    //
    //  If ``host`` and ``cluster`` are both specified, ``host`` must be a member of ``cluster``.
	Host *string
    // Cluster into which the virtual machine should be placed. 
    //
    //  If ``cluster`` and ``resourcePool`` are both specified, ``resourcePool`` must belong to ``cluster``. 
    //
    //  If ``cluster`` and ``host`` are both specified, ``host`` must be a member of ``cluster``.
	Cluster *string
}

// The ``RegisterSpec`` class contains information used to create a virtual machine from existing virtual machine files on storage. 
//
//  The location of the virtual machine files on storage must be specified by providing either VMRegisterSpec#datastore and VMRegisterSpec#path or by providing VMRegisterSpec#datastorePath. If VMRegisterSpec#datastore and VMRegisterSpec#path are map with bool value, VMRegisterSpec#datastorePath must be null, and if VMRegisterSpec#datastorePath is map with bool value, VMRegisterSpec#datastore and VMRegisterSpec#path must be null.
type VMRegisterSpec struct {
    // Identifier of the datastore on which the virtual machine's configuration state is stored.
	Datastore *string
    // Path to the virtual machine's configuration file on the datastore corresponding to {\\\\@link #datastore).
	Path *string
    // Datastore path for the virtual machine's configuration file in the format "[datastore name] path". For example "[storage1] Test-VM/Test-VM.vmx".
	DatastorePath *string
    // Virtual machine name.
	Name *string
    // Virtual machine placement information.
	Placement *VMRegisterPlacementSpec
}



func vMCreateInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["spec"] = bindings.NewReferenceType(VMCreateSpecBindingType)
	fieldNameMap["spec"] = "Spec"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func vMCreateOutputType() bindings.BindingType {
	return bindings.NewIdType([]string{"VirtualMachine"}, "")
}

func vMCreateRestMetadata() protocol.OperationRestMetadata {
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
		map[string]int{"AlreadyExists": 400,"Error": 500,"InvalidArgument": 400,"NotFound": 404,"ResourceInaccessible": 500,"ResourceInUse": 400,"ServiceUnavailable": 503,"UnableToAllocateResource": 500,"Unauthenticated": 401,"Unauthorized": 403,"Unsupported": 400})
}

func vMCloneInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["spec"] = bindings.NewReferenceType(VMCloneSpecBindingType)
	fieldNameMap["spec"] = "Spec"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func vMCloneOutputType() bindings.BindingType {
	return bindings.NewIdType([]string{"VirtualMachine"}, "")
}

func vMCloneRestMetadata() protocol.OperationRestMetadata {
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
		map[string]int{"AlreadyExists": 400,"Error": 500,"InvalidArgument": 400,"NotFound": 404,"ResourceInaccessible": 500,"ServiceUnavailable": 503,"UnableToAllocateResource": 500,"Unauthenticated": 401,"Unauthorized": 403})
}

func vMRelocateInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["vm"] = bindings.NewIdType([]string{"VirtualMachine"}, "")
	fields["spec"] = bindings.NewReferenceType(VMRelocateSpecBindingType)
	fieldNameMap["vm"] = "Vm"
	fieldNameMap["spec"] = "Spec"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func vMRelocateOutputType() bindings.BindingType {
	return bindings.NewVoidType()
}

func vMRelocateRestMetadata() protocol.OperationRestMetadata {
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
		map[string]int{"Error": 500,"InvalidArgument": 400,"NotFound": 404,"ResourceInaccessible": 500,"ServiceUnavailable": 503,"Unauthenticated": 401,"Unauthorized": 403})
}

func vMInstantCloneInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["spec"] = bindings.NewReferenceType(VMInstantCloneSpecBindingType)
	fieldNameMap["spec"] = "Spec"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func vMInstantCloneOutputType() bindings.BindingType {
	return bindings.NewIdType([]string{"VirtualMachine"}, "")
}

func vMInstantCloneRestMetadata() protocol.OperationRestMetadata {
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
		map[string]int{"AlreadyExists": 400,"Error": 500,"InvalidArgument": 400,"NotFound": 404,"ResourceInaccessible": 500,"ServiceUnavailable": 503,"UnableToAllocateResource": 500,"Unauthenticated": 401,"Unauthorized": 403})
}

func vMGetInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["vm"] = bindings.NewIdType([]string{"VirtualMachine"}, "")
	fieldNameMap["vm"] = "Vm"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func vMGetOutputType() bindings.BindingType {
	return bindings.NewReferenceType(VMInfoBindingType)
}

func vMGetRestMetadata() protocol.OperationRestMetadata {
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

func vMDeleteInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["vm"] = bindings.NewIdType([]string{"VirtualMachine"}, "")
	fieldNameMap["vm"] = "Vm"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func vMDeleteOutputType() bindings.BindingType {
	return bindings.NewVoidType()
}

func vMDeleteRestMetadata() protocol.OperationRestMetadata {
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
		map[string]int{"Error": 500,"NotAllowedInCurrentState": 400,"NotFound": 404,"ResourceBusy": 500,"ResourceInaccessible": 500,"ServiceUnavailable": 503,"Unauthenticated": 401,"Unauthorized": 403})
}

func vMListInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["filter"] = bindings.NewOptionalType(bindings.NewReferenceType(VMFilterSpecBindingType))
	fieldNameMap["filter"] = "Filter"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func vMListOutputType() bindings.BindingType {
	return bindings.NewListType(bindings.NewReferenceType(VMSummaryBindingType), reflect.TypeOf([]VMSummary{}))
}

func vMListRestMetadata() protocol.OperationRestMetadata {
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
		map[string]int{"InvalidArgument": 400,"UnableToAllocateResource": 500,"ServiceUnavailable": 503,"Unauthenticated": 401,"Unauthorized": 403})
}

func vMRegisterInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["spec"] = bindings.NewReferenceType(VMRegisterSpecBindingType)
	fieldNameMap["spec"] = "Spec"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func vMRegisterOutputType() bindings.BindingType {
	return bindings.NewIdType([]string{"VirtualMachine"}, "")
}

func vMRegisterRestMetadata() protocol.OperationRestMetadata {
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
		map[string]int{"AlreadyExists": 400,"Error": 500,"InvalidArgument": 400,"NotFound": 404,"ResourceInaccessible": 500,"ServiceUnavailable": 503,"UnableToAllocateResource": 500,"Unauthenticated": 401,"Unauthorized": 403})
}

func vMUnregisterInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["vm"] = bindings.NewIdType([]string{"VirtualMachine"}, "")
	fieldNameMap["vm"] = "Vm"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func vMUnregisterOutputType() bindings.BindingType {
	return bindings.NewVoidType()
}

func vMUnregisterRestMetadata() protocol.OperationRestMetadata {
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
		map[string]int{"NotAllowedInCurrentState": 400,"NotFound": 404,"ResourceBusy": 500,"ServiceUnavailable": 503,"Unauthenticated": 401,"Unauthorized": 403})
}


func VMInventoryPlacementSpecBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["folder"] = bindings.NewOptionalType(bindings.NewIdType([]string{"Folder"}, ""))
	fieldNameMap["folder"] = "Folder"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vcenter.VM.inventory_placement_spec", fields, reflect.TypeOf(VMInventoryPlacementSpec{}), fieldNameMap, validators)
}

func VMComputePlacementSpecBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["resource_pool"] = bindings.NewOptionalType(bindings.NewIdType([]string{"ResourcePool"}, ""))
	fieldNameMap["resource_pool"] = "ResourcePool"
	fields["host"] = bindings.NewOptionalType(bindings.NewIdType([]string{"HostSystem"}, ""))
	fieldNameMap["host"] = "Host"
	fields["cluster"] = bindings.NewOptionalType(bindings.NewIdType([]string{"ClusterComputeResource"}, ""))
	fieldNameMap["cluster"] = "Cluster"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vcenter.VM.compute_placement_spec", fields, reflect.TypeOf(VMComputePlacementSpec{}), fieldNameMap, validators)
}

func VMStoragePlacementSpecBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["datastore"] = bindings.NewOptionalType(bindings.NewIdType([]string{"Datastore"}, ""))
	fieldNameMap["datastore"] = "Datastore"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vcenter.VM.storage_placement_spec", fields, reflect.TypeOf(VMStoragePlacementSpec{}), fieldNameMap, validators)
}

func VMPlacementSpecBindingType() bindings.BindingType {
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
	fields["datastore"] = bindings.NewOptionalType(bindings.NewIdType([]string{"Datastore"}, ""))
	fieldNameMap["datastore"] = "Datastore"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vcenter.VM.placement_spec", fields, reflect.TypeOf(VMPlacementSpec{}), fieldNameMap, validators)
}

func VMStoragePolicySpecBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["policy"] = bindings.NewIdType([]string{"com.vmware.vcenter.StoragePolicy"}, "")
	fieldNameMap["policy"] = "Policy"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vcenter.VM.storage_policy_spec", fields, reflect.TypeOf(VMStoragePolicySpec{}), fieldNameMap, validators)
}

func VMCreateSpecBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["guest_OS"] = bindings.NewEnumType("com.vmware.vcenter.vm.guest_OS", reflect.TypeOf(vm.GuestOS(vm.GuestOS_DOS)))
	fieldNameMap["guest_OS"] = "GuestOS"
	fields["name"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["name"] = "Name"
	fields["placement"] = bindings.NewOptionalType(bindings.NewReferenceType(VMPlacementSpecBindingType))
	fieldNameMap["placement"] = "Placement"
	fields["hardware_version"] = bindings.NewOptionalType(bindings.NewEnumType("com.vmware.vcenter.vm.hardware.version", reflect.TypeOf(vm.HardwareVersion(vm.HardwareVersion_VMX_03))))
	fieldNameMap["hardware_version"] = "HardwareVersion"
	fields["boot"] = bindings.NewOptionalType(bindings.NewReferenceType(hardware.BootCreateSpecBindingType))
	fieldNameMap["boot"] = "Boot"
	fields["boot_devices"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewReferenceType(boot.DeviceEntryCreateSpecBindingType), reflect.TypeOf([]boot.DeviceEntryCreateSpec{})))
	fieldNameMap["boot_devices"] = "BootDevices"
	fields["cpu"] = bindings.NewOptionalType(bindings.NewReferenceType(hardware.CpuUpdateSpecBindingType))
	fieldNameMap["cpu"] = "Cpu"
	fields["memory"] = bindings.NewOptionalType(bindings.NewReferenceType(hardware.MemoryUpdateSpecBindingType))
	fieldNameMap["memory"] = "Memory"
	fields["disks"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewReferenceType(hardware.DiskCreateSpecBindingType), reflect.TypeOf([]hardware.DiskCreateSpec{})))
	fieldNameMap["disks"] = "Disks"
	fields["nics"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewReferenceType(hardware.EthernetCreateSpecBindingType), reflect.TypeOf([]hardware.EthernetCreateSpec{})))
	fieldNameMap["nics"] = "Nics"
	fields["cdroms"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewReferenceType(hardware.CdromCreateSpecBindingType), reflect.TypeOf([]hardware.CdromCreateSpec{})))
	fieldNameMap["cdroms"] = "Cdroms"
	fields["floppies"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewReferenceType(hardware.FloppyCreateSpecBindingType), reflect.TypeOf([]hardware.FloppyCreateSpec{})))
	fieldNameMap["floppies"] = "Floppies"
	fields["parallel_ports"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewReferenceType(hardware.ParallelCreateSpecBindingType), reflect.TypeOf([]hardware.ParallelCreateSpec{})))
	fieldNameMap["parallel_ports"] = "ParallelPorts"
	fields["serial_ports"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewReferenceType(hardware.SerialCreateSpecBindingType), reflect.TypeOf([]hardware.SerialCreateSpec{})))
	fieldNameMap["serial_ports"] = "SerialPorts"
	fields["sata_adapters"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewReferenceType(adapter.SataCreateSpecBindingType), reflect.TypeOf([]adapter.SataCreateSpec{})))
	fieldNameMap["sata_adapters"] = "SataAdapters"
	fields["scsi_adapters"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewReferenceType(adapter.ScsiCreateSpecBindingType), reflect.TypeOf([]adapter.ScsiCreateSpec{})))
	fieldNameMap["scsi_adapters"] = "ScsiAdapters"
	fields["nvme_adapters"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewReferenceType(adapter.NvmeCreateSpecBindingType), reflect.TypeOf([]adapter.NvmeCreateSpec{})))
	fieldNameMap["nvme_adapters"] = "NvmeAdapters"
	fields["storage_policy"] = bindings.NewOptionalType(bindings.NewReferenceType(VMStoragePolicySpecBindingType))
	fieldNameMap["storage_policy"] = "StoragePolicy"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vcenter.VM.create_spec", fields, reflect.TypeOf(VMCreateSpec{}), fieldNameMap, validators)
}

func VMInfoBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["guest_OS"] = bindings.NewEnumType("com.vmware.vcenter.vm.guest_OS", reflect.TypeOf(vm.GuestOS(vm.GuestOS_DOS)))
	fieldNameMap["guest_OS"] = "GuestOS"
	fields["name"] = bindings.NewStringType()
	fieldNameMap["name"] = "Name"
	fields["identity"] = bindings.NewOptionalType(bindings.NewReferenceType(vm.IdentityInfoBindingType))
	fieldNameMap["identity"] = "Identity"
	fields["power_state"] = bindings.NewEnumType("com.vmware.vcenter.vm.power.state", reflect.TypeOf(vm.PowerState(vm.PowerState_POWERED_OFF)))
	fieldNameMap["power_state"] = "PowerState"
	fields["instant_clone_frozen"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fieldNameMap["instant_clone_frozen"] = "InstantCloneFrozen"
	fields["hardware"] = bindings.NewReferenceType(vm.HardwareInfoBindingType)
	fieldNameMap["hardware"] = "Hardware"
	fields["boot"] = bindings.NewReferenceType(hardware.BootInfoBindingType)
	fieldNameMap["boot"] = "Boot"
	fields["boot_devices"] = bindings.NewListType(bindings.NewReferenceType(boot.DeviceEntryBindingType), reflect.TypeOf([]boot.DeviceEntry{}))
	fieldNameMap["boot_devices"] = "BootDevices"
	fields["cpu"] = bindings.NewReferenceType(hardware.CpuInfoBindingType)
	fieldNameMap["cpu"] = "Cpu"
	fields["memory"] = bindings.NewReferenceType(hardware.MemoryInfoBindingType)
	fieldNameMap["memory"] = "Memory"
	fields["disks"] = bindings.NewMapType(bindings.NewIdType([]string{"com.vmware.vcenter.vm.hardware.Disk"}, ""), bindings.NewReferenceType(hardware.DiskInfoBindingType),reflect.TypeOf(map[string]hardware.DiskInfo{}))
	fieldNameMap["disks"] = "Disks"
	fields["nics"] = bindings.NewMapType(bindings.NewIdType([]string{"com.vmware.vcenter.vm.hardware.Ethernet"}, ""), bindings.NewReferenceType(hardware.EthernetInfoBindingType),reflect.TypeOf(map[string]hardware.EthernetInfo{}))
	fieldNameMap["nics"] = "Nics"
	fields["cdroms"] = bindings.NewMapType(bindings.NewIdType([]string{"com.vmware.vcenter.vm.hardware.Cdrom"}, ""), bindings.NewReferenceType(hardware.CdromInfoBindingType),reflect.TypeOf(map[string]hardware.CdromInfo{}))
	fieldNameMap["cdroms"] = "Cdroms"
	fields["floppies"] = bindings.NewMapType(bindings.NewIdType([]string{"com.vmware.vcenter.vm.hardware.Floppy"}, ""), bindings.NewReferenceType(hardware.FloppyInfoBindingType),reflect.TypeOf(map[string]hardware.FloppyInfo{}))
	fieldNameMap["floppies"] = "Floppies"
	fields["parallel_ports"] = bindings.NewMapType(bindings.NewIdType([]string{"com.vmware.vcenter.vm.hardware.ParallelPort"}, ""), bindings.NewReferenceType(hardware.ParallelInfoBindingType),reflect.TypeOf(map[string]hardware.ParallelInfo{}))
	fieldNameMap["parallel_ports"] = "ParallelPorts"
	fields["serial_ports"] = bindings.NewMapType(bindings.NewIdType([]string{"com.vmware.vcenter.vm.hardware.SerialPort"}, ""), bindings.NewReferenceType(hardware.SerialInfoBindingType),reflect.TypeOf(map[string]hardware.SerialInfo{}))
	fieldNameMap["serial_ports"] = "SerialPorts"
	fields["sata_adapters"] = bindings.NewMapType(bindings.NewIdType([]string{"com.vmware.vcenter.vm.hardware.SataAdapter"}, ""), bindings.NewReferenceType(adapter.SataInfoBindingType),reflect.TypeOf(map[string]adapter.SataInfo{}))
	fieldNameMap["sata_adapters"] = "SataAdapters"
	fields["scsi_adapters"] = bindings.NewMapType(bindings.NewIdType([]string{"com.vmware.vcenter.vm.hardware.ScsiAdapter"}, ""), bindings.NewReferenceType(adapter.ScsiInfoBindingType),reflect.TypeOf(map[string]adapter.ScsiInfo{}))
	fieldNameMap["scsi_adapters"] = "ScsiAdapters"
	fields["nvme_adapters"] = bindings.NewOptionalType(bindings.NewMapType(bindings.NewIdType([]string{"com.vmware.vcenter.vm.hardware.NvmeAdapter"}, ""), bindings.NewReferenceType(adapter.NvmeInfoBindingType),reflect.TypeOf(map[string]adapter.NvmeInfo{})))
	fieldNameMap["nvme_adapters"] = "NvmeAdapters"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vcenter.VM.info", fields, reflect.TypeOf(VMInfo{}), fieldNameMap, validators)
}

func VMGuestCustomizationSpecBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["name"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["name"] = "Name"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vcenter.VM.guest_customization_spec", fields, reflect.TypeOf(VMGuestCustomizationSpec{}), fieldNameMap, validators)
}

func VMDiskCloneSpecBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["datastore"] = bindings.NewOptionalType(bindings.NewIdType([]string{"Datastore"}, ""))
	fieldNameMap["datastore"] = "Datastore"
	fields["storage_policy"] = bindings.NewOptionalType(bindings.NewReferenceType(hardware.DiskStoragePolicySpecBindingType))
	fieldNameMap["storage_policy"] = "StoragePolicy"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vcenter.VM.disk_clone_spec", fields, reflect.TypeOf(VMDiskCloneSpec{}), fieldNameMap, validators)
}

func VMClonePlacementSpecBindingType() bindings.BindingType {
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
	fields["datastore"] = bindings.NewOptionalType(bindings.NewIdType([]string{"Datastore"}, ""))
	fieldNameMap["datastore"] = "Datastore"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vcenter.VM.clone_placement_spec", fields, reflect.TypeOf(VMClonePlacementSpec{}), fieldNameMap, validators)
}

func VMCloneSpecBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["source"] = bindings.NewIdType([]string{"VirtualMachine"}, "")
	fieldNameMap["source"] = "Source"
	fields["name"] = bindings.NewStringType()
	fieldNameMap["name"] = "Name"
	fields["placement"] = bindings.NewOptionalType(bindings.NewReferenceType(VMClonePlacementSpecBindingType))
	fieldNameMap["placement"] = "Placement"
	fields["disks_to_remove"] = bindings.NewOptionalType(bindings.NewSetType(bindings.NewIdType([]string{"com.vmware.vcenter.vm.hardware.Disk"}, ""), reflect.TypeOf(map[string]bool{})))
	fieldNameMap["disks_to_remove"] = "DisksToRemove"
	fields["disks_to_update"] = bindings.NewOptionalType(bindings.NewMapType(bindings.NewIdType([]string{"com.vmware.vcenter.vm.hardware.Disk"}, ""), bindings.NewReferenceType(VMDiskCloneSpecBindingType),reflect.TypeOf(map[string]VMDiskCloneSpec{})))
	fieldNameMap["disks_to_update"] = "DisksToUpdate"
	fields["vm_home_storage_policy"] = bindings.NewOptionalType(bindings.NewReferenceType(VMStoragePolicySpecBindingType))
	fieldNameMap["vm_home_storage_policy"] = "VmHomeStoragePolicy"
	fields["disk_storage_policy"] = bindings.NewOptionalType(bindings.NewReferenceType(VMStoragePolicySpecBindingType))
	fieldNameMap["disk_storage_policy"] = "DiskStoragePolicy"
	fields["fallback_to_datastore_default_policy"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fieldNameMap["fallback_to_datastore_default_policy"] = "FallbackToDatastoreDefaultPolicy"
	fields["power_on"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fieldNameMap["power_on"] = "PowerOn"
	fields["guest_customization_spec"] = bindings.NewOptionalType(bindings.NewReferenceType(VMGuestCustomizationSpecBindingType))
	fieldNameMap["guest_customization_spec"] = "GuestCustomizationSpec"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vcenter.VM.clone_spec", fields, reflect.TypeOf(VMCloneSpec{}), fieldNameMap, validators)
}

func VMDiskRelocateSpecBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["datastore"] = bindings.NewOptionalType(bindings.NewIdType([]string{"Datastore"}, ""))
	fieldNameMap["datastore"] = "Datastore"
	fields["storage_policy"] = bindings.NewOptionalType(bindings.NewReferenceType(hardware.DiskStoragePolicySpecBindingType))
	fieldNameMap["storage_policy"] = "StoragePolicy"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vcenter.VM.disk_relocate_spec", fields, reflect.TypeOf(VMDiskRelocateSpec{}), fieldNameMap, validators)
}

func VMRelocatePlacementSpecBindingType() bindings.BindingType {
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
	fields["datastore"] = bindings.NewOptionalType(bindings.NewIdType([]string{"Datastore"}, ""))
	fieldNameMap["datastore"] = "Datastore"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vcenter.VM.relocate_placement_spec", fields, reflect.TypeOf(VMRelocatePlacementSpec{}), fieldNameMap, validators)
}

func VMRelocateSpecBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["placement"] = bindings.NewOptionalType(bindings.NewReferenceType(VMRelocatePlacementSpecBindingType))
	fieldNameMap["placement"] = "Placement"
	fields["disks"] = bindings.NewOptionalType(bindings.NewMapType(bindings.NewIdType([]string{"com.vmware.vcenter.vm.hardware.Disk"}, ""), bindings.NewReferenceType(VMDiskRelocateSpecBindingType),reflect.TypeOf(map[string]VMDiskRelocateSpec{})))
	fieldNameMap["disks"] = "Disks"
	fields["vm_home_storage_policy"] = bindings.NewOptionalType(bindings.NewReferenceType(VMStoragePolicySpecBindingType))
	fieldNameMap["vm_home_storage_policy"] = "VmHomeStoragePolicy"
	fields["disk_storage_policy"] = bindings.NewOptionalType(bindings.NewReferenceType(VMStoragePolicySpecBindingType))
	fieldNameMap["disk_storage_policy"] = "DiskStoragePolicy"
	fields["fallback_to_datastore_default_policy"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fieldNameMap["fallback_to_datastore_default_policy"] = "FallbackToDatastoreDefaultPolicy"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vcenter.VM.relocate_spec", fields, reflect.TypeOf(VMRelocateSpec{}), fieldNameMap, validators)
}

func VMInstantClonePlacementSpecBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["folder"] = bindings.NewOptionalType(bindings.NewIdType([]string{"Folder"}, ""))
	fieldNameMap["folder"] = "Folder"
	fields["resource_pool"] = bindings.NewOptionalType(bindings.NewIdType([]string{"ResourcePool"}, ""))
	fieldNameMap["resource_pool"] = "ResourcePool"
	fields["datastore"] = bindings.NewOptionalType(bindings.NewIdType([]string{"Datastore"}, ""))
	fieldNameMap["datastore"] = "Datastore"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vcenter.VM.instant_clone_placement_spec", fields, reflect.TypeOf(VMInstantClonePlacementSpec{}), fieldNameMap, validators)
}

func VMInstantCloneSpecBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["source"] = bindings.NewIdType([]string{"VirtualMachine"}, "")
	fieldNameMap["source"] = "Source"
	fields["name"] = bindings.NewStringType()
	fieldNameMap["name"] = "Name"
	fields["placement"] = bindings.NewOptionalType(bindings.NewReferenceType(VMInstantClonePlacementSpecBindingType))
	fieldNameMap["placement"] = "Placement"
	fields["nics_to_update"] = bindings.NewOptionalType(bindings.NewMapType(bindings.NewIdType([]string{"com.vmware.vcenter.vm.hardware.Ethernet"}, ""), bindings.NewReferenceType(hardware.EthernetUpdateSpecBindingType),reflect.TypeOf(map[string]hardware.EthernetUpdateSpec{})))
	fieldNameMap["nics_to_update"] = "NicsToUpdate"
	fields["disconnect_all_nics"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fieldNameMap["disconnect_all_nics"] = "DisconnectAllNics"
	fields["parallel_ports_to_update"] = bindings.NewOptionalType(bindings.NewMapType(bindings.NewIdType([]string{"com.vmware.vcenter.vm.hardware.ParallelPort"}, ""), bindings.NewReferenceType(hardware.ParallelUpdateSpecBindingType),reflect.TypeOf(map[string]hardware.ParallelUpdateSpec{})))
	fieldNameMap["parallel_ports_to_update"] = "ParallelPortsToUpdate"
	fields["serial_ports_to_update"] = bindings.NewOptionalType(bindings.NewMapType(bindings.NewIdType([]string{"com.vmware.vcenter.vm.hardware.SerialPort"}, ""), bindings.NewReferenceType(hardware.SerialUpdateSpecBindingType),reflect.TypeOf(map[string]hardware.SerialUpdateSpec{})))
	fieldNameMap["serial_ports_to_update"] = "SerialPortsToUpdate"
	fields["bios_uuid"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["bios_uuid"] = "BiosUuid"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vcenter.VM.instant_clone_spec", fields, reflect.TypeOf(VMInstantCloneSpec{}), fieldNameMap, validators)
}

func VMFilterSpecBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["vms"] = bindings.NewOptionalType(bindings.NewSetType(bindings.NewIdType([]string{"VirtualMachine"}, ""), reflect.TypeOf(map[string]bool{})))
	fieldNameMap["vms"] = "Vms"
	fields["names"] = bindings.NewOptionalType(bindings.NewSetType(bindings.NewStringType(), reflect.TypeOf(map[string]bool{})))
	fieldNameMap["names"] = "Names"
	fields["folders"] = bindings.NewOptionalType(bindings.NewSetType(bindings.NewIdType([]string{"Folder"}, ""), reflect.TypeOf(map[string]bool{})))
	fieldNameMap["folders"] = "Folders"
	fields["datacenters"] = bindings.NewOptionalType(bindings.NewSetType(bindings.NewIdType([]string{"Datacenter"}, ""), reflect.TypeOf(map[string]bool{})))
	fieldNameMap["datacenters"] = "Datacenters"
	fields["hosts"] = bindings.NewOptionalType(bindings.NewSetType(bindings.NewIdType([]string{"HostSystem"}, ""), reflect.TypeOf(map[string]bool{})))
	fieldNameMap["hosts"] = "Hosts"
	fields["clusters"] = bindings.NewOptionalType(bindings.NewSetType(bindings.NewIdType([]string{"ClusterComputeResource"}, ""), reflect.TypeOf(map[string]bool{})))
	fieldNameMap["clusters"] = "Clusters"
	fields["resource_pools"] = bindings.NewOptionalType(bindings.NewSetType(bindings.NewIdType([]string{"ResourcePool"}, ""), reflect.TypeOf(map[string]bool{})))
	fieldNameMap["resource_pools"] = "ResourcePools"
	fields["power_states"] = bindings.NewOptionalType(bindings.NewSetType(bindings.NewEnumType("com.vmware.vcenter.vm.power.state", reflect.TypeOf(vm.PowerState(vm.PowerState_POWERED_OFF))), reflect.TypeOf(map[vm.PowerState]bool{})))
	fieldNameMap["power_states"] = "PowerStates"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vcenter.VM.filter_spec", fields, reflect.TypeOf(VMFilterSpec{}), fieldNameMap, validators)
}

func VMSummaryBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["vm"] = bindings.NewIdType([]string{"VirtualMachine"}, "")
	fieldNameMap["vm"] = "Vm"
	fields["name"] = bindings.NewStringType()
	fieldNameMap["name"] = "Name"
	fields["power_state"] = bindings.NewEnumType("com.vmware.vcenter.vm.power.state", reflect.TypeOf(vm.PowerState(vm.PowerState_POWERED_OFF)))
	fieldNameMap["power_state"] = "PowerState"
	fields["cpu_count"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fieldNameMap["cpu_count"] = "CpuCount"
	fields["memory_size_MiB"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fieldNameMap["memory_size_MiB"] = "MemorySizeMiB"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vcenter.VM.summary", fields, reflect.TypeOf(VMSummary{}), fieldNameMap, validators)
}

func VMRegisterPlacementSpecBindingType() bindings.BindingType {
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
	return bindings.NewStructType("com.vmware.vcenter.VM.register_placement_spec", fields, reflect.TypeOf(VMRegisterPlacementSpec{}), fieldNameMap, validators)
}

func VMRegisterSpecBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["datastore"] = bindings.NewOptionalType(bindings.NewIdType([]string{"Datastore"}, ""))
	fieldNameMap["datastore"] = "Datastore"
	fields["path"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["path"] = "Path"
	fields["datastore_path"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["datastore_path"] = "DatastorePath"
	fields["name"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["name"] = "Name"
	fields["placement"] = bindings.NewOptionalType(bindings.NewReferenceType(VMRegisterPlacementSpecBindingType))
	fieldNameMap["placement"] = "Placement"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vcenter.VM.register_spec", fields, reflect.TypeOf(VMRegisterSpec{}), fieldNameMap, validators)
}

