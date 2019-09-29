/*
 * Copyright 2019 VMware, Inc.  All rights reserved.
 */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Data type definitions file for service: VM.
 * Includes binding types of a structures and enumerations defined in the service.
 * Shared by client-side stubs and server-side skeletons to ensure type
 * compatibility.
 */

package VM

import (
    "reflect"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/bindings/vcenter/vm/Hardware"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/bindings/vcenter/vm/hardware"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/bindings/vcenter/vm/hardware/Disk"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/bindings/vcenter/vm/hardware/adapter/sata"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/bindings/vcenter/vm/hardware/adapter/scsi"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/bindings/vcenter/vm/hardware/boot"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/bindings/vcenter/vm/hardware/boot/device"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/bindings/vcenter/vm/hardware/cdrom"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/bindings/vcenter/vm/hardware/cpu"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/bindings/vcenter/vm/hardware/ethernet"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/bindings/vcenter/vm/hardware/floppy"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/bindings/vcenter/vm/hardware/memory"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/bindings/vcenter/vm/hardware/parallel"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/bindings/vcenter/vm/hardware/serial"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/bindings/vcenter/vm/power"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/bindings"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/data"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/protocol"
)

// Resource type for virtual machines.
const VM_RESOURCE_TYPE = "VirtualMachine"



// The ``InventoryPlacementSpec`` class contains information used to place a virtual machine in the vCenter inventory.
 type InventoryPlacementSpec struct {
    Folder *string
}






// The ``ComputePlacementSpec`` class contains information used to place a virtual machine on compute resources.
 type ComputePlacementSpec struct {
    ResourcePool *string
    Host *string
    Cluster *string
}






// The ``StoragePlacementSpec`` class contains information used to store a virtual machine's files.
 type StoragePlacementSpec struct {
    Datastore *string
}






// The ``PlacementSpec`` class contains information used to place a virtual machine onto resources within the vCenter inventory.
 type PlacementSpec struct {
    Folder *string
    ResourcePool *string
    Host *string
    Cluster *string
    Datastore *string
}






// The ``StoragePolicySpec`` class contains information about the storage policy to be associated with a virtual machine object. This class was added in vSphere API 6.7.
 type StoragePolicySpec struct {
    Policy string
}






// Document-based creation spec.
 type CreateSpec struct {
    GuestOS GuestOS
    Name *string
    Placement *PlacementSpec
    HardwareVersion *Hardware.Version
    Boot *boot.CreateSpec
    BootDevices []device.EntryCreateSpec
    Cpu *cpu.UpdateSpec
    Memory *memory.UpdateSpec
    Disks []Disk.CreateSpec
    Nics []ethernet.CreateSpec
    Cdroms []cdrom.CreateSpec
    Floppies []floppy.CreateSpec
    ParallelPorts []parallel.CreateSpec
    SerialPorts []serial.CreateSpec
    SataAdapters []sata.CreateSpec
    ScsiAdapters []scsi.CreateSpec
    StoragePolicy *StoragePolicySpec
}






// Document-based info.
 type Info struct {
    GuestOS GuestOS
    Name string
    PowerState power.State
    Hardware hardware.Info
    Boot boot.Info
    BootDevices []device.Entry
    Cpu cpu.Info
    Memory memory.Info
    Disks map[string]Disk.Info
    Nics map[string]ethernet.Info
    Cdroms map[string]cdrom.Info
    Floppies map[string]floppy.Info
    ParallelPorts map[string]parallel.Info
    SerialPorts map[string]serial.Info
    SataAdapters map[string]sata.Info
    ScsiAdapters map[string]scsi.Info
}






// The ``FilterSpec`` class contains properties used to filter the results when listing virtual machines (see VM#list). If multiple properties are specified, only virtual machines matching all of the properties match the filter.
 type FilterSpec struct {
    Vms map[string]bool
    Names map[string]bool
    Folders map[string]bool
    Datacenters map[string]bool
    Hosts map[string]bool
    Clusters map[string]bool
    ResourcePools map[string]bool
    PowerStates map[power.State]bool
}






// The ``Summary`` class contains commonly used information about a virtual machine.
 type Summary struct {
    Vm string
    Name string
    PowerState power.State
    CpuCount *int64
    MemorySizeMiB *int64
}









func createInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["spec"] = bindings.NewReferenceType(CreateSpecBindingType)
    fieldNameMap["spec"] = "Spec"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func createOutputType() bindings.BindingType {
    return bindings.NewIdType([]string {"VirtualMachine"}, "")
}

