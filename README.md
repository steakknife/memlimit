# memlimit

Extracted from Colin Percival's awesome scrypt


## Usage

    go get github.com/steakknife/memlimit


```go
import "github.com/steakknife/memlimit"

// request 512 MiB or 30% of memory, whichever is smaller
mem, err := memlimit.MemToUse(512*1024*1024, 0.3)
if err != nil {
    // memtouse had a problem
}
// mem is in bytes
```

## Acknowledgements

Colin Percival's awesome scrypt and tarsnap
