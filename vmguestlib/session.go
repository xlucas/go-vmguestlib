package vmguestlib

/*
#cgo CFLAGS: -I../native
#cgo LDFLAGS: -L/usr/lib/vmware-tools/lib/libvmtools.so -L/usr/lib/vmware-tools/lib/libvmGuestLib.so -lvmtools -lvmGuestLib
#include <vmGuestLib.h>
#include <vmSessionId.h>
*/
import "C"

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
	NativeSession *C.VMSessionId
}

// NewSession opens a new session.
func NewSession(h *Handle) (s *Session, err error) {
	s = &Session{
		NativeSession: new(C.VMSessionId),
	}
	err = h.UpdateInfo()
	if err != nil {
		return
	}
	_, err = s.Refresh(h)
	return
}

// Refresh updates the session if the virtual machine session
// changed on the host. In this case, hasChanged is set to true.
func (s *Session) Refresh(h *Handle) (hasChanged bool, err error) {
	newSession := new(C.VMSessionId)
	e := C.VMGuestLib_GetSessionId(*h.NativeHandle, newSession)
	if e != ERROR_SUCCESS {
		err = NewError(e)
		return
	}
	hasChanged = (*newSession != *s.NativeSession)
	s.NativeSession = newSession
	return
}
