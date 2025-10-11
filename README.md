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
	err := setcred.SetCred(
		setcred.SetUid(uint32(1001)),
		setcred.SetGid(uint32(1001)),
	)
	if err != nil {
		panic(err)
	} else {
		fmt.Printf("euid: %d, egid: %d", os.Geteuid(), os.Getegid())
	}
}
```

### control

The control package can enable or disable security features
that are managed by the [HardenedBSD](https://hardenedbsd.org)
kernel on a per-file basis. Unlike other packages this one
happens to not be pure Go and requires C code to be compiled.
That's largely because HardenedBSD does not implement its
own system calls because they could conflict with FreeBSD,
and HardenedBSD regularly synchronizes updates from FreeBSD.

Given that context, HardenedBSD does not provide system calls
that can enable or disable feature state, and that leaves the
primary interface as the C libraries that HardenedBSD does
provide. In this case, that interface is
[libhbsdcontrol](https://git.hardenedbsd.org/hardenedbsd/hardenebsd).

The following example queries a list of feature names, and then proceeds
to enable, disable and restore the system default for the "mprotect"
feature. These changes are scoped to the `/usr/bin/mdo` binary:

```go
package main

import (
	"github.com/0x1eef/bsd/control"
)

func main() {
	ctx := control.New(control.Namespace("system"))
	if features, err := ctx.FeatureNames(); err != nil {
		panic(err)
	} else {
		for _, name := range features {
			fmt.Printf("feature: %s\n", name)
		}
		if err := ctx.Enable("mprotect", "/usr/bin/mdo"); err != nil {
			panic(err)
		}
		if err := ctx.Disable("mprotect", "/usr/bin/mdo"); err != nil {
			panic(err)
		}
		if err := ctx.Sysdef("mprotect", "/usr/bin/mdo"); err != nil {
			panic(err)
		}
		if status, err := ctx.Status("mprotect", "/usr/bin/mdo"); err != nil {
			fmt.Printf("The mproect has the status: %s\n", status)
		}
	}
}
```

## Install

The install process is more or less straight forward

    go get github.com/0x1eef/bsd

## Sources

* [github.com/@0x1eef](https://github.com/0x1eef/bsd#readme)
* [gitlab.com/@0x1eef](https://gitlab.com/0x1eef/bsd#about)
* [hardenedbsd.org/@0x1eef](https://git.HardenedBSD.org/0x1eef#about)

## License

[BSD Zero Clause](https://choosealicense.com/licenses/0bsd/)
<br>
See [LICENSE](./LICENSE)