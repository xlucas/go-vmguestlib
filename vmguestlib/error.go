package vmguestlib

/*
#cgo CFLAGS: -I../native
#cgo LDFLAGS: -L/usr/lib/vmware-tools/lib/libvmtools.so -L/usr/lib/vmware-tools/lib/libvmGuestLib.so -lvmtools -lvmGuestLib
#include <vmGuestLib.h>
*/
import "C"

// Error codes returned by the native VMGuestLib API
const (
	ERROR_SUCCESS             = iota // No error.
	ERROR_OTHER                      // Other error.
	ERROR_NOT_RUNNING_IN_VM          // Not running in a VM.
	ERROR_NOT_ENABLED                // GuestLib not enabled on the host.
	ERROR_NOT_AVAILABLE              // This stat is not available on the host.
	ERROR_NO_INFO                    // UpdateInfo() has never been called.
	ERROR_MEMORY                     // Not enough memory.
	ERROR_BUFFER_TOO_SMALL           // Buffer too small.
	ERROR_INVALID_HANDLE             // Handle is invalid.
	ERROR_INVALID_ARG                // One or more arguments were invalid.
	ERROR_UNSUPPORTED_VERSION        // The host doesn't support this request.
)

// Error messages associated to the native VMGuestLib API error codes
const (
	ERROR_MSG_SUCCESS             = "The function has completed successfully."
	ERROR_MSG_OTHER               = "An error has occurred. No additional information about the type of the error is available."
	ERROR_MSG_NOT_RUNNING_IN_VM   = "The program making this call is not running on a VMware virtualmachine."
	ERROR_MSG_NOT_ENABLED         = "The vSphere Guest API is not enabled on this host so these functions cannot be used."
	ERROR_MSG_NOT_AVAILABLE       = "The information requested is not available on this host."
	ERROR_MSG_NO_INFO             = "The handle data structure does not contain any information."
	ERROR_MSG_MEMORY              = "There is not enough memory available to complete the call."
	ERROR_MSG_BUFFER_TOO_SMALL    = "The buffer is too small to accommodate the function call."
	ERROR_MSG_INVALID_HANDLE      = "The handle that you used is invalid."
	ERROR_MSG_INVALID_ARG         = "One or more of the arguments passed to the function were invalid."
	ERROR_MSG_UNSUPPORTED_VERSION = "The host does not support the requested statistic."
)

var errors = map[uint]string{
	ERROR_SUCCESS:             ERROR_MSG_SUCCESS,
	ERROR_OTHER:               ERROR_MSG_OTHER,
	ERROR_NOT_RUNNING_IN_VM:   ERROR_MSG_NOT_RUNNING_IN_VM,
	ERROR_NOT_ENABLED:         ERROR_MSG_NOT_ENABLED,
	ERROR_NOT_AVAILABLE:       ERROR_MSG_NOT_AVAILABLE,
	ERROR_NO_INFO:             ERROR_MSG_NO_INFO,
	ERROR_MEMORY:              ERROR_MSG_MEMORY,
	ERROR_BUFFER_TOO_SMALL:    ERROR_MSG_BUFFER_TOO_SMALL,
	ERROR_INVALID_HANDLE:      ERROR_MSG_INVALID_HANDLE,
	ERROR_INVALID_ARG:         ERROR_MSG_INVALID_ARG,
	ERROR_UNSUPPORTED_VERSION: ERROR_MSG_UNSUPPORTED_VERSION,
}

type Error struct {
	NativeError C.VMGuestLibError
	Message     string
}

func newError(e C.VMGuestLibError) *Error {
	return &Error{
		NativeError: e,
		Message:     errors[uint(e)],
	}
}

func (e *Error) Error() string {
	return e.Message
}
