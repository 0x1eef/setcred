## About

The setcred module provides a Go interface for the
[setcred(2)](https://man.freebsd.org/setcred) system call,
a FreeBSD-specific mechanism for managing process credentials. The system call
allows programs to dynamically change their effective user ID (EUID),
effective group ID (EGID), and other credentials at runtime.

A unique and innovative aspect to setcred(2) is its ability to be extended by
MAC policies such as [mac_do(4)](https://man.freebsd.org/cgi/man.cgi?mac_do),
and in turn this can allow unprivileged users change their credentials in a way
similar to root although any credential transitions must first be added to an
allowlist.

## Examples

#### setcred

The system call either requires superuser privileges, or &ndash; as mentioned
earlier &ndash; a mac policy similar to [mac_do(4)](https://man.freebsd.org/cgi/man.cgi?mac_do)
that enables unprivileged users to execute the [setcred(2)](https://man.freebsd.org/cgi/man.cgi?setcred)
system call successfully.

The following example assumes it is being run with superuser privileges,
and when successful it changes the effective user ID and effective group
ID from 0 (root/wheel) to 1001 (an unprivileged user and group). It is
worth keeping in mind that [setcred(2)](https://man.freebsd.org/cgi/man.cgi?setcred)
is capable of changing other credentials as well:

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