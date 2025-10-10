package setcred

import (
	"syscall"
	"unsafe"
	"runtime"
)

type Option func(*setcred, *uint)

const sysid = int(591)
const setuid = uint(1) << 0
const setruid = uint(1) << 1
const setsvuid = uint(1) << 2
const setgid = uint(1) << 3
const setrgid = uint(1) << 4
const setsvgid = uint(1) << 5
const setsuppgroups = uint(1) << 6
const setmaclabel = uint(1) << 7

func SetCred(opts ...Option) error {
	creds, flags := new(opts...)
	iptr := uintptr(sysid)
	fptr := uintptr(flags)
	cptr := uintptr(unsafe.Pointer(creds))
	sptr := uintptr(unsafe.Sizeof(*creds))
	_, _, err := syscall.Syscall6(iptr, fptr, cptr, sptr, 0, 0, 0)
	if err != 0 {
		return err
	}
	runtime.KeepAlive(creds)
	for _, set := range opts {
		runtime.KeepAlive(set)
	}
	return nil
}

func SetUid(uid uint32) Option {
	return func(creds *setcred, flags *uint) {
		*flags |= setuid
		creds.sc_uid = uid
	}
}

func SetRuid(ruid uint32) Option {
	return func(creds *setcred, flags *uint) {
		*flags |= setruid
		creds.sc_ruid = ruid
	}
}

func SetSvUid(svuid uint32) Option {
	return func(creds *setcred, flags *uint) {
		*flags |= setsvuid
		creds.sc_svuid = svuid
	}
}

func SetGid(gid uint32) Option {
	return func(creds *setcred, flags *uint) {
		*flags |= setgid
		creds.sc_gid = gid
	}
}

func SetRgid(rgid uint32) Option {
	return func(creds *setcred, flags *uint) {
		*flags |= setrgid
		creds.sc_rgid = rgid
	}
}

func SetSvGid(svgid uint32) Option {
	return func(creds *setcred, flags *uint) {
		*flags |= setsvgid
		creds.sc_svgid = svgid
	}
}

func SetSuppGroups(groups ...uint32) Option {
	return func(creds *setcred, flags *uint) {
		*flags |= setsuppgroups
		creds.sc_supp_groups = uintptr(unsafe.Pointer(unsafe.SliceData(groups)))
		creds.sc_supp_groups_nb = uint32(len(groups))
	}
}
