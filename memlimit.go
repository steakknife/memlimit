package memlimit

// #cgo CFLAGS: -DHAVE_SYS_PARAM_H -DHAVE_SYSCTL_HW_USERMEM
// #cgo linux CFLAGS: -DHAVE_SYS_PARAM_H -DHAVE_SYSCTL_HW_USERMEM -DHAVE_SYS_SYSINFO_H
// #include "memlimit.h"
import "C"
import "errors"

// int memtouse(size_t, double, size_t *);
func MemToUse(maxmem uint64, maxmemfrac float64) (memlimit uint64, err error) {
    var temp C.size_t
    if C.memtouse(C.size_t(maxmem), C.double(maxmemfrac), &temp) != 0 {
        err = errors.New("memtouse error")
        return
    }
    memlimit = uint64(temp)
    return
}