func createRestMetadata() protocol.OperationRestMetadata {
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


func getInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["vm"] = bindings.NewIdType([]string {"VirtualMachine"}, "")
    fieldNameMap["vm"] = "Vm"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func getOutputType() bindings.BindingType {
    return bindings.NewReferenceType(InfoBindingType)
}

func getRestMetadata() protocol.OperationRestMetadata {
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


func deleteInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["vm"] = bindings.NewIdType([]string {"VirtualMachine"}, "")
    fieldNameMap["vm"] = "Vm"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func deleteOutputType() bindings.BindingType {
    return bindings.NewVoidType()
}

func deleteRestMetadata() protocol.OperationRestMetadata {
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


func listInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["filter"] = bindings.NewOptionalType(bindings.NewReferenceType(FilterSpecBindingType))
    fieldNameMap["filter"] = "Filter"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func listOutputType() bindings.BindingType {
    return bindings.NewListType(bindings.NewReferenceType(SummaryBindingType), reflect.TypeOf([]Summary{}))
}

func listRestMetadata() protocol.OperationRestMetadata {
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



func InventoryPlacementSpecBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["folder"] = bindings.NewOptionalType(bindings.NewIdType([]string {"Folder"}, ""))
    fieldNameMap["folder"] = "Folder"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.VM.inventory_placement_spec",fields, reflect.TypeOf(InventoryPlacementSpec{}), fieldNameMap, validators)
}

func ComputePlacementSpecBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["resource_pool"] = bindings.NewOptionalType(bindings.NewIdType([]string {"ResourcePool"}, ""))
    fieldNameMap["resource_pool"] = "ResourcePool"
    fields["host"] = bindings.NewOptionalType(bindings.NewIdType([]string {"HostSystem"}, ""))
    fieldNameMap["host"] = "Host"
    fields["cluster"] = bindings.NewOptionalType(bindings.NewIdType([]string {"ClusterComputeResource"}, ""))
    fieldNameMap["cluster"] = "Cluster"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.VM.compute_placement_spec",fields, reflect.TypeOf(ComputePlacementSpec{}), fieldNameMap, validators)
}

func StoragePlacementSpecBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["datastore"] = bindings.NewOptionalType(bindings.NewIdType([]string {"Datastore"}, ""))
    fieldNameMap["datastore"] = "Datastore"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.VM.storage_placement_spec",fields, reflect.TypeOf(StoragePlacementSpec{}), fieldNameMap, validators)
}

func PlacementSpecBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["folder"] = bindings.NewOptionalType(bindings.NewIdType([]string {"Folder"}, ""))
    fieldNameMap["folder"] = "Folder"
    fields["resource_pool"] = bindings.NewOptionalType(bindings.NewIdType([]string {"ResourcePool"}, ""))
    fieldNameMap["resource_pool"] = "ResourcePool"
    fields["host"] = bindings.NewOptionalType(bindings.NewIdType([]string {"HostSystem"}, ""))
    fieldNameMap["host"] = "Host"
    fields["cluster"] = bindings.NewOptionalType(bindings.NewIdType([]string {"ClusterComputeResource"}, ""))
    fieldNameMap["cluster"] = "Cluster"
    fields["datastore"] = bindings.NewOptionalType(bindings.NewIdType([]string {"Datastore"}, ""))
    fieldNameMap["datastore"] = "Datastore"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.VM.placement_spec",fields, reflect.TypeOf(PlacementSpec{}), fieldNameMap, validators)
}

func StoragePolicySpecBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["policy"] = bindings.NewIdType([]string {"com.vmware.vcenter.StoragePolicy"}, "")
    fieldNameMap["policy"] = "Policy"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.VM.storage_policy_spec",fields, reflect.TypeOf(StoragePolicySpec{}), fieldNameMap, validators)
}

func CreateSpecBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["guest_OS"] = bindings.NewEnumType("com.vmware.vcenter.vm.guest_OS", reflect.TypeOf(GuestOS(GuestOS_DOS)))
    fieldNameMap["guest_OS"] = "GuestOS"
    fields["name"] = bindings.NewOptionalType(bindings.NewStringType())
    fieldNameMap["name"] = "Name"
    fields["placement"] = bindings.NewOptionalType(bindings.NewReferenceType(PlacementSpecBindingType))
    fieldNameMap["placement"] = "Placement"
    fields["hardware_version"] = bindings.NewOptionalType(bindings.NewEnumType("com.vmware.vcenter.vm.hardware.version", reflect.TypeOf(Hardware.Version(Hardware.Version_VMX_03))))
    fieldNameMap["hardware_version"] = "HardwareVersion"
    fields["boot"] = bindings.NewOptionalType(bindings.NewReferenceType(boot.CreateSpecBindingType))
    fieldNameMap["boot"] = "Boot"
    fields["boot_devices"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewReferenceType(device.EntryCreateSpecBindingType), reflect.TypeOf([]device.EntryCreateSpec{})))
    fieldNameMap["boot_devices"] = "BootDevices"
    fields["cpu"] = bindings.NewOptionalType(bindings.NewReferenceType(cpu.UpdateSpecBindingType))
    fieldNameMap["cpu"] = "Cpu"
    fields["memory"] = bindings.NewOptionalType(bindings.NewReferenceType(memory.UpdateSpecBindingType))
    fieldNameMap["memory"] = "Memory"
    fields["disks"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewReferenceType(Disk.CreateSpecBindingType), reflect.TypeOf([]Disk.CreateSpec{})))
    fieldNameMap["disks"] = "Disks"
    fields["nics"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewReferenceType(ethernet.CreateSpecBindingType), reflect.TypeOf([]ethernet.CreateSpec{})))
    fieldNameMap["nics"] = "Nics"
    fields["cdroms"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewReferenceType(cdrom.CreateSpecBindingType), reflect.TypeOf([]cdrom.CreateSpec{})))
    fieldNameMap["cdroms"] = "Cdroms"
    fields["floppies"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewReferenceType(floppy.CreateSpecBindingType), reflect.TypeOf([]floppy.CreateSpec{})))
    fieldNameMap["floppies"] = "Floppies"
    fields["parallel_ports"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewReferenceType(parallel.CreateSpecBindingType), reflect.TypeOf([]parallel.CreateSpec{})))
    fieldNameMap["parallel_ports"] = "ParallelPorts"
    fields["serial_ports"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewReferenceType(serial.CreateSpecBindingType), reflect.TypeOf([]serial.CreateSpec{})))
    fieldNameMap["serial_ports"] = "SerialPorts"
    fields["sata_adapters"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewReferenceType(sata.CreateSpecBindingType), reflect.TypeOf([]sata.CreateSpec{})))
    fieldNameMap["sata_adapters"] = "SataAdapters"
    fields["scsi_adapters"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewReferenceType(scsi.CreateSpecBindingType), reflect.TypeOf([]scsi.CreateSpec{})))
    fieldNameMap["scsi_adapters"] = "ScsiAdapters"
    fields["storage_policy"] = bindings.NewOptionalType(bindings.NewReferenceType(StoragePolicySpecBindingType))
    fieldNameMap["storage_policy"] = "StoragePolicy"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.VM.create_spec",fields, reflect.TypeOf(CreateSpec{}), fieldNameMap, validators)
}

func InfoBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["guest_OS"] = bindings.NewEnumType("com.vmware.vcenter.vm.guest_OS", reflect.TypeOf(GuestOS(GuestOS_DOS)))
    fieldNameMap["guest_OS"] = "GuestOS"
    fields["name"] = bindings.NewStringType()
    fieldNameMap["name"] = "Name"
    fields["power_state"] = bindings.NewEnumType("com.vmware.vcenter.vm.power.state", reflect.TypeOf(power.State(power.State_POWERED_OFF)))
    fieldNameMap["power_state"] = "PowerState"
    fields["hardware"] = bindings.NewReferenceType(hardware.InfoBindingType)
    fieldNameMap["hardware"] = "Hardware"
    fields["boot"] = bindings.NewReferenceType(boot.InfoBindingType)
    fieldNameMap["boot"] = "Boot"
    fields["boot_devices"] = bindings.NewListType(bindings.NewReferenceType(device.EntryBindingType), reflect.TypeOf([]device.Entry{}))
    fieldNameMap["boot_devices"] = "BootDevices"
    fields["cpu"] = bindings.NewReferenceType(cpu.InfoBindingType)
    fieldNameMap["cpu"] = "Cpu"
    fields["memory"] = bindings.NewReferenceType(memory.InfoBindingType)
    fieldNameMap["memory"] = "Memory"
    fields["disks"] = bindings.NewMapType(bindings.NewIdType([]string {"com.vmware.vcenter.vm.hardware.Disk"}, ""), bindings.NewReferenceType(Disk.InfoBindingType),reflect.TypeOf(map[string]Disk.Info{}))
    fieldNameMap["disks"] = "Disks"
    fields["nics"] = bindings.NewMapType(bindings.NewIdType([]string {"com.vmware.vcenter.vm.hardware.Ethernet"}, ""), bindings.NewReferenceType(ethernet.InfoBindingType),reflect.TypeOf(map[string]ethernet.Info{}))
    fieldNameMap["nics"] = "Nics"
    fields["cdroms"] = bindings.NewMapType(bindings.NewIdType([]string {"com.vmware.vcenter.vm.hardware.Cdrom"}, ""), bindings.NewReferenceType(cdrom.InfoBindingType),reflect.TypeOf(map[string]cdrom.Info{}))
    fieldNameMap["cdroms"] = "Cdroms"
    fields["floppies"] = bindings.NewMapType(bindings.NewIdType([]string {"com.vmware.vcenter.vm.hardware.Floppy"}, ""), bindings.NewReferenceType(floppy.InfoBindingType),reflect.TypeOf(map[string]floppy.Info{}))
    fieldNameMap["floppies"] = "Floppies"
    fields["parallel_ports"] = bindings.NewMapType(bindings.NewIdType([]string {"com.vmware.vcenter.vm.hardware.ParallelPort"}, ""), bindings.NewReferenceType(parallel.InfoBindingType),reflect.TypeOf(map[string]parallel.Info{}))
    fieldNameMap["parallel_ports"] = "ParallelPorts"
    fields["serial_ports"] = bindings.NewMapType(bindings.NewIdType([]string {"com.vmware.vcenter.vm.hardware.SerialPort"}, ""), bindings.NewReferenceType(serial.InfoBindingType),reflect.TypeOf(map[string]serial.Info{}))
    fieldNameMap["serial_ports"] = "SerialPorts"
    fields["sata_adapters"] = bindings.NewMapType(bindings.NewIdType([]string {"com.vmware.vcenter.vm.hardware.SataAdapter"}, ""), bindings.NewReferenceType(sata.InfoBindingType),reflect.TypeOf(map[string]sata.Info{}))
    fieldNameMap["sata_adapters"] = "SataAdapters"
    fields["scsi_adapters"] = bindings.NewMapType(bindings.NewIdType([]string {"com.vmware.vcenter.vm.hardware.ScsiAdapter"}, ""), bindings.NewReferenceType(scsi.InfoBindingType),reflect.TypeOf(map[string]scsi.Info{}))
    fieldNameMap["scsi_adapters"] = "ScsiAdapters"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.VM.info",fields, reflect.TypeOf(Info{}), fieldNameMap, validators)
}

func FilterSpecBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["vms"] = bindings.NewOptionalType(bindings.NewSetType(bindings.NewIdType([]string {"VirtualMachine"}, ""), reflect.TypeOf(map[string]bool{})))
    fieldNameMap["vms"] = "Vms"
    fields["names"] = bindings.NewOptionalType(bindings.NewSetType(bindings.NewStringType(), reflect.TypeOf(map[string]bool{})))
    fieldNameMap["names"] = "Names"
    fields["folders"] = bindings.NewOptionalType(bindings.NewSetType(bindings.NewIdType([]string {"Folder"}, ""), reflect.TypeOf(map[string]bool{})))
    fieldNameMap["folders"] = "Folders"
    fields["datacenters"] = bindings.NewOptionalType(bindings.NewSetType(bindings.NewIdType([]string {"Datacenter"}, ""), reflect.TypeOf(map[string]bool{})))
    fieldNameMap["datacenters"] = "Datacenters"
    fields["hosts"] = bindings.NewOptionalType(bindings.NewSetType(bindings.NewIdType([]string {"HostSystem"}, ""), reflect.TypeOf(map[string]bool{})))
    fieldNameMap["hosts"] = "Hosts"
    fields["clusters"] = bindings.NewOptionalType(bindings.NewSetType(bindings.NewIdType([]string {"ClusterComputeResource"}, ""), reflect.TypeOf(map[string]bool{})))
    fieldNameMap["clusters"] = "Clusters"
    fields["resource_pools"] = bindings.NewOptionalType(bindings.NewSetType(bindings.NewIdType([]string {"ResourcePool"}, ""), reflect.TypeOf(map[string]bool{})))
    fieldNameMap["resource_pools"] = "ResourcePools"
    fields["power_states"] = bindings.NewOptionalType(bindings.NewSetType(bindings.NewEnumType("com.vmware.vcenter.vm.power.state", reflect.TypeOf(power.State(power.State_POWERED_OFF))), reflect.TypeOf(map[power.State]bool{})))
    fieldNameMap["power_states"] = "PowerStates"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.VM.filter_spec",fields, reflect.TypeOf(FilterSpec{}), fieldNameMap, validators)
}

func SummaryBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["vm"] = bindings.NewIdType([]string {"VirtualMachine"}, "")
    fieldNameMap["vm"] = "Vm"
    fields["name"] = bindings.NewStringType()
    fieldNameMap["name"] = "Name"
    fields["power_state"] = bindings.NewEnumType("com.vmware.vcenter.vm.power.state", reflect.TypeOf(power.State(power.State_POWERED_OFF)))
    fieldNameMap["power_state"] = "PowerState"
    fields["cpu_count"] = bindings.NewOptionalType(bindings.NewIntegerType())
    fieldNameMap["cpu_count"] = "CpuCount"
    fields["memory_size_MiB"] = bindings.NewOptionalType(bindings.NewIntegerType())
    fieldNameMap["memory_size_MiB"] = "MemorySizeMiB"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.VM.summary",fields, reflect.TypeOf(Summary{}), fieldNameMap, validators)
}


