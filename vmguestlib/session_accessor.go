package vmguestlib

/*
#cgo CFLAGS: -I../vendor -I../native
#cgo LDFLAGS: -L/usr/lib/vmware-tools/lib/libvmtools.so -L/usr/lib/vmware-tools/lib/libvmGuestLib.so -lvmtools -lvmGuestLib
#include <vmGuestLib.h>
#include <vmSessionId.h>
#include <vmGuestLibProxy.h>
*/
import "C"

// GetCPUReservation retrieves the minimum processing power in MHz available to the virtual machine.
func (s *Session) GetCPUReservation() (uint32, error) {
	return s.Handle.getUint32Value(C.p_uint32_f(C.VMGuestLib_GetCpuReservationMHz))
}

// GetCPULimit retrieves the maximum processing power in MHz available to the virtual machine.
func (s *Session) GetCPULimit() (uint32, error) {
	return s.Handle.getUint32Value(C.p_uint32_f(C.VMGuestLib_GetCpuLimitMHz))
}

// GetCPUShares retrieves the number of CPU shares allocated to the virtual machine.
func (s *Session) GetCPUShares() (uint32, error) {
	return s.Handle.getUint32Value(C.p_uint32_f(C.VMGuestLib_GetCpuShares))
}

// GetCPUUsed Retrieves the number of milliseconds during which the virtual machine has been using the CPU.
func (s *Session) GetCPUUsed() (uint32, error) {
	return s.Handle.getUint32Value(C.p_uint32_f(C.VMGuestLib_GetCpuUsedMs))
}

// GetHostProcessorSpeed retrieves the host processor speed.
func (s *Session) GetHostProcessorSpeed() (uint32, error) {
	return s.Handle.getUint32Value(C.p_uint32_f(C.VMGuestLib_GetHostProcessorSpeed))
}

// GetMemReservation retrieves the minimum amount of memory in MB that is available to the virtual machine.
func (s *Session) GetMemReservation() (uint32, error) {
	return s.Handle.getUint32Value(C.p_uint32_f(C.VMGuestLib_GetMemReservationMB))
}

// GetMemLimit retrieves the maximum amount of memory in MB that is available to the virtual machine.
func (s *Session) GetMemLimit() (uint32, error) {
	return s.Handle.getUint32Value(C.p_uint32_f(C.VMGuestLib_GetMemLimitMB))
}

// GetMemShares retrieves the number of memory in MB shares allocated to the virtual machine.
func (s *Session) GetMemShares() (uint32, error) {
	return s.Handle.getUint32Value(C.p_uint32_f(C.VMGuestLib_GetMemShares))
}

// GetMemMapped retrieves the mapped memory in MB size of this virtual machine.
func (s *Session) GetMemMapped() (uint32, error) {
	return s.Handle.getUint32Value(C.p_uint32_f(C.VMGuestLib_GetMemMappedMB))
}

// GetMemActive retrieves the estimated amount of memory in MB the virtual machine is actively using.
func (s *Session) GetMemActive() (uint32, error) {
	return s.Handle.getUint32Value(C.p_uint32_f(C.VMGuestLib_GetMemActiveMB))
}

// GetMemOverhead retrieves the amount of overhead memory in MB associated with this virtual machine consumed on the host system.
func (s *Session) GetMemOverhead() (uint32, error) {
	return s.Handle.getUint32Value(C.p_uint32_f(C.VMGuestLib_GetMemOverheadMB))
}

// GetMemBallooned retrieves the amount of memory in MB that has been reclaimed from this virtual machine via the VMware Memory Balloon mechanism.
func (s *Session) GetMemBallooned() (uint32, error) {
	return s.Handle.getUint32Value(C.p_uint32_f(C.VMGuestLib_GetMemBalloonedMB))
}

// GetMemSwapped retrieves the amount of memory in MB associated with this virtual machine that has been swapped by the host system.
func (s *Session) GetMemSwapped() (uint32, error) {
	return s.Handle.getUint32Value(C.p_uint32_f(C.VMGuestLib_GetMemSwappedMB))
}

// GetMemShared retrieves the amount of physical memory in MB associated with this virtual machine that is copy-on-write (COW) shared on the host.
func (s *Session) GetMemShared() (uint32, error) {
	return s.Handle.getUint32Value(C.p_uint32_f(C.VMGuestLib_GetMemSharedMB))
}

