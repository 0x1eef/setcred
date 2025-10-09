package setcred

import (
	"errors"
	"fmt"
	"syscall"
	"unsafe"
)

type Option func(*setcred, *uint)

const SYSCALL = int(591)
const SETUID = uint(1) << 0
const SETRUID = uint(1) << 1
const SETSVUID = uint(1) << 2
const SETGID = uint(1) << 3
const SETRGID = uint(1) << 4
const SETSVGID = uint(1) << 5
const SETSUPPGROUPS = uint(1) << 6
const SETMACLABEL = uint(1) << 7

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

func SetSuppGroups(groups ...uint32) Option {
	return func(creds *setcred, flags *uint) {
		*flags |= SETSUPPGROUPS
		creds.sc_supp_groups = uintptr(unsafe.Pointer(unsafe.SliceData(groups)))
		creds.sc_supp_groups_nb = uint32(len(groups))
	}
}
