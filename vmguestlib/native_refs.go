package vmguestlib

/*
#cgo CFLAGS: -I../vendor -I../native
#cgo LDFLAGS: -L/usr/lib/vmware-tools/lib/libvmtools.so -L/usr/lib/vmware-tools/lib/libvmGuestLib.so -lvmtools -lvmGuestLib
#include <vmGuestLib.h>
#include <vmGuestLibProxy.h>
*/
import "C"

var nativeGetCPULimit = C.p_uint32_f(C.VMGuestLib_GetCpuLimitMHz)
var nativeGetCPUReservation = C.p_uint32_f(C.VMGuestLib_GetCpuReservationMHz)
var nativeGetCPUShares = C.p_uint32_f(C.VMGuestLib_GetCpuShares)
var nativeGetHostProcessorSpeed = C.p_uint32_f(C.VMGuestLib_GetHostProcessorSpeed)
var nativeGetMemReservation = C.p_uint32_f(C.VMGuestLib_GetMemReservationMB)
var nativeGetMemLimit = C.p_uint32_f(C.VMGuestLib_GetMemLimitMB)
var nativeGetMemShares = C.p_uint32_f(C.VMGuestLib_GetMemShares)
var nativeGetMemMapped = C.p_uint32_f(C.VMGuestLib_GetMemMappedMB)
var nativeGetMemActive = C.p_uint32_f(C.VMGuestLib_GetMemActiveMB)
var nativeGetMemOverhead = C.p_uint32_f(C.VMGuestLib_GetMemOverheadMB)
var nativeGetMemBallooned = C.p_uint32_f(C.VMGuestLib_GetMemBalloonedMB)
var nativeGetMemSwapped = C.p_uint32_f(C.VMGuestLib_GetMemSwappedMB)
var nativeGetMemShared = C.p_uint32_f(C.VMGuestLib_GetMemSharedMB)
var nativeGetMemSharedSaved = C.p_uint32_f(C.VMGuestLib_GetMemSharedSavedMB)
var nativeGetMemUsed = C.p_uint32_f(C.VMGuestLib_GetMemUsedMB)
var nativeGetHostNumCores = C.p_uint32_f(C.VMGuestLib_GetHostNumCpuCores)
