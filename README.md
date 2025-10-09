## About

The goal of this repository is to provide idiomatic Go interfaces
for APIs that are native to the BSD family of operating systems.

## Examples

#### setcred

The [setcred(2)](https://man.freebsd.org/cgi/man.cgi?setcred) system
call can change the credentials of the current process by altering the
effective and/or real user and group IDs. The system call either requires
superuser privileges, or a mac policy similar to [mac_do(4)](https://man.freebsd.org/cgi/man.cgi?mac_do)
that enables unprivileged users to execute the [setcred(2)](https://man.freebsd.org/cgi/man.cgi?setcred)
system call successfully.

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