// GetMemSharedSaved retrieves the estimated amount of physical memory in MB on the host saved from copy-on-write (COW) shared guest physical memory in MB.
func (s *Session) GetMemSharedSaved() (uint32, error) {
	return s.Handle.getUint32Value(C.p_uint32_f(C.VMGuestLib_GetMemSharedSavedMB))
}

// GetMemUsed retrieves the estimated amount of physical host memory in MB currently consumed for this virtual machine's physical memory in MB.
func (s *Session) GetMemUsed() (uint32, error) {
	return s.Handle.getUint32Value(C.p_uint32_f(C.VMGuestLib_GetMemUsedMB))
}

// GetHostNumCPUCores retrieves the number of physical CPU cores on the host machine.
func (s *Session) GetHostNumCPUCores() (uint32, error) {
	return s.Handle.getUint32Value(C.p_uint32_f(C.VMGuestLib_GetHostNumCpuCores))
}

// GetTimeElapsed retrieves the number of milliseconds that have passed in real time since the virtual machine started running on the current host system.
func (s *Session) GetTimeElapsed() (uint64, error) {
	return s.Handle.getUint64Value(C.p_uint64_f(C.VMGuestLib_GetElapsedMs))
}

// GetCpuStolen retrieves the time in milliseconds that the VM was runnable but not scheduled to run.
func (s *Session) GetCpuStolen() (uint64, error) {
	return s.Handle.getUint64Value(C.p_uint64_f(C.VMGuestLib_GetCpuStolenMs))
}

// GetMemTargetSize retrieves the ùemory Target Size.
func (s *Session) GetMemTargetSize() (uint64, error) {
	return s.Handle.getUint64Value(C.p_uint64_f(C.VMGuestLib_GetMemTargetSizeMB))
}

// GetHostCPUUsed retrieves the total CPU time in milliseconds used by host.
func (s *Session) GetHostCPUUsed() (uint64, error) {
	return s.Handle.getUint64Value(C.p_uint64_f(C.VMGuestLib_GetHostCpuUsedMs))
}

// GetHostMemSwapped retrieves the total memory in MB swapped out on the host.
func (s *Session) GetHostMemSwapped() (uint64, error) {
	return s.Handle.getUint64Value(C.p_uint64_f(C.VMGuestLib_GetHostMemSwappedMB))
}

// GetHostMemShared retrieves the total COW (Copy-On-Write) memory in MB on host.
func (s *Session) GetHostMemShared() (uint64, error) {
	return s.Handle.getUint64Value(C.p_uint64_f(C.VMGuestLib_GetHostMemSharedMB))
}

// GetHostMemUsed retrieves the total consumed memory in MB on host.
func (s *Session) GetHostMemUsed() (uint64, error) {
	return s.Handle.getUint64Value(C.p_uint64_f(C.VMGuestLib_GetHostMemUsedMB))
}

// GetHostMemPhys retrieves the total memory in MB available to host OS kernel.
func (s *Session) GetHostMemPhys() (uint64, error) {
	return s.Handle.getUint64Value(C.p_uint64_f(C.VMGuestLib_GetHostMemPhysMB))
}

// GetHostMemPhysFree retrieves the total physical memory in MB free on host.
func (s *Session) GetHostMemPhysFree() (uint64, error) {
	return s.Handle.getUint64Value(C.p_uint64_f(C.VMGuestLib_GetHostMemPhysFreeMB))
}

// GetHostMemKernOvhd retrieves the total host kernel memory in MB overhead.
func (s *Session) GetHostMemKernOvhd() (uint64, error) {
	return s.Handle.getUint64Value(C.p_uint64_f(C.VMGuestLib_GetHostMemKernOvhdMB))
}

// GetHostMemMapped retrieves the total mapped memory in MB on host.
func (s *Session) GetHostMemMapped() (uint64, error) {
	return s.Handle.getUint64Value(C.p_uint64_f(C.VMGuestLib_GetHostMemMappedMB))
}

// GetHostMemUnmapped retrieves the total unmapped memory in MB on host.
func (s *Session) GetHostMemUnmapped() (uint64, error) {
	return s.Handle.getUint64Value(C.p_uint64_f(C.VMGuestLib_GetHostMemUnmappedMB))
}
