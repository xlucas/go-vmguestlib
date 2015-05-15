package vmguestlib

/*
#cgo CFLAGS: -I../native
#cgo LDFLAGS: -L/usr/lib/vmware-tools/lib/libvmtools.so -L/usr/lib/vmware-tools/lib/libvmGuestLib.so -lvmtools -lvmGuestLib
#include <vmGuestLib.h>
#include <vmSessionId.h>
*/
import "C"
import "time"

// A VMGuestLib session.
//
// This is used to detect changes in the "session" of a virtual
// machine. "Session" in this context refers to the particular running
// instance of this virtual machine on a given host. Moving a virtual
// machine to another host using VMotion will cause a session change
// as well as suspending and resuming a virtual machine or reverting
// to a snapshot.
//
// Events like VMotion, suspend/resume, snapshot revert are likely
// to render invalid any information previously retrieved through
// this API, so the intention of the session is to provide applications
// with a mechanism to detect those events and react accordingly,
// e.g. by refreshing and resetting any state that relies on validity
// of previously retrieved information.
//
type Session struct {
	Handle        *Handle
	NativeSession *C.VMSessionId
}

// Open fetch a new session from the given handle.
func Open(h *Handle) (s *Session, err error) {
	s = &Session{
		Handle:        h,
		NativeSession: new(C.VMSessionId),
	}
	_, err = s.RefreshInfo()
	return
}

// Close tears down the underlying handle.
func (s *Session) Close() error {
	return s.Handle.Close()
}

// NewSession allocates a new handle and calls Open().
func NewSession() (s *Session, err error) {
	h, err := NewHandle()
	if err != nil {
		return
	}
	s, err = Open(h)
	return
}

// Refresh updates the session if the virtual machine session
// changed on the host. In this case, changed is set to true.
func (s *Session) Refresh() (changed bool, err error) {
	newSession := new(C.VMSessionId)
	e := C.VMGuestLib_GetSessionId(*s.Handle.NativeHandle, newSession)
	if e != ERROR_SUCCESS {
		err = newError(e)
		return
	}
	changed = (*newSession != *s.NativeSession)
	s.NativeSession = newSession
	return
}

// RefreshInfo calls UpdateInfo() on the underlying handle
// then calls Refresh() on this session.
func (s *Session) RefreshInfo() (changed bool, err error) {
	err = s.Handle.UpdateInfo()
	if err != nil {
		return
	}
	changed, err = s.Refresh()
	return
}

// GetCpuLimitMhz retrieves the maximum processing power in MHz
// available to the virtual machine.
func (s *Session) GetCpuLimitMhz() (limit uint32, err error) {
	cLimit := new(C.uint32)
	e := C.VMGuestLib_GetCpuLimitMHz(*s.Handle.NativeHandle, cLimit)
	if e != ERROR_SUCCESS {
		err = newError(e)
	}
	limit = uint32(*cLimit)
	return
}

// GetCpuReservationMhz retrieves the minimum processing power in
// MHz available to the virtual machine.
func (s *Session) GetCpuReservationMhz() (reserved uint32, err error) {
	cReserved := new(C.uint32)
	e := C.VMGuestLib_GetCpuReservationMHz(*s.Handle.NativeHandle, cReserved)
	if e != ERROR_SUCCESS {
		err = newError(e)
	}
	reserved = uint32(*cReserved)
	return
}

// GetElaspedTime retrieves the duration since the virtual machine
// started running in the current host system.
func (s *Session) GetElaspedTime() (t time.Duration, err error) {
	elapsedMs := new(C.uint64)
	e := C.VMGuestLib_GetElapsedMs(*s.Handle.NativeHandle, elapsedMs)
	if e != ERROR_SUCCESS {
		err = newError(e)
	}
	t = time.Duration(*elapsedMs) * time.Millisecond
	return
}
