## About

The setcred module provides a Go interface for the
[setcred(2)](https://man.freebsd.org/setcred) system call &ndash;
a FreeBSD-specific system call for changing process
credentials at runtime.

## Examples

#### setcred

The system call either requires superuser privileges, or a mac policy similar
to [mac_do(4)](https://man.freebsd.org/cgi/man.cgi?mac_do) that enables unprivileged
users to execute the [setcred(2)](https://man.freebsd.org/cgi/man.cgi?setcred)
system call successfully. The following example assumes it is being run with
superuser privileges, and when successful it changes the effective user ID and
effective group ID from 0 (root/wheel) to 1001 (an unprivileged user and group):

```go
package main

import (
	"fmt"
	"os"

	"github.com/0x1eef/setcred"
)

func main() {
	err := setcred.SetCred(
		setcred.SetUid(1001),
		setcred.SetGid(1001),
	)
	if err != nil {
		panic(err)
	} else {
		fmt.Printf("euid: %d, egid: %d\n", os.Geteuid(), os.Getegid())
	}
}
```

## Install

The install process is more or less straight forward

    go get github.com/0x1eef/setcred@v0.1.0

## Sources

* [github.com/@0x1eef](https://github.com/0x1eef/setcred#readme)
* [gitlab.com/@0x1eef](https://gitlab.com/0x1eef/setcred#about)
* [hardenedbsd.org/@0x1eef](https://git.HardenedBSD.org/0x1eef/setcred#about)

## License

[BSD Zero Clause](https://choosealicense.com/licenses/0bsd/)
<br>
See [LICENSE](./LICENSE)