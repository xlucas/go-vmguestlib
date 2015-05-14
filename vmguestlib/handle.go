package vmguestlib

/*
#cgo CFLAGS: -I../native
#cgo LDFLAGS: -L/usr/lib/vmware-tools/lib/libvmtools.so -L/usr/lib/vmware-tools/lib/libvmGuestLib.so -lvmtools -lvmGuestLib
#include <vmGuestLib.h>
*/
import "C"

//
// A VMGuestLib handle.
//
// This handle provides a context for accessing all GuestLib
// state. Use OpenHandle() to get a handle for use with other
// GuestLib functions, and use CloseHandle() to release a handle
// previously acquired with OpenHandle().
//
// All of the statistics and session state are maintained per GuestLib
// handle, so operating on one GuestLib handle will not affect the
// state of another handle.
//
type Handle struct {
	NativeHandle *C.VMGuestLibHandle
}

// NewHandle opens a new handle.
func NewHandle() (h *Handle, err error) {
	h = &Handle{
		NativeHandle: new(C.VMGuestLibHandle),
	}
	e := C.VMGuestLib_OpenHandle(h.NativeHandle)
	if e != ERROR_SUCCESS {
		err = newError(e)
	}
	return h, err
}

// Close releases a previously opened handle.
func (h *Handle) Close() (err error) {
	e := C.VMGuestLib_CloseHandle(*h.NativeHandle)
	if e != ERROR_SUCCESS {
		err = newError(e)
	}
	return
}
