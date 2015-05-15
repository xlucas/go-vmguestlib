#ifndef _VM_GUEST_LIB_PROXY_H
#define _VM_GUEST_LIB_PROXY_H

#include "../vendor/vm_basic_types.h"
#include "../vendor/vmGuestLib.h"

#ifdef __cplusplus
extern "C" {
#endif

typedef VMGuestLibError (*p_uint32_f) ();

static inline VMGuestLibError proxy_uint32_f(p_uint32_f f, VMGuestLibHandle h, uint32* v)
{
    return f(h, v);
}

#ifdef __cplusplus
}
#endif

#endif /* _VM_GUEST_LIB_PROXY_H */
