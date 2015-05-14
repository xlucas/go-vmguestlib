package vmguestlib

/*
#cgo CFLAGS: -I../native
#cgo LDFLAGS: -L/usr/lib/vmware-tools/lib/libvmtools.so -L/usr/lib/vmware-tools/lib/libvmGuestLib.so -lvmtools -lvmGuestLib
#include <vmGuestLib.h>
*/
import "C"

// Error codes returned by the native VMGuestLib API
const (
	_ERROR_SUCCESS             = iota // No error.
	_ERROR_OTHER                      // Other error.
	_ERROR_NOT_RUNNING_IN_VM          // Not running in a VM.
	_ERROR_NOT_ENABLED                // GuestLib not enabled on the host.
	_ERROR_NOT_AVAILABLE              // This stat is not available on the host.
	_ERROR_NO_INFO                    // UpdateInfo() has never been called.
	_ERROR_MEMORY                     // Not enough memory.
	_ERROR_BUFFER_TOO_SMALL           // Buffer too small.
	_ERROR_INVALID_HANDLE             // Handle is invalid.
	_ERROR_INVALID_ARG                // One or more arguments were invalid.
	_ERROR_UNSUPPORTED_VERSION        // The host doesn't support this request.
)

// Error messages associated to the native VMGuestLib API error codes
const (
	_ERROR_MSG_SUCCESS             = "The function has completed successfully."
	_ERROR_MSG_OTHER               = "An error has occurred. No additional information about the type of the error is available."
	_ERROR_MSG_NOT_RUNNING_IN_VM   = "The program making this call is not running on a VMware virtualmachine."
	_ERROR_MSG_NOT_ENABLED         = "The vSphere Guest API is not enabled on this host so these functions cannot be used."
	_ERROR_MSG_NOT_AVAILABLE       = "The information requested is not available on this host."
	_ERROR_MSG_NO_INFO             = "The handle data structure does not contain any information."
	_ERROR_MSG_MEMORY              = "There is not enough memory available to complete the call."
	_ERROR_MSG_BUFFER_TOO_SMALL    = "The buffer is too small to accommodate the function call."
	_ERROR_MSG_INVALID_HANDLE      = "The handle that you used is invalid."
	_ERROR_MSG_INVALID_ARG         = "One or more of the arguments passed to the function were invalid."
	_ERROR_MSG_UNSUPPORTED_VERSION = "The host does not support the requested statistic."
)

var errors = map[uint]string{
	_ERROR_SUCCESS:             _ERROR_MSG_SUCCESS,
	_ERROR_OTHER:               _ERROR_MSG_OTHER,
	_ERROR_NOT_RUNNING_IN_VM:   _ERROR_MSG_NOT_RUNNING_IN_VM,
	_ERROR_NOT_ENABLED:         _ERROR_MSG_NOT_ENABLED,
	_ERROR_NOT_AVAILABLE:       _ERROR_MSG_NOT_AVAILABLE,
	_ERROR_NO_INFO:             _ERROR_MSG_NO_INFO,
	_ERROR_MEMORY:              _ERROR_MSG_MEMORY,
	_ERROR_BUFFER_TOO_SMALL:    _ERROR_MSG_BUFFER_TOO_SMALL,
	_ERROR_INVALID_HANDLE:      _ERROR_MSG_INVALID_HANDLE,
	_ERROR_INVALID_ARG:         _ERROR_MSG_INVALID_ARG,
	_ERROR_UNSUPPORTED_VERSION: _ERROR_MSG_UNSUPPORTED_VERSION,
}

type Error struct {
	code    uint
	message string
}

func newError(e C.VMGuestLibError) *Error {
	return &Error{
		code:    uint(e),
		message: errors[uint(e)],
	}
}

func (e *Error) Error() string {
	return e.message
}
