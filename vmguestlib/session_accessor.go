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

// GetHostProcessorSpeed retrieves the host processor speed.
func (s *Session) GetHostProcessorSpeed() (uint32, error) {
    return s.Handle.getUint32Value(C.p_uint32_f(C.VMGuestLib_GetHostProcessorSpeed))
}

// GetMemReservation retrieves the minimum amount of memory that is available to the virtual machine.
func (s *Session) GetMemReservation() (uint32, error) {
    return s.Handle.getUint32Value(C.p_uint32_f(C.VMGuestLib_GetMemReservationMB))
}

// GetMemLimit retrieves the maximum amount of memory that is available to the virtual machine.
func (s *Session) GetMemLimit() (uint32, error) {
    return s.Handle.getUint32Value(C.p_uint32_f(C.VMGuestLib_GetMemLimitMB))
}

// GetMemShares retrieves the number of memory shares allocated to the virtual machine.
func (s *Session) GetMemShares() (uint32, error) {
    return s.Handle.getUint32Value(C.p_uint32_f(C.VMGuestLib_GetMemShares))
}

// GetMemMapped retrieves the mapped memory size of this virtual machine.
func (s *Session) GetMemMapped() (uint32, error) {
    return s.Handle.getUint32Value(C.p_uint32_f(C.VMGuestLib_GetMemMappedMB))
}

// GetMemActive retrieves the estimated amount of memory the virtual machine is actively using.
func (s *Session) GetMemActive() (uint32, error) {
    return s.Handle.getUint32Value(C.p_uint32_f(C.VMGuestLib_GetMemActiveMB))
}

// GetMemOverhead retrieves the amount of overhead memory associated with this virtual machine consumed on the host system.
func (s *Session) GetMemOverhead() (uint32, error) {
    return s.Handle.getUint32Value(C.p_uint32_f(C.VMGuestLib_GetMemOverheadMB))
}

// GetMemBallooned retrieves the amount of memory that has been reclaimed from this virtual machine via the VMware Memory Balloon mechanism.
func (s *Session) GetMemBallooned() (uint32, error) {
    return s.Handle.getUint32Value(C.p_uint32_f(C.VMGuestLib_GetMemBalloonedMB))
}

// GetMemSwapped retrieves the amount of memory associated with this virtual machine that has been swapped by the host system.
func (s *Session) GetMemSwapped() (uint32, error) {
    return s.Handle.getUint32Value(C.p_uint32_f(C.VMGuestLib_GetMemSwappedMB))
}

// GetMemShared retrieves the amount of physical memory associated with this virtual machine that is copy-on-write (COW) shared on the host.
func (s *Session) GetMemShared() (uint32, error) {
    return s.Handle.getUint32Value(C.p_uint32_f(C.VMGuestLib_GetMemSharedMB))
}

// GetMemSharedSaved retrieves the estimated amount of physical memory on the host saved from copy-on-write (COW) shared guest physical memory.
func (s *Session) GetMemSharedSaved() (uint32, error) {
    return s.Handle.getUint32Value(C.p_uint32_f(C.VMGuestLib_GetMemSharedSavedMB))
}

// GetMemUsed retrieves the estimated amount of physical host memory currently consumed for this virtual machine's physical memory.
func (s *Session) GetMemUsed() (uint32, error) {
    return s.Handle.getUint32Value(C.p_uint32_f(C.VMGuestLib_GetMemUsedMB))
}

// GetHostNumCPUCores retrieves the number of physical CPU cores on the host machine.
func (s *Session) GetHostNumCPUCores() (uint32, error) {
    return s.Handle.getUint32Value(C.p_uint32_f(C.VMGuestLib_GetHostNumCpuCores))
}
