package memlimit

import (
    "testing"
)

func TestMemLimit(t *testing.T) {
    r, _ := MemToUse(1, 0.3)
    mib := uint64(1024*1024)
    if r != mib {
        t.Error("min failed")
    }

    twomib := 2*mib
    r, _ = MemToUse(twomib, 0.3)
    if r != twomib {
        t.Error("2mib failed")
    }

    tib := mib*mib
    r, _ = MemToUse(tib, 0.3)
    if r == tib {
        t.Error("tib failed")
    }
}
