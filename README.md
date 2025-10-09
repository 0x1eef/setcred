## About

The goal of this repository is to provide idiomatic Go interfaces
for APIs that are native to the BSD family of operating systems.
The [syscall](https://pkg.go.dev/syscall) module provides the
glue we need to perform system calls in pure Go, and without
writing C code.

## Examples

#### setcred

[setcred(2)](https://man.freebsd.org/cgi/man.cgi?setcred) is a
FreeBSD system call that can change the credentials of the current
process by altering the effective and/or real user and group IDs.
The system call either requires superuser privileges, or a mac policy
similar to [mac_do(4)](https://man.freebsd.org/cgi/man.cgi?mac_do)
that enables unprivileged users to execute the [setcred(2)](https://man.freebsd.org/cgi/man.cgi?setcred)
system call successfully.

The following example assumes it is being run with superuser privileges,
and when successful it changes the effective user and group IDs from 0
(root/wheel) to 1001 (an unprivileged user and group):

```go
package main

import (
	"fmt"
	"os"
	"github.com/0x1eef/bsd/setcred"
)

func main() {
	creds, flags := setcred.New(
		setcred.SetUid(uint32(1001)),
		setcred.SetGid(uint32(1001)),
	)
	if err := setcred.SetCred(creds, flags); err != nil {
		panic(err)
	} else {
		fmt.Printf("euid: %d, egid: %d", os.Geteuid(), os.Getegid())
	}
}
```

## Install

The install process is more or less straight forward

    go get github.com/0x1eef/bsd

## License

[BSD Zero Clause](https://choosealicense.com/licenses/0bsd/)
<br>
See [LICENSE](./LICENSE)