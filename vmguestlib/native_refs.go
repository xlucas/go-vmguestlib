package vmguestlib

/*
#cgo CFLAGS: -I../vendor -I../native
#cgo LDFLAGS: -L/usr/lib/vmware-tools/lib/libvmtools.so -L/usr/lib/vmware-tools/lib/libvmGuestLib.so -lvmtools -lvmGuestLib
#include <vmGuestLib.h>
#include <vmGuestLibProxy.h>
*/
import "C"

var nativeGetCPULimitMhz = C.p_uint32_f(C.VMGuestLib_GetCpuLimitMHz)
var nativeGetCPUReservationMhz = C.p_uint32_f(C.VMGuestLib_GetCpuReservationMHz)
var nativeGetCPUShares = C.p_uint32_f(C.VMGuestLib_GetCpuShares)
