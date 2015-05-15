package vmguestlib

/*
#cgo CFLAGS: -I../vendor -I../native
#cgo LDFLAGS: -L/usr/lib/vmware-tools/lib/libvmtools.so -L/usr/lib/vmware-tools/lib/libvmGuestLib.so -lvmtools -lvmGuestLib
#include <vmGuestLib.h>
#include <vmGuestLibProxy.h>
*/
import "C"

// Handle provides a context for accessing all GuestLib
// state. Use OpenHandle() to get a handle for use with other
// GuestLib functions, and use CloseHandle() to release a handle
// previously acquired with OpenHandle().
//
// All of the statistics and session state are maintained per GuestLib
// handle, so operating on one GuestLib handle will not affect the
// state of another handle.
//
type Handle struct {
	// The native VMGuestLibHandle object.
	NativeHandle *C.VMGuestLibHandle
}

// NewHandle opens a new handle.
func NewHandle() (h *Handle, err error) {
	h = &Handle{
		NativeHandle: new(C.VMGuestLibHandle),
	}
	e := C.VMGuestLib_OpenHandle(h.NativeHandle)
	if e != ErrorSuccess {
		err = newError(e)
	}
	return h, err
}

// Close releases a previously opened handle.
func (h *Handle) Close() (err error) {
	e := C.VMGuestLib_CloseHandle(*h.NativeHandle)
	if e != ErrorSuccess {
		err = newError(e)
	}
	return
}

// UpdateInfo updates the session state for the current handle.
// No locking is done internally, so multithreading environments
// should use separate handles. This is a fairly heavyweight
// function on the native API side so environments concerned
// about performance should minimize the number of calls to
// this function.
func (h *Handle) UpdateInfo() (err error) {
	e := C.VMGuestLib_UpdateInfo(*h.NativeHandle)
	if e != ErrorSuccess {
		err = newError(e)
	}
	return
}

// getUint32Value retrieves one uint32 native accessor value
// by calling the native accessor routine wrapped into a
// native proxy function pointer type and converting the C
// uint32 value to golang's relevant type.
func (h *Handle) getUint32Value(p C.p_uint32_f) (v uint32, err error) {
	nativeUint32 := new(C.uint32)
	nativeError := C.proxy_uint32_f(p, *h.NativeHandle, nativeUint32)
	if nativeError != ErrorSuccess {
		err = newError(nativeError)
		return
	}
	v = uint32(*nativeUint32)
	return
}
