package vmguestlib

/*
#cgo CFLAGS: -I../vendor -I../native
#cgo LDFLAGS: -L/usr/lib/vmware-tools/lib/libvmtools.so -L/usr/lib/vmware-tools/lib/libvmGuestLib.so -lvmtools -lvmGuestLib
#include <vmGuestLib.h>
#include <vmSessionId.h>
#include <vmGuestLibProxy.h>
*/
import "C"

// Session gives access to accessor functions through the usage of a
// particular Handle.
// Session is also used to detect changes in the "session" of a virtual
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
	// A handle object.
	Handle *Handle
	// The native VMSessionId object.
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
	if e != ErrorSuccess {
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
