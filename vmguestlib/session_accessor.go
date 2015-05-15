package vmguestlib

/*
#cgo CFLAGS: -I../vendor -I../native
#cgo LDFLAGS: -L/usr/lib/vmware-tools/lib/libvmtools.so -L/usr/lib/vmware-tools/lib/libvmGuestLib.so -lvmtools -lvmGuestLib
#include <vmGuestLib.h>
#include <vmSessionId.h>
#include <vmGuestLibProxy.h>
*/
import "C"
import "time"

// GetCPULimit retrieves the maximum processing power in MHz
// available to the virtual machine.
func (s *Session) GetCPULimit() (uint32, error) {
	return s.Handle.getUint32Value(C.p_uint32_f(C.VMGuestLib_GetCpuLimitMHz))
}

// GetCPUReservation retrieves the minimum processing power in
// MHz available to the virtual machine.
func (s *Session) GetCPUReservation() (uint32, error) {
	return s.Handle.getUint32Value(C.p_uint32_f(C.VMGuestLib_GetCpuReservationMHz))
}

// GetCPUShares retrieves the number of CPU shares allocated to the
// virtual machine.
func (s *Session) GetCPUShares() (uint32, error) {
	return s.Handle.getUint32Value(C.p_uint32_f(C.VMGuestLib_GetCpuShares))
}

// GetElaspedTime retrieves the duration since the virtual machine
// started running in the current host system.
func (s *Session) GetElaspedTime() (t time.Duration, err error) {
	nativeVal := new(C.uint64)
	e := C.VMGuestLib_GetElapsedMs(*s.Handle.NativeHandle, nativeVal)
	if e != ErrorSuccess {
		err = newError(e)
	}
	t = time.Duration(*nativeVal) * time.Millisecond
	return
}

// GetHostProcessorSpeed retrieves the host processor speed.
func (s *Session) GetHostProcessorSpeed() (uint32, error) {
	return s.Handle.getUint32Value(C.p_uint32_f(C.VMGuestLib_GetHostProcessorSpeed))
}

// GetMemReservation retrieves the minimum amount of memory
// that is available to the virtual machine.
func (s *Session) GetMemReservation() (uint32, error) {
	return s.Handle.getUint32Value(C.p_uint32_f(C.VMGuestLib_GetMemReservationMB))
}

// GetMemLimit retrieves the maximum amount of memory that
// is available to the virtual machine.
func (s *Session) GetMemLimit() (uint32, error) {
	return s.Handle.getUint32Value(C.p_uint32_f(C.VMGuestLib_GetMemLimitMB))
}

// GetMemShares retrieves the number of memory shares allocated
// to the virtual machine.
func (s *Session) GetMemShares() (uint32, error) {
	return s.Handle.getUint32Value(C.p_uint32_f(C.VMGuestLib_GetMemShares))
}
