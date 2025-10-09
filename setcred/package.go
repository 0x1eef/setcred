package setcred

import (
	"errors"
	"fmt"
	"syscall"
	"unsafe"
)

type Option func(*setcred, *uint)

func SetCred(creds *setcred, flags uint) error {
	flagsp := uintptr(flags)
	credsp := uintptr(unsafe.Pointer(creds))
	sizep := uintptr(unsafe.Sizeof(*creds))
	_, _, err := syscall.Syscall6(uintptr(SYSCALL), flagsp, credsp, sizep, 0, 0, 0)
	if err != 0 {
		return errors.New(fmt.Sprintf("errno: %d", err))
	}
	return nil
}

func SetUid(uid uint32) Option {
	return func(creds *setcred, flags *uint) {
		*flags |= SETUID
		creds.sc_uid = uid
	}
}

func SetRuid(ruid uint32) Option {
	return func(creds *setcred, flags *uint) {
		*flags |= SETRUID
		creds.sc_ruid = ruid
	}
}

func SetSvUid(svuid uint32) Option {
	return func(creds *setcred, flags *uint) {
		*flags |= SETSVUID
		creds.sc_svuid = svuid
	}
}

func SetGid(gid uint32) Option {
	return func(creds *setcred, flags *uint) {
		*flags |= SETGID
		creds.sc_gid = gid
	}
}

func SetRgid(rgid uint32) Option {
	return func(creds *setcred, flags *uint) {
		*flags |= SETRGID
		creds.sc_rgid = rgid
	}
}

func SetSvGid(svgid uint32) Option {
	return func(creds *setcred, flags *uint) {
		*flags |= SETSVGID
		creds.sc_svgid = svgid
	}
